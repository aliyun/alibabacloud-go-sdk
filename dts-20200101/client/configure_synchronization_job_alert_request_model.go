// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureSynchronizationJobAlertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *ConfigureSynchronizationJobAlertRequest
	GetAccountId() *string
	SetDelayAlertPhone(v string) *ConfigureSynchronizationJobAlertRequest
	GetDelayAlertPhone() *string
	SetDelayAlertStatus(v string) *ConfigureSynchronizationJobAlertRequest
	GetDelayAlertStatus() *string
	SetDelayOverSeconds(v string) *ConfigureSynchronizationJobAlertRequest
	GetDelayOverSeconds() *string
	SetErrorAlertPhone(v string) *ConfigureSynchronizationJobAlertRequest
	GetErrorAlertPhone() *string
	SetErrorAlertStatus(v string) *ConfigureSynchronizationJobAlertRequest
	GetErrorAlertStatus() *string
	SetOwnerId(v string) *ConfigureSynchronizationJobAlertRequest
	GetOwnerId() *string
	SetRegionId(v string) *ConfigureSynchronizationJobAlertRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ConfigureSynchronizationJobAlertRequest
	GetResourceGroupId() *string
	SetSynchronizationDirection(v string) *ConfigureSynchronizationJobAlertRequest
	GetSynchronizationDirection() *string
	SetSynchronizationJobId(v string) *ConfigureSynchronizationJobAlertRequest
	GetSynchronizationJobId() *string
}

type ConfigureSynchronizationJobAlertRequest struct {
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because it will be deprecated.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The mobile phone numbers of contacts for latency alerts. Separate multiple phone numbers with commas (,).
	//
	// Note
	//
	// This parameter is supported only on the China site (aliyun.com) and only Chinese mainland phone numbers are supported. You can specify up to 10 phone numbers.
	//
	// China site (Chinese mainland) does not support phone alerts on the China site. You can configure alert rules for DTS tasks only in the CloudMonitor console.
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
	// The threshold for triggering a latency alert. Unit: seconds. The value must be an integer. Set the threshold based on your business requirements. We recommend that you set the threshold to 10 seconds or more to avoid alert fluctuations caused by network issues or database loads.
	//
	// > This parameter is required when **DelayAlertStatus*	- is set to **enable**.
	//
	// example:
	//
	// 10
	DelayOverSeconds *string `json:"DelayOverSeconds,omitempty" xml:"DelayOverSeconds,omitempty"`
	// The mobile phone numbers of contacts for exception alerts. Separate multiple phone numbers with commas (,).
	//
	// Note
	//
	// This parameter is supported only on the China site (aliyun.com) and only Chinese mainland phone numbers are supported. You can specify up to 10 phone numbers.
	//
	// The China site does not support phone alerts on the international site (alibabacloud.com). You can configure alert rules for DTS tasks only in the CloudMonitor console.
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
	// The region ID. Specify this parameter to indicate the region where the instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
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
	// The synchronization direction. Valid values:
	//
	// - **Forward**: forward.
	//
	// - **Reverse**: reverse.
	//
	// > Default value: **Forward**.
	//
	// example:
	//
	// Forward
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	// Instance ID of the data synchronization instance. You can call the DescribeSynchronizationJobs operation to query instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dtskxz1170c10p****
	SynchronizationJobId *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
}

func (s ConfigureSynchronizationJobAlertRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigureSynchronizationJobAlertRequest) GoString() string {
	return s.String()
}

func (s *ConfigureSynchronizationJobAlertRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *ConfigureSynchronizationJobAlertRequest) GetDelayAlertPhone() *string {
	return s.DelayAlertPhone
}

func (s *ConfigureSynchronizationJobAlertRequest) GetDelayAlertStatus() *string {
	return s.DelayAlertStatus
}

func (s *ConfigureSynchronizationJobAlertRequest) GetDelayOverSeconds() *string {
	return s.DelayOverSeconds
}

func (s *ConfigureSynchronizationJobAlertRequest) GetErrorAlertPhone() *string {
	return s.ErrorAlertPhone
}

func (s *ConfigureSynchronizationJobAlertRequest) GetErrorAlertStatus() *string {
	return s.ErrorAlertStatus
}

func (s *ConfigureSynchronizationJobAlertRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ConfigureSynchronizationJobAlertRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ConfigureSynchronizationJobAlertRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ConfigureSynchronizationJobAlertRequest) GetSynchronizationDirection() *string {
	return s.SynchronizationDirection
}

func (s *ConfigureSynchronizationJobAlertRequest) GetSynchronizationJobId() *string {
	return s.SynchronizationJobId
}

func (s *ConfigureSynchronizationJobAlertRequest) SetAccountId(v string) *ConfigureSynchronizationJobAlertRequest {
	s.AccountId = &v
	return s
}

func (s *ConfigureSynchronizationJobAlertRequest) SetDelayAlertPhone(v string) *ConfigureSynchronizationJobAlertRequest {
	s.DelayAlertPhone = &v
	return s
}

func (s *ConfigureSynchronizationJobAlertRequest) SetDelayAlertStatus(v string) *ConfigureSynchronizationJobAlertRequest {
	s.DelayAlertStatus = &v
	return s
}

func (s *ConfigureSynchronizationJobAlertRequest) SetDelayOverSeconds(v string) *ConfigureSynchronizationJobAlertRequest {
	s.DelayOverSeconds = &v
	return s
}

func (s *ConfigureSynchronizationJobAlertRequest) SetErrorAlertPhone(v string) *ConfigureSynchronizationJobAlertRequest {
	s.ErrorAlertPhone = &v
	return s
}

func (s *ConfigureSynchronizationJobAlertRequest) SetErrorAlertStatus(v string) *ConfigureSynchronizationJobAlertRequest {
	s.ErrorAlertStatus = &v
	return s
}

func (s *ConfigureSynchronizationJobAlertRequest) SetOwnerId(v string) *ConfigureSynchronizationJobAlertRequest {
	s.OwnerId = &v
	return s
}

func (s *ConfigureSynchronizationJobAlertRequest) SetRegionId(v string) *ConfigureSynchronizationJobAlertRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigureSynchronizationJobAlertRequest) SetResourceGroupId(v string) *ConfigureSynchronizationJobAlertRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ConfigureSynchronizationJobAlertRequest) SetSynchronizationDirection(v string) *ConfigureSynchronizationJobAlertRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *ConfigureSynchronizationJobAlertRequest) SetSynchronizationJobId(v string) *ConfigureSynchronizationJobAlertRequest {
	s.SynchronizationJobId = &v
	return s
}

func (s *ConfigureSynchronizationJobAlertRequest) Validate() error {
	return dara.Validate(s)
}
