// This file is auto-generated, don't edit it. Thanks.
package client

import (
	gatewayclient "github.com/alibabacloud-go/alibabacloud-gateway-pop/client"
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

type AyncTradeDocumentPackageExtractSmartAppRequest struct {
	CustomExtractionRange []*string `json:"CustomExtractionRange,omitempty" xml:"CustomExtractionRange,omitempty" type:"Repeated"`
	FileName              *string   `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// This parameter is required.
	FileUrl      *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	Option       *string `json:"Option,omitempty" xml:"Option,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s AyncTradeDocumentPackageExtractSmartAppRequest) String() string {
	return tea.Prettify(s)
}

func (s AyncTradeDocumentPackageExtractSmartAppRequest) GoString() string {
	return s.String()
}

func (s *AyncTradeDocumentPackageExtractSmartAppRequest) SetCustomExtractionRange(v []*string) *AyncTradeDocumentPackageExtractSmartAppRequest {
	s.CustomExtractionRange = v
	return s
}

func (s *AyncTradeDocumentPackageExtractSmartAppRequest) SetFileName(v string) *AyncTradeDocumentPackageExtractSmartAppRequest {
	s.FileName = &v
	return s
}

func (s *AyncTradeDocumentPackageExtractSmartAppRequest) SetFileUrl(v string) *AyncTradeDocumentPackageExtractSmartAppRequest {
	s.FileUrl = &v
	return s
}

func (s *AyncTradeDocumentPackageExtractSmartAppRequest) SetOption(v string) *AyncTradeDocumentPackageExtractSmartAppRequest {
	s.Option = &v
	return s
}

func (s *AyncTradeDocumentPackageExtractSmartAppRequest) SetTemplateName(v string) *AyncTradeDocumentPackageExtractSmartAppRequest {
	s.TemplateName = &v
	return s
}

type AyncTradeDocumentPackageExtractSmartAppShrinkRequest struct {
	CustomExtractionRangeShrink *string `json:"CustomExtractionRange,omitempty" xml:"CustomExtractionRange,omitempty"`
	FileName                    *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// This parameter is required.
	FileUrl      *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	Option       *string `json:"Option,omitempty" xml:"Option,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s AyncTradeDocumentPackageExtractSmartAppShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AyncTradeDocumentPackageExtractSmartAppShrinkRequest) GoString() string {
	return s.String()
}

func (s *AyncTradeDocumentPackageExtractSmartAppShrinkRequest) SetCustomExtractionRangeShrink(v string) *AyncTradeDocumentPackageExtractSmartAppShrinkRequest {
	s.CustomExtractionRangeShrink = &v
	return s
}

func (s *AyncTradeDocumentPackageExtractSmartAppShrinkRequest) SetFileName(v string) *AyncTradeDocumentPackageExtractSmartAppShrinkRequest {
	s.FileName = &v
	return s
}

func (s *AyncTradeDocumentPackageExtractSmartAppShrinkRequest) SetFileUrl(v string) *AyncTradeDocumentPackageExtractSmartAppShrinkRequest {
	s.FileUrl = &v
	return s
}

func (s *AyncTradeDocumentPackageExtractSmartAppShrinkRequest) SetOption(v string) *AyncTradeDocumentPackageExtractSmartAppShrinkRequest {
	s.Option = &v
	return s
}

func (s *AyncTradeDocumentPackageExtractSmartAppShrinkRequest) SetTemplateName(v string) *AyncTradeDocumentPackageExtractSmartAppShrinkRequest {
	s.TemplateName = &v
	return s
}

type AyncTradeDocumentPackageExtractSmartAppResponseBody struct {
	Completed  *bool       `json:"Completed,omitempty" xml:"Completed,omitempty"`
	CreateTime *string     `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Data       interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AyncTradeDocumentPackageExtractSmartAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AyncTradeDocumentPackageExtractSmartAppResponseBody) GoString() string {
	return s.String()
}

func (s *AyncTradeDocumentPackageExtractSmartAppResponseBody) SetCompleted(v bool) *AyncTradeDocumentPackageExtractSmartAppResponseBody {
	s.Completed = &v
	return s
}

func (s *AyncTradeDocumentPackageExtractSmartAppResponseBody) SetCreateTime(v string) *AyncTradeDocumentPackageExtractSmartAppResponseBody {
	s.CreateTime = &v
	return s
}

func (s *AyncTradeDocumentPackageExtractSmartAppResponseBody) SetData(v interface{}) *AyncTradeDocumentPackageExtractSmartAppResponseBody {
	s.Data = v
	return s
}

func (s *AyncTradeDocumentPackageExtractSmartAppResponseBody) SetRequestId(v string) *AyncTradeDocumentPackageExtractSmartAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *AyncTradeDocumentPackageExtractSmartAppResponseBody) SetStatus(v string) *AyncTradeDocumentPackageExtractSmartAppResponseBody {
	s.Status = &v
	return s
}

func (s *AyncTradeDocumentPackageExtractSmartAppResponseBody) SetSuccess(v bool) *AyncTradeDocumentPackageExtractSmartAppResponseBody {
	s.Success = &v
	return s
}

type AyncTradeDocumentPackageExtractSmartAppResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AyncTradeDocumentPackageExtractSmartAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AyncTradeDocumentPackageExtractSmartAppResponse) String() string {
	return tea.Prettify(s)
}

func (s AyncTradeDocumentPackageExtractSmartAppResponse) GoString() string {
	return s.String()
}

func (s *AyncTradeDocumentPackageExtractSmartAppResponse) SetHeaders(v map[string]*string) *AyncTradeDocumentPackageExtractSmartAppResponse {
	s.Headers = v
	return s
}

func (s *AyncTradeDocumentPackageExtractSmartAppResponse) SetStatusCode(v int32) *AyncTradeDocumentPackageExtractSmartAppResponse {
	s.StatusCode = &v
	return s
}

func (s *AyncTradeDocumentPackageExtractSmartAppResponse) SetBody(v *AyncTradeDocumentPackageExtractSmartAppResponseBody) *AyncTradeDocumentPackageExtractSmartAppResponse {
	s.Body = v
	return s
}

type GetDocParserResultRequest struct {
	// example:
	//
	// docmind-20220816-1e89d65c
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	LayoutNum      *int32  `json:"LayoutNum,omitempty" xml:"LayoutNum,omitempty"`
	LayoutStepSize *int32  `json:"LayoutStepSize,omitempty" xml:"LayoutStepSize,omitempty"`
}

func (s GetDocParserResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDocParserResultRequest) GoString() string {
	return s.String()
}

func (s *GetDocParserResultRequest) SetId(v string) *GetDocParserResultRequest {
	s.Id = &v
	return s
}

func (s *GetDocParserResultRequest) SetLayoutNum(v int32) *GetDocParserResultRequest {
	s.LayoutNum = &v
	return s
}

func (s *GetDocParserResultRequest) SetLayoutStepSize(v int32) *GetDocParserResultRequest {
	s.LayoutStepSize = &v
	return s
}

