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

type AddIpRequest struct {
	IpList   *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
	PackId   *string `json:"PackId,omitempty" xml:"PackId,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s AddIpRequest) String() string {
	return tea.Prettify(s)
}

func (s AddIpRequest) GoString() string {
	return s.String()
}

func (s *AddIpRequest) SetIpList(v string) *AddIpRequest {
	s.IpList = &v
	return s
}

func (s *AddIpRequest) SetPackId(v string) *AddIpRequest {
	s.PackId = &v
	return s
}

func (s *AddIpRequest) SetSourceIp(v string) *AddIpRequest {
	s.SourceIp = &v
	return s
}

type AddIpResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddIpResponseBody) GoString() string {
	return s.String()
}

func (s *AddIpResponseBody) SetRequestId(v string) *AddIpResponseBody {
	s.RequestId = &v
	return s
}

type AddIpResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddIpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddIpResponse) String() string {
	return tea.Prettify(s)
}

func (s AddIpResponse) GoString() string {
	return s.String()
}

func (s *AddIpResponse) SetHeaders(v map[string]*string) *AddIpResponse {
	s.Headers = v
	return s
}

func (s *AddIpResponse) SetStatusCode(v int32) *AddIpResponse {
	s.StatusCode = &v
	return s
}

func (s *AddIpResponse) SetBody(v *AddIpResponseBody) *AddIpResponse {
	s.Body = v
	return s
}

type AddProductRequest struct {
	PackId   *string `json:"PackId,omitempty" xml:"PackId,omitempty"`
	Product  *string `json:"Product,omitempty" xml:"Product,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s AddProductRequest) String() string {
	return tea.Prettify(s)
}

func (s AddProductRequest) GoString() string {
	return s.String()
}

func (s *AddProductRequest) SetPackId(v string) *AddProductRequest {
	s.PackId = &v
	return s
}

func (s *AddProductRequest) SetProduct(v string) *AddProductRequest {
	s.Product = &v
	return s
}

func (s *AddProductRequest) SetSourceIp(v string) *AddProductRequest {
	s.SourceIp = &v
	return s
}

type AddProductResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddProductResponseBody) GoString() string {
	return s.String()
}

func (s *AddProductResponseBody) SetRequestId(v string) *AddProductResponseBody {
	s.RequestId = &v
	return s
}

type AddProductResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddProductResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddProductResponse) String() string {
	return tea.Prettify(s)
}

func (s AddProductResponse) GoString() string {
	return s.String()
}

func (s *AddProductResponse) SetHeaders(v map[string]*string) *AddProductResponse {
	s.Headers = v
	return s
}

func (s *AddProductResponse) SetStatusCode(v int32) *AddProductResponse {
	s.StatusCode = &v
	return s
}

func (s *AddProductResponse) SetBody(v *AddProductResponseBody) *AddProductResponse {
	s.Body = v
	return s
}

type CheckGrantRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s CheckGrantRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckGrantRequest) GoString() string {
	return s.String()
}

func (s *CheckGrantRequest) SetSourceIp(v string) *CheckGrantRequest {
	s.SourceIp = &v
	return s
}

type CheckGrantResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CheckGrantResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckGrantResponseBody) GoString() string {
	return s.String()
}

func (s *CheckGrantResponseBody) SetRequestId(v string) *CheckGrantResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckGrantResponseBody) SetStatus(v int32) *CheckGrantResponseBody {
	s.Status = &v
	return s
}

type CheckGrantResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckGrantResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckGrantResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckGrantResponse) GoString() string {
	return s.String()
}

func (s *CheckGrantResponse) SetHeaders(v map[string]*string) *CheckGrantResponse {
	s.Headers = v
	return s
}

func (s *CheckGrantResponse) SetStatusCode(v int32) *CheckGrantResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckGrantResponse) SetBody(v *CheckGrantResponseBody) *CheckGrantResponse {
	s.Body = v
	return s
}

type DeleteBlackholeRequest struct {
	Ip       *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	PackId   *string `json:"PackId,omitempty" xml:"PackId,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DeleteBlackholeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBlackholeRequest) GoString() string {
	return s.String()
}

func (s *DeleteBlackholeRequest) SetIp(v string) *DeleteBlackholeRequest {
	s.Ip = &v
	return s
}

func (s *DeleteBlackholeRequest) SetPackId(v string) *DeleteBlackholeRequest {
	s.PackId = &v
	return s
}

func (s *DeleteBlackholeRequest) SetSourceIp(v string) *DeleteBlackholeRequest {
	s.SourceIp = &v
	return s
}

type DeleteBlackholeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBlackholeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBlackholeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBlackholeResponseBody) SetRequestId(v string) *DeleteBlackholeResponseBody {
	s.RequestId = &v
	return s
}

type DeleteBlackholeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteBlackholeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteBlackholeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBlackholeResponse) GoString() string {
	return s.String()
}

func (s *DeleteBlackholeResponse) SetHeaders(v map[string]*string) *DeleteBlackholeResponse {
	s.Headers = v
	return s
}

func (s *DeleteBlackholeResponse) SetStatusCode(v int32) *DeleteBlackholeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBlackholeResponse) SetBody(v *DeleteBlackholeResponseBody) *DeleteBlackholeResponse {
	s.Body = v
	return s
}

type DeleteIpRequest struct {
	IpList   *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
	PackId   *string `json:"PackId,omitempty" xml:"PackId,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DeleteIpRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpRequest) GoString() string {
	return s.String()
}

func (s *DeleteIpRequest) SetIpList(v string) *DeleteIpRequest {
	s.IpList = &v
	return s
}

func (s *DeleteIpRequest) SetPackId(v string) *DeleteIpRequest {
	s.PackId = &v
	return s
}

func (s *DeleteIpRequest) SetSourceIp(v string) *DeleteIpRequest {
	s.SourceIp = &v
	return s
}

type DeleteIpResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIpResponseBody) SetRequestId(v string) *DeleteIpResponseBody {
	s.RequestId = &v
	return s
}

type DeleteIpResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteIpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteIpResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpResponse) GoString() string {
	return s.String()
}

func (s *DeleteIpResponse) SetHeaders(v map[string]*string) *DeleteIpResponse {
	s.Headers = v
	return s
}

func (s *DeleteIpResponse) SetStatusCode(v int32) *DeleteIpResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIpResponse) SetBody(v *DeleteIpResponseBody) *DeleteIpResponse {
	s.Body = v
	return s
}

