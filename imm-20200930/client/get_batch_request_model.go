// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBatchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetBatchRequest
	GetId() *string
	SetProjectName(v string) *GetBatchRequest
	GetProjectName() *string
}

type GetBatchRequest struct {
	// The ID of the batch processing task. For more information about how to obtain the ID, see [CreateBatch](https://help.aliyun.com/document_detail/606694.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// batch-4eb9223f-3e88-42d3-a578-3f2852******
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the project. For more information, see [CreateProject](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s GetBatchRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBatchRequest) GoString() string {
	return s.String()
}

func (s *GetBatchRequest) GetId() *string {
	return s.Id
}

func (s *GetBatchRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetBatchRequest) SetId(v string) *GetBatchRequest {
	s.Id = &v
	return s
}

func (s *GetBatchRequest) SetProjectName(v string) *GetBatchRequest {
	s.ProjectName = &v
	return s
}

func (s *GetBatchRequest) Validate() error {
	return dara.Validate(s)
}
