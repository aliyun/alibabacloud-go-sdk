// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCollectorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListCollectorsRequest
	GetInstanceId() *string
	SetName(v string) *ListCollectorsRequest
	GetName() *string
	SetPage(v int32) *ListCollectorsRequest
	GetPage() *int32
	SetResId(v string) *ListCollectorsRequest
	GetResId() *string
	SetSize(v int32) *ListCollectorsRequest
	GetSize() *int32
	SetSourceType(v string) *ListCollectorsRequest
	GetSourceType() *string
}

type ListCollectorsRequest struct {
	// The instance ID associated with the collector.
	//
	// example:
	//
	// es-cn-nif1q8auz0003****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The collector name.
	//
	// example:
	//
	// collectorName1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The page number of the returned results. Default value: 1. Minimum value: 1. Maximum value: 200.
	//
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// The collector ID.
	//
	// example:
	//
	// ct-cn-77uqof2s7rg5c****
	ResId *string `json:"resId,omitempty" xml:"resId,omitempty"`
	// The number of results per page. Default value: 20. Minimum value: 1. Maximum value: 500.
	//
	// example:
	//
	// 10
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// The type of machine on which the collector is deployed. If this parameter is not specified, all types are returned. Valid values:
	//
	// - ECS: ECS instance
	//
	// - ACK: Container Kubernetes cluster.
	//
	// example:
	//
	// ECS
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
}

func (s ListCollectorsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCollectorsRequest) GoString() string {
	return s.String()
}

func (s *ListCollectorsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCollectorsRequest) GetName() *string {
	return s.Name
}

func (s *ListCollectorsRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListCollectorsRequest) GetResId() *string {
	return s.ResId
}

func (s *ListCollectorsRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListCollectorsRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *ListCollectorsRequest) SetInstanceId(v string) *ListCollectorsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCollectorsRequest) SetName(v string) *ListCollectorsRequest {
	s.Name = &v
	return s
}

func (s *ListCollectorsRequest) SetPage(v int32) *ListCollectorsRequest {
	s.Page = &v
	return s
}

func (s *ListCollectorsRequest) SetResId(v string) *ListCollectorsRequest {
	s.ResId = &v
	return s
}

func (s *ListCollectorsRequest) SetSize(v int32) *ListCollectorsRequest {
	s.Size = &v
	return s
}

func (s *ListCollectorsRequest) SetSourceType(v string) *ListCollectorsRequest {
	s.SourceType = &v
	return s
}

func (s *ListCollectorsRequest) Validate() error {
	return dara.Validate(s)
}
