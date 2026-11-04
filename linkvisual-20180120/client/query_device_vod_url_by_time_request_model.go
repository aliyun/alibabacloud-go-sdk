// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDeviceVodUrlByTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int32) *QueryDeviceVodUrlByTimeRequest
	GetBeginTime() *int32
	SetDeviceName(v string) *QueryDeviceVodUrlByTimeRequest
	GetDeviceName() *string
	SetEnableStun(v bool) *QueryDeviceVodUrlByTimeRequest
	GetEnableStun() *bool
	SetEncryptType(v int32) *QueryDeviceVodUrlByTimeRequest
	GetEncryptType() *int32
	SetEndTime(v int32) *QueryDeviceVodUrlByTimeRequest
	GetEndTime() *int32
	SetIotId(v string) *QueryDeviceVodUrlByTimeRequest
	GetIotId() *string
	SetIotInstanceId(v string) *QueryDeviceVodUrlByTimeRequest
	GetIotInstanceId() *string
	SetPlayUnLimited(v bool) *QueryDeviceVodUrlByTimeRequest
	GetPlayUnLimited() *bool
	SetProductKey(v string) *QueryDeviceVodUrlByTimeRequest
	GetProductKey() *string
	SetScheme(v string) *QueryDeviceVodUrlByTimeRequest
	GetScheme() *string
	SetSeekTime(v int32) *QueryDeviceVodUrlByTimeRequest
	GetSeekTime() *int32
	SetShouldEncrypt(v bool) *QueryDeviceVodUrlByTimeRequest
	GetShouldEncrypt() *bool
	SetUrlValidDuration(v int32) *QueryDeviceVodUrlByTimeRequest
	GetUrlValidDuration() *int32
}

