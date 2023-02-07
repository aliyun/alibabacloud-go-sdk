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

type GetSingleDocumentExtractResultRequest struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetSingleDocumentExtractResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSingleDocumentExtractResultRequest) GoString() string {
	return s.String()
}

func (s *GetSingleDocumentExtractResultRequest) SetId(v string) *GetSingleDocumentExtractResultRequest {
	s.Id = &v
	return s
}

type GetSingleDocumentExtractResultResponseBody struct {
	Code      *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Completed *bool       `json:"Completed,omitempty" xml:"Completed,omitempty"`
	Data      interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string     `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetSingleDocumentExtractResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSingleDocumentExtractResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetSingleDocumentExtractResultResponseBody) SetCode(v string) *GetSingleDocumentExtractResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetSingleDocumentExtractResultResponseBody) SetCompleted(v bool) *GetSingleDocumentExtractResultResponseBody {
	s.Completed = &v
	return s
}

func (s *GetSingleDocumentExtractResultResponseBody) SetData(v interface{}) *GetSingleDocumentExtractResultResponseBody {
	s.Data = v
	return s
}

func (s *GetSingleDocumentExtractResultResponseBody) SetMessage(v string) *GetSingleDocumentExtractResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetSingleDocumentExtractResultResponseBody) SetRequestId(v string) *GetSingleDocumentExtractResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSingleDocumentExtractResultResponseBody) SetStatus(v string) *GetSingleDocumentExtractResultResponseBody {
	s.Status = &v
	return s
}

type GetSingleDocumentExtractResultResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSingleDocumentExtractResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSingleDocumentExtractResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSingleDocumentExtractResultResponse) GoString() string {
	return s.String()
}

func (s *GetSingleDocumentExtractResultResponse) SetHeaders(v map[string]*string) *GetSingleDocumentExtractResultResponse {
	s.Headers = v
	return s
}

func (s *GetSingleDocumentExtractResultResponse) SetStatusCode(v int32) *GetSingleDocumentExtractResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSingleDocumentExtractResultResponse) SetBody(v *GetSingleDocumentExtractResultResponseBody) *GetSingleDocumentExtractResultResponse {
	s.Body = v
	return s
}

type SubmitAirWaybillExtractJobRequest struct {
	FileName          *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileNameExtension *string `json:"FileNameExtension,omitempty" xml:"FileNameExtension,omitempty"`
	FileUrl           *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s SubmitAirWaybillExtractJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitAirWaybillExtractJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitAirWaybillExtractJobRequest) SetFileName(v string) *SubmitAirWaybillExtractJobRequest {
	s.FileName = &v
	return s
}

func (s *SubmitAirWaybillExtractJobRequest) SetFileNameExtension(v string) *SubmitAirWaybillExtractJobRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitAirWaybillExtractJobRequest) SetFileUrl(v string) *SubmitAirWaybillExtractJobRequest {
	s.FileUrl = &v
	return s
}

type SubmitAirWaybillExtractJobAdvanceRequest struct {
	FileName          *string   `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileNameExtension *string   `json:"FileNameExtension,omitempty" xml:"FileNameExtension,omitempty"`
	FileUrlObject     io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s SubmitAirWaybillExtractJobAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitAirWaybillExtractJobAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SubmitAirWaybillExtractJobAdvanceRequest) SetFileName(v string) *SubmitAirWaybillExtractJobAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *SubmitAirWaybillExtractJobAdvanceRequest) SetFileNameExtension(v string) *SubmitAirWaybillExtractJobAdvanceRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitAirWaybillExtractJobAdvanceRequest) SetFileUrlObject(v io.Reader) *SubmitAirWaybillExtractJobAdvanceRequest {
	s.FileUrlObject = v
	return s
}

type SubmitAirWaybillExtractJobResponseBody struct {
	Code      *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SubmitAirWaybillExtractJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitAirWaybillExtractJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitAirWaybillExtractJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAirWaybillExtractJobResponseBody) SetCode(v string) *SubmitAirWaybillExtractJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitAirWaybillExtractJobResponseBody) SetData(v *SubmitAirWaybillExtractJobResponseBodyData) *SubmitAirWaybillExtractJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitAirWaybillExtractJobResponseBody) SetMessage(v string) *SubmitAirWaybillExtractJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitAirWaybillExtractJobResponseBody) SetRequestId(v string) *SubmitAirWaybillExtractJobResponseBody {
	s.RequestId = &v
	return s
}

