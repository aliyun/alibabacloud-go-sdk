// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExternalUserQueryHeaders interface {
  dara.Model
  String() string
  GoString() string
  SetCommonHeaders(v map[string]*string) *ExternalUserQueryHeaders
  GetCommonHeaders() map[string]*string 
  SetXAcsBtripCorpToken(v string) *ExternalUserQueryHeaders
  GetXAcsBtripCorpToken() *string 
}

type ExternalUserQueryHeaders struct {
  CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
  // example:
  // 
  // feth00jqwls
  XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s ExternalUserQueryHeaders) String() string {
  return dara.Prettify(s)
}

func (s ExternalUserQueryHeaders) GoString() string {
  return s.String()
}

func (s *ExternalUserQueryHeaders) GetCommonHeaders() map[string]*string  {
  return s.CommonHeaders
}

func (s *ExternalUserQueryHeaders) GetXAcsBtripCorpToken() *string  {
  return s.XAcsBtripCorpToken
}

func (s *ExternalUserQueryHeaders) SetCommonHeaders(v map[string]*string) *ExternalUserQueryHeaders {
  s.CommonHeaders = v
  return s
}

func (s *ExternalUserQueryHeaders) SetXAcsBtripCorpToken(v string) *ExternalUserQueryHeaders {
  s.XAcsBtripCorpToken = &v
  return s
}

func (s *ExternalUserQueryHeaders) Validate() error {
  return dara.Validate(s)
}

