// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJobTemplateVersionDetail interface {
	dara.Model
	String() string
	GoString() string
	SetConstraints(v map[string]*string) *JobTemplateVersionDetail
	GetConstraints() map[string]*string
	SetContent(v string) *JobTemplateVersionDetail
	GetContent() *string
	SetCreatedBy(v string) *JobTemplateVersionDetail
	GetCreatedBy() *string
	SetGmtCreated(v string) *JobTemplateVersionDetail
	GetGmtCreated() *string
	SetVersion(v int32) *JobTemplateVersionDetail
	GetVersion() *int32
}

type JobTemplateVersionDetail struct {
	Constraints map[string]*string `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	// Configuration content of the job template. It supports all parameter fields of the CreateJob API and is stored as a JSON object.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// User ID that created this version
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// Creation time of this version
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// Template version number
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s JobTemplateVersionDetail) String() string {
	return dara.Prettify(s)
}

func (s JobTemplateVersionDetail) GoString() string {
	return s.String()
}

func (s *JobTemplateVersionDetail) GetConstraints() map[string]*string {
	return s.Constraints
}

func (s *JobTemplateVersionDetail) GetContent() *string {
	return s.Content
}

func (s *JobTemplateVersionDetail) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *JobTemplateVersionDetail) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *JobTemplateVersionDetail) GetVersion() *int32 {
	return s.Version
}

func (s *JobTemplateVersionDetail) SetConstraints(v map[string]*string) *JobTemplateVersionDetail {
	s.Constraints = v
	return s
}

func (s *JobTemplateVersionDetail) SetContent(v string) *JobTemplateVersionDetail {
	s.Content = &v
	return s
}

func (s *JobTemplateVersionDetail) SetCreatedBy(v string) *JobTemplateVersionDetail {
	s.CreatedBy = &v
	return s
}

func (s *JobTemplateVersionDetail) SetGmtCreated(v string) *JobTemplateVersionDetail {
	s.GmtCreated = &v
	return s
}

func (s *JobTemplateVersionDetail) SetVersion(v int32) *JobTemplateVersionDetail {
	s.Version = &v
	return s
}

func (s *JobTemplateVersionDetail) Validate() error {
	return dara.Validate(s)
}
