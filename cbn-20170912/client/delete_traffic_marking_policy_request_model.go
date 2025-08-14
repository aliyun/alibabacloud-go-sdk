// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTrafficMarkingPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteTrafficMarkingPolicyRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteTrafficMarkingPolicyRequest
	GetDryRun() *bool
	SetForce(v bool) *DeleteTrafficMarkingPolicyRequest
	GetForce() *bool
	SetOwnerAccount(v string) *DeleteTrafficMarkingPolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteTrafficMarkingPolicyRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteTrafficMarkingPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteTrafficMarkingPolicyRequest
	GetResourceOwnerId() *int64
	SetTrafficMarkingPolicyId(v string) *DeleteTrafficMarkingPolicyRequest
	GetTrafficMarkingPolicyId() *string
}

type DeleteTrafficMarkingPolicyRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among all requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the value of **RequestId*	- as the value of **ClientToken**. The value of **RequestId*	- for each API request may be different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// 	- **true**: performs a dry run. The system checks the required parameters, request syntax, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and sends the request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Whether to force delete the traffic marking policy. Valid values:
	//
	// 	- **false*	- (default): checks whether there is a traffic classification rule before deleting the traffic marking policy. If there is, the traffic marking policy cannot be deleted and an error is returned.
	//
	// 	- **true**: When you delete a traffic marking policy, all traffic classification rules are deleted by default.
	//
	// example:
	//
	// false
	Force                *bool   `json:"Force,omitempty" xml:"Force,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the traffic marking policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// tm-u9nxup5kww5po8****
	TrafficMarkingPolicyId *string `json:"TrafficMarkingPolicyId,omitempty" xml:"TrafficMarkingPolicyId,omitempty"`
}

func (s DeleteTrafficMarkingPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTrafficMarkingPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteTrafficMarkingPolicyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteTrafficMarkingPolicyRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteTrafficMarkingPolicyRequest) GetForce() *bool {
	return s.Force
}

func (s *DeleteTrafficMarkingPolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteTrafficMarkingPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteTrafficMarkingPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteTrafficMarkingPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteTrafficMarkingPolicyRequest) GetTrafficMarkingPolicyId() *string {
	return s.TrafficMarkingPolicyId
}

func (s *DeleteTrafficMarkingPolicyRequest) SetClientToken(v string) *DeleteTrafficMarkingPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteTrafficMarkingPolicyRequest) SetDryRun(v bool) *DeleteTrafficMarkingPolicyRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteTrafficMarkingPolicyRequest) SetForce(v bool) *DeleteTrafficMarkingPolicyRequest {
	s.Force = &v
	return s
}

func (s *DeleteTrafficMarkingPolicyRequest) SetOwnerAccount(v string) *DeleteTrafficMarkingPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteTrafficMarkingPolicyRequest) SetOwnerId(v int64) *DeleteTrafficMarkingPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteTrafficMarkingPolicyRequest) SetResourceOwnerAccount(v string) *DeleteTrafficMarkingPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteTrafficMarkingPolicyRequest) SetResourceOwnerId(v int64) *DeleteTrafficMarkingPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteTrafficMarkingPolicyRequest) SetTrafficMarkingPolicyId(v string) *DeleteTrafficMarkingPolicyRequest {
	s.TrafficMarkingPolicyId = &v
	return s
}

func (s *DeleteTrafficMarkingPolicyRequest) Validate() error {
	return dara.Validate(s)
}
