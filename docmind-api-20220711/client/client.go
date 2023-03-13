// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	openplatform "github.com/alibabacloud-go/openplatform-20191219/v2/client"
	fileform "github.com/alibabacloud-go/tea-fileform/service"
	oss "github.com/alibabacloud-go/tea-oss-sdk/client"
	ossutil "github.com/alibabacloud-go/tea-oss-utils/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

type GetDocStructureResultRequest struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetDocStructureResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDocStructureResultRequest) GoString() string {
	return s.String()
}

func (s *GetDocStructureResultRequest) SetId(v string) *GetDocStructureResultRequest {
	s.Id = &v
	return s
}

type GetDocStructureResultResponseBody struct {
	Code      *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Completed *bool                  `json:"Completed,omitempty" xml:"Completed,omitempty"`
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string                `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetDocStructureResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDocStructureResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocStructureResultResponseBody) SetCode(v string) *GetDocStructureResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetDocStructureResultResponseBody) SetCompleted(v bool) *GetDocStructureResultResponseBody {
	s.Completed = &v
	return s
}

func (s *GetDocStructureResultResponseBody) SetData(v map[string]interface{}) *GetDocStructureResultResponseBody {
	s.Data = v
	return s
}

func (s *GetDocStructureResultResponseBody) SetMessage(v string) *GetDocStructureResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetDocStructureResultResponseBody) SetRequestId(v string) *GetDocStructureResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocStructureResultResponseBody) SetStatus(v string) *GetDocStructureResultResponseBody {
	s.Status = &v
	return s
}

type GetDocStructureResultResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDocStructureResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDocStructureResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDocStructureResultResponse) GoString() string {
	return s.String()
}

func (s *GetDocStructureResultResponse) SetHeaders(v map[string]*string) *GetDocStructureResultResponse {
	s.Headers = v
	return s
}

func (s *GetDocStructureResultResponse) SetStatusCode(v int32) *GetDocStructureResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocStructureResultResponse) SetBody(v *GetDocStructureResultResponseBody) *GetDocStructureResultResponse {
	s.Body = v
	return s
}

type GetDocumentCompareResultRequest struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetDocumentCompareResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentCompareResultRequest) GoString() string {
	return s.String()
}

func (s *GetDocumentCompareResultRequest) SetId(v string) *GetDocumentCompareResultRequest {
	s.Id = &v
	return s
}

type GetDocumentCompareResultResponseBody struct {
	Code      *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Completed *bool       `json:"Completed,omitempty" xml:"Completed,omitempty"`
	Data      interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string     `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetDocumentCompareResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentCompareResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocumentCompareResultResponseBody) SetCode(v string) *GetDocumentCompareResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetDocumentCompareResultResponseBody) SetCompleted(v bool) *GetDocumentCompareResultResponseBody {
	s.Completed = &v
	return s
}

func (s *GetDocumentCompareResultResponseBody) SetData(v interface{}) *GetDocumentCompareResultResponseBody {
	s.Data = v
	return s
}

func (s *GetDocumentCompareResultResponseBody) SetMessage(v string) *GetDocumentCompareResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetDocumentCompareResultResponseBody) SetRequestId(v string) *GetDocumentCompareResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocumentCompareResultResponseBody) SetStatus(v string) *GetDocumentCompareResultResponseBody {
	s.Status = &v
	return s
}

type GetDocumentCompareResultResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDocumentCompareResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDocumentCompareResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentCompareResultResponse) GoString() string {
	return s.String()
}

func (s *GetDocumentCompareResultResponse) SetHeaders(v map[string]*string) *GetDocumentCompareResultResponse {
	s.Headers = v
	return s
}

func (s *GetDocumentCompareResultResponse) SetStatusCode(v int32) *GetDocumentCompareResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocumentCompareResultResponse) SetBody(v *GetDocumentCompareResultResponseBody) *GetDocumentCompareResultResponse {
	s.Body = v
	return s
}

type GetDocumentConvertResultRequest struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetDocumentConvertResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentConvertResultRequest) GoString() string {
	return s.String()
}

func (s *GetDocumentConvertResultRequest) SetId(v string) *GetDocumentConvertResultRequest {
	s.Id = &v
	return s
}

type GetDocumentConvertResultResponseBody struct {
	Code      *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Completed *bool                                       `json:"Completed,omitempty" xml:"Completed,omitempty"`
	Data      []*GetDocumentConvertResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string                                     `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetDocumentConvertResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentConvertResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocumentConvertResultResponseBody) SetCode(v string) *GetDocumentConvertResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetDocumentConvertResultResponseBody) SetCompleted(v bool) *GetDocumentConvertResultResponseBody {
	s.Completed = &v
	return s
}

func (s *GetDocumentConvertResultResponseBody) SetData(v []*GetDocumentConvertResultResponseBodyData) *GetDocumentConvertResultResponseBody {
	s.Data = v
	return s
}

func (s *GetDocumentConvertResultResponseBody) SetMessage(v string) *GetDocumentConvertResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetDocumentConvertResultResponseBody) SetRequestId(v string) *GetDocumentConvertResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocumentConvertResultResponseBody) SetStatus(v string) *GetDocumentConvertResultResponseBody {
	s.Status = &v
	return s
}

type GetDocumentConvertResultResponseBodyData struct {
	Md5  *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Size *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Url  *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetDocumentConvertResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentConvertResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDocumentConvertResultResponseBodyData) SetMd5(v string) *GetDocumentConvertResultResponseBodyData {
	s.Md5 = &v
	return s
}

func (s *GetDocumentConvertResultResponseBodyData) SetSize(v int64) *GetDocumentConvertResultResponseBodyData {
	s.Size = &v
	return s
}

func (s *GetDocumentConvertResultResponseBodyData) SetType(v string) *GetDocumentConvertResultResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetDocumentConvertResultResponseBodyData) SetUrl(v string) *GetDocumentConvertResultResponseBodyData {
	s.Url = &v
	return s
}

type GetDocumentConvertResultResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDocumentConvertResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDocumentConvertResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentConvertResultResponse) GoString() string {
	return s.String()
}

func (s *GetDocumentConvertResultResponse) SetHeaders(v map[string]*string) *GetDocumentConvertResultResponse {
	s.Headers = v
	return s
}

func (s *GetDocumentConvertResultResponse) SetStatusCode(v int32) *GetDocumentConvertResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocumentConvertResultResponse) SetBody(v *GetDocumentConvertResultResponseBody) *GetDocumentConvertResultResponse {
	s.Body = v
	return s
}

type GetDocumentExtractResultRequest struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetDocumentExtractResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentExtractResultRequest) GoString() string {
	return s.String()
}

func (s *GetDocumentExtractResultRequest) SetId(v string) *GetDocumentExtractResultRequest {
	s.Id = &v
	return s
}

type GetDocumentExtractResultResponseBody struct {
	Code      *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Completed *bool                  `json:"Completed,omitempty" xml:"Completed,omitempty"`
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string                `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetDocumentExtractResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentExtractResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocumentExtractResultResponseBody) SetCode(v string) *GetDocumentExtractResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetDocumentExtractResultResponseBody) SetCompleted(v bool) *GetDocumentExtractResultResponseBody {
	s.Completed = &v
	return s
}

func (s *GetDocumentExtractResultResponseBody) SetData(v map[string]interface{}) *GetDocumentExtractResultResponseBody {
	s.Data = v
	return s
}

func (s *GetDocumentExtractResultResponseBody) SetMessage(v string) *GetDocumentExtractResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetDocumentExtractResultResponseBody) SetRequestId(v string) *GetDocumentExtractResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocumentExtractResultResponseBody) SetStatus(v string) *GetDocumentExtractResultResponseBody {
	s.Status = &v
	return s
}

type GetDocumentExtractResultResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDocumentExtractResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDocumentExtractResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentExtractResultResponse) GoString() string {
	return s.String()
}

func (s *GetDocumentExtractResultResponse) SetHeaders(v map[string]*string) *GetDocumentExtractResultResponse {
	s.Headers = v
	return s
}

func (s *GetDocumentExtractResultResponse) SetStatusCode(v int32) *GetDocumentExtractResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocumentExtractResultResponse) SetBody(v *GetDocumentExtractResultResponseBody) *GetDocumentExtractResultResponse {
	s.Body = v
	return s
}

type GetTableUnderstandingResultRequest struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetTableUnderstandingResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTableUnderstandingResultRequest) GoString() string {
	return s.String()
}

func (s *GetTableUnderstandingResultRequest) SetId(v string) *GetTableUnderstandingResultRequest {
	s.Id = &v
	return s
}

type GetTableUnderstandingResultResponseBody struct {
	Code      *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Completed *bool                  `json:"Completed,omitempty" xml:"Completed,omitempty"`
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string                `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetTableUnderstandingResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTableUnderstandingResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableUnderstandingResultResponseBody) SetCode(v string) *GetTableUnderstandingResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetTableUnderstandingResultResponseBody) SetCompleted(v bool) *GetTableUnderstandingResultResponseBody {
	s.Completed = &v
	return s
}

func (s *GetTableUnderstandingResultResponseBody) SetData(v map[string]interface{}) *GetTableUnderstandingResultResponseBody {
	s.Data = v
	return s
}

func (s *GetTableUnderstandingResultResponseBody) SetMessage(v string) *GetTableUnderstandingResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetTableUnderstandingResultResponseBody) SetRequestId(v string) *GetTableUnderstandingResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableUnderstandingResultResponseBody) SetStatus(v string) *GetTableUnderstandingResultResponseBody {
	s.Status = &v
	return s
}

type GetTableUnderstandingResultResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTableUnderstandingResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTableUnderstandingResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTableUnderstandingResultResponse) GoString() string {
	return s.String()
}

func (s *GetTableUnderstandingResultResponse) SetHeaders(v map[string]*string) *GetTableUnderstandingResultResponse {
	s.Headers = v
	return s
}

func (s *GetTableUnderstandingResultResponse) SetStatusCode(v int32) *GetTableUnderstandingResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTableUnderstandingResultResponse) SetBody(v *GetTableUnderstandingResultResponseBody) *GetTableUnderstandingResultResponse {
	s.Body = v
	return s
}

type SubmitConvertImageToExcelJobRequest struct {
	ForceMergeExcel    *bool     `json:"ForceMergeExcel,omitempty" xml:"ForceMergeExcel,omitempty"`
	ImageNameExtension *string   `json:"ImageNameExtension,omitempty" xml:"ImageNameExtension,omitempty"`
	ImageNames         []*string `json:"ImageNames,omitempty" xml:"ImageNames,omitempty" type:"Repeated"`
	ImageUrls          []*string `json:"ImageUrls,omitempty" xml:"ImageUrls,omitempty" type:"Repeated"`
}

func (s SubmitConvertImageToExcelJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitConvertImageToExcelJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitConvertImageToExcelJobRequest) SetForceMergeExcel(v bool) *SubmitConvertImageToExcelJobRequest {
	s.ForceMergeExcel = &v
	return s
}

func (s *SubmitConvertImageToExcelJobRequest) SetImageNameExtension(v string) *SubmitConvertImageToExcelJobRequest {
	s.ImageNameExtension = &v
	return s
}

func (s *SubmitConvertImageToExcelJobRequest) SetImageNames(v []*string) *SubmitConvertImageToExcelJobRequest {
	s.ImageNames = v
	return s
}

func (s *SubmitConvertImageToExcelJobRequest) SetImageUrls(v []*string) *SubmitConvertImageToExcelJobRequest {
	s.ImageUrls = v
	return s
}

type SubmitConvertImageToExcelJobShrinkRequest struct {
	ForceMergeExcel    *bool   `json:"ForceMergeExcel,omitempty" xml:"ForceMergeExcel,omitempty"`
	ImageNameExtension *string `json:"ImageNameExtension,omitempty" xml:"ImageNameExtension,omitempty"`
	ImageNamesShrink   *string `json:"ImageNames,omitempty" xml:"ImageNames,omitempty"`
	ImageUrlsShrink    *string `json:"ImageUrls,omitempty" xml:"ImageUrls,omitempty"`
}

func (s SubmitConvertImageToExcelJobShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitConvertImageToExcelJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitConvertImageToExcelJobShrinkRequest) SetForceMergeExcel(v bool) *SubmitConvertImageToExcelJobShrinkRequest {
	s.ForceMergeExcel = &v
	return s
}

func (s *SubmitConvertImageToExcelJobShrinkRequest) SetImageNameExtension(v string) *SubmitConvertImageToExcelJobShrinkRequest {
	s.ImageNameExtension = &v
	return s
}

func (s *SubmitConvertImageToExcelJobShrinkRequest) SetImageNamesShrink(v string) *SubmitConvertImageToExcelJobShrinkRequest {
	s.ImageNamesShrink = &v
	return s
}

func (s *SubmitConvertImageToExcelJobShrinkRequest) SetImageUrlsShrink(v string) *SubmitConvertImageToExcelJobShrinkRequest {
	s.ImageUrlsShrink = &v
	return s
}

type SubmitConvertImageToExcelJobResponseBody struct {
	Code      *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SubmitConvertImageToExcelJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitConvertImageToExcelJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitConvertImageToExcelJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitConvertImageToExcelJobResponseBody) SetCode(v string) *SubmitConvertImageToExcelJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitConvertImageToExcelJobResponseBody) SetData(v *SubmitConvertImageToExcelJobResponseBodyData) *SubmitConvertImageToExcelJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitConvertImageToExcelJobResponseBody) SetMessage(v string) *SubmitConvertImageToExcelJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitConvertImageToExcelJobResponseBody) SetRequestId(v string) *SubmitConvertImageToExcelJobResponseBody {
	s.RequestId = &v
	return s
}

type SubmitConvertImageToExcelJobResponseBodyData struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SubmitConvertImageToExcelJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitConvertImageToExcelJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitConvertImageToExcelJobResponseBodyData) SetId(v string) *SubmitConvertImageToExcelJobResponseBodyData {
	s.Id = &v
	return s
}

type SubmitConvertImageToExcelJobResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitConvertImageToExcelJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitConvertImageToExcelJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitConvertImageToExcelJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitConvertImageToExcelJobResponse) SetHeaders(v map[string]*string) *SubmitConvertImageToExcelJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitConvertImageToExcelJobResponse) SetStatusCode(v int32) *SubmitConvertImageToExcelJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitConvertImageToExcelJobResponse) SetBody(v *SubmitConvertImageToExcelJobResponseBody) *SubmitConvertImageToExcelJobResponse {
	s.Body = v
	return s
}

type SubmitConvertImageToPdfJobRequest struct {
	ImageNameExtension *string   `json:"ImageNameExtension,omitempty" xml:"ImageNameExtension,omitempty"`
	ImageNames         []*string `json:"ImageNames,omitempty" xml:"ImageNames,omitempty" type:"Repeated"`
	ImageUrls          []*string `json:"ImageUrls,omitempty" xml:"ImageUrls,omitempty" type:"Repeated"`
}

func (s SubmitConvertImageToPdfJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitConvertImageToPdfJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitConvertImageToPdfJobRequest) SetImageNameExtension(v string) *SubmitConvertImageToPdfJobRequest {
	s.ImageNameExtension = &v
	return s
}

func (s *SubmitConvertImageToPdfJobRequest) SetImageNames(v []*string) *SubmitConvertImageToPdfJobRequest {
	s.ImageNames = v
	return s
}

func (s *SubmitConvertImageToPdfJobRequest) SetImageUrls(v []*string) *SubmitConvertImageToPdfJobRequest {
	s.ImageUrls = v
	return s
}

type SubmitConvertImageToPdfJobShrinkRequest struct {
	ImageNameExtension *string `json:"ImageNameExtension,omitempty" xml:"ImageNameExtension,omitempty"`
	ImageNamesShrink   *string `json:"ImageNames,omitempty" xml:"ImageNames,omitempty"`
	ImageUrlsShrink    *string `json:"ImageUrls,omitempty" xml:"ImageUrls,omitempty"`
}

func (s SubmitConvertImageToPdfJobShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitConvertImageToPdfJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitConvertImageToPdfJobShrinkRequest) SetImageNameExtension(v string) *SubmitConvertImageToPdfJobShrinkRequest {
	s.ImageNameExtension = &v
	return s
}

func (s *SubmitConvertImageToPdfJobShrinkRequest) SetImageNamesShrink(v string) *SubmitConvertImageToPdfJobShrinkRequest {
	s.ImageNamesShrink = &v
	return s
}

func (s *SubmitConvertImageToPdfJobShrinkRequest) SetImageUrlsShrink(v string) *SubmitConvertImageToPdfJobShrinkRequest {
	s.ImageUrlsShrink = &v
	return s
}

type SubmitConvertImageToPdfJobResponseBody struct {
	Code      *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SubmitConvertImageToPdfJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitConvertImageToPdfJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitConvertImageToPdfJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitConvertImageToPdfJobResponseBody) SetCode(v string) *SubmitConvertImageToPdfJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitConvertImageToPdfJobResponseBody) SetData(v *SubmitConvertImageToPdfJobResponseBodyData) *SubmitConvertImageToPdfJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitConvertImageToPdfJobResponseBody) SetMessage(v string) *SubmitConvertImageToPdfJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitConvertImageToPdfJobResponseBody) SetRequestId(v string) *SubmitConvertImageToPdfJobResponseBody {
	s.RequestId = &v
	return s
}

type SubmitConvertImageToPdfJobResponseBodyData struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SubmitConvertImageToPdfJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitConvertImageToPdfJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitConvertImageToPdfJobResponseBodyData) SetId(v string) *SubmitConvertImageToPdfJobResponseBodyData {
	s.Id = &v
	return s
}

type SubmitConvertImageToPdfJobResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitConvertImageToPdfJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitConvertImageToPdfJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitConvertImageToPdfJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitConvertImageToPdfJobResponse) SetHeaders(v map[string]*string) *SubmitConvertImageToPdfJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitConvertImageToPdfJobResponse) SetStatusCode(v int32) *SubmitConvertImageToPdfJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitConvertImageToPdfJobResponse) SetBody(v *SubmitConvertImageToPdfJobResponseBody) *SubmitConvertImageToPdfJobResponse {
	s.Body = v
	return s
}

type SubmitConvertImageToWordJobRequest struct {
	ImageNameExtension *string   `json:"ImageNameExtension,omitempty" xml:"ImageNameExtension,omitempty"`
	ImageNames         []*string `json:"ImageNames,omitempty" xml:"ImageNames,omitempty" type:"Repeated"`
	ImageUrls          []*string `json:"ImageUrls,omitempty" xml:"ImageUrls,omitempty" type:"Repeated"`
}

func (s SubmitConvertImageToWordJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitConvertImageToWordJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitConvertImageToWordJobRequest) SetImageNameExtension(v string) *SubmitConvertImageToWordJobRequest {
	s.ImageNameExtension = &v
	return s
}

func (s *SubmitConvertImageToWordJobRequest) SetImageNames(v []*string) *SubmitConvertImageToWordJobRequest {
	s.ImageNames = v
	return s
}

func (s *SubmitConvertImageToWordJobRequest) SetImageUrls(v []*string) *SubmitConvertImageToWordJobRequest {
	s.ImageUrls = v
	return s
}

type SubmitConvertImageToWordJobShrinkRequest struct {
	ImageNameExtension *string `json:"ImageNameExtension,omitempty" xml:"ImageNameExtension,omitempty"`
	ImageNamesShrink   *string `json:"ImageNames,omitempty" xml:"ImageNames,omitempty"`
	ImageUrlsShrink    *string `json:"ImageUrls,omitempty" xml:"ImageUrls,omitempty"`
}

func (s SubmitConvertImageToWordJobShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitConvertImageToWordJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitConvertImageToWordJobShrinkRequest) SetImageNameExtension(v string) *SubmitConvertImageToWordJobShrinkRequest {
	s.ImageNameExtension = &v
	return s
}

func (s *SubmitConvertImageToWordJobShrinkRequest) SetImageNamesShrink(v string) *SubmitConvertImageToWordJobShrinkRequest {
	s.ImageNamesShrink = &v
	return s
}

func (s *SubmitConvertImageToWordJobShrinkRequest) SetImageUrlsShrink(v string) *SubmitConvertImageToWordJobShrinkRequest {
	s.ImageUrlsShrink = &v
	return s
}

type SubmitConvertImageToWordJobResponseBody struct {
	Code      *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SubmitConvertImageToWordJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitConvertImageToWordJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitConvertImageToWordJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitConvertImageToWordJobResponseBody) SetCode(v string) *SubmitConvertImageToWordJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitConvertImageToWordJobResponseBody) SetData(v *SubmitConvertImageToWordJobResponseBodyData) *SubmitConvertImageToWordJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitConvertImageToWordJobResponseBody) SetMessage(v string) *SubmitConvertImageToWordJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitConvertImageToWordJobResponseBody) SetRequestId(v string) *SubmitConvertImageToWordJobResponseBody {
	s.RequestId = &v
	return s
}

type SubmitConvertImageToWordJobResponseBodyData struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SubmitConvertImageToWordJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitConvertImageToWordJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitConvertImageToWordJobResponseBodyData) SetId(v string) *SubmitConvertImageToWordJobResponseBodyData {
	s.Id = &v
	return s
}

type SubmitConvertImageToWordJobResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitConvertImageToWordJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitConvertImageToWordJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitConvertImageToWordJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitConvertImageToWordJobResponse) SetHeaders(v map[string]*string) *SubmitConvertImageToWordJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitConvertImageToWordJobResponse) SetStatusCode(v int32) *SubmitConvertImageToWordJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitConvertImageToWordJobResponse) SetBody(v *SubmitConvertImageToWordJobResponseBody) *SubmitConvertImageToWordJobResponse {
	s.Body = v
	return s
}

type SubmitConvertPdfToExcelJobRequest struct {
	FileName              *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileUrl               *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	ForceExportInnerImage *bool   `json:"ForceExportInnerImage,omitempty" xml:"ForceExportInnerImage,omitempty"`
	ForceMergeExcel       *bool   `json:"ForceMergeExcel,omitempty" xml:"ForceMergeExcel,omitempty"`
}

func (s SubmitConvertPdfToExcelJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitConvertPdfToExcelJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitConvertPdfToExcelJobRequest) SetFileName(v string) *SubmitConvertPdfToExcelJobRequest {
	s.FileName = &v
	return s
}

func (s *SubmitConvertPdfToExcelJobRequest) SetFileUrl(v string) *SubmitConvertPdfToExcelJobRequest {
	s.FileUrl = &v
	return s
}

func (s *SubmitConvertPdfToExcelJobRequest) SetForceExportInnerImage(v bool) *SubmitConvertPdfToExcelJobRequest {
	s.ForceExportInnerImage = &v
	return s
}

func (s *SubmitConvertPdfToExcelJobRequest) SetForceMergeExcel(v bool) *SubmitConvertPdfToExcelJobRequest {
	s.ForceMergeExcel = &v
	return s
}

type SubmitConvertPdfToExcelJobAdvanceRequest struct {
	FileName              *string   `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileUrlObject         io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	ForceExportInnerImage *bool     `json:"ForceExportInnerImage,omitempty" xml:"ForceExportInnerImage,omitempty"`
	ForceMergeExcel       *bool     `json:"ForceMergeExcel,omitempty" xml:"ForceMergeExcel,omitempty"`
}

func (s SubmitConvertPdfToExcelJobAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitConvertPdfToExcelJobAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SubmitConvertPdfToExcelJobAdvanceRequest) SetFileName(v string) *SubmitConvertPdfToExcelJobAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *SubmitConvertPdfToExcelJobAdvanceRequest) SetFileUrlObject(v io.Reader) *SubmitConvertPdfToExcelJobAdvanceRequest {
	s.FileUrlObject = v
	return s
}

func (s *SubmitConvertPdfToExcelJobAdvanceRequest) SetForceExportInnerImage(v bool) *SubmitConvertPdfToExcelJobAdvanceRequest {
	s.ForceExportInnerImage = &v
	return s
}

func (s *SubmitConvertPdfToExcelJobAdvanceRequest) SetForceMergeExcel(v bool) *SubmitConvertPdfToExcelJobAdvanceRequest {
	s.ForceMergeExcel = &v
	return s
}

type SubmitConvertPdfToExcelJobResponseBody struct {
	Code      *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SubmitConvertPdfToExcelJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitConvertPdfToExcelJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitConvertPdfToExcelJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitConvertPdfToExcelJobResponseBody) SetCode(v string) *SubmitConvertPdfToExcelJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitConvertPdfToExcelJobResponseBody) SetData(v *SubmitConvertPdfToExcelJobResponseBodyData) *SubmitConvertPdfToExcelJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitConvertPdfToExcelJobResponseBody) SetMessage(v string) *SubmitConvertPdfToExcelJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitConvertPdfToExcelJobResponseBody) SetRequestId(v string) *SubmitConvertPdfToExcelJobResponseBody {
	s.RequestId = &v
	return s
}

