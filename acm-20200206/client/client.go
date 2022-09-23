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

type BatchExportConfigurationsRequest struct {
	Data        *string `json:"Data,omitempty" xml:"Data,omitempty"`
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s BatchExportConfigurationsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchExportConfigurationsRequest) GoString() string {
	return s.String()
}

func (s *BatchExportConfigurationsRequest) SetData(v string) *BatchExportConfigurationsRequest {
	s.Data = &v
	return s
}

func (s *BatchExportConfigurationsRequest) SetNamespaceId(v string) *BatchExportConfigurationsRequest {
	s.NamespaceId = &v
	return s
}

type BatchExportConfigurationsResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	FileUrl   *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchExportConfigurationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchExportConfigurationsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchExportConfigurationsResponseBody) SetCode(v string) *BatchExportConfigurationsResponseBody {
	s.Code = &v
	return s
}

func (s *BatchExportConfigurationsResponseBody) SetFileUrl(v string) *BatchExportConfigurationsResponseBody {
	s.FileUrl = &v
	return s
}

func (s *BatchExportConfigurationsResponseBody) SetMessage(v string) *BatchExportConfigurationsResponseBody {
	s.Message = &v
	return s
}

func (s *BatchExportConfigurationsResponseBody) SetRequestId(v string) *BatchExportConfigurationsResponseBody {
	s.RequestId = &v
	return s
}

type BatchExportConfigurationsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchExportConfigurationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchExportConfigurationsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchExportConfigurationsResponse) GoString() string {
	return s.String()
}

func (s *BatchExportConfigurationsResponse) SetHeaders(v map[string]*string) *BatchExportConfigurationsResponse {
	s.Headers = v
	return s
}

func (s *BatchExportConfigurationsResponse) SetStatusCode(v int32) *BatchExportConfigurationsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchExportConfigurationsResponse) SetBody(v *BatchExportConfigurationsResponseBody) *BatchExportConfigurationsResponse {
	s.Body = v
	return s
}

type BatchImportConfigurationsRequest struct {
	FileUrl     *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	Policy      *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
}

func (s BatchImportConfigurationsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchImportConfigurationsRequest) GoString() string {
	return s.String()
}

func (s *BatchImportConfigurationsRequest) SetFileUrl(v string) *BatchImportConfigurationsRequest {
	s.FileUrl = &v
	return s
}

func (s *BatchImportConfigurationsRequest) SetNamespaceId(v string) *BatchImportConfigurationsRequest {
	s.NamespaceId = &v
	return s
}

func (s *BatchImportConfigurationsRequest) SetPolicy(v string) *BatchImportConfigurationsRequest {
	s.Policy = &v
	return s
}

type BatchImportConfigurationsResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchImportConfigurationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchImportConfigurationsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchImportConfigurationsResponseBody) SetCode(v string) *BatchImportConfigurationsResponseBody {
	s.Code = &v
	return s
}

func (s *BatchImportConfigurationsResponseBody) SetMessage(v string) *BatchImportConfigurationsResponseBody {
	s.Message = &v
	return s
}

func (s *BatchImportConfigurationsResponseBody) SetRequestId(v string) *BatchImportConfigurationsResponseBody {
	s.RequestId = &v
	return s
}

type BatchImportConfigurationsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchImportConfigurationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchImportConfigurationsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchImportConfigurationsResponse) GoString() string {
	return s.String()
}

func (s *BatchImportConfigurationsResponse) SetHeaders(v map[string]*string) *BatchImportConfigurationsResponse {
	s.Headers = v
	return s
}

func (s *BatchImportConfigurationsResponse) SetStatusCode(v int32) *BatchImportConfigurationsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchImportConfigurationsResponse) SetBody(v *BatchImportConfigurationsResponseBody) *BatchImportConfigurationsResponse {
	s.Body = v
	return s
}

type CheckConfigurationCloneRequest struct {
	Data          *string `json:"Data,omitempty" xml:"Data,omitempty"`
	NamespaceFrom *string `json:"NamespaceFrom,omitempty" xml:"NamespaceFrom,omitempty"`
	NamespaceTo   *string `json:"NamespaceTo,omitempty" xml:"NamespaceTo,omitempty"`
	Policy        *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
}

func (s CheckConfigurationCloneRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckConfigurationCloneRequest) GoString() string {
	return s.String()
}

func (s *CheckConfigurationCloneRequest) SetData(v string) *CheckConfigurationCloneRequest {
	s.Data = &v
	return s
}

func (s *CheckConfigurationCloneRequest) SetNamespaceFrom(v string) *CheckConfigurationCloneRequest {
	s.NamespaceFrom = &v
	return s
}

func (s *CheckConfigurationCloneRequest) SetNamespaceTo(v string) *CheckConfigurationCloneRequest {
	s.NamespaceTo = &v
	return s
}

func (s *CheckConfigurationCloneRequest) SetPolicy(v string) *CheckConfigurationCloneRequest {
	s.Policy = &v
	return s
}

type CheckConfigurationCloneResponseBody struct {
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CheckConfigurationCloneResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CheckConfigurationCloneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckConfigurationCloneResponseBody) GoString() string {
	return s.String()
}

func (s *CheckConfigurationCloneResponseBody) SetCode(v string) *CheckConfigurationCloneResponseBody {
	s.Code = &v
	return s
}

func (s *CheckConfigurationCloneResponseBody) SetMessage(v string) *CheckConfigurationCloneResponseBody {
	s.Message = &v
	return s
}

func (s *CheckConfigurationCloneResponseBody) SetRequestId(v string) *CheckConfigurationCloneResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckConfigurationCloneResponseBody) SetResult(v *CheckConfigurationCloneResponseBodyResult) *CheckConfigurationCloneResponseBody {
	s.Result = v
	return s
}

type CheckConfigurationCloneResponseBodyResult struct {
	Success      *bool                                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	SuccessItems []*CheckConfigurationCloneResponseBodyResultSuccessItems `json:"SuccessItems,omitempty" xml:"SuccessItems,omitempty" type:"Repeated"`
	TotalCount   *int32                                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s CheckConfigurationCloneResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CheckConfigurationCloneResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CheckConfigurationCloneResponseBodyResult) SetSuccess(v bool) *CheckConfigurationCloneResponseBodyResult {
	s.Success = &v
	return s
}

func (s *CheckConfigurationCloneResponseBodyResult) SetSuccessItems(v []*CheckConfigurationCloneResponseBodyResultSuccessItems) *CheckConfigurationCloneResponseBodyResult {
	s.SuccessItems = v
	return s
}

func (s *CheckConfigurationCloneResponseBodyResult) SetTotalCount(v int32) *CheckConfigurationCloneResponseBodyResult {
	s.TotalCount = &v
	return s
}

type CheckConfigurationCloneResponseBodyResultSuccessItems struct {
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	Group  *string `json:"Group,omitempty" xml:"Group,omitempty"`
}

