// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDryRunConfigRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigurationItem(v string) *DryRunConfigRuleRequest
	GetConfigurationItem() *string
	SetResourceType(v string) *DryRunConfigRuleRequest
	GetResourceType() *string
}

type DryRunConfigRuleRequest struct {
	// The complete configuration information of the resource.
	//
	// example:
	//
	// {
	//
	//   "ResourceCreationTime": 1741241360000,
	//
	//   "AccountId": 123,
	//
	//   "Configuration": {
	//
	//     "ResourceGroupId": "",
	//
	//     "Memory": 1024,
	//
	//     "InstanceChargeType": "PrePaid",
	//
	//     "Cpu": 1,
	//
	//     "OSName": "Alibaba Cloud Linux  3.2104 LTS 64 bit",
	//
	//     "InstanceNetworkType": "vpc",
	//
	//     "InnerIpAddress": {
	//
	//       "IpAddress": []
	//
	//     },
	//
	//     "ExpiredTime": "2026-05-06T16:00Z",
	//
	//     "ImageId": "aliyun_3_x64_20G_alibase_20250117.vhd",
	//
	//     "EipAddress": {
	//
	//       "AllocationId": "",
	//
	//       "IpAddress": "",
	//
	//       "InternetChargeType": ""
	//
	//     },
	//
	//     "ImageOptions": {},
	//
	//     "Status": "Running",
	//
	//     "AdditionalInfo": {},
	//
	//     "HibernationOptions": {
	//
	//       "Configured": false
	//
	//     }
	//
	//   },
	//
	//   "ResourceId": "i-bp1d8kd8ztaynb4c****",
	//
	//   "ResourceName": "****",
	//
	//   "ResourceStatus": "Running",
	//
	//   "Region": "cn-hangzhou",
	//
	//   "AvailabilityZone": "cn-hangzhou-h",
	//
	//   "ResourceType": "ACS::ECS::Instance",
	//
	//   "ResourceDeleted": 1
	//
	// }
	ConfigurationItem *string `json:"ConfigurationItem,omitempty" xml:"ConfigurationItem,omitempty"`
	// The resource type that is evaluated by the rule.
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DryRunConfigRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DryRunConfigRuleRequest) GoString() string {
	return s.String()
}

func (s *DryRunConfigRuleRequest) GetConfigurationItem() *string {
	return s.ConfigurationItem
}

func (s *DryRunConfigRuleRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DryRunConfigRuleRequest) SetConfigurationItem(v string) *DryRunConfigRuleRequest {
	s.ConfigurationItem = &v
	return s
}

func (s *DryRunConfigRuleRequest) SetResourceType(v string) *DryRunConfigRuleRequest {
	s.ResourceType = &v
	return s
}

func (s *DryRunConfigRuleRequest) Validate() error {
	return dara.Validate(s)
}