type GetDocParserResultResponseBody struct {
	Code      *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDocParserResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDocParserResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocParserResultResponseBody) SetCode(v string) *GetDocParserResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetDocParserResultResponseBody) SetData(v map[string]interface{}) *GetDocParserResultResponseBody {
	s.Data = v
	return s
}

func (s *GetDocParserResultResponseBody) SetMessage(v string) *GetDocParserResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetDocParserResultResponseBody) SetRequestId(v string) *GetDocParserResultResponseBody {
	s.RequestId = &v
	return s
}

type GetDocParserResultResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocParserResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocParserResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDocParserResultResponse) GoString() string {
	return s.String()
}

func (s *GetDocParserResultResponse) SetHeaders(v map[string]*string) *GetDocParserResultResponse {
	s.Headers = v
	return s
}

func (s *GetDocParserResultResponse) SetStatusCode(v int32) *GetDocParserResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocParserResultResponse) SetBody(v *GetDocParserResultResponseBody) *GetDocParserResultResponse {
	s.Body = v
	return s
}

type GetDocStructureResultRequest struct {
	// example:
	//
	// docmind-20220816-1e89d65c
	Id                 *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ImageStrategy      *string `json:"ImageStrategy,omitempty" xml:"ImageStrategy,omitempty"`
	RevealMarkdown     *bool   `json:"RevealMarkdown,omitempty" xml:"RevealMarkdown,omitempty"`
	UseUrlResponseBody *bool   `json:"UseUrlResponseBody,omitempty" xml:"UseUrlResponseBody,omitempty"`
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

func (s *GetDocStructureResultRequest) SetImageStrategy(v string) *GetDocStructureResultRequest {
	s.ImageStrategy = &v
	return s
}

func (s *GetDocStructureResultRequest) SetRevealMarkdown(v bool) *GetDocStructureResultRequest {
	s.RevealMarkdown = &v
	return s
}

func (s *GetDocStructureResultRequest) SetUseUrlResponseBody(v bool) *GetDocStructureResultRequest {
	s.UseUrlResponseBody = &v
	return s
}

type GetDocStructureResultResponseBody struct {
	// example:
	//
	// noPermission
	Code      *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Completed *bool                  `json:"Completed,omitempty" xml:"Completed,omitempty"`
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocStructureResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// noPermission
	Code      *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Completed *bool       `json:"Completed,omitempty" xml:"Completed,omitempty"`
	Data      interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocumentCompareResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// docmind-20220816-1e89d65c
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
	// example:
	//
	// noPermission
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// This parameter is required.
	Completed *bool                                       `json:"Completed,omitempty" xml:"Completed,omitempty"`
	Data      []*GetDocumentConvertResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	// example:
	//
	// e6d83e55df218650b9a296bfbc300076
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// example:
	//
	// 2355965
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// pdf
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// http://docmind-api-cn-hangzhou.oss-cn-hangzhou.aliyuncs.com/convert/docmind-20220816-15bc7965/0.pdf?Expires=1660722412&OSSAccessKeyId=LTAI5tFEK2uEApeeYzxNMEci&Signature=f%2FKluINWMuuVyA5w22Z1wkoRjEg%3D
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocumentConvertResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// docmind-20220816-1e89d65c
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
	// example:
	//
	// noPermission
	Code      *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Completed *bool                  `json:"Completed,omitempty" xml:"Completed,omitempty"`
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocumentExtractResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type GetPageNumRequest struct {
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s GetPageNumRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPageNumRequest) GoString() string {
	return s.String()
}

func (s *GetPageNumRequest) SetBizId(v string) *GetPageNumRequest {
	s.BizId = &v
	return s
}

type GetPageNumResponseBody struct {
	Data         *GetPageNumResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode    *string                     `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                     `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	HttpCode     *string                     `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	RequestId    *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPageNumResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPageNumResponseBody) GoString() string {
	return s.String()
}

func (s *GetPageNumResponseBody) SetData(v *GetPageNumResponseBodyData) *GetPageNumResponseBody {
	s.Data = v
	return s
}

func (s *GetPageNumResponseBody) SetErrorCode(v string) *GetPageNumResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPageNumResponseBody) SetErrorMessage(v string) *GetPageNumResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetPageNumResponseBody) SetHttpCode(v string) *GetPageNumResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetPageNumResponseBody) SetRequestId(v string) *GetPageNumResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPageNumResponseBody) SetSuccess(v bool) *GetPageNumResponseBody {
	s.Success = &v
	return s
}

type GetPageNumResponseBodyData struct {
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
}

func (s GetPageNumResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetPageNumResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPageNumResponseBodyData) SetPageNum(v int32) *GetPageNumResponseBodyData {
	s.PageNum = &v
	return s
}

type GetPageNumResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPageNumResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPageNumResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPageNumResponse) GoString() string {
	return s.String()
}

func (s *GetPageNumResponse) SetHeaders(v map[string]*string) *GetPageNumResponse {
	s.Headers = v
	return s
}

func (s *GetPageNumResponse) SetStatusCode(v int32) *GetPageNumResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPageNumResponse) SetBody(v *GetPageNumResponseBody) *GetPageNumResponse {
	s.Body = v
	return s
}

type GetTableUnderstandingResultRequest struct {
	// example:
	//
	// docmind-20220816-1e89d65c
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
	// example:
	//
	// noPermission
	Code      *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Completed *bool                  `json:"Completed,omitempty" xml:"Completed,omitempty"`
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTableUnderstandingResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type QueryDocParserStatusRequest struct {
	// example:
	//
	// docmind-20220816-1e89d65c
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s QueryDocParserStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDocParserStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryDocParserStatusRequest) SetId(v string) *QueryDocParserStatusRequest {
	s.Id = &v
	return s
}

type QueryDocParserStatusResponseBody struct {
	// example:
	//
	// noPermission
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *QueryDocParserStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryDocParserStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDocParserStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDocParserStatusResponseBody) SetCode(v string) *QueryDocParserStatusResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDocParserStatusResponseBody) SetData(v *QueryDocParserStatusResponseBodyData) *QueryDocParserStatusResponseBody {
	s.Data = v
	return s
}

func (s *QueryDocParserStatusResponseBody) SetMessage(v string) *QueryDocParserStatusResponseBody {
	s.Message = &v
	return s
}

func (s *QueryDocParserStatusResponseBody) SetRequestId(v string) *QueryDocParserStatusResponseBody {
	s.RequestId = &v
	return s
}

type QueryDocParserStatusResponseBodyData struct {
	ImageCount                *int32  `json:"ImageCount,omitempty" xml:"ImageCount,omitempty"`
	NumberOfSuccessfulParsing *int32  `json:"NumberOfSuccessfulParsing,omitempty" xml:"NumberOfSuccessfulParsing,omitempty"`
	PageCountEstimate         *int32  `json:"PageCountEstimate,omitempty" xml:"PageCountEstimate,omitempty"`
	ParagraphCount            *int32  `json:"ParagraphCount,omitempty" xml:"ParagraphCount,omitempty"`
	Status                    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TableCount                *int32  `json:"TableCount,omitempty" xml:"TableCount,omitempty"`
	Tokens                    *int64  `json:"Tokens,omitempty" xml:"Tokens,omitempty"`
}

func (s QueryDocParserStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryDocParserStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryDocParserStatusResponseBodyData) SetImageCount(v int32) *QueryDocParserStatusResponseBodyData {
	s.ImageCount = &v
	return s
}

func (s *QueryDocParserStatusResponseBodyData) SetNumberOfSuccessfulParsing(v int32) *QueryDocParserStatusResponseBodyData {
	s.NumberOfSuccessfulParsing = &v
	return s
}

func (s *QueryDocParserStatusResponseBodyData) SetPageCountEstimate(v int32) *QueryDocParserStatusResponseBodyData {
	s.PageCountEstimate = &v
	return s
}

func (s *QueryDocParserStatusResponseBodyData) SetParagraphCount(v int32) *QueryDocParserStatusResponseBodyData {
	s.ParagraphCount = &v
	return s
}

func (s *QueryDocParserStatusResponseBodyData) SetStatus(v string) *QueryDocParserStatusResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryDocParserStatusResponseBodyData) SetTableCount(v int32) *QueryDocParserStatusResponseBodyData {
	s.TableCount = &v
	return s
}

func (s *QueryDocParserStatusResponseBodyData) SetTokens(v int64) *QueryDocParserStatusResponseBodyData {
	s.Tokens = &v
	return s
}

type QueryDocParserStatusResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDocParserStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDocParserStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDocParserStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryDocParserStatusResponse) SetHeaders(v map[string]*string) *QueryDocParserStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryDocParserStatusResponse) SetStatusCode(v int32) *QueryDocParserStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDocParserStatusResponse) SetBody(v *QueryDocParserStatusResponseBody) *QueryDocParserStatusResponse {
	s.Body = v
	return s
}

type SubmitConvertImageToExcelJobRequest struct {
	ForceMergeExcel *bool `json:"ForceMergeExcel,omitempty" xml:"ForceMergeExcel,omitempty"`
	// example:
	//
	// jpg
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
	ForceMergeExcel *bool `json:"ForceMergeExcel,omitempty" xml:"ForceMergeExcel,omitempty"`
	// example:
	//
	// jpg
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
	// example:
	//
	// noPermission
	Code *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitConvertImageToExcelJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// docmind-20220810-7c5f9dd4
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitConvertImageToExcelJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type SubmitConvertImageToMarkdownJobRequest struct {
	// example:
	//
	// jpg
	ImageNameExtension *string   `json:"ImageNameExtension,omitempty" xml:"ImageNameExtension,omitempty"`
	ImageNames         []*string `json:"ImageNames,omitempty" xml:"ImageNames,omitempty" type:"Repeated"`
	ImageUrls          []*string `json:"ImageUrls,omitempty" xml:"ImageUrls,omitempty" type:"Repeated"`
}

func (s SubmitConvertImageToMarkdownJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitConvertImageToMarkdownJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitConvertImageToMarkdownJobRequest) SetImageNameExtension(v string) *SubmitConvertImageToMarkdownJobRequest {
	s.ImageNameExtension = &v
	return s
}

func (s *SubmitConvertImageToMarkdownJobRequest) SetImageNames(v []*string) *SubmitConvertImageToMarkdownJobRequest {
	s.ImageNames = v
	return s
}

func (s *SubmitConvertImageToMarkdownJobRequest) SetImageUrls(v []*string) *SubmitConvertImageToMarkdownJobRequest {
	s.ImageUrls = v
	return s
}

type SubmitConvertImageToMarkdownJobShrinkRequest struct {
	// example:
	//
	// jpg
	ImageNameExtension *string `json:"ImageNameExtension,omitempty" xml:"ImageNameExtension,omitempty"`
	ImageNamesShrink   *string `json:"ImageNames,omitempty" xml:"ImageNames,omitempty"`
	ImageUrlsShrink    *string `json:"ImageUrls,omitempty" xml:"ImageUrls,omitempty"`
}

func (s SubmitConvertImageToMarkdownJobShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitConvertImageToMarkdownJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitConvertImageToMarkdownJobShrinkRequest) SetImageNameExtension(v string) *SubmitConvertImageToMarkdownJobShrinkRequest {
	s.ImageNameExtension = &v
	return s
}

func (s *SubmitConvertImageToMarkdownJobShrinkRequest) SetImageNamesShrink(v string) *SubmitConvertImageToMarkdownJobShrinkRequest {
	s.ImageNamesShrink = &v
	return s
}

func (s *SubmitConvertImageToMarkdownJobShrinkRequest) SetImageUrlsShrink(v string) *SubmitConvertImageToMarkdownJobShrinkRequest {
	s.ImageUrlsShrink = &v
	return s
}

type SubmitConvertImageToMarkdownJobResponseBody struct {
	// example:
	//
	// noPermission
	Code *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitConvertImageToMarkdownJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitConvertImageToMarkdownJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitConvertImageToMarkdownJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitConvertImageToMarkdownJobResponseBody) SetCode(v string) *SubmitConvertImageToMarkdownJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitConvertImageToMarkdownJobResponseBody) SetData(v *SubmitConvertImageToMarkdownJobResponseBodyData) *SubmitConvertImageToMarkdownJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitConvertImageToMarkdownJobResponseBody) SetMessage(v string) *SubmitConvertImageToMarkdownJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitConvertImageToMarkdownJobResponseBody) SetRequestId(v string) *SubmitConvertImageToMarkdownJobResponseBody {
	s.RequestId = &v
	return s
}

type SubmitConvertImageToMarkdownJobResponseBodyData struct {
	// example:
	//
	// docmind-20220810-7c5f9dd4
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SubmitConvertImageToMarkdownJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitConvertImageToMarkdownJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitConvertImageToMarkdownJobResponseBodyData) SetId(v string) *SubmitConvertImageToMarkdownJobResponseBodyData {
	s.Id = &v
	return s
}

type SubmitConvertImageToMarkdownJobResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitConvertImageToMarkdownJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitConvertImageToMarkdownJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitConvertImageToMarkdownJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitConvertImageToMarkdownJobResponse) SetHeaders(v map[string]*string) *SubmitConvertImageToMarkdownJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitConvertImageToMarkdownJobResponse) SetStatusCode(v int32) *SubmitConvertImageToMarkdownJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitConvertImageToMarkdownJobResponse) SetBody(v *SubmitConvertImageToMarkdownJobResponseBody) *SubmitConvertImageToMarkdownJobResponse {
	s.Body = v
	return s
}

type SubmitConvertImageToPdfJobRequest struct {
	// example:
	//
	// JPG
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
	// example:
	//
	// JPG
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
	// example:
	//
	// noPermission
	Code *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitConvertImageToPdfJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// docmind-20220810-7c5f9dd4
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitConvertImageToPdfJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// jpg
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
	// example:
	//
	// jpg
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
	// example:
	//
	// noPermission
	Code *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitConvertImageToWordJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// docmind-20220810-7c5f9dd4
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitConvertImageToWordJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// convertPdfToExcel.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// https://gw.alipayobjects.com/os/basement_prod/598b9edf-5287-4065-9e36-464305c60698.pdf
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
	// example:
	//
	// convertPdfToExcel.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// https://gw.alipayobjects.com/os/basement_prod/598b9edf-5287-4065-9e36-464305c60698.pdf
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
	// example:
	//
	// noPermission
	Code *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitConvertPdfToExcelJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// docmind-20220816-15bc7965
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitConvertPdfToExcelJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// convertPdfToImage.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// https://gw.alipayobjects.com/os/basement_prod/598b9edf-5287-4065-9e36-464305c60698.pdf
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
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
	// example:
	//
	// convertPdfToImage.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// https://gw.alipayobjects.com/os/basement_prod/598b9edf-5287-4065-9e36-464305c60698.pdf
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
	// example:
	//
	// noPermission
	Code *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitConvertPdfToImageJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// docmind-20220810-7c5f9dd4
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitConvertPdfToImageJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type SubmitConvertPdfToMarkdownJobRequest struct {
	// example:
	//
	// convertPdfToExcel.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// https://gw.alipayobjects.com/os/basement_prod/598b9edf-5287-4065-9e36-464305c60698.pdf
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s SubmitConvertPdfToMarkdownJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitConvertPdfToMarkdownJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitConvertPdfToMarkdownJobRequest) SetFileName(v string) *SubmitConvertPdfToMarkdownJobRequest {
	s.FileName = &v
	return s
}

func (s *SubmitConvertPdfToMarkdownJobRequest) SetFileUrl(v string) *SubmitConvertPdfToMarkdownJobRequest {
	s.FileUrl = &v
	return s
}

type SubmitConvertPdfToMarkdownJobAdvanceRequest struct {
	// example:
	//
	// convertPdfToExcel.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// https://gw.alipayobjects.com/os/basement_prod/598b9edf-5287-4065-9e36-464305c60698.pdf
	FileUrlObject io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s SubmitConvertPdfToMarkdownJobAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitConvertPdfToMarkdownJobAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SubmitConvertPdfToMarkdownJobAdvanceRequest) SetFileName(v string) *SubmitConvertPdfToMarkdownJobAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *SubmitConvertPdfToMarkdownJobAdvanceRequest) SetFileUrlObject(v io.Reader) *SubmitConvertPdfToMarkdownJobAdvanceRequest {
	s.FileUrlObject = v
	return s
}

type SubmitConvertPdfToMarkdownJobResponseBody struct {
	// example:
	//
	// noPermission
	Code *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitConvertPdfToMarkdownJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitConvertPdfToMarkdownJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitConvertPdfToMarkdownJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitConvertPdfToMarkdownJobResponseBody) SetCode(v string) *SubmitConvertPdfToMarkdownJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitConvertPdfToMarkdownJobResponseBody) SetData(v *SubmitConvertPdfToMarkdownJobResponseBodyData) *SubmitConvertPdfToMarkdownJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitConvertPdfToMarkdownJobResponseBody) SetMessage(v string) *SubmitConvertPdfToMarkdownJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitConvertPdfToMarkdownJobResponseBody) SetRequestId(v string) *SubmitConvertPdfToMarkdownJobResponseBody {
	s.RequestId = &v
	return s
}

type SubmitConvertPdfToMarkdownJobResponseBodyData struct {
	// example:
	//
	// docmind-20220816-15bc7965
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SubmitConvertPdfToMarkdownJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitConvertPdfToMarkdownJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitConvertPdfToMarkdownJobResponseBodyData) SetId(v string) *SubmitConvertPdfToMarkdownJobResponseBodyData {
	s.Id = &v
	return s
}

type SubmitConvertPdfToMarkdownJobResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitConvertPdfToMarkdownJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitConvertPdfToMarkdownJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitConvertPdfToMarkdownJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitConvertPdfToMarkdownJobResponse) SetHeaders(v map[string]*string) *SubmitConvertPdfToMarkdownJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitConvertPdfToMarkdownJobResponse) SetStatusCode(v int32) *SubmitConvertPdfToMarkdownJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitConvertPdfToMarkdownJobResponse) SetBody(v *SubmitConvertPdfToMarkdownJobResponseBody) *SubmitConvertPdfToMarkdownJobResponse {
	s.Body = v
	return s
}

type SubmitConvertPdfToWordJobRequest struct {
	// example:
	//
	// covertPdfToWord.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// https://gw.alipayobjects.com/os/basement_prod/598b9edf-5287-4065-9e36-464305c60698.pdf
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
	// example:
	//
	// covertPdfToWord.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// https://gw.alipayobjects.com/os/basement_prod/598b9edf-5287-4065-9e36-464305c60698.pdf
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
	// example:
	//
	// noPermission
	Code *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitConvertPdfToWordJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// docmind-20220816-15bc7965
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitConvertPdfToWordJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type SubmitDigitalDocStructureJobRequest struct {
	// example:
	//
	// docStructure.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// pdf
	FileNameExtension *string `json:"FileNameExtension,omitempty" xml:"FileNameExtension,omitempty"`
	// example:
	//
	// https://gw.alipayobjects.com/os/basement_prod/598b9edf-5287-4065-9e36-464305c60698.pdf
	FileUrl            *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	ImageStrategy      *string `json:"ImageStrategy,omitempty" xml:"ImageStrategy,omitempty"`
	RevealMarkdown     *bool   `json:"RevealMarkdown,omitempty" xml:"RevealMarkdown,omitempty"`
	UseUrlResponseBody *bool   `json:"UseUrlResponseBody,omitempty" xml:"UseUrlResponseBody,omitempty"`
}

func (s SubmitDigitalDocStructureJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitDigitalDocStructureJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitDigitalDocStructureJobRequest) SetFileName(v string) *SubmitDigitalDocStructureJobRequest {
	s.FileName = &v
	return s
}

func (s *SubmitDigitalDocStructureJobRequest) SetFileNameExtension(v string) *SubmitDigitalDocStructureJobRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitDigitalDocStructureJobRequest) SetFileUrl(v string) *SubmitDigitalDocStructureJobRequest {
	s.FileUrl = &v
	return s
}

func (s *SubmitDigitalDocStructureJobRequest) SetImageStrategy(v string) *SubmitDigitalDocStructureJobRequest {
	s.ImageStrategy = &v
	return s
}

func (s *SubmitDigitalDocStructureJobRequest) SetRevealMarkdown(v bool) *SubmitDigitalDocStructureJobRequest {
	s.RevealMarkdown = &v
	return s
}

func (s *SubmitDigitalDocStructureJobRequest) SetUseUrlResponseBody(v bool) *SubmitDigitalDocStructureJobRequest {
	s.UseUrlResponseBody = &v
	return s
}

type SubmitDigitalDocStructureJobAdvanceRequest struct {
	// example:
	//
	// docStructure.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// pdf
	FileNameExtension *string `json:"FileNameExtension,omitempty" xml:"FileNameExtension,omitempty"`
	// example:
	//
	// https://gw.alipayobjects.com/os/basement_prod/598b9edf-5287-4065-9e36-464305c60698.pdf
	FileUrlObject      io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	ImageStrategy      *string   `json:"ImageStrategy,omitempty" xml:"ImageStrategy,omitempty"`
	RevealMarkdown     *bool     `json:"RevealMarkdown,omitempty" xml:"RevealMarkdown,omitempty"`
	UseUrlResponseBody *bool     `json:"UseUrlResponseBody,omitempty" xml:"UseUrlResponseBody,omitempty"`
}

func (s SubmitDigitalDocStructureJobAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitDigitalDocStructureJobAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SubmitDigitalDocStructureJobAdvanceRequest) SetFileName(v string) *SubmitDigitalDocStructureJobAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *SubmitDigitalDocStructureJobAdvanceRequest) SetFileNameExtension(v string) *SubmitDigitalDocStructureJobAdvanceRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitDigitalDocStructureJobAdvanceRequest) SetFileUrlObject(v io.Reader) *SubmitDigitalDocStructureJobAdvanceRequest {
	s.FileUrlObject = v
	return s
}

func (s *SubmitDigitalDocStructureJobAdvanceRequest) SetImageStrategy(v string) *SubmitDigitalDocStructureJobAdvanceRequest {
	s.ImageStrategy = &v
	return s
}

func (s *SubmitDigitalDocStructureJobAdvanceRequest) SetRevealMarkdown(v bool) *SubmitDigitalDocStructureJobAdvanceRequest {
	s.RevealMarkdown = &v
	return s
}

func (s *SubmitDigitalDocStructureJobAdvanceRequest) SetUseUrlResponseBody(v bool) *SubmitDigitalDocStructureJobAdvanceRequest {
	s.UseUrlResponseBody = &v
	return s
}

type SubmitDigitalDocStructureJobResponseBody struct {
	// example:
	//
	// noPermission
	Code *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Id   *string     `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SubmitDigitalDocStructureJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitDigitalDocStructureJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitDigitalDocStructureJobResponseBody) SetCode(v string) *SubmitDigitalDocStructureJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitDigitalDocStructureJobResponseBody) SetData(v interface{}) *SubmitDigitalDocStructureJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitDigitalDocStructureJobResponseBody) SetId(v string) *SubmitDigitalDocStructureJobResponseBody {
	s.Id = &v
	return s
}

