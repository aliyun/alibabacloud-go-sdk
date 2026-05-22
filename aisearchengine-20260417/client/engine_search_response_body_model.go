// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEngineSearchResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v int32) *EngineSearchResponseBody
  GetCode() *int32 
  SetData(v *EngineSearchResponseBodyData) *EngineSearchResponseBody
  GetData() *EngineSearchResponseBodyData 
  SetMessage(v string) *EngineSearchResponseBody
  GetMessage() *string 
}

type EngineSearchResponseBody struct {
  // example:
  // 
  // SUCCESS
  Code *int32 `json:"code,omitempty" xml:"code,omitempty"`
  Data *EngineSearchResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
  // example:
  // 
  // successful
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s EngineSearchResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EngineSearchResponseBody) GoString() string {
  return s.String()
}

func (s *EngineSearchResponseBody) GetCode() *int32  {
  return s.Code
}

func (s *EngineSearchResponseBody) GetData() *EngineSearchResponseBodyData  {
  return s.Data
}

func (s *EngineSearchResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EngineSearchResponseBody) SetCode(v int32) *EngineSearchResponseBody {
  s.Code = &v
  return s
}

func (s *EngineSearchResponseBody) SetData(v *EngineSearchResponseBodyData) *EngineSearchResponseBody {
  s.Data = v
  return s
}

func (s *EngineSearchResponseBody) SetMessage(v string) *EngineSearchResponseBody {
  s.Message = &v
  return s
}

func (s *EngineSearchResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EngineSearchResponseBodyData struct {
  // example:
  // 
  // 搜索词违禁
  ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
  // example:
  // 
  // {}
  Extra map[string]interface{} `json:"extra,omitempty" xml:"extra,omitempty"`
  Items []*EngineSearchResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
  // example:
  // 
  // 1
  Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
  // example:
  // 
  // 62CC3999-5380-5344-ABC7-46453625A910
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // example:
  // 
  // 10
  Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
  // example:
  // 
  // 200
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  // example:
  // 
  // 121
  Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
  // example:
  // 
  // {"traceId":"234234324"}
  TraceInfo map[string]interface{} `json:"traceInfo,omitempty" xml:"traceInfo,omitempty"`
}

func (s EngineSearchResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s EngineSearchResponseBodyData) GoString() string {
  return s.String()
}

func (s *EngineSearchResponseBodyData) GetErrorMessage() *string  {
  return s.ErrorMessage
}

func (s *EngineSearchResponseBodyData) GetExtra() map[string]interface{}  {
  return s.Extra
}

func (s *EngineSearchResponseBodyData) GetItems() []*EngineSearchResponseBodyDataItems  {
  return s.Items
}

func (s *EngineSearchResponseBodyData) GetPage() *int32  {
  return s.Page
}

func (s *EngineSearchResponseBodyData) GetRequestId() *string  {
  return s.RequestId
}

func (s *EngineSearchResponseBodyData) GetSize() *int32  {
  return s.Size
}

func (s *EngineSearchResponseBodyData) GetStatus() *string  {
  return s.Status
}

func (s *EngineSearchResponseBodyData) GetTotal() *int32  {
  return s.Total
}

func (s *EngineSearchResponseBodyData) GetTraceInfo() map[string]interface{}  {
  return s.TraceInfo
}

func (s *EngineSearchResponseBodyData) SetErrorMessage(v string) *EngineSearchResponseBodyData {
  s.ErrorMessage = &v
  return s
}

func (s *EngineSearchResponseBodyData) SetExtra(v map[string]interface{}) *EngineSearchResponseBodyData {
  s.Extra = v
  return s
}

func (s *EngineSearchResponseBodyData) SetItems(v []*EngineSearchResponseBodyDataItems) *EngineSearchResponseBodyData {
  s.Items = v
  return s
}

func (s *EngineSearchResponseBodyData) SetPage(v int32) *EngineSearchResponseBodyData {
  s.Page = &v
  return s
}

func (s *EngineSearchResponseBodyData) SetRequestId(v string) *EngineSearchResponseBodyData {
  s.RequestId = &v
  return s
}

func (s *EngineSearchResponseBodyData) SetSize(v int32) *EngineSearchResponseBodyData {
  s.Size = &v
  return s
}

func (s *EngineSearchResponseBodyData) SetStatus(v string) *EngineSearchResponseBodyData {
  s.Status = &v
  return s
}

func (s *EngineSearchResponseBodyData) SetTotal(v int32) *EngineSearchResponseBodyData {
  s.Total = &v
  return s
}

func (s *EngineSearchResponseBodyData) SetTraceInfo(v map[string]interface{}) *EngineSearchResponseBodyData {
  s.TraceInfo = v
  return s
}

func (s *EngineSearchResponseBodyData) Validate() error {
  if s.Items != nil {
    for _, item := range s.Items {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type EngineSearchResponseBodyDataItems struct {
  // example:
  // 
  // {"rankScore":1.0}
  Algorithm map[string]interface{} `json:"algorithm,omitempty" xml:"algorithm,omitempty"`
  // example:
  // 
  // {"titile":"大疆无人机"}
  Content map[string]interface{} `json:"content,omitempty" xml:"content,omitempty"`
  // example:
  // 
  // iceberg-tbl-b7420156-6d23-4555-8e05-12285b36d6ca
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
  // example:
  // 
  // 1.75
  Score *float32 `json:"score,omitempty" xml:"score,omitempty"`
  // example:
  // 
  // {"traceId":"12285b36d6ca"}
  TraceInfo map[string]interface{} `json:"traceInfo,omitempty" xml:"traceInfo,omitempty"`
}

func (s EngineSearchResponseBodyDataItems) String() string {
  return dara.Prettify(s)
}

func (s EngineSearchResponseBodyDataItems) GoString() string {
  return s.String()
}

func (s *EngineSearchResponseBodyDataItems) GetAlgorithm() map[string]interface{}  {
  return s.Algorithm
}

func (s *EngineSearchResponseBodyDataItems) GetContent() map[string]interface{}  {
  return s.Content
}

func (s *EngineSearchResponseBodyDataItems) GetId() *string  {
  return s.Id
}

func (s *EngineSearchResponseBodyDataItems) GetScore() *float32  {
  return s.Score
}

func (s *EngineSearchResponseBodyDataItems) GetTraceInfo() map[string]interface{}  {
  return s.TraceInfo
}

func (s *EngineSearchResponseBodyDataItems) SetAlgorithm(v map[string]interface{}) *EngineSearchResponseBodyDataItems {
  s.Algorithm = v
  return s
}

func (s *EngineSearchResponseBodyDataItems) SetContent(v map[string]interface{}) *EngineSearchResponseBodyDataItems {
  s.Content = v
  return s
}

func (s *EngineSearchResponseBodyDataItems) SetId(v string) *EngineSearchResponseBodyDataItems {
  s.Id = &v
  return s
}

func (s *EngineSearchResponseBodyDataItems) SetScore(v float32) *EngineSearchResponseBodyDataItems {
  s.Score = &v
  return s
}

func (s *EngineSearchResponseBodyDataItems) SetTraceInfo(v map[string]interface{}) *EngineSearchResponseBodyDataItems {
  s.TraceInfo = v
  return s
}

func (s *EngineSearchResponseBodyDataItems) Validate() error {
  return dara.Validate(s)
}

