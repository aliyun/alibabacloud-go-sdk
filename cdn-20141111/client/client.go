// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
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
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s *DescribeDomainBpsDataRequest) SetOwnerId(v int64) *DescribeDomainBpsDataRequest {
	s.OwnerId = &v
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
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s *DescribeDomainBpsDataByTimeStampRequest) SetOwnerId(v int64) *DescribeDomainBpsDataByTimeStampRequest {
	s.OwnerId = &v
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

type DescribeDomainFileSizeProportionDataRequest struct {
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainFileSizeProportionDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainFileSizeProportionDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainFileSizeProportionDataRequest) SetDomainName(v string) *DescribeDomainFileSizeProportionDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainFileSizeProportionDataRequest) SetEndTime(v string) *DescribeDomainFileSizeProportionDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainFileSizeProportionDataRequest) SetOwnerId(v int64) *DescribeDomainFileSizeProportionDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainFileSizeProportionDataRequest) SetSecurityToken(v string) *DescribeDomainFileSizeProportionDataRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDomainFileSizeProportionDataRequest) SetStartTime(v string) *DescribeDomainFileSizeProportionDataRequest {
	s.StartTime = &v
	return s
}

type DescribeDomainFileSizeProportionDataResponseBody struct {
	DataInterval                   *string                                                                         `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName                     *string                                                                         `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime                        *string                                                                         `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FileSizeProportionDataInterval *DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataInterval `json:"FileSizeProportionDataInterval,omitempty" xml:"FileSizeProportionDataInterval,omitempty" type:"Struct"`
	RequestId                      *string                                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime                      *string                                                                         `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainFileSizeProportionDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainFileSizeProportionDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainFileSizeProportionDataResponseBody) SetDataInterval(v string) *DescribeDomainFileSizeProportionDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainFileSizeProportionDataResponseBody) SetDomainName(v string) *DescribeDomainFileSizeProportionDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainFileSizeProportionDataResponseBody) SetEndTime(v string) *DescribeDomainFileSizeProportionDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainFileSizeProportionDataResponseBody) SetFileSizeProportionDataInterval(v *DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataInterval) *DescribeDomainFileSizeProportionDataResponseBody {
	s.FileSizeProportionDataInterval = v
	return s
}

func (s *DescribeDomainFileSizeProportionDataResponseBody) SetRequestId(v string) *DescribeDomainFileSizeProportionDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainFileSizeProportionDataResponseBody) SetStartTime(v string) *DescribeDomainFileSizeProportionDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataInterval struct {
	UsageData []*DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataInterval) SetUsageData(v []*DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageData) *DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataInterval {
	s.UsageData = v
	return s
}

type DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageData struct {
	TimeStamp *string                                                                                       `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value     *DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageDataValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageData) GoString() string {
	return s.String()
}

func (s *DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageData) SetTimeStamp(v string) *DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageData) SetValue(v *DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageDataValue) *DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageData {
	s.Value = v
	return s
}

type DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageDataValue struct {
	FileSizeProportionData []*DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageDataValueFileSizeProportionData `json:"FileSizeProportionData,omitempty" xml:"FileSizeProportionData,omitempty" type:"Repeated"`
}

func (s DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageDataValue) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageDataValue) GoString() string {
	return s.String()
}

func (s *DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageDataValue) SetFileSizeProportionData(v []*DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageDataValueFileSizeProportionData) *DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageDataValue {
	s.FileSizeProportionData = v
	return s
}

type DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageDataValueFileSizeProportionData struct {
	FileSize   *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	Proportion *string `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
}

func (s DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageDataValueFileSizeProportionData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageDataValueFileSizeProportionData) GoString() string {
	return s.String()
}

func (s *DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageDataValueFileSizeProportionData) SetFileSize(v string) *DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageDataValueFileSizeProportionData {
	s.FileSize = &v
	return s
}

func (s *DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageDataValueFileSizeProportionData) SetProportion(v string) *DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageDataValueFileSizeProportionData {
	s.Proportion = &v
	return s
}

type DescribeDomainFileSizeProportionDataResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDomainFileSizeProportionDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainFileSizeProportionDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainFileSizeProportionDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainFileSizeProportionDataResponse) SetHeaders(v map[string]*string) *DescribeDomainFileSizeProportionDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainFileSizeProportionDataResponse) SetStatusCode(v int32) *DescribeDomainFileSizeProportionDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainFileSizeProportionDataResponse) SetBody(v *DescribeDomainFileSizeProportionDataResponseBody) *DescribeDomainFileSizeProportionDataResponse {
	s.Body = v
	return s
}

type DescribeDomainFlowDataRequest struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainType     *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval       *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TimeMerge      *string `json:"TimeMerge,omitempty" xml:"TimeMerge,omitempty"`
}

func (s DescribeDomainFlowDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainFlowDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainFlowDataRequest) SetDomainName(v string) *DescribeDomainFlowDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainFlowDataRequest) SetDomainType(v string) *DescribeDomainFlowDataRequest {
	s.DomainType = &v
	return s
}

func (s *DescribeDomainFlowDataRequest) SetEndTime(v string) *DescribeDomainFlowDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainFlowDataRequest) SetInterval(v string) *DescribeDomainFlowDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainFlowDataRequest) SetIspNameEn(v string) *DescribeDomainFlowDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDomainFlowDataRequest) SetLocationNameEn(v string) *DescribeDomainFlowDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDomainFlowDataRequest) SetOwnerId(v int64) *DescribeDomainFlowDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainFlowDataRequest) SetStartTime(v string) *DescribeDomainFlowDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainFlowDataRequest) SetTimeMerge(v string) *DescribeDomainFlowDataRequest {
	s.TimeMerge = &v
	return s
}

