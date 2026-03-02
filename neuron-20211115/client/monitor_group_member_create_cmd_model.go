// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonitorGroupMemberCreateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *MonitorGroupMemberCreateCmd
	GetAccountId() *string
	SetContactIds(v []*int64) *MonitorGroupMemberCreateCmd
	GetContactIds() []*int64
}

type MonitorGroupMemberCreateCmd struct {
	// example:
	//
	// 121321412341
	AccountId  *string  `json:"accountId,omitempty" xml:"accountId,omitempty"`
	ContactIds []*int64 `json:"contactIds,omitempty" xml:"contactIds,omitempty" type:"Repeated"`
}

func (s MonitorGroupMemberCreateCmd) String() string {
	return dara.Prettify(s)
}

func (s MonitorGroupMemberCreateCmd) GoString() string {
	return s.String()
}

func (s *MonitorGroupMemberCreateCmd) GetAccountId() *string {
	return s.AccountId
}

func (s *MonitorGroupMemberCreateCmd) GetContactIds() []*int64 {
	return s.ContactIds
}

func (s *MonitorGroupMemberCreateCmd) SetAccountId(v string) *MonitorGroupMemberCreateCmd {
	s.AccountId = &v
	return s
}

func (s *MonitorGroupMemberCreateCmd) SetContactIds(v []*int64) *MonitorGroupMemberCreateCmd {
	s.ContactIds = v
	return s
}

func (s *MonitorGroupMemberCreateCmd) Validate() error {
	return dara.Validate(s)
}
