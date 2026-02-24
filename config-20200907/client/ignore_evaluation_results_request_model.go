// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIgnoreEvaluationResultsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigRuleId(v string) *IgnoreEvaluationResultsRequest
	GetConfigRuleId() *string
	SetIgnoreDate(v string) *IgnoreEvaluationResultsRequest
	GetIgnoreDate() *string
	SetReason(v string) *IgnoreEvaluationResultsRequest
	GetReason() *string
	SetResources(v []*IgnoreEvaluationResultsRequestResources) *IgnoreEvaluationResultsRequest
	GetResources() []*IgnoreEvaluationResultsRequestResources
}

type IgnoreEvaluationResultsRequest struct {
	// The rule ID.
	//
	// For more information about how to obtain the rule ID, see [ListConfigRules](https://help.aliyun.com/document_detail/169607.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cr-7e72626622af0051****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The date on which the ignored evaluation results are automatically restored.
	//
	// > If this parameter is left empty, the ignored evaluation results are not automatically restored. You must manually restore them.
	//
	// example:
	//
	// 2022-06-01
	IgnoreDate *string `json:"IgnoreDate,omitempty" xml:"IgnoreDate,omitempty"`
	// The reason for ignoring the resources.
	//
	// example:
	//
	// 无需检测。
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The list of resources to be ignored.
	//
	// This parameter is required.
	Resources []*IgnoreEvaluationResultsRequestResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s IgnoreEvaluationResultsRequest) String() string {
	return dara.Prettify(s)
}

func (s IgnoreEvaluationResultsRequest) GoString() string {
	return s.String()
}

func (s *IgnoreEvaluationResultsRequest) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *IgnoreEvaluationResultsRequest) GetIgnoreDate() *string {
	return s.IgnoreDate
}

func (s *IgnoreEvaluationResultsRequest) GetReason() *string {
	return s.Reason
}

func (s *IgnoreEvaluationResultsRequest) GetResources() []*IgnoreEvaluationResultsRequestResources {
	return s.Resources
}

func (s *IgnoreEvaluationResultsRequest) SetConfigRuleId(v string) *IgnoreEvaluationResultsRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *IgnoreEvaluationResultsRequest) SetIgnoreDate(v string) *IgnoreEvaluationResultsRequest {
	s.IgnoreDate = &v
	return s
}

func (s *IgnoreEvaluationResultsRequest) SetReason(v string) *IgnoreEvaluationResultsRequest {
	s.Reason = &v
	return s
}

func (s *IgnoreEvaluationResultsRequest) SetResources(v []*IgnoreEvaluationResultsRequestResources) *IgnoreEvaluationResultsRequest {
	s.Resources = v
	return s
}

func (s *IgnoreEvaluationResultsRequest) Validate() error {
	if s.Resources != nil {
		for _, item := range s.Resources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type IgnoreEvaluationResultsRequestResources struct {
	// The ID of the region to which the resource belongs.
	//
	// For more information about how to obtain the ID of the region to which the resource belongs, see [ListDiscoveredResources](https://help.aliyun.com/document_detail/169620.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the Alibaba Cloud account to which the resource belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100931896542****
	ResourceAccountId *int64 `json:"ResourceAccountId,omitempty" xml:"ResourceAccountId,omitempty"`
	// The resource ID.
	//
	// For more information about how to obtain the resource ID, see [ListDiscoveredResources](https://help.aliyun.com/document_detail/169620.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-hp3a3b4ztyfm2plgm****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type.
	//
	// For more information about how to obtain the resource type, see [ListDiscoveredResources](https://help.aliyun.com/document_detail/169620.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ACS::SLB::LoadBalancer
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s IgnoreEvaluationResultsRequestResources) String() string {
	return dara.Prettify(s)
}

func (s IgnoreEvaluationResultsRequestResources) GoString() string {
	return s.String()
}

func (s *IgnoreEvaluationResultsRequestResources) GetRegion() *string {
	return s.Region
}

func (s *IgnoreEvaluationResultsRequestResources) GetResourceAccountId() *int64 {
	return s.ResourceAccountId
}

func (s *IgnoreEvaluationResultsRequestResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *IgnoreEvaluationResultsRequestResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *IgnoreEvaluationResultsRequestResources) SetRegion(v string) *IgnoreEvaluationResultsRequestResources {
	s.Region = &v
	return s
}

func (s *IgnoreEvaluationResultsRequestResources) SetResourceAccountId(v int64) *IgnoreEvaluationResultsRequestResources {
	s.ResourceAccountId = &v
	return s
}

func (s *IgnoreEvaluationResultsRequestResources) SetResourceId(v string) *IgnoreEvaluationResultsRequestResources {
	s.ResourceId = &v
	return s
}

func (s *IgnoreEvaluationResultsRequestResources) SetResourceType(v string) *IgnoreEvaluationResultsRequestResources {
	s.ResourceType = &v
	return s
}

func (s *IgnoreEvaluationResultsRequestResources) Validate() error {
	return dara.Validate(s)
}
