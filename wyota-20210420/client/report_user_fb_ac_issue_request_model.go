// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportUserFbAcIssueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccount(v string) *ReportUserFbAcIssueRequest
	GetAccount() *string
	SetClientVersion(v string) *ReportUserFbAcIssueRequest
	GetClientVersion() *string
	SetErrorMsg(v string) *ReportUserFbAcIssueRequest
	GetErrorMsg() *string
	SetFileList(v []*ReportUserFbAcIssueRequestFileList) *ReportUserFbAcIssueRequest
	GetFileList() []*ReportUserFbAcIssueRequestFileList
	SetInstanceId(v string) *ReportUserFbAcIssueRequest
	GetInstanceId() *string
	SetLabels(v string) *ReportUserFbAcIssueRequest
	GetLabels() *string
	SetReservedA(v string) *ReportUserFbAcIssueRequest
	GetReservedA() *string
	SetReservedB(v string) *ReportUserFbAcIssueRequest
	GetReservedB() *string
	SetUserEmail(v string) *ReportUserFbAcIssueRequest
	GetUserEmail() *string
}

type ReportUserFbAcIssueRequest struct {
	Account       *string                               `json:"Account,omitempty" xml:"Account,omitempty"`
	ClientVersion *string                               `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	ErrorMsg      *string                               `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	FileList      []*ReportUserFbAcIssueRequestFileList `json:"FileList,omitempty" xml:"FileList,omitempty" type:"Repeated"`
	InstanceId    *string                               `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Labels        *string                               `json:"Labels,omitempty" xml:"Labels,omitempty"`
	ReservedA     *string                               `json:"ReservedA,omitempty" xml:"ReservedA,omitempty"`
	ReservedB     *string                               `json:"ReservedB,omitempty" xml:"ReservedB,omitempty"`
	UserEmail     *string                               `json:"UserEmail,omitempty" xml:"UserEmail,omitempty"`
}

func (s ReportUserFbAcIssueRequest) String() string {
	return dara.Prettify(s)
}

func (s ReportUserFbAcIssueRequest) GoString() string {
	return s.String()
}

func (s *ReportUserFbAcIssueRequest) GetAccount() *string {
	return s.Account
}

func (s *ReportUserFbAcIssueRequest) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *ReportUserFbAcIssueRequest) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ReportUserFbAcIssueRequest) GetFileList() []*ReportUserFbAcIssueRequestFileList {
	return s.FileList
}

func (s *ReportUserFbAcIssueRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReportUserFbAcIssueRequest) GetLabels() *string {
	return s.Labels
}

func (s *ReportUserFbAcIssueRequest) GetReservedA() *string {
	return s.ReservedA
}

func (s *ReportUserFbAcIssueRequest) GetReservedB() *string {
	return s.ReservedB
}

func (s *ReportUserFbAcIssueRequest) GetUserEmail() *string {
	return s.UserEmail
}

func (s *ReportUserFbAcIssueRequest) SetAccount(v string) *ReportUserFbAcIssueRequest {
	s.Account = &v
	return s
}

func (s *ReportUserFbAcIssueRequest) SetClientVersion(v string) *ReportUserFbAcIssueRequest {
	s.ClientVersion = &v
	return s
}

func (s *ReportUserFbAcIssueRequest) SetErrorMsg(v string) *ReportUserFbAcIssueRequest {
	s.ErrorMsg = &v
	return s
}

func (s *ReportUserFbAcIssueRequest) SetFileList(v []*ReportUserFbAcIssueRequestFileList) *ReportUserFbAcIssueRequest {
	s.FileList = v
	return s
}

func (s *ReportUserFbAcIssueRequest) SetInstanceId(v string) *ReportUserFbAcIssueRequest {
	s.InstanceId = &v
	return s
}

func (s *ReportUserFbAcIssueRequest) SetLabels(v string) *ReportUserFbAcIssueRequest {
	s.Labels = &v
	return s
}

func (s *ReportUserFbAcIssueRequest) SetReservedA(v string) *ReportUserFbAcIssueRequest {
	s.ReservedA = &v
	return s
}

func (s *ReportUserFbAcIssueRequest) SetReservedB(v string) *ReportUserFbAcIssueRequest {
	s.ReservedB = &v
	return s
}

func (s *ReportUserFbAcIssueRequest) SetUserEmail(v string) *ReportUserFbAcIssueRequest {
	s.UserEmail = &v
	return s
}

func (s *ReportUserFbAcIssueRequest) Validate() error {
	return dara.Validate(s)
}

type ReportUserFbAcIssueRequestFileList struct {
	// This parameter is required.
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileSize *int32  `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	FileType *int32  `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// This parameter is required.
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s ReportUserFbAcIssueRequestFileList) String() string {
	return dara.Prettify(s)
}

func (s ReportUserFbAcIssueRequestFileList) GoString() string {
	return s.String()
}

func (s *ReportUserFbAcIssueRequestFileList) GetFileName() *string {
	return s.FileName
}

func (s *ReportUserFbAcIssueRequestFileList) GetFileSize() *int32 {
	return s.FileSize
}

func (s *ReportUserFbAcIssueRequestFileList) GetFileType() *int32 {
	return s.FileType
}

func (s *ReportUserFbAcIssueRequestFileList) GetSessionId() *string {
	return s.SessionId
}

func (s *ReportUserFbAcIssueRequestFileList) SetFileName(v string) *ReportUserFbAcIssueRequestFileList {
	s.FileName = &v
	return s
}

func (s *ReportUserFbAcIssueRequestFileList) SetFileSize(v int32) *ReportUserFbAcIssueRequestFileList {
	s.FileSize = &v
	return s
}

func (s *ReportUserFbAcIssueRequestFileList) SetFileType(v int32) *ReportUserFbAcIssueRequestFileList {
	s.FileType = &v
	return s
}

func (s *ReportUserFbAcIssueRequestFileList) SetSessionId(v string) *ReportUserFbAcIssueRequestFileList {
	s.SessionId = &v
	return s
}

func (s *ReportUserFbAcIssueRequestFileList) Validate() error {
	return dara.Validate(s)
}
