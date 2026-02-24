// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevertEvaluationResultsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigRuleId(v string) *RevertEvaluationResultsRequest
	GetConfigRuleId() *string
	SetResources(v []*RevertEvaluationResultsRequestResources) *RevertEvaluationResultsRequest
	GetResources() []*RevertEvaluationResultsRequestResources
}

type RevertEvaluationResultsRequest struct {
	// This parameter is required.
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// This parameter is required.
	Resources []*RevertEvaluationResultsRequestResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s RevertEvaluationResultsRequest) String() string {
	return dara.Prettify(s)
}

func (s RevertEvaluationResultsRequest) GoString() string {
	return s.String()
}

func (s *RevertEvaluationResultsRequest) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *RevertEvaluationResultsRequest) GetResources() []*RevertEvaluationResultsRequestResources {
	return s.Resources
}

func (s *RevertEvaluationResultsRequest) SetConfigRuleId(v string) *RevertEvaluationResultsRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *RevertEvaluationResultsRequest) SetResources(v []*RevertEvaluationResultsRequestResources) *RevertEvaluationResultsRequest {
	s.Resources = v
	return s
}

func (s *RevertEvaluationResultsRequest) Validate() error {
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

type RevertEvaluationResultsRequestResources struct {
	// This parameter is required.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// This parameter is required.
	ResourceAccountId *int64 `json:"ResourceAccountId,omitempty" xml:"ResourceAccountId,omitempty"`
	// This parameter is required.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// This parameter is required.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s RevertEvaluationResultsRequestResources) String() string {
	return dara.Prettify(s)
}

func (s RevertEvaluationResultsRequestResources) GoString() string {
	return s.String()
}

func (s *RevertEvaluationResultsRequestResources) GetRegion() *string {
	return s.Region
}

func (s *RevertEvaluationResultsRequestResources) GetResourceAccountId() *int64 {
	return s.ResourceAccountId
}

func (s *RevertEvaluationResultsRequestResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *RevertEvaluationResultsRequestResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *RevertEvaluationResultsRequestResources) SetRegion(v string) *RevertEvaluationResultsRequestResources {
	s.Region = &v
	return s
}

func (s *RevertEvaluationResultsRequestResources) SetResourceAccountId(v int64) *RevertEvaluationResultsRequestResources {
	s.ResourceAccountId = &v
	return s
}

func (s *RevertEvaluationResultsRequestResources) SetResourceId(v string) *RevertEvaluationResultsRequestResources {
	s.ResourceId = &v
	return s
}

func (s *RevertEvaluationResultsRequestResources) SetResourceType(v string) *RevertEvaluationResultsRequestResources {
	s.ResourceType = &v
	return s
}

func (s *RevertEvaluationResultsRequestResources) Validate() error {
	return dara.Validate(s)
}
