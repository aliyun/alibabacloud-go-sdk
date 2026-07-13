// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProtectionPolicyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBoundResourceCategoryIdsShrink(v string) *UpdateProtectionPolicyShrinkRequest
	GetBoundResourceCategoryIdsShrink() *string
	SetClientToken(v string) *UpdateProtectionPolicyShrinkRequest
	GetClientToken() *string
	SetProtectionPolicyName(v string) *UpdateProtectionPolicyShrinkRequest
	GetProtectionPolicyName() *string
	SetSubProtectionPoliciesShrink(v string) *UpdateProtectionPolicyShrinkRequest
	GetSubProtectionPoliciesShrink() *string
}

type UpdateProtectionPolicyShrinkRequest struct {
	// The list of associated resource category IDs.
	BoundResourceCategoryIdsShrink *string `json:"BoundResourceCategoryIds,omitempty" xml:"BoundResourceCategoryIds,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters. If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// a1b2c3d4-****-****-****-a1b2c3d4f5e6
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The protection policy name.
	//
	// example:
	//
	// My***Policy
	ProtectionPolicyName *string `json:"ProtectionPolicyName,omitempty" xml:"ProtectionPolicyName,omitempty"`
	// The list of enabled sub-protection policies.
	SubProtectionPoliciesShrink *string `json:"SubProtectionPolicies,omitempty" xml:"SubProtectionPolicies,omitempty"`
}

func (s UpdateProtectionPolicyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProtectionPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateProtectionPolicyShrinkRequest) GetBoundResourceCategoryIdsShrink() *string {
	return s.BoundResourceCategoryIdsShrink
}

func (s *UpdateProtectionPolicyShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateProtectionPolicyShrinkRequest) GetProtectionPolicyName() *string {
	return s.ProtectionPolicyName
}

func (s *UpdateProtectionPolicyShrinkRequest) GetSubProtectionPoliciesShrink() *string {
	return s.SubProtectionPoliciesShrink
}

func (s *UpdateProtectionPolicyShrinkRequest) SetBoundResourceCategoryIdsShrink(v string) *UpdateProtectionPolicyShrinkRequest {
	s.BoundResourceCategoryIdsShrink = &v
	return s
}

func (s *UpdateProtectionPolicyShrinkRequest) SetClientToken(v string) *UpdateProtectionPolicyShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateProtectionPolicyShrinkRequest) SetProtectionPolicyName(v string) *UpdateProtectionPolicyShrinkRequest {
	s.ProtectionPolicyName = &v
	return s
}

func (s *UpdateProtectionPolicyShrinkRequest) SetSubProtectionPoliciesShrink(v string) *UpdateProtectionPolicyShrinkRequest {
	s.SubProtectionPoliciesShrink = &v
	return s
}

func (s *UpdateProtectionPolicyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
