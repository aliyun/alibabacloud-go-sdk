// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEngineNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *CreateEngineNamespaceRequest
	GetAcceptLanguage() *string
	SetClusterId(v string) *CreateEngineNamespaceRequest
	GetClusterId() *string
	SetDesc(v string) *CreateEngineNamespaceRequest
	GetDesc() *string
	SetId(v string) *CreateEngineNamespaceRequest
	GetId() *string
	SetInstanceId(v string) *CreateEngineNamespaceRequest
	GetInstanceId() *string
	SetName(v string) *CreateEngineNamespaceRequest
	GetName() *string
	SetServiceCount(v int32) *CreateEngineNamespaceRequest
	GetServiceCount() *int32
}

type CreateEngineNamespaceRequest struct {
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
	// The ID of the instance.
	//
	// example:
	//
	// mse-98s****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The description of the namespace.
	//
	// example:
	//
	// Development environment
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The custom ID of the namespace. If you do not specify this parameter, the automatically generated Universally Unique Identifier (UUID) is returned.
	//
	// example:
	//
	// f4fa5b81-2f26-4900-833a-7516b315ebb2
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// mse-cn-st21ri2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The display name of the namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// dev
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The maximum number of services that can run in the namespace.
	//
	// example:
	//
	// 100
	ServiceCount *int32 `json:"ServiceCount,omitempty" xml:"ServiceCount,omitempty"`
}

func (s CreateEngineNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEngineNamespaceRequest) GoString() string {
	return s.String()
}

func (s *CreateEngineNamespaceRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *CreateEngineNamespaceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateEngineNamespaceRequest) GetDesc() *string {
	return s.Desc
}

func (s *CreateEngineNamespaceRequest) GetId() *string {
	return s.Id
}

func (s *CreateEngineNamespaceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateEngineNamespaceRequest) GetName() *string {
	return s.Name
}

func (s *CreateEngineNamespaceRequest) GetServiceCount() *int32 {
	return s.ServiceCount
}

func (s *CreateEngineNamespaceRequest) SetAcceptLanguage(v string) *CreateEngineNamespaceRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *CreateEngineNamespaceRequest) SetClusterId(v string) *CreateEngineNamespaceRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateEngineNamespaceRequest) SetDesc(v string) *CreateEngineNamespaceRequest {
	s.Desc = &v
	return s
}

func (s *CreateEngineNamespaceRequest) SetId(v string) *CreateEngineNamespaceRequest {
	s.Id = &v
	return s
}

func (s *CreateEngineNamespaceRequest) SetInstanceId(v string) *CreateEngineNamespaceRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateEngineNamespaceRequest) SetName(v string) *CreateEngineNamespaceRequest {
	s.Name = &v
	return s
}

func (s *CreateEngineNamespaceRequest) SetServiceCount(v int32) *CreateEngineNamespaceRequest {
	s.ServiceCount = &v
	return s
}

func (s *CreateEngineNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
