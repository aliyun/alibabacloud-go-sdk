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

type AddZoneRequest struct {
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ZoneName        *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
	UserClientIp    *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	ProxyPattern    *string `json:"ProxyPattern,omitempty" xml:"ProxyPattern,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ZoneType        *string `json:"ZoneType,omitempty" xml:"ZoneType,omitempty"`
	ZoneTag         *string `json:"ZoneTag,omitempty" xml:"ZoneTag,omitempty"`
}

func (s AddZoneRequest) String() string {
	return tea.Prettify(s)
}

func (s AddZoneRequest) GoString() string {
	return s.String()
}

func (s *AddZoneRequest) SetLang(v string) *AddZoneRequest {
	s.Lang = &v
	return s
}

func (s *AddZoneRequest) SetZoneName(v string) *AddZoneRequest {
	s.ZoneName = &v
	return s
}

func (s *AddZoneRequest) SetUserClientIp(v string) *AddZoneRequest {
	s.UserClientIp = &v
	return s
}

func (s *AddZoneRequest) SetProxyPattern(v string) *AddZoneRequest {
	s.ProxyPattern = &v
	return s
}

func (s *AddZoneRequest) SetResourceGroupId(v string) *AddZoneRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AddZoneRequest) SetZoneType(v string) *AddZoneRequest {
	s.ZoneType = &v
	return s
}

func (s *AddZoneRequest) SetZoneTag(v string) *AddZoneRequest {
	s.ZoneTag = &v
	return s
}

type AddZoneResponseBody struct {
	ZoneName  *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
	ZoneId    *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddZoneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddZoneResponseBody) GoString() string {
	return s.String()
}

func (s *AddZoneResponseBody) SetZoneName(v string) *AddZoneResponseBody {
	s.ZoneName = &v
	return s
}

func (s *AddZoneResponseBody) SetZoneId(v string) *AddZoneResponseBody {
	s.ZoneId = &v
	return s
}

func (s *AddZoneResponseBody) SetRequestId(v string) *AddZoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddZoneResponseBody) SetSuccess(v bool) *AddZoneResponseBody {
	s.Success = &v
	return s
}

type AddZoneResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddZoneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddZoneResponse) String() string {
	return tea.Prettify(s)
}

func (s AddZoneResponse) GoString() string {
	return s.String()
}

func (s *AddZoneResponse) SetHeaders(v map[string]*string) *AddZoneResponse {
	s.Headers = v
	return s
}

func (s *AddZoneResponse) SetBody(v *AddZoneResponseBody) *AddZoneResponse {
	s.Body = v
	return s
}

type AddZoneRecordRequest struct {
	ZoneId       *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Rr           *string `json:"Rr,omitempty" xml:"Rr,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Ttl          *int32  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	Priority     *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Value        *string `json:"Value,omitempty" xml:"Value,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s AddZoneRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s AddZoneRecordRequest) GoString() string {
	return s.String()
}

func (s *AddZoneRecordRequest) SetZoneId(v string) *AddZoneRecordRequest {
	s.ZoneId = &v
	return s
}

func (s *AddZoneRecordRequest) SetLang(v string) *AddZoneRecordRequest {
	s.Lang = &v
	return s
}

func (s *AddZoneRecordRequest) SetRr(v string) *AddZoneRecordRequest {
	s.Rr = &v
	return s
}

func (s *AddZoneRecordRequest) SetType(v string) *AddZoneRecordRequest {
	s.Type = &v
	return s
}

func (s *AddZoneRecordRequest) SetTtl(v int32) *AddZoneRecordRequest {
	s.Ttl = &v
	return s
}

func (s *AddZoneRecordRequest) SetPriority(v int32) *AddZoneRecordRequest {
	s.Priority = &v
	return s
}

func (s *AddZoneRecordRequest) SetValue(v string) *AddZoneRecordRequest {
	s.Value = &v
	return s
}

func (s *AddZoneRecordRequest) SetUserClientIp(v string) *AddZoneRecordRequest {
	s.UserClientIp = &v
	return s
}

type AddZoneRecordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RecordId  *int64  `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddZoneRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddZoneRecordResponseBody) GoString() string {
	return s.String()
}

func (s *AddZoneRecordResponseBody) SetRequestId(v string) *AddZoneRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddZoneRecordResponseBody) SetRecordId(v int64) *AddZoneRecordResponseBody {
	s.RecordId = &v
	return s
}

func (s *AddZoneRecordResponseBody) SetSuccess(v bool) *AddZoneRecordResponseBody {
	s.Success = &v
	return s
}

type AddZoneRecordResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddZoneRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddZoneRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s AddZoneRecordResponse) GoString() string {
	return s.String()
}

func (s *AddZoneRecordResponse) SetHeaders(v map[string]*string) *AddZoneRecordResponse {
	s.Headers = v
	return s
}

func (s *AddZoneRecordResponse) SetBody(v *AddZoneRecordResponseBody) *AddZoneRecordResponse {
	s.Body = v
	return s
}

type BindZoneVpcRequest struct {
	Lang         *string                   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ZoneId       *string                   `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	UserClientIp *string                   `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	Vpcs         []*BindZoneVpcRequestVpcs `json:"Vpcs,omitempty" xml:"Vpcs,omitempty" type:"Repeated"`
}

func (s BindZoneVpcRequest) String() string {
	return tea.Prettify(s)
}

func (s BindZoneVpcRequest) GoString() string {
	return s.String()
}

func (s *BindZoneVpcRequest) SetLang(v string) *BindZoneVpcRequest {
	s.Lang = &v
	return s
}

func (s *BindZoneVpcRequest) SetZoneId(v string) *BindZoneVpcRequest {
	s.ZoneId = &v
	return s
}

func (s *BindZoneVpcRequest) SetUserClientIp(v string) *BindZoneVpcRequest {
	s.UserClientIp = &v
	return s
}

func (s *BindZoneVpcRequest) SetVpcs(v []*BindZoneVpcRequestVpcs) *BindZoneVpcRequest {
	s.Vpcs = v
	return s
}

type BindZoneVpcRequestVpcs struct {
	VpcId    *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s BindZoneVpcRequestVpcs) String() string {
	return tea.Prettify(s)
}

func (s BindZoneVpcRequestVpcs) GoString() string {
	return s.String()
}

func (s *BindZoneVpcRequestVpcs) SetVpcId(v string) *BindZoneVpcRequestVpcs {
	s.VpcId = &v
	return s
}

func (s *BindZoneVpcRequestVpcs) SetRegionId(v string) *BindZoneVpcRequestVpcs {
	s.RegionId = &v
	return s
}

type BindZoneVpcResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindZoneVpcResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindZoneVpcResponseBody) GoString() string {
	return s.String()
}

func (s *BindZoneVpcResponseBody) SetRequestId(v string) *BindZoneVpcResponseBody {
	s.RequestId = &v
	return s
}

type BindZoneVpcResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BindZoneVpcResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindZoneVpcResponse) String() string {
	return tea.Prettify(s)
}

func (s BindZoneVpcResponse) GoString() string {
	return s.String()
}

func (s *BindZoneVpcResponse) SetHeaders(v map[string]*string) *BindZoneVpcResponse {
	s.Headers = v
	return s
}

func (s *BindZoneVpcResponse) SetBody(v *BindZoneVpcResponseBody) *BindZoneVpcResponse {
	s.Body = v
	return s
}

type CheckZoneNameRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ZoneName     *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s CheckZoneNameRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckZoneNameRequest) GoString() string {
	return s.String()
}

func (s *CheckZoneNameRequest) SetLang(v string) *CheckZoneNameRequest {
	s.Lang = &v
	return s
}

func (s *CheckZoneNameRequest) SetZoneName(v string) *CheckZoneNameRequest {
	s.ZoneName = &v
	return s
}

func (s *CheckZoneNameRequest) SetUserClientIp(v string) *CheckZoneNameRequest {
	s.UserClientIp = &v
	return s
}

type CheckZoneNameResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Check     *bool   `json:"Check,omitempty" xml:"Check,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckZoneNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckZoneNameResponseBody) GoString() string {
	return s.String()
}

func (s *CheckZoneNameResponseBody) SetRequestId(v string) *CheckZoneNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckZoneNameResponseBody) SetCheck(v bool) *CheckZoneNameResponseBody {
	s.Check = &v
	return s
}

