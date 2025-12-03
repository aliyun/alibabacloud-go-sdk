// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCheckRunsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListCheckRunsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListCheckRunsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListCheckRunsResponseBody
	GetRequestId() *string
	SetResult(v []*ListCheckRunsResponseBodyResult) *ListCheckRunsResponseBody
	GetResult() []*ListCheckRunsResponseBodyResult
	SetSuccess(v bool) *ListCheckRunsResponseBody
	GetSuccess() *bool
}

type ListCheckRunsResponseBody struct {
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
	RequestId *string                            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListCheckRunsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListCheckRunsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCheckRunsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCheckRunsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListCheckRunsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListCheckRunsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCheckRunsResponseBody) GetResult() []*ListCheckRunsResponseBodyResult {
	return s.Result
}

func (s *ListCheckRunsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCheckRunsResponseBody) SetErrorCode(v string) *ListCheckRunsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListCheckRunsResponseBody) SetErrorMessage(v string) *ListCheckRunsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListCheckRunsResponseBody) SetRequestId(v string) *ListCheckRunsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCheckRunsResponseBody) SetResult(v []*ListCheckRunsResponseBodyResult) *ListCheckRunsResponseBody {
	s.Result = v
	return s
}

func (s *ListCheckRunsResponseBody) SetSuccess(v bool) *ListCheckRunsResponseBody {
	s.Success = &v
	return s
}

