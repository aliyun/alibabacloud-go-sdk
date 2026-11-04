// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLocalRecordDownloadByTimeJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int32) *CreateLocalRecordDownloadByTimeJobRequest
	GetBeginTime() *int32
	SetDeviceName(v string) *CreateLocalRecordDownloadByTimeJobRequest
	GetDeviceName() *string
	SetEndTime(v int32) *CreateLocalRecordDownloadByTimeJobRequest
	GetEndTime() *int32
	SetIotId(v string) *CreateLocalRecordDownloadByTimeJobRequest
	GetIotId() *string
	SetIotInstanceId(v string) *CreateLocalRecordDownloadByTimeJobRequest
	GetIotInstanceId() *string
	SetProductKey(v string) *CreateLocalRecordDownloadByTimeJobRequest
	GetProductKey() *string
	SetSpeed(v float32) *CreateLocalRecordDownloadByTimeJobRequest
	GetSpeed() *float32
}

type CreateLocalRecordDownloadByTimeJobRequest struct {
	// The start time. The value must be a 10-digit timestamp in seconds. Valid values: 0 to 2147483647.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1900000000
	BeginTime *int32 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The DeviceName of the video device. You can view the **DeviceName*	- of a device under a specified product on the **Devices*	- page in the IoT Platform console or by calling the [QueryDevice](https://help.aliyun.com/document_detail/69905.html) operation.
	//
	// 	Notice: If you specify this parameter, you must also specify **ProductKey**.
	//
	// example:
	//
	// Device01
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// The end time. The value must be a 10-digit timestamp in seconds. Valid values: 0 to 2147483647. The value must be greater than the start time specified by **BeginTime**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2100000000
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the video device. You can call the [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) operation to query the ID.
	//
	// 	Notice:
	//
	// **IotId*	- is the unique identifier of the device and corresponds to the combination of **ProductKey*	- and **DeviceName**. If you specify this parameter, you do not need to specify **ProductKey*	- or **DeviceName**. If you specify both **IotId*	- and the combination of **ProductKey*	- and **DeviceName**, **IotId*	- takes precedence.
	//
	// example:
	//
	// C47T6xwp6ms4bNlkHRWCg4****
	IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// The instance ID. You can view the ID of the current instance on the **Instance Overview*	- tab in the IoT Platform console.
	//
	// This parameter is required.
	//
	// example:
	//
	// iot-******
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	// The **ProductKey*	- of the product to which the video device belongs.
	//
	// You can view the information of all products under the current instance on the **Products*	- page in the IoT Platform console or by calling the [QueryProductList](https://help.aliyun.com/document_detail/69271.html) operation.
	//
	// 	Notice: If you specify this parameter, you must also specify **DeviceName**.
	//
	// example:
	//
	// a1Bw******
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	// The upload speed multiplier of the device. Default value: **1.0**. Valid values: **0.125*	- to **8.0**.
	//
	// example:
	//
	// 1.0
	Speed *float32 `json:"Speed,omitempty" xml:"Speed,omitempty"`
}

func (s CreateLocalRecordDownloadByTimeJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLocalRecordDownloadByTimeJobRequest) GoString() string {
	return s.String()
}

func (s *CreateLocalRecordDownloadByTimeJobRequest) GetBeginTime() *int32 {
	return s.BeginTime
}

func (s *CreateLocalRecordDownloadByTimeJobRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *CreateLocalRecordDownloadByTimeJobRequest) GetEndTime() *int32 {
	return s.EndTime
}

func (s *CreateLocalRecordDownloadByTimeJobRequest) GetIotId() *string {
	return s.IotId
}

func (s *CreateLocalRecordDownloadByTimeJobRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *CreateLocalRecordDownloadByTimeJobRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *CreateLocalRecordDownloadByTimeJobRequest) GetSpeed() *float32 {
	return s.Speed
}

func (s *CreateLocalRecordDownloadByTimeJobRequest) SetBeginTime(v int32) *CreateLocalRecordDownloadByTimeJobRequest {
	s.BeginTime = &v
	return s
}

func (s *CreateLocalRecordDownloadByTimeJobRequest) SetDeviceName(v string) *CreateLocalRecordDownloadByTimeJobRequest {
	s.DeviceName = &v
	return s
}

func (s *CreateLocalRecordDownloadByTimeJobRequest) SetEndTime(v int32) *CreateLocalRecordDownloadByTimeJobRequest {
	s.EndTime = &v
	return s
}

func (s *CreateLocalRecordDownloadByTimeJobRequest) SetIotId(v string) *CreateLocalRecordDownloadByTimeJobRequest {
	s.IotId = &v
	return s
}

func (s *CreateLocalRecordDownloadByTimeJobRequest) SetIotInstanceId(v string) *CreateLocalRecordDownloadByTimeJobRequest {
	s.IotInstanceId = &v
	return s
}

func (s *CreateLocalRecordDownloadByTimeJobRequest) SetProductKey(v string) *CreateLocalRecordDownloadByTimeJobRequest {
	s.ProductKey = &v
	return s
}

func (s *CreateLocalRecordDownloadByTimeJobRequest) SetSpeed(v float32) *CreateLocalRecordDownloadByTimeJobRequest {
	s.Speed = &v
	return s
}

func (s *CreateLocalRecordDownloadByTimeJobRequest) Validate() error {
	return dara.Validate(s)
}
