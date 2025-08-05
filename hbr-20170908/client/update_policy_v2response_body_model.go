// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolicyV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdatePolicyV2ResponseBody
	GetCode() *string
	SetMessage(v string) *UpdatePolicyV2ResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdatePolicyV2ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdatePolicyV2ResponseBody
	GetSuccess() *bool
}

type UpdatePolicyV2ResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the call is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned. If the call is successful, "successful" is returned. If the call fails, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// 	- true: The call is successful.
	//
	// 	- false: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdatePolicyV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicyV2ResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePolicyV2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdatePolicyV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdatePolicyV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePolicyV2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdatePolicyV2ResponseBody) SetCode(v string) *UpdatePolicyV2ResponseBody {
	s.Code = &v
	return s
}

func (s *UpdatePolicyV2ResponseBody) SetMessage(v string) *UpdatePolicyV2ResponseBody {
	s.Message = &v
	return s
}

func (s *UpdatePolicyV2ResponseBody) SetRequestId(v string) *UpdatePolicyV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePolicyV2ResponseBody) SetSuccess(v bool) *UpdatePolicyV2ResponseBody {
	s.Success = &v
	return s
}

func (s *UpdatePolicyV2ResponseBody) Validate() error {
	return dara.Validate(s)
}
