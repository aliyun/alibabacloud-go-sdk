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
	AbandonableCheckIdShrink *string `json:"AbandonableCheckId,omitempty" xml:"AbandonableCheckId,omitempty"`
	// The type of the deletion. Valid values:
	//
	// 	- 0: direct deletion. If the member does not have pay-as-you-go resources that are purchased within the previous 30 days, the system directly deletes the member.
	//
	// 	- 1: deletion with a silence period. If the member has pay-as-you-go resources that are purchased within the previous 30 days, the member enters a silence period of 45 days. The system starts to delete the member until the silence period ends. For more information about the silence period, see [What is the silence period for member deletion?](https://help.aliyun.com/document_detail/446079.html)
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
