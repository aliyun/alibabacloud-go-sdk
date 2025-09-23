// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigLayer4RuleAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *ConfigLayer4RuleAttributeRequest
	GetConfig() *string
	SetForwardProtocol(v string) *ConfigLayer4RuleAttributeRequest
	GetForwardProtocol() *string
	SetFrontendPort(v int32) *ConfigLayer4RuleAttributeRequest
	GetFrontendPort() *int32
	SetInstanceId(v string) *ConfigLayer4RuleAttributeRequest
	GetInstanceId() *string
	SetModule(v string) *ConfigLayer4RuleAttributeRequest
	GetModule() *string
}

type ConfigLayer4RuleAttributeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// {"Slimit":{"CpsEnable":1,"MaxconnEnable":1,"Cps":1,"Maxconn":1},"Sla":{"CpsEnable":1,"MaxconnEnable":1,"Cps":100,"Maxconn":1000},"PayloadLen":{"Min":0,"Max":6000}}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TCP
	ForwardProtocol *string `json:"ForwardProtocol,omitempty" xml:"ForwardProtocol,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 233
	FrontendPort *int32 `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-XXXXX
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Module     *string `json:"Module,omitempty" xml:"Module,omitempty"`
}

func (s ConfigLayer4RuleAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigLayer4RuleAttributeRequest) GoString() string {
	return s.String()
}

func (s *ConfigLayer4RuleAttributeRequest) GetConfig() *string {
	return s.Config
}

func (s *ConfigLayer4RuleAttributeRequest) GetForwardProtocol() *string {
	return s.ForwardProtocol
}

func (s *ConfigLayer4RuleAttributeRequest) GetFrontendPort() *int32 {
	return s.FrontendPort
}

func (s *ConfigLayer4RuleAttributeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ConfigLayer4RuleAttributeRequest) GetModule() *string {
	return s.Module
}

func (s *ConfigLayer4RuleAttributeRequest) SetConfig(v string) *ConfigLayer4RuleAttributeRequest {
	s.Config = &v
	return s
}

func (s *ConfigLayer4RuleAttributeRequest) SetForwardProtocol(v string) *ConfigLayer4RuleAttributeRequest {
	s.ForwardProtocol = &v
	return s
}

func (s *ConfigLayer4RuleAttributeRequest) SetFrontendPort(v int32) *ConfigLayer4RuleAttributeRequest {
	s.FrontendPort = &v
	return s
}

func (s *ConfigLayer4RuleAttributeRequest) SetInstanceId(v string) *ConfigLayer4RuleAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *ConfigLayer4RuleAttributeRequest) SetModule(v string) *ConfigLayer4RuleAttributeRequest {
	s.Module = &v
	return s
}

func (s *ConfigLayer4RuleAttributeRequest) Validate() error {
	return dara.Validate(s)
}
