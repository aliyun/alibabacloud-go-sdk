// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityAgentStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeSecurityAgentStatusRequest
	GetClientToken() *string
	SetInstanceId(v string) *DescribeSecurityAgentStatusRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeSecurityAgentStatusRequest
	GetRegionId() *string
}

type DescribeSecurityAgentStatusRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeSecurityAgentStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityAgentStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityAgentStatusRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeSecurityAgentStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSecurityAgentStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSecurityAgentStatusRequest) SetClientToken(v string) *DescribeSecurityAgentStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeSecurityAgentStatusRequest) SetInstanceId(v string) *DescribeSecurityAgentStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSecurityAgentStatusRequest) SetRegionId(v string) *DescribeSecurityAgentStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSecurityAgentStatusRequest) Validate() error {
	return dara.Validate(s)
}
