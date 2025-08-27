// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCostCenterModifyHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CostCenterModifyHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *CostCenterModifyHeaders
	GetXAcsBtripSoCorpToken() *string
}

type CostCenterModifyHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s CostCenterModifyHeaders) String() string {
	return dara.Prettify(s)
}

func (s CostCenterModifyHeaders) GoString() string {
	return s.String()
}

func (s *CostCenterModifyHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CostCenterModifyHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *CostCenterModifyHeaders) SetCommonHeaders(v map[string]*string) *CostCenterModifyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CostCenterModifyHeaders) SetXAcsBtripSoCorpToken(v string) *CostCenterModifyHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *CostCenterModifyHeaders) Validate() error {
	return dara.Validate(s)
}