type SubmitAirWaybillExtractJobResponseBodyData struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SubmitAirWaybillExtractJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitAirWaybillExtractJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitAirWaybillExtractJobResponseBodyData) SetId(v string) *SubmitAirWaybillExtractJobResponseBodyData {
	s.Id = &v
	return s
}

type SubmitAirWaybillExtractJobResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitAirWaybillExtractJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitAirWaybillExtractJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitAirWaybillExtractJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitAirWaybillExtractJobResponse) SetHeaders(v map[string]*string) *SubmitAirWaybillExtractJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitAirWaybillExtractJobResponse) SetStatusCode(v int32) *SubmitAirWaybillExtractJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAirWaybillExtractJobResponse) SetBody(v *SubmitAirWaybillExtractJobResponseBody) *SubmitAirWaybillExtractJobResponse {
	s.Body = v
	return s
}

type SubmitBillOfLadingExtractJobRequest struct {
	FileName          *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileNameExtension *string `json:"FileNameExtension,omitempty" xml:"FileNameExtension,omitempty"`
	FileUrl           *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s SubmitBillOfLadingExtractJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitBillOfLadingExtractJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitBillOfLadingExtractJobRequest) SetFileName(v string) *SubmitBillOfLadingExtractJobRequest {
	s.FileName = &v
	return s
}

func (s *SubmitBillOfLadingExtractJobRequest) SetFileNameExtension(v string) *SubmitBillOfLadingExtractJobRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitBillOfLadingExtractJobRequest) SetFileUrl(v string) *SubmitBillOfLadingExtractJobRequest {
	s.FileUrl = &v
	return s
}

type SubmitBillOfLadingExtractJobAdvanceRequest struct {
	FileName          *string   `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileNameExtension *string   `json:"FileNameExtension,omitempty" xml:"FileNameExtension,omitempty"`
	FileUrlObject     io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s SubmitBillOfLadingExtractJobAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitBillOfLadingExtractJobAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SubmitBillOfLadingExtractJobAdvanceRequest) SetFileName(v string) *SubmitBillOfLadingExtractJobAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *SubmitBillOfLadingExtractJobAdvanceRequest) SetFileNameExtension(v string) *SubmitBillOfLadingExtractJobAdvanceRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitBillOfLadingExtractJobAdvanceRequest) SetFileUrlObject(v io.Reader) *SubmitBillOfLadingExtractJobAdvanceRequest {
	s.FileUrlObject = v
	return s
}

type SubmitBillOfLadingExtractJobResponseBody struct {
	Code      *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SubmitBillOfLadingExtractJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitBillOfLadingExtractJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitBillOfLadingExtractJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitBillOfLadingExtractJobResponseBody) SetCode(v string) *SubmitBillOfLadingExtractJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitBillOfLadingExtractJobResponseBody) SetData(v *SubmitBillOfLadingExtractJobResponseBodyData) *SubmitBillOfLadingExtractJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitBillOfLadingExtractJobResponseBody) SetMessage(v string) *SubmitBillOfLadingExtractJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitBillOfLadingExtractJobResponseBody) SetRequestId(v string) *SubmitBillOfLadingExtractJobResponseBody {
	s.RequestId = &v
	return s
}

type SubmitBillOfLadingExtractJobResponseBodyData struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SubmitBillOfLadingExtractJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitBillOfLadingExtractJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitBillOfLadingExtractJobResponseBodyData) SetId(v string) *SubmitBillOfLadingExtractJobResponseBodyData {
	s.Id = &v
	return s
}

type SubmitBillOfLadingExtractJobResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitBillOfLadingExtractJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitBillOfLadingExtractJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitBillOfLadingExtractJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitBillOfLadingExtractJobResponse) SetHeaders(v map[string]*string) *SubmitBillOfLadingExtractJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitBillOfLadingExtractJobResponse) SetStatusCode(v int32) *SubmitBillOfLadingExtractJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitBillOfLadingExtractJobResponse) SetBody(v *SubmitBillOfLadingExtractJobResponseBody) *SubmitBillOfLadingExtractJobResponse {
	s.Body = v
	return s
}

type SubmitExportDeclarationSheetExtractJobRequest struct {
	FileName          *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileNameExtension *string `json:"FileNameExtension,omitempty" xml:"FileNameExtension,omitempty"`
	FileUrl           *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s SubmitExportDeclarationSheetExtractJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitExportDeclarationSheetExtractJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitExportDeclarationSheetExtractJobRequest) SetFileName(v string) *SubmitExportDeclarationSheetExtractJobRequest {
	s.FileName = &v
	return s
}

func (s *SubmitExportDeclarationSheetExtractJobRequest) SetFileNameExtension(v string) *SubmitExportDeclarationSheetExtractJobRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitExportDeclarationSheetExtractJobRequest) SetFileUrl(v string) *SubmitExportDeclarationSheetExtractJobRequest {
	s.FileUrl = &v
	return s
}

type SubmitExportDeclarationSheetExtractJobAdvanceRequest struct {
	FileName          *string   `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileNameExtension *string   `json:"FileNameExtension,omitempty" xml:"FileNameExtension,omitempty"`
	FileUrlObject     io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s SubmitExportDeclarationSheetExtractJobAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitExportDeclarationSheetExtractJobAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SubmitExportDeclarationSheetExtractJobAdvanceRequest) SetFileName(v string) *SubmitExportDeclarationSheetExtractJobAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *SubmitExportDeclarationSheetExtractJobAdvanceRequest) SetFileNameExtension(v string) *SubmitExportDeclarationSheetExtractJobAdvanceRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitExportDeclarationSheetExtractJobAdvanceRequest) SetFileUrlObject(v io.Reader) *SubmitExportDeclarationSheetExtractJobAdvanceRequest {
	s.FileUrlObject = v
	return s
}

