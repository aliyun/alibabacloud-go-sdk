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

type DescribeBgpPackByIpRequest struct {
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
	Ip           *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribeBgpPackByIpRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBgpPackByIpRequest) GoString() string {
	return s.String()
}

func (s *DescribeBgpPackByIpRequest) SetDdosRegionId(v string) *DescribeBgpPackByIpRequest {
	s.DdosRegionId = &v
	return s
}

func (s *DescribeBgpPackByIpRequest) SetIp(v string) *DescribeBgpPackByIpRequest {
	s.Ip = &v
	return s
}

type DescribeBgpPackByIpResponseBody struct {
	DdosbgpInfo *DescribeBgpPackByIpResponseBodyDdosbgpInfo `json:"DdosbgpInfo,omitempty" xml:"DdosbgpInfo,omitempty" type:"Struct"`
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBgpPackByIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBgpPackByIpResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBgpPackByIpResponseBody) SetDdosbgpInfo(v *DescribeBgpPackByIpResponseBodyDdosbgpInfo) *DescribeBgpPackByIpResponseBody {
	s.DdosbgpInfo = v
	return s
}

func (s *DescribeBgpPackByIpResponseBody) SetRequestId(v string) *DescribeBgpPackByIpResponseBody {
	s.RequestId = &v
	return s
}

type DescribeBgpPackByIpResponseBodyDdosbgpInfo struct {
	BaseThreshold     *int32  `json:"BaseThreshold,omitempty" xml:"BaseThreshold,omitempty"`
	DdosbgpInstanceId *string `json:"DdosbgpInstanceId,omitempty" xml:"DdosbgpInstanceId,omitempty"`
	ElasticThreshold  *int32  `json:"ElasticThreshold,omitempty" xml:"ElasticThreshold,omitempty"`
	ExpireTime        *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Ip                *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribeBgpPackByIpResponseBodyDdosbgpInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeBgpPackByIpResponseBodyDdosbgpInfo) GoString() string {
	return s.String()
}

func (s *DescribeBgpPackByIpResponseBodyDdosbgpInfo) SetBaseThreshold(v int32) *DescribeBgpPackByIpResponseBodyDdosbgpInfo {
	s.BaseThreshold = &v
	return s
}

func (s *DescribeBgpPackByIpResponseBodyDdosbgpInfo) SetDdosbgpInstanceId(v string) *DescribeBgpPackByIpResponseBodyDdosbgpInfo {
	s.DdosbgpInstanceId = &v
	return s
}

func (s *DescribeBgpPackByIpResponseBodyDdosbgpInfo) SetElasticThreshold(v int32) *DescribeBgpPackByIpResponseBodyDdosbgpInfo {
	s.ElasticThreshold = &v
	return s
}

func (s *DescribeBgpPackByIpResponseBodyDdosbgpInfo) SetExpireTime(v int64) *DescribeBgpPackByIpResponseBodyDdosbgpInfo {
	s.ExpireTime = &v
	return s
}

func (s *DescribeBgpPackByIpResponseBodyDdosbgpInfo) SetIp(v string) *DescribeBgpPackByIpResponseBodyDdosbgpInfo {
	s.Ip = &v
	return s
}

type DescribeBgpPackByIpResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeBgpPackByIpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBgpPackByIpResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBgpPackByIpResponse) GoString() string {
	return s.String()
}

func (s *DescribeBgpPackByIpResponse) SetHeaders(v map[string]*string) *DescribeBgpPackByIpResponse {
	s.Headers = v
	return s
}

func (s *DescribeBgpPackByIpResponse) SetStatusCode(v int32) *DescribeBgpPackByIpResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBgpPackByIpResponse) SetBody(v *DescribeBgpPackByIpResponseBody) *DescribeBgpPackByIpResponse {
	s.Body = v
	return s
}

type DescribeCapRequest struct {
	BegTime      *int64  `json:"BegTime,omitempty" xml:"BegTime,omitempty"`
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InternetIp   *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
}

func (s DescribeCapRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCapRequest) GoString() string {
	return s.String()
}

func (s *DescribeCapRequest) SetBegTime(v int64) *DescribeCapRequest {
	s.BegTime = &v
	return s
}

func (s *DescribeCapRequest) SetDdosRegionId(v string) *DescribeCapRequest {
	s.DdosRegionId = &v
	return s
}

func (s *DescribeCapRequest) SetInstanceId(v string) *DescribeCapRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeCapRequest) SetInstanceType(v string) *DescribeCapRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeCapRequest) SetInternetIp(v string) *DescribeCapRequest {
	s.InternetIp = &v
	return s
}

type DescribeCapResponseBody struct {
	CapUrl    *DescribeCapResponseBodyCapUrl `json:"CapUrl,omitempty" xml:"CapUrl,omitempty" type:"Struct"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCapResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCapResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCapResponseBody) SetCapUrl(v *DescribeCapResponseBodyCapUrl) *DescribeCapResponseBody {
	s.CapUrl = v
	return s
}

func (s *DescribeCapResponseBody) SetRequestId(v string) *DescribeCapResponseBody {
	s.RequestId = &v
	return s
}

type DescribeCapResponseBodyCapUrl struct {
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeCapResponseBodyCapUrl) String() string {
	return tea.Prettify(s)
}

func (s DescribeCapResponseBodyCapUrl) GoString() string {
	return s.String()
}

func (s *DescribeCapResponseBodyCapUrl) SetUrl(v string) *DescribeCapResponseBodyCapUrl {
	s.Url = &v
	return s
}

type DescribeCapResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeCapResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCapResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCapResponse) GoString() string {
	return s.String()
}

func (s *DescribeCapResponse) SetHeaders(v map[string]*string) *DescribeCapResponse {
	s.Headers = v
	return s
}

func (s *DescribeCapResponse) SetStatusCode(v int32) *DescribeCapResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCapResponse) SetBody(v *DescribeCapResponseBody) *DescribeCapResponse {
	s.Body = v
	return s
}

type DescribeDdosCountRequest struct {
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeDdosCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeDdosCountRequest) SetDdosRegionId(v string) *DescribeDdosCountRequest {
	s.DdosRegionId = &v
	return s
}

func (s *DescribeDdosCountRequest) SetInstanceType(v string) *DescribeDdosCountRequest {
	s.InstanceType = &v
	return s
}

type DescribeDdosCountResponseBody struct {
	DdosCount *DescribeDdosCountResponseBodyDdosCount `json:"DdosCount,omitempty" xml:"DdosCount,omitempty" type:"Struct"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDdosCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDdosCountResponseBody) SetDdosCount(v *DescribeDdosCountResponseBodyDdosCount) *DescribeDdosCountResponseBody {
	s.DdosCount = v
	return s
}