func (s *SubmitDigitalDocStructureJobResponseBody) SetMessage(v string) *SubmitDigitalDocStructureJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitDigitalDocStructureJobResponseBody) SetRequestId(v string) *SubmitDigitalDocStructureJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitDigitalDocStructureJobResponseBody) SetStatus(v string) *SubmitDigitalDocStructureJobResponseBody {
	s.Status = &v
	return s
}

type SubmitDigitalDocStructureJobResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitDigitalDocStructureJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitDigitalDocStructureJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitDigitalDocStructureJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitDigitalDocStructureJobResponse) SetHeaders(v map[string]*string) *SubmitDigitalDocStructureJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitDigitalDocStructureJobResponse) SetStatusCode(v int32) *SubmitDigitalDocStructureJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitDigitalDocStructureJobResponse) SetBody(v *SubmitDigitalDocStructureJobResponseBody) *SubmitDigitalDocStructureJobResponse {
	s.Body = v
	return s
}

type SubmitDocParserJobRequest struct {
	// example:
	//
	// docStructure.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// pdf
	FileNameExtension *string `json:"FileNameExtension,omitempty" xml:"FileNameExtension,omitempty"`
	// example:
	//
	// https://gw.alipayobjects.com/os/basement_prod/598b9edf-5287-4065-9e36-464305c60698.pdf
	FileUrl            *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	FormulaEnhancement *bool   `json:"FormulaEnhancement,omitempty" xml:"FormulaEnhancement,omitempty"`
}

