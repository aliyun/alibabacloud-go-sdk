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

type DescribeAdvicesRequest struct {
	AdviceId        *int64  `json:"AdviceId,omitempty" xml:"AdviceId,omitempty"`
	CheckId         *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	ExcludeAdviceId *int64  `json:"ExcludeAdviceId,omitempty" xml:"ExcludeAdviceId,omitempty"`
	Language        *string `json:"Language,omitempty" xml:"Language,omitempty"`
	Product         *string `json:"Product,omitempty" xml:"Product,omitempty"`
	ResourceId      *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s DescribeAdvicesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdvicesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAdvicesRequest) SetAdviceId(v int64) *DescribeAdvicesRequest {
	s.AdviceId = &v
	return s
}

func (s *DescribeAdvicesRequest) SetCheckId(v string) *DescribeAdvicesRequest {
	s.CheckId = &v
	return s
}

func (s *DescribeAdvicesRequest) SetExcludeAdviceId(v int64) *DescribeAdvicesRequest {
	s.ExcludeAdviceId = &v
	return s
}

func (s *DescribeAdvicesRequest) SetLanguage(v string) *DescribeAdvicesRequest {
	s.Language = &v
	return s
}

func (s *DescribeAdvicesRequest) SetProduct(v string) *DescribeAdvicesRequest {
	s.Product = &v
	return s
}

func (s *DescribeAdvicesRequest) SetResourceId(v string) *DescribeAdvicesRequest {
	s.ResourceId = &v
	return s
}

type DescribeAdvicesResponseBody struct {
	Data      *DescribeAdvicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAdvicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdvicesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAdvicesResponseBody) SetData(v *DescribeAdvicesResponseBodyData) *DescribeAdvicesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAdvicesResponseBody) SetRequestId(v string) *DescribeAdvicesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAdvicesResponseBodyData struct {
	Advice []*DescribeAdvicesResponseBodyDataAdvice `json:"Advice,omitempty" xml:"Advice,omitempty" type:"Repeated"`
}

func (s DescribeAdvicesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdvicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAdvicesResponseBodyData) SetAdvice(v []*DescribeAdvicesResponseBodyDataAdvice) *DescribeAdvicesResponseBodyData {
	s.Advice = v
	return s
}

type DescribeAdvicesResponseBodyDataAdvice struct {
	AliyunId    *int64  `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
	CheckId     *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	CheckName   *string `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreated  *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// ID
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	IsExpired  *bool   `json:"IsExpired,omitempty" xml:"IsExpired,omitempty"`
	Product    *string `json:"Product,omitempty" xml:"Product,omitempty"`
	Resource   *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Severity   *int32  `json:"Severity,omitempty" xml:"Severity,omitempty"`
}

func (s DescribeAdvicesResponseBodyDataAdvice) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdvicesResponseBodyDataAdvice) GoString() string {
	return s.String()
}

func (s *DescribeAdvicesResponseBodyDataAdvice) SetAliyunId(v int64) *DescribeAdvicesResponseBodyDataAdvice {
	s.AliyunId = &v
	return s
}

func (s *DescribeAdvicesResponseBodyDataAdvice) SetCheckId(v string) *DescribeAdvicesResponseBodyDataAdvice {
	s.CheckId = &v
	return s
}

func (s *DescribeAdvicesResponseBodyDataAdvice) SetCheckName(v string) *DescribeAdvicesResponseBodyDataAdvice {
	s.CheckName = &v
	return s
}

func (s *DescribeAdvicesResponseBodyDataAdvice) SetContent(v string) *DescribeAdvicesResponseBodyDataAdvice {
	s.Content = &v
	return s
}

func (s *DescribeAdvicesResponseBodyDataAdvice) SetDescription(v string) *DescribeAdvicesResponseBodyDataAdvice {
	s.Description = &v
	return s
}

func (s *DescribeAdvicesResponseBodyDataAdvice) SetGmtCreated(v string) *DescribeAdvicesResponseBodyDataAdvice {
	s.GmtCreated = &v
	return s
}

func (s *DescribeAdvicesResponseBodyDataAdvice) SetGmtModified(v string) *DescribeAdvicesResponseBodyDataAdvice {
	s.GmtModified = &v
	return s
}

func (s *DescribeAdvicesResponseBodyDataAdvice) SetId(v int64) *DescribeAdvicesResponseBodyDataAdvice {
	s.Id = &v
	return s
}

func (s *DescribeAdvicesResponseBodyDataAdvice) SetIsExpired(v bool) *DescribeAdvicesResponseBodyDataAdvice {
	s.IsExpired = &v
	return s
}

func (s *DescribeAdvicesResponseBodyDataAdvice) SetProduct(v string) *DescribeAdvicesResponseBodyDataAdvice {
	s.Product = &v
	return s
}

func (s *DescribeAdvicesResponseBodyDataAdvice) SetResource(v string) *DescribeAdvicesResponseBodyDataAdvice {
	s.Resource = &v
	return s
}

func (s *DescribeAdvicesResponseBodyDataAdvice) SetResourceId(v string) *DescribeAdvicesResponseBodyDataAdvice {
	s.ResourceId = &v
	return s
}

func (s *DescribeAdvicesResponseBodyDataAdvice) SetSeverity(v int32) *DescribeAdvicesResponseBodyDataAdvice {
	s.Severity = &v
	return s
}

type DescribeAdvicesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAdvicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAdvicesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdvicesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAdvicesResponse) SetHeaders(v map[string]*string) *DescribeAdvicesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAdvicesResponse) SetStatusCode(v int32) *DescribeAdvicesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAdvicesResponse) SetBody(v *DescribeAdvicesResponseBody) *DescribeAdvicesResponse {
	s.Body = v
	return s
}

type DescribeAdvicesFlatPageRequest struct {
	AdviceId   *int64  `json:"AdviceId,omitempty" xml:"AdviceId,omitempty"`
	CheckId    *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	Language   *string `json:"Language,omitempty" xml:"Language,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Product    *string `json:"Product,omitempty" xml:"Product,omitempty"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s DescribeAdvicesFlatPageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdvicesFlatPageRequest) GoString() string {
	return s.String()
}

func (s *DescribeAdvicesFlatPageRequest) SetAdviceId(v int64) *DescribeAdvicesFlatPageRequest {
	s.AdviceId = &v
	return s
}

func (s *DescribeAdvicesFlatPageRequest) SetCheckId(v string) *DescribeAdvicesFlatPageRequest {
	s.CheckId = &v
	return s
}

func (s *DescribeAdvicesFlatPageRequest) SetLanguage(v string) *DescribeAdvicesFlatPageRequest {
	s.Language = &v
	return s
}

func (s *DescribeAdvicesFlatPageRequest) SetPageNumber(v int32) *DescribeAdvicesFlatPageRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAdvicesFlatPageRequest) SetPageSize(v int32) *DescribeAdvicesFlatPageRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAdvicesFlatPageRequest) SetProduct(v string) *DescribeAdvicesFlatPageRequest {
	s.Product = &v
	return s
}

func (s *DescribeAdvicesFlatPageRequest) SetResourceId(v string) *DescribeAdvicesFlatPageRequest {
	s.ResourceId = &v
	return s
}

type DescribeAdvicesFlatPageResponseBody struct {
	Data      *DescribeAdvicesFlatPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAdvicesFlatPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdvicesFlatPageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAdvicesFlatPageResponseBody) SetData(v *DescribeAdvicesFlatPageResponseBodyData) *DescribeAdvicesFlatPageResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAdvicesFlatPageResponseBody) SetRequestId(v string) *DescribeAdvicesFlatPageResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAdvicesFlatPageResponseBodyData struct {
	PageNo   *int64                                           `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize *int64                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Result   []*DescribeAdvicesFlatPageResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Total    *int64                                           `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeAdvicesFlatPageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdvicesFlatPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAdvicesFlatPageResponseBodyData) SetPageNo(v int64) *DescribeAdvicesFlatPageResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *DescribeAdvicesFlatPageResponseBodyData) SetPageSize(v int64) *DescribeAdvicesFlatPageResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeAdvicesFlatPageResponseBodyData) SetResult(v []*DescribeAdvicesFlatPageResponseBodyDataResult) *DescribeAdvicesFlatPageResponseBodyData {
	s.Result = v
	return s
}

func (s *DescribeAdvicesFlatPageResponseBodyData) SetTotal(v int64) *DescribeAdvicesFlatPageResponseBodyData {
	s.Total = &v
	return s
}

type DescribeAdvicesFlatPageResponseBodyDataResult struct {
	AliyunId    *int64  `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
	CheckId     *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	CheckName   *string `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreated  *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	IsExpired   *bool   `json:"IsExpired,omitempty" xml:"IsExpired,omitempty"`
	Product     *string `json:"Product,omitempty" xml:"Product,omitempty"`
	Resource    *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	ResourceId  *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Severity    *int64  `json:"Severity,omitempty" xml:"Severity,omitempty"`
}

func (s DescribeAdvicesFlatPageResponseBodyDataResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdvicesFlatPageResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *DescribeAdvicesFlatPageResponseBodyDataResult) SetAliyunId(v int64) *DescribeAdvicesFlatPageResponseBodyDataResult {
	s.AliyunId = &v
	return s
}

