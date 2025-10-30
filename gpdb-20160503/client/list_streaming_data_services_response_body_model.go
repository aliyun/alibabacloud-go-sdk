// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStreamingDataServicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListStreamingDataServicesResponseBody
	GetRequestId() *string
	SetServiceItems(v []*ListStreamingDataServicesResponseBodyServiceItems) *ListStreamingDataServicesResponseBody
	GetServiceItems() []*ListStreamingDataServicesResponseBodyServiceItems
	SetTotalRecordCount(v int32) *ListStreamingDataServicesResponseBody
	GetTotalRecordCount() *int32
}

type ListStreamingDataServicesResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returns real-time data service items
	ServiceItems []*ListStreamingDataServicesResponseBodyServiceItems `json:"ServiceItems,omitempty" xml:"ServiceItems,omitempty" type:"Repeated"`
	// Total record count.
	//
	// example:
	//
	// 1
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s ListStreamingDataServicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListStreamingDataServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListStreamingDataServicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListStreamingDataServicesResponseBody) GetServiceItems() []*ListStreamingDataServicesResponseBodyServiceItems {
	return s.ServiceItems
}

func (s *ListStreamingDataServicesResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *ListStreamingDataServicesResponseBody) SetRequestId(v string) *ListStreamingDataServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStreamingDataServicesResponseBody) SetServiceItems(v []*ListStreamingDataServicesResponseBodyServiceItems) *ListStreamingDataServicesResponseBody {
	s.ServiceItems = v
	return s
}

func (s *ListStreamingDataServicesResponseBody) SetTotalRecordCount(v int32) *ListStreamingDataServicesResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *ListStreamingDataServicesResponseBody) Validate() error {
	if s.ServiceItems != nil {
		for _, item := range s.ServiceItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListStreamingDataServicesResponseBodyServiceItems struct {
	// Creation time.
	//
	// example:
	//
	// 2019-09-08T16:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Last modified time
	//
	// example:
	//
	// 2019-09-08T17:00:00Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// Service description.
	//
	// example:
	//
	// test-adbpgss
	ServiceDescription *string `json:"ServiceDescription,omitempty" xml:"ServiceDescription,omitempty"`
	// Service ID.
	//
	// example:
	//
	// 1
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// Service IP.
	//
	// example:
	//
	// 192.168.0.1
	ServiceIp *string `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	// Whether it is a managed service.
	//
	// example:
	//
	// true
	ServiceManaged *bool `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
	// Service name.
	//
	// example:
	//
	// test-adbpgss
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// Service owner ID.
	//
	// example:
	//
	// 123
	ServiceOwnerId *string `json:"ServiceOwnerId,omitempty" xml:"ServiceOwnerId,omitempty"`
	// Service port.
	//
	// example:
	//
	// 5432
	ServicePort *string `json:"ServicePort,omitempty" xml:"ServicePort,omitempty"`
	// Service specification (in CU).
	//
	// example:
	//
	// 8
	ServiceSpec *string `json:"ServiceSpec,omitempty" xml:"ServiceSpec,omitempty"`
	// Service type, with the following value:
	//
	// - **adbpgss**
	//
	// example:
	//
	// adbpgss
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// Service status, with the following values:
	//
	// - Init: Initializing
	//
	// - Running: In operation
	//
	// - Exception: Abnormal
	//
	// - Paused: Suspended
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListStreamingDataServicesResponseBodyServiceItems) String() string {
	return dara.Prettify(s)
}

func (s ListStreamingDataServicesResponseBodyServiceItems) GoString() string {
	return s.String()
}

func (s *ListStreamingDataServicesResponseBodyServiceItems) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListStreamingDataServicesResponseBodyServiceItems) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListStreamingDataServicesResponseBodyServiceItems) GetServiceDescription() *string {
	return s.ServiceDescription
}

func (s *ListStreamingDataServicesResponseBodyServiceItems) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListStreamingDataServicesResponseBodyServiceItems) GetServiceIp() *string {
	return s.ServiceIp
}

func (s *ListStreamingDataServicesResponseBodyServiceItems) GetServiceManaged() *bool {
	return s.ServiceManaged
}

func (s *ListStreamingDataServicesResponseBodyServiceItems) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListStreamingDataServicesResponseBodyServiceItems) GetServiceOwnerId() *string {
	return s.ServiceOwnerId
}

func (s *ListStreamingDataServicesResponseBodyServiceItems) GetServicePort() *string {
	return s.ServicePort
}

func (s *ListStreamingDataServicesResponseBodyServiceItems) GetServiceSpec() *string {
	return s.ServiceSpec
}

func (s *ListStreamingDataServicesResponseBodyServiceItems) GetServiceType() *string {
	return s.ServiceType
}

func (s *ListStreamingDataServicesResponseBodyServiceItems) GetStatus() *string {
	return s.Status
}

func (s *ListStreamingDataServicesResponseBodyServiceItems) SetCreateTime(v string) *ListStreamingDataServicesResponseBodyServiceItems {
	s.CreateTime = &v
	return s
}

func (s *ListStreamingDataServicesResponseBodyServiceItems) SetModifyTime(v string) *ListStreamingDataServicesResponseBodyServiceItems {
	s.ModifyTime = &v
	return s
}

func (s *ListStreamingDataServicesResponseBodyServiceItems) SetServiceDescription(v string) *ListStreamingDataServicesResponseBodyServiceItems {
	s.ServiceDescription = &v
	return s
}

func (s *ListStreamingDataServicesResponseBodyServiceItems) SetServiceId(v string) *ListStreamingDataServicesResponseBodyServiceItems {
	s.ServiceId = &v
	return s
}

func (s *ListStreamingDataServicesResponseBodyServiceItems) SetServiceIp(v string) *ListStreamingDataServicesResponseBodyServiceItems {
	s.ServiceIp = &v
	return s
}

func (s *ListStreamingDataServicesResponseBodyServiceItems) SetServiceManaged(v bool) *ListStreamingDataServicesResponseBodyServiceItems {
	s.ServiceManaged = &v
	return s
}

func (s *ListStreamingDataServicesResponseBodyServiceItems) SetServiceName(v string) *ListStreamingDataServicesResponseBodyServiceItems {
	s.ServiceName = &v
	return s
}

func (s *ListStreamingDataServicesResponseBodyServiceItems) SetServiceOwnerId(v string) *ListStreamingDataServicesResponseBodyServiceItems {
	s.ServiceOwnerId = &v
	return s
}

func (s *ListStreamingDataServicesResponseBodyServiceItems) SetServicePort(v string) *ListStreamingDataServicesResponseBodyServiceItems {
	s.ServicePort = &v
	return s
}

func (s *ListStreamingDataServicesResponseBodyServiceItems) SetServiceSpec(v string) *ListStreamingDataServicesResponseBodyServiceItems {
	s.ServiceSpec = &v
	return s
}

func (s *ListStreamingDataServicesResponseBodyServiceItems) SetServiceType(v string) *ListStreamingDataServicesResponseBodyServiceItems {
	s.ServiceType = &v
	return s
}

func (s *ListStreamingDataServicesResponseBodyServiceItems) SetStatus(v string) *ListStreamingDataServicesResponseBodyServiceItems {
	s.Status = &v
	return s
}

func (s *ListStreamingDataServicesResponseBodyServiceItems) Validate() error {
	return dara.Validate(s)
}
