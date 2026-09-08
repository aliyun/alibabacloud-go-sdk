// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCenInterRegionTrafficQosQueueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteCenInterRegionTrafficQosQueueRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteCenInterRegionTrafficQosQueueRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *DeleteCenInterRegionTrafficQosQueueRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteCenInterRegionTrafficQosQueueRequest
	GetOwnerId() *int64
	SetQosQueueId(v string) *DeleteCenInterRegionTrafficQosQueueRequest
	GetQosQueueId() *string
	SetResourceOwnerAccount(v string) *DeleteCenInterRegionTrafficQosQueueRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteCenInterRegionTrafficQosQueueRequest
	GetResourceOwnerId() *int64
}

type DeleteCenInterRegionTrafficQosQueueRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- as the **ClientToken**. The **RequestId*	- of each API request may be different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: Sends a check request without deleting the queue from the traffic scheduling policy. The system checks the required parameters, request format, and business limits. If the check fails, the corresponding error is returned. If the check passes, the error code `DryRunOperation` is returned.
	//
	// - **false*	- (default): Sends a normal request. After the check passes, the queue is directly deleted from the traffic scheduling policy.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The queue ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// qos-queue-nv2vfzqkewhk4t****
	QosQueueId           *string `json:"QosQueueId,omitempty" xml:"QosQueueId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteCenInterRegionTrafficQosQueueRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCenInterRegionTrafficQosQueueRequest) GoString() string {
	return s.String()
}

func (s *DeleteCenInterRegionTrafficQosQueueRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteCenInterRegionTrafficQosQueueRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteCenInterRegionTrafficQosQueueRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteCenInterRegionTrafficQosQueueRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteCenInterRegionTrafficQosQueueRequest) GetQosQueueId() *string {
	return s.QosQueueId
}

func (s *DeleteCenInterRegionTrafficQosQueueRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteCenInterRegionTrafficQosQueueRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteCenInterRegionTrafficQosQueueRequest) SetClientToken(v string) *DeleteCenInterRegionTrafficQosQueueRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteCenInterRegionTrafficQosQueueRequest) SetDryRun(v bool) *DeleteCenInterRegionTrafficQosQueueRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteCenInterRegionTrafficQosQueueRequest) SetOwnerAccount(v string) *DeleteCenInterRegionTrafficQosQueueRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteCenInterRegionTrafficQosQueueRequest) SetOwnerId(v int64) *DeleteCenInterRegionTrafficQosQueueRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCenInterRegionTrafficQosQueueRequest) SetQosQueueId(v string) *DeleteCenInterRegionTrafficQosQueueRequest {
	s.QosQueueId = &v
	return s
}

func (s *DeleteCenInterRegionTrafficQosQueueRequest) SetResourceOwnerAccount(v string) *DeleteCenInterRegionTrafficQosQueueRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteCenInterRegionTrafficQosQueueRequest) SetResourceOwnerId(v int64) *DeleteCenInterRegionTrafficQosQueueRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteCenInterRegionTrafficQosQueueRequest) Validate() error {
	return dara.Validate(s)
}
