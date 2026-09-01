// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcCidrBlockRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyVpcCidrBlockRequest
	GetClientToken() *string
	SetDryRun(v bool) *ModifyVpcCidrBlockRequest
	GetDryRun() *bool
	SetOriginalCidrBlock(v string) *ModifyVpcCidrBlockRequest
	GetOriginalCidrBlock() *string
	SetRegionId(v string) *ModifyVpcCidrBlockRequest
	GetRegionId() *string
	SetTargetCidrBlock(v string) *ModifyVpcCidrBlockRequest
	GetTargetCidrBlock() *string
	SetVpcId(v string) *ModifyVpcCidrBlockRequest
	GetVpcId() *string
}

type ModifyVpcCidrBlockRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The ClientToken value can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run without modifying the CIDR block of the virtual private cloud (VPC). The system checks the request for potential issues. If the check fails, the corresponding error is returned. If the check succeeds, the error code `DryRunOperation` is returned.
	//
	// - **false*	- (default): sends a Normal request. If the check succeeds, an HTTP 2xx status code is returned and the modification is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The CIDR block of the VPC to modify. Both primary and secondary CIDR blocks are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.0.0.0/8
	OriginalCidrBlock *string `json:"OriginalCidrBlock,omitempty" xml:"OriginalCidrBlock,omitempty"`
	// The ID of the region where the VPC resides.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The new CIDR block for the VPC after modification.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.0.0.0/16
	TargetCidrBlock *string `json:"TargetCidrBlock,omitempty" xml:"TargetCidrBlock,omitempty"`
	// The ID of the VPC to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1di7uewzmtvfuq8****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ModifyVpcCidrBlockRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcCidrBlockRequest) GoString() string {
	return s.String()
}

func (s *ModifyVpcCidrBlockRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyVpcCidrBlockRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyVpcCidrBlockRequest) GetOriginalCidrBlock() *string {
	return s.OriginalCidrBlock
}

func (s *ModifyVpcCidrBlockRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyVpcCidrBlockRequest) GetTargetCidrBlock() *string {
	return s.TargetCidrBlock
}

func (s *ModifyVpcCidrBlockRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ModifyVpcCidrBlockRequest) SetClientToken(v string) *ModifyVpcCidrBlockRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyVpcCidrBlockRequest) SetDryRun(v bool) *ModifyVpcCidrBlockRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyVpcCidrBlockRequest) SetOriginalCidrBlock(v string) *ModifyVpcCidrBlockRequest {
	s.OriginalCidrBlock = &v
	return s
}

func (s *ModifyVpcCidrBlockRequest) SetRegionId(v string) *ModifyVpcCidrBlockRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyVpcCidrBlockRequest) SetTargetCidrBlock(v string) *ModifyVpcCidrBlockRequest {
	s.TargetCidrBlock = &v
	return s
}

func (s *ModifyVpcCidrBlockRequest) SetVpcId(v string) *ModifyVpcCidrBlockRequest {
	s.VpcId = &v
	return s
}

func (s *ModifyVpcCidrBlockRequest) Validate() error {
	return dara.Validate(s)
}
