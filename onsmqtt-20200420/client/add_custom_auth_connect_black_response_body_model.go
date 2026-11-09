// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomAuthConnectBlackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AddCustomAuthConnectBlackResponseBody
	GetCode() *int32
	SetMessage(v string) *AddCustomAuthConnectBlackResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddCustomAuthConnectBlackResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddCustomAuthConnectBlackResponseBody
	GetSuccess() *bool
}

type AddCustomAuthConnectBlackResponseBody struct {
	// The status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82B9E503-F4A1-4F30-976F-C6999FF9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. Valid values:
	//
	// - true: The operation is successful.
	//
	// - false: The operation fails.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddCustomAuthConnectBlackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddCustomAuthConnectBlackResponseBody) GoString() string {
	return s.String()
}

func (s *AddCustomAuthConnectBlackResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AddCustomAuthConnectBlackResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddCustomAuthConnectBlackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddCustomAuthConnectBlackResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddCustomAuthConnectBlackResponseBody) SetCode(v int32) *AddCustomAuthConnectBlackResponseBody {
	s.Code = &v
	return s
}

func (s *AddCustomAuthConnectBlackResponseBody) SetMessage(v string) *AddCustomAuthConnectBlackResponseBody {
	s.Message = &v
	return s
}

func (s *AddCustomAuthConnectBlackResponseBody) SetRequestId(v string) *AddCustomAuthConnectBlackResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCustomAuthConnectBlackResponseBody) SetSuccess(v bool) *AddCustomAuthConnectBlackResponseBody {
	s.Success = &v
	return s
}

func (s *AddCustomAuthConnectBlackResponseBody) Validate() error {
	return dara.Validate(s)
}
