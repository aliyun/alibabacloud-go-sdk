// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSceneGroupTemplateHeaders interface {
  dara.Model
  String() string
  GoString() string
  SetCommonHeaders(v map[string]*string) *EnableSceneGroupTemplateHeaders
  GetCommonHeaders() map[string]*string 
  SetAccountContext(v *EnableSceneGroupTemplateHeadersAccountContext) *EnableSceneGroupTemplateHeaders
  GetAccountContext() *EnableSceneGroupTemplateHeadersAccountContext 
}

type EnableSceneGroupTemplateHeaders struct {
  CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
  AccountContext *EnableSceneGroupTemplateHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s EnableSceneGroupTemplateHeaders) String() string {
  return dara.Prettify(s)
}

func (s EnableSceneGroupTemplateHeaders) GoString() string {
  return s.String()
}

func (s *EnableSceneGroupTemplateHeaders) GetCommonHeaders() map[string]*string  {
  return s.CommonHeaders
}

func (s *EnableSceneGroupTemplateHeaders) GetAccountContext() *EnableSceneGroupTemplateHeadersAccountContext  {
  return s.AccountContext
}

func (s *EnableSceneGroupTemplateHeaders) SetCommonHeaders(v map[string]*string) *EnableSceneGroupTemplateHeaders {
  s.CommonHeaders = v
  return s
}

func (s *EnableSceneGroupTemplateHeaders) SetAccountContext(v *EnableSceneGroupTemplateHeadersAccountContext) *EnableSceneGroupTemplateHeaders {
  s.AccountContext = v
  return s
}

func (s *EnableSceneGroupTemplateHeaders) Validate() error {
  if s.AccountContext != nil {
    if err := s.AccountContext.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EnableSceneGroupTemplateHeadersAccountContext struct {
  // This parameter is required.
  // 
  // example:
  // 
  // 012345
  AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s EnableSceneGroupTemplateHeadersAccountContext) String() string {
  return dara.Prettify(s)
}

func (s EnableSceneGroupTemplateHeadersAccountContext) GoString() string {
  return s.String()
}

func (s *EnableSceneGroupTemplateHeadersAccountContext) GetAccountId() *string  {
  return s.AccountId
}

func (s *EnableSceneGroupTemplateHeadersAccountContext) SetAccountId(v string) *EnableSceneGroupTemplateHeadersAccountContext {
  s.AccountId = &v
  return s
}

func (s *EnableSceneGroupTemplateHeadersAccountContext) Validate() error {
  return dara.Validate(s)
}

