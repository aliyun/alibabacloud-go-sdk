// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStreamPushJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceName(v string) *CreateStreamPushJobRequest
	GetDeviceName() *string
	SetIotId(v string) *CreateStreamPushJobRequest
	GetIotId() *string
	SetIotInstanceId(v string) *CreateStreamPushJobRequest
	GetIotInstanceId() *string
	SetJobType(v int32) *CreateStreamPushJobRequest
	GetJobType() *int32
	SetProductKey(v string) *CreateStreamPushJobRequest
	GetProductKey() *string
	SetPushUrl(v string) *CreateStreamPushJobRequest
	GetPushUrl() *string
	SetStreamType(v int32) *CreateStreamPushJobRequest
	GetStreamType() *int32
}

type CreateStreamPushJobRequest struct {
	// The name of the device. You can view the **DeviceName*	- on the **Devices*	- page in the IoT Platform console.
	//
	// 	Notice:  If you specify this parameter, you must also specify **ProductKey**.
	//
	// example:
	//
	// camera1
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// The ID of the device. The ID is issued by IoT Platform and is the unique identifier of the device. You can call the [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) operation of IoT Platform to query the ID.
	//
	// 	Notice:  As the unique identifier of the device, **IotId*	- has a one-to-one mapping with the combination of **ProductKey*	- and **DeviceName**. If you specify this parameter, you do not need to specify **ProductKey*	- or **DeviceName**. If you specify both **IotId*	- and the combination of **ProductKey*	- and **DeviceName**, **IotId*	- takes precedence.
	//
	// example:
	//
	// d7XmBoJhAr88C6PelXDF00****
	IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// The instance ID. You can view the ID of the current instance on the **Instance Overview*	- tab in the IoT Platform console.
	//
	// This parameter is required.
	//
	// example:
	//
	// iot-cn-n6w1y59****
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	// The type of the stream pushing job. Valid values:
	//
	// - **1*	- (default): continuous job.
	//
	// - **2**: triggered job, which ends after an interruption.
	//
	// example:
	//
	// 1
	JobType *int32 `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The **ProductKey*	- of the product to which the device belongs. You can view the ProductKey on the **Products*	- page in the IoT Platform console or call the [QueryProductList](https://help.aliyun.com/document_detail/69271.html) operation to query the information about all products under the current account.
	//
	// 	Notice:  If you specify this parameter, you must also specify **DeviceName**.
	//
	// example:
	//
	// a1BwAGV****
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	// The push URL. The maximum length is 500 bytes.
	//
	// This parameter is required.
	//
	// example:
	//
	// rtmp://****
	PushUrl *string `json:"PushUrl,omitempty" xml:"PushUrl,omitempty"`
	// The type of the video stream. Valid values:
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

func (s CreateStreamPushJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStreamPushJobRequest) GoString() string {
	return s.String()
}

func (s *CreateStreamPushJobRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *CreateStreamPushJobRequest) GetIotId() *string {
	return s.IotId
}

func (s *CreateStreamPushJobRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *CreateStreamPushJobRequest) GetJobType() *int32 {
	return s.JobType
}

func (s *CreateStreamPushJobRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *CreateStreamPushJobRequest) GetPushUrl() *string {
	return s.PushUrl
}

func (s *CreateStreamPushJobRequest) GetStreamType() *int32 {
	return s.StreamType
}

func (s *CreateStreamPushJobRequest) SetDeviceName(v string) *CreateStreamPushJobRequest {
	s.DeviceName = &v
	return s
}

func (s *CreateStreamPushJobRequest) SetIotId(v string) *CreateStreamPushJobRequest {
	s.IotId = &v
	return s
}

func (s *CreateStreamPushJobRequest) SetIotInstanceId(v string) *CreateStreamPushJobRequest {
	s.IotInstanceId = &v
	return s
}

func (s *CreateStreamPushJobRequest) SetJobType(v int32) *CreateStreamPushJobRequest {
	s.JobType = &v
	return s
}

func (s *CreateStreamPushJobRequest) SetProductKey(v string) *CreateStreamPushJobRequest {
	s.ProductKey = &v
	return s
}

func (s *CreateStreamPushJobRequest) SetPushUrl(v string) *CreateStreamPushJobRequest {
	s.PushUrl = &v
	return s
}

func (s *CreateStreamPushJobRequest) SetStreamType(v int32) *CreateStreamPushJobRequest {
	s.StreamType = &v
	return s
}

func (s *CreateStreamPushJobRequest) Validate() error {
	return dara.Validate(s)
}