type SubmitConvertPdfToExcelJobResponseBodyData struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SubmitConvertPdfToExcelJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitConvertPdfToExcelJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitConvertPdfToExcelJobResponseBodyData) SetId(v string) *SubmitConvertPdfToExcelJobResponseBodyData {
	s.Id = &v
	return s
}

type SubmitConvertPdfToExcelJobResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitConvertPdfToExcelJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitConvertPdfToExcelJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitConvertPdfToExcelJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitConvertPdfToExcelJobResponse) SetHeaders(v map[string]*string) *SubmitConvertPdfToExcelJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitConvertPdfToExcelJobResponse) SetStatusCode(v int32) *SubmitConvertPdfToExcelJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitConvertPdfToExcelJobResponse) SetBody(v *SubmitConvertPdfToExcelJobResponseBody) *SubmitConvertPdfToExcelJobResponse {
	s.Body = v
	return s
}

type SubmitConvertPdfToImageJobRequest struct {
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileUrl  *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s SubmitConvertPdfToImageJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitConvertPdfToImageJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitConvertPdfToImageJobRequest) SetFileName(v string) *SubmitConvertPdfToImageJobRequest {
	s.FileName = &v
	return s
}

func (s *SubmitConvertPdfToImageJobRequest) SetFileUrl(v string) *SubmitConvertPdfToImageJobRequest {
	s.FileUrl = &v
	return s
}

type SubmitConvertPdfToImageJobAdvanceRequest struct {
	FileName      *string   `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileUrlObject io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s SubmitConvertPdfToImageJobAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitConvertPdfToImageJobAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SubmitConvertPdfToImageJobAdvanceRequest) SetFileName(v string) *SubmitConvertPdfToImageJobAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *SubmitConvertPdfToImageJobAdvanceRequest) SetFileUrlObject(v io.Reader) *SubmitConvertPdfToImageJobAdvanceRequest {
	s.FileUrlObject = v
	return s
}

type SubmitConvertPdfToImageJobResponseBody struct {
	Code      *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SubmitConvertPdfToImageJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitConvertPdfToImageJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitConvertPdfToImageJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitConvertPdfToImageJobResponseBody) SetCode(v string) *SubmitConvertPdfToImageJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitConvertPdfToImageJobResponseBody) SetData(v *SubmitConvertPdfToImageJobResponseBodyData) *SubmitConvertPdfToImageJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitConvertPdfToImageJobResponseBody) SetMessage(v string) *SubmitConvertPdfToImageJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitConvertPdfToImageJobResponseBody) SetRequestId(v string) *SubmitConvertPdfToImageJobResponseBody {
	s.RequestId = &v
	return s
}

type SubmitConvertPdfToImageJobResponseBodyData struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SubmitConvertPdfToImageJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitConvertPdfToImageJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitConvertPdfToImageJobResponseBodyData) SetId(v string) *SubmitConvertPdfToImageJobResponseBodyData {
	s.Id = &v
	return s
}

type SubmitConvertPdfToImageJobResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitConvertPdfToImageJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitConvertPdfToImageJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitConvertPdfToImageJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitConvertPdfToImageJobResponse) SetHeaders(v map[string]*string) *SubmitConvertPdfToImageJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitConvertPdfToImageJobResponse) SetStatusCode(v int32) *SubmitConvertPdfToImageJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitConvertPdfToImageJobResponse) SetBody(v *SubmitConvertPdfToImageJobResponseBody) *SubmitConvertPdfToImageJobResponse {
	s.Body = v
	return s
}

type SubmitConvertPdfToWordJobRequest struct {
	FileName              *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileUrl               *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	ForceExportInnerImage *bool   `json:"ForceExportInnerImage,omitempty" xml:"ForceExportInnerImage,omitempty"`
}

func (s SubmitConvertPdfToWordJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitConvertPdfToWordJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitConvertPdfToWordJobRequest) SetFileName(v string) *SubmitConvertPdfToWordJobRequest {
	s.FileName = &v
	return s
}

func (s *SubmitConvertPdfToWordJobRequest) SetFileUrl(v string) *SubmitConvertPdfToWordJobRequest {
	s.FileUrl = &v
	return s
}

func (s *SubmitConvertPdfToWordJobRequest) SetForceExportInnerImage(v bool) *SubmitConvertPdfToWordJobRequest {
	s.ForceExportInnerImage = &v
	return s
}

type SubmitConvertPdfToWordJobAdvanceRequest struct {
	FileName              *string   `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileUrlObject         io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	ForceExportInnerImage *bool     `json:"ForceExportInnerImage,omitempty" xml:"ForceExportInnerImage,omitempty"`
}

func (s SubmitConvertPdfToWordJobAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitConvertPdfToWordJobAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SubmitConvertPdfToWordJobAdvanceRequest) SetFileName(v string) *SubmitConvertPdfToWordJobAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *SubmitConvertPdfToWordJobAdvanceRequest) SetFileUrlObject(v io.Reader) *SubmitConvertPdfToWordJobAdvanceRequest {
	s.FileUrlObject = v
	return s
}

