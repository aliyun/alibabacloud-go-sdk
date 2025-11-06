// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcubeWhitelistsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListWhitelistResult(v *ListMcubeWhitelistsResponseBodyListWhitelistResult) *ListMcubeWhitelistsResponseBody
	GetListWhitelistResult() *ListMcubeWhitelistsResponseBodyListWhitelistResult
	SetRequestId(v string) *ListMcubeWhitelistsResponseBody
	GetRequestId() *string
	SetResultCode(v string) *ListMcubeWhitelistsResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *ListMcubeWhitelistsResponseBody
	GetResultMessage() *string
}

type ListMcubeWhitelistsResponseBody struct {
	ListWhitelistResult *ListMcubeWhitelistsResponseBodyListWhitelistResult `json:"ListWhitelistResult,omitempty" xml:"ListWhitelistResult,omitempty" type:"Struct"`
	RequestId           *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode          *string                                             `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage       *string                                             `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s ListMcubeWhitelistsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeWhitelistsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMcubeWhitelistsResponseBody) GetListWhitelistResult() *ListMcubeWhitelistsResponseBodyListWhitelistResult {
	return s.ListWhitelistResult
}

func (s *ListMcubeWhitelistsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMcubeWhitelistsResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *ListMcubeWhitelistsResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *ListMcubeWhitelistsResponseBody) SetListWhitelistResult(v *ListMcubeWhitelistsResponseBodyListWhitelistResult) *ListMcubeWhitelistsResponseBody {
	s.ListWhitelistResult = v
	return s
}

func (s *ListMcubeWhitelistsResponseBody) SetRequestId(v string) *ListMcubeWhitelistsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMcubeWhitelistsResponseBody) SetResultCode(v string) *ListMcubeWhitelistsResponseBody {
	s.ResultCode = &v
	return s
}

func (s *ListMcubeWhitelistsResponseBody) SetResultMessage(v string) *ListMcubeWhitelistsResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *ListMcubeWhitelistsResponseBody) Validate() error {
	if s.ListWhitelistResult != nil {
		if err := s.ListWhitelistResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMcubeWhitelistsResponseBodyListWhitelistResult struct {
	CurrentPage *int32                                                          `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	HasMore     *bool                                                           `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	PageSize    *int32                                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResultMsg   *string                                                         `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success     *bool                                                           `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount  *int64                                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Whitelists  []*ListMcubeWhitelistsResponseBodyListWhitelistResultWhitelists `json:"Whitelists,omitempty" xml:"Whitelists,omitempty" type:"Repeated"`
}

func (s ListMcubeWhitelistsResponseBodyListWhitelistResult) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeWhitelistsResponseBodyListWhitelistResult) GoString() string {
	return s.String()
}

func (s *ListMcubeWhitelistsResponseBodyListWhitelistResult) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListMcubeWhitelistsResponseBodyListWhitelistResult) GetHasMore() *bool {
	return s.HasMore
}

func (s *ListMcubeWhitelistsResponseBodyListWhitelistResult) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMcubeWhitelistsResponseBodyListWhitelistResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *ListMcubeWhitelistsResponseBodyListWhitelistResult) GetSuccess() *bool {
	return s.Success
}

func (s *ListMcubeWhitelistsResponseBodyListWhitelistResult) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListMcubeWhitelistsResponseBodyListWhitelistResult) GetWhitelists() []*ListMcubeWhitelistsResponseBodyListWhitelistResultWhitelists {
	return s.Whitelists
}

func (s *ListMcubeWhitelistsResponseBodyListWhitelistResult) SetCurrentPage(v int32) *ListMcubeWhitelistsResponseBodyListWhitelistResult {
	s.CurrentPage = &v
	return s
}

func (s *ListMcubeWhitelistsResponseBodyListWhitelistResult) SetHasMore(v bool) *ListMcubeWhitelistsResponseBodyListWhitelistResult {
	s.HasMore = &v
	return s
}

func (s *ListMcubeWhitelistsResponseBodyListWhitelistResult) SetPageSize(v int32) *ListMcubeWhitelistsResponseBodyListWhitelistResult {
	s.PageSize = &v
	return s
}

func (s *ListMcubeWhitelistsResponseBodyListWhitelistResult) SetResultMsg(v string) *ListMcubeWhitelistsResponseBodyListWhitelistResult {
	s.ResultMsg = &v
	return s
}

