// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppPolicyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppPolicyId(v string) *ModifyAppPolicyShrinkRequest
	GetAppPolicyId() *string
	SetProductType(v string) *ModifyAppPolicyShrinkRequest
	GetProductType() *string
	SetVideoPolicyShrink(v string) *ModifyAppPolicyShrinkRequest
	GetVideoPolicyShrink() *string
}

type ModifyAppPolicyShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pg-ee2znjktwgxu2****
	AppPolicyId *string `json:"AppPolicyId,omitempty" xml:"AppPolicyId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType       *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	VideoPolicyShrink *string `json:"VideoPolicy,omitempty" xml:"VideoPolicy,omitempty"`
}

func (s ModifyAppPolicyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppPolicyShrinkRequest) GetAppPolicyId() *string {
	return s.AppPolicyId
}

func (s *ModifyAppPolicyShrinkRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ModifyAppPolicyShrinkRequest) GetVideoPolicyShrink() *string {
	return s.VideoPolicyShrink
}

func (s *ModifyAppPolicyShrinkRequest) SetAppPolicyId(v string) *ModifyAppPolicyShrinkRequest {
	s.AppPolicyId = &v
	return s
}

func (s *ModifyAppPolicyShrinkRequest) SetProductType(v string) *ModifyAppPolicyShrinkRequest {
	s.ProductType = &v
	return s
}

func (s *ModifyAppPolicyShrinkRequest) SetVideoPolicyShrink(v string) *ModifyAppPolicyShrinkRequest {
	s.VideoPolicyShrink = &v
	return s
}

func (s *ModifyAppPolicyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
