// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomRoleHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateCustomRoleHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *CreateCustomRoleHeaders
	GetXAcsBtripCorpToken() *string
}

type CreateCustomRoleHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s CreateCustomRoleHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomRoleHeaders) GoString() string {
	return s.String()
}

func (s *CreateCustomRoleHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateCustomRoleHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *CreateCustomRoleHeaders) SetCommonHeaders(v map[string]*string) *CreateCustomRoleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateCustomRoleHeaders) SetXAcsBtripCorpToken(v string) *CreateCustomRoleHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *CreateCustomRoleHeaders) Validate() error {
	return dara.Validate(s)
}
