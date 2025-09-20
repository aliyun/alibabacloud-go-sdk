// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iElementContent interface {
  dara.Model
  String() string
  GoString() string
  SetContent(v string) *ElementContent
  GetContent() *string 
  SetTimeRange(v []*int64) *ElementContent
  GetTimeRange() []*int64 
  SetType(v string) *ElementContent
  GetType() *string 
  SetURL(v string) *ElementContent
  GetURL() *string 
}

type ElementContent struct {
  Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
  TimeRange []*int64 `json:"TimeRange,omitempty" xml:"TimeRange,omitempty" type:"Repeated"`
  Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
  URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s ElementContent) String() string {
  return dara.Prettify(s)
}

func (s ElementContent) GoString() string {
  return s.String()
}

func (s *ElementContent) GetContent() *string  {
  return s.Content
}

func (s *ElementContent) GetTimeRange() []*int64  {
  return s.TimeRange
}

func (s *ElementContent) GetType() *string  {
  return s.Type
}

func (s *ElementContent) GetURL() *string  {
  return s.URL
}

func (s *ElementContent) SetContent(v string) *ElementContent {
  s.Content = &v
  return s
}

func (s *ElementContent) SetTimeRange(v []*int64) *ElementContent {
  s.TimeRange = v
  return s
}

func (s *ElementContent) SetType(v string) *ElementContent {
  s.Type = &v
  return s
}

func (s *ElementContent) SetURL(v string) *ElementContent {
  s.URL = &v
  return s
}

func (s *ElementContent) Validate() error {
  return dara.Validate(s)
}

