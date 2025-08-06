// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaLifecycleRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *GetMediaLifecycleRuleRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetMediaLifecycleRuleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetMediaLifecycleRuleRequest
	GetResourceOwnerId() *int64
	SetResourceRealOwnerId(v int64) *GetMediaLifecycleRuleRequest
	GetResourceRealOwnerId() *int64
	SetRuleIds(v string) *GetMediaLifecycleRuleRequest
	GetRuleIds() *string
}

type GetMediaLifecycleRuleRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceRealOwnerId  *int64  `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
	// This parameter is required.
	RuleIds *string `json:"RuleIds,omitempty" xml:"RuleIds,omitempty"`
}

func (s GetMediaLifecycleRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMediaLifecycleRuleRequest) GoString() string {
	return s.String()
}

func (s *GetMediaLifecycleRuleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetMediaLifecycleRuleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetMediaLifecycleRuleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetMediaLifecycleRuleRequest) GetResourceRealOwnerId() *int64 {
	return s.ResourceRealOwnerId
}

func (s *GetMediaLifecycleRuleRequest) GetRuleIds() *string {
	return s.RuleIds
}

func (s *GetMediaLifecycleRuleRequest) SetOwnerId(v int64) *GetMediaLifecycleRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *GetMediaLifecycleRuleRequest) SetResourceOwnerAccount(v string) *GetMediaLifecycleRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetMediaLifecycleRuleRequest) SetResourceOwnerId(v int64) *GetMediaLifecycleRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetMediaLifecycleRuleRequest) SetResourceRealOwnerId(v int64) *GetMediaLifecycleRuleRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *GetMediaLifecycleRuleRequest) SetRuleIds(v string) *GetMediaLifecycleRuleRequest {
	s.RuleIds = &v
	return s
}

func (s *GetMediaLifecycleRuleRequest) Validate() error {
	return dara.Validate(s)
}
