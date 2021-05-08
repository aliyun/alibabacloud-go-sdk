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

type DescribeEndpointRequest struct {
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Password    *string `json:"Password,omitempty" xml:"Password,omitempty"`
}

func (s DescribeEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointRequest) GoString() string {
	return s.String()
}

func (s *DescribeEndpointRequest) SetId(v string) *DescribeEndpointRequest {
	s.Id = &v
	return s
}

func (s *DescribeEndpointRequest) SetServiceCode(v string) *DescribeEndpointRequest {
	s.ServiceCode = &v
	return s
}

func (s *DescribeEndpointRequest) SetPassword(v string) *DescribeEndpointRequest {
	s.Password = &v
	return s
}

type DescribeEndpointResponseBody struct {
	Type        *string                                `json:"Type,omitempty" xml:"Type,omitempty"`
	RequestId   *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Endpoint    *string                                `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	Namespace   *string                                `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	SerivceCode *string                                `json:"SerivceCode,omitempty" xml:"SerivceCode,omitempty"`
	Id          *string                                `json:"Id,omitempty" xml:"Id,omitempty"`
	Protocols   *DescribeEndpointResponseBodyProtocols `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Struct"`
}

func (s DescribeEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEndpointResponseBody) SetType(v string) *DescribeEndpointResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeEndpointResponseBody) SetRequestId(v string) *DescribeEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEndpointResponseBody) SetEndpoint(v string) *DescribeEndpointResponseBody {
	s.Endpoint = &v
	return s
}

func (s *DescribeEndpointResponseBody) SetNamespace(v string) *DescribeEndpointResponseBody {
	s.Namespace = &v
	return s
}

func (s *DescribeEndpointResponseBody) SetSerivceCode(v string) *DescribeEndpointResponseBody {
	s.SerivceCode = &v
	return s
}

func (s *DescribeEndpointResponseBody) SetId(v string) *DescribeEndpointResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeEndpointResponseBody) SetProtocols(v *DescribeEndpointResponseBodyProtocols) *DescribeEndpointResponseBody {
	s.Protocols = v
	return s
}

type DescribeEndpointResponseBodyProtocols struct {
	Protocols []*string `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Repeated"`
}

func (s DescribeEndpointResponseBodyProtocols) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointResponseBodyProtocols) GoString() string {
	return s.String()
}

func (s *DescribeEndpointResponseBodyProtocols) SetProtocols(v []*string) *DescribeEndpointResponseBodyProtocols {
	s.Protocols = v
	return s
}

type DescribeEndpointResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointResponse) GoString() string {
	return s.String()
}

func (s *DescribeEndpointResponse) SetHeaders(v map[string]*string) *DescribeEndpointResponse {
	s.Headers = v
	return s
}

func (s *DescribeEndpointResponse) SetBody(v *DescribeEndpointResponseBody) *DescribeEndpointResponse {
	s.Body = v
	return s
}

type DescribeEndpointsRequest struct {
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeEndpointsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointsRequest) GoString() string {
	return s.String()
}

func (s *DescribeEndpointsRequest) SetId(v string) *DescribeEndpointsRequest {
	s.Id = &v
	return s
}

func (s *DescribeEndpointsRequest) SetServiceCode(v string) *DescribeEndpointsRequest {
	s.ServiceCode = &v
	return s
}

func (s *DescribeEndpointsRequest) SetType(v string) *DescribeEndpointsRequest {
	s.Type = &v
	return s
}

type DescribeEndpointsResponseBody struct {
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Endpoints *DescribeEndpointsResponseBodyEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Struct"`
}

func (s DescribeEndpointsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEndpointsResponseBody) SetSuccess(v bool) *DescribeEndpointsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeEndpointsResponseBody) SetRequestId(v string) *DescribeEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEndpointsResponseBody) SetEndpoints(v *DescribeEndpointsResponseBodyEndpoints) *DescribeEndpointsResponseBody {
	s.Endpoints = v
	return s
}

type DescribeEndpointsResponseBodyEndpoints struct {
	Endpoint []*DescribeEndpointsResponseBodyEndpointsEndpoint `json:"Endpoint,omitempty" xml:"Endpoint,omitempty" type:"Repeated"`
}

func (s DescribeEndpointsResponseBodyEndpoints) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointsResponseBodyEndpoints) GoString() string {
	return s.String()
}

