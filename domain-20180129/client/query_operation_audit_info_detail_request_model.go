// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOperationAuditInfoDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuditRecordId(v int64) *QueryOperationAuditInfoDetailRequest
	GetAuditRecordId() *int64
	SetLang(v string) *QueryOperationAuditInfoDetailRequest
	GetLang() *string
}

type QueryOperationAuditInfoDetailRequest struct {
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

func (s QueryOperationAuditInfoDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryOperationAuditInfoDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryOperationAuditInfoDetailRequest) GetAuditRecordId() *int64 {
	return s.AuditRecordId
}

func (s *QueryOperationAuditInfoDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryOperationAuditInfoDetailRequest) SetAuditRecordId(v int64) *QueryOperationAuditInfoDetailRequest {
	s.AuditRecordId = &v
	return s
}

func (s *QueryOperationAuditInfoDetailRequest) SetLang(v string) *QueryOperationAuditInfoDetailRequest {
	s.Lang = &v
	return s
}

func (s *QueryOperationAuditInfoDetailRequest) Validate() error {
	return dara.Validate(s)
}
