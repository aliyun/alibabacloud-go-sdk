// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCheckRunRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *UpdateCheckRunRequest
	GetAccessToken() *string
	SetAnnotations(v []*UpdateCheckRunRequestAnnotations) *UpdateCheckRunRequest
	GetAnnotations() []*UpdateCheckRunRequestAnnotations
	SetCompletedAt(v string) *UpdateCheckRunRequest
	GetCompletedAt() *string
	SetConclusion(v string) *UpdateCheckRunRequest
	GetConclusion() *string
	SetDetailsUrl(v string) *UpdateCheckRunRequest
	GetDetailsUrl() *string
	SetExternalId(v string) *UpdateCheckRunRequest
	GetExternalId() *string
	SetName(v string) *UpdateCheckRunRequest
	GetName() *string
	SetOutput(v *UpdateCheckRunRequestOutput) *UpdateCheckRunRequest
	GetOutput() *UpdateCheckRunRequestOutput
	SetStartedAt(v string) *UpdateCheckRunRequest
	GetStartedAt() *string
	SetStatus(v string) *UpdateCheckRunRequest
	GetStatus() *string
	SetCheckRunId(v int64) *UpdateCheckRunRequest
	GetCheckRunId() *int64
	SetOrganizationId(v string) *UpdateCheckRunRequest
	GetOrganizationId() *string
	SetRepositoryIdentity(v string) *UpdateCheckRunRequest
	GetRepositoryIdentity() *string
}

type UpdateCheckRunRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string                             `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	Annotations []*UpdateCheckRunRequestAnnotations `json:"annotations,omitempty" xml:"annotations,omitempty" type:"Repeated"`
	// example:
	//
	// 2023-03-15T08:00:00Z
	CompletedAt *string `json:"completedAt,omitempty" xml:"completedAt,omitempty"`
	// example:
	//
	// success
	Conclusion *string `json:"conclusion,omitempty" xml:"conclusion,omitempty"`
	// example:
	//
	// xx
	DetailsUrl *string `json:"detailsUrl,omitempty" xml:"detailsUrl,omitempty"`
	// example:
	//
	// 42
	ExternalId *string `json:"externalId,omitempty" xml:"externalId,omitempty"`
	// example:
	//
	// my-check-ci
	Name   *string                      `json:"name,omitempty" xml:"name,omitempty"`
	Output *UpdateCheckRunRequestOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	// example:
	//
	// 2023-03-15T08:00:00Z
	StartedAt *string `json:"startedAt,omitempty" xml:"startedAt,omitempty"`
	// example:
	//
	// completed
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	CheckRunId *int64 `json:"checkRunId,omitempty" xml:"checkRunId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60de7a6852743a5162b5f957
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// This parameter is required.
	RepositoryIdentity *string `json:"repositoryIdentity,omitempty" xml:"repositoryIdentity,omitempty"`
}

func (s UpdateCheckRunRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCheckRunRequest) GoString() string {
	return s.String()
}

func (s *UpdateCheckRunRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *UpdateCheckRunRequest) GetAnnotations() []*UpdateCheckRunRequestAnnotations {
	return s.Annotations
}

func (s *UpdateCheckRunRequest) GetCompletedAt() *string {
	return s.CompletedAt
}

func (s *UpdateCheckRunRequest) GetConclusion() *string {
	return s.Conclusion
}

func (s *UpdateCheckRunRequest) GetDetailsUrl() *string {
	return s.DetailsUrl
}

func (s *UpdateCheckRunRequest) GetExternalId() *string {
	return s.ExternalId
}

func (s *UpdateCheckRunRequest) GetName() *string {
	return s.Name
}

func (s *UpdateCheckRunRequest) GetOutput() *UpdateCheckRunRequestOutput {
	return s.Output
}

func (s *UpdateCheckRunRequest) GetStartedAt() *string {
	return s.StartedAt
}

func (s *UpdateCheckRunRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateCheckRunRequest) GetCheckRunId() *int64 {
	return s.CheckRunId
}

func (s *UpdateCheckRunRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *UpdateCheckRunRequest) GetRepositoryIdentity() *string {
	return s.RepositoryIdentity
}

func (s *UpdateCheckRunRequest) SetAccessToken(v string) *UpdateCheckRunRequest {
	s.AccessToken = &v
	return s
}

func (s *UpdateCheckRunRequest) SetAnnotations(v []*UpdateCheckRunRequestAnnotations) *UpdateCheckRunRequest {
	s.Annotations = v
	return s
}

