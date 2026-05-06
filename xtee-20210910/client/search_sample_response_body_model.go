// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchSampleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SearchSampleResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *SearchSampleResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *SearchSampleResponseBody
	GetMessage() *string
	SetRequestId(v string) *SearchSampleResponseBody
	GetRequestId() *string
	SetResultObject(v []*SearchSampleResponseBodyResultObject) *SearchSampleResponseBody
	GetResultObject() []*SearchSampleResponseBodyResultObject
}

type SearchSampleResponseBody struct {
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
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject []*SearchSampleResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Repeated"`
}

func (s SearchSampleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchSampleResponseBody) GoString() string {
	return s.String()
}

func (s *SearchSampleResponseBody) GetCode() *string {
	return s.Code
}

func (s *SearchSampleResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *SearchSampleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SearchSampleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchSampleResponseBody) GetResultObject() []*SearchSampleResponseBodyResultObject {
	return s.ResultObject
}

func (s *SearchSampleResponseBody) SetCode(v string) *SearchSampleResponseBody {
	s.Code = &v
	return s
}

func (s *SearchSampleResponseBody) SetHttpStatusCode(v string) *SearchSampleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SearchSampleResponseBody) SetMessage(v string) *SearchSampleResponseBody {
	s.Message = &v
	return s
}

func (s *SearchSampleResponseBody) SetRequestId(v string) *SearchSampleResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchSampleResponseBody) SetResultObject(v []*SearchSampleResponseBodyResultObject) *SearchSampleResponseBody {
	s.ResultObject = v
	return s
}

func (s *SearchSampleResponseBody) Validate() error {
	if s.ResultObject != nil {
		for _, item := range s.ResultObject {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchSampleResponseBodyResultObject struct {
	// example:
	//
	// icekredit_model_A_2025c_1755826424_870000.csv
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// 572
	FileSize *int32 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// example:
	//
	// ios_velo
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// 325
	RowCount *int32 `json:"RowCount,omitempty" xml:"RowCount,omitempty"`
	// example:
	//
	// 9b020e69bbae49d88c07a377c3ab7a71
	SampleId *int32 `json:"SampleId,omitempty" xml:"SampleId,omitempty"`
	// example:
	//
	// Test
	SampleName *string `json:"SampleName,omitempty" xml:"SampleName,omitempty"`
	// example:
	//
	// FINANCE
	Tab *string `json:"Tab,omitempty" xml:"Tab,omitempty"`
	// example:
	//
	// 2023-12-01 12:23:34
	UploadTime *string `json:"UploadTime,omitempty" xml:"UploadTime,omitempty"`
	// example:
	//
	// ds
	UploadUserName *string `json:"UploadUserName,omitempty" xml:"UploadUserName,omitempty"`
}

func (s SearchSampleResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s SearchSampleResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *SearchSampleResponseBodyResultObject) GetFileName() *string {
	return s.FileName
}

func (s *SearchSampleResponseBodyResultObject) GetFileSize() *int32 {
	return s.FileSize
}

func (s *SearchSampleResponseBodyResultObject) GetRemark() *string {
	return s.Remark
}

func (s *SearchSampleResponseBodyResultObject) GetRowCount() *int32 {
	return s.RowCount
}

func (s *SearchSampleResponseBodyResultObject) GetSampleId() *int32 {
	return s.SampleId
}

func (s *SearchSampleResponseBodyResultObject) GetSampleName() *string {
	return s.SampleName
}

func (s *SearchSampleResponseBodyResultObject) GetTab() *string {
	return s.Tab
}

func (s *SearchSampleResponseBodyResultObject) GetUploadTime() *string {
	return s.UploadTime
}

func (s *SearchSampleResponseBodyResultObject) GetUploadUserName() *string {
	return s.UploadUserName
}

func (s *SearchSampleResponseBodyResultObject) SetFileName(v string) *SearchSampleResponseBodyResultObject {
	s.FileName = &v
	return s
}

func (s *SearchSampleResponseBodyResultObject) SetFileSize(v int32) *SearchSampleResponseBodyResultObject {
	s.FileSize = &v
	return s
}

func (s *SearchSampleResponseBodyResultObject) SetRemark(v string) *SearchSampleResponseBodyResultObject {
	s.Remark = &v
	return s
}

func (s *SearchSampleResponseBodyResultObject) SetRowCount(v int32) *SearchSampleResponseBodyResultObject {
	s.RowCount = &v
	return s
}

func (s *SearchSampleResponseBodyResultObject) SetSampleId(v int32) *SearchSampleResponseBodyResultObject {
	s.SampleId = &v
	return s
}

func (s *SearchSampleResponseBodyResultObject) SetSampleName(v string) *SearchSampleResponseBodyResultObject {
	s.SampleName = &v
	return s
}

func (s *SearchSampleResponseBodyResultObject) SetTab(v string) *SearchSampleResponseBodyResultObject {
	s.Tab = &v
	return s
}

func (s *SearchSampleResponseBodyResultObject) SetUploadTime(v string) *SearchSampleResponseBodyResultObject {
	s.UploadTime = &v
	return s
}

func (s *SearchSampleResponseBodyResultObject) SetUploadUserName(v string) *SearchSampleResponseBodyResultObject {
	s.UploadUserName = &v
	return s
}

func (s *SearchSampleResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
