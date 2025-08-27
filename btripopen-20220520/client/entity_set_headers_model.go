// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEntitySetHeaders interface {
  dara.Model
  String() string
  GoString() string
  SetCommonHeaders(v map[string]*string) *EntitySetHeaders
  GetCommonHeaders() map[string]*string 
  SetXAcsBtripSoCorpToken(v string) *EntitySetHeaders
  GetXAcsBtripSoCorpToken() *string 
}

type EntitySetHeaders struct {
  CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
  // example:
  // 
  // feth00jqwls
  XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s EntitySetHeaders) String() string {
  return dara.Prettify(s)
}

func (s EntitySetHeaders) GoString() string {
  return s.String()
}

func (s *EntitySetHeaders) GetCommonHeaders() map[string]*string  {
  return s.CommonHeaders
}

func (s *EntitySetHeaders) GetXAcsBtripSoCorpToken() *string  {
  return s.XAcsBtripSoCorpToken
}

func (s *EntitySetHeaders) SetCommonHeaders(v map[string]*string) *EntitySetHeaders {
  s.CommonHeaders = v
  return s
}

func (s *EntitySetHeaders) SetXAcsBtripSoCorpToken(v string) *EntitySetHeaders {
  s.XAcsBtripSoCorpToken = &v
  return s
}

func (s *EntitySetHeaders) Validate() error {
  return dara.Validate(s)
}

