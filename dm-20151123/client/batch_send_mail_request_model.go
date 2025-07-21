// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSendMailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *BatchSendMailRequest
	GetAccountName() *string
	SetAddressType(v int32) *BatchSendMailRequest
	GetAddressType() *int32
	SetClickTrace(v string) *BatchSendMailRequest
	GetClickTrace() *string
	SetHeaders(v string) *BatchSendMailRequest
	GetHeaders() *string
	SetIpPoolId(v string) *BatchSendMailRequest
	GetIpPoolId() *string
	SetOwnerId(v int64) *BatchSendMailRequest
	GetOwnerId() *int64
	SetReceiversName(v string) *BatchSendMailRequest
	GetReceiversName() *string
	SetReplyAddress(v string) *BatchSendMailRequest
	GetReplyAddress() *string
	SetReplyAddressAlias(v string) *BatchSendMailRequest
	GetReplyAddressAlias() *string
	SetResourceOwnerAccount(v string) *BatchSendMailRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *BatchSendMailRequest
	GetResourceOwnerId() *int64
	SetTagName(v string) *BatchSendMailRequest
	GetTagName() *string
	SetTemplateName(v string) *BatchSendMailRequest
	GetTemplateName() *string
	SetUnSubscribeFilterLevel(v string) *BatchSendMailRequest
	GetUnSubscribeFilterLevel() *string
	SetUnSubscribeLinkType(v string) *BatchSendMailRequest
	GetUnSubscribeLinkType() *string
}

