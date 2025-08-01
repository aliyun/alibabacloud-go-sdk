// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNacosServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateNacosServiceRequest
	GetAcceptLanguage() *string
	SetClusterId(v string) *UpdateNacosServiceRequest
	GetClusterId() *string
	SetGroupName(v string) *UpdateNacosServiceRequest
	GetGroupName() *string
	SetInstanceId(v string) *UpdateNacosServiceRequest
	GetInstanceId() *string
	SetNamespaceId(v string) *UpdateNacosServiceRequest
	GetNamespaceId() *string
	SetProtectThreshold(v string) *UpdateNacosServiceRequest
	GetProtectThreshold() *string
	SetServiceName(v string) *UpdateNacosServiceRequest
	GetServiceName() *string
}

type UpdateNacosServiceRequest struct {
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
	// mse-09k1q11****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
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
	// 5e3ee449-b5c0-4aee-b857-32c0acbebf26
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
	// com.dingtalk.doc.thumbnails.pdf.ThumbnailService
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s UpdateNacosServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNacosServiceRequest) GoString() string {
	return s.String()
}

func (s *UpdateNacosServiceRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateNacosServiceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateNacosServiceRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateNacosServiceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateNacosServiceRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *UpdateNacosServiceRequest) GetProtectThreshold() *string {
	return s.ProtectThreshold
}

func (s *UpdateNacosServiceRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *UpdateNacosServiceRequest) SetAcceptLanguage(v string) *UpdateNacosServiceRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateNacosServiceRequest) SetClusterId(v string) *UpdateNacosServiceRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateNacosServiceRequest) SetGroupName(v string) *UpdateNacosServiceRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateNacosServiceRequest) SetInstanceId(v string) *UpdateNacosServiceRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateNacosServiceRequest) SetNamespaceId(v string) *UpdateNacosServiceRequest {
	s.NamespaceId = &v
	return s
}

func (s *UpdateNacosServiceRequest) SetProtectThreshold(v string) *UpdateNacosServiceRequest {
	s.ProtectThreshold = &v
	return s
}

func (s *UpdateNacosServiceRequest) SetServiceName(v string) *UpdateNacosServiceRequest {
	s.ServiceName = &v
	return s
}

func (s *UpdateNacosServiceRequest) Validate() error {
	return dara.Validate(s)
}
