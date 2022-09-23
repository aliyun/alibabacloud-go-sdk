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

type AddCdnDomainRequest struct {
	CdnType         *string `json:"CdnType,omitempty" xml:"CdnType,omitempty"`
	CheckUrl        *string `json:"CheckUrl,omitempty" xml:"CheckUrl,omitempty"`
	DomainName      *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Priorities      *string `json:"Priorities,omitempty" xml:"Priorities,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Scope           *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	SourcePort      *int32  `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	SourceType      *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Sources         *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	TopLevelDomain  *string `json:"TopLevelDomain,omitempty" xml:"TopLevelDomain,omitempty"`
}

func (s AddCdnDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s AddCdnDomainRequest) GoString() string {
	return s.String()
}

func (s *AddCdnDomainRequest) SetCdnType(v string) *AddCdnDomainRequest {
	s.CdnType = &v
	return s
}

func (s *AddCdnDomainRequest) SetCheckUrl(v string) *AddCdnDomainRequest {
	s.CheckUrl = &v
	return s
}

func (s *AddCdnDomainRequest) SetDomainName(v string) *AddCdnDomainRequest {
	s.DomainName = &v
	return s
}

func (s *AddCdnDomainRequest) SetOwnerAccount(v string) *AddCdnDomainRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddCdnDomainRequest) SetOwnerId(v int64) *AddCdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *AddCdnDomainRequest) SetPriorities(v string) *AddCdnDomainRequest {
	s.Priorities = &v
	return s
}

func (s *AddCdnDomainRequest) SetRegion(v string) *AddCdnDomainRequest {
	s.Region = &v
	return s
}

func (s *AddCdnDomainRequest) SetResourceGroupId(v string) *AddCdnDomainRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AddCdnDomainRequest) SetScope(v string) *AddCdnDomainRequest {
	s.Scope = &v
	return s
}

func (s *AddCdnDomainRequest) SetSecurityToken(v string) *AddCdnDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *AddCdnDomainRequest) SetSourcePort(v int32) *AddCdnDomainRequest {
	s.SourcePort = &v
	return s
}

func (s *AddCdnDomainRequest) SetSourceType(v string) *AddCdnDomainRequest {
	s.SourceType = &v
	return s
}

func (s *AddCdnDomainRequest) SetSources(v string) *AddCdnDomainRequest {
	s.Sources = &v
	return s
}

func (s *AddCdnDomainRequest) SetTopLevelDomain(v string) *AddCdnDomainRequest {
	s.TopLevelDomain = &v
	return s
}

type AddCdnDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddCdnDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddCdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *AddCdnDomainResponseBody) SetRequestId(v string) *AddCdnDomainResponseBody {
	s.RequestId = &v
	return s
}

type AddCdnDomainResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddCdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddCdnDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s AddCdnDomainResponse) GoString() string {
	return s.String()
}

func (s *AddCdnDomainResponse) SetHeaders(v map[string]*string) *AddCdnDomainResponse {
	s.Headers = v
	return s
}

func (s *AddCdnDomainResponse) SetStatusCode(v int32) *AddCdnDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCdnDomainResponse) SetBody(v *AddCdnDomainResponseBody) *AddCdnDomainResponse {
	s.Body = v
	return s
}

type DescribeCdnDomainDetailRequest struct {
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeCdnDomainDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainDetailRequest) SetDomainName(v string) *DescribeCdnDomainDetailRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeCdnDomainDetailRequest) SetOwnerId(v int64) *DescribeCdnDomainDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCdnDomainDetailRequest) SetSecurityToken(v string) *DescribeCdnDomainDetailRequest {
	s.SecurityToken = &v
	return s
}

type DescribeCdnDomainDetailResponseBody struct {
	GetDomainDetailModel *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel `json:"GetDomainDetailModel,omitempty" xml:"GetDomainDetailModel,omitempty" type:"Struct"`
	RequestId            *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCdnDomainDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainDetailResponseBody) SetGetDomainDetailModel(v *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) *DescribeCdnDomainDetailResponseBody {
	s.GetDomainDetailModel = v
	return s
}

func (s *DescribeCdnDomainDetailResponseBody) SetRequestId(v string) *DescribeCdnDomainDetailResponseBody {
	s.RequestId = &v
	return s
}

type DescribeCdnDomainDetailResponseBodyGetDomainDetailModel struct {
	CdnType                 *string                                                              `json:"CdnType,omitempty" xml:"CdnType,omitempty"`
	Cname                   *string                                                              `json:"Cname,omitempty" xml:"Cname,omitempty"`
	Description             *string                                                              `json:"Description,omitempty" xml:"Description,omitempty"`
	DomainName              *string                                                              `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainStatus            *string                                                              `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	GmtCreated              *string                                                              `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtModified             *string                                                              `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	HttpsCname              *string                                                              `json:"HttpsCname,omitempty" xml:"HttpsCname,omitempty"`
	Region                  *string                                                              `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceGroupId         *string                                                              `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Scope                   *string                                                              `json:"Scope,omitempty" xml:"Scope,omitempty"`
	ServerCertificateStatus *string                                                              `json:"ServerCertificateStatus,omitempty" xml:"ServerCertificateStatus,omitempty"`
	SourceModels            *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModels `json:"SourceModels,omitempty" xml:"SourceModels,omitempty" type:"Struct"`
	SourcePort              *int32                                                               `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	SourceType              *string                                                              `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Sources                 *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSources      `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Struct"`
}

func (s DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) SetCdnType(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel {
	s.CdnType = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) SetCname(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel {
	s.Cname = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) SetDescription(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel {
	s.Description = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) SetDomainName(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel {
	s.DomainName = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) SetDomainStatus(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel {
	s.DomainStatus = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) SetGmtCreated(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel {
	s.GmtCreated = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) SetGmtModified(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel {
	s.GmtModified = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) SetHttpsCname(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel {
	s.HttpsCname = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) SetRegion(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel {
	s.Region = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) SetResourceGroupId(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) SetScope(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel {
	s.Scope = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) SetServerCertificateStatus(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel {
	s.ServerCertificateStatus = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) SetSourceModels(v *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModels) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel {
	s.SourceModels = v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) SetSourcePort(v int32) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel {
	s.SourcePort = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) SetSourceType(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel {
	s.SourceType = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) SetSources(v *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSources) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel {
	s.Sources = v
	return s
}

type DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModels struct {
	SourceModel []*DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel `json:"SourceModel,omitempty" xml:"SourceModel,omitempty" type:"Repeated"`
}

func (s DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModels) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModels) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModels) SetSourceModel(v []*DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModels {
	s.SourceModel = v
	return s
}

type DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel struct {
	Content  *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Enabled  *string `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel) SetContent(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel {
	s.Content = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel) SetEnabled(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel {
	s.Enabled = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel) SetPort(v int32) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel {
	s.Port = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel) SetPriority(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel {
	s.Priority = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel) SetType(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel {
	s.Type = &v
	return s
}

type DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSources struct {
	Source []*string `json:"Source,omitempty" xml:"Source,omitempty" type:"Repeated"`
}

func (s DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSources) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSources) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSources) SetSource(v []*string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSources {
	s.Source = v
	return s
}

type DescribeCdnDomainDetailResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeCdnDomainDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCdnDomainDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainDetailResponse) SetHeaders(v map[string]*string) *DescribeCdnDomainDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnDomainDetailResponse) SetStatusCode(v int32) *DescribeCdnDomainDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdnDomainDetailResponse) SetBody(v *DescribeCdnDomainDetailResponseBody) *DescribeCdnDomainDetailResponse {
	s.Body = v
	return s
}

type DescribeCdnDomainLogsRequest struct {
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	LogDay        *string `json:"LogDay,omitempty" xml:"LogDay,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber    *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeCdnDomainLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainLogsRequest) SetDomainName(v string) *DescribeCdnDomainLogsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeCdnDomainLogsRequest) SetEndTime(v string) *DescribeCdnDomainLogsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeCdnDomainLogsRequest) SetLogDay(v string) *DescribeCdnDomainLogsRequest {
	s.LogDay = &v
	return s
}

func (s *DescribeCdnDomainLogsRequest) SetOwnerId(v int64) *DescribeCdnDomainLogsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCdnDomainLogsRequest) SetPageNumber(v int64) *DescribeCdnDomainLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCdnDomainLogsRequest) SetPageSize(v int64) *DescribeCdnDomainLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCdnDomainLogsRequest) SetSecurityToken(v string) *DescribeCdnDomainLogsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeCdnDomainLogsRequest) SetStartTime(v string) *DescribeCdnDomainLogsRequest {
	s.StartTime = &v
	return s
}

type DescribeCdnDomainLogsResponseBody struct {
	DomainLogModel *DescribeCdnDomainLogsResponseBodyDomainLogModel `json:"DomainLogModel,omitempty" xml:"DomainLogModel,omitempty" type:"Struct"`
	PageNumber     *int64                                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int64                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount     *int64                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCdnDomainLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainLogsResponseBody) SetDomainLogModel(v *DescribeCdnDomainLogsResponseBodyDomainLogModel) *DescribeCdnDomainLogsResponseBody {
	s.DomainLogModel = v
	return s
}

func (s *DescribeCdnDomainLogsResponseBody) SetPageNumber(v int64) *DescribeCdnDomainLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCdnDomainLogsResponseBody) SetPageSize(v int64) *DescribeCdnDomainLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCdnDomainLogsResponseBody) SetRequestId(v string) *DescribeCdnDomainLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnDomainLogsResponseBody) SetTotalCount(v int64) *DescribeCdnDomainLogsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeCdnDomainLogsResponseBodyDomainLogModel struct {
	DomainLogDetails *DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetails `json:"DomainLogDetails,omitempty" xml:"DomainLogDetails,omitempty" type:"Struct"`
	DomainName       *string                                                          `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeCdnDomainLogsResponseBodyDomainLogModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainLogsResponseBodyDomainLogModel) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogModel) SetDomainLogDetails(v *DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetails) *DescribeCdnDomainLogsResponseBodyDomainLogModel {
	s.DomainLogDetails = v
	return s
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogModel) SetDomainName(v string) *DescribeCdnDomainLogsResponseBodyDomainLogModel {
	s.DomainName = &v
	return s
}

type DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetails struct {
	DomainLogDetail []*DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetailsDomainLogDetail `json:"DomainLogDetail,omitempty" xml:"DomainLogDetail,omitempty" type:"Repeated"`
}

func (s DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetails) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetails) SetDomainLogDetail(v []*DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetailsDomainLogDetail) *DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetails {
	s.DomainLogDetail = v
	return s
}

type DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetailsDomainLogDetail struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	LogName   *string `json:"LogName,omitempty" xml:"LogName,omitempty"`
	LogPath   *string `json:"LogPath,omitempty" xml:"LogPath,omitempty"`
	LogSize   *int64  `json:"LogSize,omitempty" xml:"LogSize,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetailsDomainLogDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetailsDomainLogDetail) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetailsDomainLogDetail) SetEndTime(v string) *DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetailsDomainLogDetail {
	s.EndTime = &v
	return s
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetailsDomainLogDetail) SetLogName(v string) *DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetailsDomainLogDetail {
	s.LogName = &v
	return s
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetailsDomainLogDetail) SetLogPath(v string) *DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetailsDomainLogDetail {
	s.LogPath = &v
	return s
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetailsDomainLogDetail) SetLogSize(v int64) *DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetailsDomainLogDetail {
	s.LogSize = &v
	return s
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetailsDomainLogDetail) SetStartTime(v string) *DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetailsDomainLogDetail {
	s.StartTime = &v
	return s
}

type DescribeCdnDomainLogsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeCdnDomainLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCdnDomainLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainLogsResponse) SetHeaders(v map[string]*string) *DescribeCdnDomainLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnDomainLogsResponse) SetStatusCode(v int32) *DescribeCdnDomainLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdnDomainLogsResponse) SetBody(v *DescribeCdnDomainLogsResponseBody) *DescribeCdnDomainLogsResponse {
	s.Body = v
	return s
}

type DescribeCdnServiceRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeCdnServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnServiceRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnServiceRequest) SetOwnerId(v int64) *DescribeCdnServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCdnServiceRequest) SetSecurityToken(v string) *DescribeCdnServiceRequest {
	s.SecurityToken = &v
	return s
}

