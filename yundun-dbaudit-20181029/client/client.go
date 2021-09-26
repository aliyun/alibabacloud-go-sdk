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

type ClearInstanceStorageRequest struct {
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StorageSpace    *string `json:"StorageSpace,omitempty" xml:"StorageSpace,omitempty"`
	StorageCategory *string `json:"StorageCategory,omitempty" xml:"StorageCategory,omitempty"`
}

func (s ClearInstanceStorageRequest) String() string {
	return tea.Prettify(s)
}

func (s ClearInstanceStorageRequest) GoString() string {
	return s.String()
}

func (s *ClearInstanceStorageRequest) SetInstanceId(v string) *ClearInstanceStorageRequest {
	s.InstanceId = &v
	return s
}

func (s *ClearInstanceStorageRequest) SetRegionId(v string) *ClearInstanceStorageRequest {
	s.RegionId = &v
	return s
}

func (s *ClearInstanceStorageRequest) SetStorageSpace(v string) *ClearInstanceStorageRequest {
	s.StorageSpace = &v
	return s
}

func (s *ClearInstanceStorageRequest) SetStorageCategory(v string) *ClearInstanceStorageRequest {
	s.StorageCategory = &v
	return s
}

type ClearInstanceStorageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClearInstanceStorageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ClearInstanceStorageResponseBody) GoString() string {
	return s.String()
}

func (s *ClearInstanceStorageResponseBody) SetRequestId(v string) *ClearInstanceStorageResponseBody {
	s.RequestId = &v
	return s
}

type ClearInstanceStorageResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ClearInstanceStorageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ClearInstanceStorageResponse) String() string {
	return tea.Prettify(s)
}

func (s ClearInstanceStorageResponse) GoString() string {
	return s.String()
}

func (s *ClearInstanceStorageResponse) SetHeaders(v map[string]*string) *ClearInstanceStorageResponse {
	s.Headers = v
	return s
}

func (s *ClearInstanceStorageResponse) SetBody(v *ClearInstanceStorageResponseBody) *ClearInstanceStorageResponse {
	s.Body = v
	return s
}

