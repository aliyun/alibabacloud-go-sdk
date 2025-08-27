// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncSingleUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SyncSingleUserResponseBody
	GetCode() *string
	SetMessage(v string) *SyncSingleUserResponseBody
	GetMessage() *string
	SetModule(v string) *SyncSingleUserResponseBody
	GetModule() *string
	SetRequestId(v string) *SyncSingleUserResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SyncSingleUserResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *SyncSingleUserResponseBody
	GetTraceId() *string
}

type SyncSingleUserResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// success
	Module *string `json:"module,omitempty" xml:"module,omitempty"`
	// example:
	//
	// A5009956-1077-52FB-B520-EA8C7E91D722
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 210f079e16603757182131635d866a
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s SyncSingleUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SyncSingleUserResponseBody) GoString() string {
	return s.String()
}

func (s *SyncSingleUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *SyncSingleUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SyncSingleUserResponseBody) GetModule() *string {
	return s.Module
}

func (s *SyncSingleUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SyncSingleUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SyncSingleUserResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *SyncSingleUserResponseBody) SetCode(v string) *SyncSingleUserResponseBody {
	s.Code = &v
	return s
}

func (s *SyncSingleUserResponseBody) SetMessage(v string) *SyncSingleUserResponseBody {
	s.Message = &v
	return s
}

func (s *SyncSingleUserResponseBody) SetModule(v string) *SyncSingleUserResponseBody {
	s.Module = &v
	return s
}

func (s *SyncSingleUserResponseBody) SetRequestId(v string) *SyncSingleUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncSingleUserResponseBody) SetSuccess(v bool) *SyncSingleUserResponseBody {
	s.Success = &v
	return s
}

func (s *SyncSingleUserResponseBody) SetTraceId(v string) *SyncSingleUserResponseBody {
	s.TraceId = &v
	return s
}

func (s *SyncSingleUserResponseBody) Validate() error {
	return dara.Validate(s)
}
