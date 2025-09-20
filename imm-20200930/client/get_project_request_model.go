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
	// The name of the project. You can obtain the name from the response of the [CreateProject](https://help.aliyun.com/document_detail/478153.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// Specifies whether to enable real-time retrieval of file statistics. Default value: false.
	//
	// 	- If you set the value to true, the returned values of FileCount and TotalFileSize in the response are valid.
	//
	// 	- If you set the value to false, the returned values of FileCount and TotalFileSize in the response are invalid or equal to 0.
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
