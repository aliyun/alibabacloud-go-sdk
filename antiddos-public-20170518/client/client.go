// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DescribeBgpPackByIpRequest struct {
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
	Ip           *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribeBgpPackByIpRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBgpPackByIpRequest) GoString() string {
	return s.String()
}

func (s *DescribeBgpPackByIpRequest) SetSourceIp(v string) *DescribeBgpPackByIpRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeBgpPackByIpRequest) SetLang(v string) *DescribeBgpPackByIpRequest {
	s.Lang = &v
	return s
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
	ExpireTime        *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	ElasticThreshold  *int32  `json:"ElasticThreshold,omitempty" xml:"ElasticThreshold,omitempty"`
	Ip                *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	BaseThreshold     *int32  `json:"BaseThreshold,omitempty" xml:"BaseThreshold,omitempty"`
	DdosbgpInstanceId *string `json:"DdosbgpInstanceId,omitempty" xml:"DdosbgpInstanceId,omitempty"`
}

func (s DescribeBgpPackByIpResponseBodyDdosbgpInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeBgpPackByIpResponseBodyDdosbgpInfo) GoString() string {
	return s.String()
}

func (s *DescribeBgpPackByIpResponseBodyDdosbgpInfo) SetExpireTime(v int64) *DescribeBgpPackByIpResponseBodyDdosbgpInfo {
	s.ExpireTime = &v
	return s
}

func (s *DescribeBgpPackByIpResponseBodyDdosbgpInfo) SetElasticThreshold(v int32) *DescribeBgpPackByIpResponseBodyDdosbgpInfo {
	s.ElasticThreshold = &v
	return s
}

func (s *DescribeBgpPackByIpResponseBodyDdosbgpInfo) SetIp(v string) *DescribeBgpPackByIpResponseBodyDdosbgpInfo {
	s.Ip = &v
	return s
}

func (s *DescribeBgpPackByIpResponseBodyDdosbgpInfo) SetBaseThreshold(v int32) *DescribeBgpPackByIpResponseBodyDdosbgpInfo {
	s.BaseThreshold = &v
	return s
}

func (s *DescribeBgpPackByIpResponseBodyDdosbgpInfo) SetDdosbgpInstanceId(v string) *DescribeBgpPackByIpResponseBodyDdosbgpInfo {
	s.DdosbgpInstanceId = &v
	return s
}

type DescribeBgpPackByIpResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBgpPackByIpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeBgpPackByIpResponse) SetBody(v *DescribeBgpPackByIpResponseBody) *DescribeBgpPackByIpResponse {
	s.Body = v
	return s
}

type DescribeBgpPackElasticThresholdRequest struct {
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
}

func (s DescribeBgpPackElasticThresholdRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBgpPackElasticThresholdRequest) GoString() string {
	return s.String()
}

func (s *DescribeBgpPackElasticThresholdRequest) SetSourceIp(v string) *DescribeBgpPackElasticThresholdRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeBgpPackElasticThresholdRequest) SetLang(v string) *DescribeBgpPackElasticThresholdRequest {
	s.Lang = &v
	return s
}

func (s *DescribeBgpPackElasticThresholdRequest) SetDdosRegionId(v string) *DescribeBgpPackElasticThresholdRequest {
	s.DdosRegionId = &v
	return s
}

