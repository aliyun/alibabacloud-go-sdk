// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTriggerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetTriggerRequest
	GetId() *string
	SetProjectName(v string) *GetTriggerRequest
	GetProjectName() *string
}

type GetTriggerRequest struct {
	// The ID of the trigger. You can obtain the ID from the response parameters of the [CreateTrigger](https://help.aliyun.com/document_detail/479912.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// trigger-9f72636a-0f0c-4baf-ae78-38b27b******
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

func (s GetTriggerRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTriggerRequest) GoString() string {
	return s.String()
}

func (s *GetTriggerRequest) GetId() *string {
	return s.Id
}

func (s *GetTriggerRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetTriggerRequest) SetId(v string) *GetTriggerRequest {
	s.Id = &v
	return s
}

func (s *GetTriggerRequest) SetProjectName(v string) *GetTriggerRequest {
	s.ProjectName = &v
	return s
}

func (s *GetTriggerRequest) Validate() error {
	return dara.Validate(s)
}
