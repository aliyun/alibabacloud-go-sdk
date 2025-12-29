// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncDeviceStatusWithAkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SyncDeviceStatusWithAkHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *SyncDeviceStatusWithAkHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *SyncDeviceStatusWithAkHeaders
	GetAuthorization() *string
}

type SyncDeviceStatusWithAkHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s SyncDeviceStatusWithAkHeaders) String() string {
	return dara.Prettify(s)
}

func (s SyncDeviceStatusWithAkHeaders) GoString() string {
	return s.String()
}

func (s *SyncDeviceStatusWithAkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SyncDeviceStatusWithAkHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *SyncDeviceStatusWithAkHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *SyncDeviceStatusWithAkHeaders) SetCommonHeaders(v map[string]*string) *SyncDeviceStatusWithAkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SyncDeviceStatusWithAkHeaders) SetXAcsAligenieAccessToken(v string) *SyncDeviceStatusWithAkHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *SyncDeviceStatusWithAkHeaders) SetAuthorization(v string) *SyncDeviceStatusWithAkHeaders {
	s.Authorization = &v
	return s
}

func (s *SyncDeviceStatusWithAkHeaders) Validate() error {
	return dara.Validate(s)
}
