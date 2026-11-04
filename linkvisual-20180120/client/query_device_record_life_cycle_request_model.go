// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDeviceRecordLifeCycleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceList(v []*string) *QueryDeviceRecordLifeCycleRequest
	GetDeviceList() []*string
}

type QueryDeviceRecordLifeCycleRequest struct {
	// The list of IoT IDs of the IPC devices. An IoT ID is a unique identifier issued by IoT Platform for a device. You can call the IoT Platform [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) operation to query the IoT ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["C47T6xwp6ms4bNlkHRWCg4****", "d7XmBoJhAr88C6PelXDF00****"]
	DeviceList []*string `json:"DeviceList,omitempty" xml:"DeviceList,omitempty" type:"Repeated"`
}

func (s QueryDeviceRecordLifeCycleRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDeviceRecordLifeCycleRequest) GoString() string {
	return s.String()
}

func (s *QueryDeviceRecordLifeCycleRequest) GetDeviceList() []*string {
	return s.DeviceList
}

func (s *QueryDeviceRecordLifeCycleRequest) SetDeviceList(v []*string) *QueryDeviceRecordLifeCycleRequest {
	s.DeviceList = v
	return s
}

func (s *QueryDeviceRecordLifeCycleRequest) Validate() error {
	return dara.Validate(s)
}
