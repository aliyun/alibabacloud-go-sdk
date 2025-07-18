// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeExtensionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *UpgradeExtensionsRequest
	GetDBInstanceId() *string
	SetDatabaseName(v string) *UpgradeExtensionsRequest
	GetDatabaseName() *string
	SetExtensions(v string) *UpgradeExtensionsRequest
	GetExtensions() *string
	SetRegionId(v string) *UpgradeExtensionsRequest
	GetRegionId() *string
}

type UpgradeExtensionsRequest struct {
	// The instance ID.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Database name.
	//
	// example:
	//
	// test01
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The extensions that you want to update. Separate multiple extensions with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// citext,dblink
	Extensions *string `json:"Extensions,omitempty" xml:"Extensions,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpgradeExtensionsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeExtensionsRequest) GoString() string {
	return s.String()
}

func (s *UpgradeExtensionsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *UpgradeExtensionsRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *UpgradeExtensionsRequest) GetExtensions() *string {
	return s.Extensions
}

func (s *UpgradeExtensionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpgradeExtensionsRequest) SetDBInstanceId(v string) *UpgradeExtensionsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpgradeExtensionsRequest) SetDatabaseName(v string) *UpgradeExtensionsRequest {
	s.DatabaseName = &v
	return s
}

func (s *UpgradeExtensionsRequest) SetExtensions(v string) *UpgradeExtensionsRequest {
	s.Extensions = &v
	return s
}

func (s *UpgradeExtensionsRequest) SetRegionId(v string) *UpgradeExtensionsRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeExtensionsRequest) Validate() error {
	return dara.Validate(s)
}
