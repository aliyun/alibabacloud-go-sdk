// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSceneGroupTemplateShrinkHeaders interface {
  dara.Model
  String() string
  GoString() string
  SetCommonHeaders(v map[string]*string) *EnableSceneGroupTemplateShrinkHeaders
  GetCommonHeaders() map[string]*string 
  SetAccountContextShrink(v string) *EnableSceneGroupTemplateShrinkHeaders
  GetAccountContextShrink() *string 
}

type EnableSceneGroupTemplateShrinkHeaders struct {
  CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
  AccountContextShrink *string `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s EnableSceneGroupTemplateShrinkHeaders) String() string {
  return dara.Prettify(s)
}

func (s EnableSceneGroupTemplateShrinkHeaders) GoString() string {
  return s.String()
}

func (s *EnableSceneGroupTemplateShrinkHeaders) GetCommonHeaders() map[string]*string  {
  return s.CommonHeaders
}

func (s *EnableSceneGroupTemplateShrinkHeaders) GetAccountContextShrink() *string  {
  return s.AccountContextShrink
}

func (s *EnableSceneGroupTemplateShrinkHeaders) SetCommonHeaders(v map[string]*string) *EnableSceneGroupTemplateShrinkHeaders {
  s.CommonHeaders = v
  return s
}

func (s *EnableSceneGroupTemplateShrinkHeaders) SetAccountContextShrink(v string) *EnableSceneGroupTemplateShrinkHeaders {
  s.AccountContextShrink = &v
  return s
}

func (s *EnableSceneGroupTemplateShrinkHeaders) Validate() error {
  return dara.Validate(s)
}

