// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	versioned1 "github.com/knative/build-pipeline/pkg/client/clientset/versioned"
)

func AnyVersioned1Interface() versioned1.Interface {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(versioned1.Interface))(nil)).Elem()))
	var nullValue versioned1.Interface
	return nullValue
}

func EqVersioned1Interface(value versioned1.Interface) versioned1.Interface {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue versioned1.Interface
	return nullValue
}
