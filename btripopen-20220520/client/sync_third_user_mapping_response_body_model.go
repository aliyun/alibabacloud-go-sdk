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
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
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