func (s *UpdateCheckRunRequest) SetCompletedAt(v string) *UpdateCheckRunRequest {
	s.CompletedAt = &v
	return s
}

func (s *UpdateCheckRunRequest) SetConclusion(v string) *UpdateCheckRunRequest {
	s.Conclusion = &v
	return s
}

func (s *UpdateCheckRunRequest) SetDetailsUrl(v string) *UpdateCheckRunRequest {
	s.DetailsUrl = &v
	return s
}

func (s *UpdateCheckRunRequest) SetExternalId(v string) *UpdateCheckRunRequest {
	s.ExternalId = &v
	return s
}

func (s *UpdateCheckRunRequest) SetName(v string) *UpdateCheckRunRequest {
	s.Name = &v
	return s
}

func (s *UpdateCheckRunRequest) SetOutput(v *UpdateCheckRunRequestOutput) *UpdateCheckRunRequest {
	s.Output = v
	return s
}

func (s *UpdateCheckRunRequest) SetStartedAt(v string) *UpdateCheckRunRequest {
	s.StartedAt = &v
	return s
}

func (s *UpdateCheckRunRequest) SetStatus(v string) *UpdateCheckRunRequest {
	s.Status = &v
	return s
}

func (s *UpdateCheckRunRequest) SetCheckRunId(v int64) *UpdateCheckRunRequest {
	s.CheckRunId = &v
	return s
}

func (s *UpdateCheckRunRequest) SetOrganizationId(v string) *UpdateCheckRunRequest {
	s.OrganizationId = &v
	return s
}

func (s *UpdateCheckRunRequest) SetRepositoryIdentity(v string) *UpdateCheckRunRequest {
	s.RepositoryIdentity = &v
	return s
}

