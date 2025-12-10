// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPromoteResourceAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *PromoteResourceAccountRequest
	GetAccountId() *string
	SetEmail(v string) *PromoteResourceAccountRequest
	GetEmail() *string
}

type PromoteResourceAccountRequest struct {
	// The ID of the resource account you want to upgrade.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The email address used to log on to the cloud account after the upgrade.
	//
	// This parameter is required.
	//
	// example:
	//
	// someone@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
}

func (s PromoteResourceAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s PromoteResourceAccountRequest) GoString() string {
	return s.String()
}

func (s *PromoteResourceAccountRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *PromoteResourceAccountRequest) GetEmail() *string {
	return s.Email
}

func (s *PromoteResourceAccountRequest) SetAccountId(v string) *PromoteResourceAccountRequest {
	s.AccountId = &v
	return s
}

func (s *PromoteResourceAccountRequest) SetEmail(v string) *PromoteResourceAccountRequest {
	s.Email = &v
	return s
}

func (s *PromoteResourceAccountRequest) Validate() error {
	return dara.Validate(s)
}
