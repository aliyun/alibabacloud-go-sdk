// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediaLifecycleRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DeleteMediaLifecycleRuleRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteMediaLifecycleRuleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteMediaLifecycleRuleRequest
	GetResourceOwnerId() *int64
	SetResourceRealOwnerId(v int64) *DeleteMediaLifecycleRuleRequest
	GetResourceRealOwnerId() *int64
	SetRuleIds(v string) *DeleteMediaLifecycleRuleRequest
	GetRuleIds() *string
}

type DeleteMediaLifecycleRuleRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceRealOwnerId  *int64  `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
	// This parameter is required.
	RuleIds *string `json:"RuleIds,omitempty" xml:"RuleIds,omitempty"`
}

func (s DeleteMediaLifecycleRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaLifecycleRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteMediaLifecycleRuleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteMediaLifecycleRuleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteMediaLifecycleRuleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteMediaLifecycleRuleRequest) GetResourceRealOwnerId() *int64 {
	return s.ResourceRealOwnerId
}

func (s *DeleteMediaLifecycleRuleRequest) GetRuleIds() *string {
	return s.RuleIds
}

func (s *DeleteMediaLifecycleRuleRequest) SetOwnerId(v int64) *DeleteMediaLifecycleRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteMediaLifecycleRuleRequest) SetResourceOwnerAccount(v string) *DeleteMediaLifecycleRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteMediaLifecycleRuleRequest) SetResourceOwnerId(v int64) *DeleteMediaLifecycleRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteMediaLifecycleRuleRequest) SetResourceRealOwnerId(v int64) *DeleteMediaLifecycleRuleRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *DeleteMediaLifecycleRuleRequest) SetRuleIds(v string) *DeleteMediaLifecycleRuleRequest {
	s.RuleIds = &v
	return s
}

func (s *DeleteMediaLifecycleRuleRequest) Validate() error {
	return dara.Validate(s)
}
