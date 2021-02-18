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

type DescribeGeoipInstanceRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeGeoipInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGeoipInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeGeoipInstanceRequest) SetLang(v string) *DescribeGeoipInstanceRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGeoipInstanceRequest) SetUserClientIp(v string) *DescribeGeoipInstanceRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeGeoipInstanceRequest) SetInstanceId(v string) *DescribeGeoipInstanceRequest {
	s.InstanceId = &v
	return s
}

type DescribeGeoipInstanceResponseBody struct {
	ExpireTimestamp *int64  `json:"ExpireTimestamp,omitempty" xml:"ExpireTimestamp,omitempty"`
	VersionCode     *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	MaxQpd          *int64  `json:"MaxQpd,omitempty" xml:"MaxQpd,omitempty"`
	MaxQps          *int64  `json:"MaxQps,omitempty" xml:"MaxQps,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ProductCode     *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	QueryCount      *int64  `json:"QueryCount,omitempty" xml:"QueryCount,omitempty"`
	ExpireTime      *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	CreateTimestamp *int64  `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
}

func (s DescribeGeoipInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGeoipInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGeoipInstanceResponseBody) SetExpireTimestamp(v int64) *DescribeGeoipInstanceResponseBody {
	s.ExpireTimestamp = &v
	return s
}

func (s *DescribeGeoipInstanceResponseBody) SetVersionCode(v string) *DescribeGeoipInstanceResponseBody {
	s.VersionCode = &v
	return s
}

func (s *DescribeGeoipInstanceResponseBody) SetMaxQpd(v int64) *DescribeGeoipInstanceResponseBody {
	s.MaxQpd = &v
	return s
}

func (s *DescribeGeoipInstanceResponseBody) SetMaxQps(v int64) *DescribeGeoipInstanceResponseBody {
	s.MaxQps = &v
	return s
}

func (s *DescribeGeoipInstanceResponseBody) SetRequestId(v string) *DescribeGeoipInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGeoipInstanceResponseBody) SetInstanceId(v string) *DescribeGeoipInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeGeoipInstanceResponseBody) SetProductCode(v string) *DescribeGeoipInstanceResponseBody {
	s.ProductCode = &v
	return s
}

func (s *DescribeGeoipInstanceResponseBody) SetCreateTime(v string) *DescribeGeoipInstanceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeGeoipInstanceResponseBody) SetQueryCount(v int64) *DescribeGeoipInstanceResponseBody {
	s.QueryCount = &v
	return s
}

func (s *DescribeGeoipInstanceResponseBody) SetExpireTime(v string) *DescribeGeoipInstanceResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeGeoipInstanceResponseBody) SetCreateTimestamp(v int64) *DescribeGeoipInstanceResponseBody {
	s.CreateTimestamp = &v
	return s
}

type DescribeGeoipInstanceResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGeoipInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGeoipInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGeoipInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeGeoipInstanceResponse) SetHeaders(v map[string]*string) *DescribeGeoipInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeGeoipInstanceResponse) SetBody(v *DescribeGeoipInstanceResponseBody) *DescribeGeoipInstanceResponse {
	s.Body = v
	return s
}

type DescribeGeoipInstanceDataInfosRequest struct {
	Lang             *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp     *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LocationDataType *string `json:"LocationDataType,omitempty" xml:"LocationDataType,omitempty"`
}

func (s DescribeGeoipInstanceDataInfosRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGeoipInstanceDataInfosRequest) GoString() string {
	return s.String()
}

func (s *DescribeGeoipInstanceDataInfosRequest) SetLang(v string) *DescribeGeoipInstanceDataInfosRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGeoipInstanceDataInfosRequest) SetUserClientIp(v string) *DescribeGeoipInstanceDataInfosRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeGeoipInstanceDataInfosRequest) SetInstanceId(v string) *DescribeGeoipInstanceDataInfosRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeGeoipInstanceDataInfosRequest) SetLocationDataType(v string) *DescribeGeoipInstanceDataInfosRequest {
	s.LocationDataType = &v
	return s
}

type DescribeGeoipInstanceDataInfosResponseBody struct {
	RequestId *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DataInfos *DescribeGeoipInstanceDataInfosResponseBodyDataInfos `json:"DataInfos,omitempty" xml:"DataInfos,omitempty" type:"Struct"`
}

func (s DescribeGeoipInstanceDataInfosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGeoipInstanceDataInfosResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGeoipInstanceDataInfosResponseBody) SetRequestId(v string) *DescribeGeoipInstanceDataInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGeoipInstanceDataInfosResponseBody) SetDataInfos(v *DescribeGeoipInstanceDataInfosResponseBodyDataInfos) *DescribeGeoipInstanceDataInfosResponseBody {
	s.DataInfos = v
	return s
}

type DescribeGeoipInstanceDataInfosResponseBodyDataInfos struct {
	DataInfo []*DescribeGeoipInstanceDataInfosResponseBodyDataInfosDataInfo `json:"DataInfo,omitempty" xml:"DataInfo,omitempty" type:"Repeated"`
}

func (s DescribeGeoipInstanceDataInfosResponseBodyDataInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeGeoipInstanceDataInfosResponseBodyDataInfos) GoString() string {
	return s.String()
}

func (s *DescribeGeoipInstanceDataInfosResponseBodyDataInfos) SetDataInfo(v []*DescribeGeoipInstanceDataInfosResponseBodyDataInfosDataInfo) *DescribeGeoipInstanceDataInfosResponseBodyDataInfos {
	s.DataInfo = v
	return s
}

type DescribeGeoipInstanceDataInfosResponseBodyDataInfosDataInfo struct {
	Type            *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateTimestamp *int64  `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	UpdateTime      *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Version         *string `json:"Version,omitempty" xml:"Version,omitempty"`
	DownloadCount   *int64  `json:"DownloadCount,omitempty" xml:"DownloadCount,omitempty"`
}

func (s DescribeGeoipInstanceDataInfosResponseBodyDataInfosDataInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeGeoipInstanceDataInfosResponseBodyDataInfosDataInfo) GoString() string {
	return s.String()
}

func (s *DescribeGeoipInstanceDataInfosResponseBodyDataInfosDataInfo) SetType(v string) *DescribeGeoipInstanceDataInfosResponseBodyDataInfosDataInfo {
	s.Type = &v
	return s
}

func (s *DescribeGeoipInstanceDataInfosResponseBodyDataInfosDataInfo) SetUpdateTimestamp(v int64) *DescribeGeoipInstanceDataInfosResponseBodyDataInfosDataInfo {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeGeoipInstanceDataInfosResponseBodyDataInfosDataInfo) SetUpdateTime(v string) *DescribeGeoipInstanceDataInfosResponseBodyDataInfosDataInfo {
	s.UpdateTime = &v
	return s
}

func (s *DescribeGeoipInstanceDataInfosResponseBodyDataInfosDataInfo) SetVersion(v string) *DescribeGeoipInstanceDataInfosResponseBodyDataInfosDataInfo {
	s.Version = &v
	return s
}

func (s *DescribeGeoipInstanceDataInfosResponseBodyDataInfosDataInfo) SetDownloadCount(v int64) *DescribeGeoipInstanceDataInfosResponseBodyDataInfosDataInfo {
	s.DownloadCount = &v
	return s
}

type DescribeGeoipInstanceDataInfosResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGeoipInstanceDataInfosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGeoipInstanceDataInfosResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGeoipInstanceDataInfosResponse) GoString() string {
	return s.String()
}

func (s *DescribeGeoipInstanceDataInfosResponse) SetHeaders(v map[string]*string) *DescribeGeoipInstanceDataInfosResponse {
	s.Headers = v
	return s
}

func (s *DescribeGeoipInstanceDataInfosResponse) SetBody(v *DescribeGeoipInstanceDataInfosResponseBody) *DescribeGeoipInstanceDataInfosResponse {
	s.Body = v
	return s
}

type DescribeGeoipInstanceDataUrlRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DataType     *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
}

func (s DescribeGeoipInstanceDataUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGeoipInstanceDataUrlRequest) GoString() string {
	return s.String()
}

func (s *DescribeGeoipInstanceDataUrlRequest) SetLang(v string) *DescribeGeoipInstanceDataUrlRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGeoipInstanceDataUrlRequest) SetUserClientIp(v string) *DescribeGeoipInstanceDataUrlRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeGeoipInstanceDataUrlRequest) SetInstanceId(v string) *DescribeGeoipInstanceDataUrlRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeGeoipInstanceDataUrlRequest) SetDataType(v string) *DescribeGeoipInstanceDataUrlRequest {
	s.DataType = &v
	return s
}

type DescribeGeoipInstanceDataUrlResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
}

func (s DescribeGeoipInstanceDataUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGeoipInstanceDataUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGeoipInstanceDataUrlResponseBody) SetRequestId(v string) *DescribeGeoipInstanceDataUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGeoipInstanceDataUrlResponseBody) SetDownloadUrl(v string) *DescribeGeoipInstanceDataUrlResponseBody {
	s.DownloadUrl = &v
	return s
}

type DescribeGeoipInstanceDataUrlResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGeoipInstanceDataUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGeoipInstanceDataUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGeoipInstanceDataUrlResponse) GoString() string {
	return s.String()
}

func (s *DescribeGeoipInstanceDataUrlResponse) SetHeaders(v map[string]*string) *DescribeGeoipInstanceDataUrlResponse {
	s.Headers = v
	return s
}

func (s *DescribeGeoipInstanceDataUrlResponse) SetBody(v *DescribeGeoipInstanceDataUrlResponseBody) *DescribeGeoipInstanceDataUrlResponse {
	s.Body = v
	return s
}

type DescribeGeoipInstancesRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s DescribeGeoipInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGeoipInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeGeoipInstancesRequest) SetLang(v string) *DescribeGeoipInstancesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGeoipInstancesRequest) SetUserClientIp(v string) *DescribeGeoipInstancesRequest {
	s.UserClientIp = &v
	return s
}

type DescribeGeoipInstancesResponseBody struct {
	RequestId      *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	GeoipInstances *DescribeGeoipInstancesResponseBodyGeoipInstances `json:"GeoipInstances,omitempty" xml:"GeoipInstances,omitempty" type:"Struct"`
}

func (s DescribeGeoipInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGeoipInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGeoipInstancesResponseBody) SetRequestId(v string) *DescribeGeoipInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGeoipInstancesResponseBody) SetGeoipInstances(v *DescribeGeoipInstancesResponseBodyGeoipInstances) *DescribeGeoipInstancesResponseBody {
	s.GeoipInstances = v
	return s
}

type DescribeGeoipInstancesResponseBodyGeoipInstances struct {
	GeoipInstance []*DescribeGeoipInstancesResponseBodyGeoipInstancesGeoipInstance `json:"GeoipInstance,omitempty" xml:"GeoipInstance,omitempty" type:"Repeated"`
}

func (s DescribeGeoipInstancesResponseBodyGeoipInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeGeoipInstancesResponseBodyGeoipInstances) GoString() string {
	return s.String()
}

func (s *DescribeGeoipInstancesResponseBodyGeoipInstances) SetGeoipInstance(v []*DescribeGeoipInstancesResponseBodyGeoipInstancesGeoipInstance) *DescribeGeoipInstancesResponseBodyGeoipInstances {
	s.GeoipInstance = v
	return s
}

type DescribeGeoipInstancesResponseBodyGeoipInstancesGeoipInstance struct {
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ExpireTimestamp *int64  `json:"ExpireTimestamp,omitempty" xml:"ExpireTimestamp,omitempty"`
	ExpireTime      *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	MaxQps          *int64  `json:"MaxQps,omitempty" xml:"MaxQps,omitempty"`
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	MaxQpd          *int64  `json:"MaxQpd,omitempty" xml:"MaxQpd,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	VersionCode     *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	CreateTimestamp *int64  `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	ProductCode     *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s DescribeGeoipInstancesResponseBodyGeoipInstancesGeoipInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeGeoipInstancesResponseBodyGeoipInstancesGeoipInstance) GoString() string {
	return s.String()
}

