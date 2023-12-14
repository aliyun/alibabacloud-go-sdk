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

type ARMSQueryDataSetRequest struct {
	DatasetId     *int64                                 `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	DateStr       *string                                `json:"DateStr,omitempty" xml:"DateStr,omitempty"`
	Dimensions    []*ARMSQueryDataSetRequestDimensions   `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	HungryMode    *bool                                  `json:"HungryMode,omitempty" xml:"HungryMode,omitempty"`
	IntervalInSec *int32                                 `json:"IntervalInSec,omitempty" xml:"IntervalInSec,omitempty"`
	IsDrillDown   *bool                                  `json:"IsDrillDown,omitempty" xml:"IsDrillDown,omitempty"`
	Limit         *int32                                 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	MaxTime       *int64                                 `json:"MaxTime,omitempty" xml:"MaxTime,omitempty"`
	Measures      []*string                              `json:"Measures,omitempty" xml:"Measures,omitempty" type:"Repeated"`
	MinTime       *int64                                 `json:"MinTime,omitempty" xml:"MinTime,omitempty"`
	OptionalDims  []*ARMSQueryDataSetRequestOptionalDims `json:"OptionalDims,omitempty" xml:"OptionalDims,omitempty" type:"Repeated"`
	OrderByKey    *string                                `json:"OrderByKey,omitempty" xml:"OrderByKey,omitempty"`
	ReduceTail    *bool                                  `json:"ReduceTail,omitempty" xml:"ReduceTail,omitempty"`
	RequiredDims  []*ARMSQueryDataSetRequestRequiredDims `json:"RequiredDims,omitempty" xml:"RequiredDims,omitempty" type:"Repeated"`
	SecurityToken *string                                `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ARMSQueryDataSetRequest) String() string {
	return tea.Prettify(s)
}

func (s ARMSQueryDataSetRequest) GoString() string {
	return s.String()
}

func (s *ARMSQueryDataSetRequest) SetDatasetId(v int64) *ARMSQueryDataSetRequest {
	s.DatasetId = &v
	return s
}

func (s *ARMSQueryDataSetRequest) SetDateStr(v string) *ARMSQueryDataSetRequest {
	s.DateStr = &v
	return s
}

func (s *ARMSQueryDataSetRequest) SetDimensions(v []*ARMSQueryDataSetRequestDimensions) *ARMSQueryDataSetRequest {
	s.Dimensions = v
	return s
}

func (s *ARMSQueryDataSetRequest) SetHungryMode(v bool) *ARMSQueryDataSetRequest {
	s.HungryMode = &v
	return s
}

func (s *ARMSQueryDataSetRequest) SetIntervalInSec(v int32) *ARMSQueryDataSetRequest {
	s.IntervalInSec = &v
	return s
}

func (s *ARMSQueryDataSetRequest) SetIsDrillDown(v bool) *ARMSQueryDataSetRequest {
	s.IsDrillDown = &v
	return s
}

func (s *ARMSQueryDataSetRequest) SetLimit(v int32) *ARMSQueryDataSetRequest {
	s.Limit = &v
	return s
}

func (s *ARMSQueryDataSetRequest) SetMaxTime(v int64) *ARMSQueryDataSetRequest {
	s.MaxTime = &v
	return s
}

func (s *ARMSQueryDataSetRequest) SetMeasures(v []*string) *ARMSQueryDataSetRequest {
	s.Measures = v
	return s
}

func (s *ARMSQueryDataSetRequest) SetMinTime(v int64) *ARMSQueryDataSetRequest {
	s.MinTime = &v
	return s
}

func (s *ARMSQueryDataSetRequest) SetOptionalDims(v []*ARMSQueryDataSetRequestOptionalDims) *ARMSQueryDataSetRequest {
	s.OptionalDims = v
	return s
}

func (s *ARMSQueryDataSetRequest) SetOrderByKey(v string) *ARMSQueryDataSetRequest {
	s.OrderByKey = &v
	return s
}

func (s *ARMSQueryDataSetRequest) SetReduceTail(v bool) *ARMSQueryDataSetRequest {
	s.ReduceTail = &v
	return s
}

func (s *ARMSQueryDataSetRequest) SetRequiredDims(v []*ARMSQueryDataSetRequestRequiredDims) *ARMSQueryDataSetRequest {
	s.RequiredDims = v
	return s
}

func (s *ARMSQueryDataSetRequest) SetSecurityToken(v string) *ARMSQueryDataSetRequest {
	s.SecurityToken = &v
	return s
}

type ARMSQueryDataSetRequestDimensions struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ARMSQueryDataSetRequestDimensions) String() string {
	return tea.Prettify(s)
}

func (s ARMSQueryDataSetRequestDimensions) GoString() string {
	return s.String()
}

func (s *ARMSQueryDataSetRequestDimensions) SetKey(v string) *ARMSQueryDataSetRequestDimensions {
	s.Key = &v
	return s
}

func (s *ARMSQueryDataSetRequestDimensions) SetType(v string) *ARMSQueryDataSetRequestDimensions {
	s.Type = &v
	return s
}

func (s *ARMSQueryDataSetRequestDimensions) SetValue(v string) *ARMSQueryDataSetRequestDimensions {
	s.Value = &v
	return s
}

type ARMSQueryDataSetRequestOptionalDims struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ARMSQueryDataSetRequestOptionalDims) String() string {
	return tea.Prettify(s)
}

func (s ARMSQueryDataSetRequestOptionalDims) GoString() string {
	return s.String()
}

func (s *ARMSQueryDataSetRequestOptionalDims) SetKey(v string) *ARMSQueryDataSetRequestOptionalDims {
	s.Key = &v
	return s
}

func (s *ARMSQueryDataSetRequestOptionalDims) SetType(v string) *ARMSQueryDataSetRequestOptionalDims {
	s.Type = &v
	return s
}

func (s *ARMSQueryDataSetRequestOptionalDims) SetValue(v string) *ARMSQueryDataSetRequestOptionalDims {
	s.Value = &v
	return s
}

type ARMSQueryDataSetRequestRequiredDims struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ARMSQueryDataSetRequestRequiredDims) String() string {
	return tea.Prettify(s)
}

func (s ARMSQueryDataSetRequestRequiredDims) GoString() string {
	return s.String()
}

func (s *ARMSQueryDataSetRequestRequiredDims) SetKey(v string) *ARMSQueryDataSetRequestRequiredDims {
	s.Key = &v
	return s
}

func (s *ARMSQueryDataSetRequestRequiredDims) SetType(v string) *ARMSQueryDataSetRequestRequiredDims {
	s.Type = &v
	return s
}

func (s *ARMSQueryDataSetRequestRequiredDims) SetValue(v string) *ARMSQueryDataSetRequestRequiredDims {
	s.Value = &v
	return s
}

type ARMSQueryDataSetResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ARMSQueryDataSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ARMSQueryDataSetResponseBody) GoString() string {
	return s.String()
}

func (s *ARMSQueryDataSetResponseBody) SetData(v string) *ARMSQueryDataSetResponseBody {
	s.Data = &v
	return s
}

func (s *ARMSQueryDataSetResponseBody) SetRequestId(v string) *ARMSQueryDataSetResponseBody {
	s.RequestId = &v
	return s
}

type ARMSQueryDataSetResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ARMSQueryDataSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ARMSQueryDataSetResponse) String() string {
	return tea.Prettify(s)
}

func (s ARMSQueryDataSetResponse) GoString() string {
	return s.String()
}

func (s *ARMSQueryDataSetResponse) SetHeaders(v map[string]*string) *ARMSQueryDataSetResponse {
	s.Headers = v
	return s
}

func (s *ARMSQueryDataSetResponse) SetStatusCode(v int32) *ARMSQueryDataSetResponse {
	s.StatusCode = &v
	return s
}

func (s *ARMSQueryDataSetResponse) SetBody(v *ARMSQueryDataSetResponseBody) *ARMSQueryDataSetResponse {
	s.Body = v
	return s
}

type CreateAlertContactRequest struct {
	ContactName         *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	DingRobotWebhookUrl *string `json:"DingRobotWebhookUrl,omitempty" xml:"DingRobotWebhookUrl,omitempty"`
	Email               *string `json:"Email,omitempty" xml:"Email,omitempty"`
	PhoneNum            *string `json:"PhoneNum,omitempty" xml:"PhoneNum,omitempty"`
	SystemNoc           *bool   `json:"SystemNoc,omitempty" xml:"SystemNoc,omitempty"`
}

func (s CreateAlertContactRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAlertContactRequest) GoString() string {
	return s.String()
}

func (s *CreateAlertContactRequest) SetContactName(v string) *CreateAlertContactRequest {
	s.ContactName = &v
	return s
}

func (s *CreateAlertContactRequest) SetDingRobotWebhookUrl(v string) *CreateAlertContactRequest {
	s.DingRobotWebhookUrl = &v
	return s
}

func (s *CreateAlertContactRequest) SetEmail(v string) *CreateAlertContactRequest {
	s.Email = &v
	return s
}

func (s *CreateAlertContactRequest) SetPhoneNum(v string) *CreateAlertContactRequest {
	s.PhoneNum = &v
	return s
}

func (s *CreateAlertContactRequest) SetSystemNoc(v bool) *CreateAlertContactRequest {
	s.SystemNoc = &v
	return s
}

type CreateAlertContactResponseBody struct {
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAlertContactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAlertContactResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAlertContactResponseBody) SetContactId(v string) *CreateAlertContactResponseBody {
	s.ContactId = &v
	return s
}

func (s *CreateAlertContactResponseBody) SetData(v string) *CreateAlertContactResponseBody {
	s.Data = &v
	return s
}

func (s *CreateAlertContactResponseBody) SetRequestId(v string) *CreateAlertContactResponseBody {
	s.RequestId = &v
	return s
}

type CreateAlertContactResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAlertContactResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAlertContactResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAlertContactResponse) GoString() string {
	return s.String()
}

func (s *CreateAlertContactResponse) SetHeaders(v map[string]*string) *CreateAlertContactResponse {
	s.Headers = v
	return s
}

func (s *CreateAlertContactResponse) SetStatusCode(v int32) *CreateAlertContactResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAlertContactResponse) SetBody(v *CreateAlertContactResponseBody) *CreateAlertContactResponse {
	s.Body = v
	return s
}

type CreateAlertContactGroupRequest struct {
	ContactGroupName *string `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
	ContactIds       *string `json:"ContactIds,omitempty" xml:"ContactIds,omitempty"`
}

