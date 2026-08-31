// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecKgGremlinResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ExecKgGremlinResponseBody
  GetCode() *string 
  SetData(v *ExecKgGremlinResponseBodyData) *ExecKgGremlinResponseBody
  GetData() *ExecKgGremlinResponseBodyData 
  SetHttpStatusCode(v int32) *ExecKgGremlinResponseBody
  GetHttpStatusCode() *int32 
  SetMessage(v string) *ExecKgGremlinResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExecKgGremlinResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecKgGremlinResponseBody
  GetSuccess() *bool 
}

type ExecKgGremlinResponseBody struct {
  // The backend response code.
  // 
  // example:
  // 
  // OK
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // The query result.
  Data *ExecKgGremlinResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s ExecKgGremlinResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecKgGremlinResponseBody) GoString() string {
  return s.String()
}

func (s *ExecKgGremlinResponseBody) GetCode() *string  {
  return s.Code
}

func (s *ExecKgGremlinResponseBody) GetData() *ExecKgGremlinResponseBodyData  {
  return s.Data
}

func (s *ExecKgGremlinResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *ExecKgGremlinResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExecKgGremlinResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecKgGremlinResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecKgGremlinResponseBody) SetCode(v string) *ExecKgGremlinResponseBody {
  s.Code = &v
  return s
}

func (s *ExecKgGremlinResponseBody) SetData(v *ExecKgGremlinResponseBodyData) *ExecKgGremlinResponseBody {
  s.Data = v
  return s
}

func (s *ExecKgGremlinResponseBody) SetHttpStatusCode(v int32) *ExecKgGremlinResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *ExecKgGremlinResponseBody) SetMessage(v string) *ExecKgGremlinResponseBody {
  s.Message = &v
  return s
}

func (s *ExecKgGremlinResponseBody) SetRequestId(v string) *ExecKgGremlinResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecKgGremlinResponseBody) SetSuccess(v bool) *ExecKgGremlinResponseBody {
  s.Success = &v
  return s
}

