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
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because this parameter will be removed in the future.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The mobile phone numbers that receive latency-related alerts. Separate mobile phone numbers with commas (,).
	//
	// >
	//
	// 	- This parameter is available only for China site (aliyun.com) users. Only mobile phone numbers in the Chinese mainland are supported. Up to 10 mobile phone numbers can be specified.
	//
	// 	- International site (alibabacloud.com) users cannot receive alerts by using mobile phones, but can [set alert rules for DTS tasks in the Cloud Monitor console](https://help.aliyun.com/document_detail/175876.html).
	//
	// example:
	//
	// 1361234****,1371234****
	DelayAlertPhone *string `json:"DelayAlertPhone,omitempty" xml:"DelayAlertPhone,omitempty"`
	// Specifies whether to monitor task latency. Valid values:
	//
	// 	- **enable**: yes
	//
	// 	- **disable**: no
	//
	// >
	//
	// 	- The default value is **enable**.
	//
	// 	- You must specify at least one of the DelayAlertStatus and **ErrorAlertStatus*	- parameters.
	//
	// example:
	//
	// enable
	DelayAlertStatus *string `json:"DelayAlertStatus,omitempty" xml:"DelayAlertStatus,omitempty"`
	// The threshold for triggering latency alerts. The unit is seconds and the value must be an integer. You can set the threshold based on your business needs. To avoid delay fluctuations caused by network and database loads, we recommend that you set the threshold to more than 10 seconds.
	//
	// >  If the **DelayAlertStatus*	- parameter is set to **enable**, this parameter must be specified.
	//
	// example:
	//
	// 10
	DelayOverSeconds *string `json:"DelayOverSeconds,omitempty" xml:"DelayOverSeconds,omitempty"`
	// The mobile phone numbers that receive status-related alerts. Separate mobile phone numbers with commas (,).
	//
	// >
	//
	// 	- This parameter is available only for China site (aliyun.com) users. Only mobile phone numbers in the Chinese mainland are supported. Up to 10 mobile phone numbers can be specified.
	//
	// 	- International site (alibabacloud.com) users cannot receive alerts by using mobile phones, but can [set alert rules for DTS tasks in the Cloud Monitor console](https://help.aliyun.com/document_detail/175876.html).
	//
	// example:
	//
	// 1361234****,1371234****
	ErrorAlertPhone *string `json:"ErrorAlertPhone,omitempty" xml:"ErrorAlertPhone,omitempty"`
	// Specifies whether to monitor task status. Valid values:
	//
	// 	- **enable**: yes
	//
	// 	- **disable**: no
	//
	// >
	//
	// 	- The default value is **enable**.
	//
	// 	- You must specify at least one of the **DelayAlertStatus*	- and ErrorAlertStatus parameters.
	//
	// 	- If the task that you monitor enters an abnormal state, an alert is triggered.
	//
	// example:
	//
	// enable
	ErrorAlertStatus *string `json:"ErrorAlertStatus,omitempty" xml:"ErrorAlertStatus,omitempty"`
	OwnerId          *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the change tracking instance resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group ID.
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
