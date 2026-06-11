// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataAgentWorkspaceMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListDataAgentWorkspaceMemberResponseBodyData) *ListDataAgentWorkspaceMemberResponseBody
	GetData() *ListDataAgentWorkspaceMemberResponseBodyData
	SetErrorCode(v string) *ListDataAgentWorkspaceMemberResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDataAgentWorkspaceMemberResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListDataAgentWorkspaceMemberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataAgentWorkspaceMemberResponseBody
	GetSuccess() *bool
}

type ListDataAgentWorkspaceMemberResponseBody struct {
	// The returned data.
	Data *ListDataAgentWorkspaceMemberResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// DMS-DA-40114
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Specified parameter is not valid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E0D2-*****-A63B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDataAgentWorkspaceMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentWorkspaceMemberResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataAgentWorkspaceMemberResponseBody) GetData() *ListDataAgentWorkspaceMemberResponseBodyData {
	return s.Data
}

func (s *ListDataAgentWorkspaceMemberResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDataAgentWorkspaceMemberResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDataAgentWorkspaceMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataAgentWorkspaceMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataAgentWorkspaceMemberResponseBody) SetData(v *ListDataAgentWorkspaceMemberResponseBodyData) *ListDataAgentWorkspaceMemberResponseBody {
	s.Data = v
	return s
}

func (s *ListDataAgentWorkspaceMemberResponseBody) SetErrorCode(v string) *ListDataAgentWorkspaceMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDataAgentWorkspaceMemberResponseBody) SetErrorMessage(v string) *ListDataAgentWorkspaceMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDataAgentWorkspaceMemberResponseBody) SetRequestId(v string) *ListDataAgentWorkspaceMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataAgentWorkspaceMemberResponseBody) SetSuccess(v bool) *ListDataAgentWorkspaceMemberResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataAgentWorkspaceMemberResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataAgentWorkspaceMemberResponseBodyData struct {
	// The data content.
	Content []*ListDataAgentWorkspaceMemberResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query.
	//
	// example:
	//
	// no use
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 50
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// no use
	TotalElements *int64 `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// no use
	TotalPages *int64 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListDataAgentWorkspaceMemberResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentWorkspaceMemberResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataAgentWorkspaceMemberResponseBodyData) GetContent() []*ListDataAgentWorkspaceMemberResponseBodyDataContent {
	return s.Content
}

func (s *ListDataAgentWorkspaceMemberResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataAgentWorkspaceMemberResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataAgentWorkspaceMemberResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListDataAgentWorkspaceMemberResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListDataAgentWorkspaceMemberResponseBodyData) GetTotalElements() *int64 {
	return s.TotalElements
}

func (s *ListDataAgentWorkspaceMemberResponseBodyData) GetTotalPages() *int64 {
	return s.TotalPages
}

func (s *ListDataAgentWorkspaceMemberResponseBodyData) SetContent(v []*ListDataAgentWorkspaceMemberResponseBodyDataContent) *ListDataAgentWorkspaceMemberResponseBodyData {
	s.Content = v
	return s
}

func (s *ListDataAgentWorkspaceMemberResponseBodyData) SetMaxResults(v int32) *ListDataAgentWorkspaceMemberResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListDataAgentWorkspaceMemberResponseBodyData) SetNextToken(v string) *ListDataAgentWorkspaceMemberResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListDataAgentWorkspaceMemberResponseBodyData) SetPageNumber(v int64) *ListDataAgentWorkspaceMemberResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListDataAgentWorkspaceMemberResponseBodyData) SetPageSize(v int64) *ListDataAgentWorkspaceMemberResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListDataAgentWorkspaceMemberResponseBodyData) SetTotalElements(v int64) *ListDataAgentWorkspaceMemberResponseBodyData {
	s.TotalElements = &v
	return s
}

func (s *ListDataAgentWorkspaceMemberResponseBodyData) SetTotalPages(v int64) *ListDataAgentWorkspaceMemberResponseBodyData {
	s.TotalPages = &v
	return s
}

func (s *ListDataAgentWorkspaceMemberResponseBodyData) Validate() error {
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

type ListDataAgentWorkspaceMemberResponseBodyDataContent struct {
	// The time when the user joined the workspace. This is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1765961516
	JoinTime *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	// The Alibaba Cloud UID of the user.
	//
	// example:
	//
	// 20282*****7591
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// The name of the user\\"s role in the workspace.
	//
	// example:
	//
	// member
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// The number of tasks that are running for the user in the workspace.
	//
	// example:
	//
	// 15
	RunningTaskNumber *int64 `json:"RunningTaskNumber,omitempty" xml:"RunningTaskNumber,omitempty"`
	// The total number of tasks initiated by the user in the workspace.
	//
	// example:
	//
	// 30
	TotalTaskNumber *int64 `json:"TotalTaskNumber,omitempty" xml:"TotalTaskNumber,omitempty"`
	// The RAM username of the user.
	//
	// example:
	//
	// yunqitest
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListDataAgentWorkspaceMemberResponseBodyDataContent) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentWorkspaceMemberResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *ListDataAgentWorkspaceMemberResponseBodyDataContent) GetJoinTime() *string {
	return s.JoinTime
}

func (s *ListDataAgentWorkspaceMemberResponseBodyDataContent) GetMemberId() *string {
	return s.MemberId
}

func (s *ListDataAgentWorkspaceMemberResponseBodyDataContent) GetRoleName() *string {
	return s.RoleName
}

func (s *ListDataAgentWorkspaceMemberResponseBodyDataContent) GetRunningTaskNumber() *int64 {
	return s.RunningTaskNumber
}

func (s *ListDataAgentWorkspaceMemberResponseBodyDataContent) GetTotalTaskNumber() *int64 {
	return s.TotalTaskNumber
}

func (s *ListDataAgentWorkspaceMemberResponseBodyDataContent) GetUserName() *string {
	return s.UserName
}

func (s *ListDataAgentWorkspaceMemberResponseBodyDataContent) SetJoinTime(v string) *ListDataAgentWorkspaceMemberResponseBodyDataContent {
	s.JoinTime = &v
	return s
}

func (s *ListDataAgentWorkspaceMemberResponseBodyDataContent) SetMemberId(v string) *ListDataAgentWorkspaceMemberResponseBodyDataContent {
	s.MemberId = &v
	return s
}

func (s *ListDataAgentWorkspaceMemberResponseBodyDataContent) SetRoleName(v string) *ListDataAgentWorkspaceMemberResponseBodyDataContent {
	s.RoleName = &v
	return s
}

func (s *ListDataAgentWorkspaceMemberResponseBodyDataContent) SetRunningTaskNumber(v int64) *ListDataAgentWorkspaceMemberResponseBodyDataContent {
	s.RunningTaskNumber = &v
	return s
}

func (s *ListDataAgentWorkspaceMemberResponseBodyDataContent) SetTotalTaskNumber(v int64) *ListDataAgentWorkspaceMemberResponseBodyDataContent {
	s.TotalTaskNumber = &v
	return s
}

func (s *ListDataAgentWorkspaceMemberResponseBodyDataContent) SetUserName(v string) *ListDataAgentWorkspaceMemberResponseBodyDataContent {
	s.UserName = &v
	return s
}

func (s *ListDataAgentWorkspaceMemberResponseBodyDataContent) Validate() error {
	return dara.Validate(s)
}
