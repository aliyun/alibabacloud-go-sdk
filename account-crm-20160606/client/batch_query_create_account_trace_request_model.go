// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchQueryCreateAccountTraceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMpk(v string) *BatchQueryCreateAccountTraceRequest
	GetMpk() *string
	SetTraceNoList(v string) *BatchQueryCreateAccountTraceRequest
	GetTraceNoList() *string
}

type BatchQueryCreateAccountTraceRequest struct {
	// This parameter is required.
	Mpk *string `json:"Mpk,omitempty" xml:"Mpk,omitempty"`
	// This parameter is required.
	TraceNoList *string `json:"TraceNoList,omitempty" xml:"TraceNoList,omitempty"`
}

func (s BatchQueryCreateAccountTraceRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryCreateAccountTraceRequest) GoString() string {
	return s.String()
}

func (s *BatchQueryCreateAccountTraceRequest) GetMpk() *string {
	return s.Mpk
}

func (s *BatchQueryCreateAccountTraceRequest) GetTraceNoList() *string {
	return s.TraceNoList
}

func (s *BatchQueryCreateAccountTraceRequest) SetMpk(v string) *BatchQueryCreateAccountTraceRequest {
	s.Mpk = &v
	return s
}

func (s *BatchQueryCreateAccountTraceRequest) SetTraceNoList(v string) *BatchQueryCreateAccountTraceRequest {
	s.TraceNoList = &v
	return s
}

func (s *BatchQueryCreateAccountTraceRequest) Validate() error {
	return dara.Validate(s)
}
