// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExtensionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DeleteExtensionRequest
	GetDBInstanceId() *string
	SetDBNames(v string) *DeleteExtensionRequest
	GetDBNames() *string
	SetExtension(v string) *DeleteExtensionRequest
	GetExtension() *string
	SetRegionId(v string) *DeleteExtensionRequest
	GetRegionId() *string
}

type DeleteExtensionRequest struct {
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
	// The name of the extension.
	//
	// This parameter is required.
	//
	// example:
	//
	// citext
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteExtensionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteExtensionRequest) GoString() string {
	return s.String()
}

func (s *DeleteExtensionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteExtensionRequest) GetDBNames() *string {
	return s.DBNames
}

func (s *DeleteExtensionRequest) GetExtension() *string {
	return s.Extension
}

func (s *DeleteExtensionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteExtensionRequest) SetDBInstanceId(v string) *DeleteExtensionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteExtensionRequest) SetDBNames(v string) *DeleteExtensionRequest {
	s.DBNames = &v
	return s
}

func (s *DeleteExtensionRequest) SetExtension(v string) *DeleteExtensionRequest {
	s.Extension = &v
	return s
}

func (s *DeleteExtensionRequest) SetRegionId(v string) *DeleteExtensionRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteExtensionRequest) Validate() error {
	return dara.Validate(s)
}
