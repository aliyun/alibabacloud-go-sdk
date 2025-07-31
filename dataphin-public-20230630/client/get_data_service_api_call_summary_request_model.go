// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceApiCallSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetDataServiceApiCallSummaryRequest
	GetEndTime() *string
	SetOpTenantId(v int64) *GetDataServiceApiCallSummaryRequest
	GetOpTenantId() *int64
	SetProjectId(v int32) *GetDataServiceApiCallSummaryRequest
	GetProjectId() *int32
	SetStartTime(v string) *GetDataServiceApiCallSummaryRequest
	GetStartTime() *string
}

type GetDataServiceApiCallSummaryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2025-06-30 20:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 102102
	ProjectId *int32 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2025-06-30 08:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetDataServiceApiCallSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceApiCallSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetDataServiceApiCallSummaryRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetDataServiceApiCallSummaryRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetDataServiceApiCallSummaryRequest) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *GetDataServiceApiCallSummaryRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetDataServiceApiCallSummaryRequest) SetEndTime(v string) *GetDataServiceApiCallSummaryRequest {
	s.EndTime = &v
	return s
}

func (s *GetDataServiceApiCallSummaryRequest) SetOpTenantId(v int64) *GetDataServiceApiCallSummaryRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetDataServiceApiCallSummaryRequest) SetProjectId(v int32) *GetDataServiceApiCallSummaryRequest {
	s.ProjectId = &v
	return s
}

func (s *GetDataServiceApiCallSummaryRequest) SetStartTime(v string) *GetDataServiceApiCallSummaryRequest {
	s.StartTime = &v
	return s
}

func (s *GetDataServiceApiCallSummaryRequest) Validate() error {
	return dara.Validate(s)
}
