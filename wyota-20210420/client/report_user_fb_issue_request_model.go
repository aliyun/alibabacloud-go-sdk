// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportUserFbIssueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ReportUserFbIssueRequest
	GetAppId() *string
	SetClientAppVersion(v string) *ReportUserFbIssueRequest
	GetClientAppVersion() *string
	SetClientId(v string) *ReportUserFbIssueRequest
	GetClientId() *string
	SetClientModel(v string) *ReportUserFbIssueRequest
	GetClientModel() *string
	SetClientOsName(v string) *ReportUserFbIssueRequest
	GetClientOsName() *string
	SetClientSn(v string) *ReportUserFbIssueRequest
	GetClientSn() *string
	SetClientVersion(v string) *ReportUserFbIssueRequest
	GetClientVersion() *string
	SetCustomerId(v string) *ReportUserFbIssueRequest
	GetCustomerId() *string
	SetDescription(v string) *ReportUserFbIssueRequest
	GetDescription() *string
	SetDesktopId(v string) *ReportUserFbIssueRequest
	GetDesktopId() *string
	SetDesktopType(v int32) *ReportUserFbIssueRequest
	GetDesktopType() *int32
	SetErrorCode(v string) *ReportUserFbIssueRequest
	GetErrorCode() *string
	SetErrorMsg(v string) *ReportUserFbIssueRequest
	GetErrorMsg() *string
	SetFbType(v int32) *ReportUserFbIssueRequest
	GetFbType() *int32
	SetFileList(v []*ReportUserFbIssueRequestFileList) *ReportUserFbIssueRequest
	GetFileList() []*ReportUserFbIssueRequestFileList
	SetIsSubstituteReport(v bool) *ReportUserFbIssueRequest
	GetIsSubstituteReport() *bool
	SetIssueLabel(v string) *ReportUserFbIssueRequest
	GetIssueLabel() *string
	SetLoginRegionId(v string) *ReportUserFbIssueRequest
	GetLoginRegionId() *string
	SetLoginToken(v string) *ReportUserFbIssueRequest
	GetLoginToken() *string
	SetOccurTime(v int64) *ReportUserFbIssueRequest
	GetOccurTime() *int64
	SetReservedA(v string) *ReportUserFbIssueRequest
	GetReservedA() *string
	SetReservedB(v string) *ReportUserFbIssueRequest
	GetReservedB() *string
	SetSessionId(v string) *ReportUserFbIssueRequest
	GetSessionId() *string
	SetTelNo(v string) *ReportUserFbIssueRequest
	GetTelNo() *string
	SetTitle(v string) *ReportUserFbIssueRequest
	GetTitle() *string
	SetUserEmail(v string) *ReportUserFbIssueRequest
	GetUserEmail() *string
	SetUserId(v string) *ReportUserFbIssueRequest
	GetUserId() *string
	SetUserName(v string) *ReportUserFbIssueRequest
	GetUserName() *string
	SetWorkspaceId(v string) *ReportUserFbIssueRequest
	GetWorkspaceId() *string
	SetWyId(v string) *ReportUserFbIssueRequest
	GetWyId() *string
}

