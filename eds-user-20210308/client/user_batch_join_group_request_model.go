// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUserBatchJoinGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndUserIds(v []*string) *UserBatchJoinGroupRequest
	GetEndUserIds() []*string
	SetGroupId(v string) *UserBatchJoinGroupRequest
	GetGroupId() *string
}

type UserBatchJoinGroupRequest struct {
	// The list of user IDs.
	EndUserIds []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	// The ID of the user group.
	//
	// This parameter is required.
	//
	// example:
	//
	// ug-12341234****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s UserBatchJoinGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UserBatchJoinGroupRequest) GoString() string {
	return s.String()
}

func (s *UserBatchJoinGroupRequest) GetEndUserIds() []*string {
	return s.EndUserIds
}

func (s *UserBatchJoinGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *UserBatchJoinGroupRequest) SetEndUserIds(v []*string) *UserBatchJoinGroupRequest {
	s.EndUserIds = v
	return s
}

func (s *UserBatchJoinGroupRequest) SetGroupId(v string) *UserBatchJoinGroupRequest {
	s.GroupId = &v
	return s
}

func (s *UserBatchJoinGroupRequest) Validate() error {
	return dara.Validate(s)
}