func (s SubmitDocParserJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocParserJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitDocParserJobRequest) SetFileName(v string) *SubmitDocParserJobRequest {
	s.FileName = &v
	return s
}

func (s *SubmitDocParserJobRequest) SetFileNameExtension(v string) *SubmitDocParserJobRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitDocParserJobRequest) SetFileUrl(v string) *SubmitDocParserJobRequest {
	s.FileUrl = &v
	return s
}

func (s *SubmitDocParserJobRequest) SetFormulaEnhancement(v bool) *SubmitDocParserJobRequest {
	s.FormulaEnhancement = &v
	return s
}

type SubmitDocParserJobAdvanceRequest struct {
	// example:
	//
	// docStructure.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// pdf
	FileNameExtension *string `json:"FileNameExtension,omitempty" xml:"FileNameExtension,omitempty"`
	// example:
	//
	// https://gw.alipayobjects.com/os/basement_prod/598b9edf-5287-4065-9e36-464305c60698.pdf
	FileUrlObject      io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	FormulaEnhancement *bool     `json:"FormulaEnhancement,omitempty" xml:"FormulaEnhancement,omitempty"`
}

func (s SubmitDocParserJobAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocParserJobAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SubmitDocParserJobAdvanceRequest) SetFileName(v string) *SubmitDocParserJobAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *SubmitDocParserJobAdvanceRequest) SetFileNameExtension(v string) *SubmitDocParserJobAdvanceRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitDocParserJobAdvanceRequest) SetFileUrlObject(v io.Reader) *SubmitDocParserJobAdvanceRequest {
	s.FileUrlObject = v
	return s
}

