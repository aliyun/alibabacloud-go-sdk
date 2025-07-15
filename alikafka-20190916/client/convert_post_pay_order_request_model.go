// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertPostPayOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDuration(v int32) *ConvertPostPayOrderRequest
	GetDuration() *int32
	SetInstanceId(v string) *ConvertPostPayOrderRequest
	GetInstanceId() *string
	SetPaidType(v int32) *ConvertPostPayOrderRequest
	GetPaidType() *int32
	SetRegionId(v string) *ConvertPostPayOrderRequest
	GetRegionId() *string
}

type ConvertPostPayOrderRequest struct {
	// The subscription duration. Unit: months. Valid values:
	//
	// 	- **1~12**
	//
	// 	- **24**
	//
	// 	- **36**
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-v0h1fgs2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PaidType   *int32  `json:"PaidType,omitempty" xml:"PaidType,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ConvertPostPayOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s ConvertPostPayOrderRequest) GoString() string {
	return s.String()
}

func (s *ConvertPostPayOrderRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *ConvertPostPayOrderRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ConvertPostPayOrderRequest) GetPaidType() *int32 {
	return s.PaidType
}

func (s *ConvertPostPayOrderRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ConvertPostPayOrderRequest) SetDuration(v int32) *ConvertPostPayOrderRequest {
	s.Duration = &v
	return s
}

func (s *ConvertPostPayOrderRequest) SetInstanceId(v string) *ConvertPostPayOrderRequest {
	s.InstanceId = &v
	return s
}

func (s *ConvertPostPayOrderRequest) SetPaidType(v int32) *ConvertPostPayOrderRequest {
	s.PaidType = &v
	return s
}

func (s *ConvertPostPayOrderRequest) SetRegionId(v string) *ConvertPostPayOrderRequest {
	s.RegionId = &v
	return s
}

func (s *ConvertPostPayOrderRequest) Validate() error {
	return dara.Validate(s)
}
