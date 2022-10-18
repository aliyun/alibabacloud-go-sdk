// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CloseTicketRequest struct {
	TicketId *string `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
}

func (s CloseTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s CloseTicketRequest) GoString() string {
	return s.String()
}

func (s *CloseTicketRequest) SetTicketId(v string) *CloseTicketRequest {
	s.TicketId = &v
	return s
}

type CloseTicketResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CloseTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloseTicketResponseBody) GoString() string {
	return s.String()
}

func (s *CloseTicketResponseBody) SetCode(v int32) *CloseTicketResponseBody {
	s.Code = &v
	return s
}

func (s *CloseTicketResponseBody) SetMessage(v string) *CloseTicketResponseBody {
	s.Message = &v
	return s
}

func (s *CloseTicketResponseBody) SetRequestId(v string) *CloseTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseTicketResponseBody) SetSuccess(v bool) *CloseTicketResponseBody {
	s.Success = &v
	return s
}

type CloseTicketResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CloseTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CloseTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s CloseTicketResponse) GoString() string {
	return s.String()
}

func (s *CloseTicketResponse) SetHeaders(v map[string]*string) *CloseTicketResponse {
	s.Headers = v
	return s
}

func (s *CloseTicketResponse) SetStatusCode(v int32) *CloseTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseTicketResponse) SetBody(v *CloseTicketResponseBody) *CloseTicketResponse {
	s.Body = v
	return s
}

type CreateTicketRequest struct {
	CategoryId   *string                        `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CreatorId    *string                        `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	Description  *string                        `json:"Description,omitempty" xml:"Description,omitempty"`
	Email        *string                        `json:"Email,omitempty" xml:"Email,omitempty"`
	FileNameList []*string                      `json:"FileNameList,omitempty" xml:"FileNameList,omitempty" type:"Repeated"`
	SecretInfo   *CreateTicketRequestSecretInfo `json:"SecretInfo,omitempty" xml:"SecretInfo,omitempty" type:"Struct"`
	Severity     *int32                         `json:"Severity,omitempty" xml:"Severity,omitempty"`
	Title        *string                        `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketRequest) GoString() string {
	return s.String()
}

func (s *CreateTicketRequest) SetCategoryId(v string) *CreateTicketRequest {
	s.CategoryId = &v
	return s
}

func (s *CreateTicketRequest) SetCreatorId(v string) *CreateTicketRequest {
	s.CreatorId = &v
	return s
}

func (s *CreateTicketRequest) SetDescription(v string) *CreateTicketRequest {
	s.Description = &v
	return s
}

func (s *CreateTicketRequest) SetEmail(v string) *CreateTicketRequest {
	s.Email = &v
	return s
}

func (s *CreateTicketRequest) SetFileNameList(v []*string) *CreateTicketRequest {
	s.FileNameList = v
	return s
}

func (s *CreateTicketRequest) SetSecretInfo(v *CreateTicketRequestSecretInfo) *CreateTicketRequest {
	s.SecretInfo = v
	return s
}

func (s *CreateTicketRequest) SetSeverity(v int32) *CreateTicketRequest {
	s.Severity = &v
	return s
}

func (s *CreateTicketRequest) SetTitle(v string) *CreateTicketRequest {
	s.Title = &v
	return s
}

type CreateTicketRequestSecretInfo struct {
	Content      *string   `json:"Content,omitempty" xml:"Content,omitempty"`
	FileNameList []*string `json:"FileNameList,omitempty" xml:"FileNameList,omitempty" type:"Repeated"`
}

func (s CreateTicketRequestSecretInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketRequestSecretInfo) GoString() string {
	return s.String()
}

func (s *CreateTicketRequestSecretInfo) SetContent(v string) *CreateTicketRequestSecretInfo {
	s.Content = &v
	return s
}

func (s *CreateTicketRequestSecretInfo) SetFileNameList(v []*string) *CreateTicketRequestSecretInfo {
	s.FileNameList = v
	return s
}

type CreateTicketShrinkRequest struct {
	CategoryId         *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CreatorId          *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Email              *string `json:"Email,omitempty" xml:"Email,omitempty"`
	FileNameListShrink *string `json:"FileNameList,omitempty" xml:"FileNameList,omitempty"`
	SecretInfoShrink   *string `json:"SecretInfo,omitempty" xml:"SecretInfo,omitempty"`
	Severity           *int32  `json:"Severity,omitempty" xml:"Severity,omitempty"`
	Title              *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateTicketShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateTicketShrinkRequest) SetCategoryId(v string) *CreateTicketShrinkRequest {
	s.CategoryId = &v
	return s
}

func (s *CreateTicketShrinkRequest) SetCreatorId(v string) *CreateTicketShrinkRequest {
	s.CreatorId = &v
	return s
}

func (s *CreateTicketShrinkRequest) SetDescription(v string) *CreateTicketShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateTicketShrinkRequest) SetEmail(v string) *CreateTicketShrinkRequest {
	s.Email = &v
	return s
}

func (s *CreateTicketShrinkRequest) SetFileNameListShrink(v string) *CreateTicketShrinkRequest {
	s.FileNameListShrink = &v
	return s
}

func (s *CreateTicketShrinkRequest) SetSecretInfoShrink(v string) *CreateTicketShrinkRequest {
	s.SecretInfoShrink = &v
	return s
}

func (s *CreateTicketShrinkRequest) SetSeverity(v int32) *CreateTicketShrinkRequest {
	s.Severity = &v
	return s
}

func (s *CreateTicketShrinkRequest) SetTitle(v string) *CreateTicketShrinkRequest {
	s.Title = &v
	return s
}

type CreateTicketResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTicketResponseBody) SetCode(v int32) *CreateTicketResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTicketResponseBody) SetData(v string) *CreateTicketResponseBody {
	s.Data = &v
	return s
}

func (s *CreateTicketResponseBody) SetMessage(v string) *CreateTicketResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTicketResponseBody) SetRequestId(v string) *CreateTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTicketResponseBody) SetSuccess(v bool) *CreateTicketResponseBody {
	s.Success = &v
	return s
}

type CreateTicketResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketResponse) GoString() string {
	return s.String()
}

func (s *CreateTicketResponse) SetHeaders(v map[string]*string) *CreateTicketResponse {
	s.Headers = v
	return s
}

func (s *CreateTicketResponse) SetStatusCode(v int32) *CreateTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTicketResponse) SetBody(v *CreateTicketResponseBody) *CreateTicketResponse {
	s.Body = v
	return s
}

type EvaluateTicketRequest struct {
	Content  *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Score    *string `json:"Score,omitempty" xml:"Score,omitempty"`
	Solved   *bool   `json:"Solved,omitempty" xml:"Solved,omitempty"`
	TicketId *string `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
}

func (s EvaluateTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s EvaluateTicketRequest) GoString() string {
	return s.String()
}

func (s *EvaluateTicketRequest) SetContent(v string) *EvaluateTicketRequest {
	s.Content = &v
	return s
}

func (s *EvaluateTicketRequest) SetScore(v string) *EvaluateTicketRequest {
	s.Score = &v
	return s
}

func (s *EvaluateTicketRequest) SetSolved(v bool) *EvaluateTicketRequest {
	s.Solved = &v
	return s
}

func (s *EvaluateTicketRequest) SetTicketId(v string) *EvaluateTicketRequest {
	s.TicketId = &v
	return s
}

type EvaluateTicketResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EvaluateTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EvaluateTicketResponseBody) GoString() string {
	return s.String()
}

func (s *EvaluateTicketResponseBody) SetCode(v int32) *EvaluateTicketResponseBody {
	s.Code = &v
	return s
}

func (s *EvaluateTicketResponseBody) SetMessage(v string) *EvaluateTicketResponseBody {
	s.Message = &v
	return s
}

func (s *EvaluateTicketResponseBody) SetRequestId(v string) *EvaluateTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *EvaluateTicketResponseBody) SetSuccess(v bool) *EvaluateTicketResponseBody {
	s.Success = &v
	return s
}

type EvaluateTicketResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EvaluateTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EvaluateTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s EvaluateTicketResponse) GoString() string {
	return s.String()
}

func (s *EvaluateTicketResponse) SetHeaders(v map[string]*string) *EvaluateTicketResponse {
	s.Headers = v
	return s
}

func (s *EvaluateTicketResponse) SetStatusCode(v int32) *EvaluateTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *EvaluateTicketResponse) SetBody(v *EvaluateTicketResponseBody) *EvaluateTicketResponse {
	s.Body = v
	return s
}

type GetAttachmentUploadUrlRequest struct {
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s GetAttachmentUploadUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAttachmentUploadUrlRequest) GoString() string {
	return s.String()
}

func (s *GetAttachmentUploadUrlRequest) SetFileName(v string) *GetAttachmentUploadUrlRequest {
	s.FileName = &v
	return s
}

type GetAttachmentUploadUrlResponseBody struct {
	Code      *int32                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetAttachmentUploadUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAttachmentUploadUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAttachmentUploadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetAttachmentUploadUrlResponseBody) SetCode(v int32) *GetAttachmentUploadUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetAttachmentUploadUrlResponseBody) SetData(v *GetAttachmentUploadUrlResponseBodyData) *GetAttachmentUploadUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetAttachmentUploadUrlResponseBody) SetMessage(v string) *GetAttachmentUploadUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetAttachmentUploadUrlResponseBody) SetRequestId(v string) *GetAttachmentUploadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAttachmentUploadUrlResponseBody) SetSuccess(v bool) *GetAttachmentUploadUrlResponseBody {
	s.Success = &v
	return s
}

type GetAttachmentUploadUrlResponseBodyData struct {
	GetSignedUrl *string `json:"GetSignedUrl,omitempty" xml:"GetSignedUrl,omitempty"`
	ObjectKey    *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
	PutSignedUrl *string `json:"PutSignedUrl,omitempty" xml:"PutSignedUrl,omitempty"`
}

func (s GetAttachmentUploadUrlResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAttachmentUploadUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAttachmentUploadUrlResponseBodyData) SetGetSignedUrl(v string) *GetAttachmentUploadUrlResponseBodyData {
	s.GetSignedUrl = &v
	return s
}

func (s *GetAttachmentUploadUrlResponseBodyData) SetObjectKey(v string) *GetAttachmentUploadUrlResponseBodyData {
	s.ObjectKey = &v
	return s
}

func (s *GetAttachmentUploadUrlResponseBodyData) SetPutSignedUrl(v string) *GetAttachmentUploadUrlResponseBodyData {
	s.PutSignedUrl = &v
	return s
}

type GetAttachmentUploadUrlResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAttachmentUploadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAttachmentUploadUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAttachmentUploadUrlResponse) GoString() string {
	return s.String()
}

func (s *GetAttachmentUploadUrlResponse) SetHeaders(v map[string]*string) *GetAttachmentUploadUrlResponse {
	s.Headers = v
	return s
}

func (s *GetAttachmentUploadUrlResponse) SetStatusCode(v int32) *GetAttachmentUploadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAttachmentUploadUrlResponse) SetBody(v *GetAttachmentUploadUrlResponseBody) *GetAttachmentUploadUrlResponse {
	s.Body = v
	return s
}

type GetMqConsumerTagResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMqConsumerTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMqConsumerTagResponseBody) GoString() string {
	return s.String()
}

func (s *GetMqConsumerTagResponseBody) SetCode(v int32) *GetMqConsumerTagResponseBody {
	s.Code = &v
	return s
}

func (s *GetMqConsumerTagResponseBody) SetData(v string) *GetMqConsumerTagResponseBody {
	s.Data = &v
	return s
}

func (s *GetMqConsumerTagResponseBody) SetMessage(v string) *GetMqConsumerTagResponseBody {
	s.Message = &v
	return s
}

func (s *GetMqConsumerTagResponseBody) SetRequestId(v string) *GetMqConsumerTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMqConsumerTagResponseBody) SetSuccess(v bool) *GetMqConsumerTagResponseBody {
	s.Success = &v
	return s
}

type GetMqConsumerTagResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMqConsumerTagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMqConsumerTagResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMqConsumerTagResponse) GoString() string {
	return s.String()
}

func (s *GetMqConsumerTagResponse) SetHeaders(v map[string]*string) *GetMqConsumerTagResponse {
	s.Headers = v
	return s
}

func (s *GetMqConsumerTagResponse) SetStatusCode(v int32) *GetMqConsumerTagResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMqConsumerTagResponse) SetBody(v *GetMqConsumerTagResponseBody) *GetMqConsumerTagResponse {
	s.Body = v
	return s
}

type GetTicketRequest struct {
	TicketId *string `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
}

