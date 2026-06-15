// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveReceiverDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomDetail(v string) *SaveReceiverDetailRequest
	GetCustomDetail() *string
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
	// The recipient\\"s email address and template parameters, in an array format.
	//
	// example:
	//
	// [{ "Email": "example@alimail.com","CustomData": {"name":"Tom","age":"30"}}]
	CustomDetail *string `json:"CustomDetail,omitempty" xml:"CustomDetail,omitempty"`
	// The recipient details. You can upload up to 500 recipients in a single request. The value is a string in a JSON array format. Each object in the array represents a recipient. For example: \\`[{ },{ },{ }...]\\`. The format for each recipient object is \\`[{"b":"birthday","e":"xxx\\@example.net","g":"gender","m":"mobile","n":"nickname","u":"name"}]\\`. If you add a duplicate recipient address, the system returns \\`"ErrorCount": 1\\`.
	//
	// The format is \\`[{ },{ },{ }...]\\`. The format of the content within each \\`{}\\` is as follows:
	//
	// [{"b":"birthday","e":"xxx\\@example.net","g":"gender","m":"mobile","n":"nickname","u":"name"}]. Pass the value as a string, not a list.
	//
	// Inserting a duplicate recipient address returns "ErrorCount": 1.
	//
	// example:
	//
	// [{"b":"birthday","e":"xxx@alibaba-inc.com","g":"gender","m":"mobile","n":"nickname","u":"name"}]
	Detail  *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the recipient list.
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

func (s *SaveReceiverDetailRequest) GetCustomDetail() *string {
	return s.CustomDetail
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

func (s *SaveReceiverDetailRequest) SetCustomDetail(v string) *SaveReceiverDetailRequest {
	s.CustomDetail = &v
	return s
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
