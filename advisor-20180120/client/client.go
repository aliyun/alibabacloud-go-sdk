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
	ClientUid       *int64  `json:"ClientUid,omitempty" xml:"ClientUid,omitempty"`
	ExcludeAdviceId *int64  `json:"ExcludeAdviceId,omitempty" xml:"ExcludeAdviceId,omitempty"`
	FilterType      *string `json:"FilterType,omitempty" xml:"FilterType,omitempty"`
	FilterValue     *string `json:"FilterValue,omitempty" xml:"FilterValue,omitempty"`
	Language        *string `json:"Language,omitempty" xml:"Language,omitempty"`
	Product         *string `json:"Product,omitempty" xml:"Product,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceId      *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Token           *string `json:"Token,omitempty" xml:"Token,omitempty"`
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

func (s *DescribeAdvicesRequest) SetClientUid(v int64) *DescribeAdvicesRequest {
	s.ClientUid = &v
	return s
}

func (s *DescribeAdvicesRequest) SetExcludeAdviceId(v int64) *DescribeAdvicesRequest {
	s.ExcludeAdviceId = &v
	return s
}

func (s *DescribeAdvicesRequest) SetFilterType(v string) *DescribeAdvicesRequest {
	s.FilterType = &v
	return s
}

func (s *DescribeAdvicesRequest) SetFilterValue(v string) *DescribeAdvicesRequest {
	s.FilterValue = &v
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

func (s *DescribeAdvicesRequest) SetRegion(v string) *DescribeAdvicesRequest {
	s.Region = &v
	return s
}

func (s *DescribeAdvicesRequest) SetResourceId(v string) *DescribeAdvicesRequest {
	s.ResourceId = &v
	return s
}

func (s *DescribeAdvicesRequest) SetToken(v string) *DescribeAdvicesRequest {
	s.Token = &v
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
	Action      *string `json:"Action,omitempty" xml:"Action,omitempty"`
	AliyunId    *int64  `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
	CheckId     *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	CheckName   *string `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Details     *string `json:"Details,omitempty" xml:"Details,omitempty"`
	GmtCreated  *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// ID
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	IsExpired  *bool   `json:"IsExpired,omitempty" xml:"IsExpired,omitempty"`
	Links      *string `json:"Links,omitempty" xml:"Links,omitempty"`
	Product    *string `json:"Product,omitempty" xml:"Product,omitempty"`
	Reason     *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
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

func (s *DescribeAdvicesResponseBodyDataAdvice) SetAction(v string) *DescribeAdvicesResponseBodyDataAdvice {
	s.Action = &v
	return s
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

func (s *DescribeAdvicesResponseBodyDataAdvice) SetDetails(v string) *DescribeAdvicesResponseBodyDataAdvice {
	s.Details = &v
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

func (s *DescribeAdvicesResponseBodyDataAdvice) SetLinks(v string) *DescribeAdvicesResponseBodyDataAdvice {
	s.Links = &v
	return s
}

func (s *DescribeAdvicesResponseBodyDataAdvice) SetProduct(v string) *DescribeAdvicesResponseBodyDataAdvice {
	s.Product = &v
	return s
}

func (s *DescribeAdvicesResponseBodyDataAdvice) SetReason(v string) *DescribeAdvicesResponseBodyDataAdvice {
	s.Reason = &v
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAdvicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DescribeAdvicesPageRequest struct {
	AdviceId        *int64  `json:"AdviceId,omitempty" xml:"AdviceId,omitempty"`
	AssociateUid    *int64  `json:"AssociateUid,omitempty" xml:"AssociateUid,omitempty"`
	CheckId         *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	ExcludeAdviceId *string `json:"ExcludeAdviceId,omitempty" xml:"ExcludeAdviceId,omitempty"`
	Language        *string `json:"Language,omitempty" xml:"Language,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Product         *string `json:"Product,omitempty" xml:"Product,omitempty"`
	ResourceId      *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
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

