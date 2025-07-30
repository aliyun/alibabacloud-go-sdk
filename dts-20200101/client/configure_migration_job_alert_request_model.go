// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureMigrationJobAlertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *ConfigureMigrationJobAlertRequest
	GetAccountId() *string
	SetDelayAlertPhone(v string) *ConfigureMigrationJobAlertRequest
	GetDelayAlertPhone() *string
	SetDelayAlertStatus(v string) *ConfigureMigrationJobAlertRequest
	GetDelayAlertStatus() *string
	SetDelayOverSeconds(v string) *ConfigureMigrationJobAlertRequest
	GetDelayOverSeconds() *string
	SetErrorAlertPhone(v string) *ConfigureMigrationJobAlertRequest
	GetErrorAlertPhone() *string
	SetErrorAlertStatus(v string) *ConfigureMigrationJobAlertRequest
	GetErrorAlertStatus() *string
	SetMigrationJobId(v string) *ConfigureMigrationJobAlertRequest
	GetMigrationJobId() *string
	SetOwnerId(v string) *ConfigureMigrationJobAlertRequest
	GetOwnerId() *string
	SetRegionId(v string) *ConfigureMigrationJobAlertRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ConfigureMigrationJobAlertRequest
	GetResourceGroupId() *string
}

type ConfigureMigrationJobAlertRequest struct {
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
	// The ID of the data migration instance. You can call the **DescribeMigrationJobs*	- operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dtslb9113qq11n****
	MigrationJobId *string `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
	OwnerId        *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the data migration instance resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
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
}

func (s ConfigureMigrationJobAlertRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigureMigrationJobAlertRequest) GoString() string {
	return s.String()
}

func (s *ConfigureMigrationJobAlertRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *ConfigureMigrationJobAlertRequest) GetDelayAlertPhone() *string {
	return s.DelayAlertPhone
}

func (s *ConfigureMigrationJobAlertRequest) GetDelayAlertStatus() *string {
	return s.DelayAlertStatus
}

func (s *ConfigureMigrationJobAlertRequest) GetDelayOverSeconds() *string {
	return s.DelayOverSeconds
}

func (s *ConfigureMigrationJobAlertRequest) GetErrorAlertPhone() *string {
	return s.ErrorAlertPhone
}

func (s *ConfigureMigrationJobAlertRequest) GetErrorAlertStatus() *string {
	return s.ErrorAlertStatus
}

func (s *ConfigureMigrationJobAlertRequest) GetMigrationJobId() *string {
	return s.MigrationJobId
}

func (s *ConfigureMigrationJobAlertRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ConfigureMigrationJobAlertRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ConfigureMigrationJobAlertRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ConfigureMigrationJobAlertRequest) SetAccountId(v string) *ConfigureMigrationJobAlertRequest {
	s.AccountId = &v
	return s
}

func (s *ConfigureMigrationJobAlertRequest) SetDelayAlertPhone(v string) *ConfigureMigrationJobAlertRequest {
	s.DelayAlertPhone = &v
	return s
}

func (s *ConfigureMigrationJobAlertRequest) SetDelayAlertStatus(v string) *ConfigureMigrationJobAlertRequest {
	s.DelayAlertStatus = &v
	return s
}

func (s *ConfigureMigrationJobAlertRequest) SetDelayOverSeconds(v string) *ConfigureMigrationJobAlertRequest {
	s.DelayOverSeconds = &v
	return s
}

func (s *ConfigureMigrationJobAlertRequest) SetErrorAlertPhone(v string) *ConfigureMigrationJobAlertRequest {
	s.ErrorAlertPhone = &v
	return s
}

func (s *ConfigureMigrationJobAlertRequest) SetErrorAlertStatus(v string) *ConfigureMigrationJobAlertRequest {
	s.ErrorAlertStatus = &v
	return s
}

func (s *ConfigureMigrationJobAlertRequest) SetMigrationJobId(v string) *ConfigureMigrationJobAlertRequest {
	s.MigrationJobId = &v
	return s
}

func (s *ConfigureMigrationJobAlertRequest) SetOwnerId(v string) *ConfigureMigrationJobAlertRequest {
	s.OwnerId = &v
	return s
}

func (s *ConfigureMigrationJobAlertRequest) SetRegionId(v string) *ConfigureMigrationJobAlertRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigureMigrationJobAlertRequest) SetResourceGroupId(v string) *ConfigureMigrationJobAlertRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ConfigureMigrationJobAlertRequest) Validate() error {
	return dara.Validate(s)
}
