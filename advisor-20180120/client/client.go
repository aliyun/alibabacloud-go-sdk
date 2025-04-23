// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type RdAccountDTO struct {
	AccountType *string             `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	Checked     *bool               `json:"Checked,omitempty" xml:"Checked,omitempty"`
	DisplayName *string             `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Id          *int64              `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string             `json:"Name,omitempty" xml:"Name,omitempty"`
	Tags        []*RdAccountDTOTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s RdAccountDTO) String() string {
	return tea.Prettify(s)
}

func (s RdAccountDTO) GoString() string {
	return s.String()
}

func (s *RdAccountDTO) SetAccountType(v string) *RdAccountDTO {
	s.AccountType = &v
	return s
}

func (s *RdAccountDTO) SetChecked(v bool) *RdAccountDTO {
	s.Checked = &v
	return s
}

func (s *RdAccountDTO) SetDisplayName(v string) *RdAccountDTO {
	s.DisplayName = &v
	return s
}

func (s *RdAccountDTO) SetId(v int64) *RdAccountDTO {
	s.Id = &v
	return s
}

func (s *RdAccountDTO) SetName(v string) *RdAccountDTO {
	s.Name = &v
	return s
}

func (s *RdAccountDTO) SetTags(v []*RdAccountDTOTags) *RdAccountDTO {
	s.Tags = v
	return s
}

type RdAccountDTOTags struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s RdAccountDTOTags) String() string {
	return tea.Prettify(s)
}

func (s RdAccountDTOTags) GoString() string {
	return s.String()
}

func (s *RdAccountDTOTags) SetTagKey(v string) *RdAccountDTOTags {
	s.TagKey = &v
	return s
}

func (s *RdAccountDTOTags) SetTagValue(v string) *RdAccountDTOTags {
	s.TagValue = &v
	return s
}

type RdAccountFolderDTO struct {
	AccountCount              *int32                `json:"AccountCount,omitempty" xml:"AccountCount,omitempty"`
	AccountList               []*RdAccountDTO       `json:"AccountList,omitempty" xml:"AccountList,omitempty" type:"Repeated"`
	FolderId                  *string               `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	FolderList                []*RdAccountFolderDTO `json:"FolderList,omitempty" xml:"FolderList,omitempty" type:"Repeated"`
	FolderName                *string               `json:"FolderName,omitempty" xml:"FolderName,omitempty"`
	ResourceDirectoryId       *string               `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	ResourceDirectoryPath     *string               `json:"ResourceDirectoryPath,omitempty" xml:"ResourceDirectoryPath,omitempty"`
	ResourceDirectoryPathName *string               `json:"ResourceDirectoryPathName,omitempty" xml:"ResourceDirectoryPathName,omitempty"`
	SelectedCount             *int32                `json:"SelectedCount,omitempty" xml:"SelectedCount,omitempty"`
}

func (s RdAccountFolderDTO) String() string {
	return tea.Prettify(s)
}

func (s RdAccountFolderDTO) GoString() string {
	return s.String()
}

func (s *RdAccountFolderDTO) SetAccountCount(v int32) *RdAccountFolderDTO {
	s.AccountCount = &v
	return s
}

func (s *RdAccountFolderDTO) SetAccountList(v []*RdAccountDTO) *RdAccountFolderDTO {
	s.AccountList = v
	return s
}

func (s *RdAccountFolderDTO) SetFolderId(v string) *RdAccountFolderDTO {
	s.FolderId = &v
	return s
}

func (s *RdAccountFolderDTO) SetFolderList(v []*RdAccountFolderDTO) *RdAccountFolderDTO {
	s.FolderList = v
	return s
}

func (s *RdAccountFolderDTO) SetFolderName(v string) *RdAccountFolderDTO {
	s.FolderName = &v
	return s
}

func (s *RdAccountFolderDTO) SetResourceDirectoryId(v string) *RdAccountFolderDTO {
	s.ResourceDirectoryId = &v
	return s
}

func (s *RdAccountFolderDTO) SetResourceDirectoryPath(v string) *RdAccountFolderDTO {
	s.ResourceDirectoryPath = &v
	return s
}

func (s *RdAccountFolderDTO) SetResourceDirectoryPathName(v string) *RdAccountFolderDTO {
	s.ResourceDirectoryPathName = &v
	return s
}

func (s *RdAccountFolderDTO) SetSelectedCount(v int32) *RdAccountFolderDTO {
	s.SelectedCount = &v
	return s
}

type DescribeAdvicesRequest struct {
	// example:
	//
	// 12345678
	AdviceId *int64 `json:"AdviceId,omitempty" xml:"AdviceId,omitempty"`
	// example:
	//
	// EcsHighCpuUtilization
	CheckId     *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	CheckPlanId *int64  `json:"CheckPlanId,omitempty" xml:"CheckPlanId,omitempty"`
	// example:
	//
	// 12345678
	ExcludeAdviceId *int64 `json:"ExcludeAdviceId,omitempty" xml:"ExcludeAdviceId,omitempty"`
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// ecs
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// example:
	//
	// i-bp67acfmxazb4p****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
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

func (s *DescribeAdvicesRequest) SetCheckPlanId(v int64) *DescribeAdvicesRequest {
	s.CheckPlanId = &v
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
	Data *DescribeAdvicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 566331F9-5AB3-550F-B745-A730331F97A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 1234567891234567
	AliyunId *int64 `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
	// example:
	//
	// EcsHighCpuUtilization
	CheckId     *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	CheckName   *string `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	CheckPlanId *int64  `json:"CheckPlanId,omitempty" xml:"CheckPlanId,omitempty"`
	// example:
	//
	// [
	//
	// 	{
	//
	// 		"key":"EcsHighCpuUtilization_xxxx",
	//
	// 		"value":xxx
	//
	// 	},
	//
	// 	{
	//
	// 		"key":"EcsHighCpuUtilization_xxxx",
	//
	// 		"value":xxx
	//
	// 	},
	//
	// 	{
	//
	// 		"key":"EcsHighCpuUtilization_xxxx",
	//
	// 		"value":xxx
	//
	// 	},
	//
	// ]
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2023-07-01 00:00:00
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// example:
	//
	// 2023-07-01 00:00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// ID
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// false
	IsExpired *bool `json:"IsExpired,omitempty" xml:"IsExpired,omitempty"`
	// example:
	//
	// ecs
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// example:
	//
	// {
	//
	// 	"resourceId": xxxx,
	//
	// 	"resourceName": xxxxxx,
	//
	// 	"regionId": xxxx,
	//
	// 	...
	//
	// }
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// example:
	//
	// i-bp67acfmxazb4p****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// 1
	Severity *int32 `json:"Severity,omitempty" xml:"Severity,omitempty"`
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

func (s *DescribeAdvicesResponseBodyDataAdvice) SetCheckPlanId(v int64) *DescribeAdvicesResponseBodyDataAdvice {
	s.CheckPlanId = &v
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
	// example:
	//
	// 12345678
	AdviceId *int64 `json:"AdviceId,omitempty" xml:"AdviceId,omitempty"`
	// example:
	//
	// EcsHighCpuUtilization
	CheckId *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// ecs
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// example:
	//
	// i-2zecnwitr2s7aca6****
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
	Data *DescribeAdvicesFlatPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 566331F9-5AB3-550F-B745-A730331F97A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int64                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Result   []*DescribeAdvicesFlatPageResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
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
	// example:
	//
	// 192895059480****
	AliyunId *int64 `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
	// example:
	//
	// EcsHighCpuUtilization
	CheckId   *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	CheckName *string `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	// example:
	//
	// [{"key":"EcsHighCpuUtilization_xxxx", "value":"xxx"}, {"key":"EcsHighCpuUtilization_xxxx", "value":"xxx"}, {"key":"EcsHighCpuUtilization_xxxx", "value":"xxx"}, ]
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2023-07-01 00:00:00
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// example:
	//
	// 2023-07-01 00:00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 40200899
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// false
	IsExpired *bool `json:"IsExpired,omitempty" xml:"IsExpired,omitempty"`
	// example:
	//
	// ecs
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// example:
	//
	// {"resourceId": "i-2zecnwitr2s7aca6****","resourceName": "ecs-20230701","regionId": "cn-hangzhou",...}
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// example:
	//
	// i-bp67acfmxazb4p****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// 1
	Severity *int64 `json:"Severity,omitempty" xml:"Severity,omitempty"`
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
	// example:
	//
	// 12345678
	AdviceId *int64 `json:"AdviceId,omitempty" xml:"AdviceId,omitempty"`
	// example:
	//
	// EcsHighCpuUtilization
	CheckId     *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	CheckPlanId *int64  `json:"CheckPlanId,omitempty" xml:"CheckPlanId,omitempty"`
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// ecs
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// example:
	//
	// i-bp67acfmxazb4p****
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

func (s *DescribeAdvicesPageRequest) SetCheckPlanId(v int64) *DescribeAdvicesPageRequest {
	s.CheckPlanId = &v
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
	Data *DescribeAdvicesPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 566331F9-5AB3-550F-B745-A730331F97A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int64                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Result   []*DescribeAdvicesPageResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
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
	// example:
	//
	// 1234567891234567
	AliyunId *int64 `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
	// example:
	//
	// EcsHighCpuUtilization
	CheckId   *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	CheckName *string `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	// example:
	//
	// [
	//
	// 	{
	//
	// 		"key":"EcsHighCpuUtilization_xxxx",
	//
	// 		"value":xxx
	//
	// 	},
	//
	// 	{
	//
	// 		"key":"EcsHighCpuUtilization_xxxx",
	//
	// 		"value":xxx
	//
	// 	},
	//
	// 	{
	//
	// 		"key":"EcsHighCpuUtilization_xxxx",
	//
	// 		"value":xxx
	//
	// 	},
	//
	// ]
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2023-07-01 00:00:00
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// example:
	//
	// 2023-07-01 00:00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// ID
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// false
	IsExpired *bool `json:"IsExpired,omitempty" xml:"IsExpired,omitempty"`
	// example:
	//
	// ecs
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// example:
	//
	// {
	//
	// 	"resourceId": xxxx,
	//
	// 	"resourceName": xxxxxx,
	//
	// 	"regionId": xxxx,
	//
	// 	...
	//
	// }
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// example:
	//
	// i-bp67acfmxazb4p****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// 1
	Severity *int64 `json:"Severity,omitempty" xml:"Severity,omitempty"`
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
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// ecs
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
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
	// example:
	//
	// 200
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DescribeAdvisorChecksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 566331F9-5AB3-550F-B745-A730331F97A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 1
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// EcsHighCpuUtilization
	Code        *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2023-07-01 00:00:00
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// example:
	//
	// 2023-07-01 00:00:00
	GmtModified   *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OperateColumn *string `json:"OperateColumn,omitempty" xml:"OperateColumn,omitempty"`
	// example:
	//
	// ecs
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Tips   *string `json:"Tips,omitempty" xml:"Tips,omitempty"`
	// example:
	//
	// [
	//
	// 	{
	//
	// 		"key":"EcsHighCpuUtilization_xxxx",
	//
	// 		"type":"DEFAULT"
	//
	// 	},
	//
	// 	{
	//
	// 		"key":"EcsHighCpuUtilization_xxxx",
	//
	// 		"type":"DEFAULT"
	//
	// 	},
	//
	// 	{
	//
	// 		"key":"EcsHighCpuUtilization_xxxx",
	//
	// 		"type":"DEFAULT"
	//
	// 	},
	//
	// ]
	ViewColumn *string `json:"ViewColumn,omitempty" xml:"ViewColumn,omitempty"`
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

type DescribeAdvisorChecksFoPagesRequest struct {
	// example:
	//
	// 11*********35
	AssumeAliyunId *int64 `json:"AssumeAliyunId,omitempty" xml:"AssumeAliyunId,omitempty"`
	// example:
	//
	// 2
	BizCategory *string `json:"BizCategory,omitempty" xml:"BizCategory,omitempty"`
	// example:
	//
	// *
	Category   *string  `json:"Category,omitempty" xml:"Category,omitempty"`
	CheckTypes []*int64 `json:"CheckTypes,omitempty" xml:"CheckTypes,omitempty" type:"Repeated"`
	// example:
	//
	// ****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// ecs
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// example:
	//
	// *
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// ***
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s DescribeAdvisorChecksFoPagesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdvisorChecksFoPagesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAdvisorChecksFoPagesRequest) SetAssumeAliyunId(v int64) *DescribeAdvisorChecksFoPagesRequest {
	s.AssumeAliyunId = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesRequest) SetBizCategory(v string) *DescribeAdvisorChecksFoPagesRequest {
	s.BizCategory = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesRequest) SetCategory(v string) *DescribeAdvisorChecksFoPagesRequest {
	s.Category = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesRequest) SetCheckTypes(v []*int64) *DescribeAdvisorChecksFoPagesRequest {
	s.CheckTypes = v
	return s
}

func (s *DescribeAdvisorChecksFoPagesRequest) SetName(v string) *DescribeAdvisorChecksFoPagesRequest {
	s.Name = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesRequest) SetPageNumber(v int32) *DescribeAdvisorChecksFoPagesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesRequest) SetPageSize(v int32) *DescribeAdvisorChecksFoPagesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesRequest) SetProduct(v string) *DescribeAdvisorChecksFoPagesRequest {
	s.Product = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesRequest) SetSource(v string) *DescribeAdvisorChecksFoPagesRequest {
	s.Source = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesRequest) SetStatus(v string) *DescribeAdvisorChecksFoPagesRequest {
	s.Status = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesRequest) SetToken(v string) *DescribeAdvisorChecksFoPagesRequest {
	s.Token = &v
	return s
}

type DescribeAdvisorChecksFoPagesShrinkRequest struct {
	// example:
	//
	// 11*********35
	AssumeAliyunId *int64 `json:"AssumeAliyunId,omitempty" xml:"AssumeAliyunId,omitempty"`
	// example:
	//
	// 2
	BizCategory *string `json:"BizCategory,omitempty" xml:"BizCategory,omitempty"`
	// example:
	//
	// *
	Category         *string `json:"Category,omitempty" xml:"Category,omitempty"`
	CheckTypesShrink *string `json:"CheckTypes,omitempty" xml:"CheckTypes,omitempty"`
	// example:
	//
	// ****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// ecs
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// example:
	//
	// *
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// ***
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s DescribeAdvisorChecksFoPagesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdvisorChecksFoPagesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeAdvisorChecksFoPagesShrinkRequest) SetAssumeAliyunId(v int64) *DescribeAdvisorChecksFoPagesShrinkRequest {
	s.AssumeAliyunId = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesShrinkRequest) SetBizCategory(v string) *DescribeAdvisorChecksFoPagesShrinkRequest {
	s.BizCategory = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesShrinkRequest) SetCategory(v string) *DescribeAdvisorChecksFoPagesShrinkRequest {
	s.Category = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesShrinkRequest) SetCheckTypesShrink(v string) *DescribeAdvisorChecksFoPagesShrinkRequest {
	s.CheckTypesShrink = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesShrinkRequest) SetName(v string) *DescribeAdvisorChecksFoPagesShrinkRequest {
	s.Name = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesShrinkRequest) SetPageNumber(v int32) *DescribeAdvisorChecksFoPagesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesShrinkRequest) SetPageSize(v int32) *DescribeAdvisorChecksFoPagesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesShrinkRequest) SetProduct(v string) *DescribeAdvisorChecksFoPagesShrinkRequest {
	s.Product = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesShrinkRequest) SetSource(v string) *DescribeAdvisorChecksFoPagesShrinkRequest {
	s.Source = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesShrinkRequest) SetStatus(v string) *DescribeAdvisorChecksFoPagesShrinkRequest {
	s.Status = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesShrinkRequest) SetToken(v string) *DescribeAdvisorChecksFoPagesShrinkRequest {
	s.Token = &v
	return s
}

type DescribeAdvisorChecksFoPagesResponseBody struct {
	// example:
	//
	// 200
	Code *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*DescribeAdvisorChecksFoPagesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 566331F9-5AB3-550F-B745-A730331F97A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAdvisorChecksFoPagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdvisorChecksFoPagesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAdvisorChecksFoPagesResponseBody) SetCode(v string) *DescribeAdvisorChecksFoPagesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBody) SetData(v []*DescribeAdvisorChecksFoPagesResponseBodyData) *DescribeAdvisorChecksFoPagesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBody) SetMessage(v string) *DescribeAdvisorChecksFoPagesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBody) SetRequestId(v string) *DescribeAdvisorChecksFoPagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBody) SetSuccess(v bool) *DescribeAdvisorChecksFoPagesResponseBody {
	s.Success = &v
	return s
}

type DescribeAdvisorChecksFoPagesResponseBodyData struct {
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Result   []*DescribeAdvisorChecksFoPagesResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeAdvisorChecksFoPagesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdvisorChecksFoPagesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyData) SetPageNo(v int32) *DescribeAdvisorChecksFoPagesResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyData) SetPageSize(v int32) *DescribeAdvisorChecksFoPagesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyData) SetResult(v []*DescribeAdvisorChecksFoPagesResponseBodyDataResult) *DescribeAdvisorChecksFoPagesResponseBodyData {
	s.Result = v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyData) SetTotal(v int32) *DescribeAdvisorChecksFoPagesResponseBodyData {
	s.Total = &v
	return s
}

type DescribeAdvisorChecksFoPagesResponseBodyDataResult struct {
	// example:
	//
	// 21
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// EcsCostLowUtilizationCheck
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// false
	ConfigSupport *string `json:"ConfigSupport,omitempty" xml:"ConfigSupport,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// false
	InspectionScope *string `json:"InspectionScope,omitempty" xml:"InspectionScope,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// [{\\"type\\":\\"template\\",\\"value\\":\\"cloudmonitor.console.aliyun.com/index.htm?custom_trace=ecs_console#/hostDetail/chart/instanceId=${Resource.resourceId}&system=Linux&region=${Resource.regionId}&aliyunhost=true\\",\\"key\\":\\"Detail\\"},{\\"type\\":\\"template\\",\\"value\\":\\"/diagnosis?product=${Product}&resourceId=${Resource.resourceId}\\",\\"key\\":\\"Refresh\\"}]
	OperateColumn *string `json:"OperateColumn,omitempty" xml:"OperateColumn,omitempty"`
	// example:
	//
	// ECS
	Product   *string `json:"Product,omitempty" xml:"Product,omitempty"`
	RiskLevel *int64  `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// example:
	//
	// Advisor
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// enabled
	Status      *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	SubCategory []*int64 `json:"SubCategory,omitempty" xml:"SubCategory,omitempty" type:"Repeated"`
	Tips        *string  `json:"Tips,omitempty" xml:"Tips,omitempty"`
	// example:
	//
	// [{\\"type\\":\\"DEFAULT\\",\\"key\\":\\"EcsCostIdleCheck_resourceId\\"},{\\"type\\":\\"DEFAULT\\",\\"key\\":\\"EcsCostIdleCheck_resourceName\\"},{\\"type\\":\\"DEFAULT\\",\\"key\\":\\"EcsCostIdleCheck_regionId\\"},{\\"type\\":\\"DEFAULT\\",\\"key\\":\\"EcsCostIdleCheck_instanceChargeType\\"},{\\"type\\":\\"DEFAULT\\",\\"key\\":\\"EcsCostIdleCheck_instanceType\\"},{\\"type\\":\\"DEFAULT\\",\\"key\\":\\"EcsCostIdleCheck_severity\\"},{\\"type\\":\\"DEFAULT\\",\\"key\\":\\"EcsCostIdleCheck_costBefore\\"},{\\"type\\":\\"DEFAULT\\",\\"key\\":\\"EcsCostIdleCheck_costAfter\\"},{\\"type\\":\\"DEFAULT\\",\\"key\\":\\"EcsCostIdleCheck_costSavings\\"},{\\"type\\":\\"DEFAULT\\",\\"key\\":\\"First_time\\"},{\\"type\\":\\"DEFAULT\\",\\"key\\":\\"Duration_time\\"}]
	ViewColumn *string `json:"ViewColumn,omitempty" xml:"ViewColumn,omitempty"`
}

