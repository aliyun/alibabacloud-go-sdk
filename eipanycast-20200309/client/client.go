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

type AllocateAnycastEipAddressRequest struct {
	Bandwidth          *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	ClientToken        *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ServiceLocation    *string `json:"ServiceLocation,omitempty" xml:"ServiceLocation,omitempty"`
}

func (s AllocateAnycastEipAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s AllocateAnycastEipAddressRequest) GoString() string {
	return s.String()
}

func (s *AllocateAnycastEipAddressRequest) SetBandwidth(v string) *AllocateAnycastEipAddressRequest {
	s.Bandwidth = &v
	return s
}

func (s *AllocateAnycastEipAddressRequest) SetClientToken(v string) *AllocateAnycastEipAddressRequest {
	s.ClientToken = &v
	return s
}

func (s *AllocateAnycastEipAddressRequest) SetDescription(v string) *AllocateAnycastEipAddressRequest {
	s.Description = &v
	return s
}

func (s *AllocateAnycastEipAddressRequest) SetInstanceChargeType(v string) *AllocateAnycastEipAddressRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *AllocateAnycastEipAddressRequest) SetInternetChargeType(v string) *AllocateAnycastEipAddressRequest {
	s.InternetChargeType = &v
	return s
}

func (s *AllocateAnycastEipAddressRequest) SetName(v string) *AllocateAnycastEipAddressRequest {
	s.Name = &v
	return s
}

func (s *AllocateAnycastEipAddressRequest) SetServiceLocation(v string) *AllocateAnycastEipAddressRequest {
	s.ServiceLocation = &v
	return s
}

type AllocateAnycastEipAddressResponseBody struct {
	AnycastId *string `json:"AnycastId,omitempty" xml:"AnycastId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllocateAnycastEipAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AllocateAnycastEipAddressResponseBody) GoString() string {
	return s.String()
}

func (s *AllocateAnycastEipAddressResponseBody) SetAnycastId(v string) *AllocateAnycastEipAddressResponseBody {
	s.AnycastId = &v
	return s
}

func (s *AllocateAnycastEipAddressResponseBody) SetOrderId(v string) *AllocateAnycastEipAddressResponseBody {
	s.OrderId = &v
	return s
}

func (s *AllocateAnycastEipAddressResponseBody) SetRequestId(v string) *AllocateAnycastEipAddressResponseBody {
	s.RequestId = &v
	return s
}

type AllocateAnycastEipAddressResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AllocateAnycastEipAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AllocateAnycastEipAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s AllocateAnycastEipAddressResponse) GoString() string {
	return s.String()
}

func (s *AllocateAnycastEipAddressResponse) SetHeaders(v map[string]*string) *AllocateAnycastEipAddressResponse {
	s.Headers = v
	return s
}

func (s *AllocateAnycastEipAddressResponse) SetStatusCode(v int32) *AllocateAnycastEipAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *AllocateAnycastEipAddressResponse) SetBody(v *AllocateAnycastEipAddressResponseBody) *AllocateAnycastEipAddressResponse {
	s.Body = v
	return s
}

type AssociateAnycastEipAddressRequest struct {
	AnycastId            *string                                          `json:"AnycastId,omitempty" xml:"AnycastId,omitempty"`
	AssociationMode      *string                                          `json:"AssociationMode,omitempty" xml:"AssociationMode,omitempty"`
	BindInstanceId       *string                                          `json:"BindInstanceId,omitempty" xml:"BindInstanceId,omitempty"`
	BindInstanceRegionId *string                                          `json:"BindInstanceRegionId,omitempty" xml:"BindInstanceRegionId,omitempty"`
	BindInstanceType     *string                                          `json:"BindInstanceType,omitempty" xml:"BindInstanceType,omitempty"`
	ClientToken          *string                                          `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun               *bool                                            `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	PopLocations         []*AssociateAnycastEipAddressRequestPopLocations `json:"PopLocations,omitempty" xml:"PopLocations,omitempty" type:"Repeated"`
	PrivateIpAddress     *string                                          `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
}

func (s AssociateAnycastEipAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s AssociateAnycastEipAddressRequest) GoString() string {
	return s.String()
}

func (s *AssociateAnycastEipAddressRequest) SetAnycastId(v string) *AssociateAnycastEipAddressRequest {
	s.AnycastId = &v
	return s
}

func (s *AssociateAnycastEipAddressRequest) SetAssociationMode(v string) *AssociateAnycastEipAddressRequest {
	s.AssociationMode = &v
	return s
}

func (s *AssociateAnycastEipAddressRequest) SetBindInstanceId(v string) *AssociateAnycastEipAddressRequest {
	s.BindInstanceId = &v
	return s
}

func (s *AssociateAnycastEipAddressRequest) SetBindInstanceRegionId(v string) *AssociateAnycastEipAddressRequest {
	s.BindInstanceRegionId = &v
	return s
}

func (s *AssociateAnycastEipAddressRequest) SetBindInstanceType(v string) *AssociateAnycastEipAddressRequest {
	s.BindInstanceType = &v
	return s
}

func (s *AssociateAnycastEipAddressRequest) SetClientToken(v string) *AssociateAnycastEipAddressRequest {
	s.ClientToken = &v
	return s
}

func (s *AssociateAnycastEipAddressRequest) SetDryRun(v bool) *AssociateAnycastEipAddressRequest {
	s.DryRun = &v
	return s
}

func (s *AssociateAnycastEipAddressRequest) SetPopLocations(v []*AssociateAnycastEipAddressRequestPopLocations) *AssociateAnycastEipAddressRequest {
	s.PopLocations = v
	return s
}

func (s *AssociateAnycastEipAddressRequest) SetPrivateIpAddress(v string) *AssociateAnycastEipAddressRequest {
	s.PrivateIpAddress = &v
	return s
}

type AssociateAnycastEipAddressRequestPopLocations struct {
	PopLocation *string `json:"PopLocation,omitempty" xml:"PopLocation,omitempty"`
}

func (s AssociateAnycastEipAddressRequestPopLocations) String() string {
	return tea.Prettify(s)
}

func (s AssociateAnycastEipAddressRequestPopLocations) GoString() string {
	return s.String()
}

func (s *AssociateAnycastEipAddressRequestPopLocations) SetPopLocation(v string) *AssociateAnycastEipAddressRequestPopLocations {
	s.PopLocation = &v
	return s
}

type AssociateAnycastEipAddressResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateAnycastEipAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssociateAnycastEipAddressResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateAnycastEipAddressResponseBody) SetRequestId(v string) *AssociateAnycastEipAddressResponseBody {
	s.RequestId = &v
	return s
}

type AssociateAnycastEipAddressResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AssociateAnycastEipAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AssociateAnycastEipAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s AssociateAnycastEipAddressResponse) GoString() string {
	return s.String()
}

func (s *AssociateAnycastEipAddressResponse) SetHeaders(v map[string]*string) *AssociateAnycastEipAddressResponse {
	s.Headers = v
	return s
}

func (s *AssociateAnycastEipAddressResponse) SetStatusCode(v int32) *AssociateAnycastEipAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateAnycastEipAddressResponse) SetBody(v *AssociateAnycastEipAddressResponseBody) *AssociateAnycastEipAddressResponse {
	s.Body = v
	return s
}

type DescribeAnycastEipAddressRequest struct {
	AnycastId      *string `json:"AnycastId,omitempty" xml:"AnycastId,omitempty"`
	BindInstanceId *string `json:"BindInstanceId,omitempty" xml:"BindInstanceId,omitempty"`
	Ip             *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribeAnycastEipAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAnycastEipAddressRequest) GoString() string {
	return s.String()
}

func (s *DescribeAnycastEipAddressRequest) SetAnycastId(v string) *DescribeAnycastEipAddressRequest {
	s.AnycastId = &v
	return s
}

func (s *DescribeAnycastEipAddressRequest) SetBindInstanceId(v string) *DescribeAnycastEipAddressRequest {
	s.BindInstanceId = &v
	return s
}

func (s *DescribeAnycastEipAddressRequest) SetIp(v string) *DescribeAnycastEipAddressRequest {
	s.Ip = &v
	return s
}