type ConfigInstanceWhiteListRequest struct {
	InstanceId *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	IpVersion  *string   `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	WhiteList  []*string `json:"WhiteList,omitempty" xml:"WhiteList,omitempty" type:"Repeated"`
}

func (s ConfigInstanceWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigInstanceWhiteListRequest) GoString() string {
	return s.String()
}

func (s *ConfigInstanceWhiteListRequest) SetInstanceId(v string) *ConfigInstanceWhiteListRequest {
	s.InstanceId = &v
	return s
}

func (s *ConfigInstanceWhiteListRequest) SetRegionId(v string) *ConfigInstanceWhiteListRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigInstanceWhiteListRequest) SetIpVersion(v string) *ConfigInstanceWhiteListRequest {
	s.IpVersion = &v
	return s
}

func (s *ConfigInstanceWhiteListRequest) SetWhiteList(v []*string) *ConfigInstanceWhiteListRequest {
	s.WhiteList = v
	return s
}

type ConfigInstanceWhiteListResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigInstanceWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigInstanceWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigInstanceWhiteListResponseBody) SetRequestId(v string) *ConfigInstanceWhiteListResponseBody {
	s.RequestId = &v
	return s
}

type ConfigInstanceWhiteListResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfigInstanceWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfigInstanceWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigInstanceWhiteListResponse) GoString() string {
	return s.String()
}

func (s *ConfigInstanceWhiteListResponse) SetHeaders(v map[string]*string) *ConfigInstanceWhiteListResponse {
	s.Headers = v
	return s
}

func (s *ConfigInstanceWhiteListResponse) SetBody(v *ConfigInstanceWhiteListResponseBody) *ConfigInstanceWhiteListResponse {
	s.Body = v
	return s
}

type DescribeAuditLogsRequest struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime     *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Sort        *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	Dir         *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	DbId        *string `json:"DbId,omitempty" xml:"DbId,omitempty"`
	QueryString *string `json:"QueryString,omitempty" xml:"QueryString,omitempty"`
	Payload     *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	LoginUser   *string `json:"LoginUser,omitempty" xml:"LoginUser,omitempty"`
	OpType      *string `json:"OpType,omitempty" xml:"OpType,omitempty"`
	Sip         *string `json:"Sip,omitempty" xml:"Sip,omitempty"`
	Dip         *string `json:"Dip,omitempty" xml:"Dip,omitempty"`
	Result      *string `json:"Result,omitempty" xml:"Result,omitempty"`
	Accessid    *string `json:"Accessid,omitempty" xml:"Accessid,omitempty"`
	Sessionid   *string `json:"Sessionid,omitempty" xml:"Sessionid,omitempty"`
	Sqlid       *string `json:"Sqlid,omitempty" xml:"Sqlid,omitempty"`
	DbType      *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	Sport       *string `json:"Sport,omitempty" xml:"Sport,omitempty"`
	Dport       *string `json:"Dport,omitempty" xml:"Dport,omitempty"`
	Smac        *string `json:"Smac,omitempty" xml:"Smac,omitempty"`
	Dmac        *string `json:"Dmac,omitempty" xml:"Dmac,omitempty"`
	DbName      *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	ClientPrg   *string `json:"ClientPrg,omitempty" xml:"ClientPrg,omitempty"`
	HostName    *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	ClientUser  *string `json:"ClientUser,omitempty" xml:"ClientUser,omitempty"`
	SqlLen      *string `json:"SqlLen,omitempty" xml:"SqlLen,omitempty"`
	EffectRow   *string `json:"EffectRow,omitempty" xml:"EffectRow,omitempty"`
	Cost        *string `json:"Cost,omitempty" xml:"Cost,omitempty"`
	ResultDesc  *string `json:"ResultDesc,omitempty" xml:"ResultDesc,omitempty"`
	DataSet     *string `json:"DataSet,omitempty" xml:"DataSet,omitempty"`
	AlarmName   *string `json:"AlarmName,omitempty" xml:"AlarmName,omitempty"`
	AlarmLevel  *string `json:"AlarmLevel,omitempty" xml:"AlarmLevel,omitempty"`
}

func (s DescribeAuditLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogsRequest) SetInstanceId(v string) *DescribeAuditLogsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetRegionId(v string) *DescribeAuditLogsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetStartTime(v string) *DescribeAuditLogsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetEndTime(v string) *DescribeAuditLogsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetCurrentPage(v int32) *DescribeAuditLogsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetPageSize(v int32) *DescribeAuditLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetSort(v string) *DescribeAuditLogsRequest {
	s.Sort = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetDir(v string) *DescribeAuditLogsRequest {
	s.Dir = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetDbId(v string) *DescribeAuditLogsRequest {
	s.DbId = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetQueryString(v string) *DescribeAuditLogsRequest {
	s.QueryString = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetPayload(v string) *DescribeAuditLogsRequest {
	s.Payload = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetLoginUser(v string) *DescribeAuditLogsRequest {
	s.LoginUser = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetOpType(v string) *DescribeAuditLogsRequest {
	s.OpType = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetSip(v string) *DescribeAuditLogsRequest {
	s.Sip = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetDip(v string) *DescribeAuditLogsRequest {
	s.Dip = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetResult(v string) *DescribeAuditLogsRequest {
	s.Result = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetAccessid(v string) *DescribeAuditLogsRequest {
	s.Accessid = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetSessionid(v string) *DescribeAuditLogsRequest {
	s.Sessionid = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetSqlid(v string) *DescribeAuditLogsRequest {
	s.Sqlid = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetDbType(v string) *DescribeAuditLogsRequest {
	s.DbType = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetSport(v string) *DescribeAuditLogsRequest {
	s.Sport = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetDport(v string) *DescribeAuditLogsRequest {
	s.Dport = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetSmac(v string) *DescribeAuditLogsRequest {
	s.Smac = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetDmac(v string) *DescribeAuditLogsRequest {
	s.Dmac = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetDbName(v string) *DescribeAuditLogsRequest {
	s.DbName = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetClientPrg(v string) *DescribeAuditLogsRequest {
	s.ClientPrg = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetHostName(v string) *DescribeAuditLogsRequest {
	s.HostName = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetClientUser(v string) *DescribeAuditLogsRequest {
	s.ClientUser = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetSqlLen(v string) *DescribeAuditLogsRequest {
	s.SqlLen = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetEffectRow(v string) *DescribeAuditLogsRequest {
	s.EffectRow = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetCost(v string) *DescribeAuditLogsRequest {
	s.Cost = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetResultDesc(v string) *DescribeAuditLogsRequest {
	s.ResultDesc = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetDataSet(v string) *DescribeAuditLogsRequest {
	s.DataSet = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetAlarmName(v string) *DescribeAuditLogsRequest {
	s.AlarmName = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetAlarmLevel(v string) *DescribeAuditLogsRequest {
	s.AlarmLevel = &v
	return s
}

type DescribeAuditLogsResponseBody struct {
	TotalCount *int32                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AuditLogs  []*DescribeAuditLogsResponseBodyAuditLogs `json:"AuditLogs,omitempty" xml:"AuditLogs,omitempty" type:"Repeated"`
}

func (s DescribeAuditLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogsResponseBody) SetTotalCount(v int32) *DescribeAuditLogsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAuditLogsResponseBody) SetRequestId(v string) *DescribeAuditLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAuditLogsResponseBody) SetAuditLogs(v []*DescribeAuditLogsResponseBodyAuditLogs) *DescribeAuditLogsResponseBody {
	s.AuditLogs = v
	return s
}

type DescribeAuditLogsResponseBodyAuditLogs struct {
	ClientUser  *string `json:"ClientUser,omitempty" xml:"ClientUser,omitempty"`
	EffectRow   *int32  `json:"EffectRow,omitempty" xml:"EffectRow,omitempty"`
	C5          *string `json:"C5,omitempty" xml:"C5,omitempty"`
	ClientPrg   *string `json:"ClientPrg,omitempty" xml:"ClientPrg,omitempty"`
	Accessid    *string `json:"Accessid,omitempty" xml:"Accessid,omitempty"`
	ResultDesc  *string `json:"ResultDesc,omitempty" xml:"ResultDesc,omitempty"`
	SqlLen      *int32  `json:"SqlLen,omitempty" xml:"SqlLen,omitempty"`
	Payload     *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	C4          *string `json:"C4,omitempty" xml:"C4,omitempty"`
	DateTime    *string `json:"DateTime,omitempty" xml:"DateTime,omitempty"`
	DbName      *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DataSet     *string `json:"DataSet,omitempty" xml:"DataSet,omitempty"`
	Sqlid       *string `json:"Sqlid,omitempty" xml:"Sqlid,omitempty"`
	RelateIp    *string `json:"RelateIp,omitempty" xml:"RelateIp,omitempty"`
	AlarmLevel  *int32  `json:"AlarmLevel,omitempty" xml:"AlarmLevel,omitempty"`
	C2          *string `json:"C2,omitempty" xml:"C2,omitempty"`
	Dip         *string `json:"Dip,omitempty" xml:"Dip,omitempty"`
	Result      *int32  `json:"Result,omitempty" xml:"Result,omitempty"`
	Cost        *int32  `json:"Cost,omitempty" xml:"Cost,omitempty"`
	RelateUser  *string `json:"RelateUser,omitempty" xml:"RelateUser,omitempty"`
	Sip         *string `json:"Sip,omitempty" xml:"Sip,omitempty"`
	C3          *string `json:"C3,omitempty" xml:"C3,omitempty"`
	HostName    *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	AlarmName   *string `json:"AlarmName,omitempty" xml:"AlarmName,omitempty"`
	PickIp      *string `json:"PickIp,omitempty" xml:"PickIp,omitempty"`
	RelateInfo  *string `json:"RelateInfo,omitempty" xml:"RelateInfo,omitempty"`
	PickUser    *string `json:"PickUser,omitempty" xml:"PickUser,omitempty"`
	OpType      *string `json:"OpType,omitempty" xml:"OpType,omitempty"`
	Sport       *int32  `json:"Sport,omitempty" xml:"Sport,omitempty"`
	DataSetSize *int32  `json:"DataSetSize,omitempty" xml:"DataSetSize,omitempty"`
	DbType      *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	AlarmFlag   *int32  `json:"AlarmFlag,omitempty" xml:"AlarmFlag,omitempty"`
	Smac        *int32  `json:"Smac,omitempty" xml:"Smac,omitempty"`
	Dport       *int32  `json:"Dport,omitempty" xml:"Dport,omitempty"`
	C1          *string `json:"C1,omitempty" xml:"C1,omitempty"`
	Dmac        *int32  `json:"Dmac,omitempty" xml:"Dmac,omitempty"`
	LoginUser   *string `json:"LoginUser,omitempty" xml:"LoginUser,omitempty"`
	Sessionid   *string `json:"Sessionid,omitempty" xml:"Sessionid,omitempty"`
}

func (s DescribeAuditLogsResponseBodyAuditLogs) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditLogsResponseBodyAuditLogs) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogsResponseBodyAuditLogs) SetClientUser(v string) *DescribeAuditLogsResponseBodyAuditLogs {
	s.ClientUser = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyAuditLogs) SetEffectRow(v int32) *DescribeAuditLogsResponseBodyAuditLogs {
	s.EffectRow = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyAuditLogs) SetC5(v string) *DescribeAuditLogsResponseBodyAuditLogs {
	s.C5 = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyAuditLogs) SetClientPrg(v string) *DescribeAuditLogsResponseBodyAuditLogs {
	s.ClientPrg = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyAuditLogs) SetAccessid(v string) *DescribeAuditLogsResponseBodyAuditLogs {
	s.Accessid = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyAuditLogs) SetResultDesc(v string) *DescribeAuditLogsResponseBodyAuditLogs {
	s.ResultDesc = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyAuditLogs) SetSqlLen(v int32) *DescribeAuditLogsResponseBodyAuditLogs {
	s.SqlLen = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyAuditLogs) SetPayload(v string) *DescribeAuditLogsResponseBodyAuditLogs {
	s.Payload = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyAuditLogs) SetC4(v string) *DescribeAuditLogsResponseBodyAuditLogs {
	s.C4 = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyAuditLogs) SetDateTime(v string) *DescribeAuditLogsResponseBodyAuditLogs {
	s.DateTime = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyAuditLogs) SetDbName(v string) *DescribeAuditLogsResponseBodyAuditLogs {
	s.DbName = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyAuditLogs) SetDataSet(v string) *DescribeAuditLogsResponseBodyAuditLogs {
	s.DataSet = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyAuditLogs) SetSqlid(v string) *DescribeAuditLogsResponseBodyAuditLogs {
	s.Sqlid = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyAuditLogs) SetRelateIp(v string) *DescribeAuditLogsResponseBodyAuditLogs {
	s.RelateIp = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyAuditLogs) SetAlarmLevel(v int32) *DescribeAuditLogsResponseBodyAuditLogs {
	s.AlarmLevel = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyAuditLogs) SetC2(v string) *DescribeAuditLogsResponseBodyAuditLogs {
	s.C2 = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyAuditLogs) SetDip(v string) *DescribeAuditLogsResponseBodyAuditLogs {
	s.Dip = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyAuditLogs) SetResult(v int32) *DescribeAuditLogsResponseBodyAuditLogs {
	s.Result = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyAuditLogs) SetCost(v int32) *DescribeAuditLogsResponseBodyAuditLogs {
	s.Cost = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyAuditLogs) SetRelateUser(v string) *DescribeAuditLogsResponseBodyAuditLogs {
	s.RelateUser = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyAuditLogs) SetSip(v string) *DescribeAuditLogsResponseBodyAuditLogs {
	s.Sip = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyAuditLogs) SetC3(v string) *DescribeAuditLogsResponseBodyAuditLogs {
	s.C3 = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyAuditLogs) SetHostName(v string) *DescribeAuditLogsResponseBodyAuditLogs {
	s.HostName = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyAuditLogs) SetAlarmName(v string) *DescribeAuditLogsResponseBodyAuditLogs {
	s.AlarmName = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyAuditLogs) SetPickIp(v string) *DescribeAuditLogsResponseBodyAuditLogs {
	s.PickIp = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyAuditLogs) SetRelateInfo(v string) *DescribeAuditLogsResponseBodyAuditLogs {
	s.RelateInfo = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyAuditLogs) SetPickUser(v string) *DescribeAuditLogsResponseBodyAuditLogs {
	s.PickUser = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyAuditLogs) SetOpType(v string) *DescribeAuditLogsResponseBodyAuditLogs {
	s.OpType = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyAuditLogs) SetSport(v int32) *DescribeAuditLogsResponseBodyAuditLogs {
	s.Sport = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyAuditLogs) SetDataSetSize(v int32) *DescribeAuditLogsResponseBodyAuditLogs {
	s.DataSetSize = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyAuditLogs) SetDbType(v string) *DescribeAuditLogsResponseBodyAuditLogs {
	s.DbType = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyAuditLogs) SetAlarmFlag(v int32) *DescribeAuditLogsResponseBodyAuditLogs {
	s.AlarmFlag = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyAuditLogs) SetSmac(v int32) *DescribeAuditLogsResponseBodyAuditLogs {
	s.Smac = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyAuditLogs) SetDport(v int32) *DescribeAuditLogsResponseBodyAuditLogs {
	s.Dport = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyAuditLogs) SetC1(v string) *DescribeAuditLogsResponseBodyAuditLogs {
	s.C1 = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyAuditLogs) SetDmac(v int32) *DescribeAuditLogsResponseBodyAuditLogs {
	s.Dmac = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyAuditLogs) SetLoginUser(v string) *DescribeAuditLogsResponseBodyAuditLogs {
	s.LoginUser = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyAuditLogs) SetSessionid(v string) *DescribeAuditLogsResponseBodyAuditLogs {
	s.Sessionid = &v
	return s
}

type DescribeAuditLogsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAuditLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAuditLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogsResponse) SetHeaders(v map[string]*string) *DescribeAuditLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAuditLogsResponse) SetBody(v *DescribeAuditLogsResponseBody) *DescribeAuditLogsResponse {
	s.Body = v
	return s
}

type DescribeInstanceAttribueRequest struct {
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeInstanceAttribueRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAttribueRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttribueRequest) SetLang(v string) *DescribeInstanceAttribueRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInstanceAttribueRequest) SetInstanceId(v string) *DescribeInstanceAttribueRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceAttribueRequest) SetRegionId(v string) *DescribeInstanceAttribueRequest {
	s.RegionId = &v
	return s
}

type DescribeInstanceAttribueResponseBody struct {
	RequestId        *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceAttribue *DescribeInstanceAttribueResponseBodyInstanceAttribue `json:"InstanceAttribue,omitempty" xml:"InstanceAttribue,omitempty" type:"Struct"`
}

func (s DescribeInstanceAttribueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAttribueResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttribueResponseBody) SetRequestId(v string) *DescribeInstanceAttribueResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceAttribueResponseBody) SetInstanceAttribue(v *DescribeInstanceAttribueResponseBodyInstanceAttribue) *DescribeInstanceAttribueResponseBody {
	s.InstanceAttribue = v
	return s
}

type DescribeInstanceAttribueResponseBodyInstanceAttribue struct {
	VpcId               *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchId           *string   `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	ExpireTime          *int64    `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	InstanceId          *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InternetEndpoint    *string   `json:"InternetEndpoint,omitempty" xml:"InternetEndpoint,omitempty"`
	RegionId            *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	IntranetEndpoint    *string   `json:"IntranetEndpoint,omitempty" xml:"IntranetEndpoint,omitempty"`
	StartTime           *int64    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	SeriesCode          *string   `json:"SeriesCode,omitempty" xml:"SeriesCode,omitempty"`
	Description         *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceStatus      *int32    `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	LicenseCode         *string   `json:"LicenseCode,omitempty" xml:"LicenseCode,omitempty"`
	PublicNetworkAccess *bool     `json:"PublicNetworkAccess,omitempty" xml:"PublicNetworkAccess,omitempty"`
	WhiteList           []*string `json:"WhiteList,omitempty" xml:"WhiteList,omitempty" type:"Repeated"`
}

