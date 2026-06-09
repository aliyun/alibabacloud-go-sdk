// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteTableResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteTableResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteTableResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteTableResponseBody
	GetSuccess() *bool
}

type DeleteTableResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 34AD682D-5B91-5773-8132-AA38C130****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTableResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTableResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteTableResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTableResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteTableResponseBody) SetCode(v string) *DeleteTableResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteTableResponseBody) SetMessage(v string) *DeleteTableResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteTableResponseBody) SetRequestId(v string) *DeleteTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTableResponseBody) SetSuccess(v bool) *DeleteTableResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteTableResponseBody) Validate() error {
	return dara.Validate(s)
}
