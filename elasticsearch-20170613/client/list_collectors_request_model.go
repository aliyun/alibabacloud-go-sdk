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
	// The ID of the resource with which the shipper is associated.
	//
	// example:
	//
	// es-cn-nif1q8auz0003****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The name of the shipper.
	//
	// example:
	//
	// collectorName1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The number of the page to return. Valid values: 1 to 200. Default value: 1.
	//
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// The ID of the shipper.
	//
	// example:
	//
	// ct-cn-77uqof2s7rg5c****
	ResId *string `json:"resId,omitempty" xml:"resId,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 500. Default value: 20.
	//
	// example:
	//
	// 10
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// The type of the machine on which the shipper is deployed. If you leave this parameter empty, shippers deployed on all types of machines are returned. Valid values:
	//
	// 	- ECS
	//
	// 	- ACK
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
