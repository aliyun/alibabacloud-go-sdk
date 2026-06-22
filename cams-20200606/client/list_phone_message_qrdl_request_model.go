// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPhoneMessageQrdlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *ListPhoneMessageQrdlRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *ListPhoneMessageQrdlRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *ListPhoneMessageQrdlRequest
	GetPhoneNumber() *string
	SetResourceOwnerAccount(v string) *ListPhoneMessageQrdlRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListPhoneMessageQrdlRequest
	GetResourceOwnerId() *int64
}

type ListPhoneMessageQrdlRequest struct {
	// example:
	//
	// 示例值
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListPhoneMessageQrdlRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPhoneMessageQrdlRequest) GoString() string {
	return s.String()
}

func (s *ListPhoneMessageQrdlRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *ListPhoneMessageQrdlRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListPhoneMessageQrdlRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *ListPhoneMessageQrdlRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListPhoneMessageQrdlRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListPhoneMessageQrdlRequest) SetCustSpaceId(v string) *ListPhoneMessageQrdlRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ListPhoneMessageQrdlRequest) SetOwnerId(v int64) *ListPhoneMessageQrdlRequest {
	s.OwnerId = &v
	return s
}

func (s *ListPhoneMessageQrdlRequest) SetPhoneNumber(v string) *ListPhoneMessageQrdlRequest {
	s.PhoneNumber = &v
	return s
}

func (s *ListPhoneMessageQrdlRequest) SetResourceOwnerAccount(v string) *ListPhoneMessageQrdlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListPhoneMessageQrdlRequest) SetResourceOwnerId(v int64) *ListPhoneMessageQrdlRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListPhoneMessageQrdlRequest) Validate() error {
	return dara.Validate(s)
}
