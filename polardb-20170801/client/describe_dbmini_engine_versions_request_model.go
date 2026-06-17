// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBMiniEngineVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArchitecture(v string) *DescribeDBMiniEngineVersionsRequest
	GetArchitecture() *string
	SetCreationCategory(v string) *DescribeDBMiniEngineVersionsRequest
	GetCreationCategory() *string
	SetDBMinorVersion(v string) *DescribeDBMiniEngineVersionsRequest
	GetDBMinorVersion() *string
	SetDBType(v string) *DescribeDBMiniEngineVersionsRequest
	GetDBType() *string
	SetDBVersion(v string) *DescribeDBMiniEngineVersionsRequest
	GetDBVersion() *string
	SetRegionId(v string) *DescribeDBMiniEngineVersionsRequest
	GetRegionId() *string
	SetZoneId(v string) *DescribeDBMiniEngineVersionsRequest
	GetZoneId() *string
}

type DescribeDBMiniEngineVersionsRequest struct {
	// The CPU architecture. Valid values:
	//
	// - **X86**
	//
	// - **ARM**
	//
	// example:
	//
	// X86
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// The product series. Valid values:
	//
	// - **Normal**: Cluster Edition (default)
	//
	// - **SENormal**: Standard Edition
	//
	// For more information about product series, see [Product series](https://help.aliyun.com/document_detail/183258.html).
	//
	// example:
	//
	// Normal
	CreationCategory *string `json:"CreationCategory,omitempty" xml:"CreationCategory,omitempty"`
	// The minor version of the database engine.
	//
	// - If `DBVersion` is set to **8.0**, valid values are:
	//
	//   - **8.0.2**
	//
	//   - **8.0.1**
	//
	// - If `DBVersion` is set to **5.7**, the valid value is **5.7.28**.
	//
	// - If `DBVersion` is set to **5.6**, the valid value is **5.6.16**.
	//
	// example:
	//
	// 8.0.1
	DBMinorVersion *string `json:"DBMinorVersion,omitempty" xml:"DBMinorVersion,omitempty"`
	// The database type. The only valid value is **MySQL**.
	//
	// - **MySQL**.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// The major version of the database engine. Valid values:
	//
	// - **8.0**
	//
	// - **5.7**
	//
	// - **5.6**
	//
	// This parameter is required.
	//
	// example:
	//
	// 5.6
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The zone.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBMiniEngineVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBMiniEngineVersionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBMiniEngineVersionsRequest) GetArchitecture() *string {
	return s.Architecture
}

func (s *DescribeDBMiniEngineVersionsRequest) GetCreationCategory() *string {
	return s.CreationCategory
}

func (s *DescribeDBMiniEngineVersionsRequest) GetDBMinorVersion() *string {
	return s.DBMinorVersion
}

func (s *DescribeDBMiniEngineVersionsRequest) GetDBType() *string {
	return s.DBType
}

func (s *DescribeDBMiniEngineVersionsRequest) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeDBMiniEngineVersionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBMiniEngineVersionsRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBMiniEngineVersionsRequest) SetArchitecture(v string) *DescribeDBMiniEngineVersionsRequest {
	s.Architecture = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsRequest) SetCreationCategory(v string) *DescribeDBMiniEngineVersionsRequest {
	s.CreationCategory = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsRequest) SetDBMinorVersion(v string) *DescribeDBMiniEngineVersionsRequest {
	s.DBMinorVersion = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsRequest) SetDBType(v string) *DescribeDBMiniEngineVersionsRequest {
	s.DBType = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsRequest) SetDBVersion(v string) *DescribeDBMiniEngineVersionsRequest {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsRequest) SetRegionId(v string) *DescribeDBMiniEngineVersionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsRequest) SetZoneId(v string) *DescribeDBMiniEngineVersionsRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsRequest) Validate() error {
	return dara.Validate(s)
}
