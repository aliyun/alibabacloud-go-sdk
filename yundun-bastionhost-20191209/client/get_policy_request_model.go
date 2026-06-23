// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetPolicyRequest
	GetInstanceId() *string
	SetPolicyId(v string) *GetPolicyRequest
	GetPolicyId() *string
	SetRegionId(v string) *GetPolicyRequest
	GetRegionId() *string
}

type GetPolicyRequest struct {
	// The ID of the bastion host instance.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to get this ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-zvp2d3syb0g
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the control policy that you want to query.
	//
	// > You can call the [ListPolicies](https://help.aliyun.com/document_detail/2758876.html) operation to get this ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The region ID of the bastion host instance.
	//
	// > For more information about region IDs, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetPolicyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetPolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *GetPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetPolicyRequest) SetInstanceId(v string) *GetPolicyRequest {
	s.InstanceId = &v
	return s
}

func (s *GetPolicyRequest) SetPolicyId(v string) *GetPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *GetPolicyRequest) SetRegionId(v string) *GetPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *GetPolicyRequest) Validate() error {
	return dara.Validate(s)
}
