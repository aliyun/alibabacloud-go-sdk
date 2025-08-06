// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddEditingProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProject(v *AddEditingProjectResponseBodyProject) *AddEditingProjectResponseBody
	GetProject() *AddEditingProjectResponseBodyProject
	SetRequestId(v string) *AddEditingProjectResponseBody
	GetRequestId() *string
}

type AddEditingProjectResponseBody struct {
	// The information about the online editing project. For more information about the structure, see [EditingProject](https://help.aliyun.com/document_detail/52839.html).
	Project *AddEditingProjectResponseBodyProject `json:"Project,omitempty" xml:"Project,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// E4EBD2BF-5EB0-4476-8829-9D94E1B1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddEditingProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddEditingProjectResponseBody) GoString() string {
	return s.String()
}

func (s *AddEditingProjectResponseBody) GetProject() *AddEditingProjectResponseBodyProject {
	return s.Project
}

func (s *AddEditingProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddEditingProjectResponseBody) SetProject(v *AddEditingProjectResponseBodyProject) *AddEditingProjectResponseBody {
	s.Project = v
	return s
}

func (s *AddEditingProjectResponseBody) SetRequestId(v string) *AddEditingProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddEditingProjectResponseBody) Validate() error {
	return dara.Validate(s)
}

type AddEditingProjectResponseBodyProject struct {
	// The time when the online editing project was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-01-11T12:00:00Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the online editing project.
	//
	// example:
	//
	// testtimeline001desciption
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the online editing project was last modified. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-01-11T13:00:00Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The ID of the online editing project.
	//
	// example:
	//
	// fb2101bf24bf4df34c4cb3187****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The status of the online editing project. Valid values:
	//
	// 	- **Normal**: the online editing project is in draft.
	//
	// 	- **Producing**: the video is being produced.
	//
	// 	- **Produced**: the video is produced.
	//
	// 	- **ProduceFailed**: the video failed to be produced.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The title of the online editing project.
	//
	// example:
	//
	// testtimeline
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s AddEditingProjectResponseBodyProject) String() string {
	return dara.Prettify(s)
}

func (s AddEditingProjectResponseBodyProject) GoString() string {
	return s.String()
}

func (s *AddEditingProjectResponseBodyProject) GetCreationTime() *string {
	return s.CreationTime
}

func (s *AddEditingProjectResponseBodyProject) GetDescription() *string {
	return s.Description
}

func (s *AddEditingProjectResponseBodyProject) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *AddEditingProjectResponseBodyProject) GetProjectId() *string {
	return s.ProjectId
}

func (s *AddEditingProjectResponseBodyProject) GetStatus() *string {
	return s.Status
}

func (s *AddEditingProjectResponseBodyProject) GetTitle() *string {
	return s.Title
}

func (s *AddEditingProjectResponseBodyProject) SetCreationTime(v string) *AddEditingProjectResponseBodyProject {
	s.CreationTime = &v
	return s
}

func (s *AddEditingProjectResponseBodyProject) SetDescription(v string) *AddEditingProjectResponseBodyProject {
	s.Description = &v
	return s
}

func (s *AddEditingProjectResponseBodyProject) SetModifiedTime(v string) *AddEditingProjectResponseBodyProject {
	s.ModifiedTime = &v
	return s
}

func (s *AddEditingProjectResponseBodyProject) SetProjectId(v string) *AddEditingProjectResponseBodyProject {
	s.ProjectId = &v
	return s
}

func (s *AddEditingProjectResponseBodyProject) SetStatus(v string) *AddEditingProjectResponseBodyProject {
	s.Status = &v
	return s
}

func (s *AddEditingProjectResponseBodyProject) SetTitle(v string) *AddEditingProjectResponseBodyProject {
	s.Title = &v
	return s
}

func (s *AddEditingProjectResponseBodyProject) Validate() error {
	return dara.Validate(s)
}
