// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceIdByIdentitiesHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListDeviceIdByIdentitiesHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ListDeviceIdByIdentitiesHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ListDeviceIdByIdentitiesHeaders
	GetAuthorization() *string
}

type ListDeviceIdByIdentitiesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListDeviceIdByIdentitiesHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceIdByIdentitiesHeaders) GoString() string {
	return s.String()
}

func (s *ListDeviceIdByIdentitiesHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListDeviceIdByIdentitiesHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ListDeviceIdByIdentitiesHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListDeviceIdByIdentitiesHeaders) SetCommonHeaders(v map[string]*string) *ListDeviceIdByIdentitiesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListDeviceIdByIdentitiesHeaders) SetXAcsAligenieAccessToken(v string) *ListDeviceIdByIdentitiesHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListDeviceIdByIdentitiesHeaders) SetAuthorization(v string) *ListDeviceIdByIdentitiesHeaders {
	s.Authorization = &v
	return s
}

func (s *ListDeviceIdByIdentitiesHeaders) Validate() error {
	return dara.Validate(s)
}
