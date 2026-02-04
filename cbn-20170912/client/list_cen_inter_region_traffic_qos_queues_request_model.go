// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCenInterRegionTrafficQosQueuesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEffectiveBandwidthFilter(v *ListCenInterRegionTrafficQosQueuesRequestEffectiveBandwidthFilter) *ListCenInterRegionTrafficQosQueuesRequest
	GetEffectiveBandwidthFilter() *ListCenInterRegionTrafficQosQueuesRequestEffectiveBandwidthFilter
	SetMaxResults(v int32) *ListCenInterRegionTrafficQosQueuesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListCenInterRegionTrafficQosQueuesRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListCenInterRegionTrafficQosQueuesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListCenInterRegionTrafficQosQueuesRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ListCenInterRegionTrafficQosQueuesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListCenInterRegionTrafficQosQueuesRequest
	GetResourceOwnerId() *int64
	SetTrafficQosPolicyId(v string) *ListCenInterRegionTrafficQosQueuesRequest
	GetTrafficQosPolicyId() *string
	SetTrafficQosQueueDescription(v string) *ListCenInterRegionTrafficQosQueuesRequest
	GetTrafficQosQueueDescription() *string
	SetTrafficQosQueueId(v string) *ListCenInterRegionTrafficQosQueuesRequest
	GetTrafficQosQueueId() *string
	SetTrafficQosQueueName(v string) *ListCenInterRegionTrafficQosQueuesRequest
	GetTrafficQosQueueName() *string
	SetTransitRouterAttachmentId(v string) *ListCenInterRegionTrafficQosQueuesRequest
	GetTransitRouterAttachmentId() *string
	SetTransitRouterId(v string) *ListCenInterRegionTrafficQosQueuesRequest
	GetTransitRouterId() *string
}

type ListCenInterRegionTrafficQosQueuesRequest struct {
	// The filter works based on the actual bandwidth. Enter a positive integer. Unit: Mbit/s.
	EffectiveBandwidthFilter *ListCenInterRegionTrafficQosQueuesRequestEffectiveBandwidthFilter `json:"EffectiveBandwidthFilter,omitempty" xml:"EffectiveBandwidthFilter,omitempty" type:"Struct"`
	// The number of entries to return on each page. Valid values: **1*	- to **100**. Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// 2ca1ed1573cb****
	NextToken            *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the QoS policy.
	//
	// example:
	//
	// qos-rnghap5gc8155x****
	TrafficQosPolicyId *string `json:"TrafficQosPolicyId,omitempty" xml:"TrafficQosPolicyId,omitempty"`
	// The description of the QoS queue.
	//
	// This parameter is optional. If you enter a description, it must be 1 to 256 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// qosQueueDescription
	TrafficQosQueueDescription *string `json:"TrafficQosQueueDescription,omitempty" xml:"TrafficQosQueueDescription,omitempty"`
	// The ID of the queue.
	//
	// example:
	//
	// qos-queue-siakjb2nn9gz5z****
	TrafficQosQueueId *string `json:"TrafficQosQueueId,omitempty" xml:"TrafficQosQueueId,omitempty"`
	// The name of the QoS queue.
	//
	// The name must be 1 to 128 characters in length, and cannot start with http:// or https://.
	//
	// example:
	//
	// qosQueueName
	TrafficQosQueueName *string `json:"TrafficQosQueueName,omitempty" xml:"TrafficQosQueueName,omitempty"`
	// The ID of the inter-region connection.
	//
	// example:
	//
	// tr-attach-a6p8voaodog5c0****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	// The ID of the transit router.
	//
	// example:
	//
	// tr-bp1rmwxnk221e3fas****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s ListCenInterRegionTrafficQosQueuesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCenInterRegionTrafficQosQueuesRequest) GoString() string {
	return s.String()
}

func (s *ListCenInterRegionTrafficQosQueuesRequest) GetEffectiveBandwidthFilter() *ListCenInterRegionTrafficQosQueuesRequestEffectiveBandwidthFilter {
	return s.EffectiveBandwidthFilter
}

func (s *ListCenInterRegionTrafficQosQueuesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCenInterRegionTrafficQosQueuesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCenInterRegionTrafficQosQueuesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListCenInterRegionTrafficQosQueuesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListCenInterRegionTrafficQosQueuesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListCenInterRegionTrafficQosQueuesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListCenInterRegionTrafficQosQueuesRequest) GetTrafficQosPolicyId() *string {
	return s.TrafficQosPolicyId
}

