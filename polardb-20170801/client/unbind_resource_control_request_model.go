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
	// >You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query region IDs.
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
	// The type of the target to unbind. Valid values: USER, DATABASE, QUERY, CONNECTION. The value is case-insensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// USER
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The value of the target to unbind. The format is the same as the TargetValue for the corresponding target type in the BindResourceControl operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_user
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
