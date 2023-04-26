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

type AssignPrivateIpAddressRequest struct {
	AssignMac          *bool   `json:"AssignMac,omitempty" xml:"AssignMac,omitempty"`
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	PrivateIpAddress   *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SkipConfig         *bool   `json:"SkipConfig,omitempty" xml:"SkipConfig,omitempty"`
	SubnetId           *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
}

func (s AssignPrivateIpAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s AssignPrivateIpAddressRequest) GoString() string {
	return s.String()
}

func (s *AssignPrivateIpAddressRequest) SetAssignMac(v bool) *AssignPrivateIpAddressRequest {
	s.AssignMac = &v
	return s
}

func (s *AssignPrivateIpAddressRequest) SetNetworkInterfaceId(v string) *AssignPrivateIpAddressRequest {
	s.NetworkInterfaceId = &v
	return s
}

func (s *AssignPrivateIpAddressRequest) SetPrivateIpAddress(v string) *AssignPrivateIpAddressRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *AssignPrivateIpAddressRequest) SetRegionId(v string) *AssignPrivateIpAddressRequest {
	s.RegionId = &v
	return s
}

func (s *AssignPrivateIpAddressRequest) SetSkipConfig(v bool) *AssignPrivateIpAddressRequest {
	s.SkipConfig = &v
	return s
}

func (s *AssignPrivateIpAddressRequest) SetSubnetId(v string) *AssignPrivateIpAddressRequest {
	s.SubnetId = &v
	return s
}

type AssignPrivateIpAddressResponseBody struct {
	Code      *int32                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   *AssignPrivateIpAddressResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssignPrivateIpAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssignPrivateIpAddressResponseBody) GoString() string {
	return s.String()
}

func (s *AssignPrivateIpAddressResponseBody) SetCode(v int32) *AssignPrivateIpAddressResponseBody {
	s.Code = &v
	return s
}

func (s *AssignPrivateIpAddressResponseBody) SetContent(v *AssignPrivateIpAddressResponseBodyContent) *AssignPrivateIpAddressResponseBody {
	s.Content = v
	return s
}

func (s *AssignPrivateIpAddressResponseBody) SetMessage(v string) *AssignPrivateIpAddressResponseBody {
	s.Message = &v
	return s
}

func (s *AssignPrivateIpAddressResponseBody) SetRequestId(v string) *AssignPrivateIpAddressResponseBody {
	s.RequestId = &v
	return s
}

type AssignPrivateIpAddressResponseBodyContent struct {
	IpName             *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
}

func (s AssignPrivateIpAddressResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s AssignPrivateIpAddressResponseBodyContent) GoString() string {
	return s.String()
}

func (s *AssignPrivateIpAddressResponseBodyContent) SetIpName(v string) *AssignPrivateIpAddressResponseBodyContent {
	s.IpName = &v
	return s
}

func (s *AssignPrivateIpAddressResponseBodyContent) SetNetworkInterfaceId(v string) *AssignPrivateIpAddressResponseBodyContent {
	s.NetworkInterfaceId = &v
	return s
}

type AssignPrivateIpAddressResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AssignPrivateIpAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AssignPrivateIpAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s AssignPrivateIpAddressResponse) GoString() string {
	return s.String()
}

func (s *AssignPrivateIpAddressResponse) SetHeaders(v map[string]*string) *AssignPrivateIpAddressResponse {
	s.Headers = v
	return s
}

func (s *AssignPrivateIpAddressResponse) SetStatusCode(v int32) *AssignPrivateIpAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *AssignPrivateIpAddressResponse) SetBody(v *AssignPrivateIpAddressResponseBody) *AssignPrivateIpAddressResponse {
	s.Body = v
	return s
}

type CreateErRequest struct {
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ErName       *string `json:"ErName,omitempty" xml:"ErName,omitempty"`
	MasterZoneId *string `json:"MasterZoneId,omitempty" xml:"MasterZoneId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateErRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateErRequest) GoString() string {
	return s.String()
}

func (s *CreateErRequest) SetDescription(v string) *CreateErRequest {
	s.Description = &v
	return s
}

func (s *CreateErRequest) SetErName(v string) *CreateErRequest {
	s.ErName = &v
	return s
}

func (s *CreateErRequest) SetMasterZoneId(v string) *CreateErRequest {
	s.MasterZoneId = &v
	return s
}

func (s *CreateErRequest) SetRegionId(v string) *CreateErRequest {
	s.RegionId = &v
	return s
}

type CreateErResponseBody struct {
	Code      *int32                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   *CreateErResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateErResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateErResponseBody) GoString() string {
	return s.String()
}

func (s *CreateErResponseBody) SetCode(v int32) *CreateErResponseBody {
	s.Code = &v
	return s
}

func (s *CreateErResponseBody) SetContent(v *CreateErResponseBodyContent) *CreateErResponseBody {
	s.Content = v
	return s
}

func (s *CreateErResponseBody) SetMessage(v string) *CreateErResponseBody {
	s.Message = &v
	return s
}

func (s *CreateErResponseBody) SetRequestId(v string) *CreateErResponseBody {
	s.RequestId = &v
	return s
}

type CreateErResponseBodyContent struct {
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
}

func (s CreateErResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s CreateErResponseBodyContent) GoString() string {
	return s.String()
}

func (s *CreateErResponseBodyContent) SetErId(v string) *CreateErResponseBodyContent {
	s.ErId = &v
	return s
}

type CreateErResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateErResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateErResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateErResponse) GoString() string {
	return s.String()
}

func (s *CreateErResponse) SetHeaders(v map[string]*string) *CreateErResponse {
	s.Headers = v
	return s
}

func (s *CreateErResponse) SetStatusCode(v int32) *CreateErResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateErResponse) SetBody(v *CreateErResponseBody) *CreateErResponse {
	s.Body = v
	return s
}

type CreateErAttachmentRequest struct {
	AutoReceiveAllRoute *bool   `json:"AutoReceiveAllRoute,omitempty" xml:"AutoReceiveAllRoute,omitempty"`
	ErAttachmentName    *string `json:"ErAttachmentName,omitempty" xml:"ErAttachmentName,omitempty"`
	ErId                *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType        *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceTenantId    *string `json:"ResourceTenantId,omitempty" xml:"ResourceTenantId,omitempty"`
}

func (s CreateErAttachmentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateErAttachmentRequest) GoString() string {
	return s.String()
}

func (s *CreateErAttachmentRequest) SetAutoReceiveAllRoute(v bool) *CreateErAttachmentRequest {
	s.AutoReceiveAllRoute = &v
	return s
}

func (s *CreateErAttachmentRequest) SetErAttachmentName(v string) *CreateErAttachmentRequest {
	s.ErAttachmentName = &v
	return s
}

func (s *CreateErAttachmentRequest) SetErId(v string) *CreateErAttachmentRequest {
	s.ErId = &v
	return s
}

func (s *CreateErAttachmentRequest) SetInstanceId(v string) *CreateErAttachmentRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateErAttachmentRequest) SetInstanceType(v string) *CreateErAttachmentRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateErAttachmentRequest) SetRegionId(v string) *CreateErAttachmentRequest {
	s.RegionId = &v
	return s
}

func (s *CreateErAttachmentRequest) SetResourceTenantId(v string) *CreateErAttachmentRequest {
	s.ResourceTenantId = &v
	return s
}

type CreateErAttachmentResponseBody struct {
	Code      *int32                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   *CreateErAttachmentResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateErAttachmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateErAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateErAttachmentResponseBody) SetCode(v int32) *CreateErAttachmentResponseBody {
	s.Code = &v
	return s
}

func (s *CreateErAttachmentResponseBody) SetContent(v *CreateErAttachmentResponseBodyContent) *CreateErAttachmentResponseBody {
	s.Content = v
	return s
}

func (s *CreateErAttachmentResponseBody) SetMessage(v string) *CreateErAttachmentResponseBody {
	s.Message = &v
	return s
}

func (s *CreateErAttachmentResponseBody) SetRequestId(v string) *CreateErAttachmentResponseBody {
	s.RequestId = &v
	return s
}

type CreateErAttachmentResponseBodyContent struct {
	ErAttachmentId *string `json:"ErAttachmentId,omitempty" xml:"ErAttachmentId,omitempty"`
}

func (s CreateErAttachmentResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s CreateErAttachmentResponseBodyContent) GoString() string {
	return s.String()
}

func (s *CreateErAttachmentResponseBodyContent) SetErAttachmentId(v string) *CreateErAttachmentResponseBodyContent {
	s.ErAttachmentId = &v
	return s
}

type CreateErAttachmentResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateErAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateErAttachmentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateErAttachmentResponse) GoString() string {
	return s.String()
}

func (s *CreateErAttachmentResponse) SetHeaders(v map[string]*string) *CreateErAttachmentResponse {
	s.Headers = v
	return s
}

func (s *CreateErAttachmentResponse) SetStatusCode(v int32) *CreateErAttachmentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateErAttachmentResponse) SetBody(v *CreateErAttachmentResponseBody) *CreateErAttachmentResponse {
	s.Body = v
	return s
}

type CreateErRouteMapRequest struct {
	Description               *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DestinationCidrBlock      *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	ErId                      *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	ReceptionInstanceId       *string `json:"ReceptionInstanceId,omitempty" xml:"ReceptionInstanceId,omitempty"`
	ReceptionInstanceOwner    *string `json:"ReceptionInstanceOwner,omitempty" xml:"ReceptionInstanceOwner,omitempty"`
	ReceptionInstanceType     *string `json:"ReceptionInstanceType,omitempty" xml:"ReceptionInstanceType,omitempty"`
	RegionId                  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RouteMapAction            *string `json:"RouteMapAction,omitempty" xml:"RouteMapAction,omitempty"`
	RouteMapNum               *int32  `json:"RouteMapNum,omitempty" xml:"RouteMapNum,omitempty"`
	TransmissionInstanceId    *string `json:"TransmissionInstanceId,omitempty" xml:"TransmissionInstanceId,omitempty"`
	TransmissionInstanceOwner *string `json:"TransmissionInstanceOwner,omitempty" xml:"TransmissionInstanceOwner,omitempty"`
	TransmissionInstanceType  *string `json:"TransmissionInstanceType,omitempty" xml:"TransmissionInstanceType,omitempty"`
}

func (s CreateErRouteMapRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateErRouteMapRequest) GoString() string {
	return s.String()
}

func (s *CreateErRouteMapRequest) SetDescription(v string) *CreateErRouteMapRequest {
	s.Description = &v
	return s
}

func (s *CreateErRouteMapRequest) SetDestinationCidrBlock(v string) *CreateErRouteMapRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *CreateErRouteMapRequest) SetErId(v string) *CreateErRouteMapRequest {
	s.ErId = &v
	return s
}

func (s *CreateErRouteMapRequest) SetReceptionInstanceId(v string) *CreateErRouteMapRequest {
	s.ReceptionInstanceId = &v
	return s
}

func (s *CreateErRouteMapRequest) SetReceptionInstanceOwner(v string) *CreateErRouteMapRequest {
	s.ReceptionInstanceOwner = &v
	return s
}

func (s *CreateErRouteMapRequest) SetReceptionInstanceType(v string) *CreateErRouteMapRequest {
	s.ReceptionInstanceType = &v
	return s
}

func (s *CreateErRouteMapRequest) SetRegionId(v string) *CreateErRouteMapRequest {
	s.RegionId = &v
	return s
}

func (s *CreateErRouteMapRequest) SetRouteMapAction(v string) *CreateErRouteMapRequest {
	s.RouteMapAction = &v
	return s
}

func (s *CreateErRouteMapRequest) SetRouteMapNum(v int32) *CreateErRouteMapRequest {
	s.RouteMapNum = &v
	return s
}

func (s *CreateErRouteMapRequest) SetTransmissionInstanceId(v string) *CreateErRouteMapRequest {
	s.TransmissionInstanceId = &v
	return s
}

func (s *CreateErRouteMapRequest) SetTransmissionInstanceOwner(v string) *CreateErRouteMapRequest {
	s.TransmissionInstanceOwner = &v
	return s
}

func (s *CreateErRouteMapRequest) SetTransmissionInstanceType(v string) *CreateErRouteMapRequest {
	s.TransmissionInstanceType = &v
	return s
}

type CreateErRouteMapResponseBody struct {
	Code      *int32                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   *CreateErRouteMapResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateErRouteMapResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateErRouteMapResponseBody) GoString() string {
	return s.String()
}

func (s *CreateErRouteMapResponseBody) SetCode(v int32) *CreateErRouteMapResponseBody {
	s.Code = &v
	return s
}

func (s *CreateErRouteMapResponseBody) SetContent(v *CreateErRouteMapResponseBodyContent) *CreateErRouteMapResponseBody {
	s.Content = v
	return s
}

func (s *CreateErRouteMapResponseBody) SetMessage(v string) *CreateErRouteMapResponseBody {
	s.Message = &v
	return s
}

func (s *CreateErRouteMapResponseBody) SetRequestId(v string) *CreateErRouteMapResponseBody {
	s.RequestId = &v
	return s
}

type CreateErRouteMapResponseBodyContent struct {
	ErRouteMapId *string `json:"ErRouteMapId,omitempty" xml:"ErRouteMapId,omitempty"`
}

func (s CreateErRouteMapResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s CreateErRouteMapResponseBodyContent) GoString() string {
	return s.String()
}

func (s *CreateErRouteMapResponseBodyContent) SetErRouteMapId(v string) *CreateErRouteMapResponseBodyContent {
	s.ErRouteMapId = &v
	return s
}

type CreateErRouteMapResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateErRouteMapResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateErRouteMapResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateErRouteMapResponse) GoString() string {
	return s.String()
}

func (s *CreateErRouteMapResponse) SetHeaders(v map[string]*string) *CreateErRouteMapResponse {
	s.Headers = v
	return s
}

func (s *CreateErRouteMapResponse) SetStatusCode(v int32) *CreateErRouteMapResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateErRouteMapResponse) SetBody(v *CreateErRouteMapResponseBody) *CreateErRouteMapResponse {
	s.Body = v
	return s
}

type CreateSubnetRequest struct {
	Cidr       *string                   `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	RegionId   *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SubnetName *string                   `json:"SubnetName,omitempty" xml:"SubnetName,omitempty"`
	Tag        []*CreateSubnetRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	Type       *string                   `json:"Type,omitempty" xml:"Type,omitempty"`
	VpdId      *string                   `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	ZoneId     *string                   `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateSubnetRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSubnetRequest) GoString() string {
	return s.String()
}

func (s *CreateSubnetRequest) SetCidr(v string) *CreateSubnetRequest {
	s.Cidr = &v
	return s
}

func (s *CreateSubnetRequest) SetRegionId(v string) *CreateSubnetRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSubnetRequest) SetSubnetName(v string) *CreateSubnetRequest {
	s.SubnetName = &v
	return s
}

func (s *CreateSubnetRequest) SetTag(v []*CreateSubnetRequestTag) *CreateSubnetRequest {
	s.Tag = v
	return s
}

func (s *CreateSubnetRequest) SetType(v string) *CreateSubnetRequest {
	s.Type = &v
	return s
}

func (s *CreateSubnetRequest) SetVpdId(v string) *CreateSubnetRequest {
	s.VpdId = &v
	return s
}

func (s *CreateSubnetRequest) SetZoneId(v string) *CreateSubnetRequest {
	s.ZoneId = &v
	return s
}

type CreateSubnetRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateSubnetRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateSubnetRequestTag) GoString() string {
	return s.String()
}

func (s *CreateSubnetRequestTag) SetKey(v string) *CreateSubnetRequestTag {
	s.Key = &v
	return s
}

func (s *CreateSubnetRequestTag) SetValue(v string) *CreateSubnetRequestTag {
	s.Value = &v
	return s
}

type CreateSubnetResponseBody struct {
	Code      *int32                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   *CreateSubnetResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSubnetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSubnetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSubnetResponseBody) SetCode(v int32) *CreateSubnetResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSubnetResponseBody) SetContent(v *CreateSubnetResponseBodyContent) *CreateSubnetResponseBody {
	s.Content = v
	return s
}

func (s *CreateSubnetResponseBody) SetMessage(v string) *CreateSubnetResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSubnetResponseBody) SetRequestId(v string) *CreateSubnetResponseBody {
	s.RequestId = &v
	return s
}

type CreateSubnetResponseBodyContent struct {
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
}

func (s CreateSubnetResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s CreateSubnetResponseBodyContent) GoString() string {
	return s.String()
}

func (s *CreateSubnetResponseBodyContent) SetSubnetId(v string) *CreateSubnetResponseBodyContent {
	s.SubnetId = &v
	return s
}

type CreateSubnetResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSubnetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSubnetResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSubnetResponse) GoString() string {
	return s.String()
}

func (s *CreateSubnetResponse) SetHeaders(v map[string]*string) *CreateSubnetResponse {
	s.Headers = v
	return s
}

func (s *CreateSubnetResponse) SetStatusCode(v int32) *CreateSubnetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSubnetResponse) SetBody(v *CreateSubnetResponseBody) *CreateSubnetResponse {
	s.Body = v
	return s
}

type CreateVccRequest struct {
	AccessCouldService *bool                  `json:"AccessCouldService,omitempty" xml:"AccessCouldService,omitempty"`
	Bandwidth          *int32                 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	BgpCidr            *string                `json:"BgpCidr,omitempty" xml:"BgpCidr,omitempty"`
	CenId              *string                `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ConnectionType     *string                `json:"ConnectionType,omitempty" xml:"ConnectionType,omitempty"`
	Description        *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	RegionId           *string                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId    *string                `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag                []*CreateVccRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	VSwitchId          *string                `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VccId              *string                `json:"VccId,omitempty" xml:"VccId,omitempty"`
	VccName            *string                `json:"VccName,omitempty" xml:"VccName,omitempty"`
	VpcId              *string                `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpdId              *string                `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	ZoneId             *string                `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateVccRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVccRequest) GoString() string {
	return s.String()
}

func (s *CreateVccRequest) SetAccessCouldService(v bool) *CreateVccRequest {
	s.AccessCouldService = &v
	return s
}

func (s *CreateVccRequest) SetBandwidth(v int32) *CreateVccRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateVccRequest) SetBgpCidr(v string) *CreateVccRequest {
	s.BgpCidr = &v
	return s
}

func (s *CreateVccRequest) SetCenId(v string) *CreateVccRequest {
	s.CenId = &v
	return s
}

func (s *CreateVccRequest) SetConnectionType(v string) *CreateVccRequest {
	s.ConnectionType = &v
	return s
}

func (s *CreateVccRequest) SetDescription(v string) *CreateVccRequest {
	s.Description = &v
	return s
}

func (s *CreateVccRequest) SetRegionId(v string) *CreateVccRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVccRequest) SetResourceGroupId(v string) *CreateVccRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateVccRequest) SetTag(v []*CreateVccRequestTag) *CreateVccRequest {
	s.Tag = v
	return s
}

func (s *CreateVccRequest) SetVSwitchId(v string) *CreateVccRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateVccRequest) SetVccId(v string) *CreateVccRequest {
	s.VccId = &v
	return s
}

func (s *CreateVccRequest) SetVccName(v string) *CreateVccRequest {
	s.VccName = &v
	return s
}

func (s *CreateVccRequest) SetVpcId(v string) *CreateVccRequest {
	s.VpcId = &v
	return s
}

func (s *CreateVccRequest) SetVpdId(v string) *CreateVccRequest {
	s.VpdId = &v
	return s
}

func (s *CreateVccRequest) SetZoneId(v string) *CreateVccRequest {
	s.ZoneId = &v
	return s
}

type CreateVccRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateVccRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateVccRequestTag) GoString() string {
	return s.String()
}

func (s *CreateVccRequestTag) SetKey(v string) *CreateVccRequestTag {
	s.Key = &v
	return s
}

func (s *CreateVccRequestTag) SetValue(v string) *CreateVccRequestTag {
	s.Value = &v
	return s
}

type CreateVccResponseBody struct {
	Code      *int32                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   *CreateVccResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVccResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVccResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVccResponseBody) SetCode(v int32) *CreateVccResponseBody {
	s.Code = &v
	return s
}

func (s *CreateVccResponseBody) SetContent(v *CreateVccResponseBodyContent) *CreateVccResponseBody {
	s.Content = v
	return s
}

func (s *CreateVccResponseBody) SetMessage(v string) *CreateVccResponseBody {
	s.Message = &v
	return s
}

func (s *CreateVccResponseBody) SetRequestId(v string) *CreateVccResponseBody {
	s.RequestId = &v
	return s
}

type CreateVccResponseBodyContent struct {
	VccId *string `json:"VccId,omitempty" xml:"VccId,omitempty"`
}

func (s CreateVccResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s CreateVccResponseBodyContent) GoString() string {
	return s.String()
}

func (s *CreateVccResponseBodyContent) SetVccId(v string) *CreateVccResponseBodyContent {
	s.VccId = &v
	return s
}

type CreateVccResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateVccResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVccResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVccResponse) GoString() string {
	return s.String()
}

func (s *CreateVccResponse) SetHeaders(v map[string]*string) *CreateVccResponse {
	s.Headers = v
	return s
}

func (s *CreateVccResponse) SetStatusCode(v int32) *CreateVccResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVccResponse) SetBody(v *CreateVccResponseBody) *CreateVccResponse {
	s.Body = v
	return s
}

type CreateVccGrantRuleRequest struct {
	ErId          *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	GrantTenantId *string `json:"GrantTenantId,omitempty" xml:"GrantTenantId,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateVccGrantRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVccGrantRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateVccGrantRuleRequest) SetErId(v string) *CreateVccGrantRuleRequest {
	s.ErId = &v
	return s
}

func (s *CreateVccGrantRuleRequest) SetGrantTenantId(v string) *CreateVccGrantRuleRequest {
	s.GrantTenantId = &v
	return s
}

func (s *CreateVccGrantRuleRequest) SetInstanceId(v string) *CreateVccGrantRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateVccGrantRuleRequest) SetRegionId(v string) *CreateVccGrantRuleRequest {
	s.RegionId = &v
	return s
}

type CreateVccGrantRuleResponseBody struct {
	Code      *int32                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   *CreateVccGrantRuleResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVccGrantRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVccGrantRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVccGrantRuleResponseBody) SetCode(v int32) *CreateVccGrantRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateVccGrantRuleResponseBody) SetContent(v *CreateVccGrantRuleResponseBodyContent) *CreateVccGrantRuleResponseBody {
	s.Content = v
	return s
}

func (s *CreateVccGrantRuleResponseBody) SetMessage(v string) *CreateVccGrantRuleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateVccGrantRuleResponseBody) SetRequestId(v string) *CreateVccGrantRuleResponseBody {
	s.RequestId = &v
	return s
}

type CreateVccGrantRuleResponseBodyContent struct {
	GrantRuleId *string `json:"GrantRuleId,omitempty" xml:"GrantRuleId,omitempty"`
}

func (s CreateVccGrantRuleResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s CreateVccGrantRuleResponseBodyContent) GoString() string {
	return s.String()
}

func (s *CreateVccGrantRuleResponseBodyContent) SetGrantRuleId(v string) *CreateVccGrantRuleResponseBodyContent {
	s.GrantRuleId = &v
	return s
}

type CreateVccGrantRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateVccGrantRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVccGrantRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVccGrantRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateVccGrantRuleResponse) SetHeaders(v map[string]*string) *CreateVccGrantRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateVccGrantRuleResponse) SetStatusCode(v int32) *CreateVccGrantRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVccGrantRuleResponse) SetBody(v *CreateVccGrantRuleResponseBody) *CreateVccGrantRuleResponse {
	s.Body = v
	return s
}

type CreateVccRouteEntryRequest struct {
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VccId                *string `json:"VccId,omitempty" xml:"VccId,omitempty"`
}

func (s CreateVccRouteEntryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVccRouteEntryRequest) GoString() string {
	return s.String()
}

func (s *CreateVccRouteEntryRequest) SetDestinationCidrBlock(v string) *CreateVccRouteEntryRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *CreateVccRouteEntryRequest) SetRegionId(v string) *CreateVccRouteEntryRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVccRouteEntryRequest) SetVccId(v string) *CreateVccRouteEntryRequest {
	s.VccId = &v
	return s
}

type CreateVccRouteEntryResponseBody struct {
	Code      *int32                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   *CreateVccRouteEntryResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVccRouteEntryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVccRouteEntryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVccRouteEntryResponseBody) SetCode(v int32) *CreateVccRouteEntryResponseBody {
	s.Code = &v
	return s
}

func (s *CreateVccRouteEntryResponseBody) SetContent(v *CreateVccRouteEntryResponseBodyContent) *CreateVccRouteEntryResponseBody {
	s.Content = v
	return s
}

func (s *CreateVccRouteEntryResponseBody) SetMessage(v string) *CreateVccRouteEntryResponseBody {
	s.Message = &v
	return s
}

func (s *CreateVccRouteEntryResponseBody) SetRequestId(v string) *CreateVccRouteEntryResponseBody {
	s.RequestId = &v
	return s
}

type CreateVccRouteEntryResponseBodyContent struct {
	VccRouteEntryId *string `json:"VccRouteEntryId,omitempty" xml:"VccRouteEntryId,omitempty"`
}

func (s CreateVccRouteEntryResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s CreateVccRouteEntryResponseBodyContent) GoString() string {
	return s.String()
}

func (s *CreateVccRouteEntryResponseBodyContent) SetVccRouteEntryId(v string) *CreateVccRouteEntryResponseBodyContent {
	s.VccRouteEntryId = &v
	return s
}

type CreateVccRouteEntryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateVccRouteEntryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVccRouteEntryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVccRouteEntryResponse) GoString() string {
	return s.String()
}

func (s *CreateVccRouteEntryResponse) SetHeaders(v map[string]*string) *CreateVccRouteEntryResponse {
	s.Headers = v
	return s
}

func (s *CreateVccRouteEntryResponse) SetStatusCode(v int32) *CreateVccRouteEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVccRouteEntryResponse) SetBody(v *CreateVccRouteEntryResponseBody) *CreateVccRouteEntryResponse {
	s.Body = v
	return s
}

type CreateVpdRequest struct {
	Cidr            *string                    `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	RegionId        *string                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string                    `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Subnets         []*CreateVpdRequestSubnets `json:"Subnets,omitempty" xml:"Subnets,omitempty" type:"Repeated"`
	Tag             []*CreateVpdRequestTag     `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	VpdName         *string                    `json:"VpdName,omitempty" xml:"VpdName,omitempty"`
}

func (s CreateVpdRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVpdRequest) GoString() string {
	return s.String()
}

func (s *CreateVpdRequest) SetCidr(v string) *CreateVpdRequest {
	s.Cidr = &v
	return s
}

func (s *CreateVpdRequest) SetRegionId(v string) *CreateVpdRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVpdRequest) SetResourceGroupId(v string) *CreateVpdRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateVpdRequest) SetSubnets(v []*CreateVpdRequestSubnets) *CreateVpdRequest {
	s.Subnets = v
	return s
}

func (s *CreateVpdRequest) SetTag(v []*CreateVpdRequestTag) *CreateVpdRequest {
	s.Tag = v
	return s
}

func (s *CreateVpdRequest) SetVpdName(v string) *CreateVpdRequest {
	s.VpdName = &v
	return s
}

type CreateVpdRequestSubnets struct {
	Cidr       *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SubnetName *string `json:"SubnetName,omitempty" xml:"SubnetName,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
	ZoneId     *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateVpdRequestSubnets) String() string {
	return tea.Prettify(s)
}

func (s CreateVpdRequestSubnets) GoString() string {
	return s.String()
}

func (s *CreateVpdRequestSubnets) SetCidr(v string) *CreateVpdRequestSubnets {
	s.Cidr = &v
	return s
}

func (s *CreateVpdRequestSubnets) SetRegionId(v string) *CreateVpdRequestSubnets {
	s.RegionId = &v
	return s
}

func (s *CreateVpdRequestSubnets) SetSubnetName(v string) *CreateVpdRequestSubnets {
	s.SubnetName = &v
	return s
}

func (s *CreateVpdRequestSubnets) SetType(v string) *CreateVpdRequestSubnets {
	s.Type = &v
	return s
}

func (s *CreateVpdRequestSubnets) SetZoneId(v string) *CreateVpdRequestSubnets {
	s.ZoneId = &v
	return s
}

type CreateVpdRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateVpdRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateVpdRequestTag) GoString() string {
	return s.String()
}

func (s *CreateVpdRequestTag) SetKey(v string) *CreateVpdRequestTag {
	s.Key = &v
	return s
}

func (s *CreateVpdRequestTag) SetValue(v string) *CreateVpdRequestTag {
	s.Value = &v
	return s
}

type CreateVpdResponseBody struct {
	Code      *int32                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   *CreateVpdResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVpdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVpdResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVpdResponseBody) SetCode(v int32) *CreateVpdResponseBody {
	s.Code = &v
	return s
}

func (s *CreateVpdResponseBody) SetContent(v *CreateVpdResponseBodyContent) *CreateVpdResponseBody {
	s.Content = v
	return s
}

func (s *CreateVpdResponseBody) SetMessage(v string) *CreateVpdResponseBody {
	s.Message = &v
	return s
}

func (s *CreateVpdResponseBody) SetRequestId(v string) *CreateVpdResponseBody {
	s.RequestId = &v
	return s
}

type CreateVpdResponseBodyContent struct {
	SubnetIds []*string `json:"SubnetIds,omitempty" xml:"SubnetIds,omitempty" type:"Repeated"`
	VpdId     *string   `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
}

func (s CreateVpdResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s CreateVpdResponseBodyContent) GoString() string {
	return s.String()
}

func (s *CreateVpdResponseBodyContent) SetSubnetIds(v []*string) *CreateVpdResponseBodyContent {
	s.SubnetIds = v
	return s
}

func (s *CreateVpdResponseBodyContent) SetVpdId(v string) *CreateVpdResponseBodyContent {
	s.VpdId = &v
	return s
}

type CreateVpdResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateVpdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVpdResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVpdResponse) GoString() string {
	return s.String()
}

func (s *CreateVpdResponse) SetHeaders(v map[string]*string) *CreateVpdResponse {
	s.Headers = v
	return s
}

func (s *CreateVpdResponse) SetStatusCode(v int32) *CreateVpdResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVpdResponse) SetBody(v *CreateVpdResponseBody) *CreateVpdResponse {
	s.Body = v
	return s
}

type CreateVpdGrantRuleRequest struct {
	ErId          *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	GrantTenantId *string `json:"GrantTenantId,omitempty" xml:"GrantTenantId,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateVpdGrantRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVpdGrantRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateVpdGrantRuleRequest) SetErId(v string) *CreateVpdGrantRuleRequest {
	s.ErId = &v
	return s
}

func (s *CreateVpdGrantRuleRequest) SetGrantTenantId(v string) *CreateVpdGrantRuleRequest {
	s.GrantTenantId = &v
	return s
}

func (s *CreateVpdGrantRuleRequest) SetInstanceId(v string) *CreateVpdGrantRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateVpdGrantRuleRequest) SetRegionId(v string) *CreateVpdGrantRuleRequest {
	s.RegionId = &v
	return s
}

type CreateVpdGrantRuleResponseBody struct {
	Code      *int32                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   *CreateVpdGrantRuleResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVpdGrantRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVpdGrantRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVpdGrantRuleResponseBody) SetCode(v int32) *CreateVpdGrantRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateVpdGrantRuleResponseBody) SetContent(v *CreateVpdGrantRuleResponseBodyContent) *CreateVpdGrantRuleResponseBody {
	s.Content = v
	return s
}

func (s *CreateVpdGrantRuleResponseBody) SetMessage(v string) *CreateVpdGrantRuleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateVpdGrantRuleResponseBody) SetRequestId(v string) *CreateVpdGrantRuleResponseBody {
	s.RequestId = &v
	return s
}

type CreateVpdGrantRuleResponseBodyContent struct {
	GrantRuleId *string `json:"GrantRuleId,omitempty" xml:"GrantRuleId,omitempty"`
}

func (s CreateVpdGrantRuleResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s CreateVpdGrantRuleResponseBodyContent) GoString() string {
	return s.String()
}

func (s *CreateVpdGrantRuleResponseBodyContent) SetGrantRuleId(v string) *CreateVpdGrantRuleResponseBodyContent {
	s.GrantRuleId = &v
	return s
}

type CreateVpdGrantRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateVpdGrantRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVpdGrantRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVpdGrantRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateVpdGrantRuleResponse) SetHeaders(v map[string]*string) *CreateVpdGrantRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateVpdGrantRuleResponse) SetStatusCode(v int32) *CreateVpdGrantRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVpdGrantRuleResponse) SetBody(v *CreateVpdGrantRuleResponseBody) *CreateVpdGrantRuleResponse {
	s.Body = v
	return s
}

type DeleteErRequest struct {
	ErId     *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteErRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteErRequest) GoString() string {
	return s.String()
}

func (s *DeleteErRequest) SetErId(v string) *DeleteErRequest {
	s.ErId = &v
	return s
}

func (s *DeleteErRequest) SetRegionId(v string) *DeleteErRequest {
	s.RegionId = &v
	return s
}

type DeleteErResponseBody struct {
	Code      *int32      `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	Message   *string     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteErResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteErResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteErResponseBody) SetCode(v int32) *DeleteErResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteErResponseBody) SetContent(v interface{}) *DeleteErResponseBody {
	s.Content = v
	return s
}

func (s *DeleteErResponseBody) SetMessage(v string) *DeleteErResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteErResponseBody) SetRequestId(v string) *DeleteErResponseBody {
	s.RequestId = &v
	return s
}

type DeleteErResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteErResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteErResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteErResponse) GoString() string {
	return s.String()
}

func (s *DeleteErResponse) SetHeaders(v map[string]*string) *DeleteErResponse {
	s.Headers = v
	return s
}

func (s *DeleteErResponse) SetStatusCode(v int32) *DeleteErResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteErResponse) SetBody(v *DeleteErResponseBody) *DeleteErResponse {
	s.Body = v
	return s
}

type DeleteErAttachmentRequest struct {
	ErAttachmentId *string `json:"ErAttachmentId,omitempty" xml:"ErAttachmentId,omitempty"`
	ErId           *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteErAttachmentRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteErAttachmentRequest) GoString() string {
	return s.String()
}

func (s *DeleteErAttachmentRequest) SetErAttachmentId(v string) *DeleteErAttachmentRequest {
	s.ErAttachmentId = &v
	return s
}

func (s *DeleteErAttachmentRequest) SetErId(v string) *DeleteErAttachmentRequest {
	s.ErId = &v
	return s
}

func (s *DeleteErAttachmentRequest) SetRegionId(v string) *DeleteErAttachmentRequest {
	s.RegionId = &v
	return s
}

type DeleteErAttachmentResponseBody struct {
	Code      *int32      `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	Message   *string     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteErAttachmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteErAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteErAttachmentResponseBody) SetCode(v int32) *DeleteErAttachmentResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteErAttachmentResponseBody) SetContent(v interface{}) *DeleteErAttachmentResponseBody {
	s.Content = v
	return s
}

func (s *DeleteErAttachmentResponseBody) SetMessage(v string) *DeleteErAttachmentResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteErAttachmentResponseBody) SetRequestId(v string) *DeleteErAttachmentResponseBody {
	s.RequestId = &v
	return s
}

type DeleteErAttachmentResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteErAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteErAttachmentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteErAttachmentResponse) GoString() string {
	return s.String()
}

func (s *DeleteErAttachmentResponse) SetHeaders(v map[string]*string) *DeleteErAttachmentResponse {
	s.Headers = v
	return s
}

func (s *DeleteErAttachmentResponse) SetStatusCode(v int32) *DeleteErAttachmentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteErAttachmentResponse) SetBody(v *DeleteErAttachmentResponseBody) *DeleteErAttachmentResponse {
	s.Body = v
	return s
}

type DeleteErRouteMapRequest struct {
	ErId          *string   `json:"ErId,omitempty" xml:"ErId,omitempty"`
	ErRouteMapIds []*string `json:"ErRouteMapIds,omitempty" xml:"ErRouteMapIds,omitempty" type:"Repeated"`
	RegionId      *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteErRouteMapRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteErRouteMapRequest) GoString() string {
	return s.String()
}

func (s *DeleteErRouteMapRequest) SetErId(v string) *DeleteErRouteMapRequest {
	s.ErId = &v
	return s
}

func (s *DeleteErRouteMapRequest) SetErRouteMapIds(v []*string) *DeleteErRouteMapRequest {
	s.ErRouteMapIds = v
	return s
}

func (s *DeleteErRouteMapRequest) SetRegionId(v string) *DeleteErRouteMapRequest {
	s.RegionId = &v
	return s
}

type DeleteErRouteMapResponseBody struct {
	Code      *int32      `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	Message   *string     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteErRouteMapResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteErRouteMapResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteErRouteMapResponseBody) SetCode(v int32) *DeleteErRouteMapResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteErRouteMapResponseBody) SetContent(v interface{}) *DeleteErRouteMapResponseBody {
	s.Content = v
	return s
}

func (s *DeleteErRouteMapResponseBody) SetMessage(v string) *DeleteErRouteMapResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteErRouteMapResponseBody) SetRequestId(v string) *DeleteErRouteMapResponseBody {
	s.RequestId = &v
	return s
}

type DeleteErRouteMapResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteErRouteMapResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteErRouteMapResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteErRouteMapResponse) GoString() string {
	return s.String()
}

func (s *DeleteErRouteMapResponse) SetHeaders(v map[string]*string) *DeleteErRouteMapResponse {
	s.Headers = v
	return s
}

func (s *DeleteErRouteMapResponse) SetStatusCode(v int32) *DeleteErRouteMapResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteErRouteMapResponse) SetBody(v *DeleteErRouteMapResponseBody) *DeleteErRouteMapResponse {
	s.Body = v
	return s
}

type DeleteSubnetRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	VpdId    *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	ZoneId   *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DeleteSubnetRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubnetRequest) GoString() string {
	return s.String()
}

func (s *DeleteSubnetRequest) SetRegionId(v string) *DeleteSubnetRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSubnetRequest) SetSubnetId(v string) *DeleteSubnetRequest {
	s.SubnetId = &v
	return s
}

func (s *DeleteSubnetRequest) SetVpdId(v string) *DeleteSubnetRequest {
	s.VpdId = &v
	return s
}

func (s *DeleteSubnetRequest) SetZoneId(v string) *DeleteSubnetRequest {
	s.ZoneId = &v
	return s
}

