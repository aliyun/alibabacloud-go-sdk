// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadForeignSampleFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UploadForeignSampleFileResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *UploadForeignSampleFileResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *UploadForeignSampleFileResponseBody
	GetMessage() *string
	SetRequestId(v string) *UploadForeignSampleFileResponseBody
	GetRequestId() *string
	SetResultObject(v *UploadForeignSampleFileResponseBodyResultObject) *UploadForeignSampleFileResponseBody
	GetResultObject() *UploadForeignSampleFileResponseBodyResultObject
}

type UploadForeignSampleFileResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId    *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *UploadForeignSampleFileResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s UploadForeignSampleFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadForeignSampleFileResponseBody) GoString() string {
	return s.String()
}

func (s *UploadForeignSampleFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *UploadForeignSampleFileResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *UploadForeignSampleFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UploadForeignSampleFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadForeignSampleFileResponseBody) GetResultObject() *UploadForeignSampleFileResponseBodyResultObject {
	return s.ResultObject
}

func (s *UploadForeignSampleFileResponseBody) SetCode(v string) *UploadForeignSampleFileResponseBody {
	s.Code = &v
	return s
}

func (s *UploadForeignSampleFileResponseBody) SetHttpStatusCode(v string) *UploadForeignSampleFileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UploadForeignSampleFileResponseBody) SetMessage(v string) *UploadForeignSampleFileResponseBody {
	s.Message = &v
	return s
}

func (s *UploadForeignSampleFileResponseBody) SetRequestId(v string) *UploadForeignSampleFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadForeignSampleFileResponseBody) SetResultObject(v *UploadForeignSampleFileResponseBodyResultObject) *UploadForeignSampleFileResponseBody {
	s.ResultObject = v
	return s
}