func (s *DescribeDdosCountResponseBody) SetRequestId(v string) *DescribeDdosCountResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDdosCountResponseBodyDdosCount struct {
	BlackholeCount *int32 `json:"BlackholeCount,omitempty" xml:"BlackholeCount,omitempty"`
	DefenseCount   *int32 `json:"DefenseCount,omitempty" xml:"DefenseCount,omitempty"`
	InstacenCount  *int32 `json:"InstacenCount,omitempty" xml:"InstacenCount,omitempty"`
}

func (s DescribeDdosCountResponseBodyDdosCount) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosCountResponseBodyDdosCount) GoString() string {
	return s.String()
}

func (s *DescribeDdosCountResponseBodyDdosCount) SetBlackholeCount(v int32) *DescribeDdosCountResponseBodyDdosCount {
	s.BlackholeCount = &v
	return s
}

func (s *DescribeDdosCountResponseBodyDdosCount) SetDefenseCount(v int32) *DescribeDdosCountResponseBodyDdosCount {
	s.DefenseCount = &v
	return s
}

func (s *DescribeDdosCountResponseBodyDdosCount) SetInstacenCount(v int32) *DescribeDdosCountResponseBodyDdosCount {
	s.InstacenCount = &v
	return s
}

type DescribeDdosCountResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDdosCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDdosCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeDdosCountResponse) SetHeaders(v map[string]*string) *DescribeDdosCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeDdosCountResponse) SetStatusCode(v int32) *DescribeDdosCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDdosCountResponse) SetBody(v *DescribeDdosCountResponseBody) *DescribeDdosCountResponse {
	s.Body = v
	return s
}

type DescribeDdosCreditRequest struct {
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
}

func (s DescribeDdosCreditRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosCreditRequest) GoString() string {
	return s.String()
}

func (s *DescribeDdosCreditRequest) SetDdosRegionId(v string) *DescribeDdosCreditRequest {
	s.DdosRegionId = &v
	return s
}

type DescribeDdosCreditResponseBody struct {
	DdosCredit *DescribeDdosCreditResponseBodyDdosCredit `json:"DdosCredit,omitempty" xml:"DdosCredit,omitempty" type:"Struct"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDdosCreditResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosCreditResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDdosCreditResponseBody) SetDdosCredit(v *DescribeDdosCreditResponseBodyDdosCredit) *DescribeDdosCreditResponseBody {
	s.DdosCredit = v
	return s
}

func (s *DescribeDdosCreditResponseBody) SetRequestId(v string) *DescribeDdosCreditResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDdosCreditResponseBody) SetSuccess(v bool) *DescribeDdosCreditResponseBody {
	s.Success = &v
	return s
}

type DescribeDdosCreditResponseBodyDdosCredit struct {
	BlackholeTime *int32  `json:"BlackholeTime,omitempty" xml:"BlackholeTime,omitempty"`
	Score         *int32  `json:"Score,omitempty" xml:"Score,omitempty"`
	ScoreLevel    *string `json:"ScoreLevel,omitempty" xml:"ScoreLevel,omitempty"`
}

func (s DescribeDdosCreditResponseBodyDdosCredit) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosCreditResponseBodyDdosCredit) GoString() string {
	return s.String()
}

func (s *DescribeDdosCreditResponseBodyDdosCredit) SetBlackholeTime(v int32) *DescribeDdosCreditResponseBodyDdosCredit {
	s.BlackholeTime = &v
	return s
}

func (s *DescribeDdosCreditResponseBodyDdosCredit) SetScore(v int32) *DescribeDdosCreditResponseBodyDdosCredit {
	s.Score = &v
	return s
}

func (s *DescribeDdosCreditResponseBodyDdosCredit) SetScoreLevel(v string) *DescribeDdosCreditResponseBodyDdosCredit {
	s.ScoreLevel = &v
	return s
}

type DescribeDdosCreditResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDdosCreditResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDdosCreditResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosCreditResponse) GoString() string {
	return s.String()
}

func (s *DescribeDdosCreditResponse) SetHeaders(v map[string]*string) *DescribeDdosCreditResponse {
	s.Headers = v
	return s
}

func (s *DescribeDdosCreditResponse) SetStatusCode(v int32) *DescribeDdosCreditResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDdosCreditResponse) SetBody(v *DescribeDdosCreditResponseBody) *DescribeDdosCreditResponse {
	s.Body = v
	return s
}

type DescribeDdosEventListRequest struct {
	CurrentPage  *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InternetIp   *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDdosEventListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosEventListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventListRequest) SetCurrentPage(v int32) *DescribeDdosEventListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDdosEventListRequest) SetDdosRegionId(v string) *DescribeDdosEventListRequest {
	s.DdosRegionId = &v
	return s
}

func (s *DescribeDdosEventListRequest) SetInstanceId(v string) *DescribeDdosEventListRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDdosEventListRequest) SetInstanceType(v string) *DescribeDdosEventListRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeDdosEventListRequest) SetInternetIp(v string) *DescribeDdosEventListRequest {
	s.InternetIp = &v
	return s
}

func (s *DescribeDdosEventListRequest) SetPageSize(v int32) *DescribeDdosEventListRequest {
	s.PageSize = &v
	return s
}

type DescribeDdosEventListResponseBody struct {
	DdosEventList *DescribeDdosEventListResponseBodyDdosEventList `json:"DdosEventList,omitempty" xml:"DdosEventList,omitempty" type:"Struct"`
	RequestId     *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total         *int32                                          `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDdosEventListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosEventListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventListResponseBody) SetDdosEventList(v *DescribeDdosEventListResponseBodyDdosEventList) *DescribeDdosEventListResponseBody {
	s.DdosEventList = v
	return s
}

