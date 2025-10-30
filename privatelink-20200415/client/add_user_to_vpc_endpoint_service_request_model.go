// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserToVpcEndpointServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AddUserToVpcEndpointServiceRequest
	GetClientToken() *string
	SetDryRun(v bool) *AddUserToVpcEndpointServiceRequest
	GetDryRun() *bool
	SetRegionId(v string) *AddUserToVpcEndpointServiceRequest
	GetRegionId() *string
	SetServiceId(v string) *AddUserToVpcEndpointServiceRequest
	GetServiceId() *string
	SetUserARN(v string) *AddUserToVpcEndpointServiceRequest
	GetUserARN() *string
	SetUserId(v int64) *AddUserToVpcEndpointServiceRequest
	GetUserId() *int64
}

type AddUserToVpcEndpointServiceRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The region ID of the endpoint service. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// eu-west-1
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The endpoint service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The whitelist in the format of Aliyun Resource Name (ARN).
	//
	// example:
	//
	// acs:ram:*:<account-id>:*
	UserARN *string `json:"UserARN,omitempty" xml:"UserARN,omitempty"`
	// The account ID that you want to add to the whitelist.
	//
	// example:
	//
	// 132193271328****
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AddUserToVpcEndpointServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s AddUserToVpcEndpointServiceRequest) GoString() string {
	return s.String()
}

func (s *AddUserToVpcEndpointServiceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AddUserToVpcEndpointServiceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *AddUserToVpcEndpointServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddUserToVpcEndpointServiceRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *AddUserToVpcEndpointServiceRequest) GetUserARN() *string {
	return s.UserARN
}

func (s *AddUserToVpcEndpointServiceRequest) GetUserId() *int64 {
	return s.UserId
}

func (s *AddUserToVpcEndpointServiceRequest) SetClientToken(v string) *AddUserToVpcEndpointServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *AddUserToVpcEndpointServiceRequest) SetDryRun(v bool) *AddUserToVpcEndpointServiceRequest {
	s.DryRun = &v
	return s
}

func (s *AddUserToVpcEndpointServiceRequest) SetRegionId(v string) *AddUserToVpcEndpointServiceRequest {
	s.RegionId = &v
	return s
}

func (s *AddUserToVpcEndpointServiceRequest) SetServiceId(v string) *AddUserToVpcEndpointServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *AddUserToVpcEndpointServiceRequest) SetUserARN(v string) *AddUserToVpcEndpointServiceRequest {
	s.UserARN = &v
	return s
}

func (s *AddUserToVpcEndpointServiceRequest) SetUserId(v int64) *AddUserToVpcEndpointServiceRequest {
	s.UserId = &v
	return s
}

func (s *AddUserToVpcEndpointServiceRequest) Validate() error {
	return dara.Validate(s)
}
