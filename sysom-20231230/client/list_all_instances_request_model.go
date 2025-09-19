// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrent(v string) *ListAllInstancesRequest
	GetCurrent() *string
	SetFilters(v string) *ListAllInstancesRequest
	GetFilters() *string
	SetInstanceType(v string) *ListAllInstancesRequest
	GetInstanceType() *string
	SetManagedType(v string) *ListAllInstancesRequest
	GetManagedType() *string
	SetMaxResults(v int32) *ListAllInstancesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAllInstancesRequest
	GetNextToken() *string
	SetPageSize(v string) *ListAllInstancesRequest
	GetPageSize() *string
	SetPluginId(v string) *ListAllInstancesRequest
	GetPluginId() *string
	SetRegion(v string) *ListAllInstancesRequest
	GetRegion() *string
}

type ListAllInstancesRequest struct {
	// example:
	//
	// 1
	Current *string `json:"current,omitempty" xml:"current,omitempty"`
	// example:
	//
	// {}
	Filters *string `json:"filters,omitempty" xml:"filters,omitempty"`
	// example:
	//
	// ecs
	InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	// example:
	//
	// managed
	ManagedType *string `json:"managedType,omitempty" xml:"managedType,omitempty"`
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// U+w1wv2R4ZWR5oZLXD0+Dp4dD+2BRJj42DLT6GrZysw=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 01fc4a0b-f199-4885-9861-b4054a310fe7
	PluginId *string `json:"pluginId,omitempty" xml:"pluginId,omitempty"`
	// example:
	//
	// cn-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
}

func (s ListAllInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAllInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListAllInstancesRequest) GetCurrent() *string {
	return s.Current
}

func (s *ListAllInstancesRequest) GetFilters() *string {
	return s.Filters
}

func (s *ListAllInstancesRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListAllInstancesRequest) GetManagedType() *string {
	return s.ManagedType
}

func (s *ListAllInstancesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAllInstancesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAllInstancesRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListAllInstancesRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *ListAllInstancesRequest) GetRegion() *string {
	return s.Region
}

func (s *ListAllInstancesRequest) SetCurrent(v string) *ListAllInstancesRequest {
	s.Current = &v
	return s
}

func (s *ListAllInstancesRequest) SetFilters(v string) *ListAllInstancesRequest {
	s.Filters = &v
	return s
}

func (s *ListAllInstancesRequest) SetInstanceType(v string) *ListAllInstancesRequest {
	s.InstanceType = &v
	return s
}

func (s *ListAllInstancesRequest) SetManagedType(v string) *ListAllInstancesRequest {
	s.ManagedType = &v
	return s
}

func (s *ListAllInstancesRequest) SetMaxResults(v int32) *ListAllInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAllInstancesRequest) SetNextToken(v string) *ListAllInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *ListAllInstancesRequest) SetPageSize(v string) *ListAllInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAllInstancesRequest) SetPluginId(v string) *ListAllInstancesRequest {
	s.PluginId = &v
	return s
}

func (s *ListAllInstancesRequest) SetRegion(v string) *ListAllInstancesRequest {
	s.Region = &v
	return s
}

func (s *ListAllInstancesRequest) Validate() error {
	return dara.Validate(s)
}
