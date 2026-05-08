// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIndividuationTextTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrowdPack(v [][]*string) *CreateIndividuationTextTaskRequest
	GetCrowdPack() [][]*string
	SetProjectId(v string) *CreateIndividuationTextTaskRequest
	GetProjectId() *string
	SetTaskName(v string) *CreateIndividuationTextTaskRequest
	GetTaskName() *string
}

type CreateIndividuationTextTaskRequest struct {
	CrowdPack [][]*string `json:"crowdPack,omitempty" xml:"crowdPack,omitempty" type:"Repeated"`
	// example:
	//
	// 840015278620459008
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	TaskName  *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
}

func (s CreateIndividuationTextTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIndividuationTextTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateIndividuationTextTaskRequest) GetCrowdPack() [][]*string {
	return s.CrowdPack
}

func (s *CreateIndividuationTextTaskRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateIndividuationTextTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateIndividuationTextTaskRequest) SetCrowdPack(v [][]*string) *CreateIndividuationTextTaskRequest {
	s.CrowdPack = v
	return s
}

func (s *CreateIndividuationTextTaskRequest) SetProjectId(v string) *CreateIndividuationTextTaskRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateIndividuationTextTaskRequest) SetTaskName(v string) *CreateIndividuationTextTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateIndividuationTextTaskRequest) Validate() error {
	return dara.Validate(s)
}
