// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckTimeDimensionStatisticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTimeStamp(v int64) *GetCheckTimeDimensionStatisticRequest
	GetEndTimeStamp() *int64
	SetStartTimeStamp(v int64) *GetCheckTimeDimensionStatisticRequest
	GetStartTimeStamp() *int64
	SetStatisticType(v string) *GetCheckTimeDimensionStatisticRequest
	GetStatisticType() *string
	SetVendors(v []*string) *GetCheckTimeDimensionStatisticRequest
	GetVendors() []*string
}

type GetCheckTimeDimensionStatisticRequest struct {
	// End time, in timestamp format.
	//
	// example:
	//
	// 1672285044000
	EndTimeStamp *int64 `json:"EndTimeStamp,omitempty" xml:"EndTimeStamp,omitempty"`
	// Start time, in timestamp format.
	//
	// example:
	//
	// 1672385044000
	StartTimeStamp *int64 `json:"StartTimeStamp,omitempty" xml:"StartTimeStamp,omitempty"`
	// Type of statistical data. Values:
	//
	// - **CheckPassRate**: Check item pass rate.
	//
	// - **AssetPassRate**: Asset pass rate.
	//
	// example:
	//
	// AssetPassRate
	StatisticType *string `json:"StatisticType,omitempty" xml:"StatisticType,omitempty"`
	// List of cloud vendors.
	Vendors []*string `json:"Vendors,omitempty" xml:"Vendors,omitempty" type:"Repeated"`
}

func (s GetCheckTimeDimensionStatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCheckTimeDimensionStatisticRequest) GoString() string {
	return s.String()
}

func (s *GetCheckTimeDimensionStatisticRequest) GetEndTimeStamp() *int64 {
	return s.EndTimeStamp
}

func (s *GetCheckTimeDimensionStatisticRequest) GetStartTimeStamp() *int64 {
	return s.StartTimeStamp
}

func (s *GetCheckTimeDimensionStatisticRequest) GetStatisticType() *string {
	return s.StatisticType
}

func (s *GetCheckTimeDimensionStatisticRequest) GetVendors() []*string {
	return s.Vendors
}

func (s *GetCheckTimeDimensionStatisticRequest) SetEndTimeStamp(v int64) *GetCheckTimeDimensionStatisticRequest {
	s.EndTimeStamp = &v
	return s
}

func (s *GetCheckTimeDimensionStatisticRequest) SetStartTimeStamp(v int64) *GetCheckTimeDimensionStatisticRequest {
	s.StartTimeStamp = &v
	return s
}

func (s *GetCheckTimeDimensionStatisticRequest) SetStatisticType(v string) *GetCheckTimeDimensionStatisticRequest {
	s.StatisticType = &v
	return s
}

func (s *GetCheckTimeDimensionStatisticRequest) SetVendors(v []*string) *GetCheckTimeDimensionStatisticRequest {
	s.Vendors = v
	return s
}

func (s *GetCheckTimeDimensionStatisticRequest) Validate() error {
	return dara.Validate(s)
}
