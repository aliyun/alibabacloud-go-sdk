// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryStreamPushJobListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *QueryStreamPushJobListRequest
	GetCurrentPage() *int32
	SetDeviceName(v string) *QueryStreamPushJobListRequest
	GetDeviceName() *string
	SetIotId(v string) *QueryStreamPushJobListRequest
	GetIotId() *string
	SetIotInstanceId(v string) *QueryStreamPushJobListRequest
	GetIotInstanceId() *string
	SetJobType(v int32) *QueryStreamPushJobListRequest
	GetJobType() *int32
	SetPageSize(v int32) *QueryStreamPushJobListRequest
	GetPageSize() *int32
	SetProductKey(v string) *QueryStreamPushJobListRequest
	GetProductKey() *string
}

type QueryStreamPushJobListRequest struct {
	// Specifies the page number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The DeviceName of the device.
	//
	// You can view the **DeviceName*	- on the **Devices*	- page in the IoT Platform console.
	//
	// 	Notice:  If you specify this parameter, you must also specify **ProductKey**.
	//
	// example:
	//
	// camera1
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// The ID of the device. The ID is the unique identifier that is issued by IoT Platform to the device. You can call the [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) operation of IoT Platform to query the ID.
	//
	// 	Notice:  **IotId*	- is the unique identifier of the device and corresponds to the combination of **ProductKey*	- and **DeviceName**. If you specify this parameter, you do not need to specify **ProductKey*	- or **DeviceName**. If you specify both **IotId*	- and the combination of **ProductKey*	- and **DeviceName**, **IotId*	- takes precedence.
	//
	// example:
	//
	// d7XmBoJhAr88C6PelXDF00****
	IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// Instance ID of the instance. You can view instance ID of the current instance on the **Instance Overview*	- tab in the IoT Platform console.
	//
	// This parameter is required.
	//
	// example:
	//
	// iot-cn-n6w1y59****
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	// The type of the stream push job. Valid values:
	//
	// - **1*	- (default): continuous job.
	//
	// - **2**: triggered job. The job ends after it is interrupted.
	//
	// example:
	//
	// 1
	JobType *int32 `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// Specifies the number of entries to return on each page. Maximum value: 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The **ProductKey*	- of the product to which the device belongs.
	//
	// You can view the **ProductKey*	- on the **Products*	- page in the IoT Platform console or call the [QueryProductList](https://help.aliyun.com/document_detail/69271.html) operation to query the information of all products under the current account.
	//
	// 	Notice:  If you specify this parameter, you must also specify **DeviceName**.
	//
	// example:
	//
	// a1BwAGV****
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s QueryStreamPushJobListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryStreamPushJobListRequest) GoString() string {
	return s.String()
}

func (s *QueryStreamPushJobListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryStreamPushJobListRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *QueryStreamPushJobListRequest) GetIotId() *string {
	return s.IotId
}

func (s *QueryStreamPushJobListRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *QueryStreamPushJobListRequest) GetJobType() *int32 {
	return s.JobType
}

func (s *QueryStreamPushJobListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryStreamPushJobListRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *QueryStreamPushJobListRequest) SetCurrentPage(v int32) *QueryStreamPushJobListRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryStreamPushJobListRequest) SetDeviceName(v string) *QueryStreamPushJobListRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryStreamPushJobListRequest) SetIotId(v string) *QueryStreamPushJobListRequest {
	s.IotId = &v
	return s
}

func (s *QueryStreamPushJobListRequest) SetIotInstanceId(v string) *QueryStreamPushJobListRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryStreamPushJobListRequest) SetJobType(v int32) *QueryStreamPushJobListRequest {
	s.JobType = &v
	return s
}

func (s *QueryStreamPushJobListRequest) SetPageSize(v int32) *QueryStreamPushJobListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryStreamPushJobListRequest) SetProductKey(v string) *QueryStreamPushJobListRequest {
	s.ProductKey = &v
	return s
}

func (s *QueryStreamPushJobListRequest) Validate() error {
	return dara.Validate(s)
}
