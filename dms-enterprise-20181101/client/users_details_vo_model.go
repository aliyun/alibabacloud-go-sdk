// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUsersDetailsVO interface {
	dara.Model
	String() string
	GoString() string
	SetApprovalSignatureBase64(v string) *UsersDetailsVO
	GetApprovalSignatureBase64() *string
	SetApprovalSqlTemplate(v string) *UsersDetailsVO
	GetApprovalSqlTemplate() *string
	SetApprovalStatus(v string) *UsersDetailsVO
	GetApprovalStatus() *string
	SetCreator(v int32) *UsersDetailsVO
	GetCreator() *int32
	SetDataReady(v int32) *UsersDetailsVO
	GetDataReady() *int32
	SetId(v int64) *UsersDetailsVO
	GetId() *int64
	SetMekid(v int64) *UsersDetailsVO
	GetMekid() *int64
	SetPathPrefix(v string) *UsersDetailsVO
	GetPathPrefix() *string
	SetResultParty(v int32) *UsersDetailsVO
	GetResultParty() *int32
	SetUid(v string) *UsersDetailsVO
	GetUid() *string
	SetUserConfirmed(v int32) *UsersDetailsVO
	GetUserConfirmed() *int32
	SetUserName(v string) *UsersDetailsVO
	GetUserName() *string
	SetUserPublicKeyPem(v string) *UsersDetailsVO
	GetUserPublicKeyPem() *string
}

type UsersDetailsVO struct {
	ApprovalSignatureBase64 *string `json:"ApprovalSignatureBase64,omitempty" xml:"ApprovalSignatureBase64,omitempty"`
	ApprovalSqlTemplate     *string `json:"ApprovalSqlTemplate,omitempty" xml:"ApprovalSqlTemplate,omitempty"`
	ApprovalStatus          *string `json:"ApprovalStatus,omitempty" xml:"ApprovalStatus,omitempty"`
	Creator                 *int32  `json:"Creator,omitempty" xml:"Creator,omitempty"`
	DataReady               *int32  `json:"DataReady,omitempty" xml:"DataReady,omitempty"`
	Id                      *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Mekid                   *int64  `json:"Mekid,omitempty" xml:"Mekid,omitempty"`
	PathPrefix              *string `json:"PathPrefix,omitempty" xml:"PathPrefix,omitempty"`
	ResultParty             *int32  `json:"ResultParty,omitempty" xml:"ResultParty,omitempty"`
	Uid                     *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	UserConfirmed           *int32  `json:"UserConfirmed,omitempty" xml:"UserConfirmed,omitempty"`
	UserName                *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserPublicKeyPem        *string `json:"UserPublicKeyPem,omitempty" xml:"UserPublicKeyPem,omitempty"`
}

func (s UsersDetailsVO) String() string {
	return dara.Prettify(s)
}

func (s UsersDetailsVO) GoString() string {
	return s.String()
}

func (s *UsersDetailsVO) GetApprovalSignatureBase64() *string {
	return s.ApprovalSignatureBase64
}

func (s *UsersDetailsVO) GetApprovalSqlTemplate() *string {
	return s.ApprovalSqlTemplate
}

func (s *UsersDetailsVO) GetApprovalStatus() *string {
	return s.ApprovalStatus
}

func (s *UsersDetailsVO) GetCreator() *int32 {
	return s.Creator
}

func (s *UsersDetailsVO) GetDataReady() *int32 {
	return s.DataReady
}

func (s *UsersDetailsVO) GetId() *int64 {
	return s.Id
}

func (s *UsersDetailsVO) GetMekid() *int64 {
	return s.Mekid
}

func (s *UsersDetailsVO) GetPathPrefix() *string {
	return s.PathPrefix
}

func (s *UsersDetailsVO) GetResultParty() *int32 {
	return s.ResultParty
}

func (s *UsersDetailsVO) GetUid() *string {
	return s.Uid
}

func (s *UsersDetailsVO) GetUserConfirmed() *int32 {
	return s.UserConfirmed
}

func (s *UsersDetailsVO) GetUserName() *string {
	return s.UserName
}

func (s *UsersDetailsVO) GetUserPublicKeyPem() *string {
	return s.UserPublicKeyPem
}

func (s *UsersDetailsVO) SetApprovalSignatureBase64(v string) *UsersDetailsVO {
	s.ApprovalSignatureBase64 = &v
	return s
}

func (s *UsersDetailsVO) SetApprovalSqlTemplate(v string) *UsersDetailsVO {
	s.ApprovalSqlTemplate = &v
	return s
}

func (s *UsersDetailsVO) SetApprovalStatus(v string) *UsersDetailsVO {
	s.ApprovalStatus = &v
	return s
}

func (s *UsersDetailsVO) SetCreator(v int32) *UsersDetailsVO {
	s.Creator = &v
	return s
}

func (s *UsersDetailsVO) SetDataReady(v int32) *UsersDetailsVO {
	s.DataReady = &v
	return s
}

func (s *UsersDetailsVO) SetId(v int64) *UsersDetailsVO {
	s.Id = &v
	return s
}

func (s *UsersDetailsVO) SetMekid(v int64) *UsersDetailsVO {
	s.Mekid = &v
	return s
}

func (s *UsersDetailsVO) SetPathPrefix(v string) *UsersDetailsVO {
	s.PathPrefix = &v
	return s
}

func (s *UsersDetailsVO) SetResultParty(v int32) *UsersDetailsVO {
	s.ResultParty = &v
	return s
}

func (s *UsersDetailsVO) SetUid(v string) *UsersDetailsVO {
	s.Uid = &v
	return s
}

func (s *UsersDetailsVO) SetUserConfirmed(v int32) *UsersDetailsVO {
	s.UserConfirmed = &v
	return s
}

func (s *UsersDetailsVO) SetUserName(v string) *UsersDetailsVO {
	s.UserName = &v
	return s
}

func (s *UsersDetailsVO) SetUserPublicKeyPem(v string) *UsersDetailsVO {
	s.UserPublicKeyPem = &v
	return s
}

func (s *UsersDetailsVO) Validate() error {
	return dara.Validate(s)
}
