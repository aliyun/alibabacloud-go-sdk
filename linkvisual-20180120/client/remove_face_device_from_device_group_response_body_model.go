// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveFaceDeviceFromDeviceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RemoveFaceDeviceFromDeviceGroupResponseBody
	GetCode() *string
	SetErrorMessage(v string) *RemoveFaceDeviceFromDeviceGroupResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *RemoveFaceDeviceFromDeviceGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RemoveFaceDeviceFromDeviceGroupResponseBody
	GetSuccess() *bool
}

type RemoveFaceDeviceFromDeviceGroupResponseBody struct {
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

func (s RemoveFaceDeviceFromDeviceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveFaceDeviceFromDeviceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveFaceDeviceFromDeviceGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *RemoveFaceDeviceFromDeviceGroupResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RemoveFaceDeviceFromDeviceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveFaceDeviceFromDeviceGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveFaceDeviceFromDeviceGroupResponseBody) SetCode(v string) *RemoveFaceDeviceFromDeviceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveFaceDeviceFromDeviceGroupResponseBody) SetErrorMessage(v string) *RemoveFaceDeviceFromDeviceGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RemoveFaceDeviceFromDeviceGroupResponseBody) SetRequestId(v string) *RemoveFaceDeviceFromDeviceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveFaceDeviceFromDeviceGroupResponseBody) SetSuccess(v bool) *RemoveFaceDeviceFromDeviceGroupResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveFaceDeviceFromDeviceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
