// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOtaTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *ListOtaTaskRequest
	GetAppInstanceGroupId() *string
	SetOtaType(v string) *ListOtaTaskRequest
	GetOtaType() *string
	SetPageNumber(v int32) *ListOtaTaskRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListOtaTaskRequest
	GetPageSize() *int32
}

type ListOtaTaskRequest struct {
	// The ID of the delivery group.
	//
	// This parameter is required.
	//
	// example:
	//
	// aig-53fvrq1oanz6c****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The type of the OTA update task.
	//
	// Valid values:
	//
	// 	- Fota: update of the system components of Alibaba Cloud Workspace
	//
	// This parameter is required.
	//
	// example:
	//
	// Fota
	OtaType *string `json:"OtaType,omitempty" xml:"OtaType,omitempty"`
	// The page number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListOtaTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOtaTaskRequest) GoString() string {
	return s.String()
}

func (s *ListOtaTaskRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *ListOtaTaskRequest) GetOtaType() *string {
	return s.OtaType
}

func (s *ListOtaTaskRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListOtaTaskRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOtaTaskRequest) SetAppInstanceGroupId(v string) *ListOtaTaskRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ListOtaTaskRequest) SetOtaType(v string) *ListOtaTaskRequest {
	s.OtaType = &v
	return s
}

func (s *ListOtaTaskRequest) SetPageNumber(v int32) *ListOtaTaskRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOtaTaskRequest) SetPageSize(v int32) *ListOtaTaskRequest {
	s.PageSize = &v
	return s
}

func (s *ListOtaTaskRequest) Validate() error {
	return dara.Validate(s)
}
