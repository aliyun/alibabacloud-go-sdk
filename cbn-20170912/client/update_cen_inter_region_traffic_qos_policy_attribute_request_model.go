// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCenInterRegionTrafficQosPolicyAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateCenInterRegionTrafficQosPolicyAttributeRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateCenInterRegionTrafficQosPolicyAttributeRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *UpdateCenInterRegionTrafficQosPolicyAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateCenInterRegionTrafficQosPolicyAttributeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpdateCenInterRegionTrafficQosPolicyAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateCenInterRegionTrafficQosPolicyAttributeRequest
	GetResourceOwnerId() *int64
	SetTrafficQosPolicyDescription(v string) *UpdateCenInterRegionTrafficQosPolicyAttributeRequest
	GetTrafficQosPolicyDescription() *string
	SetTrafficQosPolicyId(v string) *UpdateCenInterRegionTrafficQosPolicyAttributeRequest
	GetTrafficQosPolicyId() *string
	SetTrafficQosPolicyName(v string) *UpdateCenInterRegionTrafficQosPolicyAttributeRequest
	GetTrafficQosPolicyName() *string
}

type UpdateCenInterRegionTrafficQosPolicyAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// 	- **true**: performs a dry run. The system checks the required parameters, request syntax, and limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, the operation is performed.
	//
	// example:
	//
	// false
	DryRun               *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The new description of the QoS policy.
	//
	// The description must be 1 to 256 characters in length, and cannot start with http:// or https://. You can also leave this parameter empty.
	//
	// example:
	//
	// desctest
	TrafficQosPolicyDescription *string `json:"TrafficQosPolicyDescription,omitempty" xml:"TrafficQosPolicyDescription,omitempty"`
	// The ID of the QoS policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// qos-eczzew0v1kzrb5****
	TrafficQosPolicyId *string `json:"TrafficQosPolicyId,omitempty" xml:"TrafficQosPolicyId,omitempty"`
	// The new name of the QoS policy.
	//
	// The name must be 1 to 128 characters in length, and cannot start with http:// or https://. You can also leave this parameter empty.
	//
	// example:
	//
	// nametest
	TrafficQosPolicyName *string `json:"TrafficQosPolicyName,omitempty" xml:"TrafficQosPolicyName,omitempty"`
}

func (s UpdateCenInterRegionTrafficQosPolicyAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCenInterRegionTrafficQosPolicyAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateCenInterRegionTrafficQosPolicyAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCenInterRegionTrafficQosPolicyAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateCenInterRegionTrafficQosPolicyAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateCenInterRegionTrafficQosPolicyAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateCenInterRegionTrafficQosPolicyAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateCenInterRegionTrafficQosPolicyAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateCenInterRegionTrafficQosPolicyAttributeRequest) GetTrafficQosPolicyDescription() *string {
	return s.TrafficQosPolicyDescription
}

func (s *UpdateCenInterRegionTrafficQosPolicyAttributeRequest) GetTrafficQosPolicyId() *string {
	return s.TrafficQosPolicyId
}

func (s *UpdateCenInterRegionTrafficQosPolicyAttributeRequest) GetTrafficQosPolicyName() *string {
	return s.TrafficQosPolicyName
}

func (s *UpdateCenInterRegionTrafficQosPolicyAttributeRequest) SetClientToken(v string) *UpdateCenInterRegionTrafficQosPolicyAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCenInterRegionTrafficQosPolicyAttributeRequest) SetDryRun(v bool) *UpdateCenInterRegionTrafficQosPolicyAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateCenInterRegionTrafficQosPolicyAttributeRequest) SetOwnerAccount(v string) *UpdateCenInterRegionTrafficQosPolicyAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateCenInterRegionTrafficQosPolicyAttributeRequest) SetOwnerId(v int64) *UpdateCenInterRegionTrafficQosPolicyAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateCenInterRegionTrafficQosPolicyAttributeRequest) SetResourceOwnerAccount(v string) *UpdateCenInterRegionTrafficQosPolicyAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateCenInterRegionTrafficQosPolicyAttributeRequest) SetResourceOwnerId(v int64) *UpdateCenInterRegionTrafficQosPolicyAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateCenInterRegionTrafficQosPolicyAttributeRequest) SetTrafficQosPolicyDescription(v string) *UpdateCenInterRegionTrafficQosPolicyAttributeRequest {
	s.TrafficQosPolicyDescription = &v
	return s
}

func (s *UpdateCenInterRegionTrafficQosPolicyAttributeRequest) SetTrafficQosPolicyId(v string) *UpdateCenInterRegionTrafficQosPolicyAttributeRequest {
	s.TrafficQosPolicyId = &v
	return s
}

func (s *UpdateCenInterRegionTrafficQosPolicyAttributeRequest) SetTrafficQosPolicyName(v string) *UpdateCenInterRegionTrafficQosPolicyAttributeRequest {
	s.TrafficQosPolicyName = &v
	return s
}

func (s *UpdateCenInterRegionTrafficQosPolicyAttributeRequest) Validate() error {
	return dara.Validate(s)
}
