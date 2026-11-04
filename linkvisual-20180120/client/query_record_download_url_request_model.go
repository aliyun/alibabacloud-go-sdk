// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecordDownloadUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceName(v string) *QueryRecordDownloadUrlRequest
	GetDeviceName() *string
	SetFileName(v string) *QueryRecordDownloadUrlRequest
	GetFileName() *string
	SetIotId(v string) *QueryRecordDownloadUrlRequest
	GetIotId() *string
	SetIotInstanceId(v string) *QueryRecordDownloadUrlRequest
	GetIotInstanceId() *string
	SetProductKey(v string) *QueryRecordDownloadUrlRequest
	GetProductKey() *string
}

type QueryRecordDownloadUrlRequest struct {
	// The device name of the IPC device.
	//
	// You can view the **DeviceName*	- on the **Devices*	- page in the IoT Platform console.
	//
	// 	Notice:  If you specify this parameter, you must also specify **ProductKey**.
	//
	// example:
	//
	// camera1
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// The name of the recording file. You can call [QueryRecord](https://help.aliyun.com/document_detail/107922.html) to query the file name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1623999491_LYPsmZGeEN8GdKzMtYNiBXXXXXXXXX_0/c0064411-2065-400a-a13e-********
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The device ID of the IPC device. This is the unique identifier issued by IoT Platform for the device. You can call [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) of IoT Platform to query the device ID.
	//
	// 	Notice:  If you specify this parameter, you do not need to specify **ProductKey*	- or **DeviceName**. **IotId*	- is the globally unique identifier of a device and corresponds to a combination of **ProductKey*	- and **DeviceName*	- on a one-to-one basis. If you specify both **IotId*	- and the combination of **ProductKey*	- and **DeviceName**, **IotId*	- takes precedence.
	//
	// example:
	//
	// C47T6xwp6ms4bNlkHRWR****
	IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// The instance ID.
	//
	// <props="china">You can view the **ID*	- of the current instance on the **Overview*	- page in the IoT Platform console. If the ID value is displayed, you must specify this parameter. Otherwise, the call fails.
	//
	// <props="china">
	//
	// > If no ID value is displayed for the public instance, refer to [How to obtain an instance ID](https://help.aliyun.com/document_detail/267533.html).
	//
	// <props="intl">You do not need to specify this parameter.
	//
	// example:
	//
	// iot-cn-n6w1y59****
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	// The ProductKey of the product to which the IPC device belongs.
	//
	// You can view the ProductKey on the **Products*	- page in the IoT Platform console or call [QueryProductList](https://help.aliyun.com/document_detail/69271.html) to view the information about all products under the current account.
	//
	// 	Notice:  If you specify this parameter, you must also specify **DeviceName**.
	//
	// example:
	//
	// a1BwAGV****
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s QueryRecordDownloadUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRecordDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *QueryRecordDownloadUrlRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *QueryRecordDownloadUrlRequest) GetFileName() *string {
	return s.FileName
}

func (s *QueryRecordDownloadUrlRequest) GetIotId() *string {
	return s.IotId
}

func (s *QueryRecordDownloadUrlRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *QueryRecordDownloadUrlRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *QueryRecordDownloadUrlRequest) SetDeviceName(v string) *QueryRecordDownloadUrlRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryRecordDownloadUrlRequest) SetFileName(v string) *QueryRecordDownloadUrlRequest {
	s.FileName = &v
	return s
}

func (s *QueryRecordDownloadUrlRequest) SetIotId(v string) *QueryRecordDownloadUrlRequest {
	s.IotId = &v
	return s
}

func (s *QueryRecordDownloadUrlRequest) SetIotInstanceId(v string) *QueryRecordDownloadUrlRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryRecordDownloadUrlRequest) SetProductKey(v string) *QueryRecordDownloadUrlRequest {
	s.ProductKey = &v
	return s
}

func (s *QueryRecordDownloadUrlRequest) Validate() error {
	return dara.Validate(s)
}
