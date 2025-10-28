// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackChangeOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *RollbackChangeOrderResponseBody
	GetCode() *int32
	SetData(v *RollbackChangeOrderResponseBodyData) *RollbackChangeOrderResponseBody
	GetData() *RollbackChangeOrderResponseBodyData
	SetErrorCode(v string) *RollbackChangeOrderResponseBody
	GetErrorCode() *string
	SetMessage(v string) *RollbackChangeOrderResponseBody
	GetMessage() *string
	SetRequestId(v string) *RollbackChangeOrderResponseBody
	GetRequestId() *string
	SetTraceId(v string) *RollbackChangeOrderResponseBody
	GetTraceId() *string
}

type RollbackChangeOrderResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the change process.
	Data *RollbackChangeOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// B909AB1F-3763-4963-B1CE-0BDFA192****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the trace.
	//
	// example:
	//
	// 000000000000000000000000000000
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RollbackChangeOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RollbackChangeOrderResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackChangeOrderResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *RollbackChangeOrderResponseBody) GetData() *RollbackChangeOrderResponseBodyData {
	return s.Data
}

func (s *RollbackChangeOrderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RollbackChangeOrderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RollbackChangeOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RollbackChangeOrderResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *RollbackChangeOrderResponseBody) SetCode(v int32) *RollbackChangeOrderResponseBody {
	s.Code = &v
	return s
}

func (s *RollbackChangeOrderResponseBody) SetData(v *RollbackChangeOrderResponseBodyData) *RollbackChangeOrderResponseBody {
	s.Data = v
	return s
}

func (s *RollbackChangeOrderResponseBody) SetErrorCode(v string) *RollbackChangeOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RollbackChangeOrderResponseBody) SetMessage(v string) *RollbackChangeOrderResponseBody {
	s.Message = &v
	return s
}

func (s *RollbackChangeOrderResponseBody) SetRequestId(v string) *RollbackChangeOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *RollbackChangeOrderResponseBody) SetTraceId(v string) *RollbackChangeOrderResponseBody {
	s.TraceId = &v
	return s
}

func (s *RollbackChangeOrderResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RollbackChangeOrderResponseBodyData struct {
	// The ID of the change process.
	//
	// example:
	//
	// 4f40e616-cdcd-4250-a018-efd4599c****
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s RollbackChangeOrderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RollbackChangeOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *RollbackChangeOrderResponseBodyData) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *RollbackChangeOrderResponseBodyData) SetChangeOrderId(v string) *RollbackChangeOrderResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

func (s *RollbackChangeOrderResponseBodyData) Validate() error {
	return dara.Validate(s)
}
