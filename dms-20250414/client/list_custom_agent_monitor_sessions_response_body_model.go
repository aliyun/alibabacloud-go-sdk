// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomAgentMonitorSessionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListCustomAgentMonitorSessionsResponseBodyData) *ListCustomAgentMonitorSessionsResponseBody
	GetData() *ListCustomAgentMonitorSessionsResponseBodyData
	SetErrorCode(v string) *ListCustomAgentMonitorSessionsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListCustomAgentMonitorSessionsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListCustomAgentMonitorSessionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCustomAgentMonitorSessionsResponseBody
	GetSuccess() *bool
}

type ListCustomAgentMonitorSessionsResponseBody struct {
	// The response struct.
	Data *ListCustomAgentMonitorSessionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned if the call fails.
	//
	// example:
	//
	// InvalidTid
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the call fails.
	//
	// example:
	//
	// Specified parameter Tid is not valid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E0D21075-CD3E-4D98-8264-FD8AD04A63B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCustomAgentMonitorSessionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCustomAgentMonitorSessionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomAgentMonitorSessionsResponseBody) GetData() *ListCustomAgentMonitorSessionsResponseBodyData {
	return s.Data
}

func (s *ListCustomAgentMonitorSessionsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListCustomAgentMonitorSessionsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListCustomAgentMonitorSessionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCustomAgentMonitorSessionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCustomAgentMonitorSessionsResponseBody) SetData(v *ListCustomAgentMonitorSessionsResponseBodyData) *ListCustomAgentMonitorSessionsResponseBody {
	s.Data = v
	return s
}

func (s *ListCustomAgentMonitorSessionsResponseBody) SetErrorCode(v string) *ListCustomAgentMonitorSessionsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListCustomAgentMonitorSessionsResponseBody) SetErrorMessage(v string) *ListCustomAgentMonitorSessionsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListCustomAgentMonitorSessionsResponseBody) SetRequestId(v string) *ListCustomAgentMonitorSessionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomAgentMonitorSessionsResponseBody) SetSuccess(v bool) *ListCustomAgentMonitorSessionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListCustomAgentMonitorSessionsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCustomAgentMonitorSessionsResponseBodyData struct {
	// The session details list for the current page, sorted by creation time in descending order.
	Content []*ListCustomAgentMonitorSessionsResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of sessions within the filter scope.
	//
	// example:
	//
	// 5
	TotalElements *int64 `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 1
	TotalPages *int64 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListCustomAgentMonitorSessionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCustomAgentMonitorSessionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCustomAgentMonitorSessionsResponseBodyData) GetContent() []*ListCustomAgentMonitorSessionsResponseBodyDataContent {
	return s.Content
}

func (s *ListCustomAgentMonitorSessionsResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListCustomAgentMonitorSessionsResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListCustomAgentMonitorSessionsResponseBodyData) GetTotalElements() *int64 {
	return s.TotalElements
}

func (s *ListCustomAgentMonitorSessionsResponseBodyData) GetTotalPages() *int64 {
	return s.TotalPages
}

func (s *ListCustomAgentMonitorSessionsResponseBodyData) SetContent(v []*ListCustomAgentMonitorSessionsResponseBodyDataContent) *ListCustomAgentMonitorSessionsResponseBodyData {
	s.Content = v
	return s
}

func (s *ListCustomAgentMonitorSessionsResponseBodyData) SetPageNumber(v int64) *ListCustomAgentMonitorSessionsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListCustomAgentMonitorSessionsResponseBodyData) SetPageSize(v int64) *ListCustomAgentMonitorSessionsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListCustomAgentMonitorSessionsResponseBodyData) SetTotalElements(v int64) *ListCustomAgentMonitorSessionsResponseBodyData {
	s.TotalElements = &v
	return s
}

func (s *ListCustomAgentMonitorSessionsResponseBodyData) SetTotalPages(v int64) *ListCustomAgentMonitorSessionsResponseBodyData {
	s.TotalPages = &v
	return s
}

func (s *ListCustomAgentMonitorSessionsResponseBodyData) Validate() error {
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

type ListCustomAgentMonitorSessionsResponseBodyDataContent struct {
	// The Alibaba Cloud UID of the creator.
	//
	// example:
	//
	// 20372822********
	AliyunUid *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	// The display name of the creator.
	//
	// example:
	//
	// HaoY***
	CreatorUserName *string `json:"CreatorUserName,omitempty" xml:"CreatorUserName,omitempty"`
	// The custom agent ID.
	//
	// example:
	//
	// ca-4y3ca4khkcu**********ysf
	CustomAgentId *string `json:"CustomAgentId,omitempty" xml:"CustomAgentId,omitempty"`
	// The number of dislikes for the session.
	//
	// example:
	//
	// 0
	DislikeCount *int64 `json:"DislikeCount,omitempty" xml:"DislikeCount,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-03-18T10:02:04+08:00
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2025-12-11T14:04:32.000+00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The number of likes for the session.
	//
	// example:
	//
	// 2
	LikeCount *int64 `json:"LikeCount,omitempty" xml:"LikeCount,omitempty"`
	// The session ID.
	//
	// example:
	//
	// 2gum46f149******ndfxxo
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The session status. Valid values:
	//
	// - init: The session is in the initial state.
	//
	// - INITIALIZING: The session is being initialized.
	//
	// - RUNNING: The session is running.
	//
	// - IDLE: The session is idle.
	//
	// - RECOVERABLE: The session is completed and can accept further questions.
	//
	// - UNAVAILABLE: The session is completed and cannot accept further questions.
	//
	// example:
	//
	// RUNNING
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The session name.
	//
	// example:
	//
	// Please analyze this data
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The total number of turns in the session.
	//
	// example:
	//
	// 10
	TurnCount *int64 `json:"TurnCount,omitempty" xml:"TurnCount,omitempty"`
}

