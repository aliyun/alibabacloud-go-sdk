// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCarSceneQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CarSceneQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *CarSceneQueryHeaders
	GetXAcsBtripCorpToken() *string
}

type CarSceneQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwis
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s CarSceneQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s CarSceneQueryHeaders) GoString() string {
	return s.String()
}

func (s *CarSceneQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CarSceneQueryHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *CarSceneQueryHeaders) SetCommonHeaders(v map[string]*string) *CarSceneQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CarSceneQueryHeaders) SetXAcsBtripCorpToken(v string) *CarSceneQueryHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *CarSceneQueryHeaders) Validate() error {
	return dara.Validate(s)
}
