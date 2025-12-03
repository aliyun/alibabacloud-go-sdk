// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCheckRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateCheckRunResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateCheckRunResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateCheckRunResponseBody
	GetRequestId() *string
	SetResult(v *UpdateCheckRunResponseBodyResult) *UpdateCheckRunResponseBody
	GetResult() *UpdateCheckRunResponseBodyResult
	SetSuccess(v bool) *UpdateCheckRunResponseBody
	GetSuccess() *bool
}

type UpdateCheckRunResponseBody struct {
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
	Result    *UpdateCheckRunResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateCheckRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCheckRunResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCheckRunResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateCheckRunResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateCheckRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCheckRunResponseBody) GetResult() *UpdateCheckRunResponseBodyResult {
	return s.Result
}

func (s *UpdateCheckRunResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateCheckRunResponseBody) SetErrorCode(v string) *UpdateCheckRunResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateCheckRunResponseBody) SetErrorMessage(v string) *UpdateCheckRunResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateCheckRunResponseBody) SetRequestId(v string) *UpdateCheckRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCheckRunResponseBody) SetResult(v *UpdateCheckRunResponseBodyResult) *UpdateCheckRunResponseBody {
	s.Result = v
	return s
}

func (s *UpdateCheckRunResponseBody) SetSuccess(v bool) *UpdateCheckRunResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCheckRunResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateCheckRunResponseBodyResult struct {
	Annotations []*UpdateCheckRunResponseBodyResultAnnotations `json:"annotations,omitempty" xml:"annotations,omitempty" type:"Repeated"`
	CheckSuite  *UpdateCheckRunResponseBodyResultCheckSuite    `json:"checkSuite,omitempty" xml:"checkSuite,omitempty" type:"Struct"`
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
	Output *UpdateCheckRunResponseBodyResultOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
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
	Writer    *UpdateCheckRunResponseBodyResultWriter `json:"writer,omitempty" xml:"writer,omitempty" type:"Struct"`
}

func (s UpdateCheckRunResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateCheckRunResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateCheckRunResponseBodyResult) GetAnnotations() []*UpdateCheckRunResponseBodyResultAnnotations {
	return s.Annotations
}

func (s *UpdateCheckRunResponseBodyResult) GetCheckSuite() *UpdateCheckRunResponseBodyResultCheckSuite {
	return s.CheckSuite
}

func (s *UpdateCheckRunResponseBodyResult) GetCompletedAt() *string {
	return s.CompletedAt
}

func (s *UpdateCheckRunResponseBodyResult) GetConclusion() *string {
	return s.Conclusion
}

func (s *UpdateCheckRunResponseBodyResult) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *UpdateCheckRunResponseBodyResult) GetDetailsUrl() *string {
	return s.DetailsUrl
}

func (s *UpdateCheckRunResponseBodyResult) GetExternalId() *string {
	return s.ExternalId
}

func (s *UpdateCheckRunResponseBodyResult) GetHeadSha() *string {
	return s.HeadSha
}

func (s *UpdateCheckRunResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *UpdateCheckRunResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *UpdateCheckRunResponseBodyResult) GetOutput() *UpdateCheckRunResponseBodyResultOutput {
	return s.Output
}

func (s *UpdateCheckRunResponseBodyResult) GetStartedAt() *string {
	return s.StartedAt
}

func (s *UpdateCheckRunResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *UpdateCheckRunResponseBodyResult) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *UpdateCheckRunResponseBodyResult) GetWriter() *UpdateCheckRunResponseBodyResultWriter {
	return s.Writer
}

func (s *UpdateCheckRunResponseBodyResult) SetAnnotations(v []*UpdateCheckRunResponseBodyResultAnnotations) *UpdateCheckRunResponseBodyResult {
	s.Annotations = v
	return s
}

func (s *UpdateCheckRunResponseBodyResult) SetCheckSuite(v *UpdateCheckRunResponseBodyResultCheckSuite) *UpdateCheckRunResponseBodyResult {
	s.CheckSuite = v
	return s
}