func (s DescribeInstanceAttribueResponseBodyInstanceAttribue) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAttribueResponseBodyInstanceAttribue) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttribueResponseBodyInstanceAttribue) SetVpcId(v string) *DescribeInstanceAttribueResponseBodyInstanceAttribue {
	s.VpcId = &v
	return s
}

func (s *DescribeInstanceAttribueResponseBodyInstanceAttribue) SetVswitchId(v string) *DescribeInstanceAttribueResponseBodyInstanceAttribue {
	s.VswitchId = &v
	return s
}

func (s *DescribeInstanceAttribueResponseBodyInstanceAttribue) SetExpireTime(v int64) *DescribeInstanceAttribueResponseBodyInstanceAttribue {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstanceAttribueResponseBodyInstanceAttribue) SetInstanceId(v string) *DescribeInstanceAttribueResponseBodyInstanceAttribue {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceAttribueResponseBodyInstanceAttribue) SetInternetEndpoint(v string) *DescribeInstanceAttribueResponseBodyInstanceAttribue {
	s.InternetEndpoint = &v
	return s
}

func (s *DescribeInstanceAttribueResponseBodyInstanceAttribue) SetRegionId(v string) *DescribeInstanceAttribueResponseBodyInstanceAttribue {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceAttribueResponseBodyInstanceAttribue) SetIntranetEndpoint(v string) *DescribeInstanceAttribueResponseBodyInstanceAttribue {
	s.IntranetEndpoint = &v
	return s
}

func (s *DescribeInstanceAttribueResponseBodyInstanceAttribue) SetStartTime(v int64) *DescribeInstanceAttribueResponseBodyInstanceAttribue {
	s.StartTime = &v
	return s
}

func (s *DescribeInstanceAttribueResponseBodyInstanceAttribue) SetSeriesCode(v string) *DescribeInstanceAttribueResponseBodyInstanceAttribue {
	s.SeriesCode = &v
	return s
}

func (s *DescribeInstanceAttribueResponseBodyInstanceAttribue) SetDescription(v string) *DescribeInstanceAttribueResponseBodyInstanceAttribue {
	s.Description = &v
	return s
}

func (s *DescribeInstanceAttribueResponseBodyInstanceAttribue) SetInstanceStatus(v int32) *DescribeInstanceAttribueResponseBodyInstanceAttribue {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeInstanceAttribueResponseBodyInstanceAttribue) SetLicenseCode(v string) *DescribeInstanceAttribueResponseBodyInstanceAttribue {
	s.LicenseCode = &v
	return s
}

func (s *DescribeInstanceAttribueResponseBodyInstanceAttribue) SetPublicNetworkAccess(v bool) *DescribeInstanceAttribueResponseBodyInstanceAttribue {
	s.PublicNetworkAccess = &v
	return s
}

func (s *DescribeInstanceAttribueResponseBodyInstanceAttribue) SetWhiteList(v []*string) *DescribeInstanceAttribueResponseBodyInstanceAttribue {
	s.WhiteList = v
	return s
}

type DescribeInstanceAttribueResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceAttribueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceAttribueResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAttribueResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttribueResponse) SetHeaders(v map[string]*string) *DescribeInstanceAttribueResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceAttribueResponse) SetBody(v *DescribeInstanceAttribueResponseBody) *DescribeInstanceAttribueResponse {
	s.Body = v
	return s
}

type DescribeInstanceAttributeRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeInstanceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeRequest) SetInstanceId(v string) *DescribeInstanceAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceAttributeRequest) SetRegionId(v string) *DescribeInstanceAttributeRequest {
	s.RegionId = &v
	return s
}

type DescribeInstanceAttributeResponseBody struct {
	RequestId         *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceAttribute *DescribeInstanceAttributeResponseBodyInstanceAttribute `json:"InstanceAttribute,omitempty" xml:"InstanceAttribute,omitempty" type:"Struct"`
}

func (s DescribeInstanceAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponseBody) SetRequestId(v string) *DescribeInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetInstanceAttribute(v *DescribeInstanceAttributeResponseBodyInstanceAttribute) *DescribeInstanceAttributeResponseBody {
	s.InstanceAttribute = v
	return s
}

