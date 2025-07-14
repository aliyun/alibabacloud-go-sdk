// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteInstancesResponseBody
	GetCode() *string
	SetData(v *DeleteInstancesResponseBodyData) *DeleteInstancesResponseBody
	GetData() *DeleteInstancesResponseBodyData
	SetErrorCode(v string) *DeleteInstancesResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DeleteInstancesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteInstancesResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DeleteInstancesResponseBody
	GetTraceId() *string
}

type DeleteInstancesResponseBody struct {
	// example:
	//
	// 200
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DeleteInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                          `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
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
	// 0a981dd515966966104121683d****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DeleteInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstancesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteInstancesResponseBody) GetData() *DeleteInstancesResponseBodyData {
	return s.Data
}

func (s *DeleteInstancesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteInstancesResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DeleteInstancesResponseBody) SetCode(v string) *DeleteInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteInstancesResponseBody) SetData(v *DeleteInstancesResponseBodyData) *DeleteInstancesResponseBody {
	s.Data = v
	return s
}

func (s *DeleteInstancesResponseBody) SetErrorCode(v string) *DeleteInstancesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteInstancesResponseBody) SetMessage(v string) *DeleteInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteInstancesResponseBody) SetRequestId(v string) *DeleteInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstancesResponseBody) SetSuccess(v bool) *DeleteInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteInstancesResponseBody) SetTraceId(v string) *DeleteInstancesResponseBody {
	s.TraceId = &v
	return s
}

func (s *DeleteInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DeleteInstancesResponseBodyData struct {
	// example:
	//
	// 01db03d3-3ee9-48b3-b3d0-dfce2d88****
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s DeleteInstancesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteInstancesResponseBodyData) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *DeleteInstancesResponseBodyData) SetChangeOrderId(v string) *DeleteInstancesResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

func (s *DeleteInstancesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
