// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateRoleResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdateRoleResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateRoleResponseBody
	GetRequestId() *string
}

type UpdateRoleResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information that is returned.
	//
	// example:
	//
	// edit successfully
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4823-bhjf-23u4-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRoleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRoleResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateRoleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRoleResponseBody) SetCode(v int32) *UpdateRoleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateRoleResponseBody) SetMessage(v string) *UpdateRoleResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateRoleResponseBody) SetRequestId(v string) *UpdateRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
