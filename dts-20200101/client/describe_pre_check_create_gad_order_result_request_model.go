// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePreCheckCreateGadOrderResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribePreCheckCreateGadOrderResultRequest
	GetInstanceId() *string
	SetOwnerId(v string) *DescribePreCheckCreateGadOrderResultRequest
	GetOwnerId() *string
	SetRegionId(v string) *DescribePreCheckCreateGadOrderResultRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribePreCheckCreateGadOrderResultRequest
	GetResourceGroupId() *string
	SetTaskId(v string) *DescribePreCheckCreateGadOrderResultRequest
	GetTaskId() *string
}

type DescribePreCheckCreateGadOrderResultRequest struct {
	// example:
	//
	// gad-bp162d4tp0500****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId    *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// k71r16fj13g****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribePreCheckCreateGadOrderResultRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePreCheckCreateGadOrderResultRequest) GoString() string {
	return s.String()
}

func (s *DescribePreCheckCreateGadOrderResultRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribePreCheckCreateGadOrderResultRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribePreCheckCreateGadOrderResultRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePreCheckCreateGadOrderResultRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribePreCheckCreateGadOrderResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribePreCheckCreateGadOrderResultRequest) SetInstanceId(v string) *DescribePreCheckCreateGadOrderResultRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribePreCheckCreateGadOrderResultRequest) SetOwnerId(v string) *DescribePreCheckCreateGadOrderResultRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePreCheckCreateGadOrderResultRequest) SetRegionId(v string) *DescribePreCheckCreateGadOrderResultRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePreCheckCreateGadOrderResultRequest) SetResourceGroupId(v string) *DescribePreCheckCreateGadOrderResultRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribePreCheckCreateGadOrderResultRequest) SetTaskId(v string) *DescribePreCheckCreateGadOrderResultRequest {
	s.TaskId = &v
	return s
}

func (s *DescribePreCheckCreateGadOrderResultRequest) Validate() error {
	return dara.Validate(s)
}