func (s *DescribeDdosEventListResponseBody) SetRequestId(v string) *DescribeDdosEventListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDdosEventListResponseBody) SetTotal(v int32) *DescribeDdosEventListResponseBody {
	s.Total = &v
	return s
}

type DescribeDdosEventListResponseBodyDdosEventList struct {
	DdosEvent []*DescribeDdosEventListResponseBodyDdosEventListDdosEvent `json:"DdosEvent,omitempty" xml:"DdosEvent,omitempty" type:"Repeated"`
}

func (s DescribeDdosEventListResponseBodyDdosEventList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosEventListResponseBodyDdosEventList) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventListResponseBodyDdosEventList) SetDdosEvent(v []*DescribeDdosEventListResponseBodyDdosEventListDdosEvent) *DescribeDdosEventListResponseBodyDdosEventList {
	s.DdosEvent = v
	return s
}

type DescribeDdosEventListResponseBodyDdosEventListDdosEvent struct {
	DdosStatus      *string `json:"DdosStatus,omitempty" xml:"DdosStatus,omitempty"`
	DdosType        *string `json:"DdosType,omitempty" xml:"DdosType,omitempty"`
	DelayTime       *int64  `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	UnBlackholeTime *int64  `json:"UnBlackholeTime,omitempty" xml:"UnBlackholeTime,omitempty"`
}

func (s DescribeDdosEventListResponseBodyDdosEventListDdosEvent) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosEventListResponseBodyDdosEventListDdosEvent) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventListResponseBodyDdosEventListDdosEvent) SetDdosStatus(v string) *DescribeDdosEventListResponseBodyDdosEventListDdosEvent {
	s.DdosStatus = &v
	return s
}

func (s *DescribeDdosEventListResponseBodyDdosEventListDdosEvent) SetDdosType(v string) *DescribeDdosEventListResponseBodyDdosEventListDdosEvent {
	s.DdosType = &v
	return s
}

func (s *DescribeDdosEventListResponseBodyDdosEventListDdosEvent) SetDelayTime(v int64) *DescribeDdosEventListResponseBodyDdosEventListDdosEvent {
	s.DelayTime = &v
	return s
}

func (s *DescribeDdosEventListResponseBodyDdosEventListDdosEvent) SetEndTime(v int64) *DescribeDdosEventListResponseBodyDdosEventListDdosEvent {
	s.EndTime = &v
	return s
}

func (s *DescribeDdosEventListResponseBodyDdosEventListDdosEvent) SetStartTime(v int64) *DescribeDdosEventListResponseBodyDdosEventListDdosEvent {
	s.StartTime = &v
	return s
}

func (s *DescribeDdosEventListResponseBodyDdosEventListDdosEvent) SetUnBlackholeTime(v int64) *DescribeDdosEventListResponseBodyDdosEventListDdosEvent {
	s.UnBlackholeTime = &v
	return s
}

type DescribeDdosEventListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDdosEventListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDdosEventListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosEventListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventListResponse) SetHeaders(v map[string]*string) *DescribeDdosEventListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDdosEventListResponse) SetStatusCode(v int32) *DescribeDdosEventListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDdosEventListResponse) SetBody(v *DescribeDdosEventListResponseBody) *DescribeDdosEventListResponse {
	s.Body = v
	return s
}

type DescribeDdosThresholdRequest struct {
	DdosRegionId *string   `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
	DdosType     *string   `json:"DdosType,omitempty" xml:"DdosType,omitempty"`
	InstanceIds  []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	InstanceType *string   `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeDdosThresholdRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosThresholdRequest) GoString() string {
	return s.String()
}

func (s *DescribeDdosThresholdRequest) SetDdosRegionId(v string) *DescribeDdosThresholdRequest {
	s.DdosRegionId = &v
	return s
}

func (s *DescribeDdosThresholdRequest) SetDdosType(v string) *DescribeDdosThresholdRequest {
	s.DdosType = &v
	return s
}

func (s *DescribeDdosThresholdRequest) SetInstanceIds(v []*string) *DescribeDdosThresholdRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeDdosThresholdRequest) SetInstanceType(v string) *DescribeDdosThresholdRequest {
	s.InstanceType = &v
	return s
}

type DescribeDdosThresholdResponseBody struct {
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Thresholds *DescribeDdosThresholdResponseBodyThresholds `json:"Thresholds,omitempty" xml:"Thresholds,omitempty" type:"Struct"`
}

func (s DescribeDdosThresholdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosThresholdResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDdosThresholdResponseBody) SetRequestId(v string) *DescribeDdosThresholdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDdosThresholdResponseBody) SetThresholds(v *DescribeDdosThresholdResponseBodyThresholds) *DescribeDdosThresholdResponseBody {
	s.Thresholds = v
	return s
}

type DescribeDdosThresholdResponseBodyThresholds struct {
	Threshold []*DescribeDdosThresholdResponseBodyThresholdsThreshold `json:"Threshold,omitempty" xml:"Threshold,omitempty" type:"Repeated"`
}

func (s DescribeDdosThresholdResponseBodyThresholds) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosThresholdResponseBodyThresholds) GoString() string {
	return s.String()
}

func (s *DescribeDdosThresholdResponseBodyThresholds) SetThreshold(v []*DescribeDdosThresholdResponseBodyThresholdsThreshold) *DescribeDdosThresholdResponseBodyThresholds {
	s.Threshold = v
	return s
}

type DescribeDdosThresholdResponseBodyThresholdsThreshold struct {
	Bps        *int32  `json:"Bps,omitempty" xml:"Bps,omitempty"`
	DdosType   *string `json:"DdosType,omitempty" xml:"DdosType,omitempty"`
	ElasticBps *int32  `json:"ElasticBps,omitempty" xml:"ElasticBps,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	IsAuto     *bool   `json:"IsAuto,omitempty" xml:"IsAuto,omitempty"`
	MaxBps     *int32  `json:"MaxBps,omitempty" xml:"MaxBps,omitempty"`
	MaxPps     *int32  `json:"MaxPps,omitempty" xml:"MaxPps,omitempty"`
	Pps        *int32  `json:"Pps,omitempty" xml:"Pps,omitempty"`
}

