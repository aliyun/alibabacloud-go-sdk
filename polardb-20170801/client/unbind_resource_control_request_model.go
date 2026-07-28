// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindResourceControlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *UnbindResourceControlRequest
	GetDBClusterId() *string
	SetRegionId(v string) *UnbindResourceControlRequest
	GetRegionId() *string
	SetResourceControlName(v string) *UnbindResourceControlRequest
	GetResourceControlName() *string
	SetTargetType(v string) *UnbindResourceControlRequest
	GetTargetType() *string
	SetTargetValue(v string) *UnbindResourceControlRequest
	GetTargetValue() *string
}

type UnbindResourceControlRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-bp10gr51qasnl****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
	//
	// >You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query the region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource control name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-rc
	ResourceControlName *string `json:"ResourceControlName,omitempty" xml:"ResourceControlName,omitempty"`
	// The target instance type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The target value. This parameter applies to target tracking rules and prediction rules. The value of TargetValue can contain up to three decimal places and must be greater than 0.
	//
	// This parameter is required.
	//
	// example:
	//
	// 44abc95973e24ae9838713598f673535
	TargetValue *string `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
}

func (s UnbindResourceControlRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindResourceControlRequest) GoString() string {
	return s.String()
}

func (s *UnbindResourceControlRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *UnbindResourceControlRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UnbindResourceControlRequest) GetResourceControlName() *string {
	return s.ResourceControlName
}

func (s *UnbindResourceControlRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *UnbindResourceControlRequest) GetTargetValue() *string {
	return s.TargetValue
}

func (s *UnbindResourceControlRequest) SetDBClusterId(v string) *UnbindResourceControlRequest {
	s.DBClusterId = &v
	return s
}

func (s *UnbindResourceControlRequest) SetRegionId(v string) *UnbindResourceControlRequest {
	s.RegionId = &v
	return s
}

func (s *UnbindResourceControlRequest) SetResourceControlName(v string) *UnbindResourceControlRequest {
	s.ResourceControlName = &v
	return s
}

func (s *UnbindResourceControlRequest) SetTargetType(v string) *UnbindResourceControlRequest {
	s.TargetType = &v
	return s
}

func (s *UnbindResourceControlRequest) SetTargetValue(v string) *UnbindResourceControlRequest {
	s.TargetValue = &v
	return s
}

func (s *UnbindResourceControlRequest) Validate() error {
	return dara.Validate(s)
}
