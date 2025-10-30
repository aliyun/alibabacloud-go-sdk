// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveCustomPrivacyPoliciesFromBrandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBrandId(v string) *RemoveCustomPrivacyPoliciesFromBrandRequest
	GetBrandId() *string
	SetCustomPrivacyPolicyIds(v []*string) *RemoveCustomPrivacyPoliciesFromBrandRequest
	GetCustomPrivacyPolicyIds() []*string
	SetInstanceId(v string) *RemoveCustomPrivacyPoliciesFromBrandRequest
	GetInstanceId() *string
}

type RemoveCustomPrivacyPoliciesFromBrandRequest struct {
	// 品牌化Id
	//
	// This parameter is required.
	//
	// example:
	//
	// brand_xxxx
	BrandId *string `json:"BrandId,omitempty" xml:"BrandId,omitempty"`
	// 条款ID列表
	//
	// This parameter is required.
	CustomPrivacyPolicyIds []*string `json:"CustomPrivacyPolicyIds,omitempty" xml:"CustomPrivacyPolicyIds,omitempty" type:"Repeated"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s RemoveCustomPrivacyPoliciesFromBrandRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveCustomPrivacyPoliciesFromBrandRequest) GoString() string {
	return s.String()
}

func (s *RemoveCustomPrivacyPoliciesFromBrandRequest) GetBrandId() *string {
	return s.BrandId
}

func (s *RemoveCustomPrivacyPoliciesFromBrandRequest) GetCustomPrivacyPolicyIds() []*string {
	return s.CustomPrivacyPolicyIds
}

func (s *RemoveCustomPrivacyPoliciesFromBrandRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RemoveCustomPrivacyPoliciesFromBrandRequest) SetBrandId(v string) *RemoveCustomPrivacyPoliciesFromBrandRequest {
	s.BrandId = &v
	return s
}

func (s *RemoveCustomPrivacyPoliciesFromBrandRequest) SetCustomPrivacyPolicyIds(v []*string) *RemoveCustomPrivacyPoliciesFromBrandRequest {
	s.CustomPrivacyPolicyIds = v
	return s
}

func (s *RemoveCustomPrivacyPoliciesFromBrandRequest) SetInstanceId(v string) *RemoveCustomPrivacyPoliciesFromBrandRequest {
	s.InstanceId = &v
	return s
}

func (s *RemoveCustomPrivacyPoliciesFromBrandRequest) Validate() error {
	return dara.Validate(s)
}
