// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLocalFileUploadJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIotInstanceId(v string) *CreateLocalFileUploadJobRequest
	GetIotInstanceId() *string
	SetTimeSlot(v []*CreateLocalFileUploadJobRequestTimeSlot) *CreateLocalFileUploadJobRequest
	GetTimeSlot() []*CreateLocalFileUploadJobRequestTimeSlot
}

type CreateLocalFileUploadJobRequest struct {
	// The instance ID. You can view the **ID*	- of the current instance on the **Instance Overview*	- page in the IoT Platform console.
	//
	// 	Notice:
	//
	// - If the instance has an ID, you must specify the ID. Otherwise, the call fails.
	//
	// - If the instance does not have an <props="intl"><ph>**Instance Overview*	- page or </ph>ID, you do not need to specify this parameter.
	//
	//
	//
	// For more information about instances, see [Instance overview](https://help.aliyun.com/document_detail/356505.html).
	//
	// example:
	//
	// iot-cn-n6w1y59****
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	// The list of time slots.
	TimeSlot []*CreateLocalFileUploadJobRequestTimeSlot `json:"TimeSlot,omitempty" xml:"TimeSlot,omitempty" type:"Repeated"`
}

func (s CreateLocalFileUploadJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLocalFileUploadJobRequest) GoString() string {
	return s.String()
}

func (s *CreateLocalFileUploadJobRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *CreateLocalFileUploadJobRequest) GetTimeSlot() []*CreateLocalFileUploadJobRequestTimeSlot {
	return s.TimeSlot
}

func (s *CreateLocalFileUploadJobRequest) SetIotInstanceId(v string) *CreateLocalFileUploadJobRequest {
	s.IotInstanceId = &v
	return s
}

func (s *CreateLocalFileUploadJobRequest) SetTimeSlot(v []*CreateLocalFileUploadJobRequestTimeSlot) *CreateLocalFileUploadJobRequest {
	s.TimeSlot = v
	return s
}

func (s *CreateLocalFileUploadJobRequest) Validate() error {
	if s.TimeSlot != nil {
		for _, item := range s.TimeSlot {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateLocalFileUploadJobRequestTimeSlot struct {
	// The name of the video device.
	//
	// You can view the **DeviceName*	- on the **Devices*	- page in the IoT Platform console.
	//
	// 	Notice:  If you specify this parameter, you must also specify **ProductKey**.
	//
	// example:
	//
	// camera1
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// The end time for uploading the local video recording. The time is a 10-digit UNIX timestamp in seconds.
	//
	// >The maximum time range for a single local video recording upload job is 1 hour.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1631019775
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the video device. The ID is issued by IoT Platform and uniquely identifies the device. You can call the [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) operation of IoT Platform to query the ID.
	//
	// 	Notice:  If you specify this parameter, you do not need to specify **ProductKey*	- or **DeviceName**. The **IotId*	- uniquely identifies the device and has a one-to-one mapping with the combination of **ProductKey*	- and **DeviceName**. If you specify both **IotId*	- and the combination of **ProductKey*	- and **DeviceName**, **IotId*	- takes precedence.
	//
	// example:
	//
	// d7XmBoJhAr88C6PelXDF00****
	IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// The ProductKey of the product to which the video device belongs.
	//
	// You can view the ProductKey on the **Products*	- page in the IoT Platform console or call the [QueryProductList](https://help.aliyun.com/document_detail/69271.html) operation to query the information about all products under the current account.
	//
	// 	Notice:  If you specify this parameter, you must also specify **DeviceName**.
	//
	// example:
	//
	// a1BwAGV****
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	// The start time for uploading the local video recording. The time is a 10-digit UNIX timestamp in seconds.
	//
	// >The maximum time range for a single local video recording upload job is 1 hour.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1631019675
	StartTime *int32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreateLocalFileUploadJobRequestTimeSlot) String() string {
	return dara.Prettify(s)
}

func (s CreateLocalFileUploadJobRequestTimeSlot) GoString() string {
	return s.String()
}

func (s *CreateLocalFileUploadJobRequestTimeSlot) GetDeviceName() *string {
	return s.DeviceName
}

func (s *CreateLocalFileUploadJobRequestTimeSlot) GetEndTime() *int32 {
	return s.EndTime
}

func (s *CreateLocalFileUploadJobRequestTimeSlot) GetIotId() *string {
	return s.IotId
}

func (s *CreateLocalFileUploadJobRequestTimeSlot) GetProductKey() *string {
	return s.ProductKey
}

func (s *CreateLocalFileUploadJobRequestTimeSlot) GetStartTime() *int32 {
	return s.StartTime
}

func (s *CreateLocalFileUploadJobRequestTimeSlot) SetDeviceName(v string) *CreateLocalFileUploadJobRequestTimeSlot {
	s.DeviceName = &v
	return s
}

func (s *CreateLocalFileUploadJobRequestTimeSlot) SetEndTime(v int32) *CreateLocalFileUploadJobRequestTimeSlot {
	s.EndTime = &v
	return s
}

func (s *CreateLocalFileUploadJobRequestTimeSlot) SetIotId(v string) *CreateLocalFileUploadJobRequestTimeSlot {
	s.IotId = &v
	return s
}

func (s *CreateLocalFileUploadJobRequestTimeSlot) SetProductKey(v string) *CreateLocalFileUploadJobRequestTimeSlot {
	s.ProductKey = &v
	return s
}

func (s *CreateLocalFileUploadJobRequestTimeSlot) SetStartTime(v int32) *CreateLocalFileUploadJobRequestTimeSlot {
	s.StartTime = &v
	return s
}

func (s *CreateLocalFileUploadJobRequestTimeSlot) Validate() error {
	return dara.Validate(s)
}
