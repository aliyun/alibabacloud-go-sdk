// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAutomateResponseConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteAutomateResponseConfigResponseBody
	GetCode() *int32
	SetData(v string) *DeleteAutomateResponseConfigResponseBody
	GetData() *string
	SetMessage(v string) *DeleteAutomateResponseConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAutomateResponseConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteAutomateResponseConfigResponseBody
	GetSuccess() *bool
}

type DeleteAutomateResponseConfigResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAutomateResponseConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAutomateResponseConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAutomateResponseConfigResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteAutomateResponseConfigResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteAutomateResponseConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAutomateResponseConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAutomateResponseConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAutomateResponseConfigResponseBody) SetCode(v int32) *DeleteAutomateResponseConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAutomateResponseConfigResponseBody) SetData(v string) *DeleteAutomateResponseConfigResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteAutomateResponseConfigResponseBody) SetMessage(v string) *DeleteAutomateResponseConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAutomateResponseConfigResponseBody) SetRequestId(v string) *DeleteAutomateResponseConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAutomateResponseConfigResponseBody) SetSuccess(v bool) *DeleteAutomateResponseConfigResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAutomateResponseConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