type DescribeCdnServiceResponseBody struct {
	ChangingAffectTime *string                                       `json:"ChangingAffectTime,omitempty" xml:"ChangingAffectTime,omitempty"`
	ChangingChargeType *string                                       `json:"ChangingChargeType,omitempty" xml:"ChangingChargeType,omitempty"`
	InstanceId         *string                                       `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InternetChargeType *string                                       `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	OpeningTime        *string                                       `json:"OpeningTime,omitempty" xml:"OpeningTime,omitempty"`
	OperationLocks     *DescribeCdnServiceResponseBodyOperationLocks `json:"OperationLocks,omitempty" xml:"OperationLocks,omitempty" type:"Struct"`
	RequestId          *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCdnServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnServiceResponseBody) SetChangingAffectTime(v string) *DescribeCdnServiceResponseBody {
	s.ChangingAffectTime = &v
	return s
}

func (s *DescribeCdnServiceResponseBody) SetChangingChargeType(v string) *DescribeCdnServiceResponseBody {
	s.ChangingChargeType = &v
	return s
}

func (s *DescribeCdnServiceResponseBody) SetInstanceId(v string) *DescribeCdnServiceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeCdnServiceResponseBody) SetInternetChargeType(v string) *DescribeCdnServiceResponseBody {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeCdnServiceResponseBody) SetOpeningTime(v string) *DescribeCdnServiceResponseBody {
	s.OpeningTime = &v
	return s
}

func (s *DescribeCdnServiceResponseBody) SetOperationLocks(v *DescribeCdnServiceResponseBodyOperationLocks) *DescribeCdnServiceResponseBody {
	s.OperationLocks = v
	return s
}

func (s *DescribeCdnServiceResponseBody) SetRequestId(v string) *DescribeCdnServiceResponseBody {
	s.RequestId = &v
	return s
}

type DescribeCdnServiceResponseBodyOperationLocks struct {
	LockReason []*DescribeCdnServiceResponseBodyOperationLocksLockReason `json:"LockReason,omitempty" xml:"LockReason,omitempty" type:"Repeated"`
}

func (s DescribeCdnServiceResponseBodyOperationLocks) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnServiceResponseBodyOperationLocks) GoString() string {
	return s.String()
}

func (s *DescribeCdnServiceResponseBodyOperationLocks) SetLockReason(v []*DescribeCdnServiceResponseBodyOperationLocksLockReason) *DescribeCdnServiceResponseBodyOperationLocks {
	s.LockReason = v
	return s
}

type DescribeCdnServiceResponseBodyOperationLocksLockReason struct {
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
}

func (s DescribeCdnServiceResponseBodyOperationLocksLockReason) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnServiceResponseBodyOperationLocksLockReason) GoString() string {
	return s.String()
}

func (s *DescribeCdnServiceResponseBodyOperationLocksLockReason) SetLockReason(v string) *DescribeCdnServiceResponseBodyOperationLocksLockReason {
	s.LockReason = &v
	return s
}

type DescribeCdnServiceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeCdnServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCdnServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnServiceResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnServiceResponse) SetHeaders(v map[string]*string) *DescribeCdnServiceResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnServiceResponse) SetStatusCode(v int32) *DescribeCdnServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdnServiceResponse) SetBody(v *DescribeCdnServiceResponseBody) *DescribeCdnServiceResponse {
	s.Body = v
	return s
}

type DescribeDomainBpsDataRequest struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainType     *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval       *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TimeMerge      *string `json:"TimeMerge,omitempty" xml:"TimeMerge,omitempty"`
}

func (s DescribeDomainBpsDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataRequest) SetDomainName(v string) *DescribeDomainBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainBpsDataRequest) SetDomainType(v string) *DescribeDomainBpsDataRequest {
	s.DomainType = &v
	return s
}

func (s *DescribeDomainBpsDataRequest) SetEndTime(v string) *DescribeDomainBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainBpsDataRequest) SetInterval(v string) *DescribeDomainBpsDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainBpsDataRequest) SetIspNameEn(v string) *DescribeDomainBpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDomainBpsDataRequest) SetLocationNameEn(v string) *DescribeDomainBpsDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDomainBpsDataRequest) SetStartTime(v string) *DescribeDomainBpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainBpsDataRequest) SetTimeMerge(v string) *DescribeDomainBpsDataRequest {
	s.TimeMerge = &v
	return s
}

type DescribeDomainBpsDataResponseBody struct {
	BpsDataPerInterval *DescribeDomainBpsDataResponseBodyBpsDataPerInterval `json:"BpsDataPerInterval,omitempty" xml:"BpsDataPerInterval,omitempty" type:"Struct"`
	DataInterval       *string                                              `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName         *string                                              `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime            *string                                              `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IspName            *string                                              `json:"IspName,omitempty" xml:"IspName,omitempty"`
	IspNameEn          *string                                              `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationName       *string                                              `json:"LocationName,omitempty" xml:"LocationName,omitempty"`
	LocationNameEn     *string                                              `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	RequestId          *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime          *string                                              `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainBpsDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataResponseBody) SetBpsDataPerInterval(v *DescribeDomainBpsDataResponseBodyBpsDataPerInterval) *DescribeDomainBpsDataResponseBody {
	s.BpsDataPerInterval = v
	return s
}

func (s *DescribeDomainBpsDataResponseBody) SetDataInterval(v string) *DescribeDomainBpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBody) SetDomainName(v string) *DescribeDomainBpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBody) SetEndTime(v string) *DescribeDomainBpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBody) SetIspName(v string) *DescribeDomainBpsDataResponseBody {
	s.IspName = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBody) SetIspNameEn(v string) *DescribeDomainBpsDataResponseBody {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBody) SetLocationName(v string) *DescribeDomainBpsDataResponseBody {
	s.LocationName = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBody) SetLocationNameEn(v string) *DescribeDomainBpsDataResponseBody {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBody) SetRequestId(v string) *DescribeDomainBpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBody) SetStartTime(v string) *DescribeDomainBpsDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeDomainBpsDataResponseBodyBpsDataPerInterval struct {
	DataModule []*DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainBpsDataResponseBodyBpsDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainBpsDataResponseBodyBpsDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerInterval) SetDataModule(v []*DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) *DescribeDomainBpsDataResponseBodyBpsDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule struct {
	DomesticL2Value      *string `json:"DomesticL2Value,omitempty" xml:"DomesticL2Value,omitempty"`
	DomesticValue        *string `json:"DomesticValue,omitempty" xml:"DomesticValue,omitempty"`
	DynamicDomesticValue *string `json:"DynamicDomesticValue,omitempty" xml:"DynamicDomesticValue,omitempty"`
	DynamicOverseasValue *string `json:"DynamicOverseasValue,omitempty" xml:"DynamicOverseasValue,omitempty"`
	DynamicValue         *string `json:"DynamicValue,omitempty" xml:"DynamicValue,omitempty"`
	L2Value              *string `json:"L2Value,omitempty" xml:"L2Value,omitempty"`
	OverseasL2Value      *string `json:"OverseasL2Value,omitempty" xml:"OverseasL2Value,omitempty"`
	OverseasValue        *string `json:"OverseasValue,omitempty" xml:"OverseasValue,omitempty"`
	StaticDomesticValue  *string `json:"StaticDomesticValue,omitempty" xml:"StaticDomesticValue,omitempty"`
	StaticOverseasValue  *string `json:"StaticOverseasValue,omitempty" xml:"StaticOverseasValue,omitempty"`
	StaticValue          *string `json:"StaticValue,omitempty" xml:"StaticValue,omitempty"`
	TimeStamp            *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value                *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetDomesticL2Value(v string) *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.DomesticL2Value = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetDomesticValue(v string) *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.DomesticValue = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetDynamicDomesticValue(v string) *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.DynamicDomesticValue = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetDynamicOverseasValue(v string) *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.DynamicOverseasValue = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetDynamicValue(v string) *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.DynamicValue = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetL2Value(v string) *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.L2Value = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetOverseasL2Value(v string) *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.OverseasL2Value = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetOverseasValue(v string) *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.OverseasValue = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetStaticDomesticValue(v string) *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.StaticDomesticValue = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetStaticOverseasValue(v string) *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.StaticOverseasValue = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetStaticValue(v string) *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.StaticValue = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetValue(v string) *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.Value = &v
	return s
}

type DescribeDomainBpsDataResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDomainBpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainBpsDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataResponse) SetHeaders(v map[string]*string) *DescribeDomainBpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainBpsDataResponse) SetStatusCode(v int32) *DescribeDomainBpsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainBpsDataResponse) SetBody(v *DescribeDomainBpsDataResponseBody) *DescribeDomainBpsDataResponse {
	s.Body = v
	return s
}

type DescribeDomainBpsDataByTimeStampRequest struct {
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	IspNames      *string `json:"IspNames,omitempty" xml:"IspNames,omitempty"`
	LocationNames *string `json:"LocationNames,omitempty" xml:"LocationNames,omitempty"`
	TimePoint     *string `json:"TimePoint,omitempty" xml:"TimePoint,omitempty"`
}

func (s DescribeDomainBpsDataByTimeStampRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainBpsDataByTimeStampRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataByTimeStampRequest) SetDomainName(v string) *DescribeDomainBpsDataByTimeStampRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainBpsDataByTimeStampRequest) SetIspNames(v string) *DescribeDomainBpsDataByTimeStampRequest {
	s.IspNames = &v
	return s
}

func (s *DescribeDomainBpsDataByTimeStampRequest) SetLocationNames(v string) *DescribeDomainBpsDataByTimeStampRequest {
	s.LocationNames = &v
	return s
}

func (s *DescribeDomainBpsDataByTimeStampRequest) SetTimePoint(v string) *DescribeDomainBpsDataByTimeStampRequest {
	s.TimePoint = &v
	return s
}

type DescribeDomainBpsDataByTimeStampResponseBody struct {
	BpsDataList *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataList `json:"BpsDataList,omitempty" xml:"BpsDataList,omitempty" type:"Struct"`
	DomainName  *string                                                  `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	RequestId   *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TimeStamp   *string                                                  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDomainBpsDataByTimeStampResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainBpsDataByTimeStampResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataByTimeStampResponseBody) SetBpsDataList(v *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataList) *DescribeDomainBpsDataByTimeStampResponseBody {
	s.BpsDataList = v
	return s
}

func (s *DescribeDomainBpsDataByTimeStampResponseBody) SetDomainName(v string) *DescribeDomainBpsDataByTimeStampResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainBpsDataByTimeStampResponseBody) SetRequestId(v string) *DescribeDomainBpsDataByTimeStampResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainBpsDataByTimeStampResponseBody) SetTimeStamp(v string) *DescribeDomainBpsDataByTimeStampResponseBody {
	s.TimeStamp = &v
	return s
}

type DescribeDomainBpsDataByTimeStampResponseBodyBpsDataList struct {
	BpsDataModel []*DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel `json:"BpsDataModel,omitempty" xml:"BpsDataModel,omitempty" type:"Repeated"`
}

func (s DescribeDomainBpsDataByTimeStampResponseBodyBpsDataList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainBpsDataByTimeStampResponseBodyBpsDataList) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataList) SetBpsDataModel(v []*DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel) *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataList {
	s.BpsDataModel = v
	return s
}

type DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel struct {
	Bps          *int64  `json:"Bps,omitempty" xml:"Bps,omitempty"`
	IspName      *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	LocationName *string `json:"LocationName,omitempty" xml:"LocationName,omitempty"`
	TimeStamp    *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel) SetBps(v int64) *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel {
	s.Bps = &v
	return s
}

func (s *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel) SetIspName(v string) *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel {
	s.IspName = &v
	return s
}

func (s *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel) SetLocationName(v string) *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel {
	s.LocationName = &v
	return s
}

func (s *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel) SetTimeStamp(v string) *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel {
	s.TimeStamp = &v
	return s
}

type DescribeDomainBpsDataByTimeStampResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDomainBpsDataByTimeStampResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainBpsDataByTimeStampResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainBpsDataByTimeStampResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataByTimeStampResponse) SetHeaders(v map[string]*string) *DescribeDomainBpsDataByTimeStampResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainBpsDataByTimeStampResponse) SetStatusCode(v int32) *DescribeDomainBpsDataByTimeStampResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainBpsDataByTimeStampResponse) SetBody(v *DescribeDomainBpsDataByTimeStampResponseBody) *DescribeDomainBpsDataByTimeStampResponse {
	s.Body = v
	return s
}

type DescribeDomainHitRateDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval   *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainHitRateDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainHitRateDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainHitRateDataRequest) SetDomainName(v string) *DescribeDomainHitRateDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainHitRateDataRequest) SetEndTime(v string) *DescribeDomainHitRateDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainHitRateDataRequest) SetInterval(v string) *DescribeDomainHitRateDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainHitRateDataRequest) SetStartTime(v string) *DescribeDomainHitRateDataRequest {
	s.StartTime = &v
	return s
}

type DescribeDomainHitRateDataResponseBody struct {
	DataInterval    *string                                               `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName      *string                                               `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime         *string                                               `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	HitRateInterval *DescribeDomainHitRateDataResponseBodyHitRateInterval `json:"HitRateInterval,omitempty" xml:"HitRateInterval,omitempty" type:"Struct"`
	RequestId       *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime       *string                                               `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainHitRateDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainHitRateDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainHitRateDataResponseBody) SetDataInterval(v string) *DescribeDomainHitRateDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainHitRateDataResponseBody) SetDomainName(v string) *DescribeDomainHitRateDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainHitRateDataResponseBody) SetEndTime(v string) *DescribeDomainHitRateDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainHitRateDataResponseBody) SetHitRateInterval(v *DescribeDomainHitRateDataResponseBodyHitRateInterval) *DescribeDomainHitRateDataResponseBody {
	s.HitRateInterval = v
	return s
}

func (s *DescribeDomainHitRateDataResponseBody) SetRequestId(v string) *DescribeDomainHitRateDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainHitRateDataResponseBody) SetStartTime(v string) *DescribeDomainHitRateDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeDomainHitRateDataResponseBodyHitRateInterval struct {
	DataModule []*DescribeDomainHitRateDataResponseBodyHitRateIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainHitRateDataResponseBodyHitRateInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainHitRateDataResponseBodyHitRateInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainHitRateDataResponseBodyHitRateInterval) SetDataModule(v []*DescribeDomainHitRateDataResponseBodyHitRateIntervalDataModule) *DescribeDomainHitRateDataResponseBodyHitRateInterval {
	s.DataModule = v
	return s
}

type DescribeDomainHitRateDataResponseBodyHitRateIntervalDataModule struct {
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDomainHitRateDataResponseBodyHitRateIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainHitRateDataResponseBodyHitRateIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainHitRateDataResponseBodyHitRateIntervalDataModule) SetTimeStamp(v string) *DescribeDomainHitRateDataResponseBodyHitRateIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainHitRateDataResponseBodyHitRateIntervalDataModule) SetValue(v string) *DescribeDomainHitRateDataResponseBodyHitRateIntervalDataModule {
	s.Value = &v
	return s
}

type DescribeDomainHitRateDataResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDomainHitRateDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainHitRateDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainHitRateDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainHitRateDataResponse) SetHeaders(v map[string]*string) *DescribeDomainHitRateDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainHitRateDataResponse) SetStatusCode(v int32) *DescribeDomainHitRateDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainHitRateDataResponse) SetBody(v *DescribeDomainHitRateDataResponseBody) *DescribeDomainHitRateDataResponse {
	s.Body = v
	return s
}

type DescribeDomainHttpCodeDataRequest struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval       *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TimeMerge      *string `json:"TimeMerge,omitempty" xml:"TimeMerge,omitempty"`
}

func (s DescribeDomainHttpCodeDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainHttpCodeDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainHttpCodeDataRequest) SetDomainName(v string) *DescribeDomainHttpCodeDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainHttpCodeDataRequest) SetEndTime(v string) *DescribeDomainHttpCodeDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainHttpCodeDataRequest) SetInterval(v string) *DescribeDomainHttpCodeDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainHttpCodeDataRequest) SetIspNameEn(v string) *DescribeDomainHttpCodeDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDomainHttpCodeDataRequest) SetLocationNameEn(v string) *DescribeDomainHttpCodeDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDomainHttpCodeDataRequest) SetStartTime(v string) *DescribeDomainHttpCodeDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainHttpCodeDataRequest) SetTimeMerge(v string) *DescribeDomainHttpCodeDataRequest {
	s.TimeMerge = &v
	return s
}

type DescribeDomainHttpCodeDataResponseBody struct {
	DataInterval *string                                             `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName   *string                                             `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime      *string                                             `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	HttpCodeData *DescribeDomainHttpCodeDataResponseBodyHttpCodeData `json:"HttpCodeData,omitempty" xml:"HttpCodeData,omitempty" type:"Struct"`
	RequestId    *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime    *string                                             `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainHttpCodeDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainHttpCodeDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainHttpCodeDataResponseBody) SetDataInterval(v string) *DescribeDomainHttpCodeDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainHttpCodeDataResponseBody) SetDomainName(v string) *DescribeDomainHttpCodeDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainHttpCodeDataResponseBody) SetEndTime(v string) *DescribeDomainHttpCodeDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainHttpCodeDataResponseBody) SetHttpCodeData(v *DescribeDomainHttpCodeDataResponseBodyHttpCodeData) *DescribeDomainHttpCodeDataResponseBody {
	s.HttpCodeData = v
	return s
}

func (s *DescribeDomainHttpCodeDataResponseBody) SetRequestId(v string) *DescribeDomainHttpCodeDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainHttpCodeDataResponseBody) SetStartTime(v string) *DescribeDomainHttpCodeDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeDomainHttpCodeDataResponseBodyHttpCodeData struct {
	UsageData []*DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s DescribeDomainHttpCodeDataResponseBodyHttpCodeData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainHttpCodeDataResponseBodyHttpCodeData) GoString() string {
	return s.String()
}

func (s *DescribeDomainHttpCodeDataResponseBodyHttpCodeData) SetUsageData(v []*DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageData) *DescribeDomainHttpCodeDataResponseBodyHttpCodeData {
	s.UsageData = v
	return s
}

type DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageData struct {
	TimeStamp *string                                                           `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value     *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageData) GoString() string {
	return s.String()
}

