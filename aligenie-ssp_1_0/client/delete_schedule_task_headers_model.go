// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScheduleTaskHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteScheduleTaskHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *DeleteScheduleTaskHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *DeleteScheduleTaskHeaders
	GetAuthorization() *string
}

type DeleteScheduleTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s DeleteScheduleTaskHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteScheduleTaskHeaders) GoString() string {
	return s.String()
}

func (s *DeleteScheduleTaskHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteScheduleTaskHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *DeleteScheduleTaskHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *DeleteScheduleTaskHeaders) SetCommonHeaders(v map[string]*string) *DeleteScheduleTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteScheduleTaskHeaders) SetXAcsAligenieAccessToken(v string) *DeleteScheduleTaskHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *DeleteScheduleTaskHeaders) SetAuthorization(v string) *DeleteScheduleTaskHeaders {
	s.Authorization = &v
	return s
}

func (s *DeleteScheduleTaskHeaders) Validate() error {
	return dara.Validate(s)
}
