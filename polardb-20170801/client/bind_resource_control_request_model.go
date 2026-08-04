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
	// The PolarDB cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID of the PolarDB cluster.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query available regions.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the resource control rule. The name must be 1 to 63 ASCII bytes in length, start with a letter, and can contain only letters, digits, and underscores.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_rc
	ResourceControlName *string `json:"ResourceControlName,omitempty" xml:"ResourceControlName,omitempty"`
	// The type of the binding target. Valid values: USER, DATABASE, QUERY, CONNECTION. The value is case-insensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// USER
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The value of the binding target. The format depends on the value of TargetType. For more information, see the table below.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_user
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
