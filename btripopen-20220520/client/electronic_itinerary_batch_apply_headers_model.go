// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iElectronicItineraryBatchApplyHeaders interface {
  dara.Model
  String() string
  GoString() string
  SetCommonHeaders(v map[string]*string) *ElectronicItineraryBatchApplyHeaders
  GetCommonHeaders() map[string]*string 
  SetXAcsBtripSoCorpToken(v string) *ElectronicItineraryBatchApplyHeaders
  GetXAcsBtripSoCorpToken() *string 
}

type ElectronicItineraryBatchApplyHeaders struct {
  CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
  // example:
  // 
  // feth00jqwls
  XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s ElectronicItineraryBatchApplyHeaders) String() string {
  return dara.Prettify(s)
}

func (s ElectronicItineraryBatchApplyHeaders) GoString() string {
  return s.String()
}

func (s *ElectronicItineraryBatchApplyHeaders) GetCommonHeaders() map[string]*string  {
  return s.CommonHeaders
}

func (s *ElectronicItineraryBatchApplyHeaders) GetXAcsBtripSoCorpToken() *string  {
  return s.XAcsBtripSoCorpToken
}

func (s *ElectronicItineraryBatchApplyHeaders) SetCommonHeaders(v map[string]*string) *ElectronicItineraryBatchApplyHeaders {
  s.CommonHeaders = v
  return s
}

func (s *ElectronicItineraryBatchApplyHeaders) SetXAcsBtripSoCorpToken(v string) *ElectronicItineraryBatchApplyHeaders {
  s.XAcsBtripSoCorpToken = &v
  return s
}

func (s *ElectronicItineraryBatchApplyHeaders) Validate() error {
  return dara.Validate(s)
}

