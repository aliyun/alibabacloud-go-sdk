// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutContactGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PutContactGroupResponseBody
	GetCode() *string
	SetMessage(v string) *PutContactGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *PutContactGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PutContactGroupResponseBody
	GetSuccess() *bool
}

type PutContactGroupResponseBody struct {
	// The status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// Illegal parameters.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B4E30DB6-F069-5D0B-A589-2A89F7D62A57
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

func (s PutContactGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutContactGroupResponseBody) GoString() string {
	return s.String()
}

func (s *PutContactGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *PutContactGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PutContactGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutContactGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PutContactGroupResponseBody) SetCode(v string) *PutContactGroupResponseBody {
	s.Code = &v
	return s
}

func (s *PutContactGroupResponseBody) SetMessage(v string) *PutContactGroupResponseBody {
	s.Message = &v
	return s
}

func (s *PutContactGroupResponseBody) SetRequestId(v string) *PutContactGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutContactGroupResponseBody) SetSuccess(v bool) *PutContactGroupResponseBody {
	s.Success = &v
	return s
}

func (s *PutContactGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