func (s DescribeAdvisorChecksFoPagesResponseBodyDataResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdvisorChecksFoPagesResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) SetCategory(v string) *DescribeAdvisorChecksFoPagesResponseBodyDataResult {
	s.Category = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) SetCode(v string) *DescribeAdvisorChecksFoPagesResponseBodyDataResult {
	s.Code = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) SetConfigSupport(v string) *DescribeAdvisorChecksFoPagesResponseBodyDataResult {
	s.ConfigSupport = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) SetDescription(v string) *DescribeAdvisorChecksFoPagesResponseBodyDataResult {
	s.Description = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) SetInspectionScope(v string) *DescribeAdvisorChecksFoPagesResponseBodyDataResult {
	s.InspectionScope = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) SetName(v string) *DescribeAdvisorChecksFoPagesResponseBodyDataResult {
	s.Name = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) SetOperateColumn(v string) *DescribeAdvisorChecksFoPagesResponseBodyDataResult {
	s.OperateColumn = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) SetProduct(v string) *DescribeAdvisorChecksFoPagesResponseBodyDataResult {
	s.Product = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) SetRiskLevel(v int64) *DescribeAdvisorChecksFoPagesResponseBodyDataResult {
	s.RiskLevel = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) SetSource(v string) *DescribeAdvisorChecksFoPagesResponseBodyDataResult {
	s.Source = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) SetStatus(v string) *DescribeAdvisorChecksFoPagesResponseBodyDataResult {
	s.Status = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) SetSubCategory(v []*int64) *DescribeAdvisorChecksFoPagesResponseBodyDataResult {
	s.SubCategory = v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) SetTips(v string) *DescribeAdvisorChecksFoPagesResponseBodyDataResult {
	s.Tips = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponseBodyDataResult) SetViewColumn(v string) *DescribeAdvisorChecksFoPagesResponseBodyDataResult {
	s.ViewColumn = &v
	return s
}

type DescribeAdvisorChecksFoPagesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAdvisorChecksFoPagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAdvisorChecksFoPagesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdvisorChecksFoPagesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAdvisorChecksFoPagesResponse) SetHeaders(v map[string]*string) *DescribeAdvisorChecksFoPagesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponse) SetStatusCode(v int32) *DescribeAdvisorChecksFoPagesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponse) SetBody(v *DescribeAdvisorChecksFoPagesResponseBody) *DescribeAdvisorChecksFoPagesResponse {
	s.Body = v
	return s
}

type DescribeAdvisorResourcesRequest struct {
	// example:
	//
	// abcd
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// ecs
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// example:
	//
	// i-bp67acfmxazb4p****
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
	Data *DescribeAdvisorResourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 566331F9-5AB3-550F-B745-A730331F97A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Result   *DescribeAdvisorResourcesResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
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
	// example:
	//
	// {
	//
	//     "resourceId": "xxxxx",
	//
	//     "deviceAvailable": true,
	//
	//     ...
	//
	// }
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// ecs
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// i-bp67acfmxazb4p****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// ecs-20230701
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
	AssumeAliyunIdList []*int64 `json:"AssumeAliyunIdList,omitempty" xml:"AssumeAliyunIdList,omitempty" type:"Repeated"`
	// example:
	//
	// EcsCostLowUtilizationCheck
	CheckId     *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	CheckPlanId *int64  `json:"CheckPlanId,omitempty" xml:"CheckPlanId,omitempty"`
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// 6
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize            *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionIds           []*string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Repeated"`
	ResourceGroupIdList []*string `json:"ResourceGroupIdList,omitempty" xml:"ResourceGroupIdList,omitempty" type:"Repeated"`
	ResourceId          *string   `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceIds         []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// example:
	//
	// test
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// example:
	//
	// 1
	Severity  *string                                   `json:"Severity,omitempty" xml:"Severity,omitempty"`
	TagKeys   []*string                                 `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
	TagList   []*DescribeCostCheckAdvicesRequestTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	TagValues []*string                                 `json:"TagValues,omitempty" xml:"TagValues,omitempty" type:"Repeated"`
}

func (s DescribeCostCheckAdvicesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCostCheckAdvicesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCostCheckAdvicesRequest) SetAssumeAliyunIdList(v []*int64) *DescribeCostCheckAdvicesRequest {
	s.AssumeAliyunIdList = v
	return s
}

func (s *DescribeCostCheckAdvicesRequest) SetCheckId(v string) *DescribeCostCheckAdvicesRequest {
	s.CheckId = &v
	return s
}