type DescribeBgpPackElasticThresholdResponseBody struct {
	MaxThreshold *int32  `json:"MaxThreshold,omitempty" xml:"MaxThreshold,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Openable     *bool   `json:"Openable,omitempty" xml:"Openable,omitempty"`
}

func (s DescribeBgpPackElasticThresholdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBgpPackElasticThresholdResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBgpPackElasticThresholdResponseBody) SetMaxThreshold(v int32) *DescribeBgpPackElasticThresholdResponseBody {
	s.MaxThreshold = &v
	return s
}

func (s *DescribeBgpPackElasticThresholdResponseBody) SetRequestId(v string) *DescribeBgpPackElasticThresholdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBgpPackElasticThresholdResponseBody) SetOpenable(v bool) *DescribeBgpPackElasticThresholdResponseBody {
	s.Openable = &v
	return s
}

type DescribeBgpPackElasticThresholdResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBgpPackElasticThresholdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBgpPackElasticThresholdResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBgpPackElasticThresholdResponse) GoString() string {
	return s.String()
}

func (s *DescribeBgpPackElasticThresholdResponse) SetHeaders(v map[string]*string) *DescribeBgpPackElasticThresholdResponse {
	s.Headers = v
	return s
}

func (s *DescribeBgpPackElasticThresholdResponse) SetBody(v *DescribeBgpPackElasticThresholdResponseBody) *DescribeBgpPackElasticThresholdResponse {
	s.Body = v
	return s
}

type DescribeCapRequest struct {
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	BegTime      *int64  `json:"BegTime,omitempty" xml:"BegTime,omitempty"`
}

func (s DescribeCapRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCapRequest) GoString() string {
	return s.String()
}

func (s *DescribeCapRequest) SetSourceIp(v string) *DescribeCapRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCapRequest) SetLang(v string) *DescribeCapRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCapRequest) SetDdosRegionId(v string) *DescribeCapRequest {
	s.DdosRegionId = &v
	return s
}

func (s *DescribeCapRequest) SetInstanceType(v string) *DescribeCapRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeCapRequest) SetInstanceId(v string) *DescribeCapRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeCapRequest) SetBegTime(v int64) *DescribeCapRequest {
	s.BegTime = &v
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
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCapResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeCapResponse) SetBody(v *DescribeCapResponseBody) *DescribeCapResponse {
	s.Body = v
	return s
}

type DescribeCreditInfoRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeCreditInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCreditInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeCreditInfoRequest) SetSourceIp(v string) *DescribeCreditInfoRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCreditInfoRequest) SetResourceOwnerId(v int64) *DescribeCreditInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeCreditInfoResponseBody struct {
	PunishTimes   *int64    `json:"PunishTimes,omitempty" xml:"PunishTimes,omitempty"`
	LastOrderTime *int64    `json:"LastOrderTime,omitempty" xml:"LastOrderTime,omitempty"`
	LastLoginTime *int64    `json:"LastLoginTime,omitempty" xml:"LastLoginTime,omitempty"`
	RequestId     *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserLevel     *string   `json:"UserLevel,omitempty" xml:"UserLevel,omitempty"`
	BlackTimes    *int64    `json:"BlackTimes,omitempty" xml:"BlackTimes,omitempty"`
	NewCreatetime *string   `json:"NewCreatetime,omitempty" xml:"NewCreatetime,omitempty"`
	Duration      *int64    `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Productid     []*string `json:"Productid,omitempty" xml:"Productid,omitempty" type:"Repeated"`
}

func (s DescribeCreditInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCreditInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCreditInfoResponseBody) SetPunishTimes(v int64) *DescribeCreditInfoResponseBody {
	s.PunishTimes = &v
	return s
}

func (s *DescribeCreditInfoResponseBody) SetLastOrderTime(v int64) *DescribeCreditInfoResponseBody {
	s.LastOrderTime = &v
	return s
}

func (s *DescribeCreditInfoResponseBody) SetLastLoginTime(v int64) *DescribeCreditInfoResponseBody {
	s.LastLoginTime = &v
	return s
}

func (s *DescribeCreditInfoResponseBody) SetRequestId(v string) *DescribeCreditInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCreditInfoResponseBody) SetUserLevel(v string) *DescribeCreditInfoResponseBody {
	s.UserLevel = &v
	return s
}

func (s *DescribeCreditInfoResponseBody) SetBlackTimes(v int64) *DescribeCreditInfoResponseBody {
	s.BlackTimes = &v
	return s
}

func (s *DescribeCreditInfoResponseBody) SetNewCreatetime(v string) *DescribeCreditInfoResponseBody {
	s.NewCreatetime = &v
	return s
}

func (s *DescribeCreditInfoResponseBody) SetDuration(v int64) *DescribeCreditInfoResponseBody {
	s.Duration = &v
	return s
}

func (s *DescribeCreditInfoResponseBody) SetProductid(v []*string) *DescribeCreditInfoResponseBody {
	s.Productid = v
	return s
}

type DescribeCreditInfoResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCreditInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCreditInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCreditInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeCreditInfoResponse) SetHeaders(v map[string]*string) *DescribeCreditInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeCreditInfoResponse) SetBody(v *DescribeCreditInfoResponseBody) *DescribeCreditInfoResponse {
	s.Body = v
	return s
}

type DescribeDdosCountRequest struct {
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeDdosCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeDdosCountRequest) SetSourceIp(v string) *DescribeDdosCountRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDdosCountRequest) SetLang(v string) *DescribeDdosCountRequest {
	s.Lang = &v
	return s
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
	DefenseCount   *int32 `json:"DefenseCount,omitempty" xml:"DefenseCount,omitempty"`
	BlackholeCount *int32 `json:"BlackholeCount,omitempty" xml:"BlackholeCount,omitempty"`
	InstacenCount  *int32 `json:"InstacenCount,omitempty" xml:"InstacenCount,omitempty"`
}

func (s DescribeDdosCountResponseBodyDdosCount) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosCountResponseBodyDdosCount) GoString() string {
	return s.String()
}

func (s *DescribeDdosCountResponseBodyDdosCount) SetDefenseCount(v int32) *DescribeDdosCountResponseBodyDdosCount {
	s.DefenseCount = &v
	return s
}

func (s *DescribeDdosCountResponseBodyDdosCount) SetBlackholeCount(v int32) *DescribeDdosCountResponseBodyDdosCount {
	s.BlackholeCount = &v
	return s
}

