// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckThirdRightSendPlanHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CheckThirdRightSendPlanHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *CheckThirdRightSendPlanHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *CheckThirdRightSendPlanHeaders
	GetAuthorization() *string
}

type CheckThirdRightSendPlanHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s CheckThirdRightSendPlanHeaders) String() string {
	return dara.Prettify(s)
}

func (s CheckThirdRightSendPlanHeaders) GoString() string {
	return s.String()
}

func (s *CheckThirdRightSendPlanHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CheckThirdRightSendPlanHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *CheckThirdRightSendPlanHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *CheckThirdRightSendPlanHeaders) SetCommonHeaders(v map[string]*string) *CheckThirdRightSendPlanHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CheckThirdRightSendPlanHeaders) SetXAcsAligenieAccessToken(v string) *CheckThirdRightSendPlanHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *CheckThirdRightSendPlanHeaders) SetAuthorization(v string) *CheckThirdRightSendPlanHeaders {
	s.Authorization = &v
	return s
}

func (s *CheckThirdRightSendPlanHeaders) Validate() error {
	return dara.Validate(s)
}
