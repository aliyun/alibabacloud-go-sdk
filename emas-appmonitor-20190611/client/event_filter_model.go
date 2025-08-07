// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEventFilter interface {
  dara.Model
  String() string
  GoString() string
  SetKey(v string) *EventFilter
  GetKey() *string 
  SetOp(v string) *EventFilter
  GetOp() *string 
  SetSubFilters(v []*EventFilter) *EventFilter
  GetSubFilters() []*EventFilter 
  SetValues(v []*string) *EventFilter
  GetValues() []*string 
}

type EventFilter struct {
  Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
  // This parameter is required.
  Op *string `json:"Op,omitempty" xml:"Op,omitempty"`
  SubFilters []*EventFilter `json:"SubFilters,omitempty" xml:"SubFilters,omitempty" type:"Repeated"`
  Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s EventFilter) String() string {
  return dara.Prettify(s)
}

func (s EventFilter) GoString() string {
  return s.String()
}

func (s *EventFilter) GetKey() *string  {
  return s.Key
}

func (s *EventFilter) GetOp() *string  {
  return s.Op
}

func (s *EventFilter) GetSubFilters() []*EventFilter  {
  return s.SubFilters
}

func (s *EventFilter) GetValues() []*string  {
  return s.Values
}

func (s *EventFilter) SetKey(v string) *EventFilter {
  s.Key = &v
  return s
}

func (s *EventFilter) SetOp(v string) *EventFilter {
  s.Op = &v
  return s
}

func (s *EventFilter) SetSubFilters(v []*EventFilter) *EventFilter {
  s.SubFilters = v
  return s
}

func (s *EventFilter) SetValues(v []*string) *EventFilter {
  s.Values = v
  return s
}

func (s *EventFilter) Validate() error {
  return dara.Validate(s)
}

