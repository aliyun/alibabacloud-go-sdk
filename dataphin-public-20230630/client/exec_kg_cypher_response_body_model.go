// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecKgCypherResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ExecKgCypherResponseBody
  GetCode() *string 
  SetData(v *ExecKgCypherResponseBodyData) *ExecKgCypherResponseBody
  GetData() *ExecKgCypherResponseBodyData 
  SetHttpStatusCode(v int32) *ExecKgCypherResponseBody
  GetHttpStatusCode() *int32 
  SetMessage(v string) *ExecKgCypherResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExecKgCypherResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecKgCypherResponseBody
  GetSuccess() *bool 
}

type ExecKgCypherResponseBody struct {
  // The backend response code.
  // 
  // example:
  // 
  // OK
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // The query result.
  Data *ExecKgCypherResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  // The HTTP status code.
  // 
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  // The backend exception details.
  // 
  // example:
  // 
  // internal error
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // Id of the request
  // 
  // example:
  // 
  // 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the request was successful.
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExecKgCypherResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecKgCypherResponseBody) GoString() string {
  return s.String()
}

func (s *ExecKgCypherResponseBody) GetCode() *string  {
  return s.Code
}

func (s *ExecKgCypherResponseBody) GetData() *ExecKgCypherResponseBodyData  {
  return s.Data
}

func (s *ExecKgCypherResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *ExecKgCypherResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExecKgCypherResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecKgCypherResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecKgCypherResponseBody) SetCode(v string) *ExecKgCypherResponseBody {
  s.Code = &v
  return s
}

func (s *ExecKgCypherResponseBody) SetData(v *ExecKgCypherResponseBodyData) *ExecKgCypherResponseBody {
  s.Data = v
  return s
}

func (s *ExecKgCypherResponseBody) SetHttpStatusCode(v int32) *ExecKgCypherResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *ExecKgCypherResponseBody) SetMessage(v string) *ExecKgCypherResponseBody {
  s.Message = &v
  return s
}

func (s *ExecKgCypherResponseBody) SetRequestId(v string) *ExecKgCypherResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecKgCypherResponseBody) SetSuccess(v bool) *ExecKgCypherResponseBody {
  s.Success = &v
  return s
}

