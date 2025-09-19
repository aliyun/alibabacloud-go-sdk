// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLoginSwitchConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetItem(v string) *ModifyLoginSwitchConfigRequest
	GetItem() *string
	SetStatus(v int32) *ModifyLoginSwitchConfigRequest
	GetStatus() *int32
}

type ModifyLoginSwitchConfigRequest struct {
	// The type of the logon security settings that you want to enable or disable. Valid values:
	//
	// 	- **login_common_ip**: unapproved logon IP addresses
	//
	// 	- **login_common_time**: unapproved logon time ranges
	//
	// 	- **login_common_account**: unapproved logon accounts
	//
	// This parameter is required.
	//
	// example:
	//
	// login_common_account
	Item *string `json:"Item,omitempty" xml:"Item,omitempty"`
	// Specifies whether to enable the logon security settings. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyLoginSwitchConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyLoginSwitchConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyLoginSwitchConfigRequest) GetItem() *string {
	return s.Item
}

func (s *ModifyLoginSwitchConfigRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ModifyLoginSwitchConfigRequest) SetItem(v string) *ModifyLoginSwitchConfigRequest {
	s.Item = &v
	return s
}

func (s *ModifyLoginSwitchConfigRequest) SetStatus(v int32) *ModifyLoginSwitchConfigRequest {
	s.Status = &v
	return s
}

func (s *ModifyLoginSwitchConfigRequest) Validate() error {
	return dara.Validate(s)
}