func (s GetTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTicketRequest) GoString() string {
	return s.String()
}

func (s *GetTicketRequest) SetTicketId(v string) *GetTicketRequest {
	s.TicketId = &v
	return s
}

type GetTicketResponseBody struct {
	Code       *int32                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       *GetTicketResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message    *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                      `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int64                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTicketResponseBody) GoString() string {
	return s.String()
}

func (s *GetTicketResponseBody) SetCode(v int32) *GetTicketResponseBody {
	s.Code = &v
	return s
}

func (s *GetTicketResponseBody) SetData(v *GetTicketResponseBodyData) *GetTicketResponseBody {
	s.Data = v
	return s
}

func (s *GetTicketResponseBody) SetMessage(v string) *GetTicketResponseBody {
	s.Message = &v
	return s
}

func (s *GetTicketResponseBody) SetPageNumber(v int32) *GetTicketResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetTicketResponseBody) SetPageSize(v int32) *GetTicketResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetTicketResponseBody) SetRequestId(v string) *GetTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTicketResponseBody) SetSuccess(v bool) *GetTicketResponseBody {
	s.Success = &v
	return s
}

func (s *GetTicketResponseBody) SetTotalCount(v int64) *GetTicketResponseBody {
	s.TotalCount = &v
	return s
}

type GetTicketResponseBodyData struct {
	CategoryId  *string                            `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CreateTime  *int64                             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreatorId   *string                            `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	Description *string                            `json:"Description,omitempty" xml:"Description,omitempty"`
	Severity    *GetTicketResponseBodyDataSeverity `json:"Severity,omitempty" xml:"Severity,omitempty" type:"Struct"`
	Status      *GetTicketResponseBodyDataStatus   `json:"Status,omitempty" xml:"Status,omitempty" type:"Struct"`
	TicketId    *string                            `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
	Title       *string                            `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetTicketResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTicketResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTicketResponseBodyData) SetCategoryId(v string) *GetTicketResponseBodyData {
	s.CategoryId = &v
	return s
}

func (s *GetTicketResponseBodyData) SetCreateTime(v int64) *GetTicketResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetTicketResponseBodyData) SetCreatorId(v string) *GetTicketResponseBodyData {
	s.CreatorId = &v
	return s
}

func (s *GetTicketResponseBodyData) SetDescription(v string) *GetTicketResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetTicketResponseBodyData) SetSeverity(v *GetTicketResponseBodyDataSeverity) *GetTicketResponseBodyData {
	s.Severity = v
	return s
}

func (s *GetTicketResponseBodyData) SetStatus(v *GetTicketResponseBodyDataStatus) *GetTicketResponseBodyData {
	s.Status = v
	return s
}

func (s *GetTicketResponseBodyData) SetTicketId(v string) *GetTicketResponseBodyData {
	s.TicketId = &v
	return s
}

func (s *GetTicketResponseBodyData) SetTitle(v string) *GetTicketResponseBodyData {
	s.Title = &v
	return s
}

type GetTicketResponseBodyDataSeverity struct {
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTicketResponseBodyDataSeverity) String() string {
	return tea.Prettify(s)
}

func (s GetTicketResponseBodyDataSeverity) GoString() string {
	return s.String()
}

func (s *GetTicketResponseBodyDataSeverity) SetLabel(v string) *GetTicketResponseBodyDataSeverity {
	s.Label = &v
	return s
}

func (s *GetTicketResponseBodyDataSeverity) SetValue(v string) *GetTicketResponseBodyDataSeverity {
	s.Value = &v
	return s
}

type GetTicketResponseBodyDataStatus struct {
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTicketResponseBodyDataStatus) String() string {
	return tea.Prettify(s)
}

func (s GetTicketResponseBodyDataStatus) GoString() string {
	return s.String()
}

func (s *GetTicketResponseBodyDataStatus) SetLabel(v string) *GetTicketResponseBodyDataStatus {
	s.Label = &v
	return s
}

func (s *GetTicketResponseBodyDataStatus) SetValue(v string) *GetTicketResponseBodyDataStatus {
	s.Value = &v
	return s
}

type GetTicketResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTicketResponse) GoString() string {
	return s.String()
}

func (s *GetTicketResponse) SetHeaders(v map[string]*string) *GetTicketResponse {
	s.Headers = v
	return s
}

func (s *GetTicketResponse) SetStatusCode(v int32) *GetTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTicketResponse) SetBody(v *GetTicketResponseBody) *GetTicketResponse {
	s.Body = v
	return s
}

type ListCategoriesRequest struct {
	Language  *string `json:"Language,omitempty" xml:"Language,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProductId *int64  `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
}