func (s *CheckZoneNameResponseBody) SetSuccess(v bool) *CheckZoneNameResponseBody {
	s.Success = &v
	return s
}

type CheckZoneNameResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckZoneNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckZoneNameResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckZoneNameResponse) GoString() string {
	return s.String()
}

func (s *CheckZoneNameResponse) SetHeaders(v map[string]*string) *CheckZoneNameResponse {
	s.Headers = v
	return s
}

func (s *CheckZoneNameResponse) SetBody(v *CheckZoneNameResponseBody) *CheckZoneNameResponse {
	s.Body = v
	return s
}

type DeleteZoneRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ZoneId       *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s DeleteZoneRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteZoneRequest) GoString() string {
	return s.String()
}

func (s *DeleteZoneRequest) SetLang(v string) *DeleteZoneRequest {
	s.Lang = &v
	return s
}

func (s *DeleteZoneRequest) SetZoneId(v string) *DeleteZoneRequest {
	s.ZoneId = &v
	return s
}

func (s *DeleteZoneRequest) SetUserClientIp(v string) *DeleteZoneRequest {
	s.UserClientIp = &v
	return s
}

type DeleteZoneResponseBody struct {
	ZoneId    *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteZoneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteZoneResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteZoneResponseBody) SetZoneId(v string) *DeleteZoneResponseBody {
	s.ZoneId = &v
	return s
}

func (s *DeleteZoneResponseBody) SetRequestId(v string) *DeleteZoneResponseBody {
	s.RequestId = &v
	return s
}

type DeleteZoneResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteZoneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteZoneResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteZoneResponse) GoString() string {
	return s.String()
}

func (s *DeleteZoneResponse) SetHeaders(v map[string]*string) *DeleteZoneResponse {
	s.Headers = v
	return s
}

func (s *DeleteZoneResponse) SetBody(v *DeleteZoneResponseBody) *DeleteZoneResponse {
	s.Body = v
	return s
}

type DeleteZoneRecordRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	RecordId     *int64  `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s DeleteZoneRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteZoneRecordRequest) GoString() string {
	return s.String()
}

func (s *DeleteZoneRecordRequest) SetLang(v string) *DeleteZoneRecordRequest {
	s.Lang = &v
	return s
}

func (s *DeleteZoneRecordRequest) SetRecordId(v int64) *DeleteZoneRecordRequest {
	s.RecordId = &v
	return s
}

func (s *DeleteZoneRecordRequest) SetUserClientIp(v string) *DeleteZoneRecordRequest {
	s.UserClientIp = &v
	return s
}

type DeleteZoneRecordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RecordId  *int64  `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
}

func (s DeleteZoneRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteZoneRecordResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteZoneRecordResponseBody) SetRequestId(v string) *DeleteZoneRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteZoneRecordResponseBody) SetRecordId(v int64) *DeleteZoneRecordResponseBody {
	s.RecordId = &v
	return s
}

type DeleteZoneRecordResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteZoneRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteZoneRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteZoneRecordResponse) GoString() string {
	return s.String()
}

func (s *DeleteZoneRecordResponse) SetHeaders(v map[string]*string) *DeleteZoneRecordResponse {
	s.Headers = v
	return s
}

func (s *DeleteZoneRecordResponse) SetBody(v *DeleteZoneRecordResponseBody) *DeleteZoneRecordResponse {
	s.Body = v
	return s
}

type DescribeChangeLogsRequest struct {
	Keyword        *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	Lang           *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ZoneId         *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTimestamp *int64  `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	EndTimestamp   *int64  `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	EntityType     *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	UserClientIp   *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s DescribeChangeLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeChangeLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeChangeLogsRequest) SetKeyword(v string) *DescribeChangeLogsRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeChangeLogsRequest) SetLang(v string) *DescribeChangeLogsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeChangeLogsRequest) SetZoneId(v string) *DescribeChangeLogsRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeChangeLogsRequest) SetPageNumber(v int32) *DescribeChangeLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeChangeLogsRequest) SetPageSize(v int32) *DescribeChangeLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeChangeLogsRequest) SetStartTimestamp(v int64) *DescribeChangeLogsRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeChangeLogsRequest) SetEndTimestamp(v int64) *DescribeChangeLogsRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeChangeLogsRequest) SetEntityType(v string) *DescribeChangeLogsRequest {
	s.EntityType = &v
	return s
}

func (s *DescribeChangeLogsRequest) SetUserClientIp(v string) *DescribeChangeLogsRequest {
	s.UserClientIp = &v
	return s
}

type DescribeChangeLogsResponseBody struct {
	PageSize   *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ChangeLogs *DescribeChangeLogsResponseBodyChangeLogs `json:"ChangeLogs,omitempty" xml:"ChangeLogs,omitempty" type:"Struct"`
	PageNumber *int32                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalPages *int32                                    `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	TotalItems *int32                                    `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
}

func (s DescribeChangeLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeChangeLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChangeLogsResponseBody) SetPageSize(v int32) *DescribeChangeLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeChangeLogsResponseBody) SetRequestId(v string) *DescribeChangeLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeChangeLogsResponseBody) SetChangeLogs(v *DescribeChangeLogsResponseBodyChangeLogs) *DescribeChangeLogsResponseBody {
	s.ChangeLogs = v
	return s
}

func (s *DescribeChangeLogsResponseBody) SetPageNumber(v int32) *DescribeChangeLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeChangeLogsResponseBody) SetTotalPages(v int32) *DescribeChangeLogsResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeChangeLogsResponseBody) SetTotalItems(v int32) *DescribeChangeLogsResponseBody {
	s.TotalItems = &v
	return s
}

type DescribeChangeLogsResponseBodyChangeLogs struct {
	ChangeLog []*DescribeChangeLogsResponseBodyChangeLogsChangeLog `json:"ChangeLog,omitempty" xml:"ChangeLog,omitempty" type:"Repeated"`
}

func (s DescribeChangeLogsResponseBodyChangeLogs) String() string {
	return tea.Prettify(s)
}

func (s DescribeChangeLogsResponseBodyChangeLogs) GoString() string {
	return s.String()
}

func (s *DescribeChangeLogsResponseBodyChangeLogs) SetChangeLog(v []*DescribeChangeLogsResponseBodyChangeLogsChangeLog) *DescribeChangeLogsResponseBodyChangeLogs {
	s.ChangeLog = v
	return s
}

type DescribeChangeLogsResponseBodyChangeLogsChangeLog struct {
	OperTimestamp *int64  `json:"OperTimestamp,omitempty" xml:"OperTimestamp,omitempty"`
	EntityId      *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	OperObject    *string `json:"OperObject,omitempty" xml:"OperObject,omitempty"`
	OperTime      *string `json:"OperTime,omitempty" xml:"OperTime,omitempty"`
	OperIp        *string `json:"OperIp,omitempty" xml:"OperIp,omitempty"`
	OperAction    *string `json:"OperAction,omitempty" xml:"OperAction,omitempty"`
	Content       *string `json:"Content,omitempty" xml:"Content,omitempty"`
	EntityName    *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	Id            *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeChangeLogsResponseBodyChangeLogsChangeLog) String() string {
	return tea.Prettify(s)
}

func (s DescribeChangeLogsResponseBodyChangeLogsChangeLog) GoString() string {
	return s.String()
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) SetOperTimestamp(v int64) *DescribeChangeLogsResponseBodyChangeLogsChangeLog {
	s.OperTimestamp = &v
	return s
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) SetEntityId(v string) *DescribeChangeLogsResponseBodyChangeLogsChangeLog {
	s.EntityId = &v
	return s
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) SetOperObject(v string) *DescribeChangeLogsResponseBodyChangeLogsChangeLog {
	s.OperObject = &v
	return s
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) SetOperTime(v string) *DescribeChangeLogsResponseBodyChangeLogsChangeLog {
	s.OperTime = &v
	return s
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) SetOperIp(v string) *DescribeChangeLogsResponseBodyChangeLogsChangeLog {
	s.OperIp = &v
	return s
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) SetOperAction(v string) *DescribeChangeLogsResponseBodyChangeLogsChangeLog {
	s.OperAction = &v
	return s
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) SetContent(v string) *DescribeChangeLogsResponseBodyChangeLogsChangeLog {
	s.Content = &v
	return s
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) SetEntityName(v string) *DescribeChangeLogsResponseBodyChangeLogsChangeLog {
	s.EntityName = &v
	return s
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) SetId(v int64) *DescribeChangeLogsResponseBodyChangeLogsChangeLog {
	s.Id = &v
	return s
}

type DescribeChangeLogsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeChangeLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeChangeLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeChangeLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeChangeLogsResponse) SetHeaders(v map[string]*string) *DescribeChangeLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeChangeLogsResponse) SetBody(v *DescribeChangeLogsResponseBody) *DescribeChangeLogsResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	Lang             *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp     *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	AcceptLanguage   *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	AuthorizedUserId *int64  `json:"AuthorizedUserId,omitempty" xml:"AuthorizedUserId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetLang(v string) *DescribeRegionsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRegionsRequest) SetUserClientIp(v string) *DescribeRegionsRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeRegionsRequest) SetAuthorizedUserId(v int64) *DescribeRegionsRequest {
	s.AuthorizedUserId = &v
	return s
}

type DescribeRegionsResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Regions   *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
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

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
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
	LocalName      *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionName     *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetLocalName(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionName(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionsRegion {
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

type DescribeRequestGraphRequest struct {
	Lang           *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp   *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	ZoneId         *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	VpcId          *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	StartTimestamp *int64  `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	EndTimestamp   *int64  `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	BizType        *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	BizId          *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s DescribeRequestGraphRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRequestGraphRequest) GoString() string {
	return s.String()
}

func (s *DescribeRequestGraphRequest) SetLang(v string) *DescribeRequestGraphRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRequestGraphRequest) SetUserClientIp(v string) *DescribeRequestGraphRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeRequestGraphRequest) SetZoneId(v string) *DescribeRequestGraphRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeRequestGraphRequest) SetVpcId(v string) *DescribeRequestGraphRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeRequestGraphRequest) SetStartTimestamp(v int64) *DescribeRequestGraphRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeRequestGraphRequest) SetEndTimestamp(v int64) *DescribeRequestGraphRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeRequestGraphRequest) SetBizType(v string) *DescribeRequestGraphRequest {
	s.BizType = &v
	return s
}

func (s *DescribeRequestGraphRequest) SetBizId(v string) *DescribeRequestGraphRequest {
	s.BizId = &v
	return s
}

type DescribeRequestGraphResponseBody struct {
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RequestDetails *DescribeRequestGraphResponseBodyRequestDetails `json:"RequestDetails,omitempty" xml:"RequestDetails,omitempty" type:"Struct"`
}

func (s DescribeRequestGraphResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRequestGraphResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRequestGraphResponseBody) SetRequestId(v string) *DescribeRequestGraphResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRequestGraphResponseBody) SetRequestDetails(v *DescribeRequestGraphResponseBodyRequestDetails) *DescribeRequestGraphResponseBody {
	s.RequestDetails = v
	return s
}

type DescribeRequestGraphResponseBodyRequestDetails struct {
	ZoneRequestTop []*DescribeRequestGraphResponseBodyRequestDetailsZoneRequestTop `json:"ZoneRequestTop,omitempty" xml:"ZoneRequestTop,omitempty" type:"Repeated"`
}

func (s DescribeRequestGraphResponseBodyRequestDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeRequestGraphResponseBodyRequestDetails) GoString() string {
	return s.String()
}

func (s *DescribeRequestGraphResponseBodyRequestDetails) SetZoneRequestTop(v []*DescribeRequestGraphResponseBodyRequestDetailsZoneRequestTop) *DescribeRequestGraphResponseBodyRequestDetails {
	s.ZoneRequestTop = v
	return s
}

type DescribeRequestGraphResponseBodyRequestDetailsZoneRequestTop struct {
	Time         *string `json:"Time,omitempty" xml:"Time,omitempty"`
	RequestCount *int64  `json:"RequestCount,omitempty" xml:"RequestCount,omitempty"`
	Timestamp    *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeRequestGraphResponseBodyRequestDetailsZoneRequestTop) String() string {
	return tea.Prettify(s)
}

func (s DescribeRequestGraphResponseBodyRequestDetailsZoneRequestTop) GoString() string {
	return s.String()
}

func (s *DescribeRequestGraphResponseBodyRequestDetailsZoneRequestTop) SetTime(v string) *DescribeRequestGraphResponseBodyRequestDetailsZoneRequestTop {
	s.Time = &v
	return s
}

func (s *DescribeRequestGraphResponseBodyRequestDetailsZoneRequestTop) SetRequestCount(v int64) *DescribeRequestGraphResponseBodyRequestDetailsZoneRequestTop {
	s.RequestCount = &v
	return s
}

func (s *DescribeRequestGraphResponseBodyRequestDetailsZoneRequestTop) SetTimestamp(v int64) *DescribeRequestGraphResponseBodyRequestDetailsZoneRequestTop {
	s.Timestamp = &v
	return s
}

type DescribeRequestGraphResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRequestGraphResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRequestGraphResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRequestGraphResponse) GoString() string {
	return s.String()
}

func (s *DescribeRequestGraphResponse) SetHeaders(v map[string]*string) *DescribeRequestGraphResponse {
	s.Headers = v
	return s
}

func (s *DescribeRequestGraphResponse) SetBody(v *DescribeRequestGraphResponseBody) *DescribeRequestGraphResponse {
	s.Body = v
	return s
}

type DescribeStatisticSummaryRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s DescribeStatisticSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeStatisticSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeStatisticSummaryRequest) SetLang(v string) *DescribeStatisticSummaryRequest {
	s.Lang = &v
	return s
}

func (s *DescribeStatisticSummaryRequest) SetUserClientIp(v string) *DescribeStatisticSummaryRequest {
	s.UserClientIp = &v
	return s
}

type DescribeStatisticSummaryResponseBody struct {
	TotalCount      *int64                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId       *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ZoneRequestTops *DescribeStatisticSummaryResponseBodyZoneRequestTops `json:"ZoneRequestTops,omitempty" xml:"ZoneRequestTops,omitempty" type:"Struct"`
	VpcRequestTops  *DescribeStatisticSummaryResponseBodyVpcRequestTops  `json:"VpcRequestTops,omitempty" xml:"VpcRequestTops,omitempty" type:"Struct"`
}

func (s DescribeStatisticSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeStatisticSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStatisticSummaryResponseBody) SetTotalCount(v int64) *DescribeStatisticSummaryResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeStatisticSummaryResponseBody) SetRequestId(v string) *DescribeStatisticSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStatisticSummaryResponseBody) SetZoneRequestTops(v *DescribeStatisticSummaryResponseBodyZoneRequestTops) *DescribeStatisticSummaryResponseBody {
	s.ZoneRequestTops = v
	return s
}

func (s *DescribeStatisticSummaryResponseBody) SetVpcRequestTops(v *DescribeStatisticSummaryResponseBodyVpcRequestTops) *DescribeStatisticSummaryResponseBody {
	s.VpcRequestTops = v
	return s
}

type DescribeStatisticSummaryResponseBodyZoneRequestTops struct {
	ZoneRequestTop []*DescribeStatisticSummaryResponseBodyZoneRequestTopsZoneRequestTop `json:"ZoneRequestTop,omitempty" xml:"ZoneRequestTop,omitempty" type:"Repeated"`
}

func (s DescribeStatisticSummaryResponseBodyZoneRequestTops) String() string {
	return tea.Prettify(s)
}

func (s DescribeStatisticSummaryResponseBodyZoneRequestTops) GoString() string {
	return s.String()
}

func (s *DescribeStatisticSummaryResponseBodyZoneRequestTops) SetZoneRequestTop(v []*DescribeStatisticSummaryResponseBodyZoneRequestTopsZoneRequestTop) *DescribeStatisticSummaryResponseBodyZoneRequestTops {
	s.ZoneRequestTop = v
	return s
}

type DescribeStatisticSummaryResponseBodyZoneRequestTopsZoneRequestTop struct {
	RequestCount *int64  `json:"RequestCount,omitempty" xml:"RequestCount,omitempty"`
	ZoneName     *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
	BizType      *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
}

func (s DescribeStatisticSummaryResponseBodyZoneRequestTopsZoneRequestTop) String() string {
	return tea.Prettify(s)
}

func (s DescribeStatisticSummaryResponseBodyZoneRequestTopsZoneRequestTop) GoString() string {
	return s.String()
}

func (s *DescribeStatisticSummaryResponseBodyZoneRequestTopsZoneRequestTop) SetRequestCount(v int64) *DescribeStatisticSummaryResponseBodyZoneRequestTopsZoneRequestTop {
	s.RequestCount = &v
	return s
}

func (s *DescribeStatisticSummaryResponseBodyZoneRequestTopsZoneRequestTop) SetZoneName(v string) *DescribeStatisticSummaryResponseBodyZoneRequestTopsZoneRequestTop {
	s.ZoneName = &v
	return s
}

func (s *DescribeStatisticSummaryResponseBodyZoneRequestTopsZoneRequestTop) SetBizType(v string) *DescribeStatisticSummaryResponseBodyZoneRequestTopsZoneRequestTop {
	s.BizType = &v
	return s
}

type DescribeStatisticSummaryResponseBodyVpcRequestTops struct {
	VpcRequestTop []*DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop `json:"VpcRequestTop,omitempty" xml:"VpcRequestTop,omitempty" type:"Repeated"`
}

func (s DescribeStatisticSummaryResponseBodyVpcRequestTops) String() string {
	return tea.Prettify(s)
}

func (s DescribeStatisticSummaryResponseBodyVpcRequestTops) GoString() string {
	return s.String()
}

func (s *DescribeStatisticSummaryResponseBodyVpcRequestTops) SetVpcRequestTop(v []*DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop) *DescribeStatisticSummaryResponseBodyVpcRequestTops {
	s.VpcRequestTop = v
	return s
}

type DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop struct {
	VpcId        *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	RegionName   *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	TunnelId     *string `json:"TunnelId,omitempty" xml:"TunnelId,omitempty"`
	RequestCount *int64  `json:"RequestCount,omitempty" xml:"RequestCount,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop) String() string {
	return tea.Prettify(s)
}

