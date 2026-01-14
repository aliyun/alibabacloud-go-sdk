// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *DescribeApplicationMonitorResponseBody
	GetAcceleratorId() *string
	SetAddress(v string) *DescribeApplicationMonitorResponseBody
	GetAddress() *string
	SetDetectEnable(v bool) *DescribeApplicationMonitorResponseBody
	GetDetectEnable() *bool
	SetDetectThreshold(v int32) *DescribeApplicationMonitorResponseBody
	GetDetectThreshold() *int32
	SetDetectTimes(v int32) *DescribeApplicationMonitorResponseBody
	GetDetectTimes() *int32
	SetIspCityList(v []*DescribeApplicationMonitorResponseBodyIspCityList) *DescribeApplicationMonitorResponseBody
	GetIspCityList() []*DescribeApplicationMonitorResponseBodyIspCityList
	SetListenerId(v string) *DescribeApplicationMonitorResponseBody
	GetListenerId() *string
	SetOptionsJson(v string) *DescribeApplicationMonitorResponseBody
	GetOptionsJson() *string
	SetRegionId(v string) *DescribeApplicationMonitorResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeApplicationMonitorResponseBody
	GetRequestId() *string
	SetSilenceTime(v int32) *DescribeApplicationMonitorResponseBody
	GetSilenceTime() *int32
	SetState(v string) *DescribeApplicationMonitorResponseBody
	GetState() *string
	SetTaskId(v string) *DescribeApplicationMonitorResponseBody
	GetTaskId() *string
	SetTaskName(v string) *DescribeApplicationMonitorResponseBody
	GetTaskName() *string
}