type DeleteProductRequest struct {
	PackId   *string `json:"PackId,omitempty" xml:"PackId,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DeleteProductRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProductRequest) GoString() string {
	return s.String()
}

func (s *DeleteProductRequest) SetPackId(v string) *DeleteProductRequest {
	s.PackId = &v
	return s
}

func (s *DeleteProductRequest) SetSourceIp(v string) *DeleteProductRequest {
	s.SourceIp = &v
	return s
}

type DeleteProductResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProductResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProductResponseBody) SetRequestId(v string) *DeleteProductResponseBody {
	s.RequestId = &v
	return s
}

type DeleteProductResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteProductResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteProductResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProductResponse) GoString() string {
	return s.String()
}

func (s *DeleteProductResponse) SetHeaders(v map[string]*string) *DeleteProductResponse {
	s.Headers = v
	return s
}

func (s *DeleteProductResponse) SetStatusCode(v int32) *DeleteProductResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProductResponse) SetBody(v *DeleteProductResponseBody) *DeleteProductResponse {
	s.Body = v
	return s
}

type DescribeDdosEventRequest struct {
	EndTime   *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PackId    *string `json:"PackId,omitempty" xml:"PackId,omitempty"`
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	StartTime *int32  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDdosEventRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosEventRequest) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventRequest) SetEndTime(v int32) *DescribeDdosEventRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDdosEventRequest) SetPackId(v string) *DescribeDdosEventRequest {
	s.PackId = &v
	return s
}

func (s *DescribeDdosEventRequest) SetPageNo(v int32) *DescribeDdosEventRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeDdosEventRequest) SetPageSize(v int32) *DescribeDdosEventRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDdosEventRequest) SetSourceIp(v string) *DescribeDdosEventRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDdosEventRequest) SetStartTime(v int32) *DescribeDdosEventRequest {
	s.StartTime = &v
	return s
}

type DescribeDdosEventResponseBody struct {
	Events    []*DescribeDdosEventResponseBodyEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *int64                                 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDdosEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosEventResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventResponseBody) SetEvents(v []*DescribeDdosEventResponseBodyEvents) *DescribeDdosEventResponseBody {
	s.Events = v
	return s
}

func (s *DescribeDdosEventResponseBody) SetRequestId(v string) *DescribeDdosEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDdosEventResponseBody) SetTotal(v int64) *DescribeDdosEventResponseBody {
	s.Total = &v
	return s
}

type DescribeDdosEventResponseBodyEvents struct {
	EndTime   *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Ip        *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Mbps      *int32  `json:"Mbps,omitempty" xml:"Mbps,omitempty"`
	Pps       *int32  `json:"Pps,omitempty" xml:"Pps,omitempty"`
	StartTime *int32  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDdosEventResponseBodyEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosEventResponseBodyEvents) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventResponseBodyEvents) SetEndTime(v int32) *DescribeDdosEventResponseBodyEvents {
	s.EndTime = &v
	return s
}

func (s *DescribeDdosEventResponseBodyEvents) SetIp(v string) *DescribeDdosEventResponseBodyEvents {
	s.Ip = &v
	return s
}

func (s *DescribeDdosEventResponseBodyEvents) SetMbps(v int32) *DescribeDdosEventResponseBodyEvents {
	s.Mbps = &v
	return s
}

func (s *DescribeDdosEventResponseBodyEvents) SetPps(v int32) *DescribeDdosEventResponseBodyEvents {
	s.Pps = &v
	return s
}

func (s *DescribeDdosEventResponseBodyEvents) SetStartTime(v int32) *DescribeDdosEventResponseBodyEvents {
	s.StartTime = &v
	return s
}

func (s *DescribeDdosEventResponseBodyEvents) SetStatus(v string) *DescribeDdosEventResponseBodyEvents {
	s.Status = &v
	return s
}

type DescribeDdosEventResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDdosEventResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDdosEventResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosEventResponse) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventResponse) SetHeaders(v map[string]*string) *DescribeDdosEventResponse {
	s.Headers = v
	return s
}

func (s *DescribeDdosEventResponse) SetStatusCode(v int32) *DescribeDdosEventResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDdosEventResponse) SetBody(v *DescribeDdosEventResponseBody) *DescribeDdosEventResponse {
	s.Body = v
	return s
}

type DescribeInstanceListRequest struct {
	InstanceIdList *string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty"`
	PackIdList     *string `json:"PackIdList,omitempty" xml:"PackIdList,omitempty"`
	PageNo         *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SourceIp       *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeInstanceListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceListRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceListRequest) SetInstanceIdList(v string) *DescribeInstanceListRequest {
	s.InstanceIdList = &v
	return s
}

func (s *DescribeInstanceListRequest) SetPackIdList(v string) *DescribeInstanceListRequest {
	s.PackIdList = &v
	return s
}

func (s *DescribeInstanceListRequest) SetPageNo(v int32) *DescribeInstanceListRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeInstanceListRequest) SetPageSize(v int32) *DescribeInstanceListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceListRequest) SetSourceIp(v string) *DescribeInstanceListRequest {
	s.SourceIp = &v
	return s
}

type DescribeInstanceListResponseBody struct {
	InstanceList []*DescribeInstanceListResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
	RequestId    *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total        *int64                                          `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeInstanceListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceListResponseBody) SetInstanceList(v []*DescribeInstanceListResponseBodyInstanceList) *DescribeInstanceListResponseBody {
	s.InstanceList = v
	return s
}

func (s *DescribeInstanceListResponseBody) SetRequestId(v string) *DescribeInstanceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceListResponseBody) SetTotal(v int64) *DescribeInstanceListResponseBody {
	s.Total = &v
	return s
}

type DescribeInstanceListResponseBodyInstanceList struct {
	ExpireTime *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	GmtCreate  *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PackId     *string `json:"PackId,omitempty" xml:"PackId,omitempty"`
	Product    *string `json:"Product,omitempty" xml:"Product,omitempty"`
	Remark     *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeInstanceListResponseBodyInstanceList) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceListResponseBodyInstanceList) GoString() string {
	return s.String()
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetExpireTime(v int64) *DescribeInstanceListResponseBodyInstanceList {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetGmtCreate(v int64) *DescribeInstanceListResponseBodyInstanceList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetInstanceId(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetPackId(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.PackId = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetProduct(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.Product = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetRemark(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.Remark = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetStatus(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.Status = &v
	return s
}

type DescribeInstanceListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInstanceListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceListResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceListResponse) SetHeaders(v map[string]*string) *DescribeInstanceListResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceListResponse) SetStatusCode(v int32) *DescribeInstanceListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceListResponse) SetBody(v *DescribeInstanceListResponseBody) *DescribeInstanceListResponse {
	s.Body = v
	return s
}

type DescribeOnDemandInstanceRequest struct {
	// The page number of the page to return.
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page. Maximum value: **50**.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the on-demand instance that you want to query.
	//
	// >  You can call the [DescribeRegions](https://www.alibabacloud.com/help/en/ddos-protection/latest/instances-describeregions) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeOnDemandInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOnDemandInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeOnDemandInstanceRequest) SetPageNo(v int32) *DescribeOnDemandInstanceRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeOnDemandInstanceRequest) SetPageSize(v int32) *DescribeOnDemandInstanceRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeOnDemandInstanceRequest) SetRegionId(v string) *DescribeOnDemandInstanceRequest {
	s.RegionId = &v
	return s
}

type DescribeOnDemandInstanceResponseBody struct {
	// The details of the on-demand instance.
	Instances []*DescribeOnDemandInstanceResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries that were returned.
	Total *string `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeOnDemandInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOnDemandInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOnDemandInstanceResponseBody) SetInstances(v []*DescribeOnDemandInstanceResponseBodyInstances) *DescribeOnDemandInstanceResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeOnDemandInstanceResponseBody) SetRequestId(v string) *DescribeOnDemandInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOnDemandInstanceResponseBody) SetTotal(v string) *DescribeOnDemandInstanceResponseBody {
	s.Total = &v
	return s
}

type DescribeOnDemandInstanceResponseBodyInstances struct {
	// The protection status of the on-demand instance. Valid values:
	//
	// - **Defense**: The on-demand instance is protecting your assets, which indicates that traffic is routed to the on-demand instance.
	// - **UnDefense**: The on-demand instance does not protect your assets.
	DefenseStatus *string `json:"DefenseStatus,omitempty" xml:"DefenseStatus,omitempty"`
	// The ID of the on-demand instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The CIDR block of the on-demand instance.
	Ipnet []*string `json:"Ipnet,omitempty" xml:"Ipnet,omitempty" type:"Repeated"`
	// The region ID of the on-demand instance.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The description of the on-demand instance.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s DescribeOnDemandInstanceResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeOnDemandInstanceResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeOnDemandInstanceResponseBodyInstances) SetDefenseStatus(v string) *DescribeOnDemandInstanceResponseBodyInstances {
	s.DefenseStatus = &v
	return s
}

func (s *DescribeOnDemandInstanceResponseBodyInstances) SetInstanceId(v string) *DescribeOnDemandInstanceResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeOnDemandInstanceResponseBodyInstances) SetIpnet(v []*string) *DescribeOnDemandInstanceResponseBodyInstances {
	s.Ipnet = v
	return s
}

