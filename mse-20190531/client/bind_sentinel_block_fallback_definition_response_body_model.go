// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindSentinelBlockFallbackDefinitionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *BindSentinelBlockFallbackDefinitionResponseBody
	GetCode() *int32
	SetData(v bool) *BindSentinelBlockFallbackDefinitionResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *BindSentinelBlockFallbackDefinitionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *BindSentinelBlockFallbackDefinitionResponseBody
	GetMessage() *string
	SetRequestId(v string) *BindSentinelBlockFallbackDefinitionResponseBody
	GetRequestId() *string
	SetSuccess(v string) *BindSentinelBlockFallbackDefinitionResponseBody
	GetSuccess() *string
}

type BindSentinelBlockFallbackDefinitionResponseBody struct {
	// The status code. A value of 200 is returned if the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4E9FDCFE-0738-493B-B801-82BDFBCB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BindSentinelBlockFallbackDefinitionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindSentinelBlockFallbackDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *BindSentinelBlockFallbackDefinitionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *BindSentinelBlockFallbackDefinitionResponseBody) GetData() *bool {
	return s.Data
}

func (s *BindSentinelBlockFallbackDefinitionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *BindSentinelBlockFallbackDefinitionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BindSentinelBlockFallbackDefinitionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindSentinelBlockFallbackDefinitionResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *BindSentinelBlockFallbackDefinitionResponseBody) SetCode(v int32) *BindSentinelBlockFallbackDefinitionResponseBody {
	s.Code = &v
	return s
}

func (s *BindSentinelBlockFallbackDefinitionResponseBody) SetData(v bool) *BindSentinelBlockFallbackDefinitionResponseBody {
	s.Data = &v
	return s
}

func (s *BindSentinelBlockFallbackDefinitionResponseBody) SetHttpStatusCode(v int32) *BindSentinelBlockFallbackDefinitionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *BindSentinelBlockFallbackDefinitionResponseBody) SetMessage(v string) *BindSentinelBlockFallbackDefinitionResponseBody {
	s.Message = &v
	return s
}

func (s *BindSentinelBlockFallbackDefinitionResponseBody) SetRequestId(v string) *BindSentinelBlockFallbackDefinitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindSentinelBlockFallbackDefinitionResponseBody) SetSuccess(v string) *BindSentinelBlockFallbackDefinitionResponseBody {
	s.Success = &v
	return s
}

func (s *BindSentinelBlockFallbackDefinitionResponseBody) Validate() error {
	return dara.Validate(s)
}
