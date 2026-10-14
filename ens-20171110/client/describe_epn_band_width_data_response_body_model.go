// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEpnBandWidthDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMonitorData(v *DescribeEpnBandWidthDataResponseBodyMonitorData) *DescribeEpnBandWidthDataResponseBody
	GetMonitorData() *DescribeEpnBandWidthDataResponseBodyMonitorData
	SetRequestId(v string) *DescribeEpnBandWidthDataResponseBody
	GetRequestId() *string
}

type DescribeEpnBandWidthDataResponseBody struct {
	// The monitoring data of the instance.
	MonitorData *DescribeEpnBandWidthDataResponseBodyMonitorData `json:"MonitorData,omitempty" xml:"MonitorData,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1D289DAA-F6DA-5FC4-AE47-F5C8B6277BFC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEpnBandWidthDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEpnBandWidthDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEpnBandWidthDataResponseBody) GetMonitorData() *DescribeEpnBandWidthDataResponseBodyMonitorData {
	return s.MonitorData
}

func (s *DescribeEpnBandWidthDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEpnBandWidthDataResponseBody) SetMonitorData(v *DescribeEpnBandWidthDataResponseBodyMonitorData) *DescribeEpnBandWidthDataResponseBody {
	s.MonitorData = v
	return s
}

func (s *DescribeEpnBandWidthDataResponseBody) SetRequestId(v string) *DescribeEpnBandWidthDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEpnBandWidthDataResponseBody) Validate() error {
	if s.MonitorData != nil {
		if err := s.MonitorData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEpnBandWidthDataResponseBodyMonitorData struct {
	// The bandwidth monitoring data.
	BandWidthMonitorData []*DescribeEpnBandWidthDataResponseBodyMonitorDataBandWidthMonitorData `json:"BandWidthMonitorData,omitempty" xml:"BandWidthMonitorData,omitempty" type:"Repeated"`
	// The maximum downstream bandwidth within the queried time range. Unit: bit/s.
	//
	// example:
	//
	// 16817468
	MaxDownBandWidth *int64 `json:"MaxDownBandWidth,omitempty" xml:"MaxDownBandWidth,omitempty"`
	// The maximum upstream bandwidth within the queried time range. Unit: bit/s.
	//
	// example:
	//
	// 231008
	MaxUpBandWidth *int64 `json:"MaxUpBandWidth,omitempty" xml:"MaxUpBandWidth,omitempty"`
}

func (s DescribeEpnBandWidthDataResponseBodyMonitorData) String() string {
	return dara.Prettify(s)
}

func (s DescribeEpnBandWidthDataResponseBodyMonitorData) GoString() string {
	return s.String()
}

func (s *DescribeEpnBandWidthDataResponseBodyMonitorData) GetBandWidthMonitorData() []*DescribeEpnBandWidthDataResponseBodyMonitorDataBandWidthMonitorData {
	return s.BandWidthMonitorData
}

func (s *DescribeEpnBandWidthDataResponseBodyMonitorData) GetMaxDownBandWidth() *int64 {
	return s.MaxDownBandWidth
}

func (s *DescribeEpnBandWidthDataResponseBodyMonitorData) GetMaxUpBandWidth() *int64 {
	return s.MaxUpBandWidth
}

func (s *DescribeEpnBandWidthDataResponseBodyMonitorData) SetBandWidthMonitorData(v []*DescribeEpnBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) *DescribeEpnBandWidthDataResponseBodyMonitorData {
	s.BandWidthMonitorData = v
	return s
}

func (s *DescribeEpnBandWidthDataResponseBodyMonitorData) SetMaxDownBandWidth(v int64) *DescribeEpnBandWidthDataResponseBodyMonitorData {
	s.MaxDownBandWidth = &v
	return s
}

func (s *DescribeEpnBandWidthDataResponseBodyMonitorData) SetMaxUpBandWidth(v int64) *DescribeEpnBandWidthDataResponseBodyMonitorData {
	s.MaxUpBandWidth = &v
	return s
}

func (s *DescribeEpnBandWidthDataResponseBodyMonitorData) Validate() error {
	if s.BandWidthMonitorData != nil {
		for _, item := range s.BandWidthMonitorData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEpnBandWidthDataResponseBodyMonitorDataBandWidthMonitorData struct {
	// The downstream bandwidth. Unit: bit/s.
	//
	// example:
	//
	// 0
	DownBandWidth *int64 `json:"DownBandWidth,omitempty" xml:"DownBandWidth,omitempty"`
	// The inbound public network traffic of the instance. Unit: bits.
	//
	// example:
	//
	// 0
	InternetRX *int64 `json:"InternetRX,omitempty" xml:"InternetRX,omitempty"`
	// The outbound public network traffic of the instance. Unit: bits.
	//
	// example:
	//
	// 0
	InternetTX *int64 `json:"InternetTX,omitempty" xml:"InternetTX,omitempty"`
	// The timestamp of the monitoring data. The time is in UTC and in the format of yyyy-MM-DDThh:mm:ssZ.
	//
	// example:
	//
	// 2019-10-12T05:45:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The upstream bandwidth. Unit: bit/s.
	//
	// example:
	//
	// 0
	UpBandWidth *int64 `json:"UpBandWidth,omitempty" xml:"UpBandWidth,omitempty"`
}

func (s DescribeEpnBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) String() string {
	return dara.Prettify(s)
}

func (s DescribeEpnBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) GoString() string {
	return s.String()
}

func (s *DescribeEpnBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) GetDownBandWidth() *int64 {
	return s.DownBandWidth
}

func (s *DescribeEpnBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) GetInternetRX() *int64 {
	return s.InternetRX
}

func (s *DescribeEpnBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) GetInternetTX() *int64 {
	return s.InternetTX
}

func (s *DescribeEpnBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeEpnBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) GetUpBandWidth() *int64 {
	return s.UpBandWidth
}

func (s *DescribeEpnBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) SetDownBandWidth(v int64) *DescribeEpnBandWidthDataResponseBodyMonitorDataBandWidthMonitorData {
	s.DownBandWidth = &v
	return s
}

func (s *DescribeEpnBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) SetInternetRX(v int64) *DescribeEpnBandWidthDataResponseBodyMonitorDataBandWidthMonitorData {
	s.InternetRX = &v
	return s
}

func (s *DescribeEpnBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) SetInternetTX(v int64) *DescribeEpnBandWidthDataResponseBodyMonitorDataBandWidthMonitorData {
	s.InternetTX = &v
	return s
}

func (s *DescribeEpnBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) SetTimeStamp(v string) *DescribeEpnBandWidthDataResponseBodyMonitorDataBandWidthMonitorData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeEpnBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) SetUpBandWidth(v int64) *DescribeEpnBandWidthDataResponseBodyMonitorDataBandWidthMonitorData {
	s.UpBandWidth = &v
	return s
}

func (s *DescribeEpnBandWidthDataResponseBodyMonitorDataBandWidthMonitorData) Validate() error {
	return dara.Validate(s)
}
