// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataAgentWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListDataAgentWorkspaceResponseBodyData) *ListDataAgentWorkspaceResponseBody
	GetData() *ListDataAgentWorkspaceResponseBodyData
	SetErrorCode(v string) *ListDataAgentWorkspaceResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDataAgentWorkspaceResponseBody
	GetErrorMessage() *string
	SetMaxResults(v int32) *ListDataAgentWorkspaceResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataAgentWorkspaceResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDataAgentWorkspaceResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ListDataAgentWorkspaceResponseBody
	GetSuccess() *string
}

type ListDataAgentWorkspaceResponseBody struct {
	Data *ListDataAgentWorkspaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// DMS-DA-40114
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Specified parameter is not valid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// nu use
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// no use
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// E0D2-*****-A63B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDataAgentWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataAgentWorkspaceResponseBody) GetData() *ListDataAgentWorkspaceResponseBodyData {
	return s.Data
}

func (s *ListDataAgentWorkspaceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDataAgentWorkspaceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDataAgentWorkspaceResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataAgentWorkspaceResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataAgentWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataAgentWorkspaceResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ListDataAgentWorkspaceResponseBody) SetData(v *ListDataAgentWorkspaceResponseBodyData) *ListDataAgentWorkspaceResponseBody {
	s.Data = v
	return s
}

func (s *ListDataAgentWorkspaceResponseBody) SetErrorCode(v string) *ListDataAgentWorkspaceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDataAgentWorkspaceResponseBody) SetErrorMessage(v string) *ListDataAgentWorkspaceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDataAgentWorkspaceResponseBody) SetMaxResults(v int32) *ListDataAgentWorkspaceResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDataAgentWorkspaceResponseBody) SetNextToken(v string) *ListDataAgentWorkspaceResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDataAgentWorkspaceResponseBody) SetRequestId(v string) *ListDataAgentWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataAgentWorkspaceResponseBody) SetSuccess(v string) *ListDataAgentWorkspaceResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataAgentWorkspaceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataAgentWorkspaceResponseBodyData struct {
	Content []*ListDataAgentWorkspaceResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// example:
	//
	// 12
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 50
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 122
	TotalElements *int64 `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
	// example:
	//
	// 12
	TotalPages *int64 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListDataAgentWorkspaceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentWorkspaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataAgentWorkspaceResponseBodyData) GetContent() []*ListDataAgentWorkspaceResponseBodyDataContent {
	return s.Content
}

func (s *ListDataAgentWorkspaceResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListDataAgentWorkspaceResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListDataAgentWorkspaceResponseBodyData) GetTotalElements() *int64 {
	return s.TotalElements
}

func (s *ListDataAgentWorkspaceResponseBodyData) GetTotalPages() *int64 {
	return s.TotalPages
}

func (s *ListDataAgentWorkspaceResponseBodyData) SetContent(v []*ListDataAgentWorkspaceResponseBodyDataContent) *ListDataAgentWorkspaceResponseBodyData {
	s.Content = v
	return s
}

func (s *ListDataAgentWorkspaceResponseBodyData) SetPageNumber(v int64) *ListDataAgentWorkspaceResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListDataAgentWorkspaceResponseBodyData) SetPageSize(v int64) *ListDataAgentWorkspaceResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListDataAgentWorkspaceResponseBodyData) SetTotalElements(v int64) *ListDataAgentWorkspaceResponseBodyData {
	s.TotalElements = &v
	return s
}

func (s *ListDataAgentWorkspaceResponseBodyData) SetTotalPages(v int64) *ListDataAgentWorkspaceResponseBodyData {
	s.TotalPages = &v
	return s
}

func (s *ListDataAgentWorkspaceResponseBodyData) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataAgentWorkspaceResponseBodyDataContent struct {
	// example:
	//
	// 1765960516
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 20282*****7591
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// space for test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1765961516
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// admin
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// example:
	//
	// 11
	TotalMember *int64 `json:"TotalMember,omitempty" xml:"TotalMember,omitempty"`
	// example:
	//
	// 20923*****7291
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// example:
	//
	// workspaceTest
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
	// example:
	//
	// active
	WorkspaceStatus *string `json:"WorkspaceStatus,omitempty" xml:"WorkspaceStatus,omitempty"`
}

func (s ListDataAgentWorkspaceResponseBodyDataContent) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentWorkspaceResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *ListDataAgentWorkspaceResponseBodyDataContent) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListDataAgentWorkspaceResponseBodyDataContent) GetCreator() *string {
	return s.Creator
}

func (s *ListDataAgentWorkspaceResponseBodyDataContent) GetDescription() *string {
	return s.Description
}

func (s *ListDataAgentWorkspaceResponseBodyDataContent) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListDataAgentWorkspaceResponseBodyDataContent) GetRoleName() *string {
	return s.RoleName
}

func (s *ListDataAgentWorkspaceResponseBodyDataContent) GetTotalMember() *int64 {
	return s.TotalMember
}

func (s *ListDataAgentWorkspaceResponseBodyDataContent) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListDataAgentWorkspaceResponseBodyDataContent) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *ListDataAgentWorkspaceResponseBodyDataContent) GetWorkspaceStatus() *string {
	return s.WorkspaceStatus
}

func (s *ListDataAgentWorkspaceResponseBodyDataContent) SetCreateTime(v int64) *ListDataAgentWorkspaceResponseBodyDataContent {
	s.CreateTime = &v
	return s
}

func (s *ListDataAgentWorkspaceResponseBodyDataContent) SetCreator(v string) *ListDataAgentWorkspaceResponseBodyDataContent {
	s.Creator = &v
	return s
}

func (s *ListDataAgentWorkspaceResponseBodyDataContent) SetDescription(v string) *ListDataAgentWorkspaceResponseBodyDataContent {
	s.Description = &v
	return s
}

func (s *ListDataAgentWorkspaceResponseBodyDataContent) SetModifyTime(v int64) *ListDataAgentWorkspaceResponseBodyDataContent {
	s.ModifyTime = &v
	return s
}

func (s *ListDataAgentWorkspaceResponseBodyDataContent) SetRoleName(v string) *ListDataAgentWorkspaceResponseBodyDataContent {
	s.RoleName = &v
	return s
}

func (s *ListDataAgentWorkspaceResponseBodyDataContent) SetTotalMember(v int64) *ListDataAgentWorkspaceResponseBodyDataContent {
	s.TotalMember = &v
	return s
}

func (s *ListDataAgentWorkspaceResponseBodyDataContent) SetWorkspaceId(v string) *ListDataAgentWorkspaceResponseBodyDataContent {
	s.WorkspaceId = &v
	return s
}

func (s *ListDataAgentWorkspaceResponseBodyDataContent) SetWorkspaceName(v string) *ListDataAgentWorkspaceResponseBodyDataContent {
	s.WorkspaceName = &v
	return s
}

func (s *ListDataAgentWorkspaceResponseBodyDataContent) SetWorkspaceStatus(v string) *ListDataAgentWorkspaceResponseBodyDataContent {
	s.WorkspaceStatus = &v
	return s
}

func (s *ListDataAgentWorkspaceResponseBodyDataContent) Validate() error {
	return dara.Validate(s)
}
