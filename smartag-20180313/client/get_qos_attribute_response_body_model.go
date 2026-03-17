// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQosAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorConfigSmartAGCount(v int32) *GetQosAttributeResponseBody
	GetErrorConfigSmartAGCount() *int32
	SetQosCars(v []*GetQosAttributeResponseBodyQosCars) *GetQosAttributeResponseBody
	GetQosCars() []*GetQosAttributeResponseBodyQosCars
	SetQosDescription(v string) *GetQosAttributeResponseBody
	GetQosDescription() *string
	SetQosName(v string) *GetQosAttributeResponseBody
	GetQosName() *string
	SetQosPolicies(v []*GetQosAttributeResponseBodyQosPolicies) *GetQosAttributeResponseBody
	GetQosPolicies() []*GetQosAttributeResponseBodyQosPolicies
	SetRequestId(v string) *GetQosAttributeResponseBody
	GetRequestId() *string
}

type GetQosAttributeResponseBody struct {
	// The number of Smart Access Gateway (SAG) instances that have exceptional configurations.
	//
	// example:
	//
	// 1
	ErrorConfigSmartAGCount *int32 `json:"ErrorConfigSmartAGCount,omitempty" xml:"ErrorConfigSmartAGCount,omitempty"`
	// The traffic throttling rule applied to the QoS policies that have exceptional configurations.
	QosCars []*GetQosAttributeResponseBodyQosCars `json:"QosCars,omitempty" xml:"QosCars,omitempty" type:"Repeated"`
	// The description of the QoS policy.
	//
	// example:
	//
	// test
	QosDescription *string `json:"QosDescription,omitempty" xml:"QosDescription,omitempty"`
	// The name of the QoS policy.
	//
	// example:
	//
	// test
	QosName *string `json:"QosName,omitempty" xml:"QosName,omitempty"`
	// List of QoS policies based on 5-tuple.
	QosPolicies []*GetQosAttributeResponseBodyQosPolicies `json:"QosPolicies,omitempty" xml:"QosPolicies,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 91058E01-1806-45D5-B305-19E4D0A5CE04
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetQosAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQosAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetQosAttributeResponseBody) GetErrorConfigSmartAGCount() *int32 {
	return s.ErrorConfigSmartAGCount
}

func (s *GetQosAttributeResponseBody) GetQosCars() []*GetQosAttributeResponseBodyQosCars {
	return s.QosCars
}

func (s *GetQosAttributeResponseBody) GetQosDescription() *string {
	return s.QosDescription
}

func (s *GetQosAttributeResponseBody) GetQosName() *string {
	return s.QosName
}

func (s *GetQosAttributeResponseBody) GetQosPolicies() []*GetQosAttributeResponseBodyQosPolicies {
	return s.QosPolicies
}

func (s *GetQosAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQosAttributeResponseBody) SetErrorConfigSmartAGCount(v int32) *GetQosAttributeResponseBody {
	s.ErrorConfigSmartAGCount = &v
	return s
}

func (s *GetQosAttributeResponseBody) SetQosCars(v []*GetQosAttributeResponseBodyQosCars) *GetQosAttributeResponseBody {
	s.QosCars = v
	return s
}

func (s *GetQosAttributeResponseBody) SetQosDescription(v string) *GetQosAttributeResponseBody {
	s.QosDescription = &v
	return s
}

func (s *GetQosAttributeResponseBody) SetQosName(v string) *GetQosAttributeResponseBody {
	s.QosName = &v
	return s
}

func (s *GetQosAttributeResponseBody) SetQosPolicies(v []*GetQosAttributeResponseBodyQosPolicies) *GetQosAttributeResponseBody {
	s.QosPolicies = v
	return s
}

func (s *GetQosAttributeResponseBody) SetRequestId(v string) *GetQosAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQosAttributeResponseBody) Validate() error {
	if s.QosCars != nil {
		for _, item := range s.QosCars {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.QosPolicies != nil {
		for _, item := range s.QosPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetQosAttributeResponseBodyQosCars struct {
	// The type of traffic throttling. Valid values:
	//
	// 	- **Absolute**: throttles traffic based on a specific range of bandwidth.
	//
	// 	- **Percent**: throttles traffic based on a specific range of bandwidth percentage.
	//
	// example:
	//
	// Absolute
	LimitType *string `json:"LimitType,omitempty" xml:"LimitType,omitempty"`
	// The maximum bandwidth. Unit: Mbit/s.
	//
	// example:
	//
	// 2
	MaxBandwidthAbs *int32 `json:"MaxBandwidthAbs,omitempty" xml:"MaxBandwidthAbs,omitempty"`
	// The maximum bandwidth percentage that the traffic is throttled to.
	//
	// example:
	//
	// 20
	MaxBandwidthPercent *int32 `json:"MaxBandwidthPercent,omitempty" xml:"MaxBandwidthPercent,omitempty"`
	// The minimum bandwidth. Unit: Mbit/s.
	//
	// example:
	//
	// 1
	MinBandwidthAbs *int32 `json:"MinBandwidthAbs,omitempty" xml:"MinBandwidthAbs,omitempty"`
	// The minimum bandwidth percentage.
	//
	// example:
	//
	// 10
	MinBandwidthPercent *int32 `json:"MinBandwidthPercent,omitempty" xml:"MinBandwidthPercent,omitempty"`
	// Bandwidth type when traffic is throttled to a percentage of the total bandwidth of the network.
	//
	// 	- **CcnBandwidth**: Cloud Connect Network (CCN) bandwidth.
	//
	// 	- **InternetUpBandwidth**: Internet upstream bandwidth.
	//
	// example:
	//
	// InternetUpBandwidth
	PercentSourceType *string `json:"PercentSourceType,omitempty" xml:"PercentSourceType,omitempty"`
	// The priority of the traffic throttling rule.
	//
	// Valid values are from **1*	- to **3**. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The description of the traffic throttling rule.
	//
	// example:
	//
	// test
	QosCarDescription *string `json:"QosCarDescription,omitempty" xml:"QosCarDescription,omitempty"`
	// The ID of the traffic throttling rule.
	//
	// example:
	//
	// qoscar-xir1apa8ayjp56ei****
	QosCarId *string `json:"QosCarId,omitempty" xml:"QosCarId,omitempty"`
	// The name of the traffic throttling rule.
	//
	// example:
	//
	// test
	QosCarName *string `json:"QosCarName,omitempty" xml:"QosCarName,omitempty"`
}

func (s GetQosAttributeResponseBodyQosCars) String() string {
	return dara.Prettify(s)
}

func (s GetQosAttributeResponseBodyQosCars) GoString() string {
	return s.String()
}

func (s *GetQosAttributeResponseBodyQosCars) GetLimitType() *string {
	return s.LimitType
}

func (s *GetQosAttributeResponseBodyQosCars) GetMaxBandwidthAbs() *int32 {
	return s.MaxBandwidthAbs
}

func (s *GetQosAttributeResponseBodyQosCars) GetMaxBandwidthPercent() *int32 {
	return s.MaxBandwidthPercent
}

func (s *GetQosAttributeResponseBodyQosCars) GetMinBandwidthAbs() *int32 {
	return s.MinBandwidthAbs
}

func (s *GetQosAttributeResponseBodyQosCars) GetMinBandwidthPercent() *int32 {
	return s.MinBandwidthPercent
}

func (s *GetQosAttributeResponseBodyQosCars) GetPercentSourceType() *string {
	return s.PercentSourceType
}

func (s *GetQosAttributeResponseBodyQosCars) GetPriority() *int32 {
	return s.Priority
}

func (s *GetQosAttributeResponseBodyQosCars) GetQosCarDescription() *string {
	return s.QosCarDescription
}

func (s *GetQosAttributeResponseBodyQosCars) GetQosCarId() *string {
	return s.QosCarId
}

func (s *GetQosAttributeResponseBodyQosCars) GetQosCarName() *string {
	return s.QosCarName
}

func (s *GetQosAttributeResponseBodyQosCars) SetLimitType(v string) *GetQosAttributeResponseBodyQosCars {
	s.LimitType = &v
	return s
}

func (s *GetQosAttributeResponseBodyQosCars) SetMaxBandwidthAbs(v int32) *GetQosAttributeResponseBodyQosCars {
	s.MaxBandwidthAbs = &v
	return s
}

func (s *GetQosAttributeResponseBodyQosCars) SetMaxBandwidthPercent(v int32) *GetQosAttributeResponseBodyQosCars {
	s.MaxBandwidthPercent = &v
	return s
}

func (s *GetQosAttributeResponseBodyQosCars) SetMinBandwidthAbs(v int32) *GetQosAttributeResponseBodyQosCars {
	s.MinBandwidthAbs = &v
	return s
}

func (s *GetQosAttributeResponseBodyQosCars) SetMinBandwidthPercent(v int32) *GetQosAttributeResponseBodyQosCars {
	s.MinBandwidthPercent = &v
	return s
}

func (s *GetQosAttributeResponseBodyQosCars) SetPercentSourceType(v string) *GetQosAttributeResponseBodyQosCars {
	s.PercentSourceType = &v
	return s
}

func (s *GetQosAttributeResponseBodyQosCars) SetPriority(v int32) *GetQosAttributeResponseBodyQosCars {
	s.Priority = &v
	return s
}

func (s *GetQosAttributeResponseBodyQosCars) SetQosCarDescription(v string) *GetQosAttributeResponseBodyQosCars {
	s.QosCarDescription = &v
	return s
}

func (s *GetQosAttributeResponseBodyQosCars) SetQosCarId(v string) *GetQosAttributeResponseBodyQosCars {
	s.QosCarId = &v
	return s
}

func (s *GetQosAttributeResponseBodyQosCars) SetQosCarName(v string) *GetQosAttributeResponseBodyQosCars {
	s.QosCarName = &v
	return s
}

func (s *GetQosAttributeResponseBodyQosCars) Validate() error {
	return dara.Validate(s)
}

type GetQosAttributeResponseBodyQosPolicies struct {
	// The range of the destination CIDR block.
	//
	// example:
	//
	// 0.0.0.0/0
	DestCidr *string `json:"DestCidr,omitempty" xml:"DestCidr,omitempty"`
	// The range of destination ports.
	//
	// Valid values: **1*	- to **65535*	- and **-1**.
	//
	// Examples of the format of the destination port range:
	//
	// 	- **1/200**: a port range from 1 to 200.
	//
	// 	- **80/80**: port 80.
	//
	// 	- **-1/-1**: all ports.
	//
	// example:
	//
	// -1/-1
	DestPortRange *string `json:"DestPortRange,omitempty" xml:"DestPortRange,omitempty"`
	// The end time of the valid time of the 5-tuple.
	//
	// The time must be in UTC+8.
	//
	// example:
	//
	// 2021-07-29T00:00:00+0800
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The type of the protocol that is applied to the 5-tuple rule.
	//
	// example:
	//
	// ALL
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The priority of the traffic throttling rule that is applied to the 5-tuple.rule.
	//
	// A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The description of the 5-tuple.
	//
	// example:
	//
	// test
	QosPolicieDescription *string `json:"QosPolicieDescription,omitempty" xml:"QosPolicieDescription,omitempty"`
	// The name of the 5-tuple.
	//
	// example:
	//
	// test
	QosPolicieName *string `json:"QosPolicieName,omitempty" xml:"QosPolicieName,omitempty"`
	// The range of the source CIDR block.
	//
	// example:
	//
	// 0.0.0.0/0
	SourceCidr *string `json:"SourceCidr,omitempty" xml:"SourceCidr,omitempty"`
	// The range of source ports.
	//
	// Valid values: **1*	- to **65535*	- and **-1**.
	//
	// Examples of the format of the source port range:
	//
	// 	- **1/200**: a port range from 1 to 200.
	//
	// 	- **80/80**: port 80.
	//
	// 	- **-1/-1**: all ports.
	//
	// example:
	//
	// -1/-1
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
	// The start time of the valid time of the 5-tuple.
	//
	// example:
	//
	// 2021-06-21T00:00:00+0800
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetQosAttributeResponseBodyQosPolicies) String() string {
	return dara.Prettify(s)
}

func (s GetQosAttributeResponseBodyQosPolicies) GoString() string {
	return s.String()
}

func (s *GetQosAttributeResponseBodyQosPolicies) GetDestCidr() *string {
	return s.DestCidr
}

func (s *GetQosAttributeResponseBodyQosPolicies) GetDestPortRange() *string {
	return s.DestPortRange
}

func (s *GetQosAttributeResponseBodyQosPolicies) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetQosAttributeResponseBodyQosPolicies) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *GetQosAttributeResponseBodyQosPolicies) GetPriority() *int32 {
	return s.Priority
}

func (s *GetQosAttributeResponseBodyQosPolicies) GetQosPolicieDescription() *string {
	return s.QosPolicieDescription
}

func (s *GetQosAttributeResponseBodyQosPolicies) GetQosPolicieName() *string {
	return s.QosPolicieName
}

func (s *GetQosAttributeResponseBodyQosPolicies) GetSourceCidr() *string {
	return s.SourceCidr
}

func (s *GetQosAttributeResponseBodyQosPolicies) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *GetQosAttributeResponseBodyQosPolicies) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetQosAttributeResponseBodyQosPolicies) SetDestCidr(v string) *GetQosAttributeResponseBodyQosPolicies {
	s.DestCidr = &v
	return s
}

func (s *GetQosAttributeResponseBodyQosPolicies) SetDestPortRange(v string) *GetQosAttributeResponseBodyQosPolicies {
	s.DestPortRange = &v
	return s
}

func (s *GetQosAttributeResponseBodyQosPolicies) SetEndTime(v int64) *GetQosAttributeResponseBodyQosPolicies {
	s.EndTime = &v
	return s
}

func (s *GetQosAttributeResponseBodyQosPolicies) SetIpProtocol(v string) *GetQosAttributeResponseBodyQosPolicies {
	s.IpProtocol = &v
	return s
}

func (s *GetQosAttributeResponseBodyQosPolicies) SetPriority(v int32) *GetQosAttributeResponseBodyQosPolicies {
	s.Priority = &v
	return s
}

func (s *GetQosAttributeResponseBodyQosPolicies) SetQosPolicieDescription(v string) *GetQosAttributeResponseBodyQosPolicies {
	s.QosPolicieDescription = &v
	return s
}

func (s *GetQosAttributeResponseBodyQosPolicies) SetQosPolicieName(v string) *GetQosAttributeResponseBodyQosPolicies {
	s.QosPolicieName = &v
	return s
}

func (s *GetQosAttributeResponseBodyQosPolicies) SetSourceCidr(v string) *GetQosAttributeResponseBodyQosPolicies {
	s.SourceCidr = &v
	return s
}

func (s *GetQosAttributeResponseBodyQosPolicies) SetSourcePortRange(v string) *GetQosAttributeResponseBodyQosPolicies {
	s.SourcePortRange = &v
	return s
}

func (s *GetQosAttributeResponseBodyQosPolicies) SetStartTime(v int64) *GetQosAttributeResponseBodyQosPolicies {
	s.StartTime = &v
	return s
}

func (s *GetQosAttributeResponseBodyQosPolicies) Validate() error {
	return dara.Validate(s)
}
