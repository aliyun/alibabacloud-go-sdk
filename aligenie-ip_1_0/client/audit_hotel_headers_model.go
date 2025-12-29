// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuditHotelHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AuditHotelHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *AuditHotelHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *AuditHotelHeaders
	GetAuthorization() *string
}

type AuditHotelHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s AuditHotelHeaders) String() string {
	return dara.Prettify(s)
}

func (s AuditHotelHeaders) GoString() string {
	return s.String()
}

func (s *AuditHotelHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AuditHotelHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *AuditHotelHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *AuditHotelHeaders) SetCommonHeaders(v map[string]*string) *AuditHotelHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AuditHotelHeaders) SetXAcsAligenieAccessToken(v string) *AuditHotelHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *AuditHotelHeaders) SetAuthorization(v string) *AuditHotelHeaders {
	s.Authorization = &v
	return s
}

func (s *AuditHotelHeaders) Validate() error {
	return dara.Validate(s)
}
