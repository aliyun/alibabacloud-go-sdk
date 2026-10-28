// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserEniVswitchOptionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyUserEniVswitchOptionsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ModifyUserEniVswitchOptionsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyUserEniVswitchOptionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyUserEniVswitchOptionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyUserEniVswitchOptionsResponseBody
	GetSuccess() *bool
}

type ModifyUserEniVswitchOptionsResponseBody struct {
	// The HTTP status code or the error code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message. Valid values:
	//
	// 	- If the request was successful, a success message is returned.****
	//
	// 	- If the request failed, an error message is returned.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyUserEniVswitchOptionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserEniVswitchOptionsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUserEniVswitchOptionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyUserEniVswitchOptionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyUserEniVswitchOptionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyUserEniVswitchOptionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyUserEniVswitchOptionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyUserEniVswitchOptionsResponseBody) SetCode(v string) *ModifyUserEniVswitchOptionsResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyUserEniVswitchOptionsResponseBody) SetHttpStatusCode(v int32) *ModifyUserEniVswitchOptionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyUserEniVswitchOptionsResponseBody) SetMessage(v string) *ModifyUserEniVswitchOptionsResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyUserEniVswitchOptionsResponseBody) SetRequestId(v string) *ModifyUserEniVswitchOptionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyUserEniVswitchOptionsResponseBody) SetSuccess(v bool) *ModifyUserEniVswitchOptionsResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyUserEniVswitchOptionsResponseBody) Validate() error {
	return dara.Validate(s)
}
