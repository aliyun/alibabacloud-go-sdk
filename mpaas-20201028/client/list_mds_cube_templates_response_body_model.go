// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMdsCubeTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListMdsCubeTemplatesResponseBody
	GetRequestId() *string
	SetResultCode(v string) *ListMdsCubeTemplatesResponseBody
	GetResultCode() *string
	SetResultContent(v *ListMdsCubeTemplatesResponseBodyResultContent) *ListMdsCubeTemplatesResponseBody
	GetResultContent() *ListMdsCubeTemplatesResponseBodyResultContent
	SetResultMessage(v string) *ListMdsCubeTemplatesResponseBody
	GetResultMessage() *string
}

type ListMdsCubeTemplatesResponseBody struct {
	// example:
	//
	// 11E66B29-9E5E-5C10-B64E-B5A0E0F26355
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// OK
	ResultCode    *string                                        `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *ListMdsCubeTemplatesResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	// example:
	//
	// success
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s ListMdsCubeTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMdsCubeTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMdsCubeTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMdsCubeTemplatesResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *ListMdsCubeTemplatesResponseBody) GetResultContent() *ListMdsCubeTemplatesResponseBodyResultContent {
	return s.ResultContent
}

func (s *ListMdsCubeTemplatesResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *ListMdsCubeTemplatesResponseBody) SetRequestId(v string) *ListMdsCubeTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMdsCubeTemplatesResponseBody) SetResultCode(v string) *ListMdsCubeTemplatesResponseBody {
	s.ResultCode = &v
	return s
}

func (s *ListMdsCubeTemplatesResponseBody) SetResultContent(v *ListMdsCubeTemplatesResponseBodyResultContent) *ListMdsCubeTemplatesResponseBody {
	s.ResultContent = v
	return s
}

func (s *ListMdsCubeTemplatesResponseBody) SetResultMessage(v string) *ListMdsCubeTemplatesResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *ListMdsCubeTemplatesResponseBody) Validate() error {
	if s.ResultContent != nil {
		if err := s.ResultContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMdsCubeTemplatesResponseBodyResultContent struct {
	Data *ListMdsCubeTemplatesResponseBodyResultContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 11E66B29-9E5E-5C10-B64E-B5A0E0F26355
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListMdsCubeTemplatesResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s ListMdsCubeTemplatesResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *ListMdsCubeTemplatesResponseBodyResultContent) GetData() *ListMdsCubeTemplatesResponseBodyResultContentData {
	return s.Data
}

func (s *ListMdsCubeTemplatesResponseBodyResultContent) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMdsCubeTemplatesResponseBodyResultContent) SetData(v *ListMdsCubeTemplatesResponseBodyResultContentData) *ListMdsCubeTemplatesResponseBodyResultContent {
	s.Data = v
	return s
}

func (s *ListMdsCubeTemplatesResponseBodyResultContent) SetRequestId(v string) *ListMdsCubeTemplatesResponseBodyResultContent {
	s.RequestId = &v
	return s
}

func (s *ListMdsCubeTemplatesResponseBodyResultContent) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMdsCubeTemplatesResponseBodyResultContentData struct {
	Content *ListMdsCubeTemplatesResponseBodyResultContentDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// example:
	//
	// NONE
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 11E66B29-9E5E-5C10-B64E-B5A0E0F26355
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListMdsCubeTemplatesResponseBodyResultContentData) String() string {
	return dara.Prettify(s)
}

func (s ListMdsCubeTemplatesResponseBodyResultContentData) GoString() string {
	return s.String()
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentData) GetContent() *ListMdsCubeTemplatesResponseBodyResultContentDataContent {
	return s.Content
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentData) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentData) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentData) GetSuccess() *bool {
	return s.Success
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentData) SetContent(v *ListMdsCubeTemplatesResponseBodyResultContentDataContent) *ListMdsCubeTemplatesResponseBodyResultContentData {
	s.Content = v
	return s
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentData) SetErrorCode(v string) *ListMdsCubeTemplatesResponseBodyResultContentData {
	s.ErrorCode = &v
	return s
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentData) SetRequestId(v string) *ListMdsCubeTemplatesResponseBodyResultContentData {
	s.RequestId = &v
	return s
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentData) SetResultMsg(v string) *ListMdsCubeTemplatesResponseBodyResultContentData {
	s.ResultMsg = &v
	return s
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentData) SetSuccess(v bool) *ListMdsCubeTemplatesResponseBodyResultContentData {
	s.Success = &v
	return s
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentData) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMdsCubeTemplatesResponseBodyResultContentDataContent struct {
	// example:
	//
	// true
	FirstPage *bool `json:"FirstPage,omitempty" xml:"FirstPage,omitempty"`
	// example:
	//
	// 0
	FirstResult *int64 `json:"FirstResult,omitempty" xml:"FirstResult,omitempty"`
	// example:
	//
	// true
	LastPage *bool                                                           `json:"LastPage,omitempty" xml:"LastPage,omitempty"`
	List     []*ListMdsCubeTemplatesResponseBodyResultContentDataContentList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	NextPage *int64 `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	PrePage *int64 `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 1
	TotalPage *int64 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListMdsCubeTemplatesResponseBodyResultContentDataContent) String() string {
	return dara.Prettify(s)
}

func (s ListMdsCubeTemplatesResponseBodyResultContentDataContent) GoString() string {
	return s.String()
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentDataContent) GetFirstPage() *bool {
	return s.FirstPage
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentDataContent) GetFirstResult() *int64 {
	return s.FirstResult
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentDataContent) GetLastPage() *bool {
	return s.LastPage
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentDataContent) GetList() []*ListMdsCubeTemplatesResponseBodyResultContentDataContentList {
	return s.List
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentDataContent) GetNextPage() *int64 {
	return s.NextPage
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentDataContent) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentDataContent) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentDataContent) GetPrePage() *int64 {
	return s.PrePage
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentDataContent) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentDataContent) GetTotalPage() *int64 {
	return s.TotalPage
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentDataContent) SetFirstPage(v bool) *ListMdsCubeTemplatesResponseBodyResultContentDataContent {
	s.FirstPage = &v
	return s
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentDataContent) SetFirstResult(v int64) *ListMdsCubeTemplatesResponseBodyResultContentDataContent {
	s.FirstResult = &v
	return s
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentDataContent) SetLastPage(v bool) *ListMdsCubeTemplatesResponseBodyResultContentDataContent {
	s.LastPage = &v
	return s
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentDataContent) SetList(v []*ListMdsCubeTemplatesResponseBodyResultContentDataContentList) *ListMdsCubeTemplatesResponseBodyResultContentDataContent {
	s.List = v
	return s
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentDataContent) SetNextPage(v int64) *ListMdsCubeTemplatesResponseBodyResultContentDataContent {
	s.NextPage = &v
	return s
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentDataContent) SetPageNo(v int32) *ListMdsCubeTemplatesResponseBodyResultContentDataContent {
	s.PageNo = &v
	return s
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentDataContent) SetPageSize(v int32) *ListMdsCubeTemplatesResponseBodyResultContentDataContent {
	s.PageSize = &v
	return s
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentDataContent) SetPrePage(v int64) *ListMdsCubeTemplatesResponseBodyResultContentDataContent {
	s.PrePage = &v
	return s
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentDataContent) SetTotalCount(v int32) *ListMdsCubeTemplatesResponseBodyResultContentDataContent {
	s.TotalCount = &v
	return s
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentDataContent) SetTotalPage(v int64) *ListMdsCubeTemplatesResponseBodyResultContentDataContent {
	s.TotalPage = &v
	return s
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentDataContent) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMdsCubeTemplatesResponseBodyResultContentDataContentList struct {
	// example:
	//
	// 570DA89281533-default
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// example:
	//
	// 2025-10-24 11:24:34
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2025-10-24 11:24:34
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// template desc
	TemplateDesc *string `json:"TemplateDesc,omitempty" xml:"TemplateDesc,omitempty"`
	// example:
	//
	// template_test
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// template_test
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ListMdsCubeTemplatesResponseBodyResultContentDataContentList) String() string {
	return dara.Prettify(s)
}

func (s ListMdsCubeTemplatesResponseBodyResultContentDataContentList) GoString() string {
	return s.String()
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentDataContentList) GetAppCode() *string {
	return s.AppCode
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentDataContentList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentDataContentList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentDataContentList) GetId() *int64 {
	return s.Id
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentDataContentList) GetStatus() *int32 {
	return s.Status
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentDataContentList) GetTemplateDesc() *string {
	return s.TemplateDesc
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentDataContentList) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentDataContentList) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentDataContentList) SetAppCode(v string) *ListMdsCubeTemplatesResponseBodyResultContentDataContentList {
	s.AppCode = &v
	return s
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentDataContentList) SetGmtCreate(v string) *ListMdsCubeTemplatesResponseBodyResultContentDataContentList {
	s.GmtCreate = &v
	return s
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentDataContentList) SetGmtModified(v string) *ListMdsCubeTemplatesResponseBodyResultContentDataContentList {
	s.GmtModified = &v
	return s
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentDataContentList) SetId(v int64) *ListMdsCubeTemplatesResponseBodyResultContentDataContentList {
	s.Id = &v
	return s
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentDataContentList) SetStatus(v int32) *ListMdsCubeTemplatesResponseBodyResultContentDataContentList {
	s.Status = &v
	return s
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentDataContentList) SetTemplateDesc(v string) *ListMdsCubeTemplatesResponseBodyResultContentDataContentList {
	s.TemplateDesc = &v
	return s
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentDataContentList) SetTemplateId(v string) *ListMdsCubeTemplatesResponseBodyResultContentDataContentList {
	s.TemplateId = &v
	return s
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentDataContentList) SetTemplateName(v string) *ListMdsCubeTemplatesResponseBodyResultContentDataContentList {
	s.TemplateName = &v
	return s
}

func (s *ListMdsCubeTemplatesResponseBodyResultContentDataContentList) Validate() error {
	return dara.Validate(s)
}
