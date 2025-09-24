// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasetJobConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigType(v string) *ListDatasetJobConfigsRequest
	GetConfigType() *string
	SetDatasetVersion(v string) *ListDatasetJobConfigsRequest
	GetDatasetVersion() *string
	SetPageNumber(v string) *ListDatasetJobConfigsRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListDatasetJobConfigsRequest
	GetPageSize() *string
	SetWorkspaceId(v string) *ListDatasetJobConfigsRequest
	GetWorkspaceId() *string
}

type ListDatasetJobConfigsRequest struct {
	// The configuration type.
	//
	// 	- MultimodalIntelligentTag
	//
	// 	- MultimodalSemanticIndex
	//
	// example:
	//
	// MultimodalIntelligentTag
	ConfigType     *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	DatasetVersion *string `json:"DatasetVersion,omitempty" xml:"DatasetVersion,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 431514
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListDatasetJobConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetJobConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListDatasetJobConfigsRequest) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListDatasetJobConfigsRequest) GetDatasetVersion() *string {
	return s.DatasetVersion
}

func (s *ListDatasetJobConfigsRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListDatasetJobConfigsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListDatasetJobConfigsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListDatasetJobConfigsRequest) SetConfigType(v string) *ListDatasetJobConfigsRequest {
	s.ConfigType = &v
	return s
}

func (s *ListDatasetJobConfigsRequest) SetDatasetVersion(v string) *ListDatasetJobConfigsRequest {
	s.DatasetVersion = &v
	return s
}

func (s *ListDatasetJobConfigsRequest) SetPageNumber(v string) *ListDatasetJobConfigsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDatasetJobConfigsRequest) SetPageSize(v string) *ListDatasetJobConfigsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDatasetJobConfigsRequest) SetWorkspaceId(v string) *ListDatasetJobConfigsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListDatasetJobConfigsRequest) Validate() error {
	return dara.Validate(s)
}