func (s *SubmitDocParserJobAdvanceRequest) SetFormulaEnhancement(v bool) *SubmitDocParserJobAdvanceRequest {
	s.FormulaEnhancement = &v
	return s
}

type SubmitDocParserJobResponseBody struct {
	// example:
	//
	// noPermission
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitDocParserJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitDocParserJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocParserJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitDocParserJobResponseBody) SetCode(v string) *SubmitDocParserJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitDocParserJobResponseBody) SetData(v *SubmitDocParserJobResponseBodyData) *SubmitDocParserJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitDocParserJobResponseBody) SetMessage(v string) *SubmitDocParserJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitDocParserJobResponseBody) SetRequestId(v string) *SubmitDocParserJobResponseBody {
	s.RequestId = &v
	return s
}

type SubmitDocParserJobResponseBodyData struct {
	// example:
	//
	// docmind-20220816-15bc7965
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SubmitDocParserJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocParserJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitDocParserJobResponseBodyData) SetId(v string) *SubmitDocParserJobResponseBodyData {
	s.Id = &v
	return s
}

type SubmitDocParserJobResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitDocParserJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitDocParserJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocParserJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitDocParserJobResponse) SetHeaders(v map[string]*string) *SubmitDocParserJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitDocParserJobResponse) SetStatusCode(v int32) *SubmitDocParserJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitDocParserJobResponse) SetBody(v *SubmitDocParserJobResponseBody) *SubmitDocParserJobResponse {
	s.Body = v
	return s
}