func (s DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop) GoString() string {
	return s.String()
}

func (s *DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop) SetVpcId(v string) *DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop {
	s.VpcId = &v
	return s
}

func (s *DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop) SetRegionName(v string) *DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop {
	s.RegionName = &v
	return s
}

func (s *DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop) SetTunnelId(v string) *DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop {
	s.TunnelId = &v
	return s
}

func (s *DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop) SetRequestCount(v int64) *DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop {
	s.RequestCount = &v
	return s
}

func (s *DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop) SetRegionId(v string) *DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop {
	s.RegionId = &v
	return s
}

type DescribeStatisticSummaryResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeStatisticSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeStatisticSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeStatisticSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeStatisticSummaryResponse) SetHeaders(v map[string]*string) *DescribeStatisticSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeStatisticSummaryResponse) SetBody(v *DescribeStatisticSummaryResponseBody) *DescribeStatisticSummaryResponse {
	s.Body = v
	return s
}

type DescribeUserServiceStatusRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s DescribeUserServiceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserServiceStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserServiceStatusRequest) SetLang(v string) *DescribeUserServiceStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeUserServiceStatusRequest) SetUserClientIp(v string) *DescribeUserServiceStatusRequest {
	s.UserClientIp = &v
	return s
}

type DescribeUserServiceStatusResponseBody struct {
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserServiceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserServiceStatusResponseBody) SetStatus(v string) *DescribeUserServiceStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeUserServiceStatusResponseBody) SetRequestId(v string) *DescribeUserServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

type DescribeUserServiceStatusResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUserServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserServiceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserServiceStatusResponse) SetHeaders(v map[string]*string) *DescribeUserServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserServiceStatusResponse) SetBody(v *DescribeUserServiceStatusResponseBody) *DescribeUserServiceStatusResponse {
	s.Body = v
	return s
}

type DescribeZoneInfoRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ZoneId       *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s DescribeZoneInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeZoneInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeZoneInfoRequest) SetLang(v string) *DescribeZoneInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeZoneInfoRequest) SetZoneId(v string) *DescribeZoneInfoRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeZoneInfoRequest) SetUserClientIp(v string) *DescribeZoneInfoRequest {
	s.UserClientIp = &v
	return s
}

type DescribeZoneInfoResponseBody struct {
	RequestId       *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SlaveDns        *bool                                 `json:"SlaveDns,omitempty" xml:"SlaveDns,omitempty"`
	ResourceGroupId *string                               `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ZoneId          *string                               `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ProxyPattern    *string                               `json:"ProxyPattern,omitempty" xml:"ProxyPattern,omitempty"`
	CreateTime      *string                               `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ZoneType        *string                               `json:"ZoneType,omitempty" xml:"ZoneType,omitempty"`
	Remark          *string                               `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ZoneName        *string                               `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
	ZoneTag         *string                               `json:"ZoneTag,omitempty" xml:"ZoneTag,omitempty"`
	UpdateTime      *string                               `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UpdateTimestamp *int64                                `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	RecordCount     *int32                                `json:"RecordCount,omitempty" xml:"RecordCount,omitempty"`
	CreateTimestamp *int64                                `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	BindVpcs        *DescribeZoneInfoResponseBodyBindVpcs `json:"BindVpcs,omitempty" xml:"BindVpcs,omitempty" type:"Struct"`
	IsPtr           *bool                                 `json:"IsPtr,omitempty" xml:"IsPtr,omitempty"`
}

