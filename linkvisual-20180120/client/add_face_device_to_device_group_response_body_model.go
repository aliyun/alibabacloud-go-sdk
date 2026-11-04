// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFaceDeviceToDeviceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddFaceDeviceToDeviceGroupResponseBody
	GetCode() *string
	SetErrorMessage(v string) *AddFaceDeviceToDeviceGroupResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *AddFaceDeviceToDeviceGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddFaceDeviceToDeviceGroupResponseBody
	GetSuccess() *bool
}

type AddFaceDeviceToDeviceGroupResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// success
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0A13279A-5640-45E7-87AA-83750541AD0E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddFaceDeviceToDeviceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddFaceDeviceToDeviceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddFaceDeviceToDeviceGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddFaceDeviceToDeviceGroupResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *AddFaceDeviceToDeviceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddFaceDeviceToDeviceGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddFaceDeviceToDeviceGroupResponseBody) SetCode(v string) *AddFaceDeviceToDeviceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *AddFaceDeviceToDeviceGroupResponseBody) SetErrorMessage(v string) *AddFaceDeviceToDeviceGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddFaceDeviceToDeviceGroupResponseBody) SetRequestId(v string) *AddFaceDeviceToDeviceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddFaceDeviceToDeviceGroupResponseBody) SetSuccess(v bool) *AddFaceDeviceToDeviceGroupResponseBody {
	s.Success = &v
	return s
}

func (s *AddFaceDeviceToDeviceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
