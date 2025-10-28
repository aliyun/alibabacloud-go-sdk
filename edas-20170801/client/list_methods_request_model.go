// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMethodsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListMethodsRequest
	GetAppId() *string
	SetServiceName(v string) *ListMethodsRequest
	GetServiceName() *string
}

type ListMethodsRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// be213a4a-c7e4-473b-ab0****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the service.
	//
	// This parameter is required.
	//
	// example:
	//
	// Method
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s ListMethodsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMethodsRequest) GoString() string {
	return s.String()
}

func (s *ListMethodsRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListMethodsRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListMethodsRequest) SetAppId(v string) *ListMethodsRequest {
	s.AppId = &v
	return s
}

func (s *ListMethodsRequest) SetServiceName(v string) *ListMethodsRequest {
	s.ServiceName = &v
	return s
}

func (s *ListMethodsRequest) Validate() error {
	return dara.Validate(s)
}
