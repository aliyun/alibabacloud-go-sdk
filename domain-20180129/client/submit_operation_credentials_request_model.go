// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitOperationCredentialsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuditRecordId(v int64) *SubmitOperationCredentialsRequest
	GetAuditRecordId() *int64
	SetAuditType(v int32) *SubmitOperationCredentialsRequest
	GetAuditType() *int32
	SetCredentials(v string) *SubmitOperationCredentialsRequest
	GetCredentials() *string
	SetLang(v string) *SubmitOperationCredentialsRequest
	GetLang() *string
	SetRegType(v int32) *SubmitOperationCredentialsRequest
	GetRegType() *int32
}

type SubmitOperationCredentialsRequest struct {
	// example:
	//
	// 1
	AuditRecordId *int64 `json:"AuditRecordId,omitempty" xml:"AuditRecordId,omitempty"`
	// example:
	//
	// 1
	AuditType   *int32  `json:"AuditType,omitempty" xml:"AuditType,omitempty"`
	Credentials *string `json:"Credentials,omitempty" xml:"Credentials,omitempty"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 1
	RegType *int32 `json:"RegType,omitempty" xml:"RegType,omitempty"`
}

func (s SubmitOperationCredentialsRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitOperationCredentialsRequest) GoString() string {
	return s.String()
}

func (s *SubmitOperationCredentialsRequest) GetAuditRecordId() *int64 {
	return s.AuditRecordId
}

func (s *SubmitOperationCredentialsRequest) GetAuditType() *int32 {
	return s.AuditType
}

func (s *SubmitOperationCredentialsRequest) GetCredentials() *string {
	return s.Credentials
}

func (s *SubmitOperationCredentialsRequest) GetLang() *string {
	return s.Lang
}

func (s *SubmitOperationCredentialsRequest) GetRegType() *int32 {
	return s.RegType
}

func (s *SubmitOperationCredentialsRequest) SetAuditRecordId(v int64) *SubmitOperationCredentialsRequest {
	s.AuditRecordId = &v
	return s
}

func (s *SubmitOperationCredentialsRequest) SetAuditType(v int32) *SubmitOperationCredentialsRequest {
	s.AuditType = &v
	return s
}

func (s *SubmitOperationCredentialsRequest) SetCredentials(v string) *SubmitOperationCredentialsRequest {
	s.Credentials = &v
	return s
}

func (s *SubmitOperationCredentialsRequest) SetLang(v string) *SubmitOperationCredentialsRequest {
	s.Lang = &v
	return s
}

func (s *SubmitOperationCredentialsRequest) SetRegType(v int32) *SubmitOperationCredentialsRequest {
	s.RegType = &v
	return s
}

func (s *SubmitOperationCredentialsRequest) Validate() error {
	return dara.Validate(s)
}
