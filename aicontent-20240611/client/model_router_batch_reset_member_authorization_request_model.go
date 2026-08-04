// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterBatchResetMemberAuthorizationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserIds(v []*int64) *ModelRouterBatchResetMemberAuthorizationRequest
	GetUserIds() []*int64
}

type ModelRouterBatchResetMemberAuthorizationRequest struct {
	// The list of member user IDs.
	//
	// example:
	//
	// []
	UserIds []*int64 `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s ModelRouterBatchResetMemberAuthorizationRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterBatchResetMemberAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterBatchResetMemberAuthorizationRequest) GetUserIds() []*int64 {
	return s.UserIds
}

func (s *ModelRouterBatchResetMemberAuthorizationRequest) SetUserIds(v []*int64) *ModelRouterBatchResetMemberAuthorizationRequest {
	s.UserIds = v
	return s
}

func (s *ModelRouterBatchResetMemberAuthorizationRequest) Validate() error {
	return dara.Validate(s)
}