type DeleteSubnetResponseBody struct {
	Code      *int32      `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	Message   *string     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSubnetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubnetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSubnetResponseBody) SetCode(v int32) *DeleteSubnetResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSubnetResponseBody) SetContent(v interface{}) *DeleteSubnetResponseBody {
	s.Content = v
	return s
}

func (s *DeleteSubnetResponseBody) SetMessage(v string) *DeleteSubnetResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSubnetResponseBody) SetRequestId(v string) *DeleteSubnetResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSubnetResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSubnetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSubnetResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubnetResponse) GoString() string {
	return s.String()
}

func (s *DeleteSubnetResponse) SetHeaders(v map[string]*string) *DeleteSubnetResponse {
	s.Headers = v
	return s
}

func (s *DeleteSubnetResponse) SetStatusCode(v int32) *DeleteSubnetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSubnetResponse) SetBody(v *DeleteSubnetResponseBody) *DeleteSubnetResponse {
	s.Body = v
	return s
}

type DeleteVccGrantRuleRequest struct {
	ErId        *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	GrantRuleId *string `json:"GrantRuleId,omitempty" xml:"GrantRuleId,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteVccGrantRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVccGrantRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteVccGrantRuleRequest) SetErId(v string) *DeleteVccGrantRuleRequest {
	s.ErId = &v
	return s
}

func (s *DeleteVccGrantRuleRequest) SetGrantRuleId(v string) *DeleteVccGrantRuleRequest {
	s.GrantRuleId = &v
	return s
}

func (s *DeleteVccGrantRuleRequest) SetInstanceId(v string) *DeleteVccGrantRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteVccGrantRuleRequest) SetRegionId(v string) *DeleteVccGrantRuleRequest {
	s.RegionId = &v
	return s
}

type DeleteVccGrantRuleResponseBody struct {
	Code      *int32      `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	Message   *string     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVccGrantRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVccGrantRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVccGrantRuleResponseBody) SetCode(v int32) *DeleteVccGrantRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteVccGrantRuleResponseBody) SetContent(v interface{}) *DeleteVccGrantRuleResponseBody {
	s.Content = v
	return s
}

func (s *DeleteVccGrantRuleResponseBody) SetMessage(v string) *DeleteVccGrantRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteVccGrantRuleResponseBody) SetRequestId(v string) *DeleteVccGrantRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVccGrantRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteVccGrantRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVccGrantRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVccGrantRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteVccGrantRuleResponse) SetHeaders(v map[string]*string) *DeleteVccGrantRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteVccGrantRuleResponse) SetStatusCode(v int32) *DeleteVccGrantRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVccGrantRuleResponse) SetBody(v *DeleteVccGrantRuleResponseBody) *DeleteVccGrantRuleResponse {
	s.Body = v
	return s
}

type DeleteVccRouteEntryRequest struct {
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VccId                *string `json:"VccId,omitempty" xml:"VccId,omitempty"`
	VccRouteEntryId      *string `json:"VccRouteEntryId,omitempty" xml:"VccRouteEntryId,omitempty"`
}

func (s DeleteVccRouteEntryRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVccRouteEntryRequest) GoString() string {
	return s.String()
}

func (s *DeleteVccRouteEntryRequest) SetDestinationCidrBlock(v string) *DeleteVccRouteEntryRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DeleteVccRouteEntryRequest) SetRegionId(v string) *DeleteVccRouteEntryRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteVccRouteEntryRequest) SetVccId(v string) *DeleteVccRouteEntryRequest {
	s.VccId = &v
	return s
}

func (s *DeleteVccRouteEntryRequest) SetVccRouteEntryId(v string) *DeleteVccRouteEntryRequest {
	s.VccRouteEntryId = &v
	return s
}

type DeleteVccRouteEntryResponseBody struct {
	Code      *int32      `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	Message   *string     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVccRouteEntryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVccRouteEntryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVccRouteEntryResponseBody) SetCode(v int32) *DeleteVccRouteEntryResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteVccRouteEntryResponseBody) SetContent(v interface{}) *DeleteVccRouteEntryResponseBody {
	s.Content = v
	return s
}

func (s *DeleteVccRouteEntryResponseBody) SetMessage(v string) *DeleteVccRouteEntryResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteVccRouteEntryResponseBody) SetRequestId(v string) *DeleteVccRouteEntryResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVccRouteEntryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteVccRouteEntryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVccRouteEntryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVccRouteEntryResponse) GoString() string {
	return s.String()
}

func (s *DeleteVccRouteEntryResponse) SetHeaders(v map[string]*string) *DeleteVccRouteEntryResponse {
	s.Headers = v
	return s
}

func (s *DeleteVccRouteEntryResponse) SetStatusCode(v int32) *DeleteVccRouteEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVccRouteEntryResponse) SetBody(v *DeleteVccRouteEntryResponseBody) *DeleteVccRouteEntryResponse {
	s.Body = v
	return s
}

type DeleteVpdRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VpdId    *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
}

func (s DeleteVpdRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpdRequest) GoString() string {
	return s.String()
}

func (s *DeleteVpdRequest) SetRegionId(v string) *DeleteVpdRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteVpdRequest) SetVpdId(v string) *DeleteVpdRequest {
	s.VpdId = &v
	return s
}

type DeleteVpdResponseBody struct {
	Code      *int32      `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	Message   *string     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVpdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpdResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVpdResponseBody) SetCode(v int32) *DeleteVpdResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteVpdResponseBody) SetContent(v interface{}) *DeleteVpdResponseBody {
	s.Content = v
	return s
}

func (s *DeleteVpdResponseBody) SetMessage(v string) *DeleteVpdResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteVpdResponseBody) SetRequestId(v string) *DeleteVpdResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVpdResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteVpdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVpdResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpdResponse) GoString() string {
	return s.String()
}

func (s *DeleteVpdResponse) SetHeaders(v map[string]*string) *DeleteVpdResponse {
	s.Headers = v
	return s
}

func (s *DeleteVpdResponse) SetStatusCode(v int32) *DeleteVpdResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVpdResponse) SetBody(v *DeleteVpdResponseBody) *DeleteVpdResponse {
	s.Body = v
	return s
}

type DeleteVpdGrantRuleRequest struct {
	ErId          *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	GrantRuleId   *string `json:"GrantRuleId,omitempty" xml:"GrantRuleId,omitempty"`
	GrantTenantId *string `json:"GrantTenantId,omitempty" xml:"GrantTenantId,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteVpdGrantRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpdGrantRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteVpdGrantRuleRequest) SetErId(v string) *DeleteVpdGrantRuleRequest {
	s.ErId = &v
	return s
}

func (s *DeleteVpdGrantRuleRequest) SetGrantRuleId(v string) *DeleteVpdGrantRuleRequest {
	s.GrantRuleId = &v
	return s
}

func (s *DeleteVpdGrantRuleRequest) SetGrantTenantId(v string) *DeleteVpdGrantRuleRequest {
	s.GrantTenantId = &v
	return s
}

func (s *DeleteVpdGrantRuleRequest) SetInstanceId(v string) *DeleteVpdGrantRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteVpdGrantRuleRequest) SetRegionId(v string) *DeleteVpdGrantRuleRequest {
	s.RegionId = &v
	return s
}

type DeleteVpdGrantRuleResponseBody struct {
	Code      *int32      `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	Message   *string     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVpdGrantRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpdGrantRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVpdGrantRuleResponseBody) SetCode(v int32) *DeleteVpdGrantRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteVpdGrantRuleResponseBody) SetContent(v interface{}) *DeleteVpdGrantRuleResponseBody {
	s.Content = v
	return s
}

func (s *DeleteVpdGrantRuleResponseBody) SetMessage(v string) *DeleteVpdGrantRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteVpdGrantRuleResponseBody) SetRequestId(v string) *DeleteVpdGrantRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVpdGrantRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteVpdGrantRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVpdGrantRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpdGrantRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteVpdGrantRuleResponse) SetHeaders(v map[string]*string) *DeleteVpdGrantRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteVpdGrantRuleResponse) SetStatusCode(v int32) *DeleteVpdGrantRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVpdGrantRuleResponse) SetBody(v *DeleteVpdGrantRuleResponseBody) *DeleteVpdGrantRuleResponse {
	s.Body = v
	return s
}

type DescribeSlrRequest struct {
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeSlrRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlrRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlrRequest) SetResourceGroupId(v string) *DescribeSlrRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeSlrResponseBody struct {
	Code      *int32                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   *DescribeSlrResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSlrResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlrResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlrResponseBody) SetCode(v int32) *DescribeSlrResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSlrResponseBody) SetContent(v *DescribeSlrResponseBodyContent) *DescribeSlrResponseBody {
	s.Content = v
	return s
}

func (s *DescribeSlrResponseBody) SetMessage(v string) *DescribeSlrResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSlrResponseBody) SetRequestId(v string) *DescribeSlrResponseBody {
	s.RequestId = &v
	return s
}

type DescribeSlrResponseBodyContent struct {
	HasRole *bool `json:"HasRole,omitempty" xml:"HasRole,omitempty"`
}

func (s DescribeSlrResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlrResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DescribeSlrResponseBodyContent) SetHasRole(v bool) *DescribeSlrResponseBodyContent {
	s.HasRole = &v
	return s
}

type DescribeSlrResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSlrResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSlrResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlrResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlrResponse) SetHeaders(v map[string]*string) *DescribeSlrResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlrResponse) SetStatusCode(v int32) *DescribeSlrResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSlrResponse) SetBody(v *DescribeSlrResponseBody) *DescribeSlrResponse {
	s.Body = v
	return s
}

type GetErRequest struct {
	ErId     *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetErRequest) String() string {
	return tea.Prettify(s)
}

func (s GetErRequest) GoString() string {
	return s.String()
}

func (s *GetErRequest) SetErId(v string) *GetErRequest {
	s.ErId = &v
	return s
}

func (s *GetErRequest) SetRegionId(v string) *GetErRequest {
	s.RegionId = &v
	return s
}

type GetErResponseBody struct {
	Code      *int32                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   *GetErResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Message   *string                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetErResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetErResponseBody) GoString() string {
	return s.String()
}

func (s *GetErResponseBody) SetCode(v int32) *GetErResponseBody {
	s.Code = &v
	return s
}

func (s *GetErResponseBody) SetContent(v *GetErResponseBodyContent) *GetErResponseBody {
	s.Content = v
	return s
}

func (s *GetErResponseBody) SetMessage(v string) *GetErResponseBody {
	s.Message = &v
	return s
}

func (s *GetErResponseBody) SetRequestId(v string) *GetErResponseBody {
	s.RequestId = &v
	return s
}

type GetErResponseBodyContent struct {
	CreateTime    *string                                  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	ErAttachments []*GetErResponseBodyContentErAttachments `json:"ErAttachments,omitempty" xml:"ErAttachments,omitempty" type:"Repeated"`
	ErId          *string                                  `json:"ErId,omitempty" xml:"ErId,omitempty"`
	ErName        *string                                  `json:"ErName,omitempty" xml:"ErName,omitempty"`
	ErRouteEntrys []*GetErResponseBodyContentErRouteEntrys `json:"ErRouteEntrys,omitempty" xml:"ErRouteEntrys,omitempty" type:"Repeated"`
	ErRouteMaps   []*GetErResponseBodyContentErRouteMaps   `json:"ErRouteMaps,omitempty" xml:"ErRouteMaps,omitempty" type:"Repeated"`
	GmtModified   *string                                  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	MasterZoneId  *string                                  `json:"MasterZoneId,omitempty" xml:"MasterZoneId,omitempty"`
	Message       *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RegionId      *string                                  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status        *string                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	TenantId      *string                                  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetErResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s GetErResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetErResponseBodyContent) SetCreateTime(v string) *GetErResponseBodyContent {
	s.CreateTime = &v
	return s
}

func (s *GetErResponseBodyContent) SetDescription(v string) *GetErResponseBodyContent {
	s.Description = &v
	return s
}

func (s *GetErResponseBodyContent) SetErAttachments(v []*GetErResponseBodyContentErAttachments) *GetErResponseBodyContent {
	s.ErAttachments = v
	return s
}

func (s *GetErResponseBodyContent) SetErId(v string) *GetErResponseBodyContent {
	s.ErId = &v
	return s
}

func (s *GetErResponseBodyContent) SetErName(v string) *GetErResponseBodyContent {
	s.ErName = &v
	return s
}

func (s *GetErResponseBodyContent) SetErRouteEntrys(v []*GetErResponseBodyContentErRouteEntrys) *GetErResponseBodyContent {
	s.ErRouteEntrys = v
	return s
}

func (s *GetErResponseBodyContent) SetErRouteMaps(v []*GetErResponseBodyContentErRouteMaps) *GetErResponseBodyContent {
	s.ErRouteMaps = v
	return s
}

func (s *GetErResponseBodyContent) SetGmtModified(v string) *GetErResponseBodyContent {
	s.GmtModified = &v
	return s
}

func (s *GetErResponseBodyContent) SetMasterZoneId(v string) *GetErResponseBodyContent {
	s.MasterZoneId = &v
	return s
}

func (s *GetErResponseBodyContent) SetMessage(v string) *GetErResponseBodyContent {
	s.Message = &v
	return s
}

func (s *GetErResponseBodyContent) SetRegionId(v string) *GetErResponseBodyContent {
	s.RegionId = &v
	return s
}

func (s *GetErResponseBodyContent) SetStatus(v string) *GetErResponseBodyContent {
	s.Status = &v
	return s
}

func (s *GetErResponseBodyContent) SetTenantId(v string) *GetErResponseBodyContent {
	s.TenantId = &v
	return s
}

type GetErResponseBodyContentErAttachments struct {
	Across              *bool   `json:"Across,omitempty" xml:"Across,omitempty"`
	AutoReceiveAllRoute *bool   `json:"AutoReceiveAllRoute,omitempty" xml:"AutoReceiveAllRoute,omitempty"`
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ErAttachmentId      *string `json:"ErAttachmentId,omitempty" xml:"ErAttachmentId,omitempty"`
	ErAttachmentName    *string `json:"ErAttachmentName,omitempty" xml:"ErAttachmentName,omitempty"`
	ErId                *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	GmtModified         *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName        *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceType        *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Message             *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceTenantId    *string `json:"ResourceTenantId,omitempty" xml:"ResourceTenantId,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TenantId            *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetErResponseBodyContentErAttachments) String() string {
	return tea.Prettify(s)
}

func (s GetErResponseBodyContentErAttachments) GoString() string {
	return s.String()
}

func (s *GetErResponseBodyContentErAttachments) SetAcross(v bool) *GetErResponseBodyContentErAttachments {
	s.Across = &v
	return s
}

func (s *GetErResponseBodyContentErAttachments) SetAutoReceiveAllRoute(v bool) *GetErResponseBodyContentErAttachments {
	s.AutoReceiveAllRoute = &v
	return s
}

func (s *GetErResponseBodyContentErAttachments) SetCreateTime(v string) *GetErResponseBodyContentErAttachments {
	s.CreateTime = &v
	return s
}

func (s *GetErResponseBodyContentErAttachments) SetErAttachmentId(v string) *GetErResponseBodyContentErAttachments {
	s.ErAttachmentId = &v
	return s
}

func (s *GetErResponseBodyContentErAttachments) SetErAttachmentName(v string) *GetErResponseBodyContentErAttachments {
	s.ErAttachmentName = &v
	return s
}

func (s *GetErResponseBodyContentErAttachments) SetErId(v string) *GetErResponseBodyContentErAttachments {
	s.ErId = &v
	return s
}

func (s *GetErResponseBodyContentErAttachments) SetGmtModified(v string) *GetErResponseBodyContentErAttachments {
	s.GmtModified = &v
	return s
}

func (s *GetErResponseBodyContentErAttachments) SetInstanceId(v string) *GetErResponseBodyContentErAttachments {
	s.InstanceId = &v
	return s
}

func (s *GetErResponseBodyContentErAttachments) SetInstanceName(v string) *GetErResponseBodyContentErAttachments {
	s.InstanceName = &v
	return s
}

func (s *GetErResponseBodyContentErAttachments) SetInstanceType(v string) *GetErResponseBodyContentErAttachments {
	s.InstanceType = &v
	return s
}

func (s *GetErResponseBodyContentErAttachments) SetMessage(v string) *GetErResponseBodyContentErAttachments {
	s.Message = &v
	return s
}

func (s *GetErResponseBodyContentErAttachments) SetRegionId(v string) *GetErResponseBodyContentErAttachments {
	s.RegionId = &v
	return s
}

func (s *GetErResponseBodyContentErAttachments) SetResourceTenantId(v string) *GetErResponseBodyContentErAttachments {
	s.ResourceTenantId = &v
	return s
}

func (s *GetErResponseBodyContentErAttachments) SetStatus(v string) *GetErResponseBodyContentErAttachments {
	s.Status = &v
	return s
}

func (s *GetErResponseBodyContentErAttachments) SetTenantId(v string) *GetErResponseBodyContentErAttachments {
	s.TenantId = &v
	return s
}

type GetErResponseBodyContentErRouteEntrys struct {
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	ErId                 *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	ErRouteEntryId       *string `json:"ErRouteEntryId,omitempty" xml:"ErRouteEntryId,omitempty"`
	GmtModified          *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	NextHopId            *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	NextHopType          *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceTenantId     *string `json:"ResourceTenantId,omitempty" xml:"ResourceTenantId,omitempty"`
	RouteType            *string `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TenantId             *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetErResponseBodyContentErRouteEntrys) String() string {
	return tea.Prettify(s)
}

func (s GetErResponseBodyContentErRouteEntrys) GoString() string {
	return s.String()
}

func (s *GetErResponseBodyContentErRouteEntrys) SetDestinationCidrBlock(v string) *GetErResponseBodyContentErRouteEntrys {
	s.DestinationCidrBlock = &v
	return s
}

func (s *GetErResponseBodyContentErRouteEntrys) SetErId(v string) *GetErResponseBodyContentErRouteEntrys {
	s.ErId = &v
	return s
}

func (s *GetErResponseBodyContentErRouteEntrys) SetErRouteEntryId(v string) *GetErResponseBodyContentErRouteEntrys {
	s.ErRouteEntryId = &v
	return s
}

func (s *GetErResponseBodyContentErRouteEntrys) SetGmtModified(v string) *GetErResponseBodyContentErRouteEntrys {
	s.GmtModified = &v
	return s
}

func (s *GetErResponseBodyContentErRouteEntrys) SetNextHopId(v string) *GetErResponseBodyContentErRouteEntrys {
	s.NextHopId = &v
	return s
}

func (s *GetErResponseBodyContentErRouteEntrys) SetNextHopType(v string) *GetErResponseBodyContentErRouteEntrys {
	s.NextHopType = &v
	return s
}

func (s *GetErResponseBodyContentErRouteEntrys) SetRegionId(v string) *GetErResponseBodyContentErRouteEntrys {
	s.RegionId = &v
	return s
}

func (s *GetErResponseBodyContentErRouteEntrys) SetResourceTenantId(v string) *GetErResponseBodyContentErRouteEntrys {
	s.ResourceTenantId = &v
	return s
}

func (s *GetErResponseBodyContentErRouteEntrys) SetRouteType(v string) *GetErResponseBodyContentErRouteEntrys {
	s.RouteType = &v
	return s
}

func (s *GetErResponseBodyContentErRouteEntrys) SetStatus(v string) *GetErResponseBodyContentErRouteEntrys {
	s.Status = &v
	return s
}

func (s *GetErResponseBodyContentErRouteEntrys) SetTenantId(v string) *GetErResponseBodyContentErRouteEntrys {
	s.TenantId = &v
	return s
}

type GetErResponseBodyContentErRouteMaps struct {
	Action                    *string `json:"Action,omitempty" xml:"Action,omitempty"`
	CreateTime                *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description               *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DestinationCidrBlock      *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	ErId                      *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	ErRouteMapId              *string `json:"ErRouteMapId,omitempty" xml:"ErRouteMapId,omitempty"`
	ErRouteMapName            *string `json:"ErRouteMapName,omitempty" xml:"ErRouteMapName,omitempty"`
	GmtModified               *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Message                   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	ReceptionInstanceId       *string `json:"ReceptionInstanceId,omitempty" xml:"ReceptionInstanceId,omitempty"`
	ReceptionInstanceName     *string `json:"ReceptionInstanceName,omitempty" xml:"ReceptionInstanceName,omitempty"`
	ReceptionInstanceOwner    *string `json:"ReceptionInstanceOwner,omitempty" xml:"ReceptionInstanceOwner,omitempty"`
	ReceptionInstanceType     *string `json:"ReceptionInstanceType,omitempty" xml:"ReceptionInstanceType,omitempty"`
	RegionId                  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RouteMapNum               *int32  `json:"RouteMapNum,omitempty" xml:"RouteMapNum,omitempty"`
	Status                    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TenantId                  *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	TransmissionInstanceId    *string `json:"TransmissionInstanceId,omitempty" xml:"TransmissionInstanceId,omitempty"`
	TransmissionInstanceName  *string `json:"TransmissionInstanceName,omitempty" xml:"TransmissionInstanceName,omitempty"`
	TransmissionInstanceOwner *string `json:"TransmissionInstanceOwner,omitempty" xml:"TransmissionInstanceOwner,omitempty"`
	TransmissionInstanceType  *string `json:"TransmissionInstanceType,omitempty" xml:"TransmissionInstanceType,omitempty"`
}

func (s GetErResponseBodyContentErRouteMaps) String() string {
	return tea.Prettify(s)
}

func (s GetErResponseBodyContentErRouteMaps) GoString() string {
	return s.String()
}

func (s *GetErResponseBodyContentErRouteMaps) SetAction(v string) *GetErResponseBodyContentErRouteMaps {
	s.Action = &v
	return s
}

func (s *GetErResponseBodyContentErRouteMaps) SetCreateTime(v string) *GetErResponseBodyContentErRouteMaps {
	s.CreateTime = &v
	return s
}

func (s *GetErResponseBodyContentErRouteMaps) SetDescription(v string) *GetErResponseBodyContentErRouteMaps {
	s.Description = &v
	return s
}

func (s *GetErResponseBodyContentErRouteMaps) SetDestinationCidrBlock(v string) *GetErResponseBodyContentErRouteMaps {
	s.DestinationCidrBlock = &v
	return s
}

func (s *GetErResponseBodyContentErRouteMaps) SetErId(v string) *GetErResponseBodyContentErRouteMaps {
	s.ErId = &v
	return s
}

func (s *GetErResponseBodyContentErRouteMaps) SetErRouteMapId(v string) *GetErResponseBodyContentErRouteMaps {
	s.ErRouteMapId = &v
	return s
}

func (s *GetErResponseBodyContentErRouteMaps) SetErRouteMapName(v string) *GetErResponseBodyContentErRouteMaps {
	s.ErRouteMapName = &v
	return s
}

func (s *GetErResponseBodyContentErRouteMaps) SetGmtModified(v string) *GetErResponseBodyContentErRouteMaps {
	s.GmtModified = &v
	return s
}

func (s *GetErResponseBodyContentErRouteMaps) SetMessage(v string) *GetErResponseBodyContentErRouteMaps {
	s.Message = &v
	return s
}

func (s *GetErResponseBodyContentErRouteMaps) SetReceptionInstanceId(v string) *GetErResponseBodyContentErRouteMaps {
	s.ReceptionInstanceId = &v
	return s
}

func (s *GetErResponseBodyContentErRouteMaps) SetReceptionInstanceName(v string) *GetErResponseBodyContentErRouteMaps {
	s.ReceptionInstanceName = &v
	return s
}

func (s *GetErResponseBodyContentErRouteMaps) SetReceptionInstanceOwner(v string) *GetErResponseBodyContentErRouteMaps {
	s.ReceptionInstanceOwner = &v
	return s
}

func (s *GetErResponseBodyContentErRouteMaps) SetReceptionInstanceType(v string) *GetErResponseBodyContentErRouteMaps {
	s.ReceptionInstanceType = &v
	return s
}

func (s *GetErResponseBodyContentErRouteMaps) SetRegionId(v string) *GetErResponseBodyContentErRouteMaps {
	s.RegionId = &v
	return s
}

func (s *GetErResponseBodyContentErRouteMaps) SetRouteMapNum(v int32) *GetErResponseBodyContentErRouteMaps {
	s.RouteMapNum = &v
	return s
}

func (s *GetErResponseBodyContentErRouteMaps) SetStatus(v string) *GetErResponseBodyContentErRouteMaps {
	s.Status = &v
	return s
}

func (s *GetErResponseBodyContentErRouteMaps) SetTenantId(v string) *GetErResponseBodyContentErRouteMaps {
	s.TenantId = &v
	return s
}

func (s *GetErResponseBodyContentErRouteMaps) SetTransmissionInstanceId(v string) *GetErResponseBodyContentErRouteMaps {
	s.TransmissionInstanceId = &v
	return s
}

func (s *GetErResponseBodyContentErRouteMaps) SetTransmissionInstanceName(v string) *GetErResponseBodyContentErRouteMaps {
	s.TransmissionInstanceName = &v
	return s
}

func (s *GetErResponseBodyContentErRouteMaps) SetTransmissionInstanceOwner(v string) *GetErResponseBodyContentErRouteMaps {
	s.TransmissionInstanceOwner = &v
	return s
}

func (s *GetErResponseBodyContentErRouteMaps) SetTransmissionInstanceType(v string) *GetErResponseBodyContentErRouteMaps {
	s.TransmissionInstanceType = &v
	return s
}

type GetErResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetErResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetErResponse) String() string {
	return tea.Prettify(s)
}

func (s GetErResponse) GoString() string {
	return s.String()
}

func (s *GetErResponse) SetHeaders(v map[string]*string) *GetErResponse {
	s.Headers = v
	return s
}

func (s *GetErResponse) SetStatusCode(v int32) *GetErResponse {
	s.StatusCode = &v
	return s
}

func (s *GetErResponse) SetBody(v *GetErResponseBody) *GetErResponse {
	s.Body = v
	return s
}

type GetErAttachmentRequest struct {
	ErAttachmentId *string `json:"ErAttachmentId,omitempty" xml:"ErAttachmentId,omitempty"`
	ErId           *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetErAttachmentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetErAttachmentRequest) GoString() string {
	return s.String()
}

func (s *GetErAttachmentRequest) SetErAttachmentId(v string) *GetErAttachmentRequest {
	s.ErAttachmentId = &v
	return s
}

func (s *GetErAttachmentRequest) SetErId(v string) *GetErAttachmentRequest {
	s.ErId = &v
	return s
}

func (s *GetErAttachmentRequest) SetRegionId(v string) *GetErAttachmentRequest {
	s.RegionId = &v
	return s
}

type GetErAttachmentResponseBody struct {
	Code      *int32                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   *GetErAttachmentResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetErAttachmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetErAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *GetErAttachmentResponseBody) SetCode(v int32) *GetErAttachmentResponseBody {
	s.Code = &v
	return s
}

func (s *GetErAttachmentResponseBody) SetContent(v *GetErAttachmentResponseBodyContent) *GetErAttachmentResponseBody {
	s.Content = v
	return s
}

func (s *GetErAttachmentResponseBody) SetMessage(v string) *GetErAttachmentResponseBody {
	s.Message = &v
	return s
}

func (s *GetErAttachmentResponseBody) SetRequestId(v string) *GetErAttachmentResponseBody {
	s.RequestId = &v
	return s
}

type GetErAttachmentResponseBodyContent struct {
	Across              *bool   `json:"Across,omitempty" xml:"Across,omitempty"`
	AutoReceiveAllRoute *bool   `json:"AutoReceiveAllRoute,omitempty" xml:"AutoReceiveAllRoute,omitempty"`
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ErAttachmentId      *string `json:"ErAttachmentId,omitempty" xml:"ErAttachmentId,omitempty"`
	ErAttachmentName    *string `json:"ErAttachmentName,omitempty" xml:"ErAttachmentName,omitempty"`
	ErId                *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	GmtModified         *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName        *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceType        *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Message             *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceTenantId    *string `json:"ResourceTenantId,omitempty" xml:"ResourceTenantId,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TenantId            *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetErAttachmentResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s GetErAttachmentResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetErAttachmentResponseBodyContent) SetAcross(v bool) *GetErAttachmentResponseBodyContent {
	s.Across = &v
	return s
}

func (s *GetErAttachmentResponseBodyContent) SetAutoReceiveAllRoute(v bool) *GetErAttachmentResponseBodyContent {
	s.AutoReceiveAllRoute = &v
	return s
}

func (s *GetErAttachmentResponseBodyContent) SetCreateTime(v string) *GetErAttachmentResponseBodyContent {
	s.CreateTime = &v
	return s
}

func (s *GetErAttachmentResponseBodyContent) SetErAttachmentId(v string) *GetErAttachmentResponseBodyContent {
	s.ErAttachmentId = &v
	return s
}

func (s *GetErAttachmentResponseBodyContent) SetErAttachmentName(v string) *GetErAttachmentResponseBodyContent {
	s.ErAttachmentName = &v
	return s
}

func (s *GetErAttachmentResponseBodyContent) SetErId(v string) *GetErAttachmentResponseBodyContent {
	s.ErId = &v
	return s
}

func (s *GetErAttachmentResponseBodyContent) SetGmtModified(v string) *GetErAttachmentResponseBodyContent {
	s.GmtModified = &v
	return s
}

func (s *GetErAttachmentResponseBodyContent) SetInstanceId(v string) *GetErAttachmentResponseBodyContent {
	s.InstanceId = &v
	return s
}

func (s *GetErAttachmentResponseBodyContent) SetInstanceName(v string) *GetErAttachmentResponseBodyContent {
	s.InstanceName = &v
	return s
}

func (s *GetErAttachmentResponseBodyContent) SetInstanceType(v string) *GetErAttachmentResponseBodyContent {
	s.InstanceType = &v
	return s
}

func (s *GetErAttachmentResponseBodyContent) SetMessage(v string) *GetErAttachmentResponseBodyContent {
	s.Message = &v
	return s
}

func (s *GetErAttachmentResponseBodyContent) SetRegionId(v string) *GetErAttachmentResponseBodyContent {
	s.RegionId = &v
	return s
}

func (s *GetErAttachmentResponseBodyContent) SetResourceTenantId(v string) *GetErAttachmentResponseBodyContent {
	s.ResourceTenantId = &v
	return s
}

func (s *GetErAttachmentResponseBodyContent) SetStatus(v string) *GetErAttachmentResponseBodyContent {
	s.Status = &v
	return s
}

func (s *GetErAttachmentResponseBodyContent) SetTenantId(v string) *GetErAttachmentResponseBodyContent {
	s.TenantId = &v
	return s
}

type GetErAttachmentResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetErAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetErAttachmentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetErAttachmentResponse) GoString() string {
	return s.String()
}

func (s *GetErAttachmentResponse) SetHeaders(v map[string]*string) *GetErAttachmentResponse {
	s.Headers = v
	return s
}

func (s *GetErAttachmentResponse) SetStatusCode(v int32) *GetErAttachmentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetErAttachmentResponse) SetBody(v *GetErAttachmentResponseBody) *GetErAttachmentResponse {
	s.Body = v
	return s
}

type GetErRouteEntryRequest struct {
	ErId           *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	ErRouteEntryId *string `json:"ErRouteEntryId,omitempty" xml:"ErRouteEntryId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetErRouteEntryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetErRouteEntryRequest) GoString() string {
	return s.String()
}

func (s *GetErRouteEntryRequest) SetErId(v string) *GetErRouteEntryRequest {
	s.ErId = &v
	return s
}

func (s *GetErRouteEntryRequest) SetErRouteEntryId(v string) *GetErRouteEntryRequest {
	s.ErRouteEntryId = &v
	return s
}

func (s *GetErRouteEntryRequest) SetRegionId(v string) *GetErRouteEntryRequest {
	s.RegionId = &v
	return s
}

type GetErRouteEntryResponseBody struct {
	Code      *int32                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   *GetErRouteEntryResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetErRouteEntryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetErRouteEntryResponseBody) GoString() string {
	return s.String()
}

func (s *GetErRouteEntryResponseBody) SetCode(v int32) *GetErRouteEntryResponseBody {
	s.Code = &v
	return s
}

func (s *GetErRouteEntryResponseBody) SetContent(v *GetErRouteEntryResponseBodyContent) *GetErRouteEntryResponseBody {
	s.Content = v
	return s
}

func (s *GetErRouteEntryResponseBody) SetMessage(v string) *GetErRouteEntryResponseBody {
	s.Message = &v
	return s
}

func (s *GetErRouteEntryResponseBody) SetRequestId(v string) *GetErRouteEntryResponseBody {
	s.RequestId = &v
	return s
}

type GetErRouteEntryResponseBodyContent struct {
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	ErId                 *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	ErRouteEntryId       *string `json:"ErRouteEntryId,omitempty" xml:"ErRouteEntryId,omitempty"`
	GmtModified          *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	NextHopId            *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	NextHopType          *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RouteType            *string `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TenantId             *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetErRouteEntryResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s GetErRouteEntryResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetErRouteEntryResponseBodyContent) SetDestinationCidrBlock(v string) *GetErRouteEntryResponseBodyContent {
	s.DestinationCidrBlock = &v
	return s
}

func (s *GetErRouteEntryResponseBodyContent) SetErId(v string) *GetErRouteEntryResponseBodyContent {
	s.ErId = &v
	return s
}

func (s *GetErRouteEntryResponseBodyContent) SetErRouteEntryId(v string) *GetErRouteEntryResponseBodyContent {
	s.ErRouteEntryId = &v
	return s
}

func (s *GetErRouteEntryResponseBodyContent) SetGmtModified(v string) *GetErRouteEntryResponseBodyContent {
	s.GmtModified = &v
	return s
}

func (s *GetErRouteEntryResponseBodyContent) SetNextHopId(v string) *GetErRouteEntryResponseBodyContent {
	s.NextHopId = &v
	return s
}

func (s *GetErRouteEntryResponseBodyContent) SetNextHopType(v string) *GetErRouteEntryResponseBodyContent {
	s.NextHopType = &v
	return s
}

func (s *GetErRouteEntryResponseBodyContent) SetRegionId(v string) *GetErRouteEntryResponseBodyContent {
	s.RegionId = &v
	return s
}

func (s *GetErRouteEntryResponseBodyContent) SetRouteType(v string) *GetErRouteEntryResponseBodyContent {
	s.RouteType = &v
	return s
}

func (s *GetErRouteEntryResponseBodyContent) SetStatus(v string) *GetErRouteEntryResponseBodyContent {
	s.Status = &v
	return s
}

func (s *GetErRouteEntryResponseBodyContent) SetTenantId(v string) *GetErRouteEntryResponseBodyContent {
	s.TenantId = &v
	return s
}

type GetErRouteEntryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetErRouteEntryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetErRouteEntryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetErRouteEntryResponse) GoString() string {
	return s.String()
}

func (s *GetErRouteEntryResponse) SetHeaders(v map[string]*string) *GetErRouteEntryResponse {
	s.Headers = v
	return s
}

func (s *GetErRouteEntryResponse) SetStatusCode(v int32) *GetErRouteEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetErRouteEntryResponse) SetBody(v *GetErRouteEntryResponseBody) *GetErRouteEntryResponse {
	s.Body = v
	return s
}

