// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResendAsyncCreateAgAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMpk(v string) *ResendAsyncCreateAgAccountRequest
	GetMpk() *string
	SetTraceNo(v string) *ResendAsyncCreateAgAccountRequest
	GetTraceNo() *string
}

type ResendAsyncCreateAgAccountRequest struct {
	// This parameter is required.
	Mpk *string `json:"Mpk,omitempty" xml:"Mpk,omitempty"`
	// This parameter is required.
	TraceNo *string `json:"TraceNo,omitempty" xml:"TraceNo,omitempty"`
}

func (s ResendAsyncCreateAgAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s ResendAsyncCreateAgAccountRequest) GoString() string {
	return s.String()
}

func (s *ResendAsyncCreateAgAccountRequest) GetMpk() *string {
	return s.Mpk
}

func (s *ResendAsyncCreateAgAccountRequest) GetTraceNo() *string {
	return s.TraceNo
}

func (s *ResendAsyncCreateAgAccountRequest) SetMpk(v string) *ResendAsyncCreateAgAccountRequest {
	s.Mpk = &v
	return s
}

func (s *ResendAsyncCreateAgAccountRequest) SetTraceNo(v string) *ResendAsyncCreateAgAccountRequest {
	s.TraceNo = &v
	return s
}

func (s *ResendAsyncCreateAgAccountRequest) Validate() error {
	return dara.Validate(s)
}
