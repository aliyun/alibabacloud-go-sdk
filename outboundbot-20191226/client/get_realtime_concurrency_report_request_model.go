// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRealtimeConcurrencyReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetRealtimeConcurrencyReportRequest
	GetInstanceId() *string
	SetJobGroupId(v string) *GetRealtimeConcurrencyReportRequest
	GetJobGroupId() *string
	SetPageNumber(v int32) *GetRealtimeConcurrencyReportRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetRealtimeConcurrencyReportRequest
	GetPageSize() *int32
}

type GetRealtimeConcurrencyReportRequest struct {
	// example:
	//
	// 85bf7efa-a07c-498a-850e-99a5849b8589
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// c62e6789-28a8-41db-941e-171a01d3b3b9
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetRealtimeConcurrencyReportRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRealtimeConcurrencyReportRequest) GoString() string {
	return s.String()
}

func (s *GetRealtimeConcurrencyReportRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRealtimeConcurrencyReportRequest) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *GetRealtimeConcurrencyReportRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetRealtimeConcurrencyReportRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetRealtimeConcurrencyReportRequest) SetInstanceId(v string) *GetRealtimeConcurrencyReportRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRealtimeConcurrencyReportRequest) SetJobGroupId(v string) *GetRealtimeConcurrencyReportRequest {
	s.JobGroupId = &v
	return s
}

func (s *GetRealtimeConcurrencyReportRequest) SetPageNumber(v int32) *GetRealtimeConcurrencyReportRequest {
	s.PageNumber = &v
	return s
}

func (s *GetRealtimeConcurrencyReportRequest) SetPageSize(v int32) *GetRealtimeConcurrencyReportRequest {
	s.PageSize = &v
	return s
}

func (s *GetRealtimeConcurrencyReportRequest) Validate() error {
	return dara.Validate(s)
}
