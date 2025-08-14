// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCenInterRegionTrafficQosQueuesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListCenInterRegionTrafficQosQueuesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListCenInterRegionTrafficQosQueuesResponseBody
	GetRequestId() *string
	SetTrafficQosQueues(v []*ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) *ListCenInterRegionTrafficQosQueuesResponseBody
	GetTrafficQosQueues() []*ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues
}

type ListCenInterRegionTrafficQosQueuesResponseBody struct {
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// 0151fa6aa1ed****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1D1E15D2-416D-54F3-BDD9-BC27DE4C6352
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the QoS queue.
	TrafficQosQueues []*ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues `json:"TrafficQosQueues,omitempty" xml:"TrafficQosQueues,omitempty" type:"Repeated"`
}

func (s ListCenInterRegionTrafficQosQueuesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCenInterRegionTrafficQosQueuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBody) GetTrafficQosQueues() []*ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues {
	return s.TrafficQosQueues
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBody) SetNextToken(v string) *ListCenInterRegionTrafficQosQueuesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBody) SetRequestId(v string) *ListCenInterRegionTrafficQosQueuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBody) SetTrafficQosQueues(v []*ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) *ListCenInterRegionTrafficQosQueuesResponseBody {
	s.TrafficQosQueues = v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues struct {
	// The absolute bandwidth value that can be allocated to the current queue.
	//
	// A value of **1*	- indicates that the QoS queue can consume at most 1 Mbit/s of inter-region bandwidth.
	//
	// example:
	//
	// 1
	Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The Differentiated Services Code Point (DSCP) value that matches the current QoS queue.
	Dscps []*int32 `json:"Dscps,omitempty" xml:"Dscps,omitempty" type:"Repeated"`
	// The actual bandwidth of the current queue.
	//
	// example:
	//
	// 1.35
	EffectiveBandwidth *string `json:"EffectiveBandwidth,omitempty" xml:"EffectiveBandwidth,omitempty"`
	// The percentage of bandwidth that can be allocated to the current queue.
	//
	// A value of **1*	- indicates that the QoS queue can consume at most 1% of the inter-region bandwidth.
	//
	// example:
	//
	// 1
	RemainBandwidthPercent *int32 `json:"RemainBandwidthPercent,omitempty" xml:"RemainBandwidthPercent,omitempty"`
	// The status of the QoS queue. Valid values:
	//
	// 	- **Creating**
	//
	// 	- **Active**
	//
	// 	- **Deleting**
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the QoS policy.
	//
	// example:
	//
	// qos-fv2qq9yqrsjowp****
	TrafficQosPolicyId *string `json:"TrafficQosPolicyId,omitempty" xml:"TrafficQosPolicyId,omitempty"`
	// The description of the QoS queue.
	//
	// example:
	//
	// qosQueueDescription
	TrafficQosQueueDescription *string `json:"TrafficQosQueueDescription,omitempty" xml:"TrafficQosQueueDescription,omitempty"`
	// The ID of the QoS queue.
	//
	// example:
	//
	// qos-queue-siakjb2nn9gz5z****
	TrafficQosQueueId *string `json:"TrafficQosQueueId,omitempty" xml:"TrafficQosQueueId,omitempty"`
	// The name of the QoS queue.
	//
	// example:
	//
	// qosQueueName
	TrafficQosQueueName *string `json:"TrafficQosQueueName,omitempty" xml:"TrafficQosQueueName,omitempty"`
	// The ID of the inter-region connection.
	//
	// example:
	//
	// tr-attach-nzrcv25d7ezt23****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	// The ID of the transit router.
	//
	// example:
	//
	// tr-p0wwagjv6fvxt4b7y****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) String() string {
	return dara.Prettify(s)
}

func (s ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) GoString() string {
	return s.String()
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) GetBandwidth() *string {
	return s.Bandwidth
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) GetDscps() []*int32 {
	return s.Dscps
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) GetEffectiveBandwidth() *string {
	return s.EffectiveBandwidth
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) GetRemainBandwidthPercent() *int32 {
	return s.RemainBandwidthPercent
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) GetStatus() *string {
	return s.Status
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) GetTrafficQosPolicyId() *string {
	return s.TrafficQosPolicyId
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) GetTrafficQosQueueDescription() *string {
	return s.TrafficQosQueueDescription
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) GetTrafficQosQueueId() *string {
	return s.TrafficQosQueueId
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) GetTrafficQosQueueName() *string {
	return s.TrafficQosQueueName
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) SetBandwidth(v string) *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues {
	s.Bandwidth = &v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) SetDscps(v []*int32) *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues {
	s.Dscps = v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) SetEffectiveBandwidth(v string) *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues {
	s.EffectiveBandwidth = &v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) SetRemainBandwidthPercent(v int32) *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues {
	s.RemainBandwidthPercent = &v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) SetStatus(v string) *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues {
	s.Status = &v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) SetTrafficQosPolicyId(v string) *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues {
	s.TrafficQosPolicyId = &v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) SetTrafficQosQueueDescription(v string) *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues {
	s.TrafficQosQueueDescription = &v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) SetTrafficQosQueueId(v string) *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues {
	s.TrafficQosQueueId = &v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) SetTrafficQosQueueName(v string) *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues {
	s.TrafficQosQueueName = &v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) SetTransitRouterAttachmentId(v string) *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) SetTransitRouterId(v string) *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues {
	s.TransitRouterId = &v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) Validate() error {
	return dara.Validate(s)
}