func (s ListCategoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCategoriesRequest) GoString() string {
	return s.String()
}

func (s *ListCategoriesRequest) SetLanguage(v string) *ListCategoriesRequest {
	s.Language = &v
	return s
}

func (s *ListCategoriesRequest) SetName(v string) *ListCategoriesRequest {
	s.Name = &v
	return s
}

func (s *ListCategoriesRequest) SetProductId(v int64) *ListCategoriesRequest {
	s.ProductId = &v
	return s
}

type ListCategoriesResponseBody struct {
	Code      *int32                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*ListCategoriesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCategoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCategoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCategoriesResponseBody) SetCode(v int32) *ListCategoriesResponseBody {
	s.Code = &v
	return s
}

func (s *ListCategoriesResponseBody) SetData(v []*ListCategoriesResponseBodyData) *ListCategoriesResponseBody {
	s.Data = v
	return s
}

func (s *ListCategoriesResponseBody) SetMessage(v string) *ListCategoriesResponseBody {
	s.Message = &v
	return s
}

func (s *ListCategoriesResponseBody) SetRequestId(v string) *ListCategoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCategoriesResponseBody) SetSuccess(v bool) *ListCategoriesResponseBody {
	s.Success = &v
	return s
}

type ListCategoriesResponseBodyData struct {
	CategoryId   *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
}

func (s ListCategoriesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListCategoriesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCategoriesResponseBodyData) SetCategoryId(v int64) *ListCategoriesResponseBodyData {
	s.CategoryId = &v
	return s
}

func (s *ListCategoriesResponseBodyData) SetCategoryName(v string) *ListCategoriesResponseBodyData {
	s.CategoryName = &v
	return s
}

type ListCategoriesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCategoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCategoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCategoriesResponse) GoString() string {
	return s.String()
}

func (s *ListCategoriesResponse) SetHeaders(v map[string]*string) *ListCategoriesResponse {
	s.Headers = v
	return s
}

func (s *ListCategoriesResponse) SetStatusCode(v int32) *ListCategoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCategoriesResponse) SetBody(v *ListCategoriesResponseBody) *ListCategoriesResponse {
	s.Body = v
	return s
}

type ListProductsRequest struct {
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListProductsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProductsRequest) GoString() string {
	return s.String()
}

func (s *ListProductsRequest) SetLanguage(v string) *ListProductsRequest {
	s.Language = &v
	return s
}

func (s *ListProductsRequest) SetName(v string) *ListProductsRequest {
	s.Name = &v
	return s
}

type ListProductsResponseBody struct {
	Code      *int32                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*ListProductsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListProductsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProductsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBody) SetCode(v int32) *ListProductsResponseBody {
	s.Code = &v
	return s
}

func (s *ListProductsResponseBody) SetData(v []*ListProductsResponseBodyData) *ListProductsResponseBody {
	s.Data = v
	return s
}

func (s *ListProductsResponseBody) SetMessage(v string) *ListProductsResponseBody {
	s.Message = &v
	return s
}

func (s *ListProductsResponseBody) SetRequestId(v string) *ListProductsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductsResponseBody) SetSuccess(v bool) *ListProductsResponseBody {
	s.Success = &v
	return s
}

type ListProductsResponseBodyData struct {
	DirectoryId   *int64                                     `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	DirectoryName *string                                    `json:"DirectoryName,omitempty" xml:"DirectoryName,omitempty"`
	ProductList   []*ListProductsResponseBodyDataProductList `json:"ProductList,omitempty" xml:"ProductList,omitempty" type:"Repeated"`
}

func (s ListProductsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListProductsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBodyData) SetDirectoryId(v int64) *ListProductsResponseBodyData {
	s.DirectoryId = &v
	return s
}

func (s *ListProductsResponseBodyData) SetDirectoryName(v string) *ListProductsResponseBodyData {
	s.DirectoryName = &v
	return s
}

func (s *ListProductsResponseBodyData) SetProductList(v []*ListProductsResponseBodyDataProductList) *ListProductsResponseBodyData {
	s.ProductList = v
	return s
}

type ListProductsResponseBodyDataProductList struct {
	ProductId   *int64  `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
}

func (s ListProductsResponseBodyDataProductList) String() string {
	return tea.Prettify(s)
}

func (s ListProductsResponseBodyDataProductList) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBodyDataProductList) SetProductId(v int64) *ListProductsResponseBodyDataProductList {
	s.ProductId = &v
	return s
}

func (s *ListProductsResponseBodyDataProductList) SetProductName(v string) *ListProductsResponseBodyDataProductList {
	s.ProductName = &v
	return s
}

type ListProductsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListProductsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProductsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProductsResponse) GoString() string {
	return s.String()
}

func (s *ListProductsResponse) SetHeaders(v map[string]*string) *ListProductsResponse {
	s.Headers = v
	return s
}

func (s *ListProductsResponse) SetStatusCode(v int32) *ListProductsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProductsResponse) SetBody(v *ListProductsResponseBody) *ListProductsResponse {
	s.Body = v
	return s
}