func (s *DescribeAdvicesFlatPageResponseBodyDataResult) SetCheckId(v string) *DescribeAdvicesFlatPageResponseBodyDataResult {
	s.CheckId = &v
	return s
}

func (s *DescribeAdvicesFlatPageResponseBodyDataResult) SetCheckName(v string) *DescribeAdvicesFlatPageResponseBodyDataResult {
	s.CheckName = &v
	return s
}

func (s *DescribeAdvicesFlatPageResponseBodyDataResult) SetContent(v string) *DescribeAdvicesFlatPageResponseBodyDataResult {
	s.Content = &v
	return s
}

func (s *DescribeAdvicesFlatPageResponseBodyDataResult) SetDescription(v string) *DescribeAdvicesFlatPageResponseBodyDataResult {
	s.Description = &v
	return s
}

func (s *DescribeAdvicesFlatPageResponseBodyDataResult) SetGmtCreated(v string) *DescribeAdvicesFlatPageResponseBodyDataResult {
	s.GmtCreated = &v
	return s
}

func (s *DescribeAdvicesFlatPageResponseBodyDataResult) SetGmtModified(v string) *DescribeAdvicesFlatPageResponseBodyDataResult {
	s.GmtModified = &v
	return s
}

func (s *DescribeAdvicesFlatPageResponseBodyDataResult) SetId(v int64) *DescribeAdvicesFlatPageResponseBodyDataResult {
	s.Id = &v
	return s
}

func (s *DescribeAdvicesFlatPageResponseBodyDataResult) SetIsExpired(v bool) *DescribeAdvicesFlatPageResponseBodyDataResult {
	s.IsExpired = &v
	return s
}

func (s *DescribeAdvicesFlatPageResponseBodyDataResult) SetProduct(v string) *DescribeAdvicesFlatPageResponseBodyDataResult {
	s.Product = &v
	return s
}

func (s *DescribeAdvicesFlatPageResponseBodyDataResult) SetResource(v string) *DescribeAdvicesFlatPageResponseBodyDataResult {
	s.Resource = &v
	return s
}

func (s *DescribeAdvicesFlatPageResponseBodyDataResult) SetResourceId(v string) *DescribeAdvicesFlatPageResponseBodyDataResult {
	s.ResourceId = &v
	return s
}

func (s *DescribeAdvicesFlatPageResponseBodyDataResult) SetSeverity(v int64) *DescribeAdvicesFlatPageResponseBodyDataResult {
	s.Severity = &v
	return s
}

type DescribeAdvicesFlatPageResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAdvicesFlatPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAdvicesFlatPageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdvicesFlatPageResponse) GoString() string {
	return s.String()
}

func (s *DescribeAdvicesFlatPageResponse) SetHeaders(v map[string]*string) *DescribeAdvicesFlatPageResponse {
	s.Headers = v
	return s
}

func (s *DescribeAdvicesFlatPageResponse) SetStatusCode(v int32) *DescribeAdvicesFlatPageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAdvicesFlatPageResponse) SetBody(v *DescribeAdvicesFlatPageResponseBody) *DescribeAdvicesFlatPageResponse {
	s.Body = v
	return s
}

type DescribeAdvicesPageRequest struct {
	AdviceId   *int64  `json:"AdviceId,omitempty" xml:"AdviceId,omitempty"`
	CheckId    *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	Language   *string `json:"Language,omitempty" xml:"Language,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Product    *string `json:"Product,omitempty" xml:"Product,omitempty"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s DescribeAdvicesPageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdvicesPageRequest) GoString() string {
	return s.String()
}

func (s *DescribeAdvicesPageRequest) SetAdviceId(v int64) *DescribeAdvicesPageRequest {
	s.AdviceId = &v
	return s
}

func (s *DescribeAdvicesPageRequest) SetCheckId(v string) *DescribeAdvicesPageRequest {
	s.CheckId = &v
	return s
}

func (s *DescribeAdvicesPageRequest) SetLanguage(v string) *DescribeAdvicesPageRequest {
	s.Language = &v
	return s
}

func (s *DescribeAdvicesPageRequest) SetPageNumber(v int32) *DescribeAdvicesPageRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAdvicesPageRequest) SetPageSize(v int32) *DescribeAdvicesPageRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAdvicesPageRequest) SetProduct(v string) *DescribeAdvicesPageRequest {
	s.Product = &v
	return s
}

func (s *DescribeAdvicesPageRequest) SetResourceId(v string) *DescribeAdvicesPageRequest {
	s.ResourceId = &v
	return s
}

type DescribeAdvicesPageResponseBody struct {
	Data      *DescribeAdvicesPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAdvicesPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdvicesPageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAdvicesPageResponseBody) SetData(v *DescribeAdvicesPageResponseBodyData) *DescribeAdvicesPageResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAdvicesPageResponseBody) SetRequestId(v string) *DescribeAdvicesPageResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAdvicesPageResponseBodyData struct {
	PageNo   *int64                                       `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize *int64                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Result   []*DescribeAdvicesPageResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Total    *int64                                       `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeAdvicesPageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdvicesPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAdvicesPageResponseBodyData) SetPageNo(v int64) *DescribeAdvicesPageResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *DescribeAdvicesPageResponseBodyData) SetPageSize(v int64) *DescribeAdvicesPageResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeAdvicesPageResponseBodyData) SetResult(v []*DescribeAdvicesPageResponseBodyDataResult) *DescribeAdvicesPageResponseBodyData {
	s.Result = v
	return s
}

func (s *DescribeAdvicesPageResponseBodyData) SetTotal(v int64) *DescribeAdvicesPageResponseBodyData {
	s.Total = &v
	return s
}

type DescribeAdvicesPageResponseBodyDataResult struct {
	AliyunId    *int64  `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
	CheckId     *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	CheckName   *string `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreated  *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// ID
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	IsExpired  *bool   `json:"IsExpired,omitempty" xml:"IsExpired,omitempty"`
	Product    *string `json:"Product,omitempty" xml:"Product,omitempty"`
	Resource   *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Severity   *int64  `json:"Severity,omitempty" xml:"Severity,omitempty"`
}

func (s DescribeAdvicesPageResponseBodyDataResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdvicesPageResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *DescribeAdvicesPageResponseBodyDataResult) SetAliyunId(v int64) *DescribeAdvicesPageResponseBodyDataResult {
	s.AliyunId = &v
	return s
}

func (s *DescribeAdvicesPageResponseBodyDataResult) SetCheckId(v string) *DescribeAdvicesPageResponseBodyDataResult {
	s.CheckId = &v
	return s
}

func (s *DescribeAdvicesPageResponseBodyDataResult) SetCheckName(v string) *DescribeAdvicesPageResponseBodyDataResult {
	s.CheckName = &v
	return s
}

func (s *DescribeAdvicesPageResponseBodyDataResult) SetContent(v string) *DescribeAdvicesPageResponseBodyDataResult {
	s.Content = &v
	return s
}

func (s *DescribeAdvicesPageResponseBodyDataResult) SetDescription(v string) *DescribeAdvicesPageResponseBodyDataResult {
	s.Description = &v
	return s
}

func (s *DescribeAdvicesPageResponseBodyDataResult) SetGmtCreated(v string) *DescribeAdvicesPageResponseBodyDataResult {
	s.GmtCreated = &v
	return s
}

func (s *DescribeAdvicesPageResponseBodyDataResult) SetGmtModified(v string) *DescribeAdvicesPageResponseBodyDataResult {
	s.GmtModified = &v
	return s
}

func (s *DescribeAdvicesPageResponseBodyDataResult) SetId(v int64) *DescribeAdvicesPageResponseBodyDataResult {
	s.Id = &v
	return s
}

func (s *DescribeAdvicesPageResponseBodyDataResult) SetIsExpired(v bool) *DescribeAdvicesPageResponseBodyDataResult {
	s.IsExpired = &v
	return s
}

func (s *DescribeAdvicesPageResponseBodyDataResult) SetProduct(v string) *DescribeAdvicesPageResponseBodyDataResult {
	s.Product = &v
	return s
}

func (s *DescribeAdvicesPageResponseBodyDataResult) SetResource(v string) *DescribeAdvicesPageResponseBodyDataResult {
	s.Resource = &v
	return s
}

func (s *DescribeAdvicesPageResponseBodyDataResult) SetResourceId(v string) *DescribeAdvicesPageResponseBodyDataResult {
	s.ResourceId = &v
	return s
}

func (s *DescribeAdvicesPageResponseBodyDataResult) SetSeverity(v int64) *DescribeAdvicesPageResponseBodyDataResult {
	s.Severity = &v
	return s
}

type DescribeAdvicesPageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAdvicesPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAdvicesPageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdvicesPageResponse) GoString() string {
	return s.String()
}

func (s *DescribeAdvicesPageResponse) SetHeaders(v map[string]*string) *DescribeAdvicesPageResponse {
	s.Headers = v
	return s
}

func (s *DescribeAdvicesPageResponse) SetStatusCode(v int32) *DescribeAdvicesPageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAdvicesPageResponse) SetBody(v *DescribeAdvicesPageResponseBody) *DescribeAdvicesPageResponse {
	s.Body = v
	return s
}

type DescribeAdvisorChecksRequest struct {
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	Product  *string `json:"Product,omitempty" xml:"Product,omitempty"`
}