func (s *ListCenInterRegionTrafficQosQueuesRequest) GetTrafficQosQueueDescription() *string {
	return s.TrafficQosQueueDescription
}

func (s *ListCenInterRegionTrafficQosQueuesRequest) GetTrafficQosQueueId() *string {
	return s.TrafficQosQueueId
}

func (s *ListCenInterRegionTrafficQosQueuesRequest) GetTrafficQosQueueName() *string {
	return s.TrafficQosQueueName
}

func (s *ListCenInterRegionTrafficQosQueuesRequest) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *ListCenInterRegionTrafficQosQueuesRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *ListCenInterRegionTrafficQosQueuesRequest) SetEffectiveBandwidthFilter(v *ListCenInterRegionTrafficQosQueuesRequestEffectiveBandwidthFilter) *ListCenInterRegionTrafficQosQueuesRequest {
	s.EffectiveBandwidthFilter = v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesRequest) SetMaxResults(v int32) *ListCenInterRegionTrafficQosQueuesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesRequest) SetNextToken(v string) *ListCenInterRegionTrafficQosQueuesRequest {
	s.NextToken = &v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesRequest) SetOwnerAccount(v string) *ListCenInterRegionTrafficQosQueuesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesRequest) SetOwnerId(v int64) *ListCenInterRegionTrafficQosQueuesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesRequest) SetResourceOwnerAccount(v string) *ListCenInterRegionTrafficQosQueuesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesRequest) SetResourceOwnerId(v int64) *ListCenInterRegionTrafficQosQueuesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesRequest) SetTrafficQosPolicyId(v string) *ListCenInterRegionTrafficQosQueuesRequest {
	s.TrafficQosPolicyId = &v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesRequest) SetTrafficQosQueueDescription(v string) *ListCenInterRegionTrafficQosQueuesRequest {
	s.TrafficQosQueueDescription = &v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesRequest) SetTrafficQosQueueId(v string) *ListCenInterRegionTrafficQosQueuesRequest {
	s.TrafficQosQueueId = &v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesRequest) SetTrafficQosQueueName(v string) *ListCenInterRegionTrafficQosQueuesRequest {
	s.TrafficQosQueueName = &v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesRequest) SetTransitRouterAttachmentId(v string) *ListCenInterRegionTrafficQosQueuesRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesRequest) SetTransitRouterId(v string) *ListCenInterRegionTrafficQosQueuesRequest {
	s.TransitRouterId = &v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesRequest) Validate() error {
	if s.EffectiveBandwidthFilter != nil {
		if err := s.EffectiveBandwidthFilter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCenInterRegionTrafficQosQueuesRequestEffectiveBandwidthFilter struct {
	// The actual bandwidth is equal to or larger than the specified value.
	//
	// example:
	//
	// 50
	Gte *int64 `json:"Gte,omitempty" xml:"Gte,omitempty"`
	// The actual bandwidth is equal to or smaller than the specified value.
	//
	// example:
	//
	// 20
	Lte *int64 `json:"Lte,omitempty" xml:"Lte,omitempty"`
}

func (s ListCenInterRegionTrafficQosQueuesRequestEffectiveBandwidthFilter) String() string {
	return dara.Prettify(s)
}

func (s ListCenInterRegionTrafficQosQueuesRequestEffectiveBandwidthFilter) GoString() string {
	return s.String()
}

func (s *ListCenInterRegionTrafficQosQueuesRequestEffectiveBandwidthFilter) GetGte() *int64 {
	return s.Gte
}

func (s *ListCenInterRegionTrafficQosQueuesRequestEffectiveBandwidthFilter) GetLte() *int64 {
	return s.Lte
}

func (s *ListCenInterRegionTrafficQosQueuesRequestEffectiveBandwidthFilter) SetGte(v int64) *ListCenInterRegionTrafficQosQueuesRequestEffectiveBandwidthFilter {
	s.Gte = &v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesRequestEffectiveBandwidthFilter) SetLte(v int64) *ListCenInterRegionTrafficQosQueuesRequestEffectiveBandwidthFilter {
	s.Lte = &v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesRequestEffectiveBandwidthFilter) Validate() error {
	return dara.Validate(s)
}
