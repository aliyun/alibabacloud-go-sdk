// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEmployeeDetailHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryEmployeeDetailHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *QueryEmployeeDetailHeaders
	GetXAcsBtripCorpToken() *string
}

type QueryEmployeeDetailHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwis
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s QueryEmployeeDetailHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryEmployeeDetailHeaders) GoString() string {
	return s.String()
}

func (s *QueryEmployeeDetailHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryEmployeeDetailHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *QueryEmployeeDetailHeaders) SetCommonHeaders(v map[string]*string) *QueryEmployeeDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryEmployeeDetailHeaders) SetXAcsBtripCorpToken(v string) *QueryEmployeeDetailHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *QueryEmployeeDetailHeaders) Validate() error {
	return dara.Validate(s)
}