func (s DescribeDdosThresholdResponseBodyThresholdsThreshold) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosThresholdResponseBodyThresholdsThreshold) GoString() string {
	return s.String()
}

func (s *DescribeDdosThresholdResponseBodyThresholdsThreshold) SetBps(v int32) *DescribeDdosThresholdResponseBodyThresholdsThreshold {
	s.Bps = &v
	return s
}

func (s *DescribeDdosThresholdResponseBodyThresholdsThreshold) SetDdosType(v string) *DescribeDdosThresholdResponseBodyThresholdsThreshold {
	s.DdosType = &v
	return s
}

func (s *DescribeDdosThresholdResponseBodyThresholdsThreshold) SetElasticBps(v int32) *DescribeDdosThresholdResponseBodyThresholdsThreshold {
	s.ElasticBps = &v
	return s
}

func (s *DescribeDdosThresholdResponseBodyThresholdsThreshold) SetInstanceId(v string) *DescribeDdosThresholdResponseBodyThresholdsThreshold {
	s.InstanceId = &v
	return s
}

func (s *DescribeDdosThresholdResponseBodyThresholdsThreshold) SetInternetIp(v string) *DescribeDdosThresholdResponseBodyThresholdsThreshold {
	s.InternetIp = &v
	return s
}

func (s *DescribeDdosThresholdResponseBodyThresholdsThreshold) SetIsAuto(v bool) *DescribeDdosThresholdResponseBodyThresholdsThreshold {
	s.IsAuto = &v
	return s
}

func (s *DescribeDdosThresholdResponseBodyThresholdsThreshold) SetMaxBps(v int32) *DescribeDdosThresholdResponseBodyThresholdsThreshold {
	s.MaxBps = &v
	return s
}

func (s *DescribeDdosThresholdResponseBodyThresholdsThreshold) SetMaxPps(v int32) *DescribeDdosThresholdResponseBodyThresholdsThreshold {
	s.MaxPps = &v
	return s
}

func (s *DescribeDdosThresholdResponseBodyThresholdsThreshold) SetPps(v int32) *DescribeDdosThresholdResponseBodyThresholdsThreshold {
	s.Pps = &v
	return s
}

type DescribeDdosThresholdResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDdosThresholdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDdosThresholdResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosThresholdResponse) GoString() string {
	return s.String()
}

func (s *DescribeDdosThresholdResponse) SetHeaders(v map[string]*string) *DescribeDdosThresholdResponse {
	s.Headers = v
	return s
}

func (s *DescribeDdosThresholdResponse) SetStatusCode(v int32) *DescribeDdosThresholdResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDdosThresholdResponse) SetBody(v *DescribeDdosThresholdResponseBody) *DescribeDdosThresholdResponse {
	s.Body = v
	return s
}

type DescribeInstanceRequest struct {
	CurrentPage  *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
	DdosStatus   *string `json:"DdosStatus,omitempty" xml:"DdosStatus,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceIp   *string `json:"InstanceIp,omitempty" xml:"InstanceIp,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRequest) SetCurrentPage(v int32) *DescribeInstanceRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInstanceRequest) SetDdosRegionId(v string) *DescribeInstanceRequest {
	s.DdosRegionId = &v
	return s
}

func (s *DescribeInstanceRequest) SetDdosStatus(v string) *DescribeInstanceRequest {
	s.DdosStatus = &v
	return s
}

func (s *DescribeInstanceRequest) SetInstanceId(v string) *DescribeInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceRequest) SetInstanceIp(v string) *DescribeInstanceRequest {
	s.InstanceIp = &v
	return s
}

func (s *DescribeInstanceRequest) SetInstanceName(v string) *DescribeInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstanceRequest) SetInstanceType(v string) *DescribeInstanceRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstanceRequest) SetPageSize(v int32) *DescribeInstanceRequest {
	s.PageSize = &v
	return s
}

type DescribeInstanceResponseBody struct {
	InstanceList *DescribeInstanceResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Struct"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total        *int32                                    `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBody) SetInstanceList(v *DescribeInstanceResponseBodyInstanceList) *DescribeInstanceResponseBody {
	s.InstanceList = v
	return s
}

func (s *DescribeInstanceResponseBody) SetRequestId(v string) *DescribeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetTotal(v int32) *DescribeInstanceResponseBody {
	s.Total = &v
	return s
}

type DescribeInstanceResponseBodyInstanceList struct {
	Instance []*DescribeInstanceResponseBodyInstanceListInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s DescribeInstanceResponseBodyInstanceList) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyInstanceList) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyInstanceList) SetInstance(v []*DescribeInstanceResponseBodyInstanceListInstance) *DescribeInstanceResponseBodyInstanceList {
	s.Instance = v
	return s
}

type DescribeInstanceResponseBodyInstanceListInstance struct {
	BlackholeThreshold  *int32  `json:"BlackholeThreshold,omitempty" xml:"BlackholeThreshold,omitempty"`
	DefenseBpsThreshold *int32  `json:"DefenseBpsThreshold,omitempty" xml:"DefenseBpsThreshold,omitempty"`
	DefensePpsThreshold *int32  `json:"DefensePpsThreshold,omitempty" xml:"DefensePpsThreshold,omitempty"`
	ElasticThreshold    *int32  `json:"ElasticThreshold,omitempty" xml:"ElasticThreshold,omitempty"`
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceIp          *string `json:"InstanceIp,omitempty" xml:"InstanceIp,omitempty"`
	InstanceName        *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceStatus      *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	InstanceType        *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	IpVersion           *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	IsBgppack           *bool   `json:"IsBgppack,omitempty" xml:"IsBgppack,omitempty"`
}

