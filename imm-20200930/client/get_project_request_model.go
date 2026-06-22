// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectName(v string) *GetProjectRequest
	GetProjectName() *string
	SetWithStatistics(v bool) *GetProjectRequest
	GetWithStatistics() *bool
}

type GetProjectRequest struct {
	// The project name. For information about how to obtain the project name, see [创建项目](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// Specifies whether to collect file statistics. Default value: false, which indicates that file statistics are not collected.
	//
	// - File statistics are collected. The FileCount and TotalFileSize values in the returned Project struct are valid.
	//
	// - File statistics are not collected. The FileCount and TotalFileSize values in the returned Project struct may be inaccurate or both may be 0.
	//
	// 	Notice: Only files in datasets created before December 20, 2025 can be counted..
	//
	// example:
	//
	// true
	WithStatistics *bool `json:"WithStatistics,omitempty" xml:"WithStatistics,omitempty"`
}

func (s GetProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProjectRequest) GoString() string {
	return s.String()
}

func (s *GetProjectRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetProjectRequest) GetWithStatistics() *bool {
	return s.WithStatistics
}

func (s *GetProjectRequest) SetProjectName(v string) *GetProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *GetProjectRequest) SetWithStatistics(v bool) *GetProjectRequest {
	s.WithStatistics = &v
	return s
}

func (s *GetProjectRequest) Validate() error {
	return dara.Validate(s)
}