func (s DescribeZoneInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeZoneInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeZoneInfoResponseBody) SetRequestId(v string) *DescribeZoneInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetSlaveDns(v bool) *DescribeZoneInfoResponseBody {
	s.SlaveDns = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetResourceGroupId(v string) *DescribeZoneInfoResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetZoneId(v string) *DescribeZoneInfoResponseBody {
	s.ZoneId = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetProxyPattern(v string) *DescribeZoneInfoResponseBody {
	s.ProxyPattern = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetCreateTime(v string) *DescribeZoneInfoResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetZoneType(v string) *DescribeZoneInfoResponseBody {
	s.ZoneType = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetRemark(v string) *DescribeZoneInfoResponseBody {
	s.Remark = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetZoneName(v string) *DescribeZoneInfoResponseBody {
	s.ZoneName = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetZoneTag(v string) *DescribeZoneInfoResponseBody {
	s.ZoneTag = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetUpdateTime(v string) *DescribeZoneInfoResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetUpdateTimestamp(v int64) *DescribeZoneInfoResponseBody {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetRecordCount(v int32) *DescribeZoneInfoResponseBody {
	s.RecordCount = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetCreateTimestamp(v int64) *DescribeZoneInfoResponseBody {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetBindVpcs(v *DescribeZoneInfoResponseBodyBindVpcs) *DescribeZoneInfoResponseBody {
	s.BindVpcs = v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetIsPtr(v bool) *DescribeZoneInfoResponseBody {
	s.IsPtr = &v
	return s
}

type DescribeZoneInfoResponseBodyBindVpcs struct {
	Vpc []*DescribeZoneInfoResponseBodyBindVpcsVpc `json:"Vpc,omitempty" xml:"Vpc,omitempty" type:"Repeated"`
}

func (s DescribeZoneInfoResponseBodyBindVpcs) String() string {
	return tea.Prettify(s)
}

func (s DescribeZoneInfoResponseBodyBindVpcs) GoString() string {
	return s.String()
}

func (s *DescribeZoneInfoResponseBodyBindVpcs) SetVpc(v []*DescribeZoneInfoResponseBodyBindVpcsVpc) *DescribeZoneInfoResponseBodyBindVpcs {
	s.Vpc = v
	return s
}

type DescribeZoneInfoResponseBodyBindVpcsVpc struct {
	VpcName    *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
	VpcId      *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	VpcUserId  *int64  `json:"VpcUserId,omitempty" xml:"VpcUserId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeZoneInfoResponseBodyBindVpcsVpc) String() string {
	return tea.Prettify(s)
}

func (s DescribeZoneInfoResponseBodyBindVpcsVpc) GoString() string {
	return s.String()
}

func (s *DescribeZoneInfoResponseBodyBindVpcsVpc) SetVpcName(v string) *DescribeZoneInfoResponseBodyBindVpcsVpc {
	s.VpcName = &v
	return s
}

func (s *DescribeZoneInfoResponseBodyBindVpcsVpc) SetVpcId(v string) *DescribeZoneInfoResponseBodyBindVpcsVpc {
	s.VpcId = &v
	return s
}

func (s *DescribeZoneInfoResponseBodyBindVpcsVpc) SetRegionName(v string) *DescribeZoneInfoResponseBodyBindVpcsVpc {
	s.RegionName = &v
	return s
}

func (s *DescribeZoneInfoResponseBodyBindVpcsVpc) SetVpcUserId(v int64) *DescribeZoneInfoResponseBodyBindVpcsVpc {
	s.VpcUserId = &v
	return s
}

func (s *DescribeZoneInfoResponseBodyBindVpcsVpc) SetRegionId(v string) *DescribeZoneInfoResponseBodyBindVpcsVpc {
	s.RegionId = &v
	return s
}

type DescribeZoneInfoResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeZoneInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeZoneInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeZoneInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeZoneInfoResponse) SetHeaders(v map[string]*string) *DescribeZoneInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeZoneInfoResponse) SetBody(v *DescribeZoneInfoResponseBody) *DescribeZoneInfoResponse {
	s.Body = v
	return s
}

type DescribeZoneRecordsRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Keyword      *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	ZoneId       *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	Tag          *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	SearchMode   *string `json:"SearchMode,omitempty" xml:"SearchMode,omitempty"`
	OrderBy      *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	Direction    *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
}

func (s DescribeZoneRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeZoneRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeZoneRecordsRequest) SetLang(v string) *DescribeZoneRecordsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeZoneRecordsRequest) SetKeyword(v string) *DescribeZoneRecordsRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeZoneRecordsRequest) SetZoneId(v string) *DescribeZoneRecordsRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeZoneRecordsRequest) SetPageNumber(v int32) *DescribeZoneRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeZoneRecordsRequest) SetPageSize(v int32) *DescribeZoneRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeZoneRecordsRequest) SetUserClientIp(v string) *DescribeZoneRecordsRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeZoneRecordsRequest) SetTag(v string) *DescribeZoneRecordsRequest {
	s.Tag = &v
	return s
}

func (s *DescribeZoneRecordsRequest) SetSearchMode(v string) *DescribeZoneRecordsRequest {
	s.SearchMode = &v
	return s
}

func (s *DescribeZoneRecordsRequest) SetOrderBy(v string) *DescribeZoneRecordsRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeZoneRecordsRequest) SetDirection(v string) *DescribeZoneRecordsRequest {
	s.Direction = &v
	return s
}

type DescribeZoneRecordsResponseBody struct {
	PageSize   *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalPages *int32                                  `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	TotalItems *int32                                  `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	Records    *DescribeZoneRecordsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Struct"`
}

func (s DescribeZoneRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeZoneRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeZoneRecordsResponseBody) SetPageSize(v int32) *DescribeZoneRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeZoneRecordsResponseBody) SetRequestId(v string) *DescribeZoneRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeZoneRecordsResponseBody) SetPageNumber(v int32) *DescribeZoneRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeZoneRecordsResponseBody) SetTotalPages(v int32) *DescribeZoneRecordsResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeZoneRecordsResponseBody) SetTotalItems(v int32) *DescribeZoneRecordsResponseBody {
	s.TotalItems = &v
	return s
}

func (s *DescribeZoneRecordsResponseBody) SetRecords(v *DescribeZoneRecordsResponseBodyRecords) *DescribeZoneRecordsResponseBody {
	s.Records = v
	return s
}

type DescribeZoneRecordsResponseBodyRecords struct {
	Record []*DescribeZoneRecordsResponseBodyRecordsRecord `json:"Record,omitempty" xml:"Record,omitempty" type:"Repeated"`
}

func (s DescribeZoneRecordsResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeZoneRecordsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *DescribeZoneRecordsResponseBodyRecords) SetRecord(v []*DescribeZoneRecordsResponseBodyRecordsRecord) *DescribeZoneRecordsResponseBodyRecords {
	s.Record = v
	return s
}

type DescribeZoneRecordsResponseBodyRecordsRecord struct {
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Ttl      *int32  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	Remark   *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	RecordId *int64  `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	Rr       *string `json:"Rr,omitempty" xml:"Rr,omitempty"`
	Priority *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s DescribeZoneRecordsResponseBodyRecordsRecord) String() string {
	return tea.Prettify(s)
}

func (s DescribeZoneRecordsResponseBodyRecordsRecord) GoString() string {
	return s.String()
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) SetStatus(v string) *DescribeZoneRecordsResponseBodyRecordsRecord {
	s.Status = &v
	return s
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) SetType(v string) *DescribeZoneRecordsResponseBodyRecordsRecord {
	s.Type = &v
	return s
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) SetValue(v string) *DescribeZoneRecordsResponseBodyRecordsRecord {
	s.Value = &v
	return s
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) SetTtl(v int32) *DescribeZoneRecordsResponseBodyRecordsRecord {
	s.Ttl = &v
	return s
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) SetRemark(v string) *DescribeZoneRecordsResponseBodyRecordsRecord {
	s.Remark = &v
	return s
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) SetRecordId(v int64) *DescribeZoneRecordsResponseBodyRecordsRecord {
	s.RecordId = &v
	return s
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) SetRr(v string) *DescribeZoneRecordsResponseBodyRecordsRecord {
	s.Rr = &v
	return s
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) SetPriority(v int32) *DescribeZoneRecordsResponseBodyRecordsRecord {
	s.Priority = &v
	return s
}

type DescribeZoneRecordsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeZoneRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeZoneRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeZoneRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeZoneRecordsResponse) SetHeaders(v map[string]*string) *DescribeZoneRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeZoneRecordsResponse) SetBody(v *DescribeZoneRecordsResponseBody) *DescribeZoneRecordsResponse {
	s.Body = v
	return s
}

type DescribeZonesRequest struct {
	Lang            *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageNumber      *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Keyword         *string   `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	UserClientIp    *string   `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	SearchMode      *string   `json:"SearchMode,omitempty" xml:"SearchMode,omitempty"`
	QueryRegionId   *string   `json:"QueryRegionId,omitempty" xml:"QueryRegionId,omitempty"`
	QueryVpcId      *string   `json:"QueryVpcId,omitempty" xml:"QueryVpcId,omitempty"`
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	OrderBy         *string   `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	Direction       *string   `json:"Direction,omitempty" xml:"Direction,omitempty"`
	ZoneType        *string   `json:"ZoneType,omitempty" xml:"ZoneType,omitempty"`
	ZoneTag         []*string `json:"ZoneTag,omitempty" xml:"ZoneTag,omitempty" type:"Repeated"`
}

func (s DescribeZonesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeZonesRequest) SetLang(v string) *DescribeZonesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeZonesRequest) SetPageNumber(v int32) *DescribeZonesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeZonesRequest) SetPageSize(v int32) *DescribeZonesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeZonesRequest) SetKeyword(v string) *DescribeZonesRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeZonesRequest) SetUserClientIp(v string) *DescribeZonesRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeZonesRequest) SetSearchMode(v string) *DescribeZonesRequest {
	s.SearchMode = &v
	return s
}

func (s *DescribeZonesRequest) SetQueryRegionId(v string) *DescribeZonesRequest {
	s.QueryRegionId = &v
	return s
}

func (s *DescribeZonesRequest) SetQueryVpcId(v string) *DescribeZonesRequest {
	s.QueryVpcId = &v
	return s
}

func (s *DescribeZonesRequest) SetResourceGroupId(v string) *DescribeZonesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeZonesRequest) SetOrderBy(v string) *DescribeZonesRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeZonesRequest) SetDirection(v string) *DescribeZonesRequest {
	s.Direction = &v
	return s
}

func (s *DescribeZonesRequest) SetZoneType(v string) *DescribeZonesRequest {
	s.ZoneType = &v
	return s
}

func (s *DescribeZonesRequest) SetZoneTag(v []*string) *DescribeZonesRequest {
	s.ZoneTag = v
	return s
}

type DescribeZonesResponseBody struct {
	PageSize   *int32                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalPages *int32                          `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	TotalItems *int32                          `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	Zones      *DescribeZonesResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
}

func (s DescribeZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBody) SetPageSize(v int32) *DescribeZonesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeZonesResponseBody) SetRequestId(v string) *DescribeZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeZonesResponseBody) SetPageNumber(v int32) *DescribeZonesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeZonesResponseBody) SetTotalPages(v int32) *DescribeZonesResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeZonesResponseBody) SetTotalItems(v int32) *DescribeZonesResponseBody {
	s.TotalItems = &v
	return s
}