type GetErRouteMapRequest struct {
	ErId         *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	ErRouteMapId *string `json:"ErRouteMapId,omitempty" xml:"ErRouteMapId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetErRouteMapRequest) String() string {
	return tea.Prettify(s)
}

func (s GetErRouteMapRequest) GoString() string {
	return s.String()
}

func (s *GetErRouteMapRequest) SetErId(v string) *GetErRouteMapRequest {
	s.ErId = &v
	return s
}

func (s *GetErRouteMapRequest) SetErRouteMapId(v string) *GetErRouteMapRequest {
	s.ErRouteMapId = &v
	return s
}

func (s *GetErRouteMapRequest) SetRegionId(v string) *GetErRouteMapRequest {
	s.RegionId = &v
	return s
}

type GetErRouteMapResponseBody struct {
	Code      *int32                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   *GetErRouteMapResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetErRouteMapResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetErRouteMapResponseBody) GoString() string {
	return s.String()
}

func (s *GetErRouteMapResponseBody) SetCode(v int32) *GetErRouteMapResponseBody {
	s.Code = &v
	return s
}

func (s *GetErRouteMapResponseBody) SetContent(v *GetErRouteMapResponseBodyContent) *GetErRouteMapResponseBody {
	s.Content = v
	return s
}

func (s *GetErRouteMapResponseBody) SetMessage(v string) *GetErRouteMapResponseBody {
	s.Message = &v
	return s
}

func (s *GetErRouteMapResponseBody) SetRequestId(v string) *GetErRouteMapResponseBody {
	s.RequestId = &v
	return s
}

type GetErRouteMapResponseBodyContent struct {
	Action                    *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Description               *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DestinationCidrBlock      *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	ErId                      *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	ErRouteMapId              *string `json:"ErRouteMapId,omitempty" xml:"ErRouteMapId,omitempty"`
	ErRouteMapName            *string `json:"ErRouteMapName,omitempty" xml:"ErRouteMapName,omitempty"`
	GmtCreate                 *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified               *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Message                   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	ReceptionInstanceId       *string `json:"ReceptionInstanceId,omitempty" xml:"ReceptionInstanceId,omitempty"`
	ReceptionInstanceName     *string `json:"ReceptionInstanceName,omitempty" xml:"ReceptionInstanceName,omitempty"`
	ReceptionInstanceOwner    *string `json:"ReceptionInstanceOwner,omitempty" xml:"ReceptionInstanceOwner,omitempty"`
	ReceptionInstanceType     *string `json:"ReceptionInstanceType,omitempty" xml:"ReceptionInstanceType,omitempty"`
	RegionId                  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RouteMapNum               *int32  `json:"RouteMapNum,omitempty" xml:"RouteMapNum,omitempty"`
	Status                    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TenantId                  *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	TransmissionInstanceId    *string `json:"TransmissionInstanceId,omitempty" xml:"TransmissionInstanceId,omitempty"`
	TransmissionInstanceName  *string `json:"TransmissionInstanceName,omitempty" xml:"TransmissionInstanceName,omitempty"`
	TransmissionInstanceOwner *string `json:"TransmissionInstanceOwner,omitempty" xml:"TransmissionInstanceOwner,omitempty"`
	TransmissionInstanceType  *string `json:"TransmissionInstanceType,omitempty" xml:"TransmissionInstanceType,omitempty"`
}

func (s GetErRouteMapResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s GetErRouteMapResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetErRouteMapResponseBodyContent) SetAction(v string) *GetErRouteMapResponseBodyContent {
	s.Action = &v
	return s
}

func (s *GetErRouteMapResponseBodyContent) SetDescription(v string) *GetErRouteMapResponseBodyContent {
	s.Description = &v
	return s
}

func (s *GetErRouteMapResponseBodyContent) SetDestinationCidrBlock(v string) *GetErRouteMapResponseBodyContent {
	s.DestinationCidrBlock = &v
	return s
}

func (s *GetErRouteMapResponseBodyContent) SetErId(v string) *GetErRouteMapResponseBodyContent {
	s.ErId = &v
	return s
}

func (s *GetErRouteMapResponseBodyContent) SetErRouteMapId(v string) *GetErRouteMapResponseBodyContent {
	s.ErRouteMapId = &v
	return s
}

func (s *GetErRouteMapResponseBodyContent) SetErRouteMapName(v string) *GetErRouteMapResponseBodyContent {
	s.ErRouteMapName = &v
	return s
}

func (s *GetErRouteMapResponseBodyContent) SetGmtCreate(v string) *GetErRouteMapResponseBodyContent {
	s.GmtCreate = &v
	return s
}

func (s *GetErRouteMapResponseBodyContent) SetGmtModified(v string) *GetErRouteMapResponseBodyContent {
	s.GmtModified = &v
	return s
}

func (s *GetErRouteMapResponseBodyContent) SetMessage(v string) *GetErRouteMapResponseBodyContent {
	s.Message = &v
	return s
}

func (s *GetErRouteMapResponseBodyContent) SetReceptionInstanceId(v string) *GetErRouteMapResponseBodyContent {
	s.ReceptionInstanceId = &v
	return s
}

func (s *GetErRouteMapResponseBodyContent) SetReceptionInstanceName(v string) *GetErRouteMapResponseBodyContent {
	s.ReceptionInstanceName = &v
	return s
}

func (s *GetErRouteMapResponseBodyContent) SetReceptionInstanceOwner(v string) *GetErRouteMapResponseBodyContent {
	s.ReceptionInstanceOwner = &v
	return s
}

func (s *GetErRouteMapResponseBodyContent) SetReceptionInstanceType(v string) *GetErRouteMapResponseBodyContent {
	s.ReceptionInstanceType = &v
	return s
}

func (s *GetErRouteMapResponseBodyContent) SetRegionId(v string) *GetErRouteMapResponseBodyContent {
	s.RegionId = &v
	return s
}

func (s *GetErRouteMapResponseBodyContent) SetRouteMapNum(v int32) *GetErRouteMapResponseBodyContent {
	s.RouteMapNum = &v
	return s
}

func (s *GetErRouteMapResponseBodyContent) SetStatus(v string) *GetErRouteMapResponseBodyContent {
	s.Status = &v
	return s
}

func (s *GetErRouteMapResponseBodyContent) SetTenantId(v string) *GetErRouteMapResponseBodyContent {
	s.TenantId = &v
	return s
}

func (s *GetErRouteMapResponseBodyContent) SetTransmissionInstanceId(v string) *GetErRouteMapResponseBodyContent {
	s.TransmissionInstanceId = &v
	return s
}

func (s *GetErRouteMapResponseBodyContent) SetTransmissionInstanceName(v string) *GetErRouteMapResponseBodyContent {
	s.TransmissionInstanceName = &v
	return s
}

func (s *GetErRouteMapResponseBodyContent) SetTransmissionInstanceOwner(v string) *GetErRouteMapResponseBodyContent {
	s.TransmissionInstanceOwner = &v
	return s
}

func (s *GetErRouteMapResponseBodyContent) SetTransmissionInstanceType(v string) *GetErRouteMapResponseBodyContent {
	s.TransmissionInstanceType = &v
	return s
}

type GetErRouteMapResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetErRouteMapResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetErRouteMapResponse) String() string {
	return tea.Prettify(s)
}

func (s GetErRouteMapResponse) GoString() string {
	return s.String()
}

func (s *GetErRouteMapResponse) SetHeaders(v map[string]*string) *GetErRouteMapResponse {
	s.Headers = v
	return s
}

func (s *GetErRouteMapResponse) SetStatusCode(v int32) *GetErRouteMapResponse {
	s.StatusCode = &v
	return s
}

func (s *GetErRouteMapResponse) SetBody(v *GetErRouteMapResponseBody) *GetErRouteMapResponse {
	s.Body = v
	return s
}

type GetLniPrivateIpAddressRequest struct {
	IpName             *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetLniPrivateIpAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLniPrivateIpAddressRequest) GoString() string {
	return s.String()
}

func (s *GetLniPrivateIpAddressRequest) SetIpName(v string) *GetLniPrivateIpAddressRequest {
	s.IpName = &v
	return s
}

func (s *GetLniPrivateIpAddressRequest) SetNetworkInterfaceId(v string) *GetLniPrivateIpAddressRequest {
	s.NetworkInterfaceId = &v
	return s
}

func (s *GetLniPrivateIpAddressRequest) SetRegionId(v string) *GetLniPrivateIpAddressRequest {
	s.RegionId = &v
	return s
}

type GetLniPrivateIpAddressResponseBody struct {
	Code      *int32                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   *GetLniPrivateIpAddressResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLniPrivateIpAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLniPrivateIpAddressResponseBody) GoString() string {
	return s.String()
}

func (s *GetLniPrivateIpAddressResponseBody) SetCode(v int32) *GetLniPrivateIpAddressResponseBody {
	s.Code = &v
	return s
}

func (s *GetLniPrivateIpAddressResponseBody) SetContent(v *GetLniPrivateIpAddressResponseBodyContent) *GetLniPrivateIpAddressResponseBody {
	s.Content = v
	return s
}

func (s *GetLniPrivateIpAddressResponseBody) SetMessage(v string) *GetLniPrivateIpAddressResponseBody {
	s.Message = &v
	return s
}

func (s *GetLniPrivateIpAddressResponseBody) SetRequestId(v string) *GetLniPrivateIpAddressResponseBody {
	s.RequestId = &v
	return s
}

type GetLniPrivateIpAddressResponseBodyContent struct {
	GmtCreate          *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	IpAddressMac       *string `json:"IpAddressMac,omitempty" xml:"IpAddressMac,omitempty"`
	IpName             *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
	Message            *string `json:"Message,omitempty" xml:"Message,omitempty"`
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	PrivateIpAddress   *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetLniPrivateIpAddressResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s GetLniPrivateIpAddressResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetLniPrivateIpAddressResponseBodyContent) SetGmtCreate(v string) *GetLniPrivateIpAddressResponseBodyContent {
	s.GmtCreate = &v
	return s
}

func (s *GetLniPrivateIpAddressResponseBodyContent) SetIpAddressMac(v string) *GetLniPrivateIpAddressResponseBodyContent {
	s.IpAddressMac = &v
	return s
}

func (s *GetLniPrivateIpAddressResponseBodyContent) SetIpName(v string) *GetLniPrivateIpAddressResponseBodyContent {
	s.IpName = &v
	return s
}

func (s *GetLniPrivateIpAddressResponseBodyContent) SetMessage(v string) *GetLniPrivateIpAddressResponseBodyContent {
	s.Message = &v
	return s
}

func (s *GetLniPrivateIpAddressResponseBodyContent) SetNetworkInterfaceId(v string) *GetLniPrivateIpAddressResponseBodyContent {
	s.NetworkInterfaceId = &v
	return s
}

func (s *GetLniPrivateIpAddressResponseBodyContent) SetPrivateIpAddress(v string) *GetLniPrivateIpAddressResponseBodyContent {
	s.PrivateIpAddress = &v
	return s
}

func (s *GetLniPrivateIpAddressResponseBodyContent) SetRegionId(v string) *GetLniPrivateIpAddressResponseBodyContent {
	s.RegionId = &v
	return s
}

func (s *GetLniPrivateIpAddressResponseBodyContent) SetStatus(v string) *GetLniPrivateIpAddressResponseBodyContent {
	s.Status = &v
	return s
}

type GetLniPrivateIpAddressResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetLniPrivateIpAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLniPrivateIpAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLniPrivateIpAddressResponse) GoString() string {
	return s.String()
}

func (s *GetLniPrivateIpAddressResponse) SetHeaders(v map[string]*string) *GetLniPrivateIpAddressResponse {
	s.Headers = v
	return s
}

func (s *GetLniPrivateIpAddressResponse) SetStatusCode(v int32) *GetLniPrivateIpAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLniPrivateIpAddressResponse) SetBody(v *GetLniPrivateIpAddressResponseBody) *GetLniPrivateIpAddressResponse {
	s.Body = v
	return s
}

type GetNetworkInterfaceRequest struct {
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SubnetId           *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
}

func (s GetNetworkInterfaceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNetworkInterfaceRequest) GoString() string {
	return s.String()
}

func (s *GetNetworkInterfaceRequest) SetNetworkInterfaceId(v string) *GetNetworkInterfaceRequest {
	s.NetworkInterfaceId = &v
	return s
}

func (s *GetNetworkInterfaceRequest) SetRegionId(v string) *GetNetworkInterfaceRequest {
	s.RegionId = &v
	return s
}

func (s *GetNetworkInterfaceRequest) SetSubnetId(v string) *GetNetworkInterfaceRequest {
	s.SubnetId = &v
	return s
}

type GetNetworkInterfaceResponseBody struct {
	Code      *int32                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   *GetNetworkInterfaceResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNetworkInterfaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNetworkInterfaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetNetworkInterfaceResponseBody) SetCode(v int32) *GetNetworkInterfaceResponseBody {
	s.Code = &v
	return s
}

func (s *GetNetworkInterfaceResponseBody) SetContent(v *GetNetworkInterfaceResponseBodyContent) *GetNetworkInterfaceResponseBody {
	s.Content = v
	return s
}

func (s *GetNetworkInterfaceResponseBody) SetMessage(v string) *GetNetworkInterfaceResponseBody {
	s.Message = &v
	return s
}

func (s *GetNetworkInterfaceResponseBody) SetRequestId(v string) *GetNetworkInterfaceResponseBody {
	s.RequestId = &v
	return s
}

type GetNetworkInterfaceResponseBodyContent struct {
	CreateTime               *string                                                           `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Ethernet                 []*string                                                         `json:"Ethernet,omitempty" xml:"Ethernet,omitempty" type:"Repeated"`
	Gateway                  *string                                                           `json:"Gateway,omitempty" xml:"Gateway,omitempty"`
	Ip                       *string                                                           `json:"Ip,omitempty" xml:"Ip,omitempty"`
	NcType                   *string                                                           `json:"NcType,omitempty" xml:"NcType,omitempty"`
	NetworkInterfaceId       *string                                                           `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	NetworkInterfaceName     *string                                                           `json:"NetworkInterfaceName,omitempty" xml:"NetworkInterfaceName,omitempty"`
	NodeId                   *string                                                           `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	PrivateIpAddressMacGroup []*GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup `json:"PrivateIpAddressMacGroup,omitempty" xml:"PrivateIpAddressMacGroup,omitempty" type:"Repeated"`
	Quota                    *int32                                                            `json:"Quota,omitempty" xml:"Quota,omitempty"`
	RegionId                 *string                                                           `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceMac               *string                                                           `json:"ServiceMac,omitempty" xml:"ServiceMac,omitempty"`
	Status                   *string                                                           `json:"Status,omitempty" xml:"Status,omitempty"`
	SubnetBaseInfo           *GetNetworkInterfaceResponseBodyContentSubnetBaseInfo             `json:"SubnetBaseInfo,omitempty" xml:"SubnetBaseInfo,omitempty" type:"Struct"`
	TenantId                 *string                                                           `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	VpdBaseInfo              *GetNetworkInterfaceResponseBodyContentVpdBaseInfo                `json:"VpdBaseInfo,omitempty" xml:"VpdBaseInfo,omitempty" type:"Struct"`
	ZoneId                   *string                                                           `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetNetworkInterfaceResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s GetNetworkInterfaceResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetNetworkInterfaceResponseBodyContent) SetCreateTime(v string) *GetNetworkInterfaceResponseBodyContent {
	s.CreateTime = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContent) SetEthernet(v []*string) *GetNetworkInterfaceResponseBodyContent {
	s.Ethernet = v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContent) SetGateway(v string) *GetNetworkInterfaceResponseBodyContent {
	s.Gateway = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContent) SetIp(v string) *GetNetworkInterfaceResponseBodyContent {
	s.Ip = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContent) SetNcType(v string) *GetNetworkInterfaceResponseBodyContent {
	s.NcType = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContent) SetNetworkInterfaceId(v string) *GetNetworkInterfaceResponseBodyContent {
	s.NetworkInterfaceId = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContent) SetNetworkInterfaceName(v string) *GetNetworkInterfaceResponseBodyContent {
	s.NetworkInterfaceName = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContent) SetNodeId(v string) *GetNetworkInterfaceResponseBodyContent {
	s.NodeId = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContent) SetPrivateIpAddressMacGroup(v []*GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup) *GetNetworkInterfaceResponseBodyContent {
	s.PrivateIpAddressMacGroup = v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContent) SetQuota(v int32) *GetNetworkInterfaceResponseBodyContent {
	s.Quota = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContent) SetRegionId(v string) *GetNetworkInterfaceResponseBodyContent {
	s.RegionId = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContent) SetServiceMac(v string) *GetNetworkInterfaceResponseBodyContent {
	s.ServiceMac = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContent) SetStatus(v string) *GetNetworkInterfaceResponseBodyContent {
	s.Status = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContent) SetSubnetBaseInfo(v *GetNetworkInterfaceResponseBodyContentSubnetBaseInfo) *GetNetworkInterfaceResponseBodyContent {
	s.SubnetBaseInfo = v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContent) SetTenantId(v string) *GetNetworkInterfaceResponseBodyContent {
	s.TenantId = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContent) SetVpdBaseInfo(v *GetNetworkInterfaceResponseBodyContentVpdBaseInfo) *GetNetworkInterfaceResponseBodyContent {
	s.VpdBaseInfo = v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContent) SetZoneId(v string) *GetNetworkInterfaceResponseBodyContent {
	s.ZoneId = &v
	return s
}

type GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup struct {
	IpAddressMac     *string `json:"IpAddressMac,omitempty" xml:"IpAddressMac,omitempty"`
	IpName           *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
	Message          *string `json:"Message,omitempty" xml:"Message,omitempty"`
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup) String() string {
	return tea.Prettify(s)
}

func (s GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup) GoString() string {
	return s.String()
}

func (s *GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup) SetIpAddressMac(v string) *GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup {
	s.IpAddressMac = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup) SetIpName(v string) *GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup {
	s.IpName = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup) SetMessage(v string) *GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup {
	s.Message = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup) SetPrivateIpAddress(v string) *GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup {
	s.PrivateIpAddress = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup) SetStatus(v string) *GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup {
	s.Status = &v
	return s
}

type GetNetworkInterfaceResponseBodyContentSubnetBaseInfo struct {
	Cidr       *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	SubnetId   *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	SubnetName *string `json:"SubnetName,omitempty" xml:"SubnetName,omitempty"`
}

func (s GetNetworkInterfaceResponseBodyContentSubnetBaseInfo) String() string {
	return tea.Prettify(s)
}

func (s GetNetworkInterfaceResponseBodyContentSubnetBaseInfo) GoString() string {
	return s.String()
}

func (s *GetNetworkInterfaceResponseBodyContentSubnetBaseInfo) SetCidr(v string) *GetNetworkInterfaceResponseBodyContentSubnetBaseInfo {
	s.Cidr = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContentSubnetBaseInfo) SetCreateTime(v string) *GetNetworkInterfaceResponseBodyContentSubnetBaseInfo {
	s.CreateTime = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContentSubnetBaseInfo) SetSubnetId(v string) *GetNetworkInterfaceResponseBodyContentSubnetBaseInfo {
	s.SubnetId = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContentSubnetBaseInfo) SetSubnetName(v string) *GetNetworkInterfaceResponseBodyContentSubnetBaseInfo {
	s.SubnetName = &v
	return s
}

type GetNetworkInterfaceResponseBodyContentVpdBaseInfo struct {
	Cidr       *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	VpdId      *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	VpdName    *string `json:"VpdName,omitempty" xml:"VpdName,omitempty"`
}

func (s GetNetworkInterfaceResponseBodyContentVpdBaseInfo) String() string {
	return tea.Prettify(s)
}

func (s GetNetworkInterfaceResponseBodyContentVpdBaseInfo) GoString() string {
	return s.String()
}

func (s *GetNetworkInterfaceResponseBodyContentVpdBaseInfo) SetCidr(v string) *GetNetworkInterfaceResponseBodyContentVpdBaseInfo {
	s.Cidr = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContentVpdBaseInfo) SetCreateTime(v string) *GetNetworkInterfaceResponseBodyContentVpdBaseInfo {
	s.CreateTime = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContentVpdBaseInfo) SetVpdId(v string) *GetNetworkInterfaceResponseBodyContentVpdBaseInfo {
	s.VpdId = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContentVpdBaseInfo) SetVpdName(v string) *GetNetworkInterfaceResponseBodyContentVpdBaseInfo {
	s.VpdName = &v
	return s
}

type GetNetworkInterfaceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetNetworkInterfaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetNetworkInterfaceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNetworkInterfaceResponse) GoString() string {
	return s.String()
}

func (s *GetNetworkInterfaceResponse) SetHeaders(v map[string]*string) *GetNetworkInterfaceResponse {
	s.Headers = v
	return s
}

func (s *GetNetworkInterfaceResponse) SetStatusCode(v int32) *GetNetworkInterfaceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNetworkInterfaceResponse) SetBody(v *GetNetworkInterfaceResponseBody) *GetNetworkInterfaceResponse {
	s.Body = v
	return s
}

type GetSubnetRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	VpdId    *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
}

func (s GetSubnetRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSubnetRequest) GoString() string {
	return s.String()
}

func (s *GetSubnetRequest) SetRegionId(v string) *GetSubnetRequest {
	s.RegionId = &v
	return s
}

func (s *GetSubnetRequest) SetSubnetId(v string) *GetSubnetRequest {
	s.SubnetId = &v
	return s
}

func (s *GetSubnetRequest) SetVpdId(v string) *GetSubnetRequest {
	s.VpdId = &v
	return s
}

type GetSubnetResponseBody struct {
	Code      *int32                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   *GetSubnetResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSubnetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSubnetResponseBody) GoString() string {
	return s.String()
}

func (s *GetSubnetResponseBody) SetCode(v int32) *GetSubnetResponseBody {
	s.Code = &v
	return s
}

func (s *GetSubnetResponseBody) SetContent(v *GetSubnetResponseBodyContent) *GetSubnetResponseBody {
	s.Content = v
	return s
}

func (s *GetSubnetResponseBody) SetMessage(v string) *GetSubnetResponseBody {
	s.Message = &v
	return s
}

func (s *GetSubnetResponseBody) SetRequestId(v string) *GetSubnetResponseBody {
	s.RequestId = &v
	return s
}

type GetSubnetResponseBodyContent struct {
	AvailableIps    *int32                                   `json:"AvailableIps,omitempty" xml:"AvailableIps,omitempty"`
	Cidr            *string                                  `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	CreateTime      *string                                  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	GmtModified     *string                                  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	LbCount         *int64                                   `json:"LbCount,omitempty" xml:"LbCount,omitempty"`
	Message         *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	NcCount         *int64                                   `json:"NcCount,omitempty" xml:"NcCount,omitempty"`
	RegionId        *string                                  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string                                  `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Status          *string                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	SubnetId        *string                                  `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	SubnetName      *string                                  `json:"SubnetName,omitempty" xml:"SubnetName,omitempty"`
	Tags            []*GetSubnetResponseBodyContentTags      `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TenantId        *string                                  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Type            *string                                  `json:"Type,omitempty" xml:"Type,omitempty"`
	VpdBaseInfo     *GetSubnetResponseBodyContentVpdBaseInfo `json:"VpdBaseInfo,omitempty" xml:"VpdBaseInfo,omitempty" type:"Struct"`
	VpdId           *string                                  `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	ZoneId          *string                                  `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetSubnetResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s GetSubnetResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetSubnetResponseBodyContent) SetAvailableIps(v int32) *GetSubnetResponseBodyContent {
	s.AvailableIps = &v
	return s
}

func (s *GetSubnetResponseBodyContent) SetCidr(v string) *GetSubnetResponseBodyContent {
	s.Cidr = &v
	return s
}

func (s *GetSubnetResponseBodyContent) SetCreateTime(v string) *GetSubnetResponseBodyContent {
	s.CreateTime = &v
	return s
}

func (s *GetSubnetResponseBodyContent) SetGmtModified(v string) *GetSubnetResponseBodyContent {
	s.GmtModified = &v
	return s
}

func (s *GetSubnetResponseBodyContent) SetLbCount(v int64) *GetSubnetResponseBodyContent {
	s.LbCount = &v
	return s
}

func (s *GetSubnetResponseBodyContent) SetMessage(v string) *GetSubnetResponseBodyContent {
	s.Message = &v
	return s
}

func (s *GetSubnetResponseBodyContent) SetNcCount(v int64) *GetSubnetResponseBodyContent {
	s.NcCount = &v
	return s
}

func (s *GetSubnetResponseBodyContent) SetRegionId(v string) *GetSubnetResponseBodyContent {
	s.RegionId = &v
	return s
}

func (s *GetSubnetResponseBodyContent) SetResourceGroupId(v string) *GetSubnetResponseBodyContent {
	s.ResourceGroupId = &v
	return s
}

func (s *GetSubnetResponseBodyContent) SetStatus(v string) *GetSubnetResponseBodyContent {
	s.Status = &v
	return s
}

func (s *GetSubnetResponseBodyContent) SetSubnetId(v string) *GetSubnetResponseBodyContent {
	s.SubnetId = &v
	return s
}

func (s *GetSubnetResponseBodyContent) SetSubnetName(v string) *GetSubnetResponseBodyContent {
	s.SubnetName = &v
	return s
}

func (s *GetSubnetResponseBodyContent) SetTags(v []*GetSubnetResponseBodyContentTags) *GetSubnetResponseBodyContent {
	s.Tags = v
	return s
}

func (s *GetSubnetResponseBodyContent) SetTenantId(v string) *GetSubnetResponseBodyContent {
	s.TenantId = &v
	return s
}

func (s *GetSubnetResponseBodyContent) SetType(v string) *GetSubnetResponseBodyContent {
	s.Type = &v
	return s
}

func (s *GetSubnetResponseBodyContent) SetVpdBaseInfo(v *GetSubnetResponseBodyContentVpdBaseInfo) *GetSubnetResponseBodyContent {
	s.VpdBaseInfo = v
	return s
}

func (s *GetSubnetResponseBodyContent) SetVpdId(v string) *GetSubnetResponseBodyContent {
	s.VpdId = &v
	return s
}

func (s *GetSubnetResponseBodyContent) SetZoneId(v string) *GetSubnetResponseBodyContent {
	s.ZoneId = &v
	return s
}

type GetSubnetResponseBodyContentTags struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetSubnetResponseBodyContentTags) String() string {
	return tea.Prettify(s)
}

func (s GetSubnetResponseBodyContentTags) GoString() string {
	return s.String()
}

func (s *GetSubnetResponseBodyContentTags) SetTagKey(v string) *GetSubnetResponseBodyContentTags {
	s.TagKey = &v
	return s
}

func (s *GetSubnetResponseBodyContentTags) SetTagValue(v string) *GetSubnetResponseBodyContentTags {
	s.TagValue = &v
	return s
}

type GetSubnetResponseBodyContentVpdBaseInfo struct {
	Cidr       *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	VpdId      *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	VpdName    *string `json:"VpdName,omitempty" xml:"VpdName,omitempty"`
}

func (s GetSubnetResponseBodyContentVpdBaseInfo) String() string {
	return tea.Prettify(s)
}

func (s GetSubnetResponseBodyContentVpdBaseInfo) GoString() string {
	return s.String()
}

func (s *GetSubnetResponseBodyContentVpdBaseInfo) SetCidr(v string) *GetSubnetResponseBodyContentVpdBaseInfo {
	s.Cidr = &v
	return s
}

func (s *GetSubnetResponseBodyContentVpdBaseInfo) SetCreateTime(v string) *GetSubnetResponseBodyContentVpdBaseInfo {
	s.CreateTime = &v
	return s
}

func (s *GetSubnetResponseBodyContentVpdBaseInfo) SetVpdId(v string) *GetSubnetResponseBodyContentVpdBaseInfo {
	s.VpdId = &v
	return s
}

func (s *GetSubnetResponseBodyContentVpdBaseInfo) SetVpdName(v string) *GetSubnetResponseBodyContentVpdBaseInfo {
	s.VpdName = &v
	return s
}

type GetSubnetResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSubnetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSubnetResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSubnetResponse) GoString() string {
	return s.String()
}

func (s *GetSubnetResponse) SetHeaders(v map[string]*string) *GetSubnetResponse {
	s.Headers = v
	return s
}

func (s *GetSubnetResponse) SetStatusCode(v int32) *GetSubnetResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSubnetResponse) SetBody(v *GetSubnetResponseBody) *GetSubnetResponse {
	s.Body = v
	return s
}

type GetVccRequest struct {
	EnablePage *bool   `json:"EnablePage,omitempty" xml:"EnablePage,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VccId      *string `json:"VccId,omitempty" xml:"VccId,omitempty"`
}

func (s GetVccRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVccRequest) GoString() string {
	return s.String()
}

func (s *GetVccRequest) SetEnablePage(v bool) *GetVccRequest {
	s.EnablePage = &v
	return s
}

func (s *GetVccRequest) SetPageNumber(v int32) *GetVccRequest {
	s.PageNumber = &v
	return s
}

func (s *GetVccRequest) SetPageSize(v int32) *GetVccRequest {
	s.PageSize = &v
	return s
}

func (s *GetVccRequest) SetRegionId(v string) *GetVccRequest {
	s.RegionId = &v
	return s
}

func (s *GetVccRequest) SetVccId(v string) *GetVccRequest {
	s.VccId = &v
	return s
}

type GetVccResponseBody struct {
	Code      *int32                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   *GetVccResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Message   *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVccResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVccResponseBody) GoString() string {
	return s.String()
}

func (s *GetVccResponseBody) SetCode(v int32) *GetVccResponseBody {
	s.Code = &v
	return s
}

func (s *GetVccResponseBody) SetContent(v *GetVccResponseBodyContent) *GetVccResponseBody {
	s.Content = v
	return s
}

func (s *GetVccResponseBody) SetMessage(v string) *GetVccResponseBody {
	s.Message = &v
	return s
}

func (s *GetVccResponseBody) SetRequestId(v string) *GetVccResponseBody {
	s.RequestId = &v
	return s
}

type GetVccResponseBodyContent struct {
	AccessPointId      *string                                      `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	AliyunRouterInfo   []*GetVccResponseBodyContentAliyunRouterInfo `json:"AliyunRouterInfo,omitempty" xml:"AliyunRouterInfo,omitempty" type:"Repeated"`
	AttachErStatus     *bool                                        `json:"AttachErStatus,omitempty" xml:"AttachErStatus,omitempty"`
	BandwidthStr       *string                                      `json:"BandwidthStr,omitempty" xml:"BandwidthStr,omitempty"`
	BgpCidr            *string                                      `json:"BgpCidr,omitempty" xml:"BgpCidr,omitempty"`
	CenId              *string                                      `json:"CenId,omitempty" xml:"CenId,omitempty"`
	CisRouterInfo      []*GetVccResponseBodyContentCisRouterInfo    `json:"CisRouterInfo,omitempty" xml:"CisRouterInfo,omitempty" type:"Repeated"`
	CommodityCode      *string                                      `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	ConnectionType     *string                                      `json:"ConnectionType,omitempty" xml:"ConnectionType,omitempty"`
	CreateTime         *string                                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CurrentNode        *string                                      `json:"CurrentNode,omitempty" xml:"CurrentNode,omitempty"`
	Duration           *string                                      `json:"Duration,omitempty" xml:"Duration,omitempty"`
	ErInfos            []*GetVccResponseBodyContentErInfos          `json:"ErInfos,omitempty" xml:"ErInfos,omitempty" type:"Repeated"`
	ExpirationDate     *string                                      `json:"ExpirationDate,omitempty" xml:"ExpirationDate,omitempty"`
	GmtModified        *string                                      `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	InternetChargeType *string                                      `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	LineOperator       *string                                      `json:"LineOperator,omitempty" xml:"LineOperator,omitempty"`
	Message            *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	PayType            *string                                      `json:"PayType,omitempty" xml:"PayType,omitempty"`
	PortType           *string                                      `json:"PortType,omitempty" xml:"PortType,omitempty"`
	PricingCycle       *string                                      `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	RegionId           *string                                      `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId    *string                                      `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Spec               *string                                      `json:"Spec,omitempty" xml:"Spec,omitempty"`
	Status             *string                                      `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags               []*GetVccResponseBodyContentTags             `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TenantId           *string                                      `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	VSwitchId          *string                                      `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VccId              *string                                      `json:"VccId,omitempty" xml:"VccId,omitempty"`
	VccName            *string                                      `json:"VccName,omitempty" xml:"VccName,omitempty"`
	VpcId              *string                                      `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpdBaseInfo        *GetVccResponseBodyContentVpdBaseInfo        `json:"VpdBaseInfo,omitempty" xml:"VpdBaseInfo,omitempty" type:"Struct"`
	VpdId              *string                                      `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	ZoneId             *string                                      `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetVccResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s GetVccResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetVccResponseBodyContent) SetAccessPointId(v string) *GetVccResponseBodyContent {
	s.AccessPointId = &v
	return s
}

func (s *GetVccResponseBodyContent) SetAliyunRouterInfo(v []*GetVccResponseBodyContentAliyunRouterInfo) *GetVccResponseBodyContent {
	s.AliyunRouterInfo = v
	return s
}

func (s *GetVccResponseBodyContent) SetAttachErStatus(v bool) *GetVccResponseBodyContent {
	s.AttachErStatus = &v
	return s
}

func (s *GetVccResponseBodyContent) SetBandwidthStr(v string) *GetVccResponseBodyContent {
	s.BandwidthStr = &v
	return s
}

func (s *GetVccResponseBodyContent) SetBgpCidr(v string) *GetVccResponseBodyContent {
	s.BgpCidr = &v
	return s
}

func (s *GetVccResponseBodyContent) SetCenId(v string) *GetVccResponseBodyContent {
	s.CenId = &v
	return s
}

func (s *GetVccResponseBodyContent) SetCisRouterInfo(v []*GetVccResponseBodyContentCisRouterInfo) *GetVccResponseBodyContent {
	s.CisRouterInfo = v
	return s
}

func (s *GetVccResponseBodyContent) SetCommodityCode(v string) *GetVccResponseBodyContent {
	s.CommodityCode = &v
	return s
}

func (s *GetVccResponseBodyContent) SetConnectionType(v string) *GetVccResponseBodyContent {
	s.ConnectionType = &v
	return s
}

func (s *GetVccResponseBodyContent) SetCreateTime(v string) *GetVccResponseBodyContent {
	s.CreateTime = &v
	return s
}

func (s *GetVccResponseBodyContent) SetCurrentNode(v string) *GetVccResponseBodyContent {
	s.CurrentNode = &v
	return s
}

func (s *GetVccResponseBodyContent) SetDuration(v string) *GetVccResponseBodyContent {
	s.Duration = &v
	return s
}

func (s *GetVccResponseBodyContent) SetErInfos(v []*GetVccResponseBodyContentErInfos) *GetVccResponseBodyContent {
	s.ErInfos = v
	return s
}

func (s *GetVccResponseBodyContent) SetExpirationDate(v string) *GetVccResponseBodyContent {
	s.ExpirationDate = &v
	return s
}

func (s *GetVccResponseBodyContent) SetGmtModified(v string) *GetVccResponseBodyContent {
	s.GmtModified = &v
	return s
}

func (s *GetVccResponseBodyContent) SetInternetChargeType(v string) *GetVccResponseBodyContent {
	s.InternetChargeType = &v
	return s
}

func (s *GetVccResponseBodyContent) SetLineOperator(v string) *GetVccResponseBodyContent {
	s.LineOperator = &v
	return s
}

func (s *GetVccResponseBodyContent) SetMessage(v string) *GetVccResponseBodyContent {
	s.Message = &v
	return s
}

func (s *GetVccResponseBodyContent) SetPayType(v string) *GetVccResponseBodyContent {
	s.PayType = &v
	return s
}

func (s *GetVccResponseBodyContent) SetPortType(v string) *GetVccResponseBodyContent {
	s.PortType = &v
	return s
}

func (s *GetVccResponseBodyContent) SetPricingCycle(v string) *GetVccResponseBodyContent {
	s.PricingCycle = &v
	return s
}

func (s *GetVccResponseBodyContent) SetRegionId(v string) *GetVccResponseBodyContent {
	s.RegionId = &v
	return s
}

func (s *GetVccResponseBodyContent) SetResourceGroupId(v string) *GetVccResponseBodyContent {
	s.ResourceGroupId = &v
	return s
}

func (s *GetVccResponseBodyContent) SetSpec(v string) *GetVccResponseBodyContent {
	s.Spec = &v
	return s
}

func (s *GetVccResponseBodyContent) SetStatus(v string) *GetVccResponseBodyContent {
	s.Status = &v
	return s
}

func (s *GetVccResponseBodyContent) SetTags(v []*GetVccResponseBodyContentTags) *GetVccResponseBodyContent {
	s.Tags = v
	return s
}

func (s *GetVccResponseBodyContent) SetTenantId(v string) *GetVccResponseBodyContent {
	s.TenantId = &v
	return s
}

func (s *GetVccResponseBodyContent) SetVSwitchId(v string) *GetVccResponseBodyContent {
	s.VSwitchId = &v
	return s
}

func (s *GetVccResponseBodyContent) SetVccId(v string) *GetVccResponseBodyContent {
	s.VccId = &v
	return s
}

func (s *GetVccResponseBodyContent) SetVccName(v string) *GetVccResponseBodyContent {
	s.VccName = &v
	return s
}

func (s *GetVccResponseBodyContent) SetVpcId(v string) *GetVccResponseBodyContent {
	s.VpcId = &v
	return s
}

func (s *GetVccResponseBodyContent) SetVpdBaseInfo(v *GetVccResponseBodyContentVpdBaseInfo) *GetVccResponseBodyContent {
	s.VpdBaseInfo = v
	return s
}

func (s *GetVccResponseBodyContent) SetVpdId(v string) *GetVccResponseBodyContent {
	s.VpdId = &v
	return s
}

func (s *GetVccResponseBodyContent) SetZoneId(v string) *GetVccResponseBodyContent {
	s.ZoneId = &v
	return s
}

type GetVccResponseBodyContentAliyunRouterInfo struct {
	LocalGatewayIp *string `json:"LocalGatewayIp,omitempty" xml:"LocalGatewayIp,omitempty"`
	Mask           *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	PcId           *string `json:"PcId,omitempty" xml:"PcId,omitempty"`
	PeerGatewayIp  *string `json:"PeerGatewayIp,omitempty" xml:"PeerGatewayIp,omitempty"`
	VbrId          *string `json:"VbrId,omitempty" xml:"VbrId,omitempty"`
	VlanId         *string `json:"VlanId,omitempty" xml:"VlanId,omitempty"`
}

func (s GetVccResponseBodyContentAliyunRouterInfo) String() string {
	return tea.Prettify(s)
}

func (s GetVccResponseBodyContentAliyunRouterInfo) GoString() string {
	return s.String()
}

func (s *GetVccResponseBodyContentAliyunRouterInfo) SetLocalGatewayIp(v string) *GetVccResponseBodyContentAliyunRouterInfo {
	s.LocalGatewayIp = &v
	return s
}

func (s *GetVccResponseBodyContentAliyunRouterInfo) SetMask(v string) *GetVccResponseBodyContentAliyunRouterInfo {
	s.Mask = &v
	return s
}

func (s *GetVccResponseBodyContentAliyunRouterInfo) SetPcId(v string) *GetVccResponseBodyContentAliyunRouterInfo {
	s.PcId = &v
	return s
}

func (s *GetVccResponseBodyContentAliyunRouterInfo) SetPeerGatewayIp(v string) *GetVccResponseBodyContentAliyunRouterInfo {
	s.PeerGatewayIp = &v
	return s
}

func (s *GetVccResponseBodyContentAliyunRouterInfo) SetVbrId(v string) *GetVccResponseBodyContentAliyunRouterInfo {
	s.VbrId = &v
	return s
}

func (s *GetVccResponseBodyContentAliyunRouterInfo) SetVlanId(v string) *GetVccResponseBodyContentAliyunRouterInfo {
	s.VlanId = &v
	return s
}

type GetVccResponseBodyContentCisRouterInfo struct {
	CcInfos []*GetVccResponseBodyContentCisRouterInfoCcInfos `json:"CcInfos,omitempty" xml:"CcInfos,omitempty" type:"Repeated"`
	CcrId   *string                                          `json:"CcrId,omitempty" xml:"CcrId,omitempty"`
}

func (s GetVccResponseBodyContentCisRouterInfo) String() string {
	return tea.Prettify(s)
}

func (s GetVccResponseBodyContentCisRouterInfo) GoString() string {
	return s.String()
}

func (s *GetVccResponseBodyContentCisRouterInfo) SetCcInfos(v []*GetVccResponseBodyContentCisRouterInfoCcInfos) *GetVccResponseBodyContentCisRouterInfo {
	s.CcInfos = v
	return s
}

func (s *GetVccResponseBodyContentCisRouterInfo) SetCcrId(v string) *GetVccResponseBodyContentCisRouterInfo {
	s.CcrId = &v
	return s
}

type GetVccResponseBodyContentCisRouterInfoCcInfos struct {
	CcId            *string `json:"CcId,omitempty" xml:"CcId,omitempty"`
	LocalGatewayIp  *string `json:"LocalGatewayIp,omitempty" xml:"LocalGatewayIp,omitempty"`
	RemoteGatewayIp *string `json:"RemoteGatewayIp,omitempty" xml:"RemoteGatewayIp,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SubnetMask      *string `json:"SubnetMask,omitempty" xml:"SubnetMask,omitempty"`
	// vlanid
	VlanId *string `json:"VlanId,omitempty" xml:"VlanId,omitempty"`
}

func (s GetVccResponseBodyContentCisRouterInfoCcInfos) String() string {
	return tea.Prettify(s)
}

func (s GetVccResponseBodyContentCisRouterInfoCcInfos) GoString() string {
	return s.String()
}

func (s *GetVccResponseBodyContentCisRouterInfoCcInfos) SetCcId(v string) *GetVccResponseBodyContentCisRouterInfoCcInfos {
	s.CcId = &v
	return s
}

func (s *GetVccResponseBodyContentCisRouterInfoCcInfos) SetLocalGatewayIp(v string) *GetVccResponseBodyContentCisRouterInfoCcInfos {
	s.LocalGatewayIp = &v
	return s
}

func (s *GetVccResponseBodyContentCisRouterInfoCcInfos) SetRemoteGatewayIp(v string) *GetVccResponseBodyContentCisRouterInfoCcInfos {
	s.RemoteGatewayIp = &v
	return s
}

func (s *GetVccResponseBodyContentCisRouterInfoCcInfos) SetStatus(v string) *GetVccResponseBodyContentCisRouterInfoCcInfos {
	s.Status = &v
	return s
}

func (s *GetVccResponseBodyContentCisRouterInfoCcInfos) SetSubnetMask(v string) *GetVccResponseBodyContentCisRouterInfoCcInfos {
	s.SubnetMask = &v
	return s
}

func (s *GetVccResponseBodyContentCisRouterInfoCcInfos) SetVlanId(v string) *GetVccResponseBodyContentCisRouterInfoCcInfos {
	s.VlanId = &v
	return s
}

type GetVccResponseBodyContentErInfos struct {
	Connections  *int64  `json:"Connections,omitempty" xml:"Connections,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ErId         *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	ErName       *string `json:"ErName,omitempty" xml:"ErName,omitempty"`
	GmtModified  *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	MasterZoneId *string `json:"MasterZoneId,omitempty" xml:"MasterZoneId,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RouteMaps    *int64  `json:"RouteMaps,omitempty" xml:"RouteMaps,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TenantId     *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetVccResponseBodyContentErInfos) String() string {
	return tea.Prettify(s)
}