func (s *DescribeGeoipInstancesResponseBodyGeoipInstancesGeoipInstance) SetStatus(v string) *DescribeGeoipInstancesResponseBodyGeoipInstancesGeoipInstance {
	s.Status = &v
	return s
}

func (s *DescribeGeoipInstancesResponseBodyGeoipInstancesGeoipInstance) SetExpireTimestamp(v int64) *DescribeGeoipInstancesResponseBodyGeoipInstancesGeoipInstance {
	s.ExpireTimestamp = &v
	return s
}

func (s *DescribeGeoipInstancesResponseBodyGeoipInstancesGeoipInstance) SetExpireTime(v string) *DescribeGeoipInstancesResponseBodyGeoipInstancesGeoipInstance {
	s.ExpireTime = &v
	return s
}

func (s *DescribeGeoipInstancesResponseBodyGeoipInstancesGeoipInstance) SetMaxQps(v int64) *DescribeGeoipInstancesResponseBodyGeoipInstancesGeoipInstance {
	s.MaxQps = &v
	return s
}

func (s *DescribeGeoipInstancesResponseBodyGeoipInstancesGeoipInstance) SetCreateTime(v string) *DescribeGeoipInstancesResponseBodyGeoipInstancesGeoipInstance {
	s.CreateTime = &v
	return s
}

