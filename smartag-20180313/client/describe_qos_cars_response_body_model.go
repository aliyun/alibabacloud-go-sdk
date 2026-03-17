// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQosCarsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeQosCarsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeQosCarsResponseBody
	GetPageSize() *int32
	SetQosCars(v *DescribeQosCarsResponseBodyQosCars) *DescribeQosCarsResponseBody
	GetQosCars() *DescribeQosCarsResponseBodyQosCars
	SetRequestId(v string) *DescribeQosCarsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeQosCarsResponseBody
	GetTotalCount() *int32
}

type DescribeQosCarsResponseBody struct {
	// The number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 10
	PageSize *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QosCars  *DescribeQosCarsResponseBodyQosCars `json:"QosCars,omitempty" xml:"QosCars,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// B7B758A9-009E-4C9D-9618-714EAE8BA5E9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeQosCarsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeQosCarsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeQosCarsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeQosCarsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeQosCarsResponseBody) GetQosCars() *DescribeQosCarsResponseBodyQosCars {
	return s.QosCars
}

func (s *DescribeQosCarsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeQosCarsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeQosCarsResponseBody) SetPageNumber(v int32) *DescribeQosCarsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeQosCarsResponseBody) SetPageSize(v int32) *DescribeQosCarsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeQosCarsResponseBody) SetQosCars(v *DescribeQosCarsResponseBodyQosCars) *DescribeQosCarsResponseBody {
	s.QosCars = v
	return s
}

func (s *DescribeQosCarsResponseBody) SetRequestId(v string) *DescribeQosCarsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeQosCarsResponseBody) SetTotalCount(v int32) *DescribeQosCarsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeQosCarsResponseBody) Validate() error {
	if s.QosCars != nil {
		if err := s.QosCars.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeQosCarsResponseBodyQosCars struct {
	QosCar []*DescribeQosCarsResponseBodyQosCarsQosCar `json:"QosCar,omitempty" xml:"QosCar,omitempty" type:"Repeated"`
}

func (s DescribeQosCarsResponseBodyQosCars) String() string {
	return dara.Prettify(s)
}

func (s DescribeQosCarsResponseBodyQosCars) GoString() string {
	return s.String()
}

func (s *DescribeQosCarsResponseBodyQosCars) GetQosCar() []*DescribeQosCarsResponseBodyQosCarsQosCar {
	return s.QosCar
}

func (s *DescribeQosCarsResponseBodyQosCars) SetQosCar(v []*DescribeQosCarsResponseBodyQosCarsQosCar) *DescribeQosCarsResponseBodyQosCars {
	s.QosCar = v
	return s
}

func (s *DescribeQosCarsResponseBodyQosCars) Validate() error {
	if s.QosCar != nil {
		for _, item := range s.QosCar {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeQosCarsResponseBodyQosCarsQosCar struct {
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	LimitType           *string `json:"LimitType,omitempty" xml:"LimitType,omitempty"`
	MaxBandwidthAbs     *int32  `json:"MaxBandwidthAbs,omitempty" xml:"MaxBandwidthAbs,omitempty"`
	MaxBandwidthPercent *int32  `json:"MaxBandwidthPercent,omitempty" xml:"MaxBandwidthPercent,omitempty"`
	MinBandwidthAbs     *int32  `json:"MinBandwidthAbs,omitempty" xml:"MinBandwidthAbs,omitempty"`
	MinBandwidthPercent *int32  `json:"MinBandwidthPercent,omitempty" xml:"MinBandwidthPercent,omitempty"`
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PercentSourceType   *string `json:"PercentSourceType,omitempty" xml:"PercentSourceType,omitempty"`
	Priority            *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	QosCarId            *string `json:"QosCarId,omitempty" xml:"QosCarId,omitempty"`
	QosId               *string `json:"QosId,omitempty" xml:"QosId,omitempty"`
}

func (s DescribeQosCarsResponseBodyQosCarsQosCar) String() string {
	return dara.Prettify(s)
}

func (s DescribeQosCarsResponseBodyQosCarsQosCar) GoString() string {
	return s.String()
}

func (s *DescribeQosCarsResponseBodyQosCarsQosCar) GetDescription() *string {
	return s.Description
}

func (s *DescribeQosCarsResponseBodyQosCarsQosCar) GetLimitType() *string {
	return s.LimitType
}

func (s *DescribeQosCarsResponseBodyQosCarsQosCar) GetMaxBandwidthAbs() *int32 {
	return s.MaxBandwidthAbs
}

func (s *DescribeQosCarsResponseBodyQosCarsQosCar) GetMaxBandwidthPercent() *int32 {
	return s.MaxBandwidthPercent
}

func (s *DescribeQosCarsResponseBodyQosCarsQosCar) GetMinBandwidthAbs() *int32 {
	return s.MinBandwidthAbs
}

func (s *DescribeQosCarsResponseBodyQosCarsQosCar) GetMinBandwidthPercent() *int32 {
	return s.MinBandwidthPercent
}

func (s *DescribeQosCarsResponseBodyQosCarsQosCar) GetName() *string {
	return s.Name
}

func (s *DescribeQosCarsResponseBodyQosCarsQosCar) GetPercentSourceType() *string {
	return s.PercentSourceType
}

func (s *DescribeQosCarsResponseBodyQosCarsQosCar) GetPriority() *int32 {
	return s.Priority
}

func (s *DescribeQosCarsResponseBodyQosCarsQosCar) GetQosCarId() *string {
	return s.QosCarId
}

func (s *DescribeQosCarsResponseBodyQosCarsQosCar) GetQosId() *string {
	return s.QosId
}

func (s *DescribeQosCarsResponseBodyQosCarsQosCar) SetDescription(v string) *DescribeQosCarsResponseBodyQosCarsQosCar {
	s.Description = &v
	return s
}

func (s *DescribeQosCarsResponseBodyQosCarsQosCar) SetLimitType(v string) *DescribeQosCarsResponseBodyQosCarsQosCar {
	s.LimitType = &v
	return s
}

func (s *DescribeQosCarsResponseBodyQosCarsQosCar) SetMaxBandwidthAbs(v int32) *DescribeQosCarsResponseBodyQosCarsQosCar {
	s.MaxBandwidthAbs = &v
	return s
}

func (s *DescribeQosCarsResponseBodyQosCarsQosCar) SetMaxBandwidthPercent(v int32) *DescribeQosCarsResponseBodyQosCarsQosCar {
	s.MaxBandwidthPercent = &v
	return s
}

func (s *DescribeQosCarsResponseBodyQosCarsQosCar) SetMinBandwidthAbs(v int32) *DescribeQosCarsResponseBodyQosCarsQosCar {
	s.MinBandwidthAbs = &v
	return s
}

func (s *DescribeQosCarsResponseBodyQosCarsQosCar) SetMinBandwidthPercent(v int32) *DescribeQosCarsResponseBodyQosCarsQosCar {
	s.MinBandwidthPercent = &v
	return s
}

func (s *DescribeQosCarsResponseBodyQosCarsQosCar) SetName(v string) *DescribeQosCarsResponseBodyQosCarsQosCar {
	s.Name = &v
	return s
}

func (s *DescribeQosCarsResponseBodyQosCarsQosCar) SetPercentSourceType(v string) *DescribeQosCarsResponseBodyQosCarsQosCar {
	s.PercentSourceType = &v
	return s
}

func (s *DescribeQosCarsResponseBodyQosCarsQosCar) SetPriority(v int32) *DescribeQosCarsResponseBodyQosCarsQosCar {
	s.Priority = &v
	return s
}

func (s *DescribeQosCarsResponseBodyQosCarsQosCar) SetQosCarId(v string) *DescribeQosCarsResponseBodyQosCarsQosCar {
	s.QosCarId = &v
	return s
}

func (s *DescribeQosCarsResponseBodyQosCarsQosCar) SetQosId(v string) *DescribeQosCarsResponseBodyQosCarsQosCar {
	s.QosId = &v
	return s
}

func (s *DescribeQosCarsResponseBodyQosCarsQosCar) Validate() error {
	return dara.Validate(s)
}