func (s GetVccResponseBodyContentErInfos) GoString() string {
	return s.String()
}

func (s *GetVccResponseBodyContentErInfos) SetConnections(v int64) *GetVccResponseBodyContentErInfos {
	s.Connections = &v
	return s
}

func (s *GetVccResponseBodyContentErInfos) SetCreateTime(v string) *GetVccResponseBodyContentErInfos {
	s.CreateTime = &v
	return s
}

func (s *GetVccResponseBodyContentErInfos) SetDescription(v string) *GetVccResponseBodyContentErInfos {
	s.Description = &v
	return s
}

func (s *GetVccResponseBodyContentErInfos) SetErId(v string) *GetVccResponseBodyContentErInfos {
	s.ErId = &v
	return s
}

func (s *GetVccResponseBodyContentErInfos) SetErName(v string) *GetVccResponseBodyContentErInfos {
	s.ErName = &v
	return s
}

func (s *GetVccResponseBodyContentErInfos) SetGmtModified(v string) *GetVccResponseBodyContentErInfos {
	s.GmtModified = &v
	return s
}

func (s *GetVccResponseBodyContentErInfos) SetMasterZoneId(v string) *GetVccResponseBodyContentErInfos {
	s.MasterZoneId = &v
	return s
}

func (s *GetVccResponseBodyContentErInfos) SetMessage(v string) *GetVccResponseBodyContentErInfos {
	s.Message = &v
	return s
}

func (s *GetVccResponseBodyContentErInfos) SetRegionId(v string) *GetVccResponseBodyContentErInfos {
	s.RegionId = &v
	return s
}

func (s *GetVccResponseBodyContentErInfos) SetRouteMaps(v int64) *GetVccResponseBodyContentErInfos {
	s.RouteMaps = &v
	return s
}

func (s *GetVccResponseBodyContentErInfos) SetStatus(v string) *GetVccResponseBodyContentErInfos {
	s.Status = &v
	return s
}

func (s *GetVccResponseBodyContentErInfos) SetTenantId(v string) *GetVccResponseBodyContentErInfos {
	s.TenantId = &v
	return s
}

type GetVccResponseBodyContentTags struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetVccResponseBodyContentTags) String() string {
	return tea.Prettify(s)
}

func (s GetVccResponseBodyContentTags) GoString() string {
	return s.String()
}

func (s *GetVccResponseBodyContentTags) SetTagKey(v string) *GetVccResponseBodyContentTags {
	s.TagKey = &v
	return s
}

func (s *GetVccResponseBodyContentTags) SetTagValue(v string) *GetVccResponseBodyContentTags {
	s.TagValue = &v
	return s
}

type GetVccResponseBodyContentVpdBaseInfo struct {
	Cidr       *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	VpdId      *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	VpdName    *string `json:"VpdName,omitempty" xml:"VpdName,omitempty"`
}

func (s GetVccResponseBodyContentVpdBaseInfo) String() string {
	return tea.Prettify(s)
}

func (s GetVccResponseBodyContentVpdBaseInfo) GoString() string {
	return s.String()
}

func (s *GetVccResponseBodyContentVpdBaseInfo) SetCidr(v string) *GetVccResponseBodyContentVpdBaseInfo {
	s.Cidr = &v
	return s
}

func (s *GetVccResponseBodyContentVpdBaseInfo) SetCreateTime(v string) *GetVccResponseBodyContentVpdBaseInfo {
	s.CreateTime = &v
	return s
}

func (s *GetVccResponseBodyContentVpdBaseInfo) SetVpdId(v string) *GetVccResponseBodyContentVpdBaseInfo {
	s.VpdId = &v
	return s
}

func (s *GetVccResponseBodyContentVpdBaseInfo) SetVpdName(v string) *GetVccResponseBodyContentVpdBaseInfo {
	s.VpdName = &v
	return s
}

type GetVccResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetVccResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetVccResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVccResponse) GoString() string {
	return s.String()
}

func (s *GetVccResponse) SetHeaders(v map[string]*string) *GetVccResponse {
	s.Headers = v
	return s
}

func (s *GetVccResponse) SetStatusCode(v int32) *GetVccResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVccResponse) SetBody(v *GetVccResponseBody) *GetVccResponse {
	s.Body = v
	return s
}

type GetVccGrantRuleRequest struct {
	ErId          *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	GrantRuleId   *string `json:"GrantRuleId,omitempty" xml:"GrantRuleId,omitempty"`
	GrantTenantId *string `json:"GrantTenantId,omitempty" xml:"GrantTenantId,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetVccGrantRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVccGrantRuleRequest) GoString() string {
	return s.String()
}

func (s *GetVccGrantRuleRequest) SetErId(v string) *GetVccGrantRuleRequest {
	s.ErId = &v
	return s
}

func (s *GetVccGrantRuleRequest) SetGrantRuleId(v string) *GetVccGrantRuleRequest {
	s.GrantRuleId = &v
	return s
}

func (s *GetVccGrantRuleRequest) SetGrantTenantId(v string) *GetVccGrantRuleRequest {
	s.GrantTenantId = &v
	return s
}

func (s *GetVccGrantRuleRequest) SetInstanceId(v string) *GetVccGrantRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *GetVccGrantRuleRequest) SetRegionId(v string) *GetVccGrantRuleRequest {
	s.RegionId = &v
	return s
}

type GetVccGrantRuleResponseBody struct {
	Code      *int32                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   *GetVccGrantRuleResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVccGrantRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVccGrantRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetVccGrantRuleResponseBody) SetCode(v int32) *GetVccGrantRuleResponseBody {
	s.Code = &v
	return s
}

func (s *GetVccGrantRuleResponseBody) SetContent(v *GetVccGrantRuleResponseBodyContent) *GetVccGrantRuleResponseBody {
	s.Content = v
	return s
}

func (s *GetVccGrantRuleResponseBody) SetMessage(v string) *GetVccGrantRuleResponseBody {
	s.Message = &v
	return s
}

func (s *GetVccGrantRuleResponseBody) SetRequestId(v string) *GetVccGrantRuleResponseBody {
	s.RequestId = &v
	return s
}

type GetVccGrantRuleResponseBodyContent struct {
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ErId          *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	GrantRuleId   *string `json:"GrantRuleId,omitempty" xml:"GrantRuleId,omitempty"`
	GrantTenantId *string `json:"GrantTenantId,omitempty" xml:"GrantTenantId,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName  *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Product       *string `json:"Product,omitempty" xml:"Product,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TenantId      *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Used          *bool   `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s GetVccGrantRuleResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s GetVccGrantRuleResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetVccGrantRuleResponseBodyContent) SetCreateTime(v string) *GetVccGrantRuleResponseBodyContent {
	s.CreateTime = &v
	return s
}

func (s *GetVccGrantRuleResponseBodyContent) SetErId(v string) *GetVccGrantRuleResponseBodyContent {
	s.ErId = &v
	return s
}

func (s *GetVccGrantRuleResponseBodyContent) SetGrantRuleId(v string) *GetVccGrantRuleResponseBodyContent {
	s.GrantRuleId = &v
	return s
}

func (s *GetVccGrantRuleResponseBodyContent) SetGrantTenantId(v string) *GetVccGrantRuleResponseBodyContent {
	s.GrantTenantId = &v
	return s
}

func (s *GetVccGrantRuleResponseBodyContent) SetInstanceId(v string) *GetVccGrantRuleResponseBodyContent {
	s.InstanceId = &v
	return s
}

func (s *GetVccGrantRuleResponseBodyContent) SetInstanceName(v string) *GetVccGrantRuleResponseBodyContent {
	s.InstanceName = &v
	return s
}

func (s *GetVccGrantRuleResponseBodyContent) SetProduct(v string) *GetVccGrantRuleResponseBodyContent {
	s.Product = &v
	return s
}

func (s *GetVccGrantRuleResponseBodyContent) SetRegionId(v string) *GetVccGrantRuleResponseBodyContent {
	s.RegionId = &v
	return s
}

func (s *GetVccGrantRuleResponseBodyContent) SetTenantId(v string) *GetVccGrantRuleResponseBodyContent {
	s.TenantId = &v
	return s
}

func (s *GetVccGrantRuleResponseBodyContent) SetUsed(v bool) *GetVccGrantRuleResponseBodyContent {
	s.Used = &v
	return s
}

type GetVccGrantRuleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetVccGrantRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetVccGrantRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVccGrantRuleResponse) GoString() string {
	return s.String()
}

func (s *GetVccGrantRuleResponse) SetHeaders(v map[string]*string) *GetVccGrantRuleResponse {
	s.Headers = v
	return s
}

func (s *GetVccGrantRuleResponse) SetStatusCode(v int32) *GetVccGrantRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVccGrantRuleResponse) SetBody(v *GetVccGrantRuleResponseBody) *GetVccGrantRuleResponse {
	s.Body = v
	return s
}

type GetVccRouteEntryRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VccId           *string `json:"VccId,omitempty" xml:"VccId,omitempty"`
	VccRouteEntryId *string `json:"VccRouteEntryId,omitempty" xml:"VccRouteEntryId,omitempty"`
}

func (s GetVccRouteEntryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVccRouteEntryRequest) GoString() string {
	return s.String()
}

func (s *GetVccRouteEntryRequest) SetRegionId(v string) *GetVccRouteEntryRequest {
	s.RegionId = &v
	return s
}

func (s *GetVccRouteEntryRequest) SetVccId(v string) *GetVccRouteEntryRequest {
	s.VccId = &v
	return s
}

func (s *GetVccRouteEntryRequest) SetVccRouteEntryId(v string) *GetVccRouteEntryRequest {
	s.VccRouteEntryId = &v
	return s
}

type GetVccRouteEntryResponseBody struct {
	Code      *int32                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   *GetVccRouteEntryResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVccRouteEntryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVccRouteEntryResponseBody) GoString() string {
	return s.String()
}

func (s *GetVccRouteEntryResponseBody) SetCode(v int32) *GetVccRouteEntryResponseBody {
	s.Code = &v
	return s
}

func (s *GetVccRouteEntryResponseBody) SetContent(v *GetVccRouteEntryResponseBodyContent) *GetVccRouteEntryResponseBody {
	s.Content = v
	return s
}

func (s *GetVccRouteEntryResponseBody) SetMessage(v string) *GetVccRouteEntryResponseBody {
	s.Message = &v
	return s
}

func (s *GetVccRouteEntryResponseBody) SetRequestId(v string) *GetVccRouteEntryResponseBody {
	s.RequestId = &v
	return s
}

type GetVccRouteEntryResponseBodyContent struct {
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	GmtModified          *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	NextHopId            *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	NextHopType          *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RouteType            *string `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TenantId             *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	VccId                *string `json:"VccId,omitempty" xml:"VccId,omitempty"`
	VccRouteEntryId      *string `json:"VccRouteEntryId,omitempty" xml:"VccRouteEntryId,omitempty"`
}

func (s GetVccRouteEntryResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s GetVccRouteEntryResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetVccRouteEntryResponseBodyContent) SetDestinationCidrBlock(v string) *GetVccRouteEntryResponseBodyContent {
	s.DestinationCidrBlock = &v
	return s
}

func (s *GetVccRouteEntryResponseBodyContent) SetGmtModified(v string) *GetVccRouteEntryResponseBodyContent {
	s.GmtModified = &v
	return s
}

func (s *GetVccRouteEntryResponseBodyContent) SetNextHopId(v string) *GetVccRouteEntryResponseBodyContent {
	s.NextHopId = &v
	return s
}

func (s *GetVccRouteEntryResponseBodyContent) SetNextHopType(v string) *GetVccRouteEntryResponseBodyContent {
	s.NextHopType = &v
	return s
}

func (s *GetVccRouteEntryResponseBodyContent) SetRegionId(v string) *GetVccRouteEntryResponseBodyContent {
	s.RegionId = &v
	return s
}

func (s *GetVccRouteEntryResponseBodyContent) SetRouteType(v string) *GetVccRouteEntryResponseBodyContent {
	s.RouteType = &v
	return s
}

func (s *GetVccRouteEntryResponseBodyContent) SetStatus(v string) *GetVccRouteEntryResponseBodyContent {
	s.Status = &v
	return s
}

func (s *GetVccRouteEntryResponseBodyContent) SetTenantId(v string) *GetVccRouteEntryResponseBodyContent {
	s.TenantId = &v
	return s
}

func (s *GetVccRouteEntryResponseBodyContent) SetVccId(v string) *GetVccRouteEntryResponseBodyContent {
	s.VccId = &v
	return s
}

func (s *GetVccRouteEntryResponseBodyContent) SetVccRouteEntryId(v string) *GetVccRouteEntryResponseBodyContent {
	s.VccRouteEntryId = &v
	return s
}

type GetVccRouteEntryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetVccRouteEntryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetVccRouteEntryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVccRouteEntryResponse) GoString() string {
	return s.String()
}

func (s *GetVccRouteEntryResponse) SetHeaders(v map[string]*string) *GetVccRouteEntryResponse {
	s.Headers = v
	return s
}

func (s *GetVccRouteEntryResponse) SetStatusCode(v int32) *GetVccRouteEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVccRouteEntryResponse) SetBody(v *GetVccRouteEntryResponseBody) *GetVccRouteEntryResponse {
	s.Body = v
	return s
}

type GetVpdRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VpdId    *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
}

func (s GetVpdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVpdRequest) GoString() string {
	return s.String()
}

func (s *GetVpdRequest) SetRegionId(v string) *GetVpdRequest {
	s.RegionId = &v
	return s
}

func (s *GetVpdRequest) SetVpdId(v string) *GetVpdRequest {
	s.VpdId = &v
	return s
}

type GetVpdResponseBody struct {
	Code      *int32                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   *GetVpdResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Message   *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVpdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVpdResponseBody) GoString() string {
	return s.String()
}

func (s *GetVpdResponseBody) SetCode(v int32) *GetVpdResponseBody {
	s.Code = &v
	return s
}

func (s *GetVpdResponseBody) SetContent(v *GetVpdResponseBodyContent) *GetVpdResponseBody {
	s.Content = v
	return s
}

func (s *GetVpdResponseBody) SetMessage(v string) *GetVpdResponseBody {
	s.Message = &v
	return s
}

func (s *GetVpdResponseBody) SetRequestId(v string) *GetVpdResponseBody {
	s.RequestId = &v
	return s
}

type GetVpdResponseBodyContent struct {
	AttachErStatus  *bool                               `json:"AttachErStatus,omitempty" xml:"AttachErStatus,omitempty"`
	Cidr            *string                             `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	CreateTime      *string                             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ErInfos         []*GetVpdResponseBodyContentErInfos `json:"ErInfos,omitempty" xml:"ErInfos,omitempty" type:"Repeated"`
	GmtModified     *string                             `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Message         *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	NcCount         *int64                              `json:"NcCount,omitempty" xml:"NcCount,omitempty"`
	RegionId        *string                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string                             `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ServiceCidr     *string                             `json:"ServiceCidr,omitempty" xml:"ServiceCidr,omitempty"`
	Status          *string                             `json:"Status,omitempty" xml:"Status,omitempty"`
	SubnetCount     *int64                              `json:"SubnetCount,omitempty" xml:"SubnetCount,omitempty"`
	Tags            []*GetVpdResponseBodyContentTags    `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TenantId        *string                             `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	VpdId           *string                             `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	VpdName         *string                             `json:"VpdName,omitempty" xml:"VpdName,omitempty"`
}

func (s GetVpdResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s GetVpdResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetVpdResponseBodyContent) SetAttachErStatus(v bool) *GetVpdResponseBodyContent {
	s.AttachErStatus = &v
	return s
}

func (s *GetVpdResponseBodyContent) SetCidr(v string) *GetVpdResponseBodyContent {
	s.Cidr = &v
	return s
}

func (s *GetVpdResponseBodyContent) SetCreateTime(v string) *GetVpdResponseBodyContent {
	s.CreateTime = &v
	return s
}

func (s *GetVpdResponseBodyContent) SetErInfos(v []*GetVpdResponseBodyContentErInfos) *GetVpdResponseBodyContent {
	s.ErInfos = v
	return s
}

func (s *GetVpdResponseBodyContent) SetGmtModified(v string) *GetVpdResponseBodyContent {
	s.GmtModified = &v
	return s
}

func (s *GetVpdResponseBodyContent) SetMessage(v string) *GetVpdResponseBodyContent {
	s.Message = &v
	return s
}

func (s *GetVpdResponseBodyContent) SetNcCount(v int64) *GetVpdResponseBodyContent {
	s.NcCount = &v
	return s
}

func (s *GetVpdResponseBodyContent) SetRegionId(v string) *GetVpdResponseBodyContent {
	s.RegionId = &v
	return s
}

func (s *GetVpdResponseBodyContent) SetResourceGroupId(v string) *GetVpdResponseBodyContent {
	s.ResourceGroupId = &v
	return s
}

func (s *GetVpdResponseBodyContent) SetServiceCidr(v string) *GetVpdResponseBodyContent {
	s.ServiceCidr = &v
	return s
}

func (s *GetVpdResponseBodyContent) SetStatus(v string) *GetVpdResponseBodyContent {
	s.Status = &v
	return s
}

func (s *GetVpdResponseBodyContent) SetSubnetCount(v int64) *GetVpdResponseBodyContent {
	s.SubnetCount = &v
	return s
}

func (s *GetVpdResponseBodyContent) SetTags(v []*GetVpdResponseBodyContentTags) *GetVpdResponseBodyContent {
	s.Tags = v
	return s
}

func (s *GetVpdResponseBodyContent) SetTenantId(v string) *GetVpdResponseBodyContent {
	s.TenantId = &v
	return s
}

func (s *GetVpdResponseBodyContent) SetVpdId(v string) *GetVpdResponseBodyContent {
	s.VpdId = &v
	return s
}

func (s *GetVpdResponseBodyContent) SetVpdName(v string) *GetVpdResponseBodyContent {
	s.VpdName = &v
	return s
}

type GetVpdResponseBodyContentErInfos struct {
	Connections  *int64  `json:"Connections,omitempty" xml:"Connections,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ErId         *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	ErName       *string `json:"ErName,omitempty" xml:"ErName,omitempty"`
	GmtModified  *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	MasterZoneId *string `json:"MasterZoneId,omitempty" xml:"MasterZoneId,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RouteMaps    *int64  `json:"RouteMaps,omitempty" xml:"RouteMaps,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TenantId     *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetVpdResponseBodyContentErInfos) String() string {
	return tea.Prettify(s)
}

func (s GetVpdResponseBodyContentErInfos) GoString() string {
	return s.String()
}

func (s *GetVpdResponseBodyContentErInfos) SetConnections(v int64) *GetVpdResponseBodyContentErInfos {
	s.Connections = &v
	return s
}

func (s *GetVpdResponseBodyContentErInfos) SetCreateTime(v string) *GetVpdResponseBodyContentErInfos {
	s.CreateTime = &v
	return s
}

func (s *GetVpdResponseBodyContentErInfos) SetDescription(v string) *GetVpdResponseBodyContentErInfos {
	s.Description = &v
	return s
}

func (s *GetVpdResponseBodyContentErInfos) SetErId(v string) *GetVpdResponseBodyContentErInfos {
	s.ErId = &v
	return s
}

func (s *GetVpdResponseBodyContentErInfos) SetErName(v string) *GetVpdResponseBodyContentErInfos {
	s.ErName = &v
	return s
}

func (s *GetVpdResponseBodyContentErInfos) SetGmtModified(v string) *GetVpdResponseBodyContentErInfos {
	s.GmtModified = &v
	return s
}

func (s *GetVpdResponseBodyContentErInfos) SetMasterZoneId(v string) *GetVpdResponseBodyContentErInfos {
	s.MasterZoneId = &v
	return s
}

func (s *GetVpdResponseBodyContentErInfos) SetMessage(v string) *GetVpdResponseBodyContentErInfos {
	s.Message = &v
	return s
}

func (s *GetVpdResponseBodyContentErInfos) SetRegionId(v string) *GetVpdResponseBodyContentErInfos {
	s.RegionId = &v
	return s
}

func (s *GetVpdResponseBodyContentErInfos) SetRouteMaps(v int64) *GetVpdResponseBodyContentErInfos {
	s.RouteMaps = &v
	return s
}

func (s *GetVpdResponseBodyContentErInfos) SetStatus(v string) *GetVpdResponseBodyContentErInfos {
	s.Status = &v
	return s
}

func (s *GetVpdResponseBodyContentErInfos) SetTenantId(v string) *GetVpdResponseBodyContentErInfos {
	s.TenantId = &v
	return s
}

type GetVpdResponseBodyContentTags struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetVpdResponseBodyContentTags) String() string {
	return tea.Prettify(s)
}

func (s GetVpdResponseBodyContentTags) GoString() string {
	return s.String()
}

func (s *GetVpdResponseBodyContentTags) SetTagKey(v string) *GetVpdResponseBodyContentTags {
	s.TagKey = &v
	return s
}

func (s *GetVpdResponseBodyContentTags) SetTagValue(v string) *GetVpdResponseBodyContentTags {
	s.TagValue = &v
	return s
}

type GetVpdResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetVpdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetVpdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVpdResponse) GoString() string {
	return s.String()
}

func (s *GetVpdResponse) SetHeaders(v map[string]*string) *GetVpdResponse {
	s.Headers = v
	return s
}

func (s *GetVpdResponse) SetStatusCode(v int32) *GetVpdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVpdResponse) SetBody(v *GetVpdResponseBody) *GetVpdResponse {
	s.Body = v
	return s
}

type GetVpdGrantRuleRequest struct {
	ErId          *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	GrantRuleId   *string `json:"GrantRuleId,omitempty" xml:"GrantRuleId,omitempty"`
	GrantTenantId *string `json:"GrantTenantId,omitempty" xml:"GrantTenantId,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetVpdGrantRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVpdGrantRuleRequest) GoString() string {
	return s.String()
}

func (s *GetVpdGrantRuleRequest) SetErId(v string) *GetVpdGrantRuleRequest {
	s.ErId = &v
	return s
}

func (s *GetVpdGrantRuleRequest) SetGrantRuleId(v string) *GetVpdGrantRuleRequest {
	s.GrantRuleId = &v
	return s
}

func (s *GetVpdGrantRuleRequest) SetGrantTenantId(v string) *GetVpdGrantRuleRequest {
	s.GrantTenantId = &v
	return s
}

func (s *GetVpdGrantRuleRequest) SetInstanceId(v string) *GetVpdGrantRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *GetVpdGrantRuleRequest) SetRegionId(v string) *GetVpdGrantRuleRequest {
	s.RegionId = &v
	return s
}

type GetVpdGrantRuleResponseBody struct {
	Code      *int32                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   *GetVpdGrantRuleResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVpdGrantRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVpdGrantRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetVpdGrantRuleResponseBody) SetCode(v int32) *GetVpdGrantRuleResponseBody {
	s.Code = &v
	return s
}

func (s *GetVpdGrantRuleResponseBody) SetContent(v *GetVpdGrantRuleResponseBodyContent) *GetVpdGrantRuleResponseBody {
	s.Content = v
	return s
}

func (s *GetVpdGrantRuleResponseBody) SetMessage(v string) *GetVpdGrantRuleResponseBody {
	s.Message = &v
	return s
}

func (s *GetVpdGrantRuleResponseBody) SetRequestId(v string) *GetVpdGrantRuleResponseBody {
	s.RequestId = &v
	return s
}

type GetVpdGrantRuleResponseBodyContent struct {
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ErId          *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	GrantRuleId   *string `json:"GrantRuleId,omitempty" xml:"GrantRuleId,omitempty"`
	GrantTenantId *string `json:"GrantTenantId,omitempty" xml:"GrantTenantId,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName  *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Product       *string `json:"Product,omitempty" xml:"Product,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TenantId      *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Used          *bool   `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s GetVpdGrantRuleResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s GetVpdGrantRuleResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetVpdGrantRuleResponseBodyContent) SetCreateTime(v string) *GetVpdGrantRuleResponseBodyContent {
	s.CreateTime = &v
	return s
}

func (s *GetVpdGrantRuleResponseBodyContent) SetErId(v string) *GetVpdGrantRuleResponseBodyContent {
	s.ErId = &v
	return s
}

func (s *GetVpdGrantRuleResponseBodyContent) SetGrantRuleId(v string) *GetVpdGrantRuleResponseBodyContent {
	s.GrantRuleId = &v
	return s
}

func (s *GetVpdGrantRuleResponseBodyContent) SetGrantTenantId(v string) *GetVpdGrantRuleResponseBodyContent {
	s.GrantTenantId = &v
	return s
}

func (s *GetVpdGrantRuleResponseBodyContent) SetInstanceId(v string) *GetVpdGrantRuleResponseBodyContent {
	s.InstanceId = &v
	return s
}

func (s *GetVpdGrantRuleResponseBodyContent) SetInstanceName(v string) *GetVpdGrantRuleResponseBodyContent {
	s.InstanceName = &v
	return s
}

func (s *GetVpdGrantRuleResponseBodyContent) SetProduct(v string) *GetVpdGrantRuleResponseBodyContent {
	s.Product = &v
	return s
}

func (s *GetVpdGrantRuleResponseBodyContent) SetRegionId(v string) *GetVpdGrantRuleResponseBodyContent {
	s.RegionId = &v
	return s
}

func (s *GetVpdGrantRuleResponseBodyContent) SetTenantId(v string) *GetVpdGrantRuleResponseBodyContent {
	s.TenantId = &v
	return s
}

func (s *GetVpdGrantRuleResponseBodyContent) SetUsed(v bool) *GetVpdGrantRuleResponseBodyContent {
	s.Used = &v
	return s
}

type GetVpdGrantRuleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetVpdGrantRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetVpdGrantRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVpdGrantRuleResponse) GoString() string {
	return s.String()
}

func (s *GetVpdGrantRuleResponse) SetHeaders(v map[string]*string) *GetVpdGrantRuleResponse {
	s.Headers = v
	return s
}

func (s *GetVpdGrantRuleResponse) SetStatusCode(v int32) *GetVpdGrantRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVpdGrantRuleResponse) SetBody(v *GetVpdGrantRuleResponseBody) *GetVpdGrantRuleResponse {
	s.Body = v
	return s
}

type GetVpdRouteEntryRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VpdId           *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	VpdRouteEntryId *string `json:"VpdRouteEntryId,omitempty" xml:"VpdRouteEntryId,omitempty"`
}

func (s GetVpdRouteEntryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVpdRouteEntryRequest) GoString() string {
	return s.String()
}

func (s *GetVpdRouteEntryRequest) SetRegionId(v string) *GetVpdRouteEntryRequest {
	s.RegionId = &v
	return s
}

func (s *GetVpdRouteEntryRequest) SetVpdId(v string) *GetVpdRouteEntryRequest {
	s.VpdId = &v
	return s
}

func (s *GetVpdRouteEntryRequest) SetVpdRouteEntryId(v string) *GetVpdRouteEntryRequest {
	s.VpdRouteEntryId = &v
	return s
}

type GetVpdRouteEntryResponseBody struct {
	Code      *int32                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   *GetVpdRouteEntryResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVpdRouteEntryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVpdRouteEntryResponseBody) GoString() string {
	return s.String()
}

func (s *GetVpdRouteEntryResponseBody) SetCode(v int32) *GetVpdRouteEntryResponseBody {
	s.Code = &v
	return s
}

func (s *GetVpdRouteEntryResponseBody) SetContent(v *GetVpdRouteEntryResponseBodyContent) *GetVpdRouteEntryResponseBody {
	s.Content = v
	return s
}

func (s *GetVpdRouteEntryResponseBody) SetMessage(v string) *GetVpdRouteEntryResponseBody {
	s.Message = &v
	return s
}

func (s *GetVpdRouteEntryResponseBody) SetRequestId(v string) *GetVpdRouteEntryResponseBody {
	s.RequestId = &v
	return s
}

type GetVpdRouteEntryResponseBodyContent struct {
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	GmtModified          *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	NextHopId            *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	NextHopType          *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RouteType            *string `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TenantId             *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	VpdId                *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	VpdRouteEntryId      *string `json:"VpdRouteEntryId,omitempty" xml:"VpdRouteEntryId,omitempty"`
}

func (s GetVpdRouteEntryResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s GetVpdRouteEntryResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetVpdRouteEntryResponseBodyContent) SetDestinationCidrBlock(v string) *GetVpdRouteEntryResponseBodyContent {
	s.DestinationCidrBlock = &v
	return s
}

func (s *GetVpdRouteEntryResponseBodyContent) SetGmtModified(v string) *GetVpdRouteEntryResponseBodyContent {
	s.GmtModified = &v
	return s
}

func (s *GetVpdRouteEntryResponseBodyContent) SetNextHopId(v string) *GetVpdRouteEntryResponseBodyContent {
	s.NextHopId = &v
	return s
}

func (s *GetVpdRouteEntryResponseBodyContent) SetNextHopType(v string) *GetVpdRouteEntryResponseBodyContent {
	s.NextHopType = &v
	return s
}

func (s *GetVpdRouteEntryResponseBodyContent) SetRegionId(v string) *GetVpdRouteEntryResponseBodyContent {
	s.RegionId = &v
	return s
}

func (s *GetVpdRouteEntryResponseBodyContent) SetRouteType(v string) *GetVpdRouteEntryResponseBodyContent {
	s.RouteType = &v
	return s
}

func (s *GetVpdRouteEntryResponseBodyContent) SetStatus(v string) *GetVpdRouteEntryResponseBodyContent {
	s.Status = &v
	return s
}

func (s *GetVpdRouteEntryResponseBodyContent) SetTenantId(v string) *GetVpdRouteEntryResponseBodyContent {
	s.TenantId = &v
	return s
}

func (s *GetVpdRouteEntryResponseBodyContent) SetVpdId(v string) *GetVpdRouteEntryResponseBodyContent {
	s.VpdId = &v
	return s
}

func (s *GetVpdRouteEntryResponseBodyContent) SetVpdRouteEntryId(v string) *GetVpdRouteEntryResponseBodyContent {
	s.VpdRouteEntryId = &v
	return s
}

type GetVpdRouteEntryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetVpdRouteEntryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetVpdRouteEntryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVpdRouteEntryResponse) GoString() string {
	return s.String()
}

func (s *GetVpdRouteEntryResponse) SetHeaders(v map[string]*string) *GetVpdRouteEntryResponse {
	s.Headers = v
	return s
}

func (s *GetVpdRouteEntryResponse) SetStatusCode(v int32) *GetVpdRouteEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVpdRouteEntryResponse) SetBody(v *GetVpdRouteEntryResponseBody) *GetVpdRouteEntryResponse {
	s.Body = v
	return s
}

type InitializeVccRequest struct {
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s InitializeVccRequest) String() string {
	return tea.Prettify(s)
}

func (s InitializeVccRequest) GoString() string {
	return s.String()
}

func (s *InitializeVccRequest) SetResourceGroupId(v string) *InitializeVccRequest {
	s.ResourceGroupId = &v
	return s
}

type InitializeVccResponseBody struct {
	Code      *int32                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   *InitializeVccResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InitializeVccResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InitializeVccResponseBody) GoString() string {
	return s.String()
}

func (s *InitializeVccResponseBody) SetCode(v int32) *InitializeVccResponseBody {
	s.Code = &v
	return s
}

func (s *InitializeVccResponseBody) SetContent(v *InitializeVccResponseBodyContent) *InitializeVccResponseBody {
	s.Content = v
	return s
}

func (s *InitializeVccResponseBody) SetMessage(v string) *InitializeVccResponseBody {
	s.Message = &v
	return s
}

func (s *InitializeVccResponseBody) SetRequestId(v string) *InitializeVccResponseBody {
	s.RequestId = &v
	return s
}

type InitializeVccResponseBodyContent struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RoleName  *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s InitializeVccResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s InitializeVccResponseBodyContent) GoString() string {
	return s.String()
}

func (s *InitializeVccResponseBodyContent) SetRequestId(v string) *InitializeVccResponseBodyContent {
	s.RequestId = &v
	return s
}

func (s *InitializeVccResponseBodyContent) SetRoleName(v string) *InitializeVccResponseBodyContent {
	s.RoleName = &v
	return s
}

type InitializeVccResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InitializeVccResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InitializeVccResponse) String() string {
	return tea.Prettify(s)
}

func (s InitializeVccResponse) GoString() string {
	return s.String()
}

func (s *InitializeVccResponse) SetHeaders(v map[string]*string) *InitializeVccResponse {
	s.Headers = v
	return s
}

func (s *InitializeVccResponse) SetStatusCode(v int32) *InitializeVccResponse {
	s.StatusCode = &v
	return s
}

func (s *InitializeVccResponse) SetBody(v *InitializeVccResponseBody) *InitializeVccResponse {
	s.Body = v
	return s
}

type ListErAttachmentsRequest struct {
	AutoReceiveAllRoute *bool   `json:"AutoReceiveAllRoute,omitempty" xml:"AutoReceiveAllRoute,omitempty"`
	EnablePage          *bool   `json:"EnablePage,omitempty" xml:"EnablePage,omitempty"`
	ErAttachmentId      *string `json:"ErAttachmentId,omitempty" xml:"ErAttachmentId,omitempty"`
	ErAttachmentName    *string `json:"ErAttachmentName,omitempty" xml:"ErAttachmentName,omitempty"`
	ErId                *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType        *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	PageNumber          *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize            *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceTenantId    *string `json:"ResourceTenantId,omitempty" xml:"ResourceTenantId,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListErAttachmentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListErAttachmentsRequest) GoString() string {
	return s.String()
}

func (s *ListErAttachmentsRequest) SetAutoReceiveAllRoute(v bool) *ListErAttachmentsRequest {
	s.AutoReceiveAllRoute = &v
	return s
}

func (s *ListErAttachmentsRequest) SetEnablePage(v bool) *ListErAttachmentsRequest {
	s.EnablePage = &v
	return s
}

func (s *ListErAttachmentsRequest) SetErAttachmentId(v string) *ListErAttachmentsRequest {
	s.ErAttachmentId = &v
	return s
}

func (s *ListErAttachmentsRequest) SetErAttachmentName(v string) *ListErAttachmentsRequest {
	s.ErAttachmentName = &v
	return s
}

func (s *ListErAttachmentsRequest) SetErId(v string) *ListErAttachmentsRequest {
	s.ErId = &v
	return s
}

func (s *ListErAttachmentsRequest) SetInstanceId(v string) *ListErAttachmentsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListErAttachmentsRequest) SetInstanceType(v string) *ListErAttachmentsRequest {
	s.InstanceType = &v
	return s
}

func (s *ListErAttachmentsRequest) SetPageNumber(v int32) *ListErAttachmentsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListErAttachmentsRequest) SetPageSize(v int32) *ListErAttachmentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListErAttachmentsRequest) SetRegionId(v string) *ListErAttachmentsRequest {
	s.RegionId = &v
	return s
}

func (s *ListErAttachmentsRequest) SetResourceTenantId(v string) *ListErAttachmentsRequest {
	s.ResourceTenantId = &v
	return s
}

func (s *ListErAttachmentsRequest) SetStatus(v string) *ListErAttachmentsRequest {
	s.Status = &v
	return s
}

type ListErAttachmentsResponseBody struct {
	Code      *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   *ListErAttachmentsResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListErAttachmentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListErAttachmentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListErAttachmentsResponseBody) SetCode(v int32) *ListErAttachmentsResponseBody {
	s.Code = &v
	return s
}

func (s *ListErAttachmentsResponseBody) SetContent(v *ListErAttachmentsResponseBodyContent) *ListErAttachmentsResponseBody {
	s.Content = v
	return s
}

func (s *ListErAttachmentsResponseBody) SetMessage(v string) *ListErAttachmentsResponseBody {
	s.Message = &v
	return s
}

func (s *ListErAttachmentsResponseBody) SetRequestId(v string) *ListErAttachmentsResponseBody {
	s.RequestId = &v
	return s
}

type ListErAttachmentsResponseBodyContent struct {
	Data  []*ListErAttachmentsResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Total *int64                                      `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListErAttachmentsResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s ListErAttachmentsResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListErAttachmentsResponseBodyContent) SetData(v []*ListErAttachmentsResponseBodyContentData) *ListErAttachmentsResponseBodyContent {
	s.Data = v
	return s
}

func (s *ListErAttachmentsResponseBodyContent) SetTotal(v int64) *ListErAttachmentsResponseBodyContent {
	s.Total = &v
	return s
}