func (s DescribeInstanceResponseBodyInstanceListInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyInstanceListInstance) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetBlackholeThreshold(v int32) *DescribeInstanceResponseBodyInstanceListInstance {
	s.BlackholeThreshold = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetDefenseBpsThreshold(v int32) *DescribeInstanceResponseBodyInstanceListInstance {
	s.DefenseBpsThreshold = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetDefensePpsThreshold(v int32) *DescribeInstanceResponseBodyInstanceListInstance {
	s.DefensePpsThreshold = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetElasticThreshold(v int32) *DescribeInstanceResponseBodyInstanceListInstance {
	s.ElasticThreshold = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetInstanceId(v string) *DescribeInstanceResponseBodyInstanceListInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetInstanceIp(v string) *DescribeInstanceResponseBodyInstanceListInstance {
	s.InstanceIp = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetInstanceName(v string) *DescribeInstanceResponseBodyInstanceListInstance {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetInstanceStatus(v string) *DescribeInstanceResponseBodyInstanceListInstance {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetInstanceType(v string) *DescribeInstanceResponseBodyInstanceListInstance {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetIpVersion(v string) *DescribeInstanceResponseBodyInstanceListInstance {
	s.IpVersion = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceListInstance) SetIsBgppack(v bool) *DescribeInstanceResponseBodyInstanceListInstance {
	s.IsBgppack = &v
	return s
}

type DescribeInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponse) SetHeaders(v map[string]*string) *DescribeInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceResponse) SetStatusCode(v int32) *DescribeInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceResponse) SetBody(v *DescribeInstanceResponseBody) *DescribeInstanceResponse {
	s.Body = v
	return s
}

type DescribeInstanceIpAddressRequest struct {
	CurrentPage  *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
	DdosStatus   *string `json:"DdosStatus,omitempty" xml:"DdosStatus,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceIp   *string `json:"InstanceIp,omitempty" xml:"InstanceIp,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeInstanceIpAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceIpAddressRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceIpAddressRequest) SetCurrentPage(v int32) *DescribeInstanceIpAddressRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInstanceIpAddressRequest) SetDdosRegionId(v string) *DescribeInstanceIpAddressRequest {
	s.DdosRegionId = &v
	return s
}

func (s *DescribeInstanceIpAddressRequest) SetDdosStatus(v string) *DescribeInstanceIpAddressRequest {
	s.DdosStatus = &v
	return s
}

func (s *DescribeInstanceIpAddressRequest) SetInstanceId(v string) *DescribeInstanceIpAddressRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceIpAddressRequest) SetInstanceIp(v string) *DescribeInstanceIpAddressRequest {
	s.InstanceIp = &v
	return s
}

func (s *DescribeInstanceIpAddressRequest) SetInstanceName(v string) *DescribeInstanceIpAddressRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstanceIpAddressRequest) SetInstanceType(v string) *DescribeInstanceIpAddressRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstanceIpAddressRequest) SetPageSize(v int32) *DescribeInstanceIpAddressRequest {
	s.PageSize = &v
	return s
}

type DescribeInstanceIpAddressResponseBody struct {
	InstanceList []*DescribeInstanceIpAddressResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
	RequestId    *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total        *int32                                               `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeInstanceIpAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceIpAddressResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceIpAddressResponseBody) SetInstanceList(v []*DescribeInstanceIpAddressResponseBodyInstanceList) *DescribeInstanceIpAddressResponseBody {
	s.InstanceList = v
	return s
}

func (s *DescribeInstanceIpAddressResponseBody) SetRequestId(v string) *DescribeInstanceIpAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBody) SetTotal(v int32) *DescribeInstanceIpAddressResponseBody {
	s.Total = &v
	return s
}

type DescribeInstanceIpAddressResponseBodyInstanceList struct {
	InstanceId      *string                                                             `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName    *string                                                             `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceStatus  *string                                                             `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	InstanceType    *string                                                             `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	IpAddressConfig []*DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig `json:"IpAddressConfig,omitempty" xml:"IpAddressConfig,omitempty" type:"Repeated"`
}

func (s DescribeInstanceIpAddressResponseBodyInstanceList) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceIpAddressResponseBodyInstanceList) GoString() string {
	return s.String()
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceList) SetInstanceId(v string) *DescribeInstanceIpAddressResponseBodyInstanceList {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceList) SetInstanceName(v string) *DescribeInstanceIpAddressResponseBodyInstanceList {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceList) SetInstanceStatus(v string) *DescribeInstanceIpAddressResponseBodyInstanceList {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceList) SetInstanceType(v string) *DescribeInstanceIpAddressResponseBodyInstanceList {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceList) SetIpAddressConfig(v []*DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) *DescribeInstanceIpAddressResponseBodyInstanceList {
	s.IpAddressConfig = v
	return s
}

type DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig struct {
	BlackholeThreshold  *int32  `json:"BlackholeThreshold,omitempty" xml:"BlackholeThreshold,omitempty"`
	DefenseBpsThreshold *int32  `json:"DefenseBpsThreshold,omitempty" xml:"DefenseBpsThreshold,omitempty"`
	DefensePpsThreshold *int32  `json:"DefensePpsThreshold,omitempty" xml:"DefensePpsThreshold,omitempty"`
	ElasticThreshold    *int32  `json:"ElasticThreshold,omitempty" xml:"ElasticThreshold,omitempty"`
	InstanceIp          *string `json:"InstanceIp,omitempty" xml:"InstanceIp,omitempty"`
	IpStatus            *string `json:"IpStatus,omitempty" xml:"IpStatus,omitempty"`
	IpVersion           *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	IsBgppack           *bool   `json:"IsBgppack,omitempty" xml:"IsBgppack,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) GoString() string {
	return s.String()
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) SetBlackholeThreshold(v int32) *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig {
	s.BlackholeThreshold = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) SetDefenseBpsThreshold(v int32) *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig {
	s.DefenseBpsThreshold = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) SetDefensePpsThreshold(v int32) *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig {
	s.DefensePpsThreshold = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) SetElasticThreshold(v int32) *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig {
	s.ElasticThreshold = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) SetInstanceIp(v string) *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig {
	s.InstanceIp = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) SetIpStatus(v string) *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig {
	s.IpStatus = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) SetIpVersion(v string) *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig {
	s.IpVersion = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) SetIsBgppack(v bool) *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig {
	s.IsBgppack = &v
	return s
}

func (s *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig) SetRegionId(v string) *DescribeInstanceIpAddressResponseBodyInstanceListIpAddressConfig {
	s.RegionId = &v
	return s
}

type DescribeInstanceIpAddressResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInstanceIpAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceIpAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceIpAddressResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceIpAddressResponse) SetHeaders(v map[string]*string) *DescribeInstanceIpAddressResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceIpAddressResponse) SetStatusCode(v int32) *DescribeInstanceIpAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceIpAddressResponse) SetBody(v *DescribeInstanceIpAddressResponseBody) *DescribeInstanceIpAddressResponse {
	s.Body = v
	return s
}

