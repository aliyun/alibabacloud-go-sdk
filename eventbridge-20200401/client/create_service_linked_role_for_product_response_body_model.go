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
	// The returned response code. The value Success indicates that the request is successful. Other values indicate that the request failed. For more information about error codes, see Error codes.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The returned message. If the request is successful, success is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C36345A1-75F3-5A1A-BFCF-33B8271971FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. If the operation was successful, the value true is returned.
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
