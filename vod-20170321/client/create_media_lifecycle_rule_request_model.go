// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMediaLifecycleRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateMediaLifecycleRuleRequest
	GetAppId() *string
	SetOwnerId(v int64) *CreateMediaLifecycleRuleRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateMediaLifecycleRuleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateMediaLifecycleRuleRequest
	GetResourceOwnerId() *int64
	SetResourceRealOwnerId(v int64) *CreateMediaLifecycleRuleRequest
	GetResourceRealOwnerId() *int64
	SetRuleContent(v string) *CreateMediaLifecycleRuleRequest
	GetRuleContent() *string
	SetRuleType(v string) *CreateMediaLifecycleRuleRequest
	GetRuleType() *string
}

type CreateMediaLifecycleRuleRequest struct {
	AppId                *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceRealOwnerId  *int64  `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
	RuleContent          *string `json:"RuleContent,omitempty" xml:"RuleContent,omitempty"`
	RuleType             *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s CreateMediaLifecycleRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaLifecycleRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateMediaLifecycleRuleRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateMediaLifecycleRuleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateMediaLifecycleRuleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateMediaLifecycleRuleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateMediaLifecycleRuleRequest) GetResourceRealOwnerId() *int64 {
	return s.ResourceRealOwnerId
}

func (s *CreateMediaLifecycleRuleRequest) GetRuleContent() *string {
	return s.RuleContent
}

func (s *CreateMediaLifecycleRuleRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *CreateMediaLifecycleRuleRequest) SetAppId(v string) *CreateMediaLifecycleRuleRequest {
	s.AppId = &v
	return s
}

func (s *CreateMediaLifecycleRuleRequest) SetOwnerId(v int64) *CreateMediaLifecycleRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateMediaLifecycleRuleRequest) SetResourceOwnerAccount(v string) *CreateMediaLifecycleRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateMediaLifecycleRuleRequest) SetResourceOwnerId(v int64) *CreateMediaLifecycleRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateMediaLifecycleRuleRequest) SetResourceRealOwnerId(v int64) *CreateMediaLifecycleRuleRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *CreateMediaLifecycleRuleRequest) SetRuleContent(v string) *CreateMediaLifecycleRuleRequest {
	s.RuleContent = &v
	return s
}

func (s *CreateMediaLifecycleRuleRequest) SetRuleType(v string) *CreateMediaLifecycleRuleRequest {
	s.RuleType = &v
	return s
}

func (s *CreateMediaLifecycleRuleRequest) Validate() error {
	return dara.Validate(s)
}
