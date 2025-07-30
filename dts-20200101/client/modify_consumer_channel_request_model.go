// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyConsumerChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerGroupId(v string) *ModifyConsumerChannelRequest
	GetConsumerGroupId() *string
	SetConsumerGroupName(v string) *ModifyConsumerChannelRequest
	GetConsumerGroupName() *string
	SetConsumerGroupPassword(v string) *ModifyConsumerChannelRequest
	GetConsumerGroupPassword() *string
	SetConsumerGroupUserName(v string) *ModifyConsumerChannelRequest
	GetConsumerGroupUserName() *string
	SetDtsInstanceId(v string) *ModifyConsumerChannelRequest
	GetDtsInstanceId() *string
	SetDtsJobId(v string) *ModifyConsumerChannelRequest
	GetDtsJobId() *string
	SetRegionId(v string) *ModifyConsumerChannelRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ModifyConsumerChannelRequest
	GetResourceGroupId() *string
}

type ModifyConsumerChannelRequest struct {
	// The ID of the consumer group. You can call the [DescribeConsumerChannel](https://help.aliyun.com/document_detail/264169.html) operation to query the consumer group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dtsor2y66j4219****
	ConsumerGroupId *string `json:"ConsumerGroupId,omitempty" xml:"ConsumerGroupId,omitempty"`
	// The name of the consumer group. The name cannot exceed 128 characters in length. We recommend that you use an informative name for easy identification.
	//
	// example:
	//
	// dtstest
	ConsumerGroupName *string `json:"ConsumerGroupName,omitempty" xml:"ConsumerGroupName,omitempty"`
	// The new password of the consumer group.
	//
	// 	- A password must contain two or more of the following characters: uppercase letters, lowercase letters, digits, and special characters.
	//
	// 	- A password must be 8 to 32 characters in length.
	//
	// example:
	//
	// Test123456
	ConsumerGroupPassword *string `json:"ConsumerGroupPassword,omitempty" xml:"ConsumerGroupPassword,omitempty"`
	// The new username of the consumer group.
	//
	// 	- A username can contain one or more of the following character types: uppercase letters, lowercase letters, digits, and underscores (_).
	//
	// 	- A username cannot exceed 16 characters in length.
	//
	// example:
	//
	// dtstest
	ConsumerGroupUserName *string `json:"ConsumerGroupUserName,omitempty" xml:"ConsumerGroupUserName,omitempty"`
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
	// Resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ModifyConsumerChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyConsumerChannelRequest) GoString() string {
	return s.String()
}

func (s *ModifyConsumerChannelRequest) GetConsumerGroupId() *string {
	return s.ConsumerGroupId
}

func (s *ModifyConsumerChannelRequest) GetConsumerGroupName() *string {
	return s.ConsumerGroupName
}

func (s *ModifyConsumerChannelRequest) GetConsumerGroupPassword() *string {
	return s.ConsumerGroupPassword
}

func (s *ModifyConsumerChannelRequest) GetConsumerGroupUserName() *string {
	return s.ConsumerGroupUserName
}

func (s *ModifyConsumerChannelRequest) GetDtsInstanceId() *string {
	return s.DtsInstanceId
}

func (s *ModifyConsumerChannelRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *ModifyConsumerChannelRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyConsumerChannelRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyConsumerChannelRequest) SetConsumerGroupId(v string) *ModifyConsumerChannelRequest {
	s.ConsumerGroupId = &v
	return s
}

func (s *ModifyConsumerChannelRequest) SetConsumerGroupName(v string) *ModifyConsumerChannelRequest {
	s.ConsumerGroupName = &v
	return s
}

func (s *ModifyConsumerChannelRequest) SetConsumerGroupPassword(v string) *ModifyConsumerChannelRequest {
	s.ConsumerGroupPassword = &v
	return s
}

func (s *ModifyConsumerChannelRequest) SetConsumerGroupUserName(v string) *ModifyConsumerChannelRequest {
	s.ConsumerGroupUserName = &v
	return s
}

func (s *ModifyConsumerChannelRequest) SetDtsInstanceId(v string) *ModifyConsumerChannelRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *ModifyConsumerChannelRequest) SetDtsJobId(v string) *ModifyConsumerChannelRequest {
	s.DtsJobId = &v
	return s
}

func (s *ModifyConsumerChannelRequest) SetRegionId(v string) *ModifyConsumerChannelRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyConsumerChannelRequest) SetResourceGroupId(v string) *ModifyConsumerChannelRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyConsumerChannelRequest) Validate() error {
	return dara.Validate(s)
}
