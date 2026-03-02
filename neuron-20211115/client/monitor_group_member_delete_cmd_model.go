// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonitorGroupMemberDeleteCmd interface {
	dara.Model
	String() string
	GoString() string
	SetContactIds(v []*int64) *MonitorGroupMemberDeleteCmd
	GetContactIds() []*int64
	SetGroupId(v int64) *MonitorGroupMemberDeleteCmd
	GetGroupId() *int64
}

type MonitorGroupMemberDeleteCmd struct {
	ContactIds []*int64 `json:"contactIds,omitempty" xml:"contactIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	GroupId *int64 `json:"groupId,omitempty" xml:"groupId,omitempty"`
}

func (s MonitorGroupMemberDeleteCmd) String() string {
	return dara.Prettify(s)
}

func (s MonitorGroupMemberDeleteCmd) GoString() string {
	return s.String()
}

func (s *MonitorGroupMemberDeleteCmd) GetContactIds() []*int64 {
	return s.ContactIds
}

func (s *MonitorGroupMemberDeleteCmd) GetGroupId() *int64 {
	return s.GroupId
}

func (s *MonitorGroupMemberDeleteCmd) SetContactIds(v []*int64) *MonitorGroupMemberDeleteCmd {
	s.ContactIds = v
	return s
}

func (s *MonitorGroupMemberDeleteCmd) SetGroupId(v int64) *MonitorGroupMemberDeleteCmd {
	s.GroupId = &v
	return s
}

func (s *MonitorGroupMemberDeleteCmd) Validate() error {
	return dara.Validate(s)
}