func (s CheckConfigurationCloneResponseBodyResultSuccessItems) String() string {
	return tea.Prettify(s)
}

func (s CheckConfigurationCloneResponseBodyResultSuccessItems) GoString() string {
	return s.String()
}

func (s *CheckConfigurationCloneResponseBodyResultSuccessItems) SetDataId(v string) *CheckConfigurationCloneResponseBodyResultSuccessItems {
	s.DataId = &v
	return s
}

func (s *CheckConfigurationCloneResponseBodyResultSuccessItems) SetGroup(v string) *CheckConfigurationCloneResponseBodyResultSuccessItems {
	s.Group = &v
	return s
}

type CheckConfigurationCloneResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckConfigurationCloneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckConfigurationCloneResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckConfigurationCloneResponse) GoString() string {
	return s.String()
}

func (s *CheckConfigurationCloneResponse) SetHeaders(v map[string]*string) *CheckConfigurationCloneResponse {
	s.Headers = v
	return s
}

func (s *CheckConfigurationCloneResponse) SetStatusCode(v int32) *CheckConfigurationCloneResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckConfigurationCloneResponse) SetBody(v *CheckConfigurationCloneResponseBody) *CheckConfigurationCloneResponse {
	s.Body = v
	return s
}

type CheckConfigurationExportRequest struct {
	Data        *string `json:"Data,omitempty" xml:"Data,omitempty"`
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s CheckConfigurationExportRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckConfigurationExportRequest) GoString() string {
	return s.String()
}

func (s *CheckConfigurationExportRequest) SetData(v string) *CheckConfigurationExportRequest {
	s.Data = &v
	return s
}

func (s *CheckConfigurationExportRequest) SetNamespaceId(v string) *CheckConfigurationExportRequest {
	s.NamespaceId = &v
	return s
}

type CheckConfigurationExportResponseBody struct {
	Code      *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CheckConfigurationExportResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CheckConfigurationExportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckConfigurationExportResponseBody) GoString() string {
	return s.String()
}

func (s *CheckConfigurationExportResponseBody) SetCode(v string) *CheckConfigurationExportResponseBody {
	s.Code = &v
	return s
}

func (s *CheckConfigurationExportResponseBody) SetMessage(v string) *CheckConfigurationExportResponseBody {
	s.Message = &v
	return s
}

func (s *CheckConfigurationExportResponseBody) SetRequestId(v string) *CheckConfigurationExportResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckConfigurationExportResponseBody) SetResult(v *CheckConfigurationExportResponseBodyResult) *CheckConfigurationExportResponseBody {
	s.Result = v
	return s
}

type CheckConfigurationExportResponseBodyResult struct {
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s CheckConfigurationExportResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CheckConfigurationExportResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CheckConfigurationExportResponseBodyResult) SetTotalCount(v int32) *CheckConfigurationExportResponseBodyResult {
	s.TotalCount = &v
	return s
}

type CheckConfigurationExportResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckConfigurationExportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckConfigurationExportResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckConfigurationExportResponse) GoString() string {
	return s.String()
}

func (s *CheckConfigurationExportResponse) SetHeaders(v map[string]*string) *CheckConfigurationExportResponse {
	s.Headers = v
	return s
}

func (s *CheckConfigurationExportResponse) SetStatusCode(v int32) *CheckConfigurationExportResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckConfigurationExportResponse) SetBody(v *CheckConfigurationExportResponseBody) *CheckConfigurationExportResponse {
	s.Body = v
	return s
}

type CloneConfigurationRequest struct {
	Data          *string `json:"Data,omitempty" xml:"Data,omitempty"`
	NamespaceFrom *string `json:"NamespaceFrom,omitempty" xml:"NamespaceFrom,omitempty"`
	NamespaceTo   *string `json:"NamespaceTo,omitempty" xml:"NamespaceTo,omitempty"`
	Policy        *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
}

func (s CloneConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s CloneConfigurationRequest) GoString() string {
	return s.String()
}

func (s *CloneConfigurationRequest) SetData(v string) *CloneConfigurationRequest {
	s.Data = &v
	return s
}

func (s *CloneConfigurationRequest) SetNamespaceFrom(v string) *CloneConfigurationRequest {
	s.NamespaceFrom = &v
	return s
}

func (s *CloneConfigurationRequest) SetNamespaceTo(v string) *CloneConfigurationRequest {
	s.NamespaceTo = &v
	return s
}

func (s *CloneConfigurationRequest) SetPolicy(v string) *CloneConfigurationRequest {
	s.Policy = &v
	return s
}

type CloneConfigurationResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CloneConfigurationResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CloneConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloneConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *CloneConfigurationResponseBody) SetCode(v string) *CloneConfigurationResponseBody {
	s.Code = &v
	return s
}

func (s *CloneConfigurationResponseBody) SetMessage(v string) *CloneConfigurationResponseBody {
	s.Message = &v
	return s
}

func (s *CloneConfigurationResponseBody) SetRequestId(v string) *CloneConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloneConfigurationResponseBody) SetResult(v *CloneConfigurationResponseBodyResult) *CloneConfigurationResponseBody {
	s.Result = v
	return s
}

type CloneConfigurationResponseBodyResult struct {
	Success      *bool                                               `json:"Success,omitempty" xml:"Success,omitempty"`
	SuccessItems []*CloneConfigurationResponseBodyResultSuccessItems `json:"SuccessItems,omitempty" xml:"SuccessItems,omitempty" type:"Repeated"`
	TotalCount   *int32                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s CloneConfigurationResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CloneConfigurationResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CloneConfigurationResponseBodyResult) SetSuccess(v bool) *CloneConfigurationResponseBodyResult {
	s.Success = &v
	return s
}

func (s *CloneConfigurationResponseBodyResult) SetSuccessItems(v []*CloneConfigurationResponseBodyResultSuccessItems) *CloneConfigurationResponseBodyResult {
	s.SuccessItems = v
	return s
}

func (s *CloneConfigurationResponseBodyResult) SetTotalCount(v int32) *CloneConfigurationResponseBodyResult {
	s.TotalCount = &v
	return s
}

type CloneConfigurationResponseBodyResultSuccessItems struct {
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	Group  *string `json:"Group,omitempty" xml:"Group,omitempty"`
}

func (s CloneConfigurationResponseBodyResultSuccessItems) String() string {
	return tea.Prettify(s)
}

func (s CloneConfigurationResponseBodyResultSuccessItems) GoString() string {
	return s.String()
}

func (s *CloneConfigurationResponseBodyResultSuccessItems) SetDataId(v string) *CloneConfigurationResponseBodyResultSuccessItems {
	s.DataId = &v
	return s
}

