// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateTagResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateTagResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateTagResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateTagResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateTagResponseBody
	GetSuccess() *bool
	SetTagId(v string) *CreateTagResponseBody
	GetTagId() *string
}

type CreateTagResponseBody struct {
	// API status code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// API response message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// Success
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Tag ID
	//
	// example:
	//
	// 390515b5-6115-4ccf-83e2-52d5bfaf2ddf
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s CreateTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTagResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTagResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateTagResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateTagResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTagResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateTagResponseBody) GetTagId() *string {
	return s.TagId
}

func (s *CreateTagResponseBody) SetCode(v string) *CreateTagResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTagResponseBody) SetHttpStatusCode(v int32) *CreateTagResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateTagResponseBody) SetMessage(v string) *CreateTagResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTagResponseBody) SetRequestId(v string) *CreateTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTagResponseBody) SetSuccess(v bool) *CreateTagResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTagResponseBody) SetTagId(v string) *CreateTagResponseBody {
	s.TagId = &v
	return s
}

func (s *CreateTagResponseBody) Validate() error {
	return dara.Validate(s)
}
