// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureSubscriptionInstanceAlertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *ConfigureSubscriptionInstanceAlertRequest
	GetAccountId() *string
	SetDelayAlertPhone(v string) *ConfigureSubscriptionInstanceAlertRequest
	GetDelayAlertPhone() *string
	SetDelayAlertStatus(v string) *ConfigureSubscriptionInstanceAlertRequest
	GetDelayAlertStatus() *string
	SetDelayOverSeconds(v string) *ConfigureSubscriptionInstanceAlertRequest
	GetDelayOverSeconds() *string
	SetErrorAlertPhone(v string) *ConfigureSubscriptionInstanceAlertRequest
	GetErrorAlertPhone() *string
	SetErrorAlertStatus(v string) *ConfigureSubscriptionInstanceAlertRequest
	GetErrorAlertStatus() *string
	SetOwnerId(v string) *ConfigureSubscriptionInstanceAlertRequest
	GetOwnerId() *string
	SetRegionId(v string) *ConfigureSubscriptionInstanceAlertRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ConfigureSubscriptionInstanceAlertRequest
	GetResourceGroupId() *string
	SetSubscriptionInstanceId(v string) *ConfigureSubscriptionInstanceAlertRequest
	GetSubscriptionInstanceId() *string
}

type ConfigureSubscriptionInstanceAlertRequest struct {
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because it will be deprecated.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The mobile phone numbers of contacts for latency alerts. Separate multiple mobile phone numbers with commas (,).
	//
	// >- This parameter is supported only on the China site (aliyun.com). Only Chinese mainland mobile phone numbers are supported, and you can specify up to 10 mobile phone numbers.
	//
	// - China site (Chinese mainland) does not support Chinese mainland mobile phone alerts. You can only [configure alert rules for DTS tasks in CloudMonitor](https://help.aliyun.com/document_detail/175876.html).
	//
	// example:
	//
	// 1361234****,1371234****
	DelayAlertPhone *string `json:"DelayAlertPhone,omitempty" xml:"DelayAlertPhone,omitempty"`
	// Specifies whether to monitor the latency status. Valid values:
	//
	// - **enable**: yes.
	//
	// - **disable**: no.
	//
	// > - Default value: **enable**.
	//
	// - You must specify at least one of this parameter and the **ErrorAlertStatus*	- parameter.
	//
	// example:
	//
	// enable
	DelayAlertStatus *string `json:"DelayAlertStatus,omitempty" xml:"DelayAlertStatus,omitempty"`
	// The threshold for triggering a latency alert. Unit: seconds. The value must be an integer. Set the threshold based on your business requirements. To avoid alert fluctuations caused by network issues or database loads, set the threshold to 10 seconds or more.
	//
	// > This parameter is required when **DelayAlertStatus*	- is set to **enable**.
	//
	// example:
	//
	// 10
	DelayOverSeconds *string `json:"DelayOverSeconds,omitempty" xml:"DelayOverSeconds,omitempty"`
	// The mobile phone numbers of contacts for exception alerts. Separate multiple mobile phone numbers with commas (,).
	//
	// >- This parameter is supported only on the China site (aliyun.com). Only Chinese mainland mobile phone numbers are supported, and you can specify up to 10 mobile phone numbers.
	//
	// - China site (Chinese mainland) does not support Chinese mainland mobile phone alerts. You can only [configure alert rules for DTS tasks in CloudMonitor](https://help.aliyun.com/document_detail/175876.html).
	//
	// example:
	//
	// 1361234****,1371234****
	ErrorAlertPhone *string `json:"ErrorAlertPhone,omitempty" xml:"ErrorAlertPhone,omitempty"`
	// Specifies whether to monitor the exception status. Valid values:
	//
	// - **enable**: yes.
	//
	// - **disable**: no.
	//
	// > - Default value: **enable**.
	//
	// - You must specify at least one of this parameter and the **DelayAlertStatus*	- parameter.
	//
	// - After you enable exception status monitoring, an alert is triggered when an exception is detected.
	//
	// example:
	//
	// enable
	ErrorAlertStatus *string `json:"ErrorAlertStatus,omitempty" xml:"ErrorAlertStatus,omitempty"`
	OwnerId          *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the change tracking instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the change tracking instance. You can call the DescribeSubscriptionInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dtsl8zl9ek6292****
	SubscriptionInstanceId *string `json:"SubscriptionInstanceId,omitempty" xml:"SubscriptionInstanceId,omitempty"`
}

