// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportRoomControlDevicesHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ImportRoomControlDevicesHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ImportRoomControlDevicesHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ImportRoomControlDevicesHeaders
	GetAuthorization() *string
}

type ImportRoomControlDevicesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ImportRoomControlDevicesHeaders) String() string {
	return dara.Prettify(s)
}

func (s ImportRoomControlDevicesHeaders) GoString() string {
	return s.String()
}

func (s *ImportRoomControlDevicesHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ImportRoomControlDevicesHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ImportRoomControlDevicesHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ImportRoomControlDevicesHeaders) SetCommonHeaders(v map[string]*string) *ImportRoomControlDevicesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ImportRoomControlDevicesHeaders) SetXAcsAligenieAccessToken(v string) *ImportRoomControlDevicesHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ImportRoomControlDevicesHeaders) SetAuthorization(v string) *ImportRoomControlDevicesHeaders {
	s.Authorization = &v
	return s
}

func (s *ImportRoomControlDevicesHeaders) Validate() error {
	return dara.Validate(s)
}
