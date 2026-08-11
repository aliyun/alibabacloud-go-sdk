// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseOrgCreateNodeRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppName(v string) *EnterpriseOrgCreateNodeRequest
  GetAppName() *string 
  SetBizName(v string) *EnterpriseOrgCreateNodeRequest
  GetBizName() *string 
  SetExt(v map[string]interface{}) *EnterpriseOrgCreateNodeRequest
  GetExt() map[string]interface{} 
  SetIsOpenApi(v bool) *EnterpriseOrgCreateNodeRequest
  GetIsOpenApi() *bool 
  SetNodeId(v string) *EnterpriseOrgCreateNodeRequest
  GetNodeId() *string 
  SetNodeName(v string) *EnterpriseOrgCreateNodeRequest
  GetNodeName() *string 
  SetNodeType(v string) *EnterpriseOrgCreateNodeRequest
  GetNodeType() *string 
  SetOrientedEcId(v string) *EnterpriseOrgCreateNodeRequest
  GetOrientedEcId() *string 
  SetOrientedLeId(v string) *EnterpriseOrgCreateNodeRequest
  GetOrientedLeId() *string 
  SetOrientedNbId(v string) *EnterpriseOrgCreateNodeRequest
  GetOrientedNbId() *string 
  SetParentNodeId(v string) *EnterpriseOrgCreateNodeRequest
  GetParentNodeId() *string 
  SetParentNodeType(v string) *EnterpriseOrgCreateNodeRequest
  GetParentNodeType() *string 
  SetShowCompleteInfo(v bool) *EnterpriseOrgCreateNodeRequest
  GetShowCompleteInfo() *bool 
  SetTreeId(v int64) *EnterpriseOrgCreateNodeRequest
  GetTreeId() *int64 
}

type EnterpriseOrgCreateNodeRequest struct {
  AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
  BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
  Ext map[string]interface{} `json:"Ext,omitempty" xml:"Ext,omitempty"`
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

func (s EnterpriseOrgCreateNodeRequest) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseOrgCreateNodeRequest) GoString() string {
  return s.String()
}

func (s *EnterpriseOrgCreateNodeRequest) GetAppName() *string  {
  return s.AppName
}

func (s *EnterpriseOrgCreateNodeRequest) GetBizName() *string  {
  return s.BizName
}

func (s *EnterpriseOrgCreateNodeRequest) GetExt() map[string]interface{}  {
  return s.Ext
}

func (s *EnterpriseOrgCreateNodeRequest) GetIsOpenApi() *bool  {
  return s.IsOpenApi
}

func (s *EnterpriseOrgCreateNodeRequest) GetNodeId() *string  {
  return s.NodeId
}

func (s *EnterpriseOrgCreateNodeRequest) GetNodeName() *string  {
  return s.NodeName
}

func (s *EnterpriseOrgCreateNodeRequest) GetNodeType() *string  {
  return s.NodeType
}

func (s *EnterpriseOrgCreateNodeRequest) GetOrientedEcId() *string  {
  return s.OrientedEcId
}

func (s *EnterpriseOrgCreateNodeRequest) GetOrientedLeId() *string  {
  return s.OrientedLeId
}

func (s *EnterpriseOrgCreateNodeRequest) GetOrientedNbId() *string  {
  return s.OrientedNbId
}

func (s *EnterpriseOrgCreateNodeRequest) GetParentNodeId() *string  {
  return s.ParentNodeId
}

func (s *EnterpriseOrgCreateNodeRequest) GetParentNodeType() *string  {
  return s.ParentNodeType
}

func (s *EnterpriseOrgCreateNodeRequest) GetShowCompleteInfo() *bool  {
  return s.ShowCompleteInfo
}

func (s *EnterpriseOrgCreateNodeRequest) GetTreeId() *int64  {
  return s.TreeId
}

func (s *EnterpriseOrgCreateNodeRequest) SetAppName(v string) *EnterpriseOrgCreateNodeRequest {
  s.AppName = &v
  return s
}

func (s *EnterpriseOrgCreateNodeRequest) SetBizName(v string) *EnterpriseOrgCreateNodeRequest {
  s.BizName = &v
  return s
}

func (s *EnterpriseOrgCreateNodeRequest) SetExt(v map[string]interface{}) *EnterpriseOrgCreateNodeRequest {
  s.Ext = v
  return s
}

func (s *EnterpriseOrgCreateNodeRequest) SetIsOpenApi(v bool) *EnterpriseOrgCreateNodeRequest {
  s.IsOpenApi = &v
  return s
}

func (s *EnterpriseOrgCreateNodeRequest) SetNodeId(v string) *EnterpriseOrgCreateNodeRequest {
  s.NodeId = &v
  return s
}

func (s *EnterpriseOrgCreateNodeRequest) SetNodeName(v string) *EnterpriseOrgCreateNodeRequest {
  s.NodeName = &v
  return s
}

func (s *EnterpriseOrgCreateNodeRequest) SetNodeType(v string) *EnterpriseOrgCreateNodeRequest {
  s.NodeType = &v
  return s
}

func (s *EnterpriseOrgCreateNodeRequest) SetOrientedEcId(v string) *EnterpriseOrgCreateNodeRequest {
  s.OrientedEcId = &v
  return s
}

func (s *EnterpriseOrgCreateNodeRequest) SetOrientedLeId(v string) *EnterpriseOrgCreateNodeRequest {
  s.OrientedLeId = &v
  return s
}

func (s *EnterpriseOrgCreateNodeRequest) SetOrientedNbId(v string) *EnterpriseOrgCreateNodeRequest {
  s.OrientedNbId = &v
  return s
}

func (s *EnterpriseOrgCreateNodeRequest) SetParentNodeId(v string) *EnterpriseOrgCreateNodeRequest {
  s.ParentNodeId = &v
  return s
}

func (s *EnterpriseOrgCreateNodeRequest) SetParentNodeType(v string) *EnterpriseOrgCreateNodeRequest {
  s.ParentNodeType = &v
  return s
}

func (s *EnterpriseOrgCreateNodeRequest) SetShowCompleteInfo(v bool) *EnterpriseOrgCreateNodeRequest {
  s.ShowCompleteInfo = &v
  return s
}

func (s *EnterpriseOrgCreateNodeRequest) SetTreeId(v int64) *EnterpriseOrgCreateNodeRequest {
  s.TreeId = &v
  return s
}

func (s *EnterpriseOrgCreateNodeRequest) Validate() error {
  return dara.Validate(s)
}

