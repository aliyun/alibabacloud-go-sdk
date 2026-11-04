// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDevicePictureFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCaptureId(v string) *QueryDevicePictureFileRequest
	GetCaptureId() *string
	SetDeviceName(v string) *QueryDevicePictureFileRequest
	GetDeviceName() *string
	SetExpireTime(v int32) *QueryDevicePictureFileRequest
	GetExpireTime() *int32
	SetIotId(v string) *QueryDevicePictureFileRequest
	GetIotId() *string
	SetIotInstanceId(v string) *QueryDevicePictureFileRequest
	GetIotInstanceId() *string
	SetPictureType(v int32) *QueryDevicePictureFileRequest
	GetPictureType() *int32
	SetProductKey(v string) *QueryDevicePictureFileRequest
	GetProductKey() *string
	SetThumbWidth(v int32) *QueryDevicePictureFileRequest
	GetThumbWidth() *int32
}

type QueryDevicePictureFileRequest struct {
	// The ID of the image generated when a snapshot capture event occurs:
	//
	// - For active snapshot capture, see [TriggerCapturePicture](https://help.aliyun.com/document_detail/107914.html) for the image ID.
	//
	// - For snapshot capture triggered by an intelligent alert event, see [QueryDeviceEvent](https://help.aliyun.com/document_detail/107916.html) for the image ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// fdc576d03c094faea9fa10fbf48c****
	CaptureId *string `json:"CaptureId,omitempty" xml:"CaptureId,omitempty"`
	// The DeviceName of the IPC device.
	//
	// You can view the **DeviceName*	- on the **Devices*	- page in the IoT Platform console.
	//
	// 	Notice:  If you specify this parameter, you must also specify **ProductKey**.
	//
	// example:
	//
	// camera1
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// Specifies the expiration time of the image URL. Unit: minutes. Valid values: 1 to 43200.
	//
	// example:
	//
	// 60
	ExpireTime *int32 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The IotId of the IPC device. The IotId is the globally unique identifier (GUID) that is issued by IoT Platform to the device. You can call the [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) operation of IoT Platform to query the IotId.
	//
	// 	Notice:  If you specify this parameter, you do not need to specify **ProductKey*	- or **DeviceName**. The **IotId*	- is the globally unique identifier of the device and corresponds to the combination of **ProductKey*	- and **DeviceName**. If you specify both **IotId*	- and the combination of **ProductKey*	- and **DeviceName**, **IotId*	- takes precedence.
	//
	// example:
	//
	// zLZyi6aOLyOSHa9hsPyD00****
	IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// The instance ID. You can view the **ID*	- of the current instance on the **Overview*	- page in the IoT Platform console.
	//
	// 	Notice:
	//
	// - If the instance has an ID, you must specify the ID. Otherwise, the call fails.
	//
	// - If the <props="intl"><ph>**Overview*	- page is not displayed or the instance does not have an </ph>ID, you do not need to specify this parameter.
	//
	//
	//
	// For more information about instances, see [Instance overview](https://help.aliyun.com/document_detail/356505.html).
	//
	// example:
	//
	// iot-cn-n6w1y******
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	// The image type. Valid values:
	//
	// - **0**: thumbnail.
	//
	// - **1**: original image.
	//
	// - **2**: all images. This is the default value.
	//
	// example:
	//
	// 1
	PictureType *int32 `json:"PictureType,omitempty" xml:"PictureType,omitempty"`
	// The ProductKey of the product to which the IPC device belongs.
	//
	// You can view the ProductKey on the **Products*	- page in the IoT Platform console or call the [QueryProductList](https://help.aliyun.com/document_detail/69271.html) operation to query the information about all products under the current account.
	//
	// 	Notice:  If you specify this parameter, you must also specify **DeviceName**.
	//
	// example:
	//
	// a1BwAGV****
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	// Specifies the width of the target scaled image. Unit: px. Valid values: 1 to 4096.
	//
	// example:
	//
	// 100
	ThumbWidth *int32 `json:"ThumbWidth,omitempty" xml:"ThumbWidth,omitempty"`
}

func (s QueryDevicePictureFileRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDevicePictureFileRequest) GoString() string {
	return s.String()
}

func (s *QueryDevicePictureFileRequest) GetCaptureId() *string {
	return s.CaptureId
}

func (s *QueryDevicePictureFileRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *QueryDevicePictureFileRequest) GetExpireTime() *int32 {
	return s.ExpireTime
}

func (s *QueryDevicePictureFileRequest) GetIotId() *string {
	return s.IotId
}

func (s *QueryDevicePictureFileRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *QueryDevicePictureFileRequest) GetPictureType() *int32 {
	return s.PictureType
}

func (s *QueryDevicePictureFileRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *QueryDevicePictureFileRequest) GetThumbWidth() *int32 {
	return s.ThumbWidth
}

func (s *QueryDevicePictureFileRequest) SetCaptureId(v string) *QueryDevicePictureFileRequest {
	s.CaptureId = &v
	return s
}

func (s *QueryDevicePictureFileRequest) SetDeviceName(v string) *QueryDevicePictureFileRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryDevicePictureFileRequest) SetExpireTime(v int32) *QueryDevicePictureFileRequest {
	s.ExpireTime = &v
	return s
}

func (s *QueryDevicePictureFileRequest) SetIotId(v string) *QueryDevicePictureFileRequest {
	s.IotId = &v
	return s
}

func (s *QueryDevicePictureFileRequest) SetIotInstanceId(v string) *QueryDevicePictureFileRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryDevicePictureFileRequest) SetPictureType(v int32) *QueryDevicePictureFileRequest {
	s.PictureType = &v
	return s
}

func (s *QueryDevicePictureFileRequest) SetProductKey(v string) *QueryDevicePictureFileRequest {
	s.ProductKey = &v
	return s
}

func (s *QueryDevicePictureFileRequest) SetThumbWidth(v int32) *QueryDevicePictureFileRequest {
	s.ThumbWidth = &v
	return s
}

func (s *QueryDevicePictureFileRequest) Validate() error {
	return dara.Validate(s)
}
