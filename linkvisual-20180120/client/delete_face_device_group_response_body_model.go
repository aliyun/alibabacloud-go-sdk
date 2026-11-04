// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaceDeviceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteFaceDeviceGroupResponseBody
	GetCode() *string
	SetErrorMessage(v string) *DeleteFaceDeviceGroupResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteFaceDeviceGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteFaceDeviceGroupResponseBody
	GetSuccess() *bool
}

type DeleteFaceDeviceGroupResponseBody struct {
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

func (s DeleteFaceDeviceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaceDeviceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFaceDeviceGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteFaceDeviceGroupResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteFaceDeviceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFaceDeviceGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteFaceDeviceGroupResponseBody) SetCode(v string) *DeleteFaceDeviceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFaceDeviceGroupResponseBody) SetErrorMessage(v string) *DeleteFaceDeviceGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteFaceDeviceGroupResponseBody) SetRequestId(v string) *DeleteFaceDeviceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFaceDeviceGroupResponseBody) SetSuccess(v bool) *DeleteFaceDeviceGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteFaceDeviceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