func (s *DescribeAdvicesPageRequest) SetAssociateUid(v int64) *DescribeAdvicesPageRequest {
	s.AssociateUid = &v
	return s
}

func (s *DescribeAdvicesPageRequest) SetCheckId(v string) *DescribeAdvicesPageRequest {
	s.CheckId = &v
	return s
}

func (s *DescribeAdvicesPageRequest) SetExcludeAdviceId(v string) *DescribeAdvicesPageRequest {
	s.ExcludeAdviceId = &v
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
	Action      *string `json:"Action,omitempty" xml:"Action,omitempty"`
	AliyunId    *int64  `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
	CheckId     *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	CheckName   *string `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Details     *string `json:"Details,omitempty" xml:"Details,omitempty"`
	GmtCreated  *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// ID
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	IsExpired  *bool   `json:"IsExpired,omitempty" xml:"IsExpired,omitempty"`
	Links      *string `json:"Links,omitempty" xml:"Links,omitempty"`
	Product    *string `json:"Product,omitempty" xml:"Product,omitempty"`
	Reason     *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
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

func (s *DescribeAdvicesPageResponseBodyDataResult) SetAction(v string) *DescribeAdvicesPageResponseBodyDataResult {
	s.Action = &v
	return s
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

func (s *DescribeAdvicesPageResponseBodyDataResult) SetDetails(v string) *DescribeAdvicesPageResponseBodyDataResult {
	s.Details = &v
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

func (s *DescribeAdvicesPageResponseBodyDataResult) SetLinks(v string) *DescribeAdvicesPageResponseBodyDataResult {
	s.Links = &v
	return s
}

func (s *DescribeAdvicesPageResponseBodyDataResult) SetProduct(v string) *DescribeAdvicesPageResponseBodyDataResult {
	s.Product = &v
	return s
}

func (s *DescribeAdvicesPageResponseBodyDataResult) SetReason(v string) *DescribeAdvicesPageResponseBodyDataResult {
	s.Reason = &v
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAdvicesPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	CheckId    *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	Language   *string `json:"Language,omitempty" xml:"Language,omitempty"`
	Product    *string `json:"Product,omitempty" xml:"Product,omitempty"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s DescribeAdvisorChecksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdvisorChecksRequest) GoString() string {
	return s.String()
}

func (s *DescribeAdvisorChecksRequest) SetCheckId(v string) *DescribeAdvisorChecksRequest {
	s.CheckId = &v
	return s
}

func (s *DescribeAdvisorChecksRequest) SetLanguage(v string) *DescribeAdvisorChecksRequest {
	s.Language = &v
	return s
}

func (s *DescribeAdvisorChecksRequest) SetProduct(v string) *DescribeAdvisorChecksRequest {
	s.Product = &v
	return s
}

func (s *DescribeAdvisorChecksRequest) SetResourceId(v string) *DescribeAdvisorChecksRequest {
	s.ResourceId = &v
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
	Category    *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Code        *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreated  *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// ID
	Id            *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OperateColumn *string `json:"OperateColumn,omitempty" xml:"OperateColumn,omitempty"`
	Product       *string `json:"Product,omitempty" xml:"Product,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SuppResources *string `json:"SuppResources,omitempty" xml:"SuppResources,omitempty"`
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

func (s *DescribeAdvisorChecksResponseBodyDataAdvisorCheck) SetId(v int64) *DescribeAdvisorChecksResponseBodyDataAdvisorCheck {
	s.Id = &v
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

func (s *DescribeAdvisorChecksResponseBodyDataAdvisorCheck) SetSuppResources(v string) *DescribeAdvisorChecksResponseBodyDataAdvisorCheck {
	s.SuppResources = &v
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAdvisorChecksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type GetHistoryAdvicesRequest struct {
	AssociateUid *int64  `json:"AssociateUid,omitempty" xml:"AssociateUid,omitempty"`
	CheckId      *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	ClientUid    *int64  `json:"ClientUid,omitempty" xml:"ClientUid,omitempty"`
	EndDate      *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	FilterType   *string `json:"FilterType,omitempty" xml:"FilterType,omitempty"`
	FilterValue  *string `json:"FilterValue,omitempty" xml:"FilterValue,omitempty"`
	IsExpired    *bool   `json:"IsExpired,omitempty" xml:"IsExpired,omitempty"`
	Language     *string `json:"Language,omitempty" xml:"Language,omitempty"`
	PageNum      *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Product      *string `json:"Product,omitempty" xml:"Product,omitempty"`
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Reverse      *bool   `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	Severity     *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	StartDate    *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetHistoryAdvicesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHistoryAdvicesRequest) GoString() string {
	return s.String()
}

func (s *GetHistoryAdvicesRequest) SetAssociateUid(v int64) *GetHistoryAdvicesRequest {
	s.AssociateUid = &v
	return s
}

func (s *GetHistoryAdvicesRequest) SetCheckId(v string) *GetHistoryAdvicesRequest {
	s.CheckId = &v
	return s
}

func (s *GetHistoryAdvicesRequest) SetClientUid(v int64) *GetHistoryAdvicesRequest {
	s.ClientUid = &v
	return s
}

func (s *GetHistoryAdvicesRequest) SetEndDate(v string) *GetHistoryAdvicesRequest {
	s.EndDate = &v
	return s
}

func (s *GetHistoryAdvicesRequest) SetFilterType(v string) *GetHistoryAdvicesRequest {
	s.FilterType = &v
	return s
}

func (s *GetHistoryAdvicesRequest) SetFilterValue(v string) *GetHistoryAdvicesRequest {
	s.FilterValue = &v
	return s
}

func (s *GetHistoryAdvicesRequest) SetIsExpired(v bool) *GetHistoryAdvicesRequest {
	s.IsExpired = &v
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

func (s *GetHistoryAdvicesRequest) SetRegion(v string) *GetHistoryAdvicesRequest {
	s.Region = &v
	return s
}

func (s *GetHistoryAdvicesRequest) SetResourceId(v string) *GetHistoryAdvicesRequest {
	s.ResourceId = &v
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetHistoryAdvicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RefreshAdvisorCheckResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RefreshAdvisorResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

	if !tea.BoolValue(util.IsUnset(request.ClientUid)) {
		query["ClientUid"] = request.ClientUid
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeAdviceId)) {
		query["ExcludeAdviceId"] = request.ExcludeAdviceId
	}

	if !tea.BoolValue(util.IsUnset(request.FilterType)) {
		query["FilterType"] = request.FilterType
	}

	if !tea.BoolValue(util.IsUnset(request.FilterValue)) {
		query["FilterValue"] = request.FilterValue
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		query["Product"] = request.Product
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["Token"] = request.Token
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

func (client *Client) DescribeAdvicesPageWithOptions(request *DescribeAdvicesPageRequest, runtime *util.RuntimeOptions) (_result *DescribeAdvicesPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdviceId)) {
		query["AdviceId"] = request.AdviceId
	}

	if !tea.BoolValue(util.IsUnset(request.AssociateUid)) {
		query["AssociateUid"] = request.AssociateUid
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

func (client *Client) GetHistoryAdvicesWithOptions(request *GetHistoryAdvicesRequest, runtime *util.RuntimeOptions) (_result *GetHistoryAdvicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssociateUid)) {
		query["AssociateUid"] = request.AssociateUid
	}

	if !tea.BoolValue(util.IsUnset(request.CheckId)) {
		query["CheckId"] = request.CheckId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientUid)) {
		query["ClientUid"] = request.ClientUid
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.FilterType)) {
		query["FilterType"] = request.FilterType
	}

	if !tea.BoolValue(util.IsUnset(request.FilterValue)) {
		query["FilterValue"] = request.FilterValue
	}

	if !tea.BoolValue(util.IsUnset(request.IsExpired)) {
		query["IsExpired"] = request.IsExpired
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

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
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
