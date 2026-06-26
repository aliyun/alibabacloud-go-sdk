// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iElasticInterval interface {
  dara.Model
  String() string
  GoString() string
  SetEndTime(v string) *ElasticInterval
  GetEndTime() *string 
  SetStartTime(v string) *ElasticInterval
  GetStartTime() *string 
}

type ElasticInterval struct {
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ElasticInterval) String() string {
  return dara.Prettify(s)
}

func (s ElasticInterval) GoString() string {
  return s.String()
}

func (s *ElasticInterval) GetEndTime() *string  {
  return s.EndTime
}

func (s *ElasticInterval) GetStartTime() *string  {
  return s.StartTime
}

func (s *ElasticInterval) SetEndTime(v string) *ElasticInterval {
  s.EndTime = &v
  return s
}

func (s *ElasticInterval) SetStartTime(v string) *ElasticInterval {
  s.StartTime = &v
  return s
}

func (s *ElasticInterval) Validate() error {
  return dara.Validate(s)
}