func (s *CloneConfigurationResponseBodyResultSuccessItems) SetGroup(v string) *CloneConfigurationResponseBodyResultSuccessItems {
	s.Group = &v
	return s
}

type CloneConfigurationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CloneConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CloneConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s CloneConfigurationResponse) GoString() string {
	return s.String()
}

func (s *CloneConfigurationResponse) SetHeaders(v map[string]*string) *CloneConfigurationResponse {
	s.Headers = v
	return s
}

func (s *CloneConfigurationResponse) SetStatusCode(v int32) *CloneConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *CloneConfigurationResponse) SetBody(v *CloneConfigurationResponseBody) *CloneConfigurationResponse {
	s.Body = v
	return s
}

type CreateConfigurationRequest struct {
	AppName     *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	DataId      *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	Desc        *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Group       *string `json:"Group,omitempty" xml:"Group,omitempty"`
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	Tags        *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateConfigurationRequest) GoString() string {
	return s.String()
}

func (s *CreateConfigurationRequest) SetAppName(v string) *CreateConfigurationRequest {
	s.AppName = &v
	return s
}

func (s *CreateConfigurationRequest) SetContent(v string) *CreateConfigurationRequest {
	s.Content = &v
	return s
}

func (s *CreateConfigurationRequest) SetDataId(v string) *CreateConfigurationRequest {
	s.DataId = &v
	return s
}

func (s *CreateConfigurationRequest) SetDesc(v string) *CreateConfigurationRequest {
	s.Desc = &v
	return s
}

func (s *CreateConfigurationRequest) SetGroup(v string) *CreateConfigurationRequest {
	s.Group = &v
	return s
}

func (s *CreateConfigurationRequest) SetNamespaceId(v string) *CreateConfigurationRequest {
	s.NamespaceId = &v
	return s
}

func (s *CreateConfigurationRequest) SetTags(v string) *CreateConfigurationRequest {
	s.Tags = &v
	return s
}

func (s *CreateConfigurationRequest) SetType(v string) *CreateConfigurationRequest {
	s.Type = &v
	return s
}

type CreateConfigurationResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConfigurationResponseBody) SetCode(v string) *CreateConfigurationResponseBody {
	s.Code = &v
	return s
}

func (s *CreateConfigurationResponseBody) SetMessage(v string) *CreateConfigurationResponseBody {
	s.Message = &v
	return s
}

func (s *CreateConfigurationResponseBody) SetRequestId(v string) *CreateConfigurationResponseBody {
	s.RequestId = &v
	return s
}

type CreateConfigurationResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateConfigurationResponse) GoString() string {
	return s.String()
}

func (s *CreateConfigurationResponse) SetHeaders(v map[string]*string) *CreateConfigurationResponse {
	s.Headers = v
	return s
}

func (s *CreateConfigurationResponse) SetStatusCode(v int32) *CreateConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateConfigurationResponse) SetBody(v *CreateConfigurationResponseBody) *CreateConfigurationResponse {
	s.Body = v
	return s
}

type CreateNamespaceRequest struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateNamespaceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNamespaceRequest) GoString() string {
	return s.String()
}

func (s *CreateNamespaceRequest) SetName(v string) *CreateNamespaceRequest {
	s.Name = &v
	return s
}

type CreateNamespaceResponseBody struct {
	Code        *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty"`
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNamespaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNamespaceResponseBody) SetCode(v string) *CreateNamespaceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateNamespaceResponseBody) SetMessage(v string) *CreateNamespaceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateNamespaceResponseBody) SetNamespaceId(v string) *CreateNamespaceResponseBody {
	s.NamespaceId = &v
	return s
}

func (s *CreateNamespaceResponseBody) SetRequestId(v string) *CreateNamespaceResponseBody {
	s.RequestId = &v
	return s
}

type CreateNamespaceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateNamespaceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateNamespaceResponse) GoString() string {
	return s.String()
}

func (s *CreateNamespaceResponse) SetHeaders(v map[string]*string) *CreateNamespaceResponse {
	s.Headers = v
	return s
}

func (s *CreateNamespaceResponse) SetStatusCode(v int32) *CreateNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNamespaceResponse) SetBody(v *CreateNamespaceResponseBody) *CreateNamespaceResponse {
	s.Body = v
	return s
}

type DeleteConfigurationRequest struct {
	DataId      *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	Group       *string `json:"Group,omitempty" xml:"Group,omitempty"`
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s DeleteConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteConfigurationRequest) GoString() string {
	return s.String()
}

func (s *DeleteConfigurationRequest) SetDataId(v string) *DeleteConfigurationRequest {
	s.DataId = &v
	return s
}

func (s *DeleteConfigurationRequest) SetGroup(v string) *DeleteConfigurationRequest {
	s.Group = &v
	return s
}

func (s *DeleteConfigurationRequest) SetNamespaceId(v string) *DeleteConfigurationRequest {
	s.NamespaceId = &v
	return s
}

type DeleteConfigurationResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConfigurationResponseBody) SetCode(v string) *DeleteConfigurationResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteConfigurationResponseBody) SetMessage(v string) *DeleteConfigurationResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteConfigurationResponseBody) SetRequestId(v string) *DeleteConfigurationResponseBody {
	s.RequestId = &v
	return s
}

type DeleteConfigurationResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteConfigurationResponse) GoString() string {
	return s.String()
}

func (s *DeleteConfigurationResponse) SetHeaders(v map[string]*string) *DeleteConfigurationResponse {
	s.Headers = v
	return s
}

func (s *DeleteConfigurationResponse) SetStatusCode(v int32) *DeleteConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteConfigurationResponse) SetBody(v *DeleteConfigurationResponseBody) *DeleteConfigurationResponse {
	s.Body = v
	return s
}

type DeleteNamespaceRequest struct {
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s DeleteNamespaceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNamespaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteNamespaceRequest) SetNamespaceId(v string) *DeleteNamespaceRequest {
	s.NamespaceId = &v
	return s
}

type DeleteNamespaceResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNamespaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNamespaceResponseBody) SetCode(v string) *DeleteNamespaceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteNamespaceResponseBody) SetMessage(v string) *DeleteNamespaceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteNamespaceResponseBody) SetRequestId(v string) *DeleteNamespaceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteNamespaceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteNamespaceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteNamespaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteNamespaceResponse) SetHeaders(v map[string]*string) *DeleteNamespaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteNamespaceResponse) SetStatusCode(v int32) *DeleteNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNamespaceResponse) SetBody(v *DeleteNamespaceResponseBody) *DeleteNamespaceResponse {
	s.Body = v
	return s
}

type DeployConfigurationRequest struct {
	AppName     *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BetaIps     *string `json:"BetaIps,omitempty" xml:"BetaIps,omitempty"`
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	DataId      *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	Desc        *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Group       *string `json:"Group,omitempty" xml:"Group,omitempty"`
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	Tags        *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DeployConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeployConfigurationRequest) GoString() string {
	return s.String()
}

