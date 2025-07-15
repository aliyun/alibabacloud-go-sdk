// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTrafficControlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiDefault(v int32) *ModifyTrafficControlRequest
	GetApiDefault() *int32
	SetAppDefault(v int32) *ModifyTrafficControlRequest
	GetAppDefault() *int32
	SetDescription(v string) *ModifyTrafficControlRequest
	GetDescription() *string
	SetSecurityToken(v string) *ModifyTrafficControlRequest
	GetSecurityToken() *string
	SetTrafficControlId(v string) *ModifyTrafficControlRequest
	GetTrafficControlId() *string
	SetTrafficControlName(v string) *ModifyTrafficControlRequest
	GetTrafficControlName() *string
	SetTrafficControlUnit(v string) *ModifyTrafficControlRequest
	GetTrafficControlUnit() *string
	SetUserDefault(v int32) *ModifyTrafficControlRequest
	GetUserDefault() *int32
}

type ModifyTrafficControlRequest struct {
	// The default throttling value for each API.
	//
	// example:
	//
	// 10000
	ApiDefault *int32 `json:"ApiDefault,omitempty" xml:"ApiDefault,omitempty"`
	// The default throttling value for each app.
	//
	// example:
	//
	// 10000
	AppDefault *int32 `json:"AppDefault,omitempty" xml:"AppDefault,omitempty"`
	// The description of the throttling policy.
	//
	// example:
	//
	// ThrottlingTestDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The security token included in the WebSocket request header. The system uses this token to authenticate the request.
	//
	// example:
	//
	// 4223a10e-eed3-46a6-8b7c-23003f488153
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The ID of the throttling policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// tf123456
	TrafficControlId *string `json:"TrafficControlId,omitempty" xml:"TrafficControlId,omitempty"`
	// The throttling policy name. The name must be 4 to 50 characters in length and can contain letters, digits, and underscores (_). It cannot start with an underscore.
	//
	// example:
	//
	// ThrottlingTest
	TrafficControlName *string `json:"TrafficControlName,omitempty" xml:"TrafficControlName,omitempty"`
	// The unit to be used in the throttling policy. Valid values:
	//
	// 	- **SECOND**
	//
	// 	- **MINUTE**
	//
	// 	- **HOUR**
	//
	// 	- **DAY**
	//
	// example:
	//
	// HOUR
	TrafficControlUnit *string `json:"TrafficControlUnit,omitempty" xml:"TrafficControlUnit,omitempty"`
	// The default throttling value for each user.
	//
	// example:
	//
	// 10000
	UserDefault *int32 `json:"UserDefault,omitempty" xml:"UserDefault,omitempty"`
}

func (s ModifyTrafficControlRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyTrafficControlRequest) GoString() string {
	return s.String()
}

func (s *ModifyTrafficControlRequest) GetApiDefault() *int32 {
	return s.ApiDefault
}

func (s *ModifyTrafficControlRequest) GetAppDefault() *int32 {
	return s.AppDefault
}

func (s *ModifyTrafficControlRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyTrafficControlRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyTrafficControlRequest) GetTrafficControlId() *string {
	return s.TrafficControlId
}

func (s *ModifyTrafficControlRequest) GetTrafficControlName() *string {
	return s.TrafficControlName
}

func (s *ModifyTrafficControlRequest) GetTrafficControlUnit() *string {
	return s.TrafficControlUnit
}

func (s *ModifyTrafficControlRequest) GetUserDefault() *int32 {
	return s.UserDefault
}

func (s *ModifyTrafficControlRequest) SetApiDefault(v int32) *ModifyTrafficControlRequest {
	s.ApiDefault = &v
	return s
}

func (s *ModifyTrafficControlRequest) SetAppDefault(v int32) *ModifyTrafficControlRequest {
	s.AppDefault = &v
	return s
}

func (s *ModifyTrafficControlRequest) SetDescription(v string) *ModifyTrafficControlRequest {
	s.Description = &v
	return s
}

func (s *ModifyTrafficControlRequest) SetSecurityToken(v string) *ModifyTrafficControlRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyTrafficControlRequest) SetTrafficControlId(v string) *ModifyTrafficControlRequest {
	s.TrafficControlId = &v
	return s
}

func (s *ModifyTrafficControlRequest) SetTrafficControlName(v string) *ModifyTrafficControlRequest {
	s.TrafficControlName = &v
	return s
}

func (s *ModifyTrafficControlRequest) SetTrafficControlUnit(v string) *ModifyTrafficControlRequest {
	s.TrafficControlUnit = &v
	return s
}

func (s *ModifyTrafficControlRequest) SetUserDefault(v int32) *ModifyTrafficControlRequest {
	s.UserDefault = &v
	return s
}

func (s *ModifyTrafficControlRequest) Validate() error {
	return dara.Validate(s)
}
