// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableInstancePublicAccessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DisableInstancePublicAccessRequest
	GetInstanceId() *string
	SetRegionId(v string) *DisableInstancePublicAccessRequest
	GetRegionId() *string
}

type DisableInstancePublicAccessRequest struct {
	// The ID of the bastion host whose Internet access you want to disable.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-78v1gh****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DisableInstancePublicAccessRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableInstancePublicAccessRequest) GoString() string {
	return s.String()
}

func (s *DisableInstancePublicAccessRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableInstancePublicAccessRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DisableInstancePublicAccessRequest) SetInstanceId(v string) *DisableInstancePublicAccessRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableInstancePublicAccessRequest) SetRegionId(v string) *DisableInstancePublicAccessRequest {
	s.RegionId = &v
	return s
}

func (s *DisableInstancePublicAccessRequest) Validate() error {
	return dara.Validate(s)
}
