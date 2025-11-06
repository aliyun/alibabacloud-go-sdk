// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitOperationAuditInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuditInfo(v string) *SubmitOperationAuditInfoRequest
	GetAuditInfo() *string
	SetAuditType(v int32) *SubmitOperationAuditInfoRequest
	GetAuditType() *int32
	SetDomainName(v string) *SubmitOperationAuditInfoRequest
	GetDomainName() *string
	SetId(v int64) *SubmitOperationAuditInfoRequest
	GetId() *int64
	SetLang(v string) *SubmitOperationAuditInfoRequest
	GetLang() *string
}

type SubmitOperationAuditInfoRequest struct {
	AuditInfo *string `json:"AuditInfo,omitempty" xml:"AuditInfo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	AuditType *int32 `json:"AuditType,omitempty" xml:"AuditType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// example.com,example.org
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s SubmitOperationAuditInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitOperationAuditInfoRequest) GoString() string {
	return s.String()
}

func (s *SubmitOperationAuditInfoRequest) GetAuditInfo() *string {
	return s.AuditInfo
}

func (s *SubmitOperationAuditInfoRequest) GetAuditType() *int32 {
	return s.AuditType
}

func (s *SubmitOperationAuditInfoRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SubmitOperationAuditInfoRequest) GetId() *int64 {
	return s.Id
}

func (s *SubmitOperationAuditInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *SubmitOperationAuditInfoRequest) SetAuditInfo(v string) *SubmitOperationAuditInfoRequest {
	s.AuditInfo = &v
	return s
}

func (s *SubmitOperationAuditInfoRequest) SetAuditType(v int32) *SubmitOperationAuditInfoRequest {
	s.AuditType = &v
	return s
}

func (s *SubmitOperationAuditInfoRequest) SetDomainName(v string) *SubmitOperationAuditInfoRequest {
	s.DomainName = &v
	return s
}

func (s *SubmitOperationAuditInfoRequest) SetId(v int64) *SubmitOperationAuditInfoRequest {
	s.Id = &v
	return s
}

func (s *SubmitOperationAuditInfoRequest) SetLang(v string) *SubmitOperationAuditInfoRequest {
	s.Lang = &v
	return s
}

func (s *SubmitOperationAuditInfoRequest) Validate() error {
	return dara.Validate(s)
}