func (s DescribeAdvisorChecksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdvisorChecksRequest) GoString() string {
	return s.String()
}

func (s *DescribeAdvisorChecksRequest) SetLanguage(v string) *DescribeAdvisorChecksRequest {
	s.Language = &v
	return s
}

func (s *DescribeAdvisorChecksRequest) SetProduct(v string) *DescribeAdvisorChecksRequest {
	s.Product = &v
	return s
}

type DescribeAdvisorChecksResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeAdvisorChecksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAdvisorChecksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdvisorChecksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAdvisorChecksResponseBody) SetCode(v string) *DescribeAdvisorChecksResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAdvisorChecksResponseBody) SetData(v *DescribeAdvisorChecksResponseBodyData) *DescribeAdvisorChecksResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAdvisorChecksResponseBody) SetRequestId(v string) *DescribeAdvisorChecksResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAdvisorChecksResponseBodyData struct {
	AdvisorCheck []*DescribeAdvisorChecksResponseBodyDataAdvisorCheck `json:"AdvisorCheck,omitempty" xml:"AdvisorCheck,omitempty" type:"Repeated"`
}

func (s DescribeAdvisorChecksResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdvisorChecksResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAdvisorChecksResponseBodyData) SetAdvisorCheck(v []*DescribeAdvisorChecksResponseBodyDataAdvisorCheck) *DescribeAdvisorChecksResponseBodyData {
	s.AdvisorCheck = v
	return s
}

type DescribeAdvisorChecksResponseBodyDataAdvisorCheck struct {
	Category      *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Code          *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreated    *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtModified   *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OperateColumn *string `json:"OperateColumn,omitempty" xml:"OperateColumn,omitempty"`
	Product       *string `json:"Product,omitempty" xml:"Product,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Tips          *string `json:"Tips,omitempty" xml:"Tips,omitempty"`
	ViewColumn    *string `json:"ViewColumn,omitempty" xml:"ViewColumn,omitempty"`
}

func (s DescribeAdvisorChecksResponseBodyDataAdvisorCheck) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdvisorChecksResponseBodyDataAdvisorCheck) GoString() string {
	return s.String()
}

func (s *DescribeAdvisorChecksResponseBodyDataAdvisorCheck) SetCategory(v string) *DescribeAdvisorChecksResponseBodyDataAdvisorCheck {
	s.Category = &v
	return s
}

func (s *DescribeAdvisorChecksResponseBodyDataAdvisorCheck) SetCode(v string) *DescribeAdvisorChecksResponseBodyDataAdvisorCheck {
	s.Code = &v
	return s
}

func (s *DescribeAdvisorChecksResponseBodyDataAdvisorCheck) SetDescription(v string) *DescribeAdvisorChecksResponseBodyDataAdvisorCheck {
	s.Description = &v
	return s
}

func (s *DescribeAdvisorChecksResponseBodyDataAdvisorCheck) SetGmtCreated(v string) *DescribeAdvisorChecksResponseBodyDataAdvisorCheck {
	s.GmtCreated = &v
	return s
}

func (s *DescribeAdvisorChecksResponseBodyDataAdvisorCheck) SetGmtModified(v string) *DescribeAdvisorChecksResponseBodyDataAdvisorCheck {
	s.GmtModified = &v
	return s
}

func (s *DescribeAdvisorChecksResponseBodyDataAdvisorCheck) SetName(v string) *DescribeAdvisorChecksResponseBodyDataAdvisorCheck {
	s.Name = &v
	return s
}

func (s *DescribeAdvisorChecksResponseBodyDataAdvisorCheck) SetOperateColumn(v string) *DescribeAdvisorChecksResponseBodyDataAdvisorCheck {
	s.OperateColumn = &v
	return s
}

func (s *DescribeAdvisorChecksResponseBodyDataAdvisorCheck) SetProduct(v string) *DescribeAdvisorChecksResponseBodyDataAdvisorCheck {
	s.Product = &v
	return s
}

func (s *DescribeAdvisorChecksResponseBodyDataAdvisorCheck) SetStatus(v string) *DescribeAdvisorChecksResponseBodyDataAdvisorCheck {
	s.Status = &v
	return s
}

func (s *DescribeAdvisorChecksResponseBodyDataAdvisorCheck) SetTips(v string) *DescribeAdvisorChecksResponseBodyDataAdvisorCheck {
	s.Tips = &v
	return s
}

func (s *DescribeAdvisorChecksResponseBodyDataAdvisorCheck) SetViewColumn(v string) *DescribeAdvisorChecksResponseBodyDataAdvisorCheck {
	s.ViewColumn = &v
	return s
}

type DescribeAdvisorChecksResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAdvisorChecksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAdvisorChecksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdvisorChecksResponse) GoString() string {
	return s.String()
}

func (s *DescribeAdvisorChecksResponse) SetHeaders(v map[string]*string) *DescribeAdvisorChecksResponse {
	s.Headers = v
	return s
}

func (s *DescribeAdvisorChecksResponse) SetStatusCode(v int32) *DescribeAdvisorChecksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAdvisorChecksResponse) SetBody(v *DescribeAdvisorChecksResponseBody) *DescribeAdvisorChecksResponse {
	s.Body = v
	return s
}

type DescribeAdvisorResourcesRequest struct {
	Keyword    *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	Language   *string `json:"Language,omitempty" xml:"Language,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Product    *string `json:"Product,omitempty" xml:"Product,omitempty"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s DescribeAdvisorResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdvisorResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAdvisorResourcesRequest) SetKeyword(v string) *DescribeAdvisorResourcesRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeAdvisorResourcesRequest) SetLanguage(v string) *DescribeAdvisorResourcesRequest {
	s.Language = &v
	return s
}

func (s *DescribeAdvisorResourcesRequest) SetPageNumber(v int32) *DescribeAdvisorResourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAdvisorResourcesRequest) SetPageSize(v int32) *DescribeAdvisorResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAdvisorResourcesRequest) SetProduct(v string) *DescribeAdvisorResourcesRequest {
	s.Product = &v
	return s
}

func (s *DescribeAdvisorResourcesRequest) SetResourceId(v string) *DescribeAdvisorResourcesRequest {
	s.ResourceId = &v
	return s
}

type DescribeAdvisorResourcesResponseBody struct {
	Data      *DescribeAdvisorResourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAdvisorResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdvisorResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAdvisorResourcesResponseBody) SetData(v *DescribeAdvisorResourcesResponseBodyData) *DescribeAdvisorResourcesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAdvisorResourcesResponseBody) SetRequestId(v string) *DescribeAdvisorResourcesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAdvisorResourcesResponseBodyData struct {
	PageNo   *int32                                          `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Result   *DescribeAdvisorResourcesResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Total    *int32                                          `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeAdvisorResourcesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdvisorResourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAdvisorResourcesResponseBodyData) SetPageNo(v int32) *DescribeAdvisorResourcesResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *DescribeAdvisorResourcesResponseBodyData) SetPageSize(v int32) *DescribeAdvisorResourcesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeAdvisorResourcesResponseBodyData) SetResult(v *DescribeAdvisorResourcesResponseBodyDataResult) *DescribeAdvisorResourcesResponseBodyData {
	s.Result = v
	return s
}

func (s *DescribeAdvisorResourcesResponseBodyData) SetTotal(v int32) *DescribeAdvisorResourcesResponseBodyData {
	s.Total = &v
	return s
}

type DescribeAdvisorResourcesResponseBodyDataResult struct {
	Resource []*DescribeAdvisorResourcesResponseBodyDataResultResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
}

func (s DescribeAdvisorResourcesResponseBodyDataResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdvisorResourcesResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *DescribeAdvisorResourcesResponseBodyDataResult) SetResource(v []*DescribeAdvisorResourcesResponseBodyDataResultResource) *DescribeAdvisorResourcesResponseBodyDataResult {
	s.Resource = v
	return s
}

