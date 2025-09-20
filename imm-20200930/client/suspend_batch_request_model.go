// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendBatchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *SuspendBatchRequest
	GetId() *string
	SetProjectName(v string) *SuspendBatchRequest
	GetProjectName() *string
}

type SuspendBatchRequest struct {
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

func (s SuspendBatchRequest) String() string {
	return dara.Prettify(s)
}

func (s SuspendBatchRequest) GoString() string {
	return s.String()
}

func (s *SuspendBatchRequest) GetId() *string {
	return s.Id
}

func (s *SuspendBatchRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *SuspendBatchRequest) SetId(v string) *SuspendBatchRequest {
	s.Id = &v
	return s
}

func (s *SuspendBatchRequest) SetProjectName(v string) *SuspendBatchRequest {
	s.ProjectName = &v
	return s
}

func (s *SuspendBatchRequest) Validate() error {
	return dara.Validate(s)
}
