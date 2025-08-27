// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncThirdUserMappingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SyncThirdUserMappingResponseBody
	GetCode() *string
	SetMessage(v string) *SyncThirdUserMappingResponseBody
	GetMessage() *string
	SetRequestId(v string) *SyncThirdUserMappingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SyncThirdUserMappingResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *SyncThirdUserMappingResponseBody
	GetTraceId() *string
}

type SyncThirdUserMappingResponseBody struct {
	// example:
	//
	// 0
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// C61ECFF6-606B-5F66-B81D-D77369043A5F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 8465F68D-BC97-5C0F-9161-3E65919D9135
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s SyncThirdUserMappingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SyncThirdUserMappingResponseBody) GoString() string {
	return s.String()
}

func (s *SyncThirdUserMappingResponseBody) GetCode() *string {
	return s.Code
}

func (s *SyncThirdUserMappingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SyncThirdUserMappingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SyncThirdUserMappingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SyncThirdUserMappingResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *SyncThirdUserMappingResponseBody) SetCode(v string) *SyncThirdUserMappingResponseBody {
	s.Code = &v
	return s
}

func (s *SyncThirdUserMappingResponseBody) SetMessage(v string) *SyncThirdUserMappingResponseBody {
	s.Message = &v
	return s
}

func (s *SyncThirdUserMappingResponseBody) SetRequestId(v string) *SyncThirdUserMappingResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncThirdUserMappingResponseBody) SetSuccess(v bool) *SyncThirdUserMappingResponseBody {
	s.Success = &v
	return s
}

func (s *SyncThirdUserMappingResponseBody) SetTraceId(v string) *SyncThirdUserMappingResponseBody {
	s.TraceId = &v
	return s
}

func (s *SyncThirdUserMappingResponseBody) Validate() error {
	return dara.Validate(s)
}
