// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetHotlineNumberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ResetHotlineNumberResponseBody
	GetCode() *string
	SetHttpStatusCode(v int64) *ResetHotlineNumberResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *ResetHotlineNumberResponseBody
	GetMessage() *string
	SetRequestId(v string) *ResetHotlineNumberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ResetHotlineNumberResponseBody
	GetSuccess() *bool
}

type ResetHotlineNumberResponseBody struct {
	// The status code. A return value of "Success" indicates that the request succeeded.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Description of the status code.
	//
	// example:
	//
	// xxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ResetHotlineNumberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetHotlineNumberResponseBody) GoString() string {
	return s.String()
}

func (s *ResetHotlineNumberResponseBody) GetCode() *string {
	return s.Code
}

func (s *ResetHotlineNumberResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *ResetHotlineNumberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ResetHotlineNumberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetHotlineNumberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ResetHotlineNumberResponseBody) SetCode(v string) *ResetHotlineNumberResponseBody {
	s.Code = &v
	return s
}

func (s *ResetHotlineNumberResponseBody) SetHttpStatusCode(v int64) *ResetHotlineNumberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ResetHotlineNumberResponseBody) SetMessage(v string) *ResetHotlineNumberResponseBody {
	s.Message = &v
	return s
}

func (s *ResetHotlineNumberResponseBody) SetRequestId(v string) *ResetHotlineNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetHotlineNumberResponseBody) SetSuccess(v bool) *ResetHotlineNumberResponseBody {
	s.Success = &v
	return s
}

func (s *ResetHotlineNumberResponseBody) Validate() error {
	return dara.Validate(s)
}
