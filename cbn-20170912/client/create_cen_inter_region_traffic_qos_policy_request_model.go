// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCenInterRegionTrafficQosPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidthGuaranteeMode(v string) *CreateCenInterRegionTrafficQosPolicyRequest
	GetBandwidthGuaranteeMode() *string
	SetClientToken(v string) *CreateCenInterRegionTrafficQosPolicyRequest
	GetClientToken() *string
	SetConsoleDryRun(v bool) *CreateCenInterRegionTrafficQosPolicyRequest
	GetConsoleDryRun() *bool
	SetDryRun(v bool) *CreateCenInterRegionTrafficQosPolicyRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *CreateCenInterRegionTrafficQosPolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateCenInterRegionTrafficQosPolicyRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateCenInterRegionTrafficQosPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateCenInterRegionTrafficQosPolicyRequest
	GetResourceOwnerId() *int64
	SetTrafficQosPolicyDescription(v string) *CreateCenInterRegionTrafficQosPolicyRequest
	GetTrafficQosPolicyDescription() *string
	SetTrafficQosPolicyName(v string) *CreateCenInterRegionTrafficQosPolicyRequest
	GetTrafficQosPolicyName() *string
	SetTrafficQosQueues(v []*CreateCenInterRegionTrafficQosPolicyRequestTrafficQosQueues) *CreateCenInterRegionTrafficQosPolicyRequest
	GetTrafficQosQueues() []*CreateCenInterRegionTrafficQosPolicyRequestTrafficQosQueues
	SetTransitRouterAttachmentId(v string) *CreateCenInterRegionTrafficQosPolicyRequest
	GetTransitRouterAttachmentId() *string
	SetTransitRouterId(v string) *CreateCenInterRegionTrafficQosPolicyRequest
	GetTransitRouterId() *string
}

