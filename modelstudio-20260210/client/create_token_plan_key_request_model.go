// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTokenPlanKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *CreateTokenPlanKeyRequest
	GetAccountId() *string
	SetDescription(v string) *CreateTokenPlanKeyRequest
	GetDescription() *string
}

type CreateTokenPlanKeyRequest struct {
	// The account ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// acc_123456789
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The description of the key.
	//
	// example:
	//
	// APIKEY描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s CreateTokenPlanKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTokenPlanKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateTokenPlanKeyRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *CreateTokenPlanKeyRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateTokenPlanKeyRequest) SetAccountId(v string) *CreateTokenPlanKeyRequest {
	s.AccountId = &v
	return s
}

func (s *CreateTokenPlanKeyRequest) SetDescription(v string) *CreateTokenPlanKeyRequest {
	s.Description = &v
	return s
}

func (s *CreateTokenPlanKeyRequest) Validate() error {
	return dara.Validate(s)
}
