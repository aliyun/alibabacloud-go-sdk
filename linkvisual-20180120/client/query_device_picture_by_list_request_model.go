// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDevicePictureByListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceName(v string) *QueryDevicePictureByListRequest
	GetDeviceName() *string
	SetExpireTime(v int32) *QueryDevicePictureByListRequest
	GetExpireTime() *int32
	SetIotId(v string) *QueryDevicePictureByListRequest
	GetIotId() *string
	SetIotInstanceId(v string) *QueryDevicePictureByListRequest
	GetIotInstanceId() *string
	SetPictureIdList(v []*string) *QueryDevicePictureByListRequest
	GetPictureIdList() []*string
	SetPictureType(v int32) *QueryDevicePictureByListRequest
	GetPictureType() *int32
	SetProductKey(v string) *QueryDevicePictureByListRequest
	GetProductKey() *string
	SetThumbWidth(v int32) *QueryDevicePictureByListRequest
	GetThumbWidth() *int32
}

type QueryDevicePictureByListRequest struct {
	// The DeviceName of the IP camera.
	//
	// You can view the **DeviceName*	- on the **Devices*	- page in the IoT Platform console.
	//
	// 	Notice:  If you specify this parameter, you must also specify **ProductKey**.
	//
	// example:
	//
	// camera1
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// The expiration time of the image URL. Unit: minutes.
	//
	// example:
	//
	// 60
	ExpireTime *int32 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the IP camera. The ID is issued by IoT Platform and uniquely identifies the device. You can call the [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) operation of IoT Platform to query the ID.
	//
	// 	Notice:  If you specify this parameter, you do not need to specify **ProductKey*	- or **DeviceName**. **IotId*	- uniquely identifies the device and corresponds to the combination of **ProductKey*	- and **DeviceName**. If you specify both **IotId*	- and the combination of **ProductKey*	- and **DeviceName**, **IotId*	- takes precedence.
	//
	// example:
	//
	// 1CJ0****e7qiq8VkQi1000000
	IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// The instance ID.
	//
	// <props="china">You can view the **ID*	- of the current instance on the **Instance Overview*	- page in the IoT Platform console. If the instance has an ID, you must specify the ID. Otherwise, the call fails.
	//
	// <props="china">
	//
	// > If the public instance does not have an ID, see [How to obtain an instance ID](https://help.aliyun.com/document_detail/267533.html).
	//
	// <props="intl">You do not need to specify this parameter.
	//
	// example:
	//
	// iot-cn-n6w1y******
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	// The list of image IDs. Image IDs are generated when you call the [TriggerCapturePicture](https://help.aliyun.com/document_detail/107914.html) operation.
	//
	// > You can specify up to 50 image IDs in a single call.
	//
	// This parameter is required.
	//
	// example:
	//
	// cFc5MzZ****npYT1***TDFzX1hrVzA0Z****
	PictureIdList []*string `json:"PictureIdList,omitempty" xml:"PictureIdList,omitempty" type:"Repeated"`
	// The image type. Valid values:
	//
	// - **0**: thumbnail.
	//
	// - **1**: original image.
	//
	// - **2**: all. Default value.
	//
	// example:
	//
	// 2
	PictureType *int32 `json:"PictureType,omitempty" xml:"PictureType,omitempty"`
	// The ProductKey of the product to which the IP camera belongs.
	//
	// You can view the ProductKey on the **Products*	- page in the IoT Platform console or call the [QueryProductList](https://help.aliyun.com/document_detail/69271.html) operation to query the information about all products under the current account.
	//
	// 	Notice:  If you specify this parameter, you must also specify **DeviceName**.
	//
	// example:
	//
	// a1BwAGV****
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	// Specifies the width of the target thumbnail. Unit: px. Valid values: 1 to 4096.
	//
	// example:
	//
	// 100
	ThumbWidth *int32 `json:"ThumbWidth,omitempty" xml:"ThumbWidth,omitempty"`
}

func (s QueryDevicePictureByListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDevicePictureByListRequest) GoString() string {
	return s.String()
}

func (s *QueryDevicePictureByListRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *QueryDevicePictureByListRequest) GetExpireTime() *int32 {
	return s.ExpireTime
}

func (s *QueryDevicePictureByListRequest) GetIotId() *string {
	return s.IotId
}

func (s *QueryDevicePictureByListRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *QueryDevicePictureByListRequest) GetPictureIdList() []*string {
	return s.PictureIdList
}

func (s *QueryDevicePictureByListRequest) GetPictureType() *int32 {
	return s.PictureType
}

func (s *QueryDevicePictureByListRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *QueryDevicePictureByListRequest) GetThumbWidth() *int32 {
	return s.ThumbWidth
}

func (s *QueryDevicePictureByListRequest) SetDeviceName(v string) *QueryDevicePictureByListRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryDevicePictureByListRequest) SetExpireTime(v int32) *QueryDevicePictureByListRequest {
	s.ExpireTime = &v
	return s
}

func (s *QueryDevicePictureByListRequest) SetIotId(v string) *QueryDevicePictureByListRequest {
	s.IotId = &v
	return s
}

func (s *QueryDevicePictureByListRequest) SetIotInstanceId(v string) *QueryDevicePictureByListRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryDevicePictureByListRequest) SetPictureIdList(v []*string) *QueryDevicePictureByListRequest {
	s.PictureIdList = v
	return s
}

func (s *QueryDevicePictureByListRequest) SetPictureType(v int32) *QueryDevicePictureByListRequest {
	s.PictureType = &v
	return s
}

func (s *QueryDevicePictureByListRequest) SetProductKey(v string) *QueryDevicePictureByListRequest {
	s.ProductKey = &v
	return s
}

func (s *QueryDevicePictureByListRequest) SetThumbWidth(v int32) *QueryDevicePictureByListRequest {
	s.ThumbWidth = &v
	return s
}

func (s *QueryDevicePictureByListRequest) Validate() error {
	return dara.Validate(s)
}
