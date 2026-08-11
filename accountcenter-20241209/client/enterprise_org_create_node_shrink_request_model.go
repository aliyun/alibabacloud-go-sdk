// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseOrgCreateNodeShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppName(v string) *EnterpriseOrgCreateNodeShrinkRequest
  GetAppName() *string 
  SetBizName(v string) *EnterpriseOrgCreateNodeShrinkRequest
  GetBizName() *string 
  SetExtShrink(v string) *EnterpriseOrgCreateNodeShrinkRequest
  GetExtShrink() *string 
  SetIsOpenApi(v bool) *EnterpriseOrgCreateNodeShrinkRequest
  GetIsOpenApi() *bool 
  SetNodeId(v string) *EnterpriseOrgCreateNodeShrinkRequest
  GetNodeId() *string 
  SetNodeName(v string) *EnterpriseOrgCreateNodeShrinkRequest
  GetNodeName() *string 
  SetNodeType(v string) *EnterpriseOrgCreateNodeShrinkRequest
  GetNodeType() *string 
  SetOrientedEcId(v string) *EnterpriseOrgCreateNodeShrinkRequest
  GetOrientedEcId() *string 
  SetOrientedLeId(v string) *EnterpriseOrgCreateNodeShrinkRequest
  GetOrientedLeId() *string 
  SetOrientedNbId(v string) *EnterpriseOrgCreateNodeShrinkRequest
  GetOrientedNbId() *string 
  SetParentNodeId(v string) *EnterpriseOrgCreateNodeShrinkRequest
  GetParentNodeId() *string 
  SetParentNodeType(v string) *EnterpriseOrgCreateNodeShrinkRequest
  GetParentNodeType() *string 
  SetShowCompleteInfo(v bool) *EnterpriseOrgCreateNodeShrinkRequest
  GetShowCompleteInfo() *bool 
  SetTreeId(v int64) *EnterpriseOrgCreateNodeShrinkRequest
  GetTreeId() *int64 
}

type EnterpriseOrgCreateNodeShrinkRequest struct {
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
  ParentNodeId *string `json:"ParentNodeId,omitempty" xml:"ParentNodeId,omitempty"`
  ParentNodeType *string `json:"ParentNodeType,omitempty" xml:"ParentNodeType,omitempty"`
  ShowCompleteInfo *bool `json:"ShowCompleteInfo,omitempty" xml:"ShowCompleteInfo,omitempty"`
  TreeId *int64 `json:"TreeId,omitempty" xml:"TreeId,omitempty"`
}

func (s EnterpriseOrgCreateNodeShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseOrgCreateNodeShrinkRequest) GoString() string {
  return s.String()
}

func (s *EnterpriseOrgCreateNodeShrinkRequest) GetAppName() *string  {
  return s.AppName
}

func (s *EnterpriseOrgCreateNodeShrinkRequest) GetBizName() *string  {
  return s.BizName
}

func (s *EnterpriseOrgCreateNodeShrinkRequest) GetExtShrink() *string  {
  return s.ExtShrink
}

func (s *EnterpriseOrgCreateNodeShrinkRequest) GetIsOpenApi() *bool  {
  return s.IsOpenApi
}

func (s *EnterpriseOrgCreateNodeShrinkRequest) GetNodeId() *string  {
  return s.NodeId
}

func (s *EnterpriseOrgCreateNodeShrinkRequest) GetNodeName() *string  {
  return s.NodeName
}

func (s *EnterpriseOrgCreateNodeShrinkRequest) GetNodeType() *string  {
  return s.NodeType
}

func (s *EnterpriseOrgCreateNodeShrinkRequest) GetOrientedEcId() *string  {
  return s.OrientedEcId
}

func (s *EnterpriseOrgCreateNodeShrinkRequest) GetOrientedLeId() *string  {
  return s.OrientedLeId
}

func (s *EnterpriseOrgCreateNodeShrinkRequest) GetOrientedNbId() *string  {
  return s.OrientedNbId
}

func (s *EnterpriseOrgCreateNodeShrinkRequest) GetParentNodeId() *string  {
  return s.ParentNodeId
}

func (s *EnterpriseOrgCreateNodeShrinkRequest) GetParentNodeType() *string  {
  return s.ParentNodeType
}

func (s *EnterpriseOrgCreateNodeShrinkRequest) GetShowCompleteInfo() *bool  {
  return s.ShowCompleteInfo
}

func (s *EnterpriseOrgCreateNodeShrinkRequest) GetTreeId() *int64  {
  return s.TreeId
}

func (s *EnterpriseOrgCreateNodeShrinkRequest) SetAppName(v string) *EnterpriseOrgCreateNodeShrinkRequest {
  s.AppName = &v
  return s
}

func (s *EnterpriseOrgCreateNodeShrinkRequest) SetBizName(v string) *EnterpriseOrgCreateNodeShrinkRequest {
  s.BizName = &v
  return s
}

func (s *EnterpriseOrgCreateNodeShrinkRequest) SetExtShrink(v string) *EnterpriseOrgCreateNodeShrinkRequest {
  s.ExtShrink = &v
  return s
}

func (s *EnterpriseOrgCreateNodeShrinkRequest) SetIsOpenApi(v bool) *EnterpriseOrgCreateNodeShrinkRequest {
  s.IsOpenApi = &v
  return s
}

func (s *EnterpriseOrgCreateNodeShrinkRequest) SetNodeId(v string) *EnterpriseOrgCreateNodeShrinkRequest {
  s.NodeId = &v
  return s
}

func (s *EnterpriseOrgCreateNodeShrinkRequest) SetNodeName(v string) *EnterpriseOrgCreateNodeShrinkRequest {
  s.NodeName = &v
  return s
}

func (s *EnterpriseOrgCreateNodeShrinkRequest) SetNodeType(v string) *EnterpriseOrgCreateNodeShrinkRequest {
  s.NodeType = &v
  return s
}

func (s *EnterpriseOrgCreateNodeShrinkRequest) SetOrientedEcId(v string) *EnterpriseOrgCreateNodeShrinkRequest {
  s.OrientedEcId = &v
  return s
}

func (s *EnterpriseOrgCreateNodeShrinkRequest) SetOrientedLeId(v string) *EnterpriseOrgCreateNodeShrinkRequest {
  s.OrientedLeId = &v
  return s
}

func (s *EnterpriseOrgCreateNodeShrinkRequest) SetOrientedNbId(v string) *EnterpriseOrgCreateNodeShrinkRequest {
  s.OrientedNbId = &v
  return s
}

func (s *EnterpriseOrgCreateNodeShrinkRequest) SetParentNodeId(v string) *EnterpriseOrgCreateNodeShrinkRequest {
  s.ParentNodeId = &v
  return s
}

func (s *EnterpriseOrgCreateNodeShrinkRequest) SetParentNodeType(v string) *EnterpriseOrgCreateNodeShrinkRequest {
  s.ParentNodeType = &v
  return s
}

func (s *EnterpriseOrgCreateNodeShrinkRequest) SetShowCompleteInfo(v bool) *EnterpriseOrgCreateNodeShrinkRequest {
  s.ShowCompleteInfo = &v
  return s
}

func (s *EnterpriseOrgCreateNodeShrinkRequest) SetTreeId(v int64) *EnterpriseOrgCreateNodeShrinkRequest {
  s.TreeId = &v
  return s
}

func (s *EnterpriseOrgCreateNodeShrinkRequest) Validate() error {
  return dara.Validate(s)
}

