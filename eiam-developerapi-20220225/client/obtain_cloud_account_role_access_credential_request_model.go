// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObtainCloudAccountRoleAccessCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloudAccountRoleExternalId(v string) *ObtainCloudAccountRoleAccessCredentialRequest
	GetCloudAccountRoleExternalId() *string
}

type ObtainCloudAccountRoleAccessCredentialRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// acs:ram::xxx:role/role-test
	CloudAccountRoleExternalId *string `json:"cloudAccountRoleExternalId,omitempty" xml:"cloudAccountRoleExternalId,omitempty"`
}

func (s ObtainCloudAccountRoleAccessCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s ObtainCloudAccountRoleAccessCredentialRequest) GoString() string {
	return s.String()
}

func (s *ObtainCloudAccountRoleAccessCredentialRequest) GetCloudAccountRoleExternalId() *string {
	return s.CloudAccountRoleExternalId
}

func (s *ObtainCloudAccountRoleAccessCredentialRequest) SetCloudAccountRoleExternalId(v string) *ObtainCloudAccountRoleAccessCredentialRequest {
	s.CloudAccountRoleExternalId = &v
	return s
}

func (s *ObtainCloudAccountRoleAccessCredentialRequest) Validate() error {
	return dara.Validate(s)
}