type ListTicketNotesRequest struct {
	TicketId *string `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
}

func (s ListTicketNotesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTicketNotesRequest) GoString() string {
	return s.String()
}

func (s *ListTicketNotesRequest) SetTicketId(v string) *ListTicketNotesRequest {
	s.TicketId = &v
	return s
}

type ListTicketNotesResponseBody struct {
	Code      *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*ListTicketNotesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListTicketNotesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTicketNotesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTicketNotesResponseBody) SetCode(v int32) *ListTicketNotesResponseBody {
	s.Code = &v
	return s
}

func (s *ListTicketNotesResponseBody) SetData(v []*ListTicketNotesResponseBodyData) *ListTicketNotesResponseBody {
	s.Data = v
	return s
}

func (s *ListTicketNotesResponseBody) SetMessage(v string) *ListTicketNotesResponseBody {
	s.Message = &v
	return s
}

func (s *ListTicketNotesResponseBody) SetRequestId(v string) *ListTicketNotesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTicketNotesResponseBody) SetSuccess(v bool) *ListTicketNotesResponseBody {
	s.Success = &v
	return s
}

type ListTicketNotesResponseBodyData struct {
	Attachments []*ListTicketNotesResponseBodyDataAttachments `json:"Attachments,omitempty" xml:"Attachments,omitempty" type:"Repeated"`
	CreateTime  *int64                                        `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Dialog      *ListTicketNotesResponseBodyDataDialog        `json:"Dialog,omitempty" xml:"Dialog,omitempty" type:"Struct"`
	DialogId    *int64                                        `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
	Status      *int32                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	Tip         *string                                       `json:"Tip,omitempty" xml:"Tip,omitempty"`
	Type        *int32                                        `json:"Type,omitempty" xml:"Type,omitempty"`
	User        *ListTicketNotesResponseBodyDataUser          `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
}

func (s ListTicketNotesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTicketNotesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTicketNotesResponseBodyData) SetAttachments(v []*ListTicketNotesResponseBodyDataAttachments) *ListTicketNotesResponseBodyData {
	s.Attachments = v
	return s
}

func (s *ListTicketNotesResponseBodyData) SetCreateTime(v int64) *ListTicketNotesResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListTicketNotesResponseBodyData) SetDialog(v *ListTicketNotesResponseBodyDataDialog) *ListTicketNotesResponseBodyData {
	s.Dialog = v
	return s
}

func (s *ListTicketNotesResponseBodyData) SetDialogId(v int64) *ListTicketNotesResponseBodyData {
	s.DialogId = &v
	return s
}

func (s *ListTicketNotesResponseBodyData) SetStatus(v int32) *ListTicketNotesResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListTicketNotesResponseBodyData) SetTip(v string) *ListTicketNotesResponseBodyData {
	s.Tip = &v
	return s
}

func (s *ListTicketNotesResponseBodyData) SetType(v int32) *ListTicketNotesResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListTicketNotesResponseBodyData) SetUser(v *ListTicketNotesResponseBodyDataUser) *ListTicketNotesResponseBodyData {
	s.User = v
	return s
}

type ListTicketNotesResponseBodyDataAttachments struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Url  *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListTicketNotesResponseBodyDataAttachments) String() string {
	return tea.Prettify(s)
}

func (s ListTicketNotesResponseBodyDataAttachments) GoString() string {
	return s.String()
}

func (s *ListTicketNotesResponseBodyDataAttachments) SetName(v string) *ListTicketNotesResponseBodyDataAttachments {
	s.Name = &v
	return s
}

func (s *ListTicketNotesResponseBodyDataAttachments) SetUrl(v string) *ListTicketNotesResponseBodyDataAttachments {
	s.Url = &v
	return s
}

type ListTicketNotesResponseBodyDataDialog struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Schema  *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
}

func (s ListTicketNotesResponseBodyDataDialog) String() string {
	return tea.Prettify(s)
}

func (s ListTicketNotesResponseBodyDataDialog) GoString() string {
	return s.String()
}

func (s *ListTicketNotesResponseBodyDataDialog) SetContent(v string) *ListTicketNotesResponseBodyDataDialog {
	s.Content = &v
	return s
}

func (s *ListTicketNotesResponseBodyDataDialog) SetSchema(v string) *ListTicketNotesResponseBodyDataDialog {
	s.Schema = &v
	return s
}

