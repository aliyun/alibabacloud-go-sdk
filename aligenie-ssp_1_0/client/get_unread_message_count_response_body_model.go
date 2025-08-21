// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUnreadMessageCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetUnreadMessageCountResponseBody
	GetCode() *string
	SetMessage(v string) *GetUnreadMessageCountResponseBody
	GetMessage() *string
	SetResult(v int32) *GetUnreadMessageCountResponseBody
	GetResult() *int32
}

type GetUnreadMessageCountResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 10
	Result *int32 `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s GetUnreadMessageCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUnreadMessageCountResponseBody) GoString() string {
	return s.String()
}

func (s *GetUnreadMessageCountResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetUnreadMessageCountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetUnreadMessageCountResponseBody) GetResult() *int32 {
	return s.Result
}

func (s *GetUnreadMessageCountResponseBody) SetCode(v string) *GetUnreadMessageCountResponseBody {
	s.Code = &v
	return s
}

func (s *GetUnreadMessageCountResponseBody) SetMessage(v string) *GetUnreadMessageCountResponseBody {
	s.Message = &v
	return s
}

func (s *GetUnreadMessageCountResponseBody) SetResult(v int32) *GetUnreadMessageCountResponseBody {
	s.Result = &v
	return s
}

func (s *GetUnreadMessageCountResponseBody) Validate() error {
	return dara.Validate(s)
}
