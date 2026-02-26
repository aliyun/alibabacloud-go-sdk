// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRemediationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigRuleId(v string) *StartRemediationRequest
	GetConfigRuleId() *string
	SetResourceId(v string) *StartRemediationRequest
	GetResourceId() *string
	SetResourceRegionId(v string) *StartRemediationRequest
	GetResourceRegionId() *string
	SetResourceType(v string) *StartRemediationRequest
	GetResourceType() *string
}

type StartRemediationRequest struct {
	// The rule ID.
	//
	// For information about how to obtain the rule ID, see [ListConfigRules](https://help.aliyun.com/document_detail/169607.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cr-8a973ac2e2be00a2****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// example:
	//
	// vpc-xxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// example:
	//
	// ACS::VPC::VPC
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s StartRemediationRequest) String() string {
	return dara.Prettify(s)
}

func (s StartRemediationRequest) GoString() string {
	return s.String()
}

func (s *StartRemediationRequest) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *StartRemediationRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *StartRemediationRequest) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *StartRemediationRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *StartRemediationRequest) SetConfigRuleId(v string) *StartRemediationRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *StartRemediationRequest) SetResourceId(v string) *StartRemediationRequest {
	s.ResourceId = &v
	return s
}

func (s *StartRemediationRequest) SetResourceRegionId(v string) *StartRemediationRequest {
	s.ResourceRegionId = &v
	return s
}

func (s *StartRemediationRequest) SetResourceType(v string) *StartRemediationRequest {
	s.ResourceType = &v
	return s
}

func (s *StartRemediationRequest) Validate() error {
	return dara.Validate(s)
}
