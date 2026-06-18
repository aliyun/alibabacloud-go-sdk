// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddHotlineNumberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddHotlineNumberResponseBody
	GetCode() *string
	SetHttpStatusCode(v int64) *AddHotlineNumberResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *AddHotlineNumberResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddHotlineNumberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddHotlineNumberResponseBody
	GetSuccess() *bool
}

type AddHotlineNumberResponseBody struct {
	// The status code. A value of "Success" indicates that the request succeeded.
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
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EE339D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddHotlineNumberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddHotlineNumberResponseBody) GoString() string {
	return s.String()
}

func (s *AddHotlineNumberResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddHotlineNumberResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *AddHotlineNumberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddHotlineNumberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddHotlineNumberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddHotlineNumberResponseBody) SetCode(v string) *AddHotlineNumberResponseBody {
	s.Code = &v
	return s
}

func (s *AddHotlineNumberResponseBody) SetHttpStatusCode(v int64) *AddHotlineNumberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddHotlineNumberResponseBody) SetMessage(v string) *AddHotlineNumberResponseBody {
	s.Message = &v
	return s
}

func (s *AddHotlineNumberResponseBody) SetRequestId(v string) *AddHotlineNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddHotlineNumberResponseBody) SetSuccess(v bool) *AddHotlineNumberResponseBody {
	s.Success = &v
	return s
}

func (s *AddHotlineNumberResponseBody) Validate() error {
	return dara.Validate(s)
}