func (s *ListMcubeWhitelistsResponseBodyListWhitelistResult) SetSuccess(v bool) *ListMcubeWhitelistsResponseBodyListWhitelistResult {
	s.Success = &v
	return s
}

func (s *ListMcubeWhitelistsResponseBodyListWhitelistResult) SetTotalCount(v int64) *ListMcubeWhitelistsResponseBodyListWhitelistResult {
	s.TotalCount = &v
	return s
}

func (s *ListMcubeWhitelistsResponseBodyListWhitelistResult) SetWhitelists(v []*ListMcubeWhitelistsResponseBodyListWhitelistResultWhitelists) *ListMcubeWhitelistsResponseBodyListWhitelistResult {
	s.Whitelists = v
	return s
}

func (s *ListMcubeWhitelistsResponseBodyListWhitelistResult) Validate() error {
	if s.Whitelists != nil {
		for _, item := range s.Whitelists {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMcubeWhitelistsResponseBodyListWhitelistResultWhitelists struct {
	AppCode        *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	GmtCreate      *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified    *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id             *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	WhiteListCount *int64  `json:"WhiteListCount,omitempty" xml:"WhiteListCount,omitempty"`
	WhiteListName  *string `json:"WhiteListName,omitempty" xml:"WhiteListName,omitempty"`
	WhitelistType  *string `json:"WhitelistType,omitempty" xml:"WhitelistType,omitempty"`
}

func (s ListMcubeWhitelistsResponseBodyListWhitelistResultWhitelists) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeWhitelistsResponseBodyListWhitelistResultWhitelists) GoString() string {
	return s.String()
}

func (s *ListMcubeWhitelistsResponseBodyListWhitelistResultWhitelists) GetAppCode() *string {
	return s.AppCode
}

func (s *ListMcubeWhitelistsResponseBodyListWhitelistResultWhitelists) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListMcubeWhitelistsResponseBodyListWhitelistResultWhitelists) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListMcubeWhitelistsResponseBodyListWhitelistResultWhitelists) GetId() *int64 {
	return s.Id
}

func (s *ListMcubeWhitelistsResponseBodyListWhitelistResultWhitelists) GetWhiteListCount() *int64 {
	return s.WhiteListCount
}

func (s *ListMcubeWhitelistsResponseBodyListWhitelistResultWhitelists) GetWhiteListName() *string {
	return s.WhiteListName
}

func (s *ListMcubeWhitelistsResponseBodyListWhitelistResultWhitelists) GetWhitelistType() *string {
	return s.WhitelistType
}

func (s *ListMcubeWhitelistsResponseBodyListWhitelistResultWhitelists) SetAppCode(v string) *ListMcubeWhitelistsResponseBodyListWhitelistResultWhitelists {
	s.AppCode = &v
	return s
}

func (s *ListMcubeWhitelistsResponseBodyListWhitelistResultWhitelists) SetGmtCreate(v string) *ListMcubeWhitelistsResponseBodyListWhitelistResultWhitelists {
	s.GmtCreate = &v
	return s
}

func (s *ListMcubeWhitelistsResponseBodyListWhitelistResultWhitelists) SetGmtModified(v string) *ListMcubeWhitelistsResponseBodyListWhitelistResultWhitelists {
	s.GmtModified = &v
	return s
}

func (s *ListMcubeWhitelistsResponseBodyListWhitelistResultWhitelists) SetId(v int64) *ListMcubeWhitelistsResponseBodyListWhitelistResultWhitelists {
	s.Id = &v
	return s
}

func (s *ListMcubeWhitelistsResponseBodyListWhitelistResultWhitelists) SetWhiteListCount(v int64) *ListMcubeWhitelistsResponseBodyListWhitelistResultWhitelists {
	s.WhiteListCount = &v
	return s
}

func (s *ListMcubeWhitelistsResponseBodyListWhitelistResultWhitelists) SetWhiteListName(v string) *ListMcubeWhitelistsResponseBodyListWhitelistResultWhitelists {
	s.WhiteListName = &v
	return s
}

func (s *ListMcubeWhitelistsResponseBodyListWhitelistResultWhitelists) SetWhitelistType(v string) *ListMcubeWhitelistsResponseBodyListWhitelistResultWhitelists {
	s.WhitelistType = &v
	return s
}

func (s *ListMcubeWhitelistsResponseBodyListWhitelistResultWhitelists) Validate() error {
	return dara.Validate(s)
}
