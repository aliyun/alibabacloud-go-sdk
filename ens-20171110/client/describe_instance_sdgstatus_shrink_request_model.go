// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceSDGStatusShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeInstanceSDGStatusShrinkRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *DescribeInstanceSDGStatusShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeInstanceSDGStatusShrinkRequest
	GetPageSize() *int32
	SetSDGIdsShrink(v string) *DescribeInstanceSDGStatusShrinkRequest
	GetSDGIdsShrink() *string
	SetStatus(v string) *DescribeInstanceSDGStatusShrinkRequest
	GetStatus() *string
}

type DescribeInstanceSDGStatusShrinkRequest struct {
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
	SDGIdsShrink *string `json:"SDGIds,omitempty" xml:"SDGIds,omitempty"`
	// The deployment status of the SDG.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeInstanceSDGStatusShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSDGStatusShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSDGStatusShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceSDGStatusShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeInstanceSDGStatusShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstanceSDGStatusShrinkRequest) GetSDGIdsShrink() *string {
	return s.SDGIdsShrink
}

func (s *DescribeInstanceSDGStatusShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeInstanceSDGStatusShrinkRequest) SetInstanceId(v string) *DescribeInstanceSDGStatusShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceSDGStatusShrinkRequest) SetPageNumber(v int32) *DescribeInstanceSDGStatusShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstanceSDGStatusShrinkRequest) SetPageSize(v int32) *DescribeInstanceSDGStatusShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceSDGStatusShrinkRequest) SetSDGIdsShrink(v string) *DescribeInstanceSDGStatusShrinkRequest {
	s.SDGIdsShrink = &v
	return s
}

func (s *DescribeInstanceSDGStatusShrinkRequest) SetStatus(v string) *DescribeInstanceSDGStatusShrinkRequest {
	s.Status = &v
	return s
}

func (s *DescribeInstanceSDGStatusShrinkRequest) Validate() error {
	return dara.Validate(s)
}
