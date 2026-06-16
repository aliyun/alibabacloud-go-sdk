// This file is auto-generated, don't edit it. Thanks.
package multiModalGuardWs

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iModelRequestMessageEvent interface {
  dara.Model
  String() string
  GoString() string
  SetService(v string) *ModelRequestMessageEvent
  GetService() *string 
  SetServiceParameters(v string) *ModelRequestMessageEvent
  GetServiceParameters() *string 
  SetDataType(v string) *ModelRequestMessageEvent
  GetDataType() *string 
  SetSync(v bool) *ModelRequestMessageEvent
  GetSync() *bool 
  SetData(v string) *ModelRequestMessageEvent
  GetData() *string 
}

type ModelRequestMessageEvent struct {
  Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
  ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
  DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
  Sync *bool `json:"Sync,omitempty" xml:"Sync,omitempty"`
  Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s ModelRequestMessageEvent) String() string {
  return dara.Prettify(s)
}

func (s ModelRequestMessageEvent) GoString() string {
  return s.String()
}

func (s *ModelRequestMessageEvent) GetService() *string  {
  return s.Service
}

func (s *ModelRequestMessageEvent) GetServiceParameters() *string  {
  return s.ServiceParameters
}

func (s *ModelRequestMessageEvent) GetDataType() *string  {
  return s.DataType
}

func (s *ModelRequestMessageEvent) GetSync() *bool  {
  return s.Sync
}

func (s *ModelRequestMessageEvent) GetData() *string  {
  return s.Data
}

func (s *ModelRequestMessageEvent) SetService(v string) *ModelRequestMessageEvent {
  s.Service = &v
  return s
}

func (s *ModelRequestMessageEvent) SetServiceParameters(v string) *ModelRequestMessageEvent {
  s.ServiceParameters = &v
  return s
}

func (s *ModelRequestMessageEvent) SetDataType(v string) *ModelRequestMessageEvent {
  s.DataType = &v
  return s
}

func (s *ModelRequestMessageEvent) SetSync(v bool) *ModelRequestMessageEvent {
  s.Sync = &v
  return s
}

func (s *ModelRequestMessageEvent) SetData(v string) *ModelRequestMessageEvent {
  s.Data = &v
  return s
}

func (s *ModelRequestMessageEvent) Validate() error {
  return dara.Validate(s)
}

