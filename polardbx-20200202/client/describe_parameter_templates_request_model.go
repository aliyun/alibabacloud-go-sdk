// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParameterTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeParameterTemplatesRequest
	GetDBInstanceId() *string
	SetEngineVersion(v string) *DescribeParameterTemplatesRequest
	GetEngineVersion() *string
	SetParamLevel(v string) *DescribeParameterTemplatesRequest
	GetParamLevel() *string
	SetRegionId(v string) *DescribeParameterTemplatesRequest
	GetRegionId() *string
}

type DescribeParameterTemplatesRequest struct {
	// example:
	//
	// pxc-********
	DBInstanceId  *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
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

func (s DescribeParameterTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeParameterTemplatesRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeParameterTemplatesRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeParameterTemplatesRequest) GetParamLevel() *string {
	return s.ParamLevel
}

func (s *DescribeParameterTemplatesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeParameterTemplatesRequest) SetDBInstanceId(v string) *DescribeParameterTemplatesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetEngineVersion(v string) *DescribeParameterTemplatesRequest {
	s.EngineVersion = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetParamLevel(v string) *DescribeParameterTemplatesRequest {
	s.ParamLevel = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetRegionId(v string) *DescribeParameterTemplatesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
