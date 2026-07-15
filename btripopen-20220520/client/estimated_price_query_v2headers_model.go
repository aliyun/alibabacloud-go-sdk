// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEstimatedPriceQueryV2Headers interface {
  dara.Model
  String() string
  GoString() string
  SetCommonHeaders(v map[string]*string) *EstimatedPriceQueryV2Headers
  GetCommonHeaders() map[string]*string 
  SetXAcsBtripSoCorpToken(v string) *EstimatedPriceQueryV2Headers
  GetXAcsBtripSoCorpToken() *string 
}

type EstimatedPriceQueryV2Headers struct {
  CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
  XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s EstimatedPriceQueryV2Headers) String() string {
  return dara.Prettify(s)
}

func (s EstimatedPriceQueryV2Headers) GoString() string {
  return s.String()
}

func (s *EstimatedPriceQueryV2Headers) GetCommonHeaders() map[string]*string  {
  return s.CommonHeaders
}

func (s *EstimatedPriceQueryV2Headers) GetXAcsBtripSoCorpToken() *string  {
  return s.XAcsBtripSoCorpToken
}

func (s *EstimatedPriceQueryV2Headers) SetCommonHeaders(v map[string]*string) *EstimatedPriceQueryV2Headers {
  s.CommonHeaders = v
  return s
}

func (s *EstimatedPriceQueryV2Headers) SetXAcsBtripSoCorpToken(v string) *EstimatedPriceQueryV2Headers {
  s.XAcsBtripSoCorpToken = &v
  return s
}

func (s *EstimatedPriceQueryV2Headers) Validate() error {
  return dara.Validate(s)
}

