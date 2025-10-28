// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbortChangeOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AbortChangeOrderResponseBody
	GetCode() *int32
	SetData(v *AbortChangeOrderResponseBodyData) *AbortChangeOrderResponseBody
	GetData() *AbortChangeOrderResponseBodyData
	SetErrorCode(v string) *AbortChangeOrderResponseBody
	GetErrorCode() *string
	SetMessage(v string) *AbortChangeOrderResponseBody
	GetMessage() *string
	SetRequestId(v string) *AbortChangeOrderResponseBody
	GetRequestId() *string
	SetTraceId(v string) *AbortChangeOrderResponseBody
	GetTraceId() *string
}

type AbortChangeOrderResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The change process.
	Data *AbortChangeOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code that is returned.
	//
	// example:
	//
	// success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 57F146F6-3C94-****-****-A66EF4B9*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the trace.
	//
	// example:
	//
	// 0b59000b15947****55688656d****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s AbortChangeOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AbortChangeOrderResponseBody) GoString() string {
	return s.String()
}

func (s *AbortChangeOrderResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AbortChangeOrderResponseBody) GetData() *AbortChangeOrderResponseBodyData {
	return s.Data
}

func (s *AbortChangeOrderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AbortChangeOrderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AbortChangeOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AbortChangeOrderResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *AbortChangeOrderResponseBody) SetCode(v int32) *AbortChangeOrderResponseBody {
	s.Code = &v
	return s
}

func (s *AbortChangeOrderResponseBody) SetData(v *AbortChangeOrderResponseBodyData) *AbortChangeOrderResponseBody {
	s.Data = v
	return s
}

func (s *AbortChangeOrderResponseBody) SetErrorCode(v string) *AbortChangeOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AbortChangeOrderResponseBody) SetMessage(v string) *AbortChangeOrderResponseBody {
	s.Message = &v
	return s
}

func (s *AbortChangeOrderResponseBody) SetRequestId(v string) *AbortChangeOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *AbortChangeOrderResponseBody) SetTraceId(v string) *AbortChangeOrderResponseBody {
	s.TraceId = &v
	return s
}

func (s *AbortChangeOrderResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AbortChangeOrderResponseBodyData struct {
	// The ID of the change process.
	//
	// example:
	//
	// 4f038ddf-b27b-****-****-88e44375****
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s AbortChangeOrderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AbortChangeOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *AbortChangeOrderResponseBodyData) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *AbortChangeOrderResponseBodyData) SetChangeOrderId(v string) *AbortChangeOrderResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

func (s *AbortChangeOrderResponseBodyData) Validate() error {
	return dara.Validate(s)
}