type DescribeAdvisorResourcesResponseBodyDataResultResource struct {
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Product      *string `json:"Product,omitempty" xml:"Product,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
}

func (s DescribeAdvisorResourcesResponseBodyDataResultResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdvisorResourcesResponseBodyDataResultResource) GoString() string {
	return s.String()
}

func (s *DescribeAdvisorResourcesResponseBodyDataResultResource) SetData(v string) *DescribeAdvisorResourcesResponseBodyDataResultResource {
	s.Data = &v
	return s
}

func (s *DescribeAdvisorResourcesResponseBodyDataResultResource) SetProduct(v string) *DescribeAdvisorResourcesResponseBodyDataResultResource {
	s.Product = &v
	return s
}

func (s *DescribeAdvisorResourcesResponseBodyDataResultResource) SetRegionId(v string) *DescribeAdvisorResourcesResponseBodyDataResultResource {
	s.RegionId = &v
	return s
}

func (s *DescribeAdvisorResourcesResponseBodyDataResultResource) SetResourceId(v string) *DescribeAdvisorResourcesResponseBodyDataResultResource {
	s.ResourceId = &v
	return s
}

func (s *DescribeAdvisorResourcesResponseBodyDataResultResource) SetResourceName(v string) *DescribeAdvisorResourcesResponseBodyDataResultResource {
	s.ResourceName = &v
	return s
}

type DescribeAdvisorResourcesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAdvisorResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAdvisorResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdvisorResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAdvisorResourcesResponse) SetHeaders(v map[string]*string) *DescribeAdvisorResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAdvisorResourcesResponse) SetStatusCode(v int32) *DescribeAdvisorResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAdvisorResourcesResponse) SetBody(v *DescribeAdvisorResourcesResponseBody) *DescribeAdvisorResourcesResponse {
	s.Body = v
	return s
}

type DescribeCostCheckAdvicesRequest struct {
	CheckId      *string   `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	Language     *string   `json:"Language,omitempty" xml:"Language,omitempty"`
	PageNumber   *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionIds    []*string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Repeated"`
	ResourceIds  []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	ResourceName *string   `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	Severity     *string   `json:"Severity,omitempty" xml:"Severity,omitempty"`
}

func (s DescribeCostCheckAdvicesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCostCheckAdvicesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCostCheckAdvicesRequest) SetCheckId(v string) *DescribeCostCheckAdvicesRequest {
	s.CheckId = &v
	return s
}

func (s *DescribeCostCheckAdvicesRequest) SetLanguage(v string) *DescribeCostCheckAdvicesRequest {
	s.Language = &v
	return s
}

func (s *DescribeCostCheckAdvicesRequest) SetPageNumber(v int32) *DescribeCostCheckAdvicesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCostCheckAdvicesRequest) SetPageSize(v int32) *DescribeCostCheckAdvicesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCostCheckAdvicesRequest) SetRegionIds(v []*string) *DescribeCostCheckAdvicesRequest {
	s.RegionIds = v
	return s
}

func (s *DescribeCostCheckAdvicesRequest) SetResourceIds(v []*string) *DescribeCostCheckAdvicesRequest {
	s.ResourceIds = v
	return s
}

func (s *DescribeCostCheckAdvicesRequest) SetResourceName(v string) *DescribeCostCheckAdvicesRequest {
	s.ResourceName = &v
	return s
}

func (s *DescribeCostCheckAdvicesRequest) SetSeverity(v string) *DescribeCostCheckAdvicesRequest {
	s.Severity = &v
	return s
}

type DescribeCostCheckAdvicesShrinkRequest struct {
	CheckId           *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	Language          *string `json:"Language,omitempty" xml:"Language,omitempty"`
	PageNumber        *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionIdsShrink   *string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty"`
	ResourceIdsShrink *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	ResourceName      *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	Severity          *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
}

func (s DescribeCostCheckAdvicesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCostCheckAdvicesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeCostCheckAdvicesShrinkRequest) SetCheckId(v string) *DescribeCostCheckAdvicesShrinkRequest {
	s.CheckId = &v
	return s
}

func (s *DescribeCostCheckAdvicesShrinkRequest) SetLanguage(v string) *DescribeCostCheckAdvicesShrinkRequest {
	s.Language = &v
	return s
}

func (s *DescribeCostCheckAdvicesShrinkRequest) SetPageNumber(v int32) *DescribeCostCheckAdvicesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCostCheckAdvicesShrinkRequest) SetPageSize(v int32) *DescribeCostCheckAdvicesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCostCheckAdvicesShrinkRequest) SetRegionIdsShrink(v string) *DescribeCostCheckAdvicesShrinkRequest {
	s.RegionIdsShrink = &v
	return s
}

func (s *DescribeCostCheckAdvicesShrinkRequest) SetResourceIdsShrink(v string) *DescribeCostCheckAdvicesShrinkRequest {
	s.ResourceIdsShrink = &v
	return s
}

func (s *DescribeCostCheckAdvicesShrinkRequest) SetResourceName(v string) *DescribeCostCheckAdvicesShrinkRequest {
	s.ResourceName = &v
	return s
}

func (s *DescribeCostCheckAdvicesShrinkRequest) SetSeverity(v string) *DescribeCostCheckAdvicesShrinkRequest {
	s.Severity = &v
	return s
}

type DescribeCostCheckAdvicesResponseBody struct {
	Code      *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeCostCheckAdvicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCostCheckAdvicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCostCheckAdvicesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCostCheckAdvicesResponseBody) SetCode(v string) *DescribeCostCheckAdvicesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBody) SetData(v *DescribeCostCheckAdvicesResponseBodyData) *DescribeCostCheckAdvicesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBody) SetMessage(v string) *DescribeCostCheckAdvicesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBody) SetRequestId(v string) *DescribeCostCheckAdvicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBody) SetSuccess(v string) *DescribeCostCheckAdvicesResponseBody {
	s.Success = &v
	return s
}

type DescribeCostCheckAdvicesResponseBodyData struct {
	AdviceList []*DescribeCostCheckAdvicesResponseBodyDataAdviceList `json:"AdviceList,omitempty" xml:"AdviceList,omitempty" type:"Repeated"`
	CheckId    *string                                               `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	PageNumber *int32                                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCostCheckAdvicesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCostCheckAdvicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCostCheckAdvicesResponseBodyData) SetAdviceList(v []*DescribeCostCheckAdvicesResponseBodyDataAdviceList) *DescribeCostCheckAdvicesResponseBodyData {
	s.AdviceList = v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyData) SetCheckId(v string) *DescribeCostCheckAdvicesResponseBodyData {
	s.CheckId = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyData) SetPageNumber(v int32) *DescribeCostCheckAdvicesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyData) SetPageSize(v int32) *DescribeCostCheckAdvicesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyData) SetTotalCount(v int32) *DescribeCostCheckAdvicesResponseBodyData {
	s.TotalCount = &v
	return s
}

type DescribeCostCheckAdvicesResponseBodyDataAdviceList struct {
	AliyunId     *int64                                                    `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
	Content      *string                                                   `json:"Content,omitempty" xml:"Content,omitempty"`
	EndTime      *int64                                                    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	GmtDeleted   *int64                                                    `json:"GmtDeleted,omitempty" xml:"GmtDeleted,omitempty"`
	Product      *string                                                   `json:"Product,omitempty" xml:"Product,omitempty"`
	RegionId     *string                                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId   *string                                                   `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName *string                                                   `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	Severity     *string                                                   `json:"Severity,omitempty" xml:"Severity,omitempty"`
	StartTime    *int64                                                    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Tags         []*DescribeCostCheckAdvicesResponseBodyDataAdviceListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	Url          *string                                                   `json:"Url,omitempty" xml:"Url,omitempty"`
	UserName     *string                                                   `json:"UserName,omitempty" xml:"UserName,omitempty"`
	ZoneId       *string                                                   `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeCostCheckAdvicesResponseBodyDataAdviceList) String() string {
	return tea.Prettify(s)
}

func (s DescribeCostCheckAdvicesResponseBodyDataAdviceList) GoString() string {
	return s.String()
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) SetAliyunId(v int64) *DescribeCostCheckAdvicesResponseBodyDataAdviceList {
	s.AliyunId = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) SetContent(v string) *DescribeCostCheckAdvicesResponseBodyDataAdviceList {
	s.Content = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) SetEndTime(v int64) *DescribeCostCheckAdvicesResponseBodyDataAdviceList {
	s.EndTime = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) SetGmtDeleted(v int64) *DescribeCostCheckAdvicesResponseBodyDataAdviceList {
	s.GmtDeleted = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) SetProduct(v string) *DescribeCostCheckAdvicesResponseBodyDataAdviceList {
	s.Product = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) SetRegionId(v string) *DescribeCostCheckAdvicesResponseBodyDataAdviceList {
	s.RegionId = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) SetResourceId(v string) *DescribeCostCheckAdvicesResponseBodyDataAdviceList {
	s.ResourceId = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) SetResourceName(v string) *DescribeCostCheckAdvicesResponseBodyDataAdviceList {
	s.ResourceName = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) SetSeverity(v string) *DescribeCostCheckAdvicesResponseBodyDataAdviceList {
	s.Severity = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) SetStartTime(v int64) *DescribeCostCheckAdvicesResponseBodyDataAdviceList {
	s.StartTime = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) SetTags(v []*DescribeCostCheckAdvicesResponseBodyDataAdviceListTags) *DescribeCostCheckAdvicesResponseBodyDataAdviceList {
	s.Tags = v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) SetUrl(v string) *DescribeCostCheckAdvicesResponseBodyDataAdviceList {
	s.Url = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) SetUserName(v string) *DescribeCostCheckAdvicesResponseBodyDataAdviceList {
	s.UserName = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) SetZoneId(v string) *DescribeCostCheckAdvicesResponseBodyDataAdviceList {
	s.ZoneId = &v
	return s
}

type DescribeCostCheckAdvicesResponseBodyDataAdviceListTags struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeCostCheckAdvicesResponseBodyDataAdviceListTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeCostCheckAdvicesResponseBodyDataAdviceListTags) GoString() string {
	return s.String()
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceListTags) SetTagKey(v string) *DescribeCostCheckAdvicesResponseBodyDataAdviceListTags {
	s.TagKey = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceListTags) SetTagValue(v string) *DescribeCostCheckAdvicesResponseBodyDataAdviceListTags {
	s.TagValue = &v
	return s
}

type DescribeCostCheckAdvicesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCostCheckAdvicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCostCheckAdvicesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCostCheckAdvicesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCostCheckAdvicesResponse) SetHeaders(v map[string]*string) *DescribeCostCheckAdvicesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCostCheckAdvicesResponse) SetStatusCode(v int32) *DescribeCostCheckAdvicesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponse) SetBody(v *DescribeCostCheckAdvicesResponseBody) *DescribeCostCheckAdvicesResponse {
	s.Body = v
	return s
}

type DescribeCostCheckResultsRequest struct {
	CheckIds     []*string `json:"CheckIds,omitempty" xml:"CheckIds,omitempty" type:"Repeated"`
	GroupBy      *string   `json:"GroupBy,omitempty" xml:"GroupBy,omitempty"`
	Product      *string   `json:"Product,omitempty" xml:"Product,omitempty"`
	RegionIds    []*string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Repeated"`
	ResourceIds  []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	ResourceName *string   `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	Severity     *int32    `json:"Severity,omitempty" xml:"Severity,omitempty"`
}

func (s DescribeCostCheckResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCostCheckResultsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCostCheckResultsRequest) SetCheckIds(v []*string) *DescribeCostCheckResultsRequest {
	s.CheckIds = v
	return s
}

func (s *DescribeCostCheckResultsRequest) SetGroupBy(v string) *DescribeCostCheckResultsRequest {
	s.GroupBy = &v
	return s
}

func (s *DescribeCostCheckResultsRequest) SetProduct(v string) *DescribeCostCheckResultsRequest {
	s.Product = &v
	return s
}

func (s *DescribeCostCheckResultsRequest) SetRegionIds(v []*string) *DescribeCostCheckResultsRequest {
	s.RegionIds = v
	return s
}

func (s *DescribeCostCheckResultsRequest) SetResourceIds(v []*string) *DescribeCostCheckResultsRequest {
	s.ResourceIds = v
	return s
}

func (s *DescribeCostCheckResultsRequest) SetResourceName(v string) *DescribeCostCheckResultsRequest {
	s.ResourceName = &v
	return s
}

func (s *DescribeCostCheckResultsRequest) SetSeverity(v int32) *DescribeCostCheckResultsRequest {
	s.Severity = &v
	return s
}

type DescribeCostCheckResultsShrinkRequest struct {
	CheckIdsShrink    *string `json:"CheckIds,omitempty" xml:"CheckIds,omitempty"`
	GroupBy           *string `json:"GroupBy,omitempty" xml:"GroupBy,omitempty"`
	Product           *string `json:"Product,omitempty" xml:"Product,omitempty"`
	RegionIdsShrink   *string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty"`
	ResourceIdsShrink *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	ResourceName      *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	Severity          *int32  `json:"Severity,omitempty" xml:"Severity,omitempty"`
}

