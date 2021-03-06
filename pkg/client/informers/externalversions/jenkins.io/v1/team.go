// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	jenkins_io_v1 "github.com/jenkins-x/jx/pkg/apis/jenkins.io/v1"
	versioned "github.com/jenkins-x/jx/pkg/client/clientset/versioned"
	internalinterfaces "github.com/jenkins-x/jx/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/jenkins-x/jx/pkg/client/listers/jenkins.io/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// TeamInformer provides access to a shared informer and lister for
// Teams.
type TeamInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.TeamLister
}

type teamInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewTeamInformer constructs a new informer for Team type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTeamInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTeamInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredTeamInformer constructs a new informer for Team type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTeamInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.JenkinsV1().Teams(namespace).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.JenkinsV1().Teams(namespace).Watch(options)
			},
		},
		&jenkins_io_v1.Team{},
		resyncPeriod,
		indexers,
	)
}

func (f *teamInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTeamInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *teamInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&jenkins_io_v1.Team{}, f.defaultInformer)
}

func (f *teamInformer) Lister() v1.TeamLister {
	return v1.NewTeamLister(f.Informer().GetIndexer())
}
