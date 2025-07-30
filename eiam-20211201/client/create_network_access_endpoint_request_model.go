// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkAccessEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateNetworkAccessEndpointRequest
	GetClientToken() *string
	SetInstanceId(v string) *CreateNetworkAccessEndpointRequest
	GetInstanceId() *string
	SetNetworkAccessEndpointName(v string) *CreateNetworkAccessEndpointRequest
	GetNetworkAccessEndpointName() *string
	SetVSwitchIds(v []*string) *CreateNetworkAccessEndpointRequest
	GetVSwitchIds() []*string
	SetVpcId(v string) *CreateNetworkAccessEndpointRequest
	GetVpcId() *string
	SetVpcRegionId(v string) *CreateNetworkAccessEndpointRequest
	GetVpcRegionId() *string
}

type CreateNetworkAccessEndpointRequest struct {
	// Idempotent token.
	//
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Private network endpoint name.
	//
	// This parameter is required.
	//
	// example:
	//
	// eiam-vpc-access-endpoint
	NetworkAccessEndpointName *string `json:"NetworkAccessEndpointName,omitempty" xml:"NetworkAccessEndpointName,omitempty"`
	// The IDs of vSwitches.
	//
	// example:
	//
	// vsw-examplexxx
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// The ID of the VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-examplexxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The region ID of the outbound VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	VpcRegionId *string `json:"VpcRegionId,omitempty" xml:"VpcRegionId,omitempty"`
}

func (s CreateNetworkAccessEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkAccessEndpointRequest) GoString() string {
	return s.String()
}

func (s *CreateNetworkAccessEndpointRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateNetworkAccessEndpointRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateNetworkAccessEndpointRequest) GetNetworkAccessEndpointName() *string {
	return s.NetworkAccessEndpointName
}

func (s *CreateNetworkAccessEndpointRequest) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *CreateNetworkAccessEndpointRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateNetworkAccessEndpointRequest) GetVpcRegionId() *string {
	return s.VpcRegionId
}

func (s *CreateNetworkAccessEndpointRequest) SetClientToken(v string) *CreateNetworkAccessEndpointRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateNetworkAccessEndpointRequest) SetInstanceId(v string) *CreateNetworkAccessEndpointRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateNetworkAccessEndpointRequest) SetNetworkAccessEndpointName(v string) *CreateNetworkAccessEndpointRequest {
	s.NetworkAccessEndpointName = &v
	return s
}

func (s *CreateNetworkAccessEndpointRequest) SetVSwitchIds(v []*string) *CreateNetworkAccessEndpointRequest {
	s.VSwitchIds = v
	return s
}

func (s *CreateNetworkAccessEndpointRequest) SetVpcId(v string) *CreateNetworkAccessEndpointRequest {
	s.VpcId = &v
	return s
}

func (s *CreateNetworkAccessEndpointRequest) SetVpcRegionId(v string) *CreateNetworkAccessEndpointRequest {
	s.VpcRegionId = &v
	return s
}

func (s *CreateNetworkAccessEndpointRequest) Validate() error {
	return dara.Validate(s)
}
