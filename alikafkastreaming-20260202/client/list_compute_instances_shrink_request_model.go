// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListComputeInstancesShrinkRequest
	GetInstanceId() *string
	SetInstanceIdsShrink(v string) *ListComputeInstancesShrinkRequest
	GetInstanceIdsShrink() *string
	SetMaxResults(v int32) *ListComputeInstancesShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListComputeInstancesShrinkRequest
	GetNextToken() *string
	SetOrderId(v string) *ListComputeInstancesShrinkRequest
	GetOrderId() *string
	SetRegionId(v string) *ListComputeInstancesShrinkRequest
	GetRegionId() *string
}

type ListComputeInstancesShrinkRequest struct {
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	MaxResults        *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken         *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OrderId           *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListComputeInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListComputeInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListComputeInstancesShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListComputeInstancesShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *ListComputeInstancesShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListComputeInstancesShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListComputeInstancesShrinkRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *ListComputeInstancesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListComputeInstancesShrinkRequest) SetInstanceId(v string) *ListComputeInstancesShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ListComputeInstancesShrinkRequest) SetInstanceIdsShrink(v string) *ListComputeInstancesShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *ListComputeInstancesShrinkRequest) SetMaxResults(v int32) *ListComputeInstancesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListComputeInstancesShrinkRequest) SetNextToken(v string) *ListComputeInstancesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListComputeInstancesShrinkRequest) SetOrderId(v string) *ListComputeInstancesShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *ListComputeInstancesShrinkRequest) SetRegionId(v string) *ListComputeInstancesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListComputeInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
