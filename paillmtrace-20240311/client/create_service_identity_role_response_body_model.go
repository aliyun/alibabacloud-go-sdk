// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceIdentityRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateServiceIdentityRoleResponseBody
	GetCode() *string
	SetMessage(v string) *CreateServiceIdentityRoleResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateServiceIdentityRoleResponseBody
	GetRequestId() *string
	SetRoleDetails(v string) *CreateServiceIdentityRoleResponseBody
	GetRoleDetails() *string
	SetRoleName(v string) *CreateServiceIdentityRoleResponseBody
	GetRoleName() *string
}

type CreateServiceIdentityRoleResponseBody struct {
	// The error code returned if the request fails.
	//
	// example:
	//
	// InvalidInputParams
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned if the request fails.
	//
	// example:
	//
	// User don\\"t have permission to create SLR.
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
	RoleDetails *string `json:"RoleDetails,omitempty" xml:"RoleDetails,omitempty"`
	// The name of the service-linked role. Default value: AliyunServiceRoleForPaiLLMTrace.
	//
	// example:
	//
	// AliyunServiceRoleForPaiLLMTrace
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s CreateServiceIdentityRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceIdentityRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceIdentityRoleResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateServiceIdentityRoleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateServiceIdentityRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServiceIdentityRoleResponseBody) GetRoleDetails() *string {
	return s.RoleDetails
}

func (s *CreateServiceIdentityRoleResponseBody) GetRoleName() *string {
	return s.RoleName
}

func (s *CreateServiceIdentityRoleResponseBody) SetCode(v string) *CreateServiceIdentityRoleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateServiceIdentityRoleResponseBody) SetMessage(v string) *CreateServiceIdentityRoleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateServiceIdentityRoleResponseBody) SetRequestId(v string) *CreateServiceIdentityRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceIdentityRoleResponseBody) SetRoleDetails(v string) *CreateServiceIdentityRoleResponseBody {
	s.RoleDetails = &v
	return s
}

func (s *CreateServiceIdentityRoleResponseBody) SetRoleName(v string) *CreateServiceIdentityRoleResponseBody {
	s.RoleName = &v
	return s
}

func (s *CreateServiceIdentityRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
