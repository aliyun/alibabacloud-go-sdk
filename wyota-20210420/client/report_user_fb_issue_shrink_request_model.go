// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportUserFbIssueShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ReportUserFbIssueShrinkRequest
	GetAppId() *string
	SetClientAppVersion(v string) *ReportUserFbIssueShrinkRequest
	GetClientAppVersion() *string
	SetClientId(v string) *ReportUserFbIssueShrinkRequest
	GetClientId() *string
	SetClientModel(v string) *ReportUserFbIssueShrinkRequest
	GetClientModel() *string
	SetClientOsName(v string) *ReportUserFbIssueShrinkRequest
	GetClientOsName() *string
	SetClientSn(v string) *ReportUserFbIssueShrinkRequest
	GetClientSn() *string
	SetClientVersion(v string) *ReportUserFbIssueShrinkRequest
	GetClientVersion() *string
	SetCustomerId(v string) *ReportUserFbIssueShrinkRequest
	GetCustomerId() *string
	SetDescription(v string) *ReportUserFbIssueShrinkRequest
	GetDescription() *string
	SetDesktopId(v string) *ReportUserFbIssueShrinkRequest
	GetDesktopId() *string
	SetDesktopType(v int32) *ReportUserFbIssueShrinkRequest
	GetDesktopType() *int32
	SetErrorCode(v string) *ReportUserFbIssueShrinkRequest
	GetErrorCode() *string
	SetErrorMsg(v string) *ReportUserFbIssueShrinkRequest
	GetErrorMsg() *string
	SetFbType(v int32) *ReportUserFbIssueShrinkRequest
	GetFbType() *int32
	SetFileListShrink(v string) *ReportUserFbIssueShrinkRequest
	GetFileListShrink() *string
	SetIsSubstituteReport(v bool) *ReportUserFbIssueShrinkRequest
	GetIsSubstituteReport() *bool
	SetIssueLabel(v string) *ReportUserFbIssueShrinkRequest
	GetIssueLabel() *string
	SetLoginRegionId(v string) *ReportUserFbIssueShrinkRequest
	GetLoginRegionId() *string
	SetLoginToken(v string) *ReportUserFbIssueShrinkRequest
	GetLoginToken() *string
	SetOccurTime(v int64) *ReportUserFbIssueShrinkRequest
	GetOccurTime() *int64
	SetReservedA(v string) *ReportUserFbIssueShrinkRequest
	GetReservedA() *string
	SetReservedB(v string) *ReportUserFbIssueShrinkRequest
	GetReservedB() *string
	SetSessionId(v string) *ReportUserFbIssueShrinkRequest
	GetSessionId() *string
	SetTelNo(v string) *ReportUserFbIssueShrinkRequest
	GetTelNo() *string
	SetTitle(v string) *ReportUserFbIssueShrinkRequest
	GetTitle() *string
	SetUserEmail(v string) *ReportUserFbIssueShrinkRequest
	GetUserEmail() *string
	SetUserId(v string) *ReportUserFbIssueShrinkRequest
	GetUserId() *string
	SetUserName(v string) *ReportUserFbIssueShrinkRequest
	GetUserName() *string
	SetWorkspaceId(v string) *ReportUserFbIssueShrinkRequest
	GetWorkspaceId() *string
	SetWyId(v string) *ReportUserFbIssueShrinkRequest
	GetWyId() *string
}

type ReportUserFbIssueShrinkRequest struct {
	AppId              *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ClientAppVersion   *string `json:"ClientAppVersion,omitempty" xml:"ClientAppVersion,omitempty"`
	ClientId           *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientModel        *string `json:"ClientModel,omitempty" xml:"ClientModel,omitempty"`
	ClientOsName       *string `json:"ClientOsName,omitempty" xml:"ClientOsName,omitempty"`
	ClientSn           *string `json:"ClientSn,omitempty" xml:"ClientSn,omitempty"`
	ClientVersion      *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	CustomerId         *string `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DesktopId          *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	DesktopType        *int32  `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	ErrorCode          *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg           *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	FbType             *int32  `json:"FbType,omitempty" xml:"FbType,omitempty"`
	FileListShrink     *string `json:"FileList,omitempty" xml:"FileList,omitempty"`
	IsSubstituteReport *bool   `json:"IsSubstituteReport,omitempty" xml:"IsSubstituteReport,omitempty"`
	IssueLabel         *string `json:"IssueLabel,omitempty" xml:"IssueLabel,omitempty"`
	LoginRegionId      *string `json:"LoginRegionId,omitempty" xml:"LoginRegionId,omitempty"`
	LoginToken         *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	OccurTime          *int64  `json:"OccurTime,omitempty" xml:"OccurTime,omitempty"`
	ReservedA          *string `json:"ReservedA,omitempty" xml:"ReservedA,omitempty"`
	ReservedB          *string `json:"ReservedB,omitempty" xml:"ReservedB,omitempty"`
	SessionId          *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	TelNo              *string `json:"TelNo,omitempty" xml:"TelNo,omitempty"`
	Title              *string `json:"Title,omitempty" xml:"Title,omitempty"`
	UserEmail          *string `json:"UserEmail,omitempty" xml:"UserEmail,omitempty"`
	UserId             *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName           *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	WorkspaceId        *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WyId               *string `json:"WyId,omitempty" xml:"WyId,omitempty"`
}

func (s ReportUserFbIssueShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ReportUserFbIssueShrinkRequest) GoString() string {
	return s.String()
}

func (s *ReportUserFbIssueShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *ReportUserFbIssueShrinkRequest) GetClientAppVersion() *string {
	return s.ClientAppVersion
}

func (s *ReportUserFbIssueShrinkRequest) GetClientId() *string {
	return s.ClientId
}

