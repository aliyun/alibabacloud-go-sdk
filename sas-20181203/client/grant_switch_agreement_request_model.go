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
	// Indicates whether to agree to migrate the client connections from overseas servers to the Singapore center.
	//
	// example:
	//
	// true
	IsAgree *bool `json:"IsAgree,omitempty" xml:"IsAgree,omitempty"`
	// Has the user confirmed the migration of Hong Kong region data to Singapore data center
	//
	// - **true:*	- The user has confirmed that Hong Kong region data has been migrated to the Singapore data center. No notification popup needs to be displayed subsequently.
	//
	// - **false**:The user has not confirmed that Hong Kong region data has been migrated to the Singapore data center. Notification popup still needs to be displayed subsequently.
	//
	// example:
	//
	// true
	IsConfirmed *bool `json:"IsConfirmed,omitempty" xml:"IsConfirmed,omitempty"`
	// Whether to schedule the migration of data from the Hong Kong region to the Singapore data center within 24 hours. Values:
	//
	// - **true**: Schedule the switch within 24 hours.
	//
	// - **false**: Do not schedule. Users with cloud products in the Hong Kong region will be automatically migrated on March 5, 2026; users without cloud products in the Hong Kong region will be automatically migrated on November 17, 2025.
	//
	// example:
	//
	// true
	IsImmediate *bool `json:"IsImmediate,omitempty" xml:"IsImmediate,omitempty"`
	// The language type for requests and responses. The default value is **zh**. Possible values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Switching type. Possible values:
	//
	// - **sg_switch**: Migrate client connections from overseas servers to Singapore
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
