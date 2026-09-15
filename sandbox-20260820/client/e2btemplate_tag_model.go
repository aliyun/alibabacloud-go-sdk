// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iE2BTemplateTag interface {
  dara.Model
  String() string
  GoString() string
  SetBuildID(v string) *E2BTemplateTag
  GetBuildID() *string 
  SetCreatedAt(v string) *E2BTemplateTag
  GetCreatedAt() *string 
  SetTag(v string) *E2BTemplateTag
  GetTag() *string 
}

type E2BTemplateTag struct {
  BuildID *string `json:"buildID,omitempty" xml:"buildID,omitempty"`
  CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
  Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
}

func (s E2BTemplateTag) String() string {
  return dara.Prettify(s)
}

func (s E2BTemplateTag) GoString() string {
  return s.String()
}

func (s *E2BTemplateTag) GetBuildID() *string  {
  return s.BuildID
}

func (s *E2BTemplateTag) GetCreatedAt() *string  {
  return s.CreatedAt
}

func (s *E2BTemplateTag) GetTag() *string  {
  return s.Tag
}

func (s *E2BTemplateTag) SetBuildID(v string) *E2BTemplateTag {
  s.BuildID = &v
  return s
}

func (s *E2BTemplateTag) SetCreatedAt(v string) *E2BTemplateTag {
  s.CreatedAt = &v
  return s
}

func (s *E2BTemplateTag) SetTag(v string) *E2BTemplateTag {
  s.Tag = &v
  return s
}

func (s *E2BTemplateTag) Validate() error {
  return dara.Validate(s)
}

