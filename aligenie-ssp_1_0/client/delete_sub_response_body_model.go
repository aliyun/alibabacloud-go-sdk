// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSubResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteSubResponseBody
	GetCode() *int32
	SetMessage(v string) *DeleteSubResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteSubResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeleteSubResponseBody
	GetResult() *bool
}

type DeleteSubResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0AA6C15C-FD61-1E32-9881-480CC6F35A70
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteSubResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSubResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSubResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteSubResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteSubResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSubResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteSubResponseBody) SetCode(v int32) *DeleteSubResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSubResponseBody) SetMessage(v string) *DeleteSubResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSubResponseBody) SetRequestId(v string) *DeleteSubResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSubResponseBody) SetResult(v bool) *DeleteSubResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteSubResponseBody) Validate() error {
	return dara.Validate(s)
}
