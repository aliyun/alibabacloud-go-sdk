// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolicyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyIdListShrink(v string) *DeletePolicyShrinkRequest
	GetPolicyIdListShrink() *string
	SetRegionId(v string) *DeletePolicyShrinkRequest
	GetRegionId() *string
}

type DeletePolicyShrinkRequest struct {
	// List of detection policy IDs
	PolicyIdListShrink *string `json:"PolicyIdList,omitempty" xml:"PolicyIdList,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeletePolicyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeletePolicyShrinkRequest) GetPolicyIdListShrink() *string {
	return s.PolicyIdListShrink
}

func (s *DeletePolicyShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeletePolicyShrinkRequest) SetPolicyIdListShrink(v string) *DeletePolicyShrinkRequest {
	s.PolicyIdListShrink = &v
	return s
}

func (s *DeletePolicyShrinkRequest) SetRegionId(v string) *DeletePolicyShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DeletePolicyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
