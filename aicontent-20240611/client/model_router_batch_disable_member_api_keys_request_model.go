// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterBatchDisableMemberApiKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserIds(v []*int64) *ModelRouterBatchDisableMemberApiKeysRequest
	GetUserIds() []*int64
}

type ModelRouterBatchDisableMemberApiKeysRequest struct {
	// The list of member user IDs.
	//
	// example:
	//
	// []
	UserIds []*int64 `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s ModelRouterBatchDisableMemberApiKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterBatchDisableMemberApiKeysRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterBatchDisableMemberApiKeysRequest) GetUserIds() []*int64 {
	return s.UserIds
}

func (s *ModelRouterBatchDisableMemberApiKeysRequest) SetUserIds(v []*int64) *ModelRouterBatchDisableMemberApiKeysRequest {
	s.UserIds = v
	return s
}

func (s *ModelRouterBatchDisableMemberApiKeysRequest) Validate() error {
	return dara.Validate(s)
}
