// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppIds(v string) *ConfigAppRequest
	GetAppIds() *string
	SetEnable(v string) *ConfigAppRequest
	GetEnable() *string
	SetRegionId(v string) *ConfigAppRequest
	GetRegionId() *string
	SetType(v string) *ConfigAppRequest
	GetType() *string
}

type ConfigAppRequest struct {
	// The process identifier (PID) of the application. Separate multiple PIDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// iioe7jcnuk@582846f37******,atc889zkcf@d8deedfa9bf******
	AppIds *string `json:"AppIds,omitempty" xml:"AppIds,omitempty"`
	// Specifies whether to turn on or off the main switch of the ARMS agent. The monitoring stops after the switch is turned off. If you do not specify this parameter, the main switch status of the ARMS agent is queried.
	//
	// 	- `true`: turns on the switch
	//
	// 	- `false`: turns off the switch
	//
	// example:
	//
	// true
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the application. Set the value to **TRACE**.
	//
	// example:
	//
	// TRACE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ConfigAppRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigAppRequest) GoString() string {
	return s.String()
}

func (s *ConfigAppRequest) GetAppIds() *string {
	return s.AppIds
}

func (s *ConfigAppRequest) GetEnable() *string {
	return s.Enable
}

func (s *ConfigAppRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ConfigAppRequest) GetType() *string {
	return s.Type
}

func (s *ConfigAppRequest) SetAppIds(v string) *ConfigAppRequest {
	s.AppIds = &v
	return s
}

func (s *ConfigAppRequest) SetEnable(v string) *ConfigAppRequest {
	s.Enable = &v
	return s
}

func (s *ConfigAppRequest) SetRegionId(v string) *ConfigAppRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigAppRequest) SetType(v string) *ConfigAppRequest {
	s.Type = &v
	return s
}

func (s *ConfigAppRequest) Validate() error {
	return dara.Validate(s)
}
