// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSwimmingLaneGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteSwimmingLaneGroupResponseBody
	GetCode() *string
	SetErrorCode(v string) *DeleteSwimmingLaneGroupResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DeleteSwimmingLaneGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteSwimmingLaneGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteSwimmingLaneGroupResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DeleteSwimmingLaneGroupResponseBody
	GetTraceId() *string
}

type DeleteSwimmingLaneGroupResponseBody struct {
	// example:
	//
	// 200
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// ac1a0b2215622920113732501e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DeleteSwimmingLaneGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSwimmingLaneGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSwimmingLaneGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteSwimmingLaneGroupResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteSwimmingLaneGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteSwimmingLaneGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSwimmingLaneGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteSwimmingLaneGroupResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DeleteSwimmingLaneGroupResponseBody) SetCode(v string) *DeleteSwimmingLaneGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSwimmingLaneGroupResponseBody) SetErrorCode(v string) *DeleteSwimmingLaneGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteSwimmingLaneGroupResponseBody) SetMessage(v string) *DeleteSwimmingLaneGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSwimmingLaneGroupResponseBody) SetRequestId(v string) *DeleteSwimmingLaneGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSwimmingLaneGroupResponseBody) SetSuccess(v bool) *DeleteSwimmingLaneGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSwimmingLaneGroupResponseBody) SetTraceId(v string) *DeleteSwimmingLaneGroupResponseBody {
	s.TraceId = &v
	return s
}

func (s *DeleteSwimmingLaneGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
