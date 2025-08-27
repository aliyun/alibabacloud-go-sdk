// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTripTaskQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *TripTaskQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *TripTaskQueryHeaders
	GetXAcsBtripCorpToken() *string
}

type TripTaskQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s TripTaskQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s TripTaskQueryHeaders) GoString() string {
	return s.String()
}

func (s *TripTaskQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *TripTaskQueryHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *TripTaskQueryHeaders) SetCommonHeaders(v map[string]*string) *TripTaskQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TripTaskQueryHeaders) SetXAcsBtripCorpToken(v string) *TripTaskQueryHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *TripTaskQueryHeaders) Validate() error {
	return dara.Validate(s)
}
