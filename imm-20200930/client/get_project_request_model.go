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
	// The name of the project. For more information about how to obtain the project name, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// Specifies whether to collect file statistics. Default value: false.
	//
	// - true: File statistics are collected. The FileCount and TotalFileSize fields in the Project struct are accurate and valid.
	//
	// - false: File statistics are not collected. The FileCount and TotalFileSize fields in the Project struct may be inaccurate or both be 0.
	//
	// 	Notice: File statistics are supported only for datasets created before December 20, 2025.
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