func (s CreateAlertContactGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAlertContactGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateAlertContactGroupRequest) SetContactGroupName(v string) *CreateAlertContactGroupRequest {
	s.ContactGroupName = &v
	return s
}

func (s *CreateAlertContactGroupRequest) SetContactIds(v string) *CreateAlertContactGroupRequest {
	s.ContactIds = &v
	return s
}

type CreateAlertContactGroupResponseBody struct {
	ContactGroupId *string `json:"ContactGroupId,omitempty" xml:"ContactGroupId,omitempty"`
	Data           *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAlertContactGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAlertContactGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAlertContactGroupResponseBody) SetContactGroupId(v string) *CreateAlertContactGroupResponseBody {
	s.ContactGroupId = &v
	return s
}

func (s *CreateAlertContactGroupResponseBody) SetData(v string) *CreateAlertContactGroupResponseBody {
	s.Data = &v
	return s
}

func (s *CreateAlertContactGroupResponseBody) SetRequestId(v string) *CreateAlertContactGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateAlertContactGroupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAlertContactGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAlertContactGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAlertContactGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateAlertContactGroupResponse) SetHeaders(v map[string]*string) *CreateAlertContactGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateAlertContactGroupResponse) SetStatusCode(v int32) *CreateAlertContactGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAlertContactGroupResponse) SetBody(v *CreateAlertContactGroupResponseBody) *CreateAlertContactGroupResponse {
	s.Body = v
	return s
}

type GetIpOrHostsRequest struct {
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetIpOrHostsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetIpOrHostsRequest) GoString() string {
	return s.String()
}

func (s *GetIpOrHostsRequest) SetEndTime(v int64) *GetIpOrHostsRequest {
	s.EndTime = &v
	return s
}

func (s *GetIpOrHostsRequest) SetRegionId(v string) *GetIpOrHostsRequest {
	s.RegionId = &v
	return s
}

func (s *GetIpOrHostsRequest) SetServiceName(v string) *GetIpOrHostsRequest {
	s.ServiceName = &v
	return s
}

func (s *GetIpOrHostsRequest) SetStartTime(v int64) *GetIpOrHostsRequest {
	s.StartTime = &v
	return s
}

type GetIpOrHostsResponseBody struct {
	Data      *GetIpOrHostsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetIpOrHostsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIpOrHostsResponseBody) GoString() string {
	return s.String()
}

func (s *GetIpOrHostsResponseBody) SetData(v *GetIpOrHostsResponseBodyData) *GetIpOrHostsResponseBody {
	s.Data = v
	return s
}

func (s *GetIpOrHostsResponseBody) SetRequestId(v string) *GetIpOrHostsResponseBody {
	s.RequestId = &v
	return s
}

type GetIpOrHostsResponseBodyData struct {
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s GetIpOrHostsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetIpOrHostsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetIpOrHostsResponseBodyData) SetData(v []*string) *GetIpOrHostsResponseBodyData {
	s.Data = v
	return s
}

type GetIpOrHostsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetIpOrHostsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetIpOrHostsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIpOrHostsResponse) GoString() string {
	return s.String()
}

func (s *GetIpOrHostsResponse) SetHeaders(v map[string]*string) *GetIpOrHostsResponse {
	s.Headers = v
	return s
}

func (s *GetIpOrHostsResponse) SetStatusCode(v int32) *GetIpOrHostsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIpOrHostsResponse) SetBody(v *GetIpOrHostsResponseBody) *GetIpOrHostsResponse {
	s.Body = v
	return s
}

type GetServicesRequest struct {
	AppType  *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetServicesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServicesRequest) GoString() string {
	return s.String()
}

func (s *GetServicesRequest) SetAppType(v string) *GetServicesRequest {
	s.AppType = &v
	return s
}

func (s *GetServicesRequest) SetRegionId(v string) *GetServicesRequest {
	s.RegionId = &v
	return s
}

type GetServicesResponseBody struct {
	Data      *GetServicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetServicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServicesResponseBody) GoString() string {
	return s.String()
}

func (s *GetServicesResponseBody) SetData(v *GetServicesResponseBodyData) *GetServicesResponseBody {
	s.Data = v
	return s
}

func (s *GetServicesResponseBody) SetRequestId(v string) *GetServicesResponseBody {
	s.RequestId = &v
	return s
}

type GetServicesResponseBodyData struct {
	Details  *GetServicesResponseBodyDataDetails  `json:"Details,omitempty" xml:"Details,omitempty" type:"Struct"`
	Services *GetServicesResponseBodyDataServices `json:"Services,omitempty" xml:"Services,omitempty" type:"Struct"`
}

func (s GetServicesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetServicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetServicesResponseBodyData) SetDetails(v *GetServicesResponseBodyDataDetails) *GetServicesResponseBodyData {
	s.Details = v
	return s
}

func (s *GetServicesResponseBodyData) SetServices(v *GetServicesResponseBodyDataServices) *GetServicesResponseBodyData {
	s.Services = v
	return s
}

type GetServicesResponseBodyDataDetails struct {
	Details []*GetServicesResponseBodyDataDetailsDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
}

func (s GetServicesResponseBodyDataDetails) String() string {
	return tea.Prettify(s)
}

func (s GetServicesResponseBodyDataDetails) GoString() string {
	return s.String()
}

func (s *GetServicesResponseBodyDataDetails) SetDetails(v []*GetServicesResponseBodyDataDetailsDetails) *GetServicesResponseBodyDataDetails {
	s.Details = v
	return s
}

type GetServicesResponseBodyDataDetailsDetails struct {
	Pid         *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s GetServicesResponseBodyDataDetailsDetails) String() string {
	return tea.Prettify(s)
}

func (s GetServicesResponseBodyDataDetailsDetails) GoString() string {
	return s.String()
}

func (s *GetServicesResponseBodyDataDetailsDetails) SetPid(v string) *GetServicesResponseBodyDataDetailsDetails {
	s.Pid = &v
	return s
}

