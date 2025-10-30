// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIntegrationPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFeePackage(v string) *UpdateIntegrationPolicyRequest
	GetFeePackage() *string
	SetPolicyName(v string) *UpdateIntegrationPolicyRequest
	GetPolicyName() *string
	SetResourceGroupId(v string) *UpdateIntegrationPolicyRequest
	GetResourceGroupId() *string
	SetTags(v []*UpdateIntegrationPolicyRequestTags) *UpdateIntegrationPolicyRequest
	GetTags() []*UpdateIntegrationPolicyRequestTags
}

type UpdateIntegrationPolicyRequest struct {
	// Fee package type, CS_Pro/CS_Basic/empty.
	//
	// example:
	//
	// CS_Pro
	FeePackage *string `json:"feePackage,omitempty" xml:"feePackage,omitempty"`
	// Rule name, minimum 3 characters, maximum 63 characters, must start with a letter.
	//
	// example:
	//
	// metrics-inner-manage
	PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty"`
	// Resource group ID of the instance.
	//
	// example:
	//
	// rg-aekzoiafjtr7zyq
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// Resource tags.
	Tags []*UpdateIntegrationPolicyRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s UpdateIntegrationPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateIntegrationPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateIntegrationPolicyRequest) GetFeePackage() *string {
	return s.FeePackage
}

func (s *UpdateIntegrationPolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *UpdateIntegrationPolicyRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateIntegrationPolicyRequest) GetTags() []*UpdateIntegrationPolicyRequestTags {
	return s.Tags
}

func (s *UpdateIntegrationPolicyRequest) SetFeePackage(v string) *UpdateIntegrationPolicyRequest {
	s.FeePackage = &v
	return s
}

func (s *UpdateIntegrationPolicyRequest) SetPolicyName(v string) *UpdateIntegrationPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *UpdateIntegrationPolicyRequest) SetResourceGroupId(v string) *UpdateIntegrationPolicyRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateIntegrationPolicyRequest) SetTags(v []*UpdateIntegrationPolicyRequestTags) *UpdateIntegrationPolicyRequest {
	s.Tags = v
	return s
}

func (s *UpdateIntegrationPolicyRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateIntegrationPolicyRequestTags struct {
	// Tag `key` value.
	//
	// example:
	//
	// algo_bhv_expose_in_airec_exposure
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// Tag `value` value.
	//
	// example:
	//
	// [{\\"Id\\": \\"kgqie6hm\\", \\"Name\\": \\"Sheet1\\"}]
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdateIntegrationPolicyRequestTags) String() string {
	return dara.Prettify(s)
}

func (s UpdateIntegrationPolicyRequestTags) GoString() string {
	return s.String()
}

func (s *UpdateIntegrationPolicyRequestTags) GetKey() *string {
	return s.Key
}

func (s *UpdateIntegrationPolicyRequestTags) GetValue() *string {
	return s.Value
}

func (s *UpdateIntegrationPolicyRequestTags) SetKey(v string) *UpdateIntegrationPolicyRequestTags {
	s.Key = &v
	return s
}

func (s *UpdateIntegrationPolicyRequestTags) SetValue(v string) *UpdateIntegrationPolicyRequestTags {
	s.Value = &v
	return s
}

func (s *UpdateIntegrationPolicyRequestTags) Validate() error {
	return dara.Validate(s)
}