type DescribeIpDdosThresholdRequest struct {
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
	DdosType     *string `json:"DdosType,omitempty" xml:"DdosType,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InternetIp   *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
}

func (s DescribeIpDdosThresholdRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpDdosThresholdRequest) GoString() string {
	return s.String()
}

func (s *DescribeIpDdosThresholdRequest) SetDdosRegionId(v string) *DescribeIpDdosThresholdRequest {
	s.DdosRegionId = &v
	return s
}

func (s *DescribeIpDdosThresholdRequest) SetDdosType(v string) *DescribeIpDdosThresholdRequest {
	s.DdosType = &v
	return s
}

func (s *DescribeIpDdosThresholdRequest) SetInstanceId(v string) *DescribeIpDdosThresholdRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeIpDdosThresholdRequest) SetInstanceType(v string) *DescribeIpDdosThresholdRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeIpDdosThresholdRequest) SetInternetIp(v string) *DescribeIpDdosThresholdRequest {
	s.InternetIp = &v
	return s
}

type DescribeIpDdosThresholdResponseBody struct {
	// Id of the request
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Threshold *DescribeIpDdosThresholdResponseBodyThreshold `json:"Threshold,omitempty" xml:"Threshold,omitempty" type:"Struct"`
}

func (s DescribeIpDdosThresholdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpDdosThresholdResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIpDdosThresholdResponseBody) SetRequestId(v string) *DescribeIpDdosThresholdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIpDdosThresholdResponseBody) SetThreshold(v *DescribeIpDdosThresholdResponseBodyThreshold) *DescribeIpDdosThresholdResponseBody {
	s.Threshold = v
	return s
}

type DescribeIpDdosThresholdResponseBodyThreshold struct {
	Bps        *int32  `json:"Bps,omitempty" xml:"Bps,omitempty"`
	DdosType   *string `json:"DdosType,omitempty" xml:"DdosType,omitempty"`
	ElasticBps *int32  `json:"ElasticBps,omitempty" xml:"ElasticBps,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	IsAuto     *bool   `json:"IsAuto,omitempty" xml:"IsAuto,omitempty"`
	MaxBps     *int32  `json:"MaxBps,omitempty" xml:"MaxBps,omitempty"`
	MaxPps     *int32  `json:"MaxPps,omitempty" xml:"MaxPps,omitempty"`
	Pps        *int32  `json:"Pps,omitempty" xml:"Pps,omitempty"`
}

func (s DescribeIpDdosThresholdResponseBodyThreshold) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpDdosThresholdResponseBodyThreshold) GoString() string {
	return s.String()
}

func (s *DescribeIpDdosThresholdResponseBodyThreshold) SetBps(v int32) *DescribeIpDdosThresholdResponseBodyThreshold {
	s.Bps = &v
	return s
}

func (s *DescribeIpDdosThresholdResponseBodyThreshold) SetDdosType(v string) *DescribeIpDdosThresholdResponseBodyThreshold {
	s.DdosType = &v
	return s
}

func (s *DescribeIpDdosThresholdResponseBodyThreshold) SetElasticBps(v int32) *DescribeIpDdosThresholdResponseBodyThreshold {
	s.ElasticBps = &v
	return s
}

func (s *DescribeIpDdosThresholdResponseBodyThreshold) SetInstanceId(v string) *DescribeIpDdosThresholdResponseBodyThreshold {
	s.InstanceId = &v
	return s
}

func (s *DescribeIpDdosThresholdResponseBodyThreshold) SetInternetIp(v string) *DescribeIpDdosThresholdResponseBodyThreshold {
	s.InternetIp = &v
	return s
}

func (s *DescribeIpDdosThresholdResponseBodyThreshold) SetIsAuto(v bool) *DescribeIpDdosThresholdResponseBodyThreshold {
	s.IsAuto = &v
	return s
}

func (s *DescribeIpDdosThresholdResponseBodyThreshold) SetMaxBps(v int32) *DescribeIpDdosThresholdResponseBodyThreshold {
	s.MaxBps = &v
	return s
}

func (s *DescribeIpDdosThresholdResponseBodyThreshold) SetMaxPps(v int32) *DescribeIpDdosThresholdResponseBodyThreshold {
	s.MaxPps = &v
	return s
}

func (s *DescribeIpDdosThresholdResponseBodyThreshold) SetPps(v int32) *DescribeIpDdosThresholdResponseBodyThreshold {
	s.Pps = &v
	return s
}

type DescribeIpDdosThresholdResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeIpDdosThresholdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeIpDdosThresholdResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpDdosThresholdResponse) GoString() string {
	return s.String()
}

func (s *DescribeIpDdosThresholdResponse) SetHeaders(v map[string]*string) *DescribeIpDdosThresholdResponse {
	s.Headers = v
	return s
}

func (s *DescribeIpDdosThresholdResponse) SetStatusCode(v int32) *DescribeIpDdosThresholdResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIpDdosThresholdResponse) SetBody(v *DescribeIpDdosThresholdResponseBody) *DescribeIpDdosThresholdResponse {
	s.Body = v
	return s
}

type DescribeIpLocationServiceRequest struct {
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
}

func (s DescribeIpLocationServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpLocationServiceRequest) GoString() string {
	return s.String()
}

func (s *DescribeIpLocationServiceRequest) SetInternetIp(v string) *DescribeIpLocationServiceRequest {
	s.InternetIp = &v
	return s
}

type DescribeIpLocationServiceResponseBody struct {
	// instance model
	Instance *DescribeIpLocationServiceResponseBodyInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Struct"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeIpLocationServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpLocationServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIpLocationServiceResponseBody) SetInstance(v *DescribeIpLocationServiceResponseBodyInstance) *DescribeIpLocationServiceResponseBody {
	s.Instance = v
	return s
}

func (s *DescribeIpLocationServiceResponseBody) SetRequestId(v string) *DescribeIpLocationServiceResponseBody {
	s.RequestId = &v
	return s
}

