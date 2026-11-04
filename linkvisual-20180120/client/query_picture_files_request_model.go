// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPictureFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int64) *QueryPictureFilesRequest
	GetBeginTime() *int64
	SetCurrentPage(v int32) *QueryPictureFilesRequest
	GetCurrentPage() *int32
	SetDeviceName(v string) *QueryPictureFilesRequest
	GetDeviceName() *string
	SetEndTime(v int64) *QueryPictureFilesRequest
	GetEndTime() *int64
	SetIotId(v string) *QueryPictureFilesRequest
	GetIotId() *string
	SetIotInstanceId(v string) *QueryPictureFilesRequest
	GetIotInstanceId() *string
	SetPageSize(v int32) *QueryPictureFilesRequest
	GetPageSize() *int32
	SetPictureSource(v int32) *QueryPictureFilesRequest
	GetPictureSource() *int32
	SetPictureType(v int32) *QueryPictureFilesRequest
	GetPictureType() *int32
	SetProductKey(v string) *QueryPictureFilesRequest
	GetProductKey() *string
}

type QueryPictureFilesRequest struct {
	// The start time of the retrieval range. The time is a 13-digit timestamp in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1615357669000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The page number of the returned page. Pages start from 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The DeviceName of the IP camera.
	//
	// You can view the **DeviceName*	- on the **Devices*	- page of the IoT Platform console.
	//
	// 	Notice:  If you specify this parameter, you must also specify **ProductKey**.
	//
	// example:
	//
	// camera1
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// The end time of the retrieval range. The time is a 13-digit timestamp in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1615444063000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the IP camera. The ID is issued by IoT Platform and is the unique identifier of the device. You can call the [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) operation of IoT Platform to query the ID.
	//
	// 	Notice:  If you specify this parameter, you do not need to specify **ProductKey*	- or **DeviceName**. The **IotId*	- is the unique identifier of the device and corresponds to the combination of **ProductKey*	- and **DeviceName**. If you specify both **IotId*	- and the combination of **ProductKey*	- and **DeviceName**, **IotId*	- takes precedence.
	//
	// example:
	//
	// zLZyi6aOLyOSHa9hsPyD00****
	IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// Instance ID of the instance. You can view the **ID*	- of the current instance on the **Instance Overview*	- page of the IoT Platform console.
	//
	// 	Notice:
	//
	// - If the instance has an ID, you must specify instance ID. Otherwise, the call fails.
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
	// The number of entries to return on each page. Default value: 10. Maximum value: 50.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The source of the image. Valid values:
	//
	// - **0**: all events. This is the default value.
	//
	// - **1**: snapshot captured during an alert event.
	//
	// - **2**: manual snapshot.
	//
	// - **3**: other sources.
	//
	// - **4**: cloud snapshot.
	//
	// example:
	//
	// 0
	PictureSource *int32 `json:"PictureSource,omitempty" xml:"PictureSource,omitempty"`
	// The type of the image. Valid values:
	//
	// - **0**: thumbnail.
	//
	// - **1**: original image.
	//
	// - **2**: all images. This is the default value.
	//
	// example:
	//
	// 2
	PictureType *int32 `json:"PictureType,omitempty" xml:"PictureType,omitempty"`
	// The ProductKey of the product to which the IP camera belongs.
	//
	// You can view the ProductKey on the **Products*	- page of the IoT Platform console or call the [QueryProductList](https://help.aliyun.com/document_detail/69271.html) operation to query the information about all products under the current account.
	//
	// 	Notice:  If you specify this parameter, you must also specify **DeviceName**.
	//
	// example:
	//
	// a1BwAGV****
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s QueryPictureFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPictureFilesRequest) GoString() string {
	return s.String()
}

func (s *QueryPictureFilesRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *QueryPictureFilesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryPictureFilesRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *QueryPictureFilesRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryPictureFilesRequest) GetIotId() *string {
	return s.IotId
}

func (s *QueryPictureFilesRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *QueryPictureFilesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryPictureFilesRequest) GetPictureSource() *int32 {
	return s.PictureSource
}

func (s *QueryPictureFilesRequest) GetPictureType() *int32 {
	return s.PictureType
}

func (s *QueryPictureFilesRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *QueryPictureFilesRequest) SetBeginTime(v int64) *QueryPictureFilesRequest {
	s.BeginTime = &v
	return s
}

func (s *QueryPictureFilesRequest) SetCurrentPage(v int32) *QueryPictureFilesRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryPictureFilesRequest) SetDeviceName(v string) *QueryPictureFilesRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryPictureFilesRequest) SetEndTime(v int64) *QueryPictureFilesRequest {
	s.EndTime = &v
	return s
}

func (s *QueryPictureFilesRequest) SetIotId(v string) *QueryPictureFilesRequest {
	s.IotId = &v
	return s
}

func (s *QueryPictureFilesRequest) SetIotInstanceId(v string) *QueryPictureFilesRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryPictureFilesRequest) SetPageSize(v int32) *QueryPictureFilesRequest {
	s.PageSize = &v
	return s
}

func (s *QueryPictureFilesRequest) SetPictureSource(v int32) *QueryPictureFilesRequest {
	s.PictureSource = &v
	return s
}

func (s *QueryPictureFilesRequest) SetPictureType(v int32) *QueryPictureFilesRequest {
	s.PictureType = &v
	return s
}

func (s *QueryPictureFilesRequest) SetProductKey(v string) *QueryPictureFilesRequest {
	s.ProductKey = &v
	return s
}

func (s *QueryPictureFilesRequest) Validate() error {
	return dara.Validate(s)
}