func (s *DescribeOnDemandInstanceResponseBodyInstances) SetRegionId(v string) *DescribeOnDemandInstanceResponseBodyInstances {
	s.RegionId = &v
	return s
}

func (s *DescribeOnDemandInstanceResponseBodyInstances) SetRemark(v string) *DescribeOnDemandInstanceResponseBodyInstances {
	s.Remark = &v
	return s
}

type DescribeOnDemandInstanceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeOnDemandInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeOnDemandInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOnDemandInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeOnDemandInstanceResponse) SetHeaders(v map[string]*string) *DescribeOnDemandInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeOnDemandInstanceResponse) SetStatusCode(v int32) *DescribeOnDemandInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOnDemandInstanceResponse) SetBody(v *DescribeOnDemandInstanceResponseBody) *DescribeOnDemandInstanceResponse {
	s.Body = v
	return s
}

type DescribeOpEntitiesRequest struct {
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeOpEntitiesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpEntitiesRequest) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesRequest) SetCurrentPage(v int32) *DescribeOpEntitiesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetEndTime(v int64) *DescribeOpEntitiesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetLang(v string) *DescribeOpEntitiesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetPageSize(v int32) *DescribeOpEntitiesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetSourceIp(v string) *DescribeOpEntitiesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetStartTime(v int64) *DescribeOpEntitiesRequest {
	s.StartTime = &v
	return s
}

type DescribeOpEntitiesResponseBody struct {
	OpEntities []*DescribeOpEntitiesResponseBodyOpEntities `json:"OpEntities,omitempty" xml:"OpEntities,omitempty" type:"Repeated"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeOpEntitiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpEntitiesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesResponseBody) SetOpEntities(v []*DescribeOpEntitiesResponseBodyOpEntities) *DescribeOpEntitiesResponseBody {
	s.OpEntities = v
	return s
}

func (s *DescribeOpEntitiesResponseBody) SetRequestId(v string) *DescribeOpEntitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOpEntitiesResponseBody) SetTotalCount(v int32) *DescribeOpEntitiesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeOpEntitiesResponseBodyOpEntities struct {
	EntityObject *string `json:"EntityObject,omitempty" xml:"EntityObject,omitempty"`
	EntityType   *int32  `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	GmtCreate    *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	OpAccount    *string `json:"OpAccount,omitempty" xml:"OpAccount,omitempty"`
	OpAction     *int32  `json:"OpAction,omitempty" xml:"OpAction,omitempty"`
	OpDesc       *string `json:"OpDesc,omitempty" xml:"OpDesc,omitempty"`
}

