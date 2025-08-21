// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIndexControlPlayingListHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *IndexControlPlayingListHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *IndexControlPlayingListHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *IndexControlPlayingListHeaders
	GetAuthorization() *string
}

type IndexControlPlayingListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s IndexControlPlayingListHeaders) String() string {
	return dara.Prettify(s)
}

func (s IndexControlPlayingListHeaders) GoString() string {
	return s.String()
}

func (s *IndexControlPlayingListHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *IndexControlPlayingListHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *IndexControlPlayingListHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *IndexControlPlayingListHeaders) SetCommonHeaders(v map[string]*string) *IndexControlPlayingListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IndexControlPlayingListHeaders) SetXAcsAligenieAccessToken(v string) *IndexControlPlayingListHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *IndexControlPlayingListHeaders) SetAuthorization(v string) *IndexControlPlayingListHeaders {
	s.Authorization = &v
	return s
}

func (s *IndexControlPlayingListHeaders) Validate() error {
	return dara.Validate(s)
}