func (s *SubmitConvertPdfToWordJobAdvanceRequest) SetForceExportInnerImage(v bool) *SubmitConvertPdfToWordJobAdvanceRequest {
	s.ForceExportInnerImage = &v
	return s
}

type SubmitConvertPdfToWordJobResponseBody struct {
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SubmitConvertPdfToWordJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitConvertPdfToWordJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitConvertPdfToWordJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitConvertPdfToWordJobResponseBody) SetCode(v string) *SubmitConvertPdfToWordJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitConvertPdfToWordJobResponseBody) SetData(v *SubmitConvertPdfToWordJobResponseBodyData) *SubmitConvertPdfToWordJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitConvertPdfToWordJobResponseBody) SetMessage(v string) *SubmitConvertPdfToWordJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitConvertPdfToWordJobResponseBody) SetRequestId(v string) *SubmitConvertPdfToWordJobResponseBody {
	s.RequestId = &v
	return s
}

type SubmitConvertPdfToWordJobResponseBodyData struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SubmitConvertPdfToWordJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitConvertPdfToWordJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitConvertPdfToWordJobResponseBodyData) SetId(v string) *SubmitConvertPdfToWordJobResponseBodyData {
	s.Id = &v
	return s
}

type SubmitConvertPdfToWordJobResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitConvertPdfToWordJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitConvertPdfToWordJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitConvertPdfToWordJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitConvertPdfToWordJobResponse) SetHeaders(v map[string]*string) *SubmitConvertPdfToWordJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitConvertPdfToWordJobResponse) SetStatusCode(v int32) *SubmitConvertPdfToWordJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitConvertPdfToWordJobResponse) SetBody(v *SubmitConvertPdfToWordJobResponseBody) *SubmitConvertPdfToWordJobResponse {
	s.Body = v
	return s
}

type SubmitDocStructureJobRequest struct {
	FileName          *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileNameExtension *string `json:"FileNameExtension,omitempty" xml:"FileNameExtension,omitempty"`
	FileUrl           *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s SubmitDocStructureJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocStructureJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitDocStructureJobRequest) SetFileName(v string) *SubmitDocStructureJobRequest {
	s.FileName = &v
	return s
}

func (s *SubmitDocStructureJobRequest) SetFileNameExtension(v string) *SubmitDocStructureJobRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitDocStructureJobRequest) SetFileUrl(v string) *SubmitDocStructureJobRequest {
	s.FileUrl = &v
	return s
}

type SubmitDocStructureJobAdvanceRequest struct {
	FileName          *string   `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileNameExtension *string   `json:"FileNameExtension,omitempty" xml:"FileNameExtension,omitempty"`
	FileUrlObject     io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s SubmitDocStructureJobAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocStructureJobAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SubmitDocStructureJobAdvanceRequest) SetFileName(v string) *SubmitDocStructureJobAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *SubmitDocStructureJobAdvanceRequest) SetFileNameExtension(v string) *SubmitDocStructureJobAdvanceRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitDocStructureJobAdvanceRequest) SetFileUrlObject(v io.Reader) *SubmitDocStructureJobAdvanceRequest {
	s.FileUrlObject = v
	return s
}

type SubmitDocStructureJobResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SubmitDocStructureJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitDocStructureJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocStructureJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitDocStructureJobResponseBody) SetCode(v string) *SubmitDocStructureJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitDocStructureJobResponseBody) SetData(v *SubmitDocStructureJobResponseBodyData) *SubmitDocStructureJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitDocStructureJobResponseBody) SetMessage(v string) *SubmitDocStructureJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitDocStructureJobResponseBody) SetRequestId(v string) *SubmitDocStructureJobResponseBody {
	s.RequestId = &v
	return s
}

type SubmitDocStructureJobResponseBodyData struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SubmitDocStructureJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocStructureJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitDocStructureJobResponseBodyData) SetId(v string) *SubmitDocStructureJobResponseBodyData {
	s.Id = &v
	return s
}

type SubmitDocStructureJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitDocStructureJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitDocStructureJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocStructureJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitDocStructureJobResponse) SetHeaders(v map[string]*string) *SubmitDocStructureJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitDocStructureJobResponse) SetStatusCode(v int32) *SubmitDocStructureJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitDocStructureJobResponse) SetBody(v *SubmitDocStructureJobResponseBody) *SubmitDocStructureJobResponse {
	s.Body = v
	return s
}

type SubmitDocumentCompareJobRequest struct {
	CompareFileName *string `json:"CompareFileName,omitempty" xml:"CompareFileName,omitempty"`
	CompareFileUrl  *string `json:"CompareFileUrl,omitempty" xml:"CompareFileUrl,omitempty"`
	OriginFileName  *string `json:"OriginFileName,omitempty" xml:"OriginFileName,omitempty"`
	OriginFileUrl   *string `json:"OriginFileUrl,omitempty" xml:"OriginFileUrl,omitempty"`
}

func (s SubmitDocumentCompareJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocumentCompareJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitDocumentCompareJobRequest) SetCompareFileName(v string) *SubmitDocumentCompareJobRequest {
	s.CompareFileName = &v
	return s
}

func (s *SubmitDocumentCompareJobRequest) SetCompareFileUrl(v string) *SubmitDocumentCompareJobRequest {
	s.CompareFileUrl = &v
	return s
}

func (s *SubmitDocumentCompareJobRequest) SetOriginFileName(v string) *SubmitDocumentCompareJobRequest {
	s.OriginFileName = &v
	return s
}

func (s *SubmitDocumentCompareJobRequest) SetOriginFileUrl(v string) *SubmitDocumentCompareJobRequest {
	s.OriginFileUrl = &v
	return s
}

type SubmitDocumentCompareJobResponseBody struct {
	Code      *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SubmitDocumentCompareJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitDocumentCompareJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocumentCompareJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitDocumentCompareJobResponseBody) SetCode(v string) *SubmitDocumentCompareJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitDocumentCompareJobResponseBody) SetData(v *SubmitDocumentCompareJobResponseBodyData) *SubmitDocumentCompareJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitDocumentCompareJobResponseBody) SetMessage(v string) *SubmitDocumentCompareJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitDocumentCompareJobResponseBody) SetRequestId(v string) *SubmitDocumentCompareJobResponseBody {
	s.RequestId = &v
	return s
}

type SubmitDocumentCompareJobResponseBodyData struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SubmitDocumentCompareJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocumentCompareJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitDocumentCompareJobResponseBodyData) SetId(v string) *SubmitDocumentCompareJobResponseBodyData {
	s.Id = &v
	return s
}

type SubmitDocumentCompareJobResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitDocumentCompareJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitDocumentCompareJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocumentCompareJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitDocumentCompareJobResponse) SetHeaders(v map[string]*string) *SubmitDocumentCompareJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitDocumentCompareJobResponse) SetStatusCode(v int32) *SubmitDocumentCompareJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitDocumentCompareJobResponse) SetBody(v *SubmitDocumentCompareJobResponseBody) *SubmitDocumentCompareJobResponse {
	s.Body = v
	return s
}

type SubmitDocumentExtractJobRequest struct {
	FileName          *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileNameExtension *string `json:"FileNameExtension,omitempty" xml:"FileNameExtension,omitempty"`
	FileUrl           *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s SubmitDocumentExtractJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocumentExtractJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitDocumentExtractJobRequest) SetFileName(v string) *SubmitDocumentExtractJobRequest {
	s.FileName = &v
	return s
}

func (s *SubmitDocumentExtractJobRequest) SetFileNameExtension(v string) *SubmitDocumentExtractJobRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitDocumentExtractJobRequest) SetFileUrl(v string) *SubmitDocumentExtractJobRequest {
	s.FileUrl = &v
	return s
}

type SubmitDocumentExtractJobAdvanceRequest struct {
	FileName          *string   `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileNameExtension *string   `json:"FileNameExtension,omitempty" xml:"FileNameExtension,omitempty"`
	FileUrlObject     io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s SubmitDocumentExtractJobAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocumentExtractJobAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SubmitDocumentExtractJobAdvanceRequest) SetFileName(v string) *SubmitDocumentExtractJobAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *SubmitDocumentExtractJobAdvanceRequest) SetFileNameExtension(v string) *SubmitDocumentExtractJobAdvanceRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitDocumentExtractJobAdvanceRequest) SetFileUrlObject(v io.Reader) *SubmitDocumentExtractJobAdvanceRequest {
	s.FileUrlObject = v
	return s
}

