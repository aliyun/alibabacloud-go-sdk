// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlertDestinationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateAlertDestinationRequest
	GetName() *string
	SetParams(v *CreateAlertDestinationRequestParams) *CreateAlertDestinationRequest
	GetParams() *CreateAlertDestinationRequestParams
	SetSource(v string) *CreateAlertDestinationRequest
	GetSource() *string
	SetTarget(v string) *CreateAlertDestinationRequest
	GetTarget() *string
}

type CreateAlertDestinationRequest struct {
	// example:
	//
	// test_name
	Name   *string                              `json:"name,omitempty" xml:"name,omitempty"`
	Params *CreateAlertDestinationRequestParams `json:"params,omitempty" xml:"params,omitempty" type:"Struct"`
	// example:
	//
	// console
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// dingtalk
	Target *string `json:"target,omitempty" xml:"target,omitempty"`
}

func (s CreateAlertDestinationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertDestinationRequest) GoString() string {
	return s.String()
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

func (s *CreateAlertDestinationRequest) Validate() error {
	if s.Params != nil {
		if err := s.Params.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAlertDestinationRequestParams struct {
	// example:
	//
	// xxx@email.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// 1xxx
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// example:
	//
	// SECxxx
	Sec *string `json:"sec,omitempty" xml:"sec,omitempty"`
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
