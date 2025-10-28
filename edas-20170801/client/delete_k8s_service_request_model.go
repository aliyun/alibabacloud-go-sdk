// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteK8sServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteK8sServiceRequest
	GetAppId() *string
	SetName(v string) *DeleteK8sServiceRequest
	GetName() *string
}

type DeleteK8sServiceRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5a166fbd***a286-781659d9f54c
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the service.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-http
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteK8sServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteK8sServiceRequest) GoString() string {
	return s.String()
}

func (s *DeleteK8sServiceRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteK8sServiceRequest) GetName() *string {
	return s.Name
}

func (s *DeleteK8sServiceRequest) SetAppId(v string) *DeleteK8sServiceRequest {
	s.AppId = &v
	return s
}

func (s *DeleteK8sServiceRequest) SetName(v string) *DeleteK8sServiceRequest {
	s.Name = &v
	return s
}

func (s *DeleteK8sServiceRequest) Validate() error {
	return dara.Validate(s)
}
