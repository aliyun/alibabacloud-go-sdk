// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudMonitorAgentStatusesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeCloudMonitorAgentStatusesRequest
	GetClientToken() *string
	SetInstanceIds(v string) *DescribeCloudMonitorAgentStatusesRequest
	GetInstanceIds() *string
	SetRegionId(v string) *DescribeCloudMonitorAgentStatusesRequest
	GetRegionId() *string
}

type DescribeCloudMonitorAgentStatusesRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The IDs of the simple application servers. The value can be a JSON array that consists of up to 100 simple application server IDs. Separate multiple server IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// ["2ad1ae67295445f598017499dc******","2ad1ae67295445f598017499dc******"]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The region ID of the simple application servers. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCloudMonitorAgentStatusesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudMonitorAgentStatusesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudMonitorAgentStatusesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeCloudMonitorAgentStatusesRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *DescribeCloudMonitorAgentStatusesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCloudMonitorAgentStatusesRequest) SetClientToken(v string) *DescribeCloudMonitorAgentStatusesRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeCloudMonitorAgentStatusesRequest) SetInstanceIds(v string) *DescribeCloudMonitorAgentStatusesRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeCloudMonitorAgentStatusesRequest) SetRegionId(v string) *DescribeCloudMonitorAgentStatusesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCloudMonitorAgentStatusesRequest) Validate() error {
	return dara.Validate(s)
}