type DescribeDomainFlowDataResponseBody struct {
	DataInterval        *string                                                `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName          *string                                                `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime             *string                                                `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FlowDataPerInterval *DescribeDomainFlowDataResponseBodyFlowDataPerInterval `json:"FlowDataPerInterval,omitempty" xml:"FlowDataPerInterval,omitempty" type:"Struct"`
	RequestId           *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime           *string                                                `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainFlowDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainFlowDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainFlowDataResponseBody) SetDataInterval(v string) *DescribeDomainFlowDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainFlowDataResponseBody) SetDomainName(v string) *DescribeDomainFlowDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainFlowDataResponseBody) SetEndTime(v string) *DescribeDomainFlowDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainFlowDataResponseBody) SetFlowDataPerInterval(v *DescribeDomainFlowDataResponseBodyFlowDataPerInterval) *DescribeDomainFlowDataResponseBody {
	s.FlowDataPerInterval = v
	return s
}

func (s *DescribeDomainFlowDataResponseBody) SetRequestId(v string) *DescribeDomainFlowDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainFlowDataResponseBody) SetStartTime(v string) *DescribeDomainFlowDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeDomainFlowDataResponseBodyFlowDataPerInterval struct {
	DataModule []*DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainFlowDataResponseBodyFlowDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainFlowDataResponseBodyFlowDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainFlowDataResponseBodyFlowDataPerInterval) SetDataModule(v []*DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule) *DescribeDomainFlowDataResponseBodyFlowDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule struct {
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

func (s DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule) SetDomesticValue(v string) *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule {
	s.DomesticValue = &v
	return s
}

func (s *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule) SetDynamicDomesticValue(v string) *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule {
	s.DynamicDomesticValue = &v
	return s
}

func (s *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule) SetDynamicOverseasValue(v string) *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule {
	s.DynamicOverseasValue = &v
	return s
}

func (s *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule) SetDynamicValue(v string) *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule {
	s.DynamicValue = &v
	return s
}

func (s *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule) SetOverseasValue(v string) *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule {
	s.OverseasValue = &v
	return s
}

func (s *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule) SetStaticDomesticValue(v string) *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule {
	s.StaticDomesticValue = &v
	return s
}

func (s *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule) SetStaticOverseasValue(v string) *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule {
	s.StaticOverseasValue = &v
	return s
}

func (s *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule) SetStaticValue(v string) *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule {
	s.StaticValue = &v
	return s
}

func (s *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule) SetValue(v string) *DescribeDomainFlowDataResponseBodyFlowDataPerIntervalDataModule {
	s.Value = &v
	return s
}

type DescribeDomainFlowDataResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDomainFlowDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainFlowDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainFlowDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainFlowDataResponse) SetHeaders(v map[string]*string) *DescribeDomainFlowDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainFlowDataResponse) SetStatusCode(v int32) *DescribeDomainFlowDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainFlowDataResponse) SetBody(v *DescribeDomainFlowDataResponseBody) *DescribeDomainFlowDataResponse {
	s.Body = v
	return s
}

type DescribeDomainHitRateDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval   *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s *DescribeDomainHitRateDataRequest) SetOwnerId(v int64) *DescribeDomainHitRateDataRequest {
	s.OwnerId = &v
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
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s *DescribeDomainHttpCodeDataRequest) SetOwnerId(v int64) *DescribeDomainHttpCodeDataRequest {
	s.OwnerId = &v
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

type DescribeDomainISPDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainISPDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainISPDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainISPDataRequest) SetDomainName(v string) *DescribeDomainISPDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainISPDataRequest) SetEndTime(v string) *DescribeDomainISPDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainISPDataRequest) SetOwnerId(v int64) *DescribeDomainISPDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainISPDataRequest) SetStartTime(v string) *DescribeDomainISPDataRequest {
	s.StartTime = &v
	return s
}

type DescribeDomainISPDataResponseBody struct {
	DataInterval *string                                 `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName   *string                                 `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime      *string                                 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime    *string                                 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Value        *DescribeDomainISPDataResponseBodyValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s DescribeDomainISPDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainISPDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainISPDataResponseBody) SetDataInterval(v string) *DescribeDomainISPDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainISPDataResponseBody) SetDomainName(v string) *DescribeDomainISPDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainISPDataResponseBody) SetEndTime(v string) *DescribeDomainISPDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainISPDataResponseBody) SetRequestId(v string) *DescribeDomainISPDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainISPDataResponseBody) SetStartTime(v string) *DescribeDomainISPDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainISPDataResponseBody) SetValue(v *DescribeDomainISPDataResponseBodyValue) *DescribeDomainISPDataResponseBody {
	s.Value = v
	return s
}

type DescribeDomainISPDataResponseBodyValue struct {
	ISPProportionData []*DescribeDomainISPDataResponseBodyValueISPProportionData `json:"ISPProportionData,omitempty" xml:"ISPProportionData,omitempty" type:"Repeated"`
}

func (s DescribeDomainISPDataResponseBodyValue) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainISPDataResponseBodyValue) GoString() string {
	return s.String()
}

func (s *DescribeDomainISPDataResponseBodyValue) SetISPProportionData(v []*DescribeDomainISPDataResponseBodyValueISPProportionData) *DescribeDomainISPDataResponseBodyValue {
	s.ISPProportionData = v
	return s
}

