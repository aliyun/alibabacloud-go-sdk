// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProtectionPolicyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBoundResourceCategoryIdsShrink(v string) *CreateProtectionPolicyShrinkRequest
	GetBoundResourceCategoryIdsShrink() *string
	SetClientToken(v string) *CreateProtectionPolicyShrinkRequest
	GetClientToken() *string
	SetProtectionPolicyName(v string) *CreateProtectionPolicyShrinkRequest
	GetProtectionPolicyName() *string
	SetProtectionPolicyRegionId(v string) *CreateProtectionPolicyShrinkRequest
	GetProtectionPolicyRegionId() *string
	SetSubProtectionPoliciesShrink(v string) *CreateProtectionPolicyShrinkRequest
	GetSubProtectionPoliciesShrink() *string
}

type CreateProtectionPolicyShrinkRequest struct {
	// The IDs of associated resource categories.
	BoundResourceCategoryIdsShrink *string `json:"BoundResourceCategoryIds,omitempty" xml:"BoundResourceCategoryIds,omitempty"`
	// The client token used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// a1b2c3d4-****-****-****-a1b2c3d4f5e6
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The name of the protection policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// My***Policy
	ProtectionPolicyName *string `json:"ProtectionPolicyName,omitempty" xml:"ProtectionPolicyName,omitempty"`
	// The region ID of the protection policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	ProtectionPolicyRegionId *string `json:"ProtectionPolicyRegionId,omitempty" xml:"ProtectionPolicyRegionId,omitempty"`
	// The sub-protection policies.
	//
	// This parameter is required.
	SubProtectionPoliciesShrink *string `json:"SubProtectionPolicies,omitempty" xml:"SubProtectionPolicies,omitempty"`
}

func (s CreateProtectionPolicyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProtectionPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateProtectionPolicyShrinkRequest) GetBoundResourceCategoryIdsShrink() *string {
	return s.BoundResourceCategoryIdsShrink
}

func (s *CreateProtectionPolicyShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateProtectionPolicyShrinkRequest) GetProtectionPolicyName() *string {
	return s.ProtectionPolicyName
}

func (s *CreateProtectionPolicyShrinkRequest) GetProtectionPolicyRegionId() *string {
	return s.ProtectionPolicyRegionId
}

func (s *CreateProtectionPolicyShrinkRequest) GetSubProtectionPoliciesShrink() *string {
	return s.SubProtectionPoliciesShrink
}

func (s *CreateProtectionPolicyShrinkRequest) SetBoundResourceCategoryIdsShrink(v string) *CreateProtectionPolicyShrinkRequest {
	s.BoundResourceCategoryIdsShrink = &v
	return s
}

func (s *CreateProtectionPolicyShrinkRequest) SetClientToken(v string) *CreateProtectionPolicyShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateProtectionPolicyShrinkRequest) SetProtectionPolicyName(v string) *CreateProtectionPolicyShrinkRequest {
	s.ProtectionPolicyName = &v
	return s
}

func (s *CreateProtectionPolicyShrinkRequest) SetProtectionPolicyRegionId(v string) *CreateProtectionPolicyShrinkRequest {
	s.ProtectionPolicyRegionId = &v
	return s
}

func (s *CreateProtectionPolicyShrinkRequest) SetSubProtectionPoliciesShrink(v string) *CreateProtectionPolicyShrinkRequest {
	s.SubProtectionPoliciesShrink = &v
	return s
}

func (s *CreateProtectionPolicyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
