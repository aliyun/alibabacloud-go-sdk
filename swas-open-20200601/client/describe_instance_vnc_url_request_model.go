// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceVncUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeInstanceVncUrlRequest
	GetClientToken() *string
	SetInstanceId(v string) *DescribeInstanceVncUrlRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeInstanceVncUrlRequest
	GetRegionId() *string
}

type DescribeInstanceVncUrlRequest struct {
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
	// 2ad1ae67295445f598017499dc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeInstanceVncUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceVncUrlRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceVncUrlRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeInstanceVncUrlRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceVncUrlRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceVncUrlRequest) SetClientToken(v string) *DescribeInstanceVncUrlRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeInstanceVncUrlRequest) SetInstanceId(v string) *DescribeInstanceVncUrlRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceVncUrlRequest) SetRegionId(v string) *DescribeInstanceVncUrlRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceVncUrlRequest) Validate() error {
	return dara.Validate(s)
}
