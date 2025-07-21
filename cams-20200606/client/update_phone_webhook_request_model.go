// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePhoneWebhookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *UpdatePhoneWebhookRequest
	GetCustSpaceId() *string
	SetHttpFlag(v string) *UpdatePhoneWebhookRequest
	GetHttpFlag() *string
	SetOwnerId(v int64) *UpdatePhoneWebhookRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *UpdatePhoneWebhookRequest
	GetPhoneNumber() *string
	SetQueueFlag(v string) *UpdatePhoneWebhookRequest
	GetQueueFlag() *string
	SetResourceOwnerAccount(v string) *UpdatePhoneWebhookRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdatePhoneWebhookRequest
	GetResourceOwnerId() *int64
	SetStatusCallbackUrl(v string) *UpdatePhoneWebhookRequest
	GetStatusCallbackUrl() *string
	SetUpCallbackUrl(v string) *UpdatePhoneWebhookRequest
	GetUpCallbackUrl() *string
}

type UpdatePhoneWebhookRequest struct {
	// SpaceId for ISV sub clients.
	//
	// This parameter is required.
	//
	// example:
	//
	// 293483938849493****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// Whether to use HTTP to receive receipts. Value:
	//
	// 	- Y: Yes.
	//
	// 	- N: No.
	//
	// example:
	//
	// Y
	HttpFlag *string `json:"HttpFlag,omitempty" xml:"HttpFlag,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8613800001234
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// Whether to use queue method to receive receipts. Value:
	//
	// 	- Y: Yes.
	//
	// 	- N: No.
	//
	// example:
	//
	// N
	QueueFlag            *string `json:"QueueFlag,omitempty" xml:"QueueFlag,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// HTTP status report interface callback address.
	//
	// example:
	//
	// http://www.aliyun.com
	StatusCallbackUrl *string `json:"StatusCallbackUrl,omitempty" xml:"StatusCallbackUrl,omitempty"`
	// HTTP upstream message interface callback address.
	//
	// example:
	//
	// http://aliyun.com
	UpCallbackUrl *string `json:"UpCallbackUrl,omitempty" xml:"UpCallbackUrl,omitempty"`
}

func (s UpdatePhoneWebhookRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePhoneWebhookRequest) GoString() string {
	return s.String()
}

func (s *UpdatePhoneWebhookRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *UpdatePhoneWebhookRequest) GetHttpFlag() *string {
	return s.HttpFlag
}

func (s *UpdatePhoneWebhookRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdatePhoneWebhookRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *UpdatePhoneWebhookRequest) GetQueueFlag() *string {
	return s.QueueFlag
}

func (s *UpdatePhoneWebhookRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdatePhoneWebhookRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdatePhoneWebhookRequest) GetStatusCallbackUrl() *string {
	return s.StatusCallbackUrl
}

func (s *UpdatePhoneWebhookRequest) GetUpCallbackUrl() *string {
	return s.UpCallbackUrl
}

func (s *UpdatePhoneWebhookRequest) SetCustSpaceId(v string) *UpdatePhoneWebhookRequest {
	s.CustSpaceId = &v
	return s
}

func (s *UpdatePhoneWebhookRequest) SetHttpFlag(v string) *UpdatePhoneWebhookRequest {
	s.HttpFlag = &v
	return s
}

func (s *UpdatePhoneWebhookRequest) SetOwnerId(v int64) *UpdatePhoneWebhookRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdatePhoneWebhookRequest) SetPhoneNumber(v string) *UpdatePhoneWebhookRequest {
	s.PhoneNumber = &v
	return s
}

func (s *UpdatePhoneWebhookRequest) SetQueueFlag(v string) *UpdatePhoneWebhookRequest {
	s.QueueFlag = &v
	return s
}

func (s *UpdatePhoneWebhookRequest) SetResourceOwnerAccount(v string) *UpdatePhoneWebhookRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdatePhoneWebhookRequest) SetResourceOwnerId(v int64) *UpdatePhoneWebhookRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdatePhoneWebhookRequest) SetStatusCallbackUrl(v string) *UpdatePhoneWebhookRequest {
	s.StatusCallbackUrl = &v
	return s
}

func (s *UpdatePhoneWebhookRequest) SetUpCallbackUrl(v string) *UpdatePhoneWebhookRequest {
	s.UpCallbackUrl = &v
	return s
}

func (s *UpdatePhoneWebhookRequest) Validate() error {
	return dara.Validate(s)
}