type DescribeAnycastEipAddressResponseBody struct {
	AliUid                 *int64                                                         `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	AnycastEipBindInfoList []*DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList `json:"AnycastEipBindInfoList,omitempty" xml:"AnycastEipBindInfoList,omitempty" type:"Repeated"`
	AnycastId              *string                                                        `json:"AnycastId,omitempty" xml:"AnycastId,omitempty"`
	Bandwidth              *int32                                                         `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	Bid                    *string                                                        `json:"Bid,omitempty" xml:"Bid,omitempty"`
	BusinessStatus         *string                                                        `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	CreateTime             *string                                                        `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description            *string                                                        `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceChargeType     *string                                                        `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	InternetChargeType     *string                                                        `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	IpAddress              *string                                                        `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	Name                   *string                                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId              *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceLocation        *string                                                        `json:"ServiceLocation,omitempty" xml:"ServiceLocation,omitempty"`
	Status                 *string                                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags                   []*DescribeAnycastEipAddressResponseBodyTags                   `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeAnycastEipAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAnycastEipAddressResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAnycastEipAddressResponseBody) SetAliUid(v int64) *DescribeAnycastEipAddressResponseBody {
	s.AliUid = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBody) SetAnycastEipBindInfoList(v []*DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList) *DescribeAnycastEipAddressResponseBody {
	s.AnycastEipBindInfoList = v
	return s
}

func (s *DescribeAnycastEipAddressResponseBody) SetAnycastId(v string) *DescribeAnycastEipAddressResponseBody {
	s.AnycastId = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBody) SetBandwidth(v int32) *DescribeAnycastEipAddressResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBody) SetBid(v string) *DescribeAnycastEipAddressResponseBody {
	s.Bid = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBody) SetBusinessStatus(v string) *DescribeAnycastEipAddressResponseBody {
	s.BusinessStatus = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBody) SetCreateTime(v string) *DescribeAnycastEipAddressResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBody) SetDescription(v string) *DescribeAnycastEipAddressResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBody) SetInstanceChargeType(v string) *DescribeAnycastEipAddressResponseBody {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBody) SetInternetChargeType(v string) *DescribeAnycastEipAddressResponseBody {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBody) SetIpAddress(v string) *DescribeAnycastEipAddressResponseBody {
	s.IpAddress = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBody) SetName(v string) *DescribeAnycastEipAddressResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBody) SetRequestId(v string) *DescribeAnycastEipAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBody) SetServiceLocation(v string) *DescribeAnycastEipAddressResponseBody {
	s.ServiceLocation = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBody) SetStatus(v string) *DescribeAnycastEipAddressResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBody) SetTags(v []*DescribeAnycastEipAddressResponseBodyTags) *DescribeAnycastEipAddressResponseBody {
	s.Tags = v
	return s
}

type DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList struct {
	AssociationMode      *string                                                                    `json:"AssociationMode,omitempty" xml:"AssociationMode,omitempty"`
	BindInstanceId       *string                                                                    `json:"BindInstanceId,omitempty" xml:"BindInstanceId,omitempty"`
	BindInstanceRegionId *string                                                                    `json:"BindInstanceRegionId,omitempty" xml:"BindInstanceRegionId,omitempty"`
	BindInstanceType     *string                                                                    `json:"BindInstanceType,omitempty" xml:"BindInstanceType,omitempty"`
	BindTime             *string                                                                    `json:"BindTime,omitempty" xml:"BindTime,omitempty"`
	PopLocations         []*DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoListPopLocations `json:"PopLocations,omitempty" xml:"PopLocations,omitempty" type:"Repeated"`
	PrivateIpAddress     *string                                                                    `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	Status               *string                                                                    `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList) GoString() string {
	return s.String()
}

func (s *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList) SetAssociationMode(v string) *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList {
	s.AssociationMode = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList) SetBindInstanceId(v string) *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList {
	s.BindInstanceId = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList) SetBindInstanceRegionId(v string) *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList {
	s.BindInstanceRegionId = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList) SetBindInstanceType(v string) *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList {
	s.BindInstanceType = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList) SetBindTime(v string) *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList {
	s.BindTime = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList) SetPopLocations(v []*DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoListPopLocations) *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList {
	s.PopLocations = v
	return s
}

func (s *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList) SetPrivateIpAddress(v string) *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList {
	s.PrivateIpAddress = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList) SetStatus(v string) *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList {
	s.Status = &v
	return s
}

type DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoListPopLocations struct {
	PopLocation *string `json:"PopLocation,omitempty" xml:"PopLocation,omitempty"`
}

func (s DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoListPopLocations) String() string {
	return tea.Prettify(s)
}

func (s DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoListPopLocations) GoString() string {
	return s.String()
}

func (s *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoListPopLocations) SetPopLocation(v string) *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoListPopLocations {
	s.PopLocation = &v
	return s
}

type DescribeAnycastEipAddressResponseBodyTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAnycastEipAddressResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeAnycastEipAddressResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeAnycastEipAddressResponseBodyTags) SetKey(v string) *DescribeAnycastEipAddressResponseBodyTags {
	s.Key = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBodyTags) SetValue(v string) *DescribeAnycastEipAddressResponseBodyTags {
	s.Value = &v
	return s
}

type DescribeAnycastEipAddressResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAnycastEipAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAnycastEipAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAnycastEipAddressResponse) GoString() string {
	return s.String()
}

func (s *DescribeAnycastEipAddressResponse) SetHeaders(v map[string]*string) *DescribeAnycastEipAddressResponse {
	s.Headers = v
	return s
}

func (s *DescribeAnycastEipAddressResponse) SetStatusCode(v int32) *DescribeAnycastEipAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAnycastEipAddressResponse) SetBody(v *DescribeAnycastEipAddressResponseBody) *DescribeAnycastEipAddressResponse {
	s.Body = v
	return s
}

type DescribeAnycastPopLocationsRequest struct {
	ServiceLocation *string `json:"ServiceLocation,omitempty" xml:"ServiceLocation,omitempty"`
}

func (s DescribeAnycastPopLocationsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAnycastPopLocationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAnycastPopLocationsRequest) SetServiceLocation(v string) *DescribeAnycastPopLocationsRequest {
	s.ServiceLocation = &v
	return s
}

type DescribeAnycastPopLocationsResponseBody struct {
	AnycastPopLocationList []*DescribeAnycastPopLocationsResponseBodyAnycastPopLocationList `json:"AnycastPopLocationList,omitempty" xml:"AnycastPopLocationList,omitempty" type:"Repeated"`
	Count                  *string                                                          `json:"Count,omitempty" xml:"Count,omitempty"`
	RequestId              *string                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAnycastPopLocationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAnycastPopLocationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAnycastPopLocationsResponseBody) SetAnycastPopLocationList(v []*DescribeAnycastPopLocationsResponseBodyAnycastPopLocationList) *DescribeAnycastPopLocationsResponseBody {
	s.AnycastPopLocationList = v
	return s
}

func (s *DescribeAnycastPopLocationsResponseBody) SetCount(v string) *DescribeAnycastPopLocationsResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeAnycastPopLocationsResponseBody) SetRequestId(v string) *DescribeAnycastPopLocationsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAnycastPopLocationsResponseBodyAnycastPopLocationList struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
}

func (s DescribeAnycastPopLocationsResponseBodyAnycastPopLocationList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAnycastPopLocationsResponseBodyAnycastPopLocationList) GoString() string {
	return s.String()
}

func (s *DescribeAnycastPopLocationsResponseBodyAnycastPopLocationList) SetRegionId(v string) *DescribeAnycastPopLocationsResponseBodyAnycastPopLocationList {
	s.RegionId = &v
	return s
}

func (s *DescribeAnycastPopLocationsResponseBodyAnycastPopLocationList) SetRegionName(v string) *DescribeAnycastPopLocationsResponseBodyAnycastPopLocationList {
	s.RegionName = &v
	return s
}

type DescribeAnycastPopLocationsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAnycastPopLocationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAnycastPopLocationsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAnycastPopLocationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAnycastPopLocationsResponse) SetHeaders(v map[string]*string) *DescribeAnycastPopLocationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAnycastPopLocationsResponse) SetStatusCode(v int32) *DescribeAnycastPopLocationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAnycastPopLocationsResponse) SetBody(v *DescribeAnycastPopLocationsResponseBody) *DescribeAnycastPopLocationsResponse {
	s.Body = v
	return s
}

type DescribeAnycastServerRegionsRequest struct {
	ServiceLocation *string `json:"ServiceLocation,omitempty" xml:"ServiceLocation,omitempty"`
}

func (s DescribeAnycastServerRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAnycastServerRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAnycastServerRegionsRequest) SetServiceLocation(v string) *DescribeAnycastServerRegionsRequest {
	s.ServiceLocation = &v
	return s
}

type DescribeAnycastServerRegionsResponseBody struct {
	AnycastServerRegionList []*DescribeAnycastServerRegionsResponseBodyAnycastServerRegionList `json:"AnycastServerRegionList,omitempty" xml:"AnycastServerRegionList,omitempty" type:"Repeated"`
	Count                   *string                                                            `json:"Count,omitempty" xml:"Count,omitempty"`
	RequestId               *string                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAnycastServerRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAnycastServerRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAnycastServerRegionsResponseBody) SetAnycastServerRegionList(v []*DescribeAnycastServerRegionsResponseBodyAnycastServerRegionList) *DescribeAnycastServerRegionsResponseBody {
	s.AnycastServerRegionList = v
	return s
}

func (s *DescribeAnycastServerRegionsResponseBody) SetCount(v string) *DescribeAnycastServerRegionsResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeAnycastServerRegionsResponseBody) SetRequestId(v string) *DescribeAnycastServerRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAnycastServerRegionsResponseBodyAnycastServerRegionList struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
}

func (s DescribeAnycastServerRegionsResponseBodyAnycastServerRegionList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAnycastServerRegionsResponseBodyAnycastServerRegionList) GoString() string {
	return s.String()
}

func (s *DescribeAnycastServerRegionsResponseBodyAnycastServerRegionList) SetRegionId(v string) *DescribeAnycastServerRegionsResponseBodyAnycastServerRegionList {
	s.RegionId = &v
	return s
}

func (s *DescribeAnycastServerRegionsResponseBodyAnycastServerRegionList) SetRegionName(v string) *DescribeAnycastServerRegionsResponseBodyAnycastServerRegionList {
	s.RegionName = &v
	return s
}

type DescribeAnycastServerRegionsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAnycastServerRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAnycastServerRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAnycastServerRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAnycastServerRegionsResponse) SetHeaders(v map[string]*string) *DescribeAnycastServerRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAnycastServerRegionsResponse) SetStatusCode(v int32) *DescribeAnycastServerRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAnycastServerRegionsResponse) SetBody(v *DescribeAnycastServerRegionsResponseBody) *DescribeAnycastServerRegionsResponse {
	s.Body = v
	return s
}

type ListAnycastEipAddressesRequest struct {
	AnycastEipAddress  *string                               `json:"AnycastEipAddress,omitempty" xml:"AnycastEipAddress,omitempty"`
	AnycastId          *string                               `json:"AnycastId,omitempty" xml:"AnycastId,omitempty"`
	AnycastIds         []*string                             `json:"AnycastIds,omitempty" xml:"AnycastIds,omitempty" type:"Repeated"`
	BindInstanceIds    []*string                             `json:"BindInstanceIds,omitempty" xml:"BindInstanceIds,omitempty" type:"Repeated"`
	BusinessStatus     *string                               `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	InstanceChargeType *string                               `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	InternetChargeType *string                               `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	MaxResults         *int32                                `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Name               *string                               `json:"Name,omitempty" xml:"Name,omitempty"`
	NextToken          *string                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ServiceLocation    *string                               `json:"ServiceLocation,omitempty" xml:"ServiceLocation,omitempty"`
	Status             *string                               `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags               []*ListAnycastEipAddressesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListAnycastEipAddressesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAnycastEipAddressesRequest) GoString() string {
	return s.String()
}

