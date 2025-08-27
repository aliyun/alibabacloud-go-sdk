// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTripBusinessInstanceQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TripBusinessInstanceQueryResponseBody
	GetCode() *string
	SetMessage(v string) *TripBusinessInstanceQueryResponseBody
	GetMessage() *string
	SetModule(v *TripBusinessInstanceQueryResponseBodyModule) *TripBusinessInstanceQueryResponseBody
	GetModule() *TripBusinessInstanceQueryResponseBodyModule
	SetRequestId(v string) *TripBusinessInstanceQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TripBusinessInstanceQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *TripBusinessInstanceQueryResponseBody
	GetTraceId() *string
}

type TripBusinessInstanceQueryResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// module。
	Module *TripBusinessInstanceQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// C61ECFF6-606B-5F66-B81D-D77369043A5F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 210bc44416869853114684533da3c2
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s TripBusinessInstanceQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TripBusinessInstanceQueryResponseBody) GoString() string {
	return s.String()
}

func (s *TripBusinessInstanceQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *TripBusinessInstanceQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TripBusinessInstanceQueryResponseBody) GetModule() *TripBusinessInstanceQueryResponseBodyModule {
	return s.Module
}

func (s *TripBusinessInstanceQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TripBusinessInstanceQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TripBusinessInstanceQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *TripBusinessInstanceQueryResponseBody) SetCode(v string) *TripBusinessInstanceQueryResponseBody {
	s.Code = &v
	return s
}

func (s *TripBusinessInstanceQueryResponseBody) SetMessage(v string) *TripBusinessInstanceQueryResponseBody {
	s.Message = &v
	return s
}

func (s *TripBusinessInstanceQueryResponseBody) SetModule(v *TripBusinessInstanceQueryResponseBodyModule) *TripBusinessInstanceQueryResponseBody {
	s.Module = v
	return s
}

func (s *TripBusinessInstanceQueryResponseBody) SetRequestId(v string) *TripBusinessInstanceQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *TripBusinessInstanceQueryResponseBody) SetSuccess(v bool) *TripBusinessInstanceQueryResponseBody {
	s.Success = &v
	return s
}

func (s *TripBusinessInstanceQueryResponseBody) SetTraceId(v string) *TripBusinessInstanceQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *TripBusinessInstanceQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type TripBusinessInstanceQueryResponseBodyModule struct {
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// example:
	//
	// 1525104000
	GmtCreate *int64 `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// example:
	//
	// 1525104000
	GmtModified *int64 `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s TripBusinessInstanceQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s TripBusinessInstanceQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *TripBusinessInstanceQueryResponseBodyModule) GetCreator() *string {
	return s.Creator
}

func (s *TripBusinessInstanceQueryResponseBodyModule) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *TripBusinessInstanceQueryResponseBodyModule) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *TripBusinessInstanceQueryResponseBodyModule) GetStatus() *string {
	return s.Status
}

func (s *TripBusinessInstanceQueryResponseBodyModule) SetCreator(v string) *TripBusinessInstanceQueryResponseBodyModule {
	s.Creator = &v
	return s
}

func (s *TripBusinessInstanceQueryResponseBodyModule) SetGmtCreate(v int64) *TripBusinessInstanceQueryResponseBodyModule {
	s.GmtCreate = &v
	return s
}

func (s *TripBusinessInstanceQueryResponseBodyModule) SetGmtModified(v int64) *TripBusinessInstanceQueryResponseBodyModule {
	s.GmtModified = &v
	return s
}

func (s *TripBusinessInstanceQueryResponseBodyModule) SetStatus(v string) *TripBusinessInstanceQueryResponseBodyModule {
	s.Status = &v
	return s
}

func (s *TripBusinessInstanceQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
