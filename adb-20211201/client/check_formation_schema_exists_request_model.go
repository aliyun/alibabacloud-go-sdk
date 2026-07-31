// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckFormationSchemaExistsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CheckFormationSchemaExistsRequest
	GetDBClusterId() *string
	SetPrefixMode(v bool) *CheckFormationSchemaExistsRequest
	GetPrefixMode() *bool
	SetRegionId(v string) *CheckFormationSchemaExistsRequest
	GetRegionId() *string
	SetSchema(v string) *CheckFormationSchemaExistsRequest
	GetSchema() *string
}

type CheckFormationSchemaExistsRequest struct {
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the cluster IDs of all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters in a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether to enable prefix mode. Valid values:
	//
	// - true: Enable prefix mode.
	//
	// - false: Disable prefix mode.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	PrefixMode *bool `json:"PrefixMode,omitempty" xml:"PrefixMode,omitempty"`
	// RegionId
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// schema 。
	//
	// This parameter is required.
	//
	// example:
	//
	// testdb01
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
}

func (s CheckFormationSchemaExistsRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckFormationSchemaExistsRequest) GoString() string {
	return s.String()
}

func (s *CheckFormationSchemaExistsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CheckFormationSchemaExistsRequest) GetPrefixMode() *bool {
	return s.PrefixMode
}

func (s *CheckFormationSchemaExistsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CheckFormationSchemaExistsRequest) GetSchema() *string {
	return s.Schema
}

func (s *CheckFormationSchemaExistsRequest) SetDBClusterId(v string) *CheckFormationSchemaExistsRequest {
	s.DBClusterId = &v
	return s
}

func (s *CheckFormationSchemaExistsRequest) SetPrefixMode(v bool) *CheckFormationSchemaExistsRequest {
	s.PrefixMode = &v
	return s
}

func (s *CheckFormationSchemaExistsRequest) SetRegionId(v string) *CheckFormationSchemaExistsRequest {
	s.RegionId = &v
	return s
}

func (s *CheckFormationSchemaExistsRequest) SetSchema(v string) *CheckFormationSchemaExistsRequest {
	s.Schema = &v
	return s
}

func (s *CheckFormationSchemaExistsRequest) Validate() error {
	return dara.Validate(s)
}
