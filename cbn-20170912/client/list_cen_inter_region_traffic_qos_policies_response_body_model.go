// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCenInterRegionTrafficQosPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListCenInterRegionTrafficQosPoliciesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListCenInterRegionTrafficQosPoliciesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListCenInterRegionTrafficQosPoliciesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListCenInterRegionTrafficQosPoliciesResponseBody
	GetTotalCount() *int32
	SetTrafficQosPolicies(v []*ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies) *ListCenInterRegionTrafficQosPoliciesResponseBody
	GetTrafficQosPolicies() []*ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies
}

type ListCenInterRegionTrafficQosPoliciesResponseBody struct {
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start point of the query.
	//
	// 	- If **NextToken*	- was not returned in the previous query, it indicates that no additional results exist.
	//
	// 	- If **NextToken*	- was returned in the previous query, specify the value to obtain the next set of results.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 113BFD47-63DF-5D9D-972C-033FB9C360CD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// A list of QoS policies.
	TrafficQosPolicies []*ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies `json:"TrafficQosPolicies,omitempty" xml:"TrafficQosPolicies,omitempty" type:"Repeated"`
}

func (s ListCenInterRegionTrafficQosPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCenInterRegionTrafficQosPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBody) GetTrafficQosPolicies() []*ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies {
	return s.TrafficQosPolicies
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBody) SetMaxResults(v int32) *ListCenInterRegionTrafficQosPoliciesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBody) SetNextToken(v string) *ListCenInterRegionTrafficQosPoliciesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBody) SetRequestId(v string) *ListCenInterRegionTrafficQosPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBody) SetTotalCount(v int32) *ListCenInterRegionTrafficQosPoliciesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBody) SetTrafficQosPolicies(v []*ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies) *ListCenInterRegionTrafficQosPoliciesResponseBody {
	s.TrafficQosPolicies = v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBody) Validate() error {
	if s.TrafficQosPolicies != nil {
		for _, item := range s.TrafficQosPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies struct {
	// The guaranteed bandwidth mode.
	//
	// 	- **byBandwidth**: allocates absolute bandwidth values to QoS queues.
	//
	// 	- **byBandwidthPercent**: assigns bandwidth percentages to QoS queues.
	//
	// example:
	//
	// byBandwidthPercent
	BandwidthGuaranteeMode *string `json:"BandwidthGuaranteeMode,omitempty" xml:"BandwidthGuaranteeMode,omitempty"`
	// The description of the QoS policy.
	//
	// example:
	//
	// desctest
	TrafficQosPolicyDescription *string `json:"TrafficQosPolicyDescription,omitempty" xml:"TrafficQosPolicyDescription,omitempty"`
	// The ID of the QoS policy.
	//
	// example:
	//
	// qos-rnghap5gc8155x****
	TrafficQosPolicyId *string `json:"TrafficQosPolicyId,omitempty" xml:"TrafficQosPolicyId,omitempty"`
	// The name of the QoS policy.
	//
	// example:
	//
	// nametest
	TrafficQosPolicyName *string `json:"TrafficQosPolicyName,omitempty" xml:"TrafficQosPolicyName,omitempty"`
	// The status of the QoS policy.
	//
	// 	- **Creating**: The QoS policy is being created.
	//
	// 	- **Active**: The QoS policy is available.
	//
	// 	- **Modifying**: The policy is being modified.
	//
	// 	- **Deleting**: The QoS policy is being deleted.
	//
	// example:
	//
	// Creating
	TrafficQosPolicyStatus *string `json:"TrafficQosPolicyStatus,omitempty" xml:"TrafficQosPolicyStatus,omitempty"`
	// A list of queues.
	TrafficQosQueues []*ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues `json:"TrafficQosQueues,omitempty" xml:"TrafficQosQueues,omitempty" type:"Repeated"`
	// The ID of the network instance connection.
	//
	// example:
	//
	// tr-attach-q7ct7c06jpw***
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	// The ID of the transit router.
	//
	// example:
	//
	// tr-2ze4ta4v32umj0rb***
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies) GoString() string {
	return s.String()
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies) GetBandwidthGuaranteeMode() *string {
	return s.BandwidthGuaranteeMode
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies) GetTrafficQosPolicyDescription() *string {
	return s.TrafficQosPolicyDescription
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies) GetTrafficQosPolicyId() *string {
	return s.TrafficQosPolicyId
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies) GetTrafficQosPolicyName() *string {
	return s.TrafficQosPolicyName
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies) GetTrafficQosPolicyStatus() *string {
	return s.TrafficQosPolicyStatus
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies) GetTrafficQosQueues() []*ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues {
	return s.TrafficQosQueues
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies) SetBandwidthGuaranteeMode(v string) *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies {
	s.BandwidthGuaranteeMode = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies) SetTrafficQosPolicyDescription(v string) *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies {
	s.TrafficQosPolicyDescription = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies) SetTrafficQosPolicyId(v string) *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies {
	s.TrafficQosPolicyId = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies) SetTrafficQosPolicyName(v string) *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies {
	s.TrafficQosPolicyName = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies) SetTrafficQosPolicyStatus(v string) *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies {
	s.TrafficQosPolicyStatus = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies) SetTrafficQosQueues(v []*ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues) *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies {
	s.TrafficQosQueues = v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies) SetTransitRouterAttachmentId(v string) *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies) SetTransitRouterId(v string) *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies {
	s.TransitRouterId = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPolicies) Validate() error {
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

type ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues struct {
	// If the QoS queues are assigned absolute bandwidth values, this parameter indicates the absolute bandwidth value that is allocated to the queue.
	//
	// example:
	//
	// 1
	Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The differentiated services code point (DSCP) value that is used to match packets.
	Dscps []*int32 `json:"Dscps,omitempty" xml:"Dscps,omitempty" type:"Repeated"`
	// The actual bandwidth value of the current queue.
	//
	// example:
	//
	// 1.35
	EffectiveBandwidth *string `json:"EffectiveBandwidth,omitempty" xml:"EffectiveBandwidth,omitempty"`
	// The description of the queue.
	//
	// example:
	//
	// desctest
	QosQueueDescription *string `json:"QosQueueDescription,omitempty" xml:"QosQueueDescription,omitempty"`
	// The queue ID.
	//
	// example:
	//
	// qos-queue-njcrmr9fiu1jii****
	QosQueueId *string `json:"QosQueueId,omitempty" xml:"QosQueueId,omitempty"`
	// The name of the queue.
	//
	// example:
	//
	// namtest
	QosQueueName *string `json:"QosQueueName,omitempty" xml:"QosQueueName,omitempty"`
	// If the QoS queues are assigned bandwidth percentages, this parameter indicates the percentage of bandwidth that is allocated to the queue.
	//
	// example:
	//
	// 1
	RemainBandwidthPercent *int32 `json:"RemainBandwidthPercent,omitempty" xml:"RemainBandwidthPercent,omitempty"`
}

func (s ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues) String() string {
	return dara.Prettify(s)
}

func (s ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues) GoString() string {
	return s.String()
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues) GetBandwidth() *string {
	return s.Bandwidth
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues) GetDscps() []*int32 {
	return s.Dscps
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues) GetEffectiveBandwidth() *string {
	return s.EffectiveBandwidth
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues) GetQosQueueDescription() *string {
	return s.QosQueueDescription
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues) GetQosQueueId() *string {
	return s.QosQueueId
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues) GetQosQueueName() *string {
	return s.QosQueueName
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues) GetRemainBandwidthPercent() *int32 {
	return s.RemainBandwidthPercent
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues) SetBandwidth(v string) *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues {
	s.Bandwidth = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues) SetDscps(v []*int32) *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues {
	s.Dscps = v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues) SetEffectiveBandwidth(v string) *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues {
	s.EffectiveBandwidth = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues) SetQosQueueDescription(v string) *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues {
	s.QosQueueDescription = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues) SetQosQueueId(v string) *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues {
	s.QosQueueId = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues) SetQosQueueName(v string) *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues {
	s.QosQueueName = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues) SetRemainBandwidthPercent(v int32) *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues {
	s.RemainBandwidthPercent = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesResponseBodyTrafficQosPoliciesTrafficQosQueues) Validate() error {
	return dara.Validate(s)
}
