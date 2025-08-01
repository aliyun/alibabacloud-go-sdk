// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEngineNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateEngineNamespaceRequest
	GetAcceptLanguage() *string
	SetClusterId(v string) *UpdateEngineNamespaceRequest
	GetClusterId() *string
	SetDesc(v string) *UpdateEngineNamespaceRequest
	GetDesc() *string
	SetId(v string) *UpdateEngineNamespaceRequest
	GetId() *string
	SetInstanceId(v string) *UpdateEngineNamespaceRequest
	GetInstanceId() *string
	SetName(v string) *UpdateEngineNamespaceRequest
	GetName() *string
	SetServiceCount(v int32) *UpdateEngineNamespaceRequest
	GetServiceCount() *int32
}

type UpdateEngineNamespaceRequest struct {
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
	// example:
	//
	// mse-09k1q11****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The description of the cluster.
	//
	// example:
	//
	// public
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The ID of the namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 33ff74b6-d21e-4f9b-91a8-bc1ea4ef****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// mse-cn-st21ri2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of active services.
	//
	// example:
	//
	// 3
	ServiceCount *int32 `json:"ServiceCount,omitempty" xml:"ServiceCount,omitempty"`
}

func (s UpdateEngineNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEngineNamespaceRequest) GoString() string {
	return s.String()
}

func (s *UpdateEngineNamespaceRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateEngineNamespaceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateEngineNamespaceRequest) GetDesc() *string {
	return s.Desc
}

func (s *UpdateEngineNamespaceRequest) GetId() *string {
	return s.Id
}

func (s *UpdateEngineNamespaceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateEngineNamespaceRequest) GetName() *string {
	return s.Name
}

func (s *UpdateEngineNamespaceRequest) GetServiceCount() *int32 {
	return s.ServiceCount
}

func (s *UpdateEngineNamespaceRequest) SetAcceptLanguage(v string) *UpdateEngineNamespaceRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateEngineNamespaceRequest) SetClusterId(v string) *UpdateEngineNamespaceRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateEngineNamespaceRequest) SetDesc(v string) *UpdateEngineNamespaceRequest {
	s.Desc = &v
	return s
}

func (s *UpdateEngineNamespaceRequest) SetId(v string) *UpdateEngineNamespaceRequest {
	s.Id = &v
	return s
}

func (s *UpdateEngineNamespaceRequest) SetInstanceId(v string) *UpdateEngineNamespaceRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateEngineNamespaceRequest) SetName(v string) *UpdateEngineNamespaceRequest {
	s.Name = &v
	return s
}

func (s *UpdateEngineNamespaceRequest) SetServiceCount(v int32) *UpdateEngineNamespaceRequest {
	s.ServiceCount = &v
	return s
}

func (s *UpdateEngineNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
