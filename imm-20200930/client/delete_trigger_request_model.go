// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTriggerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteTriggerRequest
	GetId() *string
	SetProjectName(v string) *DeleteTriggerRequest
	GetProjectName() *string
}

type DeleteTriggerRequest struct {
	// The ID of the trigger. You can obtain the ID of the trigger from the response of the [CreateTrigger](https://help.aliyun.com/document_detail/479912.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// trigger-9f72636a-0f0c-4baf-ae78-38b27b******
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

func (s DeleteTriggerRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTriggerRequest) GoString() string {
	return s.String()
}

func (s *DeleteTriggerRequest) GetId() *string {
	return s.Id
}

func (s *DeleteTriggerRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DeleteTriggerRequest) SetId(v string) *DeleteTriggerRequest {
	s.Id = &v
	return s
}

func (s *DeleteTriggerRequest) SetProjectName(v string) *DeleteTriggerRequest {
	s.ProjectName = &v
	return s
}

func (s *DeleteTriggerRequest) Validate() error {
	return dara.Validate(s)
}
