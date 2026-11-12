// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceKeyPairRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeInstanceKeyPairRequest
	GetClientToken() *string
	SetInstanceId(v string) *DescribeInstanceKeyPairRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeInstanceKeyPairRequest
	GetRegionId() *string
}

type DescribeInstanceKeyPairRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
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

func (s DescribeInstanceKeyPairRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceKeyPairRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceKeyPairRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeInstanceKeyPairRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceKeyPairRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceKeyPairRequest) SetClientToken(v string) *DescribeInstanceKeyPairRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeInstanceKeyPairRequest) SetInstanceId(v string) *DescribeInstanceKeyPairRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceKeyPairRequest) SetRegionId(v string) *DescribeInstanceKeyPairRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceKeyPairRequest) Validate() error {
	return dara.Validate(s)
}
