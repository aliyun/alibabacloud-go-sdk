// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTransitRouteTableAggregationDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeTransitRouteTableAggregationDetailResponseBody
	GetCount() *int32
	SetData(v []*DescribeTransitRouteTableAggregationDetailResponseBodyData) *DescribeTransitRouteTableAggregationDetailResponseBody
	GetData() []*DescribeTransitRouteTableAggregationDetailResponseBodyData
	SetRequestId(v string) *DescribeTransitRouteTableAggregationDetailResponseBody
	GetRequestId() *string
	SetTotal(v int32) *DescribeTransitRouteTableAggregationDetailResponseBody
	GetTotal() *int32
}

type DescribeTransitRouteTableAggregationDetailResponseBody struct {
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The configuration of the aggregate route.
	Data []*DescribeTransitRouteTableAggregationDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 0C2EE7A8-74D4-4081-8236-CEBDE3BBCF50
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeTransitRouteTableAggregationDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransitRouteTableAggregationDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTransitRouteTableAggregationDetailResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeTransitRouteTableAggregationDetailResponseBody) GetData() []*DescribeTransitRouteTableAggregationDetailResponseBodyData {
	return s.Data
}

func (s *DescribeTransitRouteTableAggregationDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTransitRouteTableAggregationDetailResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeTransitRouteTableAggregationDetailResponseBody) SetCount(v int32) *DescribeTransitRouteTableAggregationDetailResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeTransitRouteTableAggregationDetailResponseBody) SetData(v []*DescribeTransitRouteTableAggregationDetailResponseBodyData) *DescribeTransitRouteTableAggregationDetailResponseBody {
	s.Data = v
	return s
}

func (s *DescribeTransitRouteTableAggregationDetailResponseBody) SetRequestId(v string) *DescribeTransitRouteTableAggregationDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTransitRouteTableAggregationDetailResponseBody) SetTotal(v int32) *DescribeTransitRouteTableAggregationDetailResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeTransitRouteTableAggregationDetailResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTransitRouteTableAggregationDetailResponseBodyData struct {
	// The error message returned if the configuration of the aggregate route fails.
	//
	// example:
	//
	// desctest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the virtual private cloud (VPC) for which the aggregate route is configured.
	//
	// example:
	//
	// vpc-6eh7fp9hdqa2wv85t****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The status of the aggregate route. Valid values:
	//
	// 	- **Configured**: The aggregate route is advertised to the VPC.
	//
	// 	- **Configuring**: The aggregate route is being advertised.
	//
	// 	- **ConfigFailed**: The aggregate route failed to be advertised.
	//
	// 	- **PartialConfigured**: Failed to advertise the aggregate route to some VPCs.
	//
	// 	- **Deleting**: The aggregate route is being deleted.
	//
	// example:
	//
	// Configured
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeTransitRouteTableAggregationDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransitRouteTableAggregationDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeTransitRouteTableAggregationDetailResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *DescribeTransitRouteTableAggregationDetailResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeTransitRouteTableAggregationDetailResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DescribeTransitRouteTableAggregationDetailResponseBodyData) SetDescription(v string) *DescribeTransitRouteTableAggregationDetailResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeTransitRouteTableAggregationDetailResponseBodyData) SetInstanceId(v string) *DescribeTransitRouteTableAggregationDetailResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *DescribeTransitRouteTableAggregationDetailResponseBodyData) SetStatus(v string) *DescribeTransitRouteTableAggregationDetailResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeTransitRouteTableAggregationDetailResponseBodyData) Validate() error {
	return dara.Validate(s)
}
