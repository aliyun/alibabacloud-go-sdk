// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityProjectListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetQualityProjectListRequest
	GetInstanceId() *string
	SetPageNo(v int32) *GetQualityProjectListRequest
	GetPageNo() *int32
	SetPageSize(v int32) *GetQualityProjectListRequest
	GetPageSize() *int32
	SetProjectId(v int64) *GetQualityProjectListRequest
	GetProjectId() *int64
	SetProjectName(v string) *GetQualityProjectListRequest
	GetProjectName() *string
	SetStatus(v int32) *GetQualityProjectListRequest
	GetStatus() *int32
	SetCheckFreqType(v int64) *GetQualityProjectListRequest
	GetCheckFreqType() *int64
}

type GetQualityProjectListRequest struct {
	// Artificial Intelligence Cloud Call Service (AICCS) instance ID. You can obtain it from the Artificial Intelligence Cloud Call Service console.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The current page. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// Page size. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Quality inspection job ID (supports fuzzy search).
	//
	// example:
	//
	// 15
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Quality inspection job name (supports fuzzy search).
	//
	// example:
	//
	// 质检
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// Quality inspection job status. Valid values:
	//
	// - **0**: Start
	//
	// - **1**: Shutdown
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The check frequency type. Valid values:
	//
	// - **1**: Periodic quality inspection
	//
	// - **4**: Temporary quality inspection
	//
	// example:
	//
	// 1
	CheckFreqType *int64 `json:"checkFreqType,omitempty" xml:"checkFreqType,omitempty"`
}

func (s GetQualityProjectListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQualityProjectListRequest) GoString() string {
	return s.String()
}

func (s *GetQualityProjectListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetQualityProjectListRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetQualityProjectListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetQualityProjectListRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetQualityProjectListRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetQualityProjectListRequest) GetStatus() *int32 {
	return s.Status
}

func (s *GetQualityProjectListRequest) GetCheckFreqType() *int64 {
	return s.CheckFreqType
}

func (s *GetQualityProjectListRequest) SetInstanceId(v string) *GetQualityProjectListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetQualityProjectListRequest) SetPageNo(v int32) *GetQualityProjectListRequest {
	s.PageNo = &v
	return s
}

func (s *GetQualityProjectListRequest) SetPageSize(v int32) *GetQualityProjectListRequest {
	s.PageSize = &v
	return s
}

func (s *GetQualityProjectListRequest) SetProjectId(v int64) *GetQualityProjectListRequest {
	s.ProjectId = &v
	return s
}

func (s *GetQualityProjectListRequest) SetProjectName(v string) *GetQualityProjectListRequest {
	s.ProjectName = &v
	return s
}

func (s *GetQualityProjectListRequest) SetStatus(v int32) *GetQualityProjectListRequest {
	s.Status = &v
	return s
}

func (s *GetQualityProjectListRequest) SetCheckFreqType(v int64) *GetQualityProjectListRequest {
	s.CheckFreqType = &v
	return s
}

func (s *GetQualityProjectListRequest) Validate() error {
	return dara.Validate(s)
}