type DescribeDomainISPDataResponseBodyValueISPProportionData struct {
	AvgObjectSize   *string `json:"AvgObjectSize,omitempty" xml:"AvgObjectSize,omitempty"`
	AvgResponseRate *string `json:"AvgResponseRate,omitempty" xml:"AvgResponseRate,omitempty"`
	AvgResponseTime *string `json:"AvgResponseTime,omitempty" xml:"AvgResponseTime,omitempty"`
	Bps             *string `json:"Bps,omitempty" xml:"Bps,omitempty"`
	ByteHitRate     *string `json:"ByteHitRate,omitempty" xml:"ByteHitRate,omitempty"`
	BytesProportion *string `json:"BytesProportion,omitempty" xml:"BytesProportion,omitempty"`
	ISP             *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	IspEname        *string `json:"IspEname,omitempty" xml:"IspEname,omitempty"`
	Proportion      *string `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
	Qps             *string `json:"Qps,omitempty" xml:"Qps,omitempty"`
	ReqErrRate      *string `json:"ReqErrRate,omitempty" xml:"ReqErrRate,omitempty"`
	ReqHitRate      *string `json:"ReqHitRate,omitempty" xml:"ReqHitRate,omitempty"`
	TotalBytes      *string `json:"TotalBytes,omitempty" xml:"TotalBytes,omitempty"`
	TotalQuery      *string `json:"TotalQuery,omitempty" xml:"TotalQuery,omitempty"`
}

func (s DescribeDomainISPDataResponseBodyValueISPProportionData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainISPDataResponseBodyValueISPProportionData) GoString() string {
	return s.String()
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) SetAvgObjectSize(v string) *DescribeDomainISPDataResponseBodyValueISPProportionData {
	s.AvgObjectSize = &v
	return s
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) SetAvgResponseRate(v string) *DescribeDomainISPDataResponseBodyValueISPProportionData {
	s.AvgResponseRate = &v
	return s
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) SetAvgResponseTime(v string) *DescribeDomainISPDataResponseBodyValueISPProportionData {
	s.AvgResponseTime = &v
	return s
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) SetBps(v string) *DescribeDomainISPDataResponseBodyValueISPProportionData {
	s.Bps = &v
	return s
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) SetByteHitRate(v string) *DescribeDomainISPDataResponseBodyValueISPProportionData {
	s.ByteHitRate = &v
	return s
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) SetBytesProportion(v string) *DescribeDomainISPDataResponseBodyValueISPProportionData {
	s.BytesProportion = &v
	return s
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) SetISP(v string) *DescribeDomainISPDataResponseBodyValueISPProportionData {
	s.ISP = &v
	return s
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) SetIspEname(v string) *DescribeDomainISPDataResponseBodyValueISPProportionData {
	s.IspEname = &v
	return s
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) SetProportion(v string) *DescribeDomainISPDataResponseBodyValueISPProportionData {
	s.Proportion = &v
	return s
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) SetQps(v string) *DescribeDomainISPDataResponseBodyValueISPProportionData {
	s.Qps = &v
	return s
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) SetReqErrRate(v string) *DescribeDomainISPDataResponseBodyValueISPProportionData {
	s.ReqErrRate = &v
	return s
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) SetReqHitRate(v string) *DescribeDomainISPDataResponseBodyValueISPProportionData {
	s.ReqHitRate = &v
	return s
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) SetTotalBytes(v string) *DescribeDomainISPDataResponseBodyValueISPProportionData {
	s.TotalBytes = &v
	return s
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) SetTotalQuery(v string) *DescribeDomainISPDataResponseBodyValueISPProportionData {
	s.TotalQuery = &v
	return s
}

type DescribeDomainISPDataResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDomainISPDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainISPDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainISPDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainISPDataResponse) SetHeaders(v map[string]*string) *DescribeDomainISPDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainISPDataResponse) SetStatusCode(v int32) *DescribeDomainISPDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainISPDataResponse) SetBody(v *DescribeDomainISPDataResponseBody) *DescribeDomainISPDataResponse {
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
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s *DescribeDomainQpsDataRequest) SetOwnerId(v int64) *DescribeDomainQpsDataRequest {
	s.OwnerId = &v
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

type DescribeDomainRegionDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainRegionDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRegionDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainRegionDataRequest) SetDomainName(v string) *DescribeDomainRegionDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainRegionDataRequest) SetEndTime(v string) *DescribeDomainRegionDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainRegionDataRequest) SetOwnerId(v int64) *DescribeDomainRegionDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainRegionDataRequest) SetStartTime(v string) *DescribeDomainRegionDataRequest {
	s.StartTime = &v
	return s
}

type DescribeDomainRegionDataResponseBody struct {
	DataInterval *string                                    `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName   *string                                    `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime      *string                                    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId    *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime    *string                                    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Value        *DescribeDomainRegionDataResponseBodyValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s DescribeDomainRegionDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRegionDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainRegionDataResponseBody) SetDataInterval(v string) *DescribeDomainRegionDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBody) SetDomainName(v string) *DescribeDomainRegionDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBody) SetEndTime(v string) *DescribeDomainRegionDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBody) SetRequestId(v string) *DescribeDomainRegionDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBody) SetStartTime(v string) *DescribeDomainRegionDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBody) SetValue(v *DescribeDomainRegionDataResponseBodyValue) *DescribeDomainRegionDataResponseBody {
	s.Value = v
	return s
}

type DescribeDomainRegionDataResponseBodyValue struct {
	RegionProportionData []*DescribeDomainRegionDataResponseBodyValueRegionProportionData `json:"RegionProportionData,omitempty" xml:"RegionProportionData,omitempty" type:"Repeated"`
}

func (s DescribeDomainRegionDataResponseBodyValue) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRegionDataResponseBodyValue) GoString() string {
	return s.String()
}

func (s *DescribeDomainRegionDataResponseBodyValue) SetRegionProportionData(v []*DescribeDomainRegionDataResponseBodyValueRegionProportionData) *DescribeDomainRegionDataResponseBodyValue {
	s.RegionProportionData = v
	return s
}

type DescribeDomainRegionDataResponseBodyValueRegionProportionData struct {
	AvgObjectSize   *string `json:"AvgObjectSize,omitempty" xml:"AvgObjectSize,omitempty"`
	AvgResponseRate *string `json:"AvgResponseRate,omitempty" xml:"AvgResponseRate,omitempty"`
	AvgResponseTime *string `json:"AvgResponseTime,omitempty" xml:"AvgResponseTime,omitempty"`
	Bps             *string `json:"Bps,omitempty" xml:"Bps,omitempty"`
	ByteHitRate     *string `json:"ByteHitRate,omitempty" xml:"ByteHitRate,omitempty"`
	BytesProportion *string `json:"BytesProportion,omitempty" xml:"BytesProportion,omitempty"`
	Proportion      *string `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
	Qps             *string `json:"Qps,omitempty" xml:"Qps,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RegionEname     *string `json:"RegionEname,omitempty" xml:"RegionEname,omitempty"`
	ReqErrRate      *string `json:"ReqErrRate,omitempty" xml:"ReqErrRate,omitempty"`
	ReqHitRate      *string `json:"ReqHitRate,omitempty" xml:"ReqHitRate,omitempty"`
	TotalBytes      *string `json:"TotalBytes,omitempty" xml:"TotalBytes,omitempty"`
	TotalQuery      *string `json:"TotalQuery,omitempty" xml:"TotalQuery,omitempty"`
}

func (s DescribeDomainRegionDataResponseBodyValueRegionProportionData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRegionDataResponseBodyValueRegionProportionData) GoString() string {
	return s.String()
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) SetAvgObjectSize(v string) *DescribeDomainRegionDataResponseBodyValueRegionProportionData {
	s.AvgObjectSize = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) SetAvgResponseRate(v string) *DescribeDomainRegionDataResponseBodyValueRegionProportionData {
	s.AvgResponseRate = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) SetAvgResponseTime(v string) *DescribeDomainRegionDataResponseBodyValueRegionProportionData {
	s.AvgResponseTime = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) SetBps(v string) *DescribeDomainRegionDataResponseBodyValueRegionProportionData {
	s.Bps = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) SetByteHitRate(v string) *DescribeDomainRegionDataResponseBodyValueRegionProportionData {
	s.ByteHitRate = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) SetBytesProportion(v string) *DescribeDomainRegionDataResponseBodyValueRegionProportionData {
	s.BytesProportion = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) SetProportion(v string) *DescribeDomainRegionDataResponseBodyValueRegionProportionData {
	s.Proportion = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) SetQps(v string) *DescribeDomainRegionDataResponseBodyValueRegionProportionData {
	s.Qps = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) SetRegion(v string) *DescribeDomainRegionDataResponseBodyValueRegionProportionData {
	s.Region = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) SetRegionEname(v string) *DescribeDomainRegionDataResponseBodyValueRegionProportionData {
	s.RegionEname = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) SetReqErrRate(v string) *DescribeDomainRegionDataResponseBodyValueRegionProportionData {
	s.ReqErrRate = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) SetReqHitRate(v string) *DescribeDomainRegionDataResponseBodyValueRegionProportionData {
	s.ReqHitRate = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) SetTotalBytes(v string) *DescribeDomainRegionDataResponseBodyValueRegionProportionData {
	s.TotalBytes = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) SetTotalQuery(v string) *DescribeDomainRegionDataResponseBodyValueRegionProportionData {
	s.TotalQuery = &v
	return s
}

type DescribeDomainRegionDataResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDomainRegionDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainRegionDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRegionDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainRegionDataResponse) SetHeaders(v map[string]*string) *DescribeDomainRegionDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainRegionDataResponse) SetStatusCode(v int32) *DescribeDomainRegionDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainRegionDataResponse) SetBody(v *DescribeDomainRegionDataResponseBody) *DescribeDomainRegionDataResponse {
	s.Body = v
	return s
}

type DescribeDomainReqHitRateDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval   *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s *DescribeDomainReqHitRateDataRequest) SetOwnerId(v int64) *DescribeDomainReqHitRateDataRequest {
	s.OwnerId = &v
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
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s *DescribeDomainSrcBpsDataRequest) SetOwnerId(v int64) *DescribeDomainSrcBpsDataRequest {
	s.OwnerId = &v
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
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s *DescribeDomainSrcFlowDataRequest) SetOwnerId(v int64) *DescribeDomainSrcFlowDataRequest {
	s.OwnerId = &v
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

type DescribeDomainsBySourceRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Sources       *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
}

func (s DescribeDomainsBySourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsBySourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainsBySourceRequest) SetOwnerId(v int64) *DescribeDomainsBySourceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainsBySourceRequest) SetSecurityToken(v string) *DescribeDomainsBySourceRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDomainsBySourceRequest) SetSources(v string) *DescribeDomainsBySourceRequest {
	s.Sources = &v
	return s
}

type DescribeDomainsBySourceResponseBody struct {
	DomainsList *DescribeDomainsBySourceResponseBodyDomainsList `json:"DomainsList,omitempty" xml:"DomainsList,omitempty" type:"Struct"`
	RequestId   *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Sources     *string                                         `json:"Sources,omitempty" xml:"Sources,omitempty"`
}

func (s DescribeDomainsBySourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsBySourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainsBySourceResponseBody) SetDomainsList(v *DescribeDomainsBySourceResponseBodyDomainsList) *DescribeDomainsBySourceResponseBody {
	s.DomainsList = v
	return s
}

func (s *DescribeDomainsBySourceResponseBody) SetRequestId(v string) *DescribeDomainsBySourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainsBySourceResponseBody) SetSources(v string) *DescribeDomainsBySourceResponseBody {
	s.Sources = &v
	return s
}

type DescribeDomainsBySourceResponseBodyDomainsList struct {
	DomainsData []*DescribeDomainsBySourceResponseBodyDomainsListDomainsData `json:"DomainsData,omitempty" xml:"DomainsData,omitempty" type:"Repeated"`
}

func (s DescribeDomainsBySourceResponseBodyDomainsList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsBySourceResponseBodyDomainsList) GoString() string {
	return s.String()
}

func (s *DescribeDomainsBySourceResponseBodyDomainsList) SetDomainsData(v []*DescribeDomainsBySourceResponseBodyDomainsListDomainsData) *DescribeDomainsBySourceResponseBodyDomainsList {
	s.DomainsData = v
	return s
}

type DescribeDomainsBySourceResponseBodyDomainsListDomainsData struct {
	DomainInfos *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfos `json:"DomainInfos,omitempty" xml:"DomainInfos,omitempty" type:"Struct"`
	Domains     *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomains     `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Struct"`
	Source      *string                                                               `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s DescribeDomainsBySourceResponseBodyDomainsListDomainsData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsBySourceResponseBodyDomainsListDomainsData) GoString() string {
	return s.String()
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsData) SetDomainInfos(v *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfos) *DescribeDomainsBySourceResponseBodyDomainsListDomainsData {
	s.DomainInfos = v
	return s
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsData) SetDomains(v *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomains) *DescribeDomainsBySourceResponseBodyDomainsListDomainsData {
	s.Domains = v
	return s
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsData) SetSource(v string) *DescribeDomainsBySourceResponseBodyDomainsListDomainsData {
	s.Source = &v
	return s
}

type DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfos struct {
	DomainInfo []*DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo `json:"domainInfo,omitempty" xml:"domainInfo,omitempty" type:"Repeated"`
}

func (s DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfos) GoString() string {
	return s.String()
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfos) SetDomainInfo(v []*DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo) *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfos {
	s.DomainInfo = v
	return s
}

type DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo struct {
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DomainCname *string `json:"DomainCname,omitempty" xml:"DomainCname,omitempty"`
	DomainName  *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime  *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo) GoString() string {
	return s.String()
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo) SetCreateTime(v string) *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo {
	s.CreateTime = &v
	return s
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo) SetDomainCname(v string) *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo {
	s.DomainCname = &v
	return s
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo) SetDomainName(v string) *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo) SetStatus(v string) *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo {
	s.Status = &v
	return s
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo) SetUpdateTime(v string) *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo {
	s.UpdateTime = &v
	return s
}

type DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomains struct {
	DomainNames []*string `json:"domainNames,omitempty" xml:"domainNames,omitempty" type:"Repeated"`
}

func (s DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomains) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomains) GoString() string {
	return s.String()
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomains) SetDomainNames(v []*string) *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomains {
	s.DomainNames = v
	return s
}

type DescribeDomainsBySourceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDomainsBySourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainsBySourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsBySourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainsBySourceResponse) SetHeaders(v map[string]*string) *DescribeDomainsBySourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainsBySourceResponse) SetStatusCode(v int32) *DescribeDomainsBySourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainsBySourceResponse) SetBody(v *DescribeDomainsBySourceResponseBody) *DescribeDomainsBySourceResponse {
	s.Body = v
	return s
}

type DescribeDomainsUsageByDayRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainsUsageByDayRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsUsageByDayRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainsUsageByDayRequest) SetDomainName(v string) *DescribeDomainsUsageByDayRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainsUsageByDayRequest) SetEndTime(v string) *DescribeDomainsUsageByDayRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainsUsageByDayRequest) SetOwnerId(v int64) *DescribeDomainsUsageByDayRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainsUsageByDayRequest) SetStartTime(v string) *DescribeDomainsUsageByDayRequest {
	s.StartTime = &v
	return s
}

type DescribeDomainsUsageByDayResponseBody struct {
	DataInterval *string                                           `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName   *string                                           `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime      *string                                           `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId    *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime    *string                                           `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	UsageByDays  *DescribeDomainsUsageByDayResponseBodyUsageByDays `json:"UsageByDays,omitempty" xml:"UsageByDays,omitempty" type:"Struct"`
	UsageTotal   *DescribeDomainsUsageByDayResponseBodyUsageTotal  `json:"UsageTotal,omitempty" xml:"UsageTotal,omitempty" type:"Struct"`
}

func (s DescribeDomainsUsageByDayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsUsageByDayResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainsUsageByDayResponseBody) SetDataInterval(v string) *DescribeDomainsUsageByDayResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBody) SetDomainName(v string) *DescribeDomainsUsageByDayResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBody) SetEndTime(v string) *DescribeDomainsUsageByDayResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBody) SetRequestId(v string) *DescribeDomainsUsageByDayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBody) SetStartTime(v string) *DescribeDomainsUsageByDayResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBody) SetUsageByDays(v *DescribeDomainsUsageByDayResponseBodyUsageByDays) *DescribeDomainsUsageByDayResponseBody {
	s.UsageByDays = v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBody) SetUsageTotal(v *DescribeDomainsUsageByDayResponseBodyUsageTotal) *DescribeDomainsUsageByDayResponseBody {
	s.UsageTotal = v
	return s
}

type DescribeDomainsUsageByDayResponseBodyUsageByDays struct {
	UsageByDay []*DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay `json:"UsageByDay,omitempty" xml:"UsageByDay,omitempty" type:"Repeated"`
}

func (s DescribeDomainsUsageByDayResponseBodyUsageByDays) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsUsageByDayResponseBodyUsageByDays) GoString() string {
	return s.String()
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDays) SetUsageByDay(v []*DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) *DescribeDomainsUsageByDayResponseBodyUsageByDays {
	s.UsageByDay = v
	return s
}

type DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay struct {
	BytesHitRate   *string `json:"BytesHitRate,omitempty" xml:"BytesHitRate,omitempty"`
	MaxBps         *string `json:"MaxBps,omitempty" xml:"MaxBps,omitempty"`
	MaxBpsTime     *string `json:"MaxBpsTime,omitempty" xml:"MaxBpsTime,omitempty"`
	MaxSrcBps      *string `json:"MaxSrcBps,omitempty" xml:"MaxSrcBps,omitempty"`
	MaxSrcBpsTime  *string `json:"MaxSrcBpsTime,omitempty" xml:"MaxSrcBpsTime,omitempty"`
	Qps            *string `json:"Qps,omitempty" xml:"Qps,omitempty"`
	RequestHitRate *string `json:"RequestHitRate,omitempty" xml:"RequestHitRate,omitempty"`
	TimeStamp      *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	TotalAccess    *string `json:"TotalAccess,omitempty" xml:"TotalAccess,omitempty"`
	TotalTraffic   *string `json:"TotalTraffic,omitempty" xml:"TotalTraffic,omitempty"`
}

func (s DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) GoString() string {
	return s.String()
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) SetBytesHitRate(v string) *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay {
	s.BytesHitRate = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) SetMaxBps(v string) *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay {
	s.MaxBps = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) SetMaxBpsTime(v string) *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay {
	s.MaxBpsTime = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) SetMaxSrcBps(v string) *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay {
	s.MaxSrcBps = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) SetMaxSrcBpsTime(v string) *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay {
	s.MaxSrcBpsTime = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) SetQps(v string) *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay {
	s.Qps = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) SetRequestHitRate(v string) *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay {
	s.RequestHitRate = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) SetTimeStamp(v string) *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) SetTotalAccess(v string) *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay {
	s.TotalAccess = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) SetTotalTraffic(v string) *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay {
	s.TotalTraffic = &v
	return s
}

type DescribeDomainsUsageByDayResponseBodyUsageTotal struct {
	BytesHitRate   *string `json:"BytesHitRate,omitempty" xml:"BytesHitRate,omitempty"`
	MaxBps         *string `json:"MaxBps,omitempty" xml:"MaxBps,omitempty"`
	MaxBpsTime     *string `json:"MaxBpsTime,omitempty" xml:"MaxBpsTime,omitempty"`
	MaxSrcBps      *string `json:"MaxSrcBps,omitempty" xml:"MaxSrcBps,omitempty"`
	MaxSrcBpsTime  *string `json:"MaxSrcBpsTime,omitempty" xml:"MaxSrcBpsTime,omitempty"`
	RequestHitRate *string `json:"RequestHitRate,omitempty" xml:"RequestHitRate,omitempty"`
	TotalAccess    *string `json:"TotalAccess,omitempty" xml:"TotalAccess,omitempty"`
	TotalTraffic   *string `json:"TotalTraffic,omitempty" xml:"TotalTraffic,omitempty"`
}

func (s DescribeDomainsUsageByDayResponseBodyUsageTotal) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsUsageByDayResponseBodyUsageTotal) GoString() string {
	return s.String()
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageTotal) SetBytesHitRate(v string) *DescribeDomainsUsageByDayResponseBodyUsageTotal {
	s.BytesHitRate = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageTotal) SetMaxBps(v string) *DescribeDomainsUsageByDayResponseBodyUsageTotal {
	s.MaxBps = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageTotal) SetMaxBpsTime(v string) *DescribeDomainsUsageByDayResponseBodyUsageTotal {
	s.MaxBpsTime = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageTotal) SetMaxSrcBps(v string) *DescribeDomainsUsageByDayResponseBodyUsageTotal {
	s.MaxSrcBps = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageTotal) SetMaxSrcBpsTime(v string) *DescribeDomainsUsageByDayResponseBodyUsageTotal {
	s.MaxSrcBpsTime = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageTotal) SetRequestHitRate(v string) *DescribeDomainsUsageByDayResponseBodyUsageTotal {
	s.RequestHitRate = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageTotal) SetTotalAccess(v string) *DescribeDomainsUsageByDayResponseBodyUsageTotal {
	s.TotalAccess = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageTotal) SetTotalTraffic(v string) *DescribeDomainsUsageByDayResponseBodyUsageTotal {
	s.TotalTraffic = &v
	return s
}

type DescribeDomainsUsageByDayResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDomainsUsageByDayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainsUsageByDayResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsUsageByDayResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainsUsageByDayResponse) SetHeaders(v map[string]*string) *DescribeDomainsUsageByDayResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainsUsageByDayResponse) SetStatusCode(v int32) *DescribeDomainsUsageByDayResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponse) SetBody(v *DescribeDomainsUsageByDayResponseBody) *DescribeDomainsUsageByDayResponse {
	s.Body = v
	return s
}

type DescribeRefreshQuotaRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeRefreshQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRefreshQuotaRequest) GoString() string {
	return s.String()
}

func (s *DescribeRefreshQuotaRequest) SetOwnerId(v int64) *DescribeRefreshQuotaRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRefreshQuotaRequest) SetSecurityToken(v string) *DescribeRefreshQuotaRequest {
	s.SecurityToken = &v
	return s
}

type DescribeRefreshQuotaResponseBody struct {
	BlockQuota    *string `json:"BlockQuota,omitempty" xml:"BlockQuota,omitempty"`
	BlockRemain   *string `json:"BlockRemain,omitempty" xml:"BlockRemain,omitempty"`
	DirQuota      *string `json:"DirQuota,omitempty" xml:"DirQuota,omitempty"`
	DirRemain     *string `json:"DirRemain,omitempty" xml:"DirRemain,omitempty"`
	PreloadQuota  *string `json:"PreloadQuota,omitempty" xml:"PreloadQuota,omitempty"`
	PreloadRemain *string `json:"PreloadRemain,omitempty" xml:"PreloadRemain,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UrlQuota      *string `json:"UrlQuota,omitempty" xml:"UrlQuota,omitempty"`
	UrlRemain     *string `json:"UrlRemain,omitempty" xml:"UrlRemain,omitempty"`
}

