// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSellerInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v int64) *DescribeSellerInstancesRequest
	GetInstanceId() *int64
	SetInstanceStatus(v string) *DescribeSellerInstancesRequest
	GetInstanceStatus() *string
	SetPageIndex(v int32) *DescribeSellerInstancesRequest
	GetPageIndex() *int32
	SetPageSize(v int64) *DescribeSellerInstancesRequest
	GetPageSize() *int64
	SetUserId(v int64) *DescribeSellerInstancesRequest
	GetUserId() *int64
}

type DescribeSellerInstancesRequest struct {
	// example:
	//
	// 5000002763123
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// OPENED
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 5322460655123456
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeSellerInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSellerInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSellerInstancesRequest) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *DescribeSellerInstancesRequest) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *DescribeSellerInstancesRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *DescribeSellerInstancesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeSellerInstancesRequest) GetUserId() *int64 {
	return s.UserId
}

func (s *DescribeSellerInstancesRequest) SetInstanceId(v int64) *DescribeSellerInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSellerInstancesRequest) SetInstanceStatus(v string) *DescribeSellerInstancesRequest {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeSellerInstancesRequest) SetPageIndex(v int32) *DescribeSellerInstancesRequest {
	s.PageIndex = &v
	return s
}

func (s *DescribeSellerInstancesRequest) SetPageSize(v int64) *DescribeSellerInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSellerInstancesRequest) SetUserId(v int64) *DescribeSellerInstancesRequest {
	s.UserId = &v
	return s
}

func (s *DescribeSellerInstancesRequest) Validate() error {
	return dara.Validate(s)
}
