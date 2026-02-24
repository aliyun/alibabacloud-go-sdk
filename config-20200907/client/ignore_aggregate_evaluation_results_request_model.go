// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIgnoreAggregateEvaluationResultsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *IgnoreAggregateEvaluationResultsRequest
	GetAggregatorId() *string
	SetConfigRuleId(v string) *IgnoreAggregateEvaluationResultsRequest
	GetConfigRuleId() *string
	SetIgnoreDate(v string) *IgnoreAggregateEvaluationResultsRequest
	GetIgnoreDate() *string
	SetReason(v string) *IgnoreAggregateEvaluationResultsRequest
	GetReason() *string
	SetResources(v []*IgnoreAggregateEvaluationResultsRequestResources) *IgnoreAggregateEvaluationResultsRequest
	GetResources() []*IgnoreAggregateEvaluationResultsRequestResources
}

type IgnoreAggregateEvaluationResultsRequest struct {
	// This parameter is required.
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// This parameter is required.
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	IgnoreDate   *string `json:"IgnoreDate,omitempty" xml:"IgnoreDate,omitempty"`
	Reason       *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// This parameter is required.
	Resources []*IgnoreAggregateEvaluationResultsRequestResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s IgnoreAggregateEvaluationResultsRequest) String() string {
	return dara.Prettify(s)
}

func (s IgnoreAggregateEvaluationResultsRequest) GoString() string {
	return s.String()
}

func (s *IgnoreAggregateEvaluationResultsRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *IgnoreAggregateEvaluationResultsRequest) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *IgnoreAggregateEvaluationResultsRequest) GetIgnoreDate() *string {
	return s.IgnoreDate
}

func (s *IgnoreAggregateEvaluationResultsRequest) GetReason() *string {
	return s.Reason
}

func (s *IgnoreAggregateEvaluationResultsRequest) GetResources() []*IgnoreAggregateEvaluationResultsRequestResources {
	return s.Resources
}

func (s *IgnoreAggregateEvaluationResultsRequest) SetAggregatorId(v string) *IgnoreAggregateEvaluationResultsRequest {
	s.AggregatorId = &v
	return s
}

func (s *IgnoreAggregateEvaluationResultsRequest) SetConfigRuleId(v string) *IgnoreAggregateEvaluationResultsRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *IgnoreAggregateEvaluationResultsRequest) SetIgnoreDate(v string) *IgnoreAggregateEvaluationResultsRequest {
	s.IgnoreDate = &v
	return s
}

func (s *IgnoreAggregateEvaluationResultsRequest) SetReason(v string) *IgnoreAggregateEvaluationResultsRequest {
	s.Reason = &v
	return s
}

func (s *IgnoreAggregateEvaluationResultsRequest) SetResources(v []*IgnoreAggregateEvaluationResultsRequestResources) *IgnoreAggregateEvaluationResultsRequest {
	s.Resources = v
	return s
}

func (s *IgnoreAggregateEvaluationResultsRequest) Validate() error {
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

type IgnoreAggregateEvaluationResultsRequestResources struct {
	// This parameter is required.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// This parameter is required.
	ResourceAccountId *int64 `json:"ResourceAccountId,omitempty" xml:"ResourceAccountId,omitempty"`
	// This parameter is required.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// This parameter is required.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s IgnoreAggregateEvaluationResultsRequestResources) String() string {
	return dara.Prettify(s)
}

func (s IgnoreAggregateEvaluationResultsRequestResources) GoString() string {
	return s.String()
}

func (s *IgnoreAggregateEvaluationResultsRequestResources) GetRegion() *string {
	return s.Region
}

func (s *IgnoreAggregateEvaluationResultsRequestResources) GetResourceAccountId() *int64 {
	return s.ResourceAccountId
}

func (s *IgnoreAggregateEvaluationResultsRequestResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *IgnoreAggregateEvaluationResultsRequestResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *IgnoreAggregateEvaluationResultsRequestResources) SetRegion(v string) *IgnoreAggregateEvaluationResultsRequestResources {
	s.Region = &v
	return s
}

func (s *IgnoreAggregateEvaluationResultsRequestResources) SetResourceAccountId(v int64) *IgnoreAggregateEvaluationResultsRequestResources {
	s.ResourceAccountId = &v
	return s
}

func (s *IgnoreAggregateEvaluationResultsRequestResources) SetResourceId(v string) *IgnoreAggregateEvaluationResultsRequestResources {
	s.ResourceId = &v
	return s
}

func (s *IgnoreAggregateEvaluationResultsRequestResources) SetResourceType(v string) *IgnoreAggregateEvaluationResultsRequestResources {
	s.ResourceType = &v
	return s
}

func (s *IgnoreAggregateEvaluationResultsRequestResources) Validate() error {
	return dara.Validate(s)
}
