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
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the system returns the ID of the request.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the transit router.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The description of the transit router.
	//
	// The description must be 1 to 256 characters in length, and cannot start with http:// or https://. You can also leave this parameter empty.
	//
	// example:
	//
	// testdesc
	TransitRouterDescription *string `json:"TransitRouterDescription,omitempty" xml:"TransitRouterDescription,omitempty"`
	// The transit router ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-uf654ttymmljlvh2x****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// The transit router name.
	//
	// The name must be 1 to 128 characters in length, and cannot start with http:// or https://. You can also leave this parameter empty.
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
