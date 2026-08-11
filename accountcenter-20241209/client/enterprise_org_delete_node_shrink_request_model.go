// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseOrgDeleteNodeShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppName(v string) *EnterpriseOrgDeleteNodeShrinkRequest
  GetAppName() *string 
  SetBizName(v string) *EnterpriseOrgDeleteNodeShrinkRequest
  GetBizName() *string 
  SetExtShrink(v string) *EnterpriseOrgDeleteNodeShrinkRequest
  GetExtShrink() *string 
  SetIsOpenApi(v bool) *EnterpriseOrgDeleteNodeShrinkRequest
  GetIsOpenApi() *bool 
  SetNodeId(v string) *EnterpriseOrgDeleteNodeShrinkRequest
  GetNodeId() *string 
  SetNodeType(v string) *EnterpriseOrgDeleteNodeShrinkRequest
  GetNodeType() *string 
  SetOrientedEcId(v string) *EnterpriseOrgDeleteNodeShrinkRequest
  GetOrientedEcId() *string 
  SetOrientedLeId(v string) *EnterpriseOrgDeleteNodeShrinkRequest
  GetOrientedLeId() *string 
  SetOrientedNbId(v string) *EnterpriseOrgDeleteNodeShrinkRequest
  GetOrientedNbId() *string 
  SetShowCompleteInfo(v bool) *EnterpriseOrgDeleteNodeShrinkRequest
  GetShowCompleteInfo() *bool 
  SetTreeId(v int64) *EnterpriseOrgDeleteNodeShrinkRequest
  GetTreeId() *int64 
}

type EnterpriseOrgDeleteNodeShrinkRequest struct {
  AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
  BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
  ExtShrink *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
  IsOpenApi *bool `json:"IsOpenApi,omitempty" xml:"IsOpenApi,omitempty"`
  NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
  NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
  OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
  OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
  OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
  ShowCompleteInfo *bool `json:"ShowCompleteInfo,omitempty" xml:"ShowCompleteInfo,omitempty"`
  TreeId *int64 `json:"TreeId,omitempty" xml:"TreeId,omitempty"`
}

func (s EnterpriseOrgDeleteNodeShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseOrgDeleteNodeShrinkRequest) GoString() string {
  return s.String()
}

func (s *EnterpriseOrgDeleteNodeShrinkRequest) GetAppName() *string  {
  return s.AppName
}

func (s *EnterpriseOrgDeleteNodeShrinkRequest) GetBizName() *string  {
  return s.BizName
}

func (s *EnterpriseOrgDeleteNodeShrinkRequest) GetExtShrink() *string  {
  return s.ExtShrink
}

func (s *EnterpriseOrgDeleteNodeShrinkRequest) GetIsOpenApi() *bool  {
  return s.IsOpenApi
}

func (s *EnterpriseOrgDeleteNodeShrinkRequest) GetNodeId() *string  {
  return s.NodeId
}

func (s *EnterpriseOrgDeleteNodeShrinkRequest) GetNodeType() *string  {
  return s.NodeType
}

func (s *EnterpriseOrgDeleteNodeShrinkRequest) GetOrientedEcId() *string  {
  return s.OrientedEcId
}

func (s *EnterpriseOrgDeleteNodeShrinkRequest) GetOrientedLeId() *string  {
  return s.OrientedLeId
}

func (s *EnterpriseOrgDeleteNodeShrinkRequest) GetOrientedNbId() *string  {
  return s.OrientedNbId
}

func (s *EnterpriseOrgDeleteNodeShrinkRequest) GetShowCompleteInfo() *bool  {
  return s.ShowCompleteInfo
}

func (s *EnterpriseOrgDeleteNodeShrinkRequest) GetTreeId() *int64  {
  return s.TreeId
}

func (s *EnterpriseOrgDeleteNodeShrinkRequest) SetAppName(v string) *EnterpriseOrgDeleteNodeShrinkRequest {
  s.AppName = &v
  return s
}

func (s *EnterpriseOrgDeleteNodeShrinkRequest) SetBizName(v string) *EnterpriseOrgDeleteNodeShrinkRequest {
  s.BizName = &v
  return s
}

func (s *EnterpriseOrgDeleteNodeShrinkRequest) SetExtShrink(v string) *EnterpriseOrgDeleteNodeShrinkRequest {
  s.ExtShrink = &v
  return s
}

func (s *EnterpriseOrgDeleteNodeShrinkRequest) SetIsOpenApi(v bool) *EnterpriseOrgDeleteNodeShrinkRequest {
  s.IsOpenApi = &v
  return s
}

func (s *EnterpriseOrgDeleteNodeShrinkRequest) SetNodeId(v string) *EnterpriseOrgDeleteNodeShrinkRequest {
  s.NodeId = &v
  return s
}

func (s *EnterpriseOrgDeleteNodeShrinkRequest) SetNodeType(v string) *EnterpriseOrgDeleteNodeShrinkRequest {
  s.NodeType = &v
  return s
}

func (s *EnterpriseOrgDeleteNodeShrinkRequest) SetOrientedEcId(v string) *EnterpriseOrgDeleteNodeShrinkRequest {
  s.OrientedEcId = &v
  return s
}

func (s *EnterpriseOrgDeleteNodeShrinkRequest) SetOrientedLeId(v string) *EnterpriseOrgDeleteNodeShrinkRequest {
  s.OrientedLeId = &v
  return s
}

func (s *EnterpriseOrgDeleteNodeShrinkRequest) SetOrientedNbId(v string) *EnterpriseOrgDeleteNodeShrinkRequest {
  s.OrientedNbId = &v
  return s
}

func (s *EnterpriseOrgDeleteNodeShrinkRequest) SetShowCompleteInfo(v bool) *EnterpriseOrgDeleteNodeShrinkRequest {
  s.ShowCompleteInfo = &v
  return s
}

func (s *EnterpriseOrgDeleteNodeShrinkRequest) SetTreeId(v int64) *EnterpriseOrgDeleteNodeShrinkRequest {
  s.TreeId = &v
  return s
}

func (s *EnterpriseOrgDeleteNodeShrinkRequest) Validate() error {
  return dara.Validate(s)
}

