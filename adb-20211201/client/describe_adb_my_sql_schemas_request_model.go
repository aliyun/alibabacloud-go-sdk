// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdbMySqlSchemasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeAdbMySqlSchemasRequest
	GetDBClusterId() *string
	SetRegionId(v string) *DescribeAdbMySqlSchemasRequest
	GetRegionId() *string
}

type DescribeAdbMySqlSchemasRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp1r053byu48p****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/454314.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAdbMySqlSchemasRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdbMySqlSchemasRequest) GoString() string {
	return s.String()
}

func (s *DescribeAdbMySqlSchemasRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAdbMySqlSchemasRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAdbMySqlSchemasRequest) SetDBClusterId(v string) *DescribeAdbMySqlSchemasRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAdbMySqlSchemasRequest) SetRegionId(v string) *DescribeAdbMySqlSchemasRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAdbMySqlSchemasRequest) Validate() error {
	return dara.Validate(s)
}
