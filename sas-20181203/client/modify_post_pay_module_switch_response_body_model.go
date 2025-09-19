// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPostPayModuleSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyPostPayModuleSwitchResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ModifyPostPayModuleSwitchResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyPostPayModuleSwitchResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyPostPayModuleSwitchResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyPostPayModuleSwitchResponseBody
	GetSuccess() *bool
}

type ModifyPostPayModuleSwitchResponseBody struct {
	// The response code. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D03DD0FD-**28F1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
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

func (s ModifyPostPayModuleSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyPostPayModuleSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPostPayModuleSwitchResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyPostPayModuleSwitchResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyPostPayModuleSwitchResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyPostPayModuleSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyPostPayModuleSwitchResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyPostPayModuleSwitchResponseBody) SetCode(v string) *ModifyPostPayModuleSwitchResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyPostPayModuleSwitchResponseBody) SetHttpStatusCode(v int32) *ModifyPostPayModuleSwitchResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyPostPayModuleSwitchResponseBody) SetMessage(v string) *ModifyPostPayModuleSwitchResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyPostPayModuleSwitchResponseBody) SetRequestId(v string) *ModifyPostPayModuleSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPostPayModuleSwitchResponseBody) SetSuccess(v bool) *ModifyPostPayModuleSwitchResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyPostPayModuleSwitchResponseBody) Validate() error {
	return dara.Validate(s)
}
