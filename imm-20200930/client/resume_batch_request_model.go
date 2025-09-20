// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeBatchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *ResumeBatchRequest
	GetId() *string
	SetProjectName(v string) *ResumeBatchRequest
	GetProjectName() *string
}

type ResumeBatchRequest struct {
	// The ID of the batch processing task. You can obtain the ID of the batch processing task from the response of the [CreateBatch](https://help.aliyun.com/document_detail/606694.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// batch-4eb9223f-3e88-42d3-a578-3f2852******
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the project. You can obtain the name of the project from the response of the [CreateProject](https://help.aliyun.com/document_detail/478153.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s ResumeBatchRequest) String() string {
	return dara.Prettify(s)
}

func (s ResumeBatchRequest) GoString() string {
	return s.String()
}

func (s *ResumeBatchRequest) GetId() *string {
	return s.Id
}

func (s *ResumeBatchRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *ResumeBatchRequest) SetId(v string) *ResumeBatchRequest {
	s.Id = &v
	return s
}

func (s *ResumeBatchRequest) SetProjectName(v string) *ResumeBatchRequest {
	s.ProjectName = &v
	return s
}

func (s *ResumeBatchRequest) Validate() error {
	return dara.Validate(s)
}