func (s *GetServicesResponseBodyDataDetailsDetails) SetRegionId(v string) *GetServicesResponseBodyDataDetailsDetails {
	s.RegionId = &v
	return s
}

func (s *GetServicesResponseBodyDataDetailsDetails) SetServiceName(v string) *GetServicesResponseBodyDataDetailsDetails {
	s.ServiceName = &v
	return s
}

type GetServicesResponseBodyDataServices struct {
	Services []*string `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
}

func (s GetServicesResponseBodyDataServices) String() string {
	return tea.Prettify(s)
}

func (s GetServicesResponseBodyDataServices) GoString() string {
	return s.String()
}

func (s *GetServicesResponseBodyDataServices) SetServices(v []*string) *GetServicesResponseBodyDataServices {
	s.Services = v
	return s
}

type GetServicesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetServicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetServicesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServicesResponse) GoString() string {
	return s.String()
}

func (s *GetServicesResponse) SetHeaders(v map[string]*string) *GetServicesResponse {
	s.Headers = v
	return s
}

func (s *GetServicesResponse) SetStatusCode(v int32) *GetServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServicesResponse) SetBody(v *GetServicesResponseBody) *GetServicesResponse {
	s.Body = v
	return s
}

type GetSpanNamesRequest struct {
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetSpanNamesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSpanNamesRequest) GoString() string {
	return s.String()
}

func (s *GetSpanNamesRequest) SetEndTime(v int64) *GetSpanNamesRequest {
	s.EndTime = &v
	return s
}

func (s *GetSpanNamesRequest) SetRegionId(v string) *GetSpanNamesRequest {
	s.RegionId = &v
	return s
}

func (s *GetSpanNamesRequest) SetServiceName(v string) *GetSpanNamesRequest {
	s.ServiceName = &v
	return s
}

func (s *GetSpanNamesRequest) SetStartTime(v int64) *GetSpanNamesRequest {
	s.StartTime = &v
	return s
}

type GetSpanNamesResponseBody struct {
	Data      *GetSpanNamesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSpanNamesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSpanNamesResponseBody) GoString() string {
	return s.String()
}

func (s *GetSpanNamesResponseBody) SetData(v *GetSpanNamesResponseBodyData) *GetSpanNamesResponseBody {
	s.Data = v
	return s
}

func (s *GetSpanNamesResponseBody) SetRequestId(v string) *GetSpanNamesResponseBody {
	s.RequestId = &v
	return s
}

type GetSpanNamesResponseBodyData struct {
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s GetSpanNamesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSpanNamesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSpanNamesResponseBodyData) SetData(v []*string) *GetSpanNamesResponseBodyData {
	s.Data = v
	return s
}

type GetSpanNamesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSpanNamesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSpanNamesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSpanNamesResponse) GoString() string {
	return s.String()
}

func (s *GetSpanNamesResponse) SetHeaders(v map[string]*string) *GetSpanNamesResponse {
	s.Headers = v
	return s
}

func (s *GetSpanNamesResponse) SetStatusCode(v int32) *GetSpanNamesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSpanNamesResponse) SetBody(v *GetSpanNamesResponseBody) *GetSpanNamesResponse {
	s.Body = v
	return s
}

type GetTagKeyRequest struct {
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	SpanName    *string `json:"SpanName,omitempty" xml:"SpanName,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetTagKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTagKeyRequest) GoString() string {
	return s.String()
}

func (s *GetTagKeyRequest) SetEndTime(v int64) *GetTagKeyRequest {
	s.EndTime = &v
	return s
}

func (s *GetTagKeyRequest) SetRegionId(v string) *GetTagKeyRequest {
	s.RegionId = &v
	return s
}

func (s *GetTagKeyRequest) SetServiceName(v string) *GetTagKeyRequest {
	s.ServiceName = &v
	return s
}

func (s *GetTagKeyRequest) SetSpanName(v string) *GetTagKeyRequest {
	s.SpanName = &v
	return s
}

func (s *GetTagKeyRequest) SetStartTime(v int64) *GetTagKeyRequest {
	s.StartTime = &v
	return s
}

type GetTagKeyResponseBody struct {
	Data      *GetTagKeyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTagKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTagKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GetTagKeyResponseBody) SetData(v *GetTagKeyResponseBodyData) *GetTagKeyResponseBody {
	s.Data = v
	return s
}

func (s *GetTagKeyResponseBody) SetRequestId(v string) *GetTagKeyResponseBody {
	s.RequestId = &v
	return s
}

type GetTagKeyResponseBodyData struct {
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s GetTagKeyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTagKeyResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTagKeyResponseBodyData) SetData(v []*string) *GetTagKeyResponseBodyData {
	s.Data = v
	return s
}

type GetTagKeyResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTagKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTagKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTagKeyResponse) GoString() string {
	return s.String()
}

func (s *GetTagKeyResponse) SetHeaders(v map[string]*string) *GetTagKeyResponse {
	s.Headers = v
	return s
}

func (s *GetTagKeyResponse) SetStatusCode(v int32) *GetTagKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTagKeyResponse) SetBody(v *GetTagKeyResponseBody) *GetTagKeyResponse {
	s.Body = v
	return s
}

type GetTagValRequest struct {
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	SpanName    *string `json:"SpanName,omitempty" xml:"SpanName,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TagKey      *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s GetTagValRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTagValRequest) GoString() string {
	return s.String()
}

func (s *GetTagValRequest) SetEndTime(v int64) *GetTagValRequest {
	s.EndTime = &v
	return s
}

func (s *GetTagValRequest) SetRegionId(v string) *GetTagValRequest {
	s.RegionId = &v
	return s
}

func (s *GetTagValRequest) SetServiceName(v string) *GetTagValRequest {
	s.ServiceName = &v
	return s
}

func (s *GetTagValRequest) SetSpanName(v string) *GetTagValRequest {
	s.SpanName = &v
	return s
}

func (s *GetTagValRequest) SetStartTime(v int64) *GetTagValRequest {
	s.StartTime = &v
	return s
}

func (s *GetTagValRequest) SetTagKey(v string) *GetTagValRequest {
	s.TagKey = &v
	return s
}

type GetTagValResponseBody struct {
	Data      *GetTagValResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTagValResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTagValResponseBody) GoString() string {
	return s.String()
}

func (s *GetTagValResponseBody) SetData(v *GetTagValResponseBodyData) *GetTagValResponseBody {
	s.Data = v
	return s
}

func (s *GetTagValResponseBody) SetRequestId(v string) *GetTagValResponseBody {
	s.RequestId = &v
	return s
}

type GetTagValResponseBodyData struct {
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s GetTagValResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTagValResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTagValResponseBodyData) SetData(v []*string) *GetTagValResponseBodyData {
	s.Data = v
	return s
}

type GetTagValResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTagValResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTagValResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTagValResponse) GoString() string {
	return s.String()
}

func (s *GetTagValResponse) SetHeaders(v map[string]*string) *GetTagValResponse {
	s.Headers = v
	return s
}

func (s *GetTagValResponse) SetStatusCode(v int32) *GetTagValResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTagValResponse) SetBody(v *GetTagValResponseBody) *GetTagValResponse {
	s.Body = v
	return s
}

type GetTokenRequest struct {
	AppType  *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTokenRequest) GoString() string {
	return s.String()
}

func (s *GetTokenRequest) SetAppType(v string) *GetTokenRequest {
	s.AppType = &v
	return s
}

func (s *GetTokenRequest) SetRegionId(v string) *GetTokenRequest {
	s.RegionId = &v
	return s
}

type GetTokenResponseBody struct {
	Data      *GetTokenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetTokenResponseBody) SetData(v *GetTokenResponseBodyData) *GetTokenResponseBody {
	s.Data = v
	return s
}

func (s *GetTokenResponseBody) SetRequestId(v string) *GetTokenResponseBody {
	s.RequestId = &v
	return s
}

type GetTokenResponseBodyData struct {
	Domain         *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	InternalDomain *string `json:"InternalDomain,omitempty" xml:"InternalDomain,omitempty"`
	LicenseKey     *string `json:"LicenseKey,omitempty" xml:"LicenseKey,omitempty"`
	Pid            *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
}

func (s GetTokenResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTokenResponseBodyData) SetDomain(v string) *GetTokenResponseBodyData {
	s.Domain = &v
	return s
}

func (s *GetTokenResponseBodyData) SetInternalDomain(v string) *GetTokenResponseBodyData {
	s.InternalDomain = &v
	return s
}

func (s *GetTokenResponseBodyData) SetLicenseKey(v string) *GetTokenResponseBodyData {
	s.LicenseKey = &v
	return s
}

func (s *GetTokenResponseBodyData) SetPid(v string) *GetTokenResponseBodyData {
	s.Pid = &v
	return s
}

type GetTokenResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTokenResponse) GoString() string {
	return s.String()
}

func (s *GetTokenResponse) SetHeaders(v map[string]*string) *GetTokenResponse {
	s.Headers = v
	return s
}

func (s *GetTokenResponse) SetStatusCode(v int32) *GetTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTokenResponse) SetBody(v *GetTokenResponseBody) *GetTokenResponse {
	s.Body = v
	return s
}

type GetTraceRequest struct {
	AppType  *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TraceID  *string `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
}

