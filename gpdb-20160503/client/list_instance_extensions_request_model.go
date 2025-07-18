// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceExtensionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ListInstanceExtensionsRequest
	GetDBInstanceId() *string
	SetExtension(v string) *ListInstanceExtensionsRequest
	GetExtension() *string
	SetInstallStatus(v string) *ListInstanceExtensionsRequest
	GetInstallStatus() *string
	SetPageNumber(v int32) *ListInstanceExtensionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListInstanceExtensionsRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListInstanceExtensionsRequest
	GetRegionId() *string
}

type ListInstanceExtensionsRequest struct {
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
	// The name of the extension.
	//
	// example:
	//
	// citext
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The installation status of the extension. Valid values:
	//
	// 	- installed
	//
	// 	- installing
	//
	// 	- uninstalled
	//
	// example:
	//
	// installed
	InstallStatus *string `json:"InstallStatus,omitempty" xml:"InstallStatus,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **30**
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListInstanceExtensionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceExtensionsRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceExtensionsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ListInstanceExtensionsRequest) GetExtension() *string {
	return s.Extension
}

func (s *ListInstanceExtensionsRequest) GetInstallStatus() *string {
	return s.InstallStatus
}

func (s *ListInstanceExtensionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstanceExtensionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstanceExtensionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInstanceExtensionsRequest) SetDBInstanceId(v string) *ListInstanceExtensionsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ListInstanceExtensionsRequest) SetExtension(v string) *ListInstanceExtensionsRequest {
	s.Extension = &v
	return s
}

func (s *ListInstanceExtensionsRequest) SetInstallStatus(v string) *ListInstanceExtensionsRequest {
	s.InstallStatus = &v
	return s
}

func (s *ListInstanceExtensionsRequest) SetPageNumber(v int32) *ListInstanceExtensionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceExtensionsRequest) SetPageSize(v int32) *ListInstanceExtensionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstanceExtensionsRequest) SetRegionId(v string) *ListInstanceExtensionsRequest {
	s.RegionId = &v
	return s
}

func (s *ListInstanceExtensionsRequest) Validate() error {
	return dara.Validate(s)
}
