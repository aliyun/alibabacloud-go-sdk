// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportUserFbAcIssueShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccount(v string) *ReportUserFbAcIssueShrinkRequest
	GetAccount() *string
	SetClientVersion(v string) *ReportUserFbAcIssueShrinkRequest
	GetClientVersion() *string
	SetErrorMsg(v string) *ReportUserFbAcIssueShrinkRequest
	GetErrorMsg() *string
	SetFileListShrink(v string) *ReportUserFbAcIssueShrinkRequest
	GetFileListShrink() *string
	SetInstanceId(v string) *ReportUserFbAcIssueShrinkRequest
	GetInstanceId() *string
	SetLabels(v string) *ReportUserFbAcIssueShrinkRequest
	GetLabels() *string
	SetReservedA(v string) *ReportUserFbAcIssueShrinkRequest
	GetReservedA() *string
	SetReservedB(v string) *ReportUserFbAcIssueShrinkRequest
	GetReservedB() *string
	SetUserEmail(v string) *ReportUserFbAcIssueShrinkRequest
	GetUserEmail() *string
}

type ReportUserFbAcIssueShrinkRequest struct {
	Account        *string `json:"Account,omitempty" xml:"Account,omitempty"`
	ClientVersion  *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	ErrorMsg       *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	FileListShrink *string `json:"FileList,omitempty" xml:"FileList,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Labels         *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	ReservedA      *string `json:"ReservedA,omitempty" xml:"ReservedA,omitempty"`
	ReservedB      *string `json:"ReservedB,omitempty" xml:"ReservedB,omitempty"`
	UserEmail      *string `json:"UserEmail,omitempty" xml:"UserEmail,omitempty"`
}

func (s ReportUserFbAcIssueShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ReportUserFbAcIssueShrinkRequest) GoString() string {
	return s.String()
}

func (s *ReportUserFbAcIssueShrinkRequest) GetAccount() *string {
	return s.Account
}

func (s *ReportUserFbAcIssueShrinkRequest) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *ReportUserFbAcIssueShrinkRequest) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ReportUserFbAcIssueShrinkRequest) GetFileListShrink() *string {
	return s.FileListShrink
}

func (s *ReportUserFbAcIssueShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReportUserFbAcIssueShrinkRequest) GetLabels() *string {
	return s.Labels
}

func (s *ReportUserFbAcIssueShrinkRequest) GetReservedA() *string {
	return s.ReservedA
}

func (s *ReportUserFbAcIssueShrinkRequest) GetReservedB() *string {
	return s.ReservedB
}

func (s *ReportUserFbAcIssueShrinkRequest) GetUserEmail() *string {
	return s.UserEmail
}

func (s *ReportUserFbAcIssueShrinkRequest) SetAccount(v string) *ReportUserFbAcIssueShrinkRequest {
	s.Account = &v
	return s
}

func (s *ReportUserFbAcIssueShrinkRequest) SetClientVersion(v string) *ReportUserFbAcIssueShrinkRequest {
	s.ClientVersion = &v
	return s
}

func (s *ReportUserFbAcIssueShrinkRequest) SetErrorMsg(v string) *ReportUserFbAcIssueShrinkRequest {
	s.ErrorMsg = &v
	return s
}

func (s *ReportUserFbAcIssueShrinkRequest) SetFileListShrink(v string) *ReportUserFbAcIssueShrinkRequest {
	s.FileListShrink = &v
	return s
}

func (s *ReportUserFbAcIssueShrinkRequest) SetInstanceId(v string) *ReportUserFbAcIssueShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ReportUserFbAcIssueShrinkRequest) SetLabels(v string) *ReportUserFbAcIssueShrinkRequest {
	s.Labels = &v
	return s
}

func (s *ReportUserFbAcIssueShrinkRequest) SetReservedA(v string) *ReportUserFbAcIssueShrinkRequest {
	s.ReservedA = &v
	return s
}

func (s *ReportUserFbAcIssueShrinkRequest) SetReservedB(v string) *ReportUserFbAcIssueShrinkRequest {
	s.ReservedB = &v
	return s
}

func (s *ReportUserFbAcIssueShrinkRequest) SetUserEmail(v string) *ReportUserFbAcIssueShrinkRequest {
	s.UserEmail = &v
	return s
}

func (s *ReportUserFbAcIssueShrinkRequest) Validate() error {
	return dara.Validate(s)
}
