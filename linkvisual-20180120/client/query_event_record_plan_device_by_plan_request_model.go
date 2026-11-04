// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEventRecordPlanDeviceByPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *QueryEventRecordPlanDeviceByPlanRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *QueryEventRecordPlanDeviceByPlanRequest
	GetPageSize() *int32
	SetPlanId(v string) *QueryEventRecordPlanDeviceByPlanRequest
	GetPlanId() *string
}

type QueryEventRecordPlanDeviceByPlanRequest struct {
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries to return on each page. Minimum value: 1. Maximum value: 50. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the intelligent alert event recording plan. You can call the [QueryEventRecordPlans](https://help.aliyun.com/document_detail/107907.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5b60bf0dd55944c19d7817442bc9****
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
}

func (s QueryEventRecordPlanDeviceByPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryEventRecordPlanDeviceByPlanRequest) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlanDeviceByPlanRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryEventRecordPlanDeviceByPlanRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryEventRecordPlanDeviceByPlanRequest) GetPlanId() *string {
	return s.PlanId
}

func (s *QueryEventRecordPlanDeviceByPlanRequest) SetCurrentPage(v int32) *QueryEventRecordPlanDeviceByPlanRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByPlanRequest) SetPageSize(v int32) *QueryEventRecordPlanDeviceByPlanRequest {
	s.PageSize = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByPlanRequest) SetPlanId(v string) *QueryEventRecordPlanDeviceByPlanRequest {
	s.PlanId = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByPlanRequest) Validate() error {
	return dara.Validate(s)
}
