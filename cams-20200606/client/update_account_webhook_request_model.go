// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAccountWebhookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *UpdateAccountWebhookRequest
	GetCustSpaceId() *string
	SetHttpFlag(v string) *UpdateAccountWebhookRequest
	GetHttpFlag() *string
	SetOwnerId(v int64) *UpdateAccountWebhookRequest
	GetOwnerId() *int64
	SetQueueFlag(v string) *UpdateAccountWebhookRequest
	GetQueueFlag() *string
	SetResourceOwnerAccount(v string) *UpdateAccountWebhookRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateAccountWebhookRequest
	GetResourceOwnerId() *int64
	SetStatusCallbackUrl(v string) *UpdateAccountWebhookRequest
	GetStatusCallbackUrl() *string
}

type UpdateAccountWebhookRequest struct {
	// The space ID of the RAM user within the independent software vendor (ISV) account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 293483938849493**
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// Specifies whether to use HTTP callbacks to receive message receipts. Valid values:
	//
	// 	- Y: indicates that HTTP callbacks are used to receive receipts.
	//
	// 	- N: indicates that HTTP callbacks are not used to receive receipts.
	//
	// example:
	//
	// Y
	HttpFlag *string `json:"HttpFlag,omitempty" xml:"HttpFlag,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether to use Message Service (MNS) queues to receive receipts. Valid values:
	//
	// 	- Y: indicates that MNS queues are used to receive receipts.
	//
	// 	- N: indicates that MNS queues are not used to receive receipts.
	//
	// example:
	//
	// N
	QueueFlag            *string `json:"QueueFlag,omitempty" xml:"QueueFlag,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The callback URL to which status reports are sent by using HTTP callbacks.
	//
	// example:
	//
	// http://www.aliyun.com
	StatusCallbackUrl *string `json:"StatusCallbackUrl,omitempty" xml:"StatusCallbackUrl,omitempty"`
}

func (s UpdateAccountWebhookRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAccountWebhookRequest) GoString() string {
	return s.String()
}

func (s *UpdateAccountWebhookRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *UpdateAccountWebhookRequest) GetHttpFlag() *string {
	return s.HttpFlag
}

func (s *UpdateAccountWebhookRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateAccountWebhookRequest) GetQueueFlag() *string {
	return s.QueueFlag
}

func (s *UpdateAccountWebhookRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateAccountWebhookRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateAccountWebhookRequest) GetStatusCallbackUrl() *string {
	return s.StatusCallbackUrl
}

func (s *UpdateAccountWebhookRequest) SetCustSpaceId(v string) *UpdateAccountWebhookRequest {
	s.CustSpaceId = &v
	return s
}

func (s *UpdateAccountWebhookRequest) SetHttpFlag(v string) *UpdateAccountWebhookRequest {
	s.HttpFlag = &v
	return s
}

func (s *UpdateAccountWebhookRequest) SetOwnerId(v int64) *UpdateAccountWebhookRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateAccountWebhookRequest) SetQueueFlag(v string) *UpdateAccountWebhookRequest {
	s.QueueFlag = &v
	return s
}

func (s *UpdateAccountWebhookRequest) SetResourceOwnerAccount(v string) *UpdateAccountWebhookRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateAccountWebhookRequest) SetResourceOwnerId(v int64) *UpdateAccountWebhookRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateAccountWebhookRequest) SetStatusCallbackUrl(v string) *UpdateAccountWebhookRequest {
	s.StatusCallbackUrl = &v
	return s
}

func (s *UpdateAccountWebhookRequest) Validate() error {
	return dara.Validate(s)
}
