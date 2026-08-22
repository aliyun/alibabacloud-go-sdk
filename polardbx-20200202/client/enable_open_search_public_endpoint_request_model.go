// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableOpenSearchPublicEndpointRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDBInstanceName(v string) *EnableOpenSearchPublicEndpointRequest
  GetDBInstanceName() *string 
  SetNodeType(v string) *EnableOpenSearchPublicEndpointRequest
  GetNodeType() *string 
  SetRegionId(v string) *EnableOpenSearchPublicEndpointRequest
  GetRegionId() *string 
}

type EnableOpenSearchPublicEndpointRequest struct {
  // The instance ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // pxc-********
  DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
  // The target node type. Valid values:
  // 
  // - **search**: data node. This is the default value.
  // 
  // - **dashboard**: dashboard node.
  // 
  // example:
  // 
  // search
  NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
  // The region ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-shanghai
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s EnableOpenSearchPublicEndpointRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableOpenSearchPublicEndpointRequest) GoString() string {
  return s.String()
}

func (s *EnableOpenSearchPublicEndpointRequest) GetDBInstanceName() *string  {
  return s.DBInstanceName
}

func (s *EnableOpenSearchPublicEndpointRequest) GetNodeType() *string  {
  return s.NodeType
}

func (s *EnableOpenSearchPublicEndpointRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableOpenSearchPublicEndpointRequest) SetDBInstanceName(v string) *EnableOpenSearchPublicEndpointRequest {
  s.DBInstanceName = &v
  return s
}

func (s *EnableOpenSearchPublicEndpointRequest) SetNodeType(v string) *EnableOpenSearchPublicEndpointRequest {
  s.NodeType = &v
  return s
}

func (s *EnableOpenSearchPublicEndpointRequest) SetRegionId(v string) *EnableOpenSearchPublicEndpointRequest {
  s.RegionId = &v
  return s
}

func (s *EnableOpenSearchPublicEndpointRequest) Validate() error {
  return dara.Validate(s)
}

