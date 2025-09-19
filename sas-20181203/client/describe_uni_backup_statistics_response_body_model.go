// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUniBackupStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProtectedDatabaseCount(v int32) *DescribeUniBackupStatisticsResponseBody
	GetProtectedDatabaseCount() *int32
	SetRegionCountList(v []*DescribeUniBackupStatisticsResponseBodyRegionCountList) *DescribeUniBackupStatisticsResponseBody
	GetRegionCountList() []*DescribeUniBackupStatisticsResponseBodyRegionCountList
	SetRequestId(v string) *DescribeUniBackupStatisticsResponseBody
	GetRequestId() *string
	SetRestoringTaskCount(v int32) *DescribeUniBackupStatisticsResponseBody
	GetRestoringTaskCount() *int32
	SetTotalRecoverableCount(v int32) *DescribeUniBackupStatisticsResponseBody
	GetTotalRecoverableCount() *int32
	SetTotalRestoreTaskCount(v int32) *DescribeUniBackupStatisticsResponseBody
	GetTotalRestoreTaskCount() *int32
	SetUnprotectedDatabaseCount(v int32) *DescribeUniBackupStatisticsResponseBody
	GetUnprotectedDatabaseCount() *int32
}

type DescribeUniBackupStatisticsResponseBody struct {
	// The number of protected database instances.
	//
	// example:
	//
	// 1
	ProtectedDatabaseCount *int32 `json:"ProtectedDatabaseCount,omitempty" xml:"ProtectedDatabaseCount,omitempty"`
	// The regions of the database instances.
	RegionCountList []*DescribeUniBackupStatisticsResponseBodyRegionCountList `json:"RegionCountList,omitempty" xml:"RegionCountList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 7532B7EE-7CE7-5F4D-BF04-B12447DD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of the restoration tasks that are running.
	//
	// example:
	//
	// 0
	RestoringTaskCount *int32 `json:"RestoringTaskCount,omitempty" xml:"RestoringTaskCount,omitempty"`
	// The total number of database instances that can be restored.
	//
	// example:
	//
	// 3
	TotalRecoverableCount *int32 `json:"TotalRecoverableCount,omitempty" xml:"TotalRecoverableCount,omitempty"`
	// The total number of the restoration tasks.
	//
	// example:
	//
	// 10
	TotalRestoreTaskCount *int32 `json:"TotalRestoreTaskCount,omitempty" xml:"TotalRestoreTaskCount,omitempty"`
	// The number of unprotected database instances.
	//
	// example:
	//
	// 5
	UnprotectedDatabaseCount *int32 `json:"UnprotectedDatabaseCount,omitempty" xml:"UnprotectedDatabaseCount,omitempty"`
}

func (s DescribeUniBackupStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUniBackupStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUniBackupStatisticsResponseBody) GetProtectedDatabaseCount() *int32 {
	return s.ProtectedDatabaseCount
}

func (s *DescribeUniBackupStatisticsResponseBody) GetRegionCountList() []*DescribeUniBackupStatisticsResponseBodyRegionCountList {
	return s.RegionCountList
}

func (s *DescribeUniBackupStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUniBackupStatisticsResponseBody) GetRestoringTaskCount() *int32 {
	return s.RestoringTaskCount
}

func (s *DescribeUniBackupStatisticsResponseBody) GetTotalRecoverableCount() *int32 {
	return s.TotalRecoverableCount
}

func (s *DescribeUniBackupStatisticsResponseBody) GetTotalRestoreTaskCount() *int32 {
	return s.TotalRestoreTaskCount
}

func (s *DescribeUniBackupStatisticsResponseBody) GetUnprotectedDatabaseCount() *int32 {
	return s.UnprotectedDatabaseCount
}

func (s *DescribeUniBackupStatisticsResponseBody) SetProtectedDatabaseCount(v int32) *DescribeUniBackupStatisticsResponseBody {
	s.ProtectedDatabaseCount = &v
	return s
}

func (s *DescribeUniBackupStatisticsResponseBody) SetRegionCountList(v []*DescribeUniBackupStatisticsResponseBodyRegionCountList) *DescribeUniBackupStatisticsResponseBody {
	s.RegionCountList = v
	return s
}

func (s *DescribeUniBackupStatisticsResponseBody) SetRequestId(v string) *DescribeUniBackupStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUniBackupStatisticsResponseBody) SetRestoringTaskCount(v int32) *DescribeUniBackupStatisticsResponseBody {
	s.RestoringTaskCount = &v
	return s
}

func (s *DescribeUniBackupStatisticsResponseBody) SetTotalRecoverableCount(v int32) *DescribeUniBackupStatisticsResponseBody {
	s.TotalRecoverableCount = &v
	return s
}

func (s *DescribeUniBackupStatisticsResponseBody) SetTotalRestoreTaskCount(v int32) *DescribeUniBackupStatisticsResponseBody {
	s.TotalRestoreTaskCount = &v
	return s
}

func (s *DescribeUniBackupStatisticsResponseBody) SetUnprotectedDatabaseCount(v int32) *DescribeUniBackupStatisticsResponseBody {
	s.UnprotectedDatabaseCount = &v
	return s
}

func (s *DescribeUniBackupStatisticsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeUniBackupStatisticsResponseBodyRegionCountList struct {
	// The number of database instances that are automatically scanned.
	//
	// example:
	//
	// 1
	AutomaticCount *string `json:"AutomaticCount,omitempty" xml:"AutomaticCount,omitempty"`
	// The ID of the region in which the database instance resides.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeUniBackupStatisticsResponseBodyRegionCountList) String() string {
	return dara.Prettify(s)
}

func (s DescribeUniBackupStatisticsResponseBodyRegionCountList) GoString() string {
	return s.String()
}

func (s *DescribeUniBackupStatisticsResponseBodyRegionCountList) GetAutomaticCount() *string {
	return s.AutomaticCount
}

func (s *DescribeUniBackupStatisticsResponseBodyRegionCountList) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeUniBackupStatisticsResponseBodyRegionCountList) SetAutomaticCount(v string) *DescribeUniBackupStatisticsResponseBodyRegionCountList {
	s.AutomaticCount = &v
	return s
}

func (s *DescribeUniBackupStatisticsResponseBodyRegionCountList) SetRegionId(v string) *DescribeUniBackupStatisticsResponseBodyRegionCountList {
	s.RegionId = &v
	return s
}

func (s *DescribeUniBackupStatisticsResponseBodyRegionCountList) Validate() error {
	return dara.Validate(s)
}
