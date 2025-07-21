// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePhoneMessageQrdlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *CreatePhoneMessageQrdlRequest
	GetCustSpaceId() *string
	SetGenerateQrImage(v string) *CreatePhoneMessageQrdlRequest
	GetGenerateQrImage() *string
	SetOwnerId(v int64) *CreatePhoneMessageQrdlRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *CreatePhoneMessageQrdlRequest
	GetPhoneNumber() *string
	SetPrefilledMessage(v string) *CreatePhoneMessageQrdlRequest
	GetPrefilledMessage() *string
	SetResourceOwnerAccount(v string) *CreatePhoneMessageQrdlRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreatePhoneMessageQrdlRequest
	GetResourceOwnerId() *int64
}

type CreatePhoneMessageQrdlRequest struct {
	// example:
	//
	// 示例值示例值示例值
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	GenerateQrImage *string `json:"GenerateQrImage,omitempty" xml:"GenerateQrImage,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	PrefilledMessage     *string `json:"PrefilledMessage,omitempty" xml:"PrefilledMessage,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreatePhoneMessageQrdlRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePhoneMessageQrdlRequest) GoString() string {
	return s.String()
}

func (s *CreatePhoneMessageQrdlRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *CreatePhoneMessageQrdlRequest) GetGenerateQrImage() *string {
	return s.GenerateQrImage
}

func (s *CreatePhoneMessageQrdlRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreatePhoneMessageQrdlRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *CreatePhoneMessageQrdlRequest) GetPrefilledMessage() *string {
	return s.PrefilledMessage
}

func (s *CreatePhoneMessageQrdlRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreatePhoneMessageQrdlRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreatePhoneMessageQrdlRequest) SetCustSpaceId(v string) *CreatePhoneMessageQrdlRequest {
	s.CustSpaceId = &v
	return s
}

func (s *CreatePhoneMessageQrdlRequest) SetGenerateQrImage(v string) *CreatePhoneMessageQrdlRequest {
	s.GenerateQrImage = &v
	return s
}

func (s *CreatePhoneMessageQrdlRequest) SetOwnerId(v int64) *CreatePhoneMessageQrdlRequest {
	s.OwnerId = &v
	return s
}

func (s *CreatePhoneMessageQrdlRequest) SetPhoneNumber(v string) *CreatePhoneMessageQrdlRequest {
	s.PhoneNumber = &v
	return s
}

func (s *CreatePhoneMessageQrdlRequest) SetPrefilledMessage(v string) *CreatePhoneMessageQrdlRequest {
	s.PrefilledMessage = &v
	return s
}

func (s *CreatePhoneMessageQrdlRequest) SetResourceOwnerAccount(v string) *CreatePhoneMessageQrdlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreatePhoneMessageQrdlRequest) SetResourceOwnerId(v int64) *CreatePhoneMessageQrdlRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreatePhoneMessageQrdlRequest) Validate() error {
	return dara.Validate(s)
}