func (s *UploadForeignSampleFileResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UploadForeignSampleFileResponseBodyResultObject struct {
	// example:
	//
	// 3
	BeyondBacktrackingPeriodNum *int32                                                        `json:"BeyondBacktrackingPeriodNum,omitempty" xml:"BeyondBacktrackingPeriodNum,omitempty"`
	ColumnStats                 []*UploadForeignSampleFileResponseBodyResultObjectColumnStats `json:"ColumnStats,omitempty" xml:"ColumnStats,omitempty" type:"Repeated"`
	// example:
	//
	// yyyyMMdd
	DateType *string `json:"DateType,omitempty" xml:"DateType,omitempty"`
	// example:
	//
	// Acct71b_Sample140934_md5_batch20250916.csv
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// 78
	FileSize *int32 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// example:
	//
	// saf/cpoc/34cd7959590ef568086035b956210495/1760580976089_XN_test_1016_100000.csv
	FileUrl        *string   `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	NotExistScenes []*string `json:"NotExistScenes,omitempty" xml:"NotExistScenes,omitempty" type:"Repeated"`
	// PhoneInvalidList。
	PhoneInvalidList []*string                                                   `json:"PhoneInvalidList,omitempty" xml:"PhoneInvalidList,omitempty" type:"Repeated"`
	PreviewData      *UploadForeignSampleFileResponseBodyResultObjectPreviewData `json:"PreviewData,omitempty" xml:"PreviewData,omitempty" type:"Struct"`
	// example:
	//
	// 284
	RowCount *int32 `json:"RowCount,omitempty" xml:"RowCount,omitempty"`
}

func (s UploadForeignSampleFileResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s UploadForeignSampleFileResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *UploadForeignSampleFileResponseBodyResultObject) GetBeyondBacktrackingPeriodNum() *int32 {
	return s.BeyondBacktrackingPeriodNum
}

func (s *UploadForeignSampleFileResponseBodyResultObject) GetColumnStats() []*UploadForeignSampleFileResponseBodyResultObjectColumnStats {
	return s.ColumnStats
}

func (s *UploadForeignSampleFileResponseBodyResultObject) GetDateType() *string {
	return s.DateType
}

func (s *UploadForeignSampleFileResponseBodyResultObject) GetFileName() *string {
	return s.FileName
}

func (s *UploadForeignSampleFileResponseBodyResultObject) GetFileSize() *int32 {
	return s.FileSize
}

func (s *UploadForeignSampleFileResponseBodyResultObject) GetFileUrl() *string {
	return s.FileUrl
}

func (s *UploadForeignSampleFileResponseBodyResultObject) GetNotExistScenes() []*string {
	return s.NotExistScenes
}

func (s *UploadForeignSampleFileResponseBodyResultObject) GetPhoneInvalidList() []*string {
	return s.PhoneInvalidList
}

func (s *UploadForeignSampleFileResponseBodyResultObject) GetPreviewData() *UploadForeignSampleFileResponseBodyResultObjectPreviewData {
	return s.PreviewData
}

func (s *UploadForeignSampleFileResponseBodyResultObject) GetRowCount() *int32 {
	return s.RowCount
}

func (s *UploadForeignSampleFileResponseBodyResultObject) SetBeyondBacktrackingPeriodNum(v int32) *UploadForeignSampleFileResponseBodyResultObject {
	s.BeyondBacktrackingPeriodNum = &v
	return s
}

func (s *UploadForeignSampleFileResponseBodyResultObject) SetColumnStats(v []*UploadForeignSampleFileResponseBodyResultObjectColumnStats) *UploadForeignSampleFileResponseBodyResultObject {
	s.ColumnStats = v
	return s
}

func (s *UploadForeignSampleFileResponseBodyResultObject) SetDateType(v string) *UploadForeignSampleFileResponseBodyResultObject {
	s.DateType = &v
	return s
}

func (s *UploadForeignSampleFileResponseBodyResultObject) SetFileName(v string) *UploadForeignSampleFileResponseBodyResultObject {
	s.FileName = &v
	return s
}

func (s *UploadForeignSampleFileResponseBodyResultObject) SetFileSize(v int32) *UploadForeignSampleFileResponseBodyResultObject {
	s.FileSize = &v
	return s
}

func (s *UploadForeignSampleFileResponseBodyResultObject) SetFileUrl(v string) *UploadForeignSampleFileResponseBodyResultObject {
	s.FileUrl = &v
	return s
}

func (s *UploadForeignSampleFileResponseBodyResultObject) SetNotExistScenes(v []*string) *UploadForeignSampleFileResponseBodyResultObject {
	s.NotExistScenes = v
	return s
}

func (s *UploadForeignSampleFileResponseBodyResultObject) SetPhoneInvalidList(v []*string) *UploadForeignSampleFileResponseBodyResultObject {
	s.PhoneInvalidList = v
	return s
}

func (s *UploadForeignSampleFileResponseBodyResultObject) SetPreviewData(v *UploadForeignSampleFileResponseBodyResultObjectPreviewData) *UploadForeignSampleFileResponseBodyResultObject {
	s.PreviewData = v
	return s
}

func (s *UploadForeignSampleFileResponseBodyResultObject) SetRowCount(v int32) *UploadForeignSampleFileResponseBodyResultObject {
	s.RowCount = &v
	return s
}

func (s *UploadForeignSampleFileResponseBodyResultObject) Validate() error {
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

type UploadForeignSampleFileResponseBodyResultObjectColumnStats struct {
	// example:
	//
	// 4
	DistinctNumber *int32 `json:"DistinctNumber,omitempty" xml:"DistinctNumber,omitempty"`
	// example:
	//
	// 23.87%
	DistinctRate *string `json:"DistinctRate,omitempty" xml:"DistinctRate,omitempty"`
	// example:
	//
	// repoName
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// example:
	//
	// 3
	MissNumber *int32 `json:"MissNumber,omitempty" xml:"MissNumber,omitempty"`
	// example:
	//
	// 25.32%
	MissRate *string `json:"MissRate,omitempty" xml:"MissRate,omitempty"`
	// example:
	//
	// 2
	RowNumber *int32 `json:"RowNumber,omitempty" xml:"RowNumber,omitempty"`
}

func (s UploadForeignSampleFileResponseBodyResultObjectColumnStats) String() string {
	return dara.Prettify(s)
}

func (s UploadForeignSampleFileResponseBodyResultObjectColumnStats) GoString() string {
	return s.String()
}

func (s *UploadForeignSampleFileResponseBodyResultObjectColumnStats) GetDistinctNumber() *int32 {
	return s.DistinctNumber
}

func (s *UploadForeignSampleFileResponseBodyResultObjectColumnStats) GetDistinctRate() *string {
	return s.DistinctRate
}

func (s *UploadForeignSampleFileResponseBodyResultObjectColumnStats) GetFieldName() *string {
	return s.FieldName
}

func (s *UploadForeignSampleFileResponseBodyResultObjectColumnStats) GetMissNumber() *int32 {
	return s.MissNumber
}

func (s *UploadForeignSampleFileResponseBodyResultObjectColumnStats) GetMissRate() *string {
	return s.MissRate
}

func (s *UploadForeignSampleFileResponseBodyResultObjectColumnStats) GetRowNumber() *int32 {
	return s.RowNumber
}

func (s *UploadForeignSampleFileResponseBodyResultObjectColumnStats) SetDistinctNumber(v int32) *UploadForeignSampleFileResponseBodyResultObjectColumnStats {
	s.DistinctNumber = &v
	return s
}

func (s *UploadForeignSampleFileResponseBodyResultObjectColumnStats) SetDistinctRate(v string) *UploadForeignSampleFileResponseBodyResultObjectColumnStats {
	s.DistinctRate = &v
	return s
}

func (s *UploadForeignSampleFileResponseBodyResultObjectColumnStats) SetFieldName(v string) *UploadForeignSampleFileResponseBodyResultObjectColumnStats {
	s.FieldName = &v
	return s
}

func (s *UploadForeignSampleFileResponseBodyResultObjectColumnStats) SetMissNumber(v int32) *UploadForeignSampleFileResponseBodyResultObjectColumnStats {
	s.MissNumber = &v
	return s
}

func (s *UploadForeignSampleFileResponseBodyResultObjectColumnStats) SetMissRate(v string) *UploadForeignSampleFileResponseBodyResultObjectColumnStats {
	s.MissRate = &v
	return s
}

func (s *UploadForeignSampleFileResponseBodyResultObjectColumnStats) SetRowNumber(v int32) *UploadForeignSampleFileResponseBodyResultObjectColumnStats {
	s.RowNumber = &v
	return s
}

func (s *UploadForeignSampleFileResponseBodyResultObjectColumnStats) Validate() error {
	return dara.Validate(s)
}

type UploadForeignSampleFileResponseBodyResultObjectPreviewData struct {
	Headers []*string   `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Repeated"`
	Rows    [][]*string `json:"Rows,omitempty" xml:"Rows,omitempty" type:"Repeated"`
}

func (s UploadForeignSampleFileResponseBodyResultObjectPreviewData) String() string {
	return dara.Prettify(s)
}

func (s UploadForeignSampleFileResponseBodyResultObjectPreviewData) GoString() string {
	return s.String()
}

func (s *UploadForeignSampleFileResponseBodyResultObjectPreviewData) GetHeaders() []*string {
	return s.Headers
}

func (s *UploadForeignSampleFileResponseBodyResultObjectPreviewData) GetRows() [][]*string {
	return s.Rows
}

func (s *UploadForeignSampleFileResponseBodyResultObjectPreviewData) SetHeaders(v []*string) *UploadForeignSampleFileResponseBodyResultObjectPreviewData {
	s.Headers = v
	return s
}

func (s *UploadForeignSampleFileResponseBodyResultObjectPreviewData) SetRows(v [][]*string) *UploadForeignSampleFileResponseBodyResultObjectPreviewData {
	s.Rows = v
	return s
}

func (s *UploadForeignSampleFileResponseBodyResultObjectPreviewData) Validate() error {
	return dara.Validate(s)
}