func (s DescribeOpEntitiesResponseBodyOpEntities) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpEntitiesResponseBodyOpEntities) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetEntityObject(v string) *DescribeOpEntitiesResponseBodyOpEntities {
	s.EntityObject = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetEntityType(v int32) *DescribeOpEntitiesResponseBodyOpEntities {
	s.EntityType = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetGmtCreate(v int64) *DescribeOpEntitiesResponseBodyOpEntities {
	s.GmtCreate = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetOpAccount(v string) *DescribeOpEntitiesResponseBodyOpEntities {
	s.OpAccount = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetOpAction(v int32) *DescribeOpEntitiesResponseBodyOpEntities {
	s.OpAction = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetOpDesc(v string) *DescribeOpEntitiesResponseBodyOpEntities {
	s.OpDesc = &v
	return s
}

type DescribeOpEntitiesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeOpEntitiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeOpEntitiesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpEntitiesResponse) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesResponse) SetHeaders(v map[string]*string) *DescribeOpEntitiesResponse {
	s.Headers = v
	return s
}

func (s *DescribeOpEntitiesResponse) SetStatusCode(v int32) *DescribeOpEntitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOpEntitiesResponse) SetBody(v *DescribeOpEntitiesResponseBody) *DescribeOpEntitiesResponse {
	s.Body = v
	return s
}

type DescribePackRequest struct {
	PackId   *string `json:"PackId,omitempty" xml:"PackId,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribePackRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePackRequest) GoString() string {
	return s.String()
}

func (s *DescribePackRequest) SetPackId(v string) *DescribePackRequest {
	s.PackId = &v
	return s
}

func (s *DescribePackRequest) SetSourceIp(v string) *DescribePackRequest {
	s.SourceIp = &v
	return s
}

type DescribePackResponseBody struct {
	PackInfo  *DescribePackResponseBodyPackInfo `json:"PackInfo,omitempty" xml:"PackInfo,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePackResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePackResponseBody) SetPackInfo(v *DescribePackResponseBodyPackInfo) *DescribePackResponseBody {
	s.PackInfo = v
	return s
}

func (s *DescribePackResponseBody) SetRequestId(v string) *DescribePackResponseBody {
	s.RequestId = &v
	return s
}

type DescribePackResponseBodyPackInfo struct {
	AvailableDeleteBlackholeCount *int32                                      `json:"AvailableDeleteBlackholeCount,omitempty" xml:"AvailableDeleteBlackholeCount,omitempty"`
	IpList                        []*DescribePackResponseBodyPackInfoIpList   `json:"IpList,omitempty" xml:"IpList,omitempty" type:"Repeated"`
	PackConfig                    *DescribePackResponseBodyPackInfoPackConfig `json:"PackConfig,omitempty" xml:"PackConfig,omitempty" type:"Struct"`
	Region                        *string                                     `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DescribePackResponseBodyPackInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribePackResponseBodyPackInfo) GoString() string {
	return s.String()
}

func (s *DescribePackResponseBodyPackInfo) SetAvailableDeleteBlackholeCount(v int32) *DescribePackResponseBodyPackInfo {
	s.AvailableDeleteBlackholeCount = &v
	return s
}

func (s *DescribePackResponseBodyPackInfo) SetIpList(v []*DescribePackResponseBodyPackInfoIpList) *DescribePackResponseBodyPackInfo {
	s.IpList = v
	return s
}

func (s *DescribePackResponseBodyPackInfo) SetPackConfig(v *DescribePackResponseBodyPackInfoPackConfig) *DescribePackResponseBodyPackInfo {
	s.PackConfig = v
	return s
}

func (s *DescribePackResponseBodyPackInfo) SetRegion(v string) *DescribePackResponseBodyPackInfo {
	s.Region = &v
	return s
}

type DescribePackResponseBodyPackInfoIpList struct {
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribePackResponseBodyPackInfoIpList) String() string {
	return tea.Prettify(s)
}

func (s DescribePackResponseBodyPackInfoIpList) GoString() string {
	return s.String()
}

func (s *DescribePackResponseBodyPackInfoIpList) SetIp(v string) *DescribePackResponseBodyPackInfoIpList {
	s.Ip = &v
	return s
}

type DescribePackResponseBodyPackInfoPackConfig struct {
	IpAdvanceThre *int32 `json:"IpAdvanceThre,omitempty" xml:"IpAdvanceThre,omitempty"`
	IpBasicThre   *int32 `json:"IpBasicThre,omitempty" xml:"IpBasicThre,omitempty"`
	IpSpec        *int32 `json:"IpSpec,omitempty" xml:"IpSpec,omitempty"`
	PackAdvThre   *int32 `json:"PackAdvThre,omitempty" xml:"PackAdvThre,omitempty"`
	PackBasicThre *int32 `json:"PackBasicThre,omitempty" xml:"PackBasicThre,omitempty"`
}

func (s DescribePackResponseBodyPackInfoPackConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribePackResponseBodyPackInfoPackConfig) GoString() string {
	return s.String()
}

func (s *DescribePackResponseBodyPackInfoPackConfig) SetIpAdvanceThre(v int32) *DescribePackResponseBodyPackInfoPackConfig {
	s.IpAdvanceThre = &v
	return s
}

func (s *DescribePackResponseBodyPackInfoPackConfig) SetIpBasicThre(v int32) *DescribePackResponseBodyPackInfoPackConfig {
	s.IpBasicThre = &v
	return s
}

func (s *DescribePackResponseBodyPackInfoPackConfig) SetIpSpec(v int32) *DescribePackResponseBodyPackInfoPackConfig {
	s.IpSpec = &v
	return s
}

func (s *DescribePackResponseBodyPackInfoPackConfig) SetPackAdvThre(v int32) *DescribePackResponseBodyPackInfoPackConfig {
	s.PackAdvThre = &v
	return s
}

func (s *DescribePackResponseBodyPackInfoPackConfig) SetPackBasicThre(v int32) *DescribePackResponseBodyPackInfoPackConfig {
	s.PackBasicThre = &v
	return s
}

type DescribePackResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePackResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePackResponse) GoString() string {
	return s.String()
}

func (s *DescribePackResponse) SetHeaders(v map[string]*string) *DescribePackResponse {
	s.Headers = v
	return s
}

func (s *DescribePackResponse) SetStatusCode(v int32) *DescribePackResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePackResponse) SetBody(v *DescribePackResponseBody) *DescribePackResponse {
	s.Body = v
	return s
}

type DescribePackListRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribePackListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePackListRequest) GoString() string {
	return s.String()
}

func (s *DescribePackListRequest) SetRegionId(v string) *DescribePackListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePackListRequest) SetResourceGroupId(v string) *DescribePackListRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribePackListResponseBody struct {
	PackList  []*DescribePackListResponseBodyPackList `json:"PackList,omitempty" xml:"PackList,omitempty" type:"Repeated"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePackListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePackListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePackListResponseBody) SetPackList(v []*DescribePackListResponseBodyPackList) *DescribePackListResponseBody {
	s.PackList = v
	return s
}

func (s *DescribePackListResponseBody) SetRequestId(v string) *DescribePackListResponseBody {
	s.RequestId = &v
	return s
}

type DescribePackListResponseBodyPackList struct {
	AvailableDeleteBlackholeCount *int32                                          `json:"AvailableDeleteBlackholeCount,omitempty" xml:"AvailableDeleteBlackholeCount,omitempty"`
	PackConfig                    *DescribePackListResponseBodyPackListPackConfig `json:"PackConfig,omitempty" xml:"PackConfig,omitempty" type:"Struct"`
	PackId                        *string                                         `json:"PackId,omitempty" xml:"PackId,omitempty"`
	Region                        *string                                         `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DescribePackListResponseBodyPackList) String() string {
	return tea.Prettify(s)
}

func (s DescribePackListResponseBodyPackList) GoString() string {
	return s.String()
}

func (s *DescribePackListResponseBodyPackList) SetAvailableDeleteBlackholeCount(v int32) *DescribePackListResponseBodyPackList {
	s.AvailableDeleteBlackholeCount = &v
	return s
}

func (s *DescribePackListResponseBodyPackList) SetPackConfig(v *DescribePackListResponseBodyPackListPackConfig) *DescribePackListResponseBodyPackList {
	s.PackConfig = v
	return s
}

func (s *DescribePackListResponseBodyPackList) SetPackId(v string) *DescribePackListResponseBodyPackList {
	s.PackId = &v
	return s
}

func (s *DescribePackListResponseBodyPackList) SetRegion(v string) *DescribePackListResponseBodyPackList {
	s.Region = &v
	return s
}

type DescribePackListResponseBodyPackListPackConfig struct {
	IpAdvanceThre *int32 `json:"IpAdvanceThre,omitempty" xml:"IpAdvanceThre,omitempty"`
	IpBasicThre   *int32 `json:"IpBasicThre,omitempty" xml:"IpBasicThre,omitempty"`
	IpSpec        *int32 `json:"IpSpec,omitempty" xml:"IpSpec,omitempty"`
	PackAdvThre   *int32 `json:"PackAdvThre,omitempty" xml:"PackAdvThre,omitempty"`
	PackBasicThre *int32 `json:"PackBasicThre,omitempty" xml:"PackBasicThre,omitempty"`
}

func (s DescribePackListResponseBodyPackListPackConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribePackListResponseBodyPackListPackConfig) GoString() string {
	return s.String()
}

func (s *DescribePackListResponseBodyPackListPackConfig) SetIpAdvanceThre(v int32) *DescribePackListResponseBodyPackListPackConfig {
	s.IpAdvanceThre = &v
	return s
}

func (s *DescribePackListResponseBodyPackListPackConfig) SetIpBasicThre(v int32) *DescribePackListResponseBodyPackListPackConfig {
	s.IpBasicThre = &v
	return s
}

func (s *DescribePackListResponseBodyPackListPackConfig) SetIpSpec(v int32) *DescribePackListResponseBodyPackListPackConfig {
	s.IpSpec = &v
	return s
}

func (s *DescribePackListResponseBodyPackListPackConfig) SetPackAdvThre(v int32) *DescribePackListResponseBodyPackListPackConfig {
	s.PackAdvThre = &v
	return s
}

func (s *DescribePackListResponseBodyPackListPackConfig) SetPackBasicThre(v int32) *DescribePackListResponseBodyPackListPackConfig {
	s.PackBasicThre = &v
	return s
}

type DescribePackListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePackListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePackListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePackListResponse) GoString() string {
	return s.String()
}

func (s *DescribePackListResponse) SetHeaders(v map[string]*string) *DescribePackListResponse {
	s.Headers = v
	return s
}

func (s *DescribePackListResponse) SetStatusCode(v int32) *DescribePackListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePackListResponse) SetBody(v *DescribePackListResponseBody) *DescribePackListResponse {
	s.Body = v
	return s
}

type DescribePackPaidTrafficRequest struct {
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribePackPaidTrafficRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePackPaidTrafficRequest) GoString() string {
	return s.String()
}

func (s *DescribePackPaidTrafficRequest) SetCurrentPage(v int32) *DescribePackPaidTrafficRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribePackPaidTrafficRequest) SetEndTime(v int64) *DescribePackPaidTrafficRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePackPaidTrafficRequest) SetInstanceId(v string) *DescribePackPaidTrafficRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribePackPaidTrafficRequest) SetPageSize(v int32) *DescribePackPaidTrafficRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePackPaidTrafficRequest) SetSourceIp(v string) *DescribePackPaidTrafficRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribePackPaidTrafficRequest) SetStartTime(v int64) *DescribePackPaidTrafficRequest {
	s.StartTime = &v
	return s
}

type DescribePackPaidTrafficResponseBody struct {
	PackPaidTraffics []*DescribePackPaidTrafficResponseBodyPackPaidTraffics `json:"PackPaidTraffics,omitempty" xml:"PackPaidTraffics,omitempty" type:"Repeated"`
	RequestId        *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount       *int32                                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePackPaidTrafficResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePackPaidTrafficResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePackPaidTrafficResponseBody) SetPackPaidTraffics(v []*DescribePackPaidTrafficResponseBodyPackPaidTraffics) *DescribePackPaidTrafficResponseBody {
	s.PackPaidTraffics = v
	return s
}

func (s *DescribePackPaidTrafficResponseBody) SetRequestId(v string) *DescribePackPaidTrafficResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePackPaidTrafficResponseBody) SetTotalCount(v int32) *DescribePackPaidTrafficResponseBody {
	s.TotalCount = &v
	return s
}

type DescribePackPaidTrafficResponseBodyPackPaidTraffics struct {
	BaseBandwidth    *int32   `json:"BaseBandwidth,omitempty" xml:"BaseBandwidth,omitempty"`
	ElasticBandwidth *int32   `json:"ElasticBandwidth,omitempty" xml:"ElasticBandwidth,omitempty"`
	InstanceId       *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxAttack        *float32 `json:"MaxAttack,omitempty" xml:"MaxAttack,omitempty"`
	PaidCapacity     *float32 `json:"PaidCapacity,omitempty" xml:"PaidCapacity,omitempty"`
	StartTime        *int64   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TotalCapacity    *float32 `json:"TotalCapacity,omitempty" xml:"TotalCapacity,omitempty"`
}

func (s DescribePackPaidTrafficResponseBodyPackPaidTraffics) String() string {
	return tea.Prettify(s)
}

func (s DescribePackPaidTrafficResponseBodyPackPaidTraffics) GoString() string {
	return s.String()
}

func (s *DescribePackPaidTrafficResponseBodyPackPaidTraffics) SetBaseBandwidth(v int32) *DescribePackPaidTrafficResponseBodyPackPaidTraffics {
	s.BaseBandwidth = &v
	return s
}

func (s *DescribePackPaidTrafficResponseBodyPackPaidTraffics) SetElasticBandwidth(v int32) *DescribePackPaidTrafficResponseBodyPackPaidTraffics {
	s.ElasticBandwidth = &v
	return s
}

func (s *DescribePackPaidTrafficResponseBodyPackPaidTraffics) SetInstanceId(v string) *DescribePackPaidTrafficResponseBodyPackPaidTraffics {
	s.InstanceId = &v
	return s
}

func (s *DescribePackPaidTrafficResponseBodyPackPaidTraffics) SetMaxAttack(v float32) *DescribePackPaidTrafficResponseBodyPackPaidTraffics {
	s.MaxAttack = &v
	return s
}

func (s *DescribePackPaidTrafficResponseBodyPackPaidTraffics) SetPaidCapacity(v float32) *DescribePackPaidTrafficResponseBodyPackPaidTraffics {
	s.PaidCapacity = &v
	return s
}

func (s *DescribePackPaidTrafficResponseBodyPackPaidTraffics) SetStartTime(v int64) *DescribePackPaidTrafficResponseBodyPackPaidTraffics {
	s.StartTime = &v
	return s
}

func (s *DescribePackPaidTrafficResponseBodyPackPaidTraffics) SetTotalCapacity(v float32) *DescribePackPaidTrafficResponseBodyPackPaidTraffics {
	s.TotalCapacity = &v
	return s
}

type DescribePackPaidTrafficResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePackPaidTrafficResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePackPaidTrafficResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePackPaidTrafficResponse) GoString() string {
	return s.String()
}

func (s *DescribePackPaidTrafficResponse) SetHeaders(v map[string]*string) *DescribePackPaidTrafficResponse {
	s.Headers = v
	return s
}

func (s *DescribePackPaidTrafficResponse) SetStatusCode(v int32) *DescribePackPaidTrafficResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePackPaidTrafficResponse) SetBody(v *DescribePackPaidTrafficResponseBody) *DescribePackPaidTrafficResponse {
	s.Body = v
	return s
}

type DescribeResourcePackInstancesRequest struct {
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeResourcePackInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackInstancesRequest) SetCurrentPage(v int32) *DescribeResourcePackInstancesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeResourcePackInstancesRequest) SetPageSize(v int32) *DescribeResourcePackInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeResourcePackInstancesRequest) SetSourceIp(v string) *DescribeResourcePackInstancesRequest {
	s.SourceIp = &v
	return s
}

type DescribeResourcePackInstancesResponseBody struct {
	RequestId     *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourcePacks []*DescribeResourcePackInstancesResponseBodyResourcePacks `json:"ResourcePacks,omitempty" xml:"ResourcePacks,omitempty" type:"Repeated"`
	TotalCount    *int32                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeResourcePackInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackInstancesResponseBody) SetRequestId(v string) *DescribeResourcePackInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourcePackInstancesResponseBody) SetResourcePacks(v []*DescribeResourcePackInstancesResponseBodyResourcePacks) *DescribeResourcePackInstancesResponseBody {
	s.ResourcePacks = v
	return s
}

func (s *DescribeResourcePackInstancesResponseBody) SetTotalCount(v int32) *DescribeResourcePackInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeResourcePackInstancesResponseBodyResourcePacks struct {
	CurrCapacity   *int64  `json:"CurrCapacity,omitempty" xml:"CurrCapacity,omitempty"`
	EndTime        *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InitCapacity   *int64  `json:"InitCapacity,omitempty" xml:"InitCapacity,omitempty"`
	ResourcePackId *string `json:"ResourcePackId,omitempty" xml:"ResourcePackId,omitempty"`
	StartTime      *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeResourcePackInstancesResponseBodyResourcePacks) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackInstancesResponseBodyResourcePacks) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackInstancesResponseBodyResourcePacks) SetCurrCapacity(v int64) *DescribeResourcePackInstancesResponseBodyResourcePacks {
	s.CurrCapacity = &v
	return s
}

func (s *DescribeResourcePackInstancesResponseBodyResourcePacks) SetEndTime(v int64) *DescribeResourcePackInstancesResponseBodyResourcePacks {
	s.EndTime = &v
	return s
}

func (s *DescribeResourcePackInstancesResponseBodyResourcePacks) SetInitCapacity(v int64) *DescribeResourcePackInstancesResponseBodyResourcePacks {
	s.InitCapacity = &v
	return s
}

func (s *DescribeResourcePackInstancesResponseBodyResourcePacks) SetResourcePackId(v string) *DescribeResourcePackInstancesResponseBodyResourcePacks {
	s.ResourcePackId = &v
	return s
}

func (s *DescribeResourcePackInstancesResponseBodyResourcePacks) SetStartTime(v int64) *DescribeResourcePackInstancesResponseBodyResourcePacks {
	s.StartTime = &v
	return s
}

func (s *DescribeResourcePackInstancesResponseBodyResourcePacks) SetStatus(v string) *DescribeResourcePackInstancesResponseBodyResourcePacks {
	s.Status = &v
	return s
}

type DescribeResourcePackInstancesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeResourcePackInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeResourcePackInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackInstancesResponse) SetHeaders(v map[string]*string) *DescribeResourcePackInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourcePackInstancesResponse) SetStatusCode(v int32) *DescribeResourcePackInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourcePackInstancesResponse) SetBody(v *DescribeResourcePackInstancesResponseBody) *DescribeResourcePackInstancesResponse {
	s.Body = v
	return s
}

type DescribeResourcePackStatisticsRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeResourcePackStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackStatisticsRequest) SetSourceIp(v string) *DescribeResourcePackStatisticsRequest {
	s.SourceIp = &v
	return s
}

type DescribeResourcePackStatisticsResponseBody struct {
	AvailablePackNum  *int32  `json:"AvailablePackNum,omitempty" xml:"AvailablePackNum,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCurrCapacity *int64  `json:"TotalCurrCapacity,omitempty" xml:"TotalCurrCapacity,omitempty"`
	TotalInitCapacity *int64  `json:"TotalInitCapacity,omitempty" xml:"TotalInitCapacity,omitempty"`
	TotalUsedCapacity *int64  `json:"TotalUsedCapacity,omitempty" xml:"TotalUsedCapacity,omitempty"`
}

func (s DescribeResourcePackStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackStatisticsResponseBody) SetAvailablePackNum(v int32) *DescribeResourcePackStatisticsResponseBody {
	s.AvailablePackNum = &v
	return s
}

func (s *DescribeResourcePackStatisticsResponseBody) SetRequestId(v string) *DescribeResourcePackStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourcePackStatisticsResponseBody) SetTotalCurrCapacity(v int64) *DescribeResourcePackStatisticsResponseBody {
	s.TotalCurrCapacity = &v
	return s
}

func (s *DescribeResourcePackStatisticsResponseBody) SetTotalInitCapacity(v int64) *DescribeResourcePackStatisticsResponseBody {
	s.TotalInitCapacity = &v
	return s
}

func (s *DescribeResourcePackStatisticsResponseBody) SetTotalUsedCapacity(v int64) *DescribeResourcePackStatisticsResponseBody {
	s.TotalUsedCapacity = &v
	return s
}

type DescribeResourcePackStatisticsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeResourcePackStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeResourcePackStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackStatisticsResponse) SetHeaders(v map[string]*string) *DescribeResourcePackStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourcePackStatisticsResponse) SetStatusCode(v int32) *DescribeResourcePackStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourcePackStatisticsResponse) SetBody(v *DescribeResourcePackStatisticsResponseBody) *DescribeResourcePackStatisticsResponse {
	s.Body = v
	return s
}

