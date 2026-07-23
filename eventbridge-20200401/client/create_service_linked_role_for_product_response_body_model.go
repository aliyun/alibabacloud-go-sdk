// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceLinkedRoleForProductResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateServiceLinkedRoleForProductResponseBody
	GetCode() *string
	SetHttpCode(v int32) *CreateServiceLinkedRoleForProductResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *CreateServiceLinkedRoleForProductResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateServiceLinkedRoleForProductResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateServiceLinkedRoleForProductResponseBody
	GetSuccess() *bool
}

type CreateServiceLinkedRoleForProductResponseBody struct {
	// The status code of the request. \\`Success\\` indicates that the request was successful. For more information about error codes, see the Error codes section.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code. Valid values:
	//
	// - **2xx**: a 2xx status code.
	//
	// - **3xx**: a 3xx status code.
	//
	// - **4xx**: a 4xx status code.
	//
	// - **5xx**: a 5xx status code.
	//
	// If this parameter is not specified, all HTTP status codes are queried.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The returned message. If the request is successful, \\`success\\` is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C36345A1-75F3-5A1A-BFCF-33B8271971FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. If the request is successful, \\`true\\` is returned.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateServiceLinkedRoleForProductResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceLinkedRoleForProductResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleForProductResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateServiceLinkedRoleForProductResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *CreateServiceLinkedRoleForProductResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateServiceLinkedRoleForProductResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServiceLinkedRoleForProductResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateServiceLinkedRoleForProductResponseBody) SetCode(v string) *CreateServiceLinkedRoleForProductResponseBody {
	s.Code = &v
	return s
}

func (s *CreateServiceLinkedRoleForProductResponseBody) SetHttpCode(v int32) *CreateServiceLinkedRoleForProductResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateServiceLinkedRoleForProductResponseBody) SetMessage(v string) *CreateServiceLinkedRoleForProductResponseBody {
	s.Message = &v
	return s
}

func (s *CreateServiceLinkedRoleForProductResponseBody) SetRequestId(v string) *CreateServiceLinkedRoleForProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceLinkedRoleForProductResponseBody) SetSuccess(v bool) *CreateServiceLinkedRoleForProductResponseBody {
	s.Success = &v
	return s
}

func (s *CreateServiceLinkedRoleForProductResponseBody) Validate() error {
	return dara.Validate(s)
}
