// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchRestartApplicationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchRestartApplicationsResponseBody
	GetCode() *string
	SetData(v *BatchRestartApplicationsResponseBodyData) *BatchRestartApplicationsResponseBody
	GetData() *BatchRestartApplicationsResponseBodyData
	SetErrorCode(v string) *BatchRestartApplicationsResponseBody
	GetErrorCode() *string
	SetMessage(v string) *BatchRestartApplicationsResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchRestartApplicationsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BatchRestartApplicationsResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *BatchRestartApplicationsResponseBody
	GetTraceId() *string
}

type BatchRestartApplicationsResponseBody struct {
	// example:
	//
	// 200
	Code      *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *BatchRestartApplicationsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
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
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s BatchRestartApplicationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchRestartApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchRestartApplicationsResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchRestartApplicationsResponseBody) GetData() *BatchRestartApplicationsResponseBodyData {
	return s.Data
}

func (s *BatchRestartApplicationsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *BatchRestartApplicationsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchRestartApplicationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchRestartApplicationsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchRestartApplicationsResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *BatchRestartApplicationsResponseBody) SetCode(v string) *BatchRestartApplicationsResponseBody {
	s.Code = &v
	return s
}

func (s *BatchRestartApplicationsResponseBody) SetData(v *BatchRestartApplicationsResponseBodyData) *BatchRestartApplicationsResponseBody {
	s.Data = v
	return s
}

func (s *BatchRestartApplicationsResponseBody) SetErrorCode(v string) *BatchRestartApplicationsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *BatchRestartApplicationsResponseBody) SetMessage(v string) *BatchRestartApplicationsResponseBody {
	s.Message = &v
	return s
}

func (s *BatchRestartApplicationsResponseBody) SetRequestId(v string) *BatchRestartApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchRestartApplicationsResponseBody) SetSuccess(v bool) *BatchRestartApplicationsResponseBody {
	s.Success = &v
	return s
}

func (s *BatchRestartApplicationsResponseBody) SetTraceId(v string) *BatchRestartApplicationsResponseBody {
	s.TraceId = &v
	return s
}

func (s *BatchRestartApplicationsResponseBody) Validate() error {
	return dara.Validate(s)
}

type BatchRestartApplicationsResponseBodyData struct {
	// example:
	//
	// 01db03d3-3ee9-48b3-b3d0-dfce2d88****
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s BatchRestartApplicationsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BatchRestartApplicationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchRestartApplicationsResponseBodyData) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *BatchRestartApplicationsResponseBodyData) SetChangeOrderId(v string) *BatchRestartApplicationsResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

func (s *BatchRestartApplicationsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