func (s *ListAnycastEipAddressesRequest) SetAnycastEipAddress(v string) *ListAnycastEipAddressesRequest {
	s.AnycastEipAddress = &v
	return s
}

func (s *ListAnycastEipAddressesRequest) SetAnycastId(v string) *ListAnycastEipAddressesRequest {
	s.AnycastId = &v
	return s
}

func (s *ListAnycastEipAddressesRequest) SetAnycastIds(v []*string) *ListAnycastEipAddressesRequest {
	s.AnycastIds = v
	return s
}

func (s *ListAnycastEipAddressesRequest) SetBindInstanceIds(v []*string) *ListAnycastEipAddressesRequest {
	s.BindInstanceIds = v
	return s
}

func (s *ListAnycastEipAddressesRequest) SetBusinessStatus(v string) *ListAnycastEipAddressesRequest {
	s.BusinessStatus = &v
	return s
}

func (s *ListAnycastEipAddressesRequest) SetInstanceChargeType(v string) *ListAnycastEipAddressesRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *ListAnycastEipAddressesRequest) SetInternetChargeType(v string) *ListAnycastEipAddressesRequest {
	s.InternetChargeType = &v
	return s
}

func (s *ListAnycastEipAddressesRequest) SetMaxResults(v int32) *ListAnycastEipAddressesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAnycastEipAddressesRequest) SetName(v string) *ListAnycastEipAddressesRequest {
	s.Name = &v
	return s
}

func (s *ListAnycastEipAddressesRequest) SetNextToken(v string) *ListAnycastEipAddressesRequest {
	s.NextToken = &v
	return s
}

func (s *ListAnycastEipAddressesRequest) SetServiceLocation(v string) *ListAnycastEipAddressesRequest {
	s.ServiceLocation = &v
	return s
}

func (s *ListAnycastEipAddressesRequest) SetStatus(v string) *ListAnycastEipAddressesRequest {
	s.Status = &v
	return s
}

func (s *ListAnycastEipAddressesRequest) SetTags(v []*ListAnycastEipAddressesRequestTags) *ListAnycastEipAddressesRequest {
	s.Tags = v
	return s
}

type ListAnycastEipAddressesRequestTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAnycastEipAddressesRequestTags) String() string {
	return tea.Prettify(s)
}

func (s ListAnycastEipAddressesRequestTags) GoString() string {
	return s.String()
}

func (s *ListAnycastEipAddressesRequestTags) SetKey(v string) *ListAnycastEipAddressesRequestTags {
	s.Key = &v
	return s
}

func (s *ListAnycastEipAddressesRequestTags) SetValue(v string) *ListAnycastEipAddressesRequestTags {
	s.Value = &v
	return s
}

