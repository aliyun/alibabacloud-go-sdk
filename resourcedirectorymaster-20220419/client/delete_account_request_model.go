// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAbandonableCheckId(v []*string) *DeleteAccountRequest
	GetAbandonableCheckId() []*string
	SetAccountId(v string) *DeleteAccountRequest
	GetAccountId() *string
}

type DeleteAccountRequest struct {
	// The IDs of the check items that you can choose to ignore for the member deletion.
	//
	// You can obtain the IDs from the response of the [GetAccountDeletionCheckResult](~~GetAccountDeletionCheckResult~~) operation.
	AbandonableCheckId []*string `json:"AbandonableCheckId,omitempty" xml:"AbandonableCheckId,omitempty" type:"Repeated"`
	// The Alibaba Cloud account ID of the member that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// 169946124551****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s DeleteAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccountRequest) GetAbandonableCheckId() []*string {
	return s.AbandonableCheckId
}

func (s *DeleteAccountRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *DeleteAccountRequest) SetAbandonableCheckId(v []*string) *DeleteAccountRequest {
	s.AbandonableCheckId = v
	return s
}

func (s *DeleteAccountRequest) SetAccountId(v string) *DeleteAccountRequest {
	s.AccountId = &v
	return s
}

func (s *DeleteAccountRequest) Validate() error {
	return dara.Validate(s)
}