func (s *DescribeGeoipInstancesResponseBodyGeoipInstancesGeoipInstance) SetMaxQpd(v int64) *DescribeGeoipInstancesResponseBodyGeoipInstancesGeoipInstance {
	s.MaxQpd = &v
	return s
}

func (s *DescribeGeoipInstancesResponseBodyGeoipInstancesGeoipInstance) SetInstanceId(v string) *DescribeGeoipInstancesResponseBodyGeoipInstancesGeoipInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeGeoipInstancesResponseBodyGeoipInstancesGeoipInstance) SetVersionCode(v string) *DescribeGeoipInstancesResponseBodyGeoipInstancesGeoipInstance {
	s.VersionCode = &v
	return s
}

func (s *DescribeGeoipInstancesResponseBodyGeoipInstancesGeoipInstance) SetCreateTimestamp(v int64) *DescribeGeoipInstancesResponseBodyGeoipInstancesGeoipInstance {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeGeoipInstancesResponseBodyGeoipInstancesGeoipInstance) SetProductCode(v string) *DescribeGeoipInstancesResponseBodyGeoipInstancesGeoipInstance {
	s.ProductCode = &v
	return s
}

type DescribeGeoipInstancesResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGeoipInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGeoipInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGeoipInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeGeoipInstancesResponse) SetHeaders(v map[string]*string) *DescribeGeoipInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeGeoipInstancesResponse) SetBody(v *DescribeGeoipInstancesResponseBody) *DescribeGeoipInstancesResponse {
	s.Body = v
	return s
}

type DescribeGeoipInstanceStatisticsRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	StartDate    *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	EndDate      *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
}

func (s DescribeGeoipInstanceStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGeoipInstanceStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeGeoipInstanceStatisticsRequest) SetLang(v string) *DescribeGeoipInstanceStatisticsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGeoipInstanceStatisticsRequest) SetUserClientIp(v string) *DescribeGeoipInstanceStatisticsRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeGeoipInstanceStatisticsRequest) SetInstanceId(v string) *DescribeGeoipInstanceStatisticsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeGeoipInstanceStatisticsRequest) SetStartDate(v string) *DescribeGeoipInstanceStatisticsRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeGeoipInstanceStatisticsRequest) SetEndDate(v string) *DescribeGeoipInstanceStatisticsRequest {
	s.EndDate = &v
	return s
}

type DescribeGeoipInstanceStatisticsResponseBody struct {
	RequestId  *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Statistics *DescribeGeoipInstanceStatisticsResponseBodyStatistics `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Struct"`
}

func (s DescribeGeoipInstanceStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGeoipInstanceStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGeoipInstanceStatisticsResponseBody) SetRequestId(v string) *DescribeGeoipInstanceStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGeoipInstanceStatisticsResponseBody) SetStatistics(v *DescribeGeoipInstanceStatisticsResponseBodyStatistics) *DescribeGeoipInstanceStatisticsResponseBody {
	s.Statistics = v
	return s
}

type DescribeGeoipInstanceStatisticsResponseBodyStatistics struct {
	Statistic []*DescribeGeoipInstanceStatisticsResponseBodyStatisticsStatistic `json:"Statistic,omitempty" xml:"Statistic,omitempty" type:"Repeated"`
}

