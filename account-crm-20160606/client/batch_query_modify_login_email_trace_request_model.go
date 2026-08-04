// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchQueryModifyLoginEmailTraceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMpk(v string) *BatchQueryModifyLoginEmailTraceRequest
	GetMpk() *string
	SetTraceNoList(v string) *BatchQueryModifyLoginEmailTraceRequest
	GetTraceNoList() *string
}

type BatchQueryModifyLoginEmailTraceRequest struct {
	// This parameter is required.
	Mpk *string `json:"Mpk,omitempty" xml:"Mpk,omitempty"`
	// This parameter is required.
	TraceNoList *string `json:"TraceNoList,omitempty" xml:"TraceNoList,omitempty"`
}

func (s BatchQueryModifyLoginEmailTraceRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryModifyLoginEmailTraceRequest) GoString() string {
	return s.String()
}

func (s *BatchQueryModifyLoginEmailTraceRequest) GetMpk() *string {
	return s.Mpk
}

func (s *BatchQueryModifyLoginEmailTraceRequest) GetTraceNoList() *string {
	return s.TraceNoList
}

func (s *BatchQueryModifyLoginEmailTraceRequest) SetMpk(v string) *BatchQueryModifyLoginEmailTraceRequest {
	s.Mpk = &v
	return s
}

func (s *BatchQueryModifyLoginEmailTraceRequest) SetTraceNoList(v string) *BatchQueryModifyLoginEmailTraceRequest {
	s.TraceNoList = &v
	return s
}

func (s *BatchQueryModifyLoginEmailTraceRequest) Validate() error {
	return dara.Validate(s)
}
