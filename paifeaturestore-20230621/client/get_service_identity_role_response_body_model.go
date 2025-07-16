// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceIdentityRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicy(v string) *GetServiceIdentityRoleResponseBody
	GetPolicy() *string
	SetRequestId(v string) *GetServiceIdentityRoleResponseBody
	GetRequestId() *string
	SetRoleName(v string) *GetServiceIdentityRoleResponseBody
	GetRoleName() *string
}

type GetServiceIdentityRoleResponseBody struct {
	// example:
	//
	// {
	//
	// "Version": "1",
	//
	// "Statement":[]
	//
	// }
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// 6F629E92-F64D-502D-85AA-A9C54894CA3D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// AliyunServiceRoleForPaiFeatureStore
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s GetServiceIdentityRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceIdentityRoleResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceIdentityRoleResponseBody) GetPolicy() *string {
	return s.Policy
}

func (s *GetServiceIdentityRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetServiceIdentityRoleResponseBody) GetRoleName() *string {
	return s.RoleName
}

func (s *GetServiceIdentityRoleResponseBody) SetPolicy(v string) *GetServiceIdentityRoleResponseBody {
	s.Policy = &v
	return s
}

func (s *GetServiceIdentityRoleResponseBody) SetRequestId(v string) *GetServiceIdentityRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceIdentityRoleResponseBody) SetRoleName(v string) *GetServiceIdentityRoleResponseBody {
	s.RoleName = &v
	return s
}

func (s *GetServiceIdentityRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
