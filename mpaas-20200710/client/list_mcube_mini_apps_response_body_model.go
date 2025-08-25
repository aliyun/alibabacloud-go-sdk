// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcubeMiniAppsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListMiniResult(v *ListMcubeMiniAppsResponseBodyListMiniResult) *ListMcubeMiniAppsResponseBody
	GetListMiniResult() *ListMcubeMiniAppsResponseBodyListMiniResult
	SetRequestId(v string) *ListMcubeMiniAppsResponseBody
	GetRequestId() *string
	SetResultCode(v string) *ListMcubeMiniAppsResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *ListMcubeMiniAppsResponseBody
	GetResultMessage() *string
}

type ListMcubeMiniAppsResponseBody struct {
	ListMiniResult *ListMcubeMiniAppsResponseBodyListMiniResult `json:"ListMiniResult,omitempty" xml:"ListMiniResult,omitempty" type:"Struct"`
	RequestId      *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode     *string                                      `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage  *string                                      `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s ListMcubeMiniAppsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeMiniAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMcubeMiniAppsResponseBody) GetListMiniResult() *ListMcubeMiniAppsResponseBodyListMiniResult {
	return s.ListMiniResult
}

func (s *ListMcubeMiniAppsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMcubeMiniAppsResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *ListMcubeMiniAppsResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *ListMcubeMiniAppsResponseBody) SetListMiniResult(v *ListMcubeMiniAppsResponseBodyListMiniResult) *ListMcubeMiniAppsResponseBody {
	s.ListMiniResult = v
	return s
}

func (s *ListMcubeMiniAppsResponseBody) SetRequestId(v string) *ListMcubeMiniAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMcubeMiniAppsResponseBody) SetResultCode(v string) *ListMcubeMiniAppsResponseBody {
	s.ResultCode = &v
	return s
}

func (s *ListMcubeMiniAppsResponseBody) SetResultMessage(v string) *ListMcubeMiniAppsResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *ListMcubeMiniAppsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListMcubeMiniAppsResponseBodyListMiniResult struct {
	CurrentPage     *int32                                                        `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	HasMore         *bool                                                         `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	MiniProgramList []*ListMcubeMiniAppsResponseBodyListMiniResultMiniProgramList `json:"MiniProgramList,omitempty" xml:"MiniProgramList,omitempty" type:"Repeated"`
	PageSize        *int32                                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResultMsg       *string                                                       `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success         *bool                                                         `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount      *int64                                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMcubeMiniAppsResponseBodyListMiniResult) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeMiniAppsResponseBodyListMiniResult) GoString() string {
	return s.String()
}

func (s *ListMcubeMiniAppsResponseBodyListMiniResult) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListMcubeMiniAppsResponseBodyListMiniResult) GetHasMore() *bool {
	return s.HasMore
}

func (s *ListMcubeMiniAppsResponseBodyListMiniResult) GetMiniProgramList() []*ListMcubeMiniAppsResponseBodyListMiniResultMiniProgramList {
	return s.MiniProgramList
}

func (s *ListMcubeMiniAppsResponseBodyListMiniResult) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMcubeMiniAppsResponseBodyListMiniResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *ListMcubeMiniAppsResponseBodyListMiniResult) GetSuccess() *bool {
	return s.Success
}

func (s *ListMcubeMiniAppsResponseBodyListMiniResult) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListMcubeMiniAppsResponseBodyListMiniResult) SetCurrentPage(v int32) *ListMcubeMiniAppsResponseBodyListMiniResult {
	s.CurrentPage = &v
	return s
}

func (s *ListMcubeMiniAppsResponseBodyListMiniResult) SetHasMore(v bool) *ListMcubeMiniAppsResponseBodyListMiniResult {
	s.HasMore = &v
	return s
}

func (s *ListMcubeMiniAppsResponseBodyListMiniResult) SetMiniProgramList(v []*ListMcubeMiniAppsResponseBodyListMiniResultMiniProgramList) *ListMcubeMiniAppsResponseBodyListMiniResult {
	s.MiniProgramList = v
	return s
}

func (s *ListMcubeMiniAppsResponseBodyListMiniResult) SetPageSize(v int32) *ListMcubeMiniAppsResponseBodyListMiniResult {
	s.PageSize = &v
	return s
}

func (s *ListMcubeMiniAppsResponseBodyListMiniResult) SetResultMsg(v string) *ListMcubeMiniAppsResponseBodyListMiniResult {
	s.ResultMsg = &v
	return s
}

func (s *ListMcubeMiniAppsResponseBodyListMiniResult) SetSuccess(v bool) *ListMcubeMiniAppsResponseBodyListMiniResult {
	s.Success = &v
	return s
}

func (s *ListMcubeMiniAppsResponseBodyListMiniResult) SetTotalCount(v int64) *ListMcubeMiniAppsResponseBodyListMiniResult {
	s.TotalCount = &v
	return s
}

func (s *ListMcubeMiniAppsResponseBodyListMiniResult) Validate() error {
	return dara.Validate(s)
}

type ListMcubeMiniAppsResponseBodyListMiniResultMiniProgramList struct {
	AppCode     *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	GmtCreate   *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	H5Id        *string `json:"H5Id,omitempty" xml:"H5Id,omitempty"`
	H5Name      *string `json:"H5Name,omitempty" xml:"H5Name,omitempty"`
}

func (s ListMcubeMiniAppsResponseBodyListMiniResultMiniProgramList) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeMiniAppsResponseBodyListMiniResultMiniProgramList) GoString() string {
	return s.String()
}

func (s *ListMcubeMiniAppsResponseBodyListMiniResultMiniProgramList) GetAppCode() *string {
	return s.AppCode
}

func (s *ListMcubeMiniAppsResponseBodyListMiniResultMiniProgramList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListMcubeMiniAppsResponseBodyListMiniResultMiniProgramList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListMcubeMiniAppsResponseBodyListMiniResultMiniProgramList) GetH5Id() *string {
	return s.H5Id
}

func (s *ListMcubeMiniAppsResponseBodyListMiniResultMiniProgramList) GetH5Name() *string {
	return s.H5Name
}

func (s *ListMcubeMiniAppsResponseBodyListMiniResultMiniProgramList) SetAppCode(v string) *ListMcubeMiniAppsResponseBodyListMiniResultMiniProgramList {
	s.AppCode = &v
	return s
}

func (s *ListMcubeMiniAppsResponseBodyListMiniResultMiniProgramList) SetGmtCreate(v string) *ListMcubeMiniAppsResponseBodyListMiniResultMiniProgramList {
	s.GmtCreate = &v
	return s
}

func (s *ListMcubeMiniAppsResponseBodyListMiniResultMiniProgramList) SetGmtModified(v string) *ListMcubeMiniAppsResponseBodyListMiniResultMiniProgramList {
	s.GmtModified = &v
	return s
}

func (s *ListMcubeMiniAppsResponseBodyListMiniResultMiniProgramList) SetH5Id(v string) *ListMcubeMiniAppsResponseBodyListMiniResultMiniProgramList {
	s.H5Id = &v
	return s
}

func (s *ListMcubeMiniAppsResponseBodyListMiniResultMiniProgramList) SetH5Name(v string) *ListMcubeMiniAppsResponseBodyListMiniResultMiniProgramList {
	s.H5Name = &v
	return s
}

func (s *ListMcubeMiniAppsResponseBodyListMiniResultMiniProgramList) Validate() error {
	return dara.Validate(s)
}
