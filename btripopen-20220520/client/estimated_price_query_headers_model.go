// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEstimatedPriceQueryHeaders interface {
  dara.Model
  String() string
  GoString() string
  SetCommonHeaders(v map[string]*string) *EstimatedPriceQueryHeaders
  GetCommonHeaders() map[string]*string 
  SetXAcsBtripSoCorpToken(v string) *EstimatedPriceQueryHeaders
  GetXAcsBtripSoCorpToken() *string 
}

type EstimatedPriceQueryHeaders struct {
  CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
  // example:
  // 
  // feth00jqwls
  XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s EstimatedPriceQueryHeaders) String() string {
  return dara.Prettify(s)
}

func (s EstimatedPriceQueryHeaders) GoString() string {
  return s.String()
}

func (s *EstimatedPriceQueryHeaders) GetCommonHeaders() map[string]*string  {
  return s.CommonHeaders
}

func (s *EstimatedPriceQueryHeaders) GetXAcsBtripSoCorpToken() *string  {
  return s.XAcsBtripSoCorpToken
}

func (s *EstimatedPriceQueryHeaders) SetCommonHeaders(v map[string]*string) *EstimatedPriceQueryHeaders {
  s.CommonHeaders = v
  return s
}

func (s *EstimatedPriceQueryHeaders) SetXAcsBtripSoCorpToken(v string) *EstimatedPriceQueryHeaders {
  s.XAcsBtripSoCorpToken = &v
  return s
}

func (s *EstimatedPriceQueryHeaders) Validate() error {
  return dara.Validate(s)
}

