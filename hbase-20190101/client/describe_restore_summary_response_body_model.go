// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRestoreSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHasMoreRestoreRecord(v int32) *DescribeRestoreSummaryResponseBody
	GetHasMoreRestoreRecord() *int32
	SetPageNumber(v int32) *DescribeRestoreSummaryResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRestoreSummaryResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeRestoreSummaryResponseBody
	GetRequestId() *string
	SetRescords(v *DescribeRestoreSummaryResponseBodyRescords) *DescribeRestoreSummaryResponseBody
	GetRescords() *DescribeRestoreSummaryResponseBodyRescords
	SetTotal(v int32) *DescribeRestoreSummaryResponseBody
	GetTotal() *int32
}

type DescribeRestoreSummaryResponseBody struct {
	// example:
	//
	// 0
	HasMoreRestoreRecord *int32 `json:"HasMoreRestoreRecord,omitempty" xml:"HasMoreRestoreRecord,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// AE639ED7-F0F3-4A71-911E-CF8EC088816E
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rescords  *DescribeRestoreSummaryResponseBodyRescords `json:"Rescords,omitempty" xml:"Rescords,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeRestoreSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRestoreSummaryResponseBody) GetHasMoreRestoreRecord() *int32 {
	return s.HasMoreRestoreRecord
}

func (s *DescribeRestoreSummaryResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRestoreSummaryResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRestoreSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRestoreSummaryResponseBody) GetRescords() *DescribeRestoreSummaryResponseBodyRescords {
	return s.Rescords
}

func (s *DescribeRestoreSummaryResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeRestoreSummaryResponseBody) SetHasMoreRestoreRecord(v int32) *DescribeRestoreSummaryResponseBody {
	s.HasMoreRestoreRecord = &v
	return s
}

func (s *DescribeRestoreSummaryResponseBody) SetPageNumber(v int32) *DescribeRestoreSummaryResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRestoreSummaryResponseBody) SetPageSize(v int32) *DescribeRestoreSummaryResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRestoreSummaryResponseBody) SetRequestId(v string) *DescribeRestoreSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRestoreSummaryResponseBody) SetRescords(v *DescribeRestoreSummaryResponseBodyRescords) *DescribeRestoreSummaryResponseBody {
	s.Rescords = v
	return s
}

func (s *DescribeRestoreSummaryResponseBody) SetTotal(v int32) *DescribeRestoreSummaryResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeRestoreSummaryResponseBody) Validate() error {
	if s.Rescords != nil {
		if err := s.Rescords.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRestoreSummaryResponseBodyRescords struct {
	Rescord []*DescribeRestoreSummaryResponseBodyRescordsRescord `json:"Rescord,omitempty" xml:"Rescord,omitempty" type:"Repeated"`
}

func (s DescribeRestoreSummaryResponseBodyRescords) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreSummaryResponseBodyRescords) GoString() string {
	return s.String()
}

func (s *DescribeRestoreSummaryResponseBodyRescords) GetRescord() []*DescribeRestoreSummaryResponseBodyRescordsRescord {
	return s.Rescord
}

func (s *DescribeRestoreSummaryResponseBodyRescords) SetRescord(v []*DescribeRestoreSummaryResponseBodyRescordsRescord) *DescribeRestoreSummaryResponseBodyRescords {
	s.Rescord = v
	return s
}

func (s *DescribeRestoreSummaryResponseBodyRescords) Validate() error {
	if s.Rescord != nil {
		for _, item := range s.Rescord {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRestoreSummaryResponseBodyRescordsRescord struct {
	// example:
	//
	// 1/1
	BulkLoadProcess *string `json:"BulkLoadProcess,omitempty" xml:"BulkLoadProcess,omitempty"`
	// example:
	//
	// 2020-11-05T06:45:14Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2020-11-05T06:45:51Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// example:
	//
	// 1/1
	HfileRestoreProcess *string `json:"HfileRestoreProcess,omitempty" xml:"HfileRestoreProcess,omitempty"`
	// example:
	//
	// 0/0
	LogProcess *string `json:"LogProcess,omitempty" xml:"LogProcess,omitempty"`
	// example:
	//
	// 20201105144514
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// example:
	//
	// 1/1
	SchemaProcess *string `json:"SchemaProcess,omitempty" xml:"SchemaProcess,omitempty"`
	// example:
	//
	// SUCCEEDED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeRestoreSummaryResponseBodyRescordsRescord) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreSummaryResponseBodyRescordsRescord) GoString() string {
	return s.String()
}

func (s *DescribeRestoreSummaryResponseBodyRescordsRescord) GetBulkLoadProcess() *string {
	return s.BulkLoadProcess
}

func (s *DescribeRestoreSummaryResponseBodyRescordsRescord) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeRestoreSummaryResponseBodyRescordsRescord) GetFinishTime() *string {
	return s.FinishTime
}

func (s *DescribeRestoreSummaryResponseBodyRescordsRescord) GetHfileRestoreProcess() *string {
	return s.HfileRestoreProcess
}

func (s *DescribeRestoreSummaryResponseBodyRescordsRescord) GetLogProcess() *string {
	return s.LogProcess
}

func (s *DescribeRestoreSummaryResponseBodyRescordsRescord) GetRecordId() *string {
	return s.RecordId
}

func (s *DescribeRestoreSummaryResponseBodyRescordsRescord) GetSchemaProcess() *string {
	return s.SchemaProcess
}

func (s *DescribeRestoreSummaryResponseBodyRescordsRescord) GetStatus() *string {
	return s.Status
}

func (s *DescribeRestoreSummaryResponseBodyRescordsRescord) SetBulkLoadProcess(v string) *DescribeRestoreSummaryResponseBodyRescordsRescord {
	s.BulkLoadProcess = &v
	return s
}

func (s *DescribeRestoreSummaryResponseBodyRescordsRescord) SetCreateTime(v string) *DescribeRestoreSummaryResponseBodyRescordsRescord {
	s.CreateTime = &v
	return s
}

func (s *DescribeRestoreSummaryResponseBodyRescordsRescord) SetFinishTime(v string) *DescribeRestoreSummaryResponseBodyRescordsRescord {
	s.FinishTime = &v
	return s
}

func (s *DescribeRestoreSummaryResponseBodyRescordsRescord) SetHfileRestoreProcess(v string) *DescribeRestoreSummaryResponseBodyRescordsRescord {
	s.HfileRestoreProcess = &v
	return s
}

func (s *DescribeRestoreSummaryResponseBodyRescordsRescord) SetLogProcess(v string) *DescribeRestoreSummaryResponseBodyRescordsRescord {
	s.LogProcess = &v
	return s
}

func (s *DescribeRestoreSummaryResponseBodyRescordsRescord) SetRecordId(v string) *DescribeRestoreSummaryResponseBodyRescordsRescord {
	s.RecordId = &v
	return s
}

func (s *DescribeRestoreSummaryResponseBodyRescordsRescord) SetSchemaProcess(v string) *DescribeRestoreSummaryResponseBodyRescordsRescord {
	s.SchemaProcess = &v
	return s
}

func (s *DescribeRestoreSummaryResponseBodyRescordsRescord) SetStatus(v string) *DescribeRestoreSummaryResponseBodyRescordsRescord {
	s.Status = &v
	return s
}

func (s *DescribeRestoreSummaryResponseBodyRescordsRescord) Validate() error {
	return dara.Validate(s)
}
