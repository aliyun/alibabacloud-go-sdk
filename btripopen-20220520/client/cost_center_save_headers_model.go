// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCostCenterSaveHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CostCenterSaveHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *CostCenterSaveHeaders
	GetXAcsBtripSoCorpToken() *string
}

type CostCenterSaveHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s CostCenterSaveHeaders) String() string {
	return dara.Prettify(s)
}

func (s CostCenterSaveHeaders) GoString() string {
	return s.String()
}

func (s *CostCenterSaveHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CostCenterSaveHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *CostCenterSaveHeaders) SetCommonHeaders(v map[string]*string) *CostCenterSaveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CostCenterSaveHeaders) SetXAcsBtripSoCorpToken(v string) *CostCenterSaveHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *CostCenterSaveHeaders) Validate() error {
	return dara.Validate(s)
}