type DescribeInstanceAttributeResponseBodyInstanceAttribute struct {
	VpcId               *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchId           *string   `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	ExpireTime          *int64    `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	ImageVersion        *string   `json:"ImageVersion,omitempty" xml:"ImageVersion,omitempty"`
	InstanceId          *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InternetEndpoint    *string   `json:"InternetEndpoint,omitempty" xml:"InternetEndpoint,omitempty"`
	RegionId            *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	IntranetEndpoint    *string   `json:"IntranetEndpoint,omitempty" xml:"IntranetEndpoint,omitempty"`
	StartTime           *int64    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	SeriesCode          *string   `json:"SeriesCode,omitempty" xml:"SeriesCode,omitempty"`
	Description         *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceStatus      *string   `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	LicenseCode         *string   `json:"LicenseCode,omitempty" xml:"LicenseCode,omitempty"`
	PublicNetworkAccess *bool     `json:"PublicNetworkAccess,omitempty" xml:"PublicNetworkAccess,omitempty"`
	WhiteList           []*string `json:"WhiteList,omitempty" xml:"WhiteList,omitempty" type:"Repeated"`
	Ipv6WhiteList       []*string `json:"Ipv6WhiteList,omitempty" xml:"Ipv6WhiteList,omitempty" type:"Repeated"`
}

func (s DescribeInstanceAttributeResponseBodyInstanceAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAttributeResponseBodyInstanceAttribute) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetVpcId(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.VpcId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetVswitchId(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.VswitchId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetExpireTime(v int64) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetImageVersion(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.ImageVersion = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetInstanceId(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetInternetEndpoint(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.InternetEndpoint = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetRegionId(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetIntranetEndpoint(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.IntranetEndpoint = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetStartTime(v int64) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.StartTime = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetSeriesCode(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.SeriesCode = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetDescription(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.Description = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetInstanceStatus(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetLicenseCode(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.LicenseCode = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetPublicNetworkAccess(v bool) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.PublicNetworkAccess = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetWhiteList(v []*string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.WhiteList = v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetIpv6WhiteList(v []*string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.Ipv6WhiteList = v
	return s
}

type DescribeInstanceAttributeResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponse) SetHeaders(v map[string]*string) *DescribeInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceAttributeResponse) SetBody(v *DescribeInstanceAttributeResponseBody) *DescribeInstanceAttributeResponse {
	s.Body = v
	return s
}

type DescribeInstancesRequest struct {
	PageSize        *int32                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage     *int32                         `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	RegionId        *string                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceStatus  *string                        `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	ResourceGroupId *string                        `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	InstanceId      []*string                      `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	Tag             []*DescribeInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequest) SetPageSize(v int32) *DescribeInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesRequest) SetCurrentPage(v int32) *DescribeInstancesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInstancesRequest) SetRegionId(v string) *DescribeInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceStatus(v string) *DescribeInstancesRequest {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeInstancesRequest) SetResourceGroupId(v string) *DescribeInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceId(v []*string) *DescribeInstancesRequest {
	s.InstanceId = v
	return s
}

func (s *DescribeInstancesRequest) SetTag(v []*DescribeInstancesRequestTag) *DescribeInstancesRequest {
	s.Tag = v
	return s
}

type DescribeInstancesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstancesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequestTag) SetKey(v string) *DescribeInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeInstancesRequestTag) SetValue(v string) *DescribeInstancesRequestTag {
	s.Value = &v
	return s
}

type DescribeInstancesResponseBody struct {
	TotalCount *int64                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Instances  []*DescribeInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBody) SetTotalCount(v int64) *DescribeInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetRequestId(v string) *DescribeInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetInstances(v []*DescribeInstancesResponseBodyInstances) *DescribeInstancesResponseBody {
	s.Instances = v
	return s
}

type DescribeInstancesResponseBodyInstances struct {
	VpcId               *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchId           *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	ExpireTime          *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	ImageVersion        *string `json:"ImageVersion,omitempty" xml:"ImageVersion,omitempty"`
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InternetEndpoint    *string `json:"InternetEndpoint,omitempty" xml:"InternetEndpoint,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	IntranetEndpoint    *string `json:"IntranetEndpoint,omitempty" xml:"IntranetEndpoint,omitempty"`
	StartTime           *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	SeriesCode          *string `json:"SeriesCode,omitempty" xml:"SeriesCode,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceStatus      *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	LicenseCode         *string `json:"LicenseCode,omitempty" xml:"LicenseCode,omitempty"`
	PublicNetworkAccess *bool   `json:"PublicNetworkAccess,omitempty" xml:"PublicNetworkAccess,omitempty"`
}

func (s DescribeInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstances) SetVpcId(v string) *DescribeInstancesResponseBodyInstances {
	s.VpcId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetVswitchId(v string) *DescribeInstancesResponseBodyInstances {
	s.VswitchId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetExpireTime(v int64) *DescribeInstancesResponseBodyInstances {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetImageVersion(v string) *DescribeInstancesResponseBodyInstances {
	s.ImageVersion = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInstanceId(v string) *DescribeInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInternetEndpoint(v string) *DescribeInstancesResponseBodyInstances {
	s.InternetEndpoint = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetRegionId(v string) *DescribeInstancesResponseBodyInstances {
	s.RegionId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetIntranetEndpoint(v string) *DescribeInstancesResponseBodyInstances {
	s.IntranetEndpoint = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetStartTime(v int64) *DescribeInstancesResponseBodyInstances {
	s.StartTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetSeriesCode(v string) *DescribeInstancesResponseBodyInstances {
	s.SeriesCode = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetDescription(v string) *DescribeInstancesResponseBodyInstances {
	s.Description = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInstanceStatus(v string) *DescribeInstancesResponseBodyInstances {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetLicenseCode(v string) *DescribeInstancesResponseBodyInstances {
	s.LicenseCode = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetPublicNetworkAccess(v bool) *DescribeInstancesResponseBodyInstances {
	s.PublicNetworkAccess = &v
	return s
}

type DescribeInstancesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponse) SetHeaders(v map[string]*string) *DescribeInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstancesResponse) SetBody(v *DescribeInstancesResponseBody) *DescribeInstancesResponse {
	s.Body = v
	return s
}

type DescribeInstanceStorageRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeInstanceStorageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceStorageRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStorageRequest) SetInstanceId(v string) *DescribeInstanceStorageRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceStorageRequest) SetRegionId(v string) *DescribeInstanceStorageRequest {
	s.RegionId = &v
	return s
}

type DescribeInstanceStorageResponseBody struct {
	RequestId        *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceStorages []*DescribeInstanceStorageResponseBodyInstanceStorages `json:"InstanceStorages,omitempty" xml:"InstanceStorages,omitempty" type:"Repeated"`
}

func (s DescribeInstanceStorageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceStorageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStorageResponseBody) SetRequestId(v string) *DescribeInstanceStorageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceStorageResponseBody) SetInstanceStorages(v []*DescribeInstanceStorageResponseBodyInstanceStorages) *DescribeInstanceStorageResponseBody {
	s.InstanceStorages = v
	return s
}

type DescribeInstanceStorageResponseBodyInstanceStorages struct {
	StorageTime     *int64  `json:"StorageTime,omitempty" xml:"StorageTime,omitempty"`
	StorageCapacity *int64  `json:"StorageCapacity,omitempty" xml:"StorageCapacity,omitempty"`
	StorageCategory *string `json:"StorageCategory,omitempty" xml:"StorageCategory,omitempty"`
	StorageSpace    *string `json:"StorageSpace,omitempty" xml:"StorageSpace,omitempty"`
	StorageUsed     *int64  `json:"StorageUsed,omitempty" xml:"StorageUsed,omitempty"`
}

func (s DescribeInstanceStorageResponseBodyInstanceStorages) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceStorageResponseBodyInstanceStorages) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStorageResponseBodyInstanceStorages) SetStorageTime(v int64) *DescribeInstanceStorageResponseBodyInstanceStorages {
	s.StorageTime = &v
	return s
}

func (s *DescribeInstanceStorageResponseBodyInstanceStorages) SetStorageCapacity(v int64) *DescribeInstanceStorageResponseBodyInstanceStorages {
	s.StorageCapacity = &v
	return s
}

func (s *DescribeInstanceStorageResponseBodyInstanceStorages) SetStorageCategory(v string) *DescribeInstanceStorageResponseBodyInstanceStorages {
	s.StorageCategory = &v
	return s
}

func (s *DescribeInstanceStorageResponseBodyInstanceStorages) SetStorageSpace(v string) *DescribeInstanceStorageResponseBodyInstanceStorages {
	s.StorageSpace = &v
	return s
}

func (s *DescribeInstanceStorageResponseBodyInstanceStorages) SetStorageUsed(v int64) *DescribeInstanceStorageResponseBodyInstanceStorages {
	s.StorageUsed = &v
	return s
}

type DescribeInstanceStorageResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceStorageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceStorageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceStorageResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStorageResponse) SetHeaders(v map[string]*string) *DescribeInstanceStorageResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceStorageResponse) SetBody(v *DescribeInstanceStorageResponseBody) *DescribeInstanceStorageResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeRegionsRequest) SetRegionId(v string) *DescribeRegionsRequest {
	s.RegionId = &v
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
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	LocalName      *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetLocalName(v string) *DescribeRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionId(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionId = &v
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

type DescribeRenewStatusRequest struct {
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
}

func (s DescribeRenewStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRenewStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeRenewStatusRequest) SetRegionId(v string) *DescribeRenewStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRenewStatusRequest) SetInstanceId(v []*string) *DescribeRenewStatusRequest {
	s.InstanceId = v
	return s
}

type DescribeRenewStatusResponseBody struct {
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Instances []*DescribeRenewStatusResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
}

func (s DescribeRenewStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRenewStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRenewStatusResponseBody) SetRequestId(v string) *DescribeRenewStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRenewStatusResponseBody) SetInstances(v []*DescribeRenewStatusResponseBodyInstances) *DescribeRenewStatusResponseBody {
	s.Instances = v
	return s
}

type DescribeRenewStatusResponseBodyInstances struct {
	RenewStatus *string `json:"RenewStatus,omitempty" xml:"RenewStatus,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeRenewStatusResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeRenewStatusResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeRenewStatusResponseBodyInstances) SetRenewStatus(v string) *DescribeRenewStatusResponseBodyInstances {
	s.RenewStatus = &v
	return s
}

func (s *DescribeRenewStatusResponseBodyInstances) SetInstanceId(v string) *DescribeRenewStatusResponseBodyInstances {
	s.InstanceId = &v
	return s
}

type DescribeRenewStatusResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRenewStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRenewStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRenewStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeRenewStatusResponse) SetHeaders(v map[string]*string) *DescribeRenewStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeRenewStatusResponse) SetBody(v *DescribeRenewStatusResponseBody) *DescribeRenewStatusResponse {
	s.Body = v
	return s
}

type DescribeSessionLogsRequest struct {
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	CurrentPage   *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Sort          *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	Dir           *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	DbId          *string `json:"DbId,omitempty" xml:"DbId,omitempty"`
	Sip           *string `json:"Sip,omitempty" xml:"Sip,omitempty"`
	Sport         *string `json:"Sport,omitempty" xml:"Sport,omitempty"`
	LoginUser     *string `json:"LoginUser,omitempty" xml:"LoginUser,omitempty"`
	Dip           *string `json:"Dip,omitempty" xml:"Dip,omitempty"`
	Dport         *string `json:"Dport,omitempty" xml:"Dport,omitempty"`
	Sessionid     *string `json:"Sessionid,omitempty" xml:"Sessionid,omitempty"`
	DbType        *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	Smac          *string `json:"Smac,omitempty" xml:"Smac,omitempty"`
	Dmac          *string `json:"Dmac,omitempty" xml:"Dmac,omitempty"`
	ClientPrg     *string `json:"ClientPrg,omitempty" xml:"ClientPrg,omitempty"`
	HostName      *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	ClientUser    *string `json:"ClientUser,omitempty" xml:"ClientUser,omitempty"`
	DbName        *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	SessionStatus *string `json:"SessionStatus,omitempty" xml:"SessionStatus,omitempty"`
	SqlCount      *string `json:"SqlCount,omitempty" xml:"SqlCount,omitempty"`
	ReqFlow       *string `json:"ReqFlow,omitempty" xml:"ReqFlow,omitempty"`
	RspFlow       *string `json:"RspFlow,omitempty" xml:"RspFlow,omitempty"`
}

func (s DescribeSessionLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSessionLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSessionLogsRequest) SetInstanceId(v string) *DescribeSessionLogsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSessionLogsRequest) SetRegionId(v string) *DescribeSessionLogsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSessionLogsRequest) SetStartTime(v string) *DescribeSessionLogsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSessionLogsRequest) SetEndTime(v string) *DescribeSessionLogsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSessionLogsRequest) SetCurrentPage(v int32) *DescribeSessionLogsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSessionLogsRequest) SetPageSize(v int32) *DescribeSessionLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSessionLogsRequest) SetSort(v string) *DescribeSessionLogsRequest {
	s.Sort = &v
	return s
}

func (s *DescribeSessionLogsRequest) SetDir(v string) *DescribeSessionLogsRequest {
	s.Dir = &v
	return s
}

func (s *DescribeSessionLogsRequest) SetDbId(v string) *DescribeSessionLogsRequest {
	s.DbId = &v
	return s
}

func (s *DescribeSessionLogsRequest) SetSip(v string) *DescribeSessionLogsRequest {
	s.Sip = &v
	return s
}

func (s *DescribeSessionLogsRequest) SetSport(v string) *DescribeSessionLogsRequest {
	s.Sport = &v
	return s
}

func (s *DescribeSessionLogsRequest) SetLoginUser(v string) *DescribeSessionLogsRequest {
	s.LoginUser = &v
	return s
}

func (s *DescribeSessionLogsRequest) SetDip(v string) *DescribeSessionLogsRequest {
	s.Dip = &v
	return s
}

func (s *DescribeSessionLogsRequest) SetDport(v string) *DescribeSessionLogsRequest {
	s.Dport = &v
	return s
}

func (s *DescribeSessionLogsRequest) SetSessionid(v string) *DescribeSessionLogsRequest {
	s.Sessionid = &v
	return s
}

func (s *DescribeSessionLogsRequest) SetDbType(v string) *DescribeSessionLogsRequest {
	s.DbType = &v
	return s
}

func (s *DescribeSessionLogsRequest) SetSmac(v string) *DescribeSessionLogsRequest {
	s.Smac = &v
	return s
}

func (s *DescribeSessionLogsRequest) SetDmac(v string) *DescribeSessionLogsRequest {
	s.Dmac = &v
	return s
}

func (s *DescribeSessionLogsRequest) SetClientPrg(v string) *DescribeSessionLogsRequest {
	s.ClientPrg = &v
	return s
}

func (s *DescribeSessionLogsRequest) SetHostName(v string) *DescribeSessionLogsRequest {
	s.HostName = &v
	return s
}

func (s *DescribeSessionLogsRequest) SetClientUser(v string) *DescribeSessionLogsRequest {
	s.ClientUser = &v
	return s
}

func (s *DescribeSessionLogsRequest) SetDbName(v string) *DescribeSessionLogsRequest {
	s.DbName = &v
	return s
}

func (s *DescribeSessionLogsRequest) SetSessionStatus(v string) *DescribeSessionLogsRequest {
	s.SessionStatus = &v
	return s
}

func (s *DescribeSessionLogsRequest) SetSqlCount(v string) *DescribeSessionLogsRequest {
	s.SqlCount = &v
	return s
}

func (s *DescribeSessionLogsRequest) SetReqFlow(v string) *DescribeSessionLogsRequest {
	s.ReqFlow = &v
	return s
}

func (s *DescribeSessionLogsRequest) SetRspFlow(v string) *DescribeSessionLogsRequest {
	s.RspFlow = &v
	return s
}

type DescribeSessionLogsResponseBody struct {
	TotalCount  *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SessionLogs []*DescribeSessionLogsResponseBodySessionLogs `json:"SessionLogs,omitempty" xml:"SessionLogs,omitempty" type:"Repeated"`
}

func (s DescribeSessionLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSessionLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSessionLogsResponseBody) SetTotalCount(v int32) *DescribeSessionLogsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSessionLogsResponseBody) SetRequestId(v string) *DescribeSessionLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSessionLogsResponseBody) SetSessionLogs(v []*DescribeSessionLogsResponseBodySessionLogs) *DescribeSessionLogsResponseBody {
	s.SessionLogs = v
	return s
}

type DescribeSessionLogsResponseBodySessionLogs struct {
	ClientUser    *string `json:"ClientUser,omitempty" xml:"ClientUser,omitempty"`
	SessionStatus *int32  `json:"SessionStatus,omitempty" xml:"SessionStatus,omitempty"`
	C5            *string `json:"C5,omitempty" xml:"C5,omitempty"`
	ClientPrg     *string `json:"ClientPrg,omitempty" xml:"ClientPrg,omitempty"`
	Accessid      *string `json:"Accessid,omitempty" xml:"Accessid,omitempty"`
	C4            *string `json:"C4,omitempty" xml:"C4,omitempty"`
	DbName        *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	RequestFlow   *int32  `json:"RequestFlow,omitempty" xml:"RequestFlow,omitempty"`
	ProCon        *int32  `json:"ProCon,omitempty" xml:"ProCon,omitempty"`
	C2            *string `json:"C2,omitempty" xml:"C2,omitempty"`
	Dip           *string `json:"Dip,omitempty" xml:"Dip,omitempty"`
	Sip           *string `json:"Sip,omitempty" xml:"Sip,omitempty"`
	C3            *string `json:"C3,omitempty" xml:"C3,omitempty"`
	HostName      *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	ResponseFlow  *int32  `json:"ResponseFlow,omitempty" xml:"ResponseFlow,omitempty"`
	Encode        *string `json:"Encode,omitempty" xml:"Encode,omitempty"`
	NormalEnd     *int32  `json:"NormalEnd,omitempty" xml:"NormalEnd,omitempty"`
	EndTime       *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Sport         *int32  `json:"Sport,omitempty" xml:"Sport,omitempty"`
	StartTime     *int32  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	DbType        *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	StrInfo       *string `json:"StrInfo,omitempty" xml:"StrInfo,omitempty"`
	SqlCount      *int32  `json:"SqlCount,omitempty" xml:"SqlCount,omitempty"`
	Smac          *int32  `json:"Smac,omitempty" xml:"Smac,omitempty"`
	Dport         *int32  `json:"Dport,omitempty" xml:"Dport,omitempty"`
	Dmac          *int32  `json:"Dmac,omitempty" xml:"Dmac,omitempty"`
	C1            *string `json:"C1,omitempty" xml:"C1,omitempty"`
	LoginUser     *string `json:"LoginUser,omitempty" xml:"LoginUser,omitempty"`
	Sessionid     *string `json:"Sessionid,omitempty" xml:"Sessionid,omitempty"`
}

func (s DescribeSessionLogsResponseBodySessionLogs) String() string {
	return tea.Prettify(s)
}

func (s DescribeSessionLogsResponseBodySessionLogs) GoString() string {
	return s.String()
}

