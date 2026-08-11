// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseOrgDeleteNodeRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppName(v string) *EnterpriseOrgDeleteNodeRequest
  GetAppName() *string 
  SetBizName(v string) *EnterpriseOrgDeleteNodeRequest
  GetBizName() *string 
  SetExt(v map[string]interface{}) *EnterpriseOrgDeleteNodeRequest
  GetExt() map[string]interface{} 
  SetIsOpenApi(v bool) *EnterpriseOrgDeleteNodeRequest
  GetIsOpenApi() *bool 
  SetNodeId(v string) *EnterpriseOrgDeleteNodeRequest
  GetNodeId() *string 
  SetNodeType(v string) *EnterpriseOrgDeleteNodeRequest
  GetNodeType() *string 
  SetOrientedEcId(v string) *EnterpriseOrgDeleteNodeRequest
  GetOrientedEcId() *string 
  SetOrientedLeId(v string) *EnterpriseOrgDeleteNodeRequest
  GetOrientedLeId() *string 
  SetOrientedNbId(v string) *EnterpriseOrgDeleteNodeRequest
  GetOrientedNbId() *string 
  SetShowCompleteInfo(v bool) *EnterpriseOrgDeleteNodeRequest
  GetShowCompleteInfo() *bool 
  SetTreeId(v int64) *EnterpriseOrgDeleteNodeRequest
  GetTreeId() *int64 
}

type EnterpriseOrgDeleteNodeRequest struct {
  AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
  BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
  Ext map[string]interface{} `json:"Ext,omitempty" xml:"Ext,omitempty"`
  IsOpenApi *bool `json:"IsOpenApi,omitempty" xml:"IsOpenApi,omitempty"`
  NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
  NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
  OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
  OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
  OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
  ShowCompleteInfo *bool `json:"ShowCompleteInfo,omitempty" xml:"ShowCompleteInfo,omitempty"`
  TreeId *int64 `json:"TreeId,omitempty" xml:"TreeId,omitempty"`
}

func (s EnterpriseOrgDeleteNodeRequest) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseOrgDeleteNodeRequest) GoString() string {
  return s.String()
}

func (s *EnterpriseOrgDeleteNodeRequest) GetAppName() *string  {
  return s.AppName
}

func (s *EnterpriseOrgDeleteNodeRequest) GetBizName() *string  {
  return s.BizName
}

func (s *EnterpriseOrgDeleteNodeRequest) GetExt() map[string]interface{}  {
  return s.Ext
}

func (s *EnterpriseOrgDeleteNodeRequest) GetIsOpenApi() *bool  {
  return s.IsOpenApi
}

func (s *EnterpriseOrgDeleteNodeRequest) GetNodeId() *string  {
  return s.NodeId
}

func (s *EnterpriseOrgDeleteNodeRequest) GetNodeType() *string  {
  return s.NodeType
}

func (s *EnterpriseOrgDeleteNodeRequest) GetOrientedEcId() *string  {
  return s.OrientedEcId
}

func (s *EnterpriseOrgDeleteNodeRequest) GetOrientedLeId() *string  {
  return s.OrientedLeId
}

func (s *EnterpriseOrgDeleteNodeRequest) GetOrientedNbId() *string  {
  return s.OrientedNbId
}

func (s *EnterpriseOrgDeleteNodeRequest) GetShowCompleteInfo() *bool  {
  return s.ShowCompleteInfo
}

func (s *EnterpriseOrgDeleteNodeRequest) GetTreeId() *int64  {
  return s.TreeId
}

func (s *EnterpriseOrgDeleteNodeRequest) SetAppName(v string) *EnterpriseOrgDeleteNodeRequest {
  s.AppName = &v
  return s
}

func (s *EnterpriseOrgDeleteNodeRequest) SetBizName(v string) *EnterpriseOrgDeleteNodeRequest {
  s.BizName = &v
  return s
}

func (s *EnterpriseOrgDeleteNodeRequest) SetExt(v map[string]interface{}) *EnterpriseOrgDeleteNodeRequest {
  s.Ext = v
  return s
}

func (s *EnterpriseOrgDeleteNodeRequest) SetIsOpenApi(v bool) *EnterpriseOrgDeleteNodeRequest {
  s.IsOpenApi = &v
  return s
}

func (s *EnterpriseOrgDeleteNodeRequest) SetNodeId(v string) *EnterpriseOrgDeleteNodeRequest {
  s.NodeId = &v
  return s
}

func (s *EnterpriseOrgDeleteNodeRequest) SetNodeType(v string) *EnterpriseOrgDeleteNodeRequest {
  s.NodeType = &v
  return s
}

func (s *EnterpriseOrgDeleteNodeRequest) SetOrientedEcId(v string) *EnterpriseOrgDeleteNodeRequest {
  s.OrientedEcId = &v
  return s
}

func (s *EnterpriseOrgDeleteNodeRequest) SetOrientedLeId(v string) *EnterpriseOrgDeleteNodeRequest {
  s.OrientedLeId = &v
  return s
}

func (s *EnterpriseOrgDeleteNodeRequest) SetOrientedNbId(v string) *EnterpriseOrgDeleteNodeRequest {
  s.OrientedNbId = &v
  return s
}

func (s *EnterpriseOrgDeleteNodeRequest) SetShowCompleteInfo(v bool) *EnterpriseOrgDeleteNodeRequest {
  s.ShowCompleteInfo = &v
  return s
}

func (s *EnterpriseOrgDeleteNodeRequest) SetTreeId(v int64) *EnterpriseOrgDeleteNodeRequest {
  s.TreeId = &v
  return s
}

func (s *EnterpriseOrgDeleteNodeRequest) Validate() error {
  return dara.Validate(s)
}

