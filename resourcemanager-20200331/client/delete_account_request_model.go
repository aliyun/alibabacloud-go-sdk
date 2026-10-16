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
	// The ID of the check item that you can choose to abandon and continue with the member deletion.
	//
	// You can obtain this ID from the AbandonableChecks response parameter of [GetAccountDeletionCheckResult](https://help.aliyun.com/document_detail/448775.html).
	AbandonableCheckId []*string `json:"AbandonableCheckId,omitempty" xml:"AbandonableCheckId,omitempty" type:"Repeated"`
	// The ID of the member to delete.
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