func (s DescribeCostCheckResultsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCostCheckResultsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeCostCheckResultsShrinkRequest) SetCheckIdsShrink(v string) *DescribeCostCheckResultsShrinkRequest {
	s.CheckIdsShrink = &v
	return s
}

func (s *DescribeCostCheckResultsShrinkRequest) SetGroupBy(v string) *DescribeCostCheckResultsShrinkRequest {
	s.GroupBy = &v
	return s
}

func (s *DescribeCostCheckResultsShrinkRequest) SetProduct(v string) *DescribeCostCheckResultsShrinkRequest {
	s.Product = &v
	return s
}

func (s *DescribeCostCheckResultsShrinkRequest) SetRegionIdsShrink(v string) *DescribeCostCheckResultsShrinkRequest {
	s.RegionIdsShrink = &v
	return s
}

func (s *DescribeCostCheckResultsShrinkRequest) SetResourceIdsShrink(v string) *DescribeCostCheckResultsShrinkRequest {
	s.ResourceIdsShrink = &v
	return s
}

func (s *DescribeCostCheckResultsShrinkRequest) SetResourceName(v string) *DescribeCostCheckResultsShrinkRequest {
	s.ResourceName = &v
	return s
}

func (s *DescribeCostCheckResultsShrinkRequest) SetSeverity(v int32) *DescribeCostCheckResultsShrinkRequest {
	s.Severity = &v
	return s
}

type DescribeCostCheckResultsResponseBody struct {
	Code      *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeCostCheckResultsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCostCheckResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCostCheckResultsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCostCheckResultsResponseBody) SetCode(v string) *DescribeCostCheckResultsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBody) SetData(v *DescribeCostCheckResultsResponseBodyData) *DescribeCostCheckResultsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCostCheckResultsResponseBody) SetMessage(v string) *DescribeCostCheckResultsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBody) SetRequestId(v string) *DescribeCostCheckResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBody) SetSuccess(v string) *DescribeCostCheckResultsResponseBody {
	s.Success = &v
	return s
}

type DescribeCostCheckResultsResponseBodyData struct {
	AdviceResourceCount *int32                                               `json:"AdviceResourceCount,omitempty" xml:"AdviceResourceCount,omitempty"`
	GroupBy             *string                                              `json:"GroupBy,omitempty" xml:"GroupBy,omitempty"`
	NormalCount         *int32                                               `json:"NormalCount,omitempty" xml:"NormalCount,omitempty"`
	ResourceCount       *int32                                               `json:"ResourceCount,omitempty" xml:"ResourceCount,omitempty"`
	TotalCount          *int32                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	ViewGroup           []*DescribeCostCheckResultsResponseBodyDataViewGroup `json:"ViewGroup,omitempty" xml:"ViewGroup,omitempty" type:"Repeated"`
	WarningCount        *int32                                               `json:"WarningCount,omitempty" xml:"WarningCount,omitempty"`
}

func (s DescribeCostCheckResultsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCostCheckResultsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCostCheckResultsResponseBodyData) SetAdviceResourceCount(v int32) *DescribeCostCheckResultsResponseBodyData {
	s.AdviceResourceCount = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyData) SetGroupBy(v string) *DescribeCostCheckResultsResponseBodyData {
	s.GroupBy = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyData) SetNormalCount(v int32) *DescribeCostCheckResultsResponseBodyData {
	s.NormalCount = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyData) SetResourceCount(v int32) *DescribeCostCheckResultsResponseBodyData {
	s.ResourceCount = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyData) SetTotalCount(v int32) *DescribeCostCheckResultsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyData) SetViewGroup(v []*DescribeCostCheckResultsResponseBodyDataViewGroup) *DescribeCostCheckResultsResponseBodyData {
	s.ViewGroup = v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyData) SetWarningCount(v int32) *DescribeCostCheckResultsResponseBodyData {
	s.WarningCount = &v
	return s
}

type DescribeCostCheckResultsResponseBodyDataViewGroup struct {
	CheckItems              []*DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems `json:"CheckItems,omitempty" xml:"CheckItems,omitempty" type:"Repeated"`
	GroupCode               *string                                                        `json:"GroupCode,omitempty" xml:"GroupCode,omitempty"`
	GroupCount              *int32                                                         `json:"GroupCount,omitempty" xml:"GroupCount,omitempty"`
	GroupExpectedSavingCost *float32                                                       `json:"GroupExpectedSavingCost,omitempty" xml:"GroupExpectedSavingCost,omitempty"`
	GroupName               *string                                                        `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s DescribeCostCheckResultsResponseBodyDataViewGroup) String() string {
	return tea.Prettify(s)
}

func (s DescribeCostCheckResultsResponseBodyDataViewGroup) GoString() string {
	return s.String()
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroup) SetCheckItems(v []*DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) *DescribeCostCheckResultsResponseBodyDataViewGroup {
	s.CheckItems = v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroup) SetGroupCode(v string) *DescribeCostCheckResultsResponseBodyDataViewGroup {
	s.GroupCode = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroup) SetGroupCount(v int32) *DescribeCostCheckResultsResponseBodyDataViewGroup {
	s.GroupCount = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroup) SetGroupExpectedSavingCost(v float32) *DescribeCostCheckResultsResponseBodyDataViewGroup {
	s.GroupExpectedSavingCost = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroup) SetGroupName(v string) *DescribeCostCheckResultsResponseBodyDataViewGroup {
	s.GroupName = &v
	return s
}

type DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems struct {
	AdviceCount         *int32   `json:"AdviceCount,omitempty" xml:"AdviceCount,omitempty"`
	AdviceResourceCount *int32   `json:"AdviceResourceCount,omitempty" xml:"AdviceResourceCount,omitempty"`
	Category            *string  `json:"Category,omitempty" xml:"Category,omitempty"`
	CheckId             *string  `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	CheckName           *string  `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	CurrentCost         *float32 `json:"CurrentCost,omitempty" xml:"CurrentCost,omitempty"`
	Description         *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	ExpectedSavingCost  *float32 `json:"ExpectedSavingCost,omitempty" xml:"ExpectedSavingCost,omitempty"`
	OptimizedCost       *float32 `json:"OptimizedCost,omitempty" xml:"OptimizedCost,omitempty"`
	Product             *string  `json:"Product,omitempty" xml:"Product,omitempty"`
	Severity            *int32   `json:"Severity,omitempty" xml:"Severity,omitempty"`
	Summary             *string  `json:"Summary,omitempty" xml:"Summary,omitempty"`
	Tips                *string  `json:"Tips,omitempty" xml:"Tips,omitempty"`
}

func (s DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) GoString() string {
	return s.String()
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) SetAdviceCount(v int32) *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems {
	s.AdviceCount = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) SetAdviceResourceCount(v int32) *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems {
	s.AdviceResourceCount = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) SetCategory(v string) *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems {
	s.Category = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) SetCheckId(v string) *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems {
	s.CheckId = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) SetCheckName(v string) *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems {
	s.CheckName = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) SetCurrentCost(v float32) *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems {
	s.CurrentCost = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) SetDescription(v string) *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems {
	s.Description = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) SetExpectedSavingCost(v float32) *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems {
	s.ExpectedSavingCost = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) SetOptimizedCost(v float32) *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems {
	s.OptimizedCost = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) SetProduct(v string) *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems {
	s.Product = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) SetSeverity(v int32) *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems {
	s.Severity = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) SetSummary(v string) *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems {
	s.Summary = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) SetTips(v string) *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems {
	s.Tips = &v
	return s
}

type DescribeCostCheckResultsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCostCheckResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCostCheckResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCostCheckResultsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCostCheckResultsResponse) SetHeaders(v map[string]*string) *DescribeCostCheckResultsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCostCheckResultsResponse) SetStatusCode(v int32) *DescribeCostCheckResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCostCheckResultsResponse) SetBody(v *DescribeCostCheckResultsResponseBody) *DescribeCostCheckResultsResponse {
	s.Body = v
	return s
}

type GetHistoryAdvicesRequest struct {
	EndDate   *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Language  *string `json:"Language,omitempty" xml:"Language,omitempty"`
	PageNum   *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Product   *string `json:"Product,omitempty" xml:"Product,omitempty"`
	Reverse   *bool   `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	Severity  *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetHistoryAdvicesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHistoryAdvicesRequest) GoString() string {
	return s.String()
}

