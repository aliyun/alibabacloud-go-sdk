// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportRoomGenieScenesHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ImportRoomGenieScenesHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ImportRoomGenieScenesHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ImportRoomGenieScenesHeaders
	GetAuthorization() *string
}

type ImportRoomGenieScenesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ImportRoomGenieScenesHeaders) String() string {
	return dara.Prettify(s)
}

func (s ImportRoomGenieScenesHeaders) GoString() string {
	return s.String()
}

func (s *ImportRoomGenieScenesHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ImportRoomGenieScenesHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ImportRoomGenieScenesHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ImportRoomGenieScenesHeaders) SetCommonHeaders(v map[string]*string) *ImportRoomGenieScenesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ImportRoomGenieScenesHeaders) SetXAcsAligenieAccessToken(v string) *ImportRoomGenieScenesHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ImportRoomGenieScenesHeaders) SetAuthorization(v string) *ImportRoomGenieScenesHeaders {
	s.Authorization = &v
	return s
}

func (s *ImportRoomGenieScenesHeaders) Validate() error {
	return dara.Validate(s)
}
