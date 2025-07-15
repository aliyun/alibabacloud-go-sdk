// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrafficControlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiDefault(v int32) *CreateTrafficControlRequest
	GetApiDefault() *int32
	SetAppDefault(v int32) *CreateTrafficControlRequest
	GetAppDefault() *int32
	SetDescription(v string) *CreateTrafficControlRequest
	GetDescription() *string
	SetSecurityToken(v string) *CreateTrafficControlRequest
	GetSecurityToken() *string
	SetTrafficControlName(v string) *CreateTrafficControlRequest
	GetTrafficControlName() *string
	SetTrafficControlUnit(v string) *CreateTrafficControlRequest
	GetTrafficControlUnit() *string
	SetUserDefault(v int32) *CreateTrafficControlRequest
	GetUserDefault() *int32
}

type CreateTrafficControlRequest struct {
	// The default throttling value for each API.
	//
	// This parameter is required.
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
	// 436fa39b-b3b9-40c5-ae5d-ce3e000e38c5
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The name of the throttling policy. The name must be 4 to 50 characters in length and can contain letters, digits, and underscores (_). It cannot start with an underscore.
	//
	// This parameter is required.
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
	// This parameter is required.
	//
	// example:
	//
	// MINUTE
	TrafficControlUnit *string `json:"TrafficControlUnit,omitempty" xml:"TrafficControlUnit,omitempty"`
	// The default throttling value for each user.
	//
	// example:
	//
	// 10000
	UserDefault *int32 `json:"UserDefault,omitempty" xml:"UserDefault,omitempty"`
}

func (s CreateTrafficControlRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTrafficControlRequest) GoString() string {
	return s.String()
}

func (s *CreateTrafficControlRequest) GetApiDefault() *int32 {
	return s.ApiDefault
}

func (s *CreateTrafficControlRequest) GetAppDefault() *int32 {
	return s.AppDefault
}

func (s *CreateTrafficControlRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateTrafficControlRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateTrafficControlRequest) GetTrafficControlName() *string {
	return s.TrafficControlName
}

func (s *CreateTrafficControlRequest) GetTrafficControlUnit() *string {
	return s.TrafficControlUnit
}

func (s *CreateTrafficControlRequest) GetUserDefault() *int32 {
	return s.UserDefault
}

func (s *CreateTrafficControlRequest) SetApiDefault(v int32) *CreateTrafficControlRequest {
	s.ApiDefault = &v
	return s
}

func (s *CreateTrafficControlRequest) SetAppDefault(v int32) *CreateTrafficControlRequest {
	s.AppDefault = &v
	return s
}

func (s *CreateTrafficControlRequest) SetDescription(v string) *CreateTrafficControlRequest {
	s.Description = &v
	return s
}

func (s *CreateTrafficControlRequest) SetSecurityToken(v string) *CreateTrafficControlRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateTrafficControlRequest) SetTrafficControlName(v string) *CreateTrafficControlRequest {
	s.TrafficControlName = &v
	return s
}

func (s *CreateTrafficControlRequest) SetTrafficControlUnit(v string) *CreateTrafficControlRequest {
	s.TrafficControlUnit = &v
	return s
}

func (s *CreateTrafficControlRequest) SetUserDefault(v int32) *CreateTrafficControlRequest {
	s.UserDefault = &v
	return s
}

func (s *CreateTrafficControlRequest) Validate() error {
	return dara.Validate(s)
}
