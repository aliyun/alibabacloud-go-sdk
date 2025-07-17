// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataQualityRuleTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreationSource(v string) *ListDataQualityRuleTemplatesRequest
	GetCreationSource() *string
	SetDirectoryPath(v string) *ListDataQualityRuleTemplatesRequest
	GetDirectoryPath() *string
	SetName(v string) *ListDataQualityRuleTemplatesRequest
	GetName() *string
	SetPageNumber(v int32) *ListDataQualityRuleTemplatesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataQualityRuleTemplatesRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListDataQualityRuleTemplatesRequest
	GetProjectId() *int64
}

type ListDataQualityRuleTemplatesRequest struct {
	// The source of the template. This parameter is required. Valid values:
	//
	// 	- System
	//
	// 	- UserDefined
	//
	// example:
	//
	// System
	CreationSource *string `json:"CreationSource,omitempty" xml:"CreationSource,omitempty"`
	// The directory in which the template is stored. Slashes (/) are used to separate directory levels. The name of each directory level can be up to 1,024 characters in length. It cannot contain whitespace characters or slashes (/).
	//
	// example:
	//
	// /ods/order_data
	DirectoryPath *string `json:"DirectoryPath,omitempty" xml:"DirectoryPath,omitempty"`
	// The name of the template. If you want to query a system template, set this parameter to the name of the system template. Fuzzy match is supported.
	//
	// example:
	//
	// Table rows
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The DataWorks workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListDataQualityRuleTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityRuleTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListDataQualityRuleTemplatesRequest) GetCreationSource() *string {
	return s.CreationSource
}

func (s *ListDataQualityRuleTemplatesRequest) GetDirectoryPath() *string {
	return s.DirectoryPath
}

func (s *ListDataQualityRuleTemplatesRequest) GetName() *string {
	return s.Name
}

func (s *ListDataQualityRuleTemplatesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataQualityRuleTemplatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataQualityRuleTemplatesRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDataQualityRuleTemplatesRequest) SetCreationSource(v string) *ListDataQualityRuleTemplatesRequest {
	s.CreationSource = &v
	return s
}

func (s *ListDataQualityRuleTemplatesRequest) SetDirectoryPath(v string) *ListDataQualityRuleTemplatesRequest {
	s.DirectoryPath = &v
	return s
}

func (s *ListDataQualityRuleTemplatesRequest) SetName(v string) *ListDataQualityRuleTemplatesRequest {
	s.Name = &v
	return s
}

func (s *ListDataQualityRuleTemplatesRequest) SetPageNumber(v int32) *ListDataQualityRuleTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataQualityRuleTemplatesRequest) SetPageSize(v int32) *ListDataQualityRuleTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataQualityRuleTemplatesRequest) SetProjectId(v int64) *ListDataQualityRuleTemplatesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDataQualityRuleTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