func (s *GetHistoryAdvicesRequest) SetEndDate(v string) *GetHistoryAdvicesRequest {
	s.EndDate = &v
	return s
}

func (s *GetHistoryAdvicesRequest) SetLanguage(v string) *GetHistoryAdvicesRequest {
	s.Language = &v
	return s
}

func (s *GetHistoryAdvicesRequest) SetPageNum(v int32) *GetHistoryAdvicesRequest {
	s.PageNum = &v
	return s
}

func (s *GetHistoryAdvicesRequest) SetPageSize(v int32) *GetHistoryAdvicesRequest {
	s.PageSize = &v
	return s
}

func (s *GetHistoryAdvicesRequest) SetProduct(v string) *GetHistoryAdvicesRequest {
	s.Product = &v
	return s
}

func (s *GetHistoryAdvicesRequest) SetReverse(v bool) *GetHistoryAdvicesRequest {
	s.Reverse = &v
	return s
}

func (s *GetHistoryAdvicesRequest) SetSeverity(v string) *GetHistoryAdvicesRequest {
	s.Severity = &v
	return s
}

func (s *GetHistoryAdvicesRequest) SetStartDate(v string) *GetHistoryAdvicesRequest {
	s.StartDate = &v
	return s
}

type GetHistoryAdvicesResponseBody struct {
	Data      *GetHistoryAdvicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetHistoryAdvicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHistoryAdvicesResponseBody) GoString() string {
	return s.String()
}

func (s *GetHistoryAdvicesResponseBody) SetData(v *GetHistoryAdvicesResponseBodyData) *GetHistoryAdvicesResponseBody {
	s.Data = v
	return s
}

func (s *GetHistoryAdvicesResponseBody) SetRequestId(v string) *GetHistoryAdvicesResponseBody {
	s.RequestId = &v
	return s
}

type GetHistoryAdvicesResponseBodyData struct {
	PageNo *int32                                     `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	Result []*GetHistoryAdvicesResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Total  *int32                                     `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetHistoryAdvicesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetHistoryAdvicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHistoryAdvicesResponseBodyData) SetPageNo(v int32) *GetHistoryAdvicesResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *GetHistoryAdvicesResponseBodyData) SetResult(v []*GetHistoryAdvicesResponseBodyDataResult) *GetHistoryAdvicesResponseBodyData {
	s.Result = v
	return s
}

func (s *GetHistoryAdvicesResponseBodyData) SetTotal(v int32) *GetHistoryAdvicesResponseBodyData {
	s.Total = &v
	return s
}

type GetHistoryAdvicesResponseBodyDataResult struct {
	CheckId     *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	CheckName   *string `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreated  *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	Product     *string `json:"Product,omitempty" xml:"Product,omitempty"`
	ResourceId  *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Severity    *int32  `json:"Severity,omitempty" xml:"Severity,omitempty"`
	Url         *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetHistoryAdvicesResponseBodyDataResult) String() string {
	return tea.Prettify(s)
}

func (s GetHistoryAdvicesResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *GetHistoryAdvicesResponseBodyDataResult) SetCheckId(v string) *GetHistoryAdvicesResponseBodyDataResult {
	s.CheckId = &v
	return s
}

func (s *GetHistoryAdvicesResponseBodyDataResult) SetCheckName(v string) *GetHistoryAdvicesResponseBodyDataResult {
	s.CheckName = &v
	return s
}

func (s *GetHistoryAdvicesResponseBodyDataResult) SetDescription(v string) *GetHistoryAdvicesResponseBodyDataResult {
	s.Description = &v
	return s
}

func (s *GetHistoryAdvicesResponseBodyDataResult) SetGmtCreated(v string) *GetHistoryAdvicesResponseBodyDataResult {
	s.GmtCreated = &v
	return s
}

func (s *GetHistoryAdvicesResponseBodyDataResult) SetProduct(v string) *GetHistoryAdvicesResponseBodyDataResult {
	s.Product = &v
	return s
}

func (s *GetHistoryAdvicesResponseBodyDataResult) SetResourceId(v string) *GetHistoryAdvicesResponseBodyDataResult {
	s.ResourceId = &v
	return s
}

func (s *GetHistoryAdvicesResponseBodyDataResult) SetSeverity(v int32) *GetHistoryAdvicesResponseBodyDataResult {
	s.Severity = &v
	return s
}

func (s *GetHistoryAdvicesResponseBodyDataResult) SetUrl(v string) *GetHistoryAdvicesResponseBodyDataResult {
	s.Url = &v
	return s
}

type GetHistoryAdvicesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHistoryAdvicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHistoryAdvicesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHistoryAdvicesResponse) GoString() string {
	return s.String()
}

func (s *GetHistoryAdvicesResponse) SetHeaders(v map[string]*string) *GetHistoryAdvicesResponse {
	s.Headers = v
	return s
}

func (s *GetHistoryAdvicesResponse) SetStatusCode(v int32) *GetHistoryAdvicesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHistoryAdvicesResponse) SetBody(v *GetHistoryAdvicesResponseBody) *GetHistoryAdvicesResponse {
	s.Body = v
	return s
}

type GetTaskStatusByIdRequest struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetTaskStatusByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTaskStatusByIdRequest) GoString() string {
	return s.String()
}

func (s *GetTaskStatusByIdRequest) SetTaskId(v string) *GetTaskStatusByIdRequest {
	s.TaskId = &v
	return s
}

type GetTaskStatusByIdResponseBody struct {
	Data      *GetTaskStatusByIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTaskStatusByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskStatusByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskStatusByIdResponseBody) SetData(v *GetTaskStatusByIdResponseBodyData) *GetTaskStatusByIdResponseBody {
	s.Data = v
	return s
}

func (s *GetTaskStatusByIdResponseBody) SetRequestId(v string) *GetTaskStatusByIdResponseBody {
	s.RequestId = &v
	return s
}

type GetTaskStatusByIdResponseBodyData struct {
	TaskId     *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskStatus *int32 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s GetTaskStatusByIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTaskStatusByIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTaskStatusByIdResponseBodyData) SetTaskId(v int64) *GetTaskStatusByIdResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetTaskStatusByIdResponseBodyData) SetTaskStatus(v int32) *GetTaskStatusByIdResponseBodyData {
	s.TaskStatus = &v
	return s
}

type GetTaskStatusByIdResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskStatusByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskStatusByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskStatusByIdResponse) GoString() string {
	return s.String()
}

func (s *GetTaskStatusByIdResponse) SetHeaders(v map[string]*string) *GetTaskStatusByIdResponse {
	s.Headers = v
	return s
}

func (s *GetTaskStatusByIdResponse) SetStatusCode(v int32) *GetTaskStatusByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskStatusByIdResponse) SetBody(v *GetTaskStatusByIdResponseBody) *GetTaskStatusByIdResponse {
	s.Body = v
	return s
}

type RefreshAdvisorCheckRequest struct {
	CheckId    *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	Language   *string `json:"Language,omitempty" xml:"Language,omitempty"`
	Product    *string `json:"Product,omitempty" xml:"Product,omitempty"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s RefreshAdvisorCheckRequest) String() string {
	return tea.Prettify(s)
}

func (s RefreshAdvisorCheckRequest) GoString() string {
	return s.String()
}

func (s *RefreshAdvisorCheckRequest) SetCheckId(v string) *RefreshAdvisorCheckRequest {
	s.CheckId = &v
	return s
}

func (s *RefreshAdvisorCheckRequest) SetLanguage(v string) *RefreshAdvisorCheckRequest {
	s.Language = &v
	return s
}

func (s *RefreshAdvisorCheckRequest) SetProduct(v string) *RefreshAdvisorCheckRequest {
	s.Product = &v
	return s
}

