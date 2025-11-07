// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVerifyPersonasDeviceModelStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeVerifyPersonasDeviceModelStatisticsResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeVerifyPersonasDeviceModelStatisticsResponseBodyResultObject) *DescribeVerifyPersonasDeviceModelStatisticsResponseBody
	GetResultObject() *DescribeVerifyPersonasDeviceModelStatisticsResponseBodyResultObject
}

type DescribeVerifyPersonasDeviceModelStatisticsResponseBody struct {
	// ID of this request.
	//
	// example:
	//
	// 026828A3-CC7E-5D85-85B6-08DF245C5A53
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Query result.
	ResultObject *DescribeVerifyPersonasDeviceModelStatisticsResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s DescribeVerifyPersonasDeviceModelStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyPersonasDeviceModelStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVerifyPersonasDeviceModelStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVerifyPersonasDeviceModelStatisticsResponseBody) GetResultObject() *DescribeVerifyPersonasDeviceModelStatisticsResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeVerifyPersonasDeviceModelStatisticsResponseBody) SetRequestId(v string) *DescribeVerifyPersonasDeviceModelStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVerifyPersonasDeviceModelStatisticsResponseBody) SetResultObject(v *DescribeVerifyPersonasDeviceModelStatisticsResponseBodyResultObject) *DescribeVerifyPersonasDeviceModelStatisticsResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeVerifyPersonasDeviceModelStatisticsResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVerifyPersonasDeviceModelStatisticsResponseBodyResultObject struct {
	// Total number of devices.
	//
	// example:
	//
	// 31
	AllDeviceCnt *int64 `json:"AllDeviceCnt,omitempty" xml:"AllDeviceCnt,omitempty"`
	// List of data for different phone models.
	Items []*DescribeVerifyPersonasDeviceModelStatisticsResponseBodyResultObjectItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s DescribeVerifyPersonasDeviceModelStatisticsResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyPersonasDeviceModelStatisticsResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeVerifyPersonasDeviceModelStatisticsResponseBodyResultObject) GetAllDeviceCnt() *int64 {
	return s.AllDeviceCnt
}

func (s *DescribeVerifyPersonasDeviceModelStatisticsResponseBodyResultObject) GetItems() []*DescribeVerifyPersonasDeviceModelStatisticsResponseBodyResultObjectItems {
	return s.Items
}

func (s *DescribeVerifyPersonasDeviceModelStatisticsResponseBodyResultObject) SetAllDeviceCnt(v int64) *DescribeVerifyPersonasDeviceModelStatisticsResponseBodyResultObject {
	s.AllDeviceCnt = &v
	return s
}

func (s *DescribeVerifyPersonasDeviceModelStatisticsResponseBodyResultObject) SetItems(v []*DescribeVerifyPersonasDeviceModelStatisticsResponseBodyResultObjectItems) *DescribeVerifyPersonasDeviceModelStatisticsResponseBodyResultObject {
	s.Items = v
	return s
}

func (s *DescribeVerifyPersonasDeviceModelStatisticsResponseBodyResultObject) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVerifyPersonasDeviceModelStatisticsResponseBodyResultObjectItems struct {
	// Number of devices.
	//
	// example:
	//
	// 5
	DeviceCnt *int64 `json:"DeviceCnt,omitempty" xml:"DeviceCnt,omitempty"`
	// Device model
	//
	// example:
	//
	// iPhone15,2
	DeviceModel *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	// The ratio of this device model to the total number of devices.
	//
	// example:
	//
	// 16.13
	DeviceRate *string `json:"DeviceRate,omitempty" xml:"DeviceRate,omitempty"`
}

func (s DescribeVerifyPersonasDeviceModelStatisticsResponseBodyResultObjectItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyPersonasDeviceModelStatisticsResponseBodyResultObjectItems) GoString() string {
	return s.String()
}

func (s *DescribeVerifyPersonasDeviceModelStatisticsResponseBodyResultObjectItems) GetDeviceCnt() *int64 {
	return s.DeviceCnt
}

func (s *DescribeVerifyPersonasDeviceModelStatisticsResponseBodyResultObjectItems) GetDeviceModel() *string {
	return s.DeviceModel
}

func (s *DescribeVerifyPersonasDeviceModelStatisticsResponseBodyResultObjectItems) GetDeviceRate() *string {
	return s.DeviceRate
}

func (s *DescribeVerifyPersonasDeviceModelStatisticsResponseBodyResultObjectItems) SetDeviceCnt(v int64) *DescribeVerifyPersonasDeviceModelStatisticsResponseBodyResultObjectItems {
	s.DeviceCnt = &v
	return s
}

func (s *DescribeVerifyPersonasDeviceModelStatisticsResponseBodyResultObjectItems) SetDeviceModel(v string) *DescribeVerifyPersonasDeviceModelStatisticsResponseBodyResultObjectItems {
	s.DeviceModel = &v
	return s
}

func (s *DescribeVerifyPersonasDeviceModelStatisticsResponseBodyResultObjectItems) SetDeviceRate(v string) *DescribeVerifyPersonasDeviceModelStatisticsResponseBodyResultObjectItems {
	s.DeviceRate = &v
	return s
}

func (s *DescribeVerifyPersonasDeviceModelStatisticsResponseBodyResultObjectItems) Validate() error {
	return dara.Validate(s)
}
