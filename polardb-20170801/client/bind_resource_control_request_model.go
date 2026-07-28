// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindResourceControlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *BindResourceControlRequest
	GetDBClusterId() *string
	SetRegionId(v string) *BindResourceControlRequest
	GetRegionId() *string
	SetResourceControlName(v string) *BindResourceControlRequest
	GetResourceControlName() *string
	SetTargetType(v string) *BindResourceControlRequest
	GetTargetType() *string
	SetTargetValue(v string) *BindResourceControlRequest
	GetTargetValue() *string
}

type BindResourceControlRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pm-xxxxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query available regions.
	//
	// example:
	//
	// cn-beijing
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
	// 646060ae9852280007a62545,68b696cb0a7fa600078d41af
	TargetValue *string `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
}

func (s BindResourceControlRequest) String() string {
	return dara.Prettify(s)
}

func (s BindResourceControlRequest) GoString() string {
	return s.String()
}

func (s *BindResourceControlRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *BindResourceControlRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *BindResourceControlRequest) GetResourceControlName() *string {
	return s.ResourceControlName
}

func (s *BindResourceControlRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *BindResourceControlRequest) GetTargetValue() *string {
	return s.TargetValue
}

func (s *BindResourceControlRequest) SetDBClusterId(v string) *BindResourceControlRequest {
	s.DBClusterId = &v
	return s
}

func (s *BindResourceControlRequest) SetRegionId(v string) *BindResourceControlRequest {
	s.RegionId = &v
	return s
}

func (s *BindResourceControlRequest) SetResourceControlName(v string) *BindResourceControlRequest {
	s.ResourceControlName = &v
	return s
}

func (s *BindResourceControlRequest) SetTargetType(v string) *BindResourceControlRequest {
	s.TargetType = &v
	return s
}

func (s *BindResourceControlRequest) SetTargetValue(v string) *BindResourceControlRequest {
	s.TargetValue = &v
	return s
}

func (s *BindResourceControlRequest) Validate() error {
	return dara.Validate(s)
}