type QueryDeviceVodUrlByTimeRequest struct {
	// The recording start time. The value must be a 10-digit UNIX timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1508200012
	BeginTime *int32 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
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
	EnableStun *bool   `json:"EnableStun,omitempty" xml:"EnableStun,omitempty"`
	// The encryption type. Only I-frame encryption is supported. Set the value to 0.
	//
	// example:
	//
	// 0
	EncryptType *int32 `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	// The recording end time. The value must be a 10-digit UNIX timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1582420940
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the IPC device. The ID is issued by IoT Platform and uniquely identifies the device. You can call the [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) operation of IoT Platform to query the ID.
	//
	// 	Notice:  If you specify this parameter, you do not need to specify **ProductKey*	- or **DeviceName**. The **IotId*	- uniquely identifies the device and has a one-to-one mapping with the combination of **ProductKey*	- and **DeviceName**. If you specify both **IotId*	- and the combination of **ProductKey*	- and **DeviceName**, **IotId*	- takes precedence.
	//
	// example:
	//
	// C47T6xwp6ms4bNlkHRWCg4****
	IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// The instance ID. You can view the **ID*	- of the current instance on the **Instance Overview*	- page in the IoT Platform console.
	//
	// 	Notice:
	//
	// - If the instance has an ID, you must specify this ID. Otherwise, the call fails.
	//
	// - If the instance does not have an <props="intl"><ph>**Instance Overview*	- page or </ph>ID, you do not need to specify this parameter.
	//
	//
	//
	// For more information about instances, see [Instance overview](https://help.aliyun.com/document_detail/356505.html).
	//
	// example:
	//
	// iot-cn-n6w1y****
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	// Specifies whether the playback URL can be used an unlimited number of times. Valid values:
	//
	// - **true**: The URL can be used an unlimited number of times within 10 seconds.
	//
	// - **false**: The URL can be used only once. This is the default value.
	//
	// example:
	//
	// false
	PlayUnLimited *bool `json:"PlayUnLimited,omitempty" xml:"PlayUnLimited,omitempty"`
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
	// The protocol type. Valid values:
	//
	// - **rtmp**: RTMP.
	//
	// - **flv**: HTTP-FLV.
	//
	// - **hls**: HLS.
	//
	//  > For more information about the limits of HLS, see [Cloud development](https://help.aliyun.com/document_detail/208584.html).
	//
	// example:
	//
	// rtmp
	Scheme *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
	// The offset relative to the recording start time specified by **BeginTime**. Unit: seconds. Default value: 0.
	//
	// example:
	//
	// 0
	SeekTime *int32 `json:"SeekTime,omitempty" xml:"SeekTime,omitempty"`
	// Specifies whether to encrypt the stream. Valid values:
	//
	// - **true**: Encrypts the stream.
	//
	// - **false**: Does not encrypt the stream. This is the default value.
	//
	// example:
	//
	// false
	ShouldEncrypt *bool `json:"ShouldEncrypt,omitempty" xml:"ShouldEncrypt,omitempty"`
	// The validity period of the URL. Valid values: 10 to 14400. Unit: seconds. Default value: 10.
	//
	// example:
	//
	// 60
	UrlValidDuration *int32 `json:"UrlValidDuration,omitempty" xml:"UrlValidDuration,omitempty"`
}

func (s QueryDeviceVodUrlByTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDeviceVodUrlByTimeRequest) GoString() string {
	return s.String()
}

func (s *QueryDeviceVodUrlByTimeRequest) GetBeginTime() *int32 {
	return s.BeginTime
}

func (s *QueryDeviceVodUrlByTimeRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *QueryDeviceVodUrlByTimeRequest) GetEnableStun() *bool {
	return s.EnableStun
}

func (s *QueryDeviceVodUrlByTimeRequest) GetEncryptType() *int32 {
	return s.EncryptType
}

func (s *QueryDeviceVodUrlByTimeRequest) GetEndTime() *int32 {
	return s.EndTime
}

func (s *QueryDeviceVodUrlByTimeRequest) GetIotId() *string {
	return s.IotId
}

func (s *QueryDeviceVodUrlByTimeRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *QueryDeviceVodUrlByTimeRequest) GetPlayUnLimited() *bool {
	return s.PlayUnLimited
}

func (s *QueryDeviceVodUrlByTimeRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *QueryDeviceVodUrlByTimeRequest) GetScheme() *string {
	return s.Scheme
}

func (s *QueryDeviceVodUrlByTimeRequest) GetSeekTime() *int32 {
	return s.SeekTime
}

func (s *QueryDeviceVodUrlByTimeRequest) GetShouldEncrypt() *bool {
	return s.ShouldEncrypt
}

func (s *QueryDeviceVodUrlByTimeRequest) GetUrlValidDuration() *int32 {
	return s.UrlValidDuration
}

func (s *QueryDeviceVodUrlByTimeRequest) SetBeginTime(v int32) *QueryDeviceVodUrlByTimeRequest {
	s.BeginTime = &v
	return s
}

func (s *QueryDeviceVodUrlByTimeRequest) SetDeviceName(v string) *QueryDeviceVodUrlByTimeRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryDeviceVodUrlByTimeRequest) SetEnableStun(v bool) *QueryDeviceVodUrlByTimeRequest {
	s.EnableStun = &v
	return s
}

func (s *QueryDeviceVodUrlByTimeRequest) SetEncryptType(v int32) *QueryDeviceVodUrlByTimeRequest {
	s.EncryptType = &v
	return s
}

func (s *QueryDeviceVodUrlByTimeRequest) SetEndTime(v int32) *QueryDeviceVodUrlByTimeRequest {
	s.EndTime = &v
	return s
}

func (s *QueryDeviceVodUrlByTimeRequest) SetIotId(v string) *QueryDeviceVodUrlByTimeRequest {
	s.IotId = &v
	return s
}

func (s *QueryDeviceVodUrlByTimeRequest) SetIotInstanceId(v string) *QueryDeviceVodUrlByTimeRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryDeviceVodUrlByTimeRequest) SetPlayUnLimited(v bool) *QueryDeviceVodUrlByTimeRequest {
	s.PlayUnLimited = &v
	return s
}

func (s *QueryDeviceVodUrlByTimeRequest) SetProductKey(v string) *QueryDeviceVodUrlByTimeRequest {
	s.ProductKey = &v
	return s
}

func (s *QueryDeviceVodUrlByTimeRequest) SetScheme(v string) *QueryDeviceVodUrlByTimeRequest {
	s.Scheme = &v
	return s
}

func (s *QueryDeviceVodUrlByTimeRequest) SetSeekTime(v int32) *QueryDeviceVodUrlByTimeRequest {
	s.SeekTime = &v
	return s
}

func (s *QueryDeviceVodUrlByTimeRequest) SetShouldEncrypt(v bool) *QueryDeviceVodUrlByTimeRequest {
	s.ShouldEncrypt = &v
	return s
}

func (s *QueryDeviceVodUrlByTimeRequest) SetUrlValidDuration(v int32) *QueryDeviceVodUrlByTimeRequest {
	s.UrlValidDuration = &v
	return s
}

func (s *QueryDeviceVodUrlByTimeRequest) Validate() error {
	return dara.Validate(s)
}
