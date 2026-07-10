// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveProjectManagerHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *RemoveProjectManagerHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *RemoveProjectManagerHeaders
	GetXAcsBtripCorpToken() *string
}

type RemoveProjectManagerHeaders struct {
	CommonHeaders      map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripCorpToken *string            `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s RemoveProjectManagerHeaders) String() string {
	return dara.Prettify(s)
}

func (s RemoveProjectManagerHeaders) GoString() string {
	return s.String()
}

func (s *RemoveProjectManagerHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *RemoveProjectManagerHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *RemoveProjectManagerHeaders) SetCommonHeaders(v map[string]*string) *RemoveProjectManagerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RemoveProjectManagerHeaders) SetXAcsBtripCorpToken(v string) *RemoveProjectManagerHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *RemoveProjectManagerHeaders) Validate() error {
	return dara.Validate(s)
}