type ReportUserFbIssueRequest struct {
	AppId              *string                             `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ClientAppVersion   *string                             `json:"ClientAppVersion,omitempty" xml:"ClientAppVersion,omitempty"`
	ClientId           *string                             `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientModel        *string                             `json:"ClientModel,omitempty" xml:"ClientModel,omitempty"`
	ClientOsName       *string                             `json:"ClientOsName,omitempty" xml:"ClientOsName,omitempty"`
	ClientSn           *string                             `json:"ClientSn,omitempty" xml:"ClientSn,omitempty"`
	ClientVersion      *string                             `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	CustomerId         *string                             `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	Description        *string                             `json:"Description,omitempty" xml:"Description,omitempty"`
	DesktopId          *string                             `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	DesktopType        *int32                              `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	ErrorCode          *string                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg           *string                             `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	FbType             *int32                              `json:"FbType,omitempty" xml:"FbType,omitempty"`
	FileList           []*ReportUserFbIssueRequestFileList `json:"FileList,omitempty" xml:"FileList,omitempty" type:"Repeated"`
	IsSubstituteReport *bool                               `json:"IsSubstituteReport,omitempty" xml:"IsSubstituteReport,omitempty"`
	IssueLabel         *string                             `json:"IssueLabel,omitempty" xml:"IssueLabel,omitempty"`
	LoginRegionId      *string                             `json:"LoginRegionId,omitempty" xml:"LoginRegionId,omitempty"`
	LoginToken         *string                             `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	OccurTime          *int64                              `json:"OccurTime,omitempty" xml:"OccurTime,omitempty"`
	ReservedA          *string                             `json:"ReservedA,omitempty" xml:"ReservedA,omitempty"`
	ReservedB          *string                             `json:"ReservedB,omitempty" xml:"ReservedB,omitempty"`
	SessionId          *string                             `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	TelNo              *string                             `json:"TelNo,omitempty" xml:"TelNo,omitempty"`
	Title              *string                             `json:"Title,omitempty" xml:"Title,omitempty"`
	UserEmail          *string                             `json:"UserEmail,omitempty" xml:"UserEmail,omitempty"`
	UserId             *string                             `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName           *string                             `json:"UserName,omitempty" xml:"UserName,omitempty"`
	WorkspaceId        *string                             `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WyId               *string                             `json:"WyId,omitempty" xml:"WyId,omitempty"`
}

func (s ReportUserFbIssueRequest) String() string {
	return dara.Prettify(s)
}

func (s ReportUserFbIssueRequest) GoString() string {
	return s.String()
}

func (s *ReportUserFbIssueRequest) GetAppId() *string {
	return s.AppId
}

func (s *ReportUserFbIssueRequest) GetClientAppVersion() *string {
	return s.ClientAppVersion
}

func (s *ReportUserFbIssueRequest) GetClientId() *string {
	return s.ClientId
}

func (s *ReportUserFbIssueRequest) GetClientModel() *string {
	return s.ClientModel
}

func (s *ReportUserFbIssueRequest) GetClientOsName() *string {
	return s.ClientOsName
}

func (s *ReportUserFbIssueRequest) GetClientSn() *string {
	return s.ClientSn
}

func (s *ReportUserFbIssueRequest) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *ReportUserFbIssueRequest) GetCustomerId() *string {
	return s.CustomerId
}

func (s *ReportUserFbIssueRequest) GetDescription() *string {
	return s.Description
}

func (s *ReportUserFbIssueRequest) GetDesktopId() *string {
	return s.DesktopId
}

func (s *ReportUserFbIssueRequest) GetDesktopType() *int32 {
	return s.DesktopType
}

func (s *ReportUserFbIssueRequest) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ReportUserFbIssueRequest) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ReportUserFbIssueRequest) GetFbType() *int32 {
	return s.FbType
}

func (s *ReportUserFbIssueRequest) GetFileList() []*ReportUserFbIssueRequestFileList {
	return s.FileList
}

func (s *ReportUserFbIssueRequest) GetIsSubstituteReport() *bool {
	return s.IsSubstituteReport
}

func (s *ReportUserFbIssueRequest) GetIssueLabel() *string {
	return s.IssueLabel
}

func (s *ReportUserFbIssueRequest) GetLoginRegionId() *string {
	return s.LoginRegionId
}

func (s *ReportUserFbIssueRequest) GetLoginToken() *string {
	return s.LoginToken
}

func (s *ReportUserFbIssueRequest) GetOccurTime() *int64 {
	return s.OccurTime
}

func (s *ReportUserFbIssueRequest) GetReservedA() *string {
	return s.ReservedA
}

func (s *ReportUserFbIssueRequest) GetReservedB() *string {
	return s.ReservedB
}

func (s *ReportUserFbIssueRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ReportUserFbIssueRequest) GetTelNo() *string {
	return s.TelNo
}

func (s *ReportUserFbIssueRequest) GetTitle() *string {
	return s.Title
}

func (s *ReportUserFbIssueRequest) GetUserEmail() *string {
	return s.UserEmail
}

func (s *ReportUserFbIssueRequest) GetUserId() *string {
	return s.UserId
}

func (s *ReportUserFbIssueRequest) GetUserName() *string {
	return s.UserName
}

func (s *ReportUserFbIssueRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ReportUserFbIssueRequest) GetWyId() *string {
	return s.WyId
}

func (s *ReportUserFbIssueRequest) SetAppId(v string) *ReportUserFbIssueRequest {
	s.AppId = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetClientAppVersion(v string) *ReportUserFbIssueRequest {
	s.ClientAppVersion = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetClientId(v string) *ReportUserFbIssueRequest {
	s.ClientId = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetClientModel(v string) *ReportUserFbIssueRequest {
	s.ClientModel = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetClientOsName(v string) *ReportUserFbIssueRequest {
	s.ClientOsName = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetClientSn(v string) *ReportUserFbIssueRequest {
	s.ClientSn = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetClientVersion(v string) *ReportUserFbIssueRequest {
	s.ClientVersion = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetCustomerId(v string) *ReportUserFbIssueRequest {
	s.CustomerId = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetDescription(v string) *ReportUserFbIssueRequest {
	s.Description = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetDesktopId(v string) *ReportUserFbIssueRequest {
	s.DesktopId = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetDesktopType(v int32) *ReportUserFbIssueRequest {
	s.DesktopType = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetErrorCode(v string) *ReportUserFbIssueRequest {
	s.ErrorCode = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetErrorMsg(v string) *ReportUserFbIssueRequest {
	s.ErrorMsg = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetFbType(v int32) *ReportUserFbIssueRequest {
	s.FbType = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetFileList(v []*ReportUserFbIssueRequestFileList) *ReportUserFbIssueRequest {
	s.FileList = v
	return s
}

func (s *ReportUserFbIssueRequest) SetIsSubstituteReport(v bool) *ReportUserFbIssueRequest {
	s.IsSubstituteReport = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetIssueLabel(v string) *ReportUserFbIssueRequest {
	s.IssueLabel = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetLoginRegionId(v string) *ReportUserFbIssueRequest {
	s.LoginRegionId = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetLoginToken(v string) *ReportUserFbIssueRequest {
	s.LoginToken = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetOccurTime(v int64) *ReportUserFbIssueRequest {
	s.OccurTime = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetReservedA(v string) *ReportUserFbIssueRequest {
	s.ReservedA = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetReservedB(v string) *ReportUserFbIssueRequest {
	s.ReservedB = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetSessionId(v string) *ReportUserFbIssueRequest {
	s.SessionId = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetTelNo(v string) *ReportUserFbIssueRequest {
	s.TelNo = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetTitle(v string) *ReportUserFbIssueRequest {
	s.Title = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetUserEmail(v string) *ReportUserFbIssueRequest {
	s.UserEmail = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetUserId(v string) *ReportUserFbIssueRequest {
	s.UserId = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetUserName(v string) *ReportUserFbIssueRequest {
	s.UserName = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetWorkspaceId(v string) *ReportUserFbIssueRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ReportUserFbIssueRequest) SetWyId(v string) *ReportUserFbIssueRequest {
	s.WyId = &v
	return s
}

func (s *ReportUserFbIssueRequest) Validate() error {
	return dara.Validate(s)
}

type ReportUserFbIssueRequestFileList struct {
	FileMd5 *string `json:"FileMd5,omitempty" xml:"FileMd5,omitempty"`
	// This parameter is required.
	FileName  *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileSize  *int32  `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	FileType  *int32  `json:"FileType,omitempty" xml:"FileType,omitempty"`
	OssUrl    *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s ReportUserFbIssueRequestFileList) String() string {
	return dara.Prettify(s)
}

