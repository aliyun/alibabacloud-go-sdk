// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduleTaskHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateScheduleTaskHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *CreateScheduleTaskHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *CreateScheduleTaskHeaders
	GetAuthorization() *string
}

type CreateScheduleTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s CreateScheduleTaskHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleTaskHeaders) GoString() string {
	return s.String()
}

func (s *CreateScheduleTaskHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateScheduleTaskHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *CreateScheduleTaskHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *CreateScheduleTaskHeaders) SetCommonHeaders(v map[string]*string) *CreateScheduleTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateScheduleTaskHeaders) SetXAcsAligenieAccessToken(v string) *CreateScheduleTaskHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *CreateScheduleTaskHeaders) SetAuthorization(v string) *CreateScheduleTaskHeaders {
	s.Authorization = &v
	return s
}

func (s *CreateScheduleTaskHeaders) Validate() error {
	return dara.Validate(s)
}
