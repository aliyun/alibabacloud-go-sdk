// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterBatchCreateMemberApiKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpireAt(v string) *ModelRouterBatchCreateMemberApiKeysRequest
	GetExpireAt() *string
	SetName(v string) *ModelRouterBatchCreateMemberApiKeysRequest
	GetName() *string
	SetUserIds(v []*int64) *ModelRouterBatchCreateMemberApiKeysRequest
	GetUserIds() []*int64
}

type ModelRouterBatchCreateMemberApiKeysRequest struct {
	// The expiration time in RFC 3339 format. This parameter is optional. If not specified, the key is permanently valid.
	//
	// example:
	//
	// 2027-01-01T00:00:00Z
	ExpireAt *string `json:"expireAt,omitempty" xml:"expireAt,omitempty"`
	// The key name. This parameter is optional.
	//
	// example:
	//
	// TestKey
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The list of member user IDs.
	//
	// example:
	//
	// []
	UserIds []*int64 `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s ModelRouterBatchCreateMemberApiKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterBatchCreateMemberApiKeysRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterBatchCreateMemberApiKeysRequest) GetExpireAt() *string {
	return s.ExpireAt
}

func (s *ModelRouterBatchCreateMemberApiKeysRequest) GetName() *string {
	return s.Name
}

func (s *ModelRouterBatchCreateMemberApiKeysRequest) GetUserIds() []*int64 {
	return s.UserIds
}

func (s *ModelRouterBatchCreateMemberApiKeysRequest) SetExpireAt(v string) *ModelRouterBatchCreateMemberApiKeysRequest {
	s.ExpireAt = &v
	return s
}

func (s *ModelRouterBatchCreateMemberApiKeysRequest) SetName(v string) *ModelRouterBatchCreateMemberApiKeysRequest {
	s.Name = &v
	return s
}

func (s *ModelRouterBatchCreateMemberApiKeysRequest) SetUserIds(v []*int64) *ModelRouterBatchCreateMemberApiKeysRequest {
	s.UserIds = v
	return s
}

func (s *ModelRouterBatchCreateMemberApiKeysRequest) Validate() error {
	return dara.Validate(s)
}
