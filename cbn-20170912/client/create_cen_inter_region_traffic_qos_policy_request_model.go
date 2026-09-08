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
	// The bandwidth guarantee mode. You can configure QoS queues based on absolute bandwidth values or bandwidth percentages. Valid values:
	//
	// - **byBandwidth**: configures QoS queues based on absolute bandwidth values.
	//
	// - **byBandwidthPercent*	- (default): configures QoS queues based on bandwidth percentages.
	//
	// example:
	//
	// byBandwidthPercent
	BandwidthGuaranteeMode *string `json:"BandwidthGuaranteeMode,omitempty" xml:"BandwidthGuaranteeMode,omitempty"`
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
	// - **true**: performs a dry run. The system checks the required parameters, request syntax, and business restrictions. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): performs a dry run and sends the request. If the request passes the dry run, the cross-region traffic scheduling policy is created.
	//
	// example:
	//
	// false
	ConsoleDryRun *bool `json:"ConsoleDryRun,omitempty" xml:"ConsoleDryRun,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run. The system checks the required parameters, request syntax, and business restrictions. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): performs a dry run and sends the request. If the request passes the dry run, the traffic scheduling policy is created.
	//
	// example:
	//
	// false
	DryRun               *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The description of the traffic scheduling policy.
	//
	// The description can be empty or 1 to 256 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// desctest
	TrafficQosPolicyDescription *string `json:"TrafficQosPolicyDescription,omitempty" xml:"TrafficQosPolicyDescription,omitempty"`
	// The name of the traffic scheduling policy.
	//
	// The name can be empty or 1 to 128 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// nametest
	TrafficQosPolicyName *string `json:"TrafficQosPolicyName,omitempty" xml:"TrafficQosPolicyName,omitempty"`
	// The queue information of the traffic scheduling policy.
	//
	// You can create up to 3 queues. To create more queues, call CreateCenInterRegionTrafficQosQueue.
	TrafficQosQueues []*CreateCenInterRegionTrafficQosPolicyRequestTrafficQosQueues `json:"TrafficQosQueues,omitempty" xml:"TrafficQosQueues,omitempty" type:"Repeated"`
	// The ID of the inter-region connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-attach-r6g0m3epjehw57****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	// The ID of the transit router instance.
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
	// The absolute bandwidth value that the queue can use for inter-region traffic, in Mbit/s.
	//
	// A traffic scheduling policy supports up to 10 queues, and each queue can be assigned an absolute bandwidth value.
	//
	// For example, if you enter 1, traffic that matches the queue can use up to 1 Mbit/s of inter-region bandwidth.
	//
	// > The sum of absolute bandwidth values of all allocated queues in a traffic scheduling policy cannot exceed the inter-region bandwidth value.
	//
	// example:
	//
	// 1
	Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The DSCP values that the queue matches.
	//
	// A traffic scheduling policy supports up to 3 queues, and each queue can match up to 60 DSCP values. Separate multiple DSCP values with commas (,).
	Dscps []*int32 `json:"Dscps,omitempty" xml:"Dscps,omitempty" type:"Repeated"`
	// The description of the queue.
	//
	// A traffic scheduling policy supports up to 10 queues, and each queue can have a description.
	//
	// The description can be empty or 1 to 256 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// desctest
	QosQueueDescription *string `json:"QosQueueDescription,omitempty" xml:"QosQueueDescription,omitempty"`
	// The name of the queue.
	//
	// A traffic scheduling policy supports up to 3 queues, and each queue can be assigned a name.
	//
	// The name can be empty or 1 to 128 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// nametest
	QosQueueName *string `json:"QosQueueName,omitempty" xml:"QosQueueName,omitempty"`
	// The percentage of inter-region bandwidth that the queue can use.
	//
	// A traffic scheduling policy supports up to 10 queues, and each queue can be assigned a percentage of inter-region bandwidth.
	//
	// For example, if you enter **1**, traffic that matches the queue can use up to 1% of the inter-region bandwidth.
	//
	// > The sum of bandwidth percentages of all queues in a traffic scheduling policy cannot exceed 100%.
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
