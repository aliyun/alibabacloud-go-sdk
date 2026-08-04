// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAsyncCreateAgAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMpk(v string) *CancelAsyncCreateAgAccountRequest
	GetMpk() *string
	SetTraceNo(v string) *CancelAsyncCreateAgAccountRequest
	GetTraceNo() *string
}

type CancelAsyncCreateAgAccountRequest struct {
	// This parameter is required.
	Mpk *string `json:"Mpk,omitempty" xml:"Mpk,omitempty"`
	// This parameter is required.
	TraceNo *string `json:"TraceNo,omitempty" xml:"TraceNo,omitempty"`
}

func (s CancelAsyncCreateAgAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelAsyncCreateAgAccountRequest) GoString() string {
	return s.String()
}

func (s *CancelAsyncCreateAgAccountRequest) GetMpk() *string {
	return s.Mpk
}

func (s *CancelAsyncCreateAgAccountRequest) GetTraceNo() *string {
	return s.TraceNo
}

func (s *CancelAsyncCreateAgAccountRequest) SetMpk(v string) *CancelAsyncCreateAgAccountRequest {
	s.Mpk = &v
	return s
}

func (s *CancelAsyncCreateAgAccountRequest) SetTraceNo(v string) *CancelAsyncCreateAgAccountRequest {
	s.TraceNo = &v
	return s
}

func (s *CancelAsyncCreateAgAccountRequest) Validate() error {
	return dara.Validate(s)
}
