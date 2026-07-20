// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEntityDeleteHeaders interface {
  dara.Model
  String() string
  GoString() string
  SetCommonHeaders(v map[string]*string) *EntityDeleteHeaders
  GetCommonHeaders() map[string]*string 
  SetXAcsBtripSoCorpToken(v string) *EntityDeleteHeaders
  GetXAcsBtripSoCorpToken() *string 
}

type EntityDeleteHeaders struct {
  CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
  // example:
  // 
  // feth00jqwls
  XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s EntityDeleteHeaders) String() string {
  return dara.Prettify(s)
}

func (s EntityDeleteHeaders) GoString() string {
  return s.String()
}

func (s *EntityDeleteHeaders) GetCommonHeaders() map[string]*string  {
  return s.CommonHeaders
}

func (s *EntityDeleteHeaders) GetXAcsBtripSoCorpToken() *string  {
  return s.XAcsBtripSoCorpToken
}

func (s *EntityDeleteHeaders) SetCommonHeaders(v map[string]*string) *EntityDeleteHeaders {
  s.CommonHeaders = v
  return s
}

func (s *EntityDeleteHeaders) SetXAcsBtripSoCorpToken(v string) *EntityDeleteHeaders {
  s.XAcsBtripSoCorpToken = &v
  return s
}

func (s *EntityDeleteHeaders) Validate() error {
  return dara.Validate(s)
}