type ListErAttachmentsResponseBodyContentData struct {
	Across              *bool   `json:"Across,omitempty" xml:"Across,omitempty"`
	AutoReceiveAllRoute *bool   `json:"AutoReceiveAllRoute,omitempty" xml:"AutoReceiveAllRoute,omitempty"`
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ErAttachmentId      *string `json:"ErAttachmentId,omitempty" xml:"ErAttachmentId,omitempty"`
	ErAttachmentName    *string `json:"ErAttachmentName,omitempty" xml:"ErAttachmentName,omitempty"`
	ErId                *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	GmtModified         *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName        *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceType        *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Message             *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceTenantId    *string `json:"ResourceTenantId,omitempty" xml:"ResourceTenantId,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TenantId            *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListErAttachmentsResponseBodyContentData) String() string {
	return tea.Prettify(s)
}

func (s ListErAttachmentsResponseBodyContentData) GoString() string {
	return s.String()
}

func (s *ListErAttachmentsResponseBodyContentData) SetAcross(v bool) *ListErAttachmentsResponseBodyContentData {
	s.Across = &v
	return s
}

func (s *ListErAttachmentsResponseBodyContentData) SetAutoReceiveAllRoute(v bool) *ListErAttachmentsResponseBodyContentData {
	s.AutoReceiveAllRoute = &v
	return s
}

func (s *ListErAttachmentsResponseBodyContentData) SetCreateTime(v string) *ListErAttachmentsResponseBodyContentData {
	s.CreateTime = &v
	return s
}

func (s *ListErAttachmentsResponseBodyContentData) SetErAttachmentId(v string) *ListErAttachmentsResponseBodyContentData {
	s.ErAttachmentId = &v
	return s
}

func (s *ListErAttachmentsResponseBodyContentData) SetErAttachmentName(v string) *ListErAttachmentsResponseBodyContentData {
	s.ErAttachmentName = &v
	return s
}

func (s *ListErAttachmentsResponseBodyContentData) SetErId(v string) *ListErAttachmentsResponseBodyContentData {
	s.ErId = &v
	return s
}

func (s *ListErAttachmentsResponseBodyContentData) SetGmtModified(v string) *ListErAttachmentsResponseBodyContentData {
	s.GmtModified = &v
	return s
}

func (s *ListErAttachmentsResponseBodyContentData) SetInstanceId(v string) *ListErAttachmentsResponseBodyContentData {
	s.InstanceId = &v
	return s
}

func (s *ListErAttachmentsResponseBodyContentData) SetInstanceName(v string) *ListErAttachmentsResponseBodyContentData {
	s.InstanceName = &v
	return s
}

func (s *ListErAttachmentsResponseBodyContentData) SetInstanceType(v string) *ListErAttachmentsResponseBodyContentData {
	s.InstanceType = &v
	return s
}

func (s *ListErAttachmentsResponseBodyContentData) SetMessage(v string) *ListErAttachmentsResponseBodyContentData {
	s.Message = &v
	return s
}

func (s *ListErAttachmentsResponseBodyContentData) SetRegionId(v string) *ListErAttachmentsResponseBodyContentData {
	s.RegionId = &v
	return s
}

func (s *ListErAttachmentsResponseBodyContentData) SetResourceTenantId(v string) *ListErAttachmentsResponseBodyContentData {
	s.ResourceTenantId = &v
	return s
}

func (s *ListErAttachmentsResponseBodyContentData) SetStatus(v string) *ListErAttachmentsResponseBodyContentData {
	s.Status = &v
	return s
}

func (s *ListErAttachmentsResponseBodyContentData) SetTenantId(v string) *ListErAttachmentsResponseBodyContentData {
	s.TenantId = &v
	return s
}

type ListErAttachmentsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListErAttachmentsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListErAttachmentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListErAttachmentsResponse) GoString() string {
	return s.String()
}

func (s *ListErAttachmentsResponse) SetHeaders(v map[string]*string) *ListErAttachmentsResponse {
	s.Headers = v
	return s
}

func (s *ListErAttachmentsResponse) SetStatusCode(v int32) *ListErAttachmentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListErAttachmentsResponse) SetBody(v *ListErAttachmentsResponseBody) *ListErAttachmentsResponse {
	s.Body = v
	return s
}

type ListErRouteEntriesRequest struct {
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	EnablePage           *bool   `json:"EnablePage,omitempty" xml:"EnablePage,omitempty"`
	ErId                 *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NextHopId            *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	NextHopType          *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RouteType            *string `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListErRouteEntriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListErRouteEntriesRequest) GoString() string {
	return s.String()
}

func (s *ListErRouteEntriesRequest) SetDestinationCidrBlock(v string) *ListErRouteEntriesRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *ListErRouteEntriesRequest) SetEnablePage(v bool) *ListErRouteEntriesRequest {
	s.EnablePage = &v
	return s
}

func (s *ListErRouteEntriesRequest) SetErId(v string) *ListErRouteEntriesRequest {
	s.ErId = &v
	return s
}

func (s *ListErRouteEntriesRequest) SetInstanceId(v string) *ListErRouteEntriesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListErRouteEntriesRequest) SetNextHopId(v string) *ListErRouteEntriesRequest {
	s.NextHopId = &v
	return s
}

func (s *ListErRouteEntriesRequest) SetNextHopType(v string) *ListErRouteEntriesRequest {
	s.NextHopType = &v
	return s
}

func (s *ListErRouteEntriesRequest) SetPageNumber(v int32) *ListErRouteEntriesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListErRouteEntriesRequest) SetPageSize(v int32) *ListErRouteEntriesRequest {
	s.PageSize = &v
	return s
}

func (s *ListErRouteEntriesRequest) SetRegionId(v string) *ListErRouteEntriesRequest {
	s.RegionId = &v
	return s
}

func (s *ListErRouteEntriesRequest) SetRouteType(v string) *ListErRouteEntriesRequest {
	s.RouteType = &v
	return s
}

func (s *ListErRouteEntriesRequest) SetStatus(v string) *ListErRouteEntriesRequest {
	s.Status = &v
	return s
}

type ListErRouteEntriesResponseBody struct {
	Code      *int32                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   *ListErRouteEntriesResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListErRouteEntriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListErRouteEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListErRouteEntriesResponseBody) SetCode(v int32) *ListErRouteEntriesResponseBody {
	s.Code = &v
	return s
}

func (s *ListErRouteEntriesResponseBody) SetContent(v *ListErRouteEntriesResponseBodyContent) *ListErRouteEntriesResponseBody {
	s.Content = v
	return s
}

func (s *ListErRouteEntriesResponseBody) SetMessage(v string) *ListErRouteEntriesResponseBody {
	s.Message = &v
	return s
}

func (s *ListErRouteEntriesResponseBody) SetRequestId(v string) *ListErRouteEntriesResponseBody {
	s.RequestId = &v
	return s
}

type ListErRouteEntriesResponseBodyContent struct {
	Data  []*ListErRouteEntriesResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Total *int64                                       `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListErRouteEntriesResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s ListErRouteEntriesResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListErRouteEntriesResponseBodyContent) SetData(v []*ListErRouteEntriesResponseBodyContentData) *ListErRouteEntriesResponseBodyContent {
	s.Data = v
	return s
}

func (s *ListErRouteEntriesResponseBodyContent) SetTotal(v int64) *ListErRouteEntriesResponseBodyContent {
	s.Total = &v
	return s
}

type ListErRouteEntriesResponseBodyContentData struct {
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	ErId                 *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	ErRouteEntryId       *string `json:"ErRouteEntryId,omitempty" xml:"ErRouteEntryId,omitempty"`
	GmtModified          *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	NextHopId            *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	NextHopType          *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceTenantId     *string `json:"ResourceTenantId,omitempty" xml:"ResourceTenantId,omitempty"`
	RouteType            *string `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TenantId             *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListErRouteEntriesResponseBodyContentData) String() string {
	return tea.Prettify(s)
}

func (s ListErRouteEntriesResponseBodyContentData) GoString() string {
	return s.String()
}

func (s *ListErRouteEntriesResponseBodyContentData) SetDestinationCidrBlock(v string) *ListErRouteEntriesResponseBodyContentData {
	s.DestinationCidrBlock = &v
	return s
}

func (s *ListErRouteEntriesResponseBodyContentData) SetErId(v string) *ListErRouteEntriesResponseBodyContentData {
	s.ErId = &v
	return s
}

func (s *ListErRouteEntriesResponseBodyContentData) SetErRouteEntryId(v string) *ListErRouteEntriesResponseBodyContentData {
	s.ErRouteEntryId = &v
	return s
}

func (s *ListErRouteEntriesResponseBodyContentData) SetGmtModified(v string) *ListErRouteEntriesResponseBodyContentData {
	s.GmtModified = &v
	return s
}

func (s *ListErRouteEntriesResponseBodyContentData) SetNextHopId(v string) *ListErRouteEntriesResponseBodyContentData {
	s.NextHopId = &v
	return s
}

func (s *ListErRouteEntriesResponseBodyContentData) SetNextHopType(v string) *ListErRouteEntriesResponseBodyContentData {
	s.NextHopType = &v
	return s
}

func (s *ListErRouteEntriesResponseBodyContentData) SetRegionId(v string) *ListErRouteEntriesResponseBodyContentData {
	s.RegionId = &v
	return s
}

func (s *ListErRouteEntriesResponseBodyContentData) SetResourceTenantId(v string) *ListErRouteEntriesResponseBodyContentData {
	s.ResourceTenantId = &v
	return s
}

func (s *ListErRouteEntriesResponseBodyContentData) SetRouteType(v string) *ListErRouteEntriesResponseBodyContentData {
	s.RouteType = &v
	return s
}

func (s *ListErRouteEntriesResponseBodyContentData) SetStatus(v string) *ListErRouteEntriesResponseBodyContentData {
	s.Status = &v
	return s
}

func (s *ListErRouteEntriesResponseBodyContentData) SetTenantId(v string) *ListErRouteEntriesResponseBodyContentData {
	s.TenantId = &v
	return s
}

type ListErRouteEntriesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListErRouteEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListErRouteEntriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListErRouteEntriesResponse) GoString() string {
	return s.String()
}

func (s *ListErRouteEntriesResponse) SetHeaders(v map[string]*string) *ListErRouteEntriesResponse {
	s.Headers = v
	return s
}

func (s *ListErRouteEntriesResponse) SetStatusCode(v int32) *ListErRouteEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListErRouteEntriesResponse) SetBody(v *ListErRouteEntriesResponseBody) *ListErRouteEntriesResponse {
	s.Body = v
	return s
}

type ListErRouteMapsRequest struct {
	DestinationCidrBlock     *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	EnablePage               *bool   `json:"EnablePage,omitempty" xml:"EnablePage,omitempty"`
	ErId                     *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	ErRouteMapId             *string `json:"ErRouteMapId,omitempty" xml:"ErRouteMapId,omitempty"`
	ErRouteMapNum            *int32  `json:"ErRouteMapNum,omitempty" xml:"ErRouteMapNum,omitempty"`
	PageNumber               *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize                 *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ReceptionInstanceId      *string `json:"ReceptionInstanceId,omitempty" xml:"ReceptionInstanceId,omitempty"`
	ReceptionInstanceName    *string `json:"ReceptionInstanceName,omitempty" xml:"ReceptionInstanceName,omitempty"`
	ReceptionInstanceType    *string `json:"ReceptionInstanceType,omitempty" xml:"ReceptionInstanceType,omitempty"`
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RouteMapAction           *string `json:"RouteMapAction,omitempty" xml:"RouteMapAction,omitempty"`
	TransmissionInstanceId   *string `json:"TransmissionInstanceId,omitempty" xml:"TransmissionInstanceId,omitempty"`
	TransmissionInstanceName *string `json:"TransmissionInstanceName,omitempty" xml:"TransmissionInstanceName,omitempty"`
	TransmissionInstanceType *string `json:"TransmissionInstanceType,omitempty" xml:"TransmissionInstanceType,omitempty"`
}

func (s ListErRouteMapsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListErRouteMapsRequest) GoString() string {
	return s.String()
}

func (s *ListErRouteMapsRequest) SetDestinationCidrBlock(v string) *ListErRouteMapsRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *ListErRouteMapsRequest) SetEnablePage(v bool) *ListErRouteMapsRequest {
	s.EnablePage = &v
	return s
}

func (s *ListErRouteMapsRequest) SetErId(v string) *ListErRouteMapsRequest {
	s.ErId = &v
	return s
}

func (s *ListErRouteMapsRequest) SetErRouteMapId(v string) *ListErRouteMapsRequest {
	s.ErRouteMapId = &v
	return s
}

func (s *ListErRouteMapsRequest) SetErRouteMapNum(v int32) *ListErRouteMapsRequest {
	s.ErRouteMapNum = &v
	return s
}

func (s *ListErRouteMapsRequest) SetPageNumber(v int32) *ListErRouteMapsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListErRouteMapsRequest) SetPageSize(v int32) *ListErRouteMapsRequest {
	s.PageSize = &v
	return s
}

func (s *ListErRouteMapsRequest) SetReceptionInstanceId(v string) *ListErRouteMapsRequest {
	s.ReceptionInstanceId = &v
	return s
}

func (s *ListErRouteMapsRequest) SetReceptionInstanceName(v string) *ListErRouteMapsRequest {
	s.ReceptionInstanceName = &v
	return s
}

func (s *ListErRouteMapsRequest) SetReceptionInstanceType(v string) *ListErRouteMapsRequest {
	s.ReceptionInstanceType = &v
	return s
}

func (s *ListErRouteMapsRequest) SetRegionId(v string) *ListErRouteMapsRequest {
	s.RegionId = &v
	return s
}

func (s *ListErRouteMapsRequest) SetRouteMapAction(v string) *ListErRouteMapsRequest {
	s.RouteMapAction = &v
	return s
}

func (s *ListErRouteMapsRequest) SetTransmissionInstanceId(v string) *ListErRouteMapsRequest {
	s.TransmissionInstanceId = &v
	return s
}

func (s *ListErRouteMapsRequest) SetTransmissionInstanceName(v string) *ListErRouteMapsRequest {
	s.TransmissionInstanceName = &v
	return s
}

func (s *ListErRouteMapsRequest) SetTransmissionInstanceType(v string) *ListErRouteMapsRequest {
	s.TransmissionInstanceType = &v
	return s
}

type ListErRouteMapsResponseBody struct {
	Code      *int32                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   *ListErRouteMapsResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListErRouteMapsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListErRouteMapsResponseBody) GoString() string {
	return s.String()
}

func (s *ListErRouteMapsResponseBody) SetCode(v int32) *ListErRouteMapsResponseBody {
	s.Code = &v
	return s
}

func (s *ListErRouteMapsResponseBody) SetContent(v *ListErRouteMapsResponseBodyContent) *ListErRouteMapsResponseBody {
	s.Content = v
	return s
}

func (s *ListErRouteMapsResponseBody) SetMessage(v string) *ListErRouteMapsResponseBody {
	s.Message = &v
	return s
}

func (s *ListErRouteMapsResponseBody) SetRequestId(v string) *ListErRouteMapsResponseBody {
	s.RequestId = &v
	return s
}

type ListErRouteMapsResponseBodyContent struct {
	Data  []*ListErRouteMapsResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Total *int64                                    `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListErRouteMapsResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s ListErRouteMapsResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListErRouteMapsResponseBodyContent) SetData(v []*ListErRouteMapsResponseBodyContentData) *ListErRouteMapsResponseBodyContent {
	s.Data = v
	return s
}

func (s *ListErRouteMapsResponseBodyContent) SetTotal(v int64) *ListErRouteMapsResponseBodyContent {
	s.Total = &v
	return s
}

type ListErRouteMapsResponseBodyContentData struct {
	Action                    *string `json:"Action,omitempty" xml:"Action,omitempty"`
	CreateTime                *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description               *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DestinationCidrBlock      *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	ErId                      *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	ErRouteMapId              *string `json:"ErRouteMapId,omitempty" xml:"ErRouteMapId,omitempty"`
	GmtModified               *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Message                   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	ReceptionInstanceId       *string `json:"ReceptionInstanceId,omitempty" xml:"ReceptionInstanceId,omitempty"`
	ReceptionInstanceName     *string `json:"ReceptionInstanceName,omitempty" xml:"ReceptionInstanceName,omitempty"`
	ReceptionInstanceOwner    *string `json:"ReceptionInstanceOwner,omitempty" xml:"ReceptionInstanceOwner,omitempty"`
	ReceptionInstanceType     *string `json:"ReceptionInstanceType,omitempty" xml:"ReceptionInstanceType,omitempty"`
	RegionId                  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RouteMapNum               *int32  `json:"RouteMapNum,omitempty" xml:"RouteMapNum,omitempty"`
	Status                    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TenantId                  *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	TransmissionInstanceId    *string `json:"TransmissionInstanceId,omitempty" xml:"TransmissionInstanceId,omitempty"`
	TransmissionInstanceName  *string `json:"TransmissionInstanceName,omitempty" xml:"TransmissionInstanceName,omitempty"`
	TransmissionInstanceOwner *string `json:"TransmissionInstanceOwner,omitempty" xml:"TransmissionInstanceOwner,omitempty"`
	TransmissionInstanceType  *string `json:"TransmissionInstanceType,omitempty" xml:"TransmissionInstanceType,omitempty"`
}

func (s ListErRouteMapsResponseBodyContentData) String() string {
	return tea.Prettify(s)
}

func (s ListErRouteMapsResponseBodyContentData) GoString() string {
	return s.String()
}

func (s *ListErRouteMapsResponseBodyContentData) SetAction(v string) *ListErRouteMapsResponseBodyContentData {
	s.Action = &v
	return s
}

func (s *ListErRouteMapsResponseBodyContentData) SetCreateTime(v string) *ListErRouteMapsResponseBodyContentData {
	s.CreateTime = &v
	return s
}

func (s *ListErRouteMapsResponseBodyContentData) SetDescription(v string) *ListErRouteMapsResponseBodyContentData {
	s.Description = &v
	return s
}

func (s *ListErRouteMapsResponseBodyContentData) SetDestinationCidrBlock(v string) *ListErRouteMapsResponseBodyContentData {
	s.DestinationCidrBlock = &v
	return s
}

func (s *ListErRouteMapsResponseBodyContentData) SetErId(v string) *ListErRouteMapsResponseBodyContentData {
	s.ErId = &v
	return s
}

func (s *ListErRouteMapsResponseBodyContentData) SetErRouteMapId(v string) *ListErRouteMapsResponseBodyContentData {
	s.ErRouteMapId = &v
	return s
}

func (s *ListErRouteMapsResponseBodyContentData) SetGmtModified(v string) *ListErRouteMapsResponseBodyContentData {
	s.GmtModified = &v
	return s
}

func (s *ListErRouteMapsResponseBodyContentData) SetMessage(v string) *ListErRouteMapsResponseBodyContentData {
	s.Message = &v
	return s
}

func (s *ListErRouteMapsResponseBodyContentData) SetReceptionInstanceId(v string) *ListErRouteMapsResponseBodyContentData {
	s.ReceptionInstanceId = &v
	return s
}

func (s *ListErRouteMapsResponseBodyContentData) SetReceptionInstanceName(v string) *ListErRouteMapsResponseBodyContentData {
	s.ReceptionInstanceName = &v
	return s
}

func (s *ListErRouteMapsResponseBodyContentData) SetReceptionInstanceOwner(v string) *ListErRouteMapsResponseBodyContentData {
	s.ReceptionInstanceOwner = &v
	return s
}

func (s *ListErRouteMapsResponseBodyContentData) SetReceptionInstanceType(v string) *ListErRouteMapsResponseBodyContentData {
	s.ReceptionInstanceType = &v
	return s
}

func (s *ListErRouteMapsResponseBodyContentData) SetRegionId(v string) *ListErRouteMapsResponseBodyContentData {
	s.RegionId = &v
	return s
}

func (s *ListErRouteMapsResponseBodyContentData) SetRouteMapNum(v int32) *ListErRouteMapsResponseBodyContentData {
	s.RouteMapNum = &v
	return s
}

func (s *ListErRouteMapsResponseBodyContentData) SetStatus(v string) *ListErRouteMapsResponseBodyContentData {
	s.Status = &v
	return s
}

func (s *ListErRouteMapsResponseBodyContentData) SetTenantId(v string) *ListErRouteMapsResponseBodyContentData {
	s.TenantId = &v
	return s
}

func (s *ListErRouteMapsResponseBodyContentData) SetTransmissionInstanceId(v string) *ListErRouteMapsResponseBodyContentData {
	s.TransmissionInstanceId = &v
	return s
}

func (s *ListErRouteMapsResponseBodyContentData) SetTransmissionInstanceName(v string) *ListErRouteMapsResponseBodyContentData {
	s.TransmissionInstanceName = &v
	return s
}

func (s *ListErRouteMapsResponseBodyContentData) SetTransmissionInstanceOwner(v string) *ListErRouteMapsResponseBodyContentData {
	s.TransmissionInstanceOwner = &v
	return s
}

func (s *ListErRouteMapsResponseBodyContentData) SetTransmissionInstanceType(v string) *ListErRouteMapsResponseBodyContentData {
	s.TransmissionInstanceType = &v
	return s
}

type ListErRouteMapsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListErRouteMapsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListErRouteMapsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListErRouteMapsResponse) GoString() string {
	return s.String()
}

func (s *ListErRouteMapsResponse) SetHeaders(v map[string]*string) *ListErRouteMapsResponse {
	s.Headers = v
	return s
}

func (s *ListErRouteMapsResponse) SetStatusCode(v int32) *ListErRouteMapsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListErRouteMapsResponse) SetBody(v *ListErRouteMapsResponseBody) *ListErRouteMapsResponse {
	s.Body = v
	return s
}

type ListErsRequest struct {
	EnablePage   *bool   `json:"EnablePage,omitempty" xml:"EnablePage,omitempty"`
	ErId         *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	ErName       *string `json:"ErName,omitempty" xml:"ErName,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	MasterZoneId *string `json:"MasterZoneId,omitempty" xml:"MasterZoneId,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListErsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListErsRequest) GoString() string {
	return s.String()
}

func (s *ListErsRequest) SetEnablePage(v bool) *ListErsRequest {
	s.EnablePage = &v
	return s
}

func (s *ListErsRequest) SetErId(v string) *ListErsRequest {
	s.ErId = &v
	return s
}

func (s *ListErsRequest) SetErName(v string) *ListErsRequest {
	s.ErName = &v
	return s
}

func (s *ListErsRequest) SetInstanceId(v string) *ListErsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListErsRequest) SetInstanceType(v string) *ListErsRequest {
	s.InstanceType = &v
	return s
}

func (s *ListErsRequest) SetMasterZoneId(v string) *ListErsRequest {
	s.MasterZoneId = &v
	return s
}

func (s *ListErsRequest) SetPageNumber(v int32) *ListErsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListErsRequest) SetPageSize(v int32) *ListErsRequest {
	s.PageSize = &v
	return s
}

func (s *ListErsRequest) SetRegionId(v string) *ListErsRequest {
	s.RegionId = &v
	return s
}

type ListErsResponseBody struct {
	Code      *int32                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   *ListErsResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Message   *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListErsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListErsResponseBody) GoString() string {
	return s.String()
}

func (s *ListErsResponseBody) SetCode(v int32) *ListErsResponseBody {
	s.Code = &v
	return s
}

func (s *ListErsResponseBody) SetContent(v *ListErsResponseBodyContent) *ListErsResponseBody {
	s.Content = v
	return s
}

func (s *ListErsResponseBody) SetMessage(v string) *ListErsResponseBody {
	s.Message = &v
	return s
}

func (s *ListErsResponseBody) SetRequestId(v string) *ListErsResponseBody {
	s.RequestId = &v
	return s
}

type ListErsResponseBodyContent struct {
	Data  []*ListErsResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Total *int64                            `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListErsResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s ListErsResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListErsResponseBodyContent) SetData(v []*ListErsResponseBodyContentData) *ListErsResponseBodyContent {
	s.Data = v
	return s
}

func (s *ListErsResponseBodyContent) SetTotal(v int64) *ListErsResponseBodyContent {
	s.Total = &v
	return s
}

type ListErsResponseBodyContentData struct {
	Connections  *int64  `json:"Connections,omitempty" xml:"Connections,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ErId         *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	ErName       *string `json:"ErName,omitempty" xml:"ErName,omitempty"`
	GmtModified  *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	MasterZoneId *string `json:"MasterZoneId,omitempty" xml:"MasterZoneId,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RouteMaps    *int64  `json:"RouteMaps,omitempty" xml:"RouteMaps,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TenantId     *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListErsResponseBodyContentData) String() string {
	return tea.Prettify(s)
}

func (s ListErsResponseBodyContentData) GoString() string {
	return s.String()
}

func (s *ListErsResponseBodyContentData) SetConnections(v int64) *ListErsResponseBodyContentData {
	s.Connections = &v
	return s
}

func (s *ListErsResponseBodyContentData) SetCreateTime(v string) *ListErsResponseBodyContentData {
	s.CreateTime = &v
	return s
}

func (s *ListErsResponseBodyContentData) SetDescription(v string) *ListErsResponseBodyContentData {
	s.Description = &v
	return s
}

func (s *ListErsResponseBodyContentData) SetErId(v string) *ListErsResponseBodyContentData {
	s.ErId = &v
	return s
}

func (s *ListErsResponseBodyContentData) SetErName(v string) *ListErsResponseBodyContentData {
	s.ErName = &v
	return s
}

func (s *ListErsResponseBodyContentData) SetGmtModified(v string) *ListErsResponseBodyContentData {
	s.GmtModified = &v
	return s
}

func (s *ListErsResponseBodyContentData) SetMasterZoneId(v string) *ListErsResponseBodyContentData {
	s.MasterZoneId = &v
	return s
}

func (s *ListErsResponseBodyContentData) SetMessage(v string) *ListErsResponseBodyContentData {
	s.Message = &v
	return s
}

func (s *ListErsResponseBodyContentData) SetRegionId(v string) *ListErsResponseBodyContentData {
	s.RegionId = &v
	return s
}

func (s *ListErsResponseBodyContentData) SetRouteMaps(v int64) *ListErsResponseBodyContentData {
	s.RouteMaps = &v
	return s
}

func (s *ListErsResponseBodyContentData) SetStatus(v string) *ListErsResponseBodyContentData {
	s.Status = &v
	return s
}

func (s *ListErsResponseBodyContentData) SetTenantId(v string) *ListErsResponseBodyContentData {
	s.TenantId = &v
	return s
}

type ListErsResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListErsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListErsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListErsResponse) GoString() string {
	return s.String()
}

func (s *ListErsResponse) SetHeaders(v map[string]*string) *ListErsResponse {
	s.Headers = v
	return s
}

func (s *ListErsResponse) SetStatusCode(v int32) *ListErsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListErsResponse) SetBody(v *ListErsResponseBody) *ListErsResponse {
	s.Body = v
	return s
}

type ListLniPrivateIpAddressRequest struct {
	EnablePage         *bool   `json:"EnablePage,omitempty" xml:"EnablePage,omitempty"`
	Ip                 *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	IpName             *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	PageNumber         *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListLniPrivateIpAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLniPrivateIpAddressRequest) GoString() string {
	return s.String()
}

func (s *ListLniPrivateIpAddressRequest) SetEnablePage(v bool) *ListLniPrivateIpAddressRequest {
	s.EnablePage = &v
	return s
}

func (s *ListLniPrivateIpAddressRequest) SetIp(v string) *ListLniPrivateIpAddressRequest {
	s.Ip = &v
	return s
}

func (s *ListLniPrivateIpAddressRequest) SetIpName(v string) *ListLniPrivateIpAddressRequest {
	s.IpName = &v
	return s
}

func (s *ListLniPrivateIpAddressRequest) SetNetworkInterfaceId(v string) *ListLniPrivateIpAddressRequest {
	s.NetworkInterfaceId = &v
	return s
}

func (s *ListLniPrivateIpAddressRequest) SetPageNumber(v int32) *ListLniPrivateIpAddressRequest {
	s.PageNumber = &v
	return s
}

func (s *ListLniPrivateIpAddressRequest) SetPageSize(v int32) *ListLniPrivateIpAddressRequest {
	s.PageSize = &v
	return s
}

func (s *ListLniPrivateIpAddressRequest) SetRegionId(v string) *ListLniPrivateIpAddressRequest {
	s.RegionId = &v
	return s
}

type ListLniPrivateIpAddressResponseBody struct {
	Code      *int32                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   *ListLniPrivateIpAddressResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Message   *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListLniPrivateIpAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLniPrivateIpAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ListLniPrivateIpAddressResponseBody) SetCode(v int32) *ListLniPrivateIpAddressResponseBody {
	s.Code = &v
	return s
}

func (s *ListLniPrivateIpAddressResponseBody) SetContent(v *ListLniPrivateIpAddressResponseBodyContent) *ListLniPrivateIpAddressResponseBody {
	s.Content = v
	return s
}

func (s *ListLniPrivateIpAddressResponseBody) SetMessage(v string) *ListLniPrivateIpAddressResponseBody {
	s.Message = &v
	return s
}

func (s *ListLniPrivateIpAddressResponseBody) SetRequestId(v string) *ListLniPrivateIpAddressResponseBody {
	s.RequestId = &v
	return s
}

type ListLniPrivateIpAddressResponseBodyContent struct {
	Data  []*ListLniPrivateIpAddressResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Total *int64                                            `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListLniPrivateIpAddressResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s ListLniPrivateIpAddressResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListLniPrivateIpAddressResponseBodyContent) SetData(v []*ListLniPrivateIpAddressResponseBodyContentData) *ListLniPrivateIpAddressResponseBodyContent {
	s.Data = v
	return s
}

func (s *ListLniPrivateIpAddressResponseBodyContent) SetTotal(v int64) *ListLniPrivateIpAddressResponseBodyContent {
	s.Total = &v
	return s
}

type ListLniPrivateIpAddressResponseBodyContentData struct {
	GmtCreate          *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	IpAddressMac       *string `json:"IpAddressMac,omitempty" xml:"IpAddressMac,omitempty"`
	IpName             *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
	Message            *string `json:"Message,omitempty" xml:"Message,omitempty"`
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	PrivateIpAddress   *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListLniPrivateIpAddressResponseBodyContentData) String() string {
	return tea.Prettify(s)
}

func (s ListLniPrivateIpAddressResponseBodyContentData) GoString() string {
	return s.String()
}

func (s *ListLniPrivateIpAddressResponseBodyContentData) SetGmtCreate(v string) *ListLniPrivateIpAddressResponseBodyContentData {
	s.GmtCreate = &v
	return s
}

func (s *ListLniPrivateIpAddressResponseBodyContentData) SetIpAddressMac(v string) *ListLniPrivateIpAddressResponseBodyContentData {
	s.IpAddressMac = &v
	return s
}

func (s *ListLniPrivateIpAddressResponseBodyContentData) SetIpName(v string) *ListLniPrivateIpAddressResponseBodyContentData {
	s.IpName = &v
	return s
}

func (s *ListLniPrivateIpAddressResponseBodyContentData) SetMessage(v string) *ListLniPrivateIpAddressResponseBodyContentData {
	s.Message = &v
	return s
}

func (s *ListLniPrivateIpAddressResponseBodyContentData) SetNetworkInterfaceId(v string) *ListLniPrivateIpAddressResponseBodyContentData {
	s.NetworkInterfaceId = &v
	return s
}

func (s *ListLniPrivateIpAddressResponseBodyContentData) SetPrivateIpAddress(v string) *ListLniPrivateIpAddressResponseBodyContentData {
	s.PrivateIpAddress = &v
	return s
}

func (s *ListLniPrivateIpAddressResponseBodyContentData) SetRegionId(v string) *ListLniPrivateIpAddressResponseBodyContentData {
	s.RegionId = &v
	return s
}

func (s *ListLniPrivateIpAddressResponseBodyContentData) SetStatus(v string) *ListLniPrivateIpAddressResponseBodyContentData {
	s.Status = &v
	return s
}

type ListLniPrivateIpAddressResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListLniPrivateIpAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListLniPrivateIpAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLniPrivateIpAddressResponse) GoString() string {
	return s.String()
}

func (s *ListLniPrivateIpAddressResponse) SetHeaders(v map[string]*string) *ListLniPrivateIpAddressResponse {
	s.Headers = v
	return s
}

func (s *ListLniPrivateIpAddressResponse) SetStatusCode(v int32) *ListLniPrivateIpAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLniPrivateIpAddressResponse) SetBody(v *ListLniPrivateIpAddressResponseBody) *ListLniPrivateIpAddressResponse {
	s.Body = v
	return s
}

type ListNetworkInterfacesRequest struct {
	EnablePage         *bool   `json:"EnablePage,omitempty" xml:"EnablePage,omitempty"`
	Ip                 *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	NodeId             *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	PageNumber         *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SubnetId           *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	VpdId              *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
}

func (s ListNetworkInterfacesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNetworkInterfacesRequest) GoString() string {
	return s.String()
}

func (s *ListNetworkInterfacesRequest) SetEnablePage(v bool) *ListNetworkInterfacesRequest {
	s.EnablePage = &v
	return s
}

func (s *ListNetworkInterfacesRequest) SetIp(v string) *ListNetworkInterfacesRequest {
	s.Ip = &v
	return s
}

func (s *ListNetworkInterfacesRequest) SetNetworkInterfaceId(v string) *ListNetworkInterfacesRequest {
	s.NetworkInterfaceId = &v
	return s
}

func (s *ListNetworkInterfacesRequest) SetNodeId(v string) *ListNetworkInterfacesRequest {
	s.NodeId = &v
	return s
}

func (s *ListNetworkInterfacesRequest) SetPageNumber(v int32) *ListNetworkInterfacesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNetworkInterfacesRequest) SetPageSize(v int32) *ListNetworkInterfacesRequest {
	s.PageSize = &v
	return s
}

func (s *ListNetworkInterfacesRequest) SetRegionId(v string) *ListNetworkInterfacesRequest {
	s.RegionId = &v
	return s
}

func (s *ListNetworkInterfacesRequest) SetSubnetId(v string) *ListNetworkInterfacesRequest {
	s.SubnetId = &v
	return s
}

func (s *ListNetworkInterfacesRequest) SetVpdId(v string) *ListNetworkInterfacesRequest {
	s.VpdId = &v
	return s
}

type ListNetworkInterfacesResponseBody struct {
	Code      *int32                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   *ListNetworkInterfacesResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListNetworkInterfacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNetworkInterfacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNetworkInterfacesResponseBody) SetCode(v int32) *ListNetworkInterfacesResponseBody {
	s.Code = &v
	return s
}

func (s *ListNetworkInterfacesResponseBody) SetContent(v *ListNetworkInterfacesResponseBodyContent) *ListNetworkInterfacesResponseBody {
	s.Content = v
	return s
}

func (s *ListNetworkInterfacesResponseBody) SetMessage(v string) *ListNetworkInterfacesResponseBody {
	s.Message = &v
	return s
}

func (s *ListNetworkInterfacesResponseBody) SetRequestId(v string) *ListNetworkInterfacesResponseBody {
	s.RequestId = &v
	return s
}

type ListNetworkInterfacesResponseBodyContent struct {
	Data  []*ListNetworkInterfacesResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Total *int64                                          `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListNetworkInterfacesResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s ListNetworkInterfacesResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListNetworkInterfacesResponseBodyContent) SetData(v []*ListNetworkInterfacesResponseBodyContentData) *ListNetworkInterfacesResponseBodyContent {
	s.Data = v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContent) SetTotal(v int64) *ListNetworkInterfacesResponseBodyContent {
	s.Total = &v
	return s
}

