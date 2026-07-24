// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCallOutboundInstantResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateCallOutboundInstantResponseBody
	GetCode() *string
	SetData(v *CreateCallOutboundInstantResponseBodyData) *CreateCallOutboundInstantResponseBody
	GetData() *CreateCallOutboundInstantResponseBodyData
	SetMessage(v string) *CreateCallOutboundInstantResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateCallOutboundInstantResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateCallOutboundInstantResponseBody
	GetSuccess() *bool
	SetTimestamp(v string) *CreateCallOutboundInstantResponseBody
	GetTimestamp() *string
	SetTraceId(v string) *CreateCallOutboundInstantResponseBody
	GetTraceId() *string
}

type CreateCallOutboundInstantResponseBody struct {
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateCallOutboundInstantResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	Timestamp *string                                    `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	TraceId   *string                                    `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s CreateCallOutboundInstantResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCallOutboundInstantResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCallOutboundInstantResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateCallOutboundInstantResponseBody) GetData() *CreateCallOutboundInstantResponseBodyData {
	return s.Data
}

func (s *CreateCallOutboundInstantResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateCallOutboundInstantResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCallOutboundInstantResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateCallOutboundInstantResponseBody) GetTimestamp() *string {
	return s.Timestamp
}

func (s *CreateCallOutboundInstantResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *CreateCallOutboundInstantResponseBody) SetCode(v string) *CreateCallOutboundInstantResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCallOutboundInstantResponseBody) SetData(v *CreateCallOutboundInstantResponseBodyData) *CreateCallOutboundInstantResponseBody {
	s.Data = v
	return s
}

func (s *CreateCallOutboundInstantResponseBody) SetMessage(v string) *CreateCallOutboundInstantResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCallOutboundInstantResponseBody) SetRequestId(v string) *CreateCallOutboundInstantResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCallOutboundInstantResponseBody) SetSuccess(v bool) *CreateCallOutboundInstantResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCallOutboundInstantResponseBody) SetTimestamp(v string) *CreateCallOutboundInstantResponseBody {
	s.Timestamp = &v
	return s
}

func (s *CreateCallOutboundInstantResponseBody) SetTraceId(v string) *CreateCallOutboundInstantResponseBody {
	s.TraceId = &v
	return s
}

func (s *CreateCallOutboundInstantResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCallOutboundInstantResponseBodyData struct {
	// example:
	//
	// call_123456
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// example:
	//
	// CALLING
	CallStatus *string `json:"CallStatus,omitempty" xml:"CallStatus,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9B19DF17-D5C4-5893-B8FF-4EBAADCDC6A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCallOutboundInstantResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateCallOutboundInstantResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateCallOutboundInstantResponseBodyData) GetCallId() *string {
	return s.CallId
}

func (s *CreateCallOutboundInstantResponseBodyData) GetCallStatus() *string {
	return s.CallStatus
}

func (s *CreateCallOutboundInstantResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *CreateCallOutboundInstantResponseBodyData) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCallOutboundInstantResponseBodyData) SetCallId(v string) *CreateCallOutboundInstantResponseBodyData {
	s.CallId = &v
	return s
}

func (s *CreateCallOutboundInstantResponseBodyData) SetCallStatus(v string) *CreateCallOutboundInstantResponseBodyData {
	s.CallStatus = &v
	return s
}

func (s *CreateCallOutboundInstantResponseBodyData) SetMessage(v string) *CreateCallOutboundInstantResponseBodyData {
	s.Message = &v
	return s
}

func (s *CreateCallOutboundInstantResponseBodyData) SetRequestId(v string) *CreateCallOutboundInstantResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *CreateCallOutboundInstantResponseBodyData) Validate() error {
	return dara.Validate(s)
}
