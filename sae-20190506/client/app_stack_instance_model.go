// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppStackInstance interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v int64) *AppStackInstance
	GetCreateTime() *int64
	SetEndpoints(v []*AppStackInstanceEndpoints) *AppStackInstance
	GetEndpoints() []*AppStackInstanceEndpoints
	SetInstanceId(v string) *AppStackInstance
	GetInstanceId() *string
	SetInstanceName(v string) *AppStackInstance
	GetInstanceName() *string
	SetParameters(v []*AppStackInstanceParameters) *AppStackInstance
	GetParameters() []*AppStackInstanceParameters
	SetStackId(v string) *AppStackInstance
	GetStackId() *string
	SetStatus(v string) *AppStackInstance
	GetStatus() *string
	SetUpdateTime(v int64) *AppStackInstance
	GetUpdateTime() *int64
}

type AppStackInstance struct {
	// example:
	//
	// 1706518652
	CreateTime *int64                       `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Endpoints  []*AppStackInstanceEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	// example:
	//
	// i-789y
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// palworld
	InstanceName *string                       `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Parameters   []*AppStackInstanceParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// example:
	//
	// palworld
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// example:
	//
	// WAIT
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1706518652
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s AppStackInstance) String() string {
	return dara.Prettify(s)
}

func (s AppStackInstance) GoString() string {
	return s.String()
}

func (s *AppStackInstance) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *AppStackInstance) GetEndpoints() []*AppStackInstanceEndpoints {
	return s.Endpoints
}

func (s *AppStackInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AppStackInstance) GetInstanceName() *string {
	return s.InstanceName
}

func (s *AppStackInstance) GetParameters() []*AppStackInstanceParameters {
	return s.Parameters
}

func (s *AppStackInstance) GetStackId() *string {
	return s.StackId
}

func (s *AppStackInstance) GetStatus() *string {
	return s.Status
}

func (s *AppStackInstance) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *AppStackInstance) SetCreateTime(v int64) *AppStackInstance {
	s.CreateTime = &v
	return s
}

func (s *AppStackInstance) SetEndpoints(v []*AppStackInstanceEndpoints) *AppStackInstance {
	s.Endpoints = v
	return s
}

func (s *AppStackInstance) SetInstanceId(v string) *AppStackInstance {
	s.InstanceId = &v
	return s
}

func (s *AppStackInstance) SetInstanceName(v string) *AppStackInstance {
	s.InstanceName = &v
	return s
}

func (s *AppStackInstance) SetParameters(v []*AppStackInstanceParameters) *AppStackInstance {
	s.Parameters = v
	return s
}

func (s *AppStackInstance) SetStackId(v string) *AppStackInstance {
	s.StackId = &v
	return s
}

func (s *AppStackInstance) SetStatus(v string) *AppStackInstance {
	s.Status = &v
	return s
}

func (s *AppStackInstance) SetUpdateTime(v int64) *AppStackInstance {
	s.UpdateTime = &v
	return s
}

func (s *AppStackInstance) Validate() error {
	if s.Endpoints != nil {
		for _, item := range s.Endpoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Parameters != nil {
		for _, item := range s.Parameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AppStackInstanceEndpoints struct {
	// example:
	//
	// 127.0.0.1:8211
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// example:
	//
	// game
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// UDP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s AppStackInstanceEndpoints) String() string {
	return dara.Prettify(s)
}

func (s AppStackInstanceEndpoints) GoString() string {
	return s.String()
}

func (s *AppStackInstanceEndpoints) GetAddress() *string {
	return s.Address
}

func (s *AppStackInstanceEndpoints) GetName() *string {
	return s.Name
}

func (s *AppStackInstanceEndpoints) GetProtocol() *string {
	return s.Protocol
}

func (s *AppStackInstanceEndpoints) SetAddress(v string) *AppStackInstanceEndpoints {
	s.Address = &v
	return s
}

func (s *AppStackInstanceEndpoints) SetName(v string) *AppStackInstanceEndpoints {
	s.Name = &v
	return s
}

func (s *AppStackInstanceEndpoints) SetProtocol(v string) *AppStackInstanceEndpoints {
	s.Protocol = &v
	return s
}

func (s *AppStackInstanceEndpoints) Validate() error {
	return dara.Validate(s)
}

type AppStackInstanceParameters struct {
	// example:
	//
	// regionId
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// cn-hangzhou
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AppStackInstanceParameters) String() string {
	return dara.Prettify(s)
}

func (s AppStackInstanceParameters) GoString() string {
	return s.String()
}

func (s *AppStackInstanceParameters) GetName() *string {
	return s.Name
}

func (s *AppStackInstanceParameters) GetValue() *string {
	return s.Value
}

func (s *AppStackInstanceParameters) SetName(v string) *AppStackInstanceParameters {
	s.Name = &v
	return s
}

func (s *AppStackInstanceParameters) SetValue(v string) *AppStackInstanceParameters {
	s.Value = &v
	return s
}

func (s *AppStackInstanceParameters) Validate() error {
	return dara.Validate(s)
}
