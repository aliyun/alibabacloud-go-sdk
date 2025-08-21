// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduleTaskHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetScheduleTaskHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *GetScheduleTaskHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *GetScheduleTaskHeaders
	GetAuthorization() *string
}

type GetScheduleTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetScheduleTaskHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetScheduleTaskHeaders) GoString() string {
	return s.String()
}

func (s *GetScheduleTaskHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetScheduleTaskHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *GetScheduleTaskHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetScheduleTaskHeaders) SetCommonHeaders(v map[string]*string) *GetScheduleTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetScheduleTaskHeaders) SetXAcsAligenieAccessToken(v string) *GetScheduleTaskHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetScheduleTaskHeaders) SetAuthorization(v string) *GetScheduleTaskHeaders {
	s.Authorization = &v
	return s
}

func (s *GetScheduleTaskHeaders) Validate() error {
	return dara.Validate(s)
}
