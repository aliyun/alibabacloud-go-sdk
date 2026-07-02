// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListInstancesRequest
	GetClusterId() *string
	SetCurrent(v int64) *ListInstancesRequest
	GetCurrent() *int64
	SetInstance(v string) *ListInstancesRequest
	GetInstance() *string
	SetPageSize(v int64) *ListInstancesRequest
	GetPageSize() *int64
	SetRegion(v string) *ListInstancesRequest
	GetRegion() *string
	SetStatus(v string) *ListInstancesRequest
	GetStatus() *string
}

type ListInstancesRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// xxxxx
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// The current page number. This field exists when pagination is used.
	//
	// example:
	//
	// 1
	Current *int64 `json:"current,omitempty" xml:"current,omitempty"`
	// The ECS instance ID used to filter results.
	//
	// example:
	//
	// i-8vbfd3g9vs32sfuvv38h
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// The number of entries per page. Default value: 20. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Filters instances by region.
	//
	// example:
	//
	// cn-shenzhen
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// Filters instances by status.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListInstancesRequest) GetCurrent() *int64 {
	return s.Current
}

func (s *ListInstancesRequest) GetInstance() *string {
	return s.Instance
}

func (s *ListInstancesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListInstancesRequest) GetRegion() *string {
	return s.Region
}

func (s *ListInstancesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListInstancesRequest) SetClusterId(v string) *ListInstancesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListInstancesRequest) SetCurrent(v int64) *ListInstancesRequest {
	s.Current = &v
	return s
}

func (s *ListInstancesRequest) SetInstance(v string) *ListInstancesRequest {
	s.Instance = &v
	return s
}

func (s *ListInstancesRequest) SetPageSize(v int64) *ListInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstancesRequest) SetRegion(v string) *ListInstancesRequest {
	s.Region = &v
	return s
}

func (s *ListInstancesRequest) SetStatus(v string) *ListInstancesRequest {
	s.Status = &v
	return s
}

func (s *ListInstancesRequest) Validate() error {
	return dara.Validate(s)
}
