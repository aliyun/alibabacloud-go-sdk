// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonitorContactGroupCreateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *MonitorContactGroupCreateCmd
	GetAccountId() *string
	SetContactIds(v []*int64) *MonitorContactGroupCreateCmd
	GetContactIds() []*int64
	SetEnterpriseId(v int64) *MonitorContactGroupCreateCmd
	GetEnterpriseId() *int64
	SetName(v string) *MonitorContactGroupCreateCmd
	GetName() *string
}

type MonitorContactGroupCreateCmd struct {
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
	EnterpriseId *int64 `json:"enterpriseId,omitempty" xml:"enterpriseId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 尚仁
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s MonitorContactGroupCreateCmd) String() string {
	return dara.Prettify(s)
}

func (s MonitorContactGroupCreateCmd) GoString() string {
	return s.String()
}

func (s *MonitorContactGroupCreateCmd) GetAccountId() *string {
	return s.AccountId
}

func (s *MonitorContactGroupCreateCmd) GetContactIds() []*int64 {
	return s.ContactIds
}

func (s *MonitorContactGroupCreateCmd) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *MonitorContactGroupCreateCmd) GetName() *string {
	return s.Name
}

func (s *MonitorContactGroupCreateCmd) SetAccountId(v string) *MonitorContactGroupCreateCmd {
	s.AccountId = &v
	return s
}

func (s *MonitorContactGroupCreateCmd) SetContactIds(v []*int64) *MonitorContactGroupCreateCmd {
	s.ContactIds = v
	return s
}

func (s *MonitorContactGroupCreateCmd) SetEnterpriseId(v int64) *MonitorContactGroupCreateCmd {
	s.EnterpriseId = &v
	return s
}

func (s *MonitorContactGroupCreateCmd) SetName(v string) *MonitorContactGroupCreateCmd {
	s.Name = &v
	return s
}

func (s *MonitorContactGroupCreateCmd) Validate() error {
	return dara.Validate(s)
}
