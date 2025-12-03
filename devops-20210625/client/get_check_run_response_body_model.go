// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetCheckRunResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetCheckRunResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetCheckRunResponseBody
	GetRequestId() *string
	SetResult(v *GetCheckRunResponseBodyResult) *GetCheckRunResponseBody
	GetResult() *GetCheckRunResponseBodyResult
	SetSuccess(v bool) *GetCheckRunResponseBody
	GetSuccess() *bool
}

type GetCheckRunResponseBody struct {
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
	RequestId *string                        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetCheckRunResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetCheckRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCheckRunResponseBody) GoString() string {
	return s.String()
}

func (s *GetCheckRunResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetCheckRunResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetCheckRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCheckRunResponseBody) GetResult() *GetCheckRunResponseBodyResult {
	return s.Result
}

func (s *GetCheckRunResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCheckRunResponseBody) SetErrorCode(v string) *GetCheckRunResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetCheckRunResponseBody) SetErrorMessage(v string) *GetCheckRunResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetCheckRunResponseBody) SetRequestId(v string) *GetCheckRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCheckRunResponseBody) SetResult(v *GetCheckRunResponseBodyResult) *GetCheckRunResponseBody {
	s.Result = v
	return s
}

func (s *GetCheckRunResponseBody) SetSuccess(v bool) *GetCheckRunResponseBody {
	s.Success = &v
	return s
}