func (s *DescribeDdosCountResponseBodyDdosCount) SetInstacenCount(v int32) *DescribeDdosCountResponseBodyDdosCount {
	s.InstacenCount = &v
	return s
}

type DescribeDdosCountResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDdosCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDdosCountResponse) SetBody(v *DescribeDdosCountResponseBody) *DescribeDdosCountResponse {
	s.Body = v
	return s
}

type DescribeDdosCreditRequest struct {
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
}

func (s DescribeDdosCreditRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosCreditRequest) GoString() string {
	return s.String()
}

func (s *DescribeDdosCreditRequest) SetSourceIp(v string) *DescribeDdosCreditRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDdosCreditRequest) SetLang(v string) *DescribeDdosCreditRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDdosCreditRequest) SetDdosRegionId(v string) *DescribeDdosCreditRequest {
	s.DdosRegionId = &v
	return s
}

type DescribeDdosCreditResponseBody struct {
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DdosCredit *DescribeDdosCreditResponseBodyDdosCredit `json:"DdosCredit,omitempty" xml:"DdosCredit,omitempty" type:"Struct"`
	Success    *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDdosCreditResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosCreditResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDdosCreditResponseBody) SetRequestId(v string) *DescribeDdosCreditResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDdosCreditResponseBody) SetDdosCredit(v *DescribeDdosCreditResponseBodyDdosCredit) *DescribeDdosCreditResponseBody {
	s.DdosCredit = v
	return s
}

func (s *DescribeDdosCreditResponseBody) SetSuccess(v bool) *DescribeDdosCreditResponseBody {
	s.Success = &v
	return s
}

type DescribeDdosCreditResponseBodyDdosCredit struct {
	Score         *int32  `json:"Score,omitempty" xml:"Score,omitempty"`
	ScoreLevel    *string `json:"ScoreLevel,omitempty" xml:"ScoreLevel,omitempty"`
	BlackholeTime *int32  `json:"BlackholeTime,omitempty" xml:"BlackholeTime,omitempty"`
}

func (s DescribeDdosCreditResponseBodyDdosCredit) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosCreditResponseBodyDdosCredit) GoString() string {
	return s.String()
}

func (s *DescribeDdosCreditResponseBodyDdosCredit) SetScore(v int32) *DescribeDdosCreditResponseBodyDdosCredit {
	s.Score = &v
	return s
}

func (s *DescribeDdosCreditResponseBodyDdosCredit) SetScoreLevel(v string) *DescribeDdosCreditResponseBodyDdosCredit {
	s.ScoreLevel = &v
	return s
}

func (s *DescribeDdosCreditResponseBodyDdosCredit) SetBlackholeTime(v int32) *DescribeDdosCreditResponseBodyDdosCredit {
	s.BlackholeTime = &v
	return s
}

type DescribeDdosCreditResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDdosCreditResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDdosCreditResponse) SetBody(v *DescribeDdosCreditResponseBody) *DescribeDdosCreditResponse {
	s.Body = v
	return s
}

type DescribeDdosEventListRequest struct {
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	CurrentPage  *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDdosEventListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosEventListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventListRequest) SetSourceIp(v string) *DescribeDdosEventListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDdosEventListRequest) SetLang(v string) *DescribeDdosEventListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDdosEventListRequest) SetDdosRegionId(v string) *DescribeDdosEventListRequest {
	s.DdosRegionId = &v
	return s
}

func (s *DescribeDdosEventListRequest) SetInstanceType(v string) *DescribeDdosEventListRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeDdosEventListRequest) SetInstanceId(v string) *DescribeDdosEventListRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDdosEventListRequest) SetCurrentPage(v int32) *DescribeDdosEventListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDdosEventListRequest) SetPageSize(v int32) *DescribeDdosEventListRequest {
	s.PageSize = &v
	return s
}

type DescribeDdosEventListResponseBody struct {
	DdosEventList []*DescribeDdosEventListResponseBodyDdosEventList `json:"DdosEventList,omitempty" xml:"DdosEventList,omitempty" type:"Repeated"`
	RequestId     *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total         *int32                                            `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDdosEventListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosEventListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventListResponseBody) SetDdosEventList(v []*DescribeDdosEventListResponseBodyDdosEventList) *DescribeDdosEventListResponseBody {
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
	DdosType        *string `json:"DdosType,omitempty" xml:"DdosType,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	DelayTime       *int64  `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	DdosStatus      *string `json:"DdosStatus,omitempty" xml:"DdosStatus,omitempty"`
	UnBlackholeTime *int64  `json:"UnBlackholeTime,omitempty" xml:"UnBlackholeTime,omitempty"`
}

func (s DescribeDdosEventListResponseBodyDdosEventList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosEventListResponseBodyDdosEventList) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventListResponseBodyDdosEventList) SetDdosType(v string) *DescribeDdosEventListResponseBodyDdosEventList {
	s.DdosType = &v
	return s
}