type ListTicketNotesResponseBodyDataUser struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Role *int32  `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s ListTicketNotesResponseBodyDataUser) String() string {
	return tea.Prettify(s)
}

func (s ListTicketNotesResponseBodyDataUser) GoString() string {
	return s.String()
}

func (s *ListTicketNotesResponseBodyDataUser) SetName(v string) *ListTicketNotesResponseBodyDataUser {
	s.Name = &v
	return s
}

func (s *ListTicketNotesResponseBodyDataUser) SetRole(v int32) *ListTicketNotesResponseBodyDataUser {
	s.Role = &v
	return s
}

type ListTicketNotesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTicketNotesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTicketNotesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTicketNotesResponse) GoString() string {
	return s.String()
}

func (s *ListTicketNotesResponse) SetHeaders(v map[string]*string) *ListTicketNotesResponse {
	s.Headers = v
	return s
}

func (s *ListTicketNotesResponse) SetStatusCode(v int32) *ListTicketNotesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTicketNotesResponse) SetBody(v *ListTicketNotesResponseBody) *ListTicketNotesResponse {
	s.Body = v
	return s
}

type ListTicketsRequest struct {
	EndDate      *int64    `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Keyword      *string   `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	PageNumber   *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartDate    *int64    `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	StatusList   []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
	TicketId     *string   `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
	TicketIdList []*string `json:"TicketIdList,omitempty" xml:"TicketIdList,omitempty" type:"Repeated"`
}

func (s ListTicketsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTicketsRequest) GoString() string {
	return s.String()
}

func (s *ListTicketsRequest) SetEndDate(v int64) *ListTicketsRequest {
	s.EndDate = &v
	return s
}

func (s *ListTicketsRequest) SetKeyword(v string) *ListTicketsRequest {
	s.Keyword = &v
	return s
}

func (s *ListTicketsRequest) SetPageNumber(v int32) *ListTicketsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTicketsRequest) SetPageSize(v int32) *ListTicketsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTicketsRequest) SetStartDate(v int64) *ListTicketsRequest {
	s.StartDate = &v
	return s
}

func (s *ListTicketsRequest) SetStatusList(v []*string) *ListTicketsRequest {
	s.StatusList = v
	return s
}

func (s *ListTicketsRequest) SetTicketId(v string) *ListTicketsRequest {
	s.TicketId = &v
	return s
}

func (s *ListTicketsRequest) SetTicketIdList(v []*string) *ListTicketsRequest {
	s.TicketIdList = v
	return s
}

type ListTicketsShrinkRequest struct {
	EndDate            *int64    `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Keyword            *string   `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	PageNumber         *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartDate          *int64    `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	StatusList         []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
	TicketId           *string   `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
	TicketIdListShrink *string   `json:"TicketIdList,omitempty" xml:"TicketIdList,omitempty"`
}

func (s ListTicketsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTicketsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTicketsShrinkRequest) SetEndDate(v int64) *ListTicketsShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *ListTicketsShrinkRequest) SetKeyword(v string) *ListTicketsShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListTicketsShrinkRequest) SetPageNumber(v int32) *ListTicketsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTicketsShrinkRequest) SetPageSize(v int32) *ListTicketsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListTicketsShrinkRequest) SetStartDate(v int64) *ListTicketsShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *ListTicketsShrinkRequest) SetStatusList(v []*string) *ListTicketsShrinkRequest {
	s.StatusList = v
	return s
}

func (s *ListTicketsShrinkRequest) SetTicketId(v string) *ListTicketsShrinkRequest {
	s.TicketId = &v
	return s
}

func (s *ListTicketsShrinkRequest) SetTicketIdListShrink(v string) *ListTicketsShrinkRequest {
	s.TicketIdListShrink = &v
	return s
}

type ListTicketsResponseBody struct {
	Code       *int32                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       []*ListTicketsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message    *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int64                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTicketsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTicketsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTicketsResponseBody) SetCode(v int32) *ListTicketsResponseBody {
	s.Code = &v
	return s
}

func (s *ListTicketsResponseBody) SetData(v []*ListTicketsResponseBodyData) *ListTicketsResponseBody {
	s.Data = v
	return s
}

func (s *ListTicketsResponseBody) SetMessage(v string) *ListTicketsResponseBody {
	s.Message = &v
	return s
}

func (s *ListTicketsResponseBody) SetPageNumber(v int32) *ListTicketsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTicketsResponseBody) SetPageSize(v int32) *ListTicketsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTicketsResponseBody) SetRequestId(v string) *ListTicketsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTicketsResponseBody) SetSuccess(v bool) *ListTicketsResponseBody {
	s.Success = &v
	return s
}

func (s *ListTicketsResponseBody) SetTotalCount(v int64) *ListTicketsResponseBody {
	s.TotalCount = &v
	return s
}

type ListTicketsResponseBodyData struct {
	Status   *ListTicketsResponseBodyDataStatus `json:"Status,omitempty" xml:"Status,omitempty" type:"Struct"`
	TicketId *string                            `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
	Title    *string                            `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListTicketsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTicketsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTicketsResponseBodyData) SetStatus(v *ListTicketsResponseBodyDataStatus) *ListTicketsResponseBodyData {
	s.Status = v
	return s
}

func (s *ListTicketsResponseBodyData) SetTicketId(v string) *ListTicketsResponseBodyData {
	s.TicketId = &v
	return s
}

func (s *ListTicketsResponseBodyData) SetTitle(v string) *ListTicketsResponseBodyData {
	s.Title = &v
	return s
}

type ListTicketsResponseBodyDataStatus struct {
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTicketsResponseBodyDataStatus) String() string {
	return tea.Prettify(s)
}

func (s ListTicketsResponseBodyDataStatus) GoString() string {
	return s.String()
}

func (s *ListTicketsResponseBodyDataStatus) SetLabel(v string) *ListTicketsResponseBodyDataStatus {
	s.Label = &v
	return s
}

func (s *ListTicketsResponseBodyDataStatus) SetValue(v string) *ListTicketsResponseBodyDataStatus {
	s.Value = &v
	return s
}

type ListTicketsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTicketsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTicketsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTicketsResponse) GoString() string {
	return s.String()
}

func (s *ListTicketsResponse) SetHeaders(v map[string]*string) *ListTicketsResponse {
	s.Headers = v
	return s
}

func (s *ListTicketsResponse) SetStatusCode(v int32) *ListTicketsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTicketsResponse) SetBody(v *ListTicketsResponseBody) *ListTicketsResponse {
	s.Body = v
	return s
}

type ReopenTicketRequest struct {
	Content  *string `json:"Content,omitempty" xml:"Content,omitempty"`
	TicketId *string `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
}

func (s ReopenTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s ReopenTicketRequest) GoString() string {
	return s.String()
}

func (s *ReopenTicketRequest) SetContent(v string) *ReopenTicketRequest {
	s.Content = &v
	return s
}

