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
	// The ID of the check item that you can choose to abandon and continue with the member deletion.
	//
	// You can obtain this ID from the AbandonableChecks response parameter of [GetAccountDeletionCheckResult](https://help.aliyun.com/document_detail/448775.html).
	AbandonableCheckIdShrink *string `json:"AbandonableCheckId,omitempty" xml:"AbandonableCheckId,omitempty"`
	// The ID of the member to delete.
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