func (s *DeployConfigurationRequest) SetAppName(v string) *DeployConfigurationRequest {
	s.AppName = &v
	return s
}

func (s *DeployConfigurationRequest) SetBetaIps(v string) *DeployConfigurationRequest {
	s.BetaIps = &v
	return s
}

func (s *DeployConfigurationRequest) SetContent(v string) *DeployConfigurationRequest {
	s.Content = &v
	return s
}

func (s *DeployConfigurationRequest) SetDataId(v string) *DeployConfigurationRequest {
	s.DataId = &v
	return s
}

func (s *DeployConfigurationRequest) SetDesc(v string) *DeployConfigurationRequest {
	s.Desc = &v
	return s
}

func (s *DeployConfigurationRequest) SetGroup(v string) *DeployConfigurationRequest {
	s.Group = &v
	return s
}

func (s *DeployConfigurationRequest) SetNamespaceId(v string) *DeployConfigurationRequest {
	s.NamespaceId = &v
	return s
}

func (s *DeployConfigurationRequest) SetTags(v string) *DeployConfigurationRequest {
	s.Tags = &v
	return s
}

func (s *DeployConfigurationRequest) SetType(v string) *DeployConfigurationRequest {
	s.Type = &v
	return s
}

type DeployConfigurationResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeployConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeployConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *DeployConfigurationResponseBody) SetCode(v string) *DeployConfigurationResponseBody {
	s.Code = &v
	return s
}

func (s *DeployConfigurationResponseBody) SetMessage(v string) *DeployConfigurationResponseBody {
	s.Message = &v
	return s
}

func (s *DeployConfigurationResponseBody) SetRequestId(v string) *DeployConfigurationResponseBody {
	s.RequestId = &v
	return s
}

type DeployConfigurationResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeployConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeployConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeployConfigurationResponse) GoString() string {
	return s.String()
}

func (s *DeployConfigurationResponse) SetHeaders(v map[string]*string) *DeployConfigurationResponse {
	s.Headers = v
	return s
}

func (s *DeployConfigurationResponse) SetStatusCode(v int32) *DeployConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeployConfigurationResponse) SetBody(v *DeployConfigurationResponseBody) *DeployConfigurationResponse {
	s.Body = v
	return s
}

type DescribeConfigurationRequest struct {
	DataId      *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	Group       *string `json:"Group,omitempty" xml:"Group,omitempty"`
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s DescribeConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigurationRequest) GoString() string {
	return s.String()
}

func (s *DescribeConfigurationRequest) SetDataId(v string) *DescribeConfigurationRequest {
	s.DataId = &v
	return s
}

func (s *DescribeConfigurationRequest) SetGroup(v string) *DescribeConfigurationRequest {
	s.Group = &v
	return s
}

func (s *DescribeConfigurationRequest) SetNamespaceId(v string) *DescribeConfigurationRequest {
	s.NamespaceId = &v
	return s
}

type DescribeConfigurationResponseBody struct {
	Code          *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Configuration *DescribeConfigurationResponseBodyConfiguration `json:"Configuration,omitempty" xml:"Configuration,omitempty" type:"Struct"`
	Message       *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId     *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConfigurationResponseBody) SetCode(v string) *DescribeConfigurationResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeConfigurationResponseBody) SetConfiguration(v *DescribeConfigurationResponseBodyConfiguration) *DescribeConfigurationResponseBody {
	s.Configuration = v
	return s
}

func (s *DescribeConfigurationResponseBody) SetMessage(v string) *DescribeConfigurationResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeConfigurationResponseBody) SetRequestId(v string) *DescribeConfigurationResponseBody {
	s.RequestId = &v
	return s
}

type DescribeConfigurationResponseBodyConfiguration struct {
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	DataId  *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	Desc    *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Group   *string `json:"Group,omitempty" xml:"Group,omitempty"`
	Md5     *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Tags    *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Type    *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeConfigurationResponseBodyConfiguration) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigurationResponseBodyConfiguration) GoString() string {
	return s.String()
}

func (s *DescribeConfigurationResponseBodyConfiguration) SetAppName(v string) *DescribeConfigurationResponseBodyConfiguration {
	s.AppName = &v
	return s
}

func (s *DescribeConfigurationResponseBodyConfiguration) SetContent(v string) *DescribeConfigurationResponseBodyConfiguration {
	s.Content = &v
	return s
}

func (s *DescribeConfigurationResponseBodyConfiguration) SetDataId(v string) *DescribeConfigurationResponseBodyConfiguration {
	s.DataId = &v
	return s
}

func (s *DescribeConfigurationResponseBodyConfiguration) SetDesc(v string) *DescribeConfigurationResponseBodyConfiguration {
	s.Desc = &v
	return s
}

func (s *DescribeConfigurationResponseBodyConfiguration) SetGroup(v string) *DescribeConfigurationResponseBodyConfiguration {
	s.Group = &v
	return s
}

func (s *DescribeConfigurationResponseBodyConfiguration) SetMd5(v string) *DescribeConfigurationResponseBodyConfiguration {
	s.Md5 = &v
	return s
}

func (s *DescribeConfigurationResponseBodyConfiguration) SetTags(v string) *DescribeConfigurationResponseBodyConfiguration {
	s.Tags = &v
	return s
}

func (s *DescribeConfigurationResponseBodyConfiguration) SetType(v string) *DescribeConfigurationResponseBodyConfiguration {
	s.Type = &v
	return s
}

type DescribeConfigurationResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigurationResponse) GoString() string {
	return s.String()
}

func (s *DescribeConfigurationResponse) SetHeaders(v map[string]*string) *DescribeConfigurationResponse {
	s.Headers = v
	return s
}

func (s *DescribeConfigurationResponse) SetStatusCode(v int32) *DescribeConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeConfigurationResponse) SetBody(v *DescribeConfigurationResponseBody) *DescribeConfigurationResponse {
	s.Body = v
	return s
}

type DescribeImportFileUrlRequest struct {
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
}

func (s DescribeImportFileUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeImportFileUrlRequest) GoString() string {
	return s.String()
}

func (s *DescribeImportFileUrlRequest) SetContentType(v string) *DescribeImportFileUrlRequest {
	s.ContentType = &v
	return s
}

type DescribeImportFileUrlResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	FileUrl   *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeImportFileUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeImportFileUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImportFileUrlResponseBody) SetCode(v string) *DescribeImportFileUrlResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeImportFileUrlResponseBody) SetFileUrl(v string) *DescribeImportFileUrlResponseBody {
	s.FileUrl = &v
	return s
}

func (s *DescribeImportFileUrlResponseBody) SetMessage(v string) *DescribeImportFileUrlResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeImportFileUrlResponseBody) SetRequestId(v string) *DescribeImportFileUrlResponseBody {
	s.RequestId = &v
	return s
}

type DescribeImportFileUrlResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeImportFileUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeImportFileUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeImportFileUrlResponse) GoString() string {
	return s.String()
}

func (s *DescribeImportFileUrlResponse) SetHeaders(v map[string]*string) *DescribeImportFileUrlResponse {
	s.Headers = v
	return s
}

func (s *DescribeImportFileUrlResponse) SetStatusCode(v int32) *DescribeImportFileUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImportFileUrlResponse) SetBody(v *DescribeImportFileUrlResponseBody) *DescribeImportFileUrlResponse {
	s.Body = v
	return s
}

type DescribeNamespaceRequest struct {
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s DescribeNamespaceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespaceRequest) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceRequest) SetNamespaceId(v string) *DescribeNamespaceRequest {
	s.NamespaceId = &v
	return s
}

type DescribeNamespaceResponseBody struct {
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	Namespace *DescribeNamespaceResponseBodyNamespace `json:"Namespace,omitempty" xml:"Namespace,omitempty" type:"Struct"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNamespaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceResponseBody) SetCode(v string) *DescribeNamespaceResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeNamespaceResponseBody) SetMessage(v string) *DescribeNamespaceResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeNamespaceResponseBody) SetNamespace(v *DescribeNamespaceResponseBodyNamespace) *DescribeNamespaceResponseBody {
	s.Namespace = v
	return s
}

func (s *DescribeNamespaceResponseBody) SetRequestId(v string) *DescribeNamespaceResponseBody {
	s.RequestId = &v
	return s
}

type DescribeNamespaceResponseBodyNamespace struct {
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	Endpoint  *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
}

func (s DescribeNamespaceResponseBodyNamespace) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespaceResponseBodyNamespace) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceResponseBodyNamespace) SetAccessKey(v string) *DescribeNamespaceResponseBodyNamespace {
	s.AccessKey = &v
	return s
}

func (s *DescribeNamespaceResponseBodyNamespace) SetEndpoint(v string) *DescribeNamespaceResponseBodyNamespace {
	s.Endpoint = &v
	return s
}

func (s *DescribeNamespaceResponseBodyNamespace) SetName(v string) *DescribeNamespaceResponseBodyNamespace {
	s.Name = &v
	return s
}

func (s *DescribeNamespaceResponseBodyNamespace) SetRegionId(v string) *DescribeNamespaceResponseBodyNamespace {
	s.RegionId = &v
	return s
}

func (s *DescribeNamespaceResponseBodyNamespace) SetSecretKey(v string) *DescribeNamespaceResponseBodyNamespace {
	s.SecretKey = &v
	return s
}

type DescribeNamespaceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeNamespaceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespaceResponse) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceResponse) SetHeaders(v map[string]*string) *DescribeNamespaceResponse {
	s.Headers = v
	return s
}

func (s *DescribeNamespaceResponse) SetStatusCode(v int32) *DescribeNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNamespaceResponse) SetBody(v *DescribeNamespaceResponseBody) *DescribeNamespaceResponse {
	s.Body = v
	return s
}

type DescribeNamespacesResponseBody struct {
	Code       *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	Namespaces []*DescribeNamespacesResponseBodyNamespaces `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNamespacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespacesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNamespacesResponseBody) SetCode(v string) *DescribeNamespacesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeNamespacesResponseBody) SetMessage(v string) *DescribeNamespacesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeNamespacesResponseBody) SetNamespaces(v []*DescribeNamespacesResponseBodyNamespaces) *DescribeNamespacesResponseBody {
	s.Namespaces = v
	return s
}

func (s *DescribeNamespacesResponseBody) SetRequestId(v string) *DescribeNamespacesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeNamespacesResponseBodyNamespaces struct {
	ConfigCount   *int32  `json:"ConfigCount,omitempty" xml:"ConfigCount,omitempty"`
	NamespaceId   *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	Quota         *int32  `json:"Quota,omitempty" xml:"Quota,omitempty"`
	Type          *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeNamespacesResponseBodyNamespaces) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespacesResponseBodyNamespaces) GoString() string {
	return s.String()
}

func (s *DescribeNamespacesResponseBodyNamespaces) SetConfigCount(v int32) *DescribeNamespacesResponseBodyNamespaces {
	s.ConfigCount = &v
	return s
}

func (s *DescribeNamespacesResponseBodyNamespaces) SetNamespaceId(v string) *DescribeNamespacesResponseBodyNamespaces {
	s.NamespaceId = &v
	return s
}

func (s *DescribeNamespacesResponseBodyNamespaces) SetNamespaceName(v string) *DescribeNamespacesResponseBodyNamespaces {
	s.NamespaceName = &v
	return s
}

func (s *DescribeNamespacesResponseBodyNamespaces) SetQuota(v int32) *DescribeNamespacesResponseBodyNamespaces {
	s.Quota = &v
	return s
}

func (s *DescribeNamespacesResponseBodyNamespaces) SetType(v int32) *DescribeNamespacesResponseBodyNamespaces {
	s.Type = &v
	return s
}

type DescribeNamespacesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeNamespacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeNamespacesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespacesResponse) GoString() string {
	return s.String()
}

func (s *DescribeNamespacesResponse) SetHeaders(v map[string]*string) *DescribeNamespacesResponse {
	s.Headers = v
	return s
}

func (s *DescribeNamespacesResponse) SetStatusCode(v int32) *DescribeNamespacesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNamespacesResponse) SetBody(v *DescribeNamespacesResponseBody) *DescribeNamespacesResponse {
	s.Body = v
	return s
}

