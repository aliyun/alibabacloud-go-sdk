// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCostCenterDeleteHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CostCenterDeleteHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *CostCenterDeleteHeaders
	GetXAcsBtripSoCorpToken() *string
}

type CostCenterDeleteHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s CostCenterDeleteHeaders) String() string {
	return dara.Prettify(s)
}

func (s CostCenterDeleteHeaders) GoString() string {
	return s.String()
}

func (s *CostCenterDeleteHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CostCenterDeleteHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *CostCenterDeleteHeaders) SetCommonHeaders(v map[string]*string) *CostCenterDeleteHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CostCenterDeleteHeaders) SetXAcsBtripSoCorpToken(v string) *CostCenterDeleteHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *CostCenterDeleteHeaders) Validate() error {
	return dara.Validate(s)
}