func (s ListCustomAgentMonitorSessionsResponseBodyDataContent) String() string {
	return dara.Prettify(s)
}

func (s ListCustomAgentMonitorSessionsResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *ListCustomAgentMonitorSessionsResponseBodyDataContent) GetAliyunUid() *string {
	return s.AliyunUid
}

func (s *ListCustomAgentMonitorSessionsResponseBodyDataContent) GetCreatorUserName() *string {
	return s.CreatorUserName
}

func (s *ListCustomAgentMonitorSessionsResponseBodyDataContent) GetCustomAgentId() *string {
	return s.CustomAgentId
}

func (s *ListCustomAgentMonitorSessionsResponseBodyDataContent) GetDislikeCount() *int64 {
	return s.DislikeCount
}

func (s *ListCustomAgentMonitorSessionsResponseBodyDataContent) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *ListCustomAgentMonitorSessionsResponseBodyDataContent) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListCustomAgentMonitorSessionsResponseBodyDataContent) GetLikeCount() *int64 {
	return s.LikeCount
}

func (s *ListCustomAgentMonitorSessionsResponseBodyDataContent) GetSessionId() *string {
	return s.SessionId
}

func (s *ListCustomAgentMonitorSessionsResponseBodyDataContent) GetState() *string {
	return s.State
}

func (s *ListCustomAgentMonitorSessionsResponseBodyDataContent) GetTitle() *string {
	return s.Title
}

func (s *ListCustomAgentMonitorSessionsResponseBodyDataContent) GetTurnCount() *int64 {
	return s.TurnCount
}

func (s *ListCustomAgentMonitorSessionsResponseBodyDataContent) SetAliyunUid(v string) *ListCustomAgentMonitorSessionsResponseBodyDataContent {
	s.AliyunUid = &v
	return s
}

func (s *ListCustomAgentMonitorSessionsResponseBodyDataContent) SetCreatorUserName(v string) *ListCustomAgentMonitorSessionsResponseBodyDataContent {
	s.CreatorUserName = &v
	return s
}

func (s *ListCustomAgentMonitorSessionsResponseBodyDataContent) SetCustomAgentId(v string) *ListCustomAgentMonitorSessionsResponseBodyDataContent {
	s.CustomAgentId = &v
	return s
}

func (s *ListCustomAgentMonitorSessionsResponseBodyDataContent) SetDislikeCount(v int64) *ListCustomAgentMonitorSessionsResponseBodyDataContent {
	s.DislikeCount = &v
	return s
}

func (s *ListCustomAgentMonitorSessionsResponseBodyDataContent) SetGmtCreated(v string) *ListCustomAgentMonitorSessionsResponseBodyDataContent {
	s.GmtCreated = &v
	return s
}

func (s *ListCustomAgentMonitorSessionsResponseBodyDataContent) SetGmtModified(v string) *ListCustomAgentMonitorSessionsResponseBodyDataContent {
	s.GmtModified = &v
	return s
}

func (s *ListCustomAgentMonitorSessionsResponseBodyDataContent) SetLikeCount(v int64) *ListCustomAgentMonitorSessionsResponseBodyDataContent {
	s.LikeCount = &v
	return s
}

func (s *ListCustomAgentMonitorSessionsResponseBodyDataContent) SetSessionId(v string) *ListCustomAgentMonitorSessionsResponseBodyDataContent {
	s.SessionId = &v
	return s
}

func (s *ListCustomAgentMonitorSessionsResponseBodyDataContent) SetState(v string) *ListCustomAgentMonitorSessionsResponseBodyDataContent {
	s.State = &v
	return s
}

func (s *ListCustomAgentMonitorSessionsResponseBodyDataContent) SetTitle(v string) *ListCustomAgentMonitorSessionsResponseBodyDataContent {
	s.Title = &v
	return s
}

func (s *ListCustomAgentMonitorSessionsResponseBodyDataContent) SetTurnCount(v int64) *ListCustomAgentMonitorSessionsResponseBodyDataContent {
	s.TurnCount = &v
	return s
}

func (s *ListCustomAgentMonitorSessionsResponseBodyDataContent) Validate() error {
	return dara.Validate(s)
}