func (s *DescribeDdosEventListResponseBodyDdosEventList) SetEndTime(v int64) *DescribeDdosEventListResponseBodyDdosEventList {
	s.EndTime = &v
	return s
}

func (s *DescribeDdosEventListResponseBodyDdosEventList) SetStartTime(v int64) *DescribeDdosEventListResponseBodyDdosEventList {
	s.StartTime = &v
	return s
}

func (s *DescribeDdosEventListResponseBodyDdosEventList) SetDelayTime(v int64) *DescribeDdosEventListResponseBodyDdosEventList {
	s.DelayTime = &v
	return s
}

func (s *DescribeDdosEventListResponseBodyDdosEventList) SetDdosStatus(v string) *DescribeDdosEventListResponseBodyDdosEventList {
	s.DdosStatus = &v
	return s
}

func (s *DescribeDdosEventListResponseBodyDdosEventList) SetUnBlackholeTime(v int64) *DescribeDdosEventListResponseBodyDdosEventList {
	s.UnBlackholeTime = &v
	return s
}

type DescribeDdosEventListResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDdosEventListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDdosEventListResponse) SetBody(v *DescribeDdosEventListResponseBody) *DescribeDdosEventListResponse {
	s.Body = v
	return s
}

type DescribeDdosThresholdRequest struct {
	SourceIp     *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang         *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	DdosRegionId *string   `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
	DdosType     *string   `json:"DdosType,omitempty" xml:"DdosType,omitempty"`
	InstanceType *string   `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InstanceIds  []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s DescribeDdosThresholdRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosThresholdRequest) GoString() string {
	return s.String()
}

func (s *DescribeDdosThresholdRequest) SetSourceIp(v string) *DescribeDdosThresholdRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDdosThresholdRequest) SetLang(v string) *DescribeDdosThresholdRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDdosThresholdRequest) SetDdosRegionId(v string) *DescribeDdosThresholdRequest {
	s.DdosRegionId = &v
	return s
}

func (s *DescribeDdosThresholdRequest) SetDdosType(v string) *DescribeDdosThresholdRequest {
	s.DdosType = &v
	return s
}

func (s *DescribeDdosThresholdRequest) SetInstanceType(v string) *DescribeDdosThresholdRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeDdosThresholdRequest) SetInstanceIds(v []*string) *DescribeDdosThresholdRequest {
	s.InstanceIds = v
	return s
}

