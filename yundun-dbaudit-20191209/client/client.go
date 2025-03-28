// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ClearInstanceStorageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dbaudit-cn-78v1gc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dbaudit-audit-dbaudit-cn-78v1gc****
	StorageCategory *string `json:"StorageCategory,omitempty" xml:"StorageCategory,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dbaudit-cn-78v1gc****
	StorageSpace *string `json:"StorageSpace,omitempty" xml:"StorageSpace,omitempty"`
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

func (s *ClearInstanceStorageRequest) SetLang(v string) *ClearInstanceStorageRequest {
	s.Lang = &v
	return s
}

func (s *ClearInstanceStorageRequest) SetRegionId(v string) *ClearInstanceStorageRequest {
	s.RegionId = &v
	return s
}

func (s *ClearInstanceStorageRequest) SetStorageCategory(v string) *ClearInstanceStorageRequest {
	s.StorageCategory = &v
	return s
}

func (s *ClearInstanceStorageRequest) SetStorageSpace(v string) *ClearInstanceStorageRequest {
	s.StorageSpace = &v
	return s
}

type ClearInstanceStorageResponseBody struct {
	// example:
	//
	// dbaudit-cn-78v1gc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 482EF142-BFA5-43FF-B4B0-84A4B0763639
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClearInstanceStorageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ClearInstanceStorageResponseBody) GoString() string {
	return s.String()
}

func (s *ClearInstanceStorageResponseBody) SetInstanceId(v string) *ClearInstanceStorageResponseBody {
	s.InstanceId = &v
	return s
}

func (s *ClearInstanceStorageResponseBody) SetRequestId(v string) *ClearInstanceStorageResponseBody {
	s.RequestId = &v
	return s
}

type ClearInstanceStorageResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClearInstanceStorageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ClearInstanceStorageResponse) SetStatusCode(v int32) *ClearInstanceStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *ClearInstanceStorageResponse) SetBody(v *ClearInstanceStorageResponseBody) *ClearInstanceStorageResponse {
	s.Body = v
	return s
}

type DescribeInstanceAttributeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dbaudit-cn-78v1gc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *DescribeInstanceAttributeRequest) SetLang(v string) *DescribeInstanceAttributeRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInstanceAttributeRequest) SetRegionId(v string) *DescribeInstanceAttributeRequest {
	s.RegionId = &v
	return s
}

type DescribeInstanceAttributeResponseBody struct {
	InstanceAttribute *DescribeInstanceAttributeResponseBodyInstanceAttribute `json:"InstanceAttribute,omitempty" xml:"InstanceAttribute,omitempty" type:"Struct"`
	// example:
	//
	// 28382024-466D-4641-A144-40FD0DD53766
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponseBody) SetInstanceAttribute(v *DescribeInstanceAttributeResponseBodyInstanceAttribute) *DescribeInstanceAttributeResponseBody {
	s.InstanceAttribute = v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetRequestId(v string) *DescribeInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInstanceAttributeResponseBodyInstanceAttribute struct {
	// example:
	//
	// API
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1578067200000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// dbaudit-cn-78v1gc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// RUNNING
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// example:
	//
	// tsiqvqjjlq-public.dbaudit.aliyuncs.com
	InternetEndpoint *string `json:"InternetEndpoint,omitempty" xml:"InternetEndpoint,omitempty"`
	// example:
	//
	// tsiqvqjjlq.dbaudit.aliyuncs.com
	IntranetEndpoint *string `json:"IntranetEndpoint,omitempty" xml:"IntranetEndpoint,omitempty"`
	// example:
	//
	// alpha.professional
	LicenseCode *string `json:"LicenseCode,omitempty" xml:"LicenseCode,omitempty"`
	// example:
	//
	// true
	PublicNetworkAccess *bool `json:"PublicNetworkAccess,omitempty" xml:"PublicNetworkAccess,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// alpha
	SeriesCode *string `json:"SeriesCode,omitempty" xml:"SeriesCode,omitempty"`
	// example:
	//
	// 1577437765000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// vpc-bp1c85tzgqu1bf5b****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// vsw-bp1kep1f0k5fnyfs****
	VswitchId *string   `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	WhiteList []*string `json:"WhiteList,omitempty" xml:"WhiteList,omitempty" type:"Repeated"`
}

func (s DescribeInstanceAttributeResponseBodyInstanceAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAttributeResponseBodyInstanceAttribute) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetDescription(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.Description = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetExpireTime(v int64) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetInstanceId(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetInstanceStatus(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetInternetEndpoint(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.InternetEndpoint = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetIntranetEndpoint(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.IntranetEndpoint = &v
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

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetRegionId(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetSeriesCode(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.SeriesCode = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetStartTime(v int64) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.StartTime = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetVpcId(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.VpcId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetVswitchId(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.VswitchId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetWhiteList(v []*string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.WhiteList = v
	return s
}

type DescribeInstanceAttributeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeInstanceAttributeResponse) SetStatusCode(v int32) *DescribeInstanceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceAttributeResponse) SetBody(v *DescribeInstanceAttributeResponseBody) *DescribeInstanceAttributeResponse {
	s.Body = v
	return s
}

type DescribeInstanceStorageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dbaudit-cn-78v1gc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *DescribeInstanceStorageRequest) SetLang(v string) *DescribeInstanceStorageRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInstanceStorageRequest) SetRegionId(v string) *DescribeInstanceStorageRequest {
	s.RegionId = &v
	return s
}

type DescribeInstanceStorageResponseBody struct {
	InstanceStorages []*DescribeInstanceStorageResponseBodyInstanceStorages `json:"InstanceStorages,omitempty" xml:"InstanceStorages,omitempty" type:"Repeated"`
	// example:
	//
	// 4226E2BB-EED8-4067-B31B-7F02966765B2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceStorageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceStorageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStorageResponseBody) SetInstanceStorages(v []*DescribeInstanceStorageResponseBodyInstanceStorages) *DescribeInstanceStorageResponseBody {
	s.InstanceStorages = v
	return s
}

func (s *DescribeInstanceStorageResponseBody) SetRequestId(v string) *DescribeInstanceStorageResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInstanceStorageResponseBodyInstanceStorages struct {
	// example:
	//
	// 6047313952768
	StorageCapacity *int64 `json:"StorageCapacity,omitempty" xml:"StorageCapacity,omitempty"`
	// example:
	//
	// dbaudit-audit-dbaudit-cn-78v1gc****
	StorageCategory *string `json:"StorageCategory,omitempty" xml:"StorageCategory,omitempty"`
	// example:
	//
	// dbaudit-cn-78v1gc****
	StorageSpace *string `json:"StorageSpace,omitempty" xml:"StorageSpace,omitempty"`
	// example:
	//
	// 180
	StorageTime *int64 `json:"StorageTime,omitempty" xml:"StorageTime,omitempty"`
	// example:
	//
	// 0
	StorageUsed *int64 `json:"StorageUsed,omitempty" xml:"StorageUsed,omitempty"`
}

func (s DescribeInstanceStorageResponseBodyInstanceStorages) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceStorageResponseBodyInstanceStorages) GoString() string {
	return s.String()
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

func (s *DescribeInstanceStorageResponseBodyInstanceStorages) SetStorageTime(v int64) *DescribeInstanceStorageResponseBodyInstanceStorages {
	s.StorageTime = &v
	return s
}

func (s *DescribeInstanceStorageResponseBodyInstanceStorages) SetStorageUsed(v int64) *DescribeInstanceStorageResponseBodyInstanceStorages {
	s.StorageUsed = &v
	return s
}

type DescribeInstanceStorageResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceStorageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeInstanceStorageResponse) SetStatusCode(v int32) *DescribeInstanceStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceStorageResponse) SetBody(v *DescribeInstanceStorageResponseBody) *DescribeInstanceStorageResponse {
	s.Body = v
	return s
}

type DescribeInstancesRequest struct {
	GetCenterInstance *bool `json:"GetCenterInstance,omitempty" xml:"GetCenterInstance,omitempty"`
	// example:
	//
	// dbaudit-cn-78v1gc****
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	// example:
	//
	// RUNNING
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	Lang           *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfm26ougij6****
	ResourceGroupId *string                        `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag             []*DescribeInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequest) SetGetCenterInstance(v bool) *DescribeInstancesRequest {
	s.GetCenterInstance = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceId(v []*string) *DescribeInstancesRequest {
	s.InstanceId = v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceStatus(v string) *DescribeInstancesRequest {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeInstancesRequest) SetLang(v string) *DescribeInstancesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageNumber(v int32) *DescribeInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageSize(v int32) *DescribeInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesRequest) SetRegionId(v string) *DescribeInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstancesRequest) SetResourceGroupId(v string) *DescribeInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstancesRequest) SetTag(v []*DescribeInstancesRequestTag) *DescribeInstancesRequest {
	s.Tag = v
	return s
}

