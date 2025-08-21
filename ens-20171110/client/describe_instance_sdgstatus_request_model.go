// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceSDGStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeInstanceSDGStatusRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *DescribeInstanceSDGStatusRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeInstanceSDGStatusRequest
	GetPageSize() *int32
	SetSDGIds(v []*string) *DescribeInstanceSDGStatusRequest
	GetSDGIds() []*string
	SetStatus(v string) *DescribeInstanceSDGStatusRequest
	GetStatus() *string
}

type DescribeInstanceSDGStatusRequest struct {
	// The ID of the AIC instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// aic-xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The IDs of SDGs that you want to query. By default, all SDGs are queried.
	SDGIds []*string `json:"SDGIds,omitempty" xml:"SDGIds,omitempty" type:"Repeated"`
	// The deployment status of the SDG.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeInstanceSDGStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSDGStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSDGStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceSDGStatusRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeInstanceSDGStatusRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstanceSDGStatusRequest) GetSDGIds() []*string {
	return s.SDGIds
}

func (s *DescribeInstanceSDGStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeInstanceSDGStatusRequest) SetInstanceId(v string) *DescribeInstanceSDGStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceSDGStatusRequest) SetPageNumber(v int32) *DescribeInstanceSDGStatusRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstanceSDGStatusRequest) SetPageSize(v int32) *DescribeInstanceSDGStatusRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceSDGStatusRequest) SetSDGIds(v []*string) *DescribeInstanceSDGStatusRequest {
	s.SDGIds = v
	return s
}

func (s *DescribeInstanceSDGStatusRequest) SetStatus(v string) *DescribeInstanceSDGStatusRequest {
	s.Status = &v
	return s
}

func (s *DescribeInstanceSDGStatusRequest) Validate() error {
	return dara.Validate(s)
}
