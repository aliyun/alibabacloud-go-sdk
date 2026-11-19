// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMemoryRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteMemoryRuleResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteMemoryRuleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteMemoryRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteMemoryRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteMemoryRuleResponseBody
	GetSuccess() *bool
}

type DeleteMemoryRuleResponseBody struct {
	// code
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// httpStatusCode
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// message
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// requestId
	//
	// example:
	//
	// 01A00000-0000-0000-0000-000000000001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// success
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteMemoryRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMemoryRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMemoryRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteMemoryRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteMemoryRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteMemoryRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMemoryRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMemoryRuleResponseBody) SetCode(v string) *DeleteMemoryRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMemoryRuleResponseBody) SetHttpStatusCode(v int32) *DeleteMemoryRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteMemoryRuleResponseBody) SetMessage(v string) *DeleteMemoryRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMemoryRuleResponseBody) SetRequestId(v string) *DeleteMemoryRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMemoryRuleResponseBody) SetSuccess(v bool) *DeleteMemoryRuleResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMemoryRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
