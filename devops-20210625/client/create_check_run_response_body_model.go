// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCheckRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateCheckRunResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateCheckRunResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateCheckRunResponseBody
	GetRequestId() *string
	SetResult(v *CreateCheckRunResponseBodyResult) *CreateCheckRunResponseBody
	GetResult() *CreateCheckRunResponseBodyResult
	SetSuccess(v bool) *CreateCheckRunResponseBody
	GetSuccess() *bool
}

type CreateCheckRunResponseBody struct {
	// example:
	//
	// SYSTEM_UNKNOWN_ERROR
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *CreateCheckRunResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateCheckRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCheckRunResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCheckRunResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateCheckRunResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateCheckRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCheckRunResponseBody) GetResult() *CreateCheckRunResponseBodyResult {
	return s.Result
}

func (s *CreateCheckRunResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateCheckRunResponseBody) SetErrorCode(v string) *CreateCheckRunResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateCheckRunResponseBody) SetErrorMessage(v string) *CreateCheckRunResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateCheckRunResponseBody) SetRequestId(v string) *CreateCheckRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCheckRunResponseBody) SetResult(v *CreateCheckRunResponseBodyResult) *CreateCheckRunResponseBody {
	s.Result = v
	return s
}

