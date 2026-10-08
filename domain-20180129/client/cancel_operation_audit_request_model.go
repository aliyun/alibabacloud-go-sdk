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
	// The audit record ID. You can query the audit record ID by using the [QueryOperationAuditInfoList](https://help.aliyun.com/document_detail/172568.html) API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	AuditRecordId *int64 `json:"AuditRecordId,omitempty" xml:"AuditRecordId,omitempty"`
	// The language of the error message returned by the API. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// Default value: **en**.
	//
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