func (s DescribeRefreshQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRefreshQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRefreshQuotaResponseBody) SetBlockQuota(v string) *DescribeRefreshQuotaResponseBody {
	s.BlockQuota = &v
	return s
}

func (s *DescribeRefreshQuotaResponseBody) SetBlockRemain(v string) *DescribeRefreshQuotaResponseBody {
	s.BlockRemain = &v
	return s
}

func (s *DescribeRefreshQuotaResponseBody) SetDirQuota(v string) *DescribeRefreshQuotaResponseBody {
	s.DirQuota = &v
	return s
}

func (s *DescribeRefreshQuotaResponseBody) SetDirRemain(v string) *DescribeRefreshQuotaResponseBody {
	s.DirRemain = &v
	return s
}

func (s *DescribeRefreshQuotaResponseBody) SetPreloadQuota(v string) *DescribeRefreshQuotaResponseBody {
	s.PreloadQuota = &v
	return s
}

func (s *DescribeRefreshQuotaResponseBody) SetPreloadRemain(v string) *DescribeRefreshQuotaResponseBody {
	s.PreloadRemain = &v
	return s
}

func (s *DescribeRefreshQuotaResponseBody) SetRequestId(v string) *DescribeRefreshQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRefreshQuotaResponseBody) SetUrlQuota(v string) *DescribeRefreshQuotaResponseBody {
	s.UrlQuota = &v
	return s
}