func (s GetTraceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTraceRequest) GoString() string {
	return s.String()
}

func (s *GetTraceRequest) SetAppType(v string) *GetTraceRequest {
	s.AppType = &v
	return s
}

func (s *GetTraceRequest) SetRegionId(v string) *GetTraceRequest {
	s.RegionId = &v
	return s
}

func (s *GetTraceRequest) SetTraceID(v string) *GetTraceRequest {
	s.TraceID = &v
	return s
}

type GetTraceResponseBody struct {
	Data      *GetTraceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTraceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBody) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBody) SetData(v *GetTraceResponseBodyData) *GetTraceResponseBody {
	s.Data = v
	return s
}

func (s *GetTraceResponseBody) SetRequestId(v string) *GetTraceResponseBody {
	s.RequestId = &v
	return s
}

type GetTraceResponseBodyData struct {
	CallChainInfo []*GetTraceResponseBodyDataCallChainInfo `json:"CallChainInfo,omitempty" xml:"CallChainInfo,omitempty" type:"Repeated"`
}

func (s GetTraceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodyData) SetCallChainInfo(v []*GetTraceResponseBodyDataCallChainInfo) *GetTraceResponseBodyData {
	s.CallChainInfo = v
	return s
}

type GetTraceResponseBodyDataCallChainInfo struct {
	Duration      *int64                                             `json:"Duration,omitempty" xml:"Duration,omitempty"`
	HaveStack     *bool                                              `json:"HaveStack,omitempty" xml:"HaveStack,omitempty"`
	LogEventList  *GetTraceResponseBodyDataCallChainInfoLogEventList `json:"LogEventList,omitempty" xml:"LogEventList,omitempty" type:"Struct"`
	OperationName *string                                            `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	ResultCode    *string                                            `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	RpcId         *string                                            `json:"RpcId,omitempty" xml:"RpcId,omitempty"`
	ServiceIp     *string                                            `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	ServiceName   *string                                            `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	TagEntryList  *GetTraceResponseBodyDataCallChainInfoTagEntryList `json:"TagEntryList,omitempty" xml:"TagEntryList,omitempty" type:"Struct"`
	Timestamp     *int64                                             `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	TraceID       *string                                            `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
}

func (s GetTraceResponseBodyDataCallChainInfo) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodyDataCallChainInfo) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodyDataCallChainInfo) SetDuration(v int64) *GetTraceResponseBodyDataCallChainInfo {
	s.Duration = &v
	return s
}

func (s *GetTraceResponseBodyDataCallChainInfo) SetHaveStack(v bool) *GetTraceResponseBodyDataCallChainInfo {
	s.HaveStack = &v
	return s
}

func (s *GetTraceResponseBodyDataCallChainInfo) SetLogEventList(v *GetTraceResponseBodyDataCallChainInfoLogEventList) *GetTraceResponseBodyDataCallChainInfo {
	s.LogEventList = v
	return s
}

func (s *GetTraceResponseBodyDataCallChainInfo) SetOperationName(v string) *GetTraceResponseBodyDataCallChainInfo {
	s.OperationName = &v
	return s
}

func (s *GetTraceResponseBodyDataCallChainInfo) SetResultCode(v string) *GetTraceResponseBodyDataCallChainInfo {
	s.ResultCode = &v
	return s
}

func (s *GetTraceResponseBodyDataCallChainInfo) SetRpcId(v string) *GetTraceResponseBodyDataCallChainInfo {
	s.RpcId = &v
	return s
}

func (s *GetTraceResponseBodyDataCallChainInfo) SetServiceIp(v string) *GetTraceResponseBodyDataCallChainInfo {
	s.ServiceIp = &v
	return s
}

func (s *GetTraceResponseBodyDataCallChainInfo) SetServiceName(v string) *GetTraceResponseBodyDataCallChainInfo {
	s.ServiceName = &v
	return s
}

func (s *GetTraceResponseBodyDataCallChainInfo) SetTagEntryList(v *GetTraceResponseBodyDataCallChainInfoTagEntryList) *GetTraceResponseBodyDataCallChainInfo {
	s.TagEntryList = v
	return s
}

func (s *GetTraceResponseBodyDataCallChainInfo) SetTimestamp(v int64) *GetTraceResponseBodyDataCallChainInfo {
	s.Timestamp = &v
	return s
}

func (s *GetTraceResponseBodyDataCallChainInfo) SetTraceID(v string) *GetTraceResponseBodyDataCallChainInfo {
	s.TraceID = &v
	return s
}

type GetTraceResponseBodyDataCallChainInfoLogEventList struct {
	LogEvent []*GetTraceResponseBodyDataCallChainInfoLogEventListLogEvent `json:"LogEvent,omitempty" xml:"LogEvent,omitempty" type:"Repeated"`
}

func (s GetTraceResponseBodyDataCallChainInfoLogEventList) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodyDataCallChainInfoLogEventList) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodyDataCallChainInfoLogEventList) SetLogEvent(v []*GetTraceResponseBodyDataCallChainInfoLogEventListLogEvent) *GetTraceResponseBodyDataCallChainInfoLogEventList {
	s.LogEvent = v
	return s
}

type GetTraceResponseBodyDataCallChainInfoLogEventListLogEvent struct {
	TagEntryList *GetTraceResponseBodyDataCallChainInfoLogEventListLogEventTagEntryList `json:"TagEntryList,omitempty" xml:"TagEntryList,omitempty" type:"Struct"`
	Timestamp    *int64                                                                 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s GetTraceResponseBodyDataCallChainInfoLogEventListLogEvent) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodyDataCallChainInfoLogEventListLogEvent) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodyDataCallChainInfoLogEventListLogEvent) SetTagEntryList(v *GetTraceResponseBodyDataCallChainInfoLogEventListLogEventTagEntryList) *GetTraceResponseBodyDataCallChainInfoLogEventListLogEvent {
	s.TagEntryList = v
	return s
}

func (s *GetTraceResponseBodyDataCallChainInfoLogEventListLogEvent) SetTimestamp(v int64) *GetTraceResponseBodyDataCallChainInfoLogEventListLogEvent {
	s.Timestamp = &v
	return s
}

type GetTraceResponseBodyDataCallChainInfoLogEventListLogEventTagEntryList struct {
	TagEntry []*GetTraceResponseBodyDataCallChainInfoLogEventListLogEventTagEntryListTagEntry `json:"TagEntry,omitempty" xml:"TagEntry,omitempty" type:"Repeated"`
}

func (s GetTraceResponseBodyDataCallChainInfoLogEventListLogEventTagEntryList) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodyDataCallChainInfoLogEventListLogEventTagEntryList) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodyDataCallChainInfoLogEventListLogEventTagEntryList) SetTagEntry(v []*GetTraceResponseBodyDataCallChainInfoLogEventListLogEventTagEntryListTagEntry) *GetTraceResponseBodyDataCallChainInfoLogEventListLogEventTagEntryList {
	s.TagEntry = v
	return s
}

type GetTraceResponseBodyDataCallChainInfoLogEventListLogEventTagEntryListTagEntry struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTraceResponseBodyDataCallChainInfoLogEventListLogEventTagEntryListTagEntry) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodyDataCallChainInfoLogEventListLogEventTagEntryListTagEntry) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodyDataCallChainInfoLogEventListLogEventTagEntryListTagEntry) SetKey(v string) *GetTraceResponseBodyDataCallChainInfoLogEventListLogEventTagEntryListTagEntry {
	s.Key = &v
	return s
}

