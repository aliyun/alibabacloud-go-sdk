// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveReceiverDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDetail(v string) *SaveReceiverDetailRequest
	GetDetail() *string
	SetOwnerId(v int64) *SaveReceiverDetailRequest
	GetOwnerId() *int64
	SetReceiverId(v string) *SaveReceiverDetailRequest
	GetReceiverId() *string
	SetResourceOwnerAccount(v string) *SaveReceiverDetailRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SaveReceiverDetailRequest
	GetResourceOwnerId() *int64
}

type SaveReceiverDetailRequest struct {
	// Content, supports uploading multiple recipients at once, with a limit of 500 records per upload. Each record is separated by {} and commas, example:
	//
	// [{ },{ },{ }...], the format within {} is as follows:
	//
	// [{"b":"birthday","e":"xxx@example.net","g":"gender","m":"mobile","n":"nickname","u":"name"}], when passing values, pass it as a string, not a list.
	//
	// If a duplicate recipient address is inserted, it will return "ErrorCount": 1
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"b":"birthday","e":"xxx@alibaba-inc.com","g":"gender","m":"mobile","n":"nickname","u":"name"}]
	Detail  *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Recipient list ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 34642
	ReceiverId           *string `json:"ReceiverId,omitempty" xml:"ReceiverId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SaveReceiverDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveReceiverDetailRequest) GoString() string {
	return s.String()
}

func (s *SaveReceiverDetailRequest) GetDetail() *string {
	return s.Detail
}

func (s *SaveReceiverDetailRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SaveReceiverDetailRequest) GetReceiverId() *string {
	return s.ReceiverId
}

func (s *SaveReceiverDetailRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SaveReceiverDetailRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SaveReceiverDetailRequest) SetDetail(v string) *SaveReceiverDetailRequest {
	s.Detail = &v
	return s
}

func (s *SaveReceiverDetailRequest) SetOwnerId(v int64) *SaveReceiverDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *SaveReceiverDetailRequest) SetReceiverId(v string) *SaveReceiverDetailRequest {
	s.ReceiverId = &v
	return s
}

func (s *SaveReceiverDetailRequest) SetResourceOwnerAccount(v string) *SaveReceiverDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SaveReceiverDetailRequest) SetResourceOwnerId(v int64) *SaveReceiverDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SaveReceiverDetailRequest) Validate() error {
	return dara.Validate(s)
}
