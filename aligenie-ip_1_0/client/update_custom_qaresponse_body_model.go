// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomQAResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *UpdateCustomQAResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateCustomQAResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateCustomQAResponseBody
	GetResult() *bool
	SetStatusCode(v int32) *UpdateCustomQAResponseBody
	GetStatusCode() *int32
}

type UpdateCustomQAResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 73C6***BB3E6FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s UpdateCustomQAResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomQAResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomQAResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateCustomQAResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCustomQAResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateCustomQAResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCustomQAResponseBody) SetMessage(v string) *UpdateCustomQAResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateCustomQAResponseBody) SetRequestId(v string) *UpdateCustomQAResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCustomQAResponseBody) SetResult(v bool) *UpdateCustomQAResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateCustomQAResponseBody) SetStatusCode(v int32) *UpdateCustomQAResponseBody {
	s.StatusCode = &v
	return s
}

func (s *UpdateCustomQAResponseBody) Validate() error {
	return dara.Validate(s)
}