func (s *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageData) SetTimeStamp(v string) *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageData) SetValue(v *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValue) *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageData {
	s.Value = v
	return s
}

type DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValue struct {
	CodeProportionData []*DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData `json:"CodeProportionData,omitempty" xml:"CodeProportionData,omitempty" type:"Repeated"`
}

func (s DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValue) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValue) GoString() string {
	return s.String()
}

func (s *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValue) SetCodeProportionData(v []*DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValue {
	s.CodeProportionData = v
	return s
}

type DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Count      *string `json:"Count,omitempty" xml:"Count,omitempty"`
	Proportion *string `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
}

func (s DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) GoString() string {
	return s.String()
}

func (s *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) SetCode(v string) *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData {
	s.Code = &v
	return s
}

func (s *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) SetCount(v string) *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData {
	s.Count = &v
	return s
}

func (s *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) SetProportion(v string) *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData {
	s.Proportion = &v
	return s
}

type DescribeDomainHttpCodeDataResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDomainHttpCodeDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainHttpCodeDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainHttpCodeDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainHttpCodeDataResponse) SetHeaders(v map[string]*string) *DescribeDomainHttpCodeDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainHttpCodeDataResponse) SetStatusCode(v int32) *DescribeDomainHttpCodeDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainHttpCodeDataResponse) SetBody(v *DescribeDomainHttpCodeDataResponseBody) *DescribeDomainHttpCodeDataResponse {
	s.Body = v
	return s
}

type DescribeDomainQpsDataRequest struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainType     *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval       *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TimeMerge      *string `json:"TimeMerge,omitempty" xml:"TimeMerge,omitempty"`
}

func (s DescribeDomainQpsDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainQpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsDataRequest) SetDomainName(v string) *DescribeDomainQpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainQpsDataRequest) SetDomainType(v string) *DescribeDomainQpsDataRequest {
	s.DomainType = &v
	return s
}

func (s *DescribeDomainQpsDataRequest) SetEndTime(v string) *DescribeDomainQpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainQpsDataRequest) SetInterval(v string) *DescribeDomainQpsDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainQpsDataRequest) SetIspNameEn(v string) *DescribeDomainQpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDomainQpsDataRequest) SetLocationNameEn(v string) *DescribeDomainQpsDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDomainQpsDataRequest) SetStartTime(v string) *DescribeDomainQpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainQpsDataRequest) SetTimeMerge(v string) *DescribeDomainQpsDataRequest {
	s.TimeMerge = &v
	return s
}

type DescribeDomainQpsDataResponseBody struct {
	DataInterval    *string                                           `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName      *string                                           `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime         *string                                           `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	QpsDataInterval *DescribeDomainQpsDataResponseBodyQpsDataInterval `json:"QpsDataInterval,omitempty" xml:"QpsDataInterval,omitempty" type:"Struct"`
	RequestId       *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime       *string                                           `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainQpsDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainQpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsDataResponseBody) SetDataInterval(v string) *DescribeDomainQpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBody) SetDomainName(v string) *DescribeDomainQpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBody) SetEndTime(v string) *DescribeDomainQpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBody) SetQpsDataInterval(v *DescribeDomainQpsDataResponseBodyQpsDataInterval) *DescribeDomainQpsDataResponseBody {
	s.QpsDataInterval = v
	return s
}

func (s *DescribeDomainQpsDataResponseBody) SetRequestId(v string) *DescribeDomainQpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBody) SetStartTime(v string) *DescribeDomainQpsDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeDomainQpsDataResponseBodyQpsDataInterval struct {
	DataModule []*DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainQpsDataResponseBodyQpsDataInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainQpsDataResponseBodyQpsDataInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataInterval) SetDataModule(v []*DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) *DescribeDomainQpsDataResponseBodyQpsDataInterval {
	s.DataModule = v
	return s
}

type DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule struct {
	AccDomesticValue     *string `json:"AccDomesticValue,omitempty" xml:"AccDomesticValue,omitempty"`
	AccOverseasValue     *string `json:"AccOverseasValue,omitempty" xml:"AccOverseasValue,omitempty"`
	AccValue             *string `json:"AccValue,omitempty" xml:"AccValue,omitempty"`
	DomesticValue        *string `json:"DomesticValue,omitempty" xml:"DomesticValue,omitempty"`
	DynamicDomesticValue *string `json:"DynamicDomesticValue,omitempty" xml:"DynamicDomesticValue,omitempty"`
	DynamicOverseasValue *string `json:"DynamicOverseasValue,omitempty" xml:"DynamicOverseasValue,omitempty"`
	DynamicValue         *string `json:"DynamicValue,omitempty" xml:"DynamicValue,omitempty"`
	OverseasValue        *string `json:"OverseasValue,omitempty" xml:"OverseasValue,omitempty"`
	StaticDomesticValue  *string `json:"StaticDomesticValue,omitempty" xml:"StaticDomesticValue,omitempty"`
	StaticOverseasValue  *string `json:"StaticOverseasValue,omitempty" xml:"StaticOverseasValue,omitempty"`
	StaticValue          *string `json:"StaticValue,omitempty" xml:"StaticValue,omitempty"`
	TimeStamp            *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value                *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetAccDomesticValue(v string) *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.AccDomesticValue = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetAccOverseasValue(v string) *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.AccOverseasValue = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetAccValue(v string) *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.AccValue = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetDomesticValue(v string) *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.DomesticValue = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetDynamicDomesticValue(v string) *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.DynamicDomesticValue = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetDynamicOverseasValue(v string) *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.DynamicOverseasValue = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetDynamicValue(v string) *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.DynamicValue = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetOverseasValue(v string) *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.OverseasValue = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetStaticDomesticValue(v string) *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.StaticDomesticValue = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetStaticOverseasValue(v string) *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.StaticOverseasValue = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetStaticValue(v string) *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.StaticValue = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetTimeStamp(v string) *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetValue(v string) *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.Value = &v
	return s
}

type DescribeDomainQpsDataResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDomainQpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainQpsDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainQpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsDataResponse) SetHeaders(v map[string]*string) *DescribeDomainQpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainQpsDataResponse) SetStatusCode(v int32) *DescribeDomainQpsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainQpsDataResponse) SetBody(v *DescribeDomainQpsDataResponseBody) *DescribeDomainQpsDataResponse {
	s.Body = v
	return s
}

type DescribeDomainReqHitRateDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval   *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainReqHitRateDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainReqHitRateDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainReqHitRateDataRequest) SetDomainName(v string) *DescribeDomainReqHitRateDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainReqHitRateDataRequest) SetEndTime(v string) *DescribeDomainReqHitRateDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainReqHitRateDataRequest) SetInterval(v string) *DescribeDomainReqHitRateDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainReqHitRateDataRequest) SetStartTime(v string) *DescribeDomainReqHitRateDataRequest {
	s.StartTime = &v
	return s
}

type DescribeDomainReqHitRateDataResponseBody struct {
	DataInterval       *string                                                     `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName         *string                                                     `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime            *string                                                     `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ReqHitRateInterval *DescribeDomainReqHitRateDataResponseBodyReqHitRateInterval `json:"ReqHitRateInterval,omitempty" xml:"ReqHitRateInterval,omitempty" type:"Struct"`
	RequestId          *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime          *string                                                     `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainReqHitRateDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainReqHitRateDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainReqHitRateDataResponseBody) SetDataInterval(v string) *DescribeDomainReqHitRateDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainReqHitRateDataResponseBody) SetDomainName(v string) *DescribeDomainReqHitRateDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainReqHitRateDataResponseBody) SetEndTime(v string) *DescribeDomainReqHitRateDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainReqHitRateDataResponseBody) SetReqHitRateInterval(v *DescribeDomainReqHitRateDataResponseBodyReqHitRateInterval) *DescribeDomainReqHitRateDataResponseBody {
	s.ReqHitRateInterval = v
	return s
}

func (s *DescribeDomainReqHitRateDataResponseBody) SetRequestId(v string) *DescribeDomainReqHitRateDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainReqHitRateDataResponseBody) SetStartTime(v string) *DescribeDomainReqHitRateDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeDomainReqHitRateDataResponseBodyReqHitRateInterval struct {
	DataModule []*DescribeDomainReqHitRateDataResponseBodyReqHitRateIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainReqHitRateDataResponseBodyReqHitRateInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainReqHitRateDataResponseBodyReqHitRateInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainReqHitRateDataResponseBodyReqHitRateInterval) SetDataModule(v []*DescribeDomainReqHitRateDataResponseBodyReqHitRateIntervalDataModule) *DescribeDomainReqHitRateDataResponseBodyReqHitRateInterval {
	s.DataModule = v
	return s
}

type DescribeDomainReqHitRateDataResponseBodyReqHitRateIntervalDataModule struct {
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDomainReqHitRateDataResponseBodyReqHitRateIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainReqHitRateDataResponseBodyReqHitRateIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainReqHitRateDataResponseBodyReqHitRateIntervalDataModule) SetTimeStamp(v string) *DescribeDomainReqHitRateDataResponseBodyReqHitRateIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainReqHitRateDataResponseBodyReqHitRateIntervalDataModule) SetValue(v string) *DescribeDomainReqHitRateDataResponseBodyReqHitRateIntervalDataModule {
	s.Value = &v
	return s
}

type DescribeDomainReqHitRateDataResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDomainReqHitRateDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainReqHitRateDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainReqHitRateDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainReqHitRateDataResponse) SetHeaders(v map[string]*string) *DescribeDomainReqHitRateDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainReqHitRateDataResponse) SetStatusCode(v int32) *DescribeDomainReqHitRateDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainReqHitRateDataResponse) SetBody(v *DescribeDomainReqHitRateDataResponseBody) *DescribeDomainReqHitRateDataResponse {
	s.Body = v
	return s
}

type DescribeDomainSrcBpsDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FixTimeGap *string `json:"FixTimeGap,omitempty" xml:"FixTimeGap,omitempty"`
	Interval   *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TimeMerge  *string `json:"TimeMerge,omitempty" xml:"TimeMerge,omitempty"`
}

func (s DescribeDomainSrcBpsDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcBpsDataRequest) SetDomainName(v string) *DescribeDomainSrcBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainSrcBpsDataRequest) SetEndTime(v string) *DescribeDomainSrcBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainSrcBpsDataRequest) SetFixTimeGap(v string) *DescribeDomainSrcBpsDataRequest {
	s.FixTimeGap = &v
	return s
}

func (s *DescribeDomainSrcBpsDataRequest) SetInterval(v string) *DescribeDomainSrcBpsDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainSrcBpsDataRequest) SetStartTime(v string) *DescribeDomainSrcBpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainSrcBpsDataRequest) SetTimeMerge(v string) *DescribeDomainSrcBpsDataRequest {
	s.TimeMerge = &v
	return s
}

type DescribeDomainSrcBpsDataResponseBody struct {
	DataInterval          *string                                                    `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName            *string                                                    `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime               *string                                                    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId             *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SrcBpsDataPerInterval *DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerInterval `json:"SrcBpsDataPerInterval,omitempty" xml:"SrcBpsDataPerInterval,omitempty" type:"Struct"`
	StartTime             *string                                                    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainSrcBpsDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcBpsDataResponseBody) SetDataInterval(v string) *DescribeDomainSrcBpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainSrcBpsDataResponseBody) SetDomainName(v string) *DescribeDomainSrcBpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainSrcBpsDataResponseBody) SetEndTime(v string) *DescribeDomainSrcBpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainSrcBpsDataResponseBody) SetRequestId(v string) *DescribeDomainSrcBpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainSrcBpsDataResponseBody) SetSrcBpsDataPerInterval(v *DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerInterval) *DescribeDomainSrcBpsDataResponseBody {
	s.SrcBpsDataPerInterval = v
	return s
}

func (s *DescribeDomainSrcBpsDataResponseBody) SetStartTime(v string) *DescribeDomainSrcBpsDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerInterval struct {
	DataModule []*DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerInterval) SetDataModule(v []*DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule) *DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule struct {
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule) SetValue(v string) *DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule {
	s.Value = &v
	return s
}

type DescribeDomainSrcBpsDataResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDomainSrcBpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainSrcBpsDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcBpsDataResponse) SetHeaders(v map[string]*string) *DescribeDomainSrcBpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainSrcBpsDataResponse) SetStatusCode(v int32) *DescribeDomainSrcBpsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainSrcBpsDataResponse) SetBody(v *DescribeDomainSrcBpsDataResponseBody) *DescribeDomainSrcBpsDataResponse {
	s.Body = v
	return s
}

type DescribeDomainSrcFlowDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FixTimeGap *string `json:"FixTimeGap,omitempty" xml:"FixTimeGap,omitempty"`
	Interval   *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TimeMerge  *string `json:"TimeMerge,omitempty" xml:"TimeMerge,omitempty"`
}

func (s DescribeDomainSrcFlowDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcFlowDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcFlowDataRequest) SetDomainName(v string) *DescribeDomainSrcFlowDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainSrcFlowDataRequest) SetEndTime(v string) *DescribeDomainSrcFlowDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainSrcFlowDataRequest) SetFixTimeGap(v string) *DescribeDomainSrcFlowDataRequest {
	s.FixTimeGap = &v
	return s
}

func (s *DescribeDomainSrcFlowDataRequest) SetInterval(v string) *DescribeDomainSrcFlowDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainSrcFlowDataRequest) SetStartTime(v string) *DescribeDomainSrcFlowDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainSrcFlowDataRequest) SetTimeMerge(v string) *DescribeDomainSrcFlowDataRequest {
	s.TimeMerge = &v
	return s
}

type DescribeDomainSrcFlowDataResponseBody struct {
	DataInterval           *string                                                      `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName             *string                                                      `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime                *string                                                      `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId              *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SrcFlowDataPerInterval *DescribeDomainSrcFlowDataResponseBodySrcFlowDataPerInterval `json:"SrcFlowDataPerInterval,omitempty" xml:"SrcFlowDataPerInterval,omitempty" type:"Struct"`
	StartTime              *string                                                      `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainSrcFlowDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcFlowDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcFlowDataResponseBody) SetDataInterval(v string) *DescribeDomainSrcFlowDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainSrcFlowDataResponseBody) SetDomainName(v string) *DescribeDomainSrcFlowDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainSrcFlowDataResponseBody) SetEndTime(v string) *DescribeDomainSrcFlowDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainSrcFlowDataResponseBody) SetRequestId(v string) *DescribeDomainSrcFlowDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainSrcFlowDataResponseBody) SetSrcFlowDataPerInterval(v *DescribeDomainSrcFlowDataResponseBodySrcFlowDataPerInterval) *DescribeDomainSrcFlowDataResponseBody {
	s.SrcFlowDataPerInterval = v
	return s
}

func (s *DescribeDomainSrcFlowDataResponseBody) SetStartTime(v string) *DescribeDomainSrcFlowDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeDomainSrcFlowDataResponseBodySrcFlowDataPerInterval struct {
	DataModule []*DescribeDomainSrcFlowDataResponseBodySrcFlowDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainSrcFlowDataResponseBodySrcFlowDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcFlowDataResponseBodySrcFlowDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcFlowDataResponseBodySrcFlowDataPerInterval) SetDataModule(v []*DescribeDomainSrcFlowDataResponseBodySrcFlowDataPerIntervalDataModule) *DescribeDomainSrcFlowDataResponseBodySrcFlowDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeDomainSrcFlowDataResponseBodySrcFlowDataPerIntervalDataModule struct {
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDomainSrcFlowDataResponseBodySrcFlowDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcFlowDataResponseBodySrcFlowDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcFlowDataResponseBodySrcFlowDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDomainSrcFlowDataResponseBodySrcFlowDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainSrcFlowDataResponseBodySrcFlowDataPerIntervalDataModule) SetValue(v string) *DescribeDomainSrcFlowDataResponseBodySrcFlowDataPerIntervalDataModule {
	s.Value = &v
	return s
}

type DescribeDomainSrcFlowDataResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDomainSrcFlowDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainSrcFlowDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcFlowDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcFlowDataResponse) SetHeaders(v map[string]*string) *DescribeDomainSrcFlowDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainSrcFlowDataResponse) SetStatusCode(v int32) *DescribeDomainSrcFlowDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainSrcFlowDataResponse) SetBody(v *DescribeDomainSrcFlowDataResponseBody) *DescribeDomainSrcFlowDataResponse {
	s.Body = v
	return s
}

type DescribeDomainUvDataRequest struct {
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainUvDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainUvDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainUvDataRequest) SetDomainName(v string) *DescribeDomainUvDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainUvDataRequest) SetEndTime(v string) *DescribeDomainUvDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainUvDataRequest) SetOwnerId(v int64) *DescribeDomainUvDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainUvDataRequest) SetSecurityToken(v string) *DescribeDomainUvDataRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDomainUvDataRequest) SetStartTime(v string) *DescribeDomainUvDataRequest {
	s.StartTime = &v
	return s
}

type DescribeDomainUvDataResponseBody struct {
	DataInterval   *string                                         `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName     *string                                         `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime        *string                                         `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime      *string                                         `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	UvDataInterval *DescribeDomainUvDataResponseBodyUvDataInterval `json:"UvDataInterval,omitempty" xml:"UvDataInterval,omitempty" type:"Struct"`
}

func (s DescribeDomainUvDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainUvDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainUvDataResponseBody) SetDataInterval(v string) *DescribeDomainUvDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainUvDataResponseBody) SetDomainName(v string) *DescribeDomainUvDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainUvDataResponseBody) SetEndTime(v string) *DescribeDomainUvDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainUvDataResponseBody) SetRequestId(v string) *DescribeDomainUvDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainUvDataResponseBody) SetStartTime(v string) *DescribeDomainUvDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainUvDataResponseBody) SetUvDataInterval(v *DescribeDomainUvDataResponseBodyUvDataInterval) *DescribeDomainUvDataResponseBody {
	s.UvDataInterval = v
	return s
}

type DescribeDomainUvDataResponseBodyUvDataInterval struct {
	UsageData []*DescribeDomainUvDataResponseBodyUvDataIntervalUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s DescribeDomainUvDataResponseBodyUvDataInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainUvDataResponseBodyUvDataInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainUvDataResponseBodyUvDataInterval) SetUsageData(v []*DescribeDomainUvDataResponseBodyUvDataIntervalUsageData) *DescribeDomainUvDataResponseBodyUvDataInterval {
	s.UsageData = v
	return s
}

type DescribeDomainUvDataResponseBodyUvDataIntervalUsageData struct {
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDomainUvDataResponseBodyUvDataIntervalUsageData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainUvDataResponseBodyUvDataIntervalUsageData) GoString() string {
	return s.String()
}

func (s *DescribeDomainUvDataResponseBodyUvDataIntervalUsageData) SetTimeStamp(v string) *DescribeDomainUvDataResponseBodyUvDataIntervalUsageData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainUvDataResponseBodyUvDataIntervalUsageData) SetValue(v string) *DescribeDomainUvDataResponseBodyUvDataIntervalUsageData {
	s.Value = &v
	return s
}

type DescribeDomainUvDataResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDomainUvDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainUvDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainUvDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainUvDataResponse) SetHeaders(v map[string]*string) *DescribeDomainUvDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainUvDataResponse) SetStatusCode(v int32) *DescribeDomainUvDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainUvDataResponse) SetBody(v *DescribeDomainUvDataResponseBody) *DescribeDomainUvDataResponse {
	s.Body = v
	return s
}

type DescribeUserDomainsRequest struct {
	CdnType          *string `json:"CdnType,omitempty" xml:"CdnType,omitempty"`
	CheckDomainShow  *bool   `json:"CheckDomainShow,omitempty" xml:"CheckDomainShow,omitempty"`
	DomainName       *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainSearchType *string `json:"DomainSearchType,omitempty" xml:"DomainSearchType,omitempty"`
	DomainStatus     *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber       *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceGroupId  *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityToken    *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Sources          *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
}

func (s DescribeUserDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserDomainsRequest) SetCdnType(v string) *DescribeUserDomainsRequest {
	s.CdnType = &v
	return s
}

func (s *DescribeUserDomainsRequest) SetCheckDomainShow(v bool) *DescribeUserDomainsRequest {
	s.CheckDomainShow = &v
	return s
}

func (s *DescribeUserDomainsRequest) SetDomainName(v string) *DescribeUserDomainsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeUserDomainsRequest) SetDomainSearchType(v string) *DescribeUserDomainsRequest {
	s.DomainSearchType = &v
	return s
}

func (s *DescribeUserDomainsRequest) SetDomainStatus(v string) *DescribeUserDomainsRequest {
	s.DomainStatus = &v
	return s
}

func (s *DescribeUserDomainsRequest) SetOwnerId(v int64) *DescribeUserDomainsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUserDomainsRequest) SetPageNumber(v int32) *DescribeUserDomainsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeUserDomainsRequest) SetPageSize(v int32) *DescribeUserDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeUserDomainsRequest) SetResourceGroupId(v string) *DescribeUserDomainsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeUserDomainsRequest) SetSecurityToken(v string) *DescribeUserDomainsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeUserDomainsRequest) SetSources(v string) *DescribeUserDomainsRequest {
	s.Sources = &v
	return s
}

type DescribeUserDomainsResponseBody struct {
	Domains    *DescribeUserDomainsResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Struct"`
	PageNumber *int64                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeUserDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserDomainsResponseBody) SetDomains(v *DescribeUserDomainsResponseBodyDomains) *DescribeUserDomainsResponseBody {
	s.Domains = v
	return s
}