type DescribeNamespacesWithCreateResponseBody struct {
	Code       *string                                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	Namespaces []*DescribeNamespacesWithCreateResponseBodyNamespaces `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
	RequestId  *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNamespacesWithCreateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespacesWithCreateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNamespacesWithCreateResponseBody) SetCode(v string) *DescribeNamespacesWithCreateResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeNamespacesWithCreateResponseBody) SetMessage(v string) *DescribeNamespacesWithCreateResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeNamespacesWithCreateResponseBody) SetNamespaces(v []*DescribeNamespacesWithCreateResponseBodyNamespaces) *DescribeNamespacesWithCreateResponseBody {
	s.Namespaces = v
	return s
}

func (s *DescribeNamespacesWithCreateResponseBody) SetRequestId(v string) *DescribeNamespacesWithCreateResponseBody {
	s.RequestId = &v
	return s
}

type DescribeNamespacesWithCreateResponseBodyNamespaces struct {
	ConfigCount   *int32  `json:"ConfigCount,omitempty" xml:"ConfigCount,omitempty"`
	NamespaceId   *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	Quota         *int32  `json:"Quota,omitempty" xml:"Quota,omitempty"`
	Type          *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeNamespacesWithCreateResponseBodyNamespaces) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespacesWithCreateResponseBodyNamespaces) GoString() string {
	return s.String()
}

func (s *DescribeNamespacesWithCreateResponseBodyNamespaces) SetConfigCount(v int32) *DescribeNamespacesWithCreateResponseBodyNamespaces {
	s.ConfigCount = &v
	return s
}

func (s *DescribeNamespacesWithCreateResponseBodyNamespaces) SetNamespaceId(v string) *DescribeNamespacesWithCreateResponseBodyNamespaces {
	s.NamespaceId = &v
	return s
}

func (s *DescribeNamespacesWithCreateResponseBodyNamespaces) SetNamespaceName(v string) *DescribeNamespacesWithCreateResponseBodyNamespaces {
	s.NamespaceName = &v
	return s
}

func (s *DescribeNamespacesWithCreateResponseBodyNamespaces) SetQuota(v int32) *DescribeNamespacesWithCreateResponseBodyNamespaces {
	s.Quota = &v
	return s
}

func (s *DescribeNamespacesWithCreateResponseBodyNamespaces) SetType(v int32) *DescribeNamespacesWithCreateResponseBodyNamespaces {
	s.Type = &v
	return s
}

type DescribeNamespacesWithCreateResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeNamespacesWithCreateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeNamespacesWithCreateResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespacesWithCreateResponse) GoString() string {
	return s.String()
}

func (s *DescribeNamespacesWithCreateResponse) SetHeaders(v map[string]*string) *DescribeNamespacesWithCreateResponse {
	s.Headers = v
	return s
}

func (s *DescribeNamespacesWithCreateResponse) SetStatusCode(v int32) *DescribeNamespacesWithCreateResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNamespacesWithCreateResponse) SetBody(v *DescribeNamespacesWithCreateResponseBody) *DescribeNamespacesWithCreateResponse {
	s.Body = v
	return s
}

type DescribeTraceByConfigurationRequest struct {
	DataId      *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	EndTs       *string `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	Group       *string `json:"Group,omitempty" xml:"Group,omitempty"`
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	StartTs     *string `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
}

func (s DescribeTraceByConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTraceByConfigurationRequest) GoString() string {
	return s.String()
}

func (s *DescribeTraceByConfigurationRequest) SetDataId(v string) *DescribeTraceByConfigurationRequest {
	s.DataId = &v
	return s
}

func (s *DescribeTraceByConfigurationRequest) SetEndTs(v string) *DescribeTraceByConfigurationRequest {
	s.EndTs = &v
	return s
}

func (s *DescribeTraceByConfigurationRequest) SetGroup(v string) *DescribeTraceByConfigurationRequest {
	s.Group = &v
	return s
}

func (s *DescribeTraceByConfigurationRequest) SetNamespaceId(v string) *DescribeTraceByConfigurationRequest {
	s.NamespaceId = &v
	return s
}

func (s *DescribeTraceByConfigurationRequest) SetStartTs(v string) *DescribeTraceByConfigurationRequest {
	s.StartTs = &v
	return s
}

type DescribeTraceByConfigurationResponseBody struct {
	Code      *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Traces    []*DescribeTraceByConfigurationResponseBodyTraces `json:"Traces,omitempty" xml:"Traces,omitempty" type:"Repeated"`
}

func (s DescribeTraceByConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTraceByConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTraceByConfigurationResponseBody) SetCode(v string) *DescribeTraceByConfigurationResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeTraceByConfigurationResponseBody) SetMessage(v string) *DescribeTraceByConfigurationResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeTraceByConfigurationResponseBody) SetRequestId(v string) *DescribeTraceByConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTraceByConfigurationResponseBody) SetTraces(v []*DescribeTraceByConfigurationResponseBodyTraces) *DescribeTraceByConfigurationResponseBody {
	s.Traces = v
	return s
}

type DescribeTraceByConfigurationResponseBodyTraces struct {
	EventGroups []*DescribeTraceByConfigurationResponseBodyTracesEventGroups `json:"EventGroups,omitempty" xml:"EventGroups,omitempty" type:"Repeated"`
	Timestamp   *int64                                                       `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeTraceByConfigurationResponseBodyTraces) String() string {
	return tea.Prettify(s)
}

func (s DescribeTraceByConfigurationResponseBodyTraces) GoString() string {
	return s.String()
}

func (s *DescribeTraceByConfigurationResponseBodyTraces) SetEventGroups(v []*DescribeTraceByConfigurationResponseBodyTracesEventGroups) *DescribeTraceByConfigurationResponseBodyTraces {
	s.EventGroups = v
	return s
}

func (s *DescribeTraceByConfigurationResponseBodyTraces) SetTimestamp(v int64) *DescribeTraceByConfigurationResponseBodyTraces {
	s.Timestamp = &v
	return s
}

type DescribeTraceByConfigurationResponseBodyTracesEventGroups struct {
	EventDetails []*DescribeTraceByConfigurationResponseBodyTracesEventGroupsEventDetails `json:"EventDetails,omitempty" xml:"EventDetails,omitempty" type:"Repeated"`
	EventType    *string                                                                  `json:"EventType,omitempty" xml:"EventType,omitempty"`
}

func (s DescribeTraceByConfigurationResponseBodyTracesEventGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeTraceByConfigurationResponseBodyTracesEventGroups) GoString() string {
	return s.String()
}

func (s *DescribeTraceByConfigurationResponseBodyTracesEventGroups) SetEventDetails(v []*DescribeTraceByConfigurationResponseBodyTracesEventGroupsEventDetails) *DescribeTraceByConfigurationResponseBodyTracesEventGroups {
	s.EventDetails = v
	return s
}

func (s *DescribeTraceByConfigurationResponseBodyTracesEventGroups) SetEventType(v string) *DescribeTraceByConfigurationResponseBodyTracesEventGroups {
	s.EventType = &v
	return s
}

