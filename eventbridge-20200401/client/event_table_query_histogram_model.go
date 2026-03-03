// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEventTableQueryHistogram interface {
  dara.Model
  String() string
  GoString() string
  SetCount(v int64) *EventTableQueryHistogram
  GetCount() *int64 
  SetFrom(v int64) *EventTableQueryHistogram
  GetFrom() *int64 
  SetTo(v int64) *EventTableQueryHistogram
  GetTo() *int64 
}

type EventTableQueryHistogram struct {
  Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
  From *int64 `json:"From,omitempty" xml:"From,omitempty"`
  To *int64 `json:"To,omitempty" xml:"To,omitempty"`
}

func (s EventTableQueryHistogram) String() string {
  return dara.Prettify(s)
}

func (s EventTableQueryHistogram) GoString() string {
  return s.String()
}

func (s *EventTableQueryHistogram) GetCount() *int64  {
  return s.Count
}

func (s *EventTableQueryHistogram) GetFrom() *int64  {
  return s.From
}

func (s *EventTableQueryHistogram) GetTo() *int64  {
  return s.To
}

func (s *EventTableQueryHistogram) SetCount(v int64) *EventTableQueryHistogram {
  s.Count = &v
  return s
}

func (s *EventTableQueryHistogram) SetFrom(v int64) *EventTableQueryHistogram {
  s.From = &v
  return s
}

func (s *EventTableQueryHistogram) SetTo(v int64) *EventTableQueryHistogram {
  s.To = &v
  return s
}

func (s *EventTableQueryHistogram) Validate() error {
  return dara.Validate(s)
}

