// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterBatchRenewMemberApiKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpireAt(v string) *ModelRouterBatchRenewMemberApiKeysRequest
	GetExpireAt() *string
	SetUserIds(v []*int64) *ModelRouterBatchRenewMemberApiKeysRequest
	GetUserIds() []*int64
}

type ModelRouterBatchRenewMemberApiKeysRequest struct {
	// The new expiration time in RFC 3339 format. The time must be later than the current time. If this parameter is not provided or is set to null, the API keys remain permanently valid. This parameter only modifies the validity period and does not change the enabled or disabled status.
	//
	// example:
	//
	// 2027-01-01T00:00:00+08:00
	ExpireAt *string `json:"expireAt,omitempty" xml:"expireAt,omitempty"`
	// The list of member user IDs. This operation renews all undeleted API keys of these members in the specified department.
	//
	// This parameter is required.
	//
	// example:
	//
	// []
	UserIds []*int64 `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s ModelRouterBatchRenewMemberApiKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterBatchRenewMemberApiKeysRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterBatchRenewMemberApiKeysRequest) GetExpireAt() *string {
	return s.ExpireAt
}

func (s *ModelRouterBatchRenewMemberApiKeysRequest) GetUserIds() []*int64 {
	return s.UserIds
}

func (s *ModelRouterBatchRenewMemberApiKeysRequest) SetExpireAt(v string) *ModelRouterBatchRenewMemberApiKeysRequest {
	s.ExpireAt = &v
	return s
}

func (s *ModelRouterBatchRenewMemberApiKeysRequest) SetUserIds(v []*int64) *ModelRouterBatchRenewMemberApiKeysRequest {
	s.UserIds = v
	return s
}

func (s *ModelRouterBatchRenewMemberApiKeysRequest) Validate() error {
	return dara.Validate(s)
}
