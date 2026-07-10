// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLangfuseProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateLangfuseProjectResponseBodyData) *CreateLangfuseProjectResponseBody
	GetData() *CreateLangfuseProjectResponseBodyData
	SetRequestId(v string) *CreateLangfuseProjectResponseBody
	GetRequestId() *string
}

type CreateLangfuseProjectResponseBody struct {
	// The returned result.
	Data *CreateLangfuseProjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 2C7393F1-5FD1-5CEE-A2EA-270A2CF99693
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLangfuseProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLangfuseProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLangfuseProjectResponseBody) GetData() *CreateLangfuseProjectResponseBodyData {
	return s.Data
}

func (s *CreateLangfuseProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLangfuseProjectResponseBody) SetData(v *CreateLangfuseProjectResponseBodyData) *CreateLangfuseProjectResponseBody {
	s.Data = v
	return s
}

func (s *CreateLangfuseProjectResponseBody) SetRequestId(v string) *CreateLangfuseProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLangfuseProjectResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateLangfuseProjectResponseBodyData struct {
	// The time when the Langfuse project was created.
	//
	// example:
	//
	// 2026-06-24T10:14:33Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The Langfuse project name.
	//
	// example:
	//
	// project_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The Langfuse organization ID.
	//
	// example:
	//
	// cmrbhzx930005jw2q****
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// The Langfuse project ID.
	//
	// example:
	//
	// cmrbhzx930005jw2q****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreateLangfuseProjectResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateLangfuseProjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateLangfuseProjectResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateLangfuseProjectResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateLangfuseProjectResponseBodyData) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *CreateLangfuseProjectResponseBodyData) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateLangfuseProjectResponseBodyData) SetCreatedAt(v string) *CreateLangfuseProjectResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *CreateLangfuseProjectResponseBodyData) SetName(v string) *CreateLangfuseProjectResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateLangfuseProjectResponseBodyData) SetOrganizationId(v string) *CreateLangfuseProjectResponseBodyData {
	s.OrganizationId = &v
	return s
}

func (s *CreateLangfuseProjectResponseBodyData) SetProjectId(v string) *CreateLangfuseProjectResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *CreateLangfuseProjectResponseBodyData) Validate() error {
	return dara.Validate(s)
}
