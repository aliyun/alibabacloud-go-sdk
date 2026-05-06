// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSampleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListSampleResponseBody
	GetCode() *string
	SetCurrentPage(v int32) *ListSampleResponseBody
	GetCurrentPage() *int32
	SetHttpStatusCode(v string) *ListSampleResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *ListSampleResponseBody
	GetMessage() *string
	SetPageSize(v int32) *ListSampleResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListSampleResponseBody
	GetRequestId() *string
	SetResultObject(v []*ListSampleResponseBodyResultObject) *ListSampleResponseBody
	GetResultObject() []*ListSampleResponseBodyResultObject
	SetTotalItem(v int32) *ListSampleResponseBody
	GetTotalItem() *int32
	SetTotalPage(v int32) *ListSampleResponseBody
	GetTotalPage() *int32
}

type ListSampleResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId    *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject []*ListSampleResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	TotalItem *int32 `json:"TotalItem,omitempty" xml:"TotalItem,omitempty"`
	// example:
	//
	// 1
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListSampleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSampleResponseBody) GoString() string {
	return s.String()
}

func (s *ListSampleResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListSampleResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListSampleResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *ListSampleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSampleResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSampleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSampleResponseBody) GetResultObject() []*ListSampleResponseBodyResultObject {
	return s.ResultObject
}

func (s *ListSampleResponseBody) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *ListSampleResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListSampleResponseBody) SetCode(v string) *ListSampleResponseBody {
	s.Code = &v
	return s
}

func (s *ListSampleResponseBody) SetCurrentPage(v int32) *ListSampleResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListSampleResponseBody) SetHttpStatusCode(v string) *ListSampleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListSampleResponseBody) SetMessage(v string) *ListSampleResponseBody {
	s.Message = &v
	return s
}

func (s *ListSampleResponseBody) SetPageSize(v int32) *ListSampleResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSampleResponseBody) SetRequestId(v string) *ListSampleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSampleResponseBody) SetResultObject(v []*ListSampleResponseBodyResultObject) *ListSampleResponseBody {
	s.ResultObject = v
	return s
}

func (s *ListSampleResponseBody) SetTotalItem(v int32) *ListSampleResponseBody {
	s.TotalItem = &v
	return s
}

func (s *ListSampleResponseBody) SetTotalPage(v int32) *ListSampleResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListSampleResponseBody) Validate() error {
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

type ListSampleResponseBodyResultObject struct {
	// example:
	//
	// 202604016426_2_MOB_10W.csv
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// 78
	FileSize *int32 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// example:
	//
	// vpc-gw8hs2m7qiiy4onxnjf7x
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// 284
	RowCount *int32 `json:"RowCount,omitempty" xml:"RowCount,omitempty"`
	// example:
	//
	// 7
	SampleId *int32 `json:"SampleId,omitempty" xml:"SampleId,omitempty"`
	// example:
	//
	// TEst
	SampleName *string `json:"SampleName,omitempty" xml:"SampleName,omitempty"`
	// example:
	//
	// FINANCE
	Tab *string `json:"Tab,omitempty" xml:"Tab,omitempty"`
	// example:
	//
	// 2023-04-09 12:45:23
	UploadTime *string `json:"UploadTime,omitempty" xml:"UploadTime,omitempty"`
	// example:
	//
	// mest
	UploadUserName *string `json:"UploadUserName,omitempty" xml:"UploadUserName,omitempty"`
}

func (s ListSampleResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s ListSampleResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *ListSampleResponseBodyResultObject) GetFileName() *string {
	return s.FileName
}

func (s *ListSampleResponseBodyResultObject) GetFileSize() *int32 {
	return s.FileSize
}

func (s *ListSampleResponseBodyResultObject) GetRemark() *string {
	return s.Remark
}

func (s *ListSampleResponseBodyResultObject) GetRowCount() *int32 {
	return s.RowCount
}

func (s *ListSampleResponseBodyResultObject) GetSampleId() *int32 {
	return s.SampleId
}

func (s *ListSampleResponseBodyResultObject) GetSampleName() *string {
	return s.SampleName
}

func (s *ListSampleResponseBodyResultObject) GetTab() *string {
	return s.Tab
}

func (s *ListSampleResponseBodyResultObject) GetUploadTime() *string {
	return s.UploadTime
}

func (s *ListSampleResponseBodyResultObject) GetUploadUserName() *string {
	return s.UploadUserName
}

func (s *ListSampleResponseBodyResultObject) SetFileName(v string) *ListSampleResponseBodyResultObject {
	s.FileName = &v
	return s
}

func (s *ListSampleResponseBodyResultObject) SetFileSize(v int32) *ListSampleResponseBodyResultObject {
	s.FileSize = &v
	return s
}

func (s *ListSampleResponseBodyResultObject) SetRemark(v string) *ListSampleResponseBodyResultObject {
	s.Remark = &v
	return s
}

func (s *ListSampleResponseBodyResultObject) SetRowCount(v int32) *ListSampleResponseBodyResultObject {
	s.RowCount = &v
	return s
}

func (s *ListSampleResponseBodyResultObject) SetSampleId(v int32) *ListSampleResponseBodyResultObject {
	s.SampleId = &v
	return s
}

func (s *ListSampleResponseBodyResultObject) SetSampleName(v string) *ListSampleResponseBodyResultObject {
	s.SampleName = &v
	return s
}

func (s *ListSampleResponseBodyResultObject) SetTab(v string) *ListSampleResponseBodyResultObject {
	s.Tab = &v
	return s
}

func (s *ListSampleResponseBodyResultObject) SetUploadTime(v string) *ListSampleResponseBodyResultObject {
	s.UploadTime = &v
	return s
}

func (s *ListSampleResponseBodyResultObject) SetUploadUserName(v string) *ListSampleResponseBodyResultObject {
	s.UploadUserName = &v
	return s
}

func (s *ListSampleResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