func (s DescribeGeoipInstanceStatisticsResponseBodyStatistics) String() string {
	return tea.Prettify(s)
}

func (s DescribeGeoipInstanceStatisticsResponseBodyStatistics) GoString() string {
	return s.String()
}

func (s *DescribeGeoipInstanceStatisticsResponseBodyStatistics) SetStatistic(v []*DescribeGeoipInstanceStatisticsResponseBodyStatisticsStatistic) *DescribeGeoipInstanceStatisticsResponseBodyStatistics {
	s.Statistic = v
	return s
}

type DescribeGeoipInstanceStatisticsResponseBodyStatisticsStatistic struct {
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Count     *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeGeoipInstanceStatisticsResponseBodyStatisticsStatistic) String() string {
	return tea.Prettify(s)
}

func (s DescribeGeoipInstanceStatisticsResponseBodyStatisticsStatistic) GoString() string {
	return s.String()
}

func (s *DescribeGeoipInstanceStatisticsResponseBodyStatisticsStatistic) SetTimestamp(v int64) *DescribeGeoipInstanceStatisticsResponseBodyStatisticsStatistic {
	s.Timestamp = &v
	return s
}

func (s *DescribeGeoipInstanceStatisticsResponseBodyStatisticsStatistic) SetCount(v int64) *DescribeGeoipInstanceStatisticsResponseBodyStatisticsStatistic {
	s.Count = &v
	return s
}

type DescribeGeoipInstanceStatisticsResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGeoipInstanceStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGeoipInstanceStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGeoipInstanceStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeGeoipInstanceStatisticsResponse) SetHeaders(v map[string]*string) *DescribeGeoipInstanceStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeGeoipInstanceStatisticsResponse) SetBody(v *DescribeGeoipInstanceStatisticsResponseBody) *DescribeGeoipInstanceStatisticsResponse {
	s.Body = v
	return s
}

type DescribeIpv4LocationRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	Ip           *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribeIpv4LocationRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpv4LocationRequest) GoString() string {
	return s.String()
}

func (s *DescribeIpv4LocationRequest) SetLang(v string) *DescribeIpv4LocationRequest {
	s.Lang = &v
	return s
}

func (s *DescribeIpv4LocationRequest) SetUserClientIp(v string) *DescribeIpv4LocationRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeIpv4LocationRequest) SetIp(v string) *DescribeIpv4LocationRequest {
	s.Ip = &v
	return s
}