func (s *UpdateCheckRunRequest) Validate() error {
	if s.Annotations != nil {
		for _, item := range s.Annotations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Output != nil {
		if err := s.Output.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateCheckRunRequestAnnotations struct {
	// This parameter is required.
	//
	// example:
	//
	// failure
	AnnotationLevel *string `json:"annotationLevel,omitempty" xml:"annotationLevel,omitempty"`
	// example:
	//
	// 5
	EndColumn *int64 `json:"endColumn,omitempty" xml:"endColumn,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	EndLine *int64 `json:"endLine,omitempty" xml:"endLine,omitempty"`
	// This parameter is required.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// demo/test.txt
	Path       *string `json:"path,omitempty" xml:"path,omitempty"`
	RawDetails *string `json:"rawDetails,omitempty" xml:"rawDetails,omitempty"`
	// example:
	//
	// 3
	StartColumn *int64 `json:"startColumn,omitempty" xml:"startColumn,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	StartLine *int64  `json:"startLine,omitempty" xml:"startLine,omitempty"`
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s UpdateCheckRunRequestAnnotations) String() string {
	return dara.Prettify(s)
}

func (s UpdateCheckRunRequestAnnotations) GoString() string {
	return s.String()
}

func (s *UpdateCheckRunRequestAnnotations) GetAnnotationLevel() *string {
	return s.AnnotationLevel
}

func (s *UpdateCheckRunRequestAnnotations) GetEndColumn() *int64 {
	return s.EndColumn
}

func (s *UpdateCheckRunRequestAnnotations) GetEndLine() *int64 {
	return s.EndLine
}

func (s *UpdateCheckRunRequestAnnotations) GetMessage() *string {
	return s.Message
}

func (s *UpdateCheckRunRequestAnnotations) GetPath() *string {
	return s.Path
}

func (s *UpdateCheckRunRequestAnnotations) GetRawDetails() *string {
	return s.RawDetails
}

func (s *UpdateCheckRunRequestAnnotations) GetStartColumn() *int64 {
	return s.StartColumn
}

func (s *UpdateCheckRunRequestAnnotations) GetStartLine() *int64 {
	return s.StartLine
}

func (s *UpdateCheckRunRequestAnnotations) GetTitle() *string {
	return s.Title
}

func (s *UpdateCheckRunRequestAnnotations) SetAnnotationLevel(v string) *UpdateCheckRunRequestAnnotations {
	s.AnnotationLevel = &v
	return s
}

func (s *UpdateCheckRunRequestAnnotations) SetEndColumn(v int64) *UpdateCheckRunRequestAnnotations {
	s.EndColumn = &v
	return s
}

func (s *UpdateCheckRunRequestAnnotations) SetEndLine(v int64) *UpdateCheckRunRequestAnnotations {
	s.EndLine = &v
	return s
}

func (s *UpdateCheckRunRequestAnnotations) SetMessage(v string) *UpdateCheckRunRequestAnnotations {
	s.Message = &v
	return s
}

func (s *UpdateCheckRunRequestAnnotations) SetPath(v string) *UpdateCheckRunRequestAnnotations {
	s.Path = &v
	return s
}

func (s *UpdateCheckRunRequestAnnotations) SetRawDetails(v string) *UpdateCheckRunRequestAnnotations {
	s.RawDetails = &v
	return s
}

func (s *UpdateCheckRunRequestAnnotations) SetStartColumn(v int64) *UpdateCheckRunRequestAnnotations {
	s.StartColumn = &v
	return s
}

func (s *UpdateCheckRunRequestAnnotations) SetStartLine(v int64) *UpdateCheckRunRequestAnnotations {
	s.StartLine = &v
	return s
}

func (s *UpdateCheckRunRequestAnnotations) SetTitle(v string) *UpdateCheckRunRequestAnnotations {
	s.Title = &v
	return s
}

func (s *UpdateCheckRunRequestAnnotations) Validate() error {
	return dara.Validate(s)
}

type UpdateCheckRunRequestOutput struct {
	Images []*UpdateCheckRunRequestOutputImages `json:"images,omitempty" xml:"images,omitempty" type:"Repeated"`
	// This parameter is required.
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
	Text    *string `json:"text,omitempty" xml:"text,omitempty"`
	// example:
	//
	// Mighty Readme report
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s UpdateCheckRunRequestOutput) String() string {
	return dara.Prettify(s)
}

func (s UpdateCheckRunRequestOutput) GoString() string {
	return s.String()
}

func (s *UpdateCheckRunRequestOutput) GetImages() []*UpdateCheckRunRequestOutputImages {
	return s.Images
}

func (s *UpdateCheckRunRequestOutput) GetSummary() *string {
	return s.Summary
}

func (s *UpdateCheckRunRequestOutput) GetText() *string {
	return s.Text
}

func (s *UpdateCheckRunRequestOutput) GetTitle() *string {
	return s.Title
}

func (s *UpdateCheckRunRequestOutput) SetImages(v []*UpdateCheckRunRequestOutputImages) *UpdateCheckRunRequestOutput {
	s.Images = v
	return s
}

func (s *UpdateCheckRunRequestOutput) SetSummary(v string) *UpdateCheckRunRequestOutput {
	s.Summary = &v
	return s
}

func (s *UpdateCheckRunRequestOutput) SetText(v string) *UpdateCheckRunRequestOutput {
	s.Text = &v
	return s
}

func (s *UpdateCheckRunRequestOutput) SetTitle(v string) *UpdateCheckRunRequestOutput {
	s.Title = &v
	return s
}

func (s *UpdateCheckRunRequestOutput) Validate() error {
	if s.Images != nil {
		for _, item := range s.Images {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateCheckRunRequestOutputImages struct {
	// This parameter is required.
	//
	// example:
	//
	// test-image-alt
	Alt *string `json:"alt,omitempty" xml:"alt,omitempty"`
	// example:
	//
	// test
	Caption *string `json:"caption,omitempty" xml:"caption,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxx
	ImageUrl *string `json:"imageUrl,omitempty" xml:"imageUrl,omitempty"`
}

func (s UpdateCheckRunRequestOutputImages) String() string {
	return dara.Prettify(s)
}

func (s UpdateCheckRunRequestOutputImages) GoString() string {
	return s.String()
}

func (s *UpdateCheckRunRequestOutputImages) GetAlt() *string {
	return s.Alt
}

func (s *UpdateCheckRunRequestOutputImages) GetCaption() *string {
	return s.Caption
}

func (s *UpdateCheckRunRequestOutputImages) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *UpdateCheckRunRequestOutputImages) SetAlt(v string) *UpdateCheckRunRequestOutputImages {
	s.Alt = &v
	return s
}

func (s *UpdateCheckRunRequestOutputImages) SetCaption(v string) *UpdateCheckRunRequestOutputImages {
	s.Caption = &v
	return s
}

func (s *UpdateCheckRunRequestOutputImages) SetImageUrl(v string) *UpdateCheckRunRequestOutputImages {
	s.ImageUrl = &v
	return s
}

func (s *UpdateCheckRunRequestOutputImages) Validate() error {
	return dara.Validate(s)
}
