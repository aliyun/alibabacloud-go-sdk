// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParametersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeParametersRequest
	GetDBInstanceId() *string
	SetParamLevel(v string) *DescribeParametersRequest
	GetParamLevel() *string
	SetRegionId(v string) *DescribeParametersRequest
	GetRegionId() *string
}

type DescribeParametersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// compute
	ParamLevel *string `json:"ParamLevel,omitempty" xml:"ParamLevel,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeParametersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeParametersRequest) GoString() string {
	return s.String()
}

func (s *DescribeParametersRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeParametersRequest) GetParamLevel() *string {
	return s.ParamLevel
}

func (s *DescribeParametersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeParametersRequest) SetDBInstanceId(v string) *DescribeParametersRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeParametersRequest) SetParamLevel(v string) *DescribeParametersRequest {
	s.ParamLevel = &v
	return s
}

func (s *DescribeParametersRequest) SetRegionId(v string) *DescribeParametersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeParametersRequest) Validate() error {
	return dara.Validate(s)
}
