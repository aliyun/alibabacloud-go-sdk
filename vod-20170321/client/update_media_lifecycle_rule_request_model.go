// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaLifecycleRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *UpdateMediaLifecycleRuleRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpdateMediaLifecycleRuleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateMediaLifecycleRuleRequest
	GetResourceOwnerId() *int64
	SetResourceRealOwnerId(v int64) *UpdateMediaLifecycleRuleRequest
	GetResourceRealOwnerId() *int64
	SetUpdateContent(v string) *UpdateMediaLifecycleRuleRequest
	GetUpdateContent() *string
}

type UpdateMediaLifecycleRuleRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceRealOwnerId  *int64  `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
	// This parameter is required.
	UpdateContent *string `json:"UpdateContent,omitempty" xml:"UpdateContent,omitempty"`
}

func (s UpdateMediaLifecycleRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaLifecycleRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateMediaLifecycleRuleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateMediaLifecycleRuleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateMediaLifecycleRuleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateMediaLifecycleRuleRequest) GetResourceRealOwnerId() *int64 {
	return s.ResourceRealOwnerId
}

func (s *UpdateMediaLifecycleRuleRequest) GetUpdateContent() *string {
	return s.UpdateContent
}

func (s *UpdateMediaLifecycleRuleRequest) SetOwnerId(v int64) *UpdateMediaLifecycleRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateMediaLifecycleRuleRequest) SetResourceOwnerAccount(v string) *UpdateMediaLifecycleRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateMediaLifecycleRuleRequest) SetResourceOwnerId(v int64) *UpdateMediaLifecycleRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateMediaLifecycleRuleRequest) SetResourceRealOwnerId(v int64) *UpdateMediaLifecycleRuleRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *UpdateMediaLifecycleRuleRequest) SetUpdateContent(v string) *UpdateMediaLifecycleRuleRequest {
	s.UpdateContent = &v
	return s
}

func (s *UpdateMediaLifecycleRuleRequest) Validate() error {
	return dara.Validate(s)
}