type DescribeResourcePackUsageRequest struct {
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeResourcePackUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackUsageRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackUsageRequest) SetEndTime(v int64) *DescribeResourcePackUsageRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeResourcePackUsageRequest) SetSourceIp(v string) *DescribeResourcePackUsageRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeResourcePackUsageRequest) SetStartTime(v int64) *DescribeResourcePackUsageRequest {
	s.StartTime = &v
	return s
}

type DescribeResourcePackUsageResponseBody struct {
	EndTime    *int64                                             `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval   *int64                                             `json:"Interval,omitempty" xml:"Interval,omitempty"`
	PackUsages []*DescribeResourcePackUsageResponseBodyPackUsages `json:"PackUsages,omitempty" xml:"PackUsages,omitempty" type:"Repeated"`
	RequestId  *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime  *int64                                             `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeResourcePackUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackUsageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackUsageResponseBody) SetEndTime(v int64) *DescribeResourcePackUsageResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeResourcePackUsageResponseBody) SetInterval(v int64) *DescribeResourcePackUsageResponseBody {
	s.Interval = &v
	return s
}

func (s *DescribeResourcePackUsageResponseBody) SetPackUsages(v []*DescribeResourcePackUsageResponseBodyPackUsages) *DescribeResourcePackUsageResponseBody {
	s.PackUsages = v
	return s
}