type ListAnycastEipAddressesResponseBody struct {
	AnycastList []*ListAnycastEipAddressesResponseBodyAnycastList `json:"AnycastList,omitempty" xml:"AnycastList,omitempty" type:"Repeated"`
	NextToken   *string                                           `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId   *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAnycastEipAddressesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAnycastEipAddressesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAnycastEipAddressesResponseBody) SetAnycastList(v []*ListAnycastEipAddressesResponseBodyAnycastList) *ListAnycastEipAddressesResponseBody {
	s.AnycastList = v
	return s
}

func (s *ListAnycastEipAddressesResponseBody) SetNextToken(v string) *ListAnycastEipAddressesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAnycastEipAddressesResponseBody) SetRequestId(v string) *ListAnycastEipAddressesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAnycastEipAddressesResponseBody) SetTotalCount(v int32) *ListAnycastEipAddressesResponseBody {
	s.TotalCount = &v
	return s
}

type ListAnycastEipAddressesResponseBodyAnycastList struct {
	AliUid                 *int64                                                                  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	AnycastEipBindInfoList []*ListAnycastEipAddressesResponseBodyAnycastListAnycastEipBindInfoList `json:"AnycastEipBindInfoList,omitempty" xml:"AnycastEipBindInfoList,omitempty" type:"Repeated"`
	AnycastId              *string                                                                 `json:"AnycastId,omitempty" xml:"AnycastId,omitempty"`
	Bandwidth              *int32                                                                  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	BusinessStatus         *string                                                                 `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	CreateTime             *string                                                                 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description            *string                                                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceChargeType     *string                                                                 `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	InternetChargeType     *string                                                                 `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	IpAddress              *string                                                                 `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	Name                   *string                                                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	ServiceLocation        *string                                                                 `json:"ServiceLocation,omitempty" xml:"ServiceLocation,omitempty"`
	ServiceManaged         *int32                                                                  `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
	Status                 *string                                                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags                   []*ListAnycastEipAddressesResponseBodyAnycastListTags                   `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListAnycastEipAddressesResponseBodyAnycastList) String() string {
	return tea.Prettify(s)
}

func (s ListAnycastEipAddressesResponseBodyAnycastList) GoString() string {
	return s.String()
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) SetAliUid(v int64) *ListAnycastEipAddressesResponseBodyAnycastList {
	s.AliUid = &v
	return s
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) SetAnycastEipBindInfoList(v []*ListAnycastEipAddressesResponseBodyAnycastListAnycastEipBindInfoList) *ListAnycastEipAddressesResponseBodyAnycastList {
	s.AnycastEipBindInfoList = v
	return s
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) SetAnycastId(v string) *ListAnycastEipAddressesResponseBodyAnycastList {
	s.AnycastId = &v
	return s
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) SetBandwidth(v int32) *ListAnycastEipAddressesResponseBodyAnycastList {
	s.Bandwidth = &v
	return s
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) SetBusinessStatus(v string) *ListAnycastEipAddressesResponseBodyAnycastList {
	s.BusinessStatus = &v
	return s
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) SetCreateTime(v string) *ListAnycastEipAddressesResponseBodyAnycastList {
	s.CreateTime = &v
	return s
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) SetDescription(v string) *ListAnycastEipAddressesResponseBodyAnycastList {
	s.Description = &v
	return s
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) SetInstanceChargeType(v string) *ListAnycastEipAddressesResponseBodyAnycastList {
	s.InstanceChargeType = &v
	return s
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) SetInternetChargeType(v string) *ListAnycastEipAddressesResponseBodyAnycastList {
	s.InternetChargeType = &v
	return s
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) SetIpAddress(v string) *ListAnycastEipAddressesResponseBodyAnycastList {
	s.IpAddress = &v
	return s
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) SetName(v string) *ListAnycastEipAddressesResponseBodyAnycastList {
	s.Name = &v
	return s
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) SetServiceLocation(v string) *ListAnycastEipAddressesResponseBodyAnycastList {
	s.ServiceLocation = &v
	return s
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) SetServiceManaged(v int32) *ListAnycastEipAddressesResponseBodyAnycastList {
	s.ServiceManaged = &v
	return s
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) SetStatus(v string) *ListAnycastEipAddressesResponseBodyAnycastList {
	s.Status = &v
	return s
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) SetTags(v []*ListAnycastEipAddressesResponseBodyAnycastListTags) *ListAnycastEipAddressesResponseBodyAnycastList {
	s.Tags = v
	return s
}

type ListAnycastEipAddressesResponseBodyAnycastListAnycastEipBindInfoList struct {
	BindInstanceId       *string `json:"BindInstanceId,omitempty" xml:"BindInstanceId,omitempty"`
	BindInstanceRegionId *string `json:"BindInstanceRegionId,omitempty" xml:"BindInstanceRegionId,omitempty"`
	BindInstanceType     *string `json:"BindInstanceType,omitempty" xml:"BindInstanceType,omitempty"`
	BindTime             *string `json:"BindTime,omitempty" xml:"BindTime,omitempty"`
}

func (s ListAnycastEipAddressesResponseBodyAnycastListAnycastEipBindInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListAnycastEipAddressesResponseBodyAnycastListAnycastEipBindInfoList) GoString() string {
	return s.String()
}

func (s *ListAnycastEipAddressesResponseBodyAnycastListAnycastEipBindInfoList) SetBindInstanceId(v string) *ListAnycastEipAddressesResponseBodyAnycastListAnycastEipBindInfoList {
	s.BindInstanceId = &v
	return s
}

func (s *ListAnycastEipAddressesResponseBodyAnycastListAnycastEipBindInfoList) SetBindInstanceRegionId(v string) *ListAnycastEipAddressesResponseBodyAnycastListAnycastEipBindInfoList {
	s.BindInstanceRegionId = &v
	return s
}

func (s *ListAnycastEipAddressesResponseBodyAnycastListAnycastEipBindInfoList) SetBindInstanceType(v string) *ListAnycastEipAddressesResponseBodyAnycastListAnycastEipBindInfoList {
	s.BindInstanceType = &v
	return s
}

func (s *ListAnycastEipAddressesResponseBodyAnycastListAnycastEipBindInfoList) SetBindTime(v string) *ListAnycastEipAddressesResponseBodyAnycastListAnycastEipBindInfoList {
	s.BindTime = &v
	return s
}

type ListAnycastEipAddressesResponseBodyAnycastListTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAnycastEipAddressesResponseBodyAnycastListTags) String() string {
	return tea.Prettify(s)
}

func (s ListAnycastEipAddressesResponseBodyAnycastListTags) GoString() string {
	return s.String()
}

func (s *ListAnycastEipAddressesResponseBodyAnycastListTags) SetKey(v string) *ListAnycastEipAddressesResponseBodyAnycastListTags {
	s.Key = &v
	return s
}

func (s *ListAnycastEipAddressesResponseBodyAnycastListTags) SetValue(v string) *ListAnycastEipAddressesResponseBodyAnycastListTags {
	s.Value = &v
	return s
}

type ListAnycastEipAddressesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAnycastEipAddressesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAnycastEipAddressesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAnycastEipAddressesResponse) GoString() string {
	return s.String()
}

func (s *ListAnycastEipAddressesResponse) SetHeaders(v map[string]*string) *ListAnycastEipAddressesResponse {
	s.Headers = v
	return s
}

func (s *ListAnycastEipAddressesResponse) SetStatusCode(v int32) *ListAnycastEipAddressesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAnycastEipAddressesResponse) SetBody(v *ListAnycastEipAddressesResponseBody) *ListAnycastEipAddressesResponse {
	s.Body = v
	return s
}

type ModifyAnycastEipAddressAttributeRequest struct {
	AnycastId   *string `json:"AnycastId,omitempty" xml:"AnycastId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ModifyAnycastEipAddressAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAnycastEipAddressAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyAnycastEipAddressAttributeRequest) SetAnycastId(v string) *ModifyAnycastEipAddressAttributeRequest {
	s.AnycastId = &v
	return s
}

func (s *ModifyAnycastEipAddressAttributeRequest) SetDescription(v string) *ModifyAnycastEipAddressAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyAnycastEipAddressAttributeRequest) SetName(v string) *ModifyAnycastEipAddressAttributeRequest {
	s.Name = &v
	return s
}

type ModifyAnycastEipAddressAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAnycastEipAddressAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAnycastEipAddressAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAnycastEipAddressAttributeResponseBody) SetRequestId(v string) *ModifyAnycastEipAddressAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAnycastEipAddressAttributeResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyAnycastEipAddressAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAnycastEipAddressAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAnycastEipAddressAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyAnycastEipAddressAttributeResponse) SetHeaders(v map[string]*string) *ModifyAnycastEipAddressAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyAnycastEipAddressAttributeResponse) SetStatusCode(v int32) *ModifyAnycastEipAddressAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAnycastEipAddressAttributeResponse) SetBody(v *ModifyAnycastEipAddressAttributeResponseBody) *ModifyAnycastEipAddressAttributeResponse {
	s.Body = v
	return s
}

type ModifyAnycastEipAddressSpecRequest struct {
	AnycastId *string `json:"AnycastId,omitempty" xml:"AnycastId,omitempty"`
	Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
}

func (s ModifyAnycastEipAddressSpecRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAnycastEipAddressSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyAnycastEipAddressSpecRequest) SetAnycastId(v string) *ModifyAnycastEipAddressSpecRequest {
	s.AnycastId = &v
	return s
}

