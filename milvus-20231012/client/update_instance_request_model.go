// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *UpdateInstanceRequest
	GetRegionId() *string
	SetAutoBackup(v bool) *UpdateInstanceRequest
	GetAutoBackup() *bool
	SetComponents(v []*UpdateInstanceRequestComponents) *UpdateInstanceRequest
	GetComponents() []*UpdateInstanceRequestComponents
	SetConfiguration(v string) *UpdateInstanceRequest
	GetConfiguration() *string
	SetHa(v bool) *UpdateInstanceRequest
	GetHa() *bool
	SetInstanceId(v string) *UpdateInstanceRequest
	GetInstanceId() *string
	SetInstanceName(v string) *UpdateInstanceRequest
	GetInstanceName() *string
	SetClientToken(v string) *UpdateInstanceRequest
	GetClientToken() *string
}

type UpdateInstanceRequest struct {
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// true
	AutoBackup *bool                              `json:"autoBackup,omitempty" xml:"autoBackup,omitempty"`
	Components []*UpdateInstanceRequestComponents `json:"components,omitempty" xml:"components,omitempty" type:"Repeated"`
	// example:
	//
	// rootCoord:\\n  maxDatabaseNum: 64
	Configuration *string `json:"configuration,omitempty" xml:"configuration,omitempty"`
	// example:
	//
	// true
	Ha *bool `json:"ha,omitempty" xml:"ha,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c-xxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// milvus-test
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	// example:
	//
	// xxx
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateInstanceRequest) GetAutoBackup() *bool {
	return s.AutoBackup
}

func (s *UpdateInstanceRequest) GetComponents() []*UpdateInstanceRequestComponents {
	return s.Components
}

func (s *UpdateInstanceRequest) GetConfiguration() *string {
	return s.Configuration
}

func (s *UpdateInstanceRequest) GetHa() *bool {
	return s.Ha
}

func (s *UpdateInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *UpdateInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateInstanceRequest) SetRegionId(v string) *UpdateInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateInstanceRequest) SetAutoBackup(v bool) *UpdateInstanceRequest {
	s.AutoBackup = &v
	return s
}

func (s *UpdateInstanceRequest) SetComponents(v []*UpdateInstanceRequestComponents) *UpdateInstanceRequest {
	s.Components = v
	return s
}

func (s *UpdateInstanceRequest) SetConfiguration(v string) *UpdateInstanceRequest {
	s.Configuration = &v
	return s
}

func (s *UpdateInstanceRequest) SetHa(v bool) *UpdateInstanceRequest {
	s.Ha = &v
	return s
}

func (s *UpdateInstanceRequest) SetInstanceId(v string) *UpdateInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceRequest) SetInstanceName(v string) *UpdateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *UpdateInstanceRequest) SetClientToken(v string) *UpdateInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateInstanceRequest) Validate() error {
	if s.Components != nil {
		for _, item := range s.Components {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateInstanceRequestComponents struct {
	// This parameter is required.
	//
	// example:
	//
	// 8
	CuNum *int32 `json:"cuNum,omitempty" xml:"cuNum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Replica *int32 `json:"replica,omitempty" xml:"replica,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// standalone
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateInstanceRequestComponents) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceRequestComponents) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequestComponents) GetCuNum() *int32 {
	return s.CuNum
}

func (s *UpdateInstanceRequestComponents) GetReplica() *int32 {
	return s.Replica
}

func (s *UpdateInstanceRequestComponents) GetType() *string {
	return s.Type
}

func (s *UpdateInstanceRequestComponents) SetCuNum(v int32) *UpdateInstanceRequestComponents {
	s.CuNum = &v
	return s
}

func (s *UpdateInstanceRequestComponents) SetReplica(v int32) *UpdateInstanceRequestComponents {
	s.Replica = &v
	return s
}

func (s *UpdateInstanceRequestComponents) SetType(v string) *UpdateInstanceRequestComponents {
	s.Type = &v
	return s
}

func (s *UpdateInstanceRequestComponents) Validate() error {
	return dara.Validate(s)
}
