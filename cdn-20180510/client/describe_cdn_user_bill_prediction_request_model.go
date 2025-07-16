// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnUserBillPredictionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArea(v string) *DescribeCdnUserBillPredictionRequest
	GetArea() *string
	SetDimension(v string) *DescribeCdnUserBillPredictionRequest
	GetDimension() *string
	SetEndTime(v string) *DescribeCdnUserBillPredictionRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeCdnUserBillPredictionRequest
	GetStartTime() *string
}

type DescribeCdnUserBillPredictionRequest struct {
	// The billable region. Valid values:
	//
	// 	- **CN**: the Chinese mainland
	//
	// 	- **OverSeas**: outside the Chinese mainland
	//
	// 	- **AP1**: Asia Pacific 1
	//
	// 	- **AP2**: Asia Pacific 2
	//
	// 	- **AP3**: Asia Pacific 3
	//
	// 	- **NA**: North America
	//
	// 	- **SA**: South America
	//
	// 	- **EU**: Europe
	//
	// 	- **MEAA**: Middle East and Africa
	//
	// By default, the value of this parameter is determined by the metering method that is currently used. Regions inside and outside the Chinese mainland are classified into the **CN*	- and **OverSeas*	- billable regions. Billable regions inside the Chinese mainland include **CN**. Billable regions outside the Chinese mainland include **AP1**, **AP2**, **AP3**, **NA**, **SA**, **EU**, and **MEAA**.
	//
	// > For more information about billable regions, see [Billable regions](https://help.aliyun.com/document_detail/142221.html).
	//
	// example:
	//
	// CN,OverSeas
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// The billable item. A value of flow specifies bandwidth.
	//
	// example:
	//
	// flow
	Dimension *string `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// The end time of the estimation. The default value is the current time. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// > The end time must be later than the start time.
	//
	// example:
	//
	// 2018-10-25T10:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time of the estimation. The default value is 00:00 on the first day of the current month. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2018-09-30T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeCdnUserBillPredictionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnUserBillPredictionRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserBillPredictionRequest) GetArea() *string {
	return s.Area
}

func (s *DescribeCdnUserBillPredictionRequest) GetDimension() *string {
	return s.Dimension
}

func (s *DescribeCdnUserBillPredictionRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeCdnUserBillPredictionRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeCdnUserBillPredictionRequest) SetArea(v string) *DescribeCdnUserBillPredictionRequest {
	s.Area = &v
	return s
}

func (s *DescribeCdnUserBillPredictionRequest) SetDimension(v string) *DescribeCdnUserBillPredictionRequest {
	s.Dimension = &v
	return s
}

func (s *DescribeCdnUserBillPredictionRequest) SetEndTime(v string) *DescribeCdnUserBillPredictionRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeCdnUserBillPredictionRequest) SetStartTime(v string) *DescribeCdnUserBillPredictionRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeCdnUserBillPredictionRequest) Validate() error {
	return dara.Validate(s)
}
