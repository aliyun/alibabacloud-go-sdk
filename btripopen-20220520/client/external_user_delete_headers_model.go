// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExternalUserDeleteHeaders interface {
  dara.Model
  String() string
  GoString() string
  SetCommonHeaders(v map[string]*string) *ExternalUserDeleteHeaders
  GetCommonHeaders() map[string]*string 
  SetXAcsBtripCorpToken(v string) *ExternalUserDeleteHeaders
  GetXAcsBtripCorpToken() *string 
}

type ExternalUserDeleteHeaders struct {
  CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
  // example:
  // 
  // feth00jqwls
  XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s ExternalUserDeleteHeaders) String() string {
  return dara.Prettify(s)
}

func (s ExternalUserDeleteHeaders) GoString() string {
  return s.String()
}

func (s *ExternalUserDeleteHeaders) GetCommonHeaders() map[string]*string  {
  return s.CommonHeaders
}

func (s *ExternalUserDeleteHeaders) GetXAcsBtripCorpToken() *string  {
  return s.XAcsBtripCorpToken
}

func (s *ExternalUserDeleteHeaders) SetCommonHeaders(v map[string]*string) *ExternalUserDeleteHeaders {
  s.CommonHeaders = v
  return s
}

func (s *ExternalUserDeleteHeaders) SetXAcsBtripCorpToken(v string) *ExternalUserDeleteHeaders {
  s.XAcsBtripCorpToken = &v
  return s
}

func (s *ExternalUserDeleteHeaders) Validate() error {
  return dara.Validate(s)
}

