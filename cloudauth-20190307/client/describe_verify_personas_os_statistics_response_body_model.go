// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVerifyPersonasOsStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeVerifyPersonasOsStatisticsResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeVerifyPersonasOsStatisticsResponseBodyResultObject) *DescribeVerifyPersonasOsStatisticsResponseBody
	GetResultObject() *DescribeVerifyPersonasOsStatisticsResponseBodyResultObject
}

type DescribeVerifyPersonasOsStatisticsResponseBody struct {
	// ID of this request.
	//
	// example:
	//
	// 123906BE-752B-51E3-A8FF-52F53B659CE6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Processing result.
	ResultObject *DescribeVerifyPersonasOsStatisticsResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s DescribeVerifyPersonasOsStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyPersonasOsStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVerifyPersonasOsStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVerifyPersonasOsStatisticsResponseBody) GetResultObject() *DescribeVerifyPersonasOsStatisticsResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeVerifyPersonasOsStatisticsResponseBody) SetRequestId(v string) *DescribeVerifyPersonasOsStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVerifyPersonasOsStatisticsResponseBody) SetResultObject(v *DescribeVerifyPersonasOsStatisticsResponseBodyResultObject) *DescribeVerifyPersonasOsStatisticsResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeVerifyPersonasOsStatisticsResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVerifyPersonasOsStatisticsResponseBodyResultObject struct {
	// Total number of authenticated devices.
	//
	// example:
	//
	// 24
	AllDeviceCnt *int64 `json:"AllDeviceCnt,omitempty" xml:"AllDeviceCnt,omitempty"`
	// Number of authenticated Android devices.
	//
	// example:
	//
	// 14
	DeviceAndroidCnt *int64 `json:"DeviceAndroidCnt,omitempty" xml:"DeviceAndroidCnt,omitempty"`
	// Proportion of Android devices.
	//
	// example:
	//
	// 58.33
	DeviceAndroidRate *string `json:"DeviceAndroidRate,omitempty" xml:"DeviceAndroidRate,omitempty"`
	// Number of authenticated iOS devices.
	//
	// example:
	//
	// 10
	DeviceIosCnt *int64 `json:"DeviceIosCnt,omitempty" xml:"DeviceIosCnt,omitempty"`
	// Proportion of iOS devices.
	//
	// example:
	//
	// 41.67
	DeviceIosRate *string `json:"DeviceIosRate,omitempty" xml:"DeviceIosRate,omitempty"`
}

func (s DescribeVerifyPersonasOsStatisticsResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyPersonasOsStatisticsResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeVerifyPersonasOsStatisticsResponseBodyResultObject) GetAllDeviceCnt() *int64 {
	return s.AllDeviceCnt
}

func (s *DescribeVerifyPersonasOsStatisticsResponseBodyResultObject) GetDeviceAndroidCnt() *int64 {
	return s.DeviceAndroidCnt
}

func (s *DescribeVerifyPersonasOsStatisticsResponseBodyResultObject) GetDeviceAndroidRate() *string {
	return s.DeviceAndroidRate
}

func (s *DescribeVerifyPersonasOsStatisticsResponseBodyResultObject) GetDeviceIosCnt() *int64 {
	return s.DeviceIosCnt
}

func (s *DescribeVerifyPersonasOsStatisticsResponseBodyResultObject) GetDeviceIosRate() *string {
	return s.DeviceIosRate
}

func (s *DescribeVerifyPersonasOsStatisticsResponseBodyResultObject) SetAllDeviceCnt(v int64) *DescribeVerifyPersonasOsStatisticsResponseBodyResultObject {
	s.AllDeviceCnt = &v
	return s
}

func (s *DescribeVerifyPersonasOsStatisticsResponseBodyResultObject) SetDeviceAndroidCnt(v int64) *DescribeVerifyPersonasOsStatisticsResponseBodyResultObject {
	s.DeviceAndroidCnt = &v
	return s
}

func (s *DescribeVerifyPersonasOsStatisticsResponseBodyResultObject) SetDeviceAndroidRate(v string) *DescribeVerifyPersonasOsStatisticsResponseBodyResultObject {
	s.DeviceAndroidRate = &v
	return s
}

func (s *DescribeVerifyPersonasOsStatisticsResponseBodyResultObject) SetDeviceIosCnt(v int64) *DescribeVerifyPersonasOsStatisticsResponseBodyResultObject {
	s.DeviceIosCnt = &v
	return s
}

func (s *DescribeVerifyPersonasOsStatisticsResponseBodyResultObject) SetDeviceIosRate(v string) *DescribeVerifyPersonasOsStatisticsResponseBodyResultObject {
	s.DeviceIosRate = &v
	return s
}

func (s *DescribeVerifyPersonasOsStatisticsResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
