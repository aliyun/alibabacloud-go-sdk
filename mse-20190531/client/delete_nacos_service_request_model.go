// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNacosServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DeleteNacosServiceRequest
	GetAcceptLanguage() *string
	SetGroupName(v string) *DeleteNacosServiceRequest
	GetGroupName() *string
	SetInstanceId(v string) *DeleteNacosServiceRequest
	GetInstanceId() *string
	SetNamespaceId(v string) *DeleteNacosServiceRequest
	GetNamespaceId() *string
	SetServiceName(v string) *DeleteNacosServiceRequest
	GetServiceName() *string
}

type DeleteNacosServiceRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The name of the group.
	//
	// This parameter is required.
	//
	// example:
	//
	// DEFAULT_GROUP
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// mse-cn-123456
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// 9e78a671-4b9b-4dd4-99c1-0****
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The name of the service.
	//
	// This parameter is required.
	//
	// example:
	//
	// hello_service
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s DeleteNacosServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNacosServiceRequest) GoString() string {
	return s.String()
}

func (s *DeleteNacosServiceRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DeleteNacosServiceRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *DeleteNacosServiceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteNacosServiceRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *DeleteNacosServiceRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *DeleteNacosServiceRequest) SetAcceptLanguage(v string) *DeleteNacosServiceRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DeleteNacosServiceRequest) SetGroupName(v string) *DeleteNacosServiceRequest {
	s.GroupName = &v
	return s
}

func (s *DeleteNacosServiceRequest) SetInstanceId(v string) *DeleteNacosServiceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteNacosServiceRequest) SetNamespaceId(v string) *DeleteNacosServiceRequest {
	s.NamespaceId = &v
	return s
}

func (s *DeleteNacosServiceRequest) SetServiceName(v string) *DeleteNacosServiceRequest {
	s.ServiceName = &v
	return s
}

func (s *DeleteNacosServiceRequest) Validate() error {
	return dara.Validate(s)
}
