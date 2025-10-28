// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbortAndRollbackChangeOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AbortAndRollbackChangeOrderResponseBody
	GetCode() *int32
	SetData(v *AbortAndRollbackChangeOrderResponseBodyData) *AbortAndRollbackChangeOrderResponseBody
	GetData() *AbortAndRollbackChangeOrderResponseBodyData
	SetErrorCode(v string) *AbortAndRollbackChangeOrderResponseBody
	GetErrorCode() *string
	SetMessage(v string) *AbortAndRollbackChangeOrderResponseBody
	GetMessage() *string
	SetRequestId(v string) *AbortAndRollbackChangeOrderResponseBody
	GetRequestId() *string
	SetTraceId(v string) *AbortAndRollbackChangeOrderResponseBody
	GetTraceId() *string
}

type AbortAndRollbackChangeOrderResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the change process.
	Data *AbortAndRollbackChangeOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 81E0B333-2871-****-****-B8F5FF43****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the trace.
	//
	// example:
	//
	// 210f07bf1640239405712621******
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s AbortAndRollbackChangeOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AbortAndRollbackChangeOrderResponseBody) GoString() string {
	return s.String()
}

func (s *AbortAndRollbackChangeOrderResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AbortAndRollbackChangeOrderResponseBody) GetData() *AbortAndRollbackChangeOrderResponseBodyData {
	return s.Data
}

func (s *AbortAndRollbackChangeOrderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AbortAndRollbackChangeOrderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AbortAndRollbackChangeOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AbortAndRollbackChangeOrderResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *AbortAndRollbackChangeOrderResponseBody) SetCode(v int32) *AbortAndRollbackChangeOrderResponseBody {
	s.Code = &v
	return s
}

func (s *AbortAndRollbackChangeOrderResponseBody) SetData(v *AbortAndRollbackChangeOrderResponseBodyData) *AbortAndRollbackChangeOrderResponseBody {
	s.Data = v
	return s
}

func (s *AbortAndRollbackChangeOrderResponseBody) SetErrorCode(v string) *AbortAndRollbackChangeOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AbortAndRollbackChangeOrderResponseBody) SetMessage(v string) *AbortAndRollbackChangeOrderResponseBody {
	s.Message = &v
	return s
}

func (s *AbortAndRollbackChangeOrderResponseBody) SetRequestId(v string) *AbortAndRollbackChangeOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *AbortAndRollbackChangeOrderResponseBody) SetTraceId(v string) *AbortAndRollbackChangeOrderResponseBody {
	s.TraceId = &v
	return s
}

func (s *AbortAndRollbackChangeOrderResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AbortAndRollbackChangeOrderResponseBodyData struct {
	// The ID of the change process.
	//
	// example:
	//
	// 4f40e616-cdcd-4250-a018-efd459******
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s AbortAndRollbackChangeOrderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AbortAndRollbackChangeOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *AbortAndRollbackChangeOrderResponseBodyData) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *AbortAndRollbackChangeOrderResponseBodyData) SetChangeOrderId(v string) *AbortAndRollbackChangeOrderResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

func (s *AbortAndRollbackChangeOrderResponseBodyData) Validate() error {
	return dara.Validate(s)
}
