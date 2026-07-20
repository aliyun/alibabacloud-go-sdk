// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExternalUserUpdateHeaders interface {
  dara.Model
  String() string
  GoString() string
  SetCommonHeaders(v map[string]*string) *ExternalUserUpdateHeaders
  GetCommonHeaders() map[string]*string 
  SetXAcsBtripCorpToken(v string) *ExternalUserUpdateHeaders
  GetXAcsBtripCorpToken() *string 
}

type ExternalUserUpdateHeaders struct {
  CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
  // example:
  // 
  // feth00jqwls
  XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s ExternalUserUpdateHeaders) String() string {
  return dara.Prettify(s)
}

func (s ExternalUserUpdateHeaders) GoString() string {
  return s.String()
}

func (s *ExternalUserUpdateHeaders) GetCommonHeaders() map[string]*string  {
  return s.CommonHeaders
}

func (s *ExternalUserUpdateHeaders) GetXAcsBtripCorpToken() *string  {
  return s.XAcsBtripCorpToken
}

func (s *ExternalUserUpdateHeaders) SetCommonHeaders(v map[string]*string) *ExternalUserUpdateHeaders {
  s.CommonHeaders = v
  return s
}

func (s *ExternalUserUpdateHeaders) SetXAcsBtripCorpToken(v string) *ExternalUserUpdateHeaders {
  s.XAcsBtripCorpToken = &v
  return s
}

func (s *ExternalUserUpdateHeaders) Validate() error {
  return dara.Validate(s)
}

