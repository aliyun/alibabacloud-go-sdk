// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteContactGroupResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteContactGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteContactGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteContactGroupResponseBody
	GetSuccess() *bool
}

type DeleteContactGroupResponseBody struct {
	// The status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
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
	// F722BE59-2400-4A64-9C1A-AD886AED9A69
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteContactGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteContactGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteContactGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteContactGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteContactGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteContactGroupResponseBody) SetCode(v string) *DeleteContactGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteContactGroupResponseBody) SetMessage(v string) *DeleteContactGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteContactGroupResponseBody) SetRequestId(v string) *DeleteContactGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteContactGroupResponseBody) SetSuccess(v bool) *DeleteContactGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteContactGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
