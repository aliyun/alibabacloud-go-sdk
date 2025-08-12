// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDynamicTagGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteDynamicTagGroupResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteDynamicTagGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteDynamicTagGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDynamicTagGroupResponseBody
	GetSuccess() *bool
}

type DeleteDynamicTagGroupResponseBody struct {
	// The status code.
	//
	// > The status code 200 indicates that the request was successful.
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
	// 08AAE67E-77B5-485B-9C79-D7C8C059150A
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

func (s DeleteDynamicTagGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDynamicTagGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDynamicTagGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteDynamicTagGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteDynamicTagGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDynamicTagGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDynamicTagGroupResponseBody) SetCode(v string) *DeleteDynamicTagGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDynamicTagGroupResponseBody) SetMessage(v string) *DeleteDynamicTagGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDynamicTagGroupResponseBody) SetRequestId(v string) *DeleteDynamicTagGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDynamicTagGroupResponseBody) SetSuccess(v bool) *DeleteDynamicTagGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDynamicTagGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
