// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserFbIssuesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListUserFbIssuesRequest
	GetAppId() *string
	SetClientId(v string) *ListUserFbIssuesRequest
	GetClientId() *string
	SetClientModel(v string) *ListUserFbIssuesRequest
	GetClientModel() *string
	SetClientSn(v string) *ListUserFbIssuesRequest
	GetClientSn() *string
	SetCustomerId(v string) *ListUserFbIssuesRequest
	GetCustomerId() *string
	SetDescription(v string) *ListUserFbIssuesRequest
	GetDescription() *string
	SetDesktopId(v string) *ListUserFbIssuesRequest
	GetDesktopId() *string
	SetErrorCode(v string) *ListUserFbIssuesRequest
	GetErrorCode() *string
	SetErrorMsg(v string) *ListUserFbIssuesRequest
	GetErrorMsg() *string
	SetFbType(v int32) *ListUserFbIssuesRequest
	GetFbType() *int32
	SetIssueId(v int32) *ListUserFbIssuesRequest
	GetIssueId() *int32
	SetIssueLabel(v string) *ListUserFbIssuesRequest
	GetIssueLabel() *string
	SetPageNumber(v int32) *ListUserFbIssuesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListUserFbIssuesRequest
	GetPageSize() *int32
	SetStatus(v int32) *ListUserFbIssuesRequest
	GetStatus() *int32
	SetTitle(v string) *ListUserFbIssuesRequest
	GetTitle() *string
	SetUserEmail(v string) *ListUserFbIssuesRequest
	GetUserEmail() *string
	SetUserId(v string) *ListUserFbIssuesRequest
	GetUserId() *string
	SetWasRead(v int32) *ListUserFbIssuesRequest
	GetWasRead() *int32
}

type ListUserFbIssuesRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ClientId    *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientModel *string `json:"ClientModel,omitempty" xml:"ClientModel,omitempty"`
	ClientSn    *string `json:"ClientSn,omitempty" xml:"ClientSn,omitempty"`
	CustomerId  *string `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DesktopId   *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	ErrorCode   *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg    *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	FbType      *int32  `json:"FbType,omitempty" xml:"FbType,omitempty"`
	IssueId     *int32  `json:"IssueId,omitempty" xml:"IssueId,omitempty"`
	IssueLabel  *string `json:"IssueLabel,omitempty" xml:"IssueLabel,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Status      *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
	UserEmail   *string `json:"UserEmail,omitempty" xml:"UserEmail,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WasRead     *int32  `json:"WasRead,omitempty" xml:"WasRead,omitempty"`
}

func (s ListUserFbIssuesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserFbIssuesRequest) GoString() string {
	return s.String()
}

func (s *ListUserFbIssuesRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListUserFbIssuesRequest) GetClientId() *string {
	return s.ClientId
}

func (s *ListUserFbIssuesRequest) GetClientModel() *string {
	return s.ClientModel
}

func (s *ListUserFbIssuesRequest) GetClientSn() *string {
	return s.ClientSn
}

func (s *ListUserFbIssuesRequest) GetCustomerId() *string {
	return s.CustomerId
}

func (s *ListUserFbIssuesRequest) GetDescription() *string {
	return s.Description
}

func (s *ListUserFbIssuesRequest) GetDesktopId() *string {
	return s.DesktopId
}

func (s *ListUserFbIssuesRequest) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListUserFbIssuesRequest) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ListUserFbIssuesRequest) GetFbType() *int32 {
	return s.FbType
}

func (s *ListUserFbIssuesRequest) GetIssueId() *int32 {
	return s.IssueId
}

func (s *ListUserFbIssuesRequest) GetIssueLabel() *string {
	return s.IssueLabel
}

func (s *ListUserFbIssuesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListUserFbIssuesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUserFbIssuesRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ListUserFbIssuesRequest) GetTitle() *string {
	return s.Title
}

func (s *ListUserFbIssuesRequest) GetUserEmail() *string {
	return s.UserEmail
}

func (s *ListUserFbIssuesRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListUserFbIssuesRequest) GetWasRead() *int32 {
	return s.WasRead
}

func (s *ListUserFbIssuesRequest) SetAppId(v string) *ListUserFbIssuesRequest {
	s.AppId = &v
	return s
}

func (s *ListUserFbIssuesRequest) SetClientId(v string) *ListUserFbIssuesRequest {
	s.ClientId = &v
	return s
}

func (s *ListUserFbIssuesRequest) SetClientModel(v string) *ListUserFbIssuesRequest {
	s.ClientModel = &v
	return s
}

func (s *ListUserFbIssuesRequest) SetClientSn(v string) *ListUserFbIssuesRequest {
	s.ClientSn = &v
	return s
}

func (s *ListUserFbIssuesRequest) SetCustomerId(v string) *ListUserFbIssuesRequest {
	s.CustomerId = &v
	return s
}

func (s *ListUserFbIssuesRequest) SetDescription(v string) *ListUserFbIssuesRequest {
	s.Description = &v
	return s
}

func (s *ListUserFbIssuesRequest) SetDesktopId(v string) *ListUserFbIssuesRequest {
	s.DesktopId = &v
	return s
}

func (s *ListUserFbIssuesRequest) SetErrorCode(v string) *ListUserFbIssuesRequest {
	s.ErrorCode = &v
	return s
}

func (s *ListUserFbIssuesRequest) SetErrorMsg(v string) *ListUserFbIssuesRequest {
	s.ErrorMsg = &v
	return s
}

func (s *ListUserFbIssuesRequest) SetFbType(v int32) *ListUserFbIssuesRequest {
	s.FbType = &v
	return s
}

func (s *ListUserFbIssuesRequest) SetIssueId(v int32) *ListUserFbIssuesRequest {
	s.IssueId = &v
	return s
}

func (s *ListUserFbIssuesRequest) SetIssueLabel(v string) *ListUserFbIssuesRequest {
	s.IssueLabel = &v
	return s
}

func (s *ListUserFbIssuesRequest) SetPageNumber(v int32) *ListUserFbIssuesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUserFbIssuesRequest) SetPageSize(v int32) *ListUserFbIssuesRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserFbIssuesRequest) SetStatus(v int32) *ListUserFbIssuesRequest {
	s.Status = &v
	return s
}

func (s *ListUserFbIssuesRequest) SetTitle(v string) *ListUserFbIssuesRequest {
	s.Title = &v
	return s
}

func (s *ListUserFbIssuesRequest) SetUserEmail(v string) *ListUserFbIssuesRequest {
	s.UserEmail = &v
	return s
}

func (s *ListUserFbIssuesRequest) SetUserId(v string) *ListUserFbIssuesRequest {
	s.UserId = &v
	return s
}

func (s *ListUserFbIssuesRequest) SetWasRead(v int32) *ListUserFbIssuesRequest {
	s.WasRead = &v
	return s
}

func (s *ListUserFbIssuesRequest) Validate() error {
	return dara.Validate(s)
}
