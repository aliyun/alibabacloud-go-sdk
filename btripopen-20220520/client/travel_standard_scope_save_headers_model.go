// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTravelStandardScopeSaveHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *TravelStandardScopeSaveHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *TravelStandardScopeSaveHeaders
	GetXAcsBtripCorpToken() *string
}

type TravelStandardScopeSaveHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s TravelStandardScopeSaveHeaders) String() string {
	return dara.Prettify(s)
}

func (s TravelStandardScopeSaveHeaders) GoString() string {
	return s.String()
}

func (s *TravelStandardScopeSaveHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *TravelStandardScopeSaveHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *TravelStandardScopeSaveHeaders) SetCommonHeaders(v map[string]*string) *TravelStandardScopeSaveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TravelStandardScopeSaveHeaders) SetXAcsBtripCorpToken(v string) *TravelStandardScopeSaveHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *TravelStandardScopeSaveHeaders) Validate() error {
	return dara.Validate(s)
}
