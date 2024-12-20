// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddDomainRequest struct {
	// example:
	//
	// 123456
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s AddDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s AddDomainRequest) GoString() string {
	return s.String()
}

func (s *AddDomainRequest) SetAccountId(v string) *AddDomainRequest {
	s.AccountId = &v
	return s
}

func (s *AddDomainRequest) SetDomainName(v string) *AddDomainRequest {
	s.DomainName = &v
	return s
}

type AddDomainResponseBody struct {
	// example:
	//
	// www.example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// ADA27798-6911-4B06-AF34-53F62F62XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddDomainResponseBody) GoString() string {
	return s.String()
}

func (s *AddDomainResponseBody) SetDomainName(v string) *AddDomainResponseBody {
	s.DomainName = &v
	return s
}

func (s *AddDomainResponseBody) SetRequestId(v string) *AddDomainResponseBody {
	s.RequestId = &v
	return s
}

type AddDomainResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s AddDomainResponse) GoString() string {
	return s.String()
}

func (s *AddDomainResponse) SetHeaders(v map[string]*string) *AddDomainResponse {
	s.Headers = v
	return s
}

func (s *AddDomainResponse) SetStatusCode(v int32) *AddDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDomainResponse) SetBody(v *AddDomainResponseBody) *AddDomainResponse {
	s.Body = v
	return s
}

type DeleteDomainRequest struct {
	// example:
	//
	// 12****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DeleteDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainRequest) GoString() string {
	return s.String()
}

func (s *DeleteDomainRequest) SetAccountId(v string) *DeleteDomainRequest {
	s.AccountId = &v
	return s
}

func (s *DeleteDomainRequest) SetDomainName(v string) *DeleteDomainRequest {
	s.DomainName = &v
	return s
}

type DeleteDomainResponseBody struct {
	// example:
	//
	// www.example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// FA8C2599-362D-4113-8FB4-E88A40C2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDomainResponseBody) SetDomainName(v string) *DeleteDomainResponseBody {
	s.DomainName = &v
	return s
}

func (s *DeleteDomainResponseBody) SetRequestId(v string) *DeleteDomainResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDomainResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainResponse) GoString() string {
	return s.String()
}

func (s *DeleteDomainResponse) SetHeaders(v map[string]*string) *DeleteDomainResponse {
	s.Headers = v
	return s
}

func (s *DeleteDomainResponse) SetStatusCode(v int32) *DeleteDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDomainResponse) SetBody(v *DeleteDomainResponseBody) *DeleteDomainResponse {
	s.Body = v
	return s
}

type DescribeDomainsRequest struct {
	// example:
	//
	// 123456
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainsRequest) SetAccountId(v string) *DescribeDomainsRequest {
	s.AccountId = &v
	return s
}

func (s *DescribeDomainsRequest) SetPageNumber(v int64) *DescribeDomainsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDomainsRequest) SetPageSize(v int64) *DescribeDomainsRequest {
	s.PageSize = &v
	return s
}

type DescribeDomainsResponseBody struct {
	Domains *DescribeDomainsResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// A6B3BB61-69CB-50E0-9DC0-0C1658D44A47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 12
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBody) SetDomains(v *DescribeDomainsResponseBodyDomains) *DescribeDomainsResponseBody {
	s.Domains = v
	return s
}

func (s *DescribeDomainsResponseBody) SetPageNumber(v int64) *DescribeDomainsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDomainsResponseBody) SetPageSize(v int64) *DescribeDomainsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainsResponseBody) SetRequestId(v string) *DescribeDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainsResponseBody) SetTotalCount(v int64) *DescribeDomainsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDomainsResponseBodyDomains struct {
	Domain []*DescribeDomainsResponseBodyDomainsDomain `json:"Domain,omitempty" xml:"Domain,omitempty" type:"Repeated"`
}

func (s DescribeDomainsResponseBodyDomains) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBodyDomains) SetDomain(v []*DescribeDomainsResponseBodyDomainsDomain) *DescribeDomainsResponseBodyDomains {
	s.Domain = v
	return s
}

type DescribeDomainsResponseBodyDomainsDomain struct {
	// example:
	//
	// www.aliyun.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeDomainsResponseBodyDomainsDomain) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsResponseBodyDomainsDomain) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBodyDomainsDomain) SetDomainName(v string) *DescribeDomainsResponseBodyDomainsDomain {
	s.DomainName = &v
	return s
}

type DescribeDomainsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponse) SetHeaders(v map[string]*string) *DescribeDomainsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainsResponse) SetStatusCode(v int32) *DescribeDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainsResponse) SetBody(v *DescribeDomainsResponseBody) *DescribeDomainsResponse {
	s.Body = v
	return s
}

type GetAccountInfoResponseBody struct {
	AccountInfo *GetAccountInfoResponseBodyAccountInfo `json:"AccountInfo,omitempty" xml:"AccountInfo,omitempty" type:"Struct"`
	// example:
	//
	// 50F9C40E-188D-4208-BE2C-74271337****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAccountInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccountInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccountInfoResponseBody) SetAccountInfo(v *GetAccountInfoResponseBodyAccountInfo) *GetAccountInfoResponseBody {
	s.AccountInfo = v
	return s
}

func (s *GetAccountInfoResponseBody) SetRequestId(v string) *GetAccountInfoResponseBody {
	s.RequestId = &v
	return s
}

type GetAccountInfoResponseBodyAccountInfo struct {
	// example:
	//
	// 1337****
	AccountId            *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	DohEnabled           *bool   `json:"DohEnabled,omitempty" xml:"DohEnabled,omitempty"`
	DohResolveAllEnabled *bool   `json:"DohResolveAllEnabled,omitempty" xml:"DohResolveAllEnabled,omitempty"`
	MonthDohResolveCount *int64  `json:"MonthDohResolveCount,omitempty" xml:"MonthDohResolveCount,omitempty"`
	// example:
	//
	// 1500000
	MonthFreeCount *int32 `json:"MonthFreeCount,omitempty" xml:"MonthFreeCount,omitempty"`
	// example:
	//
	// 3
	MonthHttpsResolveCount *int32 `json:"MonthHttpsResolveCount,omitempty" xml:"MonthHttpsResolveCount,omitempty"`
	// example:
	//
	// 9927326
	MonthResolveCount *int32 `json:"MonthResolveCount,omitempty" xml:"MonthResolveCount,omitempty"`
	// example:
	//
	// 0
	PackageCount *int32 `json:"PackageCount,omitempty" xml:"PackageCount,omitempty"`
	// example:
	//
	// 50F9C40E****
	SignSecret *string `json:"SignSecret,omitempty" xml:"SignSecret,omitempty"`
	// example:
	//
	// 611523
	SignedCount *int64 `json:"SignedCount,omitempty" xml:"SignedCount,omitempty"`
	// example:
	//
	// 1523
	UnsignedCount *int64 `json:"UnsignedCount,omitempty" xml:"UnsignedCount,omitempty"`
	// example:
	//
	// true
	UnsignedEnabled *bool `json:"UnsignedEnabled,omitempty" xml:"UnsignedEnabled,omitempty"`
	// example:
	//
	// 1
	UserStatus *int32 `json:"UserStatus,omitempty" xml:"UserStatus,omitempty"`
}

func (s GetAccountInfoResponseBodyAccountInfo) String() string {
	return tea.Prettify(s)
}

func (s GetAccountInfoResponseBodyAccountInfo) GoString() string {
	return s.String()
}

func (s *GetAccountInfoResponseBodyAccountInfo) SetAccountId(v string) *GetAccountInfoResponseBodyAccountInfo {
	s.AccountId = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfo) SetDohEnabled(v bool) *GetAccountInfoResponseBodyAccountInfo {
	s.DohEnabled = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfo) SetDohResolveAllEnabled(v bool) *GetAccountInfoResponseBodyAccountInfo {
	s.DohResolveAllEnabled = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfo) SetMonthDohResolveCount(v int64) *GetAccountInfoResponseBodyAccountInfo {
	s.MonthDohResolveCount = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfo) SetMonthFreeCount(v int32) *GetAccountInfoResponseBodyAccountInfo {
	s.MonthFreeCount = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfo) SetMonthHttpsResolveCount(v int32) *GetAccountInfoResponseBodyAccountInfo {
	s.MonthHttpsResolveCount = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfo) SetMonthResolveCount(v int32) *GetAccountInfoResponseBodyAccountInfo {
	s.MonthResolveCount = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfo) SetPackageCount(v int32) *GetAccountInfoResponseBodyAccountInfo {
	s.PackageCount = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfo) SetSignSecret(v string) *GetAccountInfoResponseBodyAccountInfo {
	s.SignSecret = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfo) SetSignedCount(v int64) *GetAccountInfoResponseBodyAccountInfo {
	s.SignedCount = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfo) SetUnsignedCount(v int64) *GetAccountInfoResponseBodyAccountInfo {
	s.UnsignedCount = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfo) SetUnsignedEnabled(v bool) *GetAccountInfoResponseBodyAccountInfo {
	s.UnsignedEnabled = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfo) SetUserStatus(v int32) *GetAccountInfoResponseBodyAccountInfo {
	s.UserStatus = &v
	return s
}

type GetAccountInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccountInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccountInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccountInfoResponse) GoString() string {
	return s.String()
}

func (s *GetAccountInfoResponse) SetHeaders(v map[string]*string) *GetAccountInfoResponse {
	s.Headers = v
	return s
}

func (s *GetAccountInfoResponse) SetStatusCode(v int32) *GetAccountInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccountInfoResponse) SetBody(v *GetAccountInfoResponseBody) *GetAccountInfoResponse {
	s.Body = v
	return s
}

type GetResolveCountSummaryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// day
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7
	TimeSpan *int32 `json:"TimeSpan,omitempty" xml:"TimeSpan,omitempty"`
}

func (s GetResolveCountSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResolveCountSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetResolveCountSummaryRequest) SetGranularity(v string) *GetResolveCountSummaryRequest {
	s.Granularity = &v
	return s
}

func (s *GetResolveCountSummaryRequest) SetTimeSpan(v int32) *GetResolveCountSummaryRequest {
	s.TimeSpan = &v
	return s
}

type GetResolveCountSummaryResponseBody struct {
	// example:
	//
	// 3106FFF3-3612-542A-B2CD-3CF4CD48****
	RequestId      *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResolveSummary *GetResolveCountSummaryResponseBodyResolveSummary `json:"ResolveSummary,omitempty" xml:"ResolveSummary,omitempty" type:"Struct"`
}

func (s GetResolveCountSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResolveCountSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetResolveCountSummaryResponseBody) SetRequestId(v string) *GetResolveCountSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResolveCountSummaryResponseBody) SetResolveSummary(v *GetResolveCountSummaryResponseBodyResolveSummary) *GetResolveCountSummaryResponseBody {
	s.ResolveSummary = v
	return s
}

type GetResolveCountSummaryResponseBodyResolveSummary struct {
	Doh *int64 `json:"Doh,omitempty" xml:"Doh,omitempty"`
	// example:
	//
	// 123
	Http *int64 `json:"Http,omitempty" xml:"Http,omitempty"`
	// example:
	//
	// 123
	Http6 *int64 `json:"Http6,omitempty" xml:"Http6,omitempty"`
	// example:
	//
	// 123
	Https *int64 `json:"Https,omitempty" xml:"Https,omitempty"`
	// example:
	//
	// 123
	Https6 *int64 `json:"Https6,omitempty" xml:"Https6,omitempty"`
}

func (s GetResolveCountSummaryResponseBodyResolveSummary) String() string {
	return tea.Prettify(s)
}

func (s GetResolveCountSummaryResponseBodyResolveSummary) GoString() string {
	return s.String()
}

func (s *GetResolveCountSummaryResponseBodyResolveSummary) SetDoh(v int64) *GetResolveCountSummaryResponseBodyResolveSummary {
	s.Doh = &v
	return s
}

func (s *GetResolveCountSummaryResponseBodyResolveSummary) SetHttp(v int64) *GetResolveCountSummaryResponseBodyResolveSummary {
	s.Http = &v
	return s
}

func (s *GetResolveCountSummaryResponseBodyResolveSummary) SetHttp6(v int64) *GetResolveCountSummaryResponseBodyResolveSummary {
	s.Http6 = &v
	return s
}

func (s *GetResolveCountSummaryResponseBodyResolveSummary) SetHttps(v int64) *GetResolveCountSummaryResponseBodyResolveSummary {
	s.Https = &v
	return s
}

func (s *GetResolveCountSummaryResponseBodyResolveSummary) SetHttps6(v int64) *GetResolveCountSummaryResponseBodyResolveSummary {
	s.Https6 = &v
	return s
}

type GetResolveCountSummaryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResolveCountSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResolveCountSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResolveCountSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetResolveCountSummaryResponse) SetHeaders(v map[string]*string) *GetResolveCountSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetResolveCountSummaryResponse) SetStatusCode(v int32) *GetResolveCountSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResolveCountSummaryResponse) SetBody(v *GetResolveCountSummaryResponseBody) *GetResolveCountSummaryResponse {
	s.Body = v
	return s
}

type GetResolveStatisticsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// month
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// example:
	//
	// https
	ProtocolName *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	TimeSpan *int32 `json:"TimeSpan,omitempty" xml:"TimeSpan,omitempty"`
}

func (s GetResolveStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResolveStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetResolveStatisticsRequest) SetDomainName(v string) *GetResolveStatisticsRequest {
	s.DomainName = &v
	return s
}

func (s *GetResolveStatisticsRequest) SetGranularity(v string) *GetResolveStatisticsRequest {
	s.Granularity = &v
	return s
}

func (s *GetResolveStatisticsRequest) SetProtocolName(v string) *GetResolveStatisticsRequest {
	s.ProtocolName = &v
	return s
}

func (s *GetResolveStatisticsRequest) SetTimeSpan(v int32) *GetResolveStatisticsRequest {
	s.TimeSpan = &v
	return s
}

type GetResolveStatisticsResponseBody struct {
	DataPoints *GetResolveStatisticsResponseBodyDataPoints `json:"DataPoints,omitempty" xml:"DataPoints,omitempty" type:"Struct"`
	// example:
	//
	// 50F9C40E-188D-B00B-BE2C-7427E531****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetResolveStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResolveStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetResolveStatisticsResponseBody) SetDataPoints(v *GetResolveStatisticsResponseBodyDataPoints) *GetResolveStatisticsResponseBody {
	s.DataPoints = v
	return s
}

func (s *GetResolveStatisticsResponseBody) SetRequestId(v string) *GetResolveStatisticsResponseBody {
	s.RequestId = &v
	return s
}

type GetResolveStatisticsResponseBodyDataPoints struct {
	DataPoint []*GetResolveStatisticsResponseBodyDataPointsDataPoint `json:"DataPoint,omitempty" xml:"DataPoint,omitempty" type:"Repeated"`
}

func (s GetResolveStatisticsResponseBodyDataPoints) String() string {
	return tea.Prettify(s)
}

func (s GetResolveStatisticsResponseBodyDataPoints) GoString() string {
	return s.String()
}

func (s *GetResolveStatisticsResponseBodyDataPoints) SetDataPoint(v []*GetResolveStatisticsResponseBodyDataPointsDataPoint) *GetResolveStatisticsResponseBodyDataPoints {
	s.DataPoint = v
	return s
}

type GetResolveStatisticsResponseBodyDataPointsDataPoint struct {
	// example:
	//
	// 9703
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 1488297600
	Time *int32 `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s GetResolveStatisticsResponseBodyDataPointsDataPoint) String() string {
	return tea.Prettify(s)
}

func (s GetResolveStatisticsResponseBodyDataPointsDataPoint) GoString() string {
	return s.String()
}

func (s *GetResolveStatisticsResponseBodyDataPointsDataPoint) SetCount(v int32) *GetResolveStatisticsResponseBodyDataPointsDataPoint {
	s.Count = &v
	return s
}

func (s *GetResolveStatisticsResponseBodyDataPointsDataPoint) SetTime(v int32) *GetResolveStatisticsResponseBodyDataPointsDataPoint {
	s.Time = &v
	return s
}

type GetResolveStatisticsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResolveStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResolveStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResolveStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetResolveStatisticsResponse) SetHeaders(v map[string]*string) *GetResolveStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetResolveStatisticsResponse) SetStatusCode(v int32) *GetResolveStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResolveStatisticsResponse) SetBody(v *GetResolveStatisticsResponseBody) *GetResolveStatisticsResponse {
	s.Body = v
	return s
}

type ListDomainsRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize            *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Search              *string `json:"Search,omitempty" xml:"Search,omitempty"`
	WithoutMeteringData *bool   `json:"WithoutMeteringData,omitempty" xml:"WithoutMeteringData,omitempty"`
}

func (s ListDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDomainsRequest) GoString() string {
	return s.String()
}

func (s *ListDomainsRequest) SetPageNumber(v int32) *ListDomainsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDomainsRequest) SetPageSize(v int32) *ListDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDomainsRequest) SetSearch(v string) *ListDomainsRequest {
	s.Search = &v
	return s
}

func (s *ListDomainsRequest) SetWithoutMeteringData(v bool) *ListDomainsRequest {
	s.WithoutMeteringData = &v
	return s
}

type ListDomainsResponseBody struct {
	DomainInfos *ListDomainsResponseBodyDomainInfos `json:"DomainInfos,omitempty" xml:"DomainInfos,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 5
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 50F9C40E-188D-4208-BE2C-7427E531****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 18
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDomainsResponseBody) SetDomainInfos(v *ListDomainsResponseBodyDomainInfos) *ListDomainsResponseBody {
	s.DomainInfos = v
	return s
}

