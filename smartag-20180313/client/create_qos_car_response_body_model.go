// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQosCarResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateQosCarResponseBody
	GetDescription() *string
	SetLimitType(v string) *CreateQosCarResponseBody
	GetLimitType() *string
	SetMaxBandwidthAbs(v int32) *CreateQosCarResponseBody
	GetMaxBandwidthAbs() *int32
	SetMaxBandwidthPercent(v int32) *CreateQosCarResponseBody
	GetMaxBandwidthPercent() *int32
	SetMinBandwidthAbs(v int32) *CreateQosCarResponseBody
	GetMinBandwidthAbs() *int32
	SetMinBandwidthPercent(v int32) *CreateQosCarResponseBody
	GetMinBandwidthPercent() *int32
	SetPercentSourceType(v string) *CreateQosCarResponseBody
	GetPercentSourceType() *string
	SetPriority(v int32) *CreateQosCarResponseBody
	GetPriority() *int32
	SetQosCarId(v string) *CreateQosCarResponseBody
	GetQosCarId() *string
	SetQosId(v string) *CreateQosCarResponseBody
	GetQosId() *string
	SetRequestId(v string) *CreateQosCarResponseBody
	GetRequestId() *string
}

type CreateQosCarResponseBody struct {
	// The description of the traffic throttling rule.
	//
	// example:
	//
	// Qosdesc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The type of the traffic throttling rule. Valid values:
	//
	// 	- **Absolute**: throttles traffic based on a specific range of bandwidth.
	//
	// 	- **Percent**: throttles traffic based on a specific range of bandwidth percentage.
	//
	// example:
	//
	// Percent
	LimitType *string `json:"LimitType,omitempty" xml:"LimitType,omitempty"`
	// The maximum bandwidth value. Unit: Mbit/s.
	//
	// This parameter is returned when **LimitType*	- is set to **Absolute**.
	//
	// example:
	//
	// 6
	MaxBandwidthAbs *int32 `json:"MaxBandwidthAbs,omitempty" xml:"MaxBandwidthAbs,omitempty"`
	// The maximum bandwidth percentage. Unit: percent (%).
	//
	// example:
	//
	// 90
	MaxBandwidthPercent *int32 `json:"MaxBandwidthPercent,omitempty" xml:"MaxBandwidthPercent,omitempty"`
	// The minimum bandwidth value. Unit: Mbit/s.
	//
	// This parameter is returned when **LimitType*	- is set to **Absolute**.
	//
	// example:
	//
	// 2
	MinBandwidthAbs *int32 `json:"MinBandwidthAbs,omitempty" xml:"MinBandwidthAbs,omitempty"`
	// The minimum bandwidth percentage. Unit: percent (%).
	//
	// example:
	//
	// 20
	MinBandwidthPercent *int32 `json:"MinBandwidthPercent,omitempty" xml:"MinBandwidthPercent,omitempty"`
	// The type of bandwidth when traffic is throttled based on bandwidth percentage. Valid values:
	//
	// 	- **CcnBandwidth**: CCN bandwidth
	//
	// 	- **InternetUpBandwidth**: total Internet bandwidth
	//
	// example:
	//
	// CcnBandwidth
	PercentSourceType *string `json:"PercentSourceType,omitempty" xml:"PercentSourceType,omitempty"`
	// The priority value of the traffic throttling rule.
	//
	// example:
	//
	// 2
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the traffic throttling rule.
	//
	// example:
	//
	// qoscar-n5k8g97lihlph****
	QosCarId *string `json:"QosCarId,omitempty" xml:"QosCarId,omitempty"`
	// The ID of the QoS policy.
	//
	// example:
	//
	// qos-xitd8690ucu8ro****
	QosId *string `json:"QosId,omitempty" xml:"QosId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// AC13E8FF-4D61-40AD-868E-817F2D3AC86A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateQosCarResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateQosCarResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQosCarResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CreateQosCarResponseBody) GetLimitType() *string {
	return s.LimitType
}

func (s *CreateQosCarResponseBody) GetMaxBandwidthAbs() *int32 {
	return s.MaxBandwidthAbs
}

func (s *CreateQosCarResponseBody) GetMaxBandwidthPercent() *int32 {
	return s.MaxBandwidthPercent
}

func (s *CreateQosCarResponseBody) GetMinBandwidthAbs() *int32 {
	return s.MinBandwidthAbs
}

func (s *CreateQosCarResponseBody) GetMinBandwidthPercent() *int32 {
	return s.MinBandwidthPercent
}

func (s *CreateQosCarResponseBody) GetPercentSourceType() *string {
	return s.PercentSourceType
}

func (s *CreateQosCarResponseBody) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateQosCarResponseBody) GetQosCarId() *string {
	return s.QosCarId
}

func (s *CreateQosCarResponseBody) GetQosId() *string {
	return s.QosId
}

func (s *CreateQosCarResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateQosCarResponseBody) SetDescription(v string) *CreateQosCarResponseBody {
	s.Description = &v
	return s
}

func (s *CreateQosCarResponseBody) SetLimitType(v string) *CreateQosCarResponseBody {
	s.LimitType = &v
	return s
}

func (s *CreateQosCarResponseBody) SetMaxBandwidthAbs(v int32) *CreateQosCarResponseBody {
	s.MaxBandwidthAbs = &v
	return s
}

func (s *CreateQosCarResponseBody) SetMaxBandwidthPercent(v int32) *CreateQosCarResponseBody {
	s.MaxBandwidthPercent = &v
	return s
}

func (s *CreateQosCarResponseBody) SetMinBandwidthAbs(v int32) *CreateQosCarResponseBody {
	s.MinBandwidthAbs = &v
	return s
}

func (s *CreateQosCarResponseBody) SetMinBandwidthPercent(v int32) *CreateQosCarResponseBody {
	s.MinBandwidthPercent = &v
	return s
}

func (s *CreateQosCarResponseBody) SetPercentSourceType(v string) *CreateQosCarResponseBody {
	s.PercentSourceType = &v
	return s
}

func (s *CreateQosCarResponseBody) SetPriority(v int32) *CreateQosCarResponseBody {
	s.Priority = &v
	return s
}

func (s *CreateQosCarResponseBody) SetQosCarId(v string) *CreateQosCarResponseBody {
	s.QosCarId = &v
	return s
}

func (s *CreateQosCarResponseBody) SetQosId(v string) *CreateQosCarResponseBody {
	s.QosId = &v
	return s
}

func (s *CreateQosCarResponseBody) SetRequestId(v string) *CreateQosCarResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateQosCarResponseBody) Validate() error {
	return dara.Validate(s)
}
