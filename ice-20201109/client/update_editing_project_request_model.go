// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEditingProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessStatus(v string) *UpdateEditingProjectRequest
	GetBusinessStatus() *string
	SetClipsParam(v string) *UpdateEditingProjectRequest
	GetClipsParam() *string
	SetCoverURL(v string) *UpdateEditingProjectRequest
	GetCoverURL() *string
	SetDescription(v string) *UpdateEditingProjectRequest
	GetDescription() *string
	SetProjectId(v string) *UpdateEditingProjectRequest
	GetProjectId() *string
	SetTemplateId(v string) *UpdateEditingProjectRequest
	GetTemplateId() *string
	SetTimeline(v string) *UpdateEditingProjectRequest
	GetTimeline() *string
	SetTitle(v string) *UpdateEditingProjectRequest
	GetTitle() *string
}

type UpdateEditingProjectRequest struct {
	// The business status of the project. This parameter can be ignored for general editing projects. Valid values:
	//
	// 	- Reserving
	//
	// 	- ReservationCanceled
	//
	// example:
	//
	// Reserving
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The material parameter corresponding to the template, in the JSON format. If TemplateId is specified, ClipsParam must also be specified.
	ClipsParam *string `json:"ClipsParam,omitempty" xml:"ClipsParam,omitempty"`
	// The thumbnail URL of the online editing project.
	//
	// example:
	//
	// https://****.com/6AB4D0E1E1C7446888****.png
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The description of the online editing project.
	//
	// example:
	//
	// testtimeline001desciption
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the online editing project.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****4ee4b97e27b525142a6b2****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The template ID. This parameter is used to quickly build a timeline with ease. Note: Only one of ProjectId, Timeline, and TemplateId can be specified. If TemplateId is specified, ClipsParam must also be specified.
	//
	// example:
	//
	// ****96e8864746a0b6f3****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Timeline   *string `json:"Timeline,omitempty" xml:"Timeline,omitempty"`
	// The title of the online editing project.
	//
	// example:
	//
	// testtimeline
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateEditingProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEditingProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateEditingProjectRequest) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *UpdateEditingProjectRequest) GetClipsParam() *string {
	return s.ClipsParam
}

func (s *UpdateEditingProjectRequest) GetCoverURL() *string {
	return s.CoverURL
}

func (s *UpdateEditingProjectRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateEditingProjectRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *UpdateEditingProjectRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateEditingProjectRequest) GetTimeline() *string {
	return s.Timeline
}

func (s *UpdateEditingProjectRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateEditingProjectRequest) SetBusinessStatus(v string) *UpdateEditingProjectRequest {
	s.BusinessStatus = &v
	return s
}

func (s *UpdateEditingProjectRequest) SetClipsParam(v string) *UpdateEditingProjectRequest {
	s.ClipsParam = &v
	return s
}

func (s *UpdateEditingProjectRequest) SetCoverURL(v string) *UpdateEditingProjectRequest {
	s.CoverURL = &v
	return s
}

func (s *UpdateEditingProjectRequest) SetDescription(v string) *UpdateEditingProjectRequest {
	s.Description = &v
	return s
}

func (s *UpdateEditingProjectRequest) SetProjectId(v string) *UpdateEditingProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateEditingProjectRequest) SetTemplateId(v string) *UpdateEditingProjectRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateEditingProjectRequest) SetTimeline(v string) *UpdateEditingProjectRequest {
	s.Timeline = &v
	return s
}

func (s *UpdateEditingProjectRequest) SetTitle(v string) *UpdateEditingProjectRequest {
	s.Title = &v
	return s
}

func (s *UpdateEditingProjectRequest) Validate() error {
	return dara.Validate(s)
}