type DescribeIpv4LocationResponseBody struct {
	ProvinceEn  *string `json:"ProvinceEn,omitempty" xml:"ProvinceEn,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CityEn      *string `json:"CityEn,omitempty" xml:"CityEn,omitempty"`
	Ip          *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Isp         *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	Latitude    *string `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	City        *string `json:"City,omitempty" xml:"City,omitempty"`
	County      *string `json:"County,omitempty" xml:"County,omitempty"`
	Longitude   *string `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
	CountryEn   *string `json:"CountryEn,omitempty" xml:"CountryEn,omitempty"`
	Province    *string `json:"Province,omitempty" xml:"Province,omitempty"`
	Country     *string `json:"Country,omitempty" xml:"Country,omitempty"`
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
}

func (s DescribeIpv4LocationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpv4LocationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIpv4LocationResponseBody) SetProvinceEn(v string) *DescribeIpv4LocationResponseBody {
	s.ProvinceEn = &v
	return s
}

func (s *DescribeIpv4LocationResponseBody) SetRequestId(v string) *DescribeIpv4LocationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIpv4LocationResponseBody) SetCityEn(v string) *DescribeIpv4LocationResponseBody {
	s.CityEn = &v
	return s
}

func (s *DescribeIpv4LocationResponseBody) SetIp(v string) *DescribeIpv4LocationResponseBody {
	s.Ip = &v
	return s
}

func (s *DescribeIpv4LocationResponseBody) SetIsp(v string) *DescribeIpv4LocationResponseBody {
	s.Isp = &v
	return s
}

func (s *DescribeIpv4LocationResponseBody) SetLatitude(v string) *DescribeIpv4LocationResponseBody {
	s.Latitude = &v
	return s
}

func (s *DescribeIpv4LocationResponseBody) SetCity(v string) *DescribeIpv4LocationResponseBody {
	s.City = &v
	return s
}

func (s *DescribeIpv4LocationResponseBody) SetCounty(v string) *DescribeIpv4LocationResponseBody {
	s.County = &v
	return s
}

func (s *DescribeIpv4LocationResponseBody) SetLongitude(v string) *DescribeIpv4LocationResponseBody {
	s.Longitude = &v
	return s
}

func (s *DescribeIpv4LocationResponseBody) SetCountryEn(v string) *DescribeIpv4LocationResponseBody {
	s.CountryEn = &v
	return s
}

func (s *DescribeIpv4LocationResponseBody) SetProvince(v string) *DescribeIpv4LocationResponseBody {
	s.Province = &v
	return s
}

func (s *DescribeIpv4LocationResponseBody) SetCountry(v string) *DescribeIpv4LocationResponseBody {
	s.Country = &v
	return s
}

func (s *DescribeIpv4LocationResponseBody) SetCountryCode(v string) *DescribeIpv4LocationResponseBody {
	s.CountryCode = &v
	return s
}

type DescribeIpv4LocationResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeIpv4LocationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeIpv4LocationResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpv4LocationResponse) GoString() string {
	return s.String()
}

func (s *DescribeIpv4LocationResponse) SetHeaders(v map[string]*string) *DescribeIpv4LocationResponse {
	s.Headers = v
	return s
}

func (s *DescribeIpv4LocationResponse) SetBody(v *DescribeIpv4LocationResponseBody) *DescribeIpv4LocationResponse {
	s.Body = v
	return s
}

type DescribeIpv6LocationRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	Ip           *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribeIpv6LocationRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpv6LocationRequest) GoString() string {
	return s.String()
}

func (s *DescribeIpv6LocationRequest) SetLang(v string) *DescribeIpv6LocationRequest {
	s.Lang = &v
	return s
}

func (s *DescribeIpv6LocationRequest) SetUserClientIp(v string) *DescribeIpv6LocationRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeIpv6LocationRequest) SetIp(v string) *DescribeIpv6LocationRequest {
	s.Ip = &v
	return s
}

type DescribeIpv6LocationResponseBody struct {
	ProvinceEn  *string `json:"ProvinceEn,omitempty" xml:"ProvinceEn,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CityEn      *string `json:"CityEn,omitempty" xml:"CityEn,omitempty"`
	Ip          *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Isp         *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	Latitude    *string `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	City        *string `json:"City,omitempty" xml:"City,omitempty"`
	County      *string `json:"County,omitempty" xml:"County,omitempty"`
	Longitude   *string `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
	CountryEn   *string `json:"CountryEn,omitempty" xml:"CountryEn,omitempty"`
	Province    *string `json:"Province,omitempty" xml:"Province,omitempty"`
	Country     *string `json:"Country,omitempty" xml:"Country,omitempty"`
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
}

func (s DescribeIpv6LocationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpv6LocationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIpv6LocationResponseBody) SetProvinceEn(v string) *DescribeIpv6LocationResponseBody {
	s.ProvinceEn = &v
	return s
}

func (s *DescribeIpv6LocationResponseBody) SetRequestId(v string) *DescribeIpv6LocationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIpv6LocationResponseBody) SetCityEn(v string) *DescribeIpv6LocationResponseBody {
	s.CityEn = &v
	return s
}

func (s *DescribeIpv6LocationResponseBody) SetIp(v string) *DescribeIpv6LocationResponseBody {
	s.Ip = &v
	return s
}

func (s *DescribeIpv6LocationResponseBody) SetIsp(v string) *DescribeIpv6LocationResponseBody {
	s.Isp = &v
	return s
}

func (s *DescribeIpv6LocationResponseBody) SetLatitude(v string) *DescribeIpv6LocationResponseBody {
	s.Latitude = &v
	return s
}

func (s *DescribeIpv6LocationResponseBody) SetCity(v string) *DescribeIpv6LocationResponseBody {
	s.City = &v
	return s
}

func (s *DescribeIpv6LocationResponseBody) SetCounty(v string) *DescribeIpv6LocationResponseBody {
	s.County = &v
	return s
}