func (s *DescribeUserDomainsResponseBody) SetPageNumber(v int64) *DescribeUserDomainsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeUserDomainsResponseBody) SetPageSize(v int64) *DescribeUserDomainsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeUserDomainsResponseBody) SetRequestId(v string) *DescribeUserDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserDomainsResponseBody) SetTotalCount(v int64) *DescribeUserDomainsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeUserDomainsResponseBodyDomains struct {
	PageData []*DescribeUserDomainsResponseBodyDomainsPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
}

func (s DescribeUserDomainsResponseBodyDomains) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserDomainsResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *DescribeUserDomainsResponseBodyDomains) SetPageData(v []*DescribeUserDomainsResponseBodyDomainsPageData) *DescribeUserDomainsResponseBodyDomains {
	s.PageData = v
	return s
}

type DescribeUserDomainsResponseBodyDomainsPageData struct {
	CdnType         *string                                                `json:"CdnType,omitempty" xml:"CdnType,omitempty"`
	Cname           *string                                                `json:"Cname,omitempty" xml:"Cname,omitempty"`
	Description     *string                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	DomainName      *string                                                `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainStatus    *string                                                `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	GmtCreated      *string                                                `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtModified     *string                                                `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	ResourceGroupId *string                                                `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Sandbox         *string                                                `json:"Sandbox,omitempty" xml:"Sandbox,omitempty"`
	SourceType      *string                                                `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Sources         *DescribeUserDomainsResponseBodyDomainsPageDataSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Struct"`
	SslProtocol     *string                                                `json:"SslProtocol,omitempty" xml:"SslProtocol,omitempty"`
}

