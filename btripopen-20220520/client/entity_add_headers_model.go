// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEntityAddHeaders interface {
  dara.Model
  String() string
  GoString() string
  SetCommonHeaders(v map[string]*string) *EntityAddHeaders
  GetCommonHeaders() map[string]*string 
  SetXAcsBtripSoCorpToken(v string) *EntityAddHeaders
  GetXAcsBtripSoCorpToken() *string 
}

type EntityAddHeaders struct {
  CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
  // example:
  // 
  // feth00jqwls
  XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s EntityAddHeaders) String() string {
  return dara.Prettify(s)
}

func (s EntityAddHeaders) GoString() string {
  return s.String()
}

func (s *EntityAddHeaders) GetCommonHeaders() map[string]*string  {
  return s.CommonHeaders
}

func (s *EntityAddHeaders) GetXAcsBtripSoCorpToken() *string  {
  return s.XAcsBtripSoCorpToken
}

func (s *EntityAddHeaders) SetCommonHeaders(v map[string]*string) *EntityAddHeaders {
  s.CommonHeaders = v
  return s
}

func (s *EntityAddHeaders) SetXAcsBtripSoCorpToken(v string) *EntityAddHeaders {
  s.XAcsBtripSoCorpToken = &v
  return s
}

func (s *EntityAddHeaders) Validate() error {
  return dara.Validate(s)
}