type DescribeIpLocationServiceResponseBodyInstance struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InternetIp   *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DescribeIpLocationServiceResponseBodyInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpLocationServiceResponseBodyInstance) GoString() string {
	return s.String()
}

func (s *DescribeIpLocationServiceResponseBodyInstance) SetInstanceId(v string) *DescribeIpLocationServiceResponseBodyInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeIpLocationServiceResponseBodyInstance) SetInstanceName(v string) *DescribeIpLocationServiceResponseBodyInstance {
	s.InstanceName = &v
	return s
}

func (s *DescribeIpLocationServiceResponseBodyInstance) SetInstanceType(v string) *DescribeIpLocationServiceResponseBodyInstance {
	s.InstanceType = &v
	return s
}

func (s *DescribeIpLocationServiceResponseBodyInstance) SetInternetIp(v string) *DescribeIpLocationServiceResponseBodyInstance {
	s.InternetIp = &v
	return s
}

func (s *DescribeIpLocationServiceResponseBodyInstance) SetRegion(v string) *DescribeIpLocationServiceResponseBodyInstance {
	s.Region = &v
	return s
}

type DescribeIpLocationServiceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeIpLocationServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeIpLocationServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpLocationServiceResponse) GoString() string {
	return s.String()
}

func (s *DescribeIpLocationServiceResponse) SetHeaders(v map[string]*string) *DescribeIpLocationServiceResponse {
	s.Headers = v
	return s
}

func (s *DescribeIpLocationServiceResponse) SetStatusCode(v int32) *DescribeIpLocationServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIpLocationServiceResponse) SetBody(v *DescribeIpLocationServiceResponseBody) *DescribeIpLocationServiceResponse {
	s.Body = v
	return s
}

