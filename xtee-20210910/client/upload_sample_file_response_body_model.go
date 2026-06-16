// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadSampleFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UploadSampleFileResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *UploadSampleFileResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *UploadSampleFileResponseBody
	GetMessage() *string
	SetRequestId(v string) *UploadSampleFileResponseBody
	GetRequestId() *string
	SetResultObject(v *UploadSampleFileResponseBodyResultObject) *UploadSampleFileResponseBody
	GetResultObject() *UploadSampleFileResponseBodyResultObject
}

type UploadSampleFileResponseBody struct {
	// Status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Return message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return Result.
	ResultObject *UploadSampleFileResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s UploadSampleFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadSampleFileResponseBody) GoString() string {
	return s.String()
}

func (s *UploadSampleFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *UploadSampleFileResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *UploadSampleFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UploadSampleFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadSampleFileResponseBody) GetResultObject() *UploadSampleFileResponseBodyResultObject {
	return s.ResultObject
}

func (s *UploadSampleFileResponseBody) SetCode(v string) *UploadSampleFileResponseBody {
	s.Code = &v
	return s
}

func (s *UploadSampleFileResponseBody) SetHttpStatusCode(v string) *UploadSampleFileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UploadSampleFileResponseBody) SetMessage(v string) *UploadSampleFileResponseBody {
	s.Message = &v
	return s
}

func (s *UploadSampleFileResponseBody) SetRequestId(v string) *UploadSampleFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadSampleFileResponseBody) SetResultObject(v *UploadSampleFileResponseBodyResultObject) *UploadSampleFileResponseBody {
	s.ResultObject = v
	return s
}

