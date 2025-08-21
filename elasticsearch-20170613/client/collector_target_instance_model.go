// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCollectorTargetInstance interface {
	dara.Model
	String() string
	GoString() string
	SetConfigType(v string) *CollectorTargetInstance
	GetConfigType() *string
	SetEnableMonitoring(v bool) *CollectorTargetInstance
	GetEnableMonitoring() *bool
	SetHosts(v []*string) *CollectorTargetInstance
	GetHosts() []*string
	SetInstanceId(v string) *CollectorTargetInstance
	GetInstanceId() *string
	SetInstanceType(v string) *CollectorTargetInstance
	GetInstanceType() *string
	SetPassword(v string) *CollectorTargetInstance
	GetPassword() *string
	SetProtocol(v string) *CollectorTargetInstance
	GetProtocol() *string
	SetUserName(v string) *CollectorTargetInstance
	GetUserName() *string
}

type CollectorTargetInstance struct {
	// This parameter is required.
	//
	// example:
	//
	// collectorTargetInstance
	ConfigType *string `json:"configType,omitempty" xml:"configType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	EnableMonitoring *bool     `json:"enableMonitoring,omitempty" xml:"enableMonitoring,omitempty"`
	Hosts            []*string `json:"hosts,omitempty" xml:"hosts,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// es-cn-ks8x****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// elasticsearch
	InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// password
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// username
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s CollectorTargetInstance) String() string {
	return dara.Prettify(s)
}

func (s CollectorTargetInstance) GoString() string {
	return s.String()
}

func (s *CollectorTargetInstance) GetConfigType() *string {
	return s.ConfigType
}

func (s *CollectorTargetInstance) GetEnableMonitoring() *bool {
	return s.EnableMonitoring
}

func (s *CollectorTargetInstance) GetHosts() []*string {
	return s.Hosts
}

func (s *CollectorTargetInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CollectorTargetInstance) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CollectorTargetInstance) GetPassword() *string {
	return s.Password
}

func (s *CollectorTargetInstance) GetProtocol() *string {
	return s.Protocol
}

func (s *CollectorTargetInstance) GetUserName() *string {
	return s.UserName
}

func (s *CollectorTargetInstance) SetConfigType(v string) *CollectorTargetInstance {
	s.ConfigType = &v
	return s
}

func (s *CollectorTargetInstance) SetEnableMonitoring(v bool) *CollectorTargetInstance {
	s.EnableMonitoring = &v
	return s
}

func (s *CollectorTargetInstance) SetHosts(v []*string) *CollectorTargetInstance {
	s.Hosts = v
	return s
}

func (s *CollectorTargetInstance) SetInstanceId(v string) *CollectorTargetInstance {
	s.InstanceId = &v
	return s
}

func (s *CollectorTargetInstance) SetInstanceType(v string) *CollectorTargetInstance {
	s.InstanceType = &v
	return s
}

func (s *CollectorTargetInstance) SetPassword(v string) *CollectorTargetInstance {
	s.Password = &v
	return s
}

func (s *CollectorTargetInstance) SetProtocol(v string) *CollectorTargetInstance {
	s.Protocol = &v
	return s
}

func (s *CollectorTargetInstance) SetUserName(v string) *CollectorTargetInstance {
	s.UserName = &v
	return s
}

func (s *CollectorTargetInstance) Validate() error {
	return dara.Validate(s)
}