type DescribeTraceByConfigurationResponseBodyTracesEventGroupsEventDetails struct {
	DataId     *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	Delay      *string `json:"Delay,omitempty" xml:"Delay,omitempty"`
	Event      *string `json:"Event,omitempty" xml:"Event,omitempty"`
	Ext        *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	Group      *string `json:"Group,omitempty" xml:"Group,omitempty"`
	HandleIp   *string `json:"HandleIp,omitempty" xml:"HandleIp,omitempty"`
	LogDate    *string `json:"LogDate,omitempty" xml:"LogDate,omitempty"`
	RequestIp  *string `json:"RequestIp,omitempty" xml:"RequestIp,omitempty"`
	ResponseIp *string `json:"ResponseIp,omitempty" xml:"ResponseIp,omitempty"`
	Ts         *string `json:"Ts,omitempty" xml:"Ts,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeTraceByConfigurationResponseBodyTracesEventGroupsEventDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeTraceByConfigurationResponseBodyTracesEventGroupsEventDetails) GoString() string {
	return s.String()
}

func (s *DescribeTraceByConfigurationResponseBodyTracesEventGroupsEventDetails) SetDataId(v string) *DescribeTraceByConfigurationResponseBodyTracesEventGroupsEventDetails {
	s.DataId = &v
	return s
}

func (s *DescribeTraceByConfigurationResponseBodyTracesEventGroupsEventDetails) SetDelay(v string) *DescribeTraceByConfigurationResponseBodyTracesEventGroupsEventDetails {
	s.Delay = &v
	return s
}

func (s *DescribeTraceByConfigurationResponseBodyTracesEventGroupsEventDetails) SetEvent(v string) *DescribeTraceByConfigurationResponseBodyTracesEventGroupsEventDetails {
	s.Event = &v
	return s
}

func (s *DescribeTraceByConfigurationResponseBodyTracesEventGroupsEventDetails) SetExt(v string) *DescribeTraceByConfigurationResponseBodyTracesEventGroupsEventDetails {
	s.Ext = &v
	return s
}

func (s *DescribeTraceByConfigurationResponseBodyTracesEventGroupsEventDetails) SetGroup(v string) *DescribeTraceByConfigurationResponseBodyTracesEventGroupsEventDetails {
	s.Group = &v
	return s
}

func (s *DescribeTraceByConfigurationResponseBodyTracesEventGroupsEventDetails) SetHandleIp(v string) *DescribeTraceByConfigurationResponseBodyTracesEventGroupsEventDetails {
	s.HandleIp = &v
	return s
}

func (s *DescribeTraceByConfigurationResponseBodyTracesEventGroupsEventDetails) SetLogDate(v string) *DescribeTraceByConfigurationResponseBodyTracesEventGroupsEventDetails {
	s.LogDate = &v
	return s
}

func (s *DescribeTraceByConfigurationResponseBodyTracesEventGroupsEventDetails) SetRequestIp(v string) *DescribeTraceByConfigurationResponseBodyTracesEventGroupsEventDetails {
	s.RequestIp = &v
	return s
}

func (s *DescribeTraceByConfigurationResponseBodyTracesEventGroupsEventDetails) SetResponseIp(v string) *DescribeTraceByConfigurationResponseBodyTracesEventGroupsEventDetails {
	s.ResponseIp = &v
	return s
}

func (s *DescribeTraceByConfigurationResponseBodyTracesEventGroupsEventDetails) SetTs(v string) *DescribeTraceByConfigurationResponseBodyTracesEventGroupsEventDetails {
	s.Ts = &v
	return s
}

func (s *DescribeTraceByConfigurationResponseBodyTracesEventGroupsEventDetails) SetType(v string) *DescribeTraceByConfigurationResponseBodyTracesEventGroupsEventDetails {
	s.Type = &v
	return s
}

type DescribeTraceByConfigurationResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeTraceByConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTraceByConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTraceByConfigurationResponse) GoString() string {
	return s.String()
}

func (s *DescribeTraceByConfigurationResponse) SetHeaders(v map[string]*string) *DescribeTraceByConfigurationResponse {
	s.Headers = v
	return s
}

func (s *DescribeTraceByConfigurationResponse) SetStatusCode(v int32) *DescribeTraceByConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTraceByConfigurationResponse) SetBody(v *DescribeTraceByConfigurationResponseBody) *DescribeTraceByConfigurationResponse {
	s.Body = v
	return s
}

type UpdateNamespaceRequest struct {
	NamespaceId   *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
}

func (s UpdateNamespaceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateNamespaceRequest) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceRequest) SetNamespaceId(v string) *UpdateNamespaceRequest {
	s.NamespaceId = &v
	return s
}

func (s *UpdateNamespaceRequest) SetNamespaceName(v string) *UpdateNamespaceRequest {
	s.NamespaceName = &v
	return s
}

type UpdateNamespaceResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateNamespaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceResponseBody) SetCode(v string) *UpdateNamespaceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateNamespaceResponseBody) SetMessage(v string) *UpdateNamespaceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateNamespaceResponseBody) SetRequestId(v string) *UpdateNamespaceResponseBody {
	s.RequestId = &v
	return s
}

type UpdateNamespaceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateNamespaceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateNamespaceResponse) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceResponse) SetHeaders(v map[string]*string) *UpdateNamespaceResponse {
	s.Headers = v
	return s
}

func (s *UpdateNamespaceResponse) SetStatusCode(v int32) *UpdateNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNamespaceResponse) SetBody(v *UpdateNamespaceResponseBody) *UpdateNamespaceResponse {
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
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("acm"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) BatchExportConfigurations(request *BatchExportConfigurationsRequest) (_result *BatchExportConfigurationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchExportConfigurationsResponse{}
	_body, _err := client.BatchExportConfigurationsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchExportConfigurationsWithOptions(request *BatchExportConfigurationsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BatchExportConfigurationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Data)) {
		query["Data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceId)) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchExportConfigurations"),
		Version:     tea.String("2020-02-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/diamond-ops/pop/batch/export"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchExportConfigurationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchImportConfigurations(request *BatchImportConfigurationsRequest) (_result *BatchImportConfigurationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchImportConfigurationsResponse{}
	_body, _err := client.BatchImportConfigurationsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchImportConfigurationsWithOptions(request *BatchImportConfigurationsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BatchImportConfigurationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileUrl)) {
		body["FileUrl"] = request.FileUrl
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceId)) {
		body["NamespaceId"] = request.NamespaceId
	}

	if !tea.BoolValue(util.IsUnset(request.Policy)) {
		body["Policy"] = request.Policy
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchImportConfigurations"),
		Version:     tea.String("2020-02-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/diamond-ops/pop/batch/import"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchImportConfigurationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckConfigurationClone(request *CheckConfigurationCloneRequest) (_result *CheckConfigurationCloneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CheckConfigurationCloneResponse{}
	_body, _err := client.CheckConfigurationCloneWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckConfigurationCloneWithOptions(request *CheckConfigurationCloneRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CheckConfigurationCloneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Data)) {
		body["Data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceFrom)) {
		body["NamespaceFrom"] = request.NamespaceFrom
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceTo)) {
		body["NamespaceTo"] = request.NamespaceTo
	}

	if !tea.BoolValue(util.IsUnset(request.Policy)) {
		body["Policy"] = request.Policy
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckConfigurationClone"),
		Version:     tea.String("2020-02-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/diamond-ops/pop/batch/checkForClone"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckConfigurationCloneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckConfigurationExport(request *CheckConfigurationExportRequest) (_result *CheckConfigurationExportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CheckConfigurationExportResponse{}
	_body, _err := client.CheckConfigurationExportWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckConfigurationExportWithOptions(request *CheckConfigurationExportRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CheckConfigurationExportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Data)) {
		body["Data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceId)) {
		body["NamespaceId"] = request.NamespaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckConfigurationExport"),
		Version:     tea.String("2020-02-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/diamond-ops/pop/batch/checkForExport"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckConfigurationExportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CloneConfiguration(request *CloneConfigurationRequest) (_result *CloneConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CloneConfigurationResponse{}
	_body, _err := client.CloneConfigurationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CloneConfigurationWithOptions(request *CloneConfigurationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CloneConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Data)) {
		body["Data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceFrom)) {
		body["NamespaceFrom"] = request.NamespaceFrom
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceTo)) {
		body["NamespaceTo"] = request.NamespaceTo
	}

	if !tea.BoolValue(util.IsUnset(request.Policy)) {
		body["Policy"] = request.Policy
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CloneConfiguration"),
		Version:     tea.String("2020-02-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/diamond-ops/pop/batch/clone"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CloneConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateConfiguration(request *CreateConfigurationRequest) (_result *CreateConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateConfigurationResponse{}
	_body, _err := client.CreateConfigurationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateConfigurationWithOptions(request *CreateConfigurationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.DataId)) {
		body["DataId"] = request.DataId
	}

	if !tea.BoolValue(util.IsUnset(request.Desc)) {
		body["Desc"] = request.Desc
	}

	if !tea.BoolValue(util.IsUnset(request.Group)) {
		body["Group"] = request.Group
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceId)) {
		body["NamespaceId"] = request.NamespaceId
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		body["Tags"] = request.Tags
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateConfiguration"),
		Version:     tea.String("2020-02-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/diamond-ops/pop/configuration"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateNamespace(request *CreateNamespaceRequest) (_result *CreateNamespaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateNamespaceResponse{}
	_body, _err := client.CreateNamespaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateNamespaceWithOptions(request *CreateNamespaceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateNamespaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateNamespace"),
		Version:     tea.String("2020-02-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/diamond-ops/pop/namespace"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteConfiguration(request *DeleteConfigurationRequest) (_result *DeleteConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteConfigurationResponse{}
	_body, _err := client.DeleteConfigurationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteConfigurationWithOptions(request *DeleteConfigurationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataId)) {
		query["DataId"] = request.DataId
	}

	if !tea.BoolValue(util.IsUnset(request.Group)) {
		query["Group"] = request.Group
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceId)) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteConfiguration"),
		Version:     tea.String("2020-02-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/diamond-ops/pop/configuration"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteNamespace(request *DeleteNamespaceRequest) (_result *DeleteNamespaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteNamespaceResponse{}
	_body, _err := client.DeleteNamespaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteNamespaceWithOptions(request *DeleteNamespaceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteNamespaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NamespaceId)) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteNamespace"),
		Version:     tea.String("2020-02-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/diamond-ops/pop/namespace"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeployConfiguration(request *DeployConfigurationRequest) (_result *DeployConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeployConfigurationResponse{}
	_body, _err := client.DeployConfigurationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeployConfigurationWithOptions(request *DeployConfigurationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeployConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.BetaIps)) {
		body["BetaIps"] = request.BetaIps
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.DataId)) {
		body["DataId"] = request.DataId
	}

	if !tea.BoolValue(util.IsUnset(request.Desc)) {
		body["Desc"] = request.Desc
	}

	if !tea.BoolValue(util.IsUnset(request.Group)) {
		body["Group"] = request.Group
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceId)) {
		body["NamespaceId"] = request.NamespaceId
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		body["Tags"] = request.Tags
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeployConfiguration"),
		Version:     tea.String("2020-02-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/diamond-ops/pop/configuration"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeployConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeConfiguration(request *DescribeConfigurationRequest) (_result *DescribeConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeConfigurationResponse{}
	_body, _err := client.DescribeConfigurationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeConfigurationWithOptions(request *DescribeConfigurationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataId)) {
		query["DataId"] = request.DataId
	}

	if !tea.BoolValue(util.IsUnset(request.Group)) {
		query["Group"] = request.Group
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceId)) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeConfiguration"),
		Version:     tea.String("2020-02-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/diamond-ops/pop/configuration"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeImportFileUrl(request *DescribeImportFileUrlRequest) (_result *DescribeImportFileUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeImportFileUrlResponse{}
	_body, _err := client.DescribeImportFileUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeImportFileUrlWithOptions(request *DescribeImportFileUrlRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeImportFileUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContentType)) {
		query["ContentType"] = request.ContentType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeImportFileUrl"),
		Version:     tea.String("2020-02-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/diamond-ops/pop/batch/importFileUrl"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeImportFileUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNamespace(request *DescribeNamespaceRequest) (_result *DescribeNamespaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeNamespaceResponse{}
	_body, _err := client.DescribeNamespaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNamespaceWithOptions(request *DescribeNamespaceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeNamespaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NamespaceId)) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeNamespace"),
		Version:     tea.String("2020-02-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/diamond-ops/pop/namespace"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNamespaces() (_result *DescribeNamespacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeNamespacesResponse{}
	_body, _err := client.DescribeNamespacesWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNamespacesWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeNamespacesResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeNamespaces"),
		Version:     tea.String("2020-02-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/diamond-ops/pop/namespace/list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeNamespacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNamespacesWithCreate() (_result *DescribeNamespacesWithCreateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeNamespacesWithCreateResponse{}
	_body, _err := client.DescribeNamespacesWithCreateWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNamespacesWithCreateWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeNamespacesWithCreateResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeNamespacesWithCreate"),
		Version:     tea.String("2020-02-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/diamond-ops/pop/namespace/listWithCreate"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeNamespacesWithCreateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTraceByConfiguration(request *DescribeTraceByConfigurationRequest) (_result *DescribeTraceByConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeTraceByConfigurationResponse{}
	_body, _err := client.DescribeTraceByConfigurationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTraceByConfigurationWithOptions(request *DescribeTraceByConfigurationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeTraceByConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataId)) {
		query["DataId"] = request.DataId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTs)) {
		query["EndTs"] = request.EndTs
	}

	if !tea.BoolValue(util.IsUnset(request.Group)) {
		query["Group"] = request.Group
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceId)) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTs)) {
		query["StartTs"] = request.StartTs
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTraceByConfiguration"),
		Version:     tea.String("2020-02-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/diamond-ops/pop/trace/getByConfiguration"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTraceByConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateNamespace(request *UpdateNamespaceRequest) (_result *UpdateNamespaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateNamespaceResponse{}
	_body, _err := client.UpdateNamespaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateNamespaceWithOptions(request *UpdateNamespaceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateNamespaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NamespaceId)) {
		body["NamespaceId"] = request.NamespaceId
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceName)) {
		body["NamespaceName"] = request.NamespaceName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateNamespace"),
		Version:     tea.String("2020-02-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/diamond-ops/pop/namespace"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