type DescribeDdosThresholdResponseBody struct {
	RequestId  *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Thresholds []*DescribeDdosThresholdResponseBodyThresholds `json:"Thresholds,omitempty" xml:"Thresholds,omitempty" type:"Repeated"`
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

func (s *DescribeDdosThresholdResponseBody) SetThresholds(v []*DescribeDdosThresholdResponseBodyThresholds) *DescribeDdosThresholdResponseBody {
	s.Thresholds = v
	return s
}

type DescribeDdosThresholdResponseBodyThresholds struct {
	DdosType   *string `json:"DdosType,omitempty" xml:"DdosType,omitempty"`
	IsAuto     *bool   `json:"IsAuto,omitempty" xml:"IsAuto,omitempty"`
	MaxBps     *int32  `json:"MaxBps,omitempty" xml:"MaxBps,omitempty"`
	ElasticBps *int32  `json:"ElasticBps,omitempty" xml:"ElasticBps,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Bps        *int32  `json:"Bps,omitempty" xml:"Bps,omitempty"`
	Pps        *int32  `json:"Pps,omitempty" xml:"Pps,omitempty"`
	MaxPps     *int32  `json:"MaxPps,omitempty" xml:"MaxPps,omitempty"`
}

func (s DescribeDdosThresholdResponseBodyThresholds) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosThresholdResponseBodyThresholds) GoString() string {
	return s.String()
}

func (s *DescribeDdosThresholdResponseBodyThresholds) SetDdosType(v string) *DescribeDdosThresholdResponseBodyThresholds {
	s.DdosType = &v
	return s
}

func (s *DescribeDdosThresholdResponseBodyThresholds) SetIsAuto(v bool) *DescribeDdosThresholdResponseBodyThresholds {
	s.IsAuto = &v
	return s
}

func (s *DescribeDdosThresholdResponseBodyThresholds) SetMaxBps(v int32) *DescribeDdosThresholdResponseBodyThresholds {
	s.MaxBps = &v
	return s
}

func (s *DescribeDdosThresholdResponseBodyThresholds) SetElasticBps(v int32) *DescribeDdosThresholdResponseBodyThresholds {
	s.ElasticBps = &v
	return s
}

func (s *DescribeDdosThresholdResponseBodyThresholds) SetInstanceId(v string) *DescribeDdosThresholdResponseBodyThresholds {
	s.InstanceId = &v
	return s
}

func (s *DescribeDdosThresholdResponseBodyThresholds) SetBps(v int32) *DescribeDdosThresholdResponseBodyThresholds {
	s.Bps = &v
	return s
}

func (s *DescribeDdosThresholdResponseBodyThresholds) SetPps(v int32) *DescribeDdosThresholdResponseBodyThresholds {
	s.Pps = &v
	return s
}

func (s *DescribeDdosThresholdResponseBodyThresholds) SetMaxPps(v int32) *DescribeDdosThresholdResponseBodyThresholds {
	s.MaxPps = &v
	return s
}

type DescribeDdosThresholdResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDdosThresholdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDdosThresholdResponse) SetBody(v *DescribeDdosThresholdResponseBody) *DescribeDdosThresholdResponse {
	s.Body = v
	return s
}

type DescribeFlexibleProtectionFlowRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Days     *int32  `json:"Days,omitempty" xml:"Days,omitempty"`
}

func (s DescribeFlexibleProtectionFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlexibleProtectionFlowRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlexibleProtectionFlowRequest) SetSourceIp(v string) *DescribeFlexibleProtectionFlowRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeFlexibleProtectionFlowRequest) SetLang(v string) *DescribeFlexibleProtectionFlowRequest {
	s.Lang = &v
	return s
}

func (s *DescribeFlexibleProtectionFlowRequest) SetDays(v int32) *DescribeFlexibleProtectionFlowRequest {
	s.Days = &v
	return s
}

type DescribeFlexibleProtectionFlowResponseBody struct {
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Flows     []*DescribeFlexibleProtectionFlowResponseBodyFlows `json:"Flows,omitempty" xml:"Flows,omitempty" type:"Repeated"`
}

func (s DescribeFlexibleProtectionFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlexibleProtectionFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlexibleProtectionFlowResponseBody) SetRequestId(v string) *DescribeFlexibleProtectionFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlexibleProtectionFlowResponseBody) SetFlows(v []*DescribeFlexibleProtectionFlowResponseBodyFlows) *DescribeFlexibleProtectionFlowResponseBody {
	s.Flows = v
	return s
}

type DescribeFlexibleProtectionFlowResponseBodyFlows struct {
	UsedFlow    *float32 `json:"UsedFlow,omitempty" xml:"UsedFlow,omitempty"`
	Time        *int64   `json:"Time,omitempty" xml:"Time,omitempty"`
	AddFlow     *float32 `json:"AddFlow,omitempty" xml:"AddFlow,omitempty"`
	UseableFlow *float32 `json:"UseableFlow,omitempty" xml:"UseableFlow,omitempty"`
}

func (s DescribeFlexibleProtectionFlowResponseBodyFlows) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlexibleProtectionFlowResponseBodyFlows) GoString() string {
	return s.String()
}

func (s *DescribeFlexibleProtectionFlowResponseBodyFlows) SetUsedFlow(v float32) *DescribeFlexibleProtectionFlowResponseBodyFlows {
	s.UsedFlow = &v
	return s
}

func (s *DescribeFlexibleProtectionFlowResponseBodyFlows) SetTime(v int64) *DescribeFlexibleProtectionFlowResponseBodyFlows {
	s.Time = &v
	return s
}

func (s *DescribeFlexibleProtectionFlowResponseBodyFlows) SetAddFlow(v float32) *DescribeFlexibleProtectionFlowResponseBodyFlows {
	s.AddFlow = &v
	return s
}

func (s *DescribeFlexibleProtectionFlowResponseBodyFlows) SetUseableFlow(v float32) *DescribeFlexibleProtectionFlowResponseBodyFlows {
	s.UseableFlow = &v
	return s
}

type DescribeFlexibleProtectionFlowResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFlexibleProtectionFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFlexibleProtectionFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlexibleProtectionFlowResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlexibleProtectionFlowResponse) SetHeaders(v map[string]*string) *DescribeFlexibleProtectionFlowResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlexibleProtectionFlowResponse) SetBody(v *DescribeFlexibleProtectionFlowResponseBody) *DescribeFlexibleProtectionFlowResponse {
	s.Body = v
	return s
}

type DescribeFlowgraphRequest struct {
	SourceIp         *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang             *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	EagleEyeTraceId  *string `json:"eagleEyeTraceId,omitempty" xml:"eagleEyeTraceId,omitempty"`
	EagleEyeRpcId    *string `json:"eagleEyeRpcId,omitempty" xml:"eagleEyeRpcId,omitempty"`
	EagleEyeUserData *string `json:"eagleEyeUserData,omitempty" xml:"eagleEyeUserData,omitempty"`
	DdosRegionId     *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
	InstanceType     *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Days             *int32  `json:"Days,omitempty" xml:"Days,omitempty"`
	StartTime        *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime          *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeFlowgraphRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowgraphRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowgraphRequest) SetSourceIp(v string) *DescribeFlowgraphRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeFlowgraphRequest) SetLang(v string) *DescribeFlowgraphRequest {
	s.Lang = &v
	return s
}

func (s *DescribeFlowgraphRequest) SetEagleEyeTraceId(v string) *DescribeFlowgraphRequest {
	s.EagleEyeTraceId = &v
	return s
}

func (s *DescribeFlowgraphRequest) SetEagleEyeRpcId(v string) *DescribeFlowgraphRequest {
	s.EagleEyeRpcId = &v
	return s
}

func (s *DescribeFlowgraphRequest) SetEagleEyeUserData(v string) *DescribeFlowgraphRequest {
	s.EagleEyeUserData = &v
	return s
}

func (s *DescribeFlowgraphRequest) SetDdosRegionId(v string) *DescribeFlowgraphRequest {
	s.DdosRegionId = &v
	return s
}

func (s *DescribeFlowgraphRequest) SetInstanceType(v string) *DescribeFlowgraphRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeFlowgraphRequest) SetInstanceId(v string) *DescribeFlowgraphRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeFlowgraphRequest) SetDays(v int32) *DescribeFlowgraphRequest {
	s.Days = &v
	return s
}

func (s *DescribeFlowgraphRequest) SetStartTime(v int64) *DescribeFlowgraphRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeFlowgraphRequest) SetEndTime(v int64) *DescribeFlowgraphRequest {
	s.EndTime = &v
	return s
}

type DescribeFlowgraphResponseBody struct {
	Flowgraphs []*DescribeFlowgraphResponseBodyFlowgraphs `json:"Flowgraphs,omitempty" xml:"Flowgraphs,omitempty" type:"Repeated"`
	RequestId  *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFlowgraphResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowgraphResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowgraphResponseBody) SetFlowgraphs(v []*DescribeFlowgraphResponseBodyFlowgraphs) *DescribeFlowgraphResponseBody {
	s.Flowgraphs = v
	return s
}

func (s *DescribeFlowgraphResponseBody) SetRequestId(v string) *DescribeFlowgraphResponseBody {
	s.RequestId = &v
	return s
}

type DescribeFlowgraphResponseBodyFlowgraphs struct {
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	Bps  *int64 `json:"Bps,omitempty" xml:"Bps,omitempty"`
	Pps  *int64 `json:"Pps,omitempty" xml:"Pps,omitempty"`
}

func (s DescribeFlowgraphResponseBodyFlowgraphs) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowgraphResponseBodyFlowgraphs) GoString() string {
	return s.String()
}

func (s *DescribeFlowgraphResponseBodyFlowgraphs) SetTime(v int64) *DescribeFlowgraphResponseBodyFlowgraphs {
	s.Time = &v
	return s
}

func (s *DescribeFlowgraphResponseBodyFlowgraphs) SetBps(v int64) *DescribeFlowgraphResponseBodyFlowgraphs {
	s.Bps = &v
	return s
}

func (s *DescribeFlowgraphResponseBodyFlowgraphs) SetPps(v int64) *DescribeFlowgraphResponseBodyFlowgraphs {
	s.Pps = &v
	return s
}

type DescribeFlowgraphResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFlowgraphResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFlowgraphResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowgraphResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowgraphResponse) SetHeaders(v map[string]*string) *DescribeFlowgraphResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowgraphResponse) SetBody(v *DescribeFlowgraphResponseBody) *DescribeFlowgraphResponse {
	s.Body = v
	return s
}

type DescribeRegionDdosThresholdRequest struct {
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
}

func (s DescribeRegionDdosThresholdRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionDdosThresholdRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionDdosThresholdRequest) SetSourceIp(v string) *DescribeRegionDdosThresholdRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeRegionDdosThresholdRequest) SetLang(v string) *DescribeRegionDdosThresholdRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRegionDdosThresholdRequest) SetDdosRegionId(v string) *DescribeRegionDdosThresholdRequest {
	s.DdosRegionId = &v
	return s
}

type DescribeRegionDdosThresholdResponseBody struct {
	RequestId           *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RegionDdosThreshold *DescribeRegionDdosThresholdResponseBodyRegionDdosThreshold `json:"RegionDdosThreshold,omitempty" xml:"RegionDdosThreshold,omitempty" type:"Struct"`
}

func (s DescribeRegionDdosThresholdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionDdosThresholdResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionDdosThresholdResponseBody) SetRequestId(v string) *DescribeRegionDdosThresholdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionDdosThresholdResponseBody) SetRegionDdosThreshold(v *DescribeRegionDdosThresholdResponseBodyRegionDdosThreshold) *DescribeRegionDdosThresholdResponseBody {
	s.RegionDdosThreshold = v
	return s
}

type DescribeRegionDdosThresholdResponseBodyRegionDdosThreshold struct {
	ElasticThreshold *int32 `json:"ElasticThreshold,omitempty" xml:"ElasticThreshold,omitempty"`
	BaseThreshold    *int32 `json:"BaseThreshold,omitempty" xml:"BaseThreshold,omitempty"`
}

func (s DescribeRegionDdosThresholdResponseBodyRegionDdosThreshold) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionDdosThresholdResponseBodyRegionDdosThreshold) GoString() string {
	return s.String()
}

func (s *DescribeRegionDdosThresholdResponseBodyRegionDdosThreshold) SetElasticThreshold(v int32) *DescribeRegionDdosThresholdResponseBodyRegionDdosThreshold {
	s.ElasticThreshold = &v
	return s
}

func (s *DescribeRegionDdosThresholdResponseBodyRegionDdosThreshold) SetBaseThreshold(v int32) *DescribeRegionDdosThresholdResponseBodyRegionDdosThreshold {
	s.BaseThreshold = &v
	return s
}

type DescribeRegionDdosThresholdResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRegionDdosThresholdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionDdosThresholdResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionDdosThresholdResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionDdosThresholdResponse) SetHeaders(v map[string]*string) *DescribeRegionDdosThresholdResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionDdosThresholdResponse) SetBody(v *DescribeRegionDdosThresholdResponseBody) *DescribeRegionDdosThresholdResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetSourceIp(v string) *DescribeRegionsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeRegionsRequest) SetLang(v string) *DescribeRegionsRequest {
	s.Lang = &v
	return s
}

type DescribeRegionsResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Regions   []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	RegionEnName  *string `json:"RegionEnName,omitempty" xml:"RegionEnName,omitempty"`
	RegionName    *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	RegionNo      *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	RegionNoAlias *string `json:"RegionNoAlias,omitempty" xml:"RegionNoAlias,omitempty"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionEnName(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionEnName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionName(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionNo(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionNo = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionNoAlias(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionNoAlias = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeTrafficInfoRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeTrafficInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrafficInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeTrafficInfoRequest) SetSourceIp(v string) *DescribeTrafficInfoRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeTrafficInfoRequest) SetLang(v string) *DescribeTrafficInfoRequest {
	s.Lang = &v
	return s
}

type DescribeTrafficInfoResponseBody struct {
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TrafficInfo *DescribeTrafficInfoResponseBodyTrafficInfo `json:"TrafficInfo,omitempty" xml:"TrafficInfo,omitempty" type:"Struct"`
}

func (s DescribeTrafficInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrafficInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTrafficInfoResponseBody) SetRequestId(v string) *DescribeTrafficInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTrafficInfoResponseBody) SetTrafficInfo(v *DescribeTrafficInfoResponseBodyTrafficInfo) *DescribeTrafficInfoResponseBody {
	s.TrafficInfo = v
	return s
}

type DescribeTrafficInfoResponseBodyTrafficInfo struct {
	LastUsedTraffic *int32 `json:"LastUsedTraffic,omitempty" xml:"LastUsedTraffic,omitempty"`
	AddTraffic      *int32 `json:"AddTraffic,omitempty" xml:"AddTraffic,omitempty"`
	UsableTraffic   *int32 `json:"UsableTraffic,omitempty" xml:"UsableTraffic,omitempty"`
}

func (s DescribeTrafficInfoResponseBodyTrafficInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrafficInfoResponseBodyTrafficInfo) GoString() string {
	return s.String()
}

func (s *DescribeTrafficInfoResponseBodyTrafficInfo) SetLastUsedTraffic(v int32) *DescribeTrafficInfoResponseBodyTrafficInfo {
	s.LastUsedTraffic = &v
	return s
}

func (s *DescribeTrafficInfoResponseBodyTrafficInfo) SetAddTraffic(v int32) *DescribeTrafficInfoResponseBodyTrafficInfo {
	s.AddTraffic = &v
	return s
}

func (s *DescribeTrafficInfoResponseBodyTrafficInfo) SetUsableTraffic(v int32) *DescribeTrafficInfoResponseBodyTrafficInfo {
	s.UsableTraffic = &v
	return s
}

type DescribeTrafficInfoResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTrafficInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTrafficInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrafficInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeTrafficInfoResponse) SetHeaders(v map[string]*string) *DescribeTrafficInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeTrafficInfoResponse) SetBody(v *DescribeTrafficInfoResponseBody) *DescribeTrafficInfoResponse {
	s.Body = v
	return s
}

type ModifyDdosStatusRequest struct {
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ModifyDdosStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDdosStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyDdosStatusRequest) SetSourceIp(v string) *ModifyDdosStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyDdosStatusRequest) SetLang(v string) *ModifyDdosStatusRequest {
	s.Lang = &v
	return s
}

