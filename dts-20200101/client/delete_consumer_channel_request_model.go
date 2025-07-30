// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConsumerChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerGroupId(v string) *DeleteConsumerChannelRequest
	GetConsumerGroupId() *string
	SetDtsInstanceId(v string) *DeleteConsumerChannelRequest
	GetDtsInstanceId() *string
	SetDtsJobId(v string) *DeleteConsumerChannelRequest
	GetDtsJobId() *string
	SetRegionId(v string) *DeleteConsumerChannelRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DeleteConsumerChannelRequest
	GetResourceGroupId() *string
}

type DeleteConsumerChannelRequest struct {
	// The ID of the consumer group. You can call the [DescribeConsumerChannel](https://help.aliyun.com/document_detail/264169.html) operation to query the consumer group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dtsktbb6jdn2******
	ConsumerGroupId *string `json:"ConsumerGroupId,omitempty" xml:"ConsumerGroupId,omitempty"`
	// The ID of the change tracking instance. You can call the [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) operation to query the instance ID.
	//
	// >  You must specify at least one of the **DtsInstanceId*	- and **DtsJobId*	- parameters.
	//
	// example:
	//
	// dtsboss6pn1w73****
	DtsInstanceId *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	// The ID of the change tracking task. You can call the [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) operation to query the task ID.
	//
	// >  You must specify at least one of the **DtsInstanceId*	- and **DtsJobId*	- parameters.
	//
	// example:
	//
	// boss6pn1w73****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The ID of the region where the change tracking instance resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DeleteConsumerChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteConsumerChannelRequest) GoString() string {
	return s.String()
}

func (s *DeleteConsumerChannelRequest) GetConsumerGroupId() *string {
	return s.ConsumerGroupId
}

func (s *DeleteConsumerChannelRequest) GetDtsInstanceId() *string {
	return s.DtsInstanceId
}

func (s *DeleteConsumerChannelRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *DeleteConsumerChannelRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteConsumerChannelRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeleteConsumerChannelRequest) SetConsumerGroupId(v string) *DeleteConsumerChannelRequest {
	s.ConsumerGroupId = &v
	return s
}

func (s *DeleteConsumerChannelRequest) SetDtsInstanceId(v string) *DeleteConsumerChannelRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *DeleteConsumerChannelRequest) SetDtsJobId(v string) *DeleteConsumerChannelRequest {
	s.DtsJobId = &v
	return s
}

func (s *DeleteConsumerChannelRequest) SetRegionId(v string) *DeleteConsumerChannelRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteConsumerChannelRequest) SetResourceGroupId(v string) *DeleteConsumerChannelRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteConsumerChannelRequest) Validate() error {
	return dara.Validate(s)
}
