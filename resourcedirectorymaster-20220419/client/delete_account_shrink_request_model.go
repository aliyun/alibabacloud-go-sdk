// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccountShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAbandonableCheckIdShrink(v string) *DeleteAccountShrinkRequest
	GetAbandonableCheckIdShrink() *string
	SetAccountId(v string) *DeleteAccountShrinkRequest
	GetAccountId() *string
}

type DeleteAccountShrinkRequest struct {
	// The IDs of the check items that you can choose to ignore for the member deletion.
	//
	// You can obtain the IDs from the response of the [GetAccountDeletionCheckResult](~~GetAccountDeletionCheckResult~~) operation.
	AbandonableCheckIdShrink *string `json:"AbandonableCheckId,omitempty" xml:"AbandonableCheckId,omitempty"`
	// The Alibaba Cloud account ID of the member that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// 169946124551****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s DeleteAccountShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccountShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccountShrinkRequest) GetAbandonableCheckIdShrink() *string {
	return s.AbandonableCheckIdShrink
}

func (s *DeleteAccountShrinkRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *DeleteAccountShrinkRequest) SetAbandonableCheckIdShrink(v string) *DeleteAccountShrinkRequest {
	s.AbandonableCheckIdShrink = &v
	return s
}

func (s *DeleteAccountShrinkRequest) SetAccountId(v string) *DeleteAccountShrinkRequest {
	s.AccountId = &v
	return s
}

func (s *DeleteAccountShrinkRequest) Validate() error {
	return dara.Validate(s)
}
