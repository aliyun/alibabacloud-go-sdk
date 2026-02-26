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
  // The content of the element.
  // 
  // If the value of the Type parameter is image or link, this parameter indicates the placeholder text.
  // 
  // example:
  // 
  // Text
  Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
  // The time range. The array length is fixed to 2. One element indicates the start time and the other one indicates the end time. Unit: milliseconds.
  TimeRange []*int64 `json:"TimeRange,omitempty" xml:"TimeRange,omitempty" type:"Repeated"`
  // The type of the element content.
  // 
  // Valid values:
  // 
  // 	- text
  // 
  // 	- image
  // 
  // 	- link
  // 
  // example:
  // 
  // text
  Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
  // The link to the element content. This parameter takes effect only if the Type parameter is set to image or link.
  // 
  // example:
  // 
  // http://aliyun.com
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

