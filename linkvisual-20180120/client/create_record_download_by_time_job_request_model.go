// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecordDownloadByTimeJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int32) *CreateRecordDownloadByTimeJobRequest
	GetBeginTime() *int32
	SetDeviceName(v string) *CreateRecordDownloadByTimeJobRequest
	GetDeviceName() *string
	SetEndTime(v int32) *CreateRecordDownloadByTimeJobRequest
	GetEndTime() *int32
	SetIotId(v string) *CreateRecordDownloadByTimeJobRequest
	GetIotId() *string
	SetIotInstanceId(v string) *CreateRecordDownloadByTimeJobRequest
	GetIotInstanceId() *string
	SetProductKey(v string) *CreateRecordDownloadByTimeJobRequest
	GetProductKey() *string
	SetRecordType(v int32) *CreateRecordDownloadByTimeJobRequest
	GetRecordType() *int32
	SetStreamType(v int32) *CreateRecordDownloadByTimeJobRequest
	GetStreamType() *int32
}

type CreateRecordDownloadByTimeJobRequest struct {
	// The start time. The value must be a 10-digit UNIX timestamp in seconds. Valid values: 0 to 2147483647.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1900000000
	BeginTime *int32 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The DeviceName of the video device. You can view the **DeviceName*	- of devices under a specified product on the **Devices*	- page in the [IoT Platform console](https://iot.console.aliyun.com/devices) or by calling the [QueryDevice](https://help.aliyun.com/document_detail/69905.html) operation.
	//
	// 	Notice:  If you specify this parameter, you must also specify **ProductKey**.
	//
	// example:
	//
	// Device01
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// The end time. The value must be a 10-digit UNIX timestamp in seconds. Valid values: 0 to 2147483647. The value must be greater than the start time specified by **BeginTime**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2100000000
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the video device. You can call the [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) operation to query the ID.
	//
	// 	Notice:  **IotId*	- is the unique identifier of the device and corresponds to the combination of **ProductKey*	- and **DeviceName**. If you specify this parameter, you do not need to specify **ProductKey*	- or **DeviceName**. If you specify both **IotId*	- and the combination of **ProductKey*	- and **DeviceName**, **IotId*	- takes precedence.
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
	// You can view the information about all products under the current instance on the **Products*	- page in the [IoT Platform console](https://iot.console.aliyun.com/product) or by calling the [QueryProductList](https://help.aliyun.com/document_detail/69271.html) operation.
	//
	// 	Notice:  If you specify this parameter, you must also specify **DeviceName**.
	//
	// example:
	//
	// a1Bw******
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	// The recording type. Valid values:
	//
	// - **0**: continuous recording.
	//
	// - **1**: event recording.
	//
	// - **2**: on-demand recording.
	//
	// - **4**: local-to-cloud recording.
	//
	// - **Empty**: all recordings.
	//
	// example:
	//
	// 0
	RecordType *int32 `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	// The stream type. Valid values:
	//
	// - **0*	- (default): main stream.
	//
	// - **1**: sub stream.
	//
	// example:
	//
	// 0
	StreamType *int32 `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
}

func (s CreateRecordDownloadByTimeJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRecordDownloadByTimeJobRequest) GoString() string {
	return s.String()
}

func (s *CreateRecordDownloadByTimeJobRequest) GetBeginTime() *int32 {
	return s.BeginTime
}

func (s *CreateRecordDownloadByTimeJobRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *CreateRecordDownloadByTimeJobRequest) GetEndTime() *int32 {
	return s.EndTime
}

func (s *CreateRecordDownloadByTimeJobRequest) GetIotId() *string {
	return s.IotId
}

func (s *CreateRecordDownloadByTimeJobRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *CreateRecordDownloadByTimeJobRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *CreateRecordDownloadByTimeJobRequest) GetRecordType() *int32 {
	return s.RecordType
}

func (s *CreateRecordDownloadByTimeJobRequest) GetStreamType() *int32 {
	return s.StreamType
}

func (s *CreateRecordDownloadByTimeJobRequest) SetBeginTime(v int32) *CreateRecordDownloadByTimeJobRequest {
	s.BeginTime = &v
	return s
}

func (s *CreateRecordDownloadByTimeJobRequest) SetDeviceName(v string) *CreateRecordDownloadByTimeJobRequest {
	s.DeviceName = &v
	return s
}

func (s *CreateRecordDownloadByTimeJobRequest) SetEndTime(v int32) *CreateRecordDownloadByTimeJobRequest {
	s.EndTime = &v
	return s
}

func (s *CreateRecordDownloadByTimeJobRequest) SetIotId(v string) *CreateRecordDownloadByTimeJobRequest {
	s.IotId = &v
	return s
}

func (s *CreateRecordDownloadByTimeJobRequest) SetIotInstanceId(v string) *CreateRecordDownloadByTimeJobRequest {
	s.IotInstanceId = &v
	return s
}

func (s *CreateRecordDownloadByTimeJobRequest) SetProductKey(v string) *CreateRecordDownloadByTimeJobRequest {
	s.ProductKey = &v
	return s
}

func (s *CreateRecordDownloadByTimeJobRequest) SetRecordType(v int32) *CreateRecordDownloadByTimeJobRequest {
	s.RecordType = &v
	return s
}

func (s *CreateRecordDownloadByTimeJobRequest) SetStreamType(v int32) *CreateRecordDownloadByTimeJobRequest {
	s.StreamType = &v
	return s
}

func (s *CreateRecordDownloadByTimeJobRequest) Validate() error {
	return dara.Validate(s)
}