func (s *RefreshAdvisorCheckRequest) SetResourceId(v string) *RefreshAdvisorCheckRequest {
	s.ResourceId = &v
	return s
}

type RefreshAdvisorCheckResponseBody struct {
	Data      *RefreshAdvisorCheckResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefreshAdvisorCheckResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefreshAdvisorCheckResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshAdvisorCheckResponseBody) SetData(v *RefreshAdvisorCheckResponseBodyData) *RefreshAdvisorCheckResponseBody {
	s.Data = v
	return s
}

func (s *RefreshAdvisorCheckResponseBody) SetRequestId(v string) *RefreshAdvisorCheckResponseBody {
	s.RequestId = &v
	return s
}

type RefreshAdvisorCheckResponseBodyData struct {
	Success *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId  *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RefreshAdvisorCheckResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RefreshAdvisorCheckResponseBodyData) GoString() string {
	return s.String()
}

func (s *RefreshAdvisorCheckResponseBodyData) SetSuccess(v bool) *RefreshAdvisorCheckResponseBodyData {
	s.Success = &v
	return s
}

func (s *RefreshAdvisorCheckResponseBodyData) SetTaskId(v int64) *RefreshAdvisorCheckResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *RefreshAdvisorCheckResponseBodyData) SetTraceId(v string) *RefreshAdvisorCheckResponseBodyData {
	s.TraceId = &v
	return s
}

type RefreshAdvisorCheckResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshAdvisorCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshAdvisorCheckResponse) String() string {
	return tea.Prettify(s)
}

func (s RefreshAdvisorCheckResponse) GoString() string {
	return s.String()
}

func (s *RefreshAdvisorCheckResponse) SetHeaders(v map[string]*string) *RefreshAdvisorCheckResponse {
	s.Headers = v
	return s
}

func (s *RefreshAdvisorCheckResponse) SetStatusCode(v int32) *RefreshAdvisorCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshAdvisorCheckResponse) SetBody(v *RefreshAdvisorCheckResponseBody) *RefreshAdvisorCheckResponse {
	s.Body = v
	return s
}

type RefreshAdvisorResourceRequest struct {
	Product    *string `json:"Product,omitempty" xml:"Product,omitempty"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s RefreshAdvisorResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s RefreshAdvisorResourceRequest) GoString() string {
	return s.String()
}

func (s *RefreshAdvisorResourceRequest) SetProduct(v string) *RefreshAdvisorResourceRequest {
	s.Product = &v
	return s
}

func (s *RefreshAdvisorResourceRequest) SetResourceId(v string) *RefreshAdvisorResourceRequest {
	s.ResourceId = &v
	return s
}

type RefreshAdvisorResourceResponseBody struct {
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefreshAdvisorResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefreshAdvisorResourceResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshAdvisorResourceResponseBody) SetData(v int64) *RefreshAdvisorResourceResponseBody {
	s.Data = &v
	return s
}

func (s *RefreshAdvisorResourceResponseBody) SetRequestId(v string) *RefreshAdvisorResourceResponseBody {
	s.RequestId = &v
	return s
}

type RefreshAdvisorResourceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshAdvisorResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshAdvisorResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s RefreshAdvisorResourceResponse) GoString() string {
	return s.String()
}

func (s *RefreshAdvisorResourceResponse) SetHeaders(v map[string]*string) *RefreshAdvisorResourceResponse {
	s.Headers = v
	return s
}

func (s *RefreshAdvisorResourceResponse) SetStatusCode(v int32) *RefreshAdvisorResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshAdvisorResourceResponse) SetBody(v *RefreshAdvisorResourceResponseBody) *RefreshAdvisorResourceResponse {
	s.Body = v
	return s
}

type ReportBizAlertInfoRequest struct {
	AlertDescription *string  `json:"AlertDescription,omitempty" xml:"AlertDescription,omitempty"`
	AlertDetail      *string  `json:"AlertDetail,omitempty" xml:"AlertDetail,omitempty"`
	AlertGrade       *string  `json:"AlertGrade,omitempty" xml:"AlertGrade,omitempty"`
	AlertLabels      *string  `json:"AlertLabels,omitempty" xml:"AlertLabels,omitempty"`
	AlertScene       *string  `json:"AlertScene,omitempty" xml:"AlertScene,omitempty"`
	AlertToken       *string  `json:"AlertToken,omitempty" xml:"AlertToken,omitempty"`
	AlertUid         []*int64 `json:"AlertUid,omitempty" xml:"AlertUid,omitempty" type:"Repeated"`
	Language         *string  `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s ReportBizAlertInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportBizAlertInfoRequest) GoString() string {
	return s.String()
}

func (s *ReportBizAlertInfoRequest) SetAlertDescription(v string) *ReportBizAlertInfoRequest {
	s.AlertDescription = &v
	return s
}

func (s *ReportBizAlertInfoRequest) SetAlertDetail(v string) *ReportBizAlertInfoRequest {
	s.AlertDetail = &v
	return s
}

func (s *ReportBizAlertInfoRequest) SetAlertGrade(v string) *ReportBizAlertInfoRequest {
	s.AlertGrade = &v
	return s
}

func (s *ReportBizAlertInfoRequest) SetAlertLabels(v string) *ReportBizAlertInfoRequest {
	s.AlertLabels = &v
	return s
}

func (s *ReportBizAlertInfoRequest) SetAlertScene(v string) *ReportBizAlertInfoRequest {
	s.AlertScene = &v
	return s
}

func (s *ReportBizAlertInfoRequest) SetAlertToken(v string) *ReportBizAlertInfoRequest {
	s.AlertToken = &v
	return s
}

func (s *ReportBizAlertInfoRequest) SetAlertUid(v []*int64) *ReportBizAlertInfoRequest {
	s.AlertUid = v
	return s
}

func (s *ReportBizAlertInfoRequest) SetLanguage(v string) *ReportBizAlertInfoRequest {
	s.Language = &v
	return s
}

