// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserFromVpcEndpointServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RemoveUserFromVpcEndpointServiceRequest
	GetClientToken() *string
	SetDryRun(v bool) *RemoveUserFromVpcEndpointServiceRequest
	GetDryRun() *bool
	SetRegionId(v string) *RemoveUserFromVpcEndpointServiceRequest
	GetRegionId() *string
	SetServiceId(v string) *RemoveUserFromVpcEndpointServiceRequest
	GetServiceId() *string
	SetUserARN(v string) *RemoveUserFromVpcEndpointServiceRequest
	GetUserARN() *string
	SetUserId(v int64) *RemoveUserFromVpcEndpointServiceRequest
	GetUserId() *int64
}

type RemoveUserFromVpcEndpointServiceRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the AccessKey pair, the permissions of the RAM user, and the required parameters. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The region ID of the endpoint service for which you want to remove the account ID from the whitelist. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
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
	// The account ID that you want to remove from the whitelist.
	//
	// example:
	//
	// 12345678
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s RemoveUserFromVpcEndpointServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserFromVpcEndpointServiceRequest) GoString() string {
	return s.String()
}

func (s *RemoveUserFromVpcEndpointServiceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RemoveUserFromVpcEndpointServiceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *RemoveUserFromVpcEndpointServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RemoveUserFromVpcEndpointServiceRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *RemoveUserFromVpcEndpointServiceRequest) GetUserARN() *string {
	return s.UserARN
}

func (s *RemoveUserFromVpcEndpointServiceRequest) GetUserId() *int64 {
	return s.UserId
}

func (s *RemoveUserFromVpcEndpointServiceRequest) SetClientToken(v string) *RemoveUserFromVpcEndpointServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveUserFromVpcEndpointServiceRequest) SetDryRun(v bool) *RemoveUserFromVpcEndpointServiceRequest {
	s.DryRun = &v
	return s
}

func (s *RemoveUserFromVpcEndpointServiceRequest) SetRegionId(v string) *RemoveUserFromVpcEndpointServiceRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveUserFromVpcEndpointServiceRequest) SetServiceId(v string) *RemoveUserFromVpcEndpointServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *RemoveUserFromVpcEndpointServiceRequest) SetUserARN(v string) *RemoveUserFromVpcEndpointServiceRequest {
	s.UserARN = &v
	return s
}

func (s *RemoveUserFromVpcEndpointServiceRequest) SetUserId(v int64) *RemoveUserFromVpcEndpointServiceRequest {
	s.UserId = &v
	return s
}

func (s *RemoveUserFromVpcEndpointServiceRequest) Validate() error {
	return dara.Validate(s)
}
