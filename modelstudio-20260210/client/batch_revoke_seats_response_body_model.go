// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchRevokeSeatsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchRevokeSeatsResponseBody
	GetCode() *string
	SetMessage(v string) *BatchRevokeSeatsResponseBody
	GetMessage() *string
	SetSuccess(v bool) *BatchRevokeSeatsResponseBody
	GetSuccess() *bool
}

type BatchRevokeSeatsResponseBody struct {
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
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchRevokeSeatsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchRevokeSeatsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchRevokeSeatsResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchRevokeSeatsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchRevokeSeatsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchRevokeSeatsResponseBody) SetCode(v string) *BatchRevokeSeatsResponseBody {
	s.Code = &v
	return s
}

func (s *BatchRevokeSeatsResponseBody) SetMessage(v string) *BatchRevokeSeatsResponseBody {
	s.Message = &v
	return s
}

func (s *BatchRevokeSeatsResponseBody) SetSuccess(v bool) *BatchRevokeSeatsResponseBody {
	s.Success = &v
	return s
}

func (s *BatchRevokeSeatsResponseBody) Validate() error {
	return dara.Validate(s)
}