func (s *DescribeIpv6LocationResponseBody) SetLongitude(v string) *DescribeIpv6LocationResponseBody {
	s.Longitude = &v
	return s
}

func (s *DescribeIpv6LocationResponseBody) SetCountryEn(v string) *DescribeIpv6LocationResponseBody {
	s.CountryEn = &v
	return s
}

func (s *DescribeIpv6LocationResponseBody) SetProvince(v string) *DescribeIpv6LocationResponseBody {
	s.Province = &v
	return s
}

func (s *DescribeIpv6LocationResponseBody) SetCountry(v string) *DescribeIpv6LocationResponseBody {
	s.Country = &v
	return s
}

func (s *DescribeIpv6LocationResponseBody) SetCountryCode(v string) *DescribeIpv6LocationResponseBody {
	s.CountryCode = &v
	return s
}

type DescribeIpv6LocationResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeIpv6LocationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeIpv6LocationResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpv6LocationResponse) GoString() string {
	return s.String()
}

func (s *DescribeIpv6LocationResponse) SetHeaders(v map[string]*string) *DescribeIpv6LocationResponse {
	s.Headers = v
	return s
}

func (s *DescribeIpv6LocationResponse) SetBody(v *DescribeIpv6LocationResponseBody) *DescribeIpv6LocationResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("geoip"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) DescribeGeoipInstanceWithOptions(request *DescribeGeoipInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeGeoipInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeGeoipInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeGeoipInstance"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGeoipInstance(request *DescribeGeoipInstanceRequest) (_result *DescribeGeoipInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGeoipInstanceResponse{}
	_body, _err := client.DescribeGeoipInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGeoipInstanceDataInfosWithOptions(request *DescribeGeoipInstanceDataInfosRequest, runtime *util.RuntimeOptions) (_result *DescribeGeoipInstanceDataInfosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeGeoipInstanceDataInfosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeGeoipInstanceDataInfos"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGeoipInstanceDataInfos(request *DescribeGeoipInstanceDataInfosRequest) (_result *DescribeGeoipInstanceDataInfosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGeoipInstanceDataInfosResponse{}
	_body, _err := client.DescribeGeoipInstanceDataInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGeoipInstanceDataUrlWithOptions(request *DescribeGeoipInstanceDataUrlRequest, runtime *util.RuntimeOptions) (_result *DescribeGeoipInstanceDataUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeGeoipInstanceDataUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeGeoipInstanceDataUrl"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGeoipInstanceDataUrl(request *DescribeGeoipInstanceDataUrlRequest) (_result *DescribeGeoipInstanceDataUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGeoipInstanceDataUrlResponse{}
	_body, _err := client.DescribeGeoipInstanceDataUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGeoipInstancesWithOptions(request *DescribeGeoipInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeGeoipInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeGeoipInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeGeoipInstances"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGeoipInstances(request *DescribeGeoipInstancesRequest) (_result *DescribeGeoipInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGeoipInstancesResponse{}
	_body, _err := client.DescribeGeoipInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGeoipInstanceStatisticsWithOptions(request *DescribeGeoipInstanceStatisticsRequest, runtime *util.RuntimeOptions) (_result *DescribeGeoipInstanceStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeGeoipInstanceStatisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeGeoipInstanceStatistics"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGeoipInstanceStatistics(request *DescribeGeoipInstanceStatisticsRequest) (_result *DescribeGeoipInstanceStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGeoipInstanceStatisticsResponse{}
	_body, _err := client.DescribeGeoipInstanceStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeIpv4LocationWithOptions(request *DescribeIpv4LocationRequest, runtime *util.RuntimeOptions) (_result *DescribeIpv4LocationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeIpv4LocationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeIpv4Location"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeIpv4Location(request *DescribeIpv4LocationRequest) (_result *DescribeIpv4LocationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeIpv4LocationResponse{}
	_body, _err := client.DescribeIpv4LocationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeIpv6LocationWithOptions(request *DescribeIpv6LocationRequest, runtime *util.RuntimeOptions) (_result *DescribeIpv6LocationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeIpv6LocationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeIpv6Location"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeIpv6Location(request *DescribeIpv6LocationRequest) (_result *DescribeIpv6LocationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeIpv6LocationResponse{}
	_body, _err := client.DescribeIpv6LocationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
