// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEpnBandwitdhByInternetChargeTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidthValue(v int64) *DescribeEpnBandwitdhByInternetChargeTypeResponseBody
	GetBandwidthValue() *int64
	SetInternetChargeType(v string) *DescribeEpnBandwitdhByInternetChargeTypeResponseBody
	GetInternetChargeType() *string
	SetRequestId(v string) *DescribeEpnBandwitdhByInternetChargeTypeResponseBody
	GetRequestId() *string
	SetTimeStamp(v string) *DescribeEpnBandwitdhByInternetChargeTypeResponseBody
	GetTimeStamp() *string
}

type DescribeEpnBandwitdhByInternetChargeTypeResponseBody struct {
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
	// BandwidthByDay
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 216BCED0-E055-5DDB-8E06-4084A62A4A44
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The timestamp when the monitoring data was queried. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-10-12T05:45:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeEpnBandwitdhByInternetChargeTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEpnBandwitdhByInternetChargeTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeResponseBody) GetBandwidthValue() *int64 {
	return s.BandwidthValue
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeResponseBody) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeResponseBody) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeResponseBody) SetBandwidthValue(v int64) *DescribeEpnBandwitdhByInternetChargeTypeResponseBody {
	s.BandwidthValue = &v
	return s
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeResponseBody) SetInternetChargeType(v string) *DescribeEpnBandwitdhByInternetChargeTypeResponseBody {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeResponseBody) SetRequestId(v string) *DescribeEpnBandwitdhByInternetChargeTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeResponseBody) SetTimeStamp(v string) *DescribeEpnBandwitdhByInternetChargeTypeResponseBody {
	s.TimeStamp = &v
	return s
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