func (s *GetTraceResponseBodyDataCallChainInfoLogEventListLogEventTagEntryListTagEntry) SetValue(v string) *GetTraceResponseBodyDataCallChainInfoLogEventListLogEventTagEntryListTagEntry {
	s.Value = &v
	return s
}

type GetTraceResponseBodyDataCallChainInfoTagEntryList struct {
	TagEntry []*GetTraceResponseBodyDataCallChainInfoTagEntryListTagEntry `json:"TagEntry,omitempty" xml:"TagEntry,omitempty" type:"Repeated"`
}

func (s GetTraceResponseBodyDataCallChainInfoTagEntryList) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodyDataCallChainInfoTagEntryList) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodyDataCallChainInfoTagEntryList) SetTagEntry(v []*GetTraceResponseBodyDataCallChainInfoTagEntryListTagEntry) *GetTraceResponseBodyDataCallChainInfoTagEntryList {
	s.TagEntry = v
	return s
}

type GetTraceResponseBodyDataCallChainInfoTagEntryListTagEntry struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTraceResponseBodyDataCallChainInfoTagEntryListTagEntry) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodyDataCallChainInfoTagEntryListTagEntry) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodyDataCallChainInfoTagEntryListTagEntry) SetKey(v string) *GetTraceResponseBodyDataCallChainInfoTagEntryListTagEntry {
	s.Key = &v
	return s
}

func (s *GetTraceResponseBodyDataCallChainInfoTagEntryListTagEntry) SetValue(v string) *GetTraceResponseBodyDataCallChainInfoTagEntryListTagEntry {
	s.Value = &v
	return s
}

type GetTraceResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTraceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTraceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponse) GoString() string {
	return s.String()
}

func (s *GetTraceResponse) SetHeaders(v map[string]*string) *GetTraceResponse {
	s.Headers = v
	return s
}

func (s *GetTraceResponse) SetStatusCode(v int32) *GetTraceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTraceResponse) SetBody(v *GetTraceResponseBody) *GetTraceResponse {
	s.Body = v
	return s
}

type GetTraceAppsRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetTraceAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTraceAppsRequest) GoString() string {
	return s.String()
}

func (s *GetTraceAppsRequest) SetRegionId(v string) *GetTraceAppsRequest {
	s.RegionId = &v
	return s
}

type GetTraceAppsResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTraceAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTraceAppsResponseBody) GoString() string {
	return s.String()
}

func (s *GetTraceAppsResponseBody) SetData(v string) *GetTraceAppsResponseBody {
	s.Data = &v
	return s
}

func (s *GetTraceAppsResponseBody) SetRequestId(v string) *GetTraceAppsResponseBody {
	s.RequestId = &v
	return s
}

type GetTraceAppsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTraceAppsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTraceAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTraceAppsResponse) GoString() string {
	return s.String()
}

func (s *GetTraceAppsResponse) SetHeaders(v map[string]*string) *GetTraceAppsResponse {
	s.Headers = v
	return s
}

func (s *GetTraceAppsResponse) SetStatusCode(v int32) *GetTraceAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTraceAppsResponse) SetBody(v *GetTraceAppsResponseBody) *GetTraceAppsResponse {
	s.Body = v
	return s
}

type ImportAppAlertRulesRequest struct {
	ContactGroupIds *string `json:"ContactGroupIds,omitempty" xml:"ContactGroupIds,omitempty"`
	IsAutoStart     *bool   `json:"IsAutoStart,omitempty" xml:"IsAutoStart,omitempty"`
	Pids            *string `json:"Pids,omitempty" xml:"Pids,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TemplateAlertId *string `json:"TemplateAlertId,omitempty" xml:"TemplateAlertId,omitempty"`
}

func (s ImportAppAlertRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportAppAlertRulesRequest) GoString() string {
	return s.String()
}

func (s *ImportAppAlertRulesRequest) SetContactGroupIds(v string) *ImportAppAlertRulesRequest {
	s.ContactGroupIds = &v
	return s
}

func (s *ImportAppAlertRulesRequest) SetIsAutoStart(v bool) *ImportAppAlertRulesRequest {
	s.IsAutoStart = &v
	return s
}

func (s *ImportAppAlertRulesRequest) SetPids(v string) *ImportAppAlertRulesRequest {
	s.Pids = &v
	return s
}

func (s *ImportAppAlertRulesRequest) SetRegionId(v string) *ImportAppAlertRulesRequest {
	s.RegionId = &v
	return s
}

func (s *ImportAppAlertRulesRequest) SetTemplateAlertId(v string) *ImportAppAlertRulesRequest {
	s.TemplateAlertId = &v
	return s
}

type ImportAppAlertRulesResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImportAppAlertRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportAppAlertRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ImportAppAlertRulesResponseBody) SetData(v string) *ImportAppAlertRulesResponseBody {
	s.Data = &v
	return s
}

func (s *ImportAppAlertRulesResponseBody) SetRequestId(v string) *ImportAppAlertRulesResponseBody {
	s.RequestId = &v
	return s
}

type ImportAppAlertRulesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ImportAppAlertRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ImportAppAlertRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportAppAlertRulesResponse) GoString() string {
	return s.String()
}

func (s *ImportAppAlertRulesResponse) SetHeaders(v map[string]*string) *ImportAppAlertRulesResponse {
	s.Headers = v
	return s
}

func (s *ImportAppAlertRulesResponse) SetStatusCode(v int32) *ImportAppAlertRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportAppAlertRulesResponse) SetBody(v *ImportAppAlertRulesResponseBody) *ImportAppAlertRulesResponse {
	s.Body = v
	return s
}

type MetricQueryRequest struct {
	CustomFilters  *string                      `json:"CustomFilters,omitempty" xml:"CustomFilters,omitempty"`
	Dimensions     []*string                    `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	EndTime        *int64                       `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Filters        []*MetricQueryRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	IintervalInSec *int32                       `json:"IintervalInSec,omitempty" xml:"IintervalInSec,omitempty"`
	Limit          *int32                       `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Measures       []*string                    `json:"Measures,omitempty" xml:"Measures,omitempty" type:"Repeated"`
	Metric         *string                      `json:"Metric,omitempty" xml:"Metric,omitempty"`
	Order          *string                      `json:"Order,omitempty" xml:"Order,omitempty"`
	OrderBy        *string                      `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	SecurityToken  *string                      `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StartTime      *int64                       `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Where          *string                      `json:"Where,omitempty" xml:"Where,omitempty"`
}

func (s MetricQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s MetricQueryRequest) GoString() string {
	return s.String()
}

func (s *MetricQueryRequest) SetCustomFilters(v string) *MetricQueryRequest {
	s.CustomFilters = &v
	return s
}

func (s *MetricQueryRequest) SetDimensions(v []*string) *MetricQueryRequest {
	s.Dimensions = v
	return s
}

func (s *MetricQueryRequest) SetEndTime(v int64) *MetricQueryRequest {
	s.EndTime = &v
	return s
}

func (s *MetricQueryRequest) SetFilters(v []*MetricQueryRequestFilters) *MetricQueryRequest {
	s.Filters = v
	return s
}

func (s *MetricQueryRequest) SetIintervalInSec(v int32) *MetricQueryRequest {
	s.IintervalInSec = &v
	return s
}

func (s *MetricQueryRequest) SetLimit(v int32) *MetricQueryRequest {
	s.Limit = &v
	return s
}

func (s *MetricQueryRequest) SetMeasures(v []*string) *MetricQueryRequest {
	s.Measures = v
	return s
}

func (s *MetricQueryRequest) SetMetric(v string) *MetricQueryRequest {
	s.Metric = &v
	return s
}

func (s *MetricQueryRequest) SetOrder(v string) *MetricQueryRequest {
	s.Order = &v
	return s
}

func (s *MetricQueryRequest) SetOrderBy(v string) *MetricQueryRequest {
	s.OrderBy = &v
	return s
}

func (s *MetricQueryRequest) SetSecurityToken(v string) *MetricQueryRequest {
	s.SecurityToken = &v
	return s
}

func (s *MetricQueryRequest) SetStartTime(v int64) *MetricQueryRequest {
	s.StartTime = &v
	return s
}

func (s *MetricQueryRequest) SetWhere(v string) *MetricQueryRequest {
	s.Where = &v
	return s
}

type MetricQueryRequestFilters struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s MetricQueryRequestFilters) String() string {
	return tea.Prettify(s)
}

func (s MetricQueryRequestFilters) GoString() string {
	return s.String()
}

func (s *MetricQueryRequestFilters) SetKey(v string) *MetricQueryRequestFilters {
	s.Key = &v
	return s
}

func (s *MetricQueryRequestFilters) SetValue(v string) *MetricQueryRequestFilters {
	s.Value = &v
	return s
}

type MetricQueryResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MetricQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MetricQueryResponseBody) GoString() string {
	return s.String()
}

func (s *MetricQueryResponseBody) SetData(v string) *MetricQueryResponseBody {
	s.Data = &v
	return s
}

func (s *MetricQueryResponseBody) SetRequestId(v string) *MetricQueryResponseBody {
	s.RequestId = &v
	return s
}

type MetricQueryResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *MetricQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MetricQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s MetricQueryResponse) GoString() string {
	return s.String()
}

func (s *MetricQueryResponse) SetHeaders(v map[string]*string) *MetricQueryResponse {
	s.Headers = v
	return s
}

func (s *MetricQueryResponse) SetStatusCode(v int32) *MetricQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *MetricQueryResponse) SetBody(v *MetricQueryResponseBody) *MetricQueryResponse {
	s.Body = v
	return s
}

type SearchAlertContactRequest struct {
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	PageSize    *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Phone       *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
}

func (s SearchAlertContactRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertContactRequest) GoString() string {
	return s.String()
}

func (s *SearchAlertContactRequest) SetContactName(v string) *SearchAlertContactRequest {
	s.ContactName = &v
	return s
}

func (s *SearchAlertContactRequest) SetCurrentPage(v string) *SearchAlertContactRequest {
	s.CurrentPage = &v
	return s
}

func (s *SearchAlertContactRequest) SetEmail(v string) *SearchAlertContactRequest {
	s.Email = &v
	return s
}

func (s *SearchAlertContactRequest) SetPageSize(v string) *SearchAlertContactRequest {
	s.PageSize = &v
	return s
}

func (s *SearchAlertContactRequest) SetPhone(v string) *SearchAlertContactRequest {
	s.Phone = &v
	return s
}

type SearchAlertContactResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchAlertContactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertContactResponseBody) GoString() string {
	return s.String()
}

func (s *SearchAlertContactResponseBody) SetData(v string) *SearchAlertContactResponseBody {
	s.Data = &v
	return s
}

func (s *SearchAlertContactResponseBody) SetRequestId(v string) *SearchAlertContactResponseBody {
	s.RequestId = &v
	return s
}

type SearchAlertContactResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchAlertContactResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchAlertContactResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertContactResponse) GoString() string {
	return s.String()
}

func (s *SearchAlertContactResponse) SetHeaders(v map[string]*string) *SearchAlertContactResponse {
	s.Headers = v
	return s
}

func (s *SearchAlertContactResponse) SetStatusCode(v int32) *SearchAlertContactResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchAlertContactResponse) SetBody(v *SearchAlertContactResponseBody) *SearchAlertContactResponse {
	s.Body = v
	return s
}

type SearchAlertContactGroupRequest struct {
	ContactGroupName *string `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
}

