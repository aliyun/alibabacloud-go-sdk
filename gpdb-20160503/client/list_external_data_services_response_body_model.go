// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExternalDataServicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListExternalDataServicesResponseBody
	GetPageNumber() *int32
	SetRequestId(v string) *ListExternalDataServicesResponseBody
	GetRequestId() *string
	SetServiceItems(v []*ListExternalDataServicesResponseBodyServiceItems) *ListExternalDataServicesResponseBody
	GetServiceItems() []*ListExternalDataServicesResponseBodyServiceItems
	SetTotalRecordCount(v int32) *ListExternalDataServicesResponseBody
	GetTotalRecordCount() *int32
}

type ListExternalDataServicesResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried services.
	ServiceItems []*ListExternalDataServicesResponseBodyServiceItems `json:"ServiceItems,omitempty" xml:"ServiceItems,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s ListExternalDataServicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListExternalDataServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListExternalDataServicesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListExternalDataServicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListExternalDataServicesResponseBody) GetServiceItems() []*ListExternalDataServicesResponseBodyServiceItems {
	return s.ServiceItems
}

func (s *ListExternalDataServicesResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *ListExternalDataServicesResponseBody) SetPageNumber(v int32) *ListExternalDataServicesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListExternalDataServicesResponseBody) SetRequestId(v string) *ListExternalDataServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExternalDataServicesResponseBody) SetServiceItems(v []*ListExternalDataServicesResponseBodyServiceItems) *ListExternalDataServicesResponseBody {
	s.ServiceItems = v
	return s
}

func (s *ListExternalDataServicesResponseBody) SetTotalRecordCount(v int32) *ListExternalDataServicesResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *ListExternalDataServicesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListExternalDataServicesResponseBodyServiceItems struct {
	// The time when the service was created.
	//
	// example:
	//
	// 2019-09-08T16:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the service was last modified.
	//
	// example:
	//
	// 2019-09-08T17:00:00Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The description of the service.
	//
	// example:
	//
	// test-adbpgss
	ServiceDescription *string `json:"ServiceDescription,omitempty" xml:"ServiceDescription,omitempty"`
	// The service ID.
	//
	// example:
	//
	// 1
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// test-adbpgss
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The specifications of the service. Unit: compute units (CUs).
	//
	// example:
	//
	// 2
	ServiceSpec *string `json:"ServiceSpec,omitempty" xml:"ServiceSpec,omitempty"`
	// The type of the service. Valid values:
	//
	// 	- pxf
	//
	// example:
	//
	// pxf
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The status of the service. Valid values:
	//
	// 	- Init
	//
	// 	- Running
	//
	// 	- Exception
	//
	// 	- Paused
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListExternalDataServicesResponseBodyServiceItems) String() string {
	return dara.Prettify(s)
}

func (s ListExternalDataServicesResponseBodyServiceItems) GoString() string {
	return s.String()
}

func (s *ListExternalDataServicesResponseBodyServiceItems) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListExternalDataServicesResponseBodyServiceItems) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListExternalDataServicesResponseBodyServiceItems) GetServiceDescription() *string {
	return s.ServiceDescription
}

func (s *ListExternalDataServicesResponseBodyServiceItems) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListExternalDataServicesResponseBodyServiceItems) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListExternalDataServicesResponseBodyServiceItems) GetServiceSpec() *string {
	return s.ServiceSpec
}

func (s *ListExternalDataServicesResponseBodyServiceItems) GetServiceType() *string {
	return s.ServiceType
}

func (s *ListExternalDataServicesResponseBodyServiceItems) GetStatus() *string {
	return s.Status
}

func (s *ListExternalDataServicesResponseBodyServiceItems) SetCreateTime(v string) *ListExternalDataServicesResponseBodyServiceItems {
	s.CreateTime = &v
	return s
}

func (s *ListExternalDataServicesResponseBodyServiceItems) SetModifyTime(v string) *ListExternalDataServicesResponseBodyServiceItems {
	s.ModifyTime = &v
	return s
}

func (s *ListExternalDataServicesResponseBodyServiceItems) SetServiceDescription(v string) *ListExternalDataServicesResponseBodyServiceItems {
	s.ServiceDescription = &v
	return s
}

func (s *ListExternalDataServicesResponseBodyServiceItems) SetServiceId(v string) *ListExternalDataServicesResponseBodyServiceItems {
	s.ServiceId = &v
	return s
}

func (s *ListExternalDataServicesResponseBodyServiceItems) SetServiceName(v string) *ListExternalDataServicesResponseBodyServiceItems {
	s.ServiceName = &v
	return s
}

func (s *ListExternalDataServicesResponseBodyServiceItems) SetServiceSpec(v string) *ListExternalDataServicesResponseBodyServiceItems {
	s.ServiceSpec = &v
	return s
}

func (s *ListExternalDataServicesResponseBodyServiceItems) SetServiceType(v string) *ListExternalDataServicesResponseBodyServiceItems {
	s.ServiceType = &v
	return s
}

func (s *ListExternalDataServicesResponseBodyServiceItems) SetStatus(v string) *ListExternalDataServicesResponseBodyServiceItems {
	s.Status = &v
	return s
}

func (s *ListExternalDataServicesResponseBodyServiceItems) Validate() error {
	return dara.Validate(s)
}