func (s *ModifyAnycastEipAddressSpecRequest) SetBandwidth(v string) *ModifyAnycastEipAddressSpecRequest {
	s.Bandwidth = &v
	return s
}

type ModifyAnycastEipAddressSpecResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAnycastEipAddressSpecResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAnycastEipAddressSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAnycastEipAddressSpecResponseBody) SetRequestId(v string) *ModifyAnycastEipAddressSpecResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAnycastEipAddressSpecResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyAnycastEipAddressSpecResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAnycastEipAddressSpecResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAnycastEipAddressSpecResponse) GoString() string {
	return s.String()
}

func (s *ModifyAnycastEipAddressSpecResponse) SetHeaders(v map[string]*string) *ModifyAnycastEipAddressSpecResponse {
	s.Headers = v
	return s
}

func (s *ModifyAnycastEipAddressSpecResponse) SetStatusCode(v int32) *ModifyAnycastEipAddressSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAnycastEipAddressSpecResponse) SetBody(v *ModifyAnycastEipAddressSpecResponseBody) *ModifyAnycastEipAddressSpecResponse {
	s.Body = v
	return s
}

type ReleaseAnycastEipAddressRequest struct {
	AnycastId   *string `json:"AnycastId,omitempty" xml:"AnycastId,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s ReleaseAnycastEipAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseAnycastEipAddressRequest) GoString() string {
	return s.String()
}

func (s *ReleaseAnycastEipAddressRequest) SetAnycastId(v string) *ReleaseAnycastEipAddressRequest {
	s.AnycastId = &v
	return s
}

func (s *ReleaseAnycastEipAddressRequest) SetClientToken(v string) *ReleaseAnycastEipAddressRequest {
	s.ClientToken = &v
	return s
}

type ReleaseAnycastEipAddressResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseAnycastEipAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseAnycastEipAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseAnycastEipAddressResponseBody) SetRequestId(v string) *ReleaseAnycastEipAddressResponseBody {
	s.RequestId = &v
	return s
}

type ReleaseAnycastEipAddressResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReleaseAnycastEipAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReleaseAnycastEipAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseAnycastEipAddressResponse) GoString() string {
	return s.String()
}

func (s *ReleaseAnycastEipAddressResponse) SetHeaders(v map[string]*string) *ReleaseAnycastEipAddressResponse {
	s.Headers = v
	return s
}

func (s *ReleaseAnycastEipAddressResponse) SetStatusCode(v int32) *ReleaseAnycastEipAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseAnycastEipAddressResponse) SetBody(v *ReleaseAnycastEipAddressResponseBody) *ReleaseAnycastEipAddressResponse {
	s.Body = v
	return s
}

type UnassociateAnycastEipAddressRequest struct {
	AnycastId            *string `json:"AnycastId,omitempty" xml:"AnycastId,omitempty"`
	BindInstanceId       *string `json:"BindInstanceId,omitempty" xml:"BindInstanceId,omitempty"`
	BindInstanceRegionId *string `json:"BindInstanceRegionId,omitempty" xml:"BindInstanceRegionId,omitempty"`
	BindInstanceType     *string `json:"BindInstanceType,omitempty" xml:"BindInstanceType,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun               *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	PrivateIpAddress     *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
}

func (s UnassociateAnycastEipAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s UnassociateAnycastEipAddressRequest) GoString() string {
	return s.String()
}

func (s *UnassociateAnycastEipAddressRequest) SetAnycastId(v string) *UnassociateAnycastEipAddressRequest {
	s.AnycastId = &v
	return s
}

func (s *UnassociateAnycastEipAddressRequest) SetBindInstanceId(v string) *UnassociateAnycastEipAddressRequest {
	s.BindInstanceId = &v
	return s
}

func (s *UnassociateAnycastEipAddressRequest) SetBindInstanceRegionId(v string) *UnassociateAnycastEipAddressRequest {
	s.BindInstanceRegionId = &v
	return s
}

func (s *UnassociateAnycastEipAddressRequest) SetBindInstanceType(v string) *UnassociateAnycastEipAddressRequest {
	s.BindInstanceType = &v
	return s
}

func (s *UnassociateAnycastEipAddressRequest) SetClientToken(v string) *UnassociateAnycastEipAddressRequest {
	s.ClientToken = &v
	return s
}

func (s *UnassociateAnycastEipAddressRequest) SetDryRun(v string) *UnassociateAnycastEipAddressRequest {
	s.DryRun = &v
	return s
}

func (s *UnassociateAnycastEipAddressRequest) SetPrivateIpAddress(v string) *UnassociateAnycastEipAddressRequest {
	s.PrivateIpAddress = &v
	return s
}

type UnassociateAnycastEipAddressResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnassociateAnycastEipAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnassociateAnycastEipAddressResponseBody) GoString() string {
	return s.String()
}

func (s *UnassociateAnycastEipAddressResponseBody) SetRequestId(v string) *UnassociateAnycastEipAddressResponseBody {
	s.RequestId = &v
	return s
}

type UnassociateAnycastEipAddressResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnassociateAnycastEipAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnassociateAnycastEipAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s UnassociateAnycastEipAddressResponse) GoString() string {
	return s.String()
}

func (s *UnassociateAnycastEipAddressResponse) SetHeaders(v map[string]*string) *UnassociateAnycastEipAddressResponse {
	s.Headers = v
	return s
}

func (s *UnassociateAnycastEipAddressResponse) SetStatusCode(v int32) *UnassociateAnycastEipAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *UnassociateAnycastEipAddressResponse) SetBody(v *UnassociateAnycastEipAddressResponseBody) *UnassociateAnycastEipAddressResponse {
	s.Body = v
	return s
}

type UpdateAnycastEipAddressAssociationsRequest struct {
	AnycastId             *string                                                            `json:"AnycastId,omitempty" xml:"AnycastId,omitempty"`
	AssociationMode       *string                                                            `json:"AssociationMode,omitempty" xml:"AssociationMode,omitempty"`
	BindInstanceId        *string                                                            `json:"BindInstanceId,omitempty" xml:"BindInstanceId,omitempty"`
	ClientToken           *string                                                            `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                *bool                                                              `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	PopLocationAddList    []*UpdateAnycastEipAddressAssociationsRequestPopLocationAddList    `json:"PopLocationAddList,omitempty" xml:"PopLocationAddList,omitempty" type:"Repeated"`
	PopLocationDeleteList []*UpdateAnycastEipAddressAssociationsRequestPopLocationDeleteList `json:"PopLocationDeleteList,omitempty" xml:"PopLocationDeleteList,omitempty" type:"Repeated"`
}

func (s UpdateAnycastEipAddressAssociationsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAnycastEipAddressAssociationsRequest) GoString() string {
	return s.String()
}

func (s *UpdateAnycastEipAddressAssociationsRequest) SetAnycastId(v string) *UpdateAnycastEipAddressAssociationsRequest {
	s.AnycastId = &v
	return s
}

func (s *UpdateAnycastEipAddressAssociationsRequest) SetAssociationMode(v string) *UpdateAnycastEipAddressAssociationsRequest {
	s.AssociationMode = &v
	return s
}

func (s *UpdateAnycastEipAddressAssociationsRequest) SetBindInstanceId(v string) *UpdateAnycastEipAddressAssociationsRequest {
	s.BindInstanceId = &v
	return s
}

func (s *UpdateAnycastEipAddressAssociationsRequest) SetClientToken(v string) *UpdateAnycastEipAddressAssociationsRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAnycastEipAddressAssociationsRequest) SetDryRun(v bool) *UpdateAnycastEipAddressAssociationsRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateAnycastEipAddressAssociationsRequest) SetPopLocationAddList(v []*UpdateAnycastEipAddressAssociationsRequestPopLocationAddList) *UpdateAnycastEipAddressAssociationsRequest {
	s.PopLocationAddList = v
	return s
}

func (s *UpdateAnycastEipAddressAssociationsRequest) SetPopLocationDeleteList(v []*UpdateAnycastEipAddressAssociationsRequestPopLocationDeleteList) *UpdateAnycastEipAddressAssociationsRequest {
	s.PopLocationDeleteList = v
	return s
}

type UpdateAnycastEipAddressAssociationsRequestPopLocationAddList struct {
	PopLocation *string `json:"PopLocation,omitempty" xml:"PopLocation,omitempty"`
}

func (s UpdateAnycastEipAddressAssociationsRequestPopLocationAddList) String() string {
	return tea.Prettify(s)
}

func (s UpdateAnycastEipAddressAssociationsRequestPopLocationAddList) GoString() string {
	return s.String()
}

func (s *UpdateAnycastEipAddressAssociationsRequestPopLocationAddList) SetPopLocation(v string) *UpdateAnycastEipAddressAssociationsRequestPopLocationAddList {
	s.PopLocation = &v
	return s
}

type UpdateAnycastEipAddressAssociationsRequestPopLocationDeleteList struct {
	PopLocation *string `json:"PopLocation,omitempty" xml:"PopLocation,omitempty"`
}

func (s UpdateAnycastEipAddressAssociationsRequestPopLocationDeleteList) String() string {
	return tea.Prettify(s)
}

func (s UpdateAnycastEipAddressAssociationsRequestPopLocationDeleteList) GoString() string {
	return s.String()
}

func (s *UpdateAnycastEipAddressAssociationsRequestPopLocationDeleteList) SetPopLocation(v string) *UpdateAnycastEipAddressAssociationsRequestPopLocationDeleteList {
	s.PopLocation = &v
	return s
}

type UpdateAnycastEipAddressAssociationsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAnycastEipAddressAssociationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAnycastEipAddressAssociationsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAnycastEipAddressAssociationsResponseBody) SetRequestId(v string) *UpdateAnycastEipAddressAssociationsResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAnycastEipAddressAssociationsResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAnycastEipAddressAssociationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAnycastEipAddressAssociationsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAnycastEipAddressAssociationsResponse) GoString() string {
	return s.String()
}

func (s *UpdateAnycastEipAddressAssociationsResponse) SetHeaders(v map[string]*string) *UpdateAnycastEipAddressAssociationsResponse {
	s.Headers = v
	return s
}