func (s *DescribeRefreshQuotaResponseBody) SetUrlRemain(v string) *DescribeRefreshQuotaResponseBody {
	s.UrlRemain = &v
	return s
}

type DescribeRefreshQuotaResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRefreshQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRefreshQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRefreshQuotaResponse) GoString() string {
	return s.String()
}

func (s *DescribeRefreshQuotaResponse) SetHeaders(v map[string]*string) *DescribeRefreshQuotaResponse {
	s.Headers = v
	return s
}

func (s *DescribeRefreshQuotaResponse) SetStatusCode(v int32) *DescribeRefreshQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRefreshQuotaResponse) SetBody(v *DescribeRefreshQuotaResponseBody) *DescribeRefreshQuotaResponse {
	s.Body = v
	return s
}

type DescribeTopDomainsByFlowRequest struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Limit     *int64  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Product   *string `json:"Product,omitempty" xml:"Product,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeTopDomainsByFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopDomainsByFlowRequest) GoString() string {
	return s.String()
}

func (s *DescribeTopDomainsByFlowRequest) SetEndTime(v string) *DescribeTopDomainsByFlowRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeTopDomainsByFlowRequest) SetLimit(v int64) *DescribeTopDomainsByFlowRequest {
	s.Limit = &v
	return s
}

func (s *DescribeTopDomainsByFlowRequest) SetOwnerId(v int64) *DescribeTopDomainsByFlowRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTopDomainsByFlowRequest) SetProduct(v string) *DescribeTopDomainsByFlowRequest {
	s.Product = &v
	return s
}

func (s *DescribeTopDomainsByFlowRequest) SetStartTime(v string) *DescribeTopDomainsByFlowRequest {
	s.StartTime = &v
	return s
}

type DescribeTopDomainsByFlowResponseBody struct {
	DomainCount       *int64                                          `json:"DomainCount,omitempty" xml:"DomainCount,omitempty"`
	DomainOnlineCount *int64                                          `json:"DomainOnlineCount,omitempty" xml:"DomainOnlineCount,omitempty"`
	EndTime           *string                                         `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId         *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime         *string                                         `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TopDomains        *DescribeTopDomainsByFlowResponseBodyTopDomains `json:"TopDomains,omitempty" xml:"TopDomains,omitempty" type:"Struct"`
}

func (s DescribeTopDomainsByFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopDomainsByFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTopDomainsByFlowResponseBody) SetDomainCount(v int64) *DescribeTopDomainsByFlowResponseBody {
	s.DomainCount = &v
	return s
}

func (s *DescribeTopDomainsByFlowResponseBody) SetDomainOnlineCount(v int64) *DescribeTopDomainsByFlowResponseBody {
	s.DomainOnlineCount = &v
	return s
}

func (s *DescribeTopDomainsByFlowResponseBody) SetEndTime(v string) *DescribeTopDomainsByFlowResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeTopDomainsByFlowResponseBody) SetRequestId(v string) *DescribeTopDomainsByFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTopDomainsByFlowResponseBody) SetStartTime(v string) *DescribeTopDomainsByFlowResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeTopDomainsByFlowResponseBody) SetTopDomains(v *DescribeTopDomainsByFlowResponseBodyTopDomains) *DescribeTopDomainsByFlowResponseBody {
	s.TopDomains = v
	return s
}

type DescribeTopDomainsByFlowResponseBodyTopDomains struct {
	TopDomain []*DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain `json:"TopDomain,omitempty" xml:"TopDomain,omitempty" type:"Repeated"`
}

func (s DescribeTopDomainsByFlowResponseBodyTopDomains) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopDomainsByFlowResponseBodyTopDomains) GoString() string {
	return s.String()
}

func (s *DescribeTopDomainsByFlowResponseBodyTopDomains) SetTopDomain(v []*DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain) *DescribeTopDomainsByFlowResponseBodyTopDomains {
	s.TopDomain = v
	return s
}

type DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain struct {
	DomainName     *string  `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	MaxBps         *float32 `json:"MaxBps,omitempty" xml:"MaxBps,omitempty"`
	MaxBpsTime     *string  `json:"MaxBpsTime,omitempty" xml:"MaxBpsTime,omitempty"`
	Rank           *int64   `json:"Rank,omitempty" xml:"Rank,omitempty"`
	TotalAccess    *int64   `json:"TotalAccess,omitempty" xml:"TotalAccess,omitempty"`
	TotalTraffic   *string  `json:"TotalTraffic,omitempty" xml:"TotalTraffic,omitempty"`
	TrafficPercent *string  `json:"TrafficPercent,omitempty" xml:"TrafficPercent,omitempty"`
}