type ListNetworkInterfacesResponseBodyContentData struct {
	CreateTime               *string                                                                 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Ethernet                 []*string                                                               `json:"Ethernet,omitempty" xml:"Ethernet,omitempty" type:"Repeated"`
	Gateway                  *string                                                                 `json:"Gateway,omitempty" xml:"Gateway,omitempty"`
	Ip                       *string                                                                 `json:"Ip,omitempty" xml:"Ip,omitempty"`
	NcType                   *string                                                                 `json:"NcType,omitempty" xml:"NcType,omitempty"`
	NetworkInterfaceId       *string                                                                 `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	NetworkInterfaceName     *string                                                                 `json:"NetworkInterfaceName,omitempty" xml:"NetworkInterfaceName,omitempty"`
	NodeId                   *string                                                                 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	PrivateIpAddressMacGroup []*ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup `json:"PrivateIpAddressMacGroup,omitempty" xml:"PrivateIpAddressMacGroup,omitempty" type:"Repeated"`
	Quota                    *int32                                                                  `json:"Quota,omitempty" xml:"Quota,omitempty"`
	RegionId                 *string                                                                 `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceMac               *string                                                                 `json:"ServiceMac,omitempty" xml:"ServiceMac,omitempty"`
	Status                   *string                                                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	SubnetBaseInfo           *ListNetworkInterfacesResponseBodyContentDataSubnetBaseInfo             `json:"SubnetBaseInfo,omitempty" xml:"SubnetBaseInfo,omitempty" type:"Struct"`
	TenantId                 *string                                                                 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	VpdBaseInfo              *ListNetworkInterfacesResponseBodyContentDataVpdBaseInfo                `json:"VpdBaseInfo,omitempty" xml:"VpdBaseInfo,omitempty" type:"Struct"`
	ZoneId                   *string                                                                 `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListNetworkInterfacesResponseBodyContentData) String() string {
	return tea.Prettify(s)
}

func (s ListNetworkInterfacesResponseBodyContentData) GoString() string {
	return s.String()
}

func (s *ListNetworkInterfacesResponseBodyContentData) SetCreateTime(v string) *ListNetworkInterfacesResponseBodyContentData {
	s.CreateTime = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentData) SetEthernet(v []*string) *ListNetworkInterfacesResponseBodyContentData {
	s.Ethernet = v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentData) SetGateway(v string) *ListNetworkInterfacesResponseBodyContentData {
	s.Gateway = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentData) SetIp(v string) *ListNetworkInterfacesResponseBodyContentData {
	s.Ip = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentData) SetNcType(v string) *ListNetworkInterfacesResponseBodyContentData {
	s.NcType = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentData) SetNetworkInterfaceId(v string) *ListNetworkInterfacesResponseBodyContentData {
	s.NetworkInterfaceId = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentData) SetNetworkInterfaceName(v string) *ListNetworkInterfacesResponseBodyContentData {
	s.NetworkInterfaceName = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentData) SetNodeId(v string) *ListNetworkInterfacesResponseBodyContentData {
	s.NodeId = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentData) SetPrivateIpAddressMacGroup(v []*ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup) *ListNetworkInterfacesResponseBodyContentData {
	s.PrivateIpAddressMacGroup = v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentData) SetQuota(v int32) *ListNetworkInterfacesResponseBodyContentData {
	s.Quota = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentData) SetRegionId(v string) *ListNetworkInterfacesResponseBodyContentData {
	s.RegionId = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentData) SetServiceMac(v string) *ListNetworkInterfacesResponseBodyContentData {
	s.ServiceMac = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentData) SetStatus(v string) *ListNetworkInterfacesResponseBodyContentData {
	s.Status = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentData) SetSubnetBaseInfo(v *ListNetworkInterfacesResponseBodyContentDataSubnetBaseInfo) *ListNetworkInterfacesResponseBodyContentData {
	s.SubnetBaseInfo = v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentData) SetTenantId(v string) *ListNetworkInterfacesResponseBodyContentData {
	s.TenantId = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentData) SetVpdBaseInfo(v *ListNetworkInterfacesResponseBodyContentDataVpdBaseInfo) *ListNetworkInterfacesResponseBodyContentData {
	s.VpdBaseInfo = v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentData) SetZoneId(v string) *ListNetworkInterfacesResponseBodyContentData {
	s.ZoneId = &v
	return s
}

type ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup struct {
	IpAddressMac     *string `json:"IpAddressMac,omitempty" xml:"IpAddressMac,omitempty"`
	IpName           *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
	Message          *string `json:"Message,omitempty" xml:"Message,omitempty"`
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup) String() string {
	return tea.Prettify(s)
}

func (s ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup) GoString() string {
	return s.String()
}

func (s *ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup) SetIpAddressMac(v string) *ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup {
	s.IpAddressMac = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup) SetIpName(v string) *ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup {
	s.IpName = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup) SetMessage(v string) *ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup {
	s.Message = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup) SetPrivateIpAddress(v string) *ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup {
	s.PrivateIpAddress = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup) SetStatus(v string) *ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup {
	s.Status = &v
	return s
}

type ListNetworkInterfacesResponseBodyContentDataSubnetBaseInfo struct {
	Cidr       *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	SubnetId   *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	SubnetName *string `json:"SubnetName,omitempty" xml:"SubnetName,omitempty"`
}

func (s ListNetworkInterfacesResponseBodyContentDataSubnetBaseInfo) String() string {
	return tea.Prettify(s)
}

func (s ListNetworkInterfacesResponseBodyContentDataSubnetBaseInfo) GoString() string {
	return s.String()
}

func (s *ListNetworkInterfacesResponseBodyContentDataSubnetBaseInfo) SetCidr(v string) *ListNetworkInterfacesResponseBodyContentDataSubnetBaseInfo {
	s.Cidr = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentDataSubnetBaseInfo) SetCreateTime(v string) *ListNetworkInterfacesResponseBodyContentDataSubnetBaseInfo {
	s.CreateTime = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentDataSubnetBaseInfo) SetSubnetId(v string) *ListNetworkInterfacesResponseBodyContentDataSubnetBaseInfo {
	s.SubnetId = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentDataSubnetBaseInfo) SetSubnetName(v string) *ListNetworkInterfacesResponseBodyContentDataSubnetBaseInfo {
	s.SubnetName = &v
	return s
}

type ListNetworkInterfacesResponseBodyContentDataVpdBaseInfo struct {
	Cidr       *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	VpdId      *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	VpdName    *string `json:"VpdName,omitempty" xml:"VpdName,omitempty"`
}

func (s ListNetworkInterfacesResponseBodyContentDataVpdBaseInfo) String() string {
	return tea.Prettify(s)
}

func (s ListNetworkInterfacesResponseBodyContentDataVpdBaseInfo) GoString() string {
	return s.String()
}

func (s *ListNetworkInterfacesResponseBodyContentDataVpdBaseInfo) SetCidr(v string) *ListNetworkInterfacesResponseBodyContentDataVpdBaseInfo {
	s.Cidr = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentDataVpdBaseInfo) SetCreateTime(v string) *ListNetworkInterfacesResponseBodyContentDataVpdBaseInfo {
	s.CreateTime = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentDataVpdBaseInfo) SetVpdId(v string) *ListNetworkInterfacesResponseBodyContentDataVpdBaseInfo {
	s.VpdId = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentDataVpdBaseInfo) SetVpdName(v string) *ListNetworkInterfacesResponseBodyContentDataVpdBaseInfo {
	s.VpdName = &v
	return s
}

type ListNetworkInterfacesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListNetworkInterfacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNetworkInterfacesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNetworkInterfacesResponse) GoString() string {
	return s.String()
}

func (s *ListNetworkInterfacesResponse) SetHeaders(v map[string]*string) *ListNetworkInterfacesResponse {
	s.Headers = v
	return s
}

func (s *ListNetworkInterfacesResponse) SetStatusCode(v int32) *ListNetworkInterfacesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNetworkInterfacesResponse) SetBody(v *ListNetworkInterfacesResponseBody) *ListNetworkInterfacesResponse {
	s.Body = v
	return s
}

type ListSubnetsRequest struct {
	EnablePage      *bool                    `json:"EnablePage,omitempty" xml:"EnablePage,omitempty"`
	PageNumber      *int32                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId        *string                  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string                  `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Status          *string                  `json:"Status,omitempty" xml:"Status,omitempty"`
	SubnetId        *string                  `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	SubnetName      *string                  `json:"SubnetName,omitempty" xml:"SubnetName,omitempty"`
	Tag             []*ListSubnetsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	Type            *string                  `json:"Type,omitempty" xml:"Type,omitempty"`
	VpdId           *string                  `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	ZoneId          *string                  `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListSubnetsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSubnetsRequest) GoString() string {
	return s.String()
}

func (s *ListSubnetsRequest) SetEnablePage(v bool) *ListSubnetsRequest {
	s.EnablePage = &v
	return s
}

func (s *ListSubnetsRequest) SetPageNumber(v int32) *ListSubnetsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSubnetsRequest) SetPageSize(v int32) *ListSubnetsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSubnetsRequest) SetRegionId(v string) *ListSubnetsRequest {
	s.RegionId = &v
	return s
}

func (s *ListSubnetsRequest) SetResourceGroupId(v string) *ListSubnetsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListSubnetsRequest) SetStatus(v string) *ListSubnetsRequest {
	s.Status = &v
	return s
}

func (s *ListSubnetsRequest) SetSubnetId(v string) *ListSubnetsRequest {
	s.SubnetId = &v
	return s
}

func (s *ListSubnetsRequest) SetSubnetName(v string) *ListSubnetsRequest {
	s.SubnetName = &v
	return s
}

func (s *ListSubnetsRequest) SetTag(v []*ListSubnetsRequestTag) *ListSubnetsRequest {
	s.Tag = v
	return s
}

func (s *ListSubnetsRequest) SetType(v string) *ListSubnetsRequest {
	s.Type = &v
	return s
}

func (s *ListSubnetsRequest) SetVpdId(v string) *ListSubnetsRequest {
	s.VpdId = &v
	return s
}

func (s *ListSubnetsRequest) SetZoneId(v string) *ListSubnetsRequest {
	s.ZoneId = &v
	return s
}

type ListSubnetsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListSubnetsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListSubnetsRequestTag) GoString() string {
	return s.String()
}

func (s *ListSubnetsRequestTag) SetKey(v string) *ListSubnetsRequestTag {
	s.Key = &v
	return s
}

func (s *ListSubnetsRequestTag) SetValue(v string) *ListSubnetsRequestTag {
	s.Value = &v
	return s
}

type ListSubnetsResponseBody struct {
	Code      *int32                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   *ListSubnetsResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSubnetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSubnetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubnetsResponseBody) SetCode(v int32) *ListSubnetsResponseBody {
	s.Code = &v
	return s
}

func (s *ListSubnetsResponseBody) SetContent(v *ListSubnetsResponseBodyContent) *ListSubnetsResponseBody {
	s.Content = v
	return s
}

func (s *ListSubnetsResponseBody) SetMessage(v string) *ListSubnetsResponseBody {
	s.Message = &v
	return s
}

func (s *ListSubnetsResponseBody) SetRequestId(v string) *ListSubnetsResponseBody {
	s.RequestId = &v
	return s
}

type ListSubnetsResponseBodyContent struct {
	Data  []*ListSubnetsResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Total *int64                                `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSubnetsResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s ListSubnetsResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListSubnetsResponseBodyContent) SetData(v []*ListSubnetsResponseBodyContentData) *ListSubnetsResponseBodyContent {
	s.Data = v
	return s
}

func (s *ListSubnetsResponseBodyContent) SetTotal(v int64) *ListSubnetsResponseBodyContent {
	s.Total = &v
	return s
}

type ListSubnetsResponseBodyContentData struct {
	Cidr            *string                                        `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	CreateTime      *string                                        `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	GmtModified     *string                                        `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Message         *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	NcCount         *int64                                         `json:"NcCount,omitempty" xml:"NcCount,omitempty"`
	RegionId        *string                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string                                        `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Status          *string                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	SubnetId        *string                                        `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	SubnetName      *string                                        `json:"SubnetName,omitempty" xml:"SubnetName,omitempty"`
	Tags            []*ListSubnetsResponseBodyContentDataTags      `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TenantId        *string                                        `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Type            *string                                        `json:"Type,omitempty" xml:"Type,omitempty"`
	VpdBaseInfo     *ListSubnetsResponseBodyContentDataVpdBaseInfo `json:"VpdBaseInfo,omitempty" xml:"VpdBaseInfo,omitempty" type:"Struct"`
	VpdId           *string                                        `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	ZoneId          *string                                        `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListSubnetsResponseBodyContentData) String() string {
	return tea.Prettify(s)
}

func (s ListSubnetsResponseBodyContentData) GoString() string {
	return s.String()
}

func (s *ListSubnetsResponseBodyContentData) SetCidr(v string) *ListSubnetsResponseBodyContentData {
	s.Cidr = &v
	return s
}

func (s *ListSubnetsResponseBodyContentData) SetCreateTime(v string) *ListSubnetsResponseBodyContentData {
	s.CreateTime = &v
	return s
}

func (s *ListSubnetsResponseBodyContentData) SetGmtModified(v string) *ListSubnetsResponseBodyContentData {
	s.GmtModified = &v
	return s
}

func (s *ListSubnetsResponseBodyContentData) SetMessage(v string) *ListSubnetsResponseBodyContentData {
	s.Message = &v
	return s
}

func (s *ListSubnetsResponseBodyContentData) SetNcCount(v int64) *ListSubnetsResponseBodyContentData {
	s.NcCount = &v
	return s
}

func (s *ListSubnetsResponseBodyContentData) SetRegionId(v string) *ListSubnetsResponseBodyContentData {
	s.RegionId = &v
	return s
}

func (s *ListSubnetsResponseBodyContentData) SetResourceGroupId(v string) *ListSubnetsResponseBodyContentData {
	s.ResourceGroupId = &v
	return s
}

func (s *ListSubnetsResponseBodyContentData) SetStatus(v string) *ListSubnetsResponseBodyContentData {
	s.Status = &v
	return s
}

func (s *ListSubnetsResponseBodyContentData) SetSubnetId(v string) *ListSubnetsResponseBodyContentData {
	s.SubnetId = &v
	return s
}

func (s *ListSubnetsResponseBodyContentData) SetSubnetName(v string) *ListSubnetsResponseBodyContentData {
	s.SubnetName = &v
	return s
}

func (s *ListSubnetsResponseBodyContentData) SetTags(v []*ListSubnetsResponseBodyContentDataTags) *ListSubnetsResponseBodyContentData {
	s.Tags = v
	return s
}

func (s *ListSubnetsResponseBodyContentData) SetTenantId(v string) *ListSubnetsResponseBodyContentData {
	s.TenantId = &v
	return s
}

func (s *ListSubnetsResponseBodyContentData) SetType(v string) *ListSubnetsResponseBodyContentData {
	s.Type = &v
	return s
}

func (s *ListSubnetsResponseBodyContentData) SetVpdBaseInfo(v *ListSubnetsResponseBodyContentDataVpdBaseInfo) *ListSubnetsResponseBodyContentData {
	s.VpdBaseInfo = v
	return s
}

func (s *ListSubnetsResponseBodyContentData) SetVpdId(v string) *ListSubnetsResponseBodyContentData {
	s.VpdId = &v
	return s
}

func (s *ListSubnetsResponseBodyContentData) SetZoneId(v string) *ListSubnetsResponseBodyContentData {
	s.ZoneId = &v
	return s
}

type ListSubnetsResponseBodyContentDataTags struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListSubnetsResponseBodyContentDataTags) String() string {
	return tea.Prettify(s)
}

func (s ListSubnetsResponseBodyContentDataTags) GoString() string {
	return s.String()
}

func (s *ListSubnetsResponseBodyContentDataTags) SetTagKey(v string) *ListSubnetsResponseBodyContentDataTags {
	s.TagKey = &v
	return s
}

func (s *ListSubnetsResponseBodyContentDataTags) SetTagValue(v string) *ListSubnetsResponseBodyContentDataTags {
	s.TagValue = &v
	return s
}

type ListSubnetsResponseBodyContentDataVpdBaseInfo struct {
	Cidr       *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	VpdId      *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	VpdName    *string `json:"VpdName,omitempty" xml:"VpdName,omitempty"`
}

func (s ListSubnetsResponseBodyContentDataVpdBaseInfo) String() string {
	return tea.Prettify(s)
}

func (s ListSubnetsResponseBodyContentDataVpdBaseInfo) GoString() string {
	return s.String()
}

func (s *ListSubnetsResponseBodyContentDataVpdBaseInfo) SetCidr(v string) *ListSubnetsResponseBodyContentDataVpdBaseInfo {
	s.Cidr = &v
	return s
}

func (s *ListSubnetsResponseBodyContentDataVpdBaseInfo) SetCreateTime(v string) *ListSubnetsResponseBodyContentDataVpdBaseInfo {
	s.CreateTime = &v
	return s
}

func (s *ListSubnetsResponseBodyContentDataVpdBaseInfo) SetVpdId(v string) *ListSubnetsResponseBodyContentDataVpdBaseInfo {
	s.VpdId = &v
	return s
}

func (s *ListSubnetsResponseBodyContentDataVpdBaseInfo) SetVpdName(v string) *ListSubnetsResponseBodyContentDataVpdBaseInfo {
	s.VpdName = &v
	return s
}

type ListSubnetsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSubnetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSubnetsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSubnetsResponse) GoString() string {
	return s.String()
}

func (s *ListSubnetsResponse) SetHeaders(v map[string]*string) *ListSubnetsResponse {
	s.Headers = v
	return s
}

func (s *ListSubnetsResponse) SetStatusCode(v int32) *ListSubnetsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSubnetsResponse) SetBody(v *ListSubnetsResponseBody) *ListSubnetsResponse {
	s.Body = v
	return s
}

type ListVccGrantRulesRequest struct {
	EnablePage    *bool   `json:"EnablePage,omitempty" xml:"EnablePage,omitempty"`
	ErId          *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	ForSelect     *bool   `json:"ForSelect,omitempty" xml:"ForSelect,omitempty"`
	GrantRuleId   *string `json:"GrantRuleId,omitempty" xml:"GrantRuleId,omitempty"`
	GrantTenantId *string `json:"GrantTenantId,omitempty" xml:"GrantTenantId,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName  *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListVccGrantRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVccGrantRulesRequest) GoString() string {
	return s.String()
}

func (s *ListVccGrantRulesRequest) SetEnablePage(v bool) *ListVccGrantRulesRequest {
	s.EnablePage = &v
	return s
}

func (s *ListVccGrantRulesRequest) SetErId(v string) *ListVccGrantRulesRequest {
	s.ErId = &v
	return s
}

func (s *ListVccGrantRulesRequest) SetForSelect(v bool) *ListVccGrantRulesRequest {
	s.ForSelect = &v
	return s
}

func (s *ListVccGrantRulesRequest) SetGrantRuleId(v string) *ListVccGrantRulesRequest {
	s.GrantRuleId = &v
	return s
}

func (s *ListVccGrantRulesRequest) SetGrantTenantId(v string) *ListVccGrantRulesRequest {
	s.GrantTenantId = &v
	return s
}

func (s *ListVccGrantRulesRequest) SetInstanceId(v string) *ListVccGrantRulesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListVccGrantRulesRequest) SetInstanceName(v string) *ListVccGrantRulesRequest {
	s.InstanceName = &v
	return s
}

func (s *ListVccGrantRulesRequest) SetPageNumber(v int32) *ListVccGrantRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListVccGrantRulesRequest) SetPageSize(v int32) *ListVccGrantRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListVccGrantRulesRequest) SetRegionId(v string) *ListVccGrantRulesRequest {
	s.RegionId = &v
	return s
}

type ListVccGrantRulesResponseBody struct {
	Code      *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   *ListVccGrantRulesResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListVccGrantRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVccGrantRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVccGrantRulesResponseBody) SetCode(v int32) *ListVccGrantRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListVccGrantRulesResponseBody) SetContent(v *ListVccGrantRulesResponseBodyContent) *ListVccGrantRulesResponseBody {
	s.Content = v
	return s
}

func (s *ListVccGrantRulesResponseBody) SetMessage(v string) *ListVccGrantRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListVccGrantRulesResponseBody) SetRequestId(v string) *ListVccGrantRulesResponseBody {
	s.RequestId = &v
	return s
}

type ListVccGrantRulesResponseBodyContent struct {
	Data  []*ListVccGrantRulesResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Total *int64                                      `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListVccGrantRulesResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s ListVccGrantRulesResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListVccGrantRulesResponseBodyContent) SetData(v []*ListVccGrantRulesResponseBodyContentData) *ListVccGrantRulesResponseBodyContent {
	s.Data = v
	return s
}

func (s *ListVccGrantRulesResponseBodyContent) SetTotal(v int64) *ListVccGrantRulesResponseBodyContent {
	s.Total = &v
	return s
}

type ListVccGrantRulesResponseBodyContentData struct {
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ErId          *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	GrantRuleId   *string `json:"GrantRuleId,omitempty" xml:"GrantRuleId,omitempty"`
	GrantTenantId *string `json:"GrantTenantId,omitempty" xml:"GrantTenantId,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName  *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Product       *string `json:"Product,omitempty" xml:"Product,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TenantId      *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Used          *bool   `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s ListVccGrantRulesResponseBodyContentData) String() string {
	return tea.Prettify(s)
}

func (s ListVccGrantRulesResponseBodyContentData) GoString() string {
	return s.String()
}

func (s *ListVccGrantRulesResponseBodyContentData) SetCreateTime(v string) *ListVccGrantRulesResponseBodyContentData {
	s.CreateTime = &v
	return s
}

func (s *ListVccGrantRulesResponseBodyContentData) SetErId(v string) *ListVccGrantRulesResponseBodyContentData {
	s.ErId = &v
	return s
}

func (s *ListVccGrantRulesResponseBodyContentData) SetGrantRuleId(v string) *ListVccGrantRulesResponseBodyContentData {
	s.GrantRuleId = &v
	return s
}

func (s *ListVccGrantRulesResponseBodyContentData) SetGrantTenantId(v string) *ListVccGrantRulesResponseBodyContentData {
	s.GrantTenantId = &v
	return s
}

func (s *ListVccGrantRulesResponseBodyContentData) SetInstanceId(v string) *ListVccGrantRulesResponseBodyContentData {
	s.InstanceId = &v
	return s
}

func (s *ListVccGrantRulesResponseBodyContentData) SetInstanceName(v string) *ListVccGrantRulesResponseBodyContentData {
	s.InstanceName = &v
	return s
}

func (s *ListVccGrantRulesResponseBodyContentData) SetProduct(v string) *ListVccGrantRulesResponseBodyContentData {
	s.Product = &v
	return s
}

func (s *ListVccGrantRulesResponseBodyContentData) SetRegionId(v string) *ListVccGrantRulesResponseBodyContentData {
	s.RegionId = &v
	return s
}

func (s *ListVccGrantRulesResponseBodyContentData) SetTenantId(v string) *ListVccGrantRulesResponseBodyContentData {
	s.TenantId = &v
	return s
}

func (s *ListVccGrantRulesResponseBodyContentData) SetUsed(v bool) *ListVccGrantRulesResponseBodyContentData {
	s.Used = &v
	return s
}

type ListVccGrantRulesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListVccGrantRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVccGrantRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVccGrantRulesResponse) GoString() string {
	return s.String()
}

func (s *ListVccGrantRulesResponse) SetHeaders(v map[string]*string) *ListVccGrantRulesResponse {
	s.Headers = v
	return s
}

func (s *ListVccGrantRulesResponse) SetStatusCode(v int32) *ListVccGrantRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVccGrantRulesResponse) SetBody(v *ListVccGrantRulesResponseBody) *ListVccGrantRulesResponse {
	s.Body = v
	return s
}

type ListVccRouteEntriesRequest struct {
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	EnablePage           *bool   `json:"EnablePage,omitempty" xml:"EnablePage,omitempty"`
	NextHopId            *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	NextHopType          *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RouteType            *string `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VccId                *string `json:"VccId,omitempty" xml:"VccId,omitempty"`
	VpdRouteEntryId      *string `json:"VpdRouteEntryId,omitempty" xml:"VpdRouteEntryId,omitempty"`
}

func (s ListVccRouteEntriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVccRouteEntriesRequest) GoString() string {
	return s.String()
}

func (s *ListVccRouteEntriesRequest) SetDestinationCidrBlock(v string) *ListVccRouteEntriesRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *ListVccRouteEntriesRequest) SetEnablePage(v bool) *ListVccRouteEntriesRequest {
	s.EnablePage = &v
	return s
}

func (s *ListVccRouteEntriesRequest) SetNextHopId(v string) *ListVccRouteEntriesRequest {
	s.NextHopId = &v
	return s
}

func (s *ListVccRouteEntriesRequest) SetNextHopType(v string) *ListVccRouteEntriesRequest {
	s.NextHopType = &v
	return s
}

func (s *ListVccRouteEntriesRequest) SetPageNumber(v int32) *ListVccRouteEntriesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListVccRouteEntriesRequest) SetPageSize(v int32) *ListVccRouteEntriesRequest {
	s.PageSize = &v
	return s
}

func (s *ListVccRouteEntriesRequest) SetRegionId(v string) *ListVccRouteEntriesRequest {
	s.RegionId = &v
	return s
}

func (s *ListVccRouteEntriesRequest) SetRouteType(v string) *ListVccRouteEntriesRequest {
	s.RouteType = &v
	return s
}

func (s *ListVccRouteEntriesRequest) SetStatus(v string) *ListVccRouteEntriesRequest {
	s.Status = &v
	return s
}

func (s *ListVccRouteEntriesRequest) SetVccId(v string) *ListVccRouteEntriesRequest {
	s.VccId = &v
	return s
}

func (s *ListVccRouteEntriesRequest) SetVpdRouteEntryId(v string) *ListVccRouteEntriesRequest {
	s.VpdRouteEntryId = &v
	return s
}

type ListVccRouteEntriesResponseBody struct {
	Code      *int32                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   *ListVccRouteEntriesResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListVccRouteEntriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVccRouteEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVccRouteEntriesResponseBody) SetCode(v int32) *ListVccRouteEntriesResponseBody {
	s.Code = &v
	return s
}

func (s *ListVccRouteEntriesResponseBody) SetContent(v *ListVccRouteEntriesResponseBodyContent) *ListVccRouteEntriesResponseBody {
	s.Content = v
	return s
}

func (s *ListVccRouteEntriesResponseBody) SetMessage(v string) *ListVccRouteEntriesResponseBody {
	s.Message = &v
	return s
}

func (s *ListVccRouteEntriesResponseBody) SetRequestId(v string) *ListVccRouteEntriesResponseBody {
	s.RequestId = &v
	return s
}

type ListVccRouteEntriesResponseBodyContent struct {
	Data  []*ListVccRouteEntriesResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Total *int64                                        `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListVccRouteEntriesResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s ListVccRouteEntriesResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListVccRouteEntriesResponseBodyContent) SetData(v []*ListVccRouteEntriesResponseBodyContentData) *ListVccRouteEntriesResponseBodyContent {
	s.Data = v
	return s
}

func (s *ListVccRouteEntriesResponseBodyContent) SetTotal(v int64) *ListVccRouteEntriesResponseBodyContent {
	s.Total = &v
	return s
}

type ListVccRouteEntriesResponseBodyContentData struct {
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	GmtModified          *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	NextHopId            *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	NextHopType          *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceTenantId     *string `json:"ResourceTenantId,omitempty" xml:"ResourceTenantId,omitempty"`
	RouteType            *string `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TenantId             *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	VccId                *string `json:"VccId,omitempty" xml:"VccId,omitempty"`
	VccRouteEntryId      *string `json:"VccRouteEntryId,omitempty" xml:"VccRouteEntryId,omitempty"`
}

func (s ListVccRouteEntriesResponseBodyContentData) String() string {
	return tea.Prettify(s)
}

func (s ListVccRouteEntriesResponseBodyContentData) GoString() string {
	return s.String()
}

func (s *ListVccRouteEntriesResponseBodyContentData) SetDestinationCidrBlock(v string) *ListVccRouteEntriesResponseBodyContentData {
	s.DestinationCidrBlock = &v
	return s
}

func (s *ListVccRouteEntriesResponseBodyContentData) SetGmtModified(v string) *ListVccRouteEntriesResponseBodyContentData {
	s.GmtModified = &v
	return s
}

func (s *ListVccRouteEntriesResponseBodyContentData) SetNextHopId(v string) *ListVccRouteEntriesResponseBodyContentData {
	s.NextHopId = &v
	return s
}

func (s *ListVccRouteEntriesResponseBodyContentData) SetNextHopType(v string) *ListVccRouteEntriesResponseBodyContentData {
	s.NextHopType = &v
	return s
}

func (s *ListVccRouteEntriesResponseBodyContentData) SetRegionId(v string) *ListVccRouteEntriesResponseBodyContentData {
	s.RegionId = &v
	return s
}

func (s *ListVccRouteEntriesResponseBodyContentData) SetResourceTenantId(v string) *ListVccRouteEntriesResponseBodyContentData {
	s.ResourceTenantId = &v
	return s
}

func (s *ListVccRouteEntriesResponseBodyContentData) SetRouteType(v string) *ListVccRouteEntriesResponseBodyContentData {
	s.RouteType = &v
	return s
}

func (s *ListVccRouteEntriesResponseBodyContentData) SetStatus(v string) *ListVccRouteEntriesResponseBodyContentData {
	s.Status = &v
	return s
}

func (s *ListVccRouteEntriesResponseBodyContentData) SetTenantId(v string) *ListVccRouteEntriesResponseBodyContentData {
	s.TenantId = &v
	return s
}

func (s *ListVccRouteEntriesResponseBodyContentData) SetVccId(v string) *ListVccRouteEntriesResponseBodyContentData {
	s.VccId = &v
	return s
}

func (s *ListVccRouteEntriesResponseBodyContentData) SetVccRouteEntryId(v string) *ListVccRouteEntriesResponseBodyContentData {
	s.VccRouteEntryId = &v
	return s
}

type ListVccRouteEntriesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListVccRouteEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVccRouteEntriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVccRouteEntriesResponse) GoString() string {
	return s.String()
}

func (s *ListVccRouteEntriesResponse) SetHeaders(v map[string]*string) *ListVccRouteEntriesResponse {
	s.Headers = v
	return s
}

func (s *ListVccRouteEntriesResponse) SetStatusCode(v int32) *ListVccRouteEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVccRouteEntriesResponse) SetBody(v *ListVccRouteEntriesResponseBody) *ListVccRouteEntriesResponse {
	s.Body = v
	return s
}

type ListVccsRequest struct {
	Bandwidth       *int32                `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	CenId           *string               `json:"CenId,omitempty" xml:"CenId,omitempty"`
	EnablePage      *bool                 `json:"EnablePage,omitempty" xml:"EnablePage,omitempty"`
	ExStatus        *string               `json:"ExStatus,omitempty" xml:"ExStatus,omitempty"`
	FilterErId      *string               `json:"FilterErId,omitempty" xml:"FilterErId,omitempty"`
	PageNumber      *int32                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId        *string               `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string               `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Status          *string               `json:"Status,omitempty" xml:"Status,omitempty"`
	Tag             []*ListVccsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	VccId           *string               `json:"VccId,omitempty" xml:"VccId,omitempty"`
	VpcId           *string               `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpdId           *string               `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
}

func (s ListVccsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVccsRequest) GoString() string {
	return s.String()
}

func (s *ListVccsRequest) SetBandwidth(v int32) *ListVccsRequest {
	s.Bandwidth = &v
	return s
}

func (s *ListVccsRequest) SetCenId(v string) *ListVccsRequest {
	s.CenId = &v
	return s
}

func (s *ListVccsRequest) SetEnablePage(v bool) *ListVccsRequest {
	s.EnablePage = &v
	return s
}

func (s *ListVccsRequest) SetExStatus(v string) *ListVccsRequest {
	s.ExStatus = &v
	return s
}

func (s *ListVccsRequest) SetFilterErId(v string) *ListVccsRequest {
	s.FilterErId = &v
	return s
}

func (s *ListVccsRequest) SetPageNumber(v int32) *ListVccsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListVccsRequest) SetPageSize(v int32) *ListVccsRequest {
	s.PageSize = &v
	return s
}

func (s *ListVccsRequest) SetRegionId(v string) *ListVccsRequest {
	s.RegionId = &v
	return s
}

func (s *ListVccsRequest) SetResourceGroupId(v string) *ListVccsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVccsRequest) SetStatus(v string) *ListVccsRequest {
	s.Status = &v
	return s
}

func (s *ListVccsRequest) SetTag(v []*ListVccsRequestTag) *ListVccsRequest {
	s.Tag = v
	return s
}

func (s *ListVccsRequest) SetVccId(v string) *ListVccsRequest {
	s.VccId = &v
	return s
}

func (s *ListVccsRequest) SetVpcId(v string) *ListVccsRequest {
	s.VpcId = &v
	return s
}

func (s *ListVccsRequest) SetVpdId(v string) *ListVccsRequest {
	s.VpdId = &v
	return s
}

type ListVccsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListVccsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListVccsRequestTag) GoString() string {
	return s.String()
}

func (s *ListVccsRequestTag) SetKey(v string) *ListVccsRequestTag {
	s.Key = &v
	return s
}

func (s *ListVccsRequestTag) SetValue(v string) *ListVccsRequestTag {
	s.Value = &v
	return s
}

type ListVccsResponseBody struct {
	Code      *int32                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   *ListVccsResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListVccsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVccsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVccsResponseBody) SetCode(v int32) *ListVccsResponseBody {
	s.Code = &v
	return s
}

func (s *ListVccsResponseBody) SetContent(v *ListVccsResponseBodyContent) *ListVccsResponseBody {
	s.Content = v
	return s
}

func (s *ListVccsResponseBody) SetMessage(v string) *ListVccsResponseBody {
	s.Message = &v
	return s
}

func (s *ListVccsResponseBody) SetRequestId(v string) *ListVccsResponseBody {
	s.RequestId = &v
	return s
}

type ListVccsResponseBodyContent struct {
	Data  []*ListVccsResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Total *int64                             `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListVccsResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s ListVccsResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListVccsResponseBodyContent) SetData(v []*ListVccsResponseBodyContentData) *ListVccsResponseBodyContent {
	s.Data = v
	return s
}

func (s *ListVccsResponseBodyContent) SetTotal(v int64) *ListVccsResponseBodyContent {
	s.Total = &v
	return s
}

type ListVccsResponseBodyContentData struct {
	AccessPointId   *string                                     `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	BandwidthStr    *string                                     `json:"BandwidthStr,omitempty" xml:"BandwidthStr,omitempty"`
	BgpCidr         *string                                     `json:"BgpCidr,omitempty" xml:"BgpCidr,omitempty"`
	CenId           *string                                     `json:"CenId,omitempty" xml:"CenId,omitempty"`
	CommodityCode   *string                                     `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	ConnectionType  *string                                     `json:"ConnectionType,omitempty" xml:"ConnectionType,omitempty"`
	CreateTime      *string                                     `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CurrentNode     *string                                     `json:"CurrentNode,omitempty" xml:"CurrentNode,omitempty"`
	ErInfos         []*ListVccsResponseBodyContentDataErInfos   `json:"ErInfos,omitempty" xml:"ErInfos,omitempty" type:"Repeated"`
	ExpirationDate  *string                                     `json:"ExpirationDate,omitempty" xml:"ExpirationDate,omitempty"`
	GmtModified     *string                                     `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	LineOperator    *string                                     `json:"LineOperator,omitempty" xml:"LineOperator,omitempty"`
	Message         *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	PortType        *string                                     `json:"PortType,omitempty" xml:"PortType,omitempty"`
	Rate            *float64                                    `json:"Rate,omitempty" xml:"Rate,omitempty"`
	RegionId        *string                                     `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string                                     `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Spec            *string                                     `json:"Spec,omitempty" xml:"Spec,omitempty"`
	Status          *string                                     `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags            []*ListVccsResponseBodyContentDataTags      `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TaskId          *string                                     `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TenantId        *string                                     `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	VccId           *string                                     `json:"VccId,omitempty" xml:"VccId,omitempty"`
	VccName         *string                                     `json:"VccName,omitempty" xml:"VccName,omitempty"`
	VpcId           *string                                     `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpdBaseInfo     *ListVccsResponseBodyContentDataVpdBaseInfo `json:"VpdBaseInfo,omitempty" xml:"VpdBaseInfo,omitempty" type:"Struct"`
	VpdId           *string                                     `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	ZoneId          *string                                     `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListVccsResponseBodyContentData) String() string {
	return tea.Prettify(s)
}

func (s ListVccsResponseBodyContentData) GoString() string {
	return s.String()
}

func (s *ListVccsResponseBodyContentData) SetAccessPointId(v string) *ListVccsResponseBodyContentData {
	s.AccessPointId = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetBandwidthStr(v string) *ListVccsResponseBodyContentData {
	s.BandwidthStr = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetBgpCidr(v string) *ListVccsResponseBodyContentData {
	s.BgpCidr = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetCenId(v string) *ListVccsResponseBodyContentData {
	s.CenId = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetCommodityCode(v string) *ListVccsResponseBodyContentData {
	s.CommodityCode = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetConnectionType(v string) *ListVccsResponseBodyContentData {
	s.ConnectionType = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetCreateTime(v string) *ListVccsResponseBodyContentData {
	s.CreateTime = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetCurrentNode(v string) *ListVccsResponseBodyContentData {
	s.CurrentNode = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetErInfos(v []*ListVccsResponseBodyContentDataErInfos) *ListVccsResponseBodyContentData {
	s.ErInfos = v
	return s
}

func (s *ListVccsResponseBodyContentData) SetExpirationDate(v string) *ListVccsResponseBodyContentData {
	s.ExpirationDate = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetGmtModified(v string) *ListVccsResponseBodyContentData {
	s.GmtModified = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetLineOperator(v string) *ListVccsResponseBodyContentData {
	s.LineOperator = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetMessage(v string) *ListVccsResponseBodyContentData {
	s.Message = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetPortType(v string) *ListVccsResponseBodyContentData {
	s.PortType = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetRate(v float64) *ListVccsResponseBodyContentData {
	s.Rate = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetRegionId(v string) *ListVccsResponseBodyContentData {
	s.RegionId = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetResourceGroupId(v string) *ListVccsResponseBodyContentData {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetSpec(v string) *ListVccsResponseBodyContentData {
	s.Spec = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetStatus(v string) *ListVccsResponseBodyContentData {
	s.Status = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetTags(v []*ListVccsResponseBodyContentDataTags) *ListVccsResponseBodyContentData {
	s.Tags = v
	return s
}

func (s *ListVccsResponseBodyContentData) SetTaskId(v string) *ListVccsResponseBodyContentData {
	s.TaskId = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetTenantId(v string) *ListVccsResponseBodyContentData {
	s.TenantId = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetVccId(v string) *ListVccsResponseBodyContentData {
	s.VccId = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetVccName(v string) *ListVccsResponseBodyContentData {
	s.VccName = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetVpcId(v string) *ListVccsResponseBodyContentData {
	s.VpcId = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetVpdBaseInfo(v *ListVccsResponseBodyContentDataVpdBaseInfo) *ListVccsResponseBodyContentData {
	s.VpdBaseInfo = v
	return s
}

func (s *ListVccsResponseBodyContentData) SetVpdId(v string) *ListVccsResponseBodyContentData {
	s.VpdId = &v
	return s
}

func (s *ListVccsResponseBodyContentData) SetZoneId(v string) *ListVccsResponseBodyContentData {
	s.ZoneId = &v
	return s
}

type ListVccsResponseBodyContentDataErInfos struct {
	Connections  *int64  `json:"Connections,omitempty" xml:"Connections,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ErId         *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	ErName       *string `json:"ErName,omitempty" xml:"ErName,omitempty"`
	GmtModified  *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	MasterZoneId *string `json:"MasterZoneId,omitempty" xml:"MasterZoneId,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RouteMaps    *int64  `json:"RouteMaps,omitempty" xml:"RouteMaps,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TenantId     *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListVccsResponseBodyContentDataErInfos) String() string {
	return tea.Prettify(s)
}

func (s ListVccsResponseBodyContentDataErInfos) GoString() string {
	return s.String()
}

func (s *ListVccsResponseBodyContentDataErInfos) SetConnections(v int64) *ListVccsResponseBodyContentDataErInfos {
	s.Connections = &v
	return s
}

func (s *ListVccsResponseBodyContentDataErInfos) SetCreateTime(v string) *ListVccsResponseBodyContentDataErInfos {
	s.CreateTime = &v
	return s
}

func (s *ListVccsResponseBodyContentDataErInfos) SetDescription(v string) *ListVccsResponseBodyContentDataErInfos {
	s.Description = &v
	return s
}

func (s *ListVccsResponseBodyContentDataErInfos) SetErId(v string) *ListVccsResponseBodyContentDataErInfos {
	s.ErId = &v
	return s
}

func (s *ListVccsResponseBodyContentDataErInfos) SetErName(v string) *ListVccsResponseBodyContentDataErInfos {
	s.ErName = &v
	return s
}

func (s *ListVccsResponseBodyContentDataErInfos) SetGmtModified(v string) *ListVccsResponseBodyContentDataErInfos {
	s.GmtModified = &v
	return s
}

func (s *ListVccsResponseBodyContentDataErInfos) SetMasterZoneId(v string) *ListVccsResponseBodyContentDataErInfos {
	s.MasterZoneId = &v
	return s
}

func (s *ListVccsResponseBodyContentDataErInfos) SetMessage(v string) *ListVccsResponseBodyContentDataErInfos {
	s.Message = &v
	return s
}

func (s *ListVccsResponseBodyContentDataErInfos) SetRegionId(v string) *ListVccsResponseBodyContentDataErInfos {
	s.RegionId = &v
	return s
}

func (s *ListVccsResponseBodyContentDataErInfos) SetRouteMaps(v int64) *ListVccsResponseBodyContentDataErInfos {
	s.RouteMaps = &v
	return s
}

func (s *ListVccsResponseBodyContentDataErInfos) SetStatus(v string) *ListVccsResponseBodyContentDataErInfos {
	s.Status = &v
	return s
}

func (s *ListVccsResponseBodyContentDataErInfos) SetTenantId(v string) *ListVccsResponseBodyContentDataErInfos {
	s.TenantId = &v
	return s
}

type ListVccsResponseBodyContentDataTags struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListVccsResponseBodyContentDataTags) String() string {
	return tea.Prettify(s)
}

func (s ListVccsResponseBodyContentDataTags) GoString() string {
	return s.String()
}

func (s *ListVccsResponseBodyContentDataTags) SetTagKey(v string) *ListVccsResponseBodyContentDataTags {
	s.TagKey = &v
	return s
}

func (s *ListVccsResponseBodyContentDataTags) SetTagValue(v string) *ListVccsResponseBodyContentDataTags {
	s.TagValue = &v
	return s
}

type ListVccsResponseBodyContentDataVpdBaseInfo struct {
	Cidr       *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	VpdId      *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	VpdName    *string `json:"VpdName,omitempty" xml:"VpdName,omitempty"`
}

