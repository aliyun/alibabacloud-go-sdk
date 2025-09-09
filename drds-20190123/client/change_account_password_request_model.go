// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeAccountPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *ChangeAccountPasswordRequest
	GetAccountName() *string
	SetDrdsInstanceId(v string) *ChangeAccountPasswordRequest
	GetDrdsInstanceId() *string
	SetPassword(v string) *ChangeAccountPasswordRequest
	GetPassword() *string
}

type ChangeAccountPasswordRequest struct {
	// The name of the member account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The ID of the DRDS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds********
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The new password.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
}

func (s ChangeAccountPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeAccountPasswordRequest) GoString() string {
	return s.String()
}

func (s *ChangeAccountPasswordRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *ChangeAccountPasswordRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *ChangeAccountPasswordRequest) GetPassword() *string {
	return s.Password
}

func (s *ChangeAccountPasswordRequest) SetAccountName(v string) *ChangeAccountPasswordRequest {
	s.AccountName = &v
	return s
}

func (s *ChangeAccountPasswordRequest) SetDrdsInstanceId(v string) *ChangeAccountPasswordRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *ChangeAccountPasswordRequest) SetPassword(v string) *ChangeAccountPasswordRequest {
	s.Password = &v
	return s
}

func (s *ChangeAccountPasswordRequest) Validate() error {
	return dara.Validate(s)
}
