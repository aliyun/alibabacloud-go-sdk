// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountDeletionStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *GetAccountDeletionStatusRequest
	GetAccountId() *string
}

type GetAccountDeletionStatusRequest struct {
	// The Alibaba Cloud account ID of the member.
	//
	// This parameter is required.
	//
	// example:
	//
	// 169946124551****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s GetAccountDeletionStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAccountDeletionStatusRequest) GoString() string {
	return s.String()
}

func (s *GetAccountDeletionStatusRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *GetAccountDeletionStatusRequest) SetAccountId(v string) *GetAccountDeletionStatusRequest {
	s.AccountId = &v
	return s
}

func (s *GetAccountDeletionStatusRequest) Validate() error {
	return dara.Validate(s)
}
