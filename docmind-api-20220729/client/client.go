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

type SubmitBookingNoteExtractJobRequest struct {
	FileName          *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileNameExtension *string `json:"FileNameExtension,omitempty" xml:"FileNameExtension,omitempty"`
	FileUrl           *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s SubmitBookingNoteExtractJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitBookingNoteExtractJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitBookingNoteExtractJobRequest) SetFileName(v string) *SubmitBookingNoteExtractJobRequest {
	s.FileName = &v
	return s
}

func (s *SubmitBookingNoteExtractJobRequest) SetFileNameExtension(v string) *SubmitBookingNoteExtractJobRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitBookingNoteExtractJobRequest) SetFileUrl(v string) *SubmitBookingNoteExtractJobRequest {
	s.FileUrl = &v
	return s
}

type SubmitBookingNoteExtractJobAdvanceRequest struct {
	FileName          *string   `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileNameExtension *string   `json:"FileNameExtension,omitempty" xml:"FileNameExtension,omitempty"`
	FileUrlObject     io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s SubmitBookingNoteExtractJobAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitBookingNoteExtractJobAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SubmitBookingNoteExtractJobAdvanceRequest) SetFileName(v string) *SubmitBookingNoteExtractJobAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *SubmitBookingNoteExtractJobAdvanceRequest) SetFileNameExtension(v string) *SubmitBookingNoteExtractJobAdvanceRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitBookingNoteExtractJobAdvanceRequest) SetFileUrlObject(v io.Reader) *SubmitBookingNoteExtractJobAdvanceRequest {
	s.FileUrlObject = v
	return s
}

type SubmitBookingNoteExtractJobResponseBody struct {
	Code      *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SubmitBookingNoteExtractJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitBookingNoteExtractJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitBookingNoteExtractJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitBookingNoteExtractJobResponseBody) SetCode(v string) *SubmitBookingNoteExtractJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitBookingNoteExtractJobResponseBody) SetData(v *SubmitBookingNoteExtractJobResponseBodyData) *SubmitBookingNoteExtractJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitBookingNoteExtractJobResponseBody) SetMessage(v string) *SubmitBookingNoteExtractJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitBookingNoteExtractJobResponseBody) SetRequestId(v string) *SubmitBookingNoteExtractJobResponseBody {
	s.RequestId = &v
	return s
}

type SubmitBookingNoteExtractJobResponseBodyData struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SubmitBookingNoteExtractJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitBookingNoteExtractJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitBookingNoteExtractJobResponseBodyData) SetId(v string) *SubmitBookingNoteExtractJobResponseBodyData {
	s.Id = &v
	return s
}

type SubmitBookingNoteExtractJobResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitBookingNoteExtractJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitBookingNoteExtractJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitBookingNoteExtractJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitBookingNoteExtractJobResponse) SetHeaders(v map[string]*string) *SubmitBookingNoteExtractJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitBookingNoteExtractJobResponse) SetStatusCode(v int32) *SubmitBookingNoteExtractJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitBookingNoteExtractJobResponse) SetBody(v *SubmitBookingNoteExtractJobResponseBody) *SubmitBookingNoteExtractJobResponse {
	s.Body = v
	return s
}

type SubmitCertificateOfOriginExtractJobRequest struct {
	FileName          *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileNameExtension *string `json:"FileNameExtension,omitempty" xml:"FileNameExtension,omitempty"`
	FileUrl           *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s SubmitCertificateOfOriginExtractJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitCertificateOfOriginExtractJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitCertificateOfOriginExtractJobRequest) SetFileName(v string) *SubmitCertificateOfOriginExtractJobRequest {
	s.FileName = &v
	return s
}

func (s *SubmitCertificateOfOriginExtractJobRequest) SetFileNameExtension(v string) *SubmitCertificateOfOriginExtractJobRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitCertificateOfOriginExtractJobRequest) SetFileUrl(v string) *SubmitCertificateOfOriginExtractJobRequest {
	s.FileUrl = &v
	return s
}

type SubmitCertificateOfOriginExtractJobAdvanceRequest struct {
	FileName          *string   `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileNameExtension *string   `json:"FileNameExtension,omitempty" xml:"FileNameExtension,omitempty"`
	FileUrlObject     io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s SubmitCertificateOfOriginExtractJobAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitCertificateOfOriginExtractJobAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SubmitCertificateOfOriginExtractJobAdvanceRequest) SetFileName(v string) *SubmitCertificateOfOriginExtractJobAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *SubmitCertificateOfOriginExtractJobAdvanceRequest) SetFileNameExtension(v string) *SubmitCertificateOfOriginExtractJobAdvanceRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitCertificateOfOriginExtractJobAdvanceRequest) SetFileUrlObject(v io.Reader) *SubmitCertificateOfOriginExtractJobAdvanceRequest {
	s.FileUrlObject = v
	return s
}

