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
	SetRequestId(v string) *CreateServiceIdentityRoleResponseBody
	GetRequestId() *string
	SetRoleName(v string) *CreateServiceIdentityRoleResponseBody
	GetRoleName() *string
}

type CreateServiceIdentityRoleResponseBody struct {
	// example:
	//
	// ServiceLinkedRoleAlreadyExistsErrorProblem
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// C03B2680-AC9C-59CD-93C5-8142B92537FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// AliyunServiceRoleForFeatureStore
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

func (s *CreateServiceIdentityRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServiceIdentityRoleResponseBody) GetRoleName() *string {
	return s.RoleName
}

func (s *CreateServiceIdentityRoleResponseBody) SetCode(v string) *CreateServiceIdentityRoleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateServiceIdentityRoleResponseBody) SetRequestId(v string) *CreateServiceIdentityRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceIdentityRoleResponseBody) SetRoleName(v string) *CreateServiceIdentityRoleResponseBody {
	s.RoleName = &v
	return s
}

func (s *CreateServiceIdentityRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
