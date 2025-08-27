// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCostCenterQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CostCenterQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *CostCenterQueryHeaders
	GetXAcsBtripSoCorpToken() *string
}

type CostCenterQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s CostCenterQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s CostCenterQueryHeaders) GoString() string {
	return s.String()
}

func (s *CostCenterQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CostCenterQueryHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *CostCenterQueryHeaders) SetCommonHeaders(v map[string]*string) *CostCenterQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CostCenterQueryHeaders) SetXAcsBtripSoCorpToken(v string) *CostCenterQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *CostCenterQueryHeaders) Validate() error {
	return dara.Validate(s)
}