func (s *DescribeSessionLogsResponseBodySessionLogs) SetClientUser(v string) *DescribeSessionLogsResponseBodySessionLogs {
	s.ClientUser = &v
	return s
}

func (s *DescribeSessionLogsResponseBodySessionLogs) SetSessionStatus(v int32) *DescribeSessionLogsResponseBodySessionLogs {
	s.SessionStatus = &v
	return s
}

func (s *DescribeSessionLogsResponseBodySessionLogs) SetC5(v string) *DescribeSessionLogsResponseBodySessionLogs {
	s.C5 = &v
	return s
}

func (s *DescribeSessionLogsResponseBodySessionLogs) SetClientPrg(v string) *DescribeSessionLogsResponseBodySessionLogs {
	s.ClientPrg = &v
	return s
}

func (s *DescribeSessionLogsResponseBodySessionLogs) SetAccessid(v string) *DescribeSessionLogsResponseBodySessionLogs {
	s.Accessid = &v
	return s
}

func (s *DescribeSessionLogsResponseBodySessionLogs) SetC4(v string) *DescribeSessionLogsResponseBodySessionLogs {
	s.C4 = &v
	return s
}

func (s *DescribeSessionLogsResponseBodySessionLogs) SetDbName(v string) *DescribeSessionLogsResponseBodySessionLogs {
	s.DbName = &v
	return s
}

func (s *DescribeSessionLogsResponseBodySessionLogs) SetRequestFlow(v int32) *DescribeSessionLogsResponseBodySessionLogs {
	s.RequestFlow = &v
	return s
}

func (s *DescribeSessionLogsResponseBodySessionLogs) SetProCon(v int32) *DescribeSessionLogsResponseBodySessionLogs {
	s.ProCon = &v
	return s
}

func (s *DescribeSessionLogsResponseBodySessionLogs) SetC2(v string) *DescribeSessionLogsResponseBodySessionLogs {
	s.C2 = &v
	return s
}

func (s *DescribeSessionLogsResponseBodySessionLogs) SetDip(v string) *DescribeSessionLogsResponseBodySessionLogs {
	s.Dip = &v
	return s
}

func (s *DescribeSessionLogsResponseBodySessionLogs) SetSip(v string) *DescribeSessionLogsResponseBodySessionLogs {
	s.Sip = &v
	return s
}

func (s *DescribeSessionLogsResponseBodySessionLogs) SetC3(v string) *DescribeSessionLogsResponseBodySessionLogs {
	s.C3 = &v
	return s
}

func (s *DescribeSessionLogsResponseBodySessionLogs) SetHostName(v string) *DescribeSessionLogsResponseBodySessionLogs {
	s.HostName = &v
	return s
}

func (s *DescribeSessionLogsResponseBodySessionLogs) SetResponseFlow(v int32) *DescribeSessionLogsResponseBodySessionLogs {
	s.ResponseFlow = &v
	return s
}

func (s *DescribeSessionLogsResponseBodySessionLogs) SetEncode(v string) *DescribeSessionLogsResponseBodySessionLogs {
	s.Encode = &v
	return s
}

func (s *DescribeSessionLogsResponseBodySessionLogs) SetNormalEnd(v int32) *DescribeSessionLogsResponseBodySessionLogs {
	s.NormalEnd = &v
	return s
}

func (s *DescribeSessionLogsResponseBodySessionLogs) SetEndTime(v int32) *DescribeSessionLogsResponseBodySessionLogs {
	s.EndTime = &v
	return s
}

func (s *DescribeSessionLogsResponseBodySessionLogs) SetSport(v int32) *DescribeSessionLogsResponseBodySessionLogs {
	s.Sport = &v
	return s
}

func (s *DescribeSessionLogsResponseBodySessionLogs) SetStartTime(v int32) *DescribeSessionLogsResponseBodySessionLogs {
	s.StartTime = &v
	return s
}

func (s *DescribeSessionLogsResponseBodySessionLogs) SetDbType(v string) *DescribeSessionLogsResponseBodySessionLogs {
	s.DbType = &v
	return s
}

func (s *DescribeSessionLogsResponseBodySessionLogs) SetStrInfo(v string) *DescribeSessionLogsResponseBodySessionLogs {
	s.StrInfo = &v
	return s
}

func (s *DescribeSessionLogsResponseBodySessionLogs) SetSqlCount(v int32) *DescribeSessionLogsResponseBodySessionLogs {
	s.SqlCount = &v
	return s
}

func (s *DescribeSessionLogsResponseBodySessionLogs) SetSmac(v int32) *DescribeSessionLogsResponseBodySessionLogs {
	s.Smac = &v
	return s
}

func (s *DescribeSessionLogsResponseBodySessionLogs) SetDport(v int32) *DescribeSessionLogsResponseBodySessionLogs {
	s.Dport = &v
	return s
}

func (s *DescribeSessionLogsResponseBodySessionLogs) SetDmac(v int32) *DescribeSessionLogsResponseBodySessionLogs {
	s.Dmac = &v
	return s
}

func (s *DescribeSessionLogsResponseBodySessionLogs) SetC1(v string) *DescribeSessionLogsResponseBodySessionLogs {
	s.C1 = &v
	return s
}

func (s *DescribeSessionLogsResponseBodySessionLogs) SetLoginUser(v string) *DescribeSessionLogsResponseBodySessionLogs {
	s.LoginUser = &v
	return s
}

func (s *DescribeSessionLogsResponseBodySessionLogs) SetSessionid(v string) *DescribeSessionLogsResponseBodySessionLogs {
	s.Sessionid = &v
	return s
}

type DescribeSessionLogsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSessionLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSessionLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSessionLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSessionLogsResponse) SetHeaders(v map[string]*string) *DescribeSessionLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSessionLogsResponse) SetBody(v *DescribeSessionLogsResponseBody) *DescribeSessionLogsResponse {
	s.Body = v
	return s
}

type DisableInstancePublicAccessRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DisableInstancePublicAccessRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableInstancePublicAccessRequest) GoString() string {
	return s.String()
}

func (s *DisableInstancePublicAccessRequest) SetInstanceId(v string) *DisableInstancePublicAccessRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableInstancePublicAccessRequest) SetRegionId(v string) *DisableInstancePublicAccessRequest {
	s.RegionId = &v
	return s
}

type DisableInstancePublicAccessResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableInstancePublicAccessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableInstancePublicAccessResponseBody) GoString() string {
	return s.String()
}

func (s *DisableInstancePublicAccessResponseBody) SetRequestId(v string) *DisableInstancePublicAccessResponseBody {
	s.RequestId = &v
	return s
}

type DisableInstancePublicAccessResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableInstancePublicAccessResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableInstancePublicAccessResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableInstancePublicAccessResponse) GoString() string {
	return s.String()
}

func (s *DisableInstancePublicAccessResponse) SetHeaders(v map[string]*string) *DisableInstancePublicAccessResponse {
	s.Headers = v
	return s
}

func (s *DisableInstancePublicAccessResponse) SetBody(v *DisableInstancePublicAccessResponseBody) *DisableInstancePublicAccessResponse {
	s.Body = v
	return s
}

type EnableInstancePublicAccessRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s EnableInstancePublicAccessRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableInstancePublicAccessRequest) GoString() string {
	return s.String()
}

func (s *EnableInstancePublicAccessRequest) SetInstanceId(v string) *EnableInstancePublicAccessRequest {
	s.InstanceId = &v
	return s
}

func (s *EnableInstancePublicAccessRequest) SetRegionId(v string) *EnableInstancePublicAccessRequest {
	s.RegionId = &v
	return s
}

type EnableInstancePublicAccessResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableInstancePublicAccessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableInstancePublicAccessResponseBody) GoString() string {
	return s.String()
}

func (s *EnableInstancePublicAccessResponseBody) SetRequestId(v string) *EnableInstancePublicAccessResponseBody {
	s.RequestId = &v
	return s
}

type EnableInstancePublicAccessResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableInstancePublicAccessResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableInstancePublicAccessResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableInstancePublicAccessResponse) GoString() string {
	return s.String()
}

func (s *EnableInstancePublicAccessResponse) SetHeaders(v map[string]*string) *EnableInstancePublicAccessResponse {
	s.Headers = v
	return s
}

func (s *EnableInstancePublicAccessResponse) SetBody(v *EnableInstancePublicAccessResponseBody) *EnableInstancePublicAccessResponse {
	s.Body = v
	return s
}

type GenerateUploadAuthRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GenerateUploadAuthRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateUploadAuthRequest) GoString() string {
	return s.String()
}

func (s *GenerateUploadAuthRequest) SetRegionId(v string) *GenerateUploadAuthRequest {
	s.RegionId = &v
	return s
}

func (s *GenerateUploadAuthRequest) SetInstanceId(v string) *GenerateUploadAuthRequest {
	s.InstanceId = &v
	return s
}

