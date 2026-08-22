// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableOpenSearchPublicEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DisableOpenSearchPublicEndpointRequest
	GetDBInstanceName() *string
	SetNodeType(v string) *DisableOpenSearchPublicEndpointRequest
	GetNodeType() *string
	SetRegionId(v string) *DisableOpenSearchPublicEndpointRequest
	GetRegionId() *string
}

type DisableOpenSearchPublicEndpointRequest struct {
	// The name of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The type of the target node. Valid values:
	//
	// - **search**: data node. This is the default value.
	//
	// - **dashboard**: dashboard node.
	//
	// example:
	//
	// search
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The region in which the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DisableOpenSearchPublicEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableOpenSearchPublicEndpointRequest) GoString() string {
	return s.String()
}

func (s *DisableOpenSearchPublicEndpointRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DisableOpenSearchPublicEndpointRequest) GetNodeType() *string {
	return s.NodeType
}

func (s *DisableOpenSearchPublicEndpointRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DisableOpenSearchPublicEndpointRequest) SetDBInstanceName(v string) *DisableOpenSearchPublicEndpointRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DisableOpenSearchPublicEndpointRequest) SetNodeType(v string) *DisableOpenSearchPublicEndpointRequest {
	s.NodeType = &v
	return s
}

func (s *DisableOpenSearchPublicEndpointRequest) SetRegionId(v string) *DisableOpenSearchPublicEndpointRequest {
	s.RegionId = &v
	return s
}

func (s *DisableOpenSearchPublicEndpointRequest) Validate() error {
	return dara.Validate(s)
}
