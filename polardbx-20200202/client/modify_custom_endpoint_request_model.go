// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCustomEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomEndpointId(v string) *ModifyCustomEndpointRequest
	GetCustomEndpointId() *string
	SetDBInstanceName(v string) *ModifyCustomEndpointRequest
	GetDBInstanceName() *string
	SetName(v string) *ModifyCustomEndpointRequest
	GetName() *string
	SetNodeAutoEnter(v bool) *ModifyCustomEndpointRequest
	GetNodeAutoEnter() *bool
	SetNodeIds(v string) *ModifyCustomEndpointRequest
	GetNodeIds() *string
	SetNodeRole(v string) *ModifyCustomEndpointRequest
	GetNodeRole() *string
	SetRegionId(v string) *ModifyCustomEndpointRequest
	GetRegionId() *string
}

type ModifyCustomEndpointRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cep-002
	CustomEndpointId *string `json:"CustomEndpointId,omitempty" xml:"CustomEndpointId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mydatabase
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// secondary-endpoint
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// true
	NodeAutoEnter *bool `json:"NodeAutoEnter,omitempty" xml:"NodeAutoEnter,omitempty"`
	// node ids
	//
	// example:
	//
	// node3
	NodeIds *string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty"`
	// example:
	//
	// master
	NodeRole *string `json:"NodeRole,omitempty" xml:"NodeRole,omitempty"`
	// example:
	//
	// cn-zhangjiakou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyCustomEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomEndpointRequest) GoString() string {
	return s.String()
}

func (s *ModifyCustomEndpointRequest) GetCustomEndpointId() *string {
	return s.CustomEndpointId
}

func (s *ModifyCustomEndpointRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifyCustomEndpointRequest) GetName() *string {
	return s.Name
}

func (s *ModifyCustomEndpointRequest) GetNodeAutoEnter() *bool {
	return s.NodeAutoEnter
}

func (s *ModifyCustomEndpointRequest) GetNodeIds() *string {
	return s.NodeIds
}

func (s *ModifyCustomEndpointRequest) GetNodeRole() *string {
	return s.NodeRole
}

func (s *ModifyCustomEndpointRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyCustomEndpointRequest) SetCustomEndpointId(v string) *ModifyCustomEndpointRequest {
	s.CustomEndpointId = &v
	return s
}

func (s *ModifyCustomEndpointRequest) SetDBInstanceName(v string) *ModifyCustomEndpointRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyCustomEndpointRequest) SetName(v string) *ModifyCustomEndpointRequest {
	s.Name = &v
	return s
}

func (s *ModifyCustomEndpointRequest) SetNodeAutoEnter(v bool) *ModifyCustomEndpointRequest {
	s.NodeAutoEnter = &v
	return s
}

func (s *ModifyCustomEndpointRequest) SetNodeIds(v string) *ModifyCustomEndpointRequest {
	s.NodeIds = &v
	return s
}

func (s *ModifyCustomEndpointRequest) SetNodeRole(v string) *ModifyCustomEndpointRequest {
	s.NodeRole = &v
	return s
}

func (s *ModifyCustomEndpointRequest) SetRegionId(v string) *ModifyCustomEndpointRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyCustomEndpointRequest) Validate() error {
	return dara.Validate(s)
}