type SubmitDocumentExtractJobResponseBody struct {
	Code      *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SubmitDocumentExtractJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitDocumentExtractJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocumentExtractJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitDocumentExtractJobResponseBody) SetCode(v string) *SubmitDocumentExtractJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitDocumentExtractJobResponseBody) SetData(v *SubmitDocumentExtractJobResponseBodyData) *SubmitDocumentExtractJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitDocumentExtractJobResponseBody) SetMessage(v string) *SubmitDocumentExtractJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitDocumentExtractJobResponseBody) SetRequestId(v string) *SubmitDocumentExtractJobResponseBody {
	s.RequestId = &v
	return s
}

type SubmitDocumentExtractJobResponseBodyData struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SubmitDocumentExtractJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocumentExtractJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitDocumentExtractJobResponseBodyData) SetId(v string) *SubmitDocumentExtractJobResponseBodyData {
	s.Id = &v
	return s
}

type SubmitDocumentExtractJobResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitDocumentExtractJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitDocumentExtractJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocumentExtractJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitDocumentExtractJobResponse) SetHeaders(v map[string]*string) *SubmitDocumentExtractJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitDocumentExtractJobResponse) SetStatusCode(v int32) *SubmitDocumentExtractJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitDocumentExtractJobResponse) SetBody(v *SubmitDocumentExtractJobResponseBody) *SubmitDocumentExtractJobResponse {
	s.Body = v
	return s
}

type SubmitTableUnderstandingJobRequest struct {
	FileName          *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileNameExtension *string `json:"FileNameExtension,omitempty" xml:"FileNameExtension,omitempty"`
	FileUrl           *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s SubmitTableUnderstandingJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitTableUnderstandingJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitTableUnderstandingJobRequest) SetFileName(v string) *SubmitTableUnderstandingJobRequest {
	s.FileName = &v
	return s
}

func (s *SubmitTableUnderstandingJobRequest) SetFileNameExtension(v string) *SubmitTableUnderstandingJobRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitTableUnderstandingJobRequest) SetFileUrl(v string) *SubmitTableUnderstandingJobRequest {
	s.FileUrl = &v
	return s
}