func (s *ListDomainsResponseBody) SetPageNumber(v int64) *ListDomainsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDomainsResponseBody) SetPageSize(v int64) *ListDomainsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDomainsResponseBody) SetRequestId(v string) *ListDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDomainsResponseBody) SetTotalCount(v int64) *ListDomainsResponseBody {
	s.TotalCount = &v
	return s
}

type ListDomainsResponseBodyDomainInfos struct {
	DomainInfo []*ListDomainsResponseBodyDomainInfosDomainInfo `json:"DomainInfo,omitempty" xml:"DomainInfo,omitempty" type:"Repeated"`
}

func (s ListDomainsResponseBodyDomainInfos) String() string {
	return tea.Prettify(s)
}

func (s ListDomainsResponseBodyDomainInfos) GoString() string {
	return s.String()
}

func (s *ListDomainsResponseBodyDomainInfos) SetDomainInfo(v []*ListDomainsResponseBodyDomainInfosDomainInfo) *ListDomainsResponseBodyDomainInfos {
	s.DomainInfo = v
	return s
}

type ListDomainsResponseBodyDomainInfosDomainInfo struct {
	// example:
	//
	// www.example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 10
	Resolved *int64 `json:"Resolved,omitempty" xml:"Resolved,omitempty"`
	// example:
	//
	// 20
	Resolved6   *int64 `json:"Resolved6,omitempty" xml:"Resolved6,omitempty"`
	ResolvedDoh *int64 `json:"ResolvedDoh,omitempty" xml:"ResolvedDoh,omitempty"`
	// example:
	//
	// 10
	ResolvedHttps *int64 `json:"ResolvedHttps,omitempty" xml:"ResolvedHttps,omitempty"`
	// example:
	//
	// 20
	ResolvedHttps6 *int64 `json:"ResolvedHttps6,omitempty" xml:"ResolvedHttps6,omitempty"`
	TimeModified   *int64 `json:"TimeModified,omitempty" xml:"TimeModified,omitempty"`
}

func (s ListDomainsResponseBodyDomainInfosDomainInfo) String() string {
	return tea.Prettify(s)
}

func (s ListDomainsResponseBodyDomainInfosDomainInfo) GoString() string {
	return s.String()
}

func (s *ListDomainsResponseBodyDomainInfosDomainInfo) SetDomainName(v string) *ListDomainsResponseBodyDomainInfosDomainInfo {
	s.DomainName = &v
	return s
}

func (s *ListDomainsResponseBodyDomainInfosDomainInfo) SetResolved(v int64) *ListDomainsResponseBodyDomainInfosDomainInfo {
	s.Resolved = &v
	return s
}

func (s *ListDomainsResponseBodyDomainInfosDomainInfo) SetResolved6(v int64) *ListDomainsResponseBodyDomainInfosDomainInfo {
	s.Resolved6 = &v
	return s
}

func (s *ListDomainsResponseBodyDomainInfosDomainInfo) SetResolvedDoh(v int64) *ListDomainsResponseBodyDomainInfosDomainInfo {
	s.ResolvedDoh = &v
	return s
}

func (s *ListDomainsResponseBodyDomainInfosDomainInfo) SetResolvedHttps(v int64) *ListDomainsResponseBodyDomainInfosDomainInfo {
	s.ResolvedHttps = &v
	return s
}

func (s *ListDomainsResponseBodyDomainInfosDomainInfo) SetResolvedHttps6(v int64) *ListDomainsResponseBodyDomainInfosDomainInfo {
	s.ResolvedHttps6 = &v
	return s
}

func (s *ListDomainsResponseBodyDomainInfosDomainInfo) SetTimeModified(v int64) *ListDomainsResponseBodyDomainInfosDomainInfo {
	s.TimeModified = &v
	return s
}

type ListDomainsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDomainsResponse) GoString() string {
	return s.String()
}

