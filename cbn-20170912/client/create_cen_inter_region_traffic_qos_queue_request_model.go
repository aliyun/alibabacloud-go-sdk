// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCenInterRegionTrafficQosQueueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int64) *CreateCenInterRegionTrafficQosQueueRequest
	GetBandwidth() *int64
	SetClientToken(v string) *CreateCenInterRegionTrafficQosQueueRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateCenInterRegionTrafficQosQueueRequest
	GetDryRun() *bool
	SetDscps(v []*int32) *CreateCenInterRegionTrafficQosQueueRequest
	GetDscps() []*int32
	SetOwnerAccount(v string) *CreateCenInterRegionTrafficQosQueueRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateCenInterRegionTrafficQosQueueRequest
	GetOwnerId() *int64
	SetQosQueueDescription(v string) *CreateCenInterRegionTrafficQosQueueRequest
	GetQosQueueDescription() *string
	SetQosQueueName(v string) *CreateCenInterRegionTrafficQosQueueRequest
	GetQosQueueName() *string
	SetRemainBandwidthPercent(v string) *CreateCenInterRegionTrafficQosQueueRequest
	GetRemainBandwidthPercent() *string
	SetResourceOwnerAccount(v string) *CreateCenInterRegionTrafficQosQueueRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateCenInterRegionTrafficQosQueueRequest
	GetResourceOwnerId() *int64
	SetTrafficQosPolicyId(v string) *CreateCenInterRegionTrafficQosQueueRequest
	GetTrafficQosPolicyId() *string
}

type CreateCenInterRegionTrafficQosQueueRequest struct {
	// The maximum absolute bandwidth value that can be allocated to the queue. Unit: Mbit/s.
	//
	// - The value specifies an absolute bandwidth. For example, a value of 20 specifies that the queue can consume at most 20 Mbit/s of bandwidth.
	//
	// - The sum of the bandwidth values specified for all queues that belong to the same inter-region connection cannot exceed the maximum bandwidth of the inter-region connection.
	//
	// example:
	//
	// 20
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run. The system checks the required parameters, the request format, and the service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): performs a dry run and sends the request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The differentiated services code point (DSCP) value that matches the current queue.
	//
	// You can specify at most 20 DSCP values for a queue in each call. Separate DSCP values with commas (,).
	//
	// This parameter is required.
	Dscps        []*int32 `json:"Dscps,omitempty" xml:"Dscps,omitempty" type:"Repeated"`
	OwnerAccount *string  `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64   `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The description of the queue.
	//
	// This parameter is optional. If you enter a description, it must be 1 to 256 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// desctest
	QosQueueDescription *string `json:"QosQueueDescription,omitempty" xml:"QosQueueDescription,omitempty"`
	// The name of the queue.
	//
	// The name can be empty or 1 to 128 characters in length, and cannot start with http:// or https://.
	//
	// example:
	//
	// nametest
	QosQueueName *string `json:"QosQueueName,omitempty" xml:"QosQueueName,omitempty"`
	// The maximum percentage of inter-region bandwidth that can be allocated to the queue.
	//
	// - Unit: percentage. For example, a value of 20 specifies that the queue can consume at most 20% of inter-region bandwidth.
	//
	// - The sum of the percentage values specified for all queues that belong to the same inter-region connection cannot exceed 100%.
	//
	// example:
	//
	// 20
	RemainBandwidthPercent *string `json:"RemainBandwidthPercent,omitempty" xml:"RemainBandwidthPercent,omitempty"`
	ResourceOwnerAccount   *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId        *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the QoS policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// qos-qdvybn468kaoxx****
	TrafficQosPolicyId *string `json:"TrafficQosPolicyId,omitempty" xml:"TrafficQosPolicyId,omitempty"`
}

func (s CreateCenInterRegionTrafficQosQueueRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCenInterRegionTrafficQosQueueRequest) GoString() string {
	return s.String()
}

func (s *CreateCenInterRegionTrafficQosQueueRequest) GetBandwidth() *int64 {
	return s.Bandwidth
}

func (s *CreateCenInterRegionTrafficQosQueueRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateCenInterRegionTrafficQosQueueRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateCenInterRegionTrafficQosQueueRequest) GetDscps() []*int32 {
	return s.Dscps
}

func (s *CreateCenInterRegionTrafficQosQueueRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateCenInterRegionTrafficQosQueueRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateCenInterRegionTrafficQosQueueRequest) GetQosQueueDescription() *string {
	return s.QosQueueDescription
}

func (s *CreateCenInterRegionTrafficQosQueueRequest) GetQosQueueName() *string {
	return s.QosQueueName
}

func (s *CreateCenInterRegionTrafficQosQueueRequest) GetRemainBandwidthPercent() *string {
	return s.RemainBandwidthPercent
}

func (s *CreateCenInterRegionTrafficQosQueueRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateCenInterRegionTrafficQosQueueRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateCenInterRegionTrafficQosQueueRequest) GetTrafficQosPolicyId() *string {
	return s.TrafficQosPolicyId
}

func (s *CreateCenInterRegionTrafficQosQueueRequest) SetBandwidth(v int64) *CreateCenInterRegionTrafficQosQueueRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosQueueRequest) SetClientToken(v string) *CreateCenInterRegionTrafficQosQueueRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosQueueRequest) SetDryRun(v bool) *CreateCenInterRegionTrafficQosQueueRequest {
	s.DryRun = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosQueueRequest) SetDscps(v []*int32) *CreateCenInterRegionTrafficQosQueueRequest {
	s.Dscps = v
	return s
}

func (s *CreateCenInterRegionTrafficQosQueueRequest) SetOwnerAccount(v string) *CreateCenInterRegionTrafficQosQueueRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosQueueRequest) SetOwnerId(v int64) *CreateCenInterRegionTrafficQosQueueRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosQueueRequest) SetQosQueueDescription(v string) *CreateCenInterRegionTrafficQosQueueRequest {
	s.QosQueueDescription = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosQueueRequest) SetQosQueueName(v string) *CreateCenInterRegionTrafficQosQueueRequest {
	s.QosQueueName = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosQueueRequest) SetRemainBandwidthPercent(v string) *CreateCenInterRegionTrafficQosQueueRequest {
	s.RemainBandwidthPercent = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosQueueRequest) SetResourceOwnerAccount(v string) *CreateCenInterRegionTrafficQosQueueRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosQueueRequest) SetResourceOwnerId(v int64) *CreateCenInterRegionTrafficQosQueueRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosQueueRequest) SetTrafficQosPolicyId(v string) *CreateCenInterRegionTrafficQosQueueRequest {
	s.TrafficQosPolicyId = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosQueueRequest) Validate() error {
	return dara.Validate(s)
}
