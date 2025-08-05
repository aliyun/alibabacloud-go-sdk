// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolicyV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeletePolicyV2ResponseBody
	GetCode() *string
	SetMessage(v string) *DeletePolicyV2ResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeletePolicyV2ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeletePolicyV2ResponseBody
	GetSuccess() *bool
}

type DeletePolicyV2ResponseBody struct {
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
	// 33AA3AAE-89E1-5D3A-A51D-0C0A80850F68
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

func (s DeletePolicyV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePolicyV2ResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePolicyV2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeletePolicyV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeletePolicyV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePolicyV2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeletePolicyV2ResponseBody) SetCode(v string) *DeletePolicyV2ResponseBody {
	s.Code = &v
	return s
}

func (s *DeletePolicyV2ResponseBody) SetMessage(v string) *DeletePolicyV2ResponseBody {
	s.Message = &v
	return s
}

func (s *DeletePolicyV2ResponseBody) SetRequestId(v string) *DeletePolicyV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePolicyV2ResponseBody) SetSuccess(v bool) *DeletePolicyV2ResponseBody {
	s.Success = &v
	return s
}

func (s *DeletePolicyV2ResponseBody) Validate() error {
	return dara.Validate(s)
}
