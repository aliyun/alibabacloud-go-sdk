// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevertAggregateEvaluationResultsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *RevertAggregateEvaluationResultsRequest
	GetAggregatorId() *string
	SetConfigRuleId(v string) *RevertAggregateEvaluationResultsRequest
	GetConfigRuleId() *string
	SetResources(v []*RevertAggregateEvaluationResultsRequestResources) *RevertAggregateEvaluationResultsRequest
	GetResources() []*RevertAggregateEvaluationResultsRequestResources
}

type RevertAggregateEvaluationResultsRequest struct {
	// This parameter is required.
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// This parameter is required.
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// This parameter is required.
	Resources []*RevertAggregateEvaluationResultsRequestResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s RevertAggregateEvaluationResultsRequest) String() string {
	return dara.Prettify(s)
}

func (s RevertAggregateEvaluationResultsRequest) GoString() string {
	return s.String()
}

func (s *RevertAggregateEvaluationResultsRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *RevertAggregateEvaluationResultsRequest) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *RevertAggregateEvaluationResultsRequest) GetResources() []*RevertAggregateEvaluationResultsRequestResources {
	return s.Resources
}

func (s *RevertAggregateEvaluationResultsRequest) SetAggregatorId(v string) *RevertAggregateEvaluationResultsRequest {
	s.AggregatorId = &v
	return s
}

func (s *RevertAggregateEvaluationResultsRequest) SetConfigRuleId(v string) *RevertAggregateEvaluationResultsRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *RevertAggregateEvaluationResultsRequest) SetResources(v []*RevertAggregateEvaluationResultsRequestResources) *RevertAggregateEvaluationResultsRequest {
	s.Resources = v
	return s
}

func (s *RevertAggregateEvaluationResultsRequest) Validate() error {
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

type RevertAggregateEvaluationResultsRequestResources struct {
	// This parameter is required.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// This parameter is required.
	ResourceAccountId *int64 `json:"ResourceAccountId,omitempty" xml:"ResourceAccountId,omitempty"`
	// This parameter is required.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// This parameter is required.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s RevertAggregateEvaluationResultsRequestResources) String() string {
	return dara.Prettify(s)
}

func (s RevertAggregateEvaluationResultsRequestResources) GoString() string {
	return s.String()
}

func (s *RevertAggregateEvaluationResultsRequestResources) GetRegion() *string {
	return s.Region
}

func (s *RevertAggregateEvaluationResultsRequestResources) GetResourceAccountId() *int64 {
	return s.ResourceAccountId
}

func (s *RevertAggregateEvaluationResultsRequestResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *RevertAggregateEvaluationResultsRequestResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *RevertAggregateEvaluationResultsRequestResources) SetRegion(v string) *RevertAggregateEvaluationResultsRequestResources {
	s.Region = &v
	return s
}

func (s *RevertAggregateEvaluationResultsRequestResources) SetResourceAccountId(v int64) *RevertAggregateEvaluationResultsRequestResources {
	s.ResourceAccountId = &v
	return s
}

func (s *RevertAggregateEvaluationResultsRequestResources) SetResourceId(v string) *RevertAggregateEvaluationResultsRequestResources {
	s.ResourceId = &v
	return s
}

func (s *RevertAggregateEvaluationResultsRequestResources) SetResourceType(v string) *RevertAggregateEvaluationResultsRequestResources {
	s.ResourceType = &v
	return s
}

func (s *RevertAggregateEvaluationResultsRequestResources) Validate() error {
	return dara.Validate(s)
}
