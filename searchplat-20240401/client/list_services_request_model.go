// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModelType(v string) *ListServicesRequest
	GetModelType() *string
	SetName(v string) *ListServicesRequest
	GetName() *string
	SetServiceId(v string) *ListServicesRequest
	GetServiceId() *string
	SetServiceType(v string) *ListServicesRequest
	GetServiceType() *string
}

type ListServicesRequest struct {
	// The model type. Valid values:
	//
	// - system: built-in model
	//
	// - deployment: custom deployment model.
	//
	// example:
	//
	// system
	ModelType *string `json:"modelType,omitempty" xml:"modelType,omitempty"`
	// The service name.
	//
	// example:
	//
	// 文本向量化
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The service ID.
	//
	// example:
	//
	// ops-text-embedding-001
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// The service type.
	//
	// example:
	//
	// text-embedding
	ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s ListServicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServicesRequest) GoString() string {
	return s.String()
}

func (s *ListServicesRequest) GetModelType() *string {
	return s.ModelType
}

func (s *ListServicesRequest) GetName() *string {
	return s.Name
}

func (s *ListServicesRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListServicesRequest) GetServiceType() *string {
	return s.ServiceType
}

func (s *ListServicesRequest) SetModelType(v string) *ListServicesRequest {
	s.ModelType = &v
	return s
}

func (s *ListServicesRequest) SetName(v string) *ListServicesRequest {
	s.Name = &v
	return s
}

func (s *ListServicesRequest) SetServiceId(v string) *ListServicesRequest {
	s.ServiceId = &v
	return s
}

func (s *ListServicesRequest) SetServiceType(v string) *ListServicesRequest {
	s.ServiceType = &v
	return s
}

func (s *ListServicesRequest) Validate() error {
	return dara.Validate(s)
}
