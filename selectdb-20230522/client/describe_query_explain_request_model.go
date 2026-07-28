// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQueryExplainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeQueryExplainRequest
	GetDBInstanceId() *string
	SetMode(v string) *DescribeQueryExplainRequest
	GetMode() *string
	SetQueryId(v string) *DescribeQueryExplainRequest
	GetQueryId() *string
	SetRegionId(v string) *DescribeQueryExplainRequest
	GetRegionId() *string
}

type DescribeQueryExplainRequest struct {
	// The database instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The Explain mode.
	//
	// example:
	//
	// BASIC
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The query ID.
	//
	// example:
	//
	// 40e8a568-d2a2-4f7e-a3f8-cec554ce5143
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeQueryExplainRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeQueryExplainRequest) GoString() string {
	return s.String()
}

func (s *DescribeQueryExplainRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeQueryExplainRequest) GetMode() *string {
	return s.Mode
}

func (s *DescribeQueryExplainRequest) GetQueryId() *string {
	return s.QueryId
}

func (s *DescribeQueryExplainRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeQueryExplainRequest) SetDBInstanceId(v string) *DescribeQueryExplainRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeQueryExplainRequest) SetMode(v string) *DescribeQueryExplainRequest {
	s.Mode = &v
	return s
}

func (s *DescribeQueryExplainRequest) SetQueryId(v string) *DescribeQueryExplainRequest {
	s.QueryId = &v
	return s
}

func (s *DescribeQueryExplainRequest) SetRegionId(v string) *DescribeQueryExplainRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeQueryExplainRequest) Validate() error {
	return dara.Validate(s)
}