func (s *DescribeEndpointsResponseBodyEndpoints) SetEndpoint(v []*DescribeEndpointsResponseBodyEndpointsEndpoint) *DescribeEndpointsResponseBodyEndpoints {
	s.Endpoint = v
	return s
}

type DescribeEndpointsResponseBodyEndpointsEndpoint struct {
	Type        *string                                                  `json:"Type,omitempty" xml:"Type,omitempty"`
	Namespace   *string                                                  `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	SerivceCode *string                                                  `json:"SerivceCode,omitempty" xml:"SerivceCode,omitempty"`
	Id          *string                                                  `json:"Id,omitempty" xml:"Id,omitempty"`
	Endpoint    *string                                                  `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	Protocols   *DescribeEndpointsResponseBodyEndpointsEndpointProtocols `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Struct"`
}

func (s DescribeEndpointsResponseBodyEndpointsEndpoint) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointsResponseBodyEndpointsEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeEndpointsResponseBodyEndpointsEndpoint) SetType(v string) *DescribeEndpointsResponseBodyEndpointsEndpoint {
	s.Type = &v
	return s
}

func (s *DescribeEndpointsResponseBodyEndpointsEndpoint) SetNamespace(v string) *DescribeEndpointsResponseBodyEndpointsEndpoint {
	s.Namespace = &v
	return s
}

func (s *DescribeEndpointsResponseBodyEndpointsEndpoint) SetSerivceCode(v string) *DescribeEndpointsResponseBodyEndpointsEndpoint {
	s.SerivceCode = &v
	return s
}

func (s *DescribeEndpointsResponseBodyEndpointsEndpoint) SetId(v string) *DescribeEndpointsResponseBodyEndpointsEndpoint {
	s.Id = &v
	return s
}

func (s *DescribeEndpointsResponseBodyEndpointsEndpoint) SetEndpoint(v string) *DescribeEndpointsResponseBodyEndpointsEndpoint {
	s.Endpoint = &v
	return s
}

func (s *DescribeEndpointsResponseBodyEndpointsEndpoint) SetProtocols(v *DescribeEndpointsResponseBodyEndpointsEndpointProtocols) *DescribeEndpointsResponseBodyEndpointsEndpoint {
	s.Protocols = v
	return s
}

type DescribeEndpointsResponseBodyEndpointsEndpointProtocols struct {
	Protocols []*string `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Repeated"`
}

func (s DescribeEndpointsResponseBodyEndpointsEndpointProtocols) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointsResponseBodyEndpointsEndpointProtocols) GoString() string {
	return s.String()
}

func (s *DescribeEndpointsResponseBodyEndpointsEndpointProtocols) SetProtocols(v []*string) *DescribeEndpointsResponseBodyEndpointsEndpointProtocols {
	s.Protocols = v
	return s
}

type DescribeEndpointsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEndpointsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointsResponse) GoString() string {
	return s.String()
}

func (s *DescribeEndpointsResponse) SetHeaders(v map[string]*string) *DescribeEndpointsResponse {
	s.Headers = v
	return s
}

func (s *DescribeEndpointsResponse) SetBody(v *DescribeEndpointsResponseBody) *DescribeEndpointsResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetPassword(v string) *DescribeRegionsRequest {
	s.Password = &v
	return s
}

type DescribeRegionsResponseBody struct {
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RegionIds  *DescribeRegionsResponseBodyRegionIds `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Struct"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetTotalCount(v int32) *DescribeRegionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRegionIds(v *DescribeRegionsResponseBodyRegionIds) *DescribeRegionsResponseBody {
	s.RegionIds = v
	return s
}

type DescribeRegionsResponseBodyRegionIds struct {
	RegionIds []*string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegionIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionIds) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionIds) SetRegionIds(v []*string) *DescribeRegionsResponseBodyRegionIds {
	s.RegionIds = v
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

type DescribeServicesRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
}

func (s DescribeServicesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeServicesRequest) GoString() string {
	return s.String()
}

func (s *DescribeServicesRequest) SetRegionId(v string) *DescribeServicesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeServicesRequest) SetPassword(v string) *DescribeServicesRequest {
	s.Password = &v
	return s
}

type DescribeServicesResponseBody struct {
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Services   *DescribeServicesResponseBodyServices `json:"Services,omitempty" xml:"Services,omitempty" type:"Struct"`
}

func (s DescribeServicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeServicesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServicesResponseBody) SetTotalCount(v int32) *DescribeServicesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeServicesResponseBody) SetRequestId(v string) *DescribeServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServicesResponseBody) SetServices(v *DescribeServicesResponseBodyServices) *DescribeServicesResponseBody {
	s.Services = v
	return s
}

type DescribeServicesResponseBodyServices struct {
	Services []*string `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
}

