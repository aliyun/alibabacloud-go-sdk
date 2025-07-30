// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConsumerChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerGroupName(v string) *CreateConsumerChannelRequest
	GetConsumerGroupName() *string
	SetConsumerGroupPassword(v string) *CreateConsumerChannelRequest
	GetConsumerGroupPassword() *string
	SetConsumerGroupUserName(v string) *CreateConsumerChannelRequest
	GetConsumerGroupUserName() *string
	SetDtsInstanceId(v string) *CreateConsumerChannelRequest
	GetDtsInstanceId() *string
	SetDtsJobId(v string) *CreateConsumerChannelRequest
	GetDtsJobId() *string
	SetRegionId(v string) *CreateConsumerChannelRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateConsumerChannelRequest
	GetResourceGroupId() *string
}

type CreateConsumerChannelRequest struct {
	// The name of the consumer group. The name can be up to 128 characters in length. We recommend that you use an informative name for easy identification.
	//
	// This parameter is required.
	//
	// example:
	//
	// 订阅组A
	ConsumerGroupName *string `json:"ConsumerGroupName,omitempty" xml:"ConsumerGroupName,omitempty"`
	// The password of the consumer group.
	//
	// 	- A password must contain two or more of the following characters: uppercase letters, lowercase letters, digits, and special characters.
	//
	// 	- A password must be 8 to 32 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test123456
	ConsumerGroupPassword *string `json:"ConsumerGroupPassword,omitempty" xml:"ConsumerGroupPassword,omitempty"`
	// The username of the consumer group.
	//
	// 	- A username must contain one or more of the following characters: uppercase letters, lowercase letters, digits, and underscores (_).
	//
	// 	- A username cannot exceed 16 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// dtstest
	ConsumerGroupUserName *string `json:"ConsumerGroupUserName,omitempty" xml:"ConsumerGroupUserName,omitempty"`
	// The ID of the change tracking instance. You can call the [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) operation to query the instance ID.
	//
	// >  You must specify at least one of the **DtsInstanceId*	- and **DtsJobId**. parameters.
	//
	// example:
	//
	// dtsboss6pn1w******
	DtsInstanceId *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	// The ID of the change tracking task. You can call the [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) operation to query the task ID.
	//
	// >  You must specify at least one of the **DtsInstanceId*	- and **DtsJobId**. parameters.
	//
	// example:
	//
	// boss6pn1w******
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
	// rg-aek2zx4uizich7y
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateConsumerChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateConsumerChannelRequest) GoString() string {
	return s.String()
}

func (s *CreateConsumerChannelRequest) GetConsumerGroupName() *string {
	return s.ConsumerGroupName
}

func (s *CreateConsumerChannelRequest) GetConsumerGroupPassword() *string {
	return s.ConsumerGroupPassword
}

func (s *CreateConsumerChannelRequest) GetConsumerGroupUserName() *string {
	return s.ConsumerGroupUserName
}

func (s *CreateConsumerChannelRequest) GetDtsInstanceId() *string {
	return s.DtsInstanceId
}

func (s *CreateConsumerChannelRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *CreateConsumerChannelRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateConsumerChannelRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateConsumerChannelRequest) SetConsumerGroupName(v string) *CreateConsumerChannelRequest {
	s.ConsumerGroupName = &v
	return s
}

func (s *CreateConsumerChannelRequest) SetConsumerGroupPassword(v string) *CreateConsumerChannelRequest {
	s.ConsumerGroupPassword = &v
	return s
}

func (s *CreateConsumerChannelRequest) SetConsumerGroupUserName(v string) *CreateConsumerChannelRequest {
	s.ConsumerGroupUserName = &v
	return s
}

func (s *CreateConsumerChannelRequest) SetDtsInstanceId(v string) *CreateConsumerChannelRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *CreateConsumerChannelRequest) SetDtsJobId(v string) *CreateConsumerChannelRequest {
	s.DtsJobId = &v
	return s
}

func (s *CreateConsumerChannelRequest) SetRegionId(v string) *CreateConsumerChannelRequest {
	s.RegionId = &v
	return s
}

func (s *CreateConsumerChannelRequest) SetResourceGroupId(v string) *CreateConsumerChannelRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateConsumerChannelRequest) Validate() error {
	return dara.Validate(s)
}