type ReportBizAlertInfoShrinkRequest struct {
	AlertDescription *string `json:"AlertDescription,omitempty" xml:"AlertDescription,omitempty"`
	AlertDetail      *string `json:"AlertDetail,omitempty" xml:"AlertDetail,omitempty"`
	AlertGrade       *string `json:"AlertGrade,omitempty" xml:"AlertGrade,omitempty"`
	AlertLabels      *string `json:"AlertLabels,omitempty" xml:"AlertLabels,omitempty"`
	AlertScene       *string `json:"AlertScene,omitempty" xml:"AlertScene,omitempty"`
	AlertToken       *string `json:"AlertToken,omitempty" xml:"AlertToken,omitempty"`
	AlertUidShrink   *string `json:"AlertUid,omitempty" xml:"AlertUid,omitempty"`
	Language         *string `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s ReportBizAlertInfoShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportBizAlertInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *ReportBizAlertInfoShrinkRequest) SetAlertDescription(v string) *ReportBizAlertInfoShrinkRequest {
	s.AlertDescription = &v
	return s
}

func (s *ReportBizAlertInfoShrinkRequest) SetAlertDetail(v string) *ReportBizAlertInfoShrinkRequest {
	s.AlertDetail = &v
	return s
}

func (s *ReportBizAlertInfoShrinkRequest) SetAlertGrade(v string) *ReportBizAlertInfoShrinkRequest {
	s.AlertGrade = &v
	return s
}

func (s *ReportBizAlertInfoShrinkRequest) SetAlertLabels(v string) *ReportBizAlertInfoShrinkRequest {
	s.AlertLabels = &v
	return s
}

func (s *ReportBizAlertInfoShrinkRequest) SetAlertScene(v string) *ReportBizAlertInfoShrinkRequest {
	s.AlertScene = &v
	return s
}

func (s *ReportBizAlertInfoShrinkRequest) SetAlertToken(v string) *ReportBizAlertInfoShrinkRequest {
	s.AlertToken = &v
	return s
}

func (s *ReportBizAlertInfoShrinkRequest) SetAlertUidShrink(v string) *ReportBizAlertInfoShrinkRequest {
	s.AlertUidShrink = &v
	return s
}

func (s *ReportBizAlertInfoShrinkRequest) SetLanguage(v string) *ReportBizAlertInfoShrinkRequest {
	s.Language = &v
	return s
}

type ReportBizAlertInfoResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ReportBizAlertInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReportBizAlertInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReportBizAlertInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ReportBizAlertInfoResponseBody) SetCode(v string) *ReportBizAlertInfoResponseBody {
	s.Code = &v
	return s
}

func (s *ReportBizAlertInfoResponseBody) SetData(v *ReportBizAlertInfoResponseBodyData) *ReportBizAlertInfoResponseBody {
	s.Data = v
	return s
}

func (s *ReportBizAlertInfoResponseBody) SetMessage(v string) *ReportBizAlertInfoResponseBody {
	s.Message = &v
	return s
}

func (s *ReportBizAlertInfoResponseBody) SetRequestId(v string) *ReportBizAlertInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReportBizAlertInfoResponseBody) SetSuccess(v bool) *ReportBizAlertInfoResponseBody {
	s.Success = &v
	return s
}

type ReportBizAlertInfoResponseBodyData struct {
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ReportBizAlertInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ReportBizAlertInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReportBizAlertInfoResponseBodyData) SetResult(v string) *ReportBizAlertInfoResponseBodyData {
	s.Result = &v
	return s
}

type ReportBizAlertInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReportBizAlertInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReportBizAlertInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s ReportBizAlertInfoResponse) GoString() string {
	return s.String()
}

func (s *ReportBizAlertInfoResponse) SetHeaders(v map[string]*string) *ReportBizAlertInfoResponse {
	s.Headers = v
	return s
}

func (s *ReportBizAlertInfoResponse) SetStatusCode(v int32) *ReportBizAlertInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportBizAlertInfoResponse) SetBody(v *ReportBizAlertInfoResponseBody) *ReportBizAlertInfoResponse {
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("advisor"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) DescribeAdvicesWithOptions(request *DescribeAdvicesRequest, runtime *util.RuntimeOptions) (_result *DescribeAdvicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdviceId)) {
		query["AdviceId"] = request.AdviceId
	}

	if !tea.BoolValue(util.IsUnset(request.CheckId)) {
		query["CheckId"] = request.CheckId
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeAdviceId)) {
		query["ExcludeAdviceId"] = request.ExcludeAdviceId
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		query["Product"] = request.Product
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAdvices"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAdvicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAdvices(request *DescribeAdvicesRequest) (_result *DescribeAdvicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAdvicesResponse{}
	_body, _err := client.DescribeAdvicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAdvicesFlatPageWithOptions(request *DescribeAdvicesFlatPageRequest, runtime *util.RuntimeOptions) (_result *DescribeAdvicesFlatPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdviceId)) {
		query["AdviceId"] = request.AdviceId
	}

	if !tea.BoolValue(util.IsUnset(request.CheckId)) {
		query["CheckId"] = request.CheckId
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		query["Product"] = request.Product
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAdvicesFlatPage"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAdvicesFlatPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAdvicesFlatPage(request *DescribeAdvicesFlatPageRequest) (_result *DescribeAdvicesFlatPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAdvicesFlatPageResponse{}
	_body, _err := client.DescribeAdvicesFlatPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAdvicesPageWithOptions(request *DescribeAdvicesPageRequest, runtime *util.RuntimeOptions) (_result *DescribeAdvicesPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdviceId)) {
		query["AdviceId"] = request.AdviceId
	}

	if !tea.BoolValue(util.IsUnset(request.CheckId)) {
		query["CheckId"] = request.CheckId
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		query["Product"] = request.Product
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAdvicesPage"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAdvicesPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAdvicesPage(request *DescribeAdvicesPageRequest) (_result *DescribeAdvicesPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAdvicesPageResponse{}
	_body, _err := client.DescribeAdvicesPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAdvisorChecksWithOptions(request *DescribeAdvisorChecksRequest, runtime *util.RuntimeOptions) (_result *DescribeAdvisorChecksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		query["Product"] = request.Product
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAdvisorChecks"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAdvisorChecksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAdvisorChecks(request *DescribeAdvisorChecksRequest) (_result *DescribeAdvisorChecksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAdvisorChecksResponse{}
	_body, _err := client.DescribeAdvisorChecksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAdvisorResourcesWithOptions(request *DescribeAdvisorResourcesRequest, runtime *util.RuntimeOptions) (_result *DescribeAdvisorResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		query["Product"] = request.Product
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAdvisorResources"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAdvisorResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAdvisorResources(request *DescribeAdvisorResourcesRequest) (_result *DescribeAdvisorResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAdvisorResourcesResponse{}
	_body, _err := client.DescribeAdvisorResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCostCheckAdvicesWithOptions(tmpReq *DescribeCostCheckAdvicesRequest, runtime *util.RuntimeOptions) (_result *DescribeCostCheckAdvicesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DescribeCostCheckAdvicesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RegionIds)) {
		request.RegionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RegionIds, tea.String("RegionIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceIds)) {
		request.ResourceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceIds, tea.String("ResourceIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CheckId)) {
		query["CheckId"] = request.CheckId
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionIdsShrink)) {
		query["RegionIds"] = request.RegionIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceIdsShrink)) {
		query["ResourceIds"] = request.ResourceIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceName)) {
		query["ResourceName"] = request.ResourceName
	}

	if !tea.BoolValue(util.IsUnset(request.Severity)) {
		query["Severity"] = request.Severity
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCostCheckAdvices"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCostCheckAdvicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCostCheckAdvices(request *DescribeCostCheckAdvicesRequest) (_result *DescribeCostCheckAdvicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCostCheckAdvicesResponse{}
	_body, _err := client.DescribeCostCheckAdvicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCostCheckResultsWithOptions(tmpReq *DescribeCostCheckResultsRequest, runtime *util.RuntimeOptions) (_result *DescribeCostCheckResultsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DescribeCostCheckResultsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CheckIds)) {
		request.CheckIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CheckIds, tea.String("CheckIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.RegionIds)) {
		request.RegionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RegionIds, tea.String("RegionIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceIds)) {
		request.ResourceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceIds, tea.String("ResourceIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CheckIdsShrink)) {
		query["CheckIds"] = request.CheckIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.GroupBy)) {
		query["GroupBy"] = request.GroupBy
	}

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		query["Product"] = request.Product
	}

	if !tea.BoolValue(util.IsUnset(request.RegionIdsShrink)) {
		query["RegionIds"] = request.RegionIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceIdsShrink)) {
		query["ResourceIds"] = request.ResourceIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceName)) {
		query["ResourceName"] = request.ResourceName
	}

	if !tea.BoolValue(util.IsUnset(request.Severity)) {
		query["Severity"] = request.Severity
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCostCheckResults"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCostCheckResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCostCheckResults(request *DescribeCostCheckResultsRequest) (_result *DescribeCostCheckResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCostCheckResultsResponse{}
	_body, _err := client.DescribeCostCheckResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHistoryAdvicesWithOptions(request *GetHistoryAdvicesRequest, runtime *util.RuntimeOptions) (_result *GetHistoryAdvicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		query["Product"] = request.Product
	}

	if !tea.BoolValue(util.IsUnset(request.Reverse)) {
		query["Reverse"] = request.Reverse
	}

	if !tea.BoolValue(util.IsUnset(request.Severity)) {
		query["Severity"] = request.Severity
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHistoryAdvices"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHistoryAdvicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHistoryAdvices(request *GetHistoryAdvicesRequest) (_result *GetHistoryAdvicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHistoryAdvicesResponse{}
	_body, _err := client.GetHistoryAdvicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTaskStatusByIdWithOptions(request *GetTaskStatusByIdRequest, runtime *util.RuntimeOptions) (_result *GetTaskStatusByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTaskStatusById"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTaskStatusByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTaskStatusById(request *GetTaskStatusByIdRequest) (_result *GetTaskStatusByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTaskStatusByIdResponse{}
	_body, _err := client.GetTaskStatusByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefreshAdvisorCheckWithOptions(request *RefreshAdvisorCheckRequest, runtime *util.RuntimeOptions) (_result *RefreshAdvisorCheckResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CheckId)) {
		query["CheckId"] = request.CheckId
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		query["Product"] = request.Product
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RefreshAdvisorCheck"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RefreshAdvisorCheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefreshAdvisorCheck(request *RefreshAdvisorCheckRequest) (_result *RefreshAdvisorCheckResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefreshAdvisorCheckResponse{}
	_body, _err := client.RefreshAdvisorCheckWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefreshAdvisorResourceWithOptions(request *RefreshAdvisorResourceRequest, runtime *util.RuntimeOptions) (_result *RefreshAdvisorResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Product)) {
		query["Product"] = request.Product
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RefreshAdvisorResource"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RefreshAdvisorResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefreshAdvisorResource(request *RefreshAdvisorResourceRequest) (_result *RefreshAdvisorResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefreshAdvisorResourceResponse{}
	_body, _err := client.RefreshAdvisorResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReportBizAlertInfoWithOptions(tmpReq *ReportBizAlertInfoRequest, runtime *util.RuntimeOptions) (_result *ReportBizAlertInfoResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ReportBizAlertInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AlertUid)) {
		request.AlertUidShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AlertUid, tea.String("AlertUid"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertDescription)) {
		query["AlertDescription"] = request.AlertDescription
	}

	if !tea.BoolValue(util.IsUnset(request.AlertDetail)) {
		query["AlertDetail"] = request.AlertDetail
	}

	if !tea.BoolValue(util.IsUnset(request.AlertGrade)) {
		query["AlertGrade"] = request.AlertGrade
	}

	if !tea.BoolValue(util.IsUnset(request.AlertLabels)) {
		query["AlertLabels"] = request.AlertLabels
	}

	if !tea.BoolValue(util.IsUnset(request.AlertScene)) {
		query["AlertScene"] = request.AlertScene
	}

	if !tea.BoolValue(util.IsUnset(request.AlertToken)) {
		query["AlertToken"] = request.AlertToken
	}

	if !tea.BoolValue(util.IsUnset(request.AlertUidShrink)) {
		query["AlertUid"] = request.AlertUidShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReportBizAlertInfo"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReportBizAlertInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReportBizAlertInfo(request *ReportBizAlertInfoRequest) (_result *ReportBizAlertInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReportBizAlertInfoResponse{}
	_body, _err := client.ReportBizAlertInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
