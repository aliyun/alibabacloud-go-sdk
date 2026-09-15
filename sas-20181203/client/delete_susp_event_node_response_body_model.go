// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSuspEventNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteSuspEventNodeResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteSuspEventNodeResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteSuspEventNodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteSuspEventNodeResponseBody
	GetSuccess() *bool
}

type DeleteSuspEventNodeResponseBody struct {
	// The result code. A value of **200*	- indicates success. Any other value indicates failure. You can use this field to determine the cause of a failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique request ID.
	//
	// example:
	//
	// A57C711B-AA15-55B2-8F61-4D09CEXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSuspEventNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSuspEventNodeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSuspEventNodeResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteSuspEventNodeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteSuspEventNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSuspEventNodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteSuspEventNodeResponseBody) SetCode(v string) *DeleteSuspEventNodeResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSuspEventNodeResponseBody) SetMessage(v string) *DeleteSuspEventNodeResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSuspEventNodeResponseBody) SetRequestId(v string) *DeleteSuspEventNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSuspEventNodeResponseBody) SetSuccess(v bool) *DeleteSuspEventNodeResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSuspEventNodeResponseBody) Validate() error {
	return dara.Validate(s)
}
