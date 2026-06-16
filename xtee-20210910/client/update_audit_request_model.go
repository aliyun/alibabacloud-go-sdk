// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAuditRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *UpdateAuditRequest
	GetLang() *string
	SetAuditMsg(v string) *UpdateAuditRequest
	GetAuditMsg() *string
	SetAuditRelationType(v string) *UpdateAuditRequest
	GetAuditRelationType() *string
	SetAuditStatus(v string) *UpdateAuditRequest
	GetAuditStatus() *string
	SetId(v int64) *UpdateAuditRequest
	GetId() *int64
	SetRegId(v string) *UpdateAuditRequest
	GetRegId() *string
}

type UpdateAuditRequest struct {
	// The language of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The approval comment.
	//
	// example:
	//
	// 同意
	AuditMsg *string `json:"auditMsg,omitempty" xml:"auditMsg,omitempty"`
	// The association type.
	//
	// example:
	//
	// RULE
	AuditRelationType *string `json:"auditRelationType,omitempty" xml:"auditRelationType,omitempty"`
	// The approval status.
	//
	// example:
	//
	// AGREE
	AuditStatus *string `json:"auditStatus,omitempty" xml:"auditStatus,omitempty"`
	// The ID of the approval to update.
	//
	// example:
	//
	// 376773
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// The region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s UpdateAuditRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuditRequest) GoString() string {
	return s.String()
}

func (s *UpdateAuditRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateAuditRequest) GetAuditMsg() *string {
	return s.AuditMsg
}

func (s *UpdateAuditRequest) GetAuditRelationType() *string {
	return s.AuditRelationType
}

func (s *UpdateAuditRequest) GetAuditStatus() *string {
	return s.AuditStatus
}

func (s *UpdateAuditRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateAuditRequest) GetRegId() *string {
	return s.RegId
}

func (s *UpdateAuditRequest) SetLang(v string) *UpdateAuditRequest {
	s.Lang = &v
	return s
}

func (s *UpdateAuditRequest) SetAuditMsg(v string) *UpdateAuditRequest {
	s.AuditMsg = &v
	return s
}

func (s *UpdateAuditRequest) SetAuditRelationType(v string) *UpdateAuditRequest {
	s.AuditRelationType = &v
	return s
}

func (s *UpdateAuditRequest) SetAuditStatus(v string) *UpdateAuditRequest {
	s.AuditStatus = &v
	return s
}

func (s *UpdateAuditRequest) SetId(v int64) *UpdateAuditRequest {
	s.Id = &v
	return s
}

func (s *UpdateAuditRequest) SetRegId(v string) *UpdateAuditRequest {
	s.RegId = &v
	return s
}

func (s *UpdateAuditRequest) Validate() error {
	return dara.Validate(s)
}
