// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonitorContactGroupUpdateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *MonitorContactGroupUpdateCmd
	GetAccountId() *string
	SetContactIds(v []*int64) *MonitorContactGroupUpdateCmd
	GetContactIds() []*int64
	SetId(v int64) *MonitorContactGroupUpdateCmd
	GetId() *int64
	SetName(v string) *MonitorContactGroupUpdateCmd
	GetName() *string
}

type MonitorContactGroupUpdateCmd struct {
	// example:
	//
	// 121321412341
	AccountId  *string  `json:"accountId,omitempty" xml:"accountId,omitempty"`
	ContactIds []*int64 `json:"contactIds,omitempty" xml:"contactIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 尚仁
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s MonitorContactGroupUpdateCmd) String() string {
	return dara.Prettify(s)
}

func (s MonitorContactGroupUpdateCmd) GoString() string {
	return s.String()
}

func (s *MonitorContactGroupUpdateCmd) GetAccountId() *string {
	return s.AccountId
}

func (s *MonitorContactGroupUpdateCmd) GetContactIds() []*int64 {
	return s.ContactIds
}

func (s *MonitorContactGroupUpdateCmd) GetId() *int64 {
	return s.Id
}

func (s *MonitorContactGroupUpdateCmd) GetName() *string {
	return s.Name
}

func (s *MonitorContactGroupUpdateCmd) SetAccountId(v string) *MonitorContactGroupUpdateCmd {
	s.AccountId = &v
	return s
}

func (s *MonitorContactGroupUpdateCmd) SetContactIds(v []*int64) *MonitorContactGroupUpdateCmd {
	s.ContactIds = v
	return s
}

func (s *MonitorContactGroupUpdateCmd) SetId(v int64) *MonitorContactGroupUpdateCmd {
	s.Id = &v
	return s
}

func (s *MonitorContactGroupUpdateCmd) SetName(v string) *MonitorContactGroupUpdateCmd {
	s.Name = &v
	return s
}

func (s *MonitorContactGroupUpdateCmd) Validate() error {
	return dara.Validate(s)
}
