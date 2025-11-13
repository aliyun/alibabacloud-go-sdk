// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageRepoListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeImageRepoListRequest
	GetCurrentPage() *int32
	SetFieldName(v string) *DescribeImageRepoListRequest
	GetFieldName() *string
	SetFieldValue(v string) *DescribeImageRepoListRequest
	GetFieldValue() *string
	SetOperateType(v string) *DescribeImageRepoListRequest
	GetOperateType() *string
	SetPageSize(v int32) *DescribeImageRepoListRequest
	GetPageSize() *int32
	SetRepoName(v string) *DescribeImageRepoListRequest
	GetRepoName() *string
	SetRepoNamespace(v string) *DescribeImageRepoListRequest
	GetRepoNamespace() *string
	SetSelected(v int32) *DescribeImageRepoListRequest
	GetSelected() *int32
	SetTargetType(v string) *DescribeImageRepoListRequest
	GetTargetType() *string
	SetType(v string) *DescribeImageRepoListRequest
	GetType() *string
}

type DescribeImageRepoListRequest struct {
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The name of the field that is used for the query. Valid values:
	//
	// 	- **repoName**: the name of the image repository
	//
	// 	- **repoNamespace**: the namespace to which the image repository belongs
	//
	// >  This parameter takes effect only when the **OperateType*	- parameter is set to **other**.
	//
	// example:
	//
	// repoName
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// The value of the field that is used for the query.
	//
	// >  This parameter takes effect only when the **OperateType*	- parameter is set to **other**.
	//
	// example:
	//
	// zeus
	FieldValue *string `json:"FieldValue,omitempty" xml:"FieldValue,omitempty"`
	// The type of the operation. Valid values:
	//
	// 	- **count**: counts statistics
	//
	// 	- **other**: others
	//
	// example:
	//
	// count
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// The number of entries to return on each page. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the image repository.
	//
	// example:
	//
	// script7
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The namespace to which the image repository belongs.
	//
	// example:
	//
	// libssh2
	RepoNamespace *string `json:"RepoNamespace,omitempty" xml:"RepoNamespace,omitempty"`
	// Whether it is selected. Values:
	//
	// 	- **0**: NO
	//
	// 	- **1**: YES
	//
	// example:
	//
	// 1
	Selected *int32 `json:"Selected,omitempty" xml:"Selected,omitempty"`
	// The condition by which the feature is applied. Valid values:
	//
	// 	- **image_repo**: the ID of the image repository
	//
	// This parameter is required.
	//
	// example:
	//
	// image_repo
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The type of the feature. Valid values:
	//
	// 	- **image_repo**: image repository protection
	//
	// This parameter is required.
	//
	// example:
	//
	// image_repo
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeImageRepoListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageRepoListRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageRepoListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeImageRepoListRequest) GetFieldName() *string {
	return s.FieldName
}

func (s *DescribeImageRepoListRequest) GetFieldValue() *string {
	return s.FieldValue
}

func (s *DescribeImageRepoListRequest) GetOperateType() *string {
	return s.OperateType
}

func (s *DescribeImageRepoListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImageRepoListRequest) GetRepoName() *string {
	return s.RepoName
}

func (s *DescribeImageRepoListRequest) GetRepoNamespace() *string {
	return s.RepoNamespace
}

func (s *DescribeImageRepoListRequest) GetSelected() *int32 {
	return s.Selected
}

func (s *DescribeImageRepoListRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *DescribeImageRepoListRequest) GetType() *string {
	return s.Type
}

func (s *DescribeImageRepoListRequest) SetCurrentPage(v int32) *DescribeImageRepoListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageRepoListRequest) SetFieldName(v string) *DescribeImageRepoListRequest {
	s.FieldName = &v
	return s
}

func (s *DescribeImageRepoListRequest) SetFieldValue(v string) *DescribeImageRepoListRequest {
	s.FieldValue = &v
	return s
}

func (s *DescribeImageRepoListRequest) SetOperateType(v string) *DescribeImageRepoListRequest {
	s.OperateType = &v
	return s
}

func (s *DescribeImageRepoListRequest) SetPageSize(v int32) *DescribeImageRepoListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeImageRepoListRequest) SetRepoName(v string) *DescribeImageRepoListRequest {
	s.RepoName = &v
	return s
}

func (s *DescribeImageRepoListRequest) SetRepoNamespace(v string) *DescribeImageRepoListRequest {
	s.RepoNamespace = &v
	return s
}

func (s *DescribeImageRepoListRequest) SetSelected(v int32) *DescribeImageRepoListRequest {
	s.Selected = &v
	return s
}

func (s *DescribeImageRepoListRequest) SetTargetType(v string) *DescribeImageRepoListRequest {
	s.TargetType = &v
	return s
}

func (s *DescribeImageRepoListRequest) SetType(v string) *DescribeImageRepoListRequest {
	s.Type = &v
	return s
}

func (s *DescribeImageRepoListRequest) Validate() error {
	return dara.Validate(s)
}
