// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddProjectManagerHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddProjectManagerHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *AddProjectManagerHeaders
	GetXAcsBtripCorpToken() *string
}

type AddProjectManagerHeaders struct {
	CommonHeaders      map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripCorpToken *string            `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s AddProjectManagerHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddProjectManagerHeaders) GoString() string {
	return s.String()
}

func (s *AddProjectManagerHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddProjectManagerHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *AddProjectManagerHeaders) SetCommonHeaders(v map[string]*string) *AddProjectManagerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddProjectManagerHeaders) SetXAcsBtripCorpToken(v string) *AddProjectManagerHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *AddProjectManagerHeaders) Validate() error {
	return dara.Validate(s)
}
