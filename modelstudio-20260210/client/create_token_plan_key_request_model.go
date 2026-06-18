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
	SetCallerUacAccountId(v string) *CreateTokenPlanKeyRequest
	GetCallerUacAccountId() *string
	SetDescription(v string) *CreateTokenPlanKeyRequest
	GetDescription() *string
	SetNamespaceId(v string) *CreateTokenPlanKeyRequest
	GetNamespaceId() *string
	SetWorkspaceId(v string) *CreateTokenPlanKeyRequest
	GetWorkspaceId() *string
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
	// The account ID of the caller that identifies the initiator of this call.
	//
	// example:
	//
	// acc_123456789
	CallerUacAccountId *string `json:"CallerUacAccountId,omitempty" xml:"CallerUacAccountId,omitempty"`
	// The description of the key.
	//
	// example:
	//
	// APIKEY描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The product namespace ID. For the TokenPlan product, this field is fixed to namespace-1.
	//
	// example:
	//
	// namespace-1
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ws_123456789
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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

func (s *CreateTokenPlanKeyRequest) GetCallerUacAccountId() *string {
	return s.CallerUacAccountId
}

func (s *CreateTokenPlanKeyRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateTokenPlanKeyRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *CreateTokenPlanKeyRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateTokenPlanKeyRequest) SetAccountId(v string) *CreateTokenPlanKeyRequest {
	s.AccountId = &v
	return s
}

func (s *CreateTokenPlanKeyRequest) SetCallerUacAccountId(v string) *CreateTokenPlanKeyRequest {
	s.CallerUacAccountId = &v
	return s
}

func (s *CreateTokenPlanKeyRequest) SetDescription(v string) *CreateTokenPlanKeyRequest {
	s.Description = &v
	return s
}

func (s *CreateTokenPlanKeyRequest) SetNamespaceId(v string) *CreateTokenPlanKeyRequest {
	s.NamespaceId = &v
	return s
}

func (s *CreateTokenPlanKeyRequest) SetWorkspaceId(v string) *CreateTokenPlanKeyRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateTokenPlanKeyRequest) Validate() error {
	return dara.Validate(s)
}