func (s DescribeServicesResponseBodyServices) String() string {
	return tea.Prettify(s)
}

func (s DescribeServicesResponseBodyServices) GoString() string {
	return s.String()
}

func (s *DescribeServicesResponseBodyServices) SetServices(v []*string) *DescribeServicesResponseBodyServices {
	s.Services = v
	return s
}

type DescribeServicesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeServicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeServicesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeServicesResponse) GoString() string {
	return s.String()
}

func (s *DescribeServicesResponse) SetHeaders(v map[string]*string) *DescribeServicesResponse {
	s.Headers = v
	return s
}

func (s *DescribeServicesResponse) SetBody(v *DescribeServicesResponseBody) *DescribeServicesResponse {
	s.Body = v
	return s
}

type ListEndpointsRequest struct {
	Namespace   *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	SerivceCode *string `json:"SerivceCode,omitempty" xml:"SerivceCode,omitempty"`
}

func (s ListEndpointsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEndpointsRequest) GoString() string {
	return s.String()
}

func (s *ListEndpointsRequest) SetNamespace(v string) *ListEndpointsRequest {
	s.Namespace = &v
	return s
}

func (s *ListEndpointsRequest) SetId(v string) *ListEndpointsRequest {
	s.Id = &v
	return s
}

func (s *ListEndpointsRequest) SetSerivceCode(v string) *ListEndpointsRequest {
	s.SerivceCode = &v
	return s
}

type ListEndpointsResponseBody struct {
	Success      *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	EndpointList *ListEndpointsResponseBodyEndpointList `json:"EndpointList,omitempty" xml:"EndpointList,omitempty" type:"Struct"`
}

func (s ListEndpointsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEndpointsResponseBody) SetSuccess(v bool) *ListEndpointsResponseBody {
	s.Success = &v
	return s
}

func (s *ListEndpointsResponseBody) SetRequestId(v string) *ListEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEndpointsResponseBody) SetEndpointList(v *ListEndpointsResponseBodyEndpointList) *ListEndpointsResponseBody {
	s.EndpointList = v
	return s
}

type ListEndpointsResponseBodyEndpointList struct {
	ItemEndpoint []*ListEndpointsResponseBodyEndpointListItemEndpoint `json:"ItemEndpoint,omitempty" xml:"ItemEndpoint,omitempty" type:"Repeated"`
}

func (s ListEndpointsResponseBodyEndpointList) String() string {
	return tea.Prettify(s)
}

func (s ListEndpointsResponseBodyEndpointList) GoString() string {
	return s.String()
}

func (s *ListEndpointsResponseBodyEndpointList) SetItemEndpoint(v []*ListEndpointsResponseBodyEndpointListItemEndpoint) *ListEndpointsResponseBodyEndpointList {
	s.ItemEndpoint = v
	return s
}

type ListEndpointsResponseBodyEndpointListItemEndpoint struct {
	Type      *string                                                     `json:"Type,omitempty" xml:"Type,omitempty"`
	Namespace *string                                                     `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Product   *string                                                     `json:"Product,omitempty" xml:"Product,omitempty"`
	Id        *string                                                     `json:"Id,omitempty" xml:"Id,omitempty"`
	Endpoint  *string                                                     `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	Protocols *ListEndpointsResponseBodyEndpointListItemEndpointProtocols `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Struct"`
}

func (s ListEndpointsResponseBodyEndpointListItemEndpoint) String() string {
	return tea.Prettify(s)
}

func (s ListEndpointsResponseBodyEndpointListItemEndpoint) GoString() string {
	return s.String()
}

func (s *ListEndpointsResponseBodyEndpointListItemEndpoint) SetType(v string) *ListEndpointsResponseBodyEndpointListItemEndpoint {
	s.Type = &v
	return s
}

func (s *ListEndpointsResponseBodyEndpointListItemEndpoint) SetNamespace(v string) *ListEndpointsResponseBodyEndpointListItemEndpoint {
	s.Namespace = &v
	return s
}

func (s *ListEndpointsResponseBodyEndpointListItemEndpoint) SetProduct(v string) *ListEndpointsResponseBodyEndpointListItemEndpoint {
	s.Product = &v
	return s
}