func (s ListVccsResponseBodyContentDataVpdBaseInfo) String() string {
	return tea.Prettify(s)
}

func (s ListVccsResponseBodyContentDataVpdBaseInfo) GoString() string {
	return s.String()
}

func (s *ListVccsResponseBodyContentDataVpdBaseInfo) SetCidr(v string) *ListVccsResponseBodyContentDataVpdBaseInfo {
	s.Cidr = &v
	return s
}

func (s *ListVccsResponseBodyContentDataVpdBaseInfo) SetCreateTime(v string) *ListVccsResponseBodyContentDataVpdBaseInfo {
	s.CreateTime = &v
	return s
}

func (s *ListVccsResponseBodyContentDataVpdBaseInfo) SetVpdId(v string) *ListVccsResponseBodyContentDataVpdBaseInfo {
	s.VpdId = &v
	return s
}

func (s *ListVccsResponseBodyContentDataVpdBaseInfo) SetVpdName(v string) *ListVccsResponseBodyContentDataVpdBaseInfo {
	s.VpdName = &v
	return s
}

type ListVccsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListVccsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVccsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVccsResponse) GoString() string {
	return s.String()
}

func (s *ListVccsResponse) SetHeaders(v map[string]*string) *ListVccsResponse {
	s.Headers = v
	return s
}

func (s *ListVccsResponse) SetStatusCode(v int32) *ListVccsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVccsResponse) SetBody(v *ListVccsResponseBody) *ListVccsResponse {
	s.Body = v
	return s
}

type ListVpdGrantRulesRequest struct {
	EnablePage    *bool   `json:"EnablePage,omitempty" xml:"EnablePage,omitempty"`
	ErId          *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	ForSelect     *bool   `json:"ForSelect,omitempty" xml:"ForSelect,omitempty"`
	GrantRuleId   *string `json:"GrantRuleId,omitempty" xml:"GrantRuleId,omitempty"`
	GrantTenantId *string `json:"GrantTenantId,omitempty" xml:"GrantTenantId,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName  *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListVpdGrantRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVpdGrantRulesRequest) GoString() string {
	return s.String()
}

func (s *ListVpdGrantRulesRequest) SetEnablePage(v bool) *ListVpdGrantRulesRequest {
	s.EnablePage = &v
	return s
}

func (s *ListVpdGrantRulesRequest) SetErId(v string) *ListVpdGrantRulesRequest {
	s.ErId = &v
	return s
}

func (s *ListVpdGrantRulesRequest) SetForSelect(v bool) *ListVpdGrantRulesRequest {
	s.ForSelect = &v
	return s
}

func (s *ListVpdGrantRulesRequest) SetGrantRuleId(v string) *ListVpdGrantRulesRequest {
	s.GrantRuleId = &v
	return s
}

func (s *ListVpdGrantRulesRequest) SetGrantTenantId(v string) *ListVpdGrantRulesRequest {
	s.GrantTenantId = &v
	return s
}

func (s *ListVpdGrantRulesRequest) SetInstanceId(v string) *ListVpdGrantRulesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListVpdGrantRulesRequest) SetInstanceName(v string) *ListVpdGrantRulesRequest {
	s.InstanceName = &v
	return s
}

func (s *ListVpdGrantRulesRequest) SetPageNumber(v int32) *ListVpdGrantRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListVpdGrantRulesRequest) SetPageSize(v int32) *ListVpdGrantRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListVpdGrantRulesRequest) SetRegionId(v string) *ListVpdGrantRulesRequest {
	s.RegionId = &v
	return s
}

type ListVpdGrantRulesResponseBody struct {
	Code      *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   *ListVpdGrantRulesResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListVpdGrantRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVpdGrantRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpdGrantRulesResponseBody) SetCode(v int32) *ListVpdGrantRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListVpdGrantRulesResponseBody) SetContent(v *ListVpdGrantRulesResponseBodyContent) *ListVpdGrantRulesResponseBody {
	s.Content = v
	return s
}

func (s *ListVpdGrantRulesResponseBody) SetMessage(v string) *ListVpdGrantRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListVpdGrantRulesResponseBody) SetRequestId(v string) *ListVpdGrantRulesResponseBody {
	s.RequestId = &v
	return s
}

type ListVpdGrantRulesResponseBodyContent struct {
	Data  []*ListVpdGrantRulesResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Total *int64                                      `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListVpdGrantRulesResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s ListVpdGrantRulesResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListVpdGrantRulesResponseBodyContent) SetData(v []*ListVpdGrantRulesResponseBodyContentData) *ListVpdGrantRulesResponseBodyContent {
	s.Data = v
	return s
}

func (s *ListVpdGrantRulesResponseBodyContent) SetTotal(v int64) *ListVpdGrantRulesResponseBodyContent {
	s.Total = &v
	return s
}

type ListVpdGrantRulesResponseBodyContentData struct {
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ErId          *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	GrantRuleId   *string `json:"GrantRuleId,omitempty" xml:"GrantRuleId,omitempty"`
	GrantTenantId *string `json:"GrantTenantId,omitempty" xml:"GrantTenantId,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName  *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Product       *string `json:"Product,omitempty" xml:"Product,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TenantId      *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Used          *bool   `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s ListVpdGrantRulesResponseBodyContentData) String() string {
	return tea.Prettify(s)
}

func (s ListVpdGrantRulesResponseBodyContentData) GoString() string {
	return s.String()
}

func (s *ListVpdGrantRulesResponseBodyContentData) SetCreateTime(v string) *ListVpdGrantRulesResponseBodyContentData {
	s.CreateTime = &v
	return s
}

func (s *ListVpdGrantRulesResponseBodyContentData) SetErId(v string) *ListVpdGrantRulesResponseBodyContentData {
	s.ErId = &v
	return s
}

func (s *ListVpdGrantRulesResponseBodyContentData) SetGrantRuleId(v string) *ListVpdGrantRulesResponseBodyContentData {
	s.GrantRuleId = &v
	return s
}

func (s *ListVpdGrantRulesResponseBodyContentData) SetGrantTenantId(v string) *ListVpdGrantRulesResponseBodyContentData {
	s.GrantTenantId = &v
	return s
}

func (s *ListVpdGrantRulesResponseBodyContentData) SetInstanceId(v string) *ListVpdGrantRulesResponseBodyContentData {
	s.InstanceId = &v
	return s
}

func (s *ListVpdGrantRulesResponseBodyContentData) SetInstanceName(v string) *ListVpdGrantRulesResponseBodyContentData {
	s.InstanceName = &v
	return s
}

func (s *ListVpdGrantRulesResponseBodyContentData) SetProduct(v string) *ListVpdGrantRulesResponseBodyContentData {
	s.Product = &v
	return s
}

func (s *ListVpdGrantRulesResponseBodyContentData) SetRegionId(v string) *ListVpdGrantRulesResponseBodyContentData {
	s.RegionId = &v
	return s
}

func (s *ListVpdGrantRulesResponseBodyContentData) SetTenantId(v string) *ListVpdGrantRulesResponseBodyContentData {
	s.TenantId = &v
	return s
}

func (s *ListVpdGrantRulesResponseBodyContentData) SetUsed(v bool) *ListVpdGrantRulesResponseBodyContentData {
	s.Used = &v
	return s
}

type ListVpdGrantRulesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListVpdGrantRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVpdGrantRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVpdGrantRulesResponse) GoString() string {
	return s.String()
}

func (s *ListVpdGrantRulesResponse) SetHeaders(v map[string]*string) *ListVpdGrantRulesResponse {
	s.Headers = v
	return s
}

func (s *ListVpdGrantRulesResponse) SetStatusCode(v int32) *ListVpdGrantRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVpdGrantRulesResponse) SetBody(v *ListVpdGrantRulesResponseBody) *ListVpdGrantRulesResponse {
	s.Body = v
	return s
}

type ListVpdRouteEntriesRequest struct {
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	EnablePage           *bool   `json:"EnablePage,omitempty" xml:"EnablePage,omitempty"`
	NextHopId            *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	NextHopType          *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RouteType            *string `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VpdId                *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	VpdRouteEntryId      *string `json:"VpdRouteEntryId,omitempty" xml:"VpdRouteEntryId,omitempty"`
}

func (s ListVpdRouteEntriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVpdRouteEntriesRequest) GoString() string {
	return s.String()
}

func (s *ListVpdRouteEntriesRequest) SetDestinationCidrBlock(v string) *ListVpdRouteEntriesRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *ListVpdRouteEntriesRequest) SetEnablePage(v bool) *ListVpdRouteEntriesRequest {
	s.EnablePage = &v
	return s
}

func (s *ListVpdRouteEntriesRequest) SetNextHopId(v string) *ListVpdRouteEntriesRequest {
	s.NextHopId = &v
	return s
}

func (s *ListVpdRouteEntriesRequest) SetNextHopType(v string) *ListVpdRouteEntriesRequest {
	s.NextHopType = &v
	return s
}

func (s *ListVpdRouteEntriesRequest) SetPageNumber(v int32) *ListVpdRouteEntriesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListVpdRouteEntriesRequest) SetPageSize(v int32) *ListVpdRouteEntriesRequest {
	s.PageSize = &v
	return s
}

func (s *ListVpdRouteEntriesRequest) SetRegionId(v string) *ListVpdRouteEntriesRequest {
	s.RegionId = &v
	return s
}

func (s *ListVpdRouteEntriesRequest) SetRouteType(v string) *ListVpdRouteEntriesRequest {
	s.RouteType = &v
	return s
}

func (s *ListVpdRouteEntriesRequest) SetStatus(v string) *ListVpdRouteEntriesRequest {
	s.Status = &v
	return s
}

func (s *ListVpdRouteEntriesRequest) SetVpdId(v string) *ListVpdRouteEntriesRequest {
	s.VpdId = &v
	return s
}

func (s *ListVpdRouteEntriesRequest) SetVpdRouteEntryId(v string) *ListVpdRouteEntriesRequest {
	s.VpdRouteEntryId = &v
	return s
}

type ListVpdRouteEntriesResponseBody struct {
	Code      *int32                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   *ListVpdRouteEntriesResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListVpdRouteEntriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVpdRouteEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpdRouteEntriesResponseBody) SetCode(v int32) *ListVpdRouteEntriesResponseBody {
	s.Code = &v
	return s
}

func (s *ListVpdRouteEntriesResponseBody) SetContent(v *ListVpdRouteEntriesResponseBodyContent) *ListVpdRouteEntriesResponseBody {
	s.Content = v
	return s
}

func (s *ListVpdRouteEntriesResponseBody) SetMessage(v string) *ListVpdRouteEntriesResponseBody {
	s.Message = &v
	return s
}

func (s *ListVpdRouteEntriesResponseBody) SetRequestId(v string) *ListVpdRouteEntriesResponseBody {
	s.RequestId = &v
	return s
}

type ListVpdRouteEntriesResponseBodyContent struct {
	Data  []*ListVpdRouteEntriesResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Total *int64                                        `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListVpdRouteEntriesResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s ListVpdRouteEntriesResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListVpdRouteEntriesResponseBodyContent) SetData(v []*ListVpdRouteEntriesResponseBodyContentData) *ListVpdRouteEntriesResponseBodyContent {
	s.Data = v
	return s
}

func (s *ListVpdRouteEntriesResponseBodyContent) SetTotal(v int64) *ListVpdRouteEntriesResponseBodyContent {
	s.Total = &v
	return s
}

type ListVpdRouteEntriesResponseBodyContentData struct {
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	GmtModified          *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	NextHopId            *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	NextHopType          *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceTenantId     *string `json:"ResourceTenantId,omitempty" xml:"ResourceTenantId,omitempty"`
	RouteType            *string `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TenantId             *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	VpdId                *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	VpdRouteEntryId      *string `json:"VpdRouteEntryId,omitempty" xml:"VpdRouteEntryId,omitempty"`
}

func (s ListVpdRouteEntriesResponseBodyContentData) String() string {
	return tea.Prettify(s)
}

func (s ListVpdRouteEntriesResponseBodyContentData) GoString() string {
	return s.String()
}

func (s *ListVpdRouteEntriesResponseBodyContentData) SetDestinationCidrBlock(v string) *ListVpdRouteEntriesResponseBodyContentData {
	s.DestinationCidrBlock = &v
	return s
}

func (s *ListVpdRouteEntriesResponseBodyContentData) SetGmtModified(v string) *ListVpdRouteEntriesResponseBodyContentData {
	s.GmtModified = &v
	return s
}

func (s *ListVpdRouteEntriesResponseBodyContentData) SetNextHopId(v string) *ListVpdRouteEntriesResponseBodyContentData {
	s.NextHopId = &v
	return s
}

func (s *ListVpdRouteEntriesResponseBodyContentData) SetNextHopType(v string) *ListVpdRouteEntriesResponseBodyContentData {
	s.NextHopType = &v
	return s
}

func (s *ListVpdRouteEntriesResponseBodyContentData) SetRegionId(v string) *ListVpdRouteEntriesResponseBodyContentData {
	s.RegionId = &v
	return s
}

func (s *ListVpdRouteEntriesResponseBodyContentData) SetResourceTenantId(v string) *ListVpdRouteEntriesResponseBodyContentData {
	s.ResourceTenantId = &v
	return s
}

func (s *ListVpdRouteEntriesResponseBodyContentData) SetRouteType(v string) *ListVpdRouteEntriesResponseBodyContentData {
	s.RouteType = &v
	return s
}

func (s *ListVpdRouteEntriesResponseBodyContentData) SetStatus(v string) *ListVpdRouteEntriesResponseBodyContentData {
	s.Status = &v
	return s
}

func (s *ListVpdRouteEntriesResponseBodyContentData) SetTenantId(v string) *ListVpdRouteEntriesResponseBodyContentData {
	s.TenantId = &v
	return s
}

func (s *ListVpdRouteEntriesResponseBodyContentData) SetVpdId(v string) *ListVpdRouteEntriesResponseBodyContentData {
	s.VpdId = &v
	return s
}

func (s *ListVpdRouteEntriesResponseBodyContentData) SetVpdRouteEntryId(v string) *ListVpdRouteEntriesResponseBodyContentData {
	s.VpdRouteEntryId = &v
	return s
}

type ListVpdRouteEntriesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListVpdRouteEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVpdRouteEntriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVpdRouteEntriesResponse) GoString() string {
	return s.String()
}

func (s *ListVpdRouteEntriesResponse) SetHeaders(v map[string]*string) *ListVpdRouteEntriesResponse {
	s.Headers = v
	return s
}

func (s *ListVpdRouteEntriesResponse) SetStatusCode(v int32) *ListVpdRouteEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVpdRouteEntriesResponse) SetBody(v *ListVpdRouteEntriesResponseBody) *ListVpdRouteEntriesResponse {
	s.Body = v
	return s
}

type ListVpdsRequest struct {
	EnablePage      *bool                 `json:"EnablePage,omitempty" xml:"EnablePage,omitempty"`
	FilterErId      *string               `json:"FilterErId,omitempty" xml:"FilterErId,omitempty"`
	ForSelect       *bool                 `json:"ForSelect,omitempty" xml:"ForSelect,omitempty"`
	PageNumber      *int32                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId        *string               `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string               `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Status          *string               `json:"Status,omitempty" xml:"Status,omitempty"`
	Tag             []*ListVpdsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	VpdId           *string               `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	VpdName         *string               `json:"VpdName,omitempty" xml:"VpdName,omitempty"`
	WithDependence  *bool                 `json:"WithDependence,omitempty" xml:"WithDependence,omitempty"`
	WithoutVcc      *bool                 `json:"WithoutVcc,omitempty" xml:"WithoutVcc,omitempty"`
}

func (s ListVpdsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVpdsRequest) GoString() string {
	return s.String()
}

func (s *ListVpdsRequest) SetEnablePage(v bool) *ListVpdsRequest {
	s.EnablePage = &v
	return s
}

func (s *ListVpdsRequest) SetFilterErId(v string) *ListVpdsRequest {
	s.FilterErId = &v
	return s
}

func (s *ListVpdsRequest) SetForSelect(v bool) *ListVpdsRequest {
	s.ForSelect = &v
	return s
}

func (s *ListVpdsRequest) SetPageNumber(v int32) *ListVpdsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListVpdsRequest) SetPageSize(v int32) *ListVpdsRequest {
	s.PageSize = &v
	return s
}

func (s *ListVpdsRequest) SetRegionId(v string) *ListVpdsRequest {
	s.RegionId = &v
	return s
}

func (s *ListVpdsRequest) SetResourceGroupId(v string) *ListVpdsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVpdsRequest) SetStatus(v string) *ListVpdsRequest {
	s.Status = &v
	return s
}

func (s *ListVpdsRequest) SetTag(v []*ListVpdsRequestTag) *ListVpdsRequest {
	s.Tag = v
	return s
}

func (s *ListVpdsRequest) SetVpdId(v string) *ListVpdsRequest {
	s.VpdId = &v
	return s
}

func (s *ListVpdsRequest) SetVpdName(v string) *ListVpdsRequest {
	s.VpdName = &v
	return s
}

func (s *ListVpdsRequest) SetWithDependence(v bool) *ListVpdsRequest {
	s.WithDependence = &v
	return s
}

func (s *ListVpdsRequest) SetWithoutVcc(v bool) *ListVpdsRequest {
	s.WithoutVcc = &v
	return s
}

type ListVpdsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListVpdsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListVpdsRequestTag) GoString() string {
	return s.String()
}

func (s *ListVpdsRequestTag) SetKey(v string) *ListVpdsRequestTag {
	s.Key = &v
	return s
}

func (s *ListVpdsRequestTag) SetValue(v string) *ListVpdsRequestTag {
	s.Value = &v
	return s
}

type ListVpdsResponseBody struct {
	Code      *int32                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   *ListVpdsResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListVpdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVpdsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpdsResponseBody) SetCode(v int32) *ListVpdsResponseBody {
	s.Code = &v
	return s
}

func (s *ListVpdsResponseBody) SetContent(v *ListVpdsResponseBodyContent) *ListVpdsResponseBody {
	s.Content = v
	return s
}

func (s *ListVpdsResponseBody) SetMessage(v string) *ListVpdsResponseBody {
	s.Message = &v
	return s
}

func (s *ListVpdsResponseBody) SetRequestId(v string) *ListVpdsResponseBody {
	s.RequestId = &v
	return s
}

type ListVpdsResponseBodyContent struct {
	Data  []*ListVpdsResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Total *int64                             `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListVpdsResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s ListVpdsResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListVpdsResponseBodyContent) SetData(v []*ListVpdsResponseBodyContentData) *ListVpdsResponseBodyContent {
	s.Data = v
	return s
}

func (s *ListVpdsResponseBodyContent) SetTotal(v int64) *ListVpdsResponseBodyContent {
	s.Total = &v
	return s
}

type ListVpdsResponseBodyContentData struct {
	Cidr            *string                                   `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	CreateTime      *string                                   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Dependence      map[string]interface{}                    `json:"Dependence,omitempty" xml:"Dependence,omitempty"`
	ErInfos         []*ListVpdsResponseBodyContentDataErInfos `json:"ErInfos,omitempty" xml:"ErInfos,omitempty" type:"Repeated"`
	GmtModified     *string                                   `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Message         *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	NcCount         *int32                                    `json:"NcCount,omitempty" xml:"NcCount,omitempty"`
	RegionId        *string                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string                                   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ServiceCidr     *string                                   `json:"ServiceCidr,omitempty" xml:"ServiceCidr,omitempty"`
	Status          *string                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	SubnetCount     *int32                                    `json:"SubnetCount,omitempty" xml:"SubnetCount,omitempty"`
	Tags            []*ListVpdsResponseBodyContentDataTags    `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TenantId        *string                                   `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// vpd id
	VpdId   *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	VpdName *string `json:"VpdName,omitempty" xml:"VpdName,omitempty"`
}

func (s ListVpdsResponseBodyContentData) String() string {
	return tea.Prettify(s)
}

func (s ListVpdsResponseBodyContentData) GoString() string {
	return s.String()
}

func (s *ListVpdsResponseBodyContentData) SetCidr(v string) *ListVpdsResponseBodyContentData {
	s.Cidr = &v
	return s
}

func (s *ListVpdsResponseBodyContentData) SetCreateTime(v string) *ListVpdsResponseBodyContentData {
	s.CreateTime = &v
	return s
}

func (s *ListVpdsResponseBodyContentData) SetDependence(v map[string]interface{}) *ListVpdsResponseBodyContentData {
	s.Dependence = v
	return s
}

func (s *ListVpdsResponseBodyContentData) SetErInfos(v []*ListVpdsResponseBodyContentDataErInfos) *ListVpdsResponseBodyContentData {
	s.ErInfos = v
	return s
}

func (s *ListVpdsResponseBodyContentData) SetGmtModified(v string) *ListVpdsResponseBodyContentData {
	s.GmtModified = &v
	return s
}

func (s *ListVpdsResponseBodyContentData) SetMessage(v string) *ListVpdsResponseBodyContentData {
	s.Message = &v
	return s
}

func (s *ListVpdsResponseBodyContentData) SetNcCount(v int32) *ListVpdsResponseBodyContentData {
	s.NcCount = &v
	return s
}

func (s *ListVpdsResponseBodyContentData) SetRegionId(v string) *ListVpdsResponseBodyContentData {
	s.RegionId = &v
	return s
}

func (s *ListVpdsResponseBodyContentData) SetResourceGroupId(v string) *ListVpdsResponseBodyContentData {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVpdsResponseBodyContentData) SetServiceCidr(v string) *ListVpdsResponseBodyContentData {
	s.ServiceCidr = &v
	return s
}

func (s *ListVpdsResponseBodyContentData) SetStatus(v string) *ListVpdsResponseBodyContentData {
	s.Status = &v
	return s
}

func (s *ListVpdsResponseBodyContentData) SetSubnetCount(v int32) *ListVpdsResponseBodyContentData {
	s.SubnetCount = &v
	return s
}

func (s *ListVpdsResponseBodyContentData) SetTags(v []*ListVpdsResponseBodyContentDataTags) *ListVpdsResponseBodyContentData {
	s.Tags = v
	return s
}

func (s *ListVpdsResponseBodyContentData) SetTenantId(v string) *ListVpdsResponseBodyContentData {
	s.TenantId = &v
	return s
}

func (s *ListVpdsResponseBodyContentData) SetVpdId(v string) *ListVpdsResponseBodyContentData {
	s.VpdId = &v
	return s
}

func (s *ListVpdsResponseBodyContentData) SetVpdName(v string) *ListVpdsResponseBodyContentData {
	s.VpdName = &v
	return s
}

type ListVpdsResponseBodyContentDataErInfos struct {
	Connections  *int64  `json:"Connections,omitempty" xml:"Connections,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ErId         *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	ErName       *string `json:"ErName,omitempty" xml:"ErName,omitempty"`
	GmtModified  *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	MasterZoneId *string `json:"MasterZoneId,omitempty" xml:"MasterZoneId,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RouteMaps    *int64  `json:"RouteMaps,omitempty" xml:"RouteMaps,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TenantId     *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListVpdsResponseBodyContentDataErInfos) String() string {
	return tea.Prettify(s)
}

func (s ListVpdsResponseBodyContentDataErInfos) GoString() string {
	return s.String()
}

func (s *ListVpdsResponseBodyContentDataErInfos) SetConnections(v int64) *ListVpdsResponseBodyContentDataErInfos {
	s.Connections = &v
	return s
}

func (s *ListVpdsResponseBodyContentDataErInfos) SetCreateTime(v string) *ListVpdsResponseBodyContentDataErInfos {
	s.CreateTime = &v
	return s
}

func (s *ListVpdsResponseBodyContentDataErInfos) SetDescription(v string) *ListVpdsResponseBodyContentDataErInfos {
	s.Description = &v
	return s
}

func (s *ListVpdsResponseBodyContentDataErInfos) SetErId(v string) *ListVpdsResponseBodyContentDataErInfos {
	s.ErId = &v
	return s
}

func (s *ListVpdsResponseBodyContentDataErInfos) SetErName(v string) *ListVpdsResponseBodyContentDataErInfos {
	s.ErName = &v
	return s
}

func (s *ListVpdsResponseBodyContentDataErInfos) SetGmtModified(v string) *ListVpdsResponseBodyContentDataErInfos {
	s.GmtModified = &v
	return s
}

func (s *ListVpdsResponseBodyContentDataErInfos) SetMasterZoneId(v string) *ListVpdsResponseBodyContentDataErInfos {
	s.MasterZoneId = &v
	return s
}

func (s *ListVpdsResponseBodyContentDataErInfos) SetMessage(v string) *ListVpdsResponseBodyContentDataErInfos {
	s.Message = &v
	return s
}

func (s *ListVpdsResponseBodyContentDataErInfos) SetRegionId(v string) *ListVpdsResponseBodyContentDataErInfos {
	s.RegionId = &v
	return s
}

func (s *ListVpdsResponseBodyContentDataErInfos) SetRouteMaps(v int64) *ListVpdsResponseBodyContentDataErInfos {
	s.RouteMaps = &v
	return s
}

func (s *ListVpdsResponseBodyContentDataErInfos) SetStatus(v string) *ListVpdsResponseBodyContentDataErInfos {
	s.Status = &v
	return s
}

func (s *ListVpdsResponseBodyContentDataErInfos) SetTenantId(v string) *ListVpdsResponseBodyContentDataErInfos {
	s.TenantId = &v
	return s
}

type ListVpdsResponseBodyContentDataTags struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListVpdsResponseBodyContentDataTags) String() string {
	return tea.Prettify(s)
}

func (s ListVpdsResponseBodyContentDataTags) GoString() string {
	return s.String()
}

func (s *ListVpdsResponseBodyContentDataTags) SetTagKey(v string) *ListVpdsResponseBodyContentDataTags {
	s.TagKey = &v
	return s
}

func (s *ListVpdsResponseBodyContentDataTags) SetTagValue(v string) *ListVpdsResponseBodyContentDataTags {
	s.TagValue = &v
	return s
}

type ListVpdsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListVpdsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVpdsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVpdsResponse) GoString() string {
	return s.String()
}

func (s *ListVpdsResponse) SetHeaders(v map[string]*string) *ListVpdsResponse {
	s.Headers = v
	return s
}

func (s *ListVpdsResponse) SetStatusCode(v int32) *ListVpdsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVpdsResponse) SetBody(v *ListVpdsResponseBody) *ListVpdsResponse {
	s.Body = v
	return s
}

type UnAssignPrivateIpAddressRequest struct {
	IpName             *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	PrivateIpAddress   *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SubnetId           *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
}

func (s UnAssignPrivateIpAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s UnAssignPrivateIpAddressRequest) GoString() string {
	return s.String()
}

func (s *UnAssignPrivateIpAddressRequest) SetIpName(v string) *UnAssignPrivateIpAddressRequest {
	s.IpName = &v
	return s
}

func (s *UnAssignPrivateIpAddressRequest) SetNetworkInterfaceId(v string) *UnAssignPrivateIpAddressRequest {
	s.NetworkInterfaceId = &v
	return s
}

func (s *UnAssignPrivateIpAddressRequest) SetPrivateIpAddress(v string) *UnAssignPrivateIpAddressRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *UnAssignPrivateIpAddressRequest) SetRegionId(v string) *UnAssignPrivateIpAddressRequest {
	s.RegionId = &v
	return s
}

func (s *UnAssignPrivateIpAddressRequest) SetSubnetId(v string) *UnAssignPrivateIpAddressRequest {
	s.SubnetId = &v
	return s
}

type UnAssignPrivateIpAddressResponseBody struct {
	Code      *int32                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   *UnAssignPrivateIpAddressResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnAssignPrivateIpAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnAssignPrivateIpAddressResponseBody) GoString() string {
	return s.String()
}

func (s *UnAssignPrivateIpAddressResponseBody) SetCode(v int32) *UnAssignPrivateIpAddressResponseBody {
	s.Code = &v
	return s
}

func (s *UnAssignPrivateIpAddressResponseBody) SetContent(v *UnAssignPrivateIpAddressResponseBodyContent) *UnAssignPrivateIpAddressResponseBody {
	s.Content = v
	return s
}

func (s *UnAssignPrivateIpAddressResponseBody) SetMessage(v string) *UnAssignPrivateIpAddressResponseBody {
	s.Message = &v
	return s
}

func (s *UnAssignPrivateIpAddressResponseBody) SetRequestId(v string) *UnAssignPrivateIpAddressResponseBody {
	s.RequestId = &v
	return s
}

type UnAssignPrivateIpAddressResponseBodyContent struct {
	IpName             *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
}

func (s UnAssignPrivateIpAddressResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s UnAssignPrivateIpAddressResponseBodyContent) GoString() string {
	return s.String()
}

func (s *UnAssignPrivateIpAddressResponseBodyContent) SetIpName(v string) *UnAssignPrivateIpAddressResponseBodyContent {
	s.IpName = &v
	return s
}

func (s *UnAssignPrivateIpAddressResponseBodyContent) SetNetworkInterfaceId(v string) *UnAssignPrivateIpAddressResponseBodyContent {
	s.NetworkInterfaceId = &v
	return s
}

type UnAssignPrivateIpAddressResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnAssignPrivateIpAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnAssignPrivateIpAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s UnAssignPrivateIpAddressResponse) GoString() string {
	return s.String()
}

func (s *UnAssignPrivateIpAddressResponse) SetHeaders(v map[string]*string) *UnAssignPrivateIpAddressResponse {
	s.Headers = v
	return s
}

func (s *UnAssignPrivateIpAddressResponse) SetStatusCode(v int32) *UnAssignPrivateIpAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *UnAssignPrivateIpAddressResponse) SetBody(v *UnAssignPrivateIpAddressResponseBody) *UnAssignPrivateIpAddressResponse {
	s.Body = v
	return s
}

type UpdateErRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ErId        *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	ErName      *string `json:"ErName,omitempty" xml:"ErName,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateErRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateErRequest) GoString() string {
	return s.String()
}

func (s *UpdateErRequest) SetDescription(v string) *UpdateErRequest {
	s.Description = &v
	return s
}

func (s *UpdateErRequest) SetErId(v string) *UpdateErRequest {
	s.ErId = &v
	return s
}

func (s *UpdateErRequest) SetErName(v string) *UpdateErRequest {
	s.ErName = &v
	return s
}

func (s *UpdateErRequest) SetRegionId(v string) *UpdateErRequest {
	s.RegionId = &v
	return s
}

type UpdateErResponseBody struct {
	Code      *int32                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateErResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateErResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateErResponseBody) SetCode(v int32) *UpdateErResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateErResponseBody) SetContent(v map[string]interface{}) *UpdateErResponseBody {
	s.Content = v
	return s
}

func (s *UpdateErResponseBody) SetMessage(v string) *UpdateErResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateErResponseBody) SetRequestId(v string) *UpdateErResponseBody {
	s.RequestId = &v
	return s
}

type UpdateErResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateErResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateErResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateErResponse) GoString() string {
	return s.String()
}

func (s *UpdateErResponse) SetHeaders(v map[string]*string) *UpdateErResponse {
	s.Headers = v
	return s
}

func (s *UpdateErResponse) SetStatusCode(v int32) *UpdateErResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateErResponse) SetBody(v *UpdateErResponseBody) *UpdateErResponse {
	s.Body = v
	return s
}

type UpdateErAttachmentRequest struct {
	ErAttachmentId   *string `json:"ErAttachmentId,omitempty" xml:"ErAttachmentId,omitempty"`
	ErAttachmentName *string `json:"ErAttachmentName,omitempty" xml:"ErAttachmentName,omitempty"`
	ErId             *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateErAttachmentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateErAttachmentRequest) GoString() string {
	return s.String()
}

func (s *UpdateErAttachmentRequest) SetErAttachmentId(v string) *UpdateErAttachmentRequest {
	s.ErAttachmentId = &v
	return s
}

func (s *UpdateErAttachmentRequest) SetErAttachmentName(v string) *UpdateErAttachmentRequest {
	s.ErAttachmentName = &v
	return s
}

func (s *UpdateErAttachmentRequest) SetErId(v string) *UpdateErAttachmentRequest {
	s.ErId = &v
	return s
}

func (s *UpdateErAttachmentRequest) SetRegionId(v string) *UpdateErAttachmentRequest {
	s.RegionId = &v
	return s
}

type UpdateErAttachmentResponseBody struct {
	Code      *int32                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateErAttachmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateErAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateErAttachmentResponseBody) SetCode(v int32) *UpdateErAttachmentResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateErAttachmentResponseBody) SetContent(v map[string]interface{}) *UpdateErAttachmentResponseBody {
	s.Content = v
	return s
}

func (s *UpdateErAttachmentResponseBody) SetMessage(v string) *UpdateErAttachmentResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateErAttachmentResponseBody) SetRequestId(v string) *UpdateErAttachmentResponseBody {
	s.RequestId = &v
	return s
}

type UpdateErAttachmentResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateErAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateErAttachmentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateErAttachmentResponse) GoString() string {
	return s.String()
}

func (s *UpdateErAttachmentResponse) SetHeaders(v map[string]*string) *UpdateErAttachmentResponse {
	s.Headers = v
	return s
}

func (s *UpdateErAttachmentResponse) SetStatusCode(v int32) *UpdateErAttachmentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateErAttachmentResponse) SetBody(v *UpdateErAttachmentResponseBody) *UpdateErAttachmentResponse {
	s.Body = v
	return s
}

type UpdateErRouteMapRequest struct {
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ErId         *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	ErRouteMapId *string `json:"ErRouteMapId,omitempty" xml:"ErRouteMapId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateErRouteMapRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateErRouteMapRequest) GoString() string {
	return s.String()
}

func (s *UpdateErRouteMapRequest) SetDescription(v string) *UpdateErRouteMapRequest {
	s.Description = &v
	return s
}

func (s *UpdateErRouteMapRequest) SetErId(v string) *UpdateErRouteMapRequest {
	s.ErId = &v
	return s
}

func (s *UpdateErRouteMapRequest) SetErRouteMapId(v string) *UpdateErRouteMapRequest {
	s.ErRouteMapId = &v
	return s
}

func (s *UpdateErRouteMapRequest) SetRegionId(v string) *UpdateErRouteMapRequest {
	s.RegionId = &v
	return s
}

type UpdateErRouteMapResponseBody struct {
	Code      *int32                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateErRouteMapResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateErRouteMapResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateErRouteMapResponseBody) SetCode(v int32) *UpdateErRouteMapResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateErRouteMapResponseBody) SetContent(v map[string]interface{}) *UpdateErRouteMapResponseBody {
	s.Content = v
	return s
}

func (s *UpdateErRouteMapResponseBody) SetMessage(v string) *UpdateErRouteMapResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateErRouteMapResponseBody) SetRequestId(v string) *UpdateErRouteMapResponseBody {
	s.RequestId = &v
	return s
}

type UpdateErRouteMapResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateErRouteMapResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateErRouteMapResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateErRouteMapResponse) GoString() string {
	return s.String()
}

func (s *UpdateErRouteMapResponse) SetHeaders(v map[string]*string) *UpdateErRouteMapResponse {
	s.Headers = v
	return s
}

func (s *UpdateErRouteMapResponse) SetStatusCode(v int32) *UpdateErRouteMapResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateErRouteMapResponse) SetBody(v *UpdateErRouteMapResponseBody) *UpdateErRouteMapResponse {
	s.Body = v
	return s
}

type UpdateSubnetRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SubnetId   *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	SubnetName *string `json:"SubnetName,omitempty" xml:"SubnetName,omitempty"`
	VpdId      *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	ZoneId     *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s UpdateSubnetRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubnetRequest) GoString() string {
	return s.String()
}

func (s *UpdateSubnetRequest) SetRegionId(v string) *UpdateSubnetRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateSubnetRequest) SetSubnetId(v string) *UpdateSubnetRequest {
	s.SubnetId = &v
	return s
}

func (s *UpdateSubnetRequest) SetSubnetName(v string) *UpdateSubnetRequest {
	s.SubnetName = &v
	return s
}

func (s *UpdateSubnetRequest) SetVpdId(v string) *UpdateSubnetRequest {
	s.VpdId = &v
	return s
}

func (s *UpdateSubnetRequest) SetZoneId(v string) *UpdateSubnetRequest {
	s.ZoneId = &v
	return s
}

type UpdateSubnetResponseBody struct {
	Code      *int32                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   *UpdateSubnetResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSubnetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubnetResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSubnetResponseBody) SetCode(v int32) *UpdateSubnetResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSubnetResponseBody) SetContent(v *UpdateSubnetResponseBodyContent) *UpdateSubnetResponseBody {
	s.Content = v
	return s
}

func (s *UpdateSubnetResponseBody) SetMessage(v string) *UpdateSubnetResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSubnetResponseBody) SetRequestId(v string) *UpdateSubnetResponseBody {
	s.RequestId = &v
	return s
}

type UpdateSubnetResponseBodyContent struct {
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
}

func (s UpdateSubnetResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubnetResponseBodyContent) GoString() string {
	return s.String()
}

func (s *UpdateSubnetResponseBodyContent) SetSubnetId(v string) *UpdateSubnetResponseBodyContent {
	s.SubnetId = &v
	return s
}

type UpdateSubnetResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateSubnetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSubnetResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubnetResponse) GoString() string {
	return s.String()
}

func (s *UpdateSubnetResponse) SetHeaders(v map[string]*string) *UpdateSubnetResponse {
	s.Headers = v
	return s
}

func (s *UpdateSubnetResponse) SetStatusCode(v int32) *UpdateSubnetResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSubnetResponse) SetBody(v *UpdateSubnetResponseBody) *UpdateSubnetResponse {
	s.Body = v
	return s
}

type UpdateVccRequest struct {
	Bandwidth *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VccId     *string `json:"VccId,omitempty" xml:"VccId,omitempty"`
	VccName   *string `json:"VccName,omitempty" xml:"VccName,omitempty"`
}

func (s UpdateVccRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateVccRequest) GoString() string {
	return s.String()
}

func (s *UpdateVccRequest) SetBandwidth(v int32) *UpdateVccRequest {
	s.Bandwidth = &v
	return s
}

func (s *UpdateVccRequest) SetOrderId(v string) *UpdateVccRequest {
	s.OrderId = &v
	return s
}

func (s *UpdateVccRequest) SetRegionId(v string) *UpdateVccRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateVccRequest) SetVccId(v string) *UpdateVccRequest {
	s.VccId = &v
	return s
}

func (s *UpdateVccRequest) SetVccName(v string) *UpdateVccRequest {
	s.VccName = &v
	return s
}