func (s *UploadSampleFileResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UploadSampleFileResponseBodyResultObject struct {
	// Number of jobs.
	//
	// example:
	//
	// gd
	BeyondBacktrackingPeriodNum *int32 `json:"BeyondBacktrackingPeriodNum,omitempty" xml:"BeyondBacktrackingPeriodNum,omitempty"`
	// Columns.
	ColumnStats []*UploadSampleFileResponseBodyResultObjectColumnStats `json:"ColumnStats,omitempty" xml:"ColumnStats,omitempty" type:"Repeated"`
	// Time type.
	//
	// example:
	//
	// yyyyMMdd
	DateType *string `json:"DateType,omitempty" xml:"DateType,omitempty"`
	// File name.
	//
	// example:
	//
	// Acct71b_Sample140934_md5_batch20250916.csv
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// File Size.
	//
	// example:
	//
	// 1472
	FileSize *int32 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// File URL.
	//
	// example:
	//
	// saf/cpoc/30e2aff8f355af429bbab18f776496ef/1764037297543_icekredit_model_A_2025b_1764034785_147787.csv
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// Non-existent scenarios.
	NotExistScenes []*string `json:"NotExistScenes,omitempty" xml:"NotExistScenes,omitempty" type:"Repeated"`
	// Number.
	PhoneInvalidList []*string `json:"PhoneInvalidList,omitempty" xml:"PhoneInvalidList,omitempty" type:"Repeated"`
	// Table data.
	PreviewData *UploadSampleFileResponseBodyResultObjectPreviewData `json:"PreviewData,omitempty" xml:"PreviewData,omitempty" type:"Struct"`
	// Number of rows.
	//
	// example:
	//
	// 4974
	RowCount *int32 `json:"RowCount,omitempty" xml:"RowCount,omitempty"`
}

func (s UploadSampleFileResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s UploadSampleFileResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *UploadSampleFileResponseBodyResultObject) GetBeyondBacktrackingPeriodNum() *int32 {
	return s.BeyondBacktrackingPeriodNum
}

func (s *UploadSampleFileResponseBodyResultObject) GetColumnStats() []*UploadSampleFileResponseBodyResultObjectColumnStats {
	return s.ColumnStats
}

func (s *UploadSampleFileResponseBodyResultObject) GetDateType() *string {
	return s.DateType
}

func (s *UploadSampleFileResponseBodyResultObject) GetFileName() *string {
	return s.FileName
}

func (s *UploadSampleFileResponseBodyResultObject) GetFileSize() *int32 {
	return s.FileSize
}

func (s *UploadSampleFileResponseBodyResultObject) GetFileUrl() *string {
	return s.FileUrl
}

func (s *UploadSampleFileResponseBodyResultObject) GetNotExistScenes() []*string {
	return s.NotExistScenes
}

func (s *UploadSampleFileResponseBodyResultObject) GetPhoneInvalidList() []*string {
	return s.PhoneInvalidList
}

func (s *UploadSampleFileResponseBodyResultObject) GetPreviewData() *UploadSampleFileResponseBodyResultObjectPreviewData {
	return s.PreviewData
}

func (s *UploadSampleFileResponseBodyResultObject) GetRowCount() *int32 {
	return s.RowCount
}

func (s *UploadSampleFileResponseBodyResultObject) SetBeyondBacktrackingPeriodNum(v int32) *UploadSampleFileResponseBodyResultObject {
	s.BeyondBacktrackingPeriodNum = &v
	return s
}

func (s *UploadSampleFileResponseBodyResultObject) SetColumnStats(v []*UploadSampleFileResponseBodyResultObjectColumnStats) *UploadSampleFileResponseBodyResultObject {
	s.ColumnStats = v
	return s
}

func (s *UploadSampleFileResponseBodyResultObject) SetDateType(v string) *UploadSampleFileResponseBodyResultObject {
	s.DateType = &v
	return s
}

func (s *UploadSampleFileResponseBodyResultObject) SetFileName(v string) *UploadSampleFileResponseBodyResultObject {
	s.FileName = &v
	return s
}

func (s *UploadSampleFileResponseBodyResultObject) SetFileSize(v int32) *UploadSampleFileResponseBodyResultObject {
	s.FileSize = &v
	return s
}

func (s *UploadSampleFileResponseBodyResultObject) SetFileUrl(v string) *UploadSampleFileResponseBodyResultObject {
	s.FileUrl = &v
	return s
}

func (s *UploadSampleFileResponseBodyResultObject) SetNotExistScenes(v []*string) *UploadSampleFileResponseBodyResultObject {
	s.NotExistScenes = v
	return s
}

func (s *UploadSampleFileResponseBodyResultObject) SetPhoneInvalidList(v []*string) *UploadSampleFileResponseBodyResultObject {
	s.PhoneInvalidList = v
	return s
}

func (s *UploadSampleFileResponseBodyResultObject) SetPreviewData(v *UploadSampleFileResponseBodyResultObjectPreviewData) *UploadSampleFileResponseBodyResultObject {
	s.PreviewData = v
	return s
}

func (s *UploadSampleFileResponseBodyResultObject) SetRowCount(v int32) *UploadSampleFileResponseBodyResultObject {
	s.RowCount = &v
	return s
}

func (s *UploadSampleFileResponseBodyResultObject) Validate() error {
	if s.ColumnStats != nil {
		for _, item := range s.ColumnStats {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PreviewData != nil {
		if err := s.PreviewData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UploadSampleFileResponseBodyResultObjectColumnStats struct {
	// De-duplication Count.
	//
	// example:
	//
	// 4
	DistinctNumber *int32 `json:"DistinctNumber,omitempty" xml:"DistinctNumber,omitempty"`
	// De-duplication rate.
	//
	// example:
	//
	// 28.23%
	DistinctRate *string `json:"DistinctRate,omitempty" xml:"DistinctRate,omitempty"`
	// Field Name.
	//
	// example:
	//
	// fb_org_id
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// Number of missing values.
	//
	// example:
	//
	// 5
	MissNumber *int32 `json:"MissNumber,omitempty" xml:"MissNumber,omitempty"`
	// Missing rate.
	//
	// example:
	//
	// 28.23%
	MissRate *string `json:"MissRate,omitempty" xml:"MissRate,omitempty"`
	// Row number of the record.
	//
	// example:
	//
	// 4
	RowNumber *int32 `json:"RowNumber,omitempty" xml:"RowNumber,omitempty"`
}

func (s UploadSampleFileResponseBodyResultObjectColumnStats) String() string {
	return dara.Prettify(s)
}

func (s UploadSampleFileResponseBodyResultObjectColumnStats) GoString() string {
	return s.String()
}

func (s *UploadSampleFileResponseBodyResultObjectColumnStats) GetDistinctNumber() *int32 {
	return s.DistinctNumber
}

func (s *UploadSampleFileResponseBodyResultObjectColumnStats) GetDistinctRate() *string {
	return s.DistinctRate
}

func (s *UploadSampleFileResponseBodyResultObjectColumnStats) GetFieldName() *string {
	return s.FieldName
}

func (s *UploadSampleFileResponseBodyResultObjectColumnStats) GetMissNumber() *int32 {
	return s.MissNumber
}

func (s *UploadSampleFileResponseBodyResultObjectColumnStats) GetMissRate() *string {
	return s.MissRate
}

func (s *UploadSampleFileResponseBodyResultObjectColumnStats) GetRowNumber() *int32 {
	return s.RowNumber
}

func (s *UploadSampleFileResponseBodyResultObjectColumnStats) SetDistinctNumber(v int32) *UploadSampleFileResponseBodyResultObjectColumnStats {
	s.DistinctNumber = &v
	return s
}

func (s *UploadSampleFileResponseBodyResultObjectColumnStats) SetDistinctRate(v string) *UploadSampleFileResponseBodyResultObjectColumnStats {
	s.DistinctRate = &v
	return s
}

func (s *UploadSampleFileResponseBodyResultObjectColumnStats) SetFieldName(v string) *UploadSampleFileResponseBodyResultObjectColumnStats {
	s.FieldName = &v
	return s
}

func (s *UploadSampleFileResponseBodyResultObjectColumnStats) SetMissNumber(v int32) *UploadSampleFileResponseBodyResultObjectColumnStats {
	s.MissNumber = &v
	return s
}

func (s *UploadSampleFileResponseBodyResultObjectColumnStats) SetMissRate(v string) *UploadSampleFileResponseBodyResultObjectColumnStats {
	s.MissRate = &v
	return s
}

func (s *UploadSampleFileResponseBodyResultObjectColumnStats) SetRowNumber(v int32) *UploadSampleFileResponseBodyResultObjectColumnStats {
	s.RowNumber = &v
	return s
}

func (s *UploadSampleFileResponseBodyResultObjectColumnStats) Validate() error {
	return dara.Validate(s)
}

type UploadSampleFileResponseBodyResultObjectPreviewData struct {
	// Table header.
	Headers []*string `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Repeated"`
	// Row data.
	Rows [][]*string `json:"Rows,omitempty" xml:"Rows,omitempty" type:"Repeated"`
}

func (s UploadSampleFileResponseBodyResultObjectPreviewData) String() string {
	return dara.Prettify(s)
}

func (s UploadSampleFileResponseBodyResultObjectPreviewData) GoString() string {
	return s.String()
}

func (s *UploadSampleFileResponseBodyResultObjectPreviewData) GetHeaders() []*string {
	return s.Headers
}

func (s *UploadSampleFileResponseBodyResultObjectPreviewData) GetRows() [][]*string {
	return s.Rows
}

func (s *UploadSampleFileResponseBodyResultObjectPreviewData) SetHeaders(v []*string) *UploadSampleFileResponseBodyResultObjectPreviewData {
	s.Headers = v
	return s
}

func (s *UploadSampleFileResponseBodyResultObjectPreviewData) SetRows(v [][]*string) *UploadSampleFileResponseBodyResultObjectPreviewData {
	s.Rows = v
	return s
}

func (s *UploadSampleFileResponseBodyResultObjectPreviewData) Validate() error {
	return dara.Validate(s)
}
