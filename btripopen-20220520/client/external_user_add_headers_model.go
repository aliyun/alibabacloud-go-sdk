// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExternalUserAddHeaders interface {
  dara.Model
  String() string
  GoString() string
  SetCommonHeaders(v map[string]*string) *ExternalUserAddHeaders
  GetCommonHeaders() map[string]*string 
  SetXAcsBtripCorpToken(v string) *ExternalUserAddHeaders
  GetXAcsBtripCorpToken() *string 
}

type ExternalUserAddHeaders struct {
  CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
  // example:
  // 
  // feth00jqwls
  XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s ExternalUserAddHeaders) String() string {
  return dara.Prettify(s)
}

func (s ExternalUserAddHeaders) GoString() string {
  return s.String()
}

func (s *ExternalUserAddHeaders) GetCommonHeaders() map[string]*string  {
  return s.CommonHeaders
}

func (s *ExternalUserAddHeaders) GetXAcsBtripCorpToken() *string  {
  return s.XAcsBtripCorpToken
}

func (s *ExternalUserAddHeaders) SetCommonHeaders(v map[string]*string) *ExternalUserAddHeaders {
  s.CommonHeaders = v
  return s
}

func (s *ExternalUserAddHeaders) SetXAcsBtripCorpToken(v string) *ExternalUserAddHeaders {
  s.XAcsBtripCorpToken = &v
  return s
}

func (s *ExternalUserAddHeaders) Validate() error {
  return dara.Validate(s)
}