type DescribeApplicationMonitorResponseBody struct {
	// The ID of the GA instance on which the origin probing task ran.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The URL or IP address that was probed.
	//
	// example:
	//
	// https://www.aliyun.com
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// Indicates whether the automatic diagnostics feature is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	DetectEnable *bool `json:"DetectEnable,omitempty" xml:"DetectEnable,omitempty"`
	// The threshold that is used to trigger automatic diagnostics.
	//
	// If the availability of the origin server drops below the specified threshold, the automatic diagnostics feature is triggered.
	//
	// example:
	//
	// 0
	DetectThreshold *int32 `json:"DetectThreshold,omitempty" xml:"DetectThreshold,omitempty"`
	// The number of times that are required to reach the threshold before the automatic diagnostics feature is triggered.
	//
	// example:
	//
	// 1
	DetectTimes *int32 `json:"DetectTimes,omitempty" xml:"DetectTimes,omitempty"`
	// The probe points of the Internet service provider (ISP).
	IspCityList []*DescribeApplicationMonitorResponseBodyIspCityList `json:"IspCityList,omitempty" xml:"IspCityList,omitempty" type:"Repeated"`
	// The ID of the listener on which the origin probing task ran.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The extended options of the listener protocol that is used by the origin probing task. The options vary based on the listener protocol.
	//
	// example:
	//
	// { "http_method": "get","header": "key:asd","acceptable_response_code": "500","cert_verify": true }
	OptionsJson *string `json:"OptionsJson,omitempty" xml:"OptionsJson,omitempty"`
	// The region ID of the GA instance. Set the value to **cn-hangzhou**.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The silence period of the automatic diagnostics feature. This parameter indicates the interval at which the automatic diagnostics feature is triggered. If the availability rate does not return to normal after GA triggers automatic diagnostics, GA must wait until the silence period ends before GA can trigger another automatic diagnostic.
	//
	// If the number of consecutive times that the availability rate drops below the automatic diagnostics threshold reaches the value of **DetectTimes**, the automatic diagnostics feature is triggered. The automatic diagnostics feature is not triggered again within the silence period regardless of whether the availability rate remains below the threshold. If the availability rate does not return to normal after the silence period ends, the automatic diagnostics feature is triggered again.
	//
	// Unit: seconds.
	//
	// example:
	//
	// 300
	SilenceTime *int32 `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	// The status of the origin probing task. Valid values:
	//
	// 	- **init**
	//
	// 	- **active**
	//
	// 	- **updating**
	//
	// 	- **inactive**
	//
	// 	- **deleting**
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The ID of the origin probing task.
	//
	// example:
	//
	// sm-bp1fpdjfju9k8yr1y****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the origin probing task.
	//
	// example:
	//
	// task1
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s DescribeApplicationMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationMonitorResponseBody) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *DescribeApplicationMonitorResponseBody) GetAddress() *string {
	return s.Address
}

func (s *DescribeApplicationMonitorResponseBody) GetDetectEnable() *bool {
	return s.DetectEnable
}

func (s *DescribeApplicationMonitorResponseBody) GetDetectThreshold() *int32 {
	return s.DetectThreshold
}

func (s *DescribeApplicationMonitorResponseBody) GetDetectTimes() *int32 {
	return s.DetectTimes
}

func (s *DescribeApplicationMonitorResponseBody) GetIspCityList() []*DescribeApplicationMonitorResponseBodyIspCityList {
	return s.IspCityList
}

func (s *DescribeApplicationMonitorResponseBody) GetListenerId() *string {
	return s.ListenerId
}

func (s *DescribeApplicationMonitorResponseBody) GetOptionsJson() *string {
	return s.OptionsJson
}

func (s *DescribeApplicationMonitorResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApplicationMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApplicationMonitorResponseBody) GetSilenceTime() *int32 {
	return s.SilenceTime
}

func (s *DescribeApplicationMonitorResponseBody) GetState() *string {
	return s.State
}

func (s *DescribeApplicationMonitorResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeApplicationMonitorResponseBody) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeApplicationMonitorResponseBody) SetAcceleratorId(v string) *DescribeApplicationMonitorResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *DescribeApplicationMonitorResponseBody) SetAddress(v string) *DescribeApplicationMonitorResponseBody {
	s.Address = &v
	return s
}

func (s *DescribeApplicationMonitorResponseBody) SetDetectEnable(v bool) *DescribeApplicationMonitorResponseBody {
	s.DetectEnable = &v
	return s
}

func (s *DescribeApplicationMonitorResponseBody) SetDetectThreshold(v int32) *DescribeApplicationMonitorResponseBody {
	s.DetectThreshold = &v
	return s
}

func (s *DescribeApplicationMonitorResponseBody) SetDetectTimes(v int32) *DescribeApplicationMonitorResponseBody {
	s.DetectTimes = &v
	return s
}

func (s *DescribeApplicationMonitorResponseBody) SetIspCityList(v []*DescribeApplicationMonitorResponseBodyIspCityList) *DescribeApplicationMonitorResponseBody {
	s.IspCityList = v
	return s
}

func (s *DescribeApplicationMonitorResponseBody) SetListenerId(v string) *DescribeApplicationMonitorResponseBody {
	s.ListenerId = &v
	return s
}

func (s *DescribeApplicationMonitorResponseBody) SetOptionsJson(v string) *DescribeApplicationMonitorResponseBody {
	s.OptionsJson = &v
	return s
}

func (s *DescribeApplicationMonitorResponseBody) SetRegionId(v string) *DescribeApplicationMonitorResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeApplicationMonitorResponseBody) SetRequestId(v string) *DescribeApplicationMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationMonitorResponseBody) SetSilenceTime(v int32) *DescribeApplicationMonitorResponseBody {
	s.SilenceTime = &v
	return s
}

func (s *DescribeApplicationMonitorResponseBody) SetState(v string) *DescribeApplicationMonitorResponseBody {
	s.State = &v
	return s
}

func (s *DescribeApplicationMonitorResponseBody) SetTaskId(v string) *DescribeApplicationMonitorResponseBody {
	s.TaskId = &v
	return s
}

func (s *DescribeApplicationMonitorResponseBody) SetTaskName(v string) *DescribeApplicationMonitorResponseBody {
	s.TaskName = &v
	return s
}

func (s *DescribeApplicationMonitorResponseBody) Validate() error {
	if s.IspCityList != nil {
		for _, item := range s.IspCityList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApplicationMonitorResponseBodyIspCityList struct {
	// The ID of the city in which the probe point of the ISP is deployed.
	//
	// example:
	//
	// 375
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// The name of the city in which the probe point of the ISP is deployed.
	//
	// example:
	//
	// Singapore
	CityName *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	// The probe point ID of the ISP.
	//
	// example:
	//
	// 465
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// The probe point name of the ISP.
	//
	// example:
	//
	// Alibaba
	IspName *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
}

func (s DescribeApplicationMonitorResponseBodyIspCityList) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationMonitorResponseBodyIspCityList) GoString() string {
	return s.String()
}

func (s *DescribeApplicationMonitorResponseBodyIspCityList) GetCity() *string {
	return s.City
}

func (s *DescribeApplicationMonitorResponseBodyIspCityList) GetCityName() *string {
	return s.CityName
}

func (s *DescribeApplicationMonitorResponseBodyIspCityList) GetIsp() *string {
	return s.Isp
}

func (s *DescribeApplicationMonitorResponseBodyIspCityList) GetIspName() *string {
	return s.IspName
}

func (s *DescribeApplicationMonitorResponseBodyIspCityList) SetCity(v string) *DescribeApplicationMonitorResponseBodyIspCityList {
	s.City = &v
	return s
}

func (s *DescribeApplicationMonitorResponseBodyIspCityList) SetCityName(v string) *DescribeApplicationMonitorResponseBodyIspCityList {
	s.CityName = &v
	return s
}

func (s *DescribeApplicationMonitorResponseBodyIspCityList) SetIsp(v string) *DescribeApplicationMonitorResponseBodyIspCityList {
	s.Isp = &v
	return s
}

func (s *DescribeApplicationMonitorResponseBodyIspCityList) SetIspName(v string) *DescribeApplicationMonitorResponseBodyIspCityList {
	s.IspName = &v
	return s
}

func (s *DescribeApplicationMonitorResponseBodyIspCityList) Validate() error {
	return dara.Validate(s)
}