func (s *GetCheckRunResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCheckRunResponseBodyResult struct {
	Annotations []*GetCheckRunResponseBodyResultAnnotations `json:"annotations,omitempty" xml:"annotations,omitempty" type:"Repeated"`
	CheckSuite  *GetCheckRunResponseBodyResultCheckSuite    `json:"checkSuite,omitempty" xml:"checkSuite,omitempty" type:"Struct"`
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
	// 5240
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// my-check-ci
	Name   *string                              `json:"name,omitempty" xml:"name,omitempty"`
	Output *GetCheckRunResponseBodyResultOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
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
	UpdatedAt *string                              `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	Writer    *GetCheckRunResponseBodyResultWriter `json:"writer,omitempty" xml:"writer,omitempty" type:"Struct"`
}

func (s GetCheckRunResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetCheckRunResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetCheckRunResponseBodyResult) GetAnnotations() []*GetCheckRunResponseBodyResultAnnotations {
	return s.Annotations
}

func (s *GetCheckRunResponseBodyResult) GetCheckSuite() *GetCheckRunResponseBodyResultCheckSuite {
	return s.CheckSuite
}

func (s *GetCheckRunResponseBodyResult) GetCompletedAt() *string {
	return s.CompletedAt
}

func (s *GetCheckRunResponseBodyResult) GetConclusion() *string {
	return s.Conclusion
}

func (s *GetCheckRunResponseBodyResult) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetCheckRunResponseBodyResult) GetDetailsUrl() *string {
	return s.DetailsUrl
}

func (s *GetCheckRunResponseBodyResult) GetExternalId() *string {
	return s.ExternalId
}

func (s *GetCheckRunResponseBodyResult) GetHeadSha() *string {
	return s.HeadSha
}

func (s *GetCheckRunResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *GetCheckRunResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *GetCheckRunResponseBodyResult) GetOutput() *GetCheckRunResponseBodyResultOutput {
	return s.Output
}

func (s *GetCheckRunResponseBodyResult) GetStartedAt() *string {
	return s.StartedAt
}

func (s *GetCheckRunResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *GetCheckRunResponseBodyResult) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetCheckRunResponseBodyResult) GetWriter() *GetCheckRunResponseBodyResultWriter {
	return s.Writer
}

func (s *GetCheckRunResponseBodyResult) SetAnnotations(v []*GetCheckRunResponseBodyResultAnnotations) *GetCheckRunResponseBodyResult {
	s.Annotations = v
	return s
}

func (s *GetCheckRunResponseBodyResult) SetCheckSuite(v *GetCheckRunResponseBodyResultCheckSuite) *GetCheckRunResponseBodyResult {
	s.CheckSuite = v
	return s
}

func (s *GetCheckRunResponseBodyResult) SetCompletedAt(v string) *GetCheckRunResponseBodyResult {
	s.CompletedAt = &v
	return s
}

func (s *GetCheckRunResponseBodyResult) SetConclusion(v string) *GetCheckRunResponseBodyResult {
	s.Conclusion = &v
	return s
}

func (s *GetCheckRunResponseBodyResult) SetCreatedAt(v string) *GetCheckRunResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *GetCheckRunResponseBodyResult) SetDetailsUrl(v string) *GetCheckRunResponseBodyResult {
	s.DetailsUrl = &v
	return s
}

func (s *GetCheckRunResponseBodyResult) SetExternalId(v string) *GetCheckRunResponseBodyResult {
	s.ExternalId = &v
	return s
}

func (s *GetCheckRunResponseBodyResult) SetHeadSha(v string) *GetCheckRunResponseBodyResult {
	s.HeadSha = &v
	return s
}

func (s *GetCheckRunResponseBodyResult) SetId(v int64) *GetCheckRunResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetCheckRunResponseBodyResult) SetName(v string) *GetCheckRunResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetCheckRunResponseBodyResult) SetOutput(v *GetCheckRunResponseBodyResultOutput) *GetCheckRunResponseBodyResult {
	s.Output = v
	return s
}

func (s *GetCheckRunResponseBodyResult) SetStartedAt(v string) *GetCheckRunResponseBodyResult {
	s.StartedAt = &v
	return s
}

func (s *GetCheckRunResponseBodyResult) SetStatus(v string) *GetCheckRunResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetCheckRunResponseBodyResult) SetUpdatedAt(v string) *GetCheckRunResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *GetCheckRunResponseBodyResult) SetWriter(v *GetCheckRunResponseBodyResultWriter) *GetCheckRunResponseBodyResult {
	s.Writer = v
	return s
}

func (s *GetCheckRunResponseBodyResult) Validate() error {
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

type GetCheckRunResponseBodyResultAnnotations struct {
	// example:
	//
	// warning
	AnnotationLevel *string `json:"annotationLevel,omitempty" xml:"annotationLevel,omitempty"`
	// example:
	//
	// 4
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

func (s GetCheckRunResponseBodyResultAnnotations) String() string {
	return dara.Prettify(s)
}

func (s GetCheckRunResponseBodyResultAnnotations) GoString() string {
	return s.String()
}

func (s *GetCheckRunResponseBodyResultAnnotations) GetAnnotationLevel() *string {
	return s.AnnotationLevel
}

func (s *GetCheckRunResponseBodyResultAnnotations) GetEndColumn() *int64 {
	return s.EndColumn
}

func (s *GetCheckRunResponseBodyResultAnnotations) GetEndLine() *int64 {
	return s.EndLine
}

func (s *GetCheckRunResponseBodyResultAnnotations) GetId() *int64 {
	return s.Id
}

func (s *GetCheckRunResponseBodyResultAnnotations) GetMessage() *string {
	return s.Message
}

func (s *GetCheckRunResponseBodyResultAnnotations) GetPath() *string {
	return s.Path
}

func (s *GetCheckRunResponseBodyResultAnnotations) GetRawDetails() *string {
	return s.RawDetails
}

func (s *GetCheckRunResponseBodyResultAnnotations) GetStartColumn() *int64 {
	return s.StartColumn
}

func (s *GetCheckRunResponseBodyResultAnnotations) GetStartLine() *int64 {
	return s.StartLine
}

func (s *GetCheckRunResponseBodyResultAnnotations) GetTitle() *string {
	return s.Title
}

func (s *GetCheckRunResponseBodyResultAnnotations) SetAnnotationLevel(v string) *GetCheckRunResponseBodyResultAnnotations {
	s.AnnotationLevel = &v
	return s
}

func (s *GetCheckRunResponseBodyResultAnnotations) SetEndColumn(v int64) *GetCheckRunResponseBodyResultAnnotations {
	s.EndColumn = &v
	return s
}

func (s *GetCheckRunResponseBodyResultAnnotations) SetEndLine(v int64) *GetCheckRunResponseBodyResultAnnotations {
	s.EndLine = &v
	return s
}

func (s *GetCheckRunResponseBodyResultAnnotations) SetId(v int64) *GetCheckRunResponseBodyResultAnnotations {
	s.Id = &v
	return s
}

func (s *GetCheckRunResponseBodyResultAnnotations) SetMessage(v string) *GetCheckRunResponseBodyResultAnnotations {
	s.Message = &v
	return s
}

func (s *GetCheckRunResponseBodyResultAnnotations) SetPath(v string) *GetCheckRunResponseBodyResultAnnotations {
	s.Path = &v
	return s
}

func (s *GetCheckRunResponseBodyResultAnnotations) SetRawDetails(v string) *GetCheckRunResponseBodyResultAnnotations {
	s.RawDetails = &v
	return s
}

func (s *GetCheckRunResponseBodyResultAnnotations) SetStartColumn(v int64) *GetCheckRunResponseBodyResultAnnotations {
	s.StartColumn = &v
	return s
}

func (s *GetCheckRunResponseBodyResultAnnotations) SetStartLine(v int64) *GetCheckRunResponseBodyResultAnnotations {
	s.StartLine = &v
	return s
}

func (s *GetCheckRunResponseBodyResultAnnotations) SetTitle(v string) *GetCheckRunResponseBodyResultAnnotations {
	s.Title = &v
	return s
}

func (s *GetCheckRunResponseBodyResultAnnotations) Validate() error {
	return dara.Validate(s)
}

type GetCheckRunResponseBodyResultCheckSuite struct {
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
}

func (s GetCheckRunResponseBodyResultCheckSuite) String() string {
	return dara.Prettify(s)
}

func (s GetCheckRunResponseBodyResultCheckSuite) GoString() string {
	return s.String()
}

func (s *GetCheckRunResponseBodyResultCheckSuite) GetId() *int64 {
	return s.Id
}

func (s *GetCheckRunResponseBodyResultCheckSuite) SetId(v int64) *GetCheckRunResponseBodyResultCheckSuite {
	s.Id = &v
	return s
}

func (s *GetCheckRunResponseBodyResultCheckSuite) Validate() error {
	return dara.Validate(s)
}

type GetCheckRunResponseBodyResultOutput struct {
	Images  []*GetCheckRunResponseBodyResultOutputImages `json:"images,omitempty" xml:"images,omitempty" type:"Repeated"`
	Summary *string                                      `json:"summary,omitempty" xml:"summary,omitempty"`
	Text    *string                                      `json:"text,omitempty" xml:"text,omitempty"`
	// example:
	//
	// Mighty Readme report
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetCheckRunResponseBodyResultOutput) String() string {
	return dara.Prettify(s)
}

func (s GetCheckRunResponseBodyResultOutput) GoString() string {
	return s.String()
}

func (s *GetCheckRunResponseBodyResultOutput) GetImages() []*GetCheckRunResponseBodyResultOutputImages {
	return s.Images
}

func (s *GetCheckRunResponseBodyResultOutput) GetSummary() *string {
	return s.Summary
}

func (s *GetCheckRunResponseBodyResultOutput) GetText() *string {
	return s.Text
}

func (s *GetCheckRunResponseBodyResultOutput) GetTitle() *string {
	return s.Title
}

func (s *GetCheckRunResponseBodyResultOutput) SetImages(v []*GetCheckRunResponseBodyResultOutputImages) *GetCheckRunResponseBodyResultOutput {
	s.Images = v
	return s
}

func (s *GetCheckRunResponseBodyResultOutput) SetSummary(v string) *GetCheckRunResponseBodyResultOutput {
	s.Summary = &v
	return s
}

func (s *GetCheckRunResponseBodyResultOutput) SetText(v string) *GetCheckRunResponseBodyResultOutput {
	s.Text = &v
	return s
}

func (s *GetCheckRunResponseBodyResultOutput) SetTitle(v string) *GetCheckRunResponseBodyResultOutput {
	s.Title = &v
	return s
}

func (s *GetCheckRunResponseBodyResultOutput) Validate() error {
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

type GetCheckRunResponseBodyResultOutputImages struct {
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

func (s GetCheckRunResponseBodyResultOutputImages) String() string {
	return dara.Prettify(s)
}

func (s GetCheckRunResponseBodyResultOutputImages) GoString() string {
	return s.String()
}

func (s *GetCheckRunResponseBodyResultOutputImages) GetAlt() *string {
	return s.Alt
}

func (s *GetCheckRunResponseBodyResultOutputImages) GetCaption() *string {
	return s.Caption
}

func (s *GetCheckRunResponseBodyResultOutputImages) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *GetCheckRunResponseBodyResultOutputImages) SetAlt(v string) *GetCheckRunResponseBodyResultOutputImages {
	s.Alt = &v
	return s
}

func (s *GetCheckRunResponseBodyResultOutputImages) SetCaption(v string) *GetCheckRunResponseBodyResultOutputImages {
	s.Caption = &v
	return s
}

func (s *GetCheckRunResponseBodyResultOutputImages) SetImageUrl(v string) *GetCheckRunResponseBodyResultOutputImages {
	s.ImageUrl = &v
	return s
}

func (s *GetCheckRunResponseBodyResultOutputImages) Validate() error {
	return dara.Validate(s)
}

type GetCheckRunResponseBodyResultWriter struct {
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

func (s GetCheckRunResponseBodyResultWriter) String() string {
	return dara.Prettify(s)
}

func (s GetCheckRunResponseBodyResultWriter) GoString() string {
	return s.String()
}

func (s *GetCheckRunResponseBodyResultWriter) GetId() *string {
	return s.Id
}

func (s *GetCheckRunResponseBodyResultWriter) GetLogoUrl() *string {
	return s.LogoUrl
}

func (s *GetCheckRunResponseBodyResultWriter) GetName() *string {
	return s.Name
}

func (s *GetCheckRunResponseBodyResultWriter) GetSlug() *string {
	return s.Slug
}

func (s *GetCheckRunResponseBodyResultWriter) GetType() *string {
	return s.Type
}

func (s *GetCheckRunResponseBodyResultWriter) SetId(v string) *GetCheckRunResponseBodyResultWriter {
	s.Id = &v
	return s
}

func (s *GetCheckRunResponseBodyResultWriter) SetLogoUrl(v string) *GetCheckRunResponseBodyResultWriter {
	s.LogoUrl = &v
	return s
}

func (s *GetCheckRunResponseBodyResultWriter) SetName(v string) *GetCheckRunResponseBodyResultWriter {
	s.Name = &v
	return s
}

func (s *GetCheckRunResponseBodyResultWriter) SetSlug(v string) *GetCheckRunResponseBodyResultWriter {
	s.Slug = &v
	return s
}

func (s *GetCheckRunResponseBodyResultWriter) SetType(v string) *GetCheckRunResponseBodyResultWriter {
	s.Type = &v
	return s
}

func (s *GetCheckRunResponseBodyResultWriter) Validate() error {
	return dara.Validate(s)
}
