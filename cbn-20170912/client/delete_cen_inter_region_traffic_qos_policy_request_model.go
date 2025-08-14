// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCenInterRegionTrafficQosPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteCenInterRegionTrafficQosPolicyRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteCenInterRegionTrafficQosPolicyRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *DeleteCenInterRegionTrafficQosPolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteCenInterRegionTrafficQosPolicyRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteCenInterRegionTrafficQosPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteCenInterRegionTrafficQosPolicyRequest
	GetResourceOwnerId() *int64
	SetTrafficQosPolicyId(v string) *DeleteCenInterRegionTrafficQosPolicyRequest
	GetTrafficQosPolicyId() *string
}

type DeleteCenInterRegionTrafficQosPolicyRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among all requests. The client token can contain only ASCII characters.
	//
	// >  If you do not set this parameter, **ClientToken*	- is set to the value of **RequestId**. The value of **RequestId*	- for each API request may be different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Default value: false. Valid values:
	//
	// 	- **true**: performs a dry run. The system checks the required parameters, request syntax, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and sends the request.
	//
	// example:
	//
	// false
	DryRun               *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the QoS policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// qos-eczzew0v1kzrb5****
	TrafficQosPolicyId *string `json:"TrafficQosPolicyId,omitempty" xml:"TrafficQosPolicyId,omitempty"`
}

func (s DeleteCenInterRegionTrafficQosPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCenInterRegionTrafficQosPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteCenInterRegionTrafficQosPolicyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteCenInterRegionTrafficQosPolicyRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteCenInterRegionTrafficQosPolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteCenInterRegionTrafficQosPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteCenInterRegionTrafficQosPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteCenInterRegionTrafficQosPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteCenInterRegionTrafficQosPolicyRequest) GetTrafficQosPolicyId() *string {
	return s.TrafficQosPolicyId
}

func (s *DeleteCenInterRegionTrafficQosPolicyRequest) SetClientToken(v string) *DeleteCenInterRegionTrafficQosPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteCenInterRegionTrafficQosPolicyRequest) SetDryRun(v bool) *DeleteCenInterRegionTrafficQosPolicyRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteCenInterRegionTrafficQosPolicyRequest) SetOwnerAccount(v string) *DeleteCenInterRegionTrafficQosPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteCenInterRegionTrafficQosPolicyRequest) SetOwnerId(v int64) *DeleteCenInterRegionTrafficQosPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCenInterRegionTrafficQosPolicyRequest) SetResourceOwnerAccount(v string) *DeleteCenInterRegionTrafficQosPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteCenInterRegionTrafficQosPolicyRequest) SetResourceOwnerId(v int64) *DeleteCenInterRegionTrafficQosPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteCenInterRegionTrafficQosPolicyRequest) SetTrafficQosPolicyId(v string) *DeleteCenInterRegionTrafficQosPolicyRequest {
	s.TrafficQosPolicyId = &v
	return s
}

func (s *DeleteCenInterRegionTrafficQosPolicyRequest) Validate() error {
	return dara.Validate(s)
}