type SubmitExportDeclarationSheetExtractJobResponseBody struct {
	Code      *string                                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SubmitExportDeclarationSheetExtractJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitExportDeclarationSheetExtractJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitExportDeclarationSheetExtractJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitExportDeclarationSheetExtractJobResponseBody) SetCode(v string) *SubmitExportDeclarationSheetExtractJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitExportDeclarationSheetExtractJobResponseBody) SetData(v *SubmitExportDeclarationSheetExtractJobResponseBodyData) *SubmitExportDeclarationSheetExtractJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitExportDeclarationSheetExtractJobResponseBody) SetMessage(v string) *SubmitExportDeclarationSheetExtractJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitExportDeclarationSheetExtractJobResponseBody) SetRequestId(v string) *SubmitExportDeclarationSheetExtractJobResponseBody {
	s.RequestId = &v
	return s
}

type SubmitExportDeclarationSheetExtractJobResponseBodyData struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SubmitExportDeclarationSheetExtractJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitExportDeclarationSheetExtractJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitExportDeclarationSheetExtractJobResponseBodyData) SetId(v string) *SubmitExportDeclarationSheetExtractJobResponseBodyData {
	s.Id = &v
	return s
}

type SubmitExportDeclarationSheetExtractJobResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitExportDeclarationSheetExtractJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitExportDeclarationSheetExtractJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitExportDeclarationSheetExtractJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitExportDeclarationSheetExtractJobResponse) SetHeaders(v map[string]*string) *SubmitExportDeclarationSheetExtractJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitExportDeclarationSheetExtractJobResponse) SetStatusCode(v int32) *SubmitExportDeclarationSheetExtractJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitExportDeclarationSheetExtractJobResponse) SetBody(v *SubmitExportDeclarationSheetExtractJobResponseBody) *SubmitExportDeclarationSheetExtractJobResponse {
	s.Body = v
	return s
}

type SubmitInvoiceExtractJobRequest struct {
	FileName          *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileNameExtension *string `json:"FileNameExtension,omitempty" xml:"FileNameExtension,omitempty"`
	FileUrl           *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s SubmitInvoiceExtractJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitInvoiceExtractJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitInvoiceExtractJobRequest) SetFileName(v string) *SubmitInvoiceExtractJobRequest {
	s.FileName = &v
	return s
}

