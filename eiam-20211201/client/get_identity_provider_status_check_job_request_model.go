// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentityProviderStatusCheckJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentityProviderId(v string) *GetIdentityProviderStatusCheckJobRequest
	GetIdentityProviderId() *string
	SetIdentityProviderStatusCheckJobId(v string) *GetIdentityProviderStatusCheckJobRequest
	GetIdentityProviderStatusCheckJobId() *string
	SetInstanceId(v string) *GetIdentityProviderStatusCheckJobRequest
	GetInstanceId() *string
}

type GetIdentityProviderStatusCheckJobRequest struct {
	// IDaaS的身份提供方主键id
	//
	// This parameter is required.
	//
	// example:
	//
	// idp_11111
	IdentityProviderId *string `json:"IdentityProviderId,omitempty" xml:"IdentityProviderId,omitempty"`
	// 任务ID
	//
	// This parameter is required.
	//
	// example:
	//
	// async_000xxxx
	IdentityProviderStatusCheckJobId *string `json:"IdentityProviderStatusCheckJobId,omitempty" xml:"IdentityProviderStatusCheckJobId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetIdentityProviderStatusCheckJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderStatusCheckJobRequest) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderStatusCheckJobRequest) GetIdentityProviderId() *string {
	return s.IdentityProviderId
}

func (s *GetIdentityProviderStatusCheckJobRequest) GetIdentityProviderStatusCheckJobId() *string {
	return s.IdentityProviderStatusCheckJobId
}

func (s *GetIdentityProviderStatusCheckJobRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetIdentityProviderStatusCheckJobRequest) SetIdentityProviderId(v string) *GetIdentityProviderStatusCheckJobRequest {
	s.IdentityProviderId = &v
	return s
}

func (s *GetIdentityProviderStatusCheckJobRequest) SetIdentityProviderStatusCheckJobId(v string) *GetIdentityProviderStatusCheckJobRequest {
	s.IdentityProviderStatusCheckJobId = &v
	return s
}

func (s *GetIdentityProviderStatusCheckJobRequest) SetInstanceId(v string) *GetIdentityProviderStatusCheckJobRequest {
	s.InstanceId = &v
	return s
}

func (s *GetIdentityProviderStatusCheckJobRequest) Validate() error {
	return dara.Validate(s)
}
