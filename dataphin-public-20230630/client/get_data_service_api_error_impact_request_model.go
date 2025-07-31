// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceApiErrorImpactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetDataServiceApiErrorImpactRequest
	GetEndTime() *string
	SetOpTenantId(v int64) *GetDataServiceApiErrorImpactRequest
	GetOpTenantId() *int64
	SetProjectId(v int32) *GetDataServiceApiErrorImpactRequest
	GetProjectId() *int32
	SetStartTime(v string) *GetDataServiceApiErrorImpactRequest
	GetStartTime() *string
}

type GetDataServiceApiErrorImpactRequest struct {
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

func (s GetDataServiceApiErrorImpactRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceApiErrorImpactRequest) GoString() string {
	return s.String()
}

func (s *GetDataServiceApiErrorImpactRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetDataServiceApiErrorImpactRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetDataServiceApiErrorImpactRequest) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *GetDataServiceApiErrorImpactRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetDataServiceApiErrorImpactRequest) SetEndTime(v string) *GetDataServiceApiErrorImpactRequest {
	s.EndTime = &v
	return s
}

func (s *GetDataServiceApiErrorImpactRequest) SetOpTenantId(v int64) *GetDataServiceApiErrorImpactRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetDataServiceApiErrorImpactRequest) SetProjectId(v int32) *GetDataServiceApiErrorImpactRequest {
	s.ProjectId = &v
	return s
}

func (s *GetDataServiceApiErrorImpactRequest) SetStartTime(v string) *GetDataServiceApiErrorImpactRequest {
	s.StartTime = &v
	return s
}

func (s *GetDataServiceApiErrorImpactRequest) Validate() error {
	return dara.Validate(s)
}
