// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoverPhysicalConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *RecoverPhysicalConnectionRequest
	GetDryRun() *bool
	SetInstanceId(v string) *RecoverPhysicalConnectionRequest
	GetInstanceId() *string
	SetRegionId(v string) *RecoverPhysicalConnectionRequest
	GetRegionId() *string
	SetToken(v string) *RecoverPhysicalConnectionRequest
	GetToken() *string
}

type RecoverPhysicalConnectionRequest struct {
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run without recovering access to the Express Connect circuit. The system checks the required parameters, request format, and instance status. If the check fails, the corresponding error is returned. If the check succeeds, the request ID is returned.
	//
	// - **false*	- (default): sends the request. After the request passes the check, access to the Express Connect circuit is recovered.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The instance ID of the Express Connect circuit.
	//
	// > Currently, only shared Express Connect circuits can be recovered.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-bp1mrgfbtmc9brre7****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the Express Connect circuit.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// The client generates the value of this parameter. The value must be unique among different requests and cannot exceed 64 ASCII characters in length.
	//
	// example:
	//
	// CBCE910E-D396-4944-8****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s RecoverPhysicalConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s RecoverPhysicalConnectionRequest) GoString() string {
	return s.String()
}

func (s *RecoverPhysicalConnectionRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *RecoverPhysicalConnectionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RecoverPhysicalConnectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RecoverPhysicalConnectionRequest) GetToken() *string {
	return s.Token
}

func (s *RecoverPhysicalConnectionRequest) SetDryRun(v bool) *RecoverPhysicalConnectionRequest {
	s.DryRun = &v
	return s
}

func (s *RecoverPhysicalConnectionRequest) SetInstanceId(v string) *RecoverPhysicalConnectionRequest {
	s.InstanceId = &v
	return s
}

func (s *RecoverPhysicalConnectionRequest) SetRegionId(v string) *RecoverPhysicalConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *RecoverPhysicalConnectionRequest) SetToken(v string) *RecoverPhysicalConnectionRequest {
	s.Token = &v
	return s
}

func (s *RecoverPhysicalConnectionRequest) Validate() error {
	return dara.Validate(s)
}
