// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLoginBaseConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *ModifyLoginBaseConfigRequest
	GetConfig() *string
	SetTarget(v string) *ModifyLoginBaseConfigRequest
	GetTarget() *string
	SetType(v string) *ModifyLoginBaseConfigRequest
	GetType() *string
}

type ModifyLoginBaseConfigRequest struct {
	// The details of the configuration that is used to detect unusual logons to your servers. The value of this parameter is in the JSON format and contains the following fields:
	//
	// 	- **totalCount**: the total number of servers.
	//
	// 	- **uuidCount**: the number of servers to which the configuration is applied.
	//
	// 	- **id**: the ID of the configuration.
	//
	// 	- **location**: the common logon location.
	//
	// > You must specify this field if the Type parameter is set to login_common_location.
	//
	// 	- **ip**: the common logon IP address.
	//
	// > You must specify this field if the Type parameter is set to login_common_ip.
	//
	// 	- **endTime**: the end time of the common logon time range.
	//
	// > You must specify this field if the Type parameter is set to login_common_time.
	//
	// 	- **startTime**: the start time of the common logon time range.
	//
	// > You must specify this field if the Type parameter is set to login_common_time.
	//
	// 	- **account**: the common logon account.
	//
	// > You must specify this field if the Type parameter is set to login_common_account.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"totalCount":174,"uuidCount":4,"location":"Montenegro","id":0}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The details of the server to which the configuration is applied. The value of this parameter is in the JSON format and contains the following fields:
	//
	// 	- **Target**: the UUID of the server.
	//
	// 	- **targetType**: the type of the server to which the configuration is applied. Valid values:
	//
	//     	- **uuid**: a server
	//
	//     	- **groupId**: a server group
	//
	// 	- **flag**: the operation that you want to perform on the server. Valid values:
	//
	//     	- **del**: removes the server from the configuration.
	//
	//     	- **add**: adds the server to the configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"target":"inet-7c676676-06fa-442e-90fb-b802e5d6****","targetType":"uuid","flag":"add"}]
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The logon type of the configuration to modify. Valid values:
	//
	// 	- **login_common_location**: common logon location
	//
	// 	- **login_common_ip**: common logon IP address
	//
	// 	- **login_common_time**: common logon time range
	//
	// 	- **login_common_account**: common logon account
	//
	// This parameter is required.
	//
	// example:
	//
	// login_common_location
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifyLoginBaseConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyLoginBaseConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyLoginBaseConfigRequest) GetConfig() *string {
	return s.Config
}

func (s *ModifyLoginBaseConfigRequest) GetTarget() *string {
	return s.Target
}

func (s *ModifyLoginBaseConfigRequest) GetType() *string {
	return s.Type
}

func (s *ModifyLoginBaseConfigRequest) SetConfig(v string) *ModifyLoginBaseConfigRequest {
	s.Config = &v
	return s
}

func (s *ModifyLoginBaseConfigRequest) SetTarget(v string) *ModifyLoginBaseConfigRequest {
	s.Target = &v
	return s
}

func (s *ModifyLoginBaseConfigRequest) SetType(v string) *ModifyLoginBaseConfigRequest {
	s.Type = &v
	return s
}

func (s *ModifyLoginBaseConfigRequest) Validate() error {
	return dara.Validate(s)
}