type CreateCenInterRegionTrafficQosPolicyRequest struct {
	// The bandwidth guarantee mode. You can specify an absolute bandwidth value or a bandwidth percentage. Valid values:
	//
	// - **byBandwidth**: configures the QoS queue based on an absolute bandwidth value.
	//
	// - **byBandwidthPercent*	- (default): configures the QoS queue based on a bandwidth percentage.
	//
	// example:
	//
	// byBandwidthPercent
	BandwidthGuaranteeMode *string `json:"BandwidthGuaranteeMode,omitempty" xml:"BandwidthGuaranteeMode,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to only precheck the request. Valid values:
	//
	// - **true**: prechecks the request without creating the inter-region traffic QoS policy. The system checks whether required parameters are specified, the request format, and service limits. If the request fails the precheck, the corresponding error is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends the request. If the request passes the precheck, the inter-region traffic QoS policy is created.
	//
	// example:
	//
	// false
	ConsoleDryRun *bool `json:"ConsoleDryRun,omitempty" xml:"ConsoleDryRun,omitempty"`
	// Specifies whether to only precheck the request. Valid values:
	//
	// - **true**: prechecks the request without creating the traffic QoS policy. The system checks whether required parameters are specified, the request format, and service limits. If the request fails the precheck, the corresponding error is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends the request. If the request passes the precheck, the traffic QoS policy is created.
	//
	// example:
	//
	// false
	DryRun               *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The description of the traffic QoS policy.
	//
	// The description must be 1 to 256 characters in length, and can be left empty. It cannot start with http:// or https://.
	//
	// example:
	//
	// desctest
	TrafficQosPolicyDescription *string `json:"TrafficQosPolicyDescription,omitempty" xml:"TrafficQosPolicyDescription,omitempty"`
	// The name of the traffic QoS policy.
	//
	// The name must be 1 to 128 characters in length, and can be left empty. It cannot start with http:// or https://.
	//
	// example:
	//
	// nametest
	TrafficQosPolicyName *string `json:"TrafficQosPolicyName,omitempty" xml:"TrafficQosPolicyName,omitempty"`
	// The information about the queues in the traffic QoS policy.
	//
	// You can create up to three queues. To create more queues, call the CreateCenInterRegionTrafficQosQueue operation.
	TrafficQosQueues []*CreateCenInterRegionTrafficQosPolicyRequestTrafficQosQueues `json:"TrafficQosQueues,omitempty" xml:"TrafficQosQueues,omitempty" type:"Repeated"`
	// The ID of the inter-region connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-attach-r6g0m3epjehw57****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	// The transit router instance ID, which is used for forward and routing.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-8vbuqeo5h5pu3m01d****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s CreateCenInterRegionTrafficQosPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCenInterRegionTrafficQosPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateCenInterRegionTrafficQosPolicyRequest) GetBandwidthGuaranteeMode() *string {
	return s.BandwidthGuaranteeMode
}

func (s *CreateCenInterRegionTrafficQosPolicyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateCenInterRegionTrafficQosPolicyRequest) GetConsoleDryRun() *bool {
	return s.ConsoleDryRun
}

func (s *CreateCenInterRegionTrafficQosPolicyRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateCenInterRegionTrafficQosPolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateCenInterRegionTrafficQosPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateCenInterRegionTrafficQosPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateCenInterRegionTrafficQosPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateCenInterRegionTrafficQosPolicyRequest) GetTrafficQosPolicyDescription() *string {
	return s.TrafficQosPolicyDescription
}

func (s *CreateCenInterRegionTrafficQosPolicyRequest) GetTrafficQosPolicyName() *string {
	return s.TrafficQosPolicyName
}

func (s *CreateCenInterRegionTrafficQosPolicyRequest) GetTrafficQosQueues() []*CreateCenInterRegionTrafficQosPolicyRequestTrafficQosQueues {
	return s.TrafficQosQueues
}

func (s *CreateCenInterRegionTrafficQosPolicyRequest) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *CreateCenInterRegionTrafficQosPolicyRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *CreateCenInterRegionTrafficQosPolicyRequest) SetBandwidthGuaranteeMode(v string) *CreateCenInterRegionTrafficQosPolicyRequest {
	s.BandwidthGuaranteeMode = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosPolicyRequest) SetClientToken(v string) *CreateCenInterRegionTrafficQosPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosPolicyRequest) SetConsoleDryRun(v bool) *CreateCenInterRegionTrafficQosPolicyRequest {
	s.ConsoleDryRun = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosPolicyRequest) SetDryRun(v bool) *CreateCenInterRegionTrafficQosPolicyRequest {
	s.DryRun = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosPolicyRequest) SetOwnerAccount(v string) *CreateCenInterRegionTrafficQosPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosPolicyRequest) SetOwnerId(v int64) *CreateCenInterRegionTrafficQosPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosPolicyRequest) SetResourceOwnerAccount(v string) *CreateCenInterRegionTrafficQosPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosPolicyRequest) SetResourceOwnerId(v int64) *CreateCenInterRegionTrafficQosPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosPolicyRequest) SetTrafficQosPolicyDescription(v string) *CreateCenInterRegionTrafficQosPolicyRequest {
	s.TrafficQosPolicyDescription = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosPolicyRequest) SetTrafficQosPolicyName(v string) *CreateCenInterRegionTrafficQosPolicyRequest {
	s.TrafficQosPolicyName = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosPolicyRequest) SetTrafficQosQueues(v []*CreateCenInterRegionTrafficQosPolicyRequestTrafficQosQueues) *CreateCenInterRegionTrafficQosPolicyRequest {
	s.TrafficQosQueues = v
	return s
}

func (s *CreateCenInterRegionTrafficQosPolicyRequest) SetTransitRouterAttachmentId(v string) *CreateCenInterRegionTrafficQosPolicyRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosPolicyRequest) SetTransitRouterId(v string) *CreateCenInterRegionTrafficQosPolicyRequest {
	s.TransitRouterId = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosPolicyRequest) Validate() error {
	if s.TrafficQosQueues != nil {
		for _, item := range s.TrafficQosQueues {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateCenInterRegionTrafficQosPolicyRequestTrafficQosQueues struct {
	// The absolute value of the inter-region bandwidth that the queue can use. Unit: Mbit/s.
	//
	// You can create up to 10 queues in a traffic QoS policy and specify the inter-region bandwidth that each queue can use.
	//
	// For example, if you enter 1, packets that match the queue can use up to 1 Mbit/s of the inter-region bandwidth.
	//
	// > The sum of the absolute values of the inter-region bandwidth allocated to all queues in a traffic QoS policy cannot exceed the total inter-region bandwidth.
	//
	// example:
	//
	// 1
	Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The DSCP values that match the queue.
	//
	// You can create up to three queues in a traffic QoS policy. Each queue can match up to 60 DSCP values. Separate multiple DSCP values with commas (,).
	Dscps []*int32 `json:"Dscps,omitempty" xml:"Dscps,omitempty" type:"Repeated"`
	// The description of the queue.
	//
	// You can create up to 10 queues in a traffic QoS policy and add a description to each queue.
	//
	// The description must be 1 to 256 characters in length, and can be left empty. It cannot start with http:// or https://.
	//
	// example:
	//
	// desctest
	QosQueueDescription *string `json:"QosQueueDescription,omitempty" xml:"QosQueueDescription,omitempty"`
	// The name of the queue.
	//
	// You can create up to three queues in a traffic QoS policy and specify a name for each queue.
	//
	// The name must be 1 to 128 characters in length, and can be left empty. It cannot start with http:// or https://.
	//
	// example:
	//
	// nametest
	QosQueueName *string `json:"QosQueueName,omitempty" xml:"QosQueueName,omitempty"`
	// The percentage of the inter-region bandwidth that the queue can use.
	//
	// You can create up to 10 queues in a traffic QoS policy and specify the percentage of the inter-region bandwidth that each queue can use.
	//
	// For example, if you enter **1**, packets that match the queue can use up to 1% of the inter-region bandwidth.
	//
	// > The sum of the inter-region bandwidth percentages of all queues in a traffic QoS policy cannot exceed 100%.
	//
	// example:
	//
	// 1
	RemainBandwidthPercent *string `json:"RemainBandwidthPercent,omitempty" xml:"RemainBandwidthPercent,omitempty"`
}

func (s CreateCenInterRegionTrafficQosPolicyRequestTrafficQosQueues) String() string {
	return dara.Prettify(s)
}

func (s CreateCenInterRegionTrafficQosPolicyRequestTrafficQosQueues) GoString() string {
	return s.String()
}

func (s *CreateCenInterRegionTrafficQosPolicyRequestTrafficQosQueues) GetBandwidth() *string {
	return s.Bandwidth
}

func (s *CreateCenInterRegionTrafficQosPolicyRequestTrafficQosQueues) GetDscps() []*int32 {
	return s.Dscps
}

func (s *CreateCenInterRegionTrafficQosPolicyRequestTrafficQosQueues) GetQosQueueDescription() *string {
	return s.QosQueueDescription
}

func (s *CreateCenInterRegionTrafficQosPolicyRequestTrafficQosQueues) GetQosQueueName() *string {
	return s.QosQueueName
}

func (s *CreateCenInterRegionTrafficQosPolicyRequestTrafficQosQueues) GetRemainBandwidthPercent() *string {
	return s.RemainBandwidthPercent
}

func (s *CreateCenInterRegionTrafficQosPolicyRequestTrafficQosQueues) SetBandwidth(v string) *CreateCenInterRegionTrafficQosPolicyRequestTrafficQosQueues {
	s.Bandwidth = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosPolicyRequestTrafficQosQueues) SetDscps(v []*int32) *CreateCenInterRegionTrafficQosPolicyRequestTrafficQosQueues {
	s.Dscps = v
	return s
}

func (s *CreateCenInterRegionTrafficQosPolicyRequestTrafficQosQueues) SetQosQueueDescription(v string) *CreateCenInterRegionTrafficQosPolicyRequestTrafficQosQueues {
	s.QosQueueDescription = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosPolicyRequestTrafficQosQueues) SetQosQueueName(v string) *CreateCenInterRegionTrafficQosPolicyRequestTrafficQosQueues {
	s.QosQueueName = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosPolicyRequestTrafficQosQueues) SetRemainBandwidthPercent(v string) *CreateCenInterRegionTrafficQosPolicyRequestTrafficQosQueues {
	s.RemainBandwidthPercent = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosPolicyRequestTrafficQosQueues) Validate() error {
	return dara.Validate(s)
}
