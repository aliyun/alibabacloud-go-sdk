// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlertDestinationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *UpdateAlertDestinationRequest
	GetId() *string
	SetName(v string) *UpdateAlertDestinationRequest
	GetName() *string
	SetParams(v *UpdateAlertDestinationRequestParams) *UpdateAlertDestinationRequest
	GetParams() *UpdateAlertDestinationRequestParams
	SetSource(v string) *UpdateAlertDestinationRequest
	GetSource() *string
	SetTarget(v string) *UpdateAlertDestinationRequest
	GetTarget() *string
}

type UpdateAlertDestinationRequest struct {
	// example:
	//
	// 1
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// name1
	Name   *string                              `json:"name,omitempty" xml:"name,omitempty"`
	Params *UpdateAlertDestinationRequestParams `json:"params,omitempty" xml:"params,omitempty" type:"Struct"`
	// example:
	//
	// console
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// dingtalk
	Target *string `json:"target,omitempty" xml:"target,omitempty"`
}

func (s UpdateAlertDestinationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertDestinationRequest) GoString() string {
	return s.String()
}

func (s *UpdateAlertDestinationRequest) GetId() *string {
	return s.Id
}

func (s *UpdateAlertDestinationRequest) GetName() *string {
	return s.Name
}

func (s *UpdateAlertDestinationRequest) GetParams() *UpdateAlertDestinationRequestParams {
	return s.Params
}

func (s *UpdateAlertDestinationRequest) GetSource() *string {
	return s.Source
}

func (s *UpdateAlertDestinationRequest) GetTarget() *string {
	return s.Target
}

func (s *UpdateAlertDestinationRequest) SetId(v string) *UpdateAlertDestinationRequest {
	s.Id = &v
	return s
}

func (s *UpdateAlertDestinationRequest) SetName(v string) *UpdateAlertDestinationRequest {
	s.Name = &v
	return s
}

func (s *UpdateAlertDestinationRequest) SetParams(v *UpdateAlertDestinationRequestParams) *UpdateAlertDestinationRequest {
	s.Params = v
	return s
}

func (s *UpdateAlertDestinationRequest) SetSource(v string) *UpdateAlertDestinationRequest {
	s.Source = &v
	return s
}

func (s *UpdateAlertDestinationRequest) SetTarget(v string) *UpdateAlertDestinationRequest {
	s.Target = &v
	return s
}

func (s *UpdateAlertDestinationRequest) Validate() error {
	if s.Params != nil {
		if err := s.Params.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAlertDestinationRequestParams struct {
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

func (s UpdateAlertDestinationRequestParams) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertDestinationRequestParams) GoString() string {
	return s.String()
}

func (s *UpdateAlertDestinationRequestParams) GetEmail() *string {
	return s.Email
}

func (s *UpdateAlertDestinationRequestParams) GetPhone() *string {
	return s.Phone
}

func (s *UpdateAlertDestinationRequestParams) GetSec() *string {
	return s.Sec
}

func (s *UpdateAlertDestinationRequestParams) GetWebhook() *string {
	return s.Webhook
}

func (s *UpdateAlertDestinationRequestParams) SetEmail(v string) *UpdateAlertDestinationRequestParams {
	s.Email = &v
	return s
}

func (s *UpdateAlertDestinationRequestParams) SetPhone(v string) *UpdateAlertDestinationRequestParams {
	s.Phone = &v
	return s
}

func (s *UpdateAlertDestinationRequestParams) SetSec(v string) *UpdateAlertDestinationRequestParams {
	s.Sec = &v
	return s
}

func (s *UpdateAlertDestinationRequestParams) SetWebhook(v string) *UpdateAlertDestinationRequestParams {
	s.Webhook = &v
	return s
}

func (s *UpdateAlertDestinationRequestParams) Validate() error {
	return dara.Validate(s)
}
