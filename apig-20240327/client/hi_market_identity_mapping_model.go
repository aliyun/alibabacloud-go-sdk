// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHiMarketIdentityMapping interface {
	dara.Model
	String() string
	GoString() string
	SetEmailField(v string) *HiMarketIdentityMapping
	GetEmailField() *string
	SetUserIdField(v string) *HiMarketIdentityMapping
	GetUserIdField() *string
	SetUserNameField(v string) *HiMarketIdentityMapping
	GetUserNameField() *string
}

type HiMarketIdentityMapping struct {
	EmailField    *string `json:"emailField,omitempty" xml:"emailField,omitempty"`
	UserIdField   *string `json:"userIdField,omitempty" xml:"userIdField,omitempty"`
	UserNameField *string `json:"userNameField,omitempty" xml:"userNameField,omitempty"`
}

func (s HiMarketIdentityMapping) String() string {
	return dara.Prettify(s)
}

func (s HiMarketIdentityMapping) GoString() string {
	return s.String()
}

func (s *HiMarketIdentityMapping) GetEmailField() *string {
	return s.EmailField
}

func (s *HiMarketIdentityMapping) GetUserIdField() *string {
	return s.UserIdField
}

func (s *HiMarketIdentityMapping) GetUserNameField() *string {
	return s.UserNameField
}

func (s *HiMarketIdentityMapping) SetEmailField(v string) *HiMarketIdentityMapping {
	s.EmailField = &v
	return s
}

func (s *HiMarketIdentityMapping) SetUserIdField(v string) *HiMarketIdentityMapping {
	s.UserIdField = &v
	return s
}

func (s *HiMarketIdentityMapping) SetUserNameField(v string) *HiMarketIdentityMapping {
	s.UserNameField = &v
	return s
}

func (s *HiMarketIdentityMapping) Validate() error {
	return dara.Validate(s)
}
