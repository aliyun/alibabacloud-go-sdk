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
	// The API status code or POP error code. Valid values:
	//
	// - 2xx: The request is successful.
	//
	// - 3xx: The request is redirected.
	//
	// - 4xx: A request error occurred.
	//
	// - 5xx: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	Data *BatchRestartApplicationsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code. Valid values:
	//
	// - Successful request: The ErrorCode field is not returned.
	//
	// - Failed request: The ErrorCode field is returned. For more information, see the error codes section in this topic.
	//
	// example:
	//
	// System.Upgrading
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional information. Valid values:
	//
	// - success: The request is successful.
	//
	// - Specific error code: The request is abnormal and a specific error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the applications are successfully restarted in batches. Valid values:
	//
	// - true: The applications are successfully restarted in batches.
	//
	// - false: The applications fail to be restarted in batches.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID used to accurately query call information.
	//
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchRestartApplicationsResponseBodyData struct {
	// The change order ID.
	//
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
