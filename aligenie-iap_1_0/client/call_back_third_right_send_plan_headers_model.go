// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCallBackThirdRightSendPlanHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CallBackThirdRightSendPlanHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *CallBackThirdRightSendPlanHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *CallBackThirdRightSendPlanHeaders
	GetAuthorization() *string
}

type CallBackThirdRightSendPlanHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s CallBackThirdRightSendPlanHeaders) String() string {
	return dara.Prettify(s)
}

func (s CallBackThirdRightSendPlanHeaders) GoString() string {
	return s.String()
}

func (s *CallBackThirdRightSendPlanHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CallBackThirdRightSendPlanHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *CallBackThirdRightSendPlanHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *CallBackThirdRightSendPlanHeaders) SetCommonHeaders(v map[string]*string) *CallBackThirdRightSendPlanHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CallBackThirdRightSendPlanHeaders) SetXAcsAligenieAccessToken(v string) *CallBackThirdRightSendPlanHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *CallBackThirdRightSendPlanHeaders) SetAuthorization(v string) *CallBackThirdRightSendPlanHeaders {
	s.Authorization = &v
	return s
}

func (s *CallBackThirdRightSendPlanHeaders) Validate() error {
	return dara.Validate(s)
}