func (s *ExecKgCypherResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecKgCypherResponseBodyData struct {
  // The list of edges.
  EdgeList []*ExecKgCypherResponseBodyDataEdgeList `json:"EdgeList,omitempty" xml:"EdgeList,omitempty" type:"Repeated"`
  // The transformed execution statement.
  // 
  // example:
  // 
  // MATCH p=()-[:Product]->() RETURN p, count(*) LIMIT 25
  ExecuteCypher *string `json:"ExecuteCypher,omitempty" xml:"ExecuteCypher,omitempty"`
  // The list of nodes.
  NodeList []*ExecKgCypherResponseBodyDataNodeList `json:"NodeList,omitempty" xml:"NodeList,omitempty" type:"Repeated"`
  // The list of rows.
  RowList []*ExecKgCypherResponseBodyDataRowList `json:"RowList,omitempty" xml:"RowList,omitempty" type:"Repeated"`
}

func (s ExecKgCypherResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s ExecKgCypherResponseBodyData) GoString() string {
  return s.String()
}

func (s *ExecKgCypherResponseBodyData) GetEdgeList() []*ExecKgCypherResponseBodyDataEdgeList  {
  return s.EdgeList
}

func (s *ExecKgCypherResponseBodyData) GetExecuteCypher() *string  {
  return s.ExecuteCypher
}

func (s *ExecKgCypherResponseBodyData) GetNodeList() []*ExecKgCypherResponseBodyDataNodeList  {
  return s.NodeList
}

func (s *ExecKgCypherResponseBodyData) GetRowList() []*ExecKgCypherResponseBodyDataRowList  {
  return s.RowList
}

func (s *ExecKgCypherResponseBodyData) SetEdgeList(v []*ExecKgCypherResponseBodyDataEdgeList) *ExecKgCypherResponseBodyData {
  s.EdgeList = v
  return s
}

func (s *ExecKgCypherResponseBodyData) SetExecuteCypher(v string) *ExecKgCypherResponseBodyData {
  s.ExecuteCypher = &v
  return s
}

func (s *ExecKgCypherResponseBodyData) SetNodeList(v []*ExecKgCypherResponseBodyDataNodeList) *ExecKgCypherResponseBodyData {
  s.NodeList = v
  return s
}

func (s *ExecKgCypherResponseBodyData) SetRowList(v []*ExecKgCypherResponseBodyDataRowList) *ExecKgCypherResponseBodyData {
  s.RowList = v
  return s
}

func (s *ExecKgCypherResponseBodyData) Validate() error {
  if s.EdgeList != nil {
    for _, item := range s.EdgeList {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  if s.NodeList != nil {
    for _, item := range s.NodeList {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  if s.RowList != nil {
    for _, item := range s.RowList {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type ExecKgCypherResponseBodyDataEdgeList struct {
  // The data ID of the relationship record.
  // 
  // example:
  // 
  // bcd-456
  DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
  // The list of relationship record properties.
  Properties []*ExecKgCypherResponseBodyDataEdgeListProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Repeated"`
  // The relationship type.
  // 
  // example:
  // 
  // SALE
  RelationType *string `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
  // The data ID of the source entity record.
  // 
  // example:
  // 
  // source-123
  SourceEntityDataId *string `json:"SourceEntityDataId,omitempty" xml:"SourceEntityDataId,omitempty"`
  // The source entity type.
  // 
  // example:
  // 
  // Product
  SourceEntityType *string `json:"SourceEntityType,omitempty" xml:"SourceEntityType,omitempty"`
  // The data ID of the target entity record.
  // 
  // example:
  // 
  // target-345
  TargetEntityDataId *string `json:"TargetEntityDataId,omitempty" xml:"TargetEntityDataId,omitempty"`
  // The target entity type.
  // 
  // example:
  // 
  // Shop
  TargetEntityType *string `json:"TargetEntityType,omitempty" xml:"TargetEntityType,omitempty"`
}

func (s ExecKgCypherResponseBodyDataEdgeList) String() string {
  return dara.Prettify(s)
}

func (s ExecKgCypherResponseBodyDataEdgeList) GoString() string {
  return s.String()
}

func (s *ExecKgCypherResponseBodyDataEdgeList) GetDataId() *string  {
  return s.DataId
}

func (s *ExecKgCypherResponseBodyDataEdgeList) GetProperties() []*ExecKgCypherResponseBodyDataEdgeListProperties  {
  return s.Properties
}

func (s *ExecKgCypherResponseBodyDataEdgeList) GetRelationType() *string  {
  return s.RelationType
}

func (s *ExecKgCypherResponseBodyDataEdgeList) GetSourceEntityDataId() *string  {
  return s.SourceEntityDataId
}

func (s *ExecKgCypherResponseBodyDataEdgeList) GetSourceEntityType() *string  {
  return s.SourceEntityType
}

func (s *ExecKgCypherResponseBodyDataEdgeList) GetTargetEntityDataId() *string  {
  return s.TargetEntityDataId
}

func (s *ExecKgCypherResponseBodyDataEdgeList) GetTargetEntityType() *string  {
  return s.TargetEntityType
}

func (s *ExecKgCypherResponseBodyDataEdgeList) SetDataId(v string) *ExecKgCypherResponseBodyDataEdgeList {
  s.DataId = &v
  return s
}

func (s *ExecKgCypherResponseBodyDataEdgeList) SetProperties(v []*ExecKgCypherResponseBodyDataEdgeListProperties) *ExecKgCypherResponseBodyDataEdgeList {
  s.Properties = v
  return s
}

func (s *ExecKgCypherResponseBodyDataEdgeList) SetRelationType(v string) *ExecKgCypherResponseBodyDataEdgeList {
  s.RelationType = &v
  return s
}

func (s *ExecKgCypherResponseBodyDataEdgeList) SetSourceEntityDataId(v string) *ExecKgCypherResponseBodyDataEdgeList {
  s.SourceEntityDataId = &v
  return s
}

func (s *ExecKgCypherResponseBodyDataEdgeList) SetSourceEntityType(v string) *ExecKgCypherResponseBodyDataEdgeList {
  s.SourceEntityType = &v
  return s
}

func (s *ExecKgCypherResponseBodyDataEdgeList) SetTargetEntityDataId(v string) *ExecKgCypherResponseBodyDataEdgeList {
  s.TargetEntityDataId = &v
  return s
}

func (s *ExecKgCypherResponseBodyDataEdgeList) SetTargetEntityType(v string) *ExecKgCypherResponseBodyDataEdgeList {
  s.TargetEntityType = &v
  return s
}

func (s *ExecKgCypherResponseBodyDataEdgeList) Validate() error {
  if s.Properties != nil {
    for _, item := range s.Properties {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type ExecKgCypherResponseBodyDataEdgeListProperties struct {
  // The property code.
  // 
  // example:
  // 
  // company_name
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // The property value.
  // 
  // example:
  // 
  // Alibaba
  Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ExecKgCypherResponseBodyDataEdgeListProperties) String() string {
  return dara.Prettify(s)
}

func (s ExecKgCypherResponseBodyDataEdgeListProperties) GoString() string {
  return s.String()
}

func (s *ExecKgCypherResponseBodyDataEdgeListProperties) GetCode() *string  {
  return s.Code
}

func (s *ExecKgCypherResponseBodyDataEdgeListProperties) GetValue() *string  {
  return s.Value
}

func (s *ExecKgCypherResponseBodyDataEdgeListProperties) SetCode(v string) *ExecKgCypherResponseBodyDataEdgeListProperties {
  s.Code = &v
  return s
}

func (s *ExecKgCypherResponseBodyDataEdgeListProperties) SetValue(v string) *ExecKgCypherResponseBodyDataEdgeListProperties {
  s.Value = &v
  return s
}

func (s *ExecKgCypherResponseBodyDataEdgeListProperties) Validate() error {
  return dara.Validate(s)
}

type ExecKgCypherResponseBodyDataNodeList struct {
  // The data ID of the entity record.
  // 
  // example:
  // 
  // abc-123
  DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
  // The entity type.
  // 
  // example:
  // 
  // Product
  EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
  // The list of entity record properties.
  Properties []*ExecKgCypherResponseBodyDataNodeListProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Repeated"`
}

func (s ExecKgCypherResponseBodyDataNodeList) String() string {
  return dara.Prettify(s)
}

func (s ExecKgCypherResponseBodyDataNodeList) GoString() string {
  return s.String()
}

func (s *ExecKgCypherResponseBodyDataNodeList) GetDataId() *string  {
  return s.DataId
}

func (s *ExecKgCypherResponseBodyDataNodeList) GetEntityType() *string  {
  return s.EntityType
}

func (s *ExecKgCypherResponseBodyDataNodeList) GetProperties() []*ExecKgCypherResponseBodyDataNodeListProperties  {
  return s.Properties
}

func (s *ExecKgCypherResponseBodyDataNodeList) SetDataId(v string) *ExecKgCypherResponseBodyDataNodeList {
  s.DataId = &v
  return s
}

func (s *ExecKgCypherResponseBodyDataNodeList) SetEntityType(v string) *ExecKgCypherResponseBodyDataNodeList {
  s.EntityType = &v
  return s
}

func (s *ExecKgCypherResponseBodyDataNodeList) SetProperties(v []*ExecKgCypherResponseBodyDataNodeListProperties) *ExecKgCypherResponseBodyDataNodeList {
  s.Properties = v
  return s
}

func (s *ExecKgCypherResponseBodyDataNodeList) Validate() error {
  if s.Properties != nil {
    for _, item := range s.Properties {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type ExecKgCypherResponseBodyDataNodeListProperties struct {
  // The property code.
  // 
  // example:
  // 
  // company_name
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // The property value.
  // 
  // example:
  // 
  // Alibaba
  Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ExecKgCypherResponseBodyDataNodeListProperties) String() string {
  return dara.Prettify(s)
}

func (s ExecKgCypherResponseBodyDataNodeListProperties) GoString() string {
  return s.String()
}

func (s *ExecKgCypherResponseBodyDataNodeListProperties) GetCode() *string  {
  return s.Code
}

func (s *ExecKgCypherResponseBodyDataNodeListProperties) GetValue() *string  {
  return s.Value
}

func (s *ExecKgCypherResponseBodyDataNodeListProperties) SetCode(v string) *ExecKgCypherResponseBodyDataNodeListProperties {
  s.Code = &v
  return s
}

func (s *ExecKgCypherResponseBodyDataNodeListProperties) SetValue(v string) *ExecKgCypherResponseBodyDataNodeListProperties {
  s.Value = &v
  return s
}

func (s *ExecKgCypherResponseBodyDataNodeListProperties) Validate() error {
  return dara.Validate(s)
}

type ExecKgCypherResponseBodyDataRowList struct {
  // The list of columns in the row.
  Columns []*ExecKgCypherResponseBodyDataRowListColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
}

func (s ExecKgCypherResponseBodyDataRowList) String() string {
  return dara.Prettify(s)
}

func (s ExecKgCypherResponseBodyDataRowList) GoString() string {
  return s.String()
}

func (s *ExecKgCypherResponseBodyDataRowList) GetColumns() []*ExecKgCypherResponseBodyDataRowListColumns  {
  return s.Columns
}

func (s *ExecKgCypherResponseBodyDataRowList) SetColumns(v []*ExecKgCypherResponseBodyDataRowListColumns) *ExecKgCypherResponseBodyDataRowList {
  s.Columns = v
  return s
}

func (s *ExecKgCypherResponseBodyDataRowList) Validate() error {
  if s.Columns != nil {
    for _, item := range s.Columns {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type ExecKgCypherResponseBodyDataRowListColumns struct {
  // The property code.
  // 
  // example:
  // 
  // company_name
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // The property value.
  // 
  // example:
  // 
  // Alibaba
  Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ExecKgCypherResponseBodyDataRowListColumns) String() string {
  return dara.Prettify(s)
}

func (s ExecKgCypherResponseBodyDataRowListColumns) GoString() string {
  return s.String()
}

func (s *ExecKgCypherResponseBodyDataRowListColumns) GetCode() *string  {
  return s.Code
}

func (s *ExecKgCypherResponseBodyDataRowListColumns) GetValue() *string  {
  return s.Value
}

func (s *ExecKgCypherResponseBodyDataRowListColumns) SetCode(v string) *ExecKgCypherResponseBodyDataRowListColumns {
  s.Code = &v
  return s
}

func (s *ExecKgCypherResponseBodyDataRowListColumns) SetValue(v string) *ExecKgCypherResponseBodyDataRowListColumns {
  s.Value = &v
  return s
}

func (s *ExecKgCypherResponseBodyDataRowListColumns) Validate() error {
  return dara.Validate(s)
}