func (s *ExecKgGremlinResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecKgGremlinResponseBodyData struct {
  // The list of edges.
  EdgeList []*ExecKgGremlinResponseBodyDataEdgeList `json:"EdgeList,omitempty" xml:"EdgeList,omitempty" type:"Repeated"`
  // The transformed execution statement.
  // 
  // example:
  // 
  // g.v().limit(100)
  ExecQuery *string `json:"ExecQuery,omitempty" xml:"ExecQuery,omitempty"`
  // The list of nodes.
  NodeList []*ExecKgGremlinResponseBodyDataNodeList `json:"NodeList,omitempty" xml:"NodeList,omitempty" type:"Repeated"`
  // The list of rows.
  RowList []*ExecKgGremlinResponseBodyDataRowList `json:"RowList,omitempty" xml:"RowList,omitempty" type:"Repeated"`
}

func (s ExecKgGremlinResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s ExecKgGremlinResponseBodyData) GoString() string {
  return s.String()
}

func (s *ExecKgGremlinResponseBodyData) GetEdgeList() []*ExecKgGremlinResponseBodyDataEdgeList  {
  return s.EdgeList
}

func (s *ExecKgGremlinResponseBodyData) GetExecQuery() *string  {
  return s.ExecQuery
}

func (s *ExecKgGremlinResponseBodyData) GetNodeList() []*ExecKgGremlinResponseBodyDataNodeList  {
  return s.NodeList
}

func (s *ExecKgGremlinResponseBodyData) GetRowList() []*ExecKgGremlinResponseBodyDataRowList  {
  return s.RowList
}

func (s *ExecKgGremlinResponseBodyData) SetEdgeList(v []*ExecKgGremlinResponseBodyDataEdgeList) *ExecKgGremlinResponseBodyData {
  s.EdgeList = v
  return s
}

func (s *ExecKgGremlinResponseBodyData) SetExecQuery(v string) *ExecKgGremlinResponseBodyData {
  s.ExecQuery = &v
  return s
}

func (s *ExecKgGremlinResponseBodyData) SetNodeList(v []*ExecKgGremlinResponseBodyDataNodeList) *ExecKgGremlinResponseBodyData {
  s.NodeList = v
  return s
}

func (s *ExecKgGremlinResponseBodyData) SetRowList(v []*ExecKgGremlinResponseBodyDataRowList) *ExecKgGremlinResponseBodyData {
  s.RowList = v
  return s
}

func (s *ExecKgGremlinResponseBodyData) Validate() error {
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

type ExecKgGremlinResponseBodyDataEdgeList struct {
  // The data ID of the relationship record.
  // 
  // example:
  // 
  // abcd-1235-xc
  DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
  // The list of relationship record properties.
  Properties []*ExecKgGremlinResponseBodyDataEdgeListProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Repeated"`
  // The relationship type.
  // 
  // example:
  // 
  // BUY
  RelationType *string `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
  // The data ID of the source entity record.
  // 
  // example:
  // 
  // abcd-1234
  SourceEntityDataId *string `json:"SourceEntityDataId,omitempty" xml:"SourceEntityDataId,omitempty"`
  // The source entity type.
  // 
  // example:
  // 
  // Shop
  SourceEntityType *string `json:"SourceEntityType,omitempty" xml:"SourceEntityType,omitempty"`
  // The data ID of the target entity record.
  // 
  // example:
  // 
  // abcd-2234
  TargetEntityDataId *string `json:"TargetEntityDataId,omitempty" xml:"TargetEntityDataId,omitempty"`
  // The target entity type.
  // 
  // example:
  // 
  // Product
  TargetEntityType *string `json:"TargetEntityType,omitempty" xml:"TargetEntityType,omitempty"`
}

func (s ExecKgGremlinResponseBodyDataEdgeList) String() string {
  return dara.Prettify(s)
}

func (s ExecKgGremlinResponseBodyDataEdgeList) GoString() string {
  return s.String()
}

func (s *ExecKgGremlinResponseBodyDataEdgeList) GetDataId() *string  {
  return s.DataId
}

func (s *ExecKgGremlinResponseBodyDataEdgeList) GetProperties() []*ExecKgGremlinResponseBodyDataEdgeListProperties  {
  return s.Properties
}

func (s *ExecKgGremlinResponseBodyDataEdgeList) GetRelationType() *string  {
  return s.RelationType
}

func (s *ExecKgGremlinResponseBodyDataEdgeList) GetSourceEntityDataId() *string  {
  return s.SourceEntityDataId
}

func (s *ExecKgGremlinResponseBodyDataEdgeList) GetSourceEntityType() *string  {
  return s.SourceEntityType
}

func (s *ExecKgGremlinResponseBodyDataEdgeList) GetTargetEntityDataId() *string  {
  return s.TargetEntityDataId
}

func (s *ExecKgGremlinResponseBodyDataEdgeList) GetTargetEntityType() *string  {
  return s.TargetEntityType
}

func (s *ExecKgGremlinResponseBodyDataEdgeList) SetDataId(v string) *ExecKgGremlinResponseBodyDataEdgeList {
  s.DataId = &v
  return s
}

func (s *ExecKgGremlinResponseBodyDataEdgeList) SetProperties(v []*ExecKgGremlinResponseBodyDataEdgeListProperties) *ExecKgGremlinResponseBodyDataEdgeList {
  s.Properties = v
  return s
}

func (s *ExecKgGremlinResponseBodyDataEdgeList) SetRelationType(v string) *ExecKgGremlinResponseBodyDataEdgeList {
  s.RelationType = &v
  return s
}

func (s *ExecKgGremlinResponseBodyDataEdgeList) SetSourceEntityDataId(v string) *ExecKgGremlinResponseBodyDataEdgeList {
  s.SourceEntityDataId = &v
  return s
}

func (s *ExecKgGremlinResponseBodyDataEdgeList) SetSourceEntityType(v string) *ExecKgGremlinResponseBodyDataEdgeList {
  s.SourceEntityType = &v
  return s
}

func (s *ExecKgGremlinResponseBodyDataEdgeList) SetTargetEntityDataId(v string) *ExecKgGremlinResponseBodyDataEdgeList {
  s.TargetEntityDataId = &v
  return s
}

func (s *ExecKgGremlinResponseBodyDataEdgeList) SetTargetEntityType(v string) *ExecKgGremlinResponseBodyDataEdgeList {
  s.TargetEntityType = &v
  return s
}

func (s *ExecKgGremlinResponseBodyDataEdgeList) Validate() error {
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

type ExecKgGremlinResponseBodyDataEdgeListProperties struct {
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

func (s ExecKgGremlinResponseBodyDataEdgeListProperties) String() string {
  return dara.Prettify(s)
}

func (s ExecKgGremlinResponseBodyDataEdgeListProperties) GoString() string {
  return s.String()
}

func (s *ExecKgGremlinResponseBodyDataEdgeListProperties) GetCode() *string  {
  return s.Code
}

func (s *ExecKgGremlinResponseBodyDataEdgeListProperties) GetValue() *string  {
  return s.Value
}

func (s *ExecKgGremlinResponseBodyDataEdgeListProperties) SetCode(v string) *ExecKgGremlinResponseBodyDataEdgeListProperties {
  s.Code = &v
  return s
}

func (s *ExecKgGremlinResponseBodyDataEdgeListProperties) SetValue(v string) *ExecKgGremlinResponseBodyDataEdgeListProperties {
  s.Value = &v
  return s
}

func (s *ExecKgGremlinResponseBodyDataEdgeListProperties) Validate() error {
  return dara.Validate(s)
}

type ExecKgGremlinResponseBodyDataNodeList struct {
  // The data ID of the entity record.
  // 
  // example:
  // 
  // abcd-1234-xx
  DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
  // The entity type.
  // 
  // example:
  // 
  // Product
  EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
  // The list of entity record properties.
  Properties []*ExecKgGremlinResponseBodyDataNodeListProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Repeated"`
}

func (s ExecKgGremlinResponseBodyDataNodeList) String() string {
  return dara.Prettify(s)
}

func (s ExecKgGremlinResponseBodyDataNodeList) GoString() string {
  return s.String()
}

func (s *ExecKgGremlinResponseBodyDataNodeList) GetDataId() *string  {
  return s.DataId
}

func (s *ExecKgGremlinResponseBodyDataNodeList) GetEntityType() *string  {
  return s.EntityType
}

func (s *ExecKgGremlinResponseBodyDataNodeList) GetProperties() []*ExecKgGremlinResponseBodyDataNodeListProperties  {
  return s.Properties
}

func (s *ExecKgGremlinResponseBodyDataNodeList) SetDataId(v string) *ExecKgGremlinResponseBodyDataNodeList {
  s.DataId = &v
  return s
}

func (s *ExecKgGremlinResponseBodyDataNodeList) SetEntityType(v string) *ExecKgGremlinResponseBodyDataNodeList {
  s.EntityType = &v
  return s
}

func (s *ExecKgGremlinResponseBodyDataNodeList) SetProperties(v []*ExecKgGremlinResponseBodyDataNodeListProperties) *ExecKgGremlinResponseBodyDataNodeList {
  s.Properties = v
  return s
}

func (s *ExecKgGremlinResponseBodyDataNodeList) Validate() error {
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

type ExecKgGremlinResponseBodyDataNodeListProperties struct {
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

func (s ExecKgGremlinResponseBodyDataNodeListProperties) String() string {
  return dara.Prettify(s)
}

func (s ExecKgGremlinResponseBodyDataNodeListProperties) GoString() string {
  return s.String()
}

func (s *ExecKgGremlinResponseBodyDataNodeListProperties) GetCode() *string  {
  return s.Code
}

func (s *ExecKgGremlinResponseBodyDataNodeListProperties) GetValue() *string  {
  return s.Value
}

func (s *ExecKgGremlinResponseBodyDataNodeListProperties) SetCode(v string) *ExecKgGremlinResponseBodyDataNodeListProperties {
  s.Code = &v
  return s
}

func (s *ExecKgGremlinResponseBodyDataNodeListProperties) SetValue(v string) *ExecKgGremlinResponseBodyDataNodeListProperties {
  s.Value = &v
  return s
}

func (s *ExecKgGremlinResponseBodyDataNodeListProperties) Validate() error {
  return dara.Validate(s)
}

type ExecKgGremlinResponseBodyDataRowList struct {
  // The list of columns in the row.
  Columns []*ExecKgGremlinResponseBodyDataRowListColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
}

func (s ExecKgGremlinResponseBodyDataRowList) String() string {
  return dara.Prettify(s)
}

func (s ExecKgGremlinResponseBodyDataRowList) GoString() string {
  return s.String()
}

func (s *ExecKgGremlinResponseBodyDataRowList) GetColumns() []*ExecKgGremlinResponseBodyDataRowListColumns  {
  return s.Columns
}

func (s *ExecKgGremlinResponseBodyDataRowList) SetColumns(v []*ExecKgGremlinResponseBodyDataRowListColumns) *ExecKgGremlinResponseBodyDataRowList {
  s.Columns = v
  return s
}

func (s *ExecKgGremlinResponseBodyDataRowList) Validate() error {
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

type ExecKgGremlinResponseBodyDataRowListColumns struct {
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

func (s ExecKgGremlinResponseBodyDataRowListColumns) String() string {
  return dara.Prettify(s)
}

func (s ExecKgGremlinResponseBodyDataRowListColumns) GoString() string {
  return s.String()
}

func (s *ExecKgGremlinResponseBodyDataRowListColumns) GetCode() *string  {
  return s.Code
}

func (s *ExecKgGremlinResponseBodyDataRowListColumns) GetValue() *string  {
  return s.Value
}

func (s *ExecKgGremlinResponseBodyDataRowListColumns) SetCode(v string) *ExecKgGremlinResponseBodyDataRowListColumns {
  s.Code = &v
  return s
}

func (s *ExecKgGremlinResponseBodyDataRowListColumns) SetValue(v string) *ExecKgGremlinResponseBodyDataRowListColumns {
  s.Value = &v
  return s
}

func (s *ExecKgGremlinResponseBodyDataRowListColumns) Validate() error {
  return dara.Validate(s)
}

