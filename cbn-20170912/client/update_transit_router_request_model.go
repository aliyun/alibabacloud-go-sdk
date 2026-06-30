// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTransitRouterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateTransitRouterRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateTransitRouterRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *UpdateTransitRouterRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateTransitRouterRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpdateTransitRouterRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *UpdateTransitRouterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateTransitRouterRequest
	GetResourceOwnerId() *int64
	SetTransitRouterDescription(v string) *UpdateTransitRouterRequest
	GetTransitRouterDescription() *string
	SetTransitRouterId(v string) *UpdateTransitRouterRequest
	GetTransitRouterId() *string
	SetTransitRouterName(v string) *UpdateTransitRouterRequest
	GetTransitRouterName() *string
}

type UpdateTransitRouterRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// Generate a client token to make sure that the value is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the request as the **ClientToken**. The **RequestId*	- of each request is unique.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. A dry run checks permissions and the status of the instance. Valid values:
	//
	// - **false*	- (default): Sends a normal request. After the request passes the check, the information about the TransitRouter instance is modified.
	//
	// - **true**: Sends a check request. The system checks the request for required parameters and format correctness, but does not modify the TransitRouter instance. If the check fails, an error is returned. If the check passes, the request ID is returned.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the TransitRouter instance is deployed.
	//
	// Call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The new description of the TransitRouter instance.
	//
	// The description can be empty or 1 to 256 characters in length. The description cannot start with http\\:// or https\\://.
	//
	// example:
	//
	// testdesc
	TransitRouterDescription *string `json:"TransitRouterDescription,omitempty" xml:"TransitRouterDescription,omitempty"`
	// The ID of the TransitRouter instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-uf654ttymmljlvh2x****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// The new name for the TransitRouter instance.
	//
	// The name can be empty or 1 to 128 characters in length. The name cannot start with http\\:// or https\\://.
	//
	// example:
	//
	// testname
	TransitRouterName *string `json:"TransitRouterName,omitempty" xml:"TransitRouterName,omitempty"`
}

func (s UpdateTransitRouterRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTransitRouterRequest) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateTransitRouterRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateTransitRouterRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateTransitRouterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateTransitRouterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateTransitRouterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateTransitRouterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateTransitRouterRequest) GetTransitRouterDescription() *string {
	return s.TransitRouterDescription
}

func (s *UpdateTransitRouterRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *UpdateTransitRouterRequest) GetTransitRouterName() *string {
	return s.TransitRouterName
}

func (s *UpdateTransitRouterRequest) SetClientToken(v string) *UpdateTransitRouterRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTransitRouterRequest) SetDryRun(v bool) *UpdateTransitRouterRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateTransitRouterRequest) SetOwnerAccount(v string) *UpdateTransitRouterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateTransitRouterRequest) SetOwnerId(v int64) *UpdateTransitRouterRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateTransitRouterRequest) SetRegionId(v string) *UpdateTransitRouterRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateTransitRouterRequest) SetResourceOwnerAccount(v string) *UpdateTransitRouterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateTransitRouterRequest) SetResourceOwnerId(v int64) *UpdateTransitRouterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateTransitRouterRequest) SetTransitRouterDescription(v string) *UpdateTransitRouterRequest {
	s.TransitRouterDescription = &v
	return s
}

func (s *UpdateTransitRouterRequest) SetTransitRouterId(v string) *UpdateTransitRouterRequest {
	s.TransitRouterId = &v
	return s
}

func (s *UpdateTransitRouterRequest) SetTransitRouterName(v string) *UpdateTransitRouterRequest {
	s.TransitRouterName = &v
	return s
}

func (s *UpdateTransitRouterRequest) Validate() error {
	return dara.Validate(s)
}