type UpdateVccResponseBody struct {
	Code      *int32                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   *UpdateVccResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateVccResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateVccResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVccResponseBody) SetCode(v int32) *UpdateVccResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateVccResponseBody) SetContent(v *UpdateVccResponseBodyContent) *UpdateVccResponseBody {
	s.Content = v
	return s
}

func (s *UpdateVccResponseBody) SetMessage(v string) *UpdateVccResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateVccResponseBody) SetRequestId(v string) *UpdateVccResponseBody {
	s.RequestId = &v
	return s
}

type UpdateVccResponseBodyContent struct {
	VccId *string `json:"VccId,omitempty" xml:"VccId,omitempty"`
}

func (s UpdateVccResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s UpdateVccResponseBodyContent) GoString() string {
	return s.String()
}

func (s *UpdateVccResponseBodyContent) SetVccId(v string) *UpdateVccResponseBodyContent {
	s.VccId = &v
	return s
}

type UpdateVccResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateVccResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateVccResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateVccResponse) GoString() string {
	return s.String()
}

func (s *UpdateVccResponse) SetHeaders(v map[string]*string) *UpdateVccResponse {
	s.Headers = v
	return s
}

func (s *UpdateVccResponse) SetStatusCode(v int32) *UpdateVccResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVccResponse) SetBody(v *UpdateVccResponseBody) *UpdateVccResponse {
	s.Body = v
	return s
}

type UpdateVpdRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VpdId    *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	VpdName  *string `json:"VpdName,omitempty" xml:"VpdName,omitempty"`
}

func (s UpdateVpdRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateVpdRequest) GoString() string {
	return s.String()
}

func (s *UpdateVpdRequest) SetRegionId(v string) *UpdateVpdRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateVpdRequest) SetVpdId(v string) *UpdateVpdRequest {
	s.VpdId = &v
	return s
}

func (s *UpdateVpdRequest) SetVpdName(v string) *UpdateVpdRequest {
	s.VpdName = &v
	return s
}

type UpdateVpdResponseBody struct {
	Code      *int32                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Content   *UpdateVpdResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateVpdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateVpdResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVpdResponseBody) SetCode(v int32) *UpdateVpdResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateVpdResponseBody) SetContent(v *UpdateVpdResponseBodyContent) *UpdateVpdResponseBody {
	s.Content = v
	return s
}

func (s *UpdateVpdResponseBody) SetMessage(v string) *UpdateVpdResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateVpdResponseBody) SetRequestId(v string) *UpdateVpdResponseBody {
	s.RequestId = &v
	return s
}

type UpdateVpdResponseBodyContent struct {
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
}

func (s UpdateVpdResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s UpdateVpdResponseBodyContent) GoString() string {
	return s.String()
}

func (s *UpdateVpdResponseBodyContent) SetVpdId(v string) *UpdateVpdResponseBodyContent {
	s.VpdId = &v
	return s
}

type UpdateVpdResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateVpdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateVpdResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateVpdResponse) GoString() string {
	return s.String()
}

func (s *UpdateVpdResponse) SetHeaders(v map[string]*string) *UpdateVpdResponse {
	s.Headers = v
	return s
}

func (s *UpdateVpdResponse) SetStatusCode(v int32) *UpdateVpdResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVpdResponse) SetBody(v *UpdateVpdResponseBody) *UpdateVpdResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("eflo"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AssignPrivateIpAddressWithOptions(request *AssignPrivateIpAddressRequest, runtime *util.RuntimeOptions) (_result *AssignPrivateIpAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssignMac)) {
		body["AssignMac"] = request.AssignMac
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkInterfaceId)) {
		body["NetworkInterfaceId"] = request.NetworkInterfaceId
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateIpAddress)) {
		body["PrivateIpAddress"] = request.PrivateIpAddress
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SkipConfig)) {
		body["SkipConfig"] = request.SkipConfig
	}

	if !tea.BoolValue(util.IsUnset(request.SubnetId)) {
		body["SubnetId"] = request.SubnetId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AssignPrivateIpAddress"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AssignPrivateIpAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AssignPrivateIpAddress(request *AssignPrivateIpAddressRequest) (_result *AssignPrivateIpAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssignPrivateIpAddressResponse{}
	_body, _err := client.AssignPrivateIpAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateErWithOptions(request *CreateErRequest, runtime *util.RuntimeOptions) (_result *CreateErResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ErName)) {
		body["ErName"] = request.ErName
	}

	if !tea.BoolValue(util.IsUnset(request.MasterZoneId)) {
		body["MasterZoneId"] = request.MasterZoneId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateEr"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateErResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEr(request *CreateErRequest) (_result *CreateErResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateErResponse{}
	_body, _err := client.CreateErWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateErAttachmentWithOptions(request *CreateErAttachmentRequest, runtime *util.RuntimeOptions) (_result *CreateErAttachmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoReceiveAllRoute)) {
		body["AutoReceiveAllRoute"] = request.AutoReceiveAllRoute
	}

	if !tea.BoolValue(util.IsUnset(request.ErAttachmentName)) {
		body["ErAttachmentName"] = request.ErAttachmentName
	}

	if !tea.BoolValue(util.IsUnset(request.ErId)) {
		body["ErId"] = request.ErId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		body["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceTenantId)) {
		body["ResourceTenantId"] = request.ResourceTenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateErAttachment"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateErAttachmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateErAttachment(request *CreateErAttachmentRequest) (_result *CreateErAttachmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateErAttachmentResponse{}
	_body, _err := client.CreateErAttachmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateErRouteMapWithOptions(request *CreateErRouteMapRequest, runtime *util.RuntimeOptions) (_result *CreateErRouteMapResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationCidrBlock)) {
		body["DestinationCidrBlock"] = request.DestinationCidrBlock
	}

	if !tea.BoolValue(util.IsUnset(request.ErId)) {
		body["ErId"] = request.ErId
	}

	if !tea.BoolValue(util.IsUnset(request.ReceptionInstanceId)) {
		body["ReceptionInstanceId"] = request.ReceptionInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ReceptionInstanceOwner)) {
		body["ReceptionInstanceOwner"] = request.ReceptionInstanceOwner
	}

	if !tea.BoolValue(util.IsUnset(request.ReceptionInstanceType)) {
		body["ReceptionInstanceType"] = request.ReceptionInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RouteMapAction)) {
		body["RouteMapAction"] = request.RouteMapAction
	}

	if !tea.BoolValue(util.IsUnset(request.RouteMapNum)) {
		body["RouteMapNum"] = request.RouteMapNum
	}

	if !tea.BoolValue(util.IsUnset(request.TransmissionInstanceId)) {
		body["TransmissionInstanceId"] = request.TransmissionInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TransmissionInstanceOwner)) {
		body["TransmissionInstanceOwner"] = request.TransmissionInstanceOwner
	}

	if !tea.BoolValue(util.IsUnset(request.TransmissionInstanceType)) {
		body["TransmissionInstanceType"] = request.TransmissionInstanceType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateErRouteMap"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateErRouteMapResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateErRouteMap(request *CreateErRouteMapRequest) (_result *CreateErRouteMapResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateErRouteMapResponse{}
	_body, _err := client.CreateErRouteMapWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSubnetWithOptions(request *CreateSubnetRequest, runtime *util.RuntimeOptions) (_result *CreateSubnetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cidr)) {
		body["Cidr"] = request.Cidr
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SubnetName)) {
		body["SubnetName"] = request.SubnetName
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		body["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.VpdId)) {
		body["VpdId"] = request.VpdId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		body["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSubnet"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSubnetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSubnet(request *CreateSubnetRequest) (_result *CreateSubnetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSubnetResponse{}
	_body, _err := client.CreateSubnetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVccWithOptions(request *CreateVccRequest, runtime *util.RuntimeOptions) (_result *CreateVccResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessCouldService)) {
		body["AccessCouldService"] = request.AccessCouldService
	}

	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		body["Bandwidth"] = request.Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.BgpCidr)) {
		body["BgpCidr"] = request.BgpCidr
	}

	if !tea.BoolValue(util.IsUnset(request.CenId)) {
		body["CenId"] = request.CenId
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionType)) {
		body["ConnectionType"] = request.ConnectionType
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		body["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		body["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VccId)) {
		body["VccId"] = request.VccId
	}

	if !tea.BoolValue(util.IsUnset(request.VccName)) {
		body["VccName"] = request.VccName
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.VpdId)) {
		body["VpdId"] = request.VpdId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		body["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateVcc"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVccResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVcc(request *CreateVccRequest) (_result *CreateVccResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVccResponse{}
	_body, _err := client.CreateVccWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVccGrantRuleWithOptions(request *CreateVccGrantRuleRequest, runtime *util.RuntimeOptions) (_result *CreateVccGrantRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ErId)) {
		body["ErId"] = request.ErId
	}

	if !tea.BoolValue(util.IsUnset(request.GrantTenantId)) {
		body["GrantTenantId"] = request.GrantTenantId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateVccGrantRule"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVccGrantRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVccGrantRule(request *CreateVccGrantRuleRequest) (_result *CreateVccGrantRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVccGrantRuleResponse{}
	_body, _err := client.CreateVccGrantRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVccRouteEntryWithOptions(request *CreateVccRouteEntryRequest, runtime *util.RuntimeOptions) (_result *CreateVccRouteEntryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DestinationCidrBlock)) {
		body["DestinationCidrBlock"] = request.DestinationCidrBlock
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.VccId)) {
		body["VccId"] = request.VccId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateVccRouteEntry"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVccRouteEntryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVccRouteEntry(request *CreateVccRouteEntryRequest) (_result *CreateVccRouteEntryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVccRouteEntryResponse{}
	_body, _err := client.CreateVccRouteEntryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVpdWithOptions(request *CreateVpdRequest, runtime *util.RuntimeOptions) (_result *CreateVpdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cidr)) {
		body["Cidr"] = request.Cidr
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Subnets)) {
		body["Subnets"] = request.Subnets
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		body["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VpdName)) {
		body["VpdName"] = request.VpdName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateVpd"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVpdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVpd(request *CreateVpdRequest) (_result *CreateVpdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVpdResponse{}
	_body, _err := client.CreateVpdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVpdGrantRuleWithOptions(request *CreateVpdGrantRuleRequest, runtime *util.RuntimeOptions) (_result *CreateVpdGrantRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ErId)) {
		body["ErId"] = request.ErId
	}

	if !tea.BoolValue(util.IsUnset(request.GrantTenantId)) {
		body["GrantTenantId"] = request.GrantTenantId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateVpdGrantRule"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVpdGrantRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVpdGrantRule(request *CreateVpdGrantRuleRequest) (_result *CreateVpdGrantRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVpdGrantRuleResponse{}
	_body, _err := client.CreateVpdGrantRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteErWithOptions(request *DeleteErRequest, runtime *util.RuntimeOptions) (_result *DeleteErResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ErId)) {
		body["ErId"] = request.ErId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteEr"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteErResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteEr(request *DeleteErRequest) (_result *DeleteErResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteErResponse{}
	_body, _err := client.DeleteErWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteErAttachmentWithOptions(request *DeleteErAttachmentRequest, runtime *util.RuntimeOptions) (_result *DeleteErAttachmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ErAttachmentId)) {
		body["ErAttachmentId"] = request.ErAttachmentId
	}

	if !tea.BoolValue(util.IsUnset(request.ErId)) {
		body["ErId"] = request.ErId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteErAttachment"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteErAttachmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteErAttachment(request *DeleteErAttachmentRequest) (_result *DeleteErAttachmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteErAttachmentResponse{}
	_body, _err := client.DeleteErAttachmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteErRouteMapWithOptions(request *DeleteErRouteMapRequest, runtime *util.RuntimeOptions) (_result *DeleteErRouteMapResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ErId)) {
		body["ErId"] = request.ErId
	}

	if !tea.BoolValue(util.IsUnset(request.ErRouteMapIds)) {
		body["ErRouteMapIds"] = request.ErRouteMapIds
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteErRouteMap"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteErRouteMapResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteErRouteMap(request *DeleteErRouteMapRequest) (_result *DeleteErRouteMapResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteErRouteMapResponse{}
	_body, _err := client.DeleteErRouteMapWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSubnetWithOptions(request *DeleteSubnetRequest, runtime *util.RuntimeOptions) (_result *DeleteSubnetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SubnetId)) {
		body["SubnetId"] = request.SubnetId
	}

	if !tea.BoolValue(util.IsUnset(request.VpdId)) {
		body["VpdId"] = request.VpdId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		body["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSubnet"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSubnetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSubnet(request *DeleteSubnetRequest) (_result *DeleteSubnetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSubnetResponse{}
	_body, _err := client.DeleteSubnetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVccGrantRuleWithOptions(request *DeleteVccGrantRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteVccGrantRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ErId)) {
		body["ErId"] = request.ErId
	}

	if !tea.BoolValue(util.IsUnset(request.GrantRuleId)) {
		body["GrantRuleId"] = request.GrantRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVccGrantRule"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVccGrantRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVccGrantRule(request *DeleteVccGrantRuleRequest) (_result *DeleteVccGrantRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVccGrantRuleResponse{}
	_body, _err := client.DeleteVccGrantRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVccRouteEntryWithOptions(request *DeleteVccRouteEntryRequest, runtime *util.RuntimeOptions) (_result *DeleteVccRouteEntryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DestinationCidrBlock)) {
		body["DestinationCidrBlock"] = request.DestinationCidrBlock
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.VccId)) {
		body["VccId"] = request.VccId
	}

	if !tea.BoolValue(util.IsUnset(request.VccRouteEntryId)) {
		body["VccRouteEntryId"] = request.VccRouteEntryId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVccRouteEntry"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVccRouteEntryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVccRouteEntry(request *DeleteVccRouteEntryRequest) (_result *DeleteVccRouteEntryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVccRouteEntryResponse{}
	_body, _err := client.DeleteVccRouteEntryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVpdWithOptions(request *DeleteVpdRequest, runtime *util.RuntimeOptions) (_result *DeleteVpdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.VpdId)) {
		body["VpdId"] = request.VpdId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVpd"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVpdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVpd(request *DeleteVpdRequest) (_result *DeleteVpdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVpdResponse{}
	_body, _err := client.DeleteVpdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVpdGrantRuleWithOptions(request *DeleteVpdGrantRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteVpdGrantRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ErId)) {
		body["ErId"] = request.ErId
	}

	if !tea.BoolValue(util.IsUnset(request.GrantRuleId)) {
		body["GrantRuleId"] = request.GrantRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.GrantTenantId)) {
		body["GrantTenantId"] = request.GrantTenantId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVpdGrantRule"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVpdGrantRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVpdGrantRule(request *DeleteVpdGrantRuleRequest) (_result *DeleteVpdGrantRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVpdGrantRuleResponse{}
	_body, _err := client.DeleteVpdGrantRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSlrWithOptions(request *DescribeSlrRequest, runtime *util.RuntimeOptions) (_result *DescribeSlrResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSlr"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSlrResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSlr(request *DescribeSlrRequest) (_result *DescribeSlrResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSlrResponse{}
	_body, _err := client.DescribeSlrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetErWithOptions(request *GetErRequest, runtime *util.RuntimeOptions) (_result *GetErResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ErId)) {
		body["ErId"] = request.ErId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetEr"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetErResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEr(request *GetErRequest) (_result *GetErResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetErResponse{}
	_body, _err := client.GetErWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetErAttachmentWithOptions(request *GetErAttachmentRequest, runtime *util.RuntimeOptions) (_result *GetErAttachmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ErAttachmentId)) {
		body["ErAttachmentId"] = request.ErAttachmentId
	}

	if !tea.BoolValue(util.IsUnset(request.ErId)) {
		body["ErId"] = request.ErId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetErAttachment"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetErAttachmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetErAttachment(request *GetErAttachmentRequest) (_result *GetErAttachmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetErAttachmentResponse{}
	_body, _err := client.GetErAttachmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetErRouteEntryWithOptions(request *GetErRouteEntryRequest, runtime *util.RuntimeOptions) (_result *GetErRouteEntryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ErId)) {
		body["ErId"] = request.ErId
	}

	if !tea.BoolValue(util.IsUnset(request.ErRouteEntryId)) {
		body["ErRouteEntryId"] = request.ErRouteEntryId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetErRouteEntry"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetErRouteEntryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetErRouteEntry(request *GetErRouteEntryRequest) (_result *GetErRouteEntryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetErRouteEntryResponse{}
	_body, _err := client.GetErRouteEntryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetErRouteMapWithOptions(request *GetErRouteMapRequest, runtime *util.RuntimeOptions) (_result *GetErRouteMapResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ErId)) {
		body["ErId"] = request.ErId
	}

	if !tea.BoolValue(util.IsUnset(request.ErRouteMapId)) {
		body["ErRouteMapId"] = request.ErRouteMapId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetErRouteMap"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetErRouteMapResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetErRouteMap(request *GetErRouteMapRequest) (_result *GetErRouteMapResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetErRouteMapResponse{}
	_body, _err := client.GetErRouteMapWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLniPrivateIpAddressWithOptions(request *GetLniPrivateIpAddressRequest, runtime *util.RuntimeOptions) (_result *GetLniPrivateIpAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IpName)) {
		body["IpName"] = request.IpName
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkInterfaceId)) {
		body["NetworkInterfaceId"] = request.NetworkInterfaceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLniPrivateIpAddress"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLniPrivateIpAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLniPrivateIpAddress(request *GetLniPrivateIpAddressRequest) (_result *GetLniPrivateIpAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLniPrivateIpAddressResponse{}
	_body, _err := client.GetLniPrivateIpAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetNetworkInterfaceWithOptions(request *GetNetworkInterfaceRequest, runtime *util.RuntimeOptions) (_result *GetNetworkInterfaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NetworkInterfaceId)) {
		body["NetworkInterfaceId"] = request.NetworkInterfaceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SubnetId)) {
		body["SubnetId"] = request.SubnetId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNetworkInterface"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNetworkInterfaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetNetworkInterface(request *GetNetworkInterfaceRequest) (_result *GetNetworkInterfaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetNetworkInterfaceResponse{}
	_body, _err := client.GetNetworkInterfaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSubnetWithOptions(request *GetSubnetRequest, runtime *util.RuntimeOptions) (_result *GetSubnetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SubnetId)) {
		body["SubnetId"] = request.SubnetId
	}

	if !tea.BoolValue(util.IsUnset(request.VpdId)) {
		body["VpdId"] = request.VpdId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSubnet"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSubnetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSubnet(request *GetSubnetRequest) (_result *GetSubnetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSubnetResponse{}
	_body, _err := client.GetSubnetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetVccWithOptions(request *GetVccRequest, runtime *util.RuntimeOptions) (_result *GetVccResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnablePage)) {
		body["EnablePage"] = request.EnablePage
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.VccId)) {
		body["VccId"] = request.VccId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetVcc"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetVccResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetVcc(request *GetVccRequest) (_result *GetVccResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVccResponse{}
	_body, _err := client.GetVccWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetVccGrantRuleWithOptions(request *GetVccGrantRuleRequest, runtime *util.RuntimeOptions) (_result *GetVccGrantRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ErId)) {
		body["ErId"] = request.ErId
	}

	if !tea.BoolValue(util.IsUnset(request.GrantRuleId)) {
		body["GrantRuleId"] = request.GrantRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.GrantTenantId)) {
		body["GrantTenantId"] = request.GrantTenantId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetVccGrantRule"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetVccGrantRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetVccGrantRule(request *GetVccGrantRuleRequest) (_result *GetVccGrantRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVccGrantRuleResponse{}
	_body, _err := client.GetVccGrantRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetVccRouteEntryWithOptions(request *GetVccRouteEntryRequest, runtime *util.RuntimeOptions) (_result *GetVccRouteEntryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.VccId)) {
		body["VccId"] = request.VccId
	}

	if !tea.BoolValue(util.IsUnset(request.VccRouteEntryId)) {
		body["VccRouteEntryId"] = request.VccRouteEntryId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetVccRouteEntry"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetVccRouteEntryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetVccRouteEntry(request *GetVccRouteEntryRequest) (_result *GetVccRouteEntryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVccRouteEntryResponse{}
	_body, _err := client.GetVccRouteEntryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetVpdWithOptions(request *GetVpdRequest, runtime *util.RuntimeOptions) (_result *GetVpdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.VpdId)) {
		body["VpdId"] = request.VpdId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetVpd"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetVpdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetVpd(request *GetVpdRequest) (_result *GetVpdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVpdResponse{}
	_body, _err := client.GetVpdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetVpdGrantRuleWithOptions(request *GetVpdGrantRuleRequest, runtime *util.RuntimeOptions) (_result *GetVpdGrantRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ErId)) {
		body["ErId"] = request.ErId
	}

	if !tea.BoolValue(util.IsUnset(request.GrantRuleId)) {
		body["GrantRuleId"] = request.GrantRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.GrantTenantId)) {
		body["GrantTenantId"] = request.GrantTenantId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetVpdGrantRule"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetVpdGrantRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetVpdGrantRule(request *GetVpdGrantRuleRequest) (_result *GetVpdGrantRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVpdGrantRuleResponse{}
	_body, _err := client.GetVpdGrantRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetVpdRouteEntryWithOptions(request *GetVpdRouteEntryRequest, runtime *util.RuntimeOptions) (_result *GetVpdRouteEntryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.VpdId)) {
		body["VpdId"] = request.VpdId
	}

	if !tea.BoolValue(util.IsUnset(request.VpdRouteEntryId)) {
		body["VpdRouteEntryId"] = request.VpdRouteEntryId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetVpdRouteEntry"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetVpdRouteEntryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetVpdRouteEntry(request *GetVpdRouteEntryRequest) (_result *GetVpdRouteEntryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVpdRouteEntryResponse{}
	_body, _err := client.GetVpdRouteEntryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InitializeVccWithOptions(request *InitializeVccRequest, runtime *util.RuntimeOptions) (_result *InitializeVccResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("InitializeVcc"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InitializeVccResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InitializeVcc(request *InitializeVccRequest) (_result *InitializeVccResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InitializeVccResponse{}
	_body, _err := client.InitializeVccWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListErAttachmentsWithOptions(request *ListErAttachmentsRequest, runtime *util.RuntimeOptions) (_result *ListErAttachmentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoReceiveAllRoute)) {
		body["AutoReceiveAllRoute"] = request.AutoReceiveAllRoute
	}

	if !tea.BoolValue(util.IsUnset(request.EnablePage)) {
		body["EnablePage"] = request.EnablePage
	}

	if !tea.BoolValue(util.IsUnset(request.ErAttachmentId)) {
		body["ErAttachmentId"] = request.ErAttachmentId
	}

	if !tea.BoolValue(util.IsUnset(request.ErAttachmentName)) {
		body["ErAttachmentName"] = request.ErAttachmentName
	}

	if !tea.BoolValue(util.IsUnset(request.ErId)) {
		body["ErId"] = request.ErId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		body["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceTenantId)) {
		body["ResourceTenantId"] = request.ResourceTenantId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListErAttachments"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListErAttachmentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListErAttachments(request *ListErAttachmentsRequest) (_result *ListErAttachmentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListErAttachmentsResponse{}
	_body, _err := client.ListErAttachmentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListErRouteEntriesWithOptions(request *ListErRouteEntriesRequest, runtime *util.RuntimeOptions) (_result *ListErRouteEntriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DestinationCidrBlock)) {
		body["DestinationCidrBlock"] = request.DestinationCidrBlock
	}

	if !tea.BoolValue(util.IsUnset(request.EnablePage)) {
		body["EnablePage"] = request.EnablePage
	}

	if !tea.BoolValue(util.IsUnset(request.ErId)) {
		body["ErId"] = request.ErId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NextHopId)) {
		body["NextHopId"] = request.NextHopId
	}

	if !tea.BoolValue(util.IsUnset(request.NextHopType)) {
		body["NextHopType"] = request.NextHopType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RouteType)) {
		body["RouteType"] = request.RouteType
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListErRouteEntries"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListErRouteEntriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListErRouteEntries(request *ListErRouteEntriesRequest) (_result *ListErRouteEntriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListErRouteEntriesResponse{}
	_body, _err := client.ListErRouteEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListErRouteMapsWithOptions(request *ListErRouteMapsRequest, runtime *util.RuntimeOptions) (_result *ListErRouteMapsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DestinationCidrBlock)) {
		body["DestinationCidrBlock"] = request.DestinationCidrBlock
	}

	if !tea.BoolValue(util.IsUnset(request.EnablePage)) {
		body["EnablePage"] = request.EnablePage
	}

	if !tea.BoolValue(util.IsUnset(request.ErId)) {
		body["ErId"] = request.ErId
	}

	if !tea.BoolValue(util.IsUnset(request.ErRouteMapId)) {
		body["ErRouteMapId"] = request.ErRouteMapId
	}

	if !tea.BoolValue(util.IsUnset(request.ErRouteMapNum)) {
		body["ErRouteMapNum"] = request.ErRouteMapNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ReceptionInstanceId)) {
		body["ReceptionInstanceId"] = request.ReceptionInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ReceptionInstanceName)) {
		body["ReceptionInstanceName"] = request.ReceptionInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.ReceptionInstanceType)) {
		body["ReceptionInstanceType"] = request.ReceptionInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RouteMapAction)) {
		body["RouteMapAction"] = request.RouteMapAction
	}

	if !tea.BoolValue(util.IsUnset(request.TransmissionInstanceId)) {
		body["TransmissionInstanceId"] = request.TransmissionInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TransmissionInstanceName)) {
		body["TransmissionInstanceName"] = request.TransmissionInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.TransmissionInstanceType)) {
		body["TransmissionInstanceType"] = request.TransmissionInstanceType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListErRouteMaps"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListErRouteMapsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListErRouteMaps(request *ListErRouteMapsRequest) (_result *ListErRouteMapsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListErRouteMapsResponse{}
	_body, _err := client.ListErRouteMapsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListErsWithOptions(request *ListErsRequest, runtime *util.RuntimeOptions) (_result *ListErsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnablePage)) {
		body["EnablePage"] = request.EnablePage
	}

	if !tea.BoolValue(util.IsUnset(request.ErId)) {
		body["ErId"] = request.ErId
	}

	if !tea.BoolValue(util.IsUnset(request.ErName)) {
		body["ErName"] = request.ErName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		body["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.MasterZoneId)) {
		body["MasterZoneId"] = request.MasterZoneId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListErs"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListErsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListErs(request *ListErsRequest) (_result *ListErsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListErsResponse{}
	_body, _err := client.ListErsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListLniPrivateIpAddressWithOptions(request *ListLniPrivateIpAddressRequest, runtime *util.RuntimeOptions) (_result *ListLniPrivateIpAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnablePage)) {
		body["EnablePage"] = request.EnablePage
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		body["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.IpName)) {
		body["IpName"] = request.IpName
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkInterfaceId)) {
		body["NetworkInterfaceId"] = request.NetworkInterfaceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLniPrivateIpAddress"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListLniPrivateIpAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListLniPrivateIpAddress(request *ListLniPrivateIpAddressRequest) (_result *ListLniPrivateIpAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListLniPrivateIpAddressResponse{}
	_body, _err := client.ListLniPrivateIpAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNetworkInterfacesWithOptions(request *ListNetworkInterfacesRequest, runtime *util.RuntimeOptions) (_result *ListNetworkInterfacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnablePage)) {
		body["EnablePage"] = request.EnablePage
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		body["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkInterfaceId)) {
		body["NetworkInterfaceId"] = request.NetworkInterfaceId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		body["NodeId"] = request.NodeId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SubnetId)) {
		body["SubnetId"] = request.SubnetId
	}

	if !tea.BoolValue(util.IsUnset(request.VpdId)) {
		body["VpdId"] = request.VpdId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNetworkInterfaces"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNetworkInterfacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNetworkInterfaces(request *ListNetworkInterfacesRequest) (_result *ListNetworkInterfacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNetworkInterfacesResponse{}
	_body, _err := client.ListNetworkInterfacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSubnetsWithOptions(request *ListSubnetsRequest, runtime *util.RuntimeOptions) (_result *ListSubnetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnablePage)) {
		body["EnablePage"] = request.EnablePage
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.SubnetId)) {
		body["SubnetId"] = request.SubnetId
	}

	if !tea.BoolValue(util.IsUnset(request.SubnetName)) {
		body["SubnetName"] = request.SubnetName
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		body["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.VpdId)) {
		body["VpdId"] = request.VpdId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		body["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSubnets"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSubnetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSubnets(request *ListSubnetsRequest) (_result *ListSubnetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSubnetsResponse{}
	_body, _err := client.ListSubnetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVccGrantRulesWithOptions(request *ListVccGrantRulesRequest, runtime *util.RuntimeOptions) (_result *ListVccGrantRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnablePage)) {
		body["EnablePage"] = request.EnablePage
	}

	if !tea.BoolValue(util.IsUnset(request.ErId)) {
		body["ErId"] = request.ErId
	}

	if !tea.BoolValue(util.IsUnset(request.ForSelect)) {
		body["ForSelect"] = request.ForSelect
	}

	if !tea.BoolValue(util.IsUnset(request.GrantRuleId)) {
		body["GrantRuleId"] = request.GrantRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.GrantTenantId)) {
		body["GrantTenantId"] = request.GrantTenantId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListVccGrantRules"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVccGrantRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVccGrantRules(request *ListVccGrantRulesRequest) (_result *ListVccGrantRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVccGrantRulesResponse{}
	_body, _err := client.ListVccGrantRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVccRouteEntriesWithOptions(request *ListVccRouteEntriesRequest, runtime *util.RuntimeOptions) (_result *ListVccRouteEntriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DestinationCidrBlock)) {
		body["DestinationCidrBlock"] = request.DestinationCidrBlock
	}

	if !tea.BoolValue(util.IsUnset(request.EnablePage)) {
		body["EnablePage"] = request.EnablePage
	}

	if !tea.BoolValue(util.IsUnset(request.NextHopId)) {
		body["NextHopId"] = request.NextHopId
	}

	if !tea.BoolValue(util.IsUnset(request.NextHopType)) {
		body["NextHopType"] = request.NextHopType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RouteType)) {
		body["RouteType"] = request.RouteType
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.VccId)) {
		body["VccId"] = request.VccId
	}

	if !tea.BoolValue(util.IsUnset(request.VpdRouteEntryId)) {
		body["VpdRouteEntryId"] = request.VpdRouteEntryId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListVccRouteEntries"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVccRouteEntriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVccRouteEntries(request *ListVccRouteEntriesRequest) (_result *ListVccRouteEntriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVccRouteEntriesResponse{}
	_body, _err := client.ListVccRouteEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVccsWithOptions(request *ListVccsRequest, runtime *util.RuntimeOptions) (_result *ListVccsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		body["Bandwidth"] = request.Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.CenId)) {
		body["CenId"] = request.CenId
	}

	if !tea.BoolValue(util.IsUnset(request.EnablePage)) {
		body["EnablePage"] = request.EnablePage
	}

	if !tea.BoolValue(util.IsUnset(request.ExStatus)) {
		body["ExStatus"] = request.ExStatus
	}

	if !tea.BoolValue(util.IsUnset(request.FilterErId)) {
		body["FilterErId"] = request.FilterErId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		body["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VccId)) {
		body["VccId"] = request.VccId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.VpdId)) {
		body["VpdId"] = request.VpdId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListVccs"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVccsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVccs(request *ListVccsRequest) (_result *ListVccsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVccsResponse{}
	_body, _err := client.ListVccsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVpdGrantRulesWithOptions(request *ListVpdGrantRulesRequest, runtime *util.RuntimeOptions) (_result *ListVpdGrantRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnablePage)) {
		body["EnablePage"] = request.EnablePage
	}

	if !tea.BoolValue(util.IsUnset(request.ErId)) {
		body["ErId"] = request.ErId
	}

	if !tea.BoolValue(util.IsUnset(request.ForSelect)) {
		body["ForSelect"] = request.ForSelect
	}

	if !tea.BoolValue(util.IsUnset(request.GrantRuleId)) {
		body["GrantRuleId"] = request.GrantRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.GrantTenantId)) {
		body["GrantTenantId"] = request.GrantTenantId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListVpdGrantRules"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVpdGrantRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVpdGrantRules(request *ListVpdGrantRulesRequest) (_result *ListVpdGrantRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVpdGrantRulesResponse{}
	_body, _err := client.ListVpdGrantRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVpdRouteEntriesWithOptions(request *ListVpdRouteEntriesRequest, runtime *util.RuntimeOptions) (_result *ListVpdRouteEntriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DestinationCidrBlock)) {
		body["DestinationCidrBlock"] = request.DestinationCidrBlock
	}

	if !tea.BoolValue(util.IsUnset(request.EnablePage)) {
		body["EnablePage"] = request.EnablePage
	}

	if !tea.BoolValue(util.IsUnset(request.NextHopId)) {
		body["NextHopId"] = request.NextHopId
	}

	if !tea.BoolValue(util.IsUnset(request.NextHopType)) {
		body["NextHopType"] = request.NextHopType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RouteType)) {
		body["RouteType"] = request.RouteType
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.VpdId)) {
		body["VpdId"] = request.VpdId
	}

	if !tea.BoolValue(util.IsUnset(request.VpdRouteEntryId)) {
		body["VpdRouteEntryId"] = request.VpdRouteEntryId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListVpdRouteEntries"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVpdRouteEntriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVpdRouteEntries(request *ListVpdRouteEntriesRequest) (_result *ListVpdRouteEntriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVpdRouteEntriesResponse{}
	_body, _err := client.ListVpdRouteEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVpdsWithOptions(request *ListVpdsRequest, runtime *util.RuntimeOptions) (_result *ListVpdsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnablePage)) {
		body["EnablePage"] = request.EnablePage
	}

	if !tea.BoolValue(util.IsUnset(request.FilterErId)) {
		body["FilterErId"] = request.FilterErId
	}

	if !tea.BoolValue(util.IsUnset(request.ForSelect)) {
		body["ForSelect"] = request.ForSelect
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		body["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VpdId)) {
		body["VpdId"] = request.VpdId
	}

	if !tea.BoolValue(util.IsUnset(request.VpdName)) {
		body["VpdName"] = request.VpdName
	}

	if !tea.BoolValue(util.IsUnset(request.WithDependence)) {
		body["WithDependence"] = request.WithDependence
	}

	if !tea.BoolValue(util.IsUnset(request.WithoutVcc)) {
		body["WithoutVcc"] = request.WithoutVcc
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListVpds"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVpdsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVpds(request *ListVpdsRequest) (_result *ListVpdsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVpdsResponse{}
	_body, _err := client.ListVpdsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnAssignPrivateIpAddressWithOptions(request *UnAssignPrivateIpAddressRequest, runtime *util.RuntimeOptions) (_result *UnAssignPrivateIpAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IpName)) {
		body["IpName"] = request.IpName
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkInterfaceId)) {
		body["NetworkInterfaceId"] = request.NetworkInterfaceId
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateIpAddress)) {
		body["PrivateIpAddress"] = request.PrivateIpAddress
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SubnetId)) {
		body["SubnetId"] = request.SubnetId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UnAssignPrivateIpAddress"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnAssignPrivateIpAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnAssignPrivateIpAddress(request *UnAssignPrivateIpAddressRequest) (_result *UnAssignPrivateIpAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnAssignPrivateIpAddressResponse{}
	_body, _err := client.UnAssignPrivateIpAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateErWithOptions(request *UpdateErRequest, runtime *util.RuntimeOptions) (_result *UpdateErResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ErId)) {
		body["ErId"] = request.ErId
	}

	if !tea.BoolValue(util.IsUnset(request.ErName)) {
		body["ErName"] = request.ErName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateEr"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateErResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateEr(request *UpdateErRequest) (_result *UpdateErResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateErResponse{}
	_body, _err := client.UpdateErWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateErAttachmentWithOptions(request *UpdateErAttachmentRequest, runtime *util.RuntimeOptions) (_result *UpdateErAttachmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ErAttachmentId)) {
		body["ErAttachmentId"] = request.ErAttachmentId
	}

	if !tea.BoolValue(util.IsUnset(request.ErAttachmentName)) {
		body["ErAttachmentName"] = request.ErAttachmentName
	}

	if !tea.BoolValue(util.IsUnset(request.ErId)) {
		body["ErId"] = request.ErId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateErAttachment"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateErAttachmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateErAttachment(request *UpdateErAttachmentRequest) (_result *UpdateErAttachmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateErAttachmentResponse{}
	_body, _err := client.UpdateErAttachmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateErRouteMapWithOptions(request *UpdateErRouteMapRequest, runtime *util.RuntimeOptions) (_result *UpdateErRouteMapResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ErId)) {
		body["ErId"] = request.ErId
	}

	if !tea.BoolValue(util.IsUnset(request.ErRouteMapId)) {
		body["ErRouteMapId"] = request.ErRouteMapId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateErRouteMap"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateErRouteMapResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateErRouteMap(request *UpdateErRouteMapRequest) (_result *UpdateErRouteMapResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateErRouteMapResponse{}
	_body, _err := client.UpdateErRouteMapWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSubnetWithOptions(request *UpdateSubnetRequest, runtime *util.RuntimeOptions) (_result *UpdateSubnetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SubnetId)) {
		body["SubnetId"] = request.SubnetId
	}

	if !tea.BoolValue(util.IsUnset(request.SubnetName)) {
		body["SubnetName"] = request.SubnetName
	}

	if !tea.BoolValue(util.IsUnset(request.VpdId)) {
		body["VpdId"] = request.VpdId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		body["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSubnet"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSubnetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSubnet(request *UpdateSubnetRequest) (_result *UpdateSubnetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSubnetResponse{}
	_body, _err := client.UpdateSubnetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateVccWithOptions(request *UpdateVccRequest, runtime *util.RuntimeOptions) (_result *UpdateVccResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		body["Bandwidth"] = request.Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		body["OrderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.VccId)) {
		body["VccId"] = request.VccId
	}

	if !tea.BoolValue(util.IsUnset(request.VccName)) {
		body["VccName"] = request.VccName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateVcc"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateVccResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateVcc(request *UpdateVccRequest) (_result *UpdateVccResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateVccResponse{}
	_body, _err := client.UpdateVccWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateVpdWithOptions(request *UpdateVpdRequest, runtime *util.RuntimeOptions) (_result *UpdateVpdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.VpdId)) {
		body["VpdId"] = request.VpdId
	}

	if !tea.BoolValue(util.IsUnset(request.VpdName)) {
		body["VpdName"] = request.VpdName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateVpd"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateVpdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateVpd(request *UpdateVpdRequest) (_result *UpdateVpdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateVpdResponse{}
	_body, _err := client.UpdateVpdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
