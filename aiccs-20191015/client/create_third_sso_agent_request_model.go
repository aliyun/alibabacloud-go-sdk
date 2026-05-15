// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateThirdSsoAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *CreateThirdSsoAgentRequest
	GetAccountId() *string
	SetAccountName(v string) *CreateThirdSsoAgentRequest
	GetAccountName() *string
	SetClientId(v string) *CreateThirdSsoAgentRequest
	GetClientId() *string
	SetClientToken(v string) *CreateThirdSsoAgentRequest
	GetClientToken() *string
	SetDisplayName(v string) *CreateThirdSsoAgentRequest
	GetDisplayName() *string
	SetInstanceId(v string) *CreateThirdSsoAgentRequest
	GetInstanceId() *string
	SetRoleIds(v []*int64) *CreateThirdSsoAgentRequest
	GetRoleIds() []*int64
	SetSkillGroupIds(v []*int64) *CreateThirdSsoAgentRequest
	GetSkillGroupIds() []*int64
}

type CreateThirdSsoAgentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// accountId1
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// accountName1
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 46c1341e-2648-447a-9b11-70b6a298d94d
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// 46c1341e-2648-447a-9b11-70b6a298d94d
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre***
	InstanceId    *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RoleIds       []*int64 `json:"RoleIds,omitempty" xml:"RoleIds,omitempty" type:"Repeated"`
	SkillGroupIds []*int64 `json:"SkillGroupIds,omitempty" xml:"SkillGroupIds,omitempty" type:"Repeated"`
}

func (s CreateThirdSsoAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateThirdSsoAgentRequest) GoString() string {
	return s.String()
}

func (s *CreateThirdSsoAgentRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *CreateThirdSsoAgentRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *CreateThirdSsoAgentRequest) GetClientId() *string {
	return s.ClientId
}

func (s *CreateThirdSsoAgentRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateThirdSsoAgentRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateThirdSsoAgentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateThirdSsoAgentRequest) GetRoleIds() []*int64 {
	return s.RoleIds
}

func (s *CreateThirdSsoAgentRequest) GetSkillGroupIds() []*int64 {
	return s.SkillGroupIds
}

func (s *CreateThirdSsoAgentRequest) SetAccountId(v string) *CreateThirdSsoAgentRequest {
	s.AccountId = &v
	return s
}

func (s *CreateThirdSsoAgentRequest) SetAccountName(v string) *CreateThirdSsoAgentRequest {
	s.AccountName = &v
	return s
}

func (s *CreateThirdSsoAgentRequest) SetClientId(v string) *CreateThirdSsoAgentRequest {
	s.ClientId = &v
	return s
}

func (s *CreateThirdSsoAgentRequest) SetClientToken(v string) *CreateThirdSsoAgentRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateThirdSsoAgentRequest) SetDisplayName(v string) *CreateThirdSsoAgentRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateThirdSsoAgentRequest) SetInstanceId(v string) *CreateThirdSsoAgentRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateThirdSsoAgentRequest) SetRoleIds(v []*int64) *CreateThirdSsoAgentRequest {
	s.RoleIds = v
	return s
}

func (s *CreateThirdSsoAgentRequest) SetSkillGroupIds(v []*int64) *CreateThirdSsoAgentRequest {
	s.SkillGroupIds = v
	return s
}

func (s *CreateThirdSsoAgentRequest) Validate() error {
	return dara.Validate(s)
}
