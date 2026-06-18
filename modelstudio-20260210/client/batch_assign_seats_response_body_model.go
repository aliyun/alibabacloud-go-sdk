// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchAssignSeatsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchAssignSeatsResponseBody
	GetCode() *string
	SetMessage(v string) *BatchAssignSeatsResponseBody
	GetMessage() *string
	SetSuccess(v bool) *BatchAssignSeatsResponseBody
	GetSuccess() *bool
}

type BatchAssignSeatsResponseBody struct {
	// The error code. This parameter is empty if the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message. This parameter is empty if the request was successful.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchAssignSeatsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchAssignSeatsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchAssignSeatsResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchAssignSeatsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchAssignSeatsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchAssignSeatsResponseBody) SetCode(v string) *BatchAssignSeatsResponseBody {
	s.Code = &v
	return s
}

func (s *BatchAssignSeatsResponseBody) SetMessage(v string) *BatchAssignSeatsResponseBody {
	s.Message = &v
	return s
}

func (s *BatchAssignSeatsResponseBody) SetSuccess(v bool) *BatchAssignSeatsResponseBody {
	s.Success = &v
	return s
}

func (s *BatchAssignSeatsResponseBody) Validate() error {
	return dara.Validate(s)
}