type SubmitCertificateOfOriginExtractJobResponseBody struct {
	Code      *string                                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SubmitCertificateOfOriginExtractJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitCertificateOfOriginExtractJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitCertificateOfOriginExtractJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitCertificateOfOriginExtractJobResponseBody) SetCode(v string) *SubmitCertificateOfOriginExtractJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitCertificateOfOriginExtractJobResponseBody) SetData(v *SubmitCertificateOfOriginExtractJobResponseBodyData) *SubmitCertificateOfOriginExtractJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitCertificateOfOriginExtractJobResponseBody) SetMessage(v string) *SubmitCertificateOfOriginExtractJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitCertificateOfOriginExtractJobResponseBody) SetRequestId(v string) *SubmitCertificateOfOriginExtractJobResponseBody {
	s.RequestId = &v
	return s
}

type SubmitCertificateOfOriginExtractJobResponseBodyData struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SubmitCertificateOfOriginExtractJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitCertificateOfOriginExtractJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitCertificateOfOriginExtractJobResponseBodyData) SetId(v string) *SubmitCertificateOfOriginExtractJobResponseBodyData {
	s.Id = &v
	return s
}

type SubmitCertificateOfOriginExtractJobResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitCertificateOfOriginExtractJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitCertificateOfOriginExtractJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitCertificateOfOriginExtractJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitCertificateOfOriginExtractJobResponse) SetHeaders(v map[string]*string) *SubmitCertificateOfOriginExtractJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitCertificateOfOriginExtractJobResponse) SetStatusCode(v int32) *SubmitCertificateOfOriginExtractJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitCertificateOfOriginExtractJobResponse) SetBody(v *SubmitCertificateOfOriginExtractJobResponseBody) *SubmitCertificateOfOriginExtractJobResponse {
	s.Body = v
	return s
}

type SubmitContainerLoadPlanExtractJobRequest struct {
	FileName          *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileNameExtension *string `json:"FileNameExtension,omitempty" xml:"FileNameExtension,omitempty"`
	FileUrl           *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s SubmitContainerLoadPlanExtractJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitContainerLoadPlanExtractJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitContainerLoadPlanExtractJobRequest) SetFileName(v string) *SubmitContainerLoadPlanExtractJobRequest {
	s.FileName = &v
	return s
}

func (s *SubmitContainerLoadPlanExtractJobRequest) SetFileNameExtension(v string) *SubmitContainerLoadPlanExtractJobRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitContainerLoadPlanExtractJobRequest) SetFileUrl(v string) *SubmitContainerLoadPlanExtractJobRequest {
	s.FileUrl = &v
	return s
}