func (s *ModifyDdosStatusRequest) SetDdosRegionId(v string) *ModifyDdosStatusRequest {
	s.DdosRegionId = &v
	return s
}

func (s *ModifyDdosStatusRequest) SetInstanceType(v string) *ModifyDdosStatusRequest {
	s.InstanceType = &v
	return s
}

func (s *ModifyDdosStatusRequest) SetInstanceId(v string) *ModifyDdosStatusRequest {
	s.InstanceId = &v
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDdosStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyDdosStatusResponse) SetBody(v *ModifyDdosStatusResponseBody) *ModifyDdosStatusResponse {
	s.Body = v
	return s
}

type ModifyDefenseThresholdRequest struct {
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Bps          *int32  `json:"Bps,omitempty" xml:"Bps,omitempty"`
	Pps          *int32  `json:"Pps,omitempty" xml:"Pps,omitempty"`
	IsAuto       *bool   `json:"IsAuto,omitempty" xml:"IsAuto,omitempty"`
}

func (s ModifyDefenseThresholdRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseThresholdRequest) GoString() string {
	return s.String()
}

func (s *ModifyDefenseThresholdRequest) SetSourceIp(v string) *ModifyDefenseThresholdRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyDefenseThresholdRequest) SetLang(v string) *ModifyDefenseThresholdRequest {
	s.Lang = &v
	return s
}