func (s *UpdateAnycastEipAddressAssociationsResponse) SetStatusCode(v int32) *UpdateAnycastEipAddressAssociationsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAnycastEipAddressAssociationsResponse) SetBody(v *UpdateAnycastEipAddressAssociationsResponseBody) *UpdateAnycastEipAddressAssociationsResponse {
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
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("eipanycast"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AllocateAnycastEipAddressWithOptions(request *AllocateAnycastEipAddressRequest, runtime *util.RuntimeOptions) (_result *AllocateAnycastEipAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceChargeType)) {
		query["InstanceChargeType"] = request.InstanceChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.InternetChargeType)) {
		query["InternetChargeType"] = request.InternetChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceLocation)) {
		query["ServiceLocation"] = request.ServiceLocation
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AllocateAnycastEipAddress"),
		Version:     tea.String("2020-03-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AllocateAnycastEipAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AllocateAnycastEipAddress(request *AllocateAnycastEipAddressRequest) (_result *AllocateAnycastEipAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AllocateAnycastEipAddressResponse{}
	_body, _err := client.AllocateAnycastEipAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AssociateAnycastEipAddressWithOptions(request *AssociateAnycastEipAddressRequest, runtime *util.RuntimeOptions) (_result *AssociateAnycastEipAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AnycastId)) {
		query["AnycastId"] = request.AnycastId
	}

	if !tea.BoolValue(util.IsUnset(request.AssociationMode)) {
		query["AssociationMode"] = request.AssociationMode
	}

	if !tea.BoolValue(util.IsUnset(request.BindInstanceId)) {
		query["BindInstanceId"] = request.BindInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.BindInstanceRegionId)) {
		query["BindInstanceRegionId"] = request.BindInstanceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.BindInstanceType)) {
		query["BindInstanceType"] = request.BindInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.PopLocations)) {
		query["PopLocations"] = request.PopLocations
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateIpAddress)) {
		query["PrivateIpAddress"] = request.PrivateIpAddress
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AssociateAnycastEipAddress"),
		Version:     tea.String("2020-03-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AssociateAnycastEipAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AssociateAnycastEipAddress(request *AssociateAnycastEipAddressRequest) (_result *AssociateAnycastEipAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssociateAnycastEipAddressResponse{}
	_body, _err := client.AssociateAnycastEipAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAnycastEipAddressWithOptions(request *DescribeAnycastEipAddressRequest, runtime *util.RuntimeOptions) (_result *DescribeAnycastEipAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AnycastId)) {
		query["AnycastId"] = request.AnycastId
	}

	if !tea.BoolValue(util.IsUnset(request.BindInstanceId)) {
		query["BindInstanceId"] = request.BindInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAnycastEipAddress"),
		Version:     tea.String("2020-03-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAnycastEipAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAnycastEipAddress(request *DescribeAnycastEipAddressRequest) (_result *DescribeAnycastEipAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAnycastEipAddressResponse{}
	_body, _err := client.DescribeAnycastEipAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAnycastPopLocationsWithOptions(request *DescribeAnycastPopLocationsRequest, runtime *util.RuntimeOptions) (_result *DescribeAnycastPopLocationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceLocation)) {
		query["ServiceLocation"] = request.ServiceLocation
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAnycastPopLocations"),
		Version:     tea.String("2020-03-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAnycastPopLocationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAnycastPopLocations(request *DescribeAnycastPopLocationsRequest) (_result *DescribeAnycastPopLocationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAnycastPopLocationsResponse{}
	_body, _err := client.DescribeAnycastPopLocationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAnycastServerRegionsWithOptions(request *DescribeAnycastServerRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeAnycastServerRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceLocation)) {
		query["ServiceLocation"] = request.ServiceLocation
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAnycastServerRegions"),
		Version:     tea.String("2020-03-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAnycastServerRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAnycastServerRegions(request *DescribeAnycastServerRegionsRequest) (_result *DescribeAnycastServerRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAnycastServerRegionsResponse{}
	_body, _err := client.DescribeAnycastServerRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAnycastEipAddressesWithOptions(request *ListAnycastEipAddressesRequest, runtime *util.RuntimeOptions) (_result *ListAnycastEipAddressesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AnycastEipAddress)) {
		query["AnycastEipAddress"] = request.AnycastEipAddress
	}

	if !tea.BoolValue(util.IsUnset(request.AnycastId)) {
		query["AnycastId"] = request.AnycastId
	}

	if !tea.BoolValue(util.IsUnset(request.AnycastIds)) {
		query["AnycastIds"] = request.AnycastIds
	}

	if !tea.BoolValue(util.IsUnset(request.BindInstanceIds)) {
		query["BindInstanceIds"] = request.BindInstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.BusinessStatus)) {
		query["BusinessStatus"] = request.BusinessStatus
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceChargeType)) {
		query["InstanceChargeType"] = request.InstanceChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.InternetChargeType)) {
		query["InternetChargeType"] = request.InternetChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceLocation)) {
		query["ServiceLocation"] = request.ServiceLocation
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAnycastEipAddresses"),
		Version:     tea.String("2020-03-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAnycastEipAddressesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAnycastEipAddresses(request *ListAnycastEipAddressesRequest) (_result *ListAnycastEipAddressesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAnycastEipAddressesResponse{}
	_body, _err := client.ListAnycastEipAddressesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAnycastEipAddressAttributeWithOptions(request *ModifyAnycastEipAddressAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyAnycastEipAddressAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AnycastId)) {
		query["AnycastId"] = request.AnycastId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAnycastEipAddressAttribute"),
		Version:     tea.String("2020-03-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAnycastEipAddressAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAnycastEipAddressAttribute(request *ModifyAnycastEipAddressAttributeRequest) (_result *ModifyAnycastEipAddressAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAnycastEipAddressAttributeResponse{}
	_body, _err := client.ModifyAnycastEipAddressAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAnycastEipAddressSpecWithOptions(request *ModifyAnycastEipAddressSpecRequest, runtime *util.RuntimeOptions) (_result *ModifyAnycastEipAddressSpecResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AnycastId)) {
		query["AnycastId"] = request.AnycastId
	}

	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		query["Bandwidth"] = request.Bandwidth
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAnycastEipAddressSpec"),
		Version:     tea.String("2020-03-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAnycastEipAddressSpecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAnycastEipAddressSpec(request *ModifyAnycastEipAddressSpecRequest) (_result *ModifyAnycastEipAddressSpecResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAnycastEipAddressSpecResponse{}
	_body, _err := client.ModifyAnycastEipAddressSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseAnycastEipAddressWithOptions(request *ReleaseAnycastEipAddressRequest, runtime *util.RuntimeOptions) (_result *ReleaseAnycastEipAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AnycastId)) {
		query["AnycastId"] = request.AnycastId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseAnycastEipAddress"),
		Version:     tea.String("2020-03-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleaseAnycastEipAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseAnycastEipAddress(request *ReleaseAnycastEipAddressRequest) (_result *ReleaseAnycastEipAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseAnycastEipAddressResponse{}
	_body, _err := client.ReleaseAnycastEipAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnassociateAnycastEipAddressWithOptions(request *UnassociateAnycastEipAddressRequest, runtime *util.RuntimeOptions) (_result *UnassociateAnycastEipAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AnycastId)) {
		query["AnycastId"] = request.AnycastId
	}

	if !tea.BoolValue(util.IsUnset(request.BindInstanceId)) {
		query["BindInstanceId"] = request.BindInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.BindInstanceRegionId)) {
		query["BindInstanceRegionId"] = request.BindInstanceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.BindInstanceType)) {
		query["BindInstanceType"] = request.BindInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateIpAddress)) {
		query["PrivateIpAddress"] = request.PrivateIpAddress
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnassociateAnycastEipAddress"),
		Version:     tea.String("2020-03-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnassociateAnycastEipAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnassociateAnycastEipAddress(request *UnassociateAnycastEipAddressRequest) (_result *UnassociateAnycastEipAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnassociateAnycastEipAddressResponse{}
	_body, _err := client.UnassociateAnycastEipAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAnycastEipAddressAssociationsWithOptions(request *UpdateAnycastEipAddressAssociationsRequest, runtime *util.RuntimeOptions) (_result *UpdateAnycastEipAddressAssociationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AnycastId)) {
		query["AnycastId"] = request.AnycastId
	}

	if !tea.BoolValue(util.IsUnset(request.AssociationMode)) {
		query["AssociationMode"] = request.AssociationMode
	}

	if !tea.BoolValue(util.IsUnset(request.BindInstanceId)) {
		query["BindInstanceId"] = request.BindInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.PopLocationAddList)) {
		query["PopLocationAddList"] = request.PopLocationAddList
	}

	if !tea.BoolValue(util.IsUnset(request.PopLocationDeleteList)) {
		query["PopLocationDeleteList"] = request.PopLocationDeleteList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAnycastEipAddressAssociations"),
		Version:     tea.String("2020-03-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAnycastEipAddressAssociationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAnycastEipAddressAssociations(request *UpdateAnycastEipAddressAssociationsRequest) (_result *UpdateAnycastEipAddressAssociationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAnycastEipAddressAssociationsResponse{}
	_body, _err := client.UpdateAnycastEipAddressAssociationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
