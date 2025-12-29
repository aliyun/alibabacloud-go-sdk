// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCartoonResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *AddCartoonResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddCartoonResponseBody
	GetRequestId() *string
	SetResult(v bool) *AddCartoonResponseBody
	GetResult() *bool
	SetStatusCode(v int32) *AddCartoonResponseBody
	GetStatusCode() *int32
}

type AddCartoonResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7*726E
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

func (s AddCartoonResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddCartoonResponseBody) GoString() string {
	return s.String()
}

func (s *AddCartoonResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddCartoonResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddCartoonResponseBody) GetResult() *bool {
	return s.Result
}

func (s *AddCartoonResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddCartoonResponseBody) SetMessage(v string) *AddCartoonResponseBody {
	s.Message = &v
	return s
}

func (s *AddCartoonResponseBody) SetRequestId(v string) *AddCartoonResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCartoonResponseBody) SetResult(v bool) *AddCartoonResponseBody {
	s.Result = &v
	return s
}

func (s *AddCartoonResponseBody) SetStatusCode(v int32) *AddCartoonResponseBody {
	s.StatusCode = &v
	return s
}

func (s *AddCartoonResponseBody) Validate() error {
	return dara.Validate(s)
}
