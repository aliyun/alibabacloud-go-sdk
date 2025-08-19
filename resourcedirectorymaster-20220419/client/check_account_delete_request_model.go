// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAccountDeleteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *CheckAccountDeleteRequest
	GetAccountId() *string
}

type CheckAccountDeleteRequest struct {
	// The Alibaba Cloud account ID of the member that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// 179855839641****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s CheckAccountDeleteRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckAccountDeleteRequest) GoString() string {
	return s.String()
}

func (s *CheckAccountDeleteRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *CheckAccountDeleteRequest) SetAccountId(v string) *CheckAccountDeleteRequest {
	s.AccountId = &v
	return s
}

func (s *CheckAccountDeleteRequest) Validate() error {
	return dara.Validate(s)
}
