// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSmsSignRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DeleteSmsSignRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteSmsSignRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteSmsSignRequest
	GetResourceOwnerId() *int64
	SetSignName(v string) *DeleteSmsSignRequest
	GetSignName() *string
}

type DeleteSmsSignRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The SMS signature. You can delete signatures that are in the Withdrawn, Failed, or Approved state. You cannot delete signatures that are in the Pending Approval state.
	//
	// You can call the [QuerySmsSignList](https://help.aliyun.com/document_detail/419282.html) operation to query the signatures that have been applied for under the current account, or view the signature list on the [Signature Management](https://dysms.console.aliyun.com/domestic/text/sign) page in the SMS console.
	//
	// 	Notice: Deleted SMS signatures cannot be recovered, and the signature can no longer be used to send SMS messages. Proceed with caution.
	//
	// This parameter is required.
	//
	// example:
	//
	// 阿里云
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
}

func (s DeleteSmsSignRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSmsSignRequest) GoString() string {
	return s.String()
}

func (s *DeleteSmsSignRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteSmsSignRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteSmsSignRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteSmsSignRequest) GetSignName() *string {
	return s.SignName
}

func (s *DeleteSmsSignRequest) SetOwnerId(v int64) *DeleteSmsSignRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSmsSignRequest) SetResourceOwnerAccount(v string) *DeleteSmsSignRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteSmsSignRequest) SetResourceOwnerId(v int64) *DeleteSmsSignRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteSmsSignRequest) SetSignName(v string) *DeleteSmsSignRequest {
	s.SignName = &v
	return s
}

func (s *DeleteSmsSignRequest) Validate() error {
	return dara.Validate(s)
}
