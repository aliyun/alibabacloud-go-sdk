// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEventCenterQueryEventsShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBodyShrink(v string) *EventCenterQueryEventsShrinkRequest
  GetBodyShrink() *string 
  SetBusName(v string) *EventCenterQueryEventsShrinkRequest
  GetBusName() *string 
  SetMaxResults(v int32) *EventCenterQueryEventsShrinkRequest
  GetMaxResults() *int32 
  SetNextToken(v string) *EventCenterQueryEventsShrinkRequest
  GetNextToken() *string 
}

type EventCenterQueryEventsShrinkRequest struct {
  // The request body.
  // 
  // This parameter is required.
  BodyShrink *string `json:"Body,omitempty" xml:"Body,omitempty"`
  // The name of the event bus.
  // 
  // example:
  // 
  // default
  BusName *string `json:"BusName,omitempty" xml:"BusName,omitempty"`
  // The number of entries per page. Valid values: 0 to 10000. Default value: 100.
  // 
  // example:
  // 
  // 100
  MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
  // 用来标记当前开始读取的位置。置空表示从头开始。
  // 
  // example:
  // 
  // 100
  NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s EventCenterQueryEventsShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s EventCenterQueryEventsShrinkRequest) GoString() string {
  return s.String()
}

func (s *EventCenterQueryEventsShrinkRequest) GetBodyShrink() *string  {
  return s.BodyShrink
}

func (s *EventCenterQueryEventsShrinkRequest) GetBusName() *string  {
  return s.BusName
}

func (s *EventCenterQueryEventsShrinkRequest) GetMaxResults() *int32  {
  return s.MaxResults
}

func (s *EventCenterQueryEventsShrinkRequest) GetNextToken() *string  {
  return s.NextToken
}

func (s *EventCenterQueryEventsShrinkRequest) SetBodyShrink(v string) *EventCenterQueryEventsShrinkRequest {
  s.BodyShrink = &v
  return s
}

func (s *EventCenterQueryEventsShrinkRequest) SetBusName(v string) *EventCenterQueryEventsShrinkRequest {
  s.BusName = &v
  return s
}

func (s *EventCenterQueryEventsShrinkRequest) SetMaxResults(v int32) *EventCenterQueryEventsShrinkRequest {
  s.MaxResults = &v
  return s
}

func (s *EventCenterQueryEventsShrinkRequest) SetNextToken(v string) *EventCenterQueryEventsShrinkRequest {
  s.NextToken = &v
  return s
}

func (s *EventCenterQueryEventsShrinkRequest) Validate() error {
  return dara.Validate(s)
}

