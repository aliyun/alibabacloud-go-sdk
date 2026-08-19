// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListResourcesRequest
	GetName() *string
	SetOwner(v string) *ListResourcesRequest
	GetOwner() *string
	SetPageNumber(v int32) *ListResourcesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListResourcesRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListResourcesRequest
	GetProjectId() *int64
	SetType(v string) *ListResourcesRequest
	GetType() *string
}

type ListResourcesRequest struct {
	// The name of the file resource. Fuzzy match is supported.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the owner, which is the account UID of the workspace administrator. You can log on to the Alibaba Cloud Management Console and view the account UID in the security management section of account management.
	//
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The page number of the request, used for pagination.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the workspace configuration page to obtain the workspace ID.
	//
	// This parameter specifies the DataWorks workspace for this API call.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10002
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The filter condition: resource file type.
	//
	// Valid values:
	//
	// - Python
	//
	// - Jar
	//
	// - Archive
	//
	// - File
	//
	// example:
	//
	// Python
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListResourcesRequest) GetName() *string {
	return s.Name
}

func (s *ListResourcesRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListResourcesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListResourcesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListResourcesRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListResourcesRequest) GetType() *string {
	return s.Type
}

func (s *ListResourcesRequest) SetName(v string) *ListResourcesRequest {
	s.Name = &v
	return s
}

func (s *ListResourcesRequest) SetOwner(v string) *ListResourcesRequest {
	s.Owner = &v
	return s
}

func (s *ListResourcesRequest) SetPageNumber(v int32) *ListResourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourcesRequest) SetPageSize(v int32) *ListResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourcesRequest) SetProjectId(v int64) *ListResourcesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListResourcesRequest) SetType(v string) *ListResourcesRequest {
	s.Type = &v
	return s
}

func (s *ListResourcesRequest) Validate() error {
	return dara.Validate(s)
}