func (s *DescribeZonesResponseBody) SetZones(v *DescribeZonesResponseBodyZones) *DescribeZonesResponseBody {
	s.Zones = v
	return s
}

type DescribeZonesResponseBodyZones struct {
	Zone []*DescribeZonesResponseBodyZonesZone `json:"Zone,omitempty" xml:"Zone,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBodyZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZones) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZones) SetZone(v []*DescribeZonesResponseBodyZonesZone) *DescribeZonesResponseBodyZones {
	s.Zone = v
	return s
}

type DescribeZonesResponseBodyZonesZone struct {
	UpdateTime      *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	ZoneType        *string `json:"ZoneType,omitempty" xml:"ZoneType,omitempty"`
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	RecordCount     *int32  `json:"RecordCount,omitempty" xml:"RecordCount,omitempty"`
	ZoneName        *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
	ProxyPattern    *string `json:"ProxyPattern,omitempty" xml:"ProxyPattern,omitempty"`
	UpdateTimestamp *int64  `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ZoneId          *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ZoneTag         *string `json:"ZoneTag,omitempty" xml:"ZoneTag,omitempty"`
	IsPtr           *bool   `json:"IsPtr,omitempty" xml:"IsPtr,omitempty"`
	CreateTimestamp *int64  `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
}

func (s DescribeZonesResponseBodyZonesZone) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZone) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZone) SetUpdateTime(v string) *DescribeZonesResponseBodyZonesZone {
	s.UpdateTime = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetZoneType(v string) *DescribeZonesResponseBodyZonesZone {
	s.ZoneType = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetRemark(v string) *DescribeZonesResponseBodyZonesZone {
	s.Remark = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetCreateTime(v string) *DescribeZonesResponseBodyZonesZone {
	s.CreateTime = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetRecordCount(v int32) *DescribeZonesResponseBodyZonesZone {
	s.RecordCount = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetZoneName(v string) *DescribeZonesResponseBodyZonesZone {
	s.ZoneName = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetProxyPattern(v string) *DescribeZonesResponseBodyZonesZone {
	s.ProxyPattern = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetUpdateTimestamp(v int64) *DescribeZonesResponseBodyZonesZone {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetResourceGroupId(v string) *DescribeZonesResponseBodyZonesZone {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetZoneId(v string) *DescribeZonesResponseBodyZonesZone {
	s.ZoneId = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetZoneTag(v string) *DescribeZonesResponseBodyZonesZone {
	s.ZoneTag = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetIsPtr(v bool) *DescribeZonesResponseBodyZonesZone {
	s.IsPtr = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetCreateTimestamp(v int64) *DescribeZonesResponseBodyZonesZone {
	s.CreateTimestamp = &v
	return s
}

type DescribeZonesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeZonesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponse) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponse) SetHeaders(v map[string]*string) *DescribeZonesResponse {
	s.Headers = v
	return s
}

func (s *DescribeZonesResponse) SetBody(v *DescribeZonesResponseBody) *DescribeZonesResponse {
	s.Body = v
	return s
}

type DescribeZoneVpcTreeRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s DescribeZoneVpcTreeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeZoneVpcTreeRequest) GoString() string {
	return s.String()
}

func (s *DescribeZoneVpcTreeRequest) SetLang(v string) *DescribeZoneVpcTreeRequest {
	s.Lang = &v
	return s
}

func (s *DescribeZoneVpcTreeRequest) SetUserClientIp(v string) *DescribeZoneVpcTreeRequest {
	s.UserClientIp = &v
	return s
}

type DescribeZoneVpcTreeResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Zones     *DescribeZoneVpcTreeResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
}

func (s DescribeZoneVpcTreeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeZoneVpcTreeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeZoneVpcTreeResponseBody) SetRequestId(v string) *DescribeZoneVpcTreeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBody) SetZones(v *DescribeZoneVpcTreeResponseBodyZones) *DescribeZoneVpcTreeResponseBody {
	s.Zones = v
	return s
}

type DescribeZoneVpcTreeResponseBodyZones struct {
	Zone []*DescribeZoneVpcTreeResponseBodyZonesZone `json:"Zone,omitempty" xml:"Zone,omitempty" type:"Repeated"`
}

func (s DescribeZoneVpcTreeResponseBodyZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeZoneVpcTreeResponseBodyZones) GoString() string {
	return s.String()
}

func (s *DescribeZoneVpcTreeResponseBodyZones) SetZone(v []*DescribeZoneVpcTreeResponseBodyZonesZone) *DescribeZoneVpcTreeResponseBodyZones {
	s.Zone = v
	return s
}

type DescribeZoneVpcTreeResponseBodyZonesZone struct {
	UpdateTime      *string                                       `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	ZoneType        *string                                       `json:"ZoneType,omitempty" xml:"ZoneType,omitempty"`
	Remark          *string                                       `json:"Remark,omitempty" xml:"Remark,omitempty"`
	CreateTime      *string                                       `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Vpcs            *DescribeZoneVpcTreeResponseBodyZonesZoneVpcs `json:"Vpcs,omitempty" xml:"Vpcs,omitempty" type:"Struct"`
	RecordCount     *int32                                        `json:"RecordCount,omitempty" xml:"RecordCount,omitempty"`
	ZoneName        *string                                       `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
	UpdateTimestamp *int64                                        `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	ZoneId          *string                                       `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ZoneTag         *string                                       `json:"ZoneTag,omitempty" xml:"ZoneTag,omitempty"`
	IsPtr           *bool                                         `json:"IsPtr,omitempty" xml:"IsPtr,omitempty"`
	CreateTimestamp *int64                                        `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
}

func (s DescribeZoneVpcTreeResponseBodyZonesZone) String() string {
	return tea.Prettify(s)
}

func (s DescribeZoneVpcTreeResponseBodyZonesZone) GoString() string {
	return s.String()
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetUpdateTime(v string) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.UpdateTime = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetZoneType(v string) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.ZoneType = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetRemark(v string) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.Remark = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetCreateTime(v string) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.CreateTime = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetVpcs(v *DescribeZoneVpcTreeResponseBodyZonesZoneVpcs) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.Vpcs = v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetRecordCount(v int32) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.RecordCount = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetZoneName(v string) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.ZoneName = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetUpdateTimestamp(v int64) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetZoneId(v string) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.ZoneId = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetZoneTag(v string) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.ZoneTag = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetIsPtr(v bool) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.IsPtr = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetCreateTimestamp(v int64) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.CreateTimestamp = &v
	return s
}

type DescribeZoneVpcTreeResponseBodyZonesZoneVpcs struct {
	Vpc []*DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc `json:"Vpc,omitempty" xml:"Vpc,omitempty" type:"Repeated"`
}

func (s DescribeZoneVpcTreeResponseBodyZonesZoneVpcs) String() string {
	return tea.Prettify(s)
}

func (s DescribeZoneVpcTreeResponseBodyZonesZoneVpcs) GoString() string {
	return s.String()
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZoneVpcs) SetVpc(v []*DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc) *DescribeZoneVpcTreeResponseBodyZonesZoneVpcs {
	s.Vpc = v
	return s
}

type DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc struct {
	VpcName    *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
	VpcId      *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc) String() string {
	return tea.Prettify(s)
}

func (s DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc) GoString() string {
	return s.String()
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc) SetVpcName(v string) *DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc {
	s.VpcName = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc) SetVpcId(v string) *DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc {
	s.VpcId = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc) SetRegionName(v string) *DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc {
	s.RegionName = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc) SetRegionId(v string) *DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc {
	s.RegionId = &v
	return s
}

type DescribeZoneVpcTreeResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeZoneVpcTreeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeZoneVpcTreeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeZoneVpcTreeResponse) GoString() string {
	return s.String()
}

func (s *DescribeZoneVpcTreeResponse) SetHeaders(v map[string]*string) *DescribeZoneVpcTreeResponse {
	s.Headers = v
	return s
}

func (s *DescribeZoneVpcTreeResponse) SetBody(v *DescribeZoneVpcTreeResponseBody) *DescribeZoneVpcTreeResponse {
	s.Body = v
	return s
}

type MoveResourceGroupRequest struct {
	Lang               *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceId         *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	UserClientIp       *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
}

func (s MoveResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupRequest) SetLang(v string) *MoveResourceGroupRequest {
	s.Lang = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceId(v string) *MoveResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetUserClientIp(v string) *MoveResourceGroupRequest {
	s.UserClientIp = &v
	return s
}

func (s *MoveResourceGroupRequest) SetNewResourceGroupId(v string) *MoveResourceGroupRequest {
	s.NewResourceGroupId = &v
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

type SetProxyPatternRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ZoneId       *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ProxyPattern *string `json:"ProxyPattern,omitempty" xml:"ProxyPattern,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SetProxyPatternRequest) String() string {
	return tea.Prettify(s)
}