type GenerateUploadAuthResponseBody struct {
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UploadConfig *GenerateUploadAuthResponseBodyUploadConfig `json:"UploadConfig,omitempty" xml:"UploadConfig,omitempty" type:"Struct"`
}

func (s GenerateUploadAuthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateUploadAuthResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateUploadAuthResponseBody) SetRequestId(v string) *GenerateUploadAuthResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateUploadAuthResponseBody) SetUploadConfig(v *GenerateUploadAuthResponseBodyUploadConfig) *GenerateUploadAuthResponseBody {
	s.UploadConfig = v
	return s
}

type GenerateUploadAuthResponseBodyUploadConfig struct {
	Signature  *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	FilePath   *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	Policy     *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	ExpireTime *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	UploadHost *string `json:"UploadHost,omitempty" xml:"UploadHost,omitempty"`
	AccessId   *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
}

func (s GenerateUploadAuthResponseBodyUploadConfig) String() string {
	return tea.Prettify(s)
}

func (s GenerateUploadAuthResponseBodyUploadConfig) GoString() string {
	return s.String()
}

func (s *GenerateUploadAuthResponseBodyUploadConfig) SetSignature(v string) *GenerateUploadAuthResponseBodyUploadConfig {
	s.Signature = &v
	return s
}

func (s *GenerateUploadAuthResponseBodyUploadConfig) SetFilePath(v string) *GenerateUploadAuthResponseBodyUploadConfig {
	s.FilePath = &v
	return s
}

func (s *GenerateUploadAuthResponseBodyUploadConfig) SetPolicy(v string) *GenerateUploadAuthResponseBodyUploadConfig {
	s.Policy = &v
	return s
}

func (s *GenerateUploadAuthResponseBodyUploadConfig) SetExpireTime(v int64) *GenerateUploadAuthResponseBodyUploadConfig {
	s.ExpireTime = &v
	return s
}

func (s *GenerateUploadAuthResponseBodyUploadConfig) SetUploadHost(v string) *GenerateUploadAuthResponseBodyUploadConfig {
	s.UploadHost = &v
	return s
}

func (s *GenerateUploadAuthResponseBodyUploadConfig) SetAccessId(v string) *GenerateUploadAuthResponseBodyUploadConfig {
	s.AccessId = &v
	return s
}

type GenerateUploadAuthResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GenerateUploadAuthResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateUploadAuthResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateUploadAuthResponse) GoString() string {
	return s.String()
}

func (s *GenerateUploadAuthResponse) SetHeaders(v map[string]*string) *GenerateUploadAuthResponse {
	s.Headers = v
	return s
}

func (s *GenerateUploadAuthResponse) SetBody(v *GenerateUploadAuthResponseBody) *GenerateUploadAuthResponse {
	s.Body = v
	return s
}

type GrantServiceRoleRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GrantServiceRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s GrantServiceRoleRequest) GoString() string {
	return s.String()
}

func (s *GrantServiceRoleRequest) SetRegionId(v string) *GrantServiceRoleRequest {
	s.RegionId = &v
	return s
}

type GrantServiceRoleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GrantServiceRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GrantServiceRoleResponseBody) GoString() string {
	return s.String()
}

func (s *GrantServiceRoleResponseBody) SetRequestId(v string) *GrantServiceRoleResponseBody {
	s.RequestId = &v
	return s
}

type GrantServiceRoleResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GrantServiceRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GrantServiceRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s GrantServiceRoleResponse) GoString() string {
	return s.String()
}

func (s *GrantServiceRoleResponse) SetHeaders(v map[string]*string) *GrantServiceRoleResponse {
	s.Headers = v
	return s
}

func (s *GrantServiceRoleResponse) SetBody(v *GrantServiceRoleResponseBody) *GrantServiceRoleResponse {
	s.Body = v
	return s
}

type ListTagKeysRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage  *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s ListTagKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequest) SetRegionId(v string) *ListTagKeysRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceType(v string) *ListTagKeysRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagKeysRequest) SetPageSize(v int32) *ListTagKeysRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagKeysRequest) SetCurrentPage(v int32) *ListTagKeysRequest {
	s.CurrentPage = &v
	return s
}

type ListTagKeysResponseBody struct {
	CurrentPage *int32                            `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	RequestId   *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize    *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TagKeys     []*ListTagKeysResponseBodyTagKeys `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
}

func (s ListTagKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBody) SetCurrentPage(v int32) *ListTagKeysResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListTagKeysResponseBody) SetRequestId(v string) *ListTagKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagKeysResponseBody) SetPageSize(v int32) *ListTagKeysResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTagKeysResponseBody) SetTotalCount(v int32) *ListTagKeysResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTagKeysResponseBody) SetTagKeys(v []*ListTagKeysResponseBodyTagKeys) *ListTagKeysResponseBody {
	s.TagKeys = v
	return s
}

type ListTagKeysResponseBodyTagKeys struct {
	TagCount *int32  `json:"TagCount,omitempty" xml:"TagCount,omitempty"`
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagKeysResponseBodyTagKeys) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBodyTagKeys) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBodyTagKeys) SetTagCount(v int32) *ListTagKeysResponseBodyTagKeys {
	s.TagCount = &v
	return s
}

func (s *ListTagKeysResponseBodyTagKeys) SetTagKey(v string) *ListTagKeysResponseBodyTagKeys {
	s.TagKey = &v
	return s
}

type ListTagKeysResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponse) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponse) SetHeaders(v map[string]*string) *ListTagKeysResponse {
	s.Headers = v
	return s
}

func (s *ListTagKeysResponse) SetBody(v *ListTagKeysResponseBody) *ListTagKeysResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	RegionId     *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	NextToken    *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceId   []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag          []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	NextToken    *string                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources []*ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v []*ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagValue = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagKey = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ModifyInstanceAttributeRequest struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyInstanceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttributeRequest) SetInstanceId(v string) *ModifyInstanceAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetDescription(v string) *ModifyInstanceAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetRegionId(v string) *ModifyInstanceAttributeRequest {
	s.RegionId = &v
	return s
}

type ModifyInstanceAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttributeResponseBody) SetRequestId(v string) *ModifyInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyInstanceAttributeResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyInstanceAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttributeResponse) SetHeaders(v map[string]*string) *ModifyInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceAttributeResponse) SetBody(v *ModifyInstanceAttributeResponseBody) *ModifyInstanceAttributeResponse {
	s.Body = v
	return s
}

type ModifyInstanceStorageRequest struct {
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StorageSpace    *string `json:"StorageSpace,omitempty" xml:"StorageSpace,omitempty"`
	StorageCategory *string `json:"StorageCategory,omitempty" xml:"StorageCategory,omitempty"`
	StorageTime     *int32  `json:"StorageTime,omitempty" xml:"StorageTime,omitempty"`
}

func (s ModifyInstanceStorageRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceStorageRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceStorageRequest) SetInstanceId(v string) *ModifyInstanceStorageRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceStorageRequest) SetRegionId(v string) *ModifyInstanceStorageRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceStorageRequest) SetStorageSpace(v string) *ModifyInstanceStorageRequest {
	s.StorageSpace = &v
	return s
}

func (s *ModifyInstanceStorageRequest) SetStorageCategory(v string) *ModifyInstanceStorageRequest {
	s.StorageCategory = &v
	return s
}

func (s *ModifyInstanceStorageRequest) SetStorageTime(v int32) *ModifyInstanceStorageRequest {
	s.StorageTime = &v
	return s
}

type ModifyInstanceStorageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceStorageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceStorageResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceStorageResponseBody) SetRequestId(v string) *ModifyInstanceStorageResponseBody {
	s.RequestId = &v
	return s
}

type ModifyInstanceStorageResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyInstanceStorageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyInstanceStorageResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceStorageResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceStorageResponse) SetHeaders(v map[string]*string) *ModifyInstanceStorageResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceStorageResponse) SetBody(v *ModifyInstanceStorageResponseBody) *ModifyInstanceStorageResponse {
	s.Body = v
	return s
}

type MoveResourceGroupRequest struct {
	ResourceId      *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceType    *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s MoveResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupRequest) SetResourceId(v string) *MoveResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceGroupId(v string) *MoveResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceType(v string) *MoveResourceGroupRequest {
	s.ResourceType = &v
	return s
}

func (s *MoveResourceGroupRequest) SetRegionId(v string) *MoveResourceGroupRequest {
	s.RegionId = &v
	return s
}

type MoveResourceGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MoveResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponseBody) SetRequestId(v string) *MoveResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type MoveResourceGroupResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MoveResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MoveResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponse) SetHeaders(v map[string]*string) *MoveResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *MoveResourceGroupResponse) SetBody(v *MoveResourceGroupResponseBody) *MoveResourceGroupResponse {
	s.Body = v
	return s
}

