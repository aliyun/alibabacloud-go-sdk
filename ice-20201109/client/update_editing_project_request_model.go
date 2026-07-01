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
	// The business status of the project. You can typically ignore this parameter for standard cloud editing projects. Use this parameter to modify the project\\"s reservation status:
	//
	// - `Reserving`: The project is being reserved.
	//
	// - `ReservationCanceled`: The reservation for the project is canceled.
	//
	// example:
	//
	// Reserving
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The clip parameters for the template, in JSON format. This parameter is required if you specify `TemplateId`.<props="china"> For more information about the format, see [Create and use a standard template](https://help.aliyun.com/document_detail/328557.html) and [Create and use an advanced template](https://help.aliyun.com/document_detail/291418.html).
	//
	// example:
	//
	// See the template user guide.
	ClipsParam *string `json:"ClipsParam,omitempty" xml:"ClipsParam,omitempty"`
	// The project cover.
	//
	// example:
	//
	// https://****.com/6AB4D0E1E1C7446888****.png
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The project description.
	//
	// example:
	//
	// testtimeline001desciption
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****4ee4b97e27b525142a6b2****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The template ID. Use this parameter to quickly build a timeline.	Notice: You can specify only one of `ProjectId`, `Timeline`, and `TemplateId`. If you specify `TemplateId`, `ClipsParam` is required.
	//
	// example:
	//
	// ****96e8864746a0b6f3****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The project timeline, in JSON format.
	Timeline *string `json:"Timeline,omitempty" xml:"Timeline,omitempty"`
	// The project title.
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