func (s ReportUserFbIssueRequestFileList) GoString() string {
	return s.String()
}

func (s *ReportUserFbIssueRequestFileList) GetFileMd5() *string {
	return s.FileMd5
}

func (s *ReportUserFbIssueRequestFileList) GetFileName() *string {
	return s.FileName
}

func (s *ReportUserFbIssueRequestFileList) GetFileSize() *int32 {
	return s.FileSize
}

func (s *ReportUserFbIssueRequestFileList) GetFileType() *int32 {
	return s.FileType
}

func (s *ReportUserFbIssueRequestFileList) GetOssUrl() *string {
	return s.OssUrl
}

func (s *ReportUserFbIssueRequestFileList) GetSessionId() *string {
	return s.SessionId
}

func (s *ReportUserFbIssueRequestFileList) SetFileMd5(v string) *ReportUserFbIssueRequestFileList {
	s.FileMd5 = &v
	return s
}

func (s *ReportUserFbIssueRequestFileList) SetFileName(v string) *ReportUserFbIssueRequestFileList {
	s.FileName = &v
	return s
}

func (s *ReportUserFbIssueRequestFileList) SetFileSize(v int32) *ReportUserFbIssueRequestFileList {
	s.FileSize = &v
	return s
}

func (s *ReportUserFbIssueRequestFileList) SetFileType(v int32) *ReportUserFbIssueRequestFileList {
	s.FileType = &v
	return s
}

func (s *ReportUserFbIssueRequestFileList) SetOssUrl(v string) *ReportUserFbIssueRequestFileList {
	s.OssUrl = &v
	return s
}

func (s *ReportUserFbIssueRequestFileList) SetSessionId(v string) *ReportUserFbIssueRequestFileList {
	s.SessionId = &v
	return s
}

func (s *ReportUserFbIssueRequestFileList) Validate() error {
	return dara.Validate(s)
}
