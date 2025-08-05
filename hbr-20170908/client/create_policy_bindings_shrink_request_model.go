// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicyBindingsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyBindingListShrink(v string) *CreatePolicyBindingsShrinkRequest
	GetPolicyBindingListShrink() *string
	SetPolicyId(v string) *CreatePolicyBindingsShrinkRequest
	GetPolicyId() *string
}

type CreatePolicyBindingsShrinkRequest struct {
	// The data sources that you want to bind to the backup policy.
	PolicyBindingListShrink *string `json:"PolicyBindingList,omitempty" xml:"PolicyBindingList,omitempty"`
	// The ID of the backup policy.
	//
	// example:
	//
	// po-000************8ep
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s CreatePolicyBindingsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyBindingsShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePolicyBindingsShrinkRequest) GetPolicyBindingListShrink() *string {
	return s.PolicyBindingListShrink
}

func (s *CreatePolicyBindingsShrinkRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *CreatePolicyBindingsShrinkRequest) SetPolicyBindingListShrink(v string) *CreatePolicyBindingsShrinkRequest {
	s.PolicyBindingListShrink = &v
	return s
}

func (s *CreatePolicyBindingsShrinkRequest) SetPolicyId(v string) *CreatePolicyBindingsShrinkRequest {
	s.PolicyId = &v
	return s
}

func (s *CreatePolicyBindingsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