type SubmitContainerLoadPlanExtractJobAdvanceRequest struct {
	FileName          *string   `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileNameExtension *string   `json:"FileNameExtension,omitempty" xml:"FileNameExtension,omitempty"`
	FileUrlObject     io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s SubmitContainerLoadPlanExtractJobAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitContainerLoadPlanExtractJobAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SubmitContainerLoadPlanExtractJobAdvanceRequest) SetFileName(v string) *SubmitContainerLoadPlanExtractJobAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *SubmitContainerLoadPlanExtractJobAdvanceRequest) SetFileNameExtension(v string) *SubmitContainerLoadPlanExtractJobAdvanceRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitContainerLoadPlanExtractJobAdvanceRequest) SetFileUrlObject(v io.Reader) *SubmitContainerLoadPlanExtractJobAdvanceRequest {
	s.FileUrlObject = v
	return s
}

type SubmitContainerLoadPlanExtractJobResponseBody struct {
	Code      *string                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SubmitContainerLoadPlanExtractJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitContainerLoadPlanExtractJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitContainerLoadPlanExtractJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitContainerLoadPlanExtractJobResponseBody) SetCode(v string) *SubmitContainerLoadPlanExtractJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitContainerLoadPlanExtractJobResponseBody) SetData(v *SubmitContainerLoadPlanExtractJobResponseBodyData) *SubmitContainerLoadPlanExtractJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitContainerLoadPlanExtractJobResponseBody) SetMessage(v string) *SubmitContainerLoadPlanExtractJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitContainerLoadPlanExtractJobResponseBody) SetRequestId(v string) *SubmitContainerLoadPlanExtractJobResponseBody {
	s.RequestId = &v
	return s
}

type SubmitContainerLoadPlanExtractJobResponseBodyData struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SubmitContainerLoadPlanExtractJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitContainerLoadPlanExtractJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitContainerLoadPlanExtractJobResponseBodyData) SetId(v string) *SubmitContainerLoadPlanExtractJobResponseBodyData {
	s.Id = &v
	return s
}

type SubmitContainerLoadPlanExtractJobResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitContainerLoadPlanExtractJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitContainerLoadPlanExtractJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitContainerLoadPlanExtractJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitContainerLoadPlanExtractJobResponse) SetHeaders(v map[string]*string) *SubmitContainerLoadPlanExtractJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitContainerLoadPlanExtractJobResponse) SetStatusCode(v int32) *SubmitContainerLoadPlanExtractJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitContainerLoadPlanExtractJobResponse) SetBody(v *SubmitContainerLoadPlanExtractJobResponseBody) *SubmitContainerLoadPlanExtractJobResponse {
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

type SubmitImportDeclarationSheetExtractJobRequest struct {
	FileName          *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileNameExtension *string `json:"FileNameExtension,omitempty" xml:"FileNameExtension,omitempty"`
	FileUrl           *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s SubmitImportDeclarationSheetExtractJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitImportDeclarationSheetExtractJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitImportDeclarationSheetExtractJobRequest) SetFileName(v string) *SubmitImportDeclarationSheetExtractJobRequest {
	s.FileName = &v
	return s
}

func (s *SubmitImportDeclarationSheetExtractJobRequest) SetFileNameExtension(v string) *SubmitImportDeclarationSheetExtractJobRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitImportDeclarationSheetExtractJobRequest) SetFileUrl(v string) *SubmitImportDeclarationSheetExtractJobRequest {
	s.FileUrl = &v
	return s
}

type SubmitImportDeclarationSheetExtractJobAdvanceRequest struct {
	FileName          *string   `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileNameExtension *string   `json:"FileNameExtension,omitempty" xml:"FileNameExtension,omitempty"`
	FileUrlObject     io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s SubmitImportDeclarationSheetExtractJobAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitImportDeclarationSheetExtractJobAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SubmitImportDeclarationSheetExtractJobAdvanceRequest) SetFileName(v string) *SubmitImportDeclarationSheetExtractJobAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *SubmitImportDeclarationSheetExtractJobAdvanceRequest) SetFileNameExtension(v string) *SubmitImportDeclarationSheetExtractJobAdvanceRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitImportDeclarationSheetExtractJobAdvanceRequest) SetFileUrlObject(v io.Reader) *SubmitImportDeclarationSheetExtractJobAdvanceRequest {
	s.FileUrlObject = v
	return s
}