func (s *ListEndpointsResponseBodyEndpointListItemEndpoint) SetId(v string) *ListEndpointsResponseBodyEndpointListItemEndpoint {
	s.Id = &v
	return s
}

func (s *ListEndpointsResponseBodyEndpointListItemEndpoint) SetEndpoint(v string) *ListEndpointsResponseBodyEndpointListItemEndpoint {
	s.Endpoint = &v
	return s
}

func (s *ListEndpointsResponseBodyEndpointListItemEndpoint) SetProtocols(v *ListEndpointsResponseBodyEndpointListItemEndpointProtocols) *ListEndpointsResponseBodyEndpointListItemEndpoint {
	s.Protocols = v
	return s
}

type ListEndpointsResponseBodyEndpointListItemEndpointProtocols struct {
	Protocols []*string `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Repeated"`
}

func (s ListEndpointsResponseBodyEndpointListItemEndpointProtocols) String() string {
	return tea.Prettify(s)
}

func (s ListEndpointsResponseBodyEndpointListItemEndpointProtocols) GoString() string {
	return s.String()
}

func (s *ListEndpointsResponseBodyEndpointListItemEndpointProtocols) SetProtocols(v []*string) *ListEndpointsResponseBodyEndpointListItemEndpointProtocols {
	s.Protocols = v
	return s
}

type ListEndpointsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEndpointsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEndpointsResponse) GoString() string {
	return s.String()
}

func (s *ListEndpointsResponse) SetHeaders(v map[string]*string) *ListEndpointsResponse {
	s.Headers = v
	return s
}

func (s *ListEndpointsResponse) SetBody(v *ListEndpointsResponseBody) *ListEndpointsResponse {
	s.Body = v
	return s
}

type ListEndpointsByIpRequest struct {
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s ListEndpointsByIpRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEndpointsByIpRequest) GoString() string {
	return s.String()
}

func (s *ListEndpointsByIpRequest) SetIp(v string) *ListEndpointsByIpRequest {
	s.Ip = &v
	return s
}

type ListEndpointsByIpResponseBody struct {
	Success      *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId    *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	EndpointList *ListEndpointsByIpResponseBodyEndpointList `json:"EndpointList,omitempty" xml:"EndpointList,omitempty" type:"Struct"`
}

func (s ListEndpointsByIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEndpointsByIpResponseBody) GoString() string {
	return s.String()
}

func (s *ListEndpointsByIpResponseBody) SetSuccess(v bool) *ListEndpointsByIpResponseBody {
	s.Success = &v
	return s
}

func (s *ListEndpointsByIpResponseBody) SetRequestId(v string) *ListEndpointsByIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEndpointsByIpResponseBody) SetEndpointList(v *ListEndpointsByIpResponseBodyEndpointList) *ListEndpointsByIpResponseBody {
	s.EndpointList = v
	return s
}

type ListEndpointsByIpResponseBodyEndpointList struct {
	ItemEndpoint []*ListEndpointsByIpResponseBodyEndpointListItemEndpoint `json:"ItemEndpoint,omitempty" xml:"ItemEndpoint,omitempty" type:"Repeated"`
}

func (s ListEndpointsByIpResponseBodyEndpointList) String() string {
	return tea.Prettify(s)
}

func (s ListEndpointsByIpResponseBodyEndpointList) GoString() string {
	return s.String()
}

func (s *ListEndpointsByIpResponseBodyEndpointList) SetItemEndpoint(v []*ListEndpointsByIpResponseBodyEndpointListItemEndpoint) *ListEndpointsByIpResponseBodyEndpointList {
	s.ItemEndpoint = v
	return s
}

type ListEndpointsByIpResponseBodyEndpointListItemEndpoint struct {
	Type      *string                                                         `json:"Type,omitempty" xml:"Type,omitempty"`
	Namespace *string                                                         `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Product   *string                                                         `json:"Product,omitempty" xml:"Product,omitempty"`
	Id        *string                                                         `json:"Id,omitempty" xml:"Id,omitempty"`
	Endpoint  *string                                                         `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	Protocols *ListEndpointsByIpResponseBodyEndpointListItemEndpointProtocols `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Struct"`
}

func (s ListEndpointsByIpResponseBodyEndpointListItemEndpoint) String() string {
	return tea.Prettify(s)
}

func (s ListEndpointsByIpResponseBodyEndpointListItemEndpoint) GoString() string {
	return s.String()
}