func (s *UpdateCheckRunResponseBodyResult) SetCompletedAt(v string) *UpdateCheckRunResponseBodyResult {
	s.CompletedAt = &v
	return s
}

func (s *UpdateCheckRunResponseBodyResult) SetConclusion(v string) *UpdateCheckRunResponseBodyResult {
	s.Conclusion = &v
	return s
}

func (s *UpdateCheckRunResponseBodyResult) SetCreatedAt(v string) *UpdateCheckRunResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *UpdateCheckRunResponseBodyResult) SetDetailsUrl(v string) *UpdateCheckRunResponseBodyResult {
	s.DetailsUrl = &v
	return s
}

func (s *UpdateCheckRunResponseBodyResult) SetExternalId(v string) *UpdateCheckRunResponseBodyResult {
	s.ExternalId = &v
	return s
}

func (s *UpdateCheckRunResponseBodyResult) SetHeadSha(v string) *UpdateCheckRunResponseBodyResult {
	s.HeadSha = &v
	return s
}

func (s *UpdateCheckRunResponseBodyResult) SetId(v int64) *UpdateCheckRunResponseBodyResult {
	s.Id = &v
	return s
}

func (s *UpdateCheckRunResponseBodyResult) SetName(v string) *UpdateCheckRunResponseBodyResult {
	s.Name = &v
	return s
}

func (s *UpdateCheckRunResponseBodyResult) SetOutput(v *UpdateCheckRunResponseBodyResultOutput) *UpdateCheckRunResponseBodyResult {
	s.Output = v
	return s
}

func (s *UpdateCheckRunResponseBodyResult) SetStartedAt(v string) *UpdateCheckRunResponseBodyResult {
	s.StartedAt = &v
	return s
}

func (s *UpdateCheckRunResponseBodyResult) SetStatus(v string) *UpdateCheckRunResponseBodyResult {
	s.Status = &v
	return s
}

func (s *UpdateCheckRunResponseBodyResult) SetUpdatedAt(v string) *UpdateCheckRunResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *UpdateCheckRunResponseBodyResult) SetWriter(v *UpdateCheckRunResponseBodyResultWriter) *UpdateCheckRunResponseBodyResult {
	s.Writer = v
	return s
}

