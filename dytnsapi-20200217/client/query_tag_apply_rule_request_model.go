// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTagApplyRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *QueryTagApplyRuleRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QueryTagApplyRuleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryTagApplyRuleRequest
	GetResourceOwnerId() *int64
	SetTagId(v int64) *QueryTagApplyRuleRequest
	GetTagId() *int64
}

type QueryTagApplyRuleRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tag ID.
	//
	// example:
	//
	// 61
	TagId *int64 `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s QueryTagApplyRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryTagApplyRuleRequest) GoString() string {
	return s.String()
}

func (s *QueryTagApplyRuleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryTagApplyRuleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryTagApplyRuleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryTagApplyRuleRequest) GetTagId() *int64 {
	return s.TagId
}

func (s *QueryTagApplyRuleRequest) SetOwnerId(v int64) *QueryTagApplyRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryTagApplyRuleRequest) SetResourceOwnerAccount(v string) *QueryTagApplyRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryTagApplyRuleRequest) SetResourceOwnerId(v int64) *QueryTagApplyRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryTagApplyRuleRequest) SetTagId(v int64) *QueryTagApplyRuleRequest {
	s.TagId = &v
	return s
}

func (s *QueryTagApplyRuleRequest) Validate() error {
	return dara.Validate(s)
}
