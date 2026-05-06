// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSampleDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSampleDetailResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *GetSampleDetailResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *GetSampleDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSampleDetailResponseBody
	GetRequestId() *string
	SetResultObject(v *GetSampleDetailResponseBodyResultObject) *GetSampleDetailResponseBody
	GetResultObject() *GetSampleDetailResponseBodyResultObject
}

type GetSampleDetailResponseBody struct {
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
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *GetSampleDetailResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s GetSampleDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSampleDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetSampleDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSampleDetailResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *GetSampleDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSampleDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSampleDetailResponseBody) GetResultObject() *GetSampleDetailResponseBodyResultObject {
	return s.ResultObject
}

func (s *GetSampleDetailResponseBody) SetCode(v string) *GetSampleDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetSampleDetailResponseBody) SetHttpStatusCode(v string) *GetSampleDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetSampleDetailResponseBody) SetMessage(v string) *GetSampleDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetSampleDetailResponseBody) SetRequestId(v string) *GetSampleDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSampleDetailResponseBody) SetResultObject(v *GetSampleDetailResponseBodyResultObject) *GetSampleDetailResponseBody {
	s.ResultObject = v
	return s
}

func (s *GetSampleDetailResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSampleDetailResponseBodyResultObject struct {
	ColumnStats []*GetSampleDetailResponseBodyResultObjectColumnStats `json:"ColumnStats,omitempty" xml:"ColumnStats,omitempty" type:"Repeated"`
	// example:
	//
	// month
	DateType *string `json:"DateType,omitempty" xml:"DateType,omitempty"`
	// example:
	//
	// https://cas-documents-service.oss-cn-shanghai.aliyuncs.com/Batch_Upload_Monitor_Domain.xlsx?Expires=1753755419&OSSAccessKeyId=****&Signature=****
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	FileName    *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// 199
	FileSize    *int32                                              `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	PreviewData *GetSampleDetailResponseBodyResultObjectPreviewData `json:"PreviewData,omitempty" xml:"PreviewData,omitempty" type:"Struct"`
	// example:
	//
	// ios_fb
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// 325
	RowCount *int32 `json:"RowCount,omitempty" xml:"RowCount,omitempty"`
	// example:
	//
	// 1
	SampleId *int32 `json:"SampleId,omitempty" xml:"SampleId,omitempty"`
	// example:
	//
	// SampleTest
	SampleName *string `json:"SampleName,omitempty" xml:"SampleName,omitempty"`
	// example:
	//
	// FINANCE
	Tab *string `json:"Tab,omitempty" xml:"Tab,omitempty"`
	// example:
	//
	// 2024-09-27 10:23:40
	UploadTime     *string `json:"UploadTime,omitempty" xml:"UploadTime,omitempty"`
	UploadUserName *string `json:"UploadUserName,omitempty" xml:"UploadUserName,omitempty"`
}

func (s GetSampleDetailResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s GetSampleDetailResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *GetSampleDetailResponseBodyResultObject) GetColumnStats() []*GetSampleDetailResponseBodyResultObjectColumnStats {
	return s.ColumnStats
}

func (s *GetSampleDetailResponseBodyResultObject) GetDateType() *string {
	return s.DateType
}

func (s *GetSampleDetailResponseBodyResultObject) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *GetSampleDetailResponseBodyResultObject) GetFileName() *string {
	return s.FileName
}

func (s *GetSampleDetailResponseBodyResultObject) GetFileSize() *int32 {
	return s.FileSize
}

func (s *GetSampleDetailResponseBodyResultObject) GetPreviewData() *GetSampleDetailResponseBodyResultObjectPreviewData {
	return s.PreviewData
}

func (s *GetSampleDetailResponseBodyResultObject) GetRemark() *string {
	return s.Remark
}

func (s *GetSampleDetailResponseBodyResultObject) GetRowCount() *int32 {
	return s.RowCount
}

func (s *GetSampleDetailResponseBodyResultObject) GetSampleId() *int32 {
	return s.SampleId
}

func (s *GetSampleDetailResponseBodyResultObject) GetSampleName() *string {
	return s.SampleName
}

func (s *GetSampleDetailResponseBodyResultObject) GetTab() *string {
	return s.Tab
}

func (s *GetSampleDetailResponseBodyResultObject) GetUploadTime() *string {
	return s.UploadTime
}

func (s *GetSampleDetailResponseBodyResultObject) GetUploadUserName() *string {
	return s.UploadUserName
}

func (s *GetSampleDetailResponseBodyResultObject) SetColumnStats(v []*GetSampleDetailResponseBodyResultObjectColumnStats) *GetSampleDetailResponseBodyResultObject {
	s.ColumnStats = v
	return s
}

func (s *GetSampleDetailResponseBodyResultObject) SetDateType(v string) *GetSampleDetailResponseBodyResultObject {
	s.DateType = &v
	return s
}

func (s *GetSampleDetailResponseBodyResultObject) SetDownloadUrl(v string) *GetSampleDetailResponseBodyResultObject {
	s.DownloadUrl = &v
	return s
}

func (s *GetSampleDetailResponseBodyResultObject) SetFileName(v string) *GetSampleDetailResponseBodyResultObject {
	s.FileName = &v
	return s
}

func (s *GetSampleDetailResponseBodyResultObject) SetFileSize(v int32) *GetSampleDetailResponseBodyResultObject {
	s.FileSize = &v
	return s
}

func (s *GetSampleDetailResponseBodyResultObject) SetPreviewData(v *GetSampleDetailResponseBodyResultObjectPreviewData) *GetSampleDetailResponseBodyResultObject {
	s.PreviewData = v
	return s
}

func (s *GetSampleDetailResponseBodyResultObject) SetRemark(v string) *GetSampleDetailResponseBodyResultObject {
	s.Remark = &v
	return s
}

func (s *GetSampleDetailResponseBodyResultObject) SetRowCount(v int32) *GetSampleDetailResponseBodyResultObject {
	s.RowCount = &v
	return s
}

func (s *GetSampleDetailResponseBodyResultObject) SetSampleId(v int32) *GetSampleDetailResponseBodyResultObject {
	s.SampleId = &v
	return s
}

func (s *GetSampleDetailResponseBodyResultObject) SetSampleName(v string) *GetSampleDetailResponseBodyResultObject {
	s.SampleName = &v
	return s
}

func (s *GetSampleDetailResponseBodyResultObject) SetTab(v string) *GetSampleDetailResponseBodyResultObject {
	s.Tab = &v
	return s
}

func (s *GetSampleDetailResponseBodyResultObject) SetUploadTime(v string) *GetSampleDetailResponseBodyResultObject {
	s.UploadTime = &v
	return s
}

func (s *GetSampleDetailResponseBodyResultObject) SetUploadUserName(v string) *GetSampleDetailResponseBodyResultObject {
	s.UploadUserName = &v
	return s
}

func (s *GetSampleDetailResponseBodyResultObject) Validate() error {
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

type GetSampleDetailResponseBodyResultObjectColumnStats struct {
	// example:
	//
	// 23
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

func (s GetSampleDetailResponseBodyResultObjectColumnStats) String() string {
	return dara.Prettify(s)
}

func (s GetSampleDetailResponseBodyResultObjectColumnStats) GoString() string {
	return s.String()
}

func (s *GetSampleDetailResponseBodyResultObjectColumnStats) GetDistinctNumber() *int32 {
	return s.DistinctNumber
}

func (s *GetSampleDetailResponseBodyResultObjectColumnStats) GetDistinctRate() *string {
	return s.DistinctRate
}

func (s *GetSampleDetailResponseBodyResultObjectColumnStats) GetFieldName() *string {
	return s.FieldName
}

func (s *GetSampleDetailResponseBodyResultObjectColumnStats) GetMissNumber() *int32 {
	return s.MissNumber
}

func (s *GetSampleDetailResponseBodyResultObjectColumnStats) GetMissRate() *string {
	return s.MissRate
}

func (s *GetSampleDetailResponseBodyResultObjectColumnStats) GetRowNumber() *int32 {
	return s.RowNumber
}

func (s *GetSampleDetailResponseBodyResultObjectColumnStats) SetDistinctNumber(v int32) *GetSampleDetailResponseBodyResultObjectColumnStats {
	s.DistinctNumber = &v
	return s
}

func (s *GetSampleDetailResponseBodyResultObjectColumnStats) SetDistinctRate(v string) *GetSampleDetailResponseBodyResultObjectColumnStats {
	s.DistinctRate = &v
	return s
}

func (s *GetSampleDetailResponseBodyResultObjectColumnStats) SetFieldName(v string) *GetSampleDetailResponseBodyResultObjectColumnStats {
	s.FieldName = &v
	return s
}

func (s *GetSampleDetailResponseBodyResultObjectColumnStats) SetMissNumber(v int32) *GetSampleDetailResponseBodyResultObjectColumnStats {
	s.MissNumber = &v
	return s
}

func (s *GetSampleDetailResponseBodyResultObjectColumnStats) SetMissRate(v string) *GetSampleDetailResponseBodyResultObjectColumnStats {
	s.MissRate = &v
	return s
}

func (s *GetSampleDetailResponseBodyResultObjectColumnStats) SetRowNumber(v int32) *GetSampleDetailResponseBodyResultObjectColumnStats {
	s.RowNumber = &v
	return s
}

func (s *GetSampleDetailResponseBodyResultObjectColumnStats) Validate() error {
	return dara.Validate(s)
}

type GetSampleDetailResponseBodyResultObjectPreviewData struct {
	Headers []*string   `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Repeated"`
	Rows    [][]*string `json:"Rows,omitempty" xml:"Rows,omitempty" type:"Repeated"`
}

func (s GetSampleDetailResponseBodyResultObjectPreviewData) String() string {
	return dara.Prettify(s)
}

func (s GetSampleDetailResponseBodyResultObjectPreviewData) GoString() string {
	return s.String()
}

func (s *GetSampleDetailResponseBodyResultObjectPreviewData) GetHeaders() []*string {
	return s.Headers
}

func (s *GetSampleDetailResponseBodyResultObjectPreviewData) GetRows() [][]*string {
	return s.Rows
}

func (s *GetSampleDetailResponseBodyResultObjectPreviewData) SetHeaders(v []*string) *GetSampleDetailResponseBodyResultObjectPreviewData {
	s.Headers = v
	return s
}

func (s *GetSampleDetailResponseBodyResultObjectPreviewData) SetRows(v [][]*string) *GetSampleDetailResponseBodyResultObjectPreviewData {
	s.Rows = v
	return s
}

func (s *GetSampleDetailResponseBodyResultObjectPreviewData) Validate() error {
	return dara.Validate(s)
}
