// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int32) *QueryRecordRequest
	GetBeginTime() *int32
	SetCurrentPage(v int32) *QueryRecordRequest
	GetCurrentPage() *int32
	SetDeviceName(v string) *QueryRecordRequest
	GetDeviceName() *string
	SetEndTime(v int32) *QueryRecordRequest
	GetEndTime() *int32
	SetIotId(v string) *QueryRecordRequest
	GetIotId() *string
	SetIotInstanceId(v string) *QueryRecordRequest
	GetIotInstanceId() *string
	SetNeedSnapshot(v bool) *QueryRecordRequest
	GetNeedSnapshot() *bool
	SetPageSize(v int32) *QueryRecordRequest
	GetPageSize() *int32
	SetProductKey(v string) *QueryRecordRequest
	GetProductKey() *string
	SetRecordType(v int32) *QueryRecordRequest
	GetRecordType() *int32
	SetStreamType(v int32) *QueryRecordRequest
	GetStreamType() *int32
}

type QueryRecordRequest struct {
	// The start time of the query range. The time is a 10-digit timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1508200012
	BeginTime *int32 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The DeviceName of the IP camera. You can view the **DeviceName*	- on the **Devices*	- page in the IoT Platform console.
	//
	// 	Notice:  If you specify this parameter, you must also specify **ProductKey**.
	//
	// example:
	//
	// camera1
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// The end time of the query range. The time is a 10-digit timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1582420940
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the IP camera. The ID is issued by IoT Platform and is the unique identifier of the device. You can call the [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) operation of IoT Platform to query the ID.
	//
	// 	Notice:  If you specify this parameter, you do not need to specify **ProductKey*	- or **DeviceName**. The **IotId*	- is the unique identifier of the device and corresponds to the combination of **ProductKey*	- and **DeviceName**. If you specify both **IotId*	- and the combination of **ProductKey*	- and **DeviceName**, **IotId*	- takes precedence.
	//
	// example:
	//
	// C47T6xwp6ms4bNlkHRWCg4****
	IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// The instance ID. You can view the ID of the current instance on the **Overview*	- tab in the IoT Platform console.
	//
	// 	Notice:
	//
	// - If your instance has an ID, you must specify the **ID**. Otherwise, the call fails.
	//
	// - If your instance does not have an ID, you do not need to specify this parameter.
	//
	//
	//
	// For more information about instances, see [Instance overview](https://help.aliyun.com/document_detail/356505.html).
	//
	// example:
	//
	// iot-060a****
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	// Specifies whether to return thumbnails. Valid values:
	//
	// - **true**: yes.
	//
	// - **false*	- (default): no.
	//
	// example:
	//
	// true
	NeedSnapshot *bool `json:"NeedSnapshot,omitempty" xml:"NeedSnapshot,omitempty"`
	// The number of entries to return on each page. Default value: 10. Maximum value: 50.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ProductKey of the product to which the IP camera belongs. You can view the ProductKey on the **Products*	- page in the IoT Platform console or call the [QueryProductList](https://help.aliyun.com/document_detail/69271.html) operation to query the information of all products under the current account.
	//
	// 	Notice:  If you specify this parameter, you must also specify **DeviceName**.
	//
	// example:
	//
	// a1BwAGV****
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	// The recording type. Valid values:
	//
	// - **0**: continuous recording.
	//
	// - **1**: intelligent alert event recording.
	//
	// - **2*	- (default): manual recording.
	//
	// If this parameter is left empty, all recording types are returned.
	//
	// example:
	//
	// 1
	RecordType *int32 `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	// The stream type. Valid values:
	//
	// - **0*	- (default): main stream.
	//
	// - **1**: sub stream.
	//
	// example:
	//
	// 1
	StreamType *int32 `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
}

func (s QueryRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRecordRequest) GoString() string {
	return s.String()
}

func (s *QueryRecordRequest) GetBeginTime() *int32 {
	return s.BeginTime
}

func (s *QueryRecordRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryRecordRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *QueryRecordRequest) GetEndTime() *int32 {
	return s.EndTime
}

func (s *QueryRecordRequest) GetIotId() *string {
	return s.IotId
}

func (s *QueryRecordRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *QueryRecordRequest) GetNeedSnapshot() *bool {
	return s.NeedSnapshot
}

func (s *QueryRecordRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryRecordRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *QueryRecordRequest) GetRecordType() *int32 {
	return s.RecordType
}

func (s *QueryRecordRequest) GetStreamType() *int32 {
	return s.StreamType
}

func (s *QueryRecordRequest) SetBeginTime(v int32) *QueryRecordRequest {
	s.BeginTime = &v
	return s
}

func (s *QueryRecordRequest) SetCurrentPage(v int32) *QueryRecordRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryRecordRequest) SetDeviceName(v string) *QueryRecordRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryRecordRequest) SetEndTime(v int32) *QueryRecordRequest {
	s.EndTime = &v
	return s
}

func (s *QueryRecordRequest) SetIotId(v string) *QueryRecordRequest {
	s.IotId = &v
	return s
}

func (s *QueryRecordRequest) SetIotInstanceId(v string) *QueryRecordRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryRecordRequest) SetNeedSnapshot(v bool) *QueryRecordRequest {
	s.NeedSnapshot = &v
	return s
}

func (s *QueryRecordRequest) SetPageSize(v int32) *QueryRecordRequest {
	s.PageSize = &v
	return s
}

func (s *QueryRecordRequest) SetProductKey(v string) *QueryRecordRequest {
	s.ProductKey = &v
	return s
}

func (s *QueryRecordRequest) SetRecordType(v int32) *QueryRecordRequest {
	s.RecordType = &v
	return s
}

func (s *QueryRecordRequest) SetStreamType(v int32) *QueryRecordRequest {
	s.StreamType = &v
	return s
}

func (s *QueryRecordRequest) Validate() error {
	return dara.Validate(s)
}
