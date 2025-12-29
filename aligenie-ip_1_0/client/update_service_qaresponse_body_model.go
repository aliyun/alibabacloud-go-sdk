// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceQAResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *UpdateServiceQAResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateServiceQAResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateServiceQAResponseBody
	GetResult() *bool
	SetStatusCode(v int32) *UpdateServiceQAResponseBody
	GetStatusCode() *int32
}

type UpdateServiceQAResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 73C67***6FA
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

func (s UpdateServiceQAResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceQAResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceQAResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateServiceQAResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateServiceQAResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateServiceQAResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateServiceQAResponseBody) SetMessage(v string) *UpdateServiceQAResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateServiceQAResponseBody) SetRequestId(v string) *UpdateServiceQAResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateServiceQAResponseBody) SetResult(v bool) *UpdateServiceQAResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateServiceQAResponseBody) SetStatusCode(v int32) *UpdateServiceQAResponseBody {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceQAResponseBody) Validate() error {
	return dara.Validate(s)
}
