// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryChangeAccountEmailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *RetryChangeAccountEmailRequest
	GetAccountId() *string
}

type RetryChangeAccountEmailRequest struct {
	// The Alibaba Cloud account ID of the member.
	//
	// This parameter is required.
	//
	// example:
	//
	// 181761095690****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s RetryChangeAccountEmailRequest) String() string {
	return dara.Prettify(s)
}

func (s RetryChangeAccountEmailRequest) GoString() string {
	return s.String()
}

func (s *RetryChangeAccountEmailRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *RetryChangeAccountEmailRequest) SetAccountId(v string) *RetryChangeAccountEmailRequest {
	s.AccountId = &v
	return s
}

func (s *RetryChangeAccountEmailRequest) Validate() error {
	return dara.Validate(s)
}
