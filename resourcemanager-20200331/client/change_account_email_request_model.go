// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeAccountEmailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *ChangeAccountEmailRequest
	GetAccountId() *string
	SetEmail(v string) *ChangeAccountEmailRequest
	GetEmail() *string
}

type ChangeAccountEmailRequest struct {
	// The member account ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 181761095690****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The email address of the member account.
	//
	// > The system automatically sends a verification email to the specified email address. The modification takes effect only after the verification is complete. After the verification is complete, the system synchronously updates the email address and security email address of the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// someone@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
}

func (s ChangeAccountEmailRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeAccountEmailRequest) GoString() string {
	return s.String()
}

func (s *ChangeAccountEmailRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *ChangeAccountEmailRequest) GetEmail() *string {
	return s.Email
}

func (s *ChangeAccountEmailRequest) SetAccountId(v string) *ChangeAccountEmailRequest {
	s.AccountId = &v
	return s
}

func (s *ChangeAccountEmailRequest) SetEmail(v string) *ChangeAccountEmailRequest {
	s.Email = &v
	return s
}

func (s *ChangeAccountEmailRequest) Validate() error {
	return dara.Validate(s)
}
