// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOperationOssUploadPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuditType(v int32) *GetOperationOssUploadPolicyRequest
	GetAuditType() *int32
	SetLang(v string) *GetOperationOssUploadPolicyRequest
	GetLang() *string
}

type GetOperationOssUploadPolicyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	AuditType *int32 `json:"AuditType,omitempty" xml:"AuditType,omitempty"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s GetOperationOssUploadPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOperationOssUploadPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetOperationOssUploadPolicyRequest) GetAuditType() *int32 {
	return s.AuditType
}

func (s *GetOperationOssUploadPolicyRequest) GetLang() *string {
	return s.Lang
}

func (s *GetOperationOssUploadPolicyRequest) SetAuditType(v int32) *GetOperationOssUploadPolicyRequest {
	s.AuditType = &v
	return s
}

func (s *GetOperationOssUploadPolicyRequest) SetLang(v string) *GetOperationOssUploadPolicyRequest {
	s.Lang = &v
	return s
}

func (s *GetOperationOssUploadPolicyRequest) Validate() error {
	return dara.Validate(s)
}
