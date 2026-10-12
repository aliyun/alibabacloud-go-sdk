// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRiskNotificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsMute(v bool) *UpdateRiskNotificationRequest
	GetIsMute() *bool
	SetRiskCode(v string) *UpdateRiskNotificationRequest
	GetRiskCode() *string
}

type UpdateRiskNotificationRequest struct {
	// Specifies whether to mute notifications.
	//
	// example:
	//
	// true
	IsMute *bool `json:"isMute,omitempty" xml:"isMute,omitempty"`
	// The risk item code. This parameter is required.
	//
	// The value is a numeric string (for example, 60010010002 or 50010010009). Call GetRiskNotification to query the code of the current risk item.
	//
	// example:
	//
	// 60010010002
	RiskCode *string `json:"riskCode,omitempty" xml:"riskCode,omitempty"`
}

func (s UpdateRiskNotificationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRiskNotificationRequest) GoString() string {
	return s.String()
}

func (s *UpdateRiskNotificationRequest) GetIsMute() *bool {
	return s.IsMute
}

func (s *UpdateRiskNotificationRequest) GetRiskCode() *string {
	return s.RiskCode
}

func (s *UpdateRiskNotificationRequest) SetIsMute(v bool) *UpdateRiskNotificationRequest {
	s.IsMute = &v
	return s
}

func (s *UpdateRiskNotificationRequest) SetRiskCode(v string) *UpdateRiskNotificationRequest {
	s.RiskCode = &v
	return s
}

func (s *UpdateRiskNotificationRequest) Validate() error {
	return dara.Validate(s)
}