func (s SetProxyPatternRequest) GoString() string {
	return s.String()
}

func (s *SetProxyPatternRequest) SetLang(v string) *SetProxyPatternRequest {
	s.Lang = &v
	return s
}

func (s *SetProxyPatternRequest) SetZoneId(v string) *SetProxyPatternRequest {
	s.ZoneId = &v
	return s
}

func (s *SetProxyPatternRequest) SetProxyPattern(v string) *SetProxyPatternRequest {
	s.ProxyPattern = &v
	return s
}

func (s *SetProxyPatternRequest) SetUserClientIp(v string) *SetProxyPatternRequest {
	s.UserClientIp = &v
	return s
}

type SetProxyPatternResponseBody struct {
	ZoneId    *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetProxyPatternResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetProxyPatternResponseBody) GoString() string {
	return s.String()
}

func (s *SetProxyPatternResponseBody) SetZoneId(v string) *SetProxyPatternResponseBody {
	s.ZoneId = &v
	return s
}

func (s *SetProxyPatternResponseBody) SetRequestId(v string) *SetProxyPatternResponseBody {
	s.RequestId = &v
	return s
}

type SetProxyPatternResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetProxyPatternResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetProxyPatternResponse) String() string {
	return tea.Prettify(s)
}

func (s SetProxyPatternResponse) GoString() string {
	return s.String()
}

func (s *SetProxyPatternResponse) SetHeaders(v map[string]*string) *SetProxyPatternResponse {
	s.Headers = v
	return s
}

func (s *SetProxyPatternResponse) SetBody(v *SetProxyPatternResponseBody) *SetProxyPatternResponse {
	s.Body = v
	return s
}

type SetZoneRecordStatusRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	RecordId     *int64  `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SetZoneRecordStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s SetZoneRecordStatusRequest) GoString() string {
	return s.String()
}

func (s *SetZoneRecordStatusRequest) SetLang(v string) *SetZoneRecordStatusRequest {
	s.Lang = &v
	return s
}

func (s *SetZoneRecordStatusRequest) SetRecordId(v int64) *SetZoneRecordStatusRequest {
	s.RecordId = &v
	return s
}

func (s *SetZoneRecordStatusRequest) SetStatus(v string) *SetZoneRecordStatusRequest {
	s.Status = &v
	return s
}

func (s *SetZoneRecordStatusRequest) SetUserClientIp(v string) *SetZoneRecordStatusRequest {
	s.UserClientIp = &v
	return s
}

type SetZoneRecordStatusResponseBody struct {
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RecordId  *int64  `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
}

func (s SetZoneRecordStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetZoneRecordStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetZoneRecordStatusResponseBody) SetStatus(v string) *SetZoneRecordStatusResponseBody {
	s.Status = &v
	return s
}

func (s *SetZoneRecordStatusResponseBody) SetRequestId(v string) *SetZoneRecordStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetZoneRecordStatusResponseBody) SetRecordId(v int64) *SetZoneRecordStatusResponseBody {
	s.RecordId = &v
	return s
}

type SetZoneRecordStatusResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetZoneRecordStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetZoneRecordStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s SetZoneRecordStatusResponse) GoString() string {
	return s.String()
}

func (s *SetZoneRecordStatusResponse) SetHeaders(v map[string]*string) *SetZoneRecordStatusResponse {
	s.Headers = v
	return s
}

func (s *SetZoneRecordStatusResponse) SetBody(v *SetZoneRecordStatusResponseBody) *SetZoneRecordStatusResponse {
	s.Body = v
	return s
}

type UpdateRecordRemarkRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	RecordId     *int64  `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	Remark       *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s UpdateRecordRemarkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRecordRemarkRequest) GoString() string {
	return s.String()
}

func (s *UpdateRecordRemarkRequest) SetLang(v string) *UpdateRecordRemarkRequest {
	s.Lang = &v
	return s
}

func (s *UpdateRecordRemarkRequest) SetRecordId(v int64) *UpdateRecordRemarkRequest {
	s.RecordId = &v
	return s
}

func (s *UpdateRecordRemarkRequest) SetRemark(v string) *UpdateRecordRemarkRequest {
	s.Remark = &v
	return s
}

func (s *UpdateRecordRemarkRequest) SetUserClientIp(v string) *UpdateRecordRemarkRequest {
	s.UserClientIp = &v
	return s
}

type UpdateRecordRemarkResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RecordId  *int64  `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
}

func (s UpdateRecordRemarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRecordRemarkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRecordRemarkResponseBody) SetRequestId(v string) *UpdateRecordRemarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRecordRemarkResponseBody) SetRecordId(v int64) *UpdateRecordRemarkResponseBody {
	s.RecordId = &v
	return s
}

type UpdateRecordRemarkResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateRecordRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateRecordRemarkResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRecordRemarkResponse) GoString() string {
	return s.String()
}

func (s *UpdateRecordRemarkResponse) SetHeaders(v map[string]*string) *UpdateRecordRemarkResponse {
	s.Headers = v
	return s
}

func (s *UpdateRecordRemarkResponse) SetBody(v *UpdateRecordRemarkResponseBody) *UpdateRecordRemarkResponse {
	s.Body = v
	return s
}

type UpdateZoneRecordRequest struct {
	Rr           *string `json:"Rr,omitempty" xml:"Rr,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	RecordId     *int64  `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Ttl          *int32  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	Priority     *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Value        *string `json:"Value,omitempty" xml:"Value,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s UpdateZoneRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateZoneRecordRequest) GoString() string {
	return s.String()
}

func (s *UpdateZoneRecordRequest) SetRr(v string) *UpdateZoneRecordRequest {
	s.Rr = &v
	return s
}

func (s *UpdateZoneRecordRequest) SetLang(v string) *UpdateZoneRecordRequest {
	s.Lang = &v
	return s
}

func (s *UpdateZoneRecordRequest) SetRecordId(v int64) *UpdateZoneRecordRequest {
	s.RecordId = &v
	return s
}

func (s *UpdateZoneRecordRequest) SetType(v string) *UpdateZoneRecordRequest {
	s.Type = &v
	return s
}

func (s *UpdateZoneRecordRequest) SetTtl(v int32) *UpdateZoneRecordRequest {
	s.Ttl = &v
	return s
}

func (s *UpdateZoneRecordRequest) SetPriority(v int32) *UpdateZoneRecordRequest {
	s.Priority = &v
	return s
}

func (s *UpdateZoneRecordRequest) SetValue(v string) *UpdateZoneRecordRequest {
	s.Value = &v
	return s
}

func (s *UpdateZoneRecordRequest) SetUserClientIp(v string) *UpdateZoneRecordRequest {
	s.UserClientIp = &v
	return s
}

type UpdateZoneRecordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RecordId  *int64  `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
}

func (s UpdateZoneRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateZoneRecordResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateZoneRecordResponseBody) SetRequestId(v string) *UpdateZoneRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateZoneRecordResponseBody) SetRecordId(v int64) *UpdateZoneRecordResponseBody {
	s.RecordId = &v
	return s
}

type UpdateZoneRecordResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateZoneRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateZoneRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateZoneRecordResponse) GoString() string {
	return s.String()
}

func (s *UpdateZoneRecordResponse) SetHeaders(v map[string]*string) *UpdateZoneRecordResponse {
	s.Headers = v
	return s
}

func (s *UpdateZoneRecordResponse) SetBody(v *UpdateZoneRecordResponseBody) *UpdateZoneRecordResponse {
	s.Body = v
	return s
}

type UpdateZoneRemarkRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ZoneId       *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	Remark       *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s UpdateZoneRemarkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateZoneRemarkRequest) GoString() string {
	return s.String()
}

func (s *UpdateZoneRemarkRequest) SetLang(v string) *UpdateZoneRemarkRequest {
	s.Lang = &v
	return s
}

func (s *UpdateZoneRemarkRequest) SetZoneId(v string) *UpdateZoneRemarkRequest {
	s.ZoneId = &v
	return s
}

