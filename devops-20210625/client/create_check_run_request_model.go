// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCheckRunRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *CreateCheckRunRequest
	GetAccessToken() *string
	SetAnnotations(v []*CreateCheckRunRequestAnnotations) *CreateCheckRunRequest
	GetAnnotations() []*CreateCheckRunRequestAnnotations
	SetCompletedAt(v string) *CreateCheckRunRequest
	GetCompletedAt() *string
	SetConclusion(v string) *CreateCheckRunRequest
	GetConclusion() *string
	SetDetailsUrl(v string) *CreateCheckRunRequest
	GetDetailsUrl() *string
	SetExternalId(v string) *CreateCheckRunRequest
	GetExternalId() *string
	SetHeadSha(v string) *CreateCheckRunRequest
	GetHeadSha() *string
	SetName(v string) *CreateCheckRunRequest
	GetName() *string
	SetOutput(v *CreateCheckRunRequestOutput) *CreateCheckRunRequest
	GetOutput() *CreateCheckRunRequestOutput
	SetStartedAt(v string) *CreateCheckRunRequest
	GetStartedAt() *string
	SetStatus(v string) *CreateCheckRunRequest
	GetStatus() *string
	SetOrganizationId(v string) *CreateCheckRunRequest
	GetOrganizationId() *string
	SetRepositoryIdentity(v string) *CreateCheckRunRequest
	GetRepositoryIdentity() *string
}

type CreateCheckRunRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string                             `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	Annotations []*CreateCheckRunRequestAnnotations `json:"annotations,omitempty" xml:"annotations,omitempty" type:"Repeated"`
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
	// xxx
	DetailsUrl *string `json:"detailsUrl,omitempty" xml:"detailsUrl,omitempty"`
	// example:
	//
	// 42
	ExternalId *string `json:"externalId,omitempty" xml:"externalId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 40f4ccfe019cdd4a62d4acb0c57130106fc7e1be
	HeadSha *string `json:"headSha,omitempty" xml:"headSha,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-check-ci
	Name   *string                      `json:"name,omitempty" xml:"name,omitempty"`
	Output *CreateCheckRunRequestOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
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
	// 5ebbc0228123212b59xxxxx
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// This parameter is required.
	RepositoryIdentity *string `json:"repositoryIdentity,omitempty" xml:"repositoryIdentity,omitempty"`
}

func (s CreateCheckRunRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCheckRunRequest) GoString() string {
	return s.String()
}

func (s *CreateCheckRunRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *CreateCheckRunRequest) GetAnnotations() []*CreateCheckRunRequestAnnotations {
	return s.Annotations
}

func (s *CreateCheckRunRequest) GetCompletedAt() *string {
	return s.CompletedAt
}

func (s *CreateCheckRunRequest) GetConclusion() *string {
	return s.Conclusion
}

func (s *CreateCheckRunRequest) GetDetailsUrl() *string {
	return s.DetailsUrl
}

func (s *CreateCheckRunRequest) GetExternalId() *string {
	return s.ExternalId
}

func (s *CreateCheckRunRequest) GetHeadSha() *string {
	return s.HeadSha
}

func (s *CreateCheckRunRequest) GetName() *string {
	return s.Name
}

func (s *CreateCheckRunRequest) GetOutput() *CreateCheckRunRequestOutput {
	return s.Output
}

func (s *CreateCheckRunRequest) GetStartedAt() *string {
	return s.StartedAt
}

func (s *CreateCheckRunRequest) GetStatus() *string {
	return s.Status
}

func (s *CreateCheckRunRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *CreateCheckRunRequest) GetRepositoryIdentity() *string {
	return s.RepositoryIdentity
}

func (s *CreateCheckRunRequest) SetAccessToken(v string) *CreateCheckRunRequest {
	s.AccessToken = &v
	return s
}

func (s *CreateCheckRunRequest) SetAnnotations(v []*CreateCheckRunRequestAnnotations) *CreateCheckRunRequest {
	s.Annotations = v
	return s
}

func (s *CreateCheckRunRequest) SetCompletedAt(v string) *CreateCheckRunRequest {
	s.CompletedAt = &v
	return s
}

func (s *CreateCheckRunRequest) SetConclusion(v string) *CreateCheckRunRequest {
	s.Conclusion = &v
	return s
}

func (s *CreateCheckRunRequest) SetDetailsUrl(v string) *CreateCheckRunRequest {
	s.DetailsUrl = &v
	return s
}

func (s *CreateCheckRunRequest) SetExternalId(v string) *CreateCheckRunRequest {
	s.ExternalId = &v
	return s
}

func (s *CreateCheckRunRequest) SetHeadSha(v string) *CreateCheckRunRequest {
	s.HeadSha = &v
	return s
}

func (s *CreateCheckRunRequest) SetName(v string) *CreateCheckRunRequest {
	s.Name = &v
	return s
}

func (s *CreateCheckRunRequest) SetOutput(v *CreateCheckRunRequestOutput) *CreateCheckRunRequest {
	s.Output = v
	return s
}

func (s *CreateCheckRunRequest) SetStartedAt(v string) *CreateCheckRunRequest {
	s.StartedAt = &v
	return s
}

func (s *CreateCheckRunRequest) SetStatus(v string) *CreateCheckRunRequest {
	s.Status = &v
	return s
}

func (s *CreateCheckRunRequest) SetOrganizationId(v string) *CreateCheckRunRequest {
	s.OrganizationId = &v
	return s
}

func (s *CreateCheckRunRequest) SetRepositoryIdentity(v string) *CreateCheckRunRequest {
	s.RepositoryIdentity = &v
	return s
}

func (s *CreateCheckRunRequest) Validate() error {
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

type CreateCheckRunRequestAnnotations struct {
	// This parameter is required.
	//
	// example:
	//
	// notice
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

func (s CreateCheckRunRequestAnnotations) String() string {
	return dara.Prettify(s)
}

func (s CreateCheckRunRequestAnnotations) GoString() string {
	return s.String()
}

func (s *CreateCheckRunRequestAnnotations) GetAnnotationLevel() *string {
	return s.AnnotationLevel
}

func (s *CreateCheckRunRequestAnnotations) GetEndColumn() *int64 {
	return s.EndColumn
}

func (s *CreateCheckRunRequestAnnotations) GetEndLine() *int64 {
	return s.EndLine
}

func (s *CreateCheckRunRequestAnnotations) GetMessage() *string {
	return s.Message
}

func (s *CreateCheckRunRequestAnnotations) GetPath() *string {
	return s.Path
}

func (s *CreateCheckRunRequestAnnotations) GetRawDetails() *string {
	return s.RawDetails
}

func (s *CreateCheckRunRequestAnnotations) GetStartColumn() *int64 {
	return s.StartColumn
}

func (s *CreateCheckRunRequestAnnotations) GetStartLine() *int64 {
	return s.StartLine
}

func (s *CreateCheckRunRequestAnnotations) GetTitle() *string {
	return s.Title
}

func (s *CreateCheckRunRequestAnnotations) SetAnnotationLevel(v string) *CreateCheckRunRequestAnnotations {
	s.AnnotationLevel = &v
	return s
}

func (s *CreateCheckRunRequestAnnotations) SetEndColumn(v int64) *CreateCheckRunRequestAnnotations {
	s.EndColumn = &v
	return s
}

func (s *CreateCheckRunRequestAnnotations) SetEndLine(v int64) *CreateCheckRunRequestAnnotations {
	s.EndLine = &v
	return s
}

func (s *CreateCheckRunRequestAnnotations) SetMessage(v string) *CreateCheckRunRequestAnnotations {
	s.Message = &v
	return s
}

func (s *CreateCheckRunRequestAnnotations) SetPath(v string) *CreateCheckRunRequestAnnotations {
	s.Path = &v
	return s
}

func (s *CreateCheckRunRequestAnnotations) SetRawDetails(v string) *CreateCheckRunRequestAnnotations {
	s.RawDetails = &v
	return s
}

func (s *CreateCheckRunRequestAnnotations) SetStartColumn(v int64) *CreateCheckRunRequestAnnotations {
	s.StartColumn = &v
	return s
}

func (s *CreateCheckRunRequestAnnotations) SetStartLine(v int64) *CreateCheckRunRequestAnnotations {
	s.StartLine = &v
	return s
}

func (s *CreateCheckRunRequestAnnotations) SetTitle(v string) *CreateCheckRunRequestAnnotations {
	s.Title = &v
	return s
}

func (s *CreateCheckRunRequestAnnotations) Validate() error {
	return dara.Validate(s)
}

type CreateCheckRunRequestOutput struct {
	Images []*CreateCheckRunRequestOutputImages `json:"images,omitempty" xml:"images,omitempty" type:"Repeated"`
	// This parameter is required.
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
	Text    *string `json:"text,omitempty" xml:"text,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Mighty Readme report
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateCheckRunRequestOutput) String() string {
	return dara.Prettify(s)
}