func (s *SubmitInvoiceExtractJobRequest) SetFileNameExtension(v string) *SubmitInvoiceExtractJobRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitInvoiceExtractJobRequest) SetFileUrl(v string) *SubmitInvoiceExtractJobRequest {
	s.FileUrl = &v
	return s
}

type SubmitInvoiceExtractJobAdvanceRequest struct {
	FileName          *string   `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileNameExtension *string   `json:"FileNameExtension,omitempty" xml:"FileNameExtension,omitempty"`
	FileUrlObject     io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s SubmitInvoiceExtractJobAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitInvoiceExtractJobAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SubmitInvoiceExtractJobAdvanceRequest) SetFileName(v string) *SubmitInvoiceExtractJobAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *SubmitInvoiceExtractJobAdvanceRequest) SetFileNameExtension(v string) *SubmitInvoiceExtractJobAdvanceRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitInvoiceExtractJobAdvanceRequest) SetFileUrlObject(v io.Reader) *SubmitInvoiceExtractJobAdvanceRequest {
	s.FileUrlObject = v
	return s
}

type SubmitInvoiceExtractJobResponseBody struct {
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SubmitInvoiceExtractJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitInvoiceExtractJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitInvoiceExtractJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitInvoiceExtractJobResponseBody) SetCode(v string) *SubmitInvoiceExtractJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitInvoiceExtractJobResponseBody) SetData(v *SubmitInvoiceExtractJobResponseBodyData) *SubmitInvoiceExtractJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitInvoiceExtractJobResponseBody) SetMessage(v string) *SubmitInvoiceExtractJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitInvoiceExtractJobResponseBody) SetRequestId(v string) *SubmitInvoiceExtractJobResponseBody {
	s.RequestId = &v
	return s
}

type SubmitInvoiceExtractJobResponseBodyData struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SubmitInvoiceExtractJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitInvoiceExtractJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitInvoiceExtractJobResponseBodyData) SetId(v string) *SubmitInvoiceExtractJobResponseBodyData {
	s.Id = &v
	return s
}

type SubmitInvoiceExtractJobResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitInvoiceExtractJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitInvoiceExtractJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitInvoiceExtractJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitInvoiceExtractJobResponse) SetHeaders(v map[string]*string) *SubmitInvoiceExtractJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitInvoiceExtractJobResponse) SetStatusCode(v int32) *SubmitInvoiceExtractJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitInvoiceExtractJobResponse) SetBody(v *SubmitInvoiceExtractJobResponseBody) *SubmitInvoiceExtractJobResponse {
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

func (client *Client) GetSingleDocumentExtractResultWithOptions(request *GetSingleDocumentExtractResultRequest, runtime *util.RuntimeOptions) (_result *GetSingleDocumentExtractResultResponse, _err error) {
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
		Action:      tea.String("GetSingleDocumentExtractResult"),
		Version:     tea.String("2022-07-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSingleDocumentExtractResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSingleDocumentExtractResult(request *GetSingleDocumentExtractResultRequest) (_result *GetSingleDocumentExtractResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSingleDocumentExtractResultResponse{}
	_body, _err := client.GetSingleDocumentExtractResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitAirWaybillExtractJobWithOptions(request *SubmitAirWaybillExtractJobRequest, runtime *util.RuntimeOptions) (_result *SubmitAirWaybillExtractJobResponse, _err error) {
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
		Action:      tea.String("SubmitAirWaybillExtractJob"),
		Version:     tea.String("2022-07-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitAirWaybillExtractJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitAirWaybillExtractJob(request *SubmitAirWaybillExtractJobRequest) (_result *SubmitAirWaybillExtractJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitAirWaybillExtractJobResponse{}
	_body, _err := client.SubmitAirWaybillExtractJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitAirWaybillExtractJobAdvance(request *SubmitAirWaybillExtractJobAdvanceRequest, runtime *util.RuntimeOptions) (_result *SubmitAirWaybillExtractJobResponse, _err error) {
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
	submitAirWaybillExtractJobReq := &SubmitAirWaybillExtractJobRequest{}
	openapiutil.Convert(request, submitAirWaybillExtractJobReq)
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
		submitAirWaybillExtractJobReq.FileUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	submitAirWaybillExtractJobResp, _err := client.SubmitAirWaybillExtractJobWithOptions(submitAirWaybillExtractJobReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = submitAirWaybillExtractJobResp
	return _result, _err
}

func (client *Client) SubmitBillOfLadingExtractJobWithOptions(request *SubmitBillOfLadingExtractJobRequest, runtime *util.RuntimeOptions) (_result *SubmitBillOfLadingExtractJobResponse, _err error) {
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
		Action:      tea.String("SubmitBillOfLadingExtractJob"),
		Version:     tea.String("2022-07-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitBillOfLadingExtractJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitBillOfLadingExtractJob(request *SubmitBillOfLadingExtractJobRequest) (_result *SubmitBillOfLadingExtractJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitBillOfLadingExtractJobResponse{}
	_body, _err := client.SubmitBillOfLadingExtractJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitBillOfLadingExtractJobAdvance(request *SubmitBillOfLadingExtractJobAdvanceRequest, runtime *util.RuntimeOptions) (_result *SubmitBillOfLadingExtractJobResponse, _err error) {
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
	submitBillOfLadingExtractJobReq := &SubmitBillOfLadingExtractJobRequest{}
	openapiutil.Convert(request, submitBillOfLadingExtractJobReq)
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
		submitBillOfLadingExtractJobReq.FileUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	submitBillOfLadingExtractJobResp, _err := client.SubmitBillOfLadingExtractJobWithOptions(submitBillOfLadingExtractJobReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = submitBillOfLadingExtractJobResp
	return _result, _err
}

func (client *Client) SubmitExportDeclarationSheetExtractJobWithOptions(request *SubmitExportDeclarationSheetExtractJobRequest, runtime *util.RuntimeOptions) (_result *SubmitExportDeclarationSheetExtractJobResponse, _err error) {
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
		Action:      tea.String("SubmitExportDeclarationSheetExtractJob"),
		Version:     tea.String("2022-07-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitExportDeclarationSheetExtractJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitExportDeclarationSheetExtractJob(request *SubmitExportDeclarationSheetExtractJobRequest) (_result *SubmitExportDeclarationSheetExtractJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitExportDeclarationSheetExtractJobResponse{}
	_body, _err := client.SubmitExportDeclarationSheetExtractJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitExportDeclarationSheetExtractJobAdvance(request *SubmitExportDeclarationSheetExtractJobAdvanceRequest, runtime *util.RuntimeOptions) (_result *SubmitExportDeclarationSheetExtractJobResponse, _err error) {
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
	submitExportDeclarationSheetExtractJobReq := &SubmitExportDeclarationSheetExtractJobRequest{}
	openapiutil.Convert(request, submitExportDeclarationSheetExtractJobReq)
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
		submitExportDeclarationSheetExtractJobReq.FileUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	submitExportDeclarationSheetExtractJobResp, _err := client.SubmitExportDeclarationSheetExtractJobWithOptions(submitExportDeclarationSheetExtractJobReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = submitExportDeclarationSheetExtractJobResp
	return _result, _err
}

func (client *Client) SubmitInvoiceExtractJobWithOptions(request *SubmitInvoiceExtractJobRequest, runtime *util.RuntimeOptions) (_result *SubmitInvoiceExtractJobResponse, _err error) {
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
		Action:      tea.String("SubmitInvoiceExtractJob"),
		Version:     tea.String("2022-07-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitInvoiceExtractJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitInvoiceExtractJob(request *SubmitInvoiceExtractJobRequest) (_result *SubmitInvoiceExtractJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitInvoiceExtractJobResponse{}
	_body, _err := client.SubmitInvoiceExtractJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitInvoiceExtractJobAdvance(request *SubmitInvoiceExtractJobAdvanceRequest, runtime *util.RuntimeOptions) (_result *SubmitInvoiceExtractJobResponse, _err error) {
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
	submitInvoiceExtractJobReq := &SubmitInvoiceExtractJobRequest{}
	openapiutil.Convert(request, submitInvoiceExtractJobReq)
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
		submitInvoiceExtractJobReq.FileUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	submitInvoiceExtractJobResp, _err := client.SubmitInvoiceExtractJobWithOptions(submitInvoiceExtractJobReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = submitInvoiceExtractJobResp
	return _result, _err
}
