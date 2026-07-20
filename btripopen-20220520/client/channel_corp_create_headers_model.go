// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChannelCorpCreateHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ChannelCorpCreateHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *ChannelCorpCreateHeaders
	GetXAcsBtripCorpToken() *string
}

type ChannelCorpCreateHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// aqfrefd2321
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s ChannelCorpCreateHeaders) String() string {
	return dara.Prettify(s)
}

func (s ChannelCorpCreateHeaders) GoString() string {
	return s.String()
}

func (s *ChannelCorpCreateHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ChannelCorpCreateHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *ChannelCorpCreateHeaders) SetCommonHeaders(v map[string]*string) *ChannelCorpCreateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ChannelCorpCreateHeaders) SetXAcsBtripCorpToken(v string) *ChannelCorpCreateHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *ChannelCorpCreateHeaders) Validate() error {
	return dara.Validate(s)
}