func (s *ReopenTicketRequest) SetTicketId(v string) *ReopenTicketRequest {
	s.TicketId = &v
	return s
}

type ReopenTicketResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReopenTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReopenTicketResponseBody) GoString() string {
	return s.String()
}

func (s *ReopenTicketResponseBody) SetCode(v int32) *ReopenTicketResponseBody {
	s.Code = &v
	return s
}

func (s *ReopenTicketResponseBody) SetData(v string) *ReopenTicketResponseBody {
	s.Data = &v
	return s
}

func (s *ReopenTicketResponseBody) SetMessage(v string) *ReopenTicketResponseBody {
	s.Message = &v
	return s
}

func (s *ReopenTicketResponseBody) SetRequestId(v string) *ReopenTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReopenTicketResponseBody) SetSuccess(v bool) *ReopenTicketResponseBody {
	s.Success = &v
	return s
}

type ReopenTicketResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReopenTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReopenTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s ReopenTicketResponse) GoString() string {
	return s.String()
}

func (s *ReopenTicketResponse) SetHeaders(v map[string]*string) *ReopenTicketResponse {
	s.Headers = v
	return s
}

func (s *ReopenTicketResponse) SetStatusCode(v int32) *ReopenTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *ReopenTicketResponse) SetBody(v *ReopenTicketResponseBody) *ReopenTicketResponse {
	s.Body = v
	return s
}

type ReplyTicketRequest struct {
	Content      *string   `json:"Content,omitempty" xml:"Content,omitempty"`
	Encrypt      *bool     `json:"Encrypt,omitempty" xml:"Encrypt,omitempty"`
	FileNameList []*string `json:"FileNameList,omitempty" xml:"FileNameList,omitempty" type:"Repeated"`
	TicketId     *string   `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
}

func (s ReplyTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s ReplyTicketRequest) GoString() string {
	return s.String()
}

func (s *ReplyTicketRequest) SetContent(v string) *ReplyTicketRequest {
	s.Content = &v
	return s
}

func (s *ReplyTicketRequest) SetEncrypt(v bool) *ReplyTicketRequest {
	s.Encrypt = &v
	return s
}

func (s *ReplyTicketRequest) SetFileNameList(v []*string) *ReplyTicketRequest {
	s.FileNameList = v
	return s
}

func (s *ReplyTicketRequest) SetTicketId(v string) *ReplyTicketRequest {
	s.TicketId = &v
	return s
}

type ReplyTicketShrinkRequest struct {
	Content            *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Encrypt            *bool   `json:"Encrypt,omitempty" xml:"Encrypt,omitempty"`
	FileNameListShrink *string `json:"FileNameList,omitempty" xml:"FileNameList,omitempty"`
	TicketId           *string `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
}

func (s ReplyTicketShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ReplyTicketShrinkRequest) GoString() string {
	return s.String()
}

func (s *ReplyTicketShrinkRequest) SetContent(v string) *ReplyTicketShrinkRequest {
	s.Content = &v
	return s
}

func (s *ReplyTicketShrinkRequest) SetEncrypt(v bool) *ReplyTicketShrinkRequest {
	s.Encrypt = &v
	return s
}

func (s *ReplyTicketShrinkRequest) SetFileNameListShrink(v string) *ReplyTicketShrinkRequest {
	s.FileNameListShrink = &v
	return s
}

func (s *ReplyTicketShrinkRequest) SetTicketId(v string) *ReplyTicketShrinkRequest {
	s.TicketId = &v
	return s
}

type ReplyTicketResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReplyTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReplyTicketResponseBody) GoString() string {
	return s.String()
}

func (s *ReplyTicketResponseBody) SetCode(v int32) *ReplyTicketResponseBody {
	s.Code = &v
	return s
}

func (s *ReplyTicketResponseBody) SetData(v string) *ReplyTicketResponseBody {
	s.Data = &v
	return s
}

func (s *ReplyTicketResponseBody) SetMessage(v string) *ReplyTicketResponseBody {
	s.Message = &v
	return s
}

func (s *ReplyTicketResponseBody) SetRequestId(v string) *ReplyTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReplyTicketResponseBody) SetSuccess(v bool) *ReplyTicketResponseBody {
	s.Success = &v
	return s
}

type ReplyTicketResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReplyTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReplyTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s ReplyTicketResponse) GoString() string {
	return s.String()
}

func (s *ReplyTicketResponse) SetHeaders(v map[string]*string) *ReplyTicketResponse {
	s.Headers = v
	return s
}

