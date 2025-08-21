// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPredefinedScopesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *ListPredefinedScopesRequest
	GetAppType() *string
}

type ListPredefinedScopesRequest struct {
	// The type of the application. Valid values:
	//
	// 	- WebApp
	//
	// 	- NativeApp
	//
	// 	- ServerApp
	//
	// If this parameter is empty, the permissions on all types of applications are queried.
	//
	// example:
	//
	// WebApp
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
}

func (s ListPredefinedScopesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPredefinedScopesRequest) GoString() string {
	return s.String()
}

func (s *ListPredefinedScopesRequest) GetAppType() *string {
	return s.AppType
}

func (s *ListPredefinedScopesRequest) SetAppType(v string) *ListPredefinedScopesRequest {
	s.AppType = &v
	return s
}

func (s *ListPredefinedScopesRequest) Validate() error {
	return dara.Validate(s)
}
