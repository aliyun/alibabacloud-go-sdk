// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrivilegeCmd interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *CreatePrivilegeCmd
	GetAccountId() *string
	SetRoleId(v int64) *CreatePrivilegeCmd
	GetRoleId() *int64
	SetStrategyItems(v []*StrategyItem) *CreatePrivilegeCmd
	GetStrategyItems() []*StrategyItem
	SetStrategys(v string) *CreatePrivilegeCmd
	GetStrategys() *string
}

type CreatePrivilegeCmd struct {
	AccountId     *string         `json:"accountId,omitempty" xml:"accountId,omitempty"`
	RoleId        *int64          `json:"roleId,omitempty" xml:"roleId,omitempty"`
	StrategyItems []*StrategyItem `json:"strategyItems,omitempty" xml:"strategyItems,omitempty" type:"Repeated"`
	Strategys     *string         `json:"strategys,omitempty" xml:"strategys,omitempty"`
}

func (s CreatePrivilegeCmd) String() string {
	return dara.Prettify(s)
}

func (s CreatePrivilegeCmd) GoString() string {
	return s.String()
}

func (s *CreatePrivilegeCmd) GetAccountId() *string {
	return s.AccountId
}

func (s *CreatePrivilegeCmd) GetRoleId() *int64 {
	return s.RoleId
}

func (s *CreatePrivilegeCmd) GetStrategyItems() []*StrategyItem {
	return s.StrategyItems
}

func (s *CreatePrivilegeCmd) GetStrategys() *string {
	return s.Strategys
}

func (s *CreatePrivilegeCmd) SetAccountId(v string) *CreatePrivilegeCmd {
	s.AccountId = &v
	return s
}

func (s *CreatePrivilegeCmd) SetRoleId(v int64) *CreatePrivilegeCmd {
	s.RoleId = &v
	return s
}

func (s *CreatePrivilegeCmd) SetStrategyItems(v []*StrategyItem) *CreatePrivilegeCmd {
	s.StrategyItems = v
	return s
}

func (s *CreatePrivilegeCmd) SetStrategys(v string) *CreatePrivilegeCmd {
	s.Strategys = &v
	return s
}

func (s *CreatePrivilegeCmd) Validate() error {
	if s.StrategyItems != nil {
		for _, item := range s.StrategyItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
