// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserBandWidthDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeUserBandWidthDataResponseBody
	GetCode() *int32
	SetMonitorData(v *DescribeUserBandWidthDataResponseBodyMonitorData) *DescribeUserBandWidthDataResponseBody
	GetMonitorData() *DescribeUserBandWidthDataResponseBodyMonitorData
	SetRequestId(v string) *DescribeUserBandWidthDataResponseBody
	GetRequestId() *string
}

type DescribeUserBandWidthDataResponseBody struct {
	// The returned service code. 0 indicates that the request was successful.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The monitoring data.
	MonitorData *DescribeUserBandWidthDataResponseBodyMonitorData `json:"MonitorData,omitempty" xml:"MonitorData,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 50373E71-7710-4620-8AAB-133CCE49451C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserBandWidthDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserBandWidthDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserBandWidthDataResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeUserBandWidthDataResponseBody) GetMonitorData() *DescribeUserBandWidthDataResponseBodyMonitorData {
	return s.MonitorData
}

func (s *DescribeUserBandWidthDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserBandWidthDataResponseBody) SetCode(v int32) *DescribeUserBandWidthDataResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeUserBandWidthDataResponseBody) SetMonitorData(v *DescribeUserBandWidthDataResponseBodyMonitorData) *DescribeUserBandWidthDataResponseBody {
	s.MonitorData = v
	return s
}

func (s *DescribeUserBandWidthDataResponseBody) SetRequestId(v string) *DescribeUserBandWidthDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserBandWidthDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeUserBandWidthDataResponseBodyMonitorData struct {
	// The bandwidth data.
	BandWidthMonitorData []*DescribeUserBandWidthDataResponseBodyMonitorDataBandWidthMonitorData `json:"BandWidthMonitorData,omitempty" xml:"BandWidthMonitorData,omitempty" type:"Repeated"`
	// The maximum outbound bandwidth within the queried time range. Unit: bit/s.
	//
	// example:
	//
	// 16817468
	MaxDownBandWidth *string `json:"MaxDownBandWidth,omitempty" xml:"MaxDownBandWidth,omitempty"`
	// The maximum inbound bandwidth within the queried time range. Unit: bit/s.
	//
	// example:
	//
	// 231008
	MaxUpBandWidth *string `json:"MaxUpBandWidth,omitempty" xml:"MaxUpBandWidth,omitempty"`
}

func (s DescribeUserBandWidthDataResponseBodyMonitorData) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserBandWidthDataResponseBodyMonitorData) GoString() string {
	return s.String()
}

func (s *DescribeUserBandWidthDataResponseBodyMonitorData) GetBandWidthMonitorData() []*DescribeUserBandWidthDataResponseBodyMonitorDataBandWidthMonitorData {
	return s.BandWidthMonitorData
}

func (s *DescribeUserBandWidthDataResponseBodyMonitorData) GetMaxDownBandWidth() *string {
	return s.MaxDownBandWidth
}

func (s *DescribeUserBandWidthDataResponseBodyMonitorData) GetMaxUpBandWidth() *string {
	return s.MaxUpBandWidth
}

func (s *DescribeUserBandWidthDataResponseBodyMonitorData) SetBandWidthMonitorData(v []*DescribeUserBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) *DescribeUserBandWidthDataResponseBodyMonitorData {
	s.BandWidthMonitorData = v
	return s
}

func (s *DescribeUserBandWidthDataResponseBodyMonitorData) SetMaxDownBandWidth(v string) *DescribeUserBandWidthDataResponseBodyMonitorData {
	s.MaxDownBandWidth = &v
	return s
}

func (s *DescribeUserBandWidthDataResponseBodyMonitorData) SetMaxUpBandWidth(v string) *DescribeUserBandWidthDataResponseBodyMonitorData {
	s.MaxUpBandWidth = &v
	return s
}

func (s *DescribeUserBandWidthDataResponseBodyMonitorData) Validate() error {
	return dara.Validate(s)
}

type DescribeUserBandWidthDataResponseBodyMonitorDataBandWidthMonitorData struct {
	// The outbound bandwidth. Unit: bit/s.
	//
	// example:
	//
	// 0
	DownBandWidth *int64 `json:"DownBandWidth,omitempty" xml:"DownBandWidth,omitempty"`
	// The Internet traffic to the instance. Unit: bytes.
	//
	// example:
	//
	// 0
	InternetRX *int64 `json:"InternetRX,omitempty" xml:"InternetRX,omitempty"`
	// The Internet traffic from the instance. Unit: bytes.
	//
	// example:
	//
	// 0
	InternetTX *int64 `json:"InternetTX,omitempty" xml:"InternetTX,omitempty"`
	// The timestamp when the monitoring data was queried. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-10-12T05:45:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The inbound bandwidth. Unit: bit/s.
	//
	// example:
	//
	// 0
	UpBandWidth *int64 `json:"UpBandWidth,omitempty" xml:"UpBandWidth,omitempty"`
}

func (s DescribeUserBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) GoString() string {
	return s.String()
}

func (s *DescribeUserBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) GetDownBandWidth() *int64 {
	return s.DownBandWidth
}

func (s *DescribeUserBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) GetInternetRX() *int64 {
	return s.InternetRX
}

func (s *DescribeUserBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) GetInternetTX() *int64 {
	return s.InternetTX
}

func (s *DescribeUserBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeUserBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) GetUpBandWidth() *int64 {
	return s.UpBandWidth
}

func (s *DescribeUserBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) SetDownBandWidth(v int64) *DescribeUserBandWidthDataResponseBodyMonitorDataBandWidthMonitorData {
	s.DownBandWidth = &v
	return s
}

func (s *DescribeUserBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) SetInternetRX(v int64) *DescribeUserBandWidthDataResponseBodyMonitorDataBandWidthMonitorData {
	s.InternetRX = &v
	return s
}

func (s *DescribeUserBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) SetInternetTX(v int64) *DescribeUserBandWidthDataResponseBodyMonitorDataBandWidthMonitorData {
	s.InternetTX = &v
	return s
}

func (s *DescribeUserBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) SetTimeStamp(v string) *DescribeUserBandWidthDataResponseBodyMonitorDataBandWidthMonitorData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeUserBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) SetUpBandWidth(v int64) *DescribeUserBandWidthDataResponseBodyMonitorDataBandWidthMonitorData {
	s.UpBandWidth = &v
	return s
}

func (s *DescribeUserBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) Validate() error {
	return dara.Validate(s)
}
