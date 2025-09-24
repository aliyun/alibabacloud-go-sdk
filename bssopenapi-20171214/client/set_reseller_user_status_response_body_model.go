// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetResellerUserStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SetResellerUserStatusResponseBody
	GetCode() *string
	SetData(v bool) *SetResellerUserStatusResponseBody
	GetData() *bool
	SetMessage(v string) *SetResellerUserStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *SetResellerUserStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SetResellerUserStatusResponseBody
	GetSuccess() *bool
}

type SetResellerUserStatusResponseBody struct {
	// The error code returned if the call failed. For more information, see the "Error codes" section of this topic.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message returned.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F5B803CF-94D8-43AF-ADB3-D819AAD30E27
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. A value of true indicates that the call is successful. A value of false indicates that the call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetResellerUserStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetResellerUserStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetResellerUserStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *SetResellerUserStatusResponseBody) GetData() *bool {
	return s.Data
}

func (s *SetResellerUserStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SetResellerUserStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetResellerUserStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetResellerUserStatusResponseBody) SetCode(v string) *SetResellerUserStatusResponseBody {
	s.Code = &v
	return s
}

func (s *SetResellerUserStatusResponseBody) SetData(v bool) *SetResellerUserStatusResponseBody {
	s.Data = &v
	return s
}

func (s *SetResellerUserStatusResponseBody) SetMessage(v string) *SetResellerUserStatusResponseBody {
	s.Message = &v
	return s
}

func (s *SetResellerUserStatusResponseBody) SetRequestId(v string) *SetResellerUserStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetResellerUserStatusResponseBody) SetSuccess(v bool) *SetResellerUserStatusResponseBody {
	s.Success = &v
	return s
}

func (s *SetResellerUserStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
