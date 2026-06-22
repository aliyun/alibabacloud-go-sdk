// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantSwitchAgreementRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsAgree(v bool) *GrantSwitchAgreementRequest
	GetIsAgree() *bool
	SetIsConfirmed(v bool) *GrantSwitchAgreementRequest
	GetIsConfirmed() *bool
	SetIsImmediate(v bool) *GrantSwitchAgreementRequest
	GetIsImmediate() *bool
	SetLang(v string) *GrantSwitchAgreementRequest
	GetLang() *string
	SetType(v string) *GrantSwitchAgreementRequest
	GetType() *string
}

type GrantSwitchAgreementRequest struct {
	// Specifies whether you agree to the data migration from the Hong Kong (China) region to the Singapore data center. Valid values:
	//
	// - **true**: Agree.
	//
	// - **false**: Disagree.
	//
	// example:
	//
	// true
	IsAgree *bool `json:"IsAgree,omitempty" xml:"IsAgree,omitempty"`
	// Specifies whether the user confirms that the data migration from the Hong Kong (China) region to the Singapore data center has been completed.
	//
	// - **true**: Confirmed. The user has confirmed that the data migration from the Hong Kong (China) region to the Singapore data center has been completed, and the notification pop-up window no longer needs to be displayed.
	//
	// - **false**: Not confirmed. The user has not confirmed that the data migration from the Hong Kong (China) region to the Singapore data center has been completed, and the notification pop-up window still needs to be displayed.
	IsConfirmed *bool `json:"IsConfirmed,omitempty" xml:"IsConfirmed,omitempty"`
	// Specifies whether to schedule data migration of data from the Hong Kong (China) region to the Singapore data center within 24 hours. Valid values:
	//
	// - **true**: Schedule the switch within 24 hours.
	//
	// - **false**: Do not schedule. For users who have cloud services in the Hong Kong (China) region, data migration will be automatically completed on March 5, 2026. For users who do not have cloud services in the Hong Kong (China) region, data migration will be automatically completed on November 17, 2025.
	//
	// example:
	//
	// true
	IsImmediate *bool `json:"IsImmediate,omitempty" xml:"IsImmediate,omitempty"`
	// The language type for the request and response messages. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The switch type. Valid values:
	//
	// - **sg_switch**: data migration from the Hong Kong (China) region to the Singapore data center.
	//
	// example:
	//
	// sg_switch
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GrantSwitchAgreementRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantSwitchAgreementRequest) GoString() string {
	return s.String()
}

func (s *GrantSwitchAgreementRequest) GetIsAgree() *bool {
	return s.IsAgree
}

func (s *GrantSwitchAgreementRequest) GetIsConfirmed() *bool {
	return s.IsConfirmed
}

func (s *GrantSwitchAgreementRequest) GetIsImmediate() *bool {
	return s.IsImmediate
}

func (s *GrantSwitchAgreementRequest) GetLang() *string {
	return s.Lang
}

func (s *GrantSwitchAgreementRequest) GetType() *string {
	return s.Type
}

func (s *GrantSwitchAgreementRequest) SetIsAgree(v bool) *GrantSwitchAgreementRequest {
	s.IsAgree = &v
	return s
}

func (s *GrantSwitchAgreementRequest) SetIsConfirmed(v bool) *GrantSwitchAgreementRequest {
	s.IsConfirmed = &v
	return s
}

func (s *GrantSwitchAgreementRequest) SetIsImmediate(v bool) *GrantSwitchAgreementRequest {
	s.IsImmediate = &v
	return s
}

func (s *GrantSwitchAgreementRequest) SetLang(v string) *GrantSwitchAgreementRequest {
	s.Lang = &v
	return s
}

func (s *GrantSwitchAgreementRequest) SetType(v string) *GrantSwitchAgreementRequest {
	s.Type = &v
	return s
}

func (s *GrantSwitchAgreementRequest) Validate() error {
	return dara.Validate(s)
}
