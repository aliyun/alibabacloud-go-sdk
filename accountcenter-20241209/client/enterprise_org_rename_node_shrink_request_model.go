// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseOrgRenameNodeShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppName(v string) *EnterpriseOrgRenameNodeShrinkRequest
  GetAppName() *string 
  SetBizName(v string) *EnterpriseOrgRenameNodeShrinkRequest
  GetBizName() *string 
  SetExtShrink(v string) *EnterpriseOrgRenameNodeShrinkRequest
  GetExtShrink() *string 
  SetIsOpenApi(v bool) *EnterpriseOrgRenameNodeShrinkRequest
  GetIsOpenApi() *bool 
  SetNodeId(v string) *EnterpriseOrgRenameNodeShrinkRequest
  GetNodeId() *string 
  SetNodeName(v string) *EnterpriseOrgRenameNodeShrinkRequest
  GetNodeName() *string 
  SetNodeType(v string) *EnterpriseOrgRenameNodeShrinkRequest
  GetNodeType() *string 
  SetOrientedEcId(v string) *EnterpriseOrgRenameNodeShrinkRequest
  GetOrientedEcId() *string 
  SetOrientedLeId(v string) *EnterpriseOrgRenameNodeShrinkRequest
  GetOrientedLeId() *string 
  SetOrientedNbId(v string) *EnterpriseOrgRenameNodeShrinkRequest
  GetOrientedNbId() *string 
  SetShowCompleteInfo(v bool) *EnterpriseOrgRenameNodeShrinkRequest
  GetShowCompleteInfo() *bool 
  SetTreeId(v int64) *EnterpriseOrgRenameNodeShrinkRequest
  GetTreeId() *int64 
}

type EnterpriseOrgRenameNodeShrinkRequest struct {
  AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
  BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
  ExtShrink *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
  IsOpenApi *bool `json:"IsOpenApi,omitempty" xml:"IsOpenApi,omitempty"`
  NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
  NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
  NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
  OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
  OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
  OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
  ShowCompleteInfo *bool `json:"ShowCompleteInfo,omitempty" xml:"ShowCompleteInfo,omitempty"`
  TreeId *int64 `json:"TreeId,omitempty" xml:"TreeId,omitempty"`
}

func (s EnterpriseOrgRenameNodeShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseOrgRenameNodeShrinkRequest) GoString() string {
  return s.String()
}

func (s *EnterpriseOrgRenameNodeShrinkRequest) GetAppName() *string  {
  return s.AppName
}

func (s *EnterpriseOrgRenameNodeShrinkRequest) GetBizName() *string  {
  return s.BizName
}

func (s *EnterpriseOrgRenameNodeShrinkRequest) GetExtShrink() *string  {
  return s.ExtShrink
}

func (s *EnterpriseOrgRenameNodeShrinkRequest) GetIsOpenApi() *bool  {
  return s.IsOpenApi
}

func (s *EnterpriseOrgRenameNodeShrinkRequest) GetNodeId() *string  {
  return s.NodeId
}

func (s *EnterpriseOrgRenameNodeShrinkRequest) GetNodeName() *string  {
  return s.NodeName
}

func (s *EnterpriseOrgRenameNodeShrinkRequest) GetNodeType() *string  {
  return s.NodeType
}

func (s *EnterpriseOrgRenameNodeShrinkRequest) GetOrientedEcId() *string  {
  return s.OrientedEcId
}

func (s *EnterpriseOrgRenameNodeShrinkRequest) GetOrientedLeId() *string  {
  return s.OrientedLeId
}

func (s *EnterpriseOrgRenameNodeShrinkRequest) GetOrientedNbId() *string  {
  return s.OrientedNbId
}

func (s *EnterpriseOrgRenameNodeShrinkRequest) GetShowCompleteInfo() *bool  {
  return s.ShowCompleteInfo
}

func (s *EnterpriseOrgRenameNodeShrinkRequest) GetTreeId() *int64  {
  return s.TreeId
}

func (s *EnterpriseOrgRenameNodeShrinkRequest) SetAppName(v string) *EnterpriseOrgRenameNodeShrinkRequest {
  s.AppName = &v
  return s
}

func (s *EnterpriseOrgRenameNodeShrinkRequest) SetBizName(v string) *EnterpriseOrgRenameNodeShrinkRequest {
  s.BizName = &v
  return s
}

func (s *EnterpriseOrgRenameNodeShrinkRequest) SetExtShrink(v string) *EnterpriseOrgRenameNodeShrinkRequest {
  s.ExtShrink = &v
  return s
}

func (s *EnterpriseOrgRenameNodeShrinkRequest) SetIsOpenApi(v bool) *EnterpriseOrgRenameNodeShrinkRequest {
  s.IsOpenApi = &v
  return s
}

func (s *EnterpriseOrgRenameNodeShrinkRequest) SetNodeId(v string) *EnterpriseOrgRenameNodeShrinkRequest {
  s.NodeId = &v
  return s
}

func (s *EnterpriseOrgRenameNodeShrinkRequest) SetNodeName(v string) *EnterpriseOrgRenameNodeShrinkRequest {
  s.NodeName = &v
  return s
}

func (s *EnterpriseOrgRenameNodeShrinkRequest) SetNodeType(v string) *EnterpriseOrgRenameNodeShrinkRequest {
  s.NodeType = &v
  return s
}

func (s *EnterpriseOrgRenameNodeShrinkRequest) SetOrientedEcId(v string) *EnterpriseOrgRenameNodeShrinkRequest {
  s.OrientedEcId = &v
  return s
}

func (s *EnterpriseOrgRenameNodeShrinkRequest) SetOrientedLeId(v string) *EnterpriseOrgRenameNodeShrinkRequest {
  s.OrientedLeId = &v
  return s
}

func (s *EnterpriseOrgRenameNodeShrinkRequest) SetOrientedNbId(v string) *EnterpriseOrgRenameNodeShrinkRequest {
  s.OrientedNbId = &v
  return s
}

func (s *EnterpriseOrgRenameNodeShrinkRequest) SetShowCompleteInfo(v bool) *EnterpriseOrgRenameNodeShrinkRequest {
  s.ShowCompleteInfo = &v
  return s
}

func (s *EnterpriseOrgRenameNodeShrinkRequest) SetTreeId(v int64) *EnterpriseOrgRenameNodeShrinkRequest {
  s.TreeId = &v
  return s
}

func (s *EnterpriseOrgRenameNodeShrinkRequest) Validate() error {
  return dara.Validate(s)
}

