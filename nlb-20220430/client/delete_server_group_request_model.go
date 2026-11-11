// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServerGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteServerGroupRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteServerGroupRequest
	GetDryRun() *bool
	SetRegionId(v string) *DeleteServerGroupRequest
	GetRegionId() *string
	SetServerGroupId(v string) *DeleteServerGroupRequest
	GetServerGroupId() *string
}

type DeleteServerGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate the value of this parameter on your client to ensure that the value is unique across different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request may be different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to only perform a dry run on the request. Valid values:
	//
	// - **true**: Performs a dry run without deleting the server group. The system checks whether required parameters are specified, the request format is valid, and service limits are met. If the request fails the dry run, an error is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): Performs a standard request. If the request passes the dry run, an HTTP 2xx status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the region where the Network Load Balancer (NLB) instance is deployed.
	//
	// You can invoke the [DescribeRegions](https://help.aliyun.com/document_detail/443657.html) operation to obtain region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the server group.
	//
	// This parameter is required.
	//
	// example:
	//
	// sgp-atstuj3rtoptyui****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s DeleteServerGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteServerGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteServerGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteServerGroupRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteServerGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteServerGroupRequest) GetServerGroupId() *string {
	return s.ServerGroupId
}

func (s *DeleteServerGroupRequest) SetClientToken(v string) *DeleteServerGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteServerGroupRequest) SetDryRun(v bool) *DeleteServerGroupRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteServerGroupRequest) SetRegionId(v string) *DeleteServerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteServerGroupRequest) SetServerGroupId(v string) *DeleteServerGroupRequest {
	s.ServerGroupId = &v
	return s
}

func (s *DeleteServerGroupRequest) Validate() error {
	return dara.Validate(s)
}
