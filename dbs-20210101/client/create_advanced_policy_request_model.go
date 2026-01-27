// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAdvancedPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *CreateAdvancedPolicyRequest
	GetInstanceName() *string
	SetRegionCode(v string) *CreateAdvancedPolicyRequest
	GetRegionCode() *string
}

type CreateAdvancedPolicyRequest struct {
	// The ID of the PolarDB instance.
	//
	// example:
	//
	// pc-2ze3nrr64c5****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The region in which backup sets reside.
	//
	// example:
	//
	// cn-shanghai
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
}

func (s CreateAdvancedPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAdvancedPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateAdvancedPolicyRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateAdvancedPolicyRequest) GetRegionCode() *string {
	return s.RegionCode
}

func (s *CreateAdvancedPolicyRequest) SetInstanceName(v string) *CreateAdvancedPolicyRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateAdvancedPolicyRequest) SetRegionCode(v string) *CreateAdvancedPolicyRequest {
	s.RegionCode = &v
	return s
}

func (s *CreateAdvancedPolicyRequest) Validate() error {
	return dara.Validate(s)
}
