// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUserBatchQuitGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndUserIds(v []*string) *UserBatchQuitGroupRequest
	GetEndUserIds() []*string
	SetGroupId(v string) *UserBatchQuitGroupRequest
	GetGroupId() *string
}

type UserBatchQuitGroupRequest struct {
	EndUserIds []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	// example:
	//
	// ug-lkuvalovhnlxvv****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s UserBatchQuitGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UserBatchQuitGroupRequest) GoString() string {
	return s.String()
}

func (s *UserBatchQuitGroupRequest) GetEndUserIds() []*string {
	return s.EndUserIds
}

func (s *UserBatchQuitGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *UserBatchQuitGroupRequest) SetEndUserIds(v []*string) *UserBatchQuitGroupRequest {
	s.EndUserIds = v
	return s
}

func (s *UserBatchQuitGroupRequest) SetGroupId(v string) *UserBatchQuitGroupRequest {
	s.GroupId = &v
	return s
}

func (s *UserBatchQuitGroupRequest) Validate() error {
	return dara.Validate(s)
}