type DescribeInstancesRequestTag struct {
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// testapi
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
	Instances []*DescribeInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// example:
	//
	// 12AA0353-F01D-43E2-A85C-9040F7A35D93
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 18
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBody) SetInstances(v []*DescribeInstancesResponseBodyInstances) *DescribeInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeInstancesResponseBody) SetRequestId(v string) *DescribeInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetTotalCount(v int64) *DescribeInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeInstancesResponseBodyInstances struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1578067200000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// dbaudit-cn-78v1gc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// RUNNING
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// example:
	//
	// tsiqvqjjlq-public.dbaudit.aliyuncs.com
	InternetEndpoint *string `json:"InternetEndpoint,omitempty" xml:"InternetEndpoint,omitempty"`
	// example:
	//
	// tsiqvqjjlq.dbaudit.aliyuncs.com
	IntranetEndpoint *string `json:"IntranetEndpoint,omitempty" xml:"IntranetEndpoint,omitempty"`
	// example:
	//
	// alpha.professional
	LicenseCode *string `json:"LicenseCode,omitempty" xml:"LicenseCode,omitempty"`
	// example:
	//
	// true
	PublicNetworkAccess *bool `json:"PublicNetworkAccess,omitempty" xml:"PublicNetworkAccess,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// alpha
	SeriesCode *string `json:"SeriesCode,omitempty" xml:"SeriesCode,omitempty"`
	// example:
	//
	// 1577437765000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// vpc-bp1c85tzgqu1bf5b****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// vsw-bp1kep1f0k5fnyfs****
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s DescribeInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstances) SetDescription(v string) *DescribeInstancesResponseBodyInstances {
	s.Description = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetExpireTime(v int64) *DescribeInstancesResponseBodyInstances {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInstanceId(v string) *DescribeInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInstanceStatus(v string) *DescribeInstancesResponseBodyInstances {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInternetEndpoint(v string) *DescribeInstancesResponseBodyInstances {
	s.InternetEndpoint = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetIntranetEndpoint(v string) *DescribeInstancesResponseBodyInstances {
	s.IntranetEndpoint = &v
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

func (s *DescribeInstancesResponseBodyInstances) SetRegionId(v string) *DescribeInstancesResponseBodyInstances {
	s.RegionId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetSeriesCode(v string) *DescribeInstancesResponseBodyInstances {
	s.SeriesCode = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetStartTime(v int64) *DescribeInstancesResponseBodyInstances {
	s.StartTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetVpcId(v string) *DescribeInstancesResponseBodyInstances {
	s.VpcId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetVswitchId(v string) *DescribeInstancesResponseBodyInstances {
	s.VswitchId = &v
	return s
}

type DescribeInstancesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeInstancesResponse) SetStatusCode(v int32) *DescribeInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstancesResponse) SetBody(v *DescribeInstancesResponseBody) *DescribeInstancesResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// example:
	//
	// 52EcpzBpR86EEpcc.9xyfvym3cKAXsdV2SSFZnouWTRzf1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// dbaudit-cn-78v1gc****
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// INSTANCE
	ResourceType *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// testapi
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
	// example:
	//
	// 52EcpzBpR86EEpcc.9xyfvym3cKAXsdV2SSFZnouWTRzf1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// E6A08A8A-F962-4FAD-AF0C-86B393E1F9C1
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
	// example:
	//
	// dbaudit-cn-78v1gc****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// INSTANCE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// test
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// example:
	//
	// testapi
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
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

// Summary:
//
// 清理SLS存储空间
//
// @param request - ClearInstanceStorageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClearInstanceStorageResponse
func (client *Client) ClearInstanceStorageWithOptions(request *ClearInstanceStorageRequest, runtime *util.RuntimeOptions) (_result *ClearInstanceStorageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StorageCategory)) {
		query["StorageCategory"] = request.StorageCategory
	}

	if !tea.BoolValue(util.IsUnset(request.StorageSpace)) {
		query["StorageSpace"] = request.StorageSpace
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ClearInstanceStorage"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ClearInstanceStorageResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ClearInstanceStorageResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 清理SLS存储空间
//
// @param request - ClearInstanceStorageRequest
//
// @return ClearInstanceStorageResponse
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

// Summary:
//
// 获取实例属性
//
// @param request - DescribeInstanceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceAttributeResponse
func (client *Client) DescribeInstanceAttributeWithOptions(request *DescribeInstanceAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceAttribute"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeInstanceAttributeResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeInstanceAttributeResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取实例属性
//
// @param request - DescribeInstanceAttributeRequest
//
// @return DescribeInstanceAttributeResponse
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

// Summary:
//
// 获取存储大小
//
// @param request - DescribeInstanceStorageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceStorageResponse
func (client *Client) DescribeInstanceStorageWithOptions(request *DescribeInstanceStorageRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceStorageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceStorage"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeInstanceStorageResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeInstanceStorageResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取存储大小
//
// @param request - DescribeInstanceStorageRequest
//
// @return DescribeInstanceStorageResponse
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

// Summary:
//
// 获取实例列表
//
// @param request - DescribeInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstancesResponse
func (client *Client) DescribeInstancesWithOptions(request *DescribeInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GetCenterInstance)) {
		query["GetCenterInstance"] = request.GetCenterInstance
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceStatus)) {
		query["InstanceStatus"] = request.InstanceStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
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

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstances"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeInstancesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeInstancesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取实例列表
//
// @param request - DescribeInstancesRequest
//
// @return DescribeInstancesResponse
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

// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListTagResourcesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListTagResourcesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
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
