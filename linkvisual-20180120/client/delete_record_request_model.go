// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceName(v string) *DeleteRecordRequest
	GetDeviceName() *string
	SetFileNameList(v []*string) *DeleteRecordRequest
	GetFileNameList() []*string
	SetIotId(v string) *DeleteRecordRequest
	GetIotId() *string
	SetIotInstanceId(v string) *DeleteRecordRequest
	GetIotInstanceId() *string
	SetProductKey(v string) *DeleteRecordRequest
	GetProductKey() *string
}

type DeleteRecordRequest struct {
	// The DeviceName of the IPC device. You can view the **DeviceName*	- on the **Devices*	- page in the IoT Platform console.
	//
	// 	Notice:  If you specify this parameter, you must also specify **ProductKey**.
	//
	// example:
	//
	// camera1
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// The list of names of the video files to be deleted. You can call the [QueryRecord](https://help.aliyun.com/document_detail/107922.html) operation to obtain the video file names.
	//
	// > You can delete a maximum of 50 video files in a single call.
	//
	// This parameter is required.
	//
	// example:
	//
	// file1
	FileNameList []*string `json:"FileNameList,omitempty" xml:"FileNameList,omitempty" type:"Repeated"`
	// The ID of the IPC device. The ID is issued by IoT Platform and uniquely identifies the device. You can call the [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) operation of IoT Platform to query the ID.
	//
	// 	Notice:  If you specify this parameter, you do not need to specify **ProductKey*	- or **DeviceName**. As the unique identifier of the device, **IotId*	- has a one-to-one mapping with the combination of **ProductKey*	- and **DeviceName**. If you specify both **IotId*	- and the combination of **ProductKey*	- and **DeviceName**, **IotId*	- takes precedence.
	//
	// example:
	//
	// d7XmBoJhAr88C6PelXDF00****
	IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// The instance ID. You can view the **ID*	- of the current instance on the **Instance Overview*	- page in the IoT Platform console.
	//
	// 	Notice:
	//
	// - If an ID value is available, you must specify this parameter. Otherwise, the call fails.
	//
	// - If no <props="intl"><ph>**Instance Overview*	- page or </ph>ID value is available, you do not need to specify this parameter.
	//
	//
	//
	// For more information about instances, see [Instance overview](https://help.aliyun.com/document_detail/356505.html).
	//
	// example:
	//
	// iot-060a****
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	// The ProductKey of the product to which the IPC device belongs. You can view the ProductKey on the **Products*	- page in the IoT Platform console or call the [QueryProductList](https://help.aliyun.com/document_detail/69271.html) operation to query the information about all products under the current account.
	//
	// 	Notice:  If you specify this parameter, you must also specify **DeviceName**.
	//
	// example:
	//
	// a1BwAGV****
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s DeleteRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRecordRequest) GoString() string {
	return s.String()
}

func (s *DeleteRecordRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *DeleteRecordRequest) GetFileNameList() []*string {
	return s.FileNameList
}

func (s *DeleteRecordRequest) GetIotId() *string {
	return s.IotId
}

func (s *DeleteRecordRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *DeleteRecordRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *DeleteRecordRequest) SetDeviceName(v string) *DeleteRecordRequest {
	s.DeviceName = &v
	return s
}

func (s *DeleteRecordRequest) SetFileNameList(v []*string) *DeleteRecordRequest {
	s.FileNameList = v
	return s
}

func (s *DeleteRecordRequest) SetIotId(v string) *DeleteRecordRequest {
	s.IotId = &v
	return s
}

func (s *DeleteRecordRequest) SetIotInstanceId(v string) *DeleteRecordRequest {
	s.IotInstanceId = &v
	return s
}

func (s *DeleteRecordRequest) SetProductKey(v string) *DeleteRecordRequest {
	s.ProductKey = &v
	return s
}

func (s *DeleteRecordRequest) Validate() error {
	return dara.Validate(s)
}
