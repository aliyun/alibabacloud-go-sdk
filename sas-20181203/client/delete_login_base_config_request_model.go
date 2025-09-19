// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLoginBaseConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *DeleteLoginBaseConfigRequest
	GetConfig() *string
	SetTarget(v string) *DeleteLoginBaseConfigRequest
	GetTarget() *string
	SetType(v string) *DeleteLoginBaseConfigRequest
	GetType() *string
}

type DeleteLoginBaseConfigRequest struct {
	// The content of the logon security settings to delete. The content varies based on the type of the logon security settings. Valid values:
	//
	// 	- **login_common_ip**: approved logon IP addresses
	//
	// Example: {"ip":"10.23.23.23"}.
	//
	// 	- **login_common_time**: approved logon time ranges
	//
	// Example: {"startTime":"06:00:00","endTime":"16:00:00"}.
	//
	// 	- **login_common_account**: approved logon accounts
	//
	// Example: {"account":"test_account_001"}.
	//
	// 	- **login_common_location**: approved logon locations
	//
	// Example: {"location":"Shanghai"}.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"startTime":"06:00:00","endTime":"16:00:00"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The UUID of the server whose logon security settings you want to delete.
	//
	// > You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to query the UUIDs of servers.
	//
	// example:
	//
	// 4fe8e1cd-3c37-4851-b9de-124da32c****
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The type of the logon security settings to delete. Valid values:
	//
	// 	- **login_common_ip**: approved logon IP addresses
	//
	// 	- **login_common_time**: approved logon time ranges
	//
	// 	- **login_common_account**: approved logon accounts
	//
	// 	- **login_common_location**: approved logon locations
	//
	// This parameter is required.
	//
	// example:
	//
	// login_common_time
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DeleteLoginBaseConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLoginBaseConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteLoginBaseConfigRequest) GetConfig() *string {
	return s.Config
}

func (s *DeleteLoginBaseConfigRequest) GetTarget() *string {
	return s.Target
}

func (s *DeleteLoginBaseConfigRequest) GetType() *string {
	return s.Type
}

func (s *DeleteLoginBaseConfigRequest) SetConfig(v string) *DeleteLoginBaseConfigRequest {
	s.Config = &v
	return s
}

func (s *DeleteLoginBaseConfigRequest) SetTarget(v string) *DeleteLoginBaseConfigRequest {
	s.Target = &v
	return s
}

func (s *DeleteLoginBaseConfigRequest) SetType(v string) *DeleteLoginBaseConfigRequest {
	s.Type = &v
	return s
}

func (s *DeleteLoginBaseConfigRequest) Validate() error {
	return dara.Validate(s)
}