type SubmitImportDeclarationSheetExtractJobResponseBody struct {
	Code      *string                                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SubmitImportDeclarationSheetExtractJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitImportDeclarationSheetExtractJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitImportDeclarationSheetExtractJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitImportDeclarationSheetExtractJobResponseBody) SetCode(v string) *SubmitImportDeclarationSheetExtractJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitImportDeclarationSheetExtractJobResponseBody) SetData(v *SubmitImportDeclarationSheetExtractJobResponseBodyData) *SubmitImportDeclarationSheetExtractJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitImportDeclarationSheetExtractJobResponseBody) SetMessage(v string) *SubmitImportDeclarationSheetExtractJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitImportDeclarationSheetExtractJobResponseBody) SetRequestId(v string) *SubmitImportDeclarationSheetExtractJobResponseBody {
	s.RequestId = &v
	return s
}

type SubmitImportDeclarationSheetExtractJobResponseBodyData struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SubmitImportDeclarationSheetExtractJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitImportDeclarationSheetExtractJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitImportDeclarationSheetExtractJobResponseBodyData) SetId(v string) *SubmitImportDeclarationSheetExtractJobResponseBodyData {
	s.Id = &v
	return s
}

type SubmitImportDeclarationSheetExtractJobResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitImportDeclarationSheetExtractJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitImportDeclarationSheetExtractJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitImportDeclarationSheetExtractJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitImportDeclarationSheetExtractJobResponse) SetHeaders(v map[string]*string) *SubmitImportDeclarationSheetExtractJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitImportDeclarationSheetExtractJobResponse) SetStatusCode(v int32) *SubmitImportDeclarationSheetExtractJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitImportDeclarationSheetExtractJobResponse) SetBody(v *SubmitImportDeclarationSheetExtractJobResponseBody) *SubmitImportDeclarationSheetExtractJobResponse {
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

type SubmitPackingListExtractJobRequest struct {
	FileName          *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileNameExtension *string `json:"FileNameExtension,omitempty" xml:"FileNameExtension,omitempty"`
	FileUrl           *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s SubmitPackingListExtractJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitPackingListExtractJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitPackingListExtractJobRequest) SetFileName(v string) *SubmitPackingListExtractJobRequest {
	s.FileName = &v
	return s
}

func (s *SubmitPackingListExtractJobRequest) SetFileNameExtension(v string) *SubmitPackingListExtractJobRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitPackingListExtractJobRequest) SetFileUrl(v string) *SubmitPackingListExtractJobRequest {
	s.FileUrl = &v
	return s
}

type SubmitPackingListExtractJobAdvanceRequest struct {
	FileName          *string   `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileNameExtension *string   `json:"FileNameExtension,omitempty" xml:"FileNameExtension,omitempty"`
	FileUrlObject     io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s SubmitPackingListExtractJobAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitPackingListExtractJobAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SubmitPackingListExtractJobAdvanceRequest) SetFileName(v string) *SubmitPackingListExtractJobAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *SubmitPackingListExtractJobAdvanceRequest) SetFileNameExtension(v string) *SubmitPackingListExtractJobAdvanceRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitPackingListExtractJobAdvanceRequest) SetFileUrlObject(v io.Reader) *SubmitPackingListExtractJobAdvanceRequest {
	s.FileUrlObject = v
	return s
}

type SubmitPackingListExtractJobResponseBody struct {
	Code      *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SubmitPackingListExtractJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitPackingListExtractJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitPackingListExtractJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitPackingListExtractJobResponseBody) SetCode(v string) *SubmitPackingListExtractJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitPackingListExtractJobResponseBody) SetData(v *SubmitPackingListExtractJobResponseBodyData) *SubmitPackingListExtractJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitPackingListExtractJobResponseBody) SetMessage(v string) *SubmitPackingListExtractJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitPackingListExtractJobResponseBody) SetRequestId(v string) *SubmitPackingListExtractJobResponseBody {
	s.RequestId = &v
	return s
}

type SubmitPackingListExtractJobResponseBodyData struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SubmitPackingListExtractJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitPackingListExtractJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitPackingListExtractJobResponseBodyData) SetId(v string) *SubmitPackingListExtractJobResponseBodyData {
	s.Id = &v
	return s
}

type SubmitPackingListExtractJobResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitPackingListExtractJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitPackingListExtractJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitPackingListExtractJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitPackingListExtractJobResponse) SetHeaders(v map[string]*string) *SubmitPackingListExtractJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitPackingListExtractJobResponse) SetStatusCode(v int32) *SubmitPackingListExtractJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitPackingListExtractJobResponse) SetBody(v *SubmitPackingListExtractJobResponseBody) *SubmitPackingListExtractJobResponse {
	s.Body = v
	return s
}

type SubmitSalesConfirmationExtractJobRequest struct {
	FileName          *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileNameExtension *string `json:"FileNameExtension,omitempty" xml:"FileNameExtension,omitempty"`
	FileUrl           *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s SubmitSalesConfirmationExtractJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitSalesConfirmationExtractJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitSalesConfirmationExtractJobRequest) SetFileName(v string) *SubmitSalesConfirmationExtractJobRequest {
	s.FileName = &v
	return s
}

func (s *SubmitSalesConfirmationExtractJobRequest) SetFileNameExtension(v string) *SubmitSalesConfirmationExtractJobRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitSalesConfirmationExtractJobRequest) SetFileUrl(v string) *SubmitSalesConfirmationExtractJobRequest {
	s.FileUrl = &v
	return s
}

type SubmitSalesConfirmationExtractJobAdvanceRequest struct {
	FileName          *string   `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileNameExtension *string   `json:"FileNameExtension,omitempty" xml:"FileNameExtension,omitempty"`
	FileUrlObject     io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s SubmitSalesConfirmationExtractJobAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitSalesConfirmationExtractJobAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SubmitSalesConfirmationExtractJobAdvanceRequest) SetFileName(v string) *SubmitSalesConfirmationExtractJobAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *SubmitSalesConfirmationExtractJobAdvanceRequest) SetFileNameExtension(v string) *SubmitSalesConfirmationExtractJobAdvanceRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitSalesConfirmationExtractJobAdvanceRequest) SetFileUrlObject(v io.Reader) *SubmitSalesConfirmationExtractJobAdvanceRequest {
	s.FileUrlObject = v
	return s
}

