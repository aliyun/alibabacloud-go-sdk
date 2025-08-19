// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelChangeAccountEmailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *CancelChangeAccountEmailRequest
	GetAccountId() *string
}

type CancelChangeAccountEmailRequest struct {
	// The Alibaba Cloud account ID of the member.
	//
	// This parameter is required.
	//
	// example:
	//
	// 181761095690****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s CancelChangeAccountEmailRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelChangeAccountEmailRequest) GoString() string {
	return s.String()
}

func (s *CancelChangeAccountEmailRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *CancelChangeAccountEmailRequest) SetAccountId(v string) *CancelChangeAccountEmailRequest {
	s.AccountId = &v
	return s
}

func (s *CancelChangeAccountEmailRequest) Validate() error {
	return dara.Validate(s)
}