func (s *CreateCheckRunResponseBody) SetSuccess(v bool) *CreateCheckRunResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCheckRunResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCheckRunResponseBodyResult struct {
	Annotations []*CreateCheckRunResponseBodyResultAnnotations `json:"annotations,omitempty" xml:"annotations,omitempty" type:"Repeated"`
	CheckSuite  *CreateCheckRunResponseBodyResultCheckSuite    `json:"checkSuite,omitempty" xml:"checkSuite,omitempty" type:"Struct"`
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
	// 2022-01-14T21:08:26+08:00
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// xxx
	DetailsUrl *string `json:"detailsUrl,omitempty" xml:"detailsUrl,omitempty"`
	// example:
	//
	// 42
	ExternalId *string `json:"externalId,omitempty" xml:"externalId,omitempty"`
	// example:
	//
	// 40f4ccfe019cdd4a62d4acb0c57130106fc7e1be
	HeadSha *string `json:"headSha,omitempty" xml:"headSha,omitempty"`
	// example:
	//
	// 524836
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// my-check-ci
	Name   *string                                 `json:"name,omitempty" xml:"name,omitempty"`
	Output *CreateCheckRunResponseBodyResultOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	// example:
	//
	// 2023-03-15T08:00:00Z
	StartedAt *string `json:"startedAt,omitempty" xml:"startedAt,omitempty"`
	// example:
	//
	// completed
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 2022-01-14T21:08:26+08:00
	UpdatedAt *string                                 `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	Writer    *CreateCheckRunResponseBodyResultWriter `json:"writer,omitempty" xml:"writer,omitempty" type:"Struct"`
}

func (s CreateCheckRunResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateCheckRunResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateCheckRunResponseBodyResult) GetAnnotations() []*CreateCheckRunResponseBodyResultAnnotations {
	return s.Annotations
}

func (s *CreateCheckRunResponseBodyResult) GetCheckSuite() *CreateCheckRunResponseBodyResultCheckSuite {
	return s.CheckSuite
}

func (s *CreateCheckRunResponseBodyResult) GetCompletedAt() *string {
	return s.CompletedAt
}

func (s *CreateCheckRunResponseBodyResult) GetConclusion() *string {
	return s.Conclusion
}

func (s *CreateCheckRunResponseBodyResult) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateCheckRunResponseBodyResult) GetDetailsUrl() *string {
	return s.DetailsUrl
}

func (s *CreateCheckRunResponseBodyResult) GetExternalId() *string {
	return s.ExternalId
}

func (s *CreateCheckRunResponseBodyResult) GetHeadSha() *string {
	return s.HeadSha
}

func (s *CreateCheckRunResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *CreateCheckRunResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *CreateCheckRunResponseBodyResult) GetOutput() *CreateCheckRunResponseBodyResultOutput {
	return s.Output
}

func (s *CreateCheckRunResponseBodyResult) GetStartedAt() *string {
	return s.StartedAt
}

func (s *CreateCheckRunResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *CreateCheckRunResponseBodyResult) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *CreateCheckRunResponseBodyResult) GetWriter() *CreateCheckRunResponseBodyResultWriter {
	return s.Writer
}

func (s *CreateCheckRunResponseBodyResult) SetAnnotations(v []*CreateCheckRunResponseBodyResultAnnotations) *CreateCheckRunResponseBodyResult {
	s.Annotations = v
	return s
}

func (s *CreateCheckRunResponseBodyResult) SetCheckSuite(v *CreateCheckRunResponseBodyResultCheckSuite) *CreateCheckRunResponseBodyResult {
	s.CheckSuite = v
	return s
}

func (s *CreateCheckRunResponseBodyResult) SetCompletedAt(v string) *CreateCheckRunResponseBodyResult {
	s.CompletedAt = &v
	return s
}

func (s *CreateCheckRunResponseBodyResult) SetConclusion(v string) *CreateCheckRunResponseBodyResult {
	s.Conclusion = &v
	return s
}

func (s *CreateCheckRunResponseBodyResult) SetCreatedAt(v string) *CreateCheckRunResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *CreateCheckRunResponseBodyResult) SetDetailsUrl(v string) *CreateCheckRunResponseBodyResult {
	s.DetailsUrl = &v
	return s
}

func (s *CreateCheckRunResponseBodyResult) SetExternalId(v string) *CreateCheckRunResponseBodyResult {
	s.ExternalId = &v
	return s
}

func (s *CreateCheckRunResponseBodyResult) SetHeadSha(v string) *CreateCheckRunResponseBodyResult {
	s.HeadSha = &v
	return s
}

func (s *CreateCheckRunResponseBodyResult) SetId(v int64) *CreateCheckRunResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateCheckRunResponseBodyResult) SetName(v string) *CreateCheckRunResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateCheckRunResponseBodyResult) SetOutput(v *CreateCheckRunResponseBodyResultOutput) *CreateCheckRunResponseBodyResult {
	s.Output = v
	return s
}

func (s *CreateCheckRunResponseBodyResult) SetStartedAt(v string) *CreateCheckRunResponseBodyResult {
	s.StartedAt = &v
	return s
}

func (s *CreateCheckRunResponseBodyResult) SetStatus(v string) *CreateCheckRunResponseBodyResult {
	s.Status = &v
	return s
}

func (s *CreateCheckRunResponseBodyResult) SetUpdatedAt(v string) *CreateCheckRunResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *CreateCheckRunResponseBodyResult) SetWriter(v *CreateCheckRunResponseBodyResultWriter) *CreateCheckRunResponseBodyResult {
	s.Writer = v
	return s
}

func (s *CreateCheckRunResponseBodyResult) Validate() error {
	if s.Annotations != nil {
		for _, item := range s.Annotations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.CheckSuite != nil {
		if err := s.CheckSuite.Validate(); err != nil {
			return err
		}
	}
	if s.Output != nil {
		if err := s.Output.Validate(); err != nil {
			return err
		}
	}
	if s.Writer != nil {
		if err := s.Writer.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCheckRunResponseBodyResultAnnotations struct {
	// example:
	//
	// notice
	AnnotationLevel *string `json:"annotationLevel,omitempty" xml:"annotationLevel,omitempty"`
	// example:
	//
	// 5
	EndColumn *int64 `json:"endColumn,omitempty" xml:"endColumn,omitempty"`
	// example:
	//
	// 2
	EndLine *int64 `json:"endLine,omitempty" xml:"endLine,omitempty"`
	// example:
	//
	// 2
	Id      *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// demo/test.txt
	Path       *string `json:"path,omitempty" xml:"path,omitempty"`
	RawDetails *string `json:"rawDetails,omitempty" xml:"rawDetails,omitempty"`
	// example:
	//
	// 3
	StartColumn *int64 `json:"startColumn,omitempty" xml:"startColumn,omitempty"`
	// example:
	//
	// 1
	StartLine *int64  `json:"startLine,omitempty" xml:"startLine,omitempty"`
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateCheckRunResponseBodyResultAnnotations) String() string {
	return dara.Prettify(s)
}

func (s CreateCheckRunResponseBodyResultAnnotations) GoString() string {
	return s.String()
}

func (s *CreateCheckRunResponseBodyResultAnnotations) GetAnnotationLevel() *string {
	return s.AnnotationLevel
}

func (s *CreateCheckRunResponseBodyResultAnnotations) GetEndColumn() *int64 {
	return s.EndColumn
}

func (s *CreateCheckRunResponseBodyResultAnnotations) GetEndLine() *int64 {
	return s.EndLine
}

func (s *CreateCheckRunResponseBodyResultAnnotations) GetId() *int64 {
	return s.Id
}

func (s *CreateCheckRunResponseBodyResultAnnotations) GetMessage() *string {
	return s.Message
}

func (s *CreateCheckRunResponseBodyResultAnnotations) GetPath() *string {
	return s.Path
}

func (s *CreateCheckRunResponseBodyResultAnnotations) GetRawDetails() *string {
	return s.RawDetails
}

func (s *CreateCheckRunResponseBodyResultAnnotations) GetStartColumn() *int64 {
	return s.StartColumn
}

func (s *CreateCheckRunResponseBodyResultAnnotations) GetStartLine() *int64 {
	return s.StartLine
}

func (s *CreateCheckRunResponseBodyResultAnnotations) GetTitle() *string {
	return s.Title
}

func (s *CreateCheckRunResponseBodyResultAnnotations) SetAnnotationLevel(v string) *CreateCheckRunResponseBodyResultAnnotations {
	s.AnnotationLevel = &v
	return s
}

func (s *CreateCheckRunResponseBodyResultAnnotations) SetEndColumn(v int64) *CreateCheckRunResponseBodyResultAnnotations {
	s.EndColumn = &v
	return s
}

func (s *CreateCheckRunResponseBodyResultAnnotations) SetEndLine(v int64) *CreateCheckRunResponseBodyResultAnnotations {
	s.EndLine = &v
	return s
}

func (s *CreateCheckRunResponseBodyResultAnnotations) SetId(v int64) *CreateCheckRunResponseBodyResultAnnotations {
	s.Id = &v
	return s
}

func (s *CreateCheckRunResponseBodyResultAnnotations) SetMessage(v string) *CreateCheckRunResponseBodyResultAnnotations {
	s.Message = &v
	return s
}

func (s *CreateCheckRunResponseBodyResultAnnotations) SetPath(v string) *CreateCheckRunResponseBodyResultAnnotations {
	s.Path = &v
	return s
}

func (s *CreateCheckRunResponseBodyResultAnnotations) SetRawDetails(v string) *CreateCheckRunResponseBodyResultAnnotations {
	s.RawDetails = &v
	return s
}

func (s *CreateCheckRunResponseBodyResultAnnotations) SetStartColumn(v int64) *CreateCheckRunResponseBodyResultAnnotations {
	s.StartColumn = &v
	return s
}

func (s *CreateCheckRunResponseBodyResultAnnotations) SetStartLine(v int64) *CreateCheckRunResponseBodyResultAnnotations {
	s.StartLine = &v
	return s
}

func (s *CreateCheckRunResponseBodyResultAnnotations) SetTitle(v string) *CreateCheckRunResponseBodyResultAnnotations {
	s.Title = &v
	return s
}

func (s *CreateCheckRunResponseBodyResultAnnotations) Validate() error {
	return dara.Validate(s)
}

type CreateCheckRunResponseBodyResultCheckSuite struct {
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
}

func (s CreateCheckRunResponseBodyResultCheckSuite) String() string {
	return dara.Prettify(s)
}

func (s CreateCheckRunResponseBodyResultCheckSuite) GoString() string {
	return s.String()
}

func (s *CreateCheckRunResponseBodyResultCheckSuite) GetId() *int64 {
	return s.Id
}

func (s *CreateCheckRunResponseBodyResultCheckSuite) SetId(v int64) *CreateCheckRunResponseBodyResultCheckSuite {
	s.Id = &v
	return s
}

func (s *CreateCheckRunResponseBodyResultCheckSuite) Validate() error {
	return dara.Validate(s)
}

type CreateCheckRunResponseBodyResultOutput struct {
	Images  []*CreateCheckRunResponseBodyResultOutputImages `json:"images,omitempty" xml:"images,omitempty" type:"Repeated"`
	Summary *string                                         `json:"summary,omitempty" xml:"summary,omitempty"`
	Text    *string                                         `json:"text,omitempty" xml:"text,omitempty"`
	// example:
	//
	// Mighty Readme report
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateCheckRunResponseBodyResultOutput) String() string {
	return dara.Prettify(s)
}

func (s CreateCheckRunResponseBodyResultOutput) GoString() string {
	return s.String()
}

func (s *CreateCheckRunResponseBodyResultOutput) GetImages() []*CreateCheckRunResponseBodyResultOutputImages {
	return s.Images
}

func (s *CreateCheckRunResponseBodyResultOutput) GetSummary() *string {
	return s.Summary
}

func (s *CreateCheckRunResponseBodyResultOutput) GetText() *string {
	return s.Text
}

func (s *CreateCheckRunResponseBodyResultOutput) GetTitle() *string {
	return s.Title
}

func (s *CreateCheckRunResponseBodyResultOutput) SetImages(v []*CreateCheckRunResponseBodyResultOutputImages) *CreateCheckRunResponseBodyResultOutput {
	s.Images = v
	return s
}

func (s *CreateCheckRunResponseBodyResultOutput) SetSummary(v string) *CreateCheckRunResponseBodyResultOutput {
	s.Summary = &v
	return s
}

func (s *CreateCheckRunResponseBodyResultOutput) SetText(v string) *CreateCheckRunResponseBodyResultOutput {
	s.Text = &v
	return s
}

func (s *CreateCheckRunResponseBodyResultOutput) SetTitle(v string) *CreateCheckRunResponseBodyResultOutput {
	s.Title = &v
	return s
}

func (s *CreateCheckRunResponseBodyResultOutput) Validate() error {
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

type CreateCheckRunResponseBodyResultOutputImages struct {
	// example:
	//
	// test-image-alt
	Alt *string `json:"alt,omitempty" xml:"alt,omitempty"`
	// example:
	//
	// test
	Caption *string `json:"caption,omitempty" xml:"caption,omitempty"`
	// example:
	//
	// xxx
	ImageUrl *string `json:"imageUrl,omitempty" xml:"imageUrl,omitempty"`
}

func (s CreateCheckRunResponseBodyResultOutputImages) String() string {
	return dara.Prettify(s)
}

func (s CreateCheckRunResponseBodyResultOutputImages) GoString() string {
	return s.String()
}

func (s *CreateCheckRunResponseBodyResultOutputImages) GetAlt() *string {
	return s.Alt
}

func (s *CreateCheckRunResponseBodyResultOutputImages) GetCaption() *string {
	return s.Caption
}

func (s *CreateCheckRunResponseBodyResultOutputImages) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *CreateCheckRunResponseBodyResultOutputImages) SetAlt(v string) *CreateCheckRunResponseBodyResultOutputImages {
	s.Alt = &v
	return s
}

func (s *CreateCheckRunResponseBodyResultOutputImages) SetCaption(v string) *CreateCheckRunResponseBodyResultOutputImages {
	s.Caption = &v
	return s
}

func (s *CreateCheckRunResponseBodyResultOutputImages) SetImageUrl(v string) *CreateCheckRunResponseBodyResultOutputImages {
	s.ImageUrl = &v
	return s
}

func (s *CreateCheckRunResponseBodyResultOutputImages) Validate() error {
	return dara.Validate(s)
}

type CreateCheckRunResponseBodyResultWriter struct {
	// example:
	//
	// xxx
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// xxx
	LogoUrl *string `json:"logoUrl,omitempty" xml:"logoUrl,omitempty"`
	// example:
	//
	// test-codeup
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// test-codeup
	Slug *string `json:"slug,omitempty" xml:"slug,omitempty"`
	// example:
	//
	// User
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateCheckRunResponseBodyResultWriter) String() string {
	return dara.Prettify(s)
}

func (s CreateCheckRunResponseBodyResultWriter) GoString() string {
	return s.String()
}

func (s *CreateCheckRunResponseBodyResultWriter) GetId() *string {
	return s.Id
}

func (s *CreateCheckRunResponseBodyResultWriter) GetLogoUrl() *string {
	return s.LogoUrl
}

func (s *CreateCheckRunResponseBodyResultWriter) GetName() *string {
	return s.Name
}

func (s *CreateCheckRunResponseBodyResultWriter) GetSlug() *string {
	return s.Slug
}

func (s *CreateCheckRunResponseBodyResultWriter) GetType() *string {
	return s.Type
}

func (s *CreateCheckRunResponseBodyResultWriter) SetId(v string) *CreateCheckRunResponseBodyResultWriter {
	s.Id = &v
	return s
}

func (s *CreateCheckRunResponseBodyResultWriter) SetLogoUrl(v string) *CreateCheckRunResponseBodyResultWriter {
	s.LogoUrl = &v
	return s
}

func (s *CreateCheckRunResponseBodyResultWriter) SetName(v string) *CreateCheckRunResponseBodyResultWriter {
	s.Name = &v
	return s
}

func (s *CreateCheckRunResponseBodyResultWriter) SetSlug(v string) *CreateCheckRunResponseBodyResultWriter {
	s.Slug = &v
	return s
}

func (s *CreateCheckRunResponseBodyResultWriter) SetType(v string) *CreateCheckRunResponseBodyResultWriter {
	s.Type = &v
	return s
}

func (s *CreateCheckRunResponseBodyResultWriter) Validate() error {
	return dara.Validate(s)
}
