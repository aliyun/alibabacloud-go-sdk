// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceSecurityOptionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyInstanceSecurityOptionsRequest
	GetClientToken() *string
	SetDryRun(v bool) *ModifyInstanceSecurityOptionsRequest
	GetDryRun() *bool
	SetEnableSecureBoot(v bool) *ModifyInstanceSecurityOptionsRequest
	GetEnableSecureBoot() *bool
	SetInstanceId(v string) *ModifyInstanceSecurityOptionsRequest
	GetInstanceId() *string
	SetOwnerId(v int64) *ModifyInstanceSecurityOptionsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyInstanceSecurityOptionsRequest
	GetRegionId() *string
	SetResourceOwnerId(v int64) *ModifyInstanceSecurityOptionsRequest
	GetResourceOwnerId() *int64
}

type ModifyInstanceSecurityOptionsRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The ClientToken value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://www.alibabacloud.com/help/en/ecs/developer-reference/how-to-ensure-idempotence).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run without performing the actual request. Valid values:
	//
	// - true: performs only a dry run. The secure boot setting of the instance is not modified.
	//
	// - false: performs a dry run and performs the actual request. If the request passes the dry run, the secure boot setting of the instance is modified.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to enable UEFI Secure Boot. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// Default value: false.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	EnableSecureBoot *bool `json:"EnableSecureBoot,omitempty" xml:"EnableSecureBoot,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp67acfmxazb4ph****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance. You can call [DescribeRegions](https://www.alibabacloud.com/help/en/ecs/developer-reference/api-ecs-2014-05-26-describeregions) to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyInstanceSecurityOptionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceSecurityOptionsRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceSecurityOptionsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyInstanceSecurityOptionsRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyInstanceSecurityOptionsRequest) GetEnableSecureBoot() *bool {
	return s.EnableSecureBoot
}

func (s *ModifyInstanceSecurityOptionsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceSecurityOptionsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyInstanceSecurityOptionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyInstanceSecurityOptionsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyInstanceSecurityOptionsRequest) SetClientToken(v string) *ModifyInstanceSecurityOptionsRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyInstanceSecurityOptionsRequest) SetDryRun(v bool) *ModifyInstanceSecurityOptionsRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyInstanceSecurityOptionsRequest) SetEnableSecureBoot(v bool) *ModifyInstanceSecurityOptionsRequest {
	s.EnableSecureBoot = &v
	return s
}

func (s *ModifyInstanceSecurityOptionsRequest) SetInstanceId(v string) *ModifyInstanceSecurityOptionsRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceSecurityOptionsRequest) SetOwnerId(v int64) *ModifyInstanceSecurityOptionsRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceSecurityOptionsRequest) SetRegionId(v string) *ModifyInstanceSecurityOptionsRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceSecurityOptionsRequest) SetResourceOwnerId(v int64) *ModifyInstanceSecurityOptionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceSecurityOptionsRequest) Validate() error {
	return dara.Validate(s)
}