type SubmitDocStructureJobRequest struct {
	AllowPptFormat *bool `json:"AllowPptFormat,omitempty" xml:"AllowPptFormat,omitempty"`
	// example:
	//
	// docStructure.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// pdf
	FileNameExtension *string `json:"FileNameExtension,omitempty" xml:"FileNameExtension,omitempty"`
	// example:
	//
	// https://gw.alipayobjects.com/os/basement_prod/598b9edf-5287-4065-9e36-464305c60698.pdf
	FileUrl            *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	FormulaEnhancement *bool   `json:"FormulaEnhancement,omitempty" xml:"FormulaEnhancement,omitempty"`
	StructureType      *string `json:"StructureType,omitempty" xml:"StructureType,omitempty"`
}

func (s SubmitDocStructureJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocStructureJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitDocStructureJobRequest) SetAllowPptFormat(v bool) *SubmitDocStructureJobRequest {
	s.AllowPptFormat = &v
	return s
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

func (s *SubmitDocStructureJobRequest) SetFormulaEnhancement(v bool) *SubmitDocStructureJobRequest {
	s.FormulaEnhancement = &v
	return s
}

func (s *SubmitDocStructureJobRequest) SetStructureType(v string) *SubmitDocStructureJobRequest {
	s.StructureType = &v
	return s
}

type SubmitDocStructureJobAdvanceRequest struct {
	AllowPptFormat *bool `json:"AllowPptFormat,omitempty" xml:"AllowPptFormat,omitempty"`
	// example:
	//
	// docStructure.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// pdf
	FileNameExtension *string `json:"FileNameExtension,omitempty" xml:"FileNameExtension,omitempty"`
	// example:
	//
	// https://gw.alipayobjects.com/os/basement_prod/598b9edf-5287-4065-9e36-464305c60698.pdf
	FileUrlObject      io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	FormulaEnhancement *bool     `json:"FormulaEnhancement,omitempty" xml:"FormulaEnhancement,omitempty"`
	StructureType      *string   `json:"StructureType,omitempty" xml:"StructureType,omitempty"`
}

func (s SubmitDocStructureJobAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocStructureJobAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SubmitDocStructureJobAdvanceRequest) SetAllowPptFormat(v bool) *SubmitDocStructureJobAdvanceRequest {
	s.AllowPptFormat = &v
	return s
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

func (s *SubmitDocStructureJobAdvanceRequest) SetFormulaEnhancement(v bool) *SubmitDocStructureJobAdvanceRequest {
	s.FormulaEnhancement = &v
	return s
}

func (s *SubmitDocStructureJobAdvanceRequest) SetStructureType(v string) *SubmitDocStructureJobAdvanceRequest {
	s.StructureType = &v
	return s
}

type SubmitDocStructureJobResponseBody struct {
	// example:
	//
	// noPermission
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitDocStructureJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// docmind-20220816-15bc7965
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitDocStructureJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type SubmitDocumentExtractJobRequest struct {
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// pdf
	FileNameExtension *string `json:"FileNameExtension,omitempty" xml:"FileNameExtension,omitempty"`
	// example:
	//
	// https://gw.alipayobjects.com/os/basement_prod/598b9edf-5287-4065-9e36-464305c60698.pdf
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
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
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// pdf
	FileNameExtension *string `json:"FileNameExtension,omitempty" xml:"FileNameExtension,omitempty"`
	// example:
	//
	// https://gw.alipayobjects.com/os/basement_prod/598b9edf-5287-4065-9e36-464305c60698.pdf
	FileUrlObject io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
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
	// example:
	//
	// noPermission
	Code *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitDocumentExtractJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// docmind-20220816-15bc7965
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitDocumentExtractJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// pdf
	FileNameExtension *string `json:"FileNameExtension,omitempty" xml:"FileNameExtension,omitempty"`
	// example:
	//
	// https://gw.alipayobjects.com/os/basement_prod/598b9edf-5287-4065-9e36-464305c60698.pdf
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
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
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// pdf
	FileNameExtension *string `json:"FileNameExtension,omitempty" xml:"FileNameExtension,omitempty"`
	// example:
	//
	// https://gw.alipayobjects.com/os/basement_prod/598b9edf-5287-4065-9e36-464305c60698.pdf
	FileUrlObject io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
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
	// example:
	//
	// noPermission
	Code *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitTableUnderstandingJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// docmind-20220816-15bc7965
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitTableUnderstandingJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	client.ProductId = tea.String("docmind-api")
	gatewayClient, _err := gatewayclient.NewClient()
	if _err != nil {
		return _err
	}

	client.Spi = gatewayClient
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

// Summary:
//
// 整票识别
//
// @param tmpReq - AyncTradeDocumentPackageExtractSmartAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AyncTradeDocumentPackageExtractSmartAppResponse
func (client *Client) AyncTradeDocumentPackageExtractSmartAppWithOptions(tmpReq *AyncTradeDocumentPackageExtractSmartAppRequest, runtime *util.RuntimeOptions) (_result *AyncTradeDocumentPackageExtractSmartAppResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AyncTradeDocumentPackageExtractSmartAppShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CustomExtractionRange)) {
		request.CustomExtractionRangeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomExtractionRange, tea.String("CustomExtractionRange"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomExtractionRangeShrink)) {
		query["CustomExtractionRange"] = request.CustomExtractionRangeShrink
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["FileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.FileUrl)) {
		query["FileUrl"] = request.FileUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Option)) {
		query["Option"] = request.Option
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AyncTradeDocumentPackageExtractSmartApp"),
		Version:     tea.String("2022-07-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &AyncTradeDocumentPackageExtractSmartAppResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &AyncTradeDocumentPackageExtractSmartAppResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 整票识别
//
// @param request - AyncTradeDocumentPackageExtractSmartAppRequest
//
// @return AyncTradeDocumentPackageExtractSmartAppResponse
func (client *Client) AyncTradeDocumentPackageExtractSmartApp(request *AyncTradeDocumentPackageExtractSmartAppRequest) (_result *AyncTradeDocumentPackageExtractSmartAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AyncTradeDocumentPackageExtractSmartAppResponse{}
	_body, _err := client.AyncTradeDocumentPackageExtractSmartAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 文档结构化流式接口
//
// @param request - GetDocParserResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocParserResultResponse
func (client *Client) GetDocParserResultWithOptions(request *GetDocParserResultRequest, runtime *util.RuntimeOptions) (_result *GetDocParserResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.LayoutNum)) {
		query["LayoutNum"] = request.LayoutNum
	}

	if !tea.BoolValue(util.IsUnset(request.LayoutStepSize)) {
		query["LayoutStepSize"] = request.LayoutStepSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDocParserResult"),
		Version:     tea.String("2022-07-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetDocParserResultResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetDocParserResultResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 文档结构化流式接口
//
// @param request - GetDocParserResultRequest
//
// @return GetDocParserResultResponse
func (client *Client) GetDocParserResult(request *GetDocParserResultRequest) (_result *GetDocParserResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDocParserResultResponse{}
	_body, _err := client.GetDocParserResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 文档智能解析结果查询
//
// @param request - GetDocStructureResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocStructureResultResponse
func (client *Client) GetDocStructureResultWithOptions(request *GetDocStructureResultRequest, runtime *util.RuntimeOptions) (_result *GetDocStructureResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ImageStrategy)) {
		query["ImageStrategy"] = request.ImageStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.RevealMarkdown)) {
		query["RevealMarkdown"] = request.RevealMarkdown
	}

	if !tea.BoolValue(util.IsUnset(request.UseUrlResponseBody)) {
		query["UseUrlResponseBody"] = request.UseUrlResponseBody
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetDocStructureResultResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetDocStructureResultResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 文档智能解析结果查询
//
// @param request - GetDocStructureResultRequest
//
// @return GetDocStructureResultResponse
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

// Summary:
//
// 文档对比结果查询
//
// @param request - GetDocumentCompareResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentCompareResultResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetDocumentCompareResultResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetDocumentCompareResultResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 文档对比结果查询
//
// @param request - GetDocumentCompareResultRequest
//
// @return GetDocumentCompareResultResponse
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

// Summary:
//
// 文档转换结果查询
//
// @param request - GetDocumentConvertResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentConvertResultResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetDocumentConvertResultResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetDocumentConvertResultResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 文档转换结果查询
//
// @param request - GetDocumentConvertResultRequest
//
// @return GetDocumentConvertResultResponse
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

// Summary:
//
// 文档抽取结果查询
//
// @param request - GetDocumentExtractResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentExtractResultResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetDocumentExtractResultResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetDocumentExtractResultResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 文档抽取结果查询
//
// @param request - GetDocumentExtractResultRequest
//
// @return GetDocumentExtractResultResponse
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

// Summary:
//
// openmind
//
// @param request - GetPageNumRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPageNumResponse
func (client *Client) GetPageNumWithOptions(request *GetPageNumRequest, runtime *util.RuntimeOptions) (_result *GetPageNumResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPageNum"),
		Version:     tea.String("2022-07-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetPageNumResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetPageNumResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// openmind
//
// @param request - GetPageNumRequest
//
// @return GetPageNumResponse
func (client *Client) GetPageNum(request *GetPageNumRequest) (_result *GetPageNumResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPageNumResponse{}
	_body, _err := client.GetPageNumWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 表格智能解析结果查询
//
// @param request - GetTableUnderstandingResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableUnderstandingResultResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetTableUnderstandingResultResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetTableUnderstandingResultResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 表格智能解析结果查询
//
// @param request - GetTableUnderstandingResultRequest
//
// @return GetTableUnderstandingResultResponse
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

// Summary:
//
// 获取文档智能解析处理状态
//
// @param request - QueryDocParserStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDocParserStatusResponse
func (client *Client) QueryDocParserStatusWithOptions(request *QueryDocParserStatusRequest, runtime *util.RuntimeOptions) (_result *QueryDocParserStatusResponse, _err error) {
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
		Action:      tea.String("QueryDocParserStatus"),
		Version:     tea.String("2022-07-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &QueryDocParserStatusResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &QueryDocParserStatusResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取文档智能解析处理状态
//
// @param request - QueryDocParserStatusRequest
//
// @return QueryDocParserStatusResponse
func (client *Client) QueryDocParserStatus(request *QueryDocParserStatusRequest) (_result *QueryDocParserStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDocParserStatusResponse{}
	_body, _err := client.QueryDocParserStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 图片转excel
//
// @param tmpReq - SubmitConvertImageToExcelJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitConvertImageToExcelJobResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &SubmitConvertImageToExcelJobResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &SubmitConvertImageToExcelJobResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 图片转excel
//
// @param request - SubmitConvertImageToExcelJobRequest
//
// @return SubmitConvertImageToExcelJobResponse
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

// Summary:
//
// 图片转markdown
//
// @param tmpReq - SubmitConvertImageToMarkdownJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitConvertImageToMarkdownJobResponse
func (client *Client) SubmitConvertImageToMarkdownJobWithOptions(tmpReq *SubmitConvertImageToMarkdownJobRequest, runtime *util.RuntimeOptions) (_result *SubmitConvertImageToMarkdownJobResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SubmitConvertImageToMarkdownJobShrinkRequest{}
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
		Action:      tea.String("SubmitConvertImageToMarkdownJob"),
		Version:     tea.String("2022-07-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &SubmitConvertImageToMarkdownJobResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &SubmitConvertImageToMarkdownJobResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 图片转markdown
//
// @param request - SubmitConvertImageToMarkdownJobRequest
//
// @return SubmitConvertImageToMarkdownJobResponse
func (client *Client) SubmitConvertImageToMarkdownJob(request *SubmitConvertImageToMarkdownJobRequest) (_result *SubmitConvertImageToMarkdownJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitConvertImageToMarkdownJobResponse{}
	_body, _err := client.SubmitConvertImageToMarkdownJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 图片转pdf
//
// @param tmpReq - SubmitConvertImageToPdfJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitConvertImageToPdfJobResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &SubmitConvertImageToPdfJobResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &SubmitConvertImageToPdfJobResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 图片转pdf
//
// @param request - SubmitConvertImageToPdfJobRequest
//
// @return SubmitConvertImageToPdfJobResponse
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

// Summary:
//
// 图片转word
//
// @param tmpReq - SubmitConvertImageToWordJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitConvertImageToWordJobResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &SubmitConvertImageToWordJobResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &SubmitConvertImageToWordJobResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 图片转word
//
// @param request - SubmitConvertImageToWordJobRequest
//
// @return SubmitConvertImageToWordJobResponse
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

// Summary:
//
// pdf转excel
//
// @param request - SubmitConvertPdfToExcelJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitConvertPdfToExcelJobResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &SubmitConvertPdfToExcelJobResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &SubmitConvertPdfToExcelJobResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// pdf转excel
//
// @param request - SubmitConvertPdfToExcelJobRequest
//
// @return SubmitConvertPdfToExcelJobResponse
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
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
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
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

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

// Summary:
//
// pdf转图片
//
// @param request - SubmitConvertPdfToImageJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitConvertPdfToImageJobResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &SubmitConvertPdfToImageJobResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &SubmitConvertPdfToImageJobResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// pdf转图片
//
// @param request - SubmitConvertPdfToImageJobRequest
//
// @return SubmitConvertPdfToImageJobResponse
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
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
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
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

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

// Summary:
//
// pdf转markdown
//
// @param request - SubmitConvertPdfToMarkdownJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitConvertPdfToMarkdownJobResponse
func (client *Client) SubmitConvertPdfToMarkdownJobWithOptions(request *SubmitConvertPdfToMarkdownJobRequest, runtime *util.RuntimeOptions) (_result *SubmitConvertPdfToMarkdownJobResponse, _err error) {
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
		Action:      tea.String("SubmitConvertPdfToMarkdownJob"),
		Version:     tea.String("2022-07-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &SubmitConvertPdfToMarkdownJobResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &SubmitConvertPdfToMarkdownJobResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// pdf转markdown
//
// @param request - SubmitConvertPdfToMarkdownJobRequest
//
// @return SubmitConvertPdfToMarkdownJobResponse
func (client *Client) SubmitConvertPdfToMarkdownJob(request *SubmitConvertPdfToMarkdownJobRequest) (_result *SubmitConvertPdfToMarkdownJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitConvertPdfToMarkdownJobResponse{}
	_body, _err := client.SubmitConvertPdfToMarkdownJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitConvertPdfToMarkdownJobAdvance(request *SubmitConvertPdfToMarkdownJobAdvanceRequest, runtime *util.RuntimeOptions) (_result *SubmitConvertPdfToMarkdownJobResponse, _err error) {
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
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
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
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	submitConvertPdfToMarkdownJobReq := &SubmitConvertPdfToMarkdownJobRequest{}
	openapiutil.Convert(request, submitConvertPdfToMarkdownJobReq)
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
		submitConvertPdfToMarkdownJobReq.FileUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	submitConvertPdfToMarkdownJobResp, _err := client.SubmitConvertPdfToMarkdownJobWithOptions(submitConvertPdfToMarkdownJobReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = submitConvertPdfToMarkdownJobResp
	return _result, _err
}

// Summary:
//
// pdf转word
//
// @param request - SubmitConvertPdfToWordJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitConvertPdfToWordJobResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &SubmitConvertPdfToWordJobResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &SubmitConvertPdfToWordJobResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// pdf转word
//
// @param request - SubmitConvertPdfToWordJobRequest
//
// @return SubmitConvertPdfToWordJobResponse
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
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
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
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

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

// Summary:
//
// 电子解析
//
// @param request - SubmitDigitalDocStructureJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitDigitalDocStructureJobResponse
func (client *Client) SubmitDigitalDocStructureJobWithOptions(request *SubmitDigitalDocStructureJobRequest, runtime *util.RuntimeOptions) (_result *SubmitDigitalDocStructureJobResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ImageStrategy)) {
		query["ImageStrategy"] = request.ImageStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.RevealMarkdown)) {
		query["RevealMarkdown"] = request.RevealMarkdown
	}

	if !tea.BoolValue(util.IsUnset(request.UseUrlResponseBody)) {
		query["UseUrlResponseBody"] = request.UseUrlResponseBody
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitDigitalDocStructureJob"),
		Version:     tea.String("2022-07-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &SubmitDigitalDocStructureJobResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &SubmitDigitalDocStructureJobResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 电子解析
//
// @param request - SubmitDigitalDocStructureJobRequest
//
// @return SubmitDigitalDocStructureJobResponse
func (client *Client) SubmitDigitalDocStructureJob(request *SubmitDigitalDocStructureJobRequest) (_result *SubmitDigitalDocStructureJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitDigitalDocStructureJobResponse{}
	_body, _err := client.SubmitDigitalDocStructureJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitDigitalDocStructureJobAdvance(request *SubmitDigitalDocStructureJobAdvanceRequest, runtime *util.RuntimeOptions) (_result *SubmitDigitalDocStructureJobResponse, _err error) {
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
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
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
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	submitDigitalDocStructureJobReq := &SubmitDigitalDocStructureJobRequest{}
	openapiutil.Convert(request, submitDigitalDocStructureJobReq)
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
		submitDigitalDocStructureJobReq.FileUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	submitDigitalDocStructureJobResp, _err := client.SubmitDigitalDocStructureJobWithOptions(submitDigitalDocStructureJobReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = submitDigitalDocStructureJobResp
	return _result, _err
}

// Summary:
//
// 文档智能解析流式输出
//
// @param request - SubmitDocParserJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitDocParserJobResponse
func (client *Client) SubmitDocParserJobWithOptions(request *SubmitDocParserJobRequest, runtime *util.RuntimeOptions) (_result *SubmitDocParserJobResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.FormulaEnhancement)) {
		query["FormulaEnhancement"] = request.FormulaEnhancement
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitDocParserJob"),
		Version:     tea.String("2022-07-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &SubmitDocParserJobResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &SubmitDocParserJobResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 文档智能解析流式输出
//
// @param request - SubmitDocParserJobRequest
//
// @return SubmitDocParserJobResponse
func (client *Client) SubmitDocParserJob(request *SubmitDocParserJobRequest) (_result *SubmitDocParserJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitDocParserJobResponse{}
	_body, _err := client.SubmitDocParserJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitDocParserJobAdvance(request *SubmitDocParserJobAdvanceRequest, runtime *util.RuntimeOptions) (_result *SubmitDocParserJobResponse, _err error) {
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
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
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
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	submitDocParserJobReq := &SubmitDocParserJobRequest{}
	openapiutil.Convert(request, submitDocParserJobReq)
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
		submitDocParserJobReq.FileUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	submitDocParserJobResp, _err := client.SubmitDocParserJobWithOptions(submitDocParserJobReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = submitDocParserJobResp
	return _result, _err
}

// Summary:
//
// 文档智能解析
//
// @param request - SubmitDocStructureJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitDocStructureJobResponse
func (client *Client) SubmitDocStructureJobWithOptions(request *SubmitDocStructureJobRequest, runtime *util.RuntimeOptions) (_result *SubmitDocStructureJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllowPptFormat)) {
		query["AllowPptFormat"] = request.AllowPptFormat
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["FileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.FileNameExtension)) {
		query["FileNameExtension"] = request.FileNameExtension
	}

	if !tea.BoolValue(util.IsUnset(request.FileUrl)) {
		query["FileUrl"] = request.FileUrl
	}

	if !tea.BoolValue(util.IsUnset(request.FormulaEnhancement)) {
		query["FormulaEnhancement"] = request.FormulaEnhancement
	}

	if !tea.BoolValue(util.IsUnset(request.StructureType)) {
		query["StructureType"] = request.StructureType
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &SubmitDocStructureJobResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &SubmitDocStructureJobResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 文档智能解析
//
// @param request - SubmitDocStructureJobRequest
//
// @return SubmitDocStructureJobResponse
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
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
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
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

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

// Summary:
//
// 文档抽取
//
// @param request - SubmitDocumentExtractJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitDocumentExtractJobResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &SubmitDocumentExtractJobResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &SubmitDocumentExtractJobResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 文档抽取
//
// @param request - SubmitDocumentExtractJobRequest
//
// @return SubmitDocumentExtractJobResponse
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
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
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
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

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

// Summary:
//
// 表格智能解析
//
// @param request - SubmitTableUnderstandingJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitTableUnderstandingJobResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &SubmitTableUnderstandingJobResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &SubmitTableUnderstandingJobResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 表格智能解析
//
// @param request - SubmitTableUnderstandingJobRequest
//
// @return SubmitTableUnderstandingJobResponse
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
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
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
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

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
