// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlertDestinationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *CreateAlertDestinationRequest
	GetXDebugId() *string
	SetAppId(v string) *CreateAlertDestinationRequest
	GetAppId() *string
	SetAppSecret(v string) *CreateAlertDestinationRequest
	GetAppSecret() *string
	SetGroupId(v []*string) *CreateAlertDestinationRequest
	GetGroupId() []*string
	SetImbot(v bool) *CreateAlertDestinationRequest
	GetImbot() *bool
	SetName(v string) *CreateAlertDestinationRequest
	GetName() *string
	SetParams(v *CreateAlertDestinationRequestParams) *CreateAlertDestinationRequest
	GetParams() *CreateAlertDestinationRequestParams
	SetSource(v string) *CreateAlertDestinationRequest
	GetSource() *string
	SetTarget(v string) *CreateAlertDestinationRequest
	GetTarget() *string
	SetXSysomInvokeSource(v string) *CreateAlertDestinationRequest
	GetXSysomInvokeSource() *string
}

type CreateAlertDestinationRequest struct {
	XDebugId  *string   `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	AppId     *string   `json:"app_id,omitempty" xml:"app_id,omitempty"`
	AppSecret *string   `json:"app_secret,omitempty" xml:"app_secret,omitempty"`
	GroupId   []*string `json:"group_id,omitempty" xml:"group_id,omitempty" type:"Repeated"`
	Imbot     *bool     `json:"imbot,omitempty" xml:"imbot,omitempty"`
	// The name of the alert contact.
	//
	// example:
	//
	// test_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The configuration parameters of the alert contact.
	Params *CreateAlertDestinationRequestParams `json:"params,omitempty" xml:"params,omitempty" type:"Struct"`
	// The configuration source.
	//
	// example:
	//
	// console
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The notification target. Currently, only DingTalk chatbots are supported.
	//
	// example:
	//
	// dingtalk
	Target             *string `json:"target,omitempty" xml:"target,omitempty"`
	XSysomInvokeSource *string `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s CreateAlertDestinationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertDestinationRequest) GoString() string {
	return s.String()
}

func (s *CreateAlertDestinationRequest) GetXDebugId() *string {
	return s.XDebugId
}

func (s *CreateAlertDestinationRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateAlertDestinationRequest) GetAppSecret() *string {
	return s.AppSecret
}

func (s *CreateAlertDestinationRequest) GetGroupId() []*string {
	return s.GroupId
}

func (s *CreateAlertDestinationRequest) GetImbot() *bool {
	return s.Imbot
}

func (s *CreateAlertDestinationRequest) GetName() *string {
	return s.Name
}

func (s *CreateAlertDestinationRequest) GetParams() *CreateAlertDestinationRequestParams {
	return s.Params
}

func (s *CreateAlertDestinationRequest) GetSource() *string {
	return s.Source
}

func (s *CreateAlertDestinationRequest) GetTarget() *string {
	return s.Target
}

func (s *CreateAlertDestinationRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *CreateAlertDestinationRequest) SetXDebugId(v string) *CreateAlertDestinationRequest {
	s.XDebugId = &v
	return s
}

func (s *CreateAlertDestinationRequest) SetAppId(v string) *CreateAlertDestinationRequest {
	s.AppId = &v
	return s
}

func (s *CreateAlertDestinationRequest) SetAppSecret(v string) *CreateAlertDestinationRequest {
	s.AppSecret = &v
	return s
}

func (s *CreateAlertDestinationRequest) SetGroupId(v []*string) *CreateAlertDestinationRequest {
	s.GroupId = v
	return s
}

func (s *CreateAlertDestinationRequest) SetImbot(v bool) *CreateAlertDestinationRequest {
	s.Imbot = &v
	return s
}

func (s *CreateAlertDestinationRequest) SetName(v string) *CreateAlertDestinationRequest {
	s.Name = &v
	return s
}

func (s *CreateAlertDestinationRequest) SetParams(v *CreateAlertDestinationRequestParams) *CreateAlertDestinationRequest {
	s.Params = v
	return s
}

func (s *CreateAlertDestinationRequest) SetSource(v string) *CreateAlertDestinationRequest {
	s.Source = &v
	return s
}

func (s *CreateAlertDestinationRequest) SetTarget(v string) *CreateAlertDestinationRequest {
	s.Target = &v
	return s
}

func (s *CreateAlertDestinationRequest) SetXSysomInvokeSource(v string) *CreateAlertDestinationRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *CreateAlertDestinationRequest) Validate() error {
	if s.Params != nil {
		if err := s.Params.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAlertDestinationRequestParams struct {
	// The email address.
	//
	// example:
	//
	// xxx@email.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// The phone number.
	//
	// example:
	//
	// 1xxx
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// The secret key of the chatbot.
	//
	// example:
	//
	// SECxxx
	Sec *string `json:"sec,omitempty" xml:"sec,omitempty"`
	// The webhook URL of the chatbot.
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=xxx
	Webhook *string `json:"webhook,omitempty" xml:"webhook,omitempty"`
}

func (s CreateAlertDestinationRequestParams) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertDestinationRequestParams) GoString() string {
	return s.String()
}

func (s *CreateAlertDestinationRequestParams) GetEmail() *string {
	return s.Email
}

func (s *CreateAlertDestinationRequestParams) GetPhone() *string {
	return s.Phone
}

func (s *CreateAlertDestinationRequestParams) GetSec() *string {
	return s.Sec
}

func (s *CreateAlertDestinationRequestParams) GetWebhook() *string {
	return s.Webhook
}

func (s *CreateAlertDestinationRequestParams) SetEmail(v string) *CreateAlertDestinationRequestParams {
	s.Email = &v
	return s
}

func (s *CreateAlertDestinationRequestParams) SetPhone(v string) *CreateAlertDestinationRequestParams {
	s.Phone = &v
	return s
}

func (s *CreateAlertDestinationRequestParams) SetSec(v string) *CreateAlertDestinationRequestParams {
	s.Sec = &v
	return s
}

func (s *CreateAlertDestinationRequestParams) SetWebhook(v string) *CreateAlertDestinationRequestParams {
	s.Webhook = &v
	return s
}

func (s *CreateAlertDestinationRequestParams) Validate() error {
	return dara.Validate(s)
}