func (s DescribeUserDomainsResponseBodyDomainsPageData) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserDomainsResponseBodyDomainsPageData) GoString() string {
	return s.String()
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) SetCdnType(v string) *DescribeUserDomainsResponseBodyDomainsPageData {
	s.CdnType = &v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) SetCname(v string) *DescribeUserDomainsResponseBodyDomainsPageData {
	s.Cname = &v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) SetDescription(v string) *DescribeUserDomainsResponseBodyDomainsPageData {
	s.Description = &v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) SetDomainName(v string) *DescribeUserDomainsResponseBodyDomainsPageData {
	s.DomainName = &v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) SetDomainStatus(v string) *DescribeUserDomainsResponseBodyDomainsPageData {
	s.DomainStatus = &v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) SetGmtCreated(v string) *DescribeUserDomainsResponseBodyDomainsPageData {
	s.GmtCreated = &v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) SetGmtModified(v string) *DescribeUserDomainsResponseBodyDomainsPageData {
	s.GmtModified = &v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) SetResourceGroupId(v string) *DescribeUserDomainsResponseBodyDomainsPageData {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) SetSandbox(v string) *DescribeUserDomainsResponseBodyDomainsPageData {
	s.Sandbox = &v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) SetSourceType(v string) *DescribeUserDomainsResponseBodyDomainsPageData {
	s.SourceType = &v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) SetSources(v *DescribeUserDomainsResponseBodyDomainsPageDataSources) *DescribeUserDomainsResponseBodyDomainsPageData {
	s.Sources = v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) SetSslProtocol(v string) *DescribeUserDomainsResponseBodyDomainsPageData {
	s.SslProtocol = &v
	return s
}

type DescribeUserDomainsResponseBodyDomainsPageDataSources struct {
	Source []*string `json:"Source,omitempty" xml:"Source,omitempty" type:"Repeated"`
}

func (s DescribeUserDomainsResponseBodyDomainsPageDataSources) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserDomainsResponseBodyDomainsPageDataSources) GoString() string {
	return s.String()
}

func (s *DescribeUserDomainsResponseBodyDomainsPageDataSources) SetSource(v []*string) *DescribeUserDomainsResponseBodyDomainsPageDataSources {
	s.Source = v
	return s
}

type DescribeUserDomainsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeUserDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserDomainsResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserDomainsResponse) SetHeaders(v map[string]*string) *DescribeUserDomainsResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserDomainsResponse) SetStatusCode(v int32) *DescribeUserDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserDomainsResponse) SetBody(v *DescribeUserDomainsResponseBody) *DescribeUserDomainsResponse {
	s.Body = v
	return s
}

type OpenCdnServiceRequest struct {
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken      *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s OpenCdnServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenCdnServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenCdnServiceRequest) SetInternetChargeType(v string) *OpenCdnServiceRequest {
	s.InternetChargeType = &v
	return s
}

func (s *OpenCdnServiceRequest) SetOwnerId(v int64) *OpenCdnServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *OpenCdnServiceRequest) SetSecurityToken(v string) *OpenCdnServiceRequest {
	s.SecurityToken = &v
	return s
}

type OpenCdnServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenCdnServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenCdnServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenCdnServiceResponseBody) SetRequestId(v string) *OpenCdnServiceResponseBody {
	s.RequestId = &v
	return s
}

type OpenCdnServiceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OpenCdnServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenCdnServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenCdnServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenCdnServiceResponse) SetHeaders(v map[string]*string) *OpenCdnServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenCdnServiceResponse) SetStatusCode(v int32) *OpenCdnServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenCdnServiceResponse) SetBody(v *OpenCdnServiceResponseBody) *OpenCdnServiceResponse {
	s.Body = v
	return s
}

type PushObjectCacheRequest struct {
	Area          *string `json:"Area,omitempty" xml:"Area,omitempty"`
	ObjectPath    *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s PushObjectCacheRequest) String() string {
	return tea.Prettify(s)
}

func (s PushObjectCacheRequest) GoString() string {
	return s.String()
}

func (s *PushObjectCacheRequest) SetArea(v string) *PushObjectCacheRequest {
	s.Area = &v
	return s
}

func (s *PushObjectCacheRequest) SetObjectPath(v string) *PushObjectCacheRequest {
	s.ObjectPath = &v
	return s
}

func (s *PushObjectCacheRequest) SetOwnerId(v int64) *PushObjectCacheRequest {
	s.OwnerId = &v
	return s
}

func (s *PushObjectCacheRequest) SetSecurityToken(v string) *PushObjectCacheRequest {
	s.SecurityToken = &v
	return s
}

type PushObjectCacheResponseBody struct {
	PushTaskId *string `json:"PushTaskId,omitempty" xml:"PushTaskId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PushObjectCacheResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushObjectCacheResponseBody) GoString() string {
	return s.String()
}

func (s *PushObjectCacheResponseBody) SetPushTaskId(v string) *PushObjectCacheResponseBody {
	s.PushTaskId = &v
	return s
}

func (s *PushObjectCacheResponseBody) SetRequestId(v string) *PushObjectCacheResponseBody {
	s.RequestId = &v
	return s
}

type PushObjectCacheResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PushObjectCacheResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PushObjectCacheResponse) String() string {
	return tea.Prettify(s)
}

func (s PushObjectCacheResponse) GoString() string {
	return s.String()
}

func (s *PushObjectCacheResponse) SetHeaders(v map[string]*string) *PushObjectCacheResponse {
	s.Headers = v
	return s
}

func (s *PushObjectCacheResponse) SetStatusCode(v int32) *PushObjectCacheResponse {
	s.StatusCode = &v
	return s
}

func (s *PushObjectCacheResponse) SetBody(v *PushObjectCacheResponseBody) *PushObjectCacheResponse {
	s.Body = v
	return s
}

type RefreshObjectCachesRequest struct {
	ObjectPath    *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	ObjectType    *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s RefreshObjectCachesRequest) String() string {
	return tea.Prettify(s)
}

func (s RefreshObjectCachesRequest) GoString() string {
	return s.String()
}

func (s *RefreshObjectCachesRequest) SetObjectPath(v string) *RefreshObjectCachesRequest {
	s.ObjectPath = &v
	return s
}

func (s *RefreshObjectCachesRequest) SetObjectType(v string) *RefreshObjectCachesRequest {
	s.ObjectType = &v
	return s
}

func (s *RefreshObjectCachesRequest) SetOwnerId(v int64) *RefreshObjectCachesRequest {
	s.OwnerId = &v
	return s
}

func (s *RefreshObjectCachesRequest) SetSecurityToken(v string) *RefreshObjectCachesRequest {
	s.SecurityToken = &v
	return s
}

type RefreshObjectCachesResponseBody struct {
	RefreshTaskId *string `json:"RefreshTaskId,omitempty" xml:"RefreshTaskId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefreshObjectCachesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefreshObjectCachesResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshObjectCachesResponseBody) SetRefreshTaskId(v string) *RefreshObjectCachesResponseBody {
	s.RefreshTaskId = &v
	return s
}

