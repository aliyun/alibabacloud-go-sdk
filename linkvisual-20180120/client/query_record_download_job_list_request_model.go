// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecordDownloadJobListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *QueryRecordDownloadJobListRequest
	GetCurrentPage() *int32
	SetDeviceName(v string) *QueryRecordDownloadJobListRequest
	GetDeviceName() *string
	SetIotId(v string) *QueryRecordDownloadJobListRequest
	GetIotId() *string
	SetIotInstanceId(v string) *QueryRecordDownloadJobListRequest
	GetIotInstanceId() *string
	SetPageSize(v int32) *QueryRecordDownloadJobListRequest
	GetPageSize() *int32
	SetProductKey(v string) *QueryRecordDownloadJobListRequest
	GetProductKey() *string
}

type QueryRecordDownloadJobListRequest struct {
	// Specifies the page number to return. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The DeviceName of the video device. You can view the **DeviceName*	- of devices under a specified product on the **Devices*	- page in the [IoT Platform console](https://iot.console.aliyun.com/devices) or by calling the [QueryDevice](https://help.aliyun.com/document_detail/69905.html) operation.
	//
	// 	Notice:  If you specify this parameter, you must also specify **ProductKey**.
	//
	// example:
	//
	// Device01
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// The ID of the video device. You can call the [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) operation to query the ID.
	//
	// 	Notice:  **IotId*	- is the unique identifier of the device and corresponds to the combination of **ProductKey*	- and **DeviceName**. If you specify this parameter, you do not need to specify **ProductKey*	- or **DeviceName**. If you specify both **IotId*	- and the combination of **ProductKey*	- and **DeviceName**, **IotId*	- takes precedence.
	//
	// example:
	//
	// C47T6xwp6ms4bNlkHRWCg4****
	IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// The instance ID. You can view the ID of the current instance on the **Instance Overview*	- tab in the IoT Platform console.
	//
	// This parameter is required.
	//
	// example:
	//
	// iot-2w******
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	// Specifies the number of entries to return per page. Maximum value: 50.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The **ProductKey*	- of the product to which the video device belongs.
	//
	// You can view the information of all products under the current instance on the **Products*	- page in the [IoT Platform console](https://iot.console.aliyun.com/product) or by calling the [QueryProductList](https://help.aliyun.com/document_detail/69271.html) operation.
	//
	// 	Notice:  If you specify this parameter, you must also specify **DeviceName**.
	//
	// example:
	//
	// a1Bw******
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s QueryRecordDownloadJobListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRecordDownloadJobListRequest) GoString() string {
	return s.String()
}

func (s *QueryRecordDownloadJobListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryRecordDownloadJobListRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *QueryRecordDownloadJobListRequest) GetIotId() *string {
	return s.IotId
}

func (s *QueryRecordDownloadJobListRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *QueryRecordDownloadJobListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryRecordDownloadJobListRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *QueryRecordDownloadJobListRequest) SetCurrentPage(v int32) *QueryRecordDownloadJobListRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryRecordDownloadJobListRequest) SetDeviceName(v string) *QueryRecordDownloadJobListRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryRecordDownloadJobListRequest) SetIotId(v string) *QueryRecordDownloadJobListRequest {
	s.IotId = &v
	return s
}

func (s *QueryRecordDownloadJobListRequest) SetIotInstanceId(v string) *QueryRecordDownloadJobListRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryRecordDownloadJobListRequest) SetPageSize(v int32) *QueryRecordDownloadJobListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryRecordDownloadJobListRequest) SetProductKey(v string) *QueryRecordDownloadJobListRequest {
	s.ProductKey = &v
	return s
}

func (s *QueryRecordDownloadJobListRequest) Validate() error {
	return dara.Validate(s)
}