type SubmitTableUnderstandingJobAdvanceRequest struct {
	FileName          *string   `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileNameExtension *string   `json:"FileNameExtension,omitempty" xml:"FileNameExtension,omitempty"`
	FileUrlObject     io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s SubmitTableUnderstandingJobAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitTableUnderstandingJobAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SubmitTableUnderstandingJobAdvanceRequest) SetFileName(v string) *SubmitTableUnderstandingJobAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *SubmitTableUnderstandingJobAdvanceRequest) SetFileNameExtension(v string) *SubmitTableUnderstandingJobAdvanceRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitTableUnderstandingJobAdvanceRequest) SetFileUrlObject(v io.Reader) *SubmitTableUnderstandingJobAdvanceRequest {
	s.FileUrlObject = v
	return s
}

type SubmitTableUnderstandingJobResponseBody struct {
	Code      *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SubmitTableUnderstandingJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitTableUnderstandingJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitTableUnderstandingJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitTableUnderstandingJobResponseBody) SetCode(v string) *SubmitTableUnderstandingJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitTableUnderstandingJobResponseBody) SetData(v *SubmitTableUnderstandingJobResponseBodyData) *SubmitTableUnderstandingJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitTableUnderstandingJobResponseBody) SetMessage(v string) *SubmitTableUnderstandingJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitTableUnderstandingJobResponseBody) SetRequestId(v string) *SubmitTableUnderstandingJobResponseBody {
	s.RequestId = &v
	return s
}

type SubmitTableUnderstandingJobResponseBodyData struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SubmitTableUnderstandingJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitTableUnderstandingJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitTableUnderstandingJobResponseBodyData) SetId(v string) *SubmitTableUnderstandingJobResponseBodyData {
	s.Id = &v
	return s
}

type SubmitTableUnderstandingJobResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitTableUnderstandingJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitTableUnderstandingJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitTableUnderstandingJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitTableUnderstandingJobResponse) SetHeaders(v map[string]*string) *SubmitTableUnderstandingJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitTableUnderstandingJobResponse) SetStatusCode(v int32) *SubmitTableUnderstandingJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitTableUnderstandingJobResponse) SetBody(v *SubmitTableUnderstandingJobResponseBody) *SubmitTableUnderstandingJobResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-1":              tea.String("docmind-api.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("docmind-api.aliyuncs.com"),
		"ap-south-1":                  tea.String("docmind-api.aliyuncs.com"),
		"ap-southeast-1":              tea.String("docmind-api.aliyuncs.com"),
		"ap-southeast-2":              tea.String("docmind-api.aliyuncs.com"),
		"ap-southeast-3":              tea.String("docmind-api.aliyuncs.com"),
		"ap-southeast-5":              tea.String("docmind-api.aliyuncs.com"),
		"cn-beijing":                  tea.String("docmind-api.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("docmind-api.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("docmind-api.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("docmind-api.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("docmind-api.aliyuncs.com"),
		"cn-chengdu":                  tea.String("docmind-api.aliyuncs.com"),
		"cn-edge-1":                   tea.String("docmind-api.aliyuncs.com"),
		"cn-fujian":                   tea.String("docmind-api.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("docmind-api.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("docmind-api.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("docmind-api.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("docmind-api.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("docmind-api.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("docmind-api.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("docmind-api.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("docmind-api.aliyuncs.com"),
		"cn-hongkong":                 tea.String("docmind-api.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("docmind-api.aliyuncs.com"),
		"cn-huhehaote":                tea.String("docmind-api.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("docmind-api.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("docmind-api.aliyuncs.com"),
		"cn-qingdao":                  tea.String("docmind-api.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("docmind-api.aliyuncs.com"),
		"cn-shanghai":                 tea.String("docmind-api.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("docmind-api.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("docmind-api.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("docmind-api.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("docmind-api.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("docmind-api.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("docmind-api.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("docmind-api.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("docmind-api.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("docmind-api.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("docmind-api.aliyuncs.com"),
		"cn-wuhan":                    tea.String("docmind-api.aliyuncs.com"),
		"cn-wulanchabu":               tea.String("docmind-api.aliyuncs.com"),
		"cn-yushanfang":               tea.String("docmind-api.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("docmind-api.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("docmind-api.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("docmind-api.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("docmind-api.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("docmind-api.aliyuncs.com"),
		"eu-central-1":                tea.String("docmind-api.aliyuncs.com"),
		"eu-west-1":                   tea.String("docmind-api.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("docmind-api.aliyuncs.com"),
		"me-east-1":                   tea.String("docmind-api.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("docmind-api.aliyuncs.com"),
		"us-east-1":                   tea.String("docmind-api.aliyuncs.com"),
		"us-west-1":                   tea.String("docmind-api.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("docmind-api"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDocStructureResultWithOptions(request *GetDocStructureResultRequest, runtime *util.RuntimeOptions) (_result *GetDocStructureResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDocStructureResult"),
		Version:     tea.String("2022-07-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDocStructureResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDocStructureResult(request *GetDocStructureResultRequest) (_result *GetDocStructureResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDocStructureResultResponse{}
	_body, _err := client.GetDocStructureResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDocumentCompareResultWithOptions(request *GetDocumentCompareResultRequest, runtime *util.RuntimeOptions) (_result *GetDocumentCompareResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDocumentCompareResult"),
		Version:     tea.String("2022-07-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDocumentCompareResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDocumentCompareResult(request *GetDocumentCompareResultRequest) (_result *GetDocumentCompareResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDocumentCompareResultResponse{}
	_body, _err := client.GetDocumentCompareResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDocumentConvertResultWithOptions(request *GetDocumentConvertResultRequest, runtime *util.RuntimeOptions) (_result *GetDocumentConvertResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDocumentConvertResult"),
		Version:     tea.String("2022-07-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDocumentConvertResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDocumentConvertResult(request *GetDocumentConvertResultRequest) (_result *GetDocumentConvertResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDocumentConvertResultResponse{}
	_body, _err := client.GetDocumentConvertResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDocumentExtractResultWithOptions(request *GetDocumentExtractResultRequest, runtime *util.RuntimeOptions) (_result *GetDocumentExtractResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDocumentExtractResult"),
		Version:     tea.String("2022-07-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDocumentExtractResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDocumentExtractResult(request *GetDocumentExtractResultRequest) (_result *GetDocumentExtractResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDocumentExtractResultResponse{}
	_body, _err := client.GetDocumentExtractResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTableUnderstandingResultWithOptions(request *GetTableUnderstandingResultRequest, runtime *util.RuntimeOptions) (_result *GetTableUnderstandingResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTableUnderstandingResult"),
		Version:     tea.String("2022-07-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTableUnderstandingResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTableUnderstandingResult(request *GetTableUnderstandingResultRequest) (_result *GetTableUnderstandingResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTableUnderstandingResultResponse{}
	_body, _err := client.GetTableUnderstandingResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitConvertImageToExcelJobWithOptions(tmpReq *SubmitConvertImageToExcelJobRequest, runtime *util.RuntimeOptions) (_result *SubmitConvertImageToExcelJobResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SubmitConvertImageToExcelJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ImageNames)) {
		request.ImageNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ImageNames, tea.String("ImageNames"), tea.String("simple"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ImageUrls)) {
		request.ImageUrlsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ImageUrls, tea.String("ImageUrls"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ForceMergeExcel)) {
		query["ForceMergeExcel"] = request.ForceMergeExcel
	}

	if !tea.BoolValue(util.IsUnset(request.ImageNameExtension)) {
		query["ImageNameExtension"] = request.ImageNameExtension
	}

	if !tea.BoolValue(util.IsUnset(request.ImageNamesShrink)) {
		query["ImageNames"] = request.ImageNamesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrlsShrink)) {
		query["ImageUrls"] = request.ImageUrlsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitConvertImageToExcelJob"),
		Version:     tea.String("2022-07-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitConvertImageToExcelJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitConvertImageToExcelJob(request *SubmitConvertImageToExcelJobRequest) (_result *SubmitConvertImageToExcelJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitConvertImageToExcelJobResponse{}
	_body, _err := client.SubmitConvertImageToExcelJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitConvertImageToPdfJobWithOptions(tmpReq *SubmitConvertImageToPdfJobRequest, runtime *util.RuntimeOptions) (_result *SubmitConvertImageToPdfJobResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SubmitConvertImageToPdfJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ImageNames)) {
		request.ImageNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ImageNames, tea.String("ImageNames"), tea.String("simple"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ImageUrls)) {
		request.ImageUrlsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ImageUrls, tea.String("ImageUrls"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageNameExtension)) {
		query["ImageNameExtension"] = request.ImageNameExtension
	}

	if !tea.BoolValue(util.IsUnset(request.ImageNamesShrink)) {
		query["ImageNames"] = request.ImageNamesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrlsShrink)) {
		query["ImageUrls"] = request.ImageUrlsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitConvertImageToPdfJob"),
		Version:     tea.String("2022-07-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitConvertImageToPdfJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitConvertImageToPdfJob(request *SubmitConvertImageToPdfJobRequest) (_result *SubmitConvertImageToPdfJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitConvertImageToPdfJobResponse{}
	_body, _err := client.SubmitConvertImageToPdfJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitConvertImageToWordJobWithOptions(tmpReq *SubmitConvertImageToWordJobRequest, runtime *util.RuntimeOptions) (_result *SubmitConvertImageToWordJobResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SubmitConvertImageToWordJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ImageNames)) {
		request.ImageNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ImageNames, tea.String("ImageNames"), tea.String("simple"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ImageUrls)) {
		request.ImageUrlsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ImageUrls, tea.String("ImageUrls"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageNameExtension)) {
		query["ImageNameExtension"] = request.ImageNameExtension
	}

	if !tea.BoolValue(util.IsUnset(request.ImageNamesShrink)) {
		query["ImageNames"] = request.ImageNamesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrlsShrink)) {
		query["ImageUrls"] = request.ImageUrlsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitConvertImageToWordJob"),
		Version:     tea.String("2022-07-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitConvertImageToWordJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitConvertImageToWordJob(request *SubmitConvertImageToWordJobRequest) (_result *SubmitConvertImageToWordJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitConvertImageToWordJobResponse{}
	_body, _err := client.SubmitConvertImageToWordJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitConvertPdfToExcelJobWithOptions(request *SubmitConvertPdfToExcelJobRequest, runtime *util.RuntimeOptions) (_result *SubmitConvertPdfToExcelJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["FileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.FileUrl)) {
		query["FileUrl"] = request.FileUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ForceExportInnerImage)) {
		query["ForceExportInnerImage"] = request.ForceExportInnerImage
	}

	if !tea.BoolValue(util.IsUnset(request.ForceMergeExcel)) {
		query["ForceMergeExcel"] = request.ForceMergeExcel
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitConvertPdfToExcelJob"),
		Version:     tea.String("2022-07-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitConvertPdfToExcelJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitConvertPdfToExcelJob(request *SubmitConvertPdfToExcelJobRequest) (_result *SubmitConvertPdfToExcelJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitConvertPdfToExcelJobResponse{}
	_body, _err := client.SubmitConvertPdfToExcelJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitConvertPdfToExcelJobAdvance(request *SubmitConvertPdfToExcelJobAdvanceRequest, runtime *util.RuntimeOptions) (_result *SubmitConvertPdfToExcelJobResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("docmind-api"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	submitConvertPdfToExcelJobReq := &SubmitConvertPdfToExcelJobRequest{}
	openapiutil.Convert(request, submitConvertPdfToExcelJobReq)
	if !tea.BoolValue(util.IsUnset(request.FileUrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.FileUrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		submitConvertPdfToExcelJobReq.FileUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	submitConvertPdfToExcelJobResp, _err := client.SubmitConvertPdfToExcelJobWithOptions(submitConvertPdfToExcelJobReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = submitConvertPdfToExcelJobResp
	return _result, _err
}

func (client *Client) SubmitConvertPdfToImageJobWithOptions(request *SubmitConvertPdfToImageJobRequest, runtime *util.RuntimeOptions) (_result *SubmitConvertPdfToImageJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["FileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.FileUrl)) {
		query["FileUrl"] = request.FileUrl
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitConvertPdfToImageJob"),
		Version:     tea.String("2022-07-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitConvertPdfToImageJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitConvertPdfToImageJob(request *SubmitConvertPdfToImageJobRequest) (_result *SubmitConvertPdfToImageJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitConvertPdfToImageJobResponse{}
	_body, _err := client.SubmitConvertPdfToImageJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitConvertPdfToImageJobAdvance(request *SubmitConvertPdfToImageJobAdvanceRequest, runtime *util.RuntimeOptions) (_result *SubmitConvertPdfToImageJobResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("docmind-api"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	submitConvertPdfToImageJobReq := &SubmitConvertPdfToImageJobRequest{}
	openapiutil.Convert(request, submitConvertPdfToImageJobReq)
	if !tea.BoolValue(util.IsUnset(request.FileUrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.FileUrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		submitConvertPdfToImageJobReq.FileUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	submitConvertPdfToImageJobResp, _err := client.SubmitConvertPdfToImageJobWithOptions(submitConvertPdfToImageJobReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = submitConvertPdfToImageJobResp
	return _result, _err
}

func (client *Client) SubmitConvertPdfToWordJobWithOptions(request *SubmitConvertPdfToWordJobRequest, runtime *util.RuntimeOptions) (_result *SubmitConvertPdfToWordJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["FileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.FileUrl)) {
		query["FileUrl"] = request.FileUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ForceExportInnerImage)) {
		query["ForceExportInnerImage"] = request.ForceExportInnerImage
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitConvertPdfToWordJob"),
		Version:     tea.String("2022-07-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitConvertPdfToWordJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitConvertPdfToWordJob(request *SubmitConvertPdfToWordJobRequest) (_result *SubmitConvertPdfToWordJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitConvertPdfToWordJobResponse{}
	_body, _err := client.SubmitConvertPdfToWordJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitConvertPdfToWordJobAdvance(request *SubmitConvertPdfToWordJobAdvanceRequest, runtime *util.RuntimeOptions) (_result *SubmitConvertPdfToWordJobResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("docmind-api"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	submitConvertPdfToWordJobReq := &SubmitConvertPdfToWordJobRequest{}
	openapiutil.Convert(request, submitConvertPdfToWordJobReq)
	if !tea.BoolValue(util.IsUnset(request.FileUrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.FileUrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		submitConvertPdfToWordJobReq.FileUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	submitConvertPdfToWordJobResp, _err := client.SubmitConvertPdfToWordJobWithOptions(submitConvertPdfToWordJobReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = submitConvertPdfToWordJobResp
	return _result, _err
}

func (client *Client) SubmitDocStructureJobWithOptions(request *SubmitDocStructureJobRequest, runtime *util.RuntimeOptions) (_result *SubmitDocStructureJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["FileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.FileNameExtension)) {
		query["FileNameExtension"] = request.FileNameExtension
	}

	if !tea.BoolValue(util.IsUnset(request.FileUrl)) {
		query["FileUrl"] = request.FileUrl
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitDocStructureJob"),
		Version:     tea.String("2022-07-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitDocStructureJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitDocStructureJob(request *SubmitDocStructureJobRequest) (_result *SubmitDocStructureJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitDocStructureJobResponse{}
	_body, _err := client.SubmitDocStructureJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitDocStructureJobAdvance(request *SubmitDocStructureJobAdvanceRequest, runtime *util.RuntimeOptions) (_result *SubmitDocStructureJobResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("docmind-api"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	submitDocStructureJobReq := &SubmitDocStructureJobRequest{}
	openapiutil.Convert(request, submitDocStructureJobReq)
	if !tea.BoolValue(util.IsUnset(request.FileUrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.FileUrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		submitDocStructureJobReq.FileUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	submitDocStructureJobResp, _err := client.SubmitDocStructureJobWithOptions(submitDocStructureJobReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = submitDocStructureJobResp
	return _result, _err
}

func (client *Client) SubmitDocumentCompareJobWithOptions(request *SubmitDocumentCompareJobRequest, runtime *util.RuntimeOptions) (_result *SubmitDocumentCompareJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CompareFileName)) {
		query["CompareFileName"] = request.CompareFileName
	}

	if !tea.BoolValue(util.IsUnset(request.CompareFileUrl)) {
		query["CompareFileUrl"] = request.CompareFileUrl
	}

	if !tea.BoolValue(util.IsUnset(request.OriginFileName)) {
		query["OriginFileName"] = request.OriginFileName
	}

	if !tea.BoolValue(util.IsUnset(request.OriginFileUrl)) {
		query["OriginFileUrl"] = request.OriginFileUrl
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitDocumentCompareJob"),
		Version:     tea.String("2022-07-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitDocumentCompareJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitDocumentCompareJob(request *SubmitDocumentCompareJobRequest) (_result *SubmitDocumentCompareJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitDocumentCompareJobResponse{}
	_body, _err := client.SubmitDocumentCompareJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitDocumentExtractJobWithOptions(request *SubmitDocumentExtractJobRequest, runtime *util.RuntimeOptions) (_result *SubmitDocumentExtractJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["FileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.FileNameExtension)) {
		query["FileNameExtension"] = request.FileNameExtension
	}

	if !tea.BoolValue(util.IsUnset(request.FileUrl)) {
		query["FileUrl"] = request.FileUrl
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitDocumentExtractJob"),
		Version:     tea.String("2022-07-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitDocumentExtractJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitDocumentExtractJob(request *SubmitDocumentExtractJobRequest) (_result *SubmitDocumentExtractJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitDocumentExtractJobResponse{}
	_body, _err := client.SubmitDocumentExtractJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitDocumentExtractJobAdvance(request *SubmitDocumentExtractJobAdvanceRequest, runtime *util.RuntimeOptions) (_result *SubmitDocumentExtractJobResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("docmind-api"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	submitDocumentExtractJobReq := &SubmitDocumentExtractJobRequest{}
	openapiutil.Convert(request, submitDocumentExtractJobReq)
	if !tea.BoolValue(util.IsUnset(request.FileUrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.FileUrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		submitDocumentExtractJobReq.FileUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	submitDocumentExtractJobResp, _err := client.SubmitDocumentExtractJobWithOptions(submitDocumentExtractJobReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = submitDocumentExtractJobResp
	return _result, _err
}

func (client *Client) SubmitTableUnderstandingJobWithOptions(request *SubmitTableUnderstandingJobRequest, runtime *util.RuntimeOptions) (_result *SubmitTableUnderstandingJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["FileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.FileNameExtension)) {
		query["FileNameExtension"] = request.FileNameExtension
	}

	if !tea.BoolValue(util.IsUnset(request.FileUrl)) {
		query["FileUrl"] = request.FileUrl
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitTableUnderstandingJob"),
		Version:     tea.String("2022-07-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitTableUnderstandingJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitTableUnderstandingJob(request *SubmitTableUnderstandingJobRequest) (_result *SubmitTableUnderstandingJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitTableUnderstandingJobResponse{}
	_body, _err := client.SubmitTableUnderstandingJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitTableUnderstandingJobAdvance(request *SubmitTableUnderstandingJobAdvanceRequest, runtime *util.RuntimeOptions) (_result *SubmitTableUnderstandingJobResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("docmind-api"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	submitTableUnderstandingJobReq := &SubmitTableUnderstandingJobRequest{}
	openapiutil.Convert(request, submitTableUnderstandingJobReq)
	if !tea.BoolValue(util.IsUnset(request.FileUrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.FileUrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		submitTableUnderstandingJobReq.FileUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	submitTableUnderstandingJobResp, _err := client.SubmitTableUnderstandingJobWithOptions(submitTableUnderstandingJobReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = submitTableUnderstandingJobResp
	return _result, _err
}
