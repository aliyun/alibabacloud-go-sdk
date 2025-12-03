// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerRepositoryMirrorSyncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *TriggerRepositoryMirrorSyncRequest
	GetAccessToken() *string
	SetAccount(v string) *TriggerRepositoryMirrorSyncRequest
	GetAccount() *string
	SetOrganizationId(v string) *TriggerRepositoryMirrorSyncRequest
	GetOrganizationId() *string
	SetToken(v string) *TriggerRepositoryMirrorSyncRequest
	GetToken() *string
}

type TriggerRepositoryMirrorSyncRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// example:
	//
	// test-account
	Account *string `json:"account,omitempty" xml:"account,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60de7a6852743a5162b5f957
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// example:
	//
	// asd12e44827fe2444f952e931e51xxxx
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s TriggerRepositoryMirrorSyncRequest) String() string {
	return dara.Prettify(s)
}

func (s TriggerRepositoryMirrorSyncRequest) GoString() string {
	return s.String()
}

func (s *TriggerRepositoryMirrorSyncRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *TriggerRepositoryMirrorSyncRequest) GetAccount() *string {
	return s.Account
}

func (s *TriggerRepositoryMirrorSyncRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *TriggerRepositoryMirrorSyncRequest) GetToken() *string {
	return s.Token
}

func (s *TriggerRepositoryMirrorSyncRequest) SetAccessToken(v string) *TriggerRepositoryMirrorSyncRequest {
	s.AccessToken = &v
	return s
}

func (s *TriggerRepositoryMirrorSyncRequest) SetAccount(v string) *TriggerRepositoryMirrorSyncRequest {
	s.Account = &v
	return s
}

func (s *TriggerRepositoryMirrorSyncRequest) SetOrganizationId(v string) *TriggerRepositoryMirrorSyncRequest {
	s.OrganizationId = &v
	return s
}

func (s *TriggerRepositoryMirrorSyncRequest) SetToken(v string) *TriggerRepositoryMirrorSyncRequest {
	s.Token = &v
	return s
}

func (s *TriggerRepositoryMirrorSyncRequest) Validate() error {
	return dara.Validate(s)
}
