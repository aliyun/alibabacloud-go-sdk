// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDeviceVodUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceName(v string) *QueryDeviceVodUrlRequest
	GetDeviceName() *string
	SetEnableStun(v bool) *QueryDeviceVodUrlRequest
	GetEnableStun() *bool
	SetEncryptType(v int32) *QueryDeviceVodUrlRequest
	GetEncryptType() *int32
	SetFileName(v string) *QueryDeviceVodUrlRequest
	GetFileName() *string
	SetIotId(v string) *QueryDeviceVodUrlRequest
	GetIotId() *string
	SetIotInstanceId(v string) *QueryDeviceVodUrlRequest
	GetIotInstanceId() *string
	SetPlayUnLimited(v bool) *QueryDeviceVodUrlRequest
	GetPlayUnLimited() *bool
	SetProductKey(v string) *QueryDeviceVodUrlRequest
	GetProductKey() *string
	SetScheme(v string) *QueryDeviceVodUrlRequest
	GetScheme() *string
	SetSeekTime(v int32) *QueryDeviceVodUrlRequest
	GetSeekTime() *int32
	SetShouldEncrypt(v bool) *QueryDeviceVodUrlRequest
	GetShouldEncrypt() *bool
	SetUrlValidDuration(v int32) *QueryDeviceVodUrlRequest
	GetUrlValidDuration() *int32
}

type QueryDeviceVodUrlRequest struct {
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
	EnableStun *bool   `json:"EnableStun,omitempty" xml:"EnableStun,omitempty"`
	// The encryption type. Only I-frame encryption is supported. Set the value to 0.
	//
	// example:
	//
	// 0
	EncryptType *int32 `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	// The name of the recording file. You can call the [InvokeThingService](https://help.aliyun.com/document_detail/69584.html) operation provided by IoT Platform to query the name. For request parameter examples, see **Supplementary description of request parameters*	- below.
	//
	// > The name of the recording file cannot exceed 180 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// file1
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The ID of the IP camera. The ID is issued by IoT Platform and uniquely identifies the device. You can call the [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) operation of IoT Platform to query the ID.
	//
	// 	Notice:  If you specify this parameter, you do not need to specify **ProductKey*	- or **DeviceName**. The **IotId*	- parameter uniquely identifies the device and has a one-to-one mapping with the combination of **ProductKey*	- and **DeviceName**. If you specify both **IotId*	- and the combination of **ProductKey*	- and **DeviceName**, **IotId*	- takes precedence.
	//
	// example:
	//
	// zLZyi6aOLyOSHa9hsPyD00****
	IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// The instance ID. You can view the **ID*	- of the current instance on the **Instance Overview*	- page in the IoT Platform console.
	//
	// 	Notice:
	//
	// - If an instance ID is specified, you must specify this parameter. Otherwise, the call fails.
	//
	// - If no <props="intl"><ph>**Instance Overview*	- page or </ph>instance ID is available, you do not need to specify this parameter.
	//
	//
	//
	// For more information about instances, see [Instance overview](https://help.aliyun.com/document_detail/356505.html).
	//
	// example:
	//
	// iot-cn-n6w1y59****
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
	// The offset relative to the **BeginTime*	- parameter, which specifies the start time of the recording. Unit: seconds. Default value: 0.
	//
	// example:
	//
	// 0
	SeekTime *int32 `json:"SeekTime,omitempty" xml:"SeekTime,omitempty"`
	// Specifies whether to encrypt the recording file. Valid values:
	//
	// - **true**: Encrypts the recording file.
	//
	// - **false**: Does not encrypt the recording file. This is the default value.
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

func (s QueryDeviceVodUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDeviceVodUrlRequest) GoString() string {
	return s.String()
}

func (s *QueryDeviceVodUrlRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *QueryDeviceVodUrlRequest) GetEnableStun() *bool {
	return s.EnableStun
}

func (s *QueryDeviceVodUrlRequest) GetEncryptType() *int32 {
	return s.EncryptType
}

func (s *QueryDeviceVodUrlRequest) GetFileName() *string {
	return s.FileName
}

func (s *QueryDeviceVodUrlRequest) GetIotId() *string {
	return s.IotId
}

func (s *QueryDeviceVodUrlRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *QueryDeviceVodUrlRequest) GetPlayUnLimited() *bool {
	return s.PlayUnLimited
}

func (s *QueryDeviceVodUrlRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *QueryDeviceVodUrlRequest) GetScheme() *string {
	return s.Scheme
}

func (s *QueryDeviceVodUrlRequest) GetSeekTime() *int32 {
	return s.SeekTime
}

func (s *QueryDeviceVodUrlRequest) GetShouldEncrypt() *bool {
	return s.ShouldEncrypt
}

func (s *QueryDeviceVodUrlRequest) GetUrlValidDuration() *int32 {
	return s.UrlValidDuration
}

func (s *QueryDeviceVodUrlRequest) SetDeviceName(v string) *QueryDeviceVodUrlRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryDeviceVodUrlRequest) SetEnableStun(v bool) *QueryDeviceVodUrlRequest {
	s.EnableStun = &v
	return s
}

func (s *QueryDeviceVodUrlRequest) SetEncryptType(v int32) *QueryDeviceVodUrlRequest {
	s.EncryptType = &v
	return s
}

func (s *QueryDeviceVodUrlRequest) SetFileName(v string) *QueryDeviceVodUrlRequest {
	s.FileName = &v
	return s
}

func (s *QueryDeviceVodUrlRequest) SetIotId(v string) *QueryDeviceVodUrlRequest {
	s.IotId = &v
	return s
}

func (s *QueryDeviceVodUrlRequest) SetIotInstanceId(v string) *QueryDeviceVodUrlRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryDeviceVodUrlRequest) SetPlayUnLimited(v bool) *QueryDeviceVodUrlRequest {
	s.PlayUnLimited = &v
	return s
}

func (s *QueryDeviceVodUrlRequest) SetProductKey(v string) *QueryDeviceVodUrlRequest {
	s.ProductKey = &v
	return s
}

func (s *QueryDeviceVodUrlRequest) SetScheme(v string) *QueryDeviceVodUrlRequest {
	s.Scheme = &v
	return s
}

func (s *QueryDeviceVodUrlRequest) SetSeekTime(v int32) *QueryDeviceVodUrlRequest {
	s.SeekTime = &v
	return s
}

func (s *QueryDeviceVodUrlRequest) SetShouldEncrypt(v bool) *QueryDeviceVodUrlRequest {
	s.ShouldEncrypt = &v
	return s
}

func (s *QueryDeviceVodUrlRequest) SetUrlValidDuration(v int32) *QueryDeviceVodUrlRequest {
	s.UrlValidDuration = &v
	return s
}

func (s *QueryDeviceVodUrlRequest) Validate() error {
	return dara.Validate(s)
}
