// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserFbAcIssuesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccount(v string) *ListUserFbAcIssuesRequest
	GetAccount() *string
	SetClientVersion(v string) *ListUserFbAcIssuesRequest
	GetClientVersion() *string
	SetErrorMessage(v string) *ListUserFbAcIssuesRequest
	GetErrorMessage() *string
	SetInstanceId(v string) *ListUserFbAcIssuesRequest
	GetInstanceId() *string
	SetIssueId(v string) *ListUserFbAcIssuesRequest
	GetIssueId() *string
	SetLabel(v string) *ListUserFbAcIssuesRequest
	GetLabel() *string
	SetReservedA(v string) *ListUserFbAcIssuesRequest
	GetReservedA() *string
	SetReservedB(v string) *ListUserFbAcIssuesRequest
	GetReservedB() *string
	SetUserEmail(v string) *ListUserFbAcIssuesRequest
	GetUserEmail() *string
}

type ListUserFbAcIssuesRequest struct {
	Account       *string `json:"Account,omitempty" xml:"Account,omitempty"`
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	ErrorMessage  *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IssueId       *string `json:"IssueId,omitempty" xml:"IssueId,omitempty"`
	Label         *string `json:"Label,omitempty" xml:"Label,omitempty"`
	ReservedA     *string `json:"ReservedA,omitempty" xml:"ReservedA,omitempty"`
	ReservedB     *string `json:"ReservedB,omitempty" xml:"ReservedB,omitempty"`
	UserEmail     *string `json:"UserEmail,omitempty" xml:"UserEmail,omitempty"`
}

func (s ListUserFbAcIssuesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserFbAcIssuesRequest) GoString() string {
	return s.String()
}

func (s *ListUserFbAcIssuesRequest) GetAccount() *string {
	return s.Account
}

func (s *ListUserFbAcIssuesRequest) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *ListUserFbAcIssuesRequest) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListUserFbAcIssuesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListUserFbAcIssuesRequest) GetIssueId() *string {
	return s.IssueId
}

func (s *ListUserFbAcIssuesRequest) GetLabel() *string {
	return s.Label
}

func (s *ListUserFbAcIssuesRequest) GetReservedA() *string {
	return s.ReservedA
}

func (s *ListUserFbAcIssuesRequest) GetReservedB() *string {
	return s.ReservedB
}

func (s *ListUserFbAcIssuesRequest) GetUserEmail() *string {
	return s.UserEmail
}

func (s *ListUserFbAcIssuesRequest) SetAccount(v string) *ListUserFbAcIssuesRequest {
	s.Account = &v
	return s
}

func (s *ListUserFbAcIssuesRequest) SetClientVersion(v string) *ListUserFbAcIssuesRequest {
	s.ClientVersion = &v
	return s
}

func (s *ListUserFbAcIssuesRequest) SetErrorMessage(v string) *ListUserFbAcIssuesRequest {
	s.ErrorMessage = &v
	return s
}

func (s *ListUserFbAcIssuesRequest) SetInstanceId(v string) *ListUserFbAcIssuesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListUserFbAcIssuesRequest) SetIssueId(v string) *ListUserFbAcIssuesRequest {
	s.IssueId = &v
	return s
}

func (s *ListUserFbAcIssuesRequest) SetLabel(v string) *ListUserFbAcIssuesRequest {
	s.Label = &v
	return s
}

func (s *ListUserFbAcIssuesRequest) SetReservedA(v string) *ListUserFbAcIssuesRequest {
	s.ReservedA = &v
	return s
}

func (s *ListUserFbAcIssuesRequest) SetReservedB(v string) *ListUserFbAcIssuesRequest {
	s.ReservedB = &v
	return s
}

func (s *ListUserFbAcIssuesRequest) SetUserEmail(v string) *ListUserFbAcIssuesRequest {
	s.UserEmail = &v
	return s
}

func (s *ListUserFbAcIssuesRequest) Validate() error {
	return dara.Validate(s)
}