func (s CreateCheckRunRequestOutput) GoString() string {
	return s.String()
}

func (s *CreateCheckRunRequestOutput) GetImages() []*CreateCheckRunRequestOutputImages {
	return s.Images
}

func (s *CreateCheckRunRequestOutput) GetSummary() *string {
	return s.Summary
}

func (s *CreateCheckRunRequestOutput) GetText() *string {
	return s.Text
}

func (s *CreateCheckRunRequestOutput) GetTitle() *string {
	return s.Title
}

func (s *CreateCheckRunRequestOutput) SetImages(v []*CreateCheckRunRequestOutputImages) *CreateCheckRunRequestOutput {
	s.Images = v
	return s
}

func (s *CreateCheckRunRequestOutput) SetSummary(v string) *CreateCheckRunRequestOutput {
	s.Summary = &v
	return s
}

func (s *CreateCheckRunRequestOutput) SetText(v string) *CreateCheckRunRequestOutput {
	s.Text = &v
	return s
}

func (s *CreateCheckRunRequestOutput) SetTitle(v string) *CreateCheckRunRequestOutput {
	s.Title = &v
	return s
}

func (s *CreateCheckRunRequestOutput) Validate() error {
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

type CreateCheckRunRequestOutputImages struct {
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

func (s CreateCheckRunRequestOutputImages) String() string {
	return dara.Prettify(s)
}

func (s CreateCheckRunRequestOutputImages) GoString() string {
	return s.String()
}

func (s *CreateCheckRunRequestOutputImages) GetAlt() *string {
	return s.Alt
}

func (s *CreateCheckRunRequestOutputImages) GetCaption() *string {
	return s.Caption
}

func (s *CreateCheckRunRequestOutputImages) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *CreateCheckRunRequestOutputImages) SetAlt(v string) *CreateCheckRunRequestOutputImages {
	s.Alt = &v
	return s
}

func (s *CreateCheckRunRequestOutputImages) SetCaption(v string) *CreateCheckRunRequestOutputImages {
	s.Caption = &v
	return s
}

func (s *CreateCheckRunRequestOutputImages) SetImageUrl(v string) *CreateCheckRunRequestOutputImages {
	s.ImageUrl = &v
	return s
}

func (s *CreateCheckRunRequestOutputImages) Validate() error {
	return dara.Validate(s)
}
