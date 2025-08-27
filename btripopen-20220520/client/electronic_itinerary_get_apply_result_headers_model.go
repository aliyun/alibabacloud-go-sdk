// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iElectronicItineraryGetApplyResultHeaders interface {
  dara.Model
  String() string
  GoString() string
  SetCommonHeaders(v map[string]*string) *ElectronicItineraryGetApplyResultHeaders
  GetCommonHeaders() map[string]*string 
  SetXAcsBtripSoCorpToken(v string) *ElectronicItineraryGetApplyResultHeaders
  GetXAcsBtripSoCorpToken() *string 
}

type ElectronicItineraryGetApplyResultHeaders struct {
  CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
  // example:
  // 
  // feth00jqwls
  XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s ElectronicItineraryGetApplyResultHeaders) String() string {
  return dara.Prettify(s)
}

func (s ElectronicItineraryGetApplyResultHeaders) GoString() string {
  return s.String()
}

func (s *ElectronicItineraryGetApplyResultHeaders) GetCommonHeaders() map[string]*string  {
  return s.CommonHeaders
}

func (s *ElectronicItineraryGetApplyResultHeaders) GetXAcsBtripSoCorpToken() *string  {
  return s.XAcsBtripSoCorpToken
}

func (s *ElectronicItineraryGetApplyResultHeaders) SetCommonHeaders(v map[string]*string) *ElectronicItineraryGetApplyResultHeaders {
  s.CommonHeaders = v
  return s
}

func (s *ElectronicItineraryGetApplyResultHeaders) SetXAcsBtripSoCorpToken(v string) *ElectronicItineraryGetApplyResultHeaders {
  s.XAcsBtripSoCorpToken = &v
  return s
}

func (s *ElectronicItineraryGetApplyResultHeaders) Validate() error {
  return dara.Validate(s)
}

