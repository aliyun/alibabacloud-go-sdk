// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCurrentConcurrencyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCurrentConcurrencyResponseBody
	GetCode() *string
	SetCurrentConcurrency(v int32) *GetCurrentConcurrencyResponseBody
	GetCurrentConcurrency() *int32
	SetHttpStatusCode(v int32) *GetCurrentConcurrencyResponseBody
	GetHttpStatusCode() *int32
	SetInstanceId(v string) *GetCurrentConcurrencyResponseBody
	GetInstanceId() *string
	SetMaxConcurrentConversation(v int32) *GetCurrentConcurrencyResponseBody
	GetMaxConcurrentConversation() *int32
	SetMessage(v string) *GetCurrentConcurrencyResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCurrentConcurrencyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCurrentConcurrencyResponseBody
	GetSuccess() *bool
}

type GetCurrentConcurrencyResponseBody struct {
	// API status code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Current concurrency
	//
	// example:
	//
	// 12
	CurrentConcurrency *int32 `json:"CurrentConcurrency,omitempty" xml:"CurrentConcurrency,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Instance ID
	//
	// example:
	//
	// 12f3dd08-0c55-44ce-9b64-e69d35ed3a76
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Maximum concurrency
	//
	// example:
	//
	// 5
	MaxConcurrentConversation *int32 `json:"MaxConcurrentConversation,omitempty" xml:"MaxConcurrentConversation,omitempty"`
	// API message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCurrentConcurrencyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCurrentConcurrencyResponseBody) GoString() string {
	return s.String()
}

func (s *GetCurrentConcurrencyResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCurrentConcurrencyResponseBody) GetCurrentConcurrency() *int32 {
	return s.CurrentConcurrency
}

func (s *GetCurrentConcurrencyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetCurrentConcurrencyResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetCurrentConcurrencyResponseBody) GetMaxConcurrentConversation() *int32 {
	return s.MaxConcurrentConversation
}

func (s *GetCurrentConcurrencyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCurrentConcurrencyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCurrentConcurrencyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCurrentConcurrencyResponseBody) SetCode(v string) *GetCurrentConcurrencyResponseBody {
	s.Code = &v
	return s
}

func (s *GetCurrentConcurrencyResponseBody) SetCurrentConcurrency(v int32) *GetCurrentConcurrencyResponseBody {
	s.CurrentConcurrency = &v
	return s
}

func (s *GetCurrentConcurrencyResponseBody) SetHttpStatusCode(v int32) *GetCurrentConcurrencyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetCurrentConcurrencyResponseBody) SetInstanceId(v string) *GetCurrentConcurrencyResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetCurrentConcurrencyResponseBody) SetMaxConcurrentConversation(v int32) *GetCurrentConcurrencyResponseBody {
	s.MaxConcurrentConversation = &v
	return s
}

func (s *GetCurrentConcurrencyResponseBody) SetMessage(v string) *GetCurrentConcurrencyResponseBody {
	s.Message = &v
	return s
}

func (s *GetCurrentConcurrencyResponseBody) SetRequestId(v string) *GetCurrentConcurrencyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCurrentConcurrencyResponseBody) SetSuccess(v bool) *GetCurrentConcurrencyResponseBody {
	s.Success = &v
	return s
}

func (s *GetCurrentConcurrencyResponseBody) Validate() error {
	return dara.Validate(s)
}
