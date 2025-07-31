// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceApiCallTrendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetDataServiceApiCallTrendRequest
	GetEndTime() *string
	SetOpTenantId(v int64) *GetDataServiceApiCallTrendRequest
	GetOpTenantId() *int64
	SetProjectId(v int32) *GetDataServiceApiCallTrendRequest
	GetProjectId() *int32
	SetStartTime(v string) *GetDataServiceApiCallTrendRequest
	GetStartTime() *string
}

type GetDataServiceApiCallTrendRequest struct {
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

func (s GetDataServiceApiCallTrendRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceApiCallTrendRequest) GoString() string {
	return s.String()
}

func (s *GetDataServiceApiCallTrendRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetDataServiceApiCallTrendRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetDataServiceApiCallTrendRequest) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *GetDataServiceApiCallTrendRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetDataServiceApiCallTrendRequest) SetEndTime(v string) *GetDataServiceApiCallTrendRequest {
	s.EndTime = &v
	return s
}

func (s *GetDataServiceApiCallTrendRequest) SetOpTenantId(v int64) *GetDataServiceApiCallTrendRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetDataServiceApiCallTrendRequest) SetProjectId(v int32) *GetDataServiceApiCallTrendRequest {
	s.ProjectId = &v
	return s
}

func (s *GetDataServiceApiCallTrendRequest) SetStartTime(v string) *GetDataServiceApiCallTrendRequest {
	s.StartTime = &v
	return s
}

func (s *GetDataServiceApiCallTrendRequest) Validate() error {
	return dara.Validate(s)
}