func (s SearchAlertContactGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertContactGroupRequest) GoString() string {
	return s.String()
}

func (s *SearchAlertContactGroupRequest) SetContactGroupName(v string) *SearchAlertContactGroupRequest {
	s.ContactGroupName = &v
	return s
}

type SearchAlertContactGroupResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchAlertContactGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertContactGroupResponseBody) GoString() string {
	return s.String()
}

func (s *SearchAlertContactGroupResponseBody) SetData(v string) *SearchAlertContactGroupResponseBody {
	s.Data = &v
	return s
}

func (s *SearchAlertContactGroupResponseBody) SetRequestId(v string) *SearchAlertContactGroupResponseBody {
	s.RequestId = &v
	return s
}

type SearchAlertContactGroupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchAlertContactGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchAlertContactGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertContactGroupResponse) GoString() string {
	return s.String()
}

func (s *SearchAlertContactGroupResponse) SetHeaders(v map[string]*string) *SearchAlertContactGroupResponse {
	s.Headers = v
	return s
}

func (s *SearchAlertContactGroupResponse) SetStatusCode(v int32) *SearchAlertContactGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchAlertContactGroupResponse) SetBody(v *SearchAlertContactGroupResponseBody) *SearchAlertContactGroupResponse {
	s.Body = v
	return s
}

type SearchTraceCountRequest struct {
	AppType       *string                       `json:"AppType,omitempty" xml:"AppType,omitempty"`
	EndTime       *int64                        `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MinDuration   *int64                        `json:"MinDuration,omitempty" xml:"MinDuration,omitempty"`
	OperationName *string                       `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	RegionId      *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceIp     *string                       `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	ServiceName   *string                       `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	StartTime     *int64                        `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Tag           []*SearchTraceCountRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s SearchTraceCountRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchTraceCountRequest) GoString() string {
	return s.String()
}

func (s *SearchTraceCountRequest) SetAppType(v string) *SearchTraceCountRequest {
	s.AppType = &v
	return s
}

func (s *SearchTraceCountRequest) SetEndTime(v int64) *SearchTraceCountRequest {
	s.EndTime = &v
	return s
}

func (s *SearchTraceCountRequest) SetMinDuration(v int64) *SearchTraceCountRequest {
	s.MinDuration = &v
	return s
}

func (s *SearchTraceCountRequest) SetOperationName(v string) *SearchTraceCountRequest {
	s.OperationName = &v
	return s
}

func (s *SearchTraceCountRequest) SetRegionId(v string) *SearchTraceCountRequest {
	s.RegionId = &v
	return s
}

func (s *SearchTraceCountRequest) SetServiceIp(v string) *SearchTraceCountRequest {
	s.ServiceIp = &v
	return s
}

func (s *SearchTraceCountRequest) SetServiceName(v string) *SearchTraceCountRequest {
	s.ServiceName = &v
	return s
}

func (s *SearchTraceCountRequest) SetStartTime(v int64) *SearchTraceCountRequest {
	s.StartTime = &v
	return s
}

func (s *SearchTraceCountRequest) SetTag(v []*SearchTraceCountRequestTag) *SearchTraceCountRequest {
	s.Tag = v
	return s
}

type SearchTraceCountRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SearchTraceCountRequestTag) String() string {
	return tea.Prettify(s)
}

func (s SearchTraceCountRequestTag) GoString() string {
	return s.String()
}

func (s *SearchTraceCountRequestTag) SetKey(v string) *SearchTraceCountRequestTag {
	s.Key = &v
	return s
}

func (s *SearchTraceCountRequestTag) SetValue(v string) *SearchTraceCountRequestTag {
	s.Value = &v
	return s
}

type SearchTraceCountResponseBody struct {
	Data      *SearchTraceCountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchTraceCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchTraceCountResponseBody) GoString() string {
	return s.String()
}

func (s *SearchTraceCountResponseBody) SetData(v *SearchTraceCountResponseBodyData) *SearchTraceCountResponseBody {
	s.Data = v
	return s
}

func (s *SearchTraceCountResponseBody) SetRequestId(v string) *SearchTraceCountResponseBody {
	s.RequestId = &v
	return s
}

type SearchTraceCountResponseBodyData struct {
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s SearchTraceCountResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchTraceCountResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchTraceCountResponseBodyData) SetCount(v int64) *SearchTraceCountResponseBodyData {
	s.Count = &v
	return s
}

type SearchTraceCountResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchTraceCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchTraceCountResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchTraceCountResponse) GoString() string {
	return s.String()
}

func (s *SearchTraceCountResponse) SetHeaders(v map[string]*string) *SearchTraceCountResponse {
	s.Headers = v
	return s
}

func (s *SearchTraceCountResponse) SetStatusCode(v int32) *SearchTraceCountResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchTraceCountResponse) SetBody(v *SearchTraceCountResponseBody) *SearchTraceCountResponse {
	s.Body = v
	return s
}

type SearchTracesRequest struct {
	AppType       *string                   `json:"AppType,omitempty" xml:"AppType,omitempty"`
	EndTime       *int64                    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MinDuration   *int64                    `json:"MinDuration,omitempty" xml:"MinDuration,omitempty"`
	OperationName *string                   `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	PageIndex     *int32                    `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize      *int32                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId      *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Reverse       *bool                     `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	ServiceIp     *string                   `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	ServiceName   *string                   `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	StartTime     *int64                    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Tag           []*SearchTracesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s SearchTracesRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesRequest) GoString() string {
	return s.String()
}

