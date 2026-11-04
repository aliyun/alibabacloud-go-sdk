// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLiveStreamingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCacheDuration(v int32) *QueryLiveStreamingRequest
	GetCacheDuration() *int32
	SetDeviceName(v string) *QueryLiveStreamingRequest
	GetDeviceName() *string
	SetEnableStun(v bool) *QueryLiveStreamingRequest
	GetEnableStun() *bool
	SetEncryptType(v int32) *QueryLiveStreamingRequest
	GetEncryptType() *int32
	SetForceIFrame(v bool) *QueryLiveStreamingRequest
	GetForceIFrame() *bool
	SetIotId(v string) *QueryLiveStreamingRequest
	GetIotId() *string
	SetIotInstanceId(v string) *QueryLiveStreamingRequest
	GetIotInstanceId() *string
	SetPlayUnLimited(v bool) *QueryLiveStreamingRequest
	GetPlayUnLimited() *bool
	SetProductKey(v string) *QueryLiveStreamingRequest
	GetProductKey() *string
	SetScheme(v string) *QueryLiveStreamingRequest
	GetScheme() *string
	SetShouldEncrypt(v bool) *QueryLiveStreamingRequest
	GetShouldEncrypt() *bool
	SetStreamType(v int32) *QueryLiveStreamingRequest
	GetStreamType() *int32
	SetUrlValidDuration(v int32) *QueryLiveStreamingRequest
	GetUrlValidDuration() *int32
}

type QueryLiveStreamingRequest struct {
	// The duration of video cached data. Valid values: 0 to 10000. Unit: milliseconds. Default value: 0.
	//
	// example:
	//
	// 0
	CacheDuration *int32 `json:"CacheDuration,omitempty" xml:"CacheDuration,omitempty"`
	// The DeviceName of the IPC device. You can view the **DeviceName*	- on the **Devices*	- page in the IoT Platform console.
	//
	// 	Notice:  If you specify this parameter, you must also specify **ProductKey**.
	//
	// example:
	//
	// camera1
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	EnableStun *bool   `json:"EnableStun,omitempty" xml:"EnableStun,omitempty"`
	// The encryption type. Valid values:
	//
	// - **1*	- (default): Keyframe encryption.
	//
	// - **0**: No encryption.
	//
	// example:
	//
	// 1
	EncryptType *int32 `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	// Specifies whether to send a force I-frame command to the device. Valid values:
	//
	// - **true**: Sent.
	//
	// - **false*	- (default): Not sent.
	//
	// example:
	//
	// false
	ForceIFrame *bool `json:"ForceIFrame,omitempty" xml:"ForceIFrame,omitempty"`
	// The IotId of the IPC device. The IotId is the globally unique identifier that is issued by IoT Platform to the device. You can call the [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) operation of IoT Platform to query the IotId.
	//
	// 	Notice:  If you specify this parameter, you do not need to specify **ProductKey*	- or **DeviceName**. The **IotId*	- is the globally unique identifier of the device and corresponds to the combination of **ProductKey*	- and **DeviceName**. If you specify both **IotId*	- and the combination of **ProductKey*	- and **DeviceName**, **IotId*	- takes precedence.
	//
	// example:
	//
	// C47T6xwp6ms4bNlkHRWCg4****
	IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// The instance ID. You can view the **ID*	- of the current instance on the **Instance Overview*	- page in the IoT Platform console.
	//
	// 	Notice:
	//
	// - If an ID value is displayed, you must specify this parameter. Otherwise, the call fails.
	//
	// - If the <props="intl"><ph>**Instance Overview*	- page or </ph>ID value is not displayed, you do not need to specify this parameter.
	//
	//
	//
	// For more information about instances, see [Instance overview](https://help.aliyun.com/document_detail/356505.html).
	//
	// example:
	//
	// iot-cn-n6w1y59****
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	// Specifies whether the playback URL can be used for an unlimited number of times within the validity period. Valid values:
	//
	// - **true**: Allowed.
	//
	// - **false*	- (default): Not allowed.
	//
	// example:
	//
	// false
	PlayUnLimited *bool `json:"PlayUnLimited,omitempty" xml:"PlayUnLimited,omitempty"`
	// The ProductKey of the product to which the IPC device belongs. You can view the ProductKey on the **Products*	- page in the IoT Platform console or call the [QueryProductList](https://help.aliyun.com/document_detail/69271.html) operation to query the information about all products under the current account.
	//
	// 	Notice:  If you specify this parameter, you must also specify **DeviceName**.
	//
	// example:
	//
	// a1BwAGV****
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	// The playback protocol type. Valid values:
	//
	// - **rtmp*	- (default): RTMP.
	//
	// - **hls**: HLS.
	//
	// - **flv**: HTTP-FLV.
	//
	// example:
	//
	// rtmp
	Scheme *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
	// Specifies whether to encrypt the stream. Valid values:
	//
	// - **true**: Encrypted.
	//
	// - **false*	- (default): Not encrypted.
	//
	// example:
	//
	// true
	ShouldEncrypt *bool `json:"ShouldEncrypt,omitempty" xml:"ShouldEncrypt,omitempty"`
	// The stream type. Valid values:
	//
	// - **0*	- (default): Main stream.
	//
	// - **1**: Sub stream.
	//
	// example:
	//
	// 0
	StreamType *int32 `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	// The validity period of the URL. Valid values: 10 to 14400. Unit: seconds. Default value: 10.
	//
	// example:
	//
	// 60
	UrlValidDuration *int32 `json:"UrlValidDuration,omitempty" xml:"UrlValidDuration,omitempty"`
}