type SubmitSalesConfirmationExtractJobResponseBody struct {
	Code      *string                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SubmitSalesConfirmationExtractJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitSalesConfirmationExtractJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitSalesConfirmationExtractJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSalesConfirmationExtractJobResponseBody) SetCode(v string) *SubmitSalesConfirmationExtractJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitSalesConfirmationExtractJobResponseBody) SetData(v *SubmitSalesConfirmationExtractJobResponseBodyData) *SubmitSalesConfirmationExtractJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitSalesConfirmationExtractJobResponseBody) SetMessage(v string) *SubmitSalesConfirmationExtractJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitSalesConfirmationExtractJobResponseBody) SetRequestId(v string) *SubmitSalesConfirmationExtractJobResponseBody {
	s.RequestId = &v
	return s
}

type SubmitSalesConfirmationExtractJobResponseBodyData struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SubmitSalesConfirmationExtractJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitSalesConfirmationExtractJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitSalesConfirmationExtractJobResponseBodyData) SetId(v string) *SubmitSalesConfirmationExtractJobResponseBodyData {
	s.Id = &v
	return s
}

type SubmitSalesConfirmationExtractJobResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitSalesConfirmationExtractJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitSalesConfirmationExtractJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitSalesConfirmationExtractJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitSalesConfirmationExtractJobResponse) SetHeaders(v map[string]*string) *SubmitSalesConfirmationExtractJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitSalesConfirmationExtractJobResponse) SetStatusCode(v int32) *SubmitSalesConfirmationExtractJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitSalesConfirmationExtractJobResponse) SetBody(v *SubmitSalesConfirmationExtractJobResponseBody) *SubmitSalesConfirmationExtractJobResponse {
	s.Body = v
	return s
}

type SubmitSeaWaybillExtractJobRequest struct {
	FileName          *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileNameExtension *string `json:"FileNameExtension,omitempty" xml:"FileNameExtension,omitempty"`
	FileUrl           *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s SubmitSeaWaybillExtractJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitSeaWaybillExtractJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitSeaWaybillExtractJobRequest) SetFileName(v string) *SubmitSeaWaybillExtractJobRequest {
	s.FileName = &v
	return s
}