func (s *RefreshObjectCachesResponseBody) SetRequestId(v string) *RefreshObjectCachesResponseBody {
	s.RequestId = &v
	return s
}

type RefreshObjectCachesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RefreshObjectCachesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RefreshObjectCachesResponse) String() string {
	return tea.Prettify(s)
}

func (s RefreshObjectCachesResponse) GoString() string {
	return s.String()
}

func (s *RefreshObjectCachesResponse) SetHeaders(v map[string]*string) *RefreshObjectCachesResponse {
	s.Headers = v
	return s
}

func (s *RefreshObjectCachesResponse) SetStatusCode(v int32) *RefreshObjectCachesResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshObjectCachesResponse) SetBody(v *RefreshObjectCachesResponseBody) *RefreshObjectCachesResponse {
	s.Body = v
	return s
}

type TestDescribeDomainBpsDataRequest struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainType     *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval       *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TimeMerge      *string `json:"TimeMerge,omitempty" xml:"TimeMerge,omitempty"`
}

func (s TestDescribeDomainBpsDataRequest) String() string {
	return tea.Prettify(s)
}

func (s TestDescribeDomainBpsDataRequest) GoString() string {
	return s.String()
}

func (s *TestDescribeDomainBpsDataRequest) SetDomainName(v string) *TestDescribeDomainBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *TestDescribeDomainBpsDataRequest) SetDomainType(v string) *TestDescribeDomainBpsDataRequest {
	s.DomainType = &v
	return s
}

func (s *TestDescribeDomainBpsDataRequest) SetEndTime(v string) *TestDescribeDomainBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *TestDescribeDomainBpsDataRequest) SetInterval(v string) *TestDescribeDomainBpsDataRequest {
	s.Interval = &v
	return s
}

func (s *TestDescribeDomainBpsDataRequest) SetIspNameEn(v string) *TestDescribeDomainBpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *TestDescribeDomainBpsDataRequest) SetLocationNameEn(v string) *TestDescribeDomainBpsDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *TestDescribeDomainBpsDataRequest) SetStartTime(v string) *TestDescribeDomainBpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *TestDescribeDomainBpsDataRequest) SetTimeMerge(v string) *TestDescribeDomainBpsDataRequest {
	s.TimeMerge = &v
	return s
}

type TestDescribeDomainBpsDataResponseBody struct {
	BpsDataPerInterval *TestDescribeDomainBpsDataResponseBodyBpsDataPerInterval `json:"BpsDataPerInterval,omitempty" xml:"BpsDataPerInterval,omitempty" type:"Struct"`
	DataInterval       *string                                                  `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName         *string                                                  `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime            *string                                                  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IspName            *string                                                  `json:"IspName,omitempty" xml:"IspName,omitempty"`
	IspNameEn          *string                                                  `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationName       *string                                                  `json:"LocationName,omitempty" xml:"LocationName,omitempty"`
	LocationNameEn     *string                                                  `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	RequestId          *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime          *string                                                  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s TestDescribeDomainBpsDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TestDescribeDomainBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *TestDescribeDomainBpsDataResponseBody) SetBpsDataPerInterval(v *TestDescribeDomainBpsDataResponseBodyBpsDataPerInterval) *TestDescribeDomainBpsDataResponseBody {
	s.BpsDataPerInterval = v
	return s
}

func (s *TestDescribeDomainBpsDataResponseBody) SetDataInterval(v string) *TestDescribeDomainBpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *TestDescribeDomainBpsDataResponseBody) SetDomainName(v string) *TestDescribeDomainBpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *TestDescribeDomainBpsDataResponseBody) SetEndTime(v string) *TestDescribeDomainBpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *TestDescribeDomainBpsDataResponseBody) SetIspName(v string) *TestDescribeDomainBpsDataResponseBody {
	s.IspName = &v
	return s
}

func (s *TestDescribeDomainBpsDataResponseBody) SetIspNameEn(v string) *TestDescribeDomainBpsDataResponseBody {
	s.IspNameEn = &v
	return s
}

func (s *TestDescribeDomainBpsDataResponseBody) SetLocationName(v string) *TestDescribeDomainBpsDataResponseBody {
	s.LocationName = &v
	return s
}

func (s *TestDescribeDomainBpsDataResponseBody) SetLocationNameEn(v string) *TestDescribeDomainBpsDataResponseBody {
	s.LocationNameEn = &v
	return s
}

func (s *TestDescribeDomainBpsDataResponseBody) SetRequestId(v string) *TestDescribeDomainBpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *TestDescribeDomainBpsDataResponseBody) SetStartTime(v string) *TestDescribeDomainBpsDataResponseBody {
	s.StartTime = &v
	return s
}

type TestDescribeDomainBpsDataResponseBodyBpsDataPerInterval struct {
	DataModule []*TestDescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s TestDescribeDomainBpsDataResponseBodyBpsDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s TestDescribeDomainBpsDataResponseBodyBpsDataPerInterval) GoString() string {
	return s.String()
}

func (s *TestDescribeDomainBpsDataResponseBodyBpsDataPerInterval) SetDataModule(v []*TestDescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) *TestDescribeDomainBpsDataResponseBodyBpsDataPerInterval {
	s.DataModule = v
	return s
}

type TestDescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule struct {
	DomesticL2Value      *string `json:"DomesticL2Value,omitempty" xml:"DomesticL2Value,omitempty"`
	DomesticValue        *string `json:"DomesticValue,omitempty" xml:"DomesticValue,omitempty"`
	DynamicDomesticValue *string `json:"DynamicDomesticValue,omitempty" xml:"DynamicDomesticValue,omitempty"`
	DynamicOverseasValue *string `json:"DynamicOverseasValue,omitempty" xml:"DynamicOverseasValue,omitempty"`
	DynamicValue         *string `json:"DynamicValue,omitempty" xml:"DynamicValue,omitempty"`
	L2Value              *string `json:"L2Value,omitempty" xml:"L2Value,omitempty"`
	OverseasL2Value      *string `json:"OverseasL2Value,omitempty" xml:"OverseasL2Value,omitempty"`
	OverseasValue        *string `json:"OverseasValue,omitempty" xml:"OverseasValue,omitempty"`
	StaticDomesticValue  *string `json:"StaticDomesticValue,omitempty" xml:"StaticDomesticValue,omitempty"`
	StaticOverseasValue  *string `json:"StaticOverseasValue,omitempty" xml:"StaticOverseasValue,omitempty"`
	StaticValue          *string `json:"StaticValue,omitempty" xml:"StaticValue,omitempty"`
	TimeStamp            *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value                *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TestDescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s TestDescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *TestDescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetDomesticL2Value(v string) *TestDescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.DomesticL2Value = &v
	return s
}

func (s *TestDescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetDomesticValue(v string) *TestDescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.DomesticValue = &v
	return s
}

func (s *TestDescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetDynamicDomesticValue(v string) *TestDescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.DynamicDomesticValue = &v
	return s
}

func (s *TestDescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetDynamicOverseasValue(v string) *TestDescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.DynamicOverseasValue = &v
	return s
}

func (s *TestDescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetDynamicValue(v string) *TestDescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.DynamicValue = &v
	return s
}

func (s *TestDescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetL2Value(v string) *TestDescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.L2Value = &v
	return s
}

func (s *TestDescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetOverseasL2Value(v string) *TestDescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.OverseasL2Value = &v
	return s
}

func (s *TestDescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetOverseasValue(v string) *TestDescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.OverseasValue = &v
	return s
}

func (s *TestDescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetStaticDomesticValue(v string) *TestDescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.StaticDomesticValue = &v
	return s
}

func (s *TestDescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetStaticOverseasValue(v string) *TestDescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.StaticOverseasValue = &v
	return s
}

func (s *TestDescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetStaticValue(v string) *TestDescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.StaticValue = &v
	return s
}

func (s *TestDescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetTimeStamp(v string) *TestDescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *TestDescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetValue(v string) *TestDescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.Value = &v
	return s
}

type TestDescribeDomainBpsDataResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TestDescribeDomainBpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TestDescribeDomainBpsDataResponse) String() string {
	return tea.Prettify(s)
}

func (s TestDescribeDomainBpsDataResponse) GoString() string {
	return s.String()
}

func (s *TestDescribeDomainBpsDataResponse) SetHeaders(v map[string]*string) *TestDescribeDomainBpsDataResponse {
	s.Headers = v
	return s
}