func (s ConfigureSubscriptionInstanceAlertRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigureSubscriptionInstanceAlertRequest) GoString() string {
	return s.String()
}

func (s *ConfigureSubscriptionInstanceAlertRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *ConfigureSubscriptionInstanceAlertRequest) GetDelayAlertPhone() *string {
	return s.DelayAlertPhone
}

func (s *ConfigureSubscriptionInstanceAlertRequest) GetDelayAlertStatus() *string {
	return s.DelayAlertStatus
}

func (s *ConfigureSubscriptionInstanceAlertRequest) GetDelayOverSeconds() *string {
	return s.DelayOverSeconds
}

func (s *ConfigureSubscriptionInstanceAlertRequest) GetErrorAlertPhone() *string {
	return s.ErrorAlertPhone
}

func (s *ConfigureSubscriptionInstanceAlertRequest) GetErrorAlertStatus() *string {
	return s.ErrorAlertStatus
}

func (s *ConfigureSubscriptionInstanceAlertRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ConfigureSubscriptionInstanceAlertRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ConfigureSubscriptionInstanceAlertRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ConfigureSubscriptionInstanceAlertRequest) GetSubscriptionInstanceId() *string {
	return s.SubscriptionInstanceId
}

func (s *ConfigureSubscriptionInstanceAlertRequest) SetAccountId(v string) *ConfigureSubscriptionInstanceAlertRequest {
	s.AccountId = &v
	return s
}

func (s *ConfigureSubscriptionInstanceAlertRequest) SetDelayAlertPhone(v string) *ConfigureSubscriptionInstanceAlertRequest {
	s.DelayAlertPhone = &v
	return s
}

func (s *ConfigureSubscriptionInstanceAlertRequest) SetDelayAlertStatus(v string) *ConfigureSubscriptionInstanceAlertRequest {
	s.DelayAlertStatus = &v
	return s
}

func (s *ConfigureSubscriptionInstanceAlertRequest) SetDelayOverSeconds(v string) *ConfigureSubscriptionInstanceAlertRequest {
	s.DelayOverSeconds = &v
	return s
}

func (s *ConfigureSubscriptionInstanceAlertRequest) SetErrorAlertPhone(v string) *ConfigureSubscriptionInstanceAlertRequest {
	s.ErrorAlertPhone = &v
	return s
}

func (s *ConfigureSubscriptionInstanceAlertRequest) SetErrorAlertStatus(v string) *ConfigureSubscriptionInstanceAlertRequest {
	s.ErrorAlertStatus = &v
	return s
}

func (s *ConfigureSubscriptionInstanceAlertRequest) SetOwnerId(v string) *ConfigureSubscriptionInstanceAlertRequest {
	s.OwnerId = &v
	return s
}

func (s *ConfigureSubscriptionInstanceAlertRequest) SetRegionId(v string) *ConfigureSubscriptionInstanceAlertRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigureSubscriptionInstanceAlertRequest) SetResourceGroupId(v string) *ConfigureSubscriptionInstanceAlertRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ConfigureSubscriptionInstanceAlertRequest) SetSubscriptionInstanceId(v string) *ConfigureSubscriptionInstanceAlertRequest {
	s.SubscriptionInstanceId = &v
	return s
}

func (s *ConfigureSubscriptionInstanceAlertRequest) Validate() error {
	return dara.Validate(s)
}
