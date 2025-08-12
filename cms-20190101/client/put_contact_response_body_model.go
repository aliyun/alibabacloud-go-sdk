// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutContactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PutContactResponseBody
	GetCode() *string
	SetMessage(v string) *PutContactResponseBody
	GetMessage() *string
	SetRequestId(v string) *PutContactResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PutContactResponseBody
	GetSuccess() *bool
}

type PutContactResponseBody struct {
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
	// The Request is not authorization.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 181C406E-9DE4-484C-9C61-37AE9A1A12EE
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

func (s PutContactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutContactResponseBody) GoString() string {
	return s.String()
}

func (s *PutContactResponseBody) GetCode() *string {
	return s.Code
}

func (s *PutContactResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PutContactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutContactResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PutContactResponseBody) SetCode(v string) *PutContactResponseBody {
	s.Code = &v
	return s
}

func (s *PutContactResponseBody) SetMessage(v string) *PutContactResponseBody {
	s.Message = &v
	return s
}

func (s *PutContactResponseBody) SetRequestId(v string) *PutContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutContactResponseBody) SetSuccess(v bool) *PutContactResponseBody {
	s.Success = &v
	return s
}

func (s *PutContactResponseBody) Validate() error {
	return dara.Validate(s)
}