func (s *SearchTracesRequest) SetAppType(v string) *SearchTracesRequest {
	s.AppType = &v
	return s
}

func (s *SearchTracesRequest) SetEndTime(v int64) *SearchTracesRequest {
	s.EndTime = &v
	return s
}

func (s *SearchTracesRequest) SetMinDuration(v int64) *SearchTracesRequest {
	s.MinDuration = &v
	return s
}

func (s *SearchTracesRequest) SetOperationName(v string) *SearchTracesRequest {
	s.OperationName = &v
	return s
}

func (s *SearchTracesRequest) SetPageIndex(v int32) *SearchTracesRequest {
	s.PageIndex = &v
	return s
}

func (s *SearchTracesRequest) SetPageSize(v int32) *SearchTracesRequest {
	s.PageSize = &v
	return s
}

func (s *SearchTracesRequest) SetRegionId(v string) *SearchTracesRequest {
	s.RegionId = &v
	return s
}

func (s *SearchTracesRequest) SetReverse(v bool) *SearchTracesRequest {
	s.Reverse = &v
	return s
}

func (s *SearchTracesRequest) SetServiceIp(v string) *SearchTracesRequest {
	s.ServiceIp = &v
	return s
}

func (s *SearchTracesRequest) SetServiceName(v string) *SearchTracesRequest {
	s.ServiceName = &v
	return s
}

func (s *SearchTracesRequest) SetStartTime(v int64) *SearchTracesRequest {
	s.StartTime = &v
	return s
}

func (s *SearchTracesRequest) SetTag(v []*SearchTracesRequestTag) *SearchTracesRequest {
	s.Tag = v
	return s
}

type SearchTracesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SearchTracesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesRequestTag) GoString() string {
	return s.String()
}

func (s *SearchTracesRequestTag) SetKey(v string) *SearchTracesRequestTag {
	s.Key = &v
	return s
}

func (s *SearchTracesRequestTag) SetValue(v string) *SearchTracesRequestTag {
	s.Value = &v
	return s
}

type SearchTracesResponseBody struct {
	Data      *SearchTracesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchTracesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchTracesResponseBody) SetData(v *SearchTracesResponseBodyData) *SearchTracesResponseBody {
	s.Data = v
	return s
}

func (s *SearchTracesResponseBody) SetRequestId(v string) *SearchTracesResponseBody {
	s.RequestId = &v
	return s
}

type SearchTracesResponseBodyData struct {
	TraceInfo []*SearchTracesResponseBodyDataTraceInfo `json:"TraceInfo,omitempty" xml:"TraceInfo,omitempty" type:"Repeated"`
}

func (s SearchTracesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchTracesResponseBodyData) SetTraceInfo(v []*SearchTracesResponseBodyDataTraceInfo) *SearchTracesResponseBodyData {
	s.TraceInfo = v
	return s
}