type BatchSendMailRequest struct {
	// The sending address configured in the management console.
	//
	// This parameter is required.
	//
	// example:
	//
	// test@example.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// - 0: Random account
	//
	// - 1: Sending address
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	AddressType *int32 `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// - 1: Enable data tracking function
	//
	// - 0 (default): Disable data tracking function
	//
	// example:
	//
	// 0
	ClickTrace *string `json:"ClickTrace,omitempty" xml:"ClickTrace,omitempty"`
	// Currently, the standard fields that can be added to the email header are Message-ID, List-Unsubscribe, and List-Unsubscribe-Post. Standard fields will overwrite the existing values in the email header, while non-standard fields must start with X-User- and will be appended to the email header. Currently, up to 10 headers can be passed in JSON format, and both standard and non-standard fields must comply with the syntax requirements for headers.
	//
	// example:
	//
	// {
	//
	//   "Message-ID": "<msg0001@example.com>",
	//
	//   "X-User-UID1": "UID-1-000001",
	//
	//   "X-User-UID2": "UID-2-000001"
	//
	// }
	Headers *string `json:"Headers,omitempty" xml:"Headers,omitempty"`
	// dedicated IP pool ID. Users who have purchased an dedicated IP can use this parameter to specify the outgoing IP for this send operation.
	//
	// example:
	//
	// xxx
	IpPoolId *string `json:"IpPoolId,omitempty" xml:"IpPoolId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the recipient list that has been created and uploaded with recipients. Note: The recipient list should not be deleted until at least 10 minutes after the task is triggered, otherwise it may cause sending failure.
	//
	// This parameter is required.
	//
	// example:
	//
	// test2
	ReceiversName *string `json:"ReceiversName,omitempty" xml:"ReceiversName,omitempty"`
	// Reply address
	//
	// example:
	//
	// test2***@example.net
	ReplyAddress *string `json:"ReplyAddress,omitempty" xml:"ReplyAddress,omitempty"`
	// Alias for the reply address
	//
	// example:
	//
	// Lucy
	ReplyAddressAlias    *string `json:"ReplyAddressAlias,omitempty" xml:"ReplyAddressAlias,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Email tag name.
	//
	// example:
	//
	// test3
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	// The name of the template that has been created and approved in advance.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// Filtering level. Refer to the [Unsubscribe Function Link Generation and Filtering Mechanism](https://help.aliyun.com/document_detail/2689048.html) document.
	//
	// - disabled: No filtering
	//
	// - default: Use the default strategy, bulk addresses use sender address-level filtering
	//
	// - mailfrom: Sender address-level filtering
	//
	// - mailfrom_domain: Sender domain-level filtering
	//
	// - edm_id: Account-level filtering
	//
	// example:
	//
	// mailfrom_domain
	UnSubscribeFilterLevel *string `json:"UnSubscribeFilterLevel,omitempty" xml:"UnSubscribeFilterLevel,omitempty"`
	// The type of generated unsubscribe link. Refer to the [Unsubscribe Function Link Generation and Filtering Mechanism](https://help.aliyun.com/document_detail/2689048.html) document.
	//
	// - disabled: Do not generate
	//
	// - default: Use the default strategy: Generate an unsubscribe link when a bulk-type sending address sends to specific domains, such as those containing keywords like "gmail", "yahoo",
	//
	// "google", "aol.com", "hotmail",
	//
	// "outlook", "ymail.com", etc.
	//
	// - zh-cn: Generate, for future content preparation
	//
	// - en-us: Generate, for future content preparation
	//
	// example:
	//
	// default
	UnSubscribeLinkType *string `json:"UnSubscribeLinkType,omitempty" xml:"UnSubscribeLinkType,omitempty"`
}

func (s BatchSendMailRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchSendMailRequest) GoString() string {
	return s.String()
}

func (s *BatchSendMailRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *BatchSendMailRequest) GetAddressType() *int32 {
	return s.AddressType
}

func (s *BatchSendMailRequest) GetClickTrace() *string {
	return s.ClickTrace
}

func (s *BatchSendMailRequest) GetHeaders() *string {
	return s.Headers
}

func (s *BatchSendMailRequest) GetIpPoolId() *string {
	return s.IpPoolId
}

func (s *BatchSendMailRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BatchSendMailRequest) GetReceiversName() *string {
	return s.ReceiversName
}

func (s *BatchSendMailRequest) GetReplyAddress() *string {
	return s.ReplyAddress
}

func (s *BatchSendMailRequest) GetReplyAddressAlias() *string {
	return s.ReplyAddressAlias
}

func (s *BatchSendMailRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *BatchSendMailRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *BatchSendMailRequest) GetTagName() *string {
	return s.TagName
}

func (s *BatchSendMailRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *BatchSendMailRequest) GetUnSubscribeFilterLevel() *string {
	return s.UnSubscribeFilterLevel
}

func (s *BatchSendMailRequest) GetUnSubscribeLinkType() *string {
	return s.UnSubscribeLinkType
}

func (s *BatchSendMailRequest) SetAccountName(v string) *BatchSendMailRequest {
	s.AccountName = &v
	return s
}

func (s *BatchSendMailRequest) SetAddressType(v int32) *BatchSendMailRequest {
	s.AddressType = &v
	return s
}

func (s *BatchSendMailRequest) SetClickTrace(v string) *BatchSendMailRequest {
	s.ClickTrace = &v
	return s
}

func (s *BatchSendMailRequest) SetHeaders(v string) *BatchSendMailRequest {
	s.Headers = &v
	return s
}

func (s *BatchSendMailRequest) SetIpPoolId(v string) *BatchSendMailRequest {
	s.IpPoolId = &v
	return s
}

func (s *BatchSendMailRequest) SetOwnerId(v int64) *BatchSendMailRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchSendMailRequest) SetReceiversName(v string) *BatchSendMailRequest {
	s.ReceiversName = &v
	return s
}

func (s *BatchSendMailRequest) SetReplyAddress(v string) *BatchSendMailRequest {
	s.ReplyAddress = &v
	return s
}

func (s *BatchSendMailRequest) SetReplyAddressAlias(v string) *BatchSendMailRequest {
	s.ReplyAddressAlias = &v
	return s
}

func (s *BatchSendMailRequest) SetResourceOwnerAccount(v string) *BatchSendMailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BatchSendMailRequest) SetResourceOwnerId(v int64) *BatchSendMailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BatchSendMailRequest) SetTagName(v string) *BatchSendMailRequest {
	s.TagName = &v
	return s
}

func (s *BatchSendMailRequest) SetTemplateName(v string) *BatchSendMailRequest {
	s.TemplateName = &v
	return s
}

func (s *BatchSendMailRequest) SetUnSubscribeFilterLevel(v string) *BatchSendMailRequest {
	s.UnSubscribeFilterLevel = &v
	return s
}

func (s *BatchSendMailRequest) SetUnSubscribeLinkType(v string) *BatchSendMailRequest {
	s.UnSubscribeLinkType = &v
	return s
}

func (s *BatchSendMailRequest) Validate() error {
	return dara.Validate(s)
}
