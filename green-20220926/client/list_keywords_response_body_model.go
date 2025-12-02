// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKeywordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListKeywordsResponseBody
	GetCode() *int32
	SetData(v *ListKeywordsResponseBodyData) *ListKeywordsResponseBody
	GetData() *ListKeywordsResponseBodyData
	SetMsg(v string) *ListKeywordsResponseBody
	GetMsg() *string
	SetRequestId(v string) *ListKeywordsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListKeywordsResponseBody
	GetSuccess() *bool
}

type ListKeywordsResponseBody struct {
	// Error code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Returned data.
	Data *ListKeywordsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Further description of the error code.
	//
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// ID assigned by the backend to uniquely identify a request. Can be used for troubleshooting.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Success flag.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListKeywordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListKeywordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListKeywordsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListKeywordsResponseBody) GetData() *ListKeywordsResponseBodyData {
	return s.Data
}

func (s *ListKeywordsResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *ListKeywordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListKeywordsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListKeywordsResponseBody) SetCode(v int32) *ListKeywordsResponseBody {
	s.Code = &v
	return s
}

func (s *ListKeywordsResponseBody) SetData(v *ListKeywordsResponseBodyData) *ListKeywordsResponseBody {
	s.Data = v
	return s
}

func (s *ListKeywordsResponseBody) SetMsg(v string) *ListKeywordsResponseBody {
	s.Msg = &v
	return s
}

func (s *ListKeywordsResponseBody) SetRequestId(v string) *ListKeywordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKeywordsResponseBody) SetSuccess(v bool) *ListKeywordsResponseBody {
	s.Success = &v
	return s
}

func (s *ListKeywordsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListKeywordsResponseBodyData struct {
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Data of the current page.
	Items []*ListKeywordsResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// Page size.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Total count.
	//
	// example:
	//
	// 23
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListKeywordsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListKeywordsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListKeywordsResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListKeywordsResponseBodyData) GetItems() []*ListKeywordsResponseBodyDataItems {
	return s.Items
}

func (s *ListKeywordsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListKeywordsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListKeywordsResponseBodyData) SetCurrentPage(v int32) *ListKeywordsResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *ListKeywordsResponseBodyData) SetItems(v []*ListKeywordsResponseBodyDataItems) *ListKeywordsResponseBodyData {
	s.Items = v
	return s
}

func (s *ListKeywordsResponseBodyData) SetPageSize(v int32) *ListKeywordsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListKeywordsResponseBodyData) SetTotalCount(v int32) *ListKeywordsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListKeywordsResponseBodyData) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListKeywordsResponseBodyDataItems struct {
	// Creation time.
	//
	// example:
	//
	// 2023-06-03 14:43:03
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// Modification time.
	//
	// example:
	//
	// 2023-06-03 14:43:03
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Primary key ID.
	//
	// example:
	//
	// 112
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Keyword library ID.
	//
	// example:
	//
	// custom_xxxx
	KeywordLibId *string `json:"KeywordLibId,omitempty" xml:"KeywordLibId,omitempty"`
	// Keyword data ID.
	//
	// example:
	//
	// 4205334
	KeywordMd5Id *int64 `json:"KeywordMd5Id,omitempty" xml:"KeywordMd5Id,omitempty"`
	// Keyword.
	//
	// example:
	//
	// 测试词
	Word *string `json:"Word,omitempty" xml:"Word,omitempty"`
}

func (s ListKeywordsResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s ListKeywordsResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListKeywordsResponseBodyDataItems) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListKeywordsResponseBodyDataItems) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListKeywordsResponseBodyDataItems) GetId() *int64 {
	return s.Id
}

func (s *ListKeywordsResponseBodyDataItems) GetKeywordLibId() *string {
	return s.KeywordLibId
}

func (s *ListKeywordsResponseBodyDataItems) GetKeywordMd5Id() *int64 {
	return s.KeywordMd5Id
}

func (s *ListKeywordsResponseBodyDataItems) GetWord() *string {
	return s.Word
}

func (s *ListKeywordsResponseBodyDataItems) SetGmtCreate(v string) *ListKeywordsResponseBodyDataItems {
	s.GmtCreate = &v
	return s
}

func (s *ListKeywordsResponseBodyDataItems) SetGmtModified(v string) *ListKeywordsResponseBodyDataItems {
	s.GmtModified = &v
	return s
}

func (s *ListKeywordsResponseBodyDataItems) SetId(v int64) *ListKeywordsResponseBodyDataItems {
	s.Id = &v
	return s
}

func (s *ListKeywordsResponseBodyDataItems) SetKeywordLibId(v string) *ListKeywordsResponseBodyDataItems {
	s.KeywordLibId = &v
	return s
}

func (s *ListKeywordsResponseBodyDataItems) SetKeywordMd5Id(v int64) *ListKeywordsResponseBodyDataItems {
	s.KeywordMd5Id = &v
	return s
}

func (s *ListKeywordsResponseBodyDataItems) SetWord(v string) *ListKeywordsResponseBodyDataItems {
	s.Word = &v
	return s
}

func (s *ListKeywordsResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
