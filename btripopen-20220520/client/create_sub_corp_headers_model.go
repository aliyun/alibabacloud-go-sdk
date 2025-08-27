// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSubCorpHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateSubCorpHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *CreateSubCorpHeaders
	GetXAcsBtripCorpToken() *string
}

type CreateSubCorpHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s CreateSubCorpHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateSubCorpHeaders) GoString() string {
	return s.String()
}

func (s *CreateSubCorpHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateSubCorpHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *CreateSubCorpHeaders) SetCommonHeaders(v map[string]*string) *CreateSubCorpHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateSubCorpHeaders) SetXAcsBtripCorpToken(v string) *CreateSubCorpHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *CreateSubCorpHeaders) Validate() error {
	return dara.Validate(s)
}
