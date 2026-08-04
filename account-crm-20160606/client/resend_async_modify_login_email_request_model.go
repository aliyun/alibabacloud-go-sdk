// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResendAsyncModifyLoginEmailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMpk(v string) *ResendAsyncModifyLoginEmailRequest
	GetMpk() *string
	SetTraceNo(v string) *ResendAsyncModifyLoginEmailRequest
	GetTraceNo() *string
}

type ResendAsyncModifyLoginEmailRequest struct {
	// This parameter is required.
	Mpk *string `json:"Mpk,omitempty" xml:"Mpk,omitempty"`
	// This parameter is required.
	TraceNo *string `json:"TraceNo,omitempty" xml:"TraceNo,omitempty"`
}

func (s ResendAsyncModifyLoginEmailRequest) String() string {
	return dara.Prettify(s)
}

func (s ResendAsyncModifyLoginEmailRequest) GoString() string {
	return s.String()
}

func (s *ResendAsyncModifyLoginEmailRequest) GetMpk() *string {
	return s.Mpk
}

func (s *ResendAsyncModifyLoginEmailRequest) GetTraceNo() *string {
	return s.TraceNo
}

func (s *ResendAsyncModifyLoginEmailRequest) SetMpk(v string) *ResendAsyncModifyLoginEmailRequest {
	s.Mpk = &v
	return s
}

func (s *ResendAsyncModifyLoginEmailRequest) SetTraceNo(v string) *ResendAsyncModifyLoginEmailRequest {
	s.TraceNo = &v
	return s
}

func (s *ResendAsyncModifyLoginEmailRequest) Validate() error {
	return dara.Validate(s)
}
