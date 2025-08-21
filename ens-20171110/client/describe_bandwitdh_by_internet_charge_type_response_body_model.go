// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBandwitdhByInternetChargeTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidthValue(v int64) *DescribeBandwitdhByInternetChargeTypeResponseBody
	GetBandwidthValue() *int64
	SetInternetChargeType(v string) *DescribeBandwitdhByInternetChargeTypeResponseBody
	GetInternetChargeType() *string
	SetRequestId(v string) *DescribeBandwitdhByInternetChargeTypeResponseBody
	GetRequestId() *string
	SetTimeStamp(v string) *DescribeBandwitdhByInternetChargeTypeResponseBody
	GetTimeStamp() *string
}

type DescribeBandwitdhByInternetChargeTypeResponseBody struct {
	// The bandwidth. Unit: bit/s.
	//
	// example:
	//
	// 123
	BandwidthValue *int64 `json:"BandwidthValue,omitempty" xml:"BandwidthValue,omitempty"`
	// The metering method. Valid values:
	//
	// 	- BandwidthByDay: Pay by daily peak bandwidth
	//
	// 	- 95BandwidthByMonth: Pay by monthly 95th percentile bandwidth
	//
	// 	- PayByBandwidth4thMonth: Pay by monthly fourth peak bandwidth
	//
	// 	- PayByBandwidth: Pay by fixed bandwidth
	//
	// You can specify only one metering method for network usage and cannot overwrite the existing metering method.
	//
	// example:
	//
	// 95BandwidthByMonth
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// 08027633-8501-5A36-B90D-F7C38B5EC75D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The timestamp. The time follows the ISO 8601 standard. The time is displayed in UTC. Example: 2016-10-20T04:00:00Z.
	//
	// example:
	//
	// 2019-10-12T05:45:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeBandwitdhByInternetChargeTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBandwitdhByInternetChargeTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBandwitdhByInternetChargeTypeResponseBody) GetBandwidthValue() *int64 {
	return s.BandwidthValue
}

func (s *DescribeBandwitdhByInternetChargeTypeResponseBody) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribeBandwitdhByInternetChargeTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBandwitdhByInternetChargeTypeResponseBody) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeBandwitdhByInternetChargeTypeResponseBody) SetBandwidthValue(v int64) *DescribeBandwitdhByInternetChargeTypeResponseBody {
	s.BandwidthValue = &v
	return s
}

func (s *DescribeBandwitdhByInternetChargeTypeResponseBody) SetInternetChargeType(v string) *DescribeBandwitdhByInternetChargeTypeResponseBody {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeBandwitdhByInternetChargeTypeResponseBody) SetRequestId(v string) *DescribeBandwitdhByInternetChargeTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBandwitdhByInternetChargeTypeResponseBody) SetTimeStamp(v string) *DescribeBandwitdhByInternetChargeTypeResponseBody {
	s.TimeStamp = &v
	return s
}

func (s *DescribeBandwitdhByInternetChargeTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