func (s *TestDescribeDomainBpsDataResponse) SetStatusCode(v int32) *TestDescribeDomainBpsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *TestDescribeDomainBpsDataResponse) SetBody(v *TestDescribeDomainBpsDataResponseBody) *TestDescribeDomainBpsDataResponse {
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
		"ap-northeast-1": tea.String("cdn.ap-southeast-1.aliyuncs.com"),
		"ap-south-1":     tea.String("cdn.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-1": tea.String("cdn.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2": tea.String("cdn.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-3": tea.String("cdn.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-5": tea.String("cdn.ap-southeast-1.aliyuncs.com"),
		"eu-central-1":   tea.String("cdn.ap-southeast-1.aliyuncs.com"),
		"eu-west-1":      tea.String("cdn.ap-southeast-1.aliyuncs.com"),
		"me-east-1":      tea.String("cdn.ap-southeast-1.aliyuncs.com"),
		"us-east-1":      tea.String("cdn.ap-southeast-1.aliyuncs.com"),
		"us-west-1":      tea.String("cdn.ap-southeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("cdn"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddCdnDomainWithOptions(request *AddCdnDomainRequest, runtime *util.RuntimeOptions) (_result *AddCdnDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CdnType)) {
		query["CdnType"] = request.CdnType
	}

	if !tea.BoolValue(util.IsUnset(request.CheckUrl)) {
		query["CheckUrl"] = request.CheckUrl
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Priorities)) {
		query["Priorities"] = request.Priorities
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		query["Scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.SourcePort)) {
		query["SourcePort"] = request.SourcePort
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Sources)) {
		query["Sources"] = request.Sources
	}

	if !tea.BoolValue(util.IsUnset(request.TopLevelDomain)) {
		query["TopLevelDomain"] = request.TopLevelDomain
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddCdnDomain"),
		Version:     tea.String("2014-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddCdnDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddCdnDomain(request *AddCdnDomainRequest) (_result *AddCdnDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddCdnDomainResponse{}
	_body, _err := client.AddCdnDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCdnDomainDetailWithOptions(request *DescribeCdnDomainDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeCdnDomainDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCdnDomainDetail"),
		Version:     tea.String("2014-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCdnDomainDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCdnDomainDetail(request *DescribeCdnDomainDetailRequest) (_result *DescribeCdnDomainDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCdnDomainDetailResponse{}
	_body, _err := client.DescribeCdnDomainDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCdnDomainLogsWithOptions(request *DescribeCdnDomainLogsRequest, runtime *util.RuntimeOptions) (_result *DescribeCdnDomainLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.LogDay)) {
		query["LogDay"] = request.LogDay
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCdnDomainLogs"),
		Version:     tea.String("2014-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCdnDomainLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCdnDomainLogs(request *DescribeCdnDomainLogsRequest) (_result *DescribeCdnDomainLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCdnDomainLogsResponse{}
	_body, _err := client.DescribeCdnDomainLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCdnServiceWithOptions(request *DescribeCdnServiceRequest, runtime *util.RuntimeOptions) (_result *DescribeCdnServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCdnService"),
		Version:     tea.String("2014-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCdnServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCdnService(request *DescribeCdnServiceRequest) (_result *DescribeCdnServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCdnServiceResponse{}
	_body, _err := client.DescribeCdnServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainBpsDataWithOptions(request *DescribeDomainBpsDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainBpsDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.DomainType)) {
		query["DomainType"] = request.DomainType
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.IspNameEn)) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.LocationNameEn)) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TimeMerge)) {
		query["TimeMerge"] = request.TimeMerge
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomainBpsData"),
		Version:     tea.String("2014-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDomainBpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainBpsData(request *DescribeDomainBpsDataRequest) (_result *DescribeDomainBpsDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainBpsDataResponse{}
	_body, _err := client.DescribeDomainBpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainBpsDataByTimeStampWithOptions(request *DescribeDomainBpsDataByTimeStampRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainBpsDataByTimeStampResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.IspNames)) {
		query["IspNames"] = request.IspNames
	}

	if !tea.BoolValue(util.IsUnset(request.LocationNames)) {
		query["LocationNames"] = request.LocationNames
	}

	if !tea.BoolValue(util.IsUnset(request.TimePoint)) {
		query["TimePoint"] = request.TimePoint
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomainBpsDataByTimeStamp"),
		Version:     tea.String("2014-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDomainBpsDataByTimeStampResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainBpsDataByTimeStamp(request *DescribeDomainBpsDataByTimeStampRequest) (_result *DescribeDomainBpsDataByTimeStampResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainBpsDataByTimeStampResponse{}
	_body, _err := client.DescribeDomainBpsDataByTimeStampWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainHitRateDataWithOptions(request *DescribeDomainHitRateDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainHitRateDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomainHitRateData"),
		Version:     tea.String("2014-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDomainHitRateDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainHitRateData(request *DescribeDomainHitRateDataRequest) (_result *DescribeDomainHitRateDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainHitRateDataResponse{}
	_body, _err := client.DescribeDomainHitRateDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainHttpCodeDataWithOptions(request *DescribeDomainHttpCodeDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainHttpCodeDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.IspNameEn)) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.LocationNameEn)) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TimeMerge)) {
		query["TimeMerge"] = request.TimeMerge
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomainHttpCodeData"),
		Version:     tea.String("2014-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDomainHttpCodeDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainHttpCodeData(request *DescribeDomainHttpCodeDataRequest) (_result *DescribeDomainHttpCodeDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainHttpCodeDataResponse{}
	_body, _err := client.DescribeDomainHttpCodeDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainQpsDataWithOptions(request *DescribeDomainQpsDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainQpsDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.DomainType)) {
		query["DomainType"] = request.DomainType
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.IspNameEn)) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.LocationNameEn)) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TimeMerge)) {
		query["TimeMerge"] = request.TimeMerge
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomainQpsData"),
		Version:     tea.String("2014-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDomainQpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainQpsData(request *DescribeDomainQpsDataRequest) (_result *DescribeDomainQpsDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainQpsDataResponse{}
	_body, _err := client.DescribeDomainQpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainReqHitRateDataWithOptions(request *DescribeDomainReqHitRateDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainReqHitRateDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomainReqHitRateData"),
		Version:     tea.String("2014-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDomainReqHitRateDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainReqHitRateData(request *DescribeDomainReqHitRateDataRequest) (_result *DescribeDomainReqHitRateDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainReqHitRateDataResponse{}
	_body, _err := client.DescribeDomainReqHitRateDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainSrcBpsDataWithOptions(request *DescribeDomainSrcBpsDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainSrcBpsDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.FixTimeGap)) {
		query["FixTimeGap"] = request.FixTimeGap
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TimeMerge)) {
		query["TimeMerge"] = request.TimeMerge
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomainSrcBpsData"),
		Version:     tea.String("2014-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDomainSrcBpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainSrcBpsData(request *DescribeDomainSrcBpsDataRequest) (_result *DescribeDomainSrcBpsDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainSrcBpsDataResponse{}
	_body, _err := client.DescribeDomainSrcBpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainSrcFlowDataWithOptions(request *DescribeDomainSrcFlowDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainSrcFlowDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.FixTimeGap)) {
		query["FixTimeGap"] = request.FixTimeGap
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TimeMerge)) {
		query["TimeMerge"] = request.TimeMerge
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomainSrcFlowData"),
		Version:     tea.String("2014-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDomainSrcFlowDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainSrcFlowData(request *DescribeDomainSrcFlowDataRequest) (_result *DescribeDomainSrcFlowDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainSrcFlowDataResponse{}
	_body, _err := client.DescribeDomainSrcFlowDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainUvDataWithOptions(request *DescribeDomainUvDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainUvDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomainUvData"),
		Version:     tea.String("2014-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDomainUvDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainUvData(request *DescribeDomainUvDataRequest) (_result *DescribeDomainUvDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainUvDataResponse{}
	_body, _err := client.DescribeDomainUvDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserDomainsWithOptions(request *DescribeUserDomainsRequest, runtime *util.RuntimeOptions) (_result *DescribeUserDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CdnType)) {
		query["CdnType"] = request.CdnType
	}

	if !tea.BoolValue(util.IsUnset(request.CheckDomainShow)) {
		query["CheckDomainShow"] = request.CheckDomainShow
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.DomainSearchType)) {
		query["DomainSearchType"] = request.DomainSearchType
	}

	if !tea.BoolValue(util.IsUnset(request.DomainStatus)) {
		query["DomainStatus"] = request.DomainStatus
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Sources)) {
		query["Sources"] = request.Sources
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUserDomains"),
		Version:     tea.String("2014-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUserDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserDomains(request *DescribeUserDomainsRequest) (_result *DescribeUserDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserDomainsResponse{}
	_body, _err := client.DescribeUserDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenCdnServiceWithOptions(request *OpenCdnServiceRequest, runtime *util.RuntimeOptions) (_result *OpenCdnServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InternetChargeType)) {
		query["InternetChargeType"] = request.InternetChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OpenCdnService"),
		Version:     tea.String("2014-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenCdnServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenCdnService(request *OpenCdnServiceRequest) (_result *OpenCdnServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenCdnServiceResponse{}
	_body, _err := client.OpenCdnServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PushObjectCacheWithOptions(request *PushObjectCacheRequest, runtime *util.RuntimeOptions) (_result *PushObjectCacheResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Area)) {
		query["Area"] = request.Area
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectPath)) {
		query["ObjectPath"] = request.ObjectPath
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PushObjectCache"),
		Version:     tea.String("2014-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PushObjectCacheResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PushObjectCache(request *PushObjectCacheRequest) (_result *PushObjectCacheResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PushObjectCacheResponse{}
	_body, _err := client.PushObjectCacheWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefreshObjectCachesWithOptions(request *RefreshObjectCachesRequest, runtime *util.RuntimeOptions) (_result *RefreshObjectCachesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ObjectPath)) {
		query["ObjectPath"] = request.ObjectPath
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectType)) {
		query["ObjectType"] = request.ObjectType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RefreshObjectCaches"),
		Version:     tea.String("2014-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RefreshObjectCachesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefreshObjectCaches(request *RefreshObjectCachesRequest) (_result *RefreshObjectCachesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefreshObjectCachesResponse{}
	_body, _err := client.RefreshObjectCachesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TestDescribeDomainBpsDataWithOptions(request *TestDescribeDomainBpsDataRequest, runtime *util.RuntimeOptions) (_result *TestDescribeDomainBpsDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.DomainType)) {
		query["DomainType"] = request.DomainType
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.IspNameEn)) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.LocationNameEn)) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TimeMerge)) {
		query["TimeMerge"] = request.TimeMerge
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TestDescribeDomainBpsData"),
		Version:     tea.String("2014-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TestDescribeDomainBpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TestDescribeDomainBpsData(request *TestDescribeDomainBpsDataRequest) (_result *TestDescribeDomainBpsDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TestDescribeDomainBpsDataResponse{}
	_body, _err := client.TestDescribeDomainBpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