type DescribeRegionsResponseBody struct {
	Regions   *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	Region []*DescribeRegionsResponseBodyRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRegion(v []*DescribeRegionsResponseBodyRegionsRegion) *DescribeRegionsResponseBodyRegions {
	s.Region = v
	return s
}

type DescribeRegionsResponseBodyRegionsRegion struct {
	RegionEnName  *string `json:"RegionEnName,omitempty" xml:"RegionEnName,omitempty"`
	RegionName    *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	RegionNo      *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	RegionNoAlias *string `json:"RegionNoAlias,omitempty" xml:"RegionNoAlias,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionEnName(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionEnName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionName(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionNo(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionNo = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionNoAlias(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionNoAlias = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type ModifyDdosStatusRequest struct {
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s ModifyDdosStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDdosStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyDdosStatusRequest) SetDdosRegionId(v string) *ModifyDdosStatusRequest {
	s.DdosRegionId = &v
	return s
}

func (s *ModifyDdosStatusRequest) SetInstanceId(v string) *ModifyDdosStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDdosStatusRequest) SetInstanceType(v string) *ModifyDdosStatusRequest {
	s.InstanceType = &v
	return s
}

type ModifyDdosStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDdosStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDdosStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDdosStatusResponseBody) SetRequestId(v string) *ModifyDdosStatusResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDdosStatusResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDdosStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDdosStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDdosStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyDdosStatusResponse) SetHeaders(v map[string]*string) *ModifyDdosStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyDdosStatusResponse) SetStatusCode(v int32) *ModifyDdosStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDdosStatusResponse) SetBody(v *ModifyDdosStatusResponseBody) *ModifyDdosStatusResponse {
	s.Body = v
	return s
}

type ModifyDefenseThresholdRequest struct {
	Bps          *int32  `json:"Bps,omitempty" xml:"Bps,omitempty"`
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InternetIp   *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	IsAuto       *bool   `json:"IsAuto,omitempty" xml:"IsAuto,omitempty"`
	Pps          *int32  `json:"Pps,omitempty" xml:"Pps,omitempty"`
}

func (s ModifyDefenseThresholdRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseThresholdRequest) GoString() string {
	return s.String()
}

func (s *ModifyDefenseThresholdRequest) SetBps(v int32) *ModifyDefenseThresholdRequest {
	s.Bps = &v
	return s
}

func (s *ModifyDefenseThresholdRequest) SetDdosRegionId(v string) *ModifyDefenseThresholdRequest {
	s.DdosRegionId = &v
	return s
}

func (s *ModifyDefenseThresholdRequest) SetInstanceId(v string) *ModifyDefenseThresholdRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDefenseThresholdRequest) SetInstanceType(v string) *ModifyDefenseThresholdRequest {
	s.InstanceType = &v
	return s
}

func (s *ModifyDefenseThresholdRequest) SetInternetIp(v string) *ModifyDefenseThresholdRequest {
	s.InternetIp = &v
	return s
}

func (s *ModifyDefenseThresholdRequest) SetIsAuto(v bool) *ModifyDefenseThresholdRequest {
	s.IsAuto = &v
	return s
}

func (s *ModifyDefenseThresholdRequest) SetPps(v int32) *ModifyDefenseThresholdRequest {
	s.Pps = &v
	return s
}

type ModifyDefenseThresholdResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDefenseThresholdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseThresholdResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDefenseThresholdResponseBody) SetRequestId(v string) *ModifyDefenseThresholdResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDefenseThresholdResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDefenseThresholdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDefenseThresholdResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseThresholdResponse) GoString() string {
	return s.String()
}

func (s *ModifyDefenseThresholdResponse) SetHeaders(v map[string]*string) *ModifyDefenseThresholdResponse {
	s.Headers = v
	return s
}

func (s *ModifyDefenseThresholdResponse) SetStatusCode(v int32) *ModifyDefenseThresholdResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDefenseThresholdResponse) SetBody(v *ModifyDefenseThresholdResponseBody) *ModifyDefenseThresholdResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("antiddos-public"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) DescribeBgpPackByIpWithOptions(request *DescribeBgpPackByIpRequest, runtime *util.RuntimeOptions) (_result *DescribeBgpPackByIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DdosRegionId)) {
		query["DdosRegionId"] = request.DdosRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBgpPackByIp"),
		Version:     tea.String("2017-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBgpPackByIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBgpPackByIp(request *DescribeBgpPackByIpRequest) (_result *DescribeBgpPackByIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBgpPackByIpResponse{}
	_body, _err := client.DescribeBgpPackByIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCapWithOptions(request *DescribeCapRequest, runtime *util.RuntimeOptions) (_result *DescribeCapResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BegTime)) {
		query["BegTime"] = request.BegTime
	}

	if !tea.BoolValue(util.IsUnset(request.DdosRegionId)) {
		query["DdosRegionId"] = request.DdosRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.InternetIp)) {
		query["InternetIp"] = request.InternetIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCap"),
		Version:     tea.String("2017-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCapResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCap(request *DescribeCapRequest) (_result *DescribeCapResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCapResponse{}
	_body, _err := client.DescribeCapWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDdosCountWithOptions(request *DescribeDdosCountRequest, runtime *util.RuntimeOptions) (_result *DescribeDdosCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DdosRegionId)) {
		query["DdosRegionId"] = request.DdosRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDdosCount"),
		Version:     tea.String("2017-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDdosCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDdosCount(request *DescribeDdosCountRequest) (_result *DescribeDdosCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDdosCountResponse{}
	_body, _err := client.DescribeDdosCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDdosCreditWithOptions(request *DescribeDdosCreditRequest, runtime *util.RuntimeOptions) (_result *DescribeDdosCreditResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DdosRegionId)) {
		query["DdosRegionId"] = request.DdosRegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDdosCredit"),
		Version:     tea.String("2017-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDdosCreditResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDdosCredit(request *DescribeDdosCreditRequest) (_result *DescribeDdosCreditResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDdosCreditResponse{}
	_body, _err := client.DescribeDdosCreditWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDdosEventListWithOptions(request *DescribeDdosEventListRequest, runtime *util.RuntimeOptions) (_result *DescribeDdosEventListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DdosRegionId)) {
		query["DdosRegionId"] = request.DdosRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.InternetIp)) {
		query["InternetIp"] = request.InternetIp
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDdosEventList"),
		Version:     tea.String("2017-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDdosEventListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDdosEventList(request *DescribeDdosEventListRequest) (_result *DescribeDdosEventListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDdosEventListResponse{}
	_body, _err := client.DescribeDdosEventListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDdosThresholdWithOptions(request *DescribeDdosThresholdRequest, runtime *util.RuntimeOptions) (_result *DescribeDdosThresholdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DdosRegionId)) {
		query["DdosRegionId"] = request.DdosRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.DdosType)) {
		query["DdosType"] = request.DdosType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDdosThreshold"),
		Version:     tea.String("2017-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDdosThresholdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDdosThreshold(request *DescribeDdosThresholdRequest) (_result *DescribeDdosThresholdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDdosThresholdResponse{}
	_body, _err := client.DescribeDdosThresholdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceWithOptions(request *DescribeInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DdosRegionId)) {
		query["DdosRegionId"] = request.DdosRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.DdosStatus)) {
		query["DdosStatus"] = request.DdosStatus
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIp)) {
		query["InstanceIp"] = request.InstanceIp
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstance"),
		Version:     tea.String("2017-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstance(request *DescribeInstanceRequest) (_result *DescribeInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.DescribeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceIpAddressWithOptions(request *DescribeInstanceIpAddressRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceIpAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DdosRegionId)) {
		query["DdosRegionId"] = request.DdosRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.DdosStatus)) {
		query["DdosStatus"] = request.DdosStatus
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIp)) {
		query["InstanceIp"] = request.InstanceIp
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceIpAddress"),
		Version:     tea.String("2017-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceIpAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceIpAddress(request *DescribeInstanceIpAddressRequest) (_result *DescribeInstanceIpAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceIpAddressResponse{}
	_body, _err := client.DescribeInstanceIpAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeIpDdosThresholdWithOptions(request *DescribeIpDdosThresholdRequest, runtime *util.RuntimeOptions) (_result *DescribeIpDdosThresholdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DdosRegionId)) {
		query["DdosRegionId"] = request.DdosRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.DdosType)) {
		query["DdosType"] = request.DdosType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.InternetIp)) {
		query["InternetIp"] = request.InternetIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeIpDdosThreshold"),
		Version:     tea.String("2017-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeIpDdosThresholdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeIpDdosThreshold(request *DescribeIpDdosThresholdRequest) (_result *DescribeIpDdosThresholdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeIpDdosThresholdResponse{}
	_body, _err := client.DescribeIpDdosThresholdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeIpLocationServiceWithOptions(request *DescribeIpLocationServiceRequest, runtime *util.RuntimeOptions) (_result *DescribeIpLocationServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InternetIp)) {
		query["InternetIp"] = request.InternetIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeIpLocationService"),
		Version:     tea.String("2017-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeIpLocationServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeIpLocationService(request *DescribeIpLocationServiceRequest) (_result *DescribeIpLocationServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeIpLocationServiceResponse{}
	_body, _err := client.DescribeIpLocationServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2017-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions() (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDdosStatusWithOptions(request *ModifyDdosStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyDdosStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DdosRegionId)) {
		query["DdosRegionId"] = request.DdosRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDdosStatus"),
		Version:     tea.String("2017-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDdosStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDdosStatus(request *ModifyDdosStatusRequest) (_result *ModifyDdosStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDdosStatusResponse{}
	_body, _err := client.ModifyDdosStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDefenseThresholdWithOptions(request *ModifyDefenseThresholdRequest, runtime *util.RuntimeOptions) (_result *ModifyDefenseThresholdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bps)) {
		query["Bps"] = request.Bps
	}

	if !tea.BoolValue(util.IsUnset(request.DdosRegionId)) {
		query["DdosRegionId"] = request.DdosRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.InternetIp)) {
		query["InternetIp"] = request.InternetIp
	}

	if !tea.BoolValue(util.IsUnset(request.IsAuto)) {
		query["IsAuto"] = request.IsAuto
	}

	if !tea.BoolValue(util.IsUnset(request.Pps)) {
		query["Pps"] = request.Pps
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDefenseThreshold"),
		Version:     tea.String("2017-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDefenseThresholdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDefenseThreshold(request *ModifyDefenseThresholdRequest) (_result *ModifyDefenseThresholdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDefenseThresholdResponse{}
	_body, _err := client.ModifyDefenseThresholdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