func (s *UpdateCheckRunResponseBodyResult) Validate() error {
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

type UpdateCheckRunResponseBodyResultAnnotations struct {
	// example:
	//
	// warning
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
	// 1
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

func (s UpdateCheckRunResponseBodyResultAnnotations) String() string {
	return dara.Prettify(s)
}

func (s UpdateCheckRunResponseBodyResultAnnotations) GoString() string {
	return s.String()
}

func (s *UpdateCheckRunResponseBodyResultAnnotations) GetAnnotationLevel() *string {
	return s.AnnotationLevel
}

func (s *UpdateCheckRunResponseBodyResultAnnotations) GetEndColumn() *int64 {
	return s.EndColumn
}

func (s *UpdateCheckRunResponseBodyResultAnnotations) GetEndLine() *int64 {
	return s.EndLine
}

func (s *UpdateCheckRunResponseBodyResultAnnotations) GetId() *int64 {
	return s.Id
}

func (s *UpdateCheckRunResponseBodyResultAnnotations) GetMessage() *string {
	return s.Message
}

func (s *UpdateCheckRunResponseBodyResultAnnotations) GetPath() *string {
	return s.Path
}

func (s *UpdateCheckRunResponseBodyResultAnnotations) GetRawDetails() *string {
	return s.RawDetails
}

func (s *UpdateCheckRunResponseBodyResultAnnotations) GetStartColumn() *int64 {
	return s.StartColumn
}

func (s *UpdateCheckRunResponseBodyResultAnnotations) GetStartLine() *int64 {
	return s.StartLine
}

func (s *UpdateCheckRunResponseBodyResultAnnotations) GetTitle() *string {
	return s.Title
}

func (s *UpdateCheckRunResponseBodyResultAnnotations) SetAnnotationLevel(v string) *UpdateCheckRunResponseBodyResultAnnotations {
	s.AnnotationLevel = &v
	return s
}

func (s *UpdateCheckRunResponseBodyResultAnnotations) SetEndColumn(v int64) *UpdateCheckRunResponseBodyResultAnnotations {
	s.EndColumn = &v
	return s
}

func (s *UpdateCheckRunResponseBodyResultAnnotations) SetEndLine(v int64) *UpdateCheckRunResponseBodyResultAnnotations {
	s.EndLine = &v
	return s
}

func (s *UpdateCheckRunResponseBodyResultAnnotations) SetId(v int64) *UpdateCheckRunResponseBodyResultAnnotations {
	s.Id = &v
	return s
}

func (s *UpdateCheckRunResponseBodyResultAnnotations) SetMessage(v string) *UpdateCheckRunResponseBodyResultAnnotations {
	s.Message = &v
	return s
}

func (s *UpdateCheckRunResponseBodyResultAnnotations) SetPath(v string) *UpdateCheckRunResponseBodyResultAnnotations {
	s.Path = &v
	return s
}

func (s *UpdateCheckRunResponseBodyResultAnnotations) SetRawDetails(v string) *UpdateCheckRunResponseBodyResultAnnotations {
	s.RawDetails = &v
	return s
}

func (s *UpdateCheckRunResponseBodyResultAnnotations) SetStartColumn(v int64) *UpdateCheckRunResponseBodyResultAnnotations {
	s.StartColumn = &v
	return s
}

func (s *UpdateCheckRunResponseBodyResultAnnotations) SetStartLine(v int64) *UpdateCheckRunResponseBodyResultAnnotations {
	s.StartLine = &v
	return s
}

func (s *UpdateCheckRunResponseBodyResultAnnotations) SetTitle(v string) *UpdateCheckRunResponseBodyResultAnnotations {
	s.Title = &v
	return s
}

func (s *UpdateCheckRunResponseBodyResultAnnotations) Validate() error {
	return dara.Validate(s)
}

type UpdateCheckRunResponseBodyResultCheckSuite struct {
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
}

func (s UpdateCheckRunResponseBodyResultCheckSuite) String() string {
	return dara.Prettify(s)
}

func (s UpdateCheckRunResponseBodyResultCheckSuite) GoString() string {
	return s.String()
}

func (s *UpdateCheckRunResponseBodyResultCheckSuite) GetId() *int64 {
	return s.Id
}

func (s *UpdateCheckRunResponseBodyResultCheckSuite) SetId(v int64) *UpdateCheckRunResponseBodyResultCheckSuite {
	s.Id = &v
	return s
}

func (s *UpdateCheckRunResponseBodyResultCheckSuite) Validate() error {
	return dara.Validate(s)
}

type UpdateCheckRunResponseBodyResultOutput struct {
	Images  []*UpdateCheckRunResponseBodyResultOutputImages `json:"images,omitempty" xml:"images,omitempty" type:"Repeated"`
	Summary *string                                         `json:"summary,omitempty" xml:"summary,omitempty"`
	Text    *string                                         `json:"text,omitempty" xml:"text,omitempty"`
	// example:
	//
	// Mighty Readme report
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s UpdateCheckRunResponseBodyResultOutput) String() string {
	return dara.Prettify(s)
}

func (s UpdateCheckRunResponseBodyResultOutput) GoString() string {
	return s.String()
}

func (s *UpdateCheckRunResponseBodyResultOutput) GetImages() []*UpdateCheckRunResponseBodyResultOutputImages {
	return s.Images
}

func (s *UpdateCheckRunResponseBodyResultOutput) GetSummary() *string {
	return s.Summary
}

func (s *UpdateCheckRunResponseBodyResultOutput) GetText() *string {
	return s.Text
}

func (s *UpdateCheckRunResponseBodyResultOutput) GetTitle() *string {
	return s.Title
}

func (s *UpdateCheckRunResponseBodyResultOutput) SetImages(v []*UpdateCheckRunResponseBodyResultOutputImages) *UpdateCheckRunResponseBodyResultOutput {
	s.Images = v
	return s
}

func (s *UpdateCheckRunResponseBodyResultOutput) SetSummary(v string) *UpdateCheckRunResponseBodyResultOutput {
	s.Summary = &v
	return s
}

func (s *UpdateCheckRunResponseBodyResultOutput) SetText(v string) *UpdateCheckRunResponseBodyResultOutput {
	s.Text = &v
	return s
}

func (s *UpdateCheckRunResponseBodyResultOutput) SetTitle(v string) *UpdateCheckRunResponseBodyResultOutput {
	s.Title = &v
	return s
}

func (s *UpdateCheckRunResponseBodyResultOutput) Validate() error {
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

type UpdateCheckRunResponseBodyResultOutputImages struct {
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

func (s UpdateCheckRunResponseBodyResultOutputImages) String() string {
	return dara.Prettify(s)
}

func (s UpdateCheckRunResponseBodyResultOutputImages) GoString() string {
	return s.String()
}

func (s *UpdateCheckRunResponseBodyResultOutputImages) GetAlt() *string {
	return s.Alt
}

func (s *UpdateCheckRunResponseBodyResultOutputImages) GetCaption() *string {
	return s.Caption
}

func (s *UpdateCheckRunResponseBodyResultOutputImages) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *UpdateCheckRunResponseBodyResultOutputImages) SetAlt(v string) *UpdateCheckRunResponseBodyResultOutputImages {
	s.Alt = &v
	return s
}

func (s *UpdateCheckRunResponseBodyResultOutputImages) SetCaption(v string) *UpdateCheckRunResponseBodyResultOutputImages {
	s.Caption = &v
	return s
}

func (s *UpdateCheckRunResponseBodyResultOutputImages) SetImageUrl(v string) *UpdateCheckRunResponseBodyResultOutputImages {
	s.ImageUrl = &v
	return s
}

func (s *UpdateCheckRunResponseBodyResultOutputImages) Validate() error {
	return dara.Validate(s)
}

type UpdateCheckRunResponseBodyResultWriter struct {
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

func (s UpdateCheckRunResponseBodyResultWriter) String() string {
	return dara.Prettify(s)
}

func (s UpdateCheckRunResponseBodyResultWriter) GoString() string {
	return s.String()
}

func (s *UpdateCheckRunResponseBodyResultWriter) GetId() *string {
	return s.Id
}

func (s *UpdateCheckRunResponseBodyResultWriter) GetLogoUrl() *string {
	return s.LogoUrl
}

func (s *UpdateCheckRunResponseBodyResultWriter) GetName() *string {
	return s.Name
}

func (s *UpdateCheckRunResponseBodyResultWriter) GetSlug() *string {
	return s.Slug
}

func (s *UpdateCheckRunResponseBodyResultWriter) GetType() *string {
	return s.Type
}

func (s *UpdateCheckRunResponseBodyResultWriter) SetId(v string) *UpdateCheckRunResponseBodyResultWriter {
	s.Id = &v
	return s
}

func (s *UpdateCheckRunResponseBodyResultWriter) SetLogoUrl(v string) *UpdateCheckRunResponseBodyResultWriter {
	s.LogoUrl = &v
	return s
}

func (s *UpdateCheckRunResponseBodyResultWriter) SetName(v string) *UpdateCheckRunResponseBodyResultWriter {
	s.Name = &v
	return s
}

func (s *UpdateCheckRunResponseBodyResultWriter) SetSlug(v string) *UpdateCheckRunResponseBodyResultWriter {
	s.Slug = &v
	return s
}

func (s *UpdateCheckRunResponseBodyResultWriter) SetType(v string) *UpdateCheckRunResponseBodyResultWriter {
	s.Type = &v
	return s
}

func (s *UpdateCheckRunResponseBodyResultWriter) Validate() error {
	return dara.Validate(s)
}
