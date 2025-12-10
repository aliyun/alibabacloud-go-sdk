// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountDeletionCheckResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *GetAccountDeletionCheckResultRequest
	GetAccountId() *string
}

type GetAccountDeletionCheckResultRequest struct {
	// The ID of the member that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// 179855839641****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s GetAccountDeletionCheckResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAccountDeletionCheckResultRequest) GoString() string {
	return s.String()
}

func (s *GetAccountDeletionCheckResultRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *GetAccountDeletionCheckResultRequest) SetAccountId(v string) *GetAccountDeletionCheckResultRequest {
	s.AccountId = &v
	return s
}

func (s *GetAccountDeletionCheckResultRequest) Validate() error {
	return dara.Validate(s)
}
