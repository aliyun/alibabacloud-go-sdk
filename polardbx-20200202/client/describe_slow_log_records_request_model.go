// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlowLogRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCharacterType(v string) *DescribeSlowLogRecordsRequest
	GetCharacterType() *string
	SetDBInstanceName(v string) *DescribeSlowLogRecordsRequest
	GetDBInstanceName() *string
	SetDBName(v string) *DescribeSlowLogRecordsRequest
	GetDBName() *string
	SetDBNodeIds(v string) *DescribeSlowLogRecordsRequest
	GetDBNodeIds() *string
	SetEndTime(v string) *DescribeSlowLogRecordsRequest
	GetEndTime() *string
	SetPage(v int32) *DescribeSlowLogRecordsRequest
	GetPage() *int32
	SetPageSize(v int32) *DescribeSlowLogRecordsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeSlowLogRecordsRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeSlowLogRecordsRequest
	GetStartTime() *string
}

type DescribeSlowLogRecordsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// polarx_cn
	CharacterType *string `json:"CharacterType,omitempty" xml:"CharacterType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-bjxxxxxxxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// testdb
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// example:
	//
	// pxc-i-mezcj4ejdz
	DBNodeIds *string `json:"DBNodeIds,omitempty" xml:"DBNodeIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024-11-22T02:22Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024-10-09T02:26
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSlowLogRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsRequest) GetCharacterType() *string {
	return s.CharacterType
}

func (s *DescribeSlowLogRecordsRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeSlowLogRecordsRequest) GetDBName() *string {
	return s.DBName
}

func (s *DescribeSlowLogRecordsRequest) GetDBNodeIds() *string {
	return s.DBNodeIds
}

func (s *DescribeSlowLogRecordsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSlowLogRecordsRequest) GetPage() *int32 {
	return s.Page
}

func (s *DescribeSlowLogRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSlowLogRecordsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSlowLogRecordsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSlowLogRecordsRequest) SetCharacterType(v string) *DescribeSlowLogRecordsRequest {
	s.CharacterType = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetDBInstanceName(v string) *DescribeSlowLogRecordsRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetDBName(v string) *DescribeSlowLogRecordsRequest {
	s.DBName = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetDBNodeIds(v string) *DescribeSlowLogRecordsRequest {
	s.DBNodeIds = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetEndTime(v string) *DescribeSlowLogRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetPage(v int32) *DescribeSlowLogRecordsRequest {
	s.Page = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetPageSize(v int32) *DescribeSlowLogRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetRegionId(v string) *DescribeSlowLogRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetStartTime(v string) *DescribeSlowLogRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) Validate() error {
	return dara.Validate(s)
}
