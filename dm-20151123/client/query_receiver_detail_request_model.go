// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryReceiverDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyWord(v string) *QueryReceiverDetailRequest
	GetKeyWord() *string
	SetNextStart(v string) *QueryReceiverDetailRequest
	GetNextStart() *string
	SetOwnerId(v int64) *QueryReceiverDetailRequest
	GetOwnerId() *int64
	SetPageSize(v int32) *QueryReceiverDetailRequest
	GetPageSize() *int32
	SetReceiverId(v string) *QueryReceiverDetailRequest
	GetReceiverId() *string
	SetResourceOwnerAccount(v string) *QueryReceiverDetailRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryReceiverDetailRequest
	GetResourceOwnerId() *int64
}

type QueryReceiverDetailRequest struct {
	// Recipient address, length 0-50
	//
	// example:
	//
	// b***@example.net
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	// Starting position for the next item, default: 0
	//
	// example:
	//
	// 0
	NextStart *string `json:"NextStart,omitempty" xml:"NextStart,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Number of items per page, default: 10
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Recipient list ID (returned when creating a recipient list using the CreateReceiver API).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1235
	ReceiverId           *string `json:"ReceiverId,omitempty" xml:"ReceiverId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryReceiverDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryReceiverDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryReceiverDetailRequest) GetKeyWord() *string {
	return s.KeyWord
}

func (s *QueryReceiverDetailRequest) GetNextStart() *string {
	return s.NextStart
}

func (s *QueryReceiverDetailRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryReceiverDetailRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryReceiverDetailRequest) GetReceiverId() *string {
	return s.ReceiverId
}

func (s *QueryReceiverDetailRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryReceiverDetailRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryReceiverDetailRequest) SetKeyWord(v string) *QueryReceiverDetailRequest {
	s.KeyWord = &v
	return s
}

func (s *QueryReceiverDetailRequest) SetNextStart(v string) *QueryReceiverDetailRequest {
	s.NextStart = &v
	return s
}

func (s *QueryReceiverDetailRequest) SetOwnerId(v int64) *QueryReceiverDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryReceiverDetailRequest) SetPageSize(v int32) *QueryReceiverDetailRequest {
	s.PageSize = &v
	return s
}

func (s *QueryReceiverDetailRequest) SetReceiverId(v string) *QueryReceiverDetailRequest {
	s.ReceiverId = &v
	return s
}

func (s *QueryReceiverDetailRequest) SetResourceOwnerAccount(v string) *QueryReceiverDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryReceiverDetailRequest) SetResourceOwnerId(v int64) *QueryReceiverDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryReceiverDetailRequest) Validate() error {
	return dara.Validate(s)
}
