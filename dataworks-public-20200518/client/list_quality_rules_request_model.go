// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQualityRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEntityId(v int64) *ListQualityRulesRequest
	GetEntityId() *int64
	SetPageNumber(v int32) *ListQualityRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListQualityRulesRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListQualityRulesRequest
	GetProjectId() *int64
	SetProjectName(v string) *ListQualityRulesRequest
	GetProjectName() *string
}

type ListQualityRulesRequest struct {
	// The ID of the partition filter expression. You can call the [GetQualityEntity](https://help.aliyun.com/document_detail/174003.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 20.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 12345
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the compute engine or data source. You can obtain the name from data source configurations.
	//
	// This parameter is required.
	//
	// example:
	//
	// autotest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s ListQualityRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQualityRulesRequest) GoString() string {
	return s.String()
}

func (s *ListQualityRulesRequest) GetEntityId() *int64 {
	return s.EntityId
}

func (s *ListQualityRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListQualityRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListQualityRulesRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListQualityRulesRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListQualityRulesRequest) SetEntityId(v int64) *ListQualityRulesRequest {
	s.EntityId = &v
	return s
}

func (s *ListQualityRulesRequest) SetPageNumber(v int32) *ListQualityRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListQualityRulesRequest) SetPageSize(v int32) *ListQualityRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListQualityRulesRequest) SetProjectId(v int64) *ListQualityRulesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListQualityRulesRequest) SetProjectName(v string) *ListQualityRulesRequest {
	s.ProjectName = &v
	return s
}

func (s *ListQualityRulesRequest) Validate() error {
	return dara.Validate(s)
}
