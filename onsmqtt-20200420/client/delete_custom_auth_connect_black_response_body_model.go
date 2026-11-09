// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomAuthConnectBlackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteCustomAuthConnectBlackResponseBody
	GetCode() *int32
	SetMessage(v string) *DeleteCustomAuthConnectBlackResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteCustomAuthConnectBlackResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteCustomAuthConnectBlackResponseBody
	GetSuccess() *bool
}

type DeleteCustomAuthConnectBlackResponseBody struct {
	// The status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result message of the call.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 317076B7-F946-46BC-A98F-4CF9777C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. A value of true indicates success, and a value of false indicates failure.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCustomAuthConnectBlackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomAuthConnectBlackResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomAuthConnectBlackResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteCustomAuthConnectBlackResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteCustomAuthConnectBlackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomAuthConnectBlackResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCustomAuthConnectBlackResponseBody) SetCode(v int32) *DeleteCustomAuthConnectBlackResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCustomAuthConnectBlackResponseBody) SetMessage(v string) *DeleteCustomAuthConnectBlackResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCustomAuthConnectBlackResponseBody) SetRequestId(v string) *DeleteCustomAuthConnectBlackResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomAuthConnectBlackResponseBody) SetSuccess(v bool) *DeleteCustomAuthConnectBlackResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCustomAuthConnectBlackResponseBody) Validate() error {
	return dara.Validate(s)
}