func (s *DescribeCostCheckAdvicesRequest) SetCheckPlanId(v int64) *DescribeCostCheckAdvicesRequest {
	s.CheckPlanId = &v
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

func (s *DescribeCostCheckAdvicesRequest) SetResourceGroupIdList(v []*string) *DescribeCostCheckAdvicesRequest {
	s.ResourceGroupIdList = v
	return s
}

func (s *DescribeCostCheckAdvicesRequest) SetResourceId(v string) *DescribeCostCheckAdvicesRequest {
	s.ResourceId = &v
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

func (s *DescribeCostCheckAdvicesRequest) SetTagKeys(v []*string) *DescribeCostCheckAdvicesRequest {
	s.TagKeys = v
	return s
}

func (s *DescribeCostCheckAdvicesRequest) SetTagList(v []*DescribeCostCheckAdvicesRequestTagList) *DescribeCostCheckAdvicesRequest {
	s.TagList = v
	return s
}

func (s *DescribeCostCheckAdvicesRequest) SetTagValues(v []*string) *DescribeCostCheckAdvicesRequest {
	s.TagValues = v
	return s
}

type DescribeCostCheckAdvicesRequestTagList struct {
	// example:
	//
	// ecs_***_shanghai
	TagKey   *string   `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue []*string `json:"TagValue,omitempty" xml:"TagValue,omitempty" type:"Repeated"`
}

func (s DescribeCostCheckAdvicesRequestTagList) String() string {
	return tea.Prettify(s)
}

func (s DescribeCostCheckAdvicesRequestTagList) GoString() string {
	return s.String()
}

func (s *DescribeCostCheckAdvicesRequestTagList) SetTagKey(v string) *DescribeCostCheckAdvicesRequestTagList {
	s.TagKey = &v
	return s
}

func (s *DescribeCostCheckAdvicesRequestTagList) SetTagValue(v []*string) *DescribeCostCheckAdvicesRequestTagList {
	s.TagValue = v
	return s
}

type DescribeCostCheckAdvicesShrinkRequest struct {
	AssumeAliyunIdListShrink *string `json:"AssumeAliyunIdList,omitempty" xml:"AssumeAliyunIdList,omitempty"`
	// example:
	//
	// EcsCostLowUtilizationCheck
	CheckId     *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	CheckPlanId *int64  `json:"CheckPlanId,omitempty" xml:"CheckPlanId,omitempty"`
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// 6
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize                  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionIdsShrink           *string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty"`
	ResourceGroupIdListShrink *string `json:"ResourceGroupIdList,omitempty" xml:"ResourceGroupIdList,omitempty"`
	ResourceId                *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceIdsShrink         *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// example:
	//
	// test
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// example:
	//
	// 1
	Severity        *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	TagKeysShrink   *string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty"`
	TagListShrink   *string `json:"TagList,omitempty" xml:"TagList,omitempty"`
	TagValuesShrink *string `json:"TagValues,omitempty" xml:"TagValues,omitempty"`
}

func (s DescribeCostCheckAdvicesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCostCheckAdvicesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeCostCheckAdvicesShrinkRequest) SetAssumeAliyunIdListShrink(v string) *DescribeCostCheckAdvicesShrinkRequest {
	s.AssumeAliyunIdListShrink = &v
	return s
}

func (s *DescribeCostCheckAdvicesShrinkRequest) SetCheckId(v string) *DescribeCostCheckAdvicesShrinkRequest {
	s.CheckId = &v
	return s
}

func (s *DescribeCostCheckAdvicesShrinkRequest) SetCheckPlanId(v int64) *DescribeCostCheckAdvicesShrinkRequest {
	s.CheckPlanId = &v
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

func (s *DescribeCostCheckAdvicesShrinkRequest) SetResourceGroupIdListShrink(v string) *DescribeCostCheckAdvicesShrinkRequest {
	s.ResourceGroupIdListShrink = &v
	return s
}

func (s *DescribeCostCheckAdvicesShrinkRequest) SetResourceId(v string) *DescribeCostCheckAdvicesShrinkRequest {
	s.ResourceId = &v
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

func (s *DescribeCostCheckAdvicesShrinkRequest) SetTagKeysShrink(v string) *DescribeCostCheckAdvicesShrinkRequest {
	s.TagKeysShrink = &v
	return s
}

func (s *DescribeCostCheckAdvicesShrinkRequest) SetTagListShrink(v string) *DescribeCostCheckAdvicesShrinkRequest {
	s.TagListShrink = &v
	return s
}

func (s *DescribeCostCheckAdvicesShrinkRequest) SetTagValuesShrink(v string) *DescribeCostCheckAdvicesShrinkRequest {
	s.TagValuesShrink = &v
	return s
}

type DescribeCostCheckAdvicesResponseBody struct {
	// example:
	//
	// 200
	Code *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DescribeCostCheckAdvicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 566331F9-****-550F-B745-A730331F97A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// example:
	//
	// EcsHighCpuUtilization
	CheckId *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// example:
	//
	// 4
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 0
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// example:
	//
	// 1
	AccountFolderId *string `json:"AccountFolderId,omitempty" xml:"AccountFolderId,omitempty"`
	// example:
	//
	// 1
	AccountFolderName *string `json:"AccountFolderName,omitempty" xml:"AccountFolderName,omitempty"`
	// example:
	//
	// 111******767
	AliyunId *int64 `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
	// example:
	//
	// {\\"Domains\\": [{\\"Status\\": \\"success\\", \\"\\": \\"cn\\", \\"DomainName\\": \\"www.****.com\\", Region\\"Desc\\": \\"ok\\"}]}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Email
	//
	// example:
	//
	// xxx
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// 2025-03-05T02:02:00Z
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 2025-03-05T02:02:00Z
	GmtDeleted *int64 `json:"GmtDeleted,omitempty" xml:"GmtDeleted,omitempty"`
	// example:
	//
	// ecs
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// i-2ze5*****ef7d2lk63in
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// 1200***bles_df
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// example:
	//
	// 1
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// example:
	//
	// 2025-02-04T16:00:00Z
	StartTime *int64                                                    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Tags      []*DescribeCostCheckAdvicesResponseBodyDataAdviceListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	Url       *string                                                   `json:"Url,omitempty" xml:"Url,omitempty"`
	// example:
	//
	// test
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// example:
	//
	// cn-beijing-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeCostCheckAdvicesResponseBodyDataAdviceList) String() string {
	return tea.Prettify(s)
}

func (s DescribeCostCheckAdvicesResponseBodyDataAdviceList) GoString() string {
	return s.String()
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) SetAccountFolderId(v string) *DescribeCostCheckAdvicesResponseBodyDataAdviceList {
	s.AccountFolderId = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) SetAccountFolderName(v string) *DescribeCostCheckAdvicesResponseBodyDataAdviceList {
	s.AccountFolderName = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) SetAliyunId(v int64) *DescribeCostCheckAdvicesResponseBodyDataAdviceList {
	s.AliyunId = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) SetContent(v string) *DescribeCostCheckAdvicesResponseBodyDataAdviceList {
	s.Content = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) SetEmail(v string) *DescribeCostCheckAdvicesResponseBodyDataAdviceList {
	s.Email = &v
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
	// example:
	//
	// autoTest-7
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// example:
	//
	// basic
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
	AssumeAliyunIdList []*int64  `json:"AssumeAliyunIdList,omitempty" xml:"AssumeAliyunIdList,omitempty" type:"Repeated"`
	CheckIds           []*string `json:"CheckIds,omitempty" xml:"CheckIds,omitempty" type:"Repeated"`
	CheckPlanId        *int64    `json:"CheckPlanId,omitempty" xml:"CheckPlanId,omitempty"`
	// example:
	//
	// Category
	GroupBy *string `json:"GroupBy,omitempty" xml:"GroupBy,omitempty"`
	// example:
	//
	// ecs
	Product             *string   `json:"Product,omitempty" xml:"Product,omitempty"`
	RegionIds           []*string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Repeated"`
	ResourceGroupIdList []*string `json:"ResourceGroupIdList,omitempty" xml:"ResourceGroupIdList,omitempty" type:"Repeated"`
	ResourceId          *string   `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceIds         []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// example:
	//
	// SYNC_********_TASK
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// example:
	//
	// 1
	Severity  *int32                                    `json:"Severity,omitempty" xml:"Severity,omitempty"`
	TagKeys   []*string                                 `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
	TagList   []*DescribeCostCheckResultsRequestTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	TagValues []*string                                 `json:"TagValues,omitempty" xml:"TagValues,omitempty" type:"Repeated"`
}

func (s DescribeCostCheckResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCostCheckResultsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCostCheckResultsRequest) SetAssumeAliyunIdList(v []*int64) *DescribeCostCheckResultsRequest {
	s.AssumeAliyunIdList = v
	return s
}

func (s *DescribeCostCheckResultsRequest) SetCheckIds(v []*string) *DescribeCostCheckResultsRequest {
	s.CheckIds = v
	return s
}

func (s *DescribeCostCheckResultsRequest) SetCheckPlanId(v int64) *DescribeCostCheckResultsRequest {
	s.CheckPlanId = &v
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

func (s *DescribeCostCheckResultsRequest) SetResourceGroupIdList(v []*string) *DescribeCostCheckResultsRequest {
	s.ResourceGroupIdList = v
	return s
}

func (s *DescribeCostCheckResultsRequest) SetResourceId(v string) *DescribeCostCheckResultsRequest {
	s.ResourceId = &v
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

func (s *DescribeCostCheckResultsRequest) SetTagKeys(v []*string) *DescribeCostCheckResultsRequest {
	s.TagKeys = v
	return s
}

func (s *DescribeCostCheckResultsRequest) SetTagList(v []*DescribeCostCheckResultsRequestTagList) *DescribeCostCheckResultsRequest {
	s.TagList = v
	return s
}

func (s *DescribeCostCheckResultsRequest) SetTagValues(v []*string) *DescribeCostCheckResultsRequest {
	s.TagValues = v
	return s
}

type DescribeCostCheckResultsRequestTagList struct {
	// example:
	//
	// ERP
	TagKey   *string   `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue []*string `json:"TagValue,omitempty" xml:"TagValue,omitempty" type:"Repeated"`
}

func (s DescribeCostCheckResultsRequestTagList) String() string {
	return tea.Prettify(s)
}

func (s DescribeCostCheckResultsRequestTagList) GoString() string {
	return s.String()
}

func (s *DescribeCostCheckResultsRequestTagList) SetTagKey(v string) *DescribeCostCheckResultsRequestTagList {
	s.TagKey = &v
	return s
}

func (s *DescribeCostCheckResultsRequestTagList) SetTagValue(v []*string) *DescribeCostCheckResultsRequestTagList {
	s.TagValue = v
	return s
}

type DescribeCostCheckResultsShrinkRequest struct {
	AssumeAliyunIdListShrink *string `json:"AssumeAliyunIdList,omitempty" xml:"AssumeAliyunIdList,omitempty"`
	CheckIdsShrink           *string `json:"CheckIds,omitempty" xml:"CheckIds,omitempty"`
	CheckPlanId              *int64  `json:"CheckPlanId,omitempty" xml:"CheckPlanId,omitempty"`
	// example:
	//
	// Category
	GroupBy *string `json:"GroupBy,omitempty" xml:"GroupBy,omitempty"`
	// example:
	//
	// ecs
	Product                   *string `json:"Product,omitempty" xml:"Product,omitempty"`
	RegionIdsShrink           *string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty"`
	ResourceGroupIdListShrink *string `json:"ResourceGroupIdList,omitempty" xml:"ResourceGroupIdList,omitempty"`
	ResourceId                *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceIdsShrink         *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// example:
	//
	// SYNC_********_TASK
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// example:
	//
	// 1
	Severity        *int32  `json:"Severity,omitempty" xml:"Severity,omitempty"`
	TagKeysShrink   *string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty"`
	TagListShrink   *string `json:"TagList,omitempty" xml:"TagList,omitempty"`
	TagValuesShrink *string `json:"TagValues,omitempty" xml:"TagValues,omitempty"`
}

func (s DescribeCostCheckResultsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCostCheckResultsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeCostCheckResultsShrinkRequest) SetAssumeAliyunIdListShrink(v string) *DescribeCostCheckResultsShrinkRequest {
	s.AssumeAliyunIdListShrink = &v
	return s
}

func (s *DescribeCostCheckResultsShrinkRequest) SetCheckIdsShrink(v string) *DescribeCostCheckResultsShrinkRequest {
	s.CheckIdsShrink = &v
	return s
}

func (s *DescribeCostCheckResultsShrinkRequest) SetCheckPlanId(v int64) *DescribeCostCheckResultsShrinkRequest {
	s.CheckPlanId = &v
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

func (s *DescribeCostCheckResultsShrinkRequest) SetResourceGroupIdListShrink(v string) *DescribeCostCheckResultsShrinkRequest {
	s.ResourceGroupIdListShrink = &v
	return s
}

func (s *DescribeCostCheckResultsShrinkRequest) SetResourceId(v string) *DescribeCostCheckResultsShrinkRequest {
	s.ResourceId = &v
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

func (s *DescribeCostCheckResultsShrinkRequest) SetTagKeysShrink(v string) *DescribeCostCheckResultsShrinkRequest {
	s.TagKeysShrink = &v
	return s
}

func (s *DescribeCostCheckResultsShrinkRequest) SetTagListShrink(v string) *DescribeCostCheckResultsShrinkRequest {
	s.TagListShrink = &v
	return s
}

func (s *DescribeCostCheckResultsShrinkRequest) SetTagValuesShrink(v string) *DescribeCostCheckResultsShrinkRequest {
	s.TagValuesShrink = &v
	return s
}

type DescribeCostCheckResultsResponseBody struct {
	// example:
	//
	// 200
	Code *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DescribeCostCheckResultsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 566331F9-****-550F-B745-A730331F97A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
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
	AdviceResourceCount *int32 `json:"AdviceResourceCount,omitempty" xml:"AdviceResourceCount,omitempty"`
	// example:
	//
	// Category
	GroupBy *string `json:"GroupBy,omitempty" xml:"GroupBy,omitempty"`
	// example:
	//
	// 1
	NormalCount *int32 `json:"NormalCount,omitempty" xml:"NormalCount,omitempty"`
	// example:
	//
	// 76
	ResourceCount *int32 `json:"ResourceCount,omitempty" xml:"ResourceCount,omitempty"`
	// example:
	//
	// 4
	TotalCount *int32                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	ViewGroup  []*DescribeCostCheckResultsResponseBodyDataViewGroup `json:"ViewGroup,omitempty" xml:"ViewGroup,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	WarningCount *int32 `json:"WarningCount,omitempty" xml:"WarningCount,omitempty"`
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
	CheckItems []*DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems `json:"CheckItems,omitempty" xml:"CheckItems,omitempty" type:"Repeated"`
	// example:
	//
	// 22
	GroupCode *string `json:"GroupCode,omitempty" xml:"GroupCode,omitempty"`
	// example:
	//
	// 0
	GroupCount *int32 `json:"GroupCount,omitempty" xml:"GroupCount,omitempty"`
	// example:
	//
	// 1
	GroupExpectedSavingCost *float32 `json:"GroupExpectedSavingCost,omitempty" xml:"GroupExpectedSavingCost,omitempty"`
	// example:
	//
	// aut***8ainRh1
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
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
	// example:
	//
	// 100
	AdviceCount *int32 `json:"AdviceCount,omitempty" xml:"AdviceCount,omitempty"`
	// example:
	//
	// 1
	AdviceResourceCount *int32 `json:"AdviceResourceCount,omitempty" xml:"AdviceResourceCount,omitempty"`
	// example:
	//
	// 4
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// EbsCostIdleCheck
	CheckId   *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	CheckName *string `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	// example:
	//
	// 1
	CurrentCost *float32 `json:"CurrentCost,omitempty" xml:"CurrentCost,omitempty"`
	Description *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1
	ExpectedSavingCost *float32 `json:"ExpectedSavingCost,omitempty" xml:"ExpectedSavingCost,omitempty"`
	// example:
	//
	// 1
	OptimizedCost *float32 `json:"OptimizedCost,omitempty" xml:"OptimizedCost,omitempty"`
	// example:
	//
	// slb
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// example:
	//
	// 1
	Severity *int32 `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// example:
	//
	// true
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	Tips    *string `json:"Tips,omitempty" xml:"Tips,omitempty"`
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

type DescribeCostOptimizationOverviewRequest struct {
	// example:
	//
	// 11***********35
	AssumeAliyunId     *int64   `json:"AssumeAliyunId,omitempty" xml:"AssumeAliyunId,omitempty"`
	AssumeAliyunIdList []*int64 `json:"AssumeAliyunIdList,omitempty" xml:"AssumeAliyunIdList,omitempty" type:"Repeated"`
	CheckPlanId        *int64   `json:"CheckPlanId,omitempty" xml:"CheckPlanId,omitempty"`
	// example:
	//
	// ***
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s DescribeCostOptimizationOverviewRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCostOptimizationOverviewRequest) GoString() string {
	return s.String()
}

func (s *DescribeCostOptimizationOverviewRequest) SetAssumeAliyunId(v int64) *DescribeCostOptimizationOverviewRequest {
	s.AssumeAliyunId = &v
	return s
}

func (s *DescribeCostOptimizationOverviewRequest) SetAssumeAliyunIdList(v []*int64) *DescribeCostOptimizationOverviewRequest {
	s.AssumeAliyunIdList = v
	return s
}

func (s *DescribeCostOptimizationOverviewRequest) SetCheckPlanId(v int64) *DescribeCostOptimizationOverviewRequest {
	s.CheckPlanId = &v
	return s
}

func (s *DescribeCostOptimizationOverviewRequest) SetToken(v string) *DescribeCostOptimizationOverviewRequest {
	s.Token = &v
	return s
}

type DescribeCostOptimizationOverviewShrinkRequest struct {
	// example:
	//
	// 11***********35
	AssumeAliyunId           *int64  `json:"AssumeAliyunId,omitempty" xml:"AssumeAliyunId,omitempty"`
	AssumeAliyunIdListShrink *string `json:"AssumeAliyunIdList,omitempty" xml:"AssumeAliyunIdList,omitempty"`
	CheckPlanId              *int64  `json:"CheckPlanId,omitempty" xml:"CheckPlanId,omitempty"`
	// example:
	//
	// ***
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s DescribeCostOptimizationOverviewShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCostOptimizationOverviewShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeCostOptimizationOverviewShrinkRequest) SetAssumeAliyunId(v int64) *DescribeCostOptimizationOverviewShrinkRequest {
	s.AssumeAliyunId = &v
	return s
}

func (s *DescribeCostOptimizationOverviewShrinkRequest) SetAssumeAliyunIdListShrink(v string) *DescribeCostOptimizationOverviewShrinkRequest {
	s.AssumeAliyunIdListShrink = &v
	return s
}

func (s *DescribeCostOptimizationOverviewShrinkRequest) SetCheckPlanId(v int64) *DescribeCostOptimizationOverviewShrinkRequest {
	s.CheckPlanId = &v
	return s
}

func (s *DescribeCostOptimizationOverviewShrinkRequest) SetToken(v string) *DescribeCostOptimizationOverviewShrinkRequest {
	s.Token = &v
	return s
}

type DescribeCostOptimizationOverviewResponseBody struct {
	AccessDeniedDetail *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DescribeCostOptimizationOverviewResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Internal service issue. Detail:.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 566331F9-****-550F-B745-A730331F97A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCostOptimizationOverviewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCostOptimizationOverviewResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCostOptimizationOverviewResponseBody) SetAccessDeniedDetail(v *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail) *DescribeCostOptimizationOverviewResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBody) SetCode(v string) *DescribeCostOptimizationOverviewResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBody) SetData(v *DescribeCostOptimizationOverviewResponseBodyData) *DescribeCostOptimizationOverviewResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBody) SetMessage(v string) *DescribeCostOptimizationOverviewResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBody) SetRequestId(v string) *DescribeCostOptimizationOverviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBody) SetSuccess(v bool) *DescribeCostOptimizationOverviewResponseBody {
	s.Success = &v
	return s
}

type DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail struct {
	// AuthAction
	//
	// example:
	//
	// null
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// example:
	//
	// null
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// example:
	//
	// null
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// example:
	//
	// null
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// example:
	//
	// *****
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// example:
	//
	// null
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// example:
	//
	// PauseNotify
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

type DescribeCostOptimizationOverviewResponseBodyData struct {
	// example:
	//
	// 100
	BillingCycleDate *string `json:"BillingCycleDate,omitempty" xml:"BillingCycleDate,omitempty"`
	// example:
	//
	// 100
	CurrentBillingCost *string `json:"CurrentBillingCost,omitempty" xml:"CurrentBillingCost,omitempty"`
	// example:
	//
	// 100
	ExpectedSavingCost *string `json:"ExpectedSavingCost,omitempty" xml:"ExpectedSavingCost,omitempty"`
	// example:
	//
	// 2023-07-01 00:00:00
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 100
	OptCheckItemNum *string `json:"OptCheckItemNum,omitempty" xml:"OptCheckItemNum,omitempty"`
	// example:
	//
	// 100
	OptResourceNum         *string `json:"OptResourceNum,omitempty" xml:"OptResourceNum,omitempty"`
	ProcessedResourceCount *string `json:"ProcessedResourceCount,omitempty" xml:"ProcessedResourceCount,omitempty"`
	ProcessedSaveAmount    *string `json:"ProcessedSaveAmount,omitempty" xml:"ProcessedSaveAmount,omitempty"`
	// example:
	//
	// 95***135
	TaskId                   *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	WaitProcessResourceCount *string `json:"WaitProcessResourceCount,omitempty" xml:"WaitProcessResourceCount,omitempty"`
}

func (s DescribeCostOptimizationOverviewResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCostOptimizationOverviewResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCostOptimizationOverviewResponseBodyData) SetBillingCycleDate(v string) *DescribeCostOptimizationOverviewResponseBodyData {
	s.BillingCycleDate = &v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBodyData) SetCurrentBillingCost(v string) *DescribeCostOptimizationOverviewResponseBodyData {
	s.CurrentBillingCost = &v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBodyData) SetExpectedSavingCost(v string) *DescribeCostOptimizationOverviewResponseBodyData {
	s.ExpectedSavingCost = &v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBodyData) SetGmtModified(v int64) *DescribeCostOptimizationOverviewResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBodyData) SetOptCheckItemNum(v string) *DescribeCostOptimizationOverviewResponseBodyData {
	s.OptCheckItemNum = &v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBodyData) SetOptResourceNum(v string) *DescribeCostOptimizationOverviewResponseBodyData {
	s.OptResourceNum = &v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBodyData) SetProcessedResourceCount(v string) *DescribeCostOptimizationOverviewResponseBodyData {
	s.ProcessedResourceCount = &v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBodyData) SetProcessedSaveAmount(v string) *DescribeCostOptimizationOverviewResponseBodyData {
	s.ProcessedSaveAmount = &v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBodyData) SetTaskId(v int64) *DescribeCostOptimizationOverviewResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBodyData) SetWaitProcessResourceCount(v string) *DescribeCostOptimizationOverviewResponseBodyData {
	s.WaitProcessResourceCount = &v
	return s
}

type DescribeCostOptimizationOverviewResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCostOptimizationOverviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCostOptimizationOverviewResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCostOptimizationOverviewResponse) GoString() string {
	return s.String()
}

func (s *DescribeCostOptimizationOverviewResponse) SetHeaders(v map[string]*string) *DescribeCostOptimizationOverviewResponse {
	s.Headers = v
	return s
}

func (s *DescribeCostOptimizationOverviewResponse) SetStatusCode(v int32) *DescribeCostOptimizationOverviewResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCostOptimizationOverviewResponse) SetBody(v *DescribeCostOptimizationOverviewResponseBody) *DescribeCostOptimizationOverviewResponse {
	s.Body = v
	return s
}

type GetHistoryAdvicesRequest struct {
	// example:
	//
	// 2023-07-01
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// ecs
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// example:
	//
	// true
	Reverse *bool `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	// example:
	//
	// 1
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// example:
	//
	// 2023-07-01
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
	Data *GetHistoryAdvicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 566331F9-5AB3-550F-B745-A730331F97A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 1
	PageNo *int32                                     `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	Result []*GetHistoryAdvicesResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
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
	// example:
	//
	// EcsHighCpuUtilization
	CheckId     *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	CheckName   *string `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2023-07-01 00:00:00
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// example:
	//
	// ecs
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// example:
	//
	// i-bp67acfmxazb4p****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// 1
	Severity *int32  `json:"Severity,omitempty" xml:"Severity,omitempty"`
	Url      *string `json:"Url,omitempty" xml:"Url,omitempty"`
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

type GetInspectProgressRequest struct {
	// example:
	//
	// 14********37
	AssumeAliyunId *int64 `json:"AssumeAliyunId,omitempty" xml:"AssumeAliyunId,omitempty"`
	// example:
	//
	// 95***135
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// ***
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetInspectProgressRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInspectProgressRequest) GoString() string {
	return s.String()
}

func (s *GetInspectProgressRequest) SetAssumeAliyunId(v int64) *GetInspectProgressRequest {
	s.AssumeAliyunId = &v
	return s
}

func (s *GetInspectProgressRequest) SetTaskId(v int64) *GetInspectProgressRequest {
	s.TaskId = &v
	return s
}

func (s *GetInspectProgressRequest) SetToken(v string) *GetInspectProgressRequest {
	s.Token = &v
	return s
}

type GetInspectProgressResponseBody struct {
	// example:
	//
	// 200
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetInspectProgressResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 566331F9-****-550F-B745-A730331F97A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInspectProgressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInspectProgressResponseBody) GoString() string {
	return s.String()
}

func (s *GetInspectProgressResponseBody) SetCode(v string) *GetInspectProgressResponseBody {
	s.Code = &v
	return s
}

func (s *GetInspectProgressResponseBody) SetData(v *GetInspectProgressResponseBodyData) *GetInspectProgressResponseBody {
	s.Data = v
	return s
}

func (s *GetInspectProgressResponseBody) SetMessage(v string) *GetInspectProgressResponseBody {
	s.Message = &v
	return s
}

func (s *GetInspectProgressResponseBody) SetRequestId(v string) *GetInspectProgressResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInspectProgressResponseBody) SetSuccess(v bool) *GetInspectProgressResponseBody {
	s.Success = &v
	return s
}

type GetInspectProgressResponseBodyData struct {
	// example:
	//
	// 100
	AllSubtaskCount *int32 `json:"AllSubtaskCount,omitempty" xml:"AllSubtaskCount,omitempty"`
	// example:
	//
	// True
	Finish *bool `json:"Finish,omitempty" xml:"Finish,omitempty"`
	// example:
	//
	// 1
	FinishRate *float64 `json:"FinishRate,omitempty" xml:"FinishRate,omitempty"`
	// example:
	//
	// 1
	FinishSubtaskCount *int32 `json:"FinishSubtaskCount,omitempty" xml:"FinishSubtaskCount,omitempty"`
	// example:
	//
	// 111
	LastInspectDate *int64 `json:"LastInspectDate,omitempty" xml:"LastInspectDate,omitempty"`
	// example:
	//
	// 95***135
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 1
	UsedTime *int64 `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
}

func (s GetInspectProgressResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetInspectProgressResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInspectProgressResponseBodyData) SetAllSubtaskCount(v int32) *GetInspectProgressResponseBodyData {
	s.AllSubtaskCount = &v
	return s
}

func (s *GetInspectProgressResponseBodyData) SetFinish(v bool) *GetInspectProgressResponseBodyData {
	s.Finish = &v
	return s
}

func (s *GetInspectProgressResponseBodyData) SetFinishRate(v float64) *GetInspectProgressResponseBodyData {
	s.FinishRate = &v
	return s
}

func (s *GetInspectProgressResponseBodyData) SetFinishSubtaskCount(v int32) *GetInspectProgressResponseBodyData {
	s.FinishSubtaskCount = &v
	return s
}

func (s *GetInspectProgressResponseBodyData) SetLastInspectDate(v int64) *GetInspectProgressResponseBodyData {
	s.LastInspectDate = &v
	return s
}

func (s *GetInspectProgressResponseBodyData) SetTaskId(v int64) *GetInspectProgressResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetInspectProgressResponseBodyData) SetUsedTime(v int64) *GetInspectProgressResponseBodyData {
	s.UsedTime = &v
	return s
}

type GetInspectProgressResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInspectProgressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInspectProgressResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInspectProgressResponse) GoString() string {
	return s.String()
}

func (s *GetInspectProgressResponse) SetHeaders(v map[string]*string) *GetInspectProgressResponse {
	s.Headers = v
	return s
}

func (s *GetInspectProgressResponse) SetStatusCode(v int32) *GetInspectProgressResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInspectProgressResponse) SetBody(v *GetInspectProgressResponseBody) *GetInspectProgressResponse {
	s.Body = v
	return s
}

type GetProductListRequest struct {
	// example:
	//
	// ****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetProductListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProductListRequest) GoString() string {
	return s.String()
}

func (s *GetProductListRequest) SetToken(v string) *GetProductListRequest {
	s.Token = &v
	return s
}

type GetProductListResponseBody struct {
	AccessDeniedDetail *GetProductListResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*GetProductListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 566331F9-****-550F-B745-A730331F97A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// *
	UserMessage *string `json:"UserMessage,omitempty" xml:"UserMessage,omitempty"`
}

func (s GetProductListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProductListResponseBody) GoString() string {
	return s.String()
}

func (s *GetProductListResponseBody) SetAccessDeniedDetail(v *GetProductListResponseBodyAccessDeniedDetail) *GetProductListResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *GetProductListResponseBody) SetCode(v string) *GetProductListResponseBody {
	s.Code = &v
	return s
}

func (s *GetProductListResponseBody) SetData(v []*GetProductListResponseBodyData) *GetProductListResponseBody {
	s.Data = v
	return s
}

func (s *GetProductListResponseBody) SetMessage(v string) *GetProductListResponseBody {
	s.Message = &v
	return s
}

func (s *GetProductListResponseBody) SetRequestId(v string) *GetProductListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProductListResponseBody) SetSuccess(v bool) *GetProductListResponseBody {
	s.Success = &v
	return s
}

func (s *GetProductListResponseBody) SetUserMessage(v string) *GetProductListResponseBody {
	s.UserMessage = &v
	return s
}

type GetProductListResponseBodyAccessDeniedDetail struct {
	// example:
	//
	// *
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// example:
	//
	// *
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// example:
	//
	// *
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// example:
	//
	// *
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// example:
	//
	// ****
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// example:
	//
	// *
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// example:
	//
	// *
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s GetProductListResponseBodyAccessDeniedDetail) String() string {
	return tea.Prettify(s)
}

func (s GetProductListResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *GetProductListResponseBodyAccessDeniedDetail) SetAuthAction(v string) *GetProductListResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *GetProductListResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *GetProductListResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *GetProductListResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *GetProductListResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *GetProductListResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *GetProductListResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *GetProductListResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *GetProductListResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *GetProductListResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *GetProductListResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *GetProductListResponseBodyAccessDeniedDetail) SetPolicyType(v string) *GetProductListResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

type GetProductListResponseBodyData struct {
	Category    *string                                      `json:"Category,omitempty" xml:"Category,omitempty"`
	ProductList []*GetProductListResponseBodyDataProductList `json:"ProductList,omitempty" xml:"ProductList,omitempty" type:"Repeated"`
}

func (s GetProductListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetProductListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetProductListResponseBodyData) SetCategory(v string) *GetProductListResponseBodyData {
	s.Category = &v
	return s
}

func (s *GetProductListResponseBodyData) SetProductList(v []*GetProductListResponseBodyDataProductList) *GetProductListResponseBodyData {
	s.ProductList = v
	return s
}

type GetProductListResponseBodyDataProductList struct {
	NewLabel *string `json:"NewLabel,omitempty" xml:"NewLabel,omitempty"`
	// example:
	//
	// hologres
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
}

func (s GetProductListResponseBodyDataProductList) String() string {
	return tea.Prettify(s)
}

func (s GetProductListResponseBodyDataProductList) GoString() string {
	return s.String()
}

func (s *GetProductListResponseBodyDataProductList) SetNewLabel(v string) *GetProductListResponseBodyDataProductList {
	s.NewLabel = &v
	return s
}

func (s *GetProductListResponseBodyDataProductList) SetProduct(v string) *GetProductListResponseBodyDataProductList {
	s.Product = &v
	return s
}

type GetProductListResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProductListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProductListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProductListResponse) GoString() string {
	return s.String()
}

func (s *GetProductListResponse) SetHeaders(v map[string]*string) *GetProductListResponse {
	s.Headers = v
	return s
}

func (s *GetProductListResponse) SetStatusCode(v int32) *GetProductListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProductListResponse) SetBody(v *GetProductListResponseBody) *GetProductListResponse {
	s.Body = v
	return s
}

type GetTaskStatusByIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 95906135
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
	Data *GetTaskStatusByIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 566331F9-5AB3-550F-B745-A730331F97A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 95906135
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 1
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
	AssumeAliyunId *int64 `json:"AssumeAliyunId,omitempty" xml:"AssumeAliyunId,omitempty"`
	// example:
	//
	// EcsHighCpuUtilization
	CheckId     *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	CheckPlanId *int64  `json:"CheckPlanId,omitempty" xml:"CheckPlanId,omitempty"`
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// ecs
	Product               *string                                            `json:"Product,omitempty" xml:"Product,omitempty"`
	ResourceDimensionList []*RefreshAdvisorCheckRequestResourceDimensionList `json:"ResourceDimensionList,omitempty" xml:"ResourceDimensionList,omitempty" type:"Repeated"`
	// example:
	//
	// i-bp67acfmxazb4p****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Token      *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s RefreshAdvisorCheckRequest) String() string {
	return tea.Prettify(s)
}

func (s RefreshAdvisorCheckRequest) GoString() string {
	return s.String()
}

func (s *RefreshAdvisorCheckRequest) SetAssumeAliyunId(v int64) *RefreshAdvisorCheckRequest {
	s.AssumeAliyunId = &v
	return s
}

func (s *RefreshAdvisorCheckRequest) SetCheckId(v string) *RefreshAdvisorCheckRequest {
	s.CheckId = &v
	return s
}

func (s *RefreshAdvisorCheckRequest) SetCheckPlanId(v int64) *RefreshAdvisorCheckRequest {
	s.CheckPlanId = &v
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

func (s *RefreshAdvisorCheckRequest) SetResourceDimensionList(v []*RefreshAdvisorCheckRequestResourceDimensionList) *RefreshAdvisorCheckRequest {
	s.ResourceDimensionList = v
	return s
}

func (s *RefreshAdvisorCheckRequest) SetResourceId(v string) *RefreshAdvisorCheckRequest {
	s.ResourceId = &v
	return s
}

func (s *RefreshAdvisorCheckRequest) SetToken(v string) *RefreshAdvisorCheckRequest {
	s.Token = &v
	return s
}

type RefreshAdvisorCheckRequestResourceDimensionList struct {
	Cost         *bool   `json:"Cost,omitempty" xml:"Cost,omitempty"`
	Performance  *bool   `json:"Performance,omitempty" xml:"Performance,omitempty"`
	Product      *string `json:"Product,omitempty" xml:"Product,omitempty"`
	ProductName  *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	Reliablility *bool   `json:"Reliablility,omitempty" xml:"Reliablility,omitempty"`
	Security     *bool   `json:"Security,omitempty" xml:"Security,omitempty"`
	Service      *bool   `json:"Service,omitempty" xml:"Service,omitempty"`
}

func (s RefreshAdvisorCheckRequestResourceDimensionList) String() string {
	return tea.Prettify(s)
}

func (s RefreshAdvisorCheckRequestResourceDimensionList) GoString() string {
	return s.String()
}

func (s *RefreshAdvisorCheckRequestResourceDimensionList) SetCost(v bool) *RefreshAdvisorCheckRequestResourceDimensionList {
	s.Cost = &v
	return s
}

func (s *RefreshAdvisorCheckRequestResourceDimensionList) SetPerformance(v bool) *RefreshAdvisorCheckRequestResourceDimensionList {
	s.Performance = &v
	return s
}

func (s *RefreshAdvisorCheckRequestResourceDimensionList) SetProduct(v string) *RefreshAdvisorCheckRequestResourceDimensionList {
	s.Product = &v
	return s
}

func (s *RefreshAdvisorCheckRequestResourceDimensionList) SetProductName(v string) *RefreshAdvisorCheckRequestResourceDimensionList {
	s.ProductName = &v
	return s
}

func (s *RefreshAdvisorCheckRequestResourceDimensionList) SetReliablility(v bool) *RefreshAdvisorCheckRequestResourceDimensionList {
	s.Reliablility = &v
	return s
}

func (s *RefreshAdvisorCheckRequestResourceDimensionList) SetSecurity(v bool) *RefreshAdvisorCheckRequestResourceDimensionList {
	s.Security = &v
	return s
}

func (s *RefreshAdvisorCheckRequestResourceDimensionList) SetService(v bool) *RefreshAdvisorCheckRequestResourceDimensionList {
	s.Service = &v
	return s
}

type RefreshAdvisorCheckShrinkRequest struct {
	AssumeAliyunId *int64 `json:"AssumeAliyunId,omitempty" xml:"AssumeAliyunId,omitempty"`
	// example:
	//
	// EcsHighCpuUtilization
	CheckId     *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	CheckPlanId *int64  `json:"CheckPlanId,omitempty" xml:"CheckPlanId,omitempty"`
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// ecs
	Product                     *string `json:"Product,omitempty" xml:"Product,omitempty"`
	ResourceDimensionListShrink *string `json:"ResourceDimensionList,omitempty" xml:"ResourceDimensionList,omitempty"`
	// example:
	//
	// i-bp67acfmxazb4p****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Token      *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s RefreshAdvisorCheckShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RefreshAdvisorCheckShrinkRequest) GoString() string {
	return s.String()
}

func (s *RefreshAdvisorCheckShrinkRequest) SetAssumeAliyunId(v int64) *RefreshAdvisorCheckShrinkRequest {
	s.AssumeAliyunId = &v
	return s
}

func (s *RefreshAdvisorCheckShrinkRequest) SetCheckId(v string) *RefreshAdvisorCheckShrinkRequest {
	s.CheckId = &v
	return s
}

func (s *RefreshAdvisorCheckShrinkRequest) SetCheckPlanId(v int64) *RefreshAdvisorCheckShrinkRequest {
	s.CheckPlanId = &v
	return s
}

func (s *RefreshAdvisorCheckShrinkRequest) SetLanguage(v string) *RefreshAdvisorCheckShrinkRequest {
	s.Language = &v
	return s
}

func (s *RefreshAdvisorCheckShrinkRequest) SetProduct(v string) *RefreshAdvisorCheckShrinkRequest {
	s.Product = &v
	return s
}

func (s *RefreshAdvisorCheckShrinkRequest) SetResourceDimensionListShrink(v string) *RefreshAdvisorCheckShrinkRequest {
	s.ResourceDimensionListShrink = &v
	return s
}

func (s *RefreshAdvisorCheckShrinkRequest) SetResourceId(v string) *RefreshAdvisorCheckShrinkRequest {
	s.ResourceId = &v
	return s
}

func (s *RefreshAdvisorCheckShrinkRequest) SetToken(v string) *RefreshAdvisorCheckShrinkRequest {
	s.Token = &v
	return s
}

type RefreshAdvisorCheckResponseBody struct {
	Data *RefreshAdvisorCheckResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 566331F9-5AB3-550F-B745-A730331F97A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 12345678
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// manual-1be17af1121b4974822e69daee4f2481
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RefreshAdvisorCheckResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RefreshAdvisorCheckResponseBodyData) GoString() string {
	return s.String()
}

func (s *RefreshAdvisorCheckResponseBodyData) SetMessage(v string) *RefreshAdvisorCheckResponseBodyData {
	s.Message = &v
	return s
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

type RefreshAdvisorCostCheckRequest struct {
	AssumeAliyunIdList []*int64  `json:"AssumeAliyunIdList,omitempty" xml:"AssumeAliyunIdList,omitempty" type:"Repeated"`
	CheckIds           []*string `json:"CheckIds,omitempty" xml:"CheckIds,omitempty" type:"Repeated"`
	CheckPlanId        *int64    `json:"CheckPlanId,omitempty" xml:"CheckPlanId,omitempty"`
	// example:
	//
	// ecs
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// example:
	//
	// false
	RefreshResource *bool     `json:"RefreshResource,omitempty" xml:"RefreshResource,omitempty"`
	ResourceIds     []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
}

func (s RefreshAdvisorCostCheckRequest) String() string {
	return tea.Prettify(s)
}

func (s RefreshAdvisorCostCheckRequest) GoString() string {
	return s.String()
}

func (s *RefreshAdvisorCostCheckRequest) SetAssumeAliyunIdList(v []*int64) *RefreshAdvisorCostCheckRequest {
	s.AssumeAliyunIdList = v
	return s
}

func (s *RefreshAdvisorCostCheckRequest) SetCheckIds(v []*string) *RefreshAdvisorCostCheckRequest {
	s.CheckIds = v
	return s
}

func (s *RefreshAdvisorCostCheckRequest) SetCheckPlanId(v int64) *RefreshAdvisorCostCheckRequest {
	s.CheckPlanId = &v
	return s
}

func (s *RefreshAdvisorCostCheckRequest) SetProduct(v string) *RefreshAdvisorCostCheckRequest {
	s.Product = &v
	return s
}

func (s *RefreshAdvisorCostCheckRequest) SetRefreshResource(v bool) *RefreshAdvisorCostCheckRequest {
	s.RefreshResource = &v
	return s
}

func (s *RefreshAdvisorCostCheckRequest) SetResourceIds(v []*string) *RefreshAdvisorCostCheckRequest {
	s.ResourceIds = v
	return s
}

type RefreshAdvisorCostCheckShrinkRequest struct {
	AssumeAliyunIdListShrink *string `json:"AssumeAliyunIdList,omitempty" xml:"AssumeAliyunIdList,omitempty"`
	CheckIdsShrink           *string `json:"CheckIds,omitempty" xml:"CheckIds,omitempty"`
	CheckPlanId              *int64  `json:"CheckPlanId,omitempty" xml:"CheckPlanId,omitempty"`
	// example:
	//
	// ecs
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// example:
	//
	// false
	RefreshResource   *bool   `json:"RefreshResource,omitempty" xml:"RefreshResource,omitempty"`
	ResourceIdsShrink *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
}

func (s RefreshAdvisorCostCheckShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RefreshAdvisorCostCheckShrinkRequest) GoString() string {
	return s.String()
}

func (s *RefreshAdvisorCostCheckShrinkRequest) SetAssumeAliyunIdListShrink(v string) *RefreshAdvisorCostCheckShrinkRequest {
	s.AssumeAliyunIdListShrink = &v
	return s
}

func (s *RefreshAdvisorCostCheckShrinkRequest) SetCheckIdsShrink(v string) *RefreshAdvisorCostCheckShrinkRequest {
	s.CheckIdsShrink = &v
	return s
}

func (s *RefreshAdvisorCostCheckShrinkRequest) SetCheckPlanId(v int64) *RefreshAdvisorCostCheckShrinkRequest {
	s.CheckPlanId = &v
	return s
}

func (s *RefreshAdvisorCostCheckShrinkRequest) SetProduct(v string) *RefreshAdvisorCostCheckShrinkRequest {
	s.Product = &v
	return s
}

func (s *RefreshAdvisorCostCheckShrinkRequest) SetRefreshResource(v bool) *RefreshAdvisorCostCheckShrinkRequest {
	s.RefreshResource = &v
	return s
}

func (s *RefreshAdvisorCostCheckShrinkRequest) SetResourceIdsShrink(v string) *RefreshAdvisorCostCheckShrinkRequest {
	s.ResourceIdsShrink = &v
	return s
}

type RefreshAdvisorCostCheckResponseBody struct {
	// example:
	//
	// 200
	Code *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *RefreshAdvisorCostCheckResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 566331F9-****-550F-B745-A730331F97A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RefreshAdvisorCostCheckResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefreshAdvisorCostCheckResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshAdvisorCostCheckResponseBody) SetCode(v string) *RefreshAdvisorCostCheckResponseBody {
	s.Code = &v
	return s
}

func (s *RefreshAdvisorCostCheckResponseBody) SetData(v *RefreshAdvisorCostCheckResponseBodyData) *RefreshAdvisorCostCheckResponseBody {
	s.Data = v
	return s
}

func (s *RefreshAdvisorCostCheckResponseBody) SetMessage(v string) *RefreshAdvisorCostCheckResponseBody {
	s.Message = &v
	return s
}

func (s *RefreshAdvisorCostCheckResponseBody) SetRequestId(v string) *RefreshAdvisorCostCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshAdvisorCostCheckResponseBody) SetSuccess(v bool) *RefreshAdvisorCostCheckResponseBody {
	s.Success = &v
	return s
}

type RefreshAdvisorCostCheckResponseBodyData struct {
	// example:
	//
	// c-wl*****n0g
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// example:
	//
	// 11***********35
	ManagerTaskId *int64 `json:"ManagerTaskId,omitempty" xml:"ManagerTaskId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 959***135
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RefreshAdvisorCostCheckResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RefreshAdvisorCostCheckResponseBodyData) GoString() string {
	return s.String()
}

func (s *RefreshAdvisorCostCheckResponseBodyData) SetCommandId(v string) *RefreshAdvisorCostCheckResponseBodyData {
	s.CommandId = &v
	return s
}

func (s *RefreshAdvisorCostCheckResponseBodyData) SetManagerTaskId(v int64) *RefreshAdvisorCostCheckResponseBodyData {
	s.ManagerTaskId = &v
	return s
}

func (s *RefreshAdvisorCostCheckResponseBodyData) SetSuccess(v bool) *RefreshAdvisorCostCheckResponseBodyData {
	s.Success = &v
	return s
}

func (s *RefreshAdvisorCostCheckResponseBodyData) SetTaskId(v int64) *RefreshAdvisorCostCheckResponseBodyData {
	s.TaskId = &v
	return s
}

type RefreshAdvisorCostCheckResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshAdvisorCostCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshAdvisorCostCheckResponse) String() string {
	return tea.Prettify(s)
}

func (s RefreshAdvisorCostCheckResponse) GoString() string {
	return s.String()
}

func (s *RefreshAdvisorCostCheckResponse) SetHeaders(v map[string]*string) *RefreshAdvisorCostCheckResponse {
	s.Headers = v
	return s
}

func (s *RefreshAdvisorCostCheckResponse) SetStatusCode(v int32) *RefreshAdvisorCostCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshAdvisorCostCheckResponse) SetBody(v *RefreshAdvisorCostCheckResponseBody) *RefreshAdvisorCostCheckResponse {
	s.Body = v
	return s
}

type RefreshAdvisorResourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ecs
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// example:
	//
	// i-bp67acfmxazb4p****
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
	// example:
	//
	// 12345678
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 566331F9-5AB3-550F-B745-A730331F97A9
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
	AlertDescription *string `json:"AlertDescription,omitempty" xml:"AlertDescription,omitempty"`
	// This parameter is required.
	AlertDetail *string `json:"AlertDetail,omitempty" xml:"AlertDetail,omitempty"`
	AlertGrade  *string `json:"AlertGrade,omitempty" xml:"AlertGrade,omitempty"`
	AlertLabels *string `json:"AlertLabels,omitempty" xml:"AlertLabels,omitempty"`
	// This parameter is required.
	AlertScene *string `json:"AlertScene,omitempty" xml:"AlertScene,omitempty"`
	// This parameter is required.
	AlertToken *string  `json:"AlertToken,omitempty" xml:"AlertToken,omitempty"`
	AlertUid   []*int64 `json:"AlertUid,omitempty" xml:"AlertUid,omitempty" type:"Repeated"`
	Language   *string  `json:"Language,omitempty" xml:"Language,omitempty"`
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
	// This parameter is required.
	AlertDetail *string `json:"AlertDetail,omitempty" xml:"AlertDetail,omitempty"`
	AlertGrade  *string `json:"AlertGrade,omitempty" xml:"AlertGrade,omitempty"`
	AlertLabels *string `json:"AlertLabels,omitempty" xml:"AlertLabels,omitempty"`
	// This parameter is required.
	AlertScene *string `json:"AlertScene,omitempty" xml:"AlertScene,omitempty"`
	// This parameter is required.
	AlertToken     *string `json:"AlertToken,omitempty" xml:"AlertToken,omitempty"`
	AlertUidShrink *string `json:"AlertUid,omitempty" xml:"AlertUid,omitempty"`
	Language       *string `json:"Language,omitempty" xml:"Language,omitempty"`
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

// Summary:
//
// 根据多个维度获取用户最新的巡检结果，全量返回-openApi
//
// @param request - DescribeAdvicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAdvicesResponse
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

	if !tea.BoolValue(util.IsUnset(request.CheckPlanId)) {
		query["CheckPlanId"] = request.CheckPlanId
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

// Summary:
//
// 根据多个维度获取用户最新的巡检结果，全量返回-openApi
//
// @param request - DescribeAdvicesRequest
//
// @return DescribeAdvicesResponse
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

// Summary:
//
// # DescribeAdvicesFlat分页
//
// @param request - DescribeAdvicesFlatPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAdvicesFlatPageResponse
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

// Summary:
//
// # DescribeAdvicesFlat分页
//
// @param request - DescribeAdvicesFlatPageRequest
//
// @return DescribeAdvicesFlatPageResponse
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

// Summary:
//
// # DescribeAdvices分页
//
// @param request - DescribeAdvicesPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAdvicesPageResponse
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

	if !tea.BoolValue(util.IsUnset(request.CheckPlanId)) {
		query["CheckPlanId"] = request.CheckPlanId
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

// Summary:
//
// # DescribeAdvices分页
//
// @param request - DescribeAdvicesPageRequest
//
// @return DescribeAdvicesPageResponse
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

// @param request - DescribeAdvisorChecksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAdvisorChecksResponse
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

// @param request - DescribeAdvisorChecksRequest
//
// @return DescribeAdvisorChecksResponse
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

// Summary:
//
// 巡检项设置-分页
//
// @param tmpReq - DescribeAdvisorChecksFoPagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAdvisorChecksFoPagesResponse
func (client *Client) DescribeAdvisorChecksFoPagesWithOptions(tmpReq *DescribeAdvisorChecksFoPagesRequest, runtime *util.RuntimeOptions) (_result *DescribeAdvisorChecksFoPagesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DescribeAdvisorChecksFoPagesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CheckTypes)) {
		request.CheckTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CheckTypes, tea.String("CheckTypes"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssumeAliyunId)) {
		query["AssumeAliyunId"] = request.AssumeAliyunId
	}

	if !tea.BoolValue(util.IsUnset(request.BizCategory)) {
		query["BizCategory"] = request.BizCategory
	}

	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.CheckTypesShrink)) {
		query["CheckTypes"] = request.CheckTypesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
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

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["Token"] = request.Token
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAdvisorChecksFoPages"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAdvisorChecksFoPagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 巡检项设置-分页
//
// @param request - DescribeAdvisorChecksFoPagesRequest
//
// @return DescribeAdvisorChecksFoPagesResponse
func (client *Client) DescribeAdvisorChecksFoPages(request *DescribeAdvisorChecksFoPagesRequest) (_result *DescribeAdvisorChecksFoPagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAdvisorChecksFoPagesResponse{}
	_body, _err := client.DescribeAdvisorChecksFoPagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeAdvisorResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAdvisorResourcesResponse
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

// @param request - DescribeAdvisorResourcesRequest
//
// @return DescribeAdvisorResourcesResponse
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

// Summary:
//
// 查询成本优化结果详情
//
// @param tmpReq - DescribeCostCheckAdvicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCostCheckAdvicesResponse
func (client *Client) DescribeCostCheckAdvicesWithOptions(tmpReq *DescribeCostCheckAdvicesRequest, runtime *util.RuntimeOptions) (_result *DescribeCostCheckAdvicesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DescribeCostCheckAdvicesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AssumeAliyunIdList)) {
		request.AssumeAliyunIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AssumeAliyunIdList, tea.String("AssumeAliyunIdList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.RegionIds)) {
		request.RegionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RegionIds, tea.String("RegionIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceGroupIdList)) {
		request.ResourceGroupIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceGroupIdList, tea.String("ResourceGroupIdList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceIds)) {
		request.ResourceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceIds, tea.String("ResourceIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TagKeys)) {
		request.TagKeysShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagKeys, tea.String("TagKeys"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TagList)) {
		request.TagListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagList, tea.String("TagList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TagValues)) {
		request.TagValuesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagValues, tea.String("TagValues"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssumeAliyunIdListShrink)) {
		query["AssumeAliyunIdList"] = request.AssumeAliyunIdListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.CheckId)) {
		query["CheckId"] = request.CheckId
	}

	if !tea.BoolValue(util.IsUnset(request.CheckPlanId)) {
		query["CheckPlanId"] = request.CheckPlanId
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

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupIdListShrink)) {
		query["ResourceGroupIdList"] = request.ResourceGroupIdListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
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

	if !tea.BoolValue(util.IsUnset(request.TagKeysShrink)) {
		query["TagKeys"] = request.TagKeysShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TagListShrink)) {
		query["TagList"] = request.TagListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TagValuesShrink)) {
		query["TagValues"] = request.TagValuesShrink
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

// Summary:
//
// 查询成本优化结果详情
//
// @param request - DescribeCostCheckAdvicesRequest
//
// @return DescribeCostCheckAdvicesResponse
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

// Summary:
//
// 查询巡检项聚合成本优化结果概览
//
// @param tmpReq - DescribeCostCheckResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCostCheckResultsResponse
func (client *Client) DescribeCostCheckResultsWithOptions(tmpReq *DescribeCostCheckResultsRequest, runtime *util.RuntimeOptions) (_result *DescribeCostCheckResultsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DescribeCostCheckResultsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AssumeAliyunIdList)) {
		request.AssumeAliyunIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AssumeAliyunIdList, tea.String("AssumeAliyunIdList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.CheckIds)) {
		request.CheckIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CheckIds, tea.String("CheckIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.RegionIds)) {
		request.RegionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RegionIds, tea.String("RegionIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceGroupIdList)) {
		request.ResourceGroupIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceGroupIdList, tea.String("ResourceGroupIdList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceIds)) {
		request.ResourceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceIds, tea.String("ResourceIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TagKeys)) {
		request.TagKeysShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagKeys, tea.String("TagKeys"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TagList)) {
		request.TagListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagList, tea.String("TagList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TagValues)) {
		request.TagValuesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagValues, tea.String("TagValues"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssumeAliyunIdListShrink)) {
		query["AssumeAliyunIdList"] = request.AssumeAliyunIdListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.CheckIdsShrink)) {
		query["CheckIds"] = request.CheckIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.CheckPlanId)) {
		query["CheckPlanId"] = request.CheckPlanId
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

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupIdListShrink)) {
		query["ResourceGroupIdList"] = request.ResourceGroupIdListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
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

	if !tea.BoolValue(util.IsUnset(request.TagKeysShrink)) {
		query["TagKeys"] = request.TagKeysShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TagListShrink)) {
		query["TagList"] = request.TagListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TagValuesShrink)) {
		query["TagValues"] = request.TagValuesShrink
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

// Summary:
//
// 查询巡检项聚合成本优化结果概览
//
// @param request - DescribeCostCheckResultsRequest
//
// @return DescribeCostCheckResultsResponse
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

// Summary:
//
// 成本优化-概览
//
// @param tmpReq - DescribeCostOptimizationOverviewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCostOptimizationOverviewResponse
func (client *Client) DescribeCostOptimizationOverviewWithOptions(tmpReq *DescribeCostOptimizationOverviewRequest, runtime *util.RuntimeOptions) (_result *DescribeCostOptimizationOverviewResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DescribeCostOptimizationOverviewShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AssumeAliyunIdList)) {
		request.AssumeAliyunIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AssumeAliyunIdList, tea.String("AssumeAliyunIdList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssumeAliyunId)) {
		query["AssumeAliyunId"] = request.AssumeAliyunId
	}

	if !tea.BoolValue(util.IsUnset(request.AssumeAliyunIdListShrink)) {
		query["AssumeAliyunIdList"] = request.AssumeAliyunIdListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.CheckPlanId)) {
		query["CheckPlanId"] = request.CheckPlanId
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["Token"] = request.Token
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCostOptimizationOverview"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCostOptimizationOverviewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 成本优化-概览
//
// @param request - DescribeCostOptimizationOverviewRequest
//
// @return DescribeCostOptimizationOverviewResponse
func (client *Client) DescribeCostOptimizationOverview(request *DescribeCostOptimizationOverviewRequest) (_result *DescribeCostOptimizationOverviewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCostOptimizationOverviewResponse{}
	_body, _err := client.DescribeCostOptimizationOverviewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetHistoryAdvicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHistoryAdvicesResponse
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

// @param request - GetHistoryAdvicesRequest
//
// @return GetHistoryAdvicesResponse
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

// Summary:
//
// 获取任务执行进度(普通用户、RD单账号)
//
// @param request - GetInspectProgressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInspectProgressResponse
func (client *Client) GetInspectProgressWithOptions(request *GetInspectProgressRequest, runtime *util.RuntimeOptions) (_result *GetInspectProgressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssumeAliyunId)) {
		query["AssumeAliyunId"] = request.AssumeAliyunId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["Token"] = request.Token
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInspectProgress"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInspectProgressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取任务执行进度(普通用户、RD单账号)
//
// @param request - GetInspectProgressRequest
//
// @return GetInspectProgressResponse
func (client *Client) GetInspectProgress(request *GetInspectProgressRequest) (_result *GetInspectProgressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInspectProgressResponse{}
	_body, _err := client.GetInspectProgressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取云产品的列表
//
// @param request - GetProductListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProductListResponse
func (client *Client) GetProductListWithOptions(request *GetProductListRequest, runtime *util.RuntimeOptions) (_result *GetProductListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["Token"] = request.Token
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetProductList"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProductListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取云产品的列表
//
// @param request - GetProductListRequest
//
// @return GetProductListResponse
func (client *Client) GetProductList(request *GetProductListRequest) (_result *GetProductListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetProductListResponse{}
	_body, _err := client.GetProductListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据id获取任务状态
//
// @param request - GetTaskStatusByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskStatusByIdResponse
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

// Summary:
//
// 根据id获取任务状态
//
// @param request - GetTaskStatusByIdRequest
//
// @return GetTaskStatusByIdResponse
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

// Summary:
//
// 触发立即巡检
//
// @param tmpReq - RefreshAdvisorCheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshAdvisorCheckResponse
func (client *Client) RefreshAdvisorCheckWithOptions(tmpReq *RefreshAdvisorCheckRequest, runtime *util.RuntimeOptions) (_result *RefreshAdvisorCheckResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RefreshAdvisorCheckShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceDimensionList)) {
		request.ResourceDimensionListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceDimensionList, tea.String("ResourceDimensionList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssumeAliyunId)) {
		query["AssumeAliyunId"] = request.AssumeAliyunId
	}

	if !tea.BoolValue(util.IsUnset(request.CheckId)) {
		query["CheckId"] = request.CheckId
	}

	if !tea.BoolValue(util.IsUnset(request.CheckPlanId)) {
		query["CheckPlanId"] = request.CheckPlanId
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

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["Token"] = request.Token
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceDimensionListShrink)) {
		body["ResourceDimensionList"] = request.ResourceDimensionListShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
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

// Summary:
//
// 触发立即巡检
//
// @param request - RefreshAdvisorCheckRequest
//
// @return RefreshAdvisorCheckResponse
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

// Summary:
//
// 发起成本优化巡检
//
// @param tmpReq - RefreshAdvisorCostCheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshAdvisorCostCheckResponse
func (client *Client) RefreshAdvisorCostCheckWithOptions(tmpReq *RefreshAdvisorCostCheckRequest, runtime *util.RuntimeOptions) (_result *RefreshAdvisorCostCheckResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RefreshAdvisorCostCheckShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AssumeAliyunIdList)) {
		request.AssumeAliyunIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AssumeAliyunIdList, tea.String("AssumeAliyunIdList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.CheckIds)) {
		request.CheckIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CheckIds, tea.String("CheckIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceIds)) {
		request.ResourceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceIds, tea.String("ResourceIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssumeAliyunIdListShrink)) {
		query["AssumeAliyunIdList"] = request.AssumeAliyunIdListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.CheckIdsShrink)) {
		query["CheckIds"] = request.CheckIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.CheckPlanId)) {
		query["CheckPlanId"] = request.CheckPlanId
	}

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		query["Product"] = request.Product
	}

	if !tea.BoolValue(util.IsUnset(request.RefreshResource)) {
		query["RefreshResource"] = request.RefreshResource
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceIdsShrink)) {
		query["ResourceIds"] = request.ResourceIdsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RefreshAdvisorCostCheck"),
		Version:     tea.String("2018-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RefreshAdvisorCostCheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发起成本优化巡检
//
// @param request - RefreshAdvisorCostCheckRequest
//
// @return RefreshAdvisorCostCheckResponse
func (client *Client) RefreshAdvisorCostCheck(request *RefreshAdvisorCostCheckRequest) (_result *RefreshAdvisorCostCheckResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefreshAdvisorCostCheckResponse{}
	_body, _err := client.RefreshAdvisorCostCheckWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # RefreshAdvisorResource
//
// @param request - RefreshAdvisorResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshAdvisorResourceResponse
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

// Summary:
//
// # RefreshAdvisorResource
//
// @param request - RefreshAdvisorResourceRequest
//
// @return RefreshAdvisorResourceResponse
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

// Summary:
//
// 上报用户业务报警信息
//
// @param tmpReq - ReportBizAlertInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReportBizAlertInfoResponse
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

// Summary:
//
// 上报用户业务报警信息
//
// @param request - ReportBizAlertInfoRequest
//
// @return ReportBizAlertInfoResponse
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
