// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecordPlanDeviceByPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *QueryRecordPlanDeviceByPlanRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *QueryRecordPlanDeviceByPlanRequest
	GetPageSize() *int32
	SetPlanId(v string) *QueryRecordPlanDeviceByPlanRequest
	GetPlanId() *string
}

type QueryRecordPlanDeviceByPlanRequest struct {
	// The page number to return. Pages start from page 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page. Minimum value: 1. Maximum value: 50. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the recording plan. You can call [QueryRecordPlans](https://help.aliyun.com/document_detail/107896.html) to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5b60bf0dd55944c19d7817442bc9****
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
}

func (s QueryRecordPlanDeviceByPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRecordPlanDeviceByPlanRequest) GoString() string {
	return s.String()
}

func (s *QueryRecordPlanDeviceByPlanRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryRecordPlanDeviceByPlanRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryRecordPlanDeviceByPlanRequest) GetPlanId() *string {
	return s.PlanId
}

func (s *QueryRecordPlanDeviceByPlanRequest) SetCurrentPage(v int32) *QueryRecordPlanDeviceByPlanRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryRecordPlanDeviceByPlanRequest) SetPageSize(v int32) *QueryRecordPlanDeviceByPlanRequest {
	s.PageSize = &v
	return s
}

func (s *QueryRecordPlanDeviceByPlanRequest) SetPlanId(v string) *QueryRecordPlanDeviceByPlanRequest {
	s.PlanId = &v
	return s
}

func (s *QueryRecordPlanDeviceByPlanRequest) Validate() error {
	return dara.Validate(s)
}
