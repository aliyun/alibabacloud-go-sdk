// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecordUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceName(v string) *QueryRecordUrlRequest
	GetDeviceName() *string
	SetFileName(v string) *QueryRecordUrlRequest
	GetFileName() *string
	SetIotId(v string) *QueryRecordUrlRequest
	GetIotId() *string
	SetIotInstanceId(v string) *QueryRecordUrlRequest
	GetIotInstanceId() *string
	SetProductKey(v string) *QueryRecordUrlRequest
	GetProductKey() *string
	SetSeekTime(v int32) *QueryRecordUrlRequest
	GetSeekTime() *int32
	SetUrlValidDuration(v int32) *QueryRecordUrlRequest
	GetUrlValidDuration() *int32
}

type QueryRecordUrlRequest struct {
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
	// The name of the cloud recording file. You can call the [QueryRecord](https://help.aliyun.com/document_detail/107922.html) operation to query the file name.
	//
	// This parameter is required.
	//
	// example:
	//
	// file1
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The IotId of the IP camera. The IotId is the unique identifier that is issued by IoT Platform to the device. You can call the IoT Platform [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) operation to query the IotId.
	//
	// 	Notice:  If you specify this parameter, you do not need to specify **ProductKey*	- or **DeviceName**. The **IotId*	- is the unique identifier of the device and corresponds to the combination of **ProductKey*	- and **DeviceName**. If you specify both **IotId*	- and the combination of **ProductKey*	- and **DeviceName**, **IotId*	- takes precedence.
	//
	// example:
	//
	// VrJwPGjC1sJlqPjZA3cxg4****
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
	// iot-cn-n6w1y****
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
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
	// The playback offset time of the recording, in seconds. Default value: 0.
	//
	// example:
	//
	// 0
	SeekTime *int32 `json:"SeekTime,omitempty" xml:"SeekTime,omitempty"`
	// The validity period of the URL, in seconds. Valid values: 10 to 3600.
	//
	// example:
	//
	// 10
	UrlValidDuration *int32 `json:"UrlValidDuration,omitempty" xml:"UrlValidDuration,omitempty"`
}

func (s QueryRecordUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRecordUrlRequest) GoString() string {
	return s.String()
}

func (s *QueryRecordUrlRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *QueryRecordUrlRequest) GetFileName() *string {
	return s.FileName
}

func (s *QueryRecordUrlRequest) GetIotId() *string {
	return s.IotId
}

func (s *QueryRecordUrlRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *QueryRecordUrlRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *QueryRecordUrlRequest) GetSeekTime() *int32 {
	return s.SeekTime
}

func (s *QueryRecordUrlRequest) GetUrlValidDuration() *int32 {
	return s.UrlValidDuration
}

func (s *QueryRecordUrlRequest) SetDeviceName(v string) *QueryRecordUrlRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryRecordUrlRequest) SetFileName(v string) *QueryRecordUrlRequest {
	s.FileName = &v
	return s
}

func (s *QueryRecordUrlRequest) SetIotId(v string) *QueryRecordUrlRequest {
	s.IotId = &v
	return s
}

func (s *QueryRecordUrlRequest) SetIotInstanceId(v string) *QueryRecordUrlRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryRecordUrlRequest) SetProductKey(v string) *QueryRecordUrlRequest {
	s.ProductKey = &v
	return s
}

func (s *QueryRecordUrlRequest) SetSeekTime(v int32) *QueryRecordUrlRequest {
	s.SeekTime = &v
	return s
}

func (s *QueryRecordUrlRequest) SetUrlValidDuration(v int32) *QueryRecordUrlRequest {
	s.UrlValidDuration = &v
	return s
}

func (s *QueryRecordUrlRequest) Validate() error {
	return dara.Validate(s)
}