func (s DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain) GoString() string {
	return s.String()
}

func (s *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetDomainName(v string) *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.DomainName = &v
	return s
}

func (s *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetMaxBps(v float32) *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.MaxBps = &v
	return s
}

func (s *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetMaxBpsTime(v string) *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.MaxBpsTime = &v
	return s
}

func (s *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetRank(v int64) *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.Rank = &v
	return s
}

func (s *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetTotalAccess(v int64) *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.TotalAccess = &v
	return s
}

func (s *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetTotalTraffic(v string) *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.TotalTraffic = &v
	return s
}

func (s *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetTrafficPercent(v string) *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.TrafficPercent = &v
	return s
}

type DescribeTopDomainsByFlowResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeTopDomainsByFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTopDomainsByFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopDomainsByFlowResponse) GoString() string {
	return s.String()
}

func (s *DescribeTopDomainsByFlowResponse) SetHeaders(v map[string]*string) *DescribeTopDomainsByFlowResponse {
	s.Headers = v
	return s
}

func (s *DescribeTopDomainsByFlowResponse) SetStatusCode(v int32) *DescribeTopDomainsByFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTopDomainsByFlowResponse) SetBody(v *DescribeTopDomainsByFlowResponseBody) *DescribeTopDomainsByFlowResponse {
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

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
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

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
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

func (client *Client) DescribeDomainFileSizeProportionDataWithOptions(request *DescribeDomainFileSizeProportionDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainFileSizeProportionDataResponse, _err error) {
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
		Action:      tea.String("DescribeDomainFileSizeProportionData"),
		Version:     tea.String("2014-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDomainFileSizeProportionDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainFileSizeProportionData(request *DescribeDomainFileSizeProportionDataRequest) (_result *DescribeDomainFileSizeProportionDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainFileSizeProportionDataResponse{}
	_body, _err := client.DescribeDomainFileSizeProportionDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainFlowDataWithOptions(request *DescribeDomainFlowDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainFlowDataResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
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
		Action:      tea.String("DescribeDomainFlowData"),
		Version:     tea.String("2014-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDomainFlowDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainFlowData(request *DescribeDomainFlowDataRequest) (_result *DescribeDomainFlowDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainFlowDataResponse{}
	_body, _err := client.DescribeDomainFlowDataWithOptions(request, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
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

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
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

func (client *Client) DescribeDomainISPDataWithOptions(request *DescribeDomainISPDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainISPDataResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomainISPData"),
		Version:     tea.String("2014-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDomainISPDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainISPData(request *DescribeDomainISPDataRequest) (_result *DescribeDomainISPDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainISPDataResponse{}
	_body, _err := client.DescribeDomainISPDataWithOptions(request, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
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

func (client *Client) DescribeDomainRegionDataWithOptions(request *DescribeDomainRegionDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainRegionDataResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomainRegionData"),
		Version:     tea.String("2014-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDomainRegionDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainRegionData(request *DescribeDomainRegionDataRequest) (_result *DescribeDomainRegionDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainRegionDataResponse{}
	_body, _err := client.DescribeDomainRegionDataWithOptions(request, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
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

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
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

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
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

func (client *Client) DescribeDomainsBySourceWithOptions(request *DescribeDomainsBySourceRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainsBySourceResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Sources)) {
		query["Sources"] = request.Sources
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomainsBySource"),
		Version:     tea.String("2014-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDomainsBySourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainsBySource(request *DescribeDomainsBySourceRequest) (_result *DescribeDomainsBySourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainsBySourceResponse{}
	_body, _err := client.DescribeDomainsBySourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainsUsageByDayWithOptions(request *DescribeDomainsUsageByDayRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainsUsageByDayResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomainsUsageByDay"),
		Version:     tea.String("2014-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDomainsUsageByDayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainsUsageByDay(request *DescribeDomainsUsageByDayRequest) (_result *DescribeDomainsUsageByDayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainsUsageByDayResponse{}
	_body, _err := client.DescribeDomainsUsageByDayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRefreshQuotaWithOptions(request *DescribeRefreshQuotaRequest, runtime *util.RuntimeOptions) (_result *DescribeRefreshQuotaResponse, _err error) {
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
		Action:      tea.String("DescribeRefreshQuota"),
		Version:     tea.String("2014-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRefreshQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRefreshQuota(request *DescribeRefreshQuotaRequest) (_result *DescribeRefreshQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRefreshQuotaResponse{}
	_body, _err := client.DescribeRefreshQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTopDomainsByFlowWithOptions(request *DescribeTopDomainsByFlowRequest, runtime *util.RuntimeOptions) (_result *DescribeTopDomainsByFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		query["Product"] = request.Product
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTopDomainsByFlow"),
		Version:     tea.String("2014-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTopDomainsByFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTopDomainsByFlow(request *DescribeTopDomainsByFlowRequest) (_result *DescribeTopDomainsByFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTopDomainsByFlowResponse{}
	_body, _err := client.DescribeTopDomainsByFlowWithOptions(request, runtime)
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
