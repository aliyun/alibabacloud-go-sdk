// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstancePasswordsSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeInstancePasswordsSettingRequest
	GetClientToken() *string
	SetInstanceId(v string) *DescribeInstancePasswordsSettingRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeInstancePasswordsSettingRequest
	GetRegionId() *string
}

type DescribeInstancePasswordsSettingRequest struct {
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
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/2361076.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeInstancePasswordsSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancePasswordsSettingRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancePasswordsSettingRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeInstancePasswordsSettingRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstancePasswordsSettingRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstancePasswordsSettingRequest) SetClientToken(v string) *DescribeInstancePasswordsSettingRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeInstancePasswordsSettingRequest) SetInstanceId(v string) *DescribeInstancePasswordsSettingRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancePasswordsSettingRequest) SetRegionId(v string) *DescribeInstancePasswordsSettingRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstancePasswordsSettingRequest) Validate() error {
	return dara.Validate(s)
}
