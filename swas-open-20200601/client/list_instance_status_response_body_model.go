// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceStatuses(v []*ListInstanceStatusResponseBodyInstanceStatuses) *ListInstanceStatusResponseBody
	GetInstanceStatuses() []*ListInstanceStatusResponseBodyInstanceStatuses
	SetPageNumber(v int32) *ListInstanceStatusResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListInstanceStatusResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListInstanceStatusResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListInstanceStatusResponseBody
	GetTotalCount() *int32
}

type ListInstanceStatusResponseBody struct {
	// The ID of the simple application server.
	InstanceStatuses []*ListInstanceStatusResponseBodyInstanceStatuses `json:"InstanceStatuses,omitempty" xml:"InstanceStatuses,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 54
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstanceStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceStatusResponseBody) GetInstanceStatuses() []*ListInstanceStatusResponseBodyInstanceStatuses {
	return s.InstanceStatuses
}

func (s *ListInstanceStatusResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstanceStatusResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstanceStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceStatusResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListInstanceStatusResponseBody) SetInstanceStatuses(v []*ListInstanceStatusResponseBodyInstanceStatuses) *ListInstanceStatusResponseBody {
	s.InstanceStatuses = v
	return s
}

func (s *ListInstanceStatusResponseBody) SetPageNumber(v int32) *ListInstanceStatusResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceStatusResponseBody) SetPageSize(v int32) *ListInstanceStatusResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListInstanceStatusResponseBody) SetRequestId(v string) *ListInstanceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceStatusResponseBody) SetTotalCount(v int32) *ListInstanceStatusResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListInstanceStatusResponseBody) Validate() error {
	if s.InstanceStatuses != nil {
		for _, item := range s.InstanceStatuses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstanceStatusResponseBodyInstanceStatuses struct {
	// The ID of the simple application server.
	//
	// example:
	//
	// a9a6474b935d41bcb531250bb5d****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The status of the simple application server. Valid values:
	//
	// - Pending
	//
	// - Starting
	//
	// - Running
	//
	// - Stopping
	//
	// - Stopped
	//
	// - Resetting
	//
	// - Upgrading
	//
	// - Disabled
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListInstanceStatusResponseBodyInstanceStatuses) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceStatusResponseBodyInstanceStatuses) GoString() string {
	return s.String()
}

func (s *ListInstanceStatusResponseBodyInstanceStatuses) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstanceStatusResponseBodyInstanceStatuses) GetStatus() *string {
	return s.Status
}

func (s *ListInstanceStatusResponseBodyInstanceStatuses) SetInstanceId(v string) *ListInstanceStatusResponseBodyInstanceStatuses {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceStatusResponseBodyInstanceStatuses) SetStatus(v string) *ListInstanceStatusResponseBodyInstanceStatuses {
	s.Status = &v
	return s
}

func (s *ListInstanceStatusResponseBodyInstanceStatuses) Validate() error {
	return dara.Validate(s)
}
