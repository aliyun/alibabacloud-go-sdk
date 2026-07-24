// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListComputeInstancesRequest
	GetInstanceId() *string
	SetInstanceIds(v []*string) *ListComputeInstancesRequest
	GetInstanceIds() []*string
	SetMaxResults(v int32) *ListComputeInstancesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListComputeInstancesRequest
	GetNextToken() *string
	SetOrderId(v string) *ListComputeInstancesRequest
	GetOrderId() *string
	SetRegionId(v string) *ListComputeInstancesRequest
	GetRegionId() *string
}

type ListComputeInstancesRequest struct {
	InstanceId  *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	MaxResults  *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OrderId     *string   `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListComputeInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListComputeInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListComputeInstancesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListComputeInstancesRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *ListComputeInstancesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListComputeInstancesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListComputeInstancesRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *ListComputeInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListComputeInstancesRequest) SetInstanceId(v string) *ListComputeInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListComputeInstancesRequest) SetInstanceIds(v []*string) *ListComputeInstancesRequest {
	s.InstanceIds = v
	return s
}

func (s *ListComputeInstancesRequest) SetMaxResults(v int32) *ListComputeInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListComputeInstancesRequest) SetNextToken(v string) *ListComputeInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *ListComputeInstancesRequest) SetOrderId(v string) *ListComputeInstancesRequest {
	s.OrderId = &v
	return s
}

func (s *ListComputeInstancesRequest) SetRegionId(v string) *ListComputeInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *ListComputeInstancesRequest) Validate() error {
	return dara.Validate(s)
}