func (s *ListCheckRunsResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCheckRunsResponseBodyResult struct {
	Annotations []*ListCheckRunsResponseBodyResultAnnotations `json:"annotations,omitempty" xml:"annotations,omitempty" type:"Repeated"`
	CheckSuite  *ListCheckRunsResponseBodyResultCheckSuite    `json:"checkSuite,omitempty" xml:"checkSuite,omitempty" type:"Struct"`
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
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// my-check-ci
	Name   *string                                `json:"name,omitempty" xml:"name,omitempty"`
	Output *ListCheckRunsResponseBodyResultOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
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
	UpdatedAt *string                                `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	Writer    *ListCheckRunsResponseBodyResultWriter `json:"writer,omitempty" xml:"writer,omitempty" type:"Struct"`
}

func (s ListCheckRunsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListCheckRunsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListCheckRunsResponseBodyResult) GetAnnotations() []*ListCheckRunsResponseBodyResultAnnotations {
	return s.Annotations
}

func (s *ListCheckRunsResponseBodyResult) GetCheckSuite() *ListCheckRunsResponseBodyResultCheckSuite {
	return s.CheckSuite
}

func (s *ListCheckRunsResponseBodyResult) GetCompletedAt() *string {
	return s.CompletedAt
}

func (s *ListCheckRunsResponseBodyResult) GetConclusion() *string {
	return s.Conclusion
}

func (s *ListCheckRunsResponseBodyResult) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListCheckRunsResponseBodyResult) GetDetailsUrl() *string {
	return s.DetailsUrl
}

func (s *ListCheckRunsResponseBodyResult) GetExternalId() *string {
	return s.ExternalId
}

func (s *ListCheckRunsResponseBodyResult) GetHeadSha() *string {
	return s.HeadSha
}

func (s *ListCheckRunsResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *ListCheckRunsResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListCheckRunsResponseBodyResult) GetOutput() *ListCheckRunsResponseBodyResultOutput {
	return s.Output
}

func (s *ListCheckRunsResponseBodyResult) GetStartedAt() *string {
	return s.StartedAt
}

func (s *ListCheckRunsResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *ListCheckRunsResponseBodyResult) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListCheckRunsResponseBodyResult) GetWriter() *ListCheckRunsResponseBodyResultWriter {
	return s.Writer
}

func (s *ListCheckRunsResponseBodyResult) SetAnnotations(v []*ListCheckRunsResponseBodyResultAnnotations) *ListCheckRunsResponseBodyResult {
	s.Annotations = v
	return s
}

func (s *ListCheckRunsResponseBodyResult) SetCheckSuite(v *ListCheckRunsResponseBodyResultCheckSuite) *ListCheckRunsResponseBodyResult {
	s.CheckSuite = v
	return s
}

func (s *ListCheckRunsResponseBodyResult) SetCompletedAt(v string) *ListCheckRunsResponseBodyResult {
	s.CompletedAt = &v
	return s
}

func (s *ListCheckRunsResponseBodyResult) SetConclusion(v string) *ListCheckRunsResponseBodyResult {
	s.Conclusion = &v
	return s
}

func (s *ListCheckRunsResponseBodyResult) SetCreatedAt(v string) *ListCheckRunsResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *ListCheckRunsResponseBodyResult) SetDetailsUrl(v string) *ListCheckRunsResponseBodyResult {
	s.DetailsUrl = &v
	return s
}

func (s *ListCheckRunsResponseBodyResult) SetExternalId(v string) *ListCheckRunsResponseBodyResult {
	s.ExternalId = &v
	return s
}

func (s *ListCheckRunsResponseBodyResult) SetHeadSha(v string) *ListCheckRunsResponseBodyResult {
	s.HeadSha = &v
	return s
}

func (s *ListCheckRunsResponseBodyResult) SetId(v int64) *ListCheckRunsResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListCheckRunsResponseBodyResult) SetName(v string) *ListCheckRunsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListCheckRunsResponseBodyResult) SetOutput(v *ListCheckRunsResponseBodyResultOutput) *ListCheckRunsResponseBodyResult {
	s.Output = v
	return s
}

func (s *ListCheckRunsResponseBodyResult) SetStartedAt(v string) *ListCheckRunsResponseBodyResult {
	s.StartedAt = &v
	return s
}

func (s *ListCheckRunsResponseBodyResult) SetStatus(v string) *ListCheckRunsResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListCheckRunsResponseBodyResult) SetUpdatedAt(v string) *ListCheckRunsResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *ListCheckRunsResponseBodyResult) SetWriter(v *ListCheckRunsResponseBodyResultWriter) *ListCheckRunsResponseBodyResult {
	s.Writer = v
	return s
}

func (s *ListCheckRunsResponseBodyResult) Validate() error {
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

type ListCheckRunsResponseBodyResultAnnotations struct {
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
	// 11806
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

func (s ListCheckRunsResponseBodyResultAnnotations) String() string {
	return dara.Prettify(s)
}

func (s ListCheckRunsResponseBodyResultAnnotations) GoString() string {
	return s.String()
}

func (s *ListCheckRunsResponseBodyResultAnnotations) GetAnnotationLevel() *string {
	return s.AnnotationLevel
}

func (s *ListCheckRunsResponseBodyResultAnnotations) GetEndColumn() *int64 {
	return s.EndColumn
}

func (s *ListCheckRunsResponseBodyResultAnnotations) GetEndLine() *int64 {
	return s.EndLine
}

func (s *ListCheckRunsResponseBodyResultAnnotations) GetId() *int64 {
	return s.Id
}

func (s *ListCheckRunsResponseBodyResultAnnotations) GetMessage() *string {
	return s.Message
}

func (s *ListCheckRunsResponseBodyResultAnnotations) GetPath() *string {
	return s.Path
}

func (s *ListCheckRunsResponseBodyResultAnnotations) GetRawDetails() *string {
	return s.RawDetails
}

func (s *ListCheckRunsResponseBodyResultAnnotations) GetStartColumn() *int64 {
	return s.StartColumn
}

func (s *ListCheckRunsResponseBodyResultAnnotations) GetStartLine() *int64 {
	return s.StartLine
}

func (s *ListCheckRunsResponseBodyResultAnnotations) GetTitle() *string {
	return s.Title
}

func (s *ListCheckRunsResponseBodyResultAnnotations) SetAnnotationLevel(v string) *ListCheckRunsResponseBodyResultAnnotations {
	s.AnnotationLevel = &v
	return s
}

func (s *ListCheckRunsResponseBodyResultAnnotations) SetEndColumn(v int64) *ListCheckRunsResponseBodyResultAnnotations {
	s.EndColumn = &v
	return s
}

func (s *ListCheckRunsResponseBodyResultAnnotations) SetEndLine(v int64) *ListCheckRunsResponseBodyResultAnnotations {
	s.EndLine = &v
	return s
}

func (s *ListCheckRunsResponseBodyResultAnnotations) SetId(v int64) *ListCheckRunsResponseBodyResultAnnotations {
	s.Id = &v
	return s
}

func (s *ListCheckRunsResponseBodyResultAnnotations) SetMessage(v string) *ListCheckRunsResponseBodyResultAnnotations {
	s.Message = &v
	return s
}

func (s *ListCheckRunsResponseBodyResultAnnotations) SetPath(v string) *ListCheckRunsResponseBodyResultAnnotations {
	s.Path = &v
	return s
}

func (s *ListCheckRunsResponseBodyResultAnnotations) SetRawDetails(v string) *ListCheckRunsResponseBodyResultAnnotations {
	s.RawDetails = &v
	return s
}

func (s *ListCheckRunsResponseBodyResultAnnotations) SetStartColumn(v int64) *ListCheckRunsResponseBodyResultAnnotations {
	s.StartColumn = &v
	return s
}

func (s *ListCheckRunsResponseBodyResultAnnotations) SetStartLine(v int64) *ListCheckRunsResponseBodyResultAnnotations {
	s.StartLine = &v
	return s
}

func (s *ListCheckRunsResponseBodyResultAnnotations) SetTitle(v string) *ListCheckRunsResponseBodyResultAnnotations {
	s.Title = &v
	return s
}

func (s *ListCheckRunsResponseBodyResultAnnotations) Validate() error {
	return dara.Validate(s)
}

type ListCheckRunsResponseBodyResultCheckSuite struct {
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
}

func (s ListCheckRunsResponseBodyResultCheckSuite) String() string {
	return dara.Prettify(s)
}

func (s ListCheckRunsResponseBodyResultCheckSuite) GoString() string {
	return s.String()
}

func (s *ListCheckRunsResponseBodyResultCheckSuite) GetId() *int64 {
	return s.Id
}

func (s *ListCheckRunsResponseBodyResultCheckSuite) SetId(v int64) *ListCheckRunsResponseBodyResultCheckSuite {
	s.Id = &v
	return s
}

func (s *ListCheckRunsResponseBodyResultCheckSuite) Validate() error {
	return dara.Validate(s)
}

type ListCheckRunsResponseBodyResultOutput struct {
	Images  []*ListCheckRunsResponseBodyResultOutputImages `json:"images,omitempty" xml:"images,omitempty" type:"Repeated"`
	Summary *string                                        `json:"summary,omitempty" xml:"summary,omitempty"`
	Text    *string                                        `json:"text,omitempty" xml:"text,omitempty"`
	// example:
	//
	// Mighty Readme report
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ListCheckRunsResponseBodyResultOutput) String() string {
	return dara.Prettify(s)
}

func (s ListCheckRunsResponseBodyResultOutput) GoString() string {
	return s.String()
}

func (s *ListCheckRunsResponseBodyResultOutput) GetImages() []*ListCheckRunsResponseBodyResultOutputImages {
	return s.Images
}

func (s *ListCheckRunsResponseBodyResultOutput) GetSummary() *string {
	return s.Summary
}

func (s *ListCheckRunsResponseBodyResultOutput) GetText() *string {
	return s.Text
}

func (s *ListCheckRunsResponseBodyResultOutput) GetTitle() *string {
	return s.Title
}

func (s *ListCheckRunsResponseBodyResultOutput) SetImages(v []*ListCheckRunsResponseBodyResultOutputImages) *ListCheckRunsResponseBodyResultOutput {
	s.Images = v
	return s
}

func (s *ListCheckRunsResponseBodyResultOutput) SetSummary(v string) *ListCheckRunsResponseBodyResultOutput {
	s.Summary = &v
	return s
}

func (s *ListCheckRunsResponseBodyResultOutput) SetText(v string) *ListCheckRunsResponseBodyResultOutput {
	s.Text = &v
	return s
}

func (s *ListCheckRunsResponseBodyResultOutput) SetTitle(v string) *ListCheckRunsResponseBodyResultOutput {
	s.Title = &v
	return s
}

func (s *ListCheckRunsResponseBodyResultOutput) Validate() error {
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

type ListCheckRunsResponseBodyResultOutputImages struct {
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

func (s ListCheckRunsResponseBodyResultOutputImages) String() string {
	return dara.Prettify(s)
}

func (s ListCheckRunsResponseBodyResultOutputImages) GoString() string {
	return s.String()
}

func (s *ListCheckRunsResponseBodyResultOutputImages) GetAlt() *string {
	return s.Alt
}

func (s *ListCheckRunsResponseBodyResultOutputImages) GetCaption() *string {
	return s.Caption
}

func (s *ListCheckRunsResponseBodyResultOutputImages) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ListCheckRunsResponseBodyResultOutputImages) SetAlt(v string) *ListCheckRunsResponseBodyResultOutputImages {
	s.Alt = &v
	return s
}

func (s *ListCheckRunsResponseBodyResultOutputImages) SetCaption(v string) *ListCheckRunsResponseBodyResultOutputImages {
	s.Caption = &v
	return s
}

func (s *ListCheckRunsResponseBodyResultOutputImages) SetImageUrl(v string) *ListCheckRunsResponseBodyResultOutputImages {
	s.ImageUrl = &v
	return s
}

func (s *ListCheckRunsResponseBodyResultOutputImages) Validate() error {
	return dara.Validate(s)
}

type ListCheckRunsResponseBodyResultWriter struct {
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

func (s ListCheckRunsResponseBodyResultWriter) String() string {
	return dara.Prettify(s)
}

func (s ListCheckRunsResponseBodyResultWriter) GoString() string {
	return s.String()
}

func (s *ListCheckRunsResponseBodyResultWriter) GetId() *string {
	return s.Id
}

func (s *ListCheckRunsResponseBodyResultWriter) GetLogoUrl() *string {
	return s.LogoUrl
}

func (s *ListCheckRunsResponseBodyResultWriter) GetName() *string {
	return s.Name
}

func (s *ListCheckRunsResponseBodyResultWriter) GetSlug() *string {
	return s.Slug
}

func (s *ListCheckRunsResponseBodyResultWriter) GetType() *string {
	return s.Type
}

func (s *ListCheckRunsResponseBodyResultWriter) SetId(v string) *ListCheckRunsResponseBodyResultWriter {
	s.Id = &v
	return s
}

func (s *ListCheckRunsResponseBodyResultWriter) SetLogoUrl(v string) *ListCheckRunsResponseBodyResultWriter {
	s.LogoUrl = &v
	return s
}

func (s *ListCheckRunsResponseBodyResultWriter) SetName(v string) *ListCheckRunsResponseBodyResultWriter {
	s.Name = &v
	return s
}

func (s *ListCheckRunsResponseBodyResultWriter) SetSlug(v string) *ListCheckRunsResponseBodyResultWriter {
	s.Slug = &v
	return s
}

func (s *ListCheckRunsResponseBodyResultWriter) SetType(v string) *ListCheckRunsResponseBodyResultWriter {
	s.Type = &v
	return s
}

func (s *ListCheckRunsResponseBodyResultWriter) Validate() error {
	return dara.Validate(s)
}