func (s QueryLiveStreamingRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryLiveStreamingRequest) GoString() string {
	return s.String()
}

func (s *QueryLiveStreamingRequest) GetCacheDuration() *int32 {
	return s.CacheDuration
}

func (s *QueryLiveStreamingRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *QueryLiveStreamingRequest) GetEnableStun() *bool {
	return s.EnableStun
}

func (s *QueryLiveStreamingRequest) GetEncryptType() *int32 {
	return s.EncryptType
}

func (s *QueryLiveStreamingRequest) GetForceIFrame() *bool {
	return s.ForceIFrame
}

func (s *QueryLiveStreamingRequest) GetIotId() *string {
	return s.IotId
}

func (s *QueryLiveStreamingRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *QueryLiveStreamingRequest) GetPlayUnLimited() *bool {
	return s.PlayUnLimited
}

func (s *QueryLiveStreamingRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *QueryLiveStreamingRequest) GetScheme() *string {
	return s.Scheme
}

func (s *QueryLiveStreamingRequest) GetShouldEncrypt() *bool {
	return s.ShouldEncrypt
}

func (s *QueryLiveStreamingRequest) GetStreamType() *int32 {
	return s.StreamType
}

func (s *QueryLiveStreamingRequest) GetUrlValidDuration() *int32 {
	return s.UrlValidDuration
}

func (s *QueryLiveStreamingRequest) SetCacheDuration(v int32) *QueryLiveStreamingRequest {
	s.CacheDuration = &v
	return s
}

func (s *QueryLiveStreamingRequest) SetDeviceName(v string) *QueryLiveStreamingRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryLiveStreamingRequest) SetEnableStun(v bool) *QueryLiveStreamingRequest {
	s.EnableStun = &v
	return s
}

func (s *QueryLiveStreamingRequest) SetEncryptType(v int32) *QueryLiveStreamingRequest {
	s.EncryptType = &v
	return s
}

func (s *QueryLiveStreamingRequest) SetForceIFrame(v bool) *QueryLiveStreamingRequest {
	s.ForceIFrame = &v
	return s
}

func (s *QueryLiveStreamingRequest) SetIotId(v string) *QueryLiveStreamingRequest {
	s.IotId = &v
	return s
}

func (s *QueryLiveStreamingRequest) SetIotInstanceId(v string) *QueryLiveStreamingRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryLiveStreamingRequest) SetPlayUnLimited(v bool) *QueryLiveStreamingRequest {
	s.PlayUnLimited = &v
	return s
}

func (s *QueryLiveStreamingRequest) SetProductKey(v string) *QueryLiveStreamingRequest {
	s.ProductKey = &v
	return s
}

func (s *QueryLiveStreamingRequest) SetScheme(v string) *QueryLiveStreamingRequest {
	s.Scheme = &v
	return s
}

func (s *QueryLiveStreamingRequest) SetShouldEncrypt(v bool) *QueryLiveStreamingRequest {
	s.ShouldEncrypt = &v
	return s
}

func (s *QueryLiveStreamingRequest) SetStreamType(v int32) *QueryLiveStreamingRequest {
	s.StreamType = &v
	return s
}

func (s *QueryLiveStreamingRequest) SetUrlValidDuration(v int32) *QueryLiveStreamingRequest {
	s.UrlValidDuration = &v
	return s
}

func (s *QueryLiveStreamingRequest) Validate() error {
	return dara.Validate(s)
}