type SearchTracesResponseBodyDataTraceInfo struct {
	Duration      *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	OperationName *string `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	ServiceIp     *string `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	ServiceName   *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	Timestamp     *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	TraceID       *string `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
}

func (s SearchTracesResponseBodyDataTraceInfo) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesResponseBodyDataTraceInfo) GoString() string {
	return s.String()
}

func (s *SearchTracesResponseBodyDataTraceInfo) SetDuration(v int32) *SearchTracesResponseBodyDataTraceInfo {
	s.Duration = &v
	return s
}

func (s *SearchTracesResponseBodyDataTraceInfo) SetOperationName(v string) *SearchTracesResponseBodyDataTraceInfo {
	s.OperationName = &v
	return s
}

func (s *SearchTracesResponseBodyDataTraceInfo) SetServiceIp(v string) *SearchTracesResponseBodyDataTraceInfo {
	s.ServiceIp = &v
	return s
}

func (s *SearchTracesResponseBodyDataTraceInfo) SetServiceName(v string) *SearchTracesResponseBodyDataTraceInfo {
	s.ServiceName = &v
	return s
}

func (s *SearchTracesResponseBodyDataTraceInfo) SetTimestamp(v int64) *SearchTracesResponseBodyDataTraceInfo {
	s.Timestamp = &v
	return s
}

func (s *SearchTracesResponseBodyDataTraceInfo) SetTraceID(v string) *SearchTracesResponseBodyDataTraceInfo {
	s.TraceID = &v
	return s
}

type SearchTracesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchTracesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchTracesResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesResponse) GoString() string {
	return s.String()
}

func (s *SearchTracesResponse) SetHeaders(v map[string]*string) *SearchTracesResponse {
	s.Headers = v
	return s
}

func (s *SearchTracesResponse) SetStatusCode(v int32) *SearchTracesResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchTracesResponse) SetBody(v *SearchTracesResponseBody) *SearchTracesResponse {
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
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-2-pop":          tea.String("arms.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("arms.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("arms.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("arms.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("arms.aliyuncs.com"),
		"cn-edge-1":                   tea.String("arms.aliyuncs.com"),
		"cn-fujian":                   tea.String("arms.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("arms.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("arms.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("arms.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("arms.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("arms.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("arms.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("arms.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("arms.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("arms.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("arms.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("arms.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("arms.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("arms.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("arms.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("arms.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("arms.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("arms.aliyuncs.com"),
		"cn-wuhan":                    tea.String("arms.aliyuncs.com"),
		"cn-yushanfang":               tea.String("arms.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("arms.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("arms.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("arms.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("arms.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("arms.aliyuncs.com"),
		"me-east-1":                   tea.String("arms.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("arms.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("arms"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) ARMSQueryDataSetWithOptions(request *ARMSQueryDataSetRequest, runtime *util.RuntimeOptions) (_result *ARMSQueryDataSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetId)) {
		query["DatasetId"] = request.DatasetId
	}

	if !tea.BoolValue(util.IsUnset(request.DateStr)) {
		query["DateStr"] = request.DateStr
	}

	if !tea.BoolValue(util.IsUnset(request.Dimensions)) {
		query["Dimensions"] = request.Dimensions
	}

	if !tea.BoolValue(util.IsUnset(request.HungryMode)) {
		query["HungryMode"] = request.HungryMode
	}

	if !tea.BoolValue(util.IsUnset(request.IntervalInSec)) {
		query["IntervalInSec"] = request.IntervalInSec
	}

	if !tea.BoolValue(util.IsUnset(request.IsDrillDown)) {
		query["IsDrillDown"] = request.IsDrillDown
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.MaxTime)) {
		query["MaxTime"] = request.MaxTime
	}

	if !tea.BoolValue(util.IsUnset(request.Measures)) {
		query["Measures"] = request.Measures
	}

	if !tea.BoolValue(util.IsUnset(request.MinTime)) {
		query["MinTime"] = request.MinTime
	}

	if !tea.BoolValue(util.IsUnset(request.OptionalDims)) {
		query["OptionalDims"] = request.OptionalDims
	}

	if !tea.BoolValue(util.IsUnset(request.OrderByKey)) {
		query["OrderByKey"] = request.OrderByKey
	}

	if !tea.BoolValue(util.IsUnset(request.ReduceTail)) {
		query["ReduceTail"] = request.ReduceTail
	}

	if !tea.BoolValue(util.IsUnset(request.RequiredDims)) {
		query["RequiredDims"] = request.RequiredDims
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ARMSQueryDataSet"),
		Version:     tea.String("2019-02-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ARMSQueryDataSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ARMSQueryDataSet(request *ARMSQueryDataSetRequest) (_result *ARMSQueryDataSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ARMSQueryDataSetResponse{}
	_body, _err := client.ARMSQueryDataSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAlertContactWithOptions(request *CreateAlertContactRequest, runtime *util.RuntimeOptions) (_result *CreateAlertContactResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactName)) {
		query["ContactName"] = request.ContactName
	}

	if !tea.BoolValue(util.IsUnset(request.DingRobotWebhookUrl)) {
		query["DingRobotWebhookUrl"] = request.DingRobotWebhookUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNum)) {
		query["PhoneNum"] = request.PhoneNum
	}

	if !tea.BoolValue(util.IsUnset(request.SystemNoc)) {
		query["SystemNoc"] = request.SystemNoc
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAlertContact"),
		Version:     tea.String("2019-02-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAlertContactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAlertContact(request *CreateAlertContactRequest) (_result *CreateAlertContactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAlertContactResponse{}
	_body, _err := client.CreateAlertContactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAlertContactGroupWithOptions(request *CreateAlertContactGroupRequest, runtime *util.RuntimeOptions) (_result *CreateAlertContactGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactGroupName)) {
		query["ContactGroupName"] = request.ContactGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.ContactIds)) {
		query["ContactIds"] = request.ContactIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAlertContactGroup"),
		Version:     tea.String("2019-02-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAlertContactGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAlertContactGroup(request *CreateAlertContactGroupRequest) (_result *CreateAlertContactGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAlertContactGroupResponse{}
	_body, _err := client.CreateAlertContactGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetIpOrHostsWithOptions(request *GetIpOrHostsRequest, runtime *util.RuntimeOptions) (_result *GetIpOrHostsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetIpOrHosts"),
		Version:     tea.String("2019-02-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetIpOrHostsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetIpOrHosts(request *GetIpOrHostsRequest) (_result *GetIpOrHostsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetIpOrHostsResponse{}
	_body, _err := client.GetIpOrHostsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetServicesWithOptions(request *GetServicesRequest, runtime *util.RuntimeOptions) (_result *GetServicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["AppType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetServices"),
		Version:     tea.String("2019-02-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetServices(request *GetServicesRequest) (_result *GetServicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetServicesResponse{}
	_body, _err := client.GetServicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSpanNamesWithOptions(request *GetSpanNamesRequest, runtime *util.RuntimeOptions) (_result *GetSpanNamesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSpanNames"),
		Version:     tea.String("2019-02-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSpanNamesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSpanNames(request *GetSpanNamesRequest) (_result *GetSpanNamesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSpanNamesResponse{}
	_body, _err := client.GetSpanNamesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTagKeyWithOptions(request *GetTagKeyRequest, runtime *util.RuntimeOptions) (_result *GetTagKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	if !tea.BoolValue(util.IsUnset(request.SpanName)) {
		query["SpanName"] = request.SpanName
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTagKey"),
		Version:     tea.String("2019-02-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTagKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTagKey(request *GetTagKeyRequest) (_result *GetTagKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTagKeyResponse{}
	_body, _err := client.GetTagKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTagValWithOptions(request *GetTagValRequest, runtime *util.RuntimeOptions) (_result *GetTagValResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	if !tea.BoolValue(util.IsUnset(request.SpanName)) {
		query["SpanName"] = request.SpanName
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTagVal"),
		Version:     tea.String("2019-02-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTagValResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTagVal(request *GetTagValRequest) (_result *GetTagValResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTagValResponse{}
	_body, _err := client.GetTagValWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTokenWithOptions(request *GetTokenRequest, runtime *util.RuntimeOptions) (_result *GetTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["AppType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetToken"),
		Version:     tea.String("2019-02-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetToken(request *GetTokenRequest) (_result *GetTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTokenResponse{}
	_body, _err := client.GetTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTraceWithOptions(request *GetTraceRequest, runtime *util.RuntimeOptions) (_result *GetTraceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["AppType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TraceID)) {
		query["TraceID"] = request.TraceID
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTrace"),
		Version:     tea.String("2019-02-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTraceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTrace(request *GetTraceRequest) (_result *GetTraceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTraceResponse{}
	_body, _err := client.GetTraceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTraceAppsWithOptions(request *GetTraceAppsRequest, runtime *util.RuntimeOptions) (_result *GetTraceAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTraceApps"),
		Version:     tea.String("2019-02-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTraceAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTraceApps(request *GetTraceAppsRequest) (_result *GetTraceAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTraceAppsResponse{}
	_body, _err := client.GetTraceAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImportAppAlertRulesWithOptions(request *ImportAppAlertRulesRequest, runtime *util.RuntimeOptions) (_result *ImportAppAlertRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactGroupIds)) {
		query["ContactGroupIds"] = request.ContactGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.IsAutoStart)) {
		query["IsAutoStart"] = request.IsAutoStart
	}

	if !tea.BoolValue(util.IsUnset(request.Pids)) {
		query["Pids"] = request.Pids
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateAlertId)) {
		query["TemplateAlertId"] = request.TemplateAlertId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ImportAppAlertRules"),
		Version:     tea.String("2019-02-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ImportAppAlertRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImportAppAlertRules(request *ImportAppAlertRulesRequest) (_result *ImportAppAlertRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImportAppAlertRulesResponse{}
	_body, _err := client.ImportAppAlertRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MetricQueryWithOptions(request *MetricQueryRequest, runtime *util.RuntimeOptions) (_result *MetricQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomFilters)) {
		query["CustomFilters"] = request.CustomFilters
	}

	if !tea.BoolValue(util.IsUnset(request.Dimensions)) {
		query["Dimensions"] = request.Dimensions
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Filters)) {
		query["Filters"] = request.Filters
	}

	if !tea.BoolValue(util.IsUnset(request.IintervalInSec)) {
		query["IintervalInSec"] = request.IintervalInSec
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Measures)) {
		query["Measures"] = request.Measures
	}

	if !tea.BoolValue(util.IsUnset(request.Metric)) {
		query["Metric"] = request.Metric
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["OrderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Where)) {
		query["Where"] = request.Where
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("MetricQuery"),
		Version:     tea.String("2019-02-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MetricQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MetricQuery(request *MetricQueryRequest) (_result *MetricQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MetricQueryResponse{}
	_body, _err := client.MetricQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchAlertContactWithOptions(request *SearchAlertContactRequest, runtime *util.RuntimeOptions) (_result *SearchAlertContactResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactName)) {
		query["ContactName"] = request.ContactName
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Phone)) {
		query["Phone"] = request.Phone
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchAlertContact"),
		Version:     tea.String("2019-02-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchAlertContactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchAlertContact(request *SearchAlertContactRequest) (_result *SearchAlertContactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchAlertContactResponse{}
	_body, _err := client.SearchAlertContactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchAlertContactGroupWithOptions(request *SearchAlertContactGroupRequest, runtime *util.RuntimeOptions) (_result *SearchAlertContactGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactGroupName)) {
		query["ContactGroupName"] = request.ContactGroupName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchAlertContactGroup"),
		Version:     tea.String("2019-02-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchAlertContactGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchAlertContactGroup(request *SearchAlertContactGroupRequest) (_result *SearchAlertContactGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchAlertContactGroupResponse{}
	_body, _err := client.SearchAlertContactGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchTraceCountWithOptions(request *SearchTraceCountRequest, runtime *util.RuntimeOptions) (_result *SearchTraceCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["AppType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MinDuration)) {
		query["MinDuration"] = request.MinDuration
	}

	if !tea.BoolValue(util.IsUnset(request.OperationName)) {
		query["OperationName"] = request.OperationName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceIp)) {
		query["ServiceIp"] = request.ServiceIp
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchTraceCount"),
		Version:     tea.String("2019-02-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchTraceCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchTraceCount(request *SearchTraceCountRequest) (_result *SearchTraceCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchTraceCountResponse{}
	_body, _err := client.SearchTraceCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchTracesWithOptions(request *SearchTracesRequest, runtime *util.RuntimeOptions) (_result *SearchTracesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["AppType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MinDuration)) {
		query["MinDuration"] = request.MinDuration
	}

	if !tea.BoolValue(util.IsUnset(request.OperationName)) {
		query["OperationName"] = request.OperationName
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Reverse)) {
		query["Reverse"] = request.Reverse
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceIp)) {
		query["ServiceIp"] = request.ServiceIp
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchTraces"),
		Version:     tea.String("2019-02-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchTracesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchTraces(request *SearchTracesRequest) (_result *SearchTracesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchTracesResponse{}
	_body, _err := client.SearchTracesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
