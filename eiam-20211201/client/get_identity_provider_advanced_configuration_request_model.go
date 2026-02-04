// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentityProviderAdvancedConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentityProviderId(v string) *GetIdentityProviderAdvancedConfigurationRequest
	GetIdentityProviderId() *string
	SetInstanceId(v string) *GetIdentityProviderAdvancedConfigurationRequest
	GetInstanceId() *string
}

type GetIdentityProviderAdvancedConfigurationRequest struct {
	// IDaaS的身份提供方主键id
	//
	// This parameter is required.
	//
	// example:
	//
	// idp_11111
	IdentityProviderId *string `json:"IdentityProviderId,omitempty" xml:"IdentityProviderId,omitempty"`
	// IDaaS EIAM的实例id
	//
	// This parameter is required.
	//
	// example:
	//
	// eiam-111ccc1111
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetIdentityProviderAdvancedConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderAdvancedConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderAdvancedConfigurationRequest) GetIdentityProviderId() *string {
	return s.IdentityProviderId
}

func (s *GetIdentityProviderAdvancedConfigurationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetIdentityProviderAdvancedConfigurationRequest) SetIdentityProviderId(v string) *GetIdentityProviderAdvancedConfigurationRequest {
	s.IdentityProviderId = &v
	return s
}

func (s *GetIdentityProviderAdvancedConfigurationRequest) SetInstanceId(v string) *GetIdentityProviderAdvancedConfigurationRequest {
	s.InstanceId = &v
	return s
}

func (s *GetIdentityProviderAdvancedConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
