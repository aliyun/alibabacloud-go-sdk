// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReopenInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ReopenInstanceResponseBody
	GetCode() *int32
	SetMessage(v string) *ReopenInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReopenInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReopenInstanceResponseBody
	GetSuccess() *bool
}

type ReopenInstanceResponseBody struct {
	// The returned HTTP status code. If the request is successful, 200 is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 252820E1-A2E6-45F2-B4C9-1056B8CE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReopenInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReopenInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ReopenInstanceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ReopenInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReopenInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReopenInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReopenInstanceResponseBody) SetCode(v int32) *ReopenInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *ReopenInstanceResponseBody) SetMessage(v string) *ReopenInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *ReopenInstanceResponseBody) SetRequestId(v string) *ReopenInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReopenInstanceResponseBody) SetSuccess(v bool) *ReopenInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *ReopenInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