func (s *ListEndpointsByIpResponseBodyEndpointListItemEndpoint) SetType(v string) *ListEndpointsByIpResponseBodyEndpointListItemEndpoint {
	s.Type = &v
	return s
}

func (s *ListEndpointsByIpResponseBodyEndpointListItemEndpoint) SetNamespace(v string) *ListEndpointsByIpResponseBodyEndpointListItemEndpoint {
	s.Namespace = &v
	return s
}

func (s *ListEndpointsByIpResponseBodyEndpointListItemEndpoint) SetProduct(v string) *ListEndpointsByIpResponseBodyEndpointListItemEndpoint {
	s.Product = &v
	return s
}

func (s *ListEndpointsByIpResponseBodyEndpointListItemEndpoint) SetId(v string) *ListEndpointsByIpResponseBodyEndpointListItemEndpoint {
	s.Id = &v
	return s
}

func (s *ListEndpointsByIpResponseBodyEndpointListItemEndpoint) SetEndpoint(v string) *ListEndpointsByIpResponseBodyEndpointListItemEndpoint {
	s.Endpoint = &v
	return s
}

func (s *ListEndpointsByIpResponseBodyEndpointListItemEndpoint) SetProtocols(v *ListEndpointsByIpResponseBodyEndpointListItemEndpointProtocols) *ListEndpointsByIpResponseBodyEndpointListItemEndpoint {
	s.Protocols = v
	return s
}

type ListEndpointsByIpResponseBodyEndpointListItemEndpointProtocols struct {
	Protocols []*string `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Repeated"`
}

func (s ListEndpointsByIpResponseBodyEndpointListItemEndpointProtocols) String() string {
	return tea.Prettify(s)
}

func (s ListEndpointsByIpResponseBodyEndpointListItemEndpointProtocols) GoString() string {
	return s.String()
}

func (s *ListEndpointsByIpResponseBodyEndpointListItemEndpointProtocols) SetProtocols(v []*string) *ListEndpointsByIpResponseBodyEndpointListItemEndpointProtocols {
	s.Protocols = v
	return s
}

type ListEndpointsByIpResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListEndpointsByIpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEndpointsByIpResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEndpointsByIpResponse) GoString() string {
	return s.String()
}

func (s *ListEndpointsByIpResponse) SetHeaders(v map[string]*string) *ListEndpointsByIpResponse {
	s.Headers = v
	return s
}

func (s *ListEndpointsByIpResponse) SetBody(v *ListEndpointsByIpResponseBody) *ListEndpointsByIpResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("location"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) DescribeEndpointWithOptions(request *DescribeEndpointRequest, runtime *util.RuntimeOptions) (_result *DescribeEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeEndpointResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeEndpoint"), tea.String("2015-06-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEndpoint(request *DescribeEndpointRequest) (_result *DescribeEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEndpointResponse{}
	_body, _err := client.DescribeEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEndpointsWithOptions(request *DescribeEndpointsRequest, runtime *util.RuntimeOptions) (_result *DescribeEndpointsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeEndpointsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeEndpoints"), tea.String("2015-06-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEndpoints(request *DescribeEndpointsRequest) (_result *DescribeEndpointsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEndpointsResponse{}
	_body, _err := client.DescribeEndpointsWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DescribeRegions"), tea.String("2015-06-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeServicesWithOptions(request *DescribeServicesRequest, runtime *util.RuntimeOptions) (_result *DescribeServicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeServicesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeServices"), tea.String("2015-06-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeServices(request *DescribeServicesRequest) (_result *DescribeServicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeServicesResponse{}
	_body, _err := client.DescribeServicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEndpointsWithOptions(request *ListEndpointsRequest, runtime *util.RuntimeOptions) (_result *ListEndpointsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListEndpointsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListEndpoints"), tea.String("2015-06-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEndpoints(request *ListEndpointsRequest) (_result *ListEndpointsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEndpointsResponse{}
	_body, _err := client.ListEndpointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEndpointsByIpWithOptions(request *ListEndpointsByIpRequest, runtime *util.RuntimeOptions) (_result *ListEndpointsByIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListEndpointsByIpResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListEndpointsByIp"), tea.String("2015-06-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEndpointsByIp(request *ListEndpointsByIpRequest) (_result *ListEndpointsByIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEndpointsByIpResponse{}
	_body, _err := client.ListEndpointsByIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
