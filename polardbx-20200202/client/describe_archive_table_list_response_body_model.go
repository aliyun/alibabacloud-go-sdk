// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeArchiveTableListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeArchiveTableListResponseBodyData) *DescribeArchiveTableListResponseBody
	GetData() *DescribeArchiveTableListResponseBodyData
	SetRequestId(v string) *DescribeArchiveTableListResponseBody
	GetRequestId() *string
}

type DescribeArchiveTableListResponseBody struct {
	Data      *DescribeArchiveTableListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeArchiveTableListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeArchiveTableListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeArchiveTableListResponseBody) GetData() *DescribeArchiveTableListResponseBodyData {
	return s.Data
}

func (s *DescribeArchiveTableListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeArchiveTableListResponseBody) SetData(v *DescribeArchiveTableListResponseBodyData) *DescribeArchiveTableListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeArchiveTableListResponseBody) SetRequestId(v string) *DescribeArchiveTableListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeArchiveTableListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeArchiveTableListResponseBodyData struct {
	PageIndex         *int64                                            `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize          *int64                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PausedCount       *int32                                            `json:"PausedCount,omitempty" xml:"PausedCount,omitempty"`
	RunningCount      *int32                                            `json:"RunningCount,omitempty" xml:"RunningCount,omitempty"`
	SuccessCount      *int32                                            `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	Tables            []*DescribeArchiveTableListResponseBodyDataTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	TobeArchivedConut *int32                                            `json:"TobeArchivedConut,omitempty" xml:"TobeArchivedConut,omitempty"`
	Total             *int64                                            `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeArchiveTableListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeArchiveTableListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeArchiveTableListResponseBodyData) GetPageIndex() *int64 {
	return s.PageIndex
}

func (s *DescribeArchiveTableListResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeArchiveTableListResponseBodyData) GetPausedCount() *int32 {
	return s.PausedCount
}

func (s *DescribeArchiveTableListResponseBodyData) GetRunningCount() *int32 {
	return s.RunningCount
}

func (s *DescribeArchiveTableListResponseBodyData) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *DescribeArchiveTableListResponseBodyData) GetTables() []*DescribeArchiveTableListResponseBodyDataTables {
	return s.Tables
}

func (s *DescribeArchiveTableListResponseBodyData) GetTobeArchivedConut() *int32 {
	return s.TobeArchivedConut
}

func (s *DescribeArchiveTableListResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeArchiveTableListResponseBodyData) SetPageIndex(v int64) *DescribeArchiveTableListResponseBodyData {
	s.PageIndex = &v
	return s
}

func (s *DescribeArchiveTableListResponseBodyData) SetPageSize(v int64) *DescribeArchiveTableListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeArchiveTableListResponseBodyData) SetPausedCount(v int32) *DescribeArchiveTableListResponseBodyData {
	s.PausedCount = &v
	return s
}

func (s *DescribeArchiveTableListResponseBodyData) SetRunningCount(v int32) *DescribeArchiveTableListResponseBodyData {
	s.RunningCount = &v
	return s
}

func (s *DescribeArchiveTableListResponseBodyData) SetSuccessCount(v int32) *DescribeArchiveTableListResponseBodyData {
	s.SuccessCount = &v
	return s
}

func (s *DescribeArchiveTableListResponseBodyData) SetTables(v []*DescribeArchiveTableListResponseBodyDataTables) *DescribeArchiveTableListResponseBodyData {
	s.Tables = v
	return s
}

func (s *DescribeArchiveTableListResponseBodyData) SetTobeArchivedConut(v int32) *DescribeArchiveTableListResponseBodyData {
	s.TobeArchivedConut = &v
	return s
}

func (s *DescribeArchiveTableListResponseBodyData) SetTotal(v int64) *DescribeArchiveTableListResponseBodyData {
	s.Total = &v
	return s
}

func (s *DescribeArchiveTableListResponseBodyData) Validate() error {
	if s.Tables != nil {
		for _, item := range s.Tables {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeArchiveTableListResponseBodyDataTables struct {
	ArchiveStatus          *string  `json:"ArchiveStatus,omitempty" xml:"ArchiveStatus,omitempty"`
	CreatedDate            *int64   `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	FileCount              *int32   `json:"FileCount,omitempty" xml:"FileCount,omitempty"`
	LastSuccessArchiveTime *int64   `json:"LastSuccessArchiveTime,omitempty" xml:"LastSuccessArchiveTime,omitempty"`
	SchemaName             *string  `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	SpaceSize              *float64 `json:"SpaceSize,omitempty" xml:"SpaceSize,omitempty"`
	TableName              *string  `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeArchiveTableListResponseBodyDataTables) String() string {
	return dara.Prettify(s)
}

func (s DescribeArchiveTableListResponseBodyDataTables) GoString() string {
	return s.String()
}

func (s *DescribeArchiveTableListResponseBodyDataTables) GetArchiveStatus() *string {
	return s.ArchiveStatus
}

func (s *DescribeArchiveTableListResponseBodyDataTables) GetCreatedDate() *int64 {
	return s.CreatedDate
}

func (s *DescribeArchiveTableListResponseBodyDataTables) GetFileCount() *int32 {
	return s.FileCount
}

func (s *DescribeArchiveTableListResponseBodyDataTables) GetLastSuccessArchiveTime() *int64 {
	return s.LastSuccessArchiveTime
}

func (s *DescribeArchiveTableListResponseBodyDataTables) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DescribeArchiveTableListResponseBodyDataTables) GetSpaceSize() *float64 {
	return s.SpaceSize
}

func (s *DescribeArchiveTableListResponseBodyDataTables) GetTableName() *string {
	return s.TableName
}

func (s *DescribeArchiveTableListResponseBodyDataTables) SetArchiveStatus(v string) *DescribeArchiveTableListResponseBodyDataTables {
	s.ArchiveStatus = &v
	return s
}

func (s *DescribeArchiveTableListResponseBodyDataTables) SetCreatedDate(v int64) *DescribeArchiveTableListResponseBodyDataTables {
	s.CreatedDate = &v
	return s
}

func (s *DescribeArchiveTableListResponseBodyDataTables) SetFileCount(v int32) *DescribeArchiveTableListResponseBodyDataTables {
	s.FileCount = &v
	return s
}

func (s *DescribeArchiveTableListResponseBodyDataTables) SetLastSuccessArchiveTime(v int64) *DescribeArchiveTableListResponseBodyDataTables {
	s.LastSuccessArchiveTime = &v
	return s
}

func (s *DescribeArchiveTableListResponseBodyDataTables) SetSchemaName(v string) *DescribeArchiveTableListResponseBodyDataTables {
	s.SchemaName = &v
	return s
}

func (s *DescribeArchiveTableListResponseBodyDataTables) SetSpaceSize(v float64) *DescribeArchiveTableListResponseBodyDataTables {
	s.SpaceSize = &v
	return s
}

func (s *DescribeArchiveTableListResponseBodyDataTables) SetTableName(v string) *DescribeArchiveTableListResponseBodyDataTables {
	s.TableName = &v
	return s
}

func (s *DescribeArchiveTableListResponseBodyDataTables) Validate() error {
	return dara.Validate(s)
}
