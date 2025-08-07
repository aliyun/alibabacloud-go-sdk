// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateActivationCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunOrderId(v string) *CreateActivationCodeRequest
	GetAliyunOrderId() *string
	SetDescription(v string) *CreateActivationCodeRequest
	GetDescription() *string
	SetMacAddress(v string) *CreateActivationCodeRequest
	GetMacAddress() *string
	SetName(v string) *CreateActivationCodeRequest
	GetName() *string
	SetOwnerAccount(v string) *CreateActivationCodeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateActivationCodeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateActivationCodeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateActivationCodeRequest
	GetResourceOwnerId() *int64
	SetSystemIdentifier(v string) *CreateActivationCodeRequest
	GetSystemIdentifier() *string
}

type CreateActivationCodeRequest struct {
	// The Alibaba Cloud order ID (including the virtual order ID).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2233****445566
	AliyunOrderId *string `json:"AliyunOrderId,omitempty" xml:"AliyunOrderId,omitempty"`
	// The description of the activation code.
	//
	// example:
	//
	// testCode
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The MAC address.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12:34:56:78:98:00
	MacAddress *string `json:"MacAddress,omitempty" xml:"MacAddress,omitempty"`
	// The name of the activation code. The name can contain only letters, digits, underscores (_), and hyphens (-). The activation code file downloaded from the console is named based on this name.
	//
	// This parameter is required.
	//
	// example:
	//
	// testName
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The system identifier of the database. This parameter is required if you set AllowEmptySystemIdentifier to false.
	//
	// example:
	//
	// 1234567890123456
	SystemIdentifier *string `json:"SystemIdentifier,omitempty" xml:"SystemIdentifier,omitempty"`
}

func (s CreateActivationCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateActivationCodeRequest) GoString() string {
	return s.String()
}

func (s *CreateActivationCodeRequest) GetAliyunOrderId() *string {
	return s.AliyunOrderId
}

func (s *CreateActivationCodeRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateActivationCodeRequest) GetMacAddress() *string {
	return s.MacAddress
}

func (s *CreateActivationCodeRequest) GetName() *string {
	return s.Name
}

func (s *CreateActivationCodeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateActivationCodeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateActivationCodeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateActivationCodeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateActivationCodeRequest) GetSystemIdentifier() *string {
	return s.SystemIdentifier
}

func (s *CreateActivationCodeRequest) SetAliyunOrderId(v string) *CreateActivationCodeRequest {
	s.AliyunOrderId = &v
	return s
}

func (s *CreateActivationCodeRequest) SetDescription(v string) *CreateActivationCodeRequest {
	s.Description = &v
	return s
}

func (s *CreateActivationCodeRequest) SetMacAddress(v string) *CreateActivationCodeRequest {
	s.MacAddress = &v
	return s
}

func (s *CreateActivationCodeRequest) SetName(v string) *CreateActivationCodeRequest {
	s.Name = &v
	return s
}

func (s *CreateActivationCodeRequest) SetOwnerAccount(v string) *CreateActivationCodeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateActivationCodeRequest) SetOwnerId(v int64) *CreateActivationCodeRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateActivationCodeRequest) SetResourceOwnerAccount(v string) *CreateActivationCodeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateActivationCodeRequest) SetResourceOwnerId(v int64) *CreateActivationCodeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateActivationCodeRequest) SetSystemIdentifier(v string) *CreateActivationCodeRequest {
	s.SystemIdentifier = &v
	return s
}

func (s *CreateActivationCodeRequest) Validate() error {
	return dara.Validate(s)
}
