// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBatchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteBatchRequest
	GetId() *string
	SetProjectName(v string) *DeleteBatchRequest
	GetProjectName() *string
}

type DeleteBatchRequest struct {
	// The ID of the batch processing task.
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

func (s DeleteBatchRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBatchRequest) GoString() string {
	return s.String()
}

func (s *DeleteBatchRequest) GetId() *string {
	return s.Id
}

func (s *DeleteBatchRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DeleteBatchRequest) SetId(v string) *DeleteBatchRequest {
	s.Id = &v
	return s
}

func (s *DeleteBatchRequest) SetProjectName(v string) *DeleteBatchRequest {
	s.ProjectName = &v
	return s
}

func (s *DeleteBatchRequest) Validate() error {
	return dara.Validate(s)
}
