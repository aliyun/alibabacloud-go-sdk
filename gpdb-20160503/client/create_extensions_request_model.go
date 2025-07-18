// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExtensionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *CreateExtensionsRequest
	GetDBInstanceId() *string
	SetDBNames(v string) *CreateExtensionsRequest
	GetDBNames() *string
	SetExtensions(v string) *CreateExtensionsRequest
	GetExtensions() *string
	SetRegionId(v string) *CreateExtensionsRequest
	GetRegionId() *string
}

type CreateExtensionsRequest struct {
	// The instance ID.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the IDs of all AnalyticDB for PostgreSQL instances in a specific region.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo1
	DBNames *string `json:"DBNames,omitempty" xml:"DBNames,omitempty"`
	// The name of the extension that you want to install. Separate multiple extension names with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// citext, dblink
	Extensions *string `json:"Extensions,omitempty" xml:"Extensions,omitempty"`
	// The region ID.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateExtensionsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateExtensionsRequest) GoString() string {
	return s.String()
}

func (s *CreateExtensionsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateExtensionsRequest) GetDBNames() *string {
	return s.DBNames
}

func (s *CreateExtensionsRequest) GetExtensions() *string {
	return s.Extensions
}

func (s *CreateExtensionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateExtensionsRequest) SetDBInstanceId(v string) *CreateExtensionsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateExtensionsRequest) SetDBNames(v string) *CreateExtensionsRequest {
	s.DBNames = &v
	return s
}

func (s *CreateExtensionsRequest) SetExtensions(v string) *CreateExtensionsRequest {
	s.Extensions = &v
	return s
}

func (s *CreateExtensionsRequest) SetRegionId(v string) *CreateExtensionsRequest {
	s.RegionId = &v
	return s
}

func (s *CreateExtensionsRequest) Validate() error {
	return dara.Validate(s)
}
