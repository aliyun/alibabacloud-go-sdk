// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseOrgRenameNodeRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppName(v string) *EnterpriseOrgRenameNodeRequest
  GetAppName() *string 
  SetBizName(v string) *EnterpriseOrgRenameNodeRequest
  GetBizName() *string 
  SetExt(v map[string]interface{}) *EnterpriseOrgRenameNodeRequest
  GetExt() map[string]interface{} 
  SetIsOpenApi(v bool) *EnterpriseOrgRenameNodeRequest
  GetIsOpenApi() *bool 
  SetNodeId(v string) *EnterpriseOrgRenameNodeRequest
  GetNodeId() *string 
  SetNodeName(v string) *EnterpriseOrgRenameNodeRequest
  GetNodeName() *string 
  SetNodeType(v string) *EnterpriseOrgRenameNodeRequest
  GetNodeType() *string 
  SetOrientedEcId(v string) *EnterpriseOrgRenameNodeRequest
  GetOrientedEcId() *string 
  SetOrientedLeId(v string) *EnterpriseOrgRenameNodeRequest
  GetOrientedLeId() *string 
  SetOrientedNbId(v string) *EnterpriseOrgRenameNodeRequest
  GetOrientedNbId() *string 
  SetShowCompleteInfo(v bool) *EnterpriseOrgRenameNodeRequest
  GetShowCompleteInfo() *bool 
  SetTreeId(v int64) *EnterpriseOrgRenameNodeRequest
  GetTreeId() *int64 
}

type EnterpriseOrgRenameNodeRequest struct {
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
  ShowCompleteInfo *bool `json:"ShowCompleteInfo,omitempty" xml:"ShowCompleteInfo,omitempty"`
  TreeId *int64 `json:"TreeId,omitempty" xml:"TreeId,omitempty"`
}

func (s EnterpriseOrgRenameNodeRequest) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseOrgRenameNodeRequest) GoString() string {
  return s.String()
}

func (s *EnterpriseOrgRenameNodeRequest) GetAppName() *string  {
  return s.AppName
}

func (s *EnterpriseOrgRenameNodeRequest) GetBizName() *string  {
  return s.BizName
}

func (s *EnterpriseOrgRenameNodeRequest) GetExt() map[string]interface{}  {
  return s.Ext
}

func (s *EnterpriseOrgRenameNodeRequest) GetIsOpenApi() *bool  {
  return s.IsOpenApi
}

func (s *EnterpriseOrgRenameNodeRequest) GetNodeId() *string  {
  return s.NodeId
}

func (s *EnterpriseOrgRenameNodeRequest) GetNodeName() *string  {
  return s.NodeName
}

func (s *EnterpriseOrgRenameNodeRequest) GetNodeType() *string  {
  return s.NodeType
}

func (s *EnterpriseOrgRenameNodeRequest) GetOrientedEcId() *string  {
  return s.OrientedEcId
}

func (s *EnterpriseOrgRenameNodeRequest) GetOrientedLeId() *string  {
  return s.OrientedLeId
}

func (s *EnterpriseOrgRenameNodeRequest) GetOrientedNbId() *string  {
  return s.OrientedNbId
}

func (s *EnterpriseOrgRenameNodeRequest) GetShowCompleteInfo() *bool  {
  return s.ShowCompleteInfo
}

func (s *EnterpriseOrgRenameNodeRequest) GetTreeId() *int64  {
  return s.TreeId
}

func (s *EnterpriseOrgRenameNodeRequest) SetAppName(v string) *EnterpriseOrgRenameNodeRequest {
  s.AppName = &v
  return s
}

func (s *EnterpriseOrgRenameNodeRequest) SetBizName(v string) *EnterpriseOrgRenameNodeRequest {
  s.BizName = &v
  return s
}

func (s *EnterpriseOrgRenameNodeRequest) SetExt(v map[string]interface{}) *EnterpriseOrgRenameNodeRequest {
  s.Ext = v
  return s
}

func (s *EnterpriseOrgRenameNodeRequest) SetIsOpenApi(v bool) *EnterpriseOrgRenameNodeRequest {
  s.IsOpenApi = &v
  return s
}

func (s *EnterpriseOrgRenameNodeRequest) SetNodeId(v string) *EnterpriseOrgRenameNodeRequest {
  s.NodeId = &v
  return s
}

func (s *EnterpriseOrgRenameNodeRequest) SetNodeName(v string) *EnterpriseOrgRenameNodeRequest {
  s.NodeName = &v
  return s
}

func (s *EnterpriseOrgRenameNodeRequest) SetNodeType(v string) *EnterpriseOrgRenameNodeRequest {
  s.NodeType = &v
  return s
}

func (s *EnterpriseOrgRenameNodeRequest) SetOrientedEcId(v string) *EnterpriseOrgRenameNodeRequest {
  s.OrientedEcId = &v
  return s
}

func (s *EnterpriseOrgRenameNodeRequest) SetOrientedLeId(v string) *EnterpriseOrgRenameNodeRequest {
  s.OrientedLeId = &v
  return s
}

func (s *EnterpriseOrgRenameNodeRequest) SetOrientedNbId(v string) *EnterpriseOrgRenameNodeRequest {
  s.OrientedNbId = &v
  return s
}

func (s *EnterpriseOrgRenameNodeRequest) SetShowCompleteInfo(v bool) *EnterpriseOrgRenameNodeRequest {
  s.ShowCompleteInfo = &v
  return s
}

func (s *EnterpriseOrgRenameNodeRequest) SetTreeId(v int64) *EnterpriseOrgRenameNodeRequest {
  s.TreeId = &v
  return s
}

func (s *EnterpriseOrgRenameNodeRequest) Validate() error {
  return dara.Validate(s)
}