func (s *ReportUserFbIssueShrinkRequest) GetClientModel() *string {
	return s.ClientModel
}

func (s *ReportUserFbIssueShrinkRequest) GetClientOsName() *string {
	return s.ClientOsName
}

func (s *ReportUserFbIssueShrinkRequest) GetClientSn() *string {
	return s.ClientSn
}

func (s *ReportUserFbIssueShrinkRequest) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *ReportUserFbIssueShrinkRequest) GetCustomerId() *string {
	return s.CustomerId
}

func (s *ReportUserFbIssueShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *ReportUserFbIssueShrinkRequest) GetDesktopId() *string {
	return s.DesktopId
}

func (s *ReportUserFbIssueShrinkRequest) GetDesktopType() *int32 {
	return s.DesktopType
}

func (s *ReportUserFbIssueShrinkRequest) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ReportUserFbIssueShrinkRequest) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ReportUserFbIssueShrinkRequest) GetFbType() *int32 {
	return s.FbType
}

func (s *ReportUserFbIssueShrinkRequest) GetFileListShrink() *string {
	return s.FileListShrink
}

func (s *ReportUserFbIssueShrinkRequest) GetIsSubstituteReport() *bool {
	return s.IsSubstituteReport
}

func (s *ReportUserFbIssueShrinkRequest) GetIssueLabel() *string {
	return s.IssueLabel
}

func (s *ReportUserFbIssueShrinkRequest) GetLoginRegionId() *string {
	return s.LoginRegionId
}

func (s *ReportUserFbIssueShrinkRequest) GetLoginToken() *string {
	return s.LoginToken
}

func (s *ReportUserFbIssueShrinkRequest) GetOccurTime() *int64 {
	return s.OccurTime
}

func (s *ReportUserFbIssueShrinkRequest) GetReservedA() *string {
	return s.ReservedA
}

func (s *ReportUserFbIssueShrinkRequest) GetReservedB() *string {
	return s.ReservedB
}

func (s *ReportUserFbIssueShrinkRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ReportUserFbIssueShrinkRequest) GetTelNo() *string {
	return s.TelNo
}

func (s *ReportUserFbIssueShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *ReportUserFbIssueShrinkRequest) GetUserEmail() *string {
	return s.UserEmail
}

func (s *ReportUserFbIssueShrinkRequest) GetUserId() *string {
	return s.UserId
}

func (s *ReportUserFbIssueShrinkRequest) GetUserName() *string {
	return s.UserName
}

func (s *ReportUserFbIssueShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ReportUserFbIssueShrinkRequest) GetWyId() *string {
	return s.WyId
}

func (s *ReportUserFbIssueShrinkRequest) SetAppId(v string) *ReportUserFbIssueShrinkRequest {
	s.AppId = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetClientAppVersion(v string) *ReportUserFbIssueShrinkRequest {
	s.ClientAppVersion = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetClientId(v string) *ReportUserFbIssueShrinkRequest {
	s.ClientId = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetClientModel(v string) *ReportUserFbIssueShrinkRequest {
	s.ClientModel = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetClientOsName(v string) *ReportUserFbIssueShrinkRequest {
	s.ClientOsName = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetClientSn(v string) *ReportUserFbIssueShrinkRequest {
	s.ClientSn = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetClientVersion(v string) *ReportUserFbIssueShrinkRequest {
	s.ClientVersion = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetCustomerId(v string) *ReportUserFbIssueShrinkRequest {
	s.CustomerId = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetDescription(v string) *ReportUserFbIssueShrinkRequest {
	s.Description = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetDesktopId(v string) *ReportUserFbIssueShrinkRequest {
	s.DesktopId = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetDesktopType(v int32) *ReportUserFbIssueShrinkRequest {
	s.DesktopType = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetErrorCode(v string) *ReportUserFbIssueShrinkRequest {
	s.ErrorCode = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetErrorMsg(v string) *ReportUserFbIssueShrinkRequest {
	s.ErrorMsg = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetFbType(v int32) *ReportUserFbIssueShrinkRequest {
	s.FbType = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetFileListShrink(v string) *ReportUserFbIssueShrinkRequest {
	s.FileListShrink = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetIsSubstituteReport(v bool) *ReportUserFbIssueShrinkRequest {
	s.IsSubstituteReport = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetIssueLabel(v string) *ReportUserFbIssueShrinkRequest {
	s.IssueLabel = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetLoginRegionId(v string) *ReportUserFbIssueShrinkRequest {
	s.LoginRegionId = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetLoginToken(v string) *ReportUserFbIssueShrinkRequest {
	s.LoginToken = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetOccurTime(v int64) *ReportUserFbIssueShrinkRequest {
	s.OccurTime = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetReservedA(v string) *ReportUserFbIssueShrinkRequest {
	s.ReservedA = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetReservedB(v string) *ReportUserFbIssueShrinkRequest {
	s.ReservedB = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetSessionId(v string) *ReportUserFbIssueShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetTelNo(v string) *ReportUserFbIssueShrinkRequest {
	s.TelNo = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetTitle(v string) *ReportUserFbIssueShrinkRequest {
	s.Title = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetUserEmail(v string) *ReportUserFbIssueShrinkRequest {
	s.UserEmail = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetUserId(v string) *ReportUserFbIssueShrinkRequest {
	s.UserId = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetUserName(v string) *ReportUserFbIssueShrinkRequest {
	s.UserName = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetWorkspaceId(v string) *ReportUserFbIssueShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) SetWyId(v string) *ReportUserFbIssueShrinkRequest {
	s.WyId = &v
	return s
}

func (s *ReportUserFbIssueShrinkRequest) Validate() error {
	return dara.Validate(s)
}