func (s *DescribeResourcePackUsageResponseBody) SetRequestId(v string) *DescribeResourcePackUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourcePackUsageResponseBody) SetStartTime(v int64) *DescribeResourcePackUsageResponseBody {
	s.StartTime = &v
	return s
}

type DescribeResourcePackUsageResponseBodyPackUsages struct {
	Time    *int64   `json:"Time,omitempty" xml:"Time,omitempty"`
	Traffic *float32 `json:"Traffic,omitempty" xml:"Traffic,omitempty"`
}

func (s DescribeResourcePackUsageResponseBodyPackUsages) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackUsageResponseBodyPackUsages) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackUsageResponseBodyPackUsages) SetTime(v int64) *DescribeResourcePackUsageResponseBodyPackUsages {
	s.Time = &v
	return s
}

func (s *DescribeResourcePackUsageResponseBodyPackUsages) SetTraffic(v float32) *DescribeResourcePackUsageResponseBodyPackUsages {
	s.Traffic = &v
	return s
}

type DescribeResourcePackUsageResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeResourcePackUsageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeResourcePackUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackUsageResponse) SetHeaders(v map[string]*string) *DescribeResourcePackUsageResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourcePackUsageResponse) SetStatusCode(v int32) *DescribeResourcePackUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourcePackUsageResponse) SetBody(v *DescribeResourcePackUsageResponseBody) *DescribeResourcePackUsageResponse {
	s.Body = v
	return s
}

type DescribeTopTrafficRequest struct {
	// The end of the time range to query. This value is a UNIX timestamp. Unit: seconds.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the on-demand instance.
	//
	// >  You can call the [DescribeOnDemandInstance](~~152120~~) operation to query the IDs of all on-demand instances.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The CIDR block of the on-demand instance that you want to query.
	Ipnet *string `json:"Ipnet,omitempty" xml:"Ipnet,omitempty"`
	// The number of the page to return. Default value: **1**.
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page. Default value: **10**. Maximum value: **50**.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the on-demand instance.
	//
	// >  You can call the [DescribeRegions](~~118703~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the on-demand instance belongs in Resource Management.
	//
	// If you do not specify this parameter, the instance belongs to the default resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The number of IP addresses from which the most traffic is forwarded. Default value: **1**, which indicates the IP address from which the most traffic is forwarded.
	Rn *int32 `json:"Rn,omitempty" xml:"Rn,omitempty"`
	// The beginning of the time range to query. This value is a UNIX timestamp. Unit: seconds.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeTopTrafficRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopTrafficRequest) GoString() string {
	return s.String()
}

func (s *DescribeTopTrafficRequest) SetEndTime(v string) *DescribeTopTrafficRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeTopTrafficRequest) SetInstanceId(v string) *DescribeTopTrafficRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTopTrafficRequest) SetIpnet(v string) *DescribeTopTrafficRequest {
	s.Ipnet = &v
	return s
}

func (s *DescribeTopTrafficRequest) SetPageNo(v int32) *DescribeTopTrafficRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeTopTrafficRequest) SetPageSize(v int32) *DescribeTopTrafficRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTopTrafficRequest) SetRegionId(v string) *DescribeTopTrafficRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTopTrafficRequest) SetResourceGroupId(v string) *DescribeTopTrafficRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeTopTrafficRequest) SetRn(v int32) *DescribeTopTrafficRequest {
	s.Rn = &v
	return s
}

func (s *DescribeTopTrafficRequest) SetStartTime(v string) *DescribeTopTrafficRequest {
	s.StartTime = &v
	return s
}

type DescribeTopTrafficResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
	// The information about the traffic that is forwarded by the on-demand instance.
	TrafficList []*DescribeTopTrafficResponseBodyTrafficList `json:"TrafficList,omitempty" xml:"TrafficList,omitempty" type:"Repeated"`
}

func (s DescribeTopTrafficResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopTrafficResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTopTrafficResponseBody) SetRequestId(v string) *DescribeTopTrafficResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTopTrafficResponseBody) SetTotal(v int64) *DescribeTopTrafficResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeTopTrafficResponseBody) SetTrafficList(v []*DescribeTopTrafficResponseBodyTrafficList) *DescribeTopTrafficResponseBody {
	s.TrafficList = v
	return s
}

type DescribeTopTrafficResponseBodyTrafficList struct {
	// The attack traffic. Unit: Kbit/s.
	AttackBps *int32 `json:"AttackBps,omitempty" xml:"AttackBps,omitempty"`
	// The number of attack data packets. Unit: packets per second (pps).
	AttackPps *int32 `json:"AttackPps,omitempty" xml:"AttackPps,omitempty"`
	// The total traffic. Unit: Kbit/s.
	Bps *int32 `json:"Bps,omitempty" xml:"Bps,omitempty"`
	// The IP address from which the most traffic is forwarded by the on-demand instance.
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The total number of data packets. Unit: pps.
	Pps *int32 `json:"Pps,omitempty" xml:"Pps,omitempty"`
}

func (s DescribeTopTrafficResponseBodyTrafficList) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopTrafficResponseBodyTrafficList) GoString() string {
	return s.String()
}

func (s *DescribeTopTrafficResponseBodyTrafficList) SetAttackBps(v int32) *DescribeTopTrafficResponseBodyTrafficList {
	s.AttackBps = &v
	return s
}

func (s *DescribeTopTrafficResponseBodyTrafficList) SetAttackPps(v int32) *DescribeTopTrafficResponseBodyTrafficList {
	s.AttackPps = &v
	return s
}

func (s *DescribeTopTrafficResponseBodyTrafficList) SetBps(v int32) *DescribeTopTrafficResponseBodyTrafficList {
	s.Bps = &v
	return s
}

func (s *DescribeTopTrafficResponseBodyTrafficList) SetIp(v string) *DescribeTopTrafficResponseBodyTrafficList {
	s.Ip = &v
	return s
}

func (s *DescribeTopTrafficResponseBodyTrafficList) SetPps(v int32) *DescribeTopTrafficResponseBodyTrafficList {
	s.Pps = &v
	return s
}

type DescribeTopTrafficResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeTopTrafficResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTopTrafficResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopTrafficResponse) GoString() string {
	return s.String()
}

func (s *DescribeTopTrafficResponse) SetHeaders(v map[string]*string) *DescribeTopTrafficResponse {
	s.Headers = v
	return s
}

