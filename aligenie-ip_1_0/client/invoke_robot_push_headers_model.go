// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeRobotPushHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InvokeRobotPushHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *InvokeRobotPushHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *InvokeRobotPushHeaders
	GetAuthorization() *string
}

type InvokeRobotPushHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s InvokeRobotPushHeaders) String() string {
	return dara.Prettify(s)
}

func (s InvokeRobotPushHeaders) GoString() string {
	return s.String()
}

func (s *InvokeRobotPushHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InvokeRobotPushHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *InvokeRobotPushHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *InvokeRobotPushHeaders) SetCommonHeaders(v map[string]*string) *InvokeRobotPushHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InvokeRobotPushHeaders) SetXAcsAligenieAccessToken(v string) *InvokeRobotPushHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *InvokeRobotPushHeaders) SetAuthorization(v string) *InvokeRobotPushHeaders {
	s.Authorization = &v
	return s
}

func (s *InvokeRobotPushHeaders) Validate() error {
	return dara.Validate(s)
}
