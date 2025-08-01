// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceIdentityRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetServiceIdentityRoleResponseBody
	GetCode() *string
	SetMessage(v string) *GetServiceIdentityRoleResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetServiceIdentityRoleResponseBody
	GetRequestId() *string
	SetRoleDetail(v string) *GetServiceIdentityRoleResponseBody
	GetRoleDetail() *string
	SetRoleName(v string) *GetServiceIdentityRoleResponseBody
	GetRoleName() *string
}

type GetServiceIdentityRoleResponseBody struct {
	// The internal error code. This parameter is returned only when an error occurs.
	//
	// example:
	//
	// EntityNotExist
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message. This parameter is returned only when an error occurs.
	//
	// example:
	//
	// Serivce role does not exit.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6A87228C-969A-1381-98CF-AE07AE630FA5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The role details.
	//
	// example:
	//
	// AliyunServiceRoleForPaiLLMTrace
	RoleDetail *string `json:"RoleDetail,omitempty" xml:"RoleDetail,omitempty"`
	// The name of the service-linked role. Default value: AliyunServiceRoleForPaiLLMTrace.
	//
	// example:
	//
	// AliyunServiceRoleForPaiLLMTrace
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s GetServiceIdentityRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceIdentityRoleResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceIdentityRoleResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetServiceIdentityRoleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetServiceIdentityRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetServiceIdentityRoleResponseBody) GetRoleDetail() *string {
	return s.RoleDetail
}

func (s *GetServiceIdentityRoleResponseBody) GetRoleName() *string {
	return s.RoleName
}

func (s *GetServiceIdentityRoleResponseBody) SetCode(v string) *GetServiceIdentityRoleResponseBody {
	s.Code = &v
	return s
}

func (s *GetServiceIdentityRoleResponseBody) SetMessage(v string) *GetServiceIdentityRoleResponseBody {
	s.Message = &v
	return s
}

func (s *GetServiceIdentityRoleResponseBody) SetRequestId(v string) *GetServiceIdentityRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceIdentityRoleResponseBody) SetRoleDetail(v string) *GetServiceIdentityRoleResponseBody {
	s.RoleDetail = &v
	return s
}

func (s *GetServiceIdentityRoleResponseBody) SetRoleName(v string) *GetServiceIdentityRoleResponseBody {
	s.RoleName = &v
	return s
}

func (s *GetServiceIdentityRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