func (s *ListDomainsResponse) SetHeaders(v map[string]*string) *ListDomainsResponse {
	s.Headers = v
	return s
}

func (s *ListDomainsResponse) SetStatusCode(v int32) *ListDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDomainsResponse) SetBody(v *ListDomainsResponseBody) *ListDomainsResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("httpdns"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 添加域名
//
// @param request - AddDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDomainResponse
func (client *Client) AddDomainWithOptions(request *AddDomainRequest, runtime *util.RuntimeOptions) (_result *AddDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddDomain"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加域名
//
// @param request - AddDomainRequest
//
// @return AddDomainResponse
func (client *Client) AddDomain(request *AddDomainRequest) (_result *AddDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddDomainResponse{}
	_body, _err := client.AddDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除域名
//
// @param request - DeleteDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDomainResponse
func (client *Client) DeleteDomainWithOptions(request *DeleteDomainRequest, runtime *util.RuntimeOptions) (_result *DeleteDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDomain"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除域名
//
// @param request - DeleteDomainRequest
//
// @return DeleteDomainResponse
func (client *Client) DeleteDomain(request *DeleteDomainRequest) (_result *DeleteDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDomainResponse{}
	_body, _err := client.DeleteDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainsResponse
func (client *Client) DescribeDomainsWithOptions(request *DescribeDomainsRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomains"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeDomainsRequest
//
// @return DescribeDomainsResponse
func (client *Client) DescribeDomains(request *DescribeDomainsRequest) (_result *DescribeDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainsResponse{}
	_body, _err := client.DescribeDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户信息包含配置项
//
// @param request - GetAccountInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccountInfoResponse
func (client *Client) GetAccountInfoWithOptions(runtime *util.RuntimeOptions) (_result *GetAccountInfoResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetAccountInfo"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAccountInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户信息包含配置项
//
// @return GetAccountInfoResponse
func (client *Client) GetAccountInfo() (_result *GetAccountInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAccountInfoResponse{}
	_body, _err := client.GetAccountInfoWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 解析次数概览
//
// @param request - GetResolveCountSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResolveCountSummaryResponse
func (client *Client) GetResolveCountSummaryWithOptions(request *GetResolveCountSummaryRequest, runtime *util.RuntimeOptions) (_result *GetResolveCountSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Granularity)) {
		query["Granularity"] = request.Granularity
	}

	if !tea.BoolValue(util.IsUnset(request.TimeSpan)) {
		query["TimeSpan"] = request.TimeSpan
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResolveCountSummary"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResolveCountSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解析次数概览
//
// @param request - GetResolveCountSummaryRequest
//
// @return GetResolveCountSummaryResponse
func (client *Client) GetResolveCountSummary(request *GetResolveCountSummaryRequest) (_result *GetResolveCountSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResolveCountSummaryResponse{}
	_body, _err := client.GetResolveCountSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetResolveStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResolveStatisticsResponse
func (client *Client) GetResolveStatisticsWithOptions(request *GetResolveStatisticsRequest, runtime *util.RuntimeOptions) (_result *GetResolveStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Granularity)) {
		query["Granularity"] = request.Granularity
	}

	if !tea.BoolValue(util.IsUnset(request.ProtocolName)) {
		query["ProtocolName"] = request.ProtocolName
	}

	if !tea.BoolValue(util.IsUnset(request.TimeSpan)) {
		query["TimeSpan"] = request.TimeSpan
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResolveStatistics"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResolveStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetResolveStatisticsRequest
//
// @return GetResolveStatisticsResponse
func (client *Client) GetResolveStatistics(request *GetResolveStatisticsRequest) (_result *GetResolveStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResolveStatisticsResponse{}
	_body, _err := client.GetResolveStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列举域名以及解析统计信息
//
// @param request - ListDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDomainsResponse
func (client *Client) ListDomainsWithOptions(request *ListDomainsRequest, runtime *util.RuntimeOptions) (_result *ListDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Search)) {
		query["Search"] = request.Search
	}

	if !tea.BoolValue(util.IsUnset(request.WithoutMeteringData)) {
		query["WithoutMeteringData"] = request.WithoutMeteringData
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDomains"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举域名以及解析统计信息
//
// @param request - ListDomainsRequest
//
// @return ListDomainsResponse
func (client *Client) ListDomains(request *ListDomainsRequest) (_result *ListDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDomainsResponse{}
	_body, _err := client.ListDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