func (s *ReplyTicketResponse) SetStatusCode(v int32) *ReplyTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *ReplyTicketResponse) SetBody(v *ReplyTicketResponseBody) *ReplyTicketResponse {
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
	client.EndpointRule = tea.String("central")
	client.EndpointMap = map[string]*string{
		"ap-northeast-1": tea.String("workorder.ap-northeast-1.aliyuncs.com"),
		"ap-southeast-1": tea.String("workorder.ap-southeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("workorder"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CloseTicketWithOptions(request *CloseTicketRequest, runtime *util.RuntimeOptions) (_result *CloseTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TicketId)) {
		body["TicketId"] = request.TicketId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CloseTicket"),
		Version:     tea.String("2021-06-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CloseTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CloseTicket(request *CloseTicketRequest) (_result *CloseTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CloseTicketResponse{}
	_body, _err := client.CloseTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTicketWithOptions(tmpReq *CreateTicketRequest, runtime *util.RuntimeOptions) (_result *CreateTicketResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateTicketShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.FileNameList)) {
		request.FileNameListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FileNameList, tea.String("FileNameList"), tea.String("simple"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.SecretInfo))) {
		request.SecretInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.SecretInfo), tea.String("SecretInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SecretInfoShrink)) {
		query["SecretInfo"] = request.SecretInfoShrink
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CategoryId)) {
		body["CategoryId"] = request.CategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.CreatorId)) {
		body["CreatorId"] = request.CreatorId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		body["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.FileNameListShrink)) {
		body["FileNameList"] = request.FileNameListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Severity)) {
		body["Severity"] = request.Severity
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTicket"),
		Version:     tea.String("2021-06-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTicket(request *CreateTicketRequest) (_result *CreateTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTicketResponse{}
	_body, _err := client.CreateTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EvaluateTicketWithOptions(request *EvaluateTicketRequest, runtime *util.RuntimeOptions) (_result *EvaluateTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Score)) {
		body["Score"] = request.Score
	}

	if !tea.BoolValue(util.IsUnset(request.Solved)) {
		body["Solved"] = request.Solved
	}

	if !tea.BoolValue(util.IsUnset(request.TicketId)) {
		body["TicketId"] = request.TicketId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EvaluateTicket"),
		Version:     tea.String("2021-06-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EvaluateTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EvaluateTicket(request *EvaluateTicketRequest) (_result *EvaluateTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EvaluateTicketResponse{}
	_body, _err := client.EvaluateTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAttachmentUploadUrlWithOptions(request *GetAttachmentUploadUrlRequest, runtime *util.RuntimeOptions) (_result *GetAttachmentUploadUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		body["FileName"] = request.FileName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAttachmentUploadUrl"),
		Version:     tea.String("2021-06-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAttachmentUploadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAttachmentUploadUrl(request *GetAttachmentUploadUrlRequest) (_result *GetAttachmentUploadUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAttachmentUploadUrlResponse{}
	_body, _err := client.GetAttachmentUploadUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMqConsumerTagWithOptions(runtime *util.RuntimeOptions) (_result *GetMqConsumerTagResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetMqConsumerTag"),
		Version:     tea.String("2021-06-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMqConsumerTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMqConsumerTag() (_result *GetMqConsumerTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMqConsumerTagResponse{}
	_body, _err := client.GetMqConsumerTagWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTicketWithOptions(request *GetTicketRequest, runtime *util.RuntimeOptions) (_result *GetTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TicketId)) {
		body["TicketId"] = request.TicketId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTicket"),
		Version:     tea.String("2021-06-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTicket(request *GetTicketRequest) (_result *GetTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTicketResponse{}
	_body, _err := client.GetTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCategoriesWithOptions(request *ListCategoriesRequest, runtime *util.RuntimeOptions) (_result *ListCategoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		body["ProductId"] = request.ProductId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCategories"),
		Version:     tea.String("2021-06-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCategoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCategories(request *ListCategoriesRequest) (_result *ListCategoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCategoriesResponse{}
	_body, _err := client.ListCategoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProductsWithOptions(request *ListProductsRequest, runtime *util.RuntimeOptions) (_result *ListProductsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProducts"),
		Version:     tea.String("2021-06-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProductsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProducts(request *ListProductsRequest) (_result *ListProductsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListProductsResponse{}
	_body, _err := client.ListProductsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTicketNotesWithOptions(request *ListTicketNotesRequest, runtime *util.RuntimeOptions) (_result *ListTicketNotesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TicketId)) {
		query["TicketId"] = request.TicketId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTicketNotes"),
		Version:     tea.String("2021-06-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTicketNotesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTicketNotes(request *ListTicketNotesRequest) (_result *ListTicketNotesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTicketNotesResponse{}
	_body, _err := client.ListTicketNotesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTicketsWithOptions(tmpReq *ListTicketsRequest, runtime *util.RuntimeOptions) (_result *ListTicketsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListTicketsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.TicketIdList)) {
		request.TicketIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TicketIdList, tea.String("TicketIdList"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		body["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.StatusList)) {
		body["StatusList"] = request.StatusList
	}

	if !tea.BoolValue(util.IsUnset(request.TicketId)) {
		body["TicketId"] = request.TicketId
	}

	if !tea.BoolValue(util.IsUnset(request.TicketIdListShrink)) {
		body["TicketIdList"] = request.TicketIdListShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTickets"),
		Version:     tea.String("2021-06-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTicketsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTickets(request *ListTicketsRequest) (_result *ListTicketsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTicketsResponse{}
	_body, _err := client.ListTicketsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReopenTicketWithOptions(request *ReopenTicketRequest, runtime *util.RuntimeOptions) (_result *ReopenTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.TicketId)) {
		body["TicketId"] = request.TicketId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ReopenTicket"),
		Version:     tea.String("2021-06-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReopenTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReopenTicket(request *ReopenTicketRequest) (_result *ReopenTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReopenTicketResponse{}
	_body, _err := client.ReopenTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReplyTicketWithOptions(tmpReq *ReplyTicketRequest, runtime *util.RuntimeOptions) (_result *ReplyTicketResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ReplyTicketShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.FileNameList)) {
		request.FileNameListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FileNameList, tea.String("FileNameList"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileNameListShrink)) {
		query["FileNameList"] = request.FileNameListShrink
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Encrypt)) {
		body["Encrypt"] = request.Encrypt
	}

	if !tea.BoolValue(util.IsUnset(request.TicketId)) {
		body["TicketId"] = request.TicketId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ReplyTicket"),
		Version:     tea.String("2021-06-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReplyTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReplyTicket(request *ReplyTicketRequest) (_result *ReplyTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReplyTicketResponse{}
	_body, _err := client.ReplyTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
