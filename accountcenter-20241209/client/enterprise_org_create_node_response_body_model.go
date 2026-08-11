// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseOrgCreateNodeResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnterpriseOrgCreateNodeResponseBody
  GetCode() *string 
  SetData(v *EnterpriseOrgCreateNodeResponseBodyData) *EnterpriseOrgCreateNodeResponseBody
  GetData() *EnterpriseOrgCreateNodeResponseBodyData 
  SetMessage(v string) *EnterpriseOrgCreateNodeResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnterpriseOrgCreateNodeResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnterpriseOrgCreateNodeResponseBody
  GetSuccess() *bool 
}

type EnterpriseOrgCreateNodeResponseBody struct {
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Data *EnterpriseOrgCreateNodeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseOrgCreateNodeResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseOrgCreateNodeResponseBody) GoString() string {
  return s.String()
}

func (s *EnterpriseOrgCreateNodeResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnterpriseOrgCreateNodeResponseBody) GetData() *EnterpriseOrgCreateNodeResponseBodyData  {
  return s.Data
}

func (s *EnterpriseOrgCreateNodeResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnterpriseOrgCreateNodeResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseOrgCreateNodeResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnterpriseOrgCreateNodeResponseBody) SetCode(v string) *EnterpriseOrgCreateNodeResponseBody {
  s.Code = &v
  return s
}

func (s *EnterpriseOrgCreateNodeResponseBody) SetData(v *EnterpriseOrgCreateNodeResponseBodyData) *EnterpriseOrgCreateNodeResponseBody {
  s.Data = v
  return s
}

func (s *EnterpriseOrgCreateNodeResponseBody) SetMessage(v string) *EnterpriseOrgCreateNodeResponseBody {
  s.Message = &v
  return s
}

func (s *EnterpriseOrgCreateNodeResponseBody) SetRequestId(v string) *EnterpriseOrgCreateNodeResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnterpriseOrgCreateNodeResponseBody) SetSuccess(v bool) *EnterpriseOrgCreateNodeResponseBody {
  s.Success = &v
  return s
}

func (s *EnterpriseOrgCreateNodeResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EnterpriseOrgCreateNodeResponseBodyData struct {
  GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
  Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
  NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
  NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
  NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
  ParentNodeId *string `json:"ParentNodeId,omitempty" xml:"ParentNodeId,omitempty"`
  ParentNodeType *string `json:"ParentNodeType,omitempty" xml:"ParentNodeType,omitempty"`
  TreeId *int64 `json:"TreeId,omitempty" xml:"TreeId,omitempty"`
}

func (s EnterpriseOrgCreateNodeResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseOrgCreateNodeResponseBodyData) GoString() string {
  return s.String()
}

func (s *EnterpriseOrgCreateNodeResponseBodyData) GetGmtCreate() *string  {
  return s.GmtCreate
}

func (s *EnterpriseOrgCreateNodeResponseBodyData) GetId() *int64  {
  return s.Id
}

func (s *EnterpriseOrgCreateNodeResponseBodyData) GetNodeId() *string  {
  return s.NodeId
}

func (s *EnterpriseOrgCreateNodeResponseBodyData) GetNodeName() *string  {
  return s.NodeName
}

func (s *EnterpriseOrgCreateNodeResponseBodyData) GetNodeType() *string  {
  return s.NodeType
}

func (s *EnterpriseOrgCreateNodeResponseBodyData) GetParentNodeId() *string  {
  return s.ParentNodeId
}

func (s *EnterpriseOrgCreateNodeResponseBodyData) GetParentNodeType() *string  {
  return s.ParentNodeType
}

func (s *EnterpriseOrgCreateNodeResponseBodyData) GetTreeId() *int64  {
  return s.TreeId
}

func (s *EnterpriseOrgCreateNodeResponseBodyData) SetGmtCreate(v string) *EnterpriseOrgCreateNodeResponseBodyData {
  s.GmtCreate = &v
  return s
}

func (s *EnterpriseOrgCreateNodeResponseBodyData) SetId(v int64) *EnterpriseOrgCreateNodeResponseBodyData {
  s.Id = &v
  return s
}

func (s *EnterpriseOrgCreateNodeResponseBodyData) SetNodeId(v string) *EnterpriseOrgCreateNodeResponseBodyData {
  s.NodeId = &v
  return s
}

func (s *EnterpriseOrgCreateNodeResponseBodyData) SetNodeName(v string) *EnterpriseOrgCreateNodeResponseBodyData {
  s.NodeName = &v
  return s
}

func (s *EnterpriseOrgCreateNodeResponseBodyData) SetNodeType(v string) *EnterpriseOrgCreateNodeResponseBodyData {
  s.NodeType = &v
  return s
}

func (s *EnterpriseOrgCreateNodeResponseBodyData) SetParentNodeId(v string) *EnterpriseOrgCreateNodeResponseBodyData {
  s.ParentNodeId = &v
  return s
}

func (s *EnterpriseOrgCreateNodeResponseBodyData) SetParentNodeType(v string) *EnterpriseOrgCreateNodeResponseBodyData {
  s.ParentNodeType = &v
  return s
}

func (s *EnterpriseOrgCreateNodeResponseBodyData) SetTreeId(v int64) *EnterpriseOrgCreateNodeResponseBodyData {
  s.TreeId = &v
  return s
}

func (s *EnterpriseOrgCreateNodeResponseBodyData) Validate() error {
  return dara.Validate(s)
}