func (s *UpdateZoneRemarkRequest) SetRemark(v string) *UpdateZoneRemarkRequest {
	s.Remark = &v
	return s
}

func (s *UpdateZoneRemarkRequest) SetUserClientIp(v string) *UpdateZoneRemarkRequest {
	s.UserClientIp = &v
	return s
}

type UpdateZoneRemarkResponseBody struct {
	ZoneId    *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateZoneRemarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateZoneRemarkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateZoneRemarkResponseBody) SetZoneId(v string) *UpdateZoneRemarkResponseBody {
	s.ZoneId = &v
	return s
}

func (s *UpdateZoneRemarkResponseBody) SetRequestId(v string) *UpdateZoneRemarkResponseBody {
	s.RequestId = &v
	return s
}

type UpdateZoneRemarkResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateZoneRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateZoneRemarkResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateZoneRemarkResponse) GoString() string {
	return s.String()
}

func (s *UpdateZoneRemarkResponse) SetHeaders(v map[string]*string) *UpdateZoneRemarkResponse {
	s.Headers = v
	return s
}

func (s *UpdateZoneRemarkResponse) SetBody(v *UpdateZoneRemarkResponseBody) *UpdateZoneRemarkResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("pvtz"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddZoneWithOptions(request *AddZoneRequest, runtime *util.RuntimeOptions) (_result *AddZoneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddZoneResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddZone"), tea.String("2018-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddZone(request *AddZoneRequest) (_result *AddZoneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddZoneResponse{}
	_body, _err := client.AddZoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddZoneRecordWithOptions(request *AddZoneRecordRequest, runtime *util.RuntimeOptions) (_result *AddZoneRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddZoneRecordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddZoneRecord"), tea.String("2018-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddZoneRecord(request *AddZoneRecordRequest) (_result *AddZoneRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddZoneRecordResponse{}
	_body, _err := client.AddZoneRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindZoneVpcWithOptions(request *BindZoneVpcRequest, runtime *util.RuntimeOptions) (_result *BindZoneVpcResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BindZoneVpcResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BindZoneVpc"), tea.String("2018-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindZoneVpc(request *BindZoneVpcRequest) (_result *BindZoneVpcResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindZoneVpcResponse{}
	_body, _err := client.BindZoneVpcWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckZoneNameWithOptions(request *CheckZoneNameRequest, runtime *util.RuntimeOptions) (_result *CheckZoneNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CheckZoneNameResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CheckZoneName"), tea.String("2018-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckZoneName(request *CheckZoneNameRequest) (_result *CheckZoneNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckZoneNameResponse{}
	_body, _err := client.CheckZoneNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteZoneWithOptions(request *DeleteZoneRequest, runtime *util.RuntimeOptions) (_result *DeleteZoneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteZoneResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteZone"), tea.String("2018-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteZone(request *DeleteZoneRequest) (_result *DeleteZoneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteZoneResponse{}
	_body, _err := client.DeleteZoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteZoneRecordWithOptions(request *DeleteZoneRecordRequest, runtime *util.RuntimeOptions) (_result *DeleteZoneRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteZoneRecordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteZoneRecord"), tea.String("2018-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteZoneRecord(request *DeleteZoneRecordRequest) (_result *DeleteZoneRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteZoneRecordResponse{}
	_body, _err := client.DeleteZoneRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeChangeLogsWithOptions(request *DescribeChangeLogsRequest, runtime *util.RuntimeOptions) (_result *DescribeChangeLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeChangeLogsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeChangeLogs"), tea.String("2018-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeChangeLogs(request *DescribeChangeLogsRequest) (_result *DescribeChangeLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeChangeLogsResponse{}
	_body, _err := client.DescribeChangeLogsWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DescribeRegions"), tea.String("2018-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeRequestGraphWithOptions(request *DescribeRequestGraphRequest, runtime *util.RuntimeOptions) (_result *DescribeRequestGraphResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRequestGraphResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRequestGraph"), tea.String("2018-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRequestGraph(request *DescribeRequestGraphRequest) (_result *DescribeRequestGraphResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRequestGraphResponse{}
	_body, _err := client.DescribeRequestGraphWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeStatisticSummaryWithOptions(request *DescribeStatisticSummaryRequest, runtime *util.RuntimeOptions) (_result *DescribeStatisticSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeStatisticSummaryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeStatisticSummary"), tea.String("2018-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeStatisticSummary(request *DescribeStatisticSummaryRequest) (_result *DescribeStatisticSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeStatisticSummaryResponse{}
	_body, _err := client.DescribeStatisticSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserServiceStatusWithOptions(request *DescribeUserServiceStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeUserServiceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeUserServiceStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeUserServiceStatus"), tea.String("2018-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserServiceStatus(request *DescribeUserServiceStatusRequest) (_result *DescribeUserServiceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserServiceStatusResponse{}
	_body, _err := client.DescribeUserServiceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeZoneInfoWithOptions(request *DescribeZoneInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeZoneInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeZoneInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeZoneInfo"), tea.String("2018-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeZoneInfo(request *DescribeZoneInfoRequest) (_result *DescribeZoneInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeZoneInfoResponse{}
	_body, _err := client.DescribeZoneInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeZoneRecordsWithOptions(request *DescribeZoneRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeZoneRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeZoneRecordsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeZoneRecords"), tea.String("2018-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeZoneRecords(request *DescribeZoneRecordsRequest) (_result *DescribeZoneRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeZoneRecordsResponse{}
	_body, _err := client.DescribeZoneRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeZonesWithOptions(request *DescribeZonesRequest, runtime *util.RuntimeOptions) (_result *DescribeZonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeZonesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeZones"), tea.String("2018-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeZones(request *DescribeZonesRequest) (_result *DescribeZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeZonesResponse{}
	_body, _err := client.DescribeZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeZoneVpcTreeWithOptions(request *DescribeZoneVpcTreeRequest, runtime *util.RuntimeOptions) (_result *DescribeZoneVpcTreeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeZoneVpcTreeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeZoneVpcTree"), tea.String("2018-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeZoneVpcTree(request *DescribeZoneVpcTreeRequest) (_result *DescribeZoneVpcTreeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeZoneVpcTreeResponse{}
	_body, _err := client.DescribeZoneVpcTreeWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("MoveResourceGroup"), tea.String("2018-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) SetProxyPatternWithOptions(request *SetProxyPatternRequest, runtime *util.RuntimeOptions) (_result *SetProxyPatternResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetProxyPatternResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetProxyPattern"), tea.String("2018-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetProxyPattern(request *SetProxyPatternRequest) (_result *SetProxyPatternResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetProxyPatternResponse{}
	_body, _err := client.SetProxyPatternWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetZoneRecordStatusWithOptions(request *SetZoneRecordStatusRequest, runtime *util.RuntimeOptions) (_result *SetZoneRecordStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetZoneRecordStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetZoneRecordStatus"), tea.String("2018-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetZoneRecordStatus(request *SetZoneRecordStatusRequest) (_result *SetZoneRecordStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetZoneRecordStatusResponse{}
	_body, _err := client.SetZoneRecordStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateRecordRemarkWithOptions(request *UpdateRecordRemarkRequest, runtime *util.RuntimeOptions) (_result *UpdateRecordRemarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateRecordRemarkResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateRecordRemark"), tea.String("2018-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateRecordRemark(request *UpdateRecordRemarkRequest) (_result *UpdateRecordRemarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateRecordRemarkResponse{}
	_body, _err := client.UpdateRecordRemarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateZoneRecordWithOptions(request *UpdateZoneRecordRequest, runtime *util.RuntimeOptions) (_result *UpdateZoneRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateZoneRecordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateZoneRecord"), tea.String("2018-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateZoneRecord(request *UpdateZoneRecordRequest) (_result *UpdateZoneRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateZoneRecordResponse{}
	_body, _err := client.UpdateZoneRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateZoneRemarkWithOptions(request *UpdateZoneRemarkRequest, runtime *util.RuntimeOptions) (_result *UpdateZoneRemarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateZoneRemarkResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateZoneRemark"), tea.String("2018-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateZoneRemark(request *UpdateZoneRemarkRequest) (_result *UpdateZoneRemarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateZoneRemarkResponse{}
	_body, _err := client.UpdateZoneRemarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