type RefundInstanceRequest struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
}

func (s RefundInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RefundInstanceRequest) GoString() string {
	return s.String()
}

func (s *RefundInstanceRequest) SetInstanceId(v string) *RefundInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RefundInstanceRequest) SetRegionId(v string) *RefundInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *RefundInstanceRequest) SetServiceCode(v string) *RefundInstanceRequest {
	s.ServiceCode = &v
	return s
}

type RefundInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefundInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefundInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RefundInstanceResponseBody) SetRequestId(v string) *RefundInstanceResponseBody {
	s.RequestId = &v
	return s
}

type RefundInstanceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RefundInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RefundInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RefundInstanceResponse) GoString() string {
	return s.String()
}

func (s *RefundInstanceResponse) SetHeaders(v map[string]*string) *RefundInstanceResponse {
	s.Headers = v
	return s
}

func (s *RefundInstanceResponse) SetBody(v *RefundInstanceResponseBody) *RefundInstanceResponse {
	s.Body = v
	return s
}

type StartInstanceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	VswitchId  *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StartInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartInstanceRequest) SetInstanceId(v string) *StartInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *StartInstanceRequest) SetVswitchId(v string) *StartInstanceRequest {
	s.VswitchId = &v
	return s
}

func (s *StartInstanceRequest) SetRegionId(v string) *StartInstanceRequest {
	s.RegionId = &v
	return s
}

type StartInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartInstanceResponseBody) SetRequestId(v string) *StartInstanceResponseBody {
	s.RequestId = &v
	return s
}

type StartInstanceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartInstanceResponse) SetHeaders(v map[string]*string) *StartInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartInstanceResponse) SetBody(v *StartInstanceResponseBody) *StartInstanceResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	RegionId     *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ResourceId   []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag          []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	All          *bool     `json:"All,omitempty" xml:"All,omitempty"`
	ResourceId   []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	TagKey       []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UntagResourcesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
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
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("yundun-dbaudit"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) ClearInstanceStorageWithOptions(request *ClearInstanceStorageRequest, runtime *util.RuntimeOptions) (_result *ClearInstanceStorageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ClearInstanceStorageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ClearInstanceStorage"), tea.String("2018-10-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ClearInstanceStorage(request *ClearInstanceStorageRequest) (_result *ClearInstanceStorageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ClearInstanceStorageResponse{}
	_body, _err := client.ClearInstanceStorageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfigInstanceWhiteListWithOptions(request *ConfigInstanceWhiteListRequest, runtime *util.RuntimeOptions) (_result *ConfigInstanceWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConfigInstanceWhiteListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConfigInstanceWhiteList"), tea.String("2018-10-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigInstanceWhiteList(request *ConfigInstanceWhiteListRequest) (_result *ConfigInstanceWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigInstanceWhiteListResponse{}
	_body, _err := client.ConfigInstanceWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAuditLogsWithOptions(request *DescribeAuditLogsRequest, runtime *util.RuntimeOptions) (_result *DescribeAuditLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAuditLogsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAuditLogs"), tea.String("2018-10-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAuditLogs(request *DescribeAuditLogsRequest) (_result *DescribeAuditLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAuditLogsResponse{}
	_body, _err := client.DescribeAuditLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceAttribueWithOptions(request *DescribeInstanceAttribueRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceAttribueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstanceAttribueResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstanceAttribue"), tea.String("2018-10-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceAttribue(request *DescribeInstanceAttribueRequest) (_result *DescribeInstanceAttribueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceAttribueResponse{}
	_body, _err := client.DescribeInstanceAttribueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceAttributeWithOptions(request *DescribeInstanceAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstanceAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstanceAttribute"), tea.String("2018-10-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceAttribute(request *DescribeInstanceAttributeRequest) (_result *DescribeInstanceAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceAttributeResponse{}
	_body, _err := client.DescribeInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstancesWithOptions(request *DescribeInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstances"), tea.String("2018-10-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstances(request *DescribeInstancesRequest) (_result *DescribeInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.DescribeInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceStorageWithOptions(request *DescribeInstanceStorageRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceStorageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstanceStorageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstanceStorage"), tea.String("2018-10-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceStorage(request *DescribeInstanceStorageRequest) (_result *DescribeInstanceStorageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceStorageResponse{}
	_body, _err := client.DescribeInstanceStorageWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DescribeRegions"), tea.String("2018-10-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeRenewStatusWithOptions(request *DescribeRenewStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeRenewStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRenewStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRenewStatus"), tea.String("2018-10-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRenewStatus(request *DescribeRenewStatusRequest) (_result *DescribeRenewStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRenewStatusResponse{}
	_body, _err := client.DescribeRenewStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSessionLogsWithOptions(request *DescribeSessionLogsRequest, runtime *util.RuntimeOptions) (_result *DescribeSessionLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSessionLogsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSessionLogs"), tea.String("2018-10-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSessionLogs(request *DescribeSessionLogsRequest) (_result *DescribeSessionLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSessionLogsResponse{}
	_body, _err := client.DescribeSessionLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableInstancePublicAccessWithOptions(request *DisableInstancePublicAccessRequest, runtime *util.RuntimeOptions) (_result *DisableInstancePublicAccessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DisableInstancePublicAccessResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DisableInstancePublicAccess"), tea.String("2018-10-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableInstancePublicAccess(request *DisableInstancePublicAccessRequest) (_result *DisableInstancePublicAccessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableInstancePublicAccessResponse{}
	_body, _err := client.DisableInstancePublicAccessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableInstancePublicAccessWithOptions(request *EnableInstancePublicAccessRequest, runtime *util.RuntimeOptions) (_result *EnableInstancePublicAccessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableInstancePublicAccessResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableInstancePublicAccess"), tea.String("2018-10-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableInstancePublicAccess(request *EnableInstancePublicAccessRequest) (_result *EnableInstancePublicAccessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableInstancePublicAccessResponse{}
	_body, _err := client.EnableInstancePublicAccessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateUploadAuthWithOptions(request *GenerateUploadAuthRequest, runtime *util.RuntimeOptions) (_result *GenerateUploadAuthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GenerateUploadAuthResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GenerateUploadAuth"), tea.String("2018-10-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateUploadAuth(request *GenerateUploadAuthRequest) (_result *GenerateUploadAuthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateUploadAuthResponse{}
	_body, _err := client.GenerateUploadAuthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GrantServiceRoleWithOptions(request *GrantServiceRoleRequest, runtime *util.RuntimeOptions) (_result *GrantServiceRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GrantServiceRoleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GrantServiceRole"), tea.String("2018-10-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GrantServiceRole(request *GrantServiceRoleRequest) (_result *GrantServiceRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GrantServiceRoleResponse{}
	_body, _err := client.GrantServiceRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagKeysWithOptions(request *ListTagKeysRequest, runtime *util.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagKeysResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagKeys"), tea.String("2018-10-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagKeys(request *ListTagKeysRequest) (_result *ListTagKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagKeysResponse{}
	_body, _err := client.ListTagKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagResources"), tea.String("2018-10-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyInstanceAttributeWithOptions(request *ModifyInstanceAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyInstanceAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyInstanceAttribute"), tea.String("2018-10-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyInstanceAttribute(request *ModifyInstanceAttributeRequest) (_result *ModifyInstanceAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceAttributeResponse{}
	_body, _err := client.ModifyInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyInstanceStorageWithOptions(request *ModifyInstanceStorageRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceStorageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyInstanceStorageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyInstanceStorage"), tea.String("2018-10-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyInstanceStorage(request *ModifyInstanceStorageRequest) (_result *ModifyInstanceStorageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceStorageResponse{}
	_body, _err := client.ModifyInstanceStorageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MoveResourceGroupWithOptions(request *MoveResourceGroupRequest, runtime *util.RuntimeOptions) (_result *MoveResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("MoveResourceGroup"), tea.String("2018-10-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MoveResourceGroup(request *MoveResourceGroupRequest) (_result *MoveResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.MoveResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefundInstanceWithOptions(request *RefundInstanceRequest, runtime *util.RuntimeOptions) (_result *RefundInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RefundInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RefundInstance"), tea.String("2018-10-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefundInstance(request *RefundInstanceRequest) (_result *RefundInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefundInstanceResponse{}
	_body, _err := client.RefundInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartInstanceWithOptions(request *StartInstanceRequest, runtime *util.RuntimeOptions) (_result *StartInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartInstance"), tea.String("2018-10-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartInstance(request *StartInstanceRequest) (_result *StartInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartInstanceResponse{}
	_body, _err := client.StartInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TagResources"), tea.String("2018-10-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UntagResources"), tea.String("2018-10-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