func (s *SubmitSeaWaybillExtractJobRequest) SetFileNameExtension(v string) *SubmitSeaWaybillExtractJobRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitSeaWaybillExtractJobRequest) SetFileUrl(v string) *SubmitSeaWaybillExtractJobRequest {
	s.FileUrl = &v
	return s
}

type SubmitSeaWaybillExtractJobAdvanceRequest struct {
	FileName          *string   `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileNameExtension *string   `json:"FileNameExtension,omitempty" xml:"FileNameExtension,omitempty"`
	FileUrlObject     io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s SubmitSeaWaybillExtractJobAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitSeaWaybillExtractJobAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SubmitSeaWaybillExtractJobAdvanceRequest) SetFileName(v string) *SubmitSeaWaybillExtractJobAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *SubmitSeaWaybillExtractJobAdvanceRequest) SetFileNameExtension(v string) *SubmitSeaWaybillExtractJobAdvanceRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitSeaWaybillExtractJobAdvanceRequest) SetFileUrlObject(v io.Reader) *SubmitSeaWaybillExtractJobAdvanceRequest {
	s.FileUrlObject = v
	return s
}

type SubmitSeaWaybillExtractJobResponseBody struct {
	Code      *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SubmitSeaWaybillExtractJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitSeaWaybillExtractJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitSeaWaybillExtractJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSeaWaybillExtractJobResponseBody) SetCode(v string) *SubmitSeaWaybillExtractJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitSeaWaybillExtractJobResponseBody) SetData(v *SubmitSeaWaybillExtractJobResponseBodyData) *SubmitSeaWaybillExtractJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitSeaWaybillExtractJobResponseBody) SetMessage(v string) *SubmitSeaWaybillExtractJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitSeaWaybillExtractJobResponseBody) SetRequestId(v string) *SubmitSeaWaybillExtractJobResponseBody {
	s.RequestId = &v
	return s
}

type SubmitSeaWaybillExtractJobResponseBodyData struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SubmitSeaWaybillExtractJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitSeaWaybillExtractJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitSeaWaybillExtractJobResponseBodyData) SetId(v string) *SubmitSeaWaybillExtractJobResponseBodyData {
	s.Id = &v
	return s
}

type SubmitSeaWaybillExtractJobResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitSeaWaybillExtractJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitSeaWaybillExtractJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitSeaWaybillExtractJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitSeaWaybillExtractJobResponse) SetHeaders(v map[string]*string) *SubmitSeaWaybillExtractJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitSeaWaybillExtractJobResponse) SetStatusCode(v int32) *SubmitSeaWaybillExtractJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitSeaWaybillExtractJobResponse) SetBody(v *SubmitSeaWaybillExtractJobResponseBody) *SubmitSeaWaybillExtractJobResponse {
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

func (client *Client) SubmitBookingNoteExtractJobWithOptions(request *SubmitBookingNoteExtractJobRequest, runtime *util.RuntimeOptions) (_result *SubmitBookingNoteExtractJobResponse, _err error) {
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
		Action:      tea.String("SubmitBookingNoteExtractJob"),
		Version:     tea.String("2022-07-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitBookingNoteExtractJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitBookingNoteExtractJob(request *SubmitBookingNoteExtractJobRequest) (_result *SubmitBookingNoteExtractJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitBookingNoteExtractJobResponse{}
	_body, _err := client.SubmitBookingNoteExtractJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitBookingNoteExtractJobAdvance(request *SubmitBookingNoteExtractJobAdvanceRequest, runtime *util.RuntimeOptions) (_result *SubmitBookingNoteExtractJobResponse, _err error) {
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
	submitBookingNoteExtractJobReq := &SubmitBookingNoteExtractJobRequest{}
	openapiutil.Convert(request, submitBookingNoteExtractJobReq)
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
		submitBookingNoteExtractJobReq.FileUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	submitBookingNoteExtractJobResp, _err := client.SubmitBookingNoteExtractJobWithOptions(submitBookingNoteExtractJobReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = submitBookingNoteExtractJobResp
	return _result, _err
}

func (client *Client) SubmitCertificateOfOriginExtractJobWithOptions(request *SubmitCertificateOfOriginExtractJobRequest, runtime *util.RuntimeOptions) (_result *SubmitCertificateOfOriginExtractJobResponse, _err error) {
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
		Action:      tea.String("SubmitCertificateOfOriginExtractJob"),
		Version:     tea.String("2022-07-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitCertificateOfOriginExtractJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitCertificateOfOriginExtractJob(request *SubmitCertificateOfOriginExtractJobRequest) (_result *SubmitCertificateOfOriginExtractJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitCertificateOfOriginExtractJobResponse{}
	_body, _err := client.SubmitCertificateOfOriginExtractJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitCertificateOfOriginExtractJobAdvance(request *SubmitCertificateOfOriginExtractJobAdvanceRequest, runtime *util.RuntimeOptions) (_result *SubmitCertificateOfOriginExtractJobResponse, _err error) {
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
	submitCertificateOfOriginExtractJobReq := &SubmitCertificateOfOriginExtractJobRequest{}
	openapiutil.Convert(request, submitCertificateOfOriginExtractJobReq)
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
		submitCertificateOfOriginExtractJobReq.FileUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	submitCertificateOfOriginExtractJobResp, _err := client.SubmitCertificateOfOriginExtractJobWithOptions(submitCertificateOfOriginExtractJobReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = submitCertificateOfOriginExtractJobResp
	return _result, _err
}

func (client *Client) SubmitContainerLoadPlanExtractJobWithOptions(request *SubmitContainerLoadPlanExtractJobRequest, runtime *util.RuntimeOptions) (_result *SubmitContainerLoadPlanExtractJobResponse, _err error) {
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
		Action:      tea.String("SubmitContainerLoadPlanExtractJob"),
		Version:     tea.String("2022-07-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitContainerLoadPlanExtractJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitContainerLoadPlanExtractJob(request *SubmitContainerLoadPlanExtractJobRequest) (_result *SubmitContainerLoadPlanExtractJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitContainerLoadPlanExtractJobResponse{}
	_body, _err := client.SubmitContainerLoadPlanExtractJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitContainerLoadPlanExtractJobAdvance(request *SubmitContainerLoadPlanExtractJobAdvanceRequest, runtime *util.RuntimeOptions) (_result *SubmitContainerLoadPlanExtractJobResponse, _err error) {
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
	submitContainerLoadPlanExtractJobReq := &SubmitContainerLoadPlanExtractJobRequest{}
	openapiutil.Convert(request, submitContainerLoadPlanExtractJobReq)
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
		submitContainerLoadPlanExtractJobReq.FileUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	submitContainerLoadPlanExtractJobResp, _err := client.SubmitContainerLoadPlanExtractJobWithOptions(submitContainerLoadPlanExtractJobReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = submitContainerLoadPlanExtractJobResp
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

func (client *Client) SubmitImportDeclarationSheetExtractJobWithOptions(request *SubmitImportDeclarationSheetExtractJobRequest, runtime *util.RuntimeOptions) (_result *SubmitImportDeclarationSheetExtractJobResponse, _err error) {
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
		Action:      tea.String("SubmitImportDeclarationSheetExtractJob"),
		Version:     tea.String("2022-07-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitImportDeclarationSheetExtractJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitImportDeclarationSheetExtractJob(request *SubmitImportDeclarationSheetExtractJobRequest) (_result *SubmitImportDeclarationSheetExtractJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitImportDeclarationSheetExtractJobResponse{}
	_body, _err := client.SubmitImportDeclarationSheetExtractJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitImportDeclarationSheetExtractJobAdvance(request *SubmitImportDeclarationSheetExtractJobAdvanceRequest, runtime *util.RuntimeOptions) (_result *SubmitImportDeclarationSheetExtractJobResponse, _err error) {
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
	submitImportDeclarationSheetExtractJobReq := &SubmitImportDeclarationSheetExtractJobRequest{}
	openapiutil.Convert(request, submitImportDeclarationSheetExtractJobReq)
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
		submitImportDeclarationSheetExtractJobReq.FileUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	submitImportDeclarationSheetExtractJobResp, _err := client.SubmitImportDeclarationSheetExtractJobWithOptions(submitImportDeclarationSheetExtractJobReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = submitImportDeclarationSheetExtractJobResp
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

func (client *Client) SubmitPackingListExtractJobWithOptions(request *SubmitPackingListExtractJobRequest, runtime *util.RuntimeOptions) (_result *SubmitPackingListExtractJobResponse, _err error) {
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
		Action:      tea.String("SubmitPackingListExtractJob"),
		Version:     tea.String("2022-07-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitPackingListExtractJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitPackingListExtractJob(request *SubmitPackingListExtractJobRequest) (_result *SubmitPackingListExtractJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitPackingListExtractJobResponse{}
	_body, _err := client.SubmitPackingListExtractJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitPackingListExtractJobAdvance(request *SubmitPackingListExtractJobAdvanceRequest, runtime *util.RuntimeOptions) (_result *SubmitPackingListExtractJobResponse, _err error) {
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
	submitPackingListExtractJobReq := &SubmitPackingListExtractJobRequest{}
	openapiutil.Convert(request, submitPackingListExtractJobReq)
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
		submitPackingListExtractJobReq.FileUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	submitPackingListExtractJobResp, _err := client.SubmitPackingListExtractJobWithOptions(submitPackingListExtractJobReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = submitPackingListExtractJobResp
	return _result, _err
}

func (client *Client) SubmitSalesConfirmationExtractJobWithOptions(request *SubmitSalesConfirmationExtractJobRequest, runtime *util.RuntimeOptions) (_result *SubmitSalesConfirmationExtractJobResponse, _err error) {
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
		Action:      tea.String("SubmitSalesConfirmationExtractJob"),
		Version:     tea.String("2022-07-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitSalesConfirmationExtractJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitSalesConfirmationExtractJob(request *SubmitSalesConfirmationExtractJobRequest) (_result *SubmitSalesConfirmationExtractJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitSalesConfirmationExtractJobResponse{}
	_body, _err := client.SubmitSalesConfirmationExtractJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitSalesConfirmationExtractJobAdvance(request *SubmitSalesConfirmationExtractJobAdvanceRequest, runtime *util.RuntimeOptions) (_result *SubmitSalesConfirmationExtractJobResponse, _err error) {
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
	submitSalesConfirmationExtractJobReq := &SubmitSalesConfirmationExtractJobRequest{}
	openapiutil.Convert(request, submitSalesConfirmationExtractJobReq)
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
		submitSalesConfirmationExtractJobReq.FileUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	submitSalesConfirmationExtractJobResp, _err := client.SubmitSalesConfirmationExtractJobWithOptions(submitSalesConfirmationExtractJobReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = submitSalesConfirmationExtractJobResp
	return _result, _err
}

func (client *Client) SubmitSeaWaybillExtractJobWithOptions(request *SubmitSeaWaybillExtractJobRequest, runtime *util.RuntimeOptions) (_result *SubmitSeaWaybillExtractJobResponse, _err error) {
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
		Action:      tea.String("SubmitSeaWaybillExtractJob"),
		Version:     tea.String("2022-07-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitSeaWaybillExtractJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitSeaWaybillExtractJob(request *SubmitSeaWaybillExtractJobRequest) (_result *SubmitSeaWaybillExtractJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitSeaWaybillExtractJobResponse{}
	_body, _err := client.SubmitSeaWaybillExtractJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitSeaWaybillExtractJobAdvance(request *SubmitSeaWaybillExtractJobAdvanceRequest, runtime *util.RuntimeOptions) (_result *SubmitSeaWaybillExtractJobResponse, _err error) {
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
	submitSeaWaybillExtractJobReq := &SubmitSeaWaybillExtractJobRequest{}
	openapiutil.Convert(request, submitSeaWaybillExtractJobReq)
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
		submitSeaWaybillExtractJobReq.FileUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	submitSeaWaybillExtractJobResp, _err := client.SubmitSeaWaybillExtractJobWithOptions(submitSeaWaybillExtractJobReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = submitSeaWaybillExtractJobResp
	return _result, _err
}
