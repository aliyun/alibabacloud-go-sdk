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
	// Specifies whether to perform only a dry run on the request. Valid values:
	//
	// - **true**: performs only a dry run and does not restore the connection of the Express Connect circuit. The system checks whether the required parameters are specified, the request format is valid, and the instance status is normal. If the request fails the dry run, an error message is returned. If the request passes the dry run, the request ID is returned.
	//
	// - **false*	- (default): performs a dry run and restores the connection of the Express Connect circuit if the request passes the dry run.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The instance ID of the Express Connect circuit to recover.
	//
	// > Currently, only the connection of shared Express Connect circuits can be recovered.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-bp1mrgfbtmc9brre7****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the Express Connect circuit.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
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
