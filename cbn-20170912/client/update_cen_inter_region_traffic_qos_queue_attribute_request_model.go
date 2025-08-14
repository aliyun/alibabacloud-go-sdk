// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCenInterRegionTrafficQosQueueAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int64) *UpdateCenInterRegionTrafficQosQueueAttributeRequest
	GetBandwidth() *int64
	SetClientToken(v string) *UpdateCenInterRegionTrafficQosQueueAttributeRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateCenInterRegionTrafficQosQueueAttributeRequest
	GetDryRun() *bool
	SetDscps(v []*int32) *UpdateCenInterRegionTrafficQosQueueAttributeRequest
	GetDscps() []*int32
	SetOwnerAccount(v string) *UpdateCenInterRegionTrafficQosQueueAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateCenInterRegionTrafficQosQueueAttributeRequest
	GetOwnerId() *int64
	SetQosQueueDescription(v string) *UpdateCenInterRegionTrafficQosQueueAttributeRequest
	GetQosQueueDescription() *string
	SetQosQueueId(v string) *UpdateCenInterRegionTrafficQosQueueAttributeRequest
	GetQosQueueId() *string
	SetQosQueueName(v string) *UpdateCenInterRegionTrafficQosQueueAttributeRequest
	GetQosQueueName() *string
	SetRemainBandwidthPercent(v string) *UpdateCenInterRegionTrafficQosQueueAttributeRequest
	GetRemainBandwidthPercent() *string
	SetResourceOwnerAccount(v string) *UpdateCenInterRegionTrafficQosQueueAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateCenInterRegionTrafficQosQueueAttributeRequest
	GetResourceOwnerId() *int64
}

type UpdateCenInterRegionTrafficQosQueueAttributeRequest struct {
	// The absolute bandwidth value that can be allocated to the current queue. Unit: Mbit/s.
	//
	// Enter a number. You do not need to enter a unit.
	//
	// example:
	//
	// 1
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The differentiated services code point (DSCP) value used to match packets in the queue.
	Dscps        []*int32 `json:"Dscps,omitempty" xml:"Dscps,omitempty" type:"Repeated"`
	OwnerAccount *string  `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64   `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The new description of the queue.
	//
	// The description must be 1 to 256 characters in length, and cannot start with http:// or https://. You can also leave this parameter empty.
	//
	// example:
	//
	// desctest
	QosQueueDescription *string `json:"QosQueueDescription,omitempty" xml:"QosQueueDescription,omitempty"`
	// The queue ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// qos-queue-nv2vfzqkewhk4t****
	QosQueueId *string `json:"QosQueueId,omitempty" xml:"QosQueueId,omitempty"`
	// The new name of the queue.
	//
	// The name must be 1 to 128 characters in length, and cannot start with http:// or https://. You can also leave this parameter empty.
	//
	// example:
	//
	// nametest
	QosQueueName *string `json:"QosQueueName,omitempty" xml:"QosQueueName,omitempty"`
	// The percentage of bandwidth that can be allocated to the current queue.
	//
	// Enter a number. You do not need to enter a percent sign (%).
	//
	// example:
	//
	// 1
	RemainBandwidthPercent *string `json:"RemainBandwidthPercent,omitempty" xml:"RemainBandwidthPercent,omitempty"`
	ResourceOwnerAccount   *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId        *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdateCenInterRegionTrafficQosQueueAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCenInterRegionTrafficQosQueueAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeRequest) GetBandwidth() *int64 {
	return s.Bandwidth
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeRequest) GetDscps() []*int32 {
	return s.Dscps
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeRequest) GetQosQueueDescription() *string {
	return s.QosQueueDescription
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeRequest) GetQosQueueId() *string {
	return s.QosQueueId
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeRequest) GetQosQueueName() *string {
	return s.QosQueueName
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeRequest) GetRemainBandwidthPercent() *string {
	return s.RemainBandwidthPercent
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeRequest) SetBandwidth(v int64) *UpdateCenInterRegionTrafficQosQueueAttributeRequest {
	s.Bandwidth = &v
	return s
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeRequest) SetClientToken(v string) *UpdateCenInterRegionTrafficQosQueueAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeRequest) SetDryRun(v bool) *UpdateCenInterRegionTrafficQosQueueAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeRequest) SetDscps(v []*int32) *UpdateCenInterRegionTrafficQosQueueAttributeRequest {
	s.Dscps = v
	return s
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeRequest) SetOwnerAccount(v string) *UpdateCenInterRegionTrafficQosQueueAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeRequest) SetOwnerId(v int64) *UpdateCenInterRegionTrafficQosQueueAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeRequest) SetQosQueueDescription(v string) *UpdateCenInterRegionTrafficQosQueueAttributeRequest {
	s.QosQueueDescription = &v
	return s
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeRequest) SetQosQueueId(v string) *UpdateCenInterRegionTrafficQosQueueAttributeRequest {
	s.QosQueueId = &v
	return s
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeRequest) SetQosQueueName(v string) *UpdateCenInterRegionTrafficQosQueueAttributeRequest {
	s.QosQueueName = &v
	return s
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeRequest) SetRemainBandwidthPercent(v string) *UpdateCenInterRegionTrafficQosQueueAttributeRequest {
	s.RemainBandwidthPercent = &v
	return s
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeRequest) SetResourceOwnerAccount(v string) *UpdateCenInterRegionTrafficQosQueueAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeRequest) SetResourceOwnerId(v int64) *UpdateCenInterRegionTrafficQosQueueAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeRequest) Validate() error {
	return dara.Validate(s)
}
