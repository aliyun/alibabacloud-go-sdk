// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelOperationAuditRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuditRecordId(v int64) *CancelOperationAuditRequest
	GetAuditRecordId() *int64
	SetLang(v string) *CancelOperationAuditRequest
	GetLang() *string
}

type CancelOperationAuditRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	AuditRecordId *int64 `json:"AuditRecordId,omitempty" xml:"AuditRecordId,omitempty"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s CancelOperationAuditRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelOperationAuditRequest) GoString() string {
	return s.String()
}

func (s *CancelOperationAuditRequest) GetAuditRecordId() *int64 {
	return s.AuditRecordId
}

func (s *CancelOperationAuditRequest) GetLang() *string {
	return s.Lang
}

func (s *CancelOperationAuditRequest) SetAuditRecordId(v int64) *CancelOperationAuditRequest {
	s.AuditRecordId = &v
	return s
}

func (s *CancelOperationAuditRequest) SetLang(v string) *CancelOperationAuditRequest {
	s.Lang = &v
	return s
}

func (s *CancelOperationAuditRequest) Validate() error {
	return dara.Validate(s)
}
