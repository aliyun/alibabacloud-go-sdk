// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePhoneMessageQrdlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *DeletePhoneMessageQrdlRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *DeletePhoneMessageQrdlRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *DeletePhoneMessageQrdlRequest
	GetPhoneNumber() *string
	SetQrdlCode(v string) *DeletePhoneMessageQrdlRequest
	GetQrdlCode() *string
	SetResourceOwnerAccount(v string) *DeletePhoneMessageQrdlRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeletePhoneMessageQrdlRequest
	GetResourceOwnerId() *int64
}

type DeletePhoneMessageQrdlRequest struct {
	// example:
	//
	// 示例值示例值
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	QrdlCode             *string `json:"QrdlCode,omitempty" xml:"QrdlCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeletePhoneMessageQrdlRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePhoneMessageQrdlRequest) GoString() string {
	return s.String()
}

func (s *DeletePhoneMessageQrdlRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *DeletePhoneMessageQrdlRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeletePhoneMessageQrdlRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *DeletePhoneMessageQrdlRequest) GetQrdlCode() *string {
	return s.QrdlCode
}

func (s *DeletePhoneMessageQrdlRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeletePhoneMessageQrdlRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeletePhoneMessageQrdlRequest) SetCustSpaceId(v string) *DeletePhoneMessageQrdlRequest {
	s.CustSpaceId = &v
	return s
}

func (s *DeletePhoneMessageQrdlRequest) SetOwnerId(v int64) *DeletePhoneMessageQrdlRequest {
	s.OwnerId = &v
	return s
}

func (s *DeletePhoneMessageQrdlRequest) SetPhoneNumber(v string) *DeletePhoneMessageQrdlRequest {
	s.PhoneNumber = &v
	return s
}

func (s *DeletePhoneMessageQrdlRequest) SetQrdlCode(v string) *DeletePhoneMessageQrdlRequest {
	s.QrdlCode = &v
	return s
}

func (s *DeletePhoneMessageQrdlRequest) SetResourceOwnerAccount(v string) *DeletePhoneMessageQrdlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeletePhoneMessageQrdlRequest) SetResourceOwnerId(v int64) *DeletePhoneMessageQrdlRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeletePhoneMessageQrdlRequest) Validate() error {
	return dara.Validate(s)
}
