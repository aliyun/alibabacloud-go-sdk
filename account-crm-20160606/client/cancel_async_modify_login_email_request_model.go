// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAsyncModifyLoginEmailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMpk(v string) *CancelAsyncModifyLoginEmailRequest
	GetMpk() *string
	SetTraceNo(v string) *CancelAsyncModifyLoginEmailRequest
	GetTraceNo() *string
}

type CancelAsyncModifyLoginEmailRequest struct {
	// This parameter is required.
	Mpk *string `json:"Mpk,omitempty" xml:"Mpk,omitempty"`
	// This parameter is required.
	TraceNo *string `json:"TraceNo,omitempty" xml:"TraceNo,omitempty"`
}

func (s CancelAsyncModifyLoginEmailRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelAsyncModifyLoginEmailRequest) GoString() string {
	return s.String()
}

func (s *CancelAsyncModifyLoginEmailRequest) GetMpk() *string {
	return s.Mpk
}

func (s *CancelAsyncModifyLoginEmailRequest) GetTraceNo() *string {
	return s.TraceNo
}

func (s *CancelAsyncModifyLoginEmailRequest) SetMpk(v string) *CancelAsyncModifyLoginEmailRequest {
	s.Mpk = &v
	return s
}

func (s *CancelAsyncModifyLoginEmailRequest) SetTraceNo(v string) *CancelAsyncModifyLoginEmailRequest {
	s.TraceNo = &v
	return s
}

func (s *CancelAsyncModifyLoginEmailRequest) Validate() error {
	return dara.Validate(s)
}
