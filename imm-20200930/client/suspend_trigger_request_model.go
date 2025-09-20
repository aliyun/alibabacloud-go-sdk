// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendTriggerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *SuspendTriggerRequest
	GetId() *string
	SetProjectName(v string) *SuspendTriggerRequest
	GetProjectName() *string
}

type SuspendTriggerRequest struct {
	// The ID of the trigger.[](~~479912~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// trigger-9f72636a-0f0c-4baf-ae78-38b27b******
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the project.[](~~478153~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s SuspendTriggerRequest) String() string {
	return dara.Prettify(s)
}

func (s SuspendTriggerRequest) GoString() string {
	return s.String()
}

func (s *SuspendTriggerRequest) GetId() *string {
	return s.Id
}

func (s *SuspendTriggerRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *SuspendTriggerRequest) SetId(v string) *SuspendTriggerRequest {
	s.Id = &v
	return s
}

func (s *SuspendTriggerRequest) SetProjectName(v string) *SuspendTriggerRequest {
	s.ProjectName = &v
	return s
}

func (s *SuspendTriggerRequest) Validate() error {
	return dara.Validate(s)
}
