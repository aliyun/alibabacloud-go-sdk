// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNacosServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *CreateNacosServiceRequest
	GetAcceptLanguage() *string
	SetClusterId(v string) *CreateNacosServiceRequest
	GetClusterId() *string
	SetEphemeral(v bool) *CreateNacosServiceRequest
	GetEphemeral() *bool
	SetGroupName(v string) *CreateNacosServiceRequest
	GetGroupName() *string
	SetInstanceId(v string) *CreateNacosServiceRequest
	GetInstanceId() *string
	SetNamespaceId(v string) *CreateNacosServiceRequest
	GetNamespaceId() *string
	SetProtectThreshold(v string) *CreateNacosServiceRequest
	GetProtectThreshold() *string
	SetServiceName(v string) *CreateNacosServiceRequest
	GetServiceName() *string
}

type CreateNacosServiceRequest struct {
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
	// The ID of the cluster.
	//
	// > This operation contains both the InstanceId and ClusterId parameters. You must specify one of them.
	//
	// example:
	//
	// mse-3691a080
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether the instance is marked as a temporary node. Valid values:
	//
	// 	- `true`: yes
	//
	// 	- `false`: no
	//
	// example:
	//
	// true
	Ephemeral *bool `json:"Ephemeral,omitempty" xml:"Ephemeral,omitempty"`
	// The name of the group.
	//
	// example:
	//
	// DEFAULT_GROUP
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The ID of the instance.
	//
	// > This operation contains both the InstanceId and ClusterId parameters. You must specify one of them.
	//
	// example:
	//
	// mse-cn-st21ri2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// production
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The protection threshold.
	//
	// example:
	//
	// 0
	ProtectThreshold *string `json:"ProtectThreshold,omitempty" xml:"ProtectThreshold,omitempty"`
	// The name of the service.
	//
	// This parameter is required.
	//
	// example:
	//
	// com.alibabacloud.hipstershop.cartserviceapi.service.CartService
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s CreateNacosServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNacosServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateNacosServiceRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *CreateNacosServiceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateNacosServiceRequest) GetEphemeral() *bool {
	return s.Ephemeral
}

func (s *CreateNacosServiceRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateNacosServiceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateNacosServiceRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *CreateNacosServiceRequest) GetProtectThreshold() *string {
	return s.ProtectThreshold
}

func (s *CreateNacosServiceRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *CreateNacosServiceRequest) SetAcceptLanguage(v string) *CreateNacosServiceRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *CreateNacosServiceRequest) SetClusterId(v string) *CreateNacosServiceRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateNacosServiceRequest) SetEphemeral(v bool) *CreateNacosServiceRequest {
	s.Ephemeral = &v
	return s
}

func (s *CreateNacosServiceRequest) SetGroupName(v string) *CreateNacosServiceRequest {
	s.GroupName = &v
	return s
}

func (s *CreateNacosServiceRequest) SetInstanceId(v string) *CreateNacosServiceRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateNacosServiceRequest) SetNamespaceId(v string) *CreateNacosServiceRequest {
	s.NamespaceId = &v
	return s
}

func (s *CreateNacosServiceRequest) SetProtectThreshold(v string) *CreateNacosServiceRequest {
	s.ProtectThreshold = &v
	return s
}

func (s *CreateNacosServiceRequest) SetServiceName(v string) *CreateNacosServiceRequest {
	s.ServiceName = &v
	return s
}

func (s *CreateNacosServiceRequest) Validate() error {
	return dara.Validate(s)
}