func (s *DescribeTopTrafficResponse) SetStatusCode(v int32) *DescribeTopTrafficResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTopTrafficResponse) SetBody(v *DescribeTopTrafficResponseBody) *DescribeTopTrafficResponse {
	s.Body = v
	return s
}

type DescribeTrafficRequest struct {
	EndTime   *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval  *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	StartTime *int32  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeTrafficRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrafficRequest) GoString() string {
	return s.String()
}

func (s *DescribeTrafficRequest) SetEndTime(v int32) *DescribeTrafficRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeTrafficRequest) SetInterval(v int32) *DescribeTrafficRequest {
	s.Interval = &v
	return s
}

func (s *DescribeTrafficRequest) SetName(v string) *DescribeTrafficRequest {
	s.Name = &v
	return s
}

func (s *DescribeTrafficRequest) SetSourceIp(v string) *DescribeTrafficRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeTrafficRequest) SetStartTime(v int32) *DescribeTrafficRequest {
	s.StartTime = &v
	return s
}

type DescribeTrafficResponseBody struct {
	FlowList  []*DescribeTrafficResponseBodyFlowList `json:"FlowList,omitempty" xml:"FlowList,omitempty" type:"Repeated"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTrafficResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrafficResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTrafficResponseBody) SetFlowList(v []*DescribeTrafficResponseBodyFlowList) *DescribeTrafficResponseBody {
	s.FlowList = v
	return s
}

func (s *DescribeTrafficResponseBody) SetRequestId(v string) *DescribeTrafficResponseBody {
	s.RequestId = &v
	return s
}

type DescribeTrafficResponseBodyFlowList struct {
	FlowType *string `json:"FlowType,omitempty" xml:"FlowType,omitempty"`
	Kbps     *int32  `json:"Kbps,omitempty" xml:"Kbps,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Pps      *int32  `json:"Pps,omitempty" xml:"Pps,omitempty"`
	Time     *int32  `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribeTrafficResponseBodyFlowList) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrafficResponseBodyFlowList) GoString() string {
	return s.String()
}

func (s *DescribeTrafficResponseBodyFlowList) SetFlowType(v string) *DescribeTrafficResponseBodyFlowList {
	s.FlowType = &v
	return s
}

func (s *DescribeTrafficResponseBodyFlowList) SetKbps(v int32) *DescribeTrafficResponseBodyFlowList {
	s.Kbps = &v
	return s
}

func (s *DescribeTrafficResponseBodyFlowList) SetName(v string) *DescribeTrafficResponseBodyFlowList {
	s.Name = &v
	return s
}

func (s *DescribeTrafficResponseBodyFlowList) SetPps(v int32) *DescribeTrafficResponseBodyFlowList {
	s.Pps = &v
	return s
}

func (s *DescribeTrafficResponseBodyFlowList) SetTime(v int32) *DescribeTrafficResponseBodyFlowList {
	s.Time = &v
	return s
}

type DescribeTrafficResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeTrafficResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTrafficResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrafficResponse) GoString() string {
	return s.String()
}

func (s *DescribeTrafficResponse) SetHeaders(v map[string]*string) *DescribeTrafficResponse {
	s.Headers = v
	return s
}

func (s *DescribeTrafficResponse) SetStatusCode(v int32) *DescribeTrafficResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTrafficResponse) SetBody(v *DescribeTrafficResponseBody) *DescribeTrafficResponse {
	s.Body = v
	return s
}

type ModifyOnDemaondDefenseStatusRequest struct {
	// The protection status of the on-demand instance. Valid values:
	//
	// *   **Defense**: enables the on-demand instance.
	// *   **UnDefense**: disables the on-demand instance.
	DefenseStatus *string `json:"DefenseStatus,omitempty" xml:"DefenseStatus,omitempty"`
	// The ID of the on-demand instance.
	//
	// >  You can call the [DescribeOnDemandInstance](~~152120~~) operation to query the IDs of all on-demand instances.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the on-demand instance.
	//
	// >  You can call the [DescribeRegions](~~118703~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyOnDemaondDefenseStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyOnDemaondDefenseStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyOnDemaondDefenseStatusRequest) SetDefenseStatus(v string) *ModifyOnDemaondDefenseStatusRequest {
	s.DefenseStatus = &v
	return s
}

func (s *ModifyOnDemaondDefenseStatusRequest) SetInstanceId(v string) *ModifyOnDemaondDefenseStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyOnDemaondDefenseStatusRequest) SetRegionId(v string) *ModifyOnDemaondDefenseStatusRequest {
	s.RegionId = &v
	return s
}

type ModifyOnDemaondDefenseStatusResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyOnDemaondDefenseStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyOnDemaondDefenseStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyOnDemaondDefenseStatusResponseBody) SetRequestId(v string) *ModifyOnDemaondDefenseStatusResponseBody {
	s.RequestId = &v
	return s
}

type ModifyOnDemaondDefenseStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyOnDemaondDefenseStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyOnDemaondDefenseStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyOnDemaondDefenseStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyOnDemaondDefenseStatusResponse) SetHeaders(v map[string]*string) *ModifyOnDemaondDefenseStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyOnDemaondDefenseStatusResponse) SetStatusCode(v int32) *ModifyOnDemaondDefenseStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyOnDemaondDefenseStatusResponse) SetBody(v *ModifyOnDemaondDefenseStatusResponseBody) *ModifyOnDemaondDefenseStatusResponse {
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
		"cn-qingdao":            tea.String("ddosbgp.aliyuncs.com"),
		"cn-beijing":            tea.String("ddosbgp.aliyuncs.com"),
		"cn-zhangjiakou":        tea.String("ddosbgp.aliyuncs.com"),
		"cn-huhehaote":          tea.String("ddosbgp.aliyuncs.com"),
		"cn-hangzhou":           tea.String("ddosbgp.aliyuncs.com"),
		"cn-shanghai":           tea.String("ddosbgp.aliyuncs.com"),
		"cn-shenzhen":           tea.String("ddosbgp.aliyuncs.com"),
		"ap-northeast-1":        tea.String("ddosbgp.ap-southeast-1.aliyuncs.com"),
		"ap-south-1":            tea.String("ddosbgp.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":        tea.String("ddosbgp.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-3":        tea.String("ddosbgp.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-5":        tea.String("ddosbgp.ap-southeast-1.aliyuncs.com"),
		"cn-chengdu":            tea.String("ddosbgp.aliyuncs.com"),
		"eu-central-1":          tea.String("ddosbgp.ap-southeast-1.aliyuncs.com"),
		"eu-west-1":             tea.String("ddosbgp.ap-southeast-1.aliyuncs.com"),
		"me-east-1":             tea.String("ddosbgp.ap-southeast-1.aliyuncs.com"),
		"cn-hangzhou-finance":   tea.String("ddosbgp.aliyuncs.com"),
		"cn-shenzhen-finance-1": tea.String("ddosbgp.aliyuncs.com"),
		"cn-shanghai-finance-1": tea.String("ddosbgp.aliyuncs.com"),
		"cn-north-2-gov-1":      tea.String("ddosbgp.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ddosbgp"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddIpWithOptions(request *AddIpRequest, runtime *util.RuntimeOptions) (_result *AddIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IpList)) {
		query["IpList"] = request.IpList
	}

	if !tea.BoolValue(util.IsUnset(request.PackId)) {
		query["PackId"] = request.PackId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddIp"),
		Version:     tea.String("2017-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddIp(request *AddIpRequest) (_result *AddIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddIpResponse{}
	_body, _err := client.AddIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddProductWithOptions(request *AddProductRequest, runtime *util.RuntimeOptions) (_result *AddProductResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PackId)) {
		query["PackId"] = request.PackId
	}

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		query["Product"] = request.Product
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddProduct"),
		Version:     tea.String("2017-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddProduct(request *AddProductRequest) (_result *AddProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddProductResponse{}
	_body, _err := client.AddProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckGrantWithOptions(request *CheckGrantRequest, runtime *util.RuntimeOptions) (_result *CheckGrantResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckGrant"),
		Version:     tea.String("2017-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckGrantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckGrant(request *CheckGrantRequest) (_result *CheckGrantResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckGrantResponse{}
	_body, _err := client.CheckGrantWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBlackholeWithOptions(request *DeleteBlackholeRequest, runtime *util.RuntimeOptions) (_result *DeleteBlackholeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.PackId)) {
		query["PackId"] = request.PackId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteBlackhole"),
		Version:     tea.String("2017-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteBlackholeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBlackhole(request *DeleteBlackholeRequest) (_result *DeleteBlackholeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteBlackholeResponse{}
	_body, _err := client.DeleteBlackholeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteIpWithOptions(request *DeleteIpRequest, runtime *util.RuntimeOptions) (_result *DeleteIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IpList)) {
		query["IpList"] = request.IpList
	}

	if !tea.BoolValue(util.IsUnset(request.PackId)) {
		query["PackId"] = request.PackId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIp"),
		Version:     tea.String("2017-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteIp(request *DeleteIpRequest) (_result *DeleteIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteIpResponse{}
	_body, _err := client.DeleteIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProductWithOptions(request *DeleteProductRequest, runtime *util.RuntimeOptions) (_result *DeleteProductResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PackId)) {
		query["PackId"] = request.PackId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteProduct"),
		Version:     tea.String("2017-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProduct(request *DeleteProductRequest) (_result *DeleteProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteProductResponse{}
	_body, _err := client.DeleteProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDdosEventWithOptions(request *DescribeDdosEventRequest, runtime *util.RuntimeOptions) (_result *DescribeDdosEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PackId)) {
		query["PackId"] = request.PackId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDdosEvent"),
		Version:     tea.String("2017-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDdosEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDdosEvent(request *DescribeDdosEventRequest) (_result *DescribeDdosEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDdosEventResponse{}
	_body, _err := client.DescribeDdosEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceListWithOptions(request *DescribeInstanceListRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIdList)) {
		query["InstanceIdList"] = request.InstanceIdList
	}

	if !tea.BoolValue(util.IsUnset(request.PackIdList)) {
		query["PackIdList"] = request.PackIdList
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceList"),
		Version:     tea.String("2017-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceList(request *DescribeInstanceListRequest) (_result *DescribeInstanceListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceListResponse{}
	_body, _err := client.DescribeInstanceListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Queries the information about on-demand instances, such as whether an on-demand instance is enabled and the CIDR block of each on-demand instance.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeOnDemandInstanceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeOnDemandInstanceResponse
 */
func (client *Client) DescribeOnDemandInstanceWithOptions(request *DescribeOnDemandInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeOnDemandInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOnDemandInstance"),
		Version:     tea.String("2017-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOnDemandInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Queries the information about on-demand instances, such as whether an on-demand instance is enabled and the CIDR block of each on-demand instance.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeOnDemandInstanceRequest
 * @return DescribeOnDemandInstanceResponse
 */
func (client *Client) DescribeOnDemandInstance(request *DescribeOnDemandInstanceRequest) (_result *DescribeOnDemandInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOnDemandInstanceResponse{}
	_body, _err := client.DescribeOnDemandInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOpEntitiesWithOptions(request *DescribeOpEntitiesRequest, runtime *util.RuntimeOptions) (_result *DescribeOpEntitiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOpEntities"),
		Version:     tea.String("2017-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOpEntitiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOpEntities(request *DescribeOpEntitiesRequest) (_result *DescribeOpEntitiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOpEntitiesResponse{}
	_body, _err := client.DescribeOpEntitiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePackWithOptions(request *DescribePackRequest, runtime *util.RuntimeOptions) (_result *DescribePackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PackId)) {
		query["PackId"] = request.PackId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePack"),
		Version:     tea.String("2017-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePack(request *DescribePackRequest) (_result *DescribePackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePackResponse{}
	_body, _err := client.DescribePackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePackListWithOptions(request *DescribePackListRequest, runtime *util.RuntimeOptions) (_result *DescribePackListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePackList"),
		Version:     tea.String("2017-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePackListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePackList(request *DescribePackListRequest) (_result *DescribePackListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePackListResponse{}
	_body, _err := client.DescribePackListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePackPaidTrafficWithOptions(request *DescribePackPaidTrafficRequest, runtime *util.RuntimeOptions) (_result *DescribePackPaidTrafficResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePackPaidTraffic"),
		Version:     tea.String("2017-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePackPaidTrafficResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePackPaidTraffic(request *DescribePackPaidTrafficRequest) (_result *DescribePackPaidTrafficResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePackPaidTrafficResponse{}
	_body, _err := client.DescribePackPaidTrafficWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeResourcePackInstancesWithOptions(request *DescribeResourcePackInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeResourcePackInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeResourcePackInstances"),
		Version:     tea.String("2017-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeResourcePackInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeResourcePackInstances(request *DescribeResourcePackInstancesRequest) (_result *DescribeResourcePackInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeResourcePackInstancesResponse{}
	_body, _err := client.DescribeResourcePackInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeResourcePackStatisticsWithOptions(request *DescribeResourcePackStatisticsRequest, runtime *util.RuntimeOptions) (_result *DescribeResourcePackStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeResourcePackStatistics"),
		Version:     tea.String("2017-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeResourcePackStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeResourcePackStatistics(request *DescribeResourcePackStatisticsRequest) (_result *DescribeResourcePackStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeResourcePackStatisticsResponse{}
	_body, _err := client.DescribeResourcePackStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeResourcePackUsageWithOptions(request *DescribeResourcePackUsageRequest, runtime *util.RuntimeOptions) (_result *DescribeResourcePackUsageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeResourcePackUsage"),
		Version:     tea.String("2017-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeResourcePackUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeResourcePackUsage(request *DescribeResourcePackUsageRequest) (_result *DescribeResourcePackUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeResourcePackUsageResponse{}
	_body, _err := client.DescribeResourcePackUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DescribeTopTraffic operation to query the top N IP addresses from which the most traffic is forwarded by an on-demand instance within a specific period.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeTopTrafficRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeTopTrafficResponse
 */
func (client *Client) DescribeTopTrafficWithOptions(request *DescribeTopTrafficRequest, runtime *util.RuntimeOptions) (_result *DescribeTopTrafficResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Ipnet)) {
		query["Ipnet"] = request.Ipnet
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Rn)) {
		query["Rn"] = request.Rn
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTopTraffic"),
		Version:     tea.String("2017-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTopTrafficResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DescribeTopTraffic operation to query the top N IP addresses from which the most traffic is forwarded by an on-demand instance within a specific period.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeTopTrafficRequest
 * @return DescribeTopTrafficResponse
 */
func (client *Client) DescribeTopTraffic(request *DescribeTopTrafficRequest) (_result *DescribeTopTrafficResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTopTrafficResponse{}
	_body, _err := client.DescribeTopTrafficWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTrafficWithOptions(request *DescribeTrafficRequest, runtime *util.RuntimeOptions) (_result *DescribeTrafficResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTraffic"),
		Version:     tea.String("2017-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTrafficResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTraffic(request *DescribeTrafficRequest) (_result *DescribeTrafficResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTrafficResponse{}
	_body, _err := client.DescribeTrafficWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyOnDemaondDefenseStatusWithOptions(request *ModifyOnDemaondDefenseStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyOnDemaondDefenseStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DefenseStatus)) {
		query["DefenseStatus"] = request.DefenseStatus
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyOnDemaondDefenseStatus"),
		Version:     tea.String("2017-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyOnDemaondDefenseStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyOnDemaondDefenseStatus(request *ModifyOnDemaondDefenseStatusRequest) (_result *ModifyOnDemaondDefenseStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyOnDemaondDefenseStatusResponse{}
	_body, _err := client.ModifyOnDemaondDefenseStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
