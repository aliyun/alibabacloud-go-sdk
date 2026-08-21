// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecordRemarkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateRecordRemarkRequest
	GetClientToken() *string
	SetLang(v string) *UpdateRecordRemarkRequest
	GetLang() *string
	SetRecordId(v int64) *UpdateRecordRemarkRequest
	GetRecordId() *int64
	SetRemark(v string) *UpdateRecordRemarkRequest
	GetRemark() *string
}

type UpdateRecordRemarkRequest struct {
	// A client token used to ensure the idempotence of a request. The client generates this value, which must be unique across different requests. The token can be up to 64 ASCII characters in length.
	//
	// example:
	//
	// 6447728c8578e66aacf062d2df4446dc
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The language of the response. Valid values:
	//
	// - zh: Chinese.
	//
	// - en: English.
	//
	// Default value: en.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the DNS record.
	//
	// This parameter is required.
	//
	// example:
	//
	// 202991****
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The remarks.
	//
	// example:
	//
	// test record
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s UpdateRecordRemarkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecordRemarkRequest) GoString() string {
	return s.String()
}

func (s *UpdateRecordRemarkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateRecordRemarkRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateRecordRemarkRequest) GetRecordId() *int64 {
	return s.RecordId
}

func (s *UpdateRecordRemarkRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateRecordRemarkRequest) SetClientToken(v string) *UpdateRecordRemarkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateRecordRemarkRequest) SetLang(v string) *UpdateRecordRemarkRequest {
	s.Lang = &v
	return s
}

func (s *UpdateRecordRemarkRequest) SetRecordId(v int64) *UpdateRecordRemarkRequest {
	s.RecordId = &v
	return s
}

func (s *UpdateRecordRemarkRequest) SetRemark(v string) *UpdateRecordRemarkRequest {
	s.Remark = &v
	return s
}

func (s *UpdateRecordRemarkRequest) Validate() error {
	return dara.Validate(s)
}