func (s *ModifyDefenseThresholdRequest) SetDdosRegionId(v string) *ModifyDefenseThresholdRequest {
	s.DdosRegionId = &v
	return s
}

func (s *ModifyDefenseThresholdRequest) SetInstanceType(v string) *ModifyDefenseThresholdRequest {
	s.InstanceType = &v
	return s
}

func (s *ModifyDefenseThresholdRequest) SetInstanceId(v string) *ModifyDefenseThresholdRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDefenseThresholdRequest) SetBps(v int32) *ModifyDefenseThresholdRequest {
	s.Bps = &v
	return s
}

func (s *ModifyDefenseThresholdRequest) SetPps(v int32) *ModifyDefenseThresholdRequest {
	s.Pps = &v
	return s
}

func (s *ModifyDefenseThresholdRequest) SetIsAuto(v bool) *ModifyDefenseThresholdRequest {
	s.IsAuto = &v
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDefenseThresholdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBgpPackByIpResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBgpPackByIp"), tea.String("2017-05-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeBgpPackElasticThresholdWithOptions(request *DescribeBgpPackElasticThresholdRequest, runtime *util.RuntimeOptions) (_result *DescribeBgpPackElasticThresholdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBgpPackElasticThresholdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBgpPackElasticThreshold"), tea.String("2017-05-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBgpPackElasticThreshold(request *DescribeBgpPackElasticThresholdRequest) (_result *DescribeBgpPackElasticThresholdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBgpPackElasticThresholdResponse{}
	_body, _err := client.DescribeBgpPackElasticThresholdWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCapResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCap"), tea.String("2017-05-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeCreditInfoWithOptions(request *DescribeCreditInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeCreditInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCreditInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCreditInfo"), tea.String("2017-05-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCreditInfo(request *DescribeCreditInfoRequest) (_result *DescribeCreditInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCreditInfoResponse{}
	_body, _err := client.DescribeCreditInfoWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDdosCountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDdosCount"), tea.String("2017-05-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDdosCreditResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDdosCredit"), tea.String("2017-05-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDdosEventListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDdosEventList"), tea.String("2017-05-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDdosThresholdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDdosThreshold"), tea.String("2017-05-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeFlexibleProtectionFlowWithOptions(request *DescribeFlexibleProtectionFlowRequest, runtime *util.RuntimeOptions) (_result *DescribeFlexibleProtectionFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFlexibleProtectionFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFlexibleProtectionFlow"), tea.String("2017-05-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFlexibleProtectionFlow(request *DescribeFlexibleProtectionFlowRequest) (_result *DescribeFlexibleProtectionFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFlexibleProtectionFlowResponse{}
	_body, _err := client.DescribeFlexibleProtectionFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFlowgraphWithOptions(request *DescribeFlowgraphRequest, runtime *util.RuntimeOptions) (_result *DescribeFlowgraphResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFlowgraphResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFlowgraph"), tea.String("2017-05-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFlowgraph(request *DescribeFlowgraphRequest) (_result *DescribeFlowgraphResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFlowgraphResponse{}
	_body, _err := client.DescribeFlowgraphWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionDdosThresholdWithOptions(request *DescribeRegionDdosThresholdRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionDdosThresholdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRegionDdosThresholdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRegionDdosThreshold"), tea.String("2017-05-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegionDdosThreshold(request *DescribeRegionDdosThresholdRequest) (_result *DescribeRegionDdosThresholdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionDdosThresholdResponse{}
	_body, _err := client.DescribeRegionDdosThresholdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRegions"), tea.String("2017-05-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTrafficInfoWithOptions(request *DescribeTrafficInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeTrafficInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeTrafficInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeTrafficInfo"), tea.String("2017-05-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTrafficInfo(request *DescribeTrafficInfoRequest) (_result *DescribeTrafficInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTrafficInfoResponse{}
	_body, _err := client.DescribeTrafficInfoWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDdosStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDdosStatus"), tea.String("2017-05-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDefenseThresholdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDefenseThreshold"), tea.String("2017-05-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
