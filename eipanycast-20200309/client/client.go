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
	// The maximum bandwidth of the Anycast EIP. Unit: Mbit/s.
	//
	// Valid values: **200** to **1000**.
	//
	// Default value: **1000**.
	//
	// > The maximum bandwidth is not a guaranteed service and is for reference only.
	Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID** as the **client token**. The **request ID** may be different for each request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the Anycast EIP.
	//
	// The description must be 0 to 256 characters in length and cannot start with `http://` or `https://`.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The billing method of the Anycast EIP.
	//
	// Set the value to **PostPaid**, which specifies the pay-as-you-go billing method.
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The metering method of the Anycast EIP.
	//
	// Set the value to **PayByTraffic**, which specifies the pay-by-data-transfer metering method.
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The name of the Anycast EIP.
	//
	// The name must be 0 to 128 characters in length, and can contain letters, digits, underscores (\_), and hyphens (-). It must start with a letter.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the resource group to which the instance belongs.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The access area of the Anycast EIP.
	//
	// Set the value to **international**, which specifies the areas outside the Chinese mainland.
	ServiceLocation *string `json:"ServiceLocation,omitempty" xml:"ServiceLocation,omitempty"`
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

func (s *AllocateAnycastEipAddressRequest) SetResourceGroupId(v string) *AllocateAnycastEipAddressRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AllocateAnycastEipAddressRequest) SetServiceLocation(v string) *AllocateAnycastEipAddressRequest {
	s.ServiceLocation = &v
	return s
}

type AllocateAnycastEipAddressResponseBody struct {
	// The ID of the Anycast EIP.
	AnycastId *string `json:"AnycastId,omitempty" xml:"AnycastId,omitempty"`
	// The order ID.
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AllocateAnycastEipAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the Anycast EIP.
	AnycastId *string `json:"AnycastId,omitempty" xml:"AnycastId,omitempty"`
	// The association mode. Valid values:
	//
	// *   **Default**: the default mode. In this mode, the endpoint to be associated serves as the default origin server.
	// *   **Normal**: the standard mode. In this mode, the endpoint to be associated serves as a standard origin server.
	//
	// > You can associate endpoints in multiple regions with an Anycast EIP. However, only one endpoint can serve as the default origin server. Others serve as standard origin servers. If you do not specify or add an access point, requests are forwarded to the default origin server.\
	//
	//
	// *   If this is your first time to associate an Anycast EIP with an endpoint, set the value to **Default**.
	// *   If not, you can also set the value to **Default**, which specifies a new default origin server. In this case, the previous origin server functions as a standard origin server.
	AssociationMode *string `json:"AssociationMode,omitempty" xml:"AssociationMode,omitempty"`
	// The ID of the endpoint with which you want to associate the Anycast EIP.
	BindInstanceId *string `json:"BindInstanceId,omitempty" xml:"BindInstanceId,omitempty"`
	// The ID of the region where the endpoint is deployed.
	//
	// You can associate Anycast EIPs only with endpoints in specific regions. You can call the [DescribeAnycastServerRegions](~~171939~~) operation to query the region IDs.
	BindInstanceRegionId *string `json:"BindInstanceRegionId,omitempty" xml:"BindInstanceRegionId,omitempty"`
	// The type of endpoint with which you want to associate the Anycast EIP. Valid values:
	//
	// *   **SlbInstance**: internal-facing Server Load Balancer (SLB) instance that is deployed in a virtual private cloud (VPC)
	// *   **NetworkInterface**: elastic network interface (ENI)
	BindInstanceType *string `json:"BindInstanceType,omitempty" xml:"BindInstanceType,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID** as the **client token**. The **request ID** may be different for each request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// *   **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	// *   **false**(default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The information about the access points in associated access areas when you associate an Anycast EIP with an endpoint.
	//
	// If this is your first time to associate an Anycast EIP with an endpoint, ignore this parameter. The system automatically associates all access areas.
	//
	// You can call the [DescribeAnycastPopLocations](~~171938~~) operation to query information about access points in supported access areas.
	PopLocations []*AssociateAnycastEipAddressRequestPopLocations `json:"PopLocations,omitempty" xml:"PopLocations,omitempty" type:"Repeated"`
	// The secondary private IP address of the ENI with which you want to associate the Anycast EIP.
	//
	// This parameter is valid only when you set **BindInstanceType** to **NetworkInterface**. If you do not set this parameter, the primary private IP address of the ENI is used.
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
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
	// The information about the access points in associated access areas when you associate an Anycast EIP with an endpoint.
	//
	// If this is your first time to associate an Anycast EIP with an endpoint, ignore this parameter. The system automatically associates all access areas.
	//
	// You can call the [DescribeAnycastPopLocations](~~171938~~) operation to query information about access points in supported access areas.
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
	// The request ID.
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateAnycastEipAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the Anycast EIP.
	//
	// > You must specify **Ip** or **AnycastId**.
	AnycastId *string `json:"AnycastId,omitempty" xml:"AnycastId,omitempty"`
	// The ID of the endpoint with which the Anycast EIP is associated.
	BindInstanceId *string `json:"BindInstanceId,omitempty" xml:"BindInstanceId,omitempty"`
	// The IP address of the Anycast EIP.
	//
	// > You must specify **Ip** or **AnycastId**.
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
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
	// The ID of the account to which the Anycast EIP belongs.
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The information about the endpoint with which the Anycast EIP is associated.
	AnycastEipBindInfoList []*DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList `json:"AnycastEipBindInfoList,omitempty" xml:"AnycastEipBindInfoList,omitempty" type:"Repeated"`
	// The ID of the Anycast EIP.
	AnycastId *string `json:"AnycastId,omitempty" xml:"AnycastId,omitempty"`
	// The maximum bandwidth of the Anycast EIP. Unit: Mbit/s.
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The BID of the account to which the Anycast EIP belongs.
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// The service status of the Anycast EIP. Valid values:
	//
	// *   **Normal**
	// *   **FinancialLocked**
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The point in time at which the Anycast EIP was created.
	//
	// The time follows the ISO8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time is displayed in UTC.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the Anycast EIP.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The billing method of the Anycast EIP.
	//
	// Only **PostPaid** may be returned, which indicates the pay-as-you-go billing method.
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The metering method of the Anycast EIP.
	//
	// Only **PayByTraffic** may be returned, which indicates the pay-by-data-transfer metering method.
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The IP address of the Anycast EIP.
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The name of the Anycast EIP.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the instance belongs.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The access area of the Anycast EIP.
	//
	// Only **international** may be returned, which indicates the areas outside the Chinese mainland.
	ServiceLocation *string `json:"ServiceLocation,omitempty" xml:"ServiceLocation,omitempty"`
	// The status of the Anycast EIP.
	//
	// *   **Associating**
	// *   **Unassociating**
	// *   **Allocated**
	// *   **Associated**
	// *   **Modifying**
	// *   **Releasing**
	// *   **Released**
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The information about the tags.
	Tags []*DescribeAnycastEipAddressResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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

func (s *DescribeAnycastEipAddressResponseBody) SetResourceGroupId(v string) *DescribeAnycastEipAddressResponseBody {
	s.ResourceGroupId = &v
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
	// The association mode. Valid values:
	//
	// *   **Default**: the default mode. In this mode, the associated endpoint serves as the default origin server.
	// *   **Normal**: the standard mode. In this mode, the associated endpoint serves as a standard origin server.
	AssociationMode *string `json:"AssociationMode,omitempty" xml:"AssociationMode,omitempty"`
	// The ID of the endpoint with which the Anycast EIP is associated.
	BindInstanceId *string `json:"BindInstanceId,omitempty" xml:"BindInstanceId,omitempty"`
	// The ID of the region in which the endpoint is deployed.
	BindInstanceRegionId *string `json:"BindInstanceRegionId,omitempty" xml:"BindInstanceRegionId,omitempty"`
	// The type of endpoint with which the Anycast EIP is associated. Valid values:
	//
	// *   **SlbInstance**: a CLB instance in a VPC.
	// *   **NetworkInterface**: an elastic network interface (ENI).
	BindInstanceType *string `json:"BindInstanceType,omitempty" xml:"BindInstanceType,omitempty"`
	// The time when the Anycast EIP was associated.
	//
	// The time follows the ISO 8601 standard in the `YYYY-MM-DDThh:mm:ssZ` format. The time is displayed in UTC.
	BindTime *string `json:"BindTime,omitempty" xml:"BindTime,omitempty"`
	// The information about the access points in associated access areas when you associate an Anycast EIP with a cloud resource.
	//
	// If this is your first time associating an Anycast EIP with an endpoint, the system returns information about access points in all access areas.
	PopLocations []*DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoListPopLocations `json:"PopLocations,omitempty" xml:"PopLocations,omitempty" type:"Repeated"`
	// The secondary private IP address of the associated ENI.
	//
	// This parameter is valid only when **BindInstanceType** is set to **NetworkInterface**.
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The status of the endpoint. Valid values:
	//
	// *   **BINDING**
	// *   **BINDED**
	// *   **UNBINDING**
	// *   **DELETED**
	// *   **MODIFYING**
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	// The information about the access points in associated access areas when you associate an Anycast EIP with a cloud resource.
	//
	// If this is your first time associating an Anycast EIP with an endpoint, the system returns information about access points in all access areas.
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
	// The tag key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAnycastEipAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The access area of the Anycast elastic IP address (EIP).
	//
	// Set the value to **international**, which specifies the areas outside the Chinese mainland.
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
	// The list of access points in the specified access area.
	AnycastPopLocationList []*DescribeAnycastPopLocationsResponseBodyAnycastPopLocationList `json:"AnycastPopLocationList,omitempty" xml:"AnycastPopLocationList,omitempty" type:"Repeated"`
	// The number of access points.
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The ID of the region where the access point is deployed.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the region where the access point is deployed.
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAnycastPopLocationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The access area from which you use the Anycast EIP to communicate with the Internet.
	//
	// Set the value to **international**, which specifies the areas outside the Chinese mainland.
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
	// The list of regions where you can associate Anycast EIPs with endpoints.
	AnycastServerRegionList []*DescribeAnycastServerRegionsResponseBodyAnycastServerRegionList `json:"AnycastServerRegionList,omitempty" xml:"AnycastServerRegionList,omitempty" type:"Repeated"`
	// The number of returned entries.
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the region.
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAnycastServerRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The IP address that is allocated to the Anycast EIP.
	AnycastEipAddress *string `json:"AnycastEipAddress,omitempty" xml:"AnycastEipAddress,omitempty"`
	// The ID of the Anycast EIP.
	//
	// >  To ensure the accuracy of the query result, we do not recommend that you specify **AnycastIds** and **AnycastId** at the same time.
	AnycastId *string `json:"AnycastId,omitempty" xml:"AnycastId,omitempty"`
	// The IDs of Anycast EIPs.
	//
	// You can enter at most 50 Anycast EIP IDs.
	//
	// >  To ensure the accuracy of the query result, we do not recommend that you specify **AnycastIds** and **AnycastId** at the same time.
	AnycastIds []*string `json:"AnycastIds,omitempty" xml:"AnycastIds,omitempty" type:"Repeated"`
	// The IDs of the cloud resources with which the Anycast EIPs are associated.
	//
	// You can enter at most 100 cloud resource IDs.
	BindInstanceIds []*string `json:"BindInstanceIds,omitempty" xml:"BindInstanceIds,omitempty" type:"Repeated"`
	// The service status of the Anycast EIP. Valid values:
	//
	// *   **Normal**
	// *   **FinancialLocked**
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The billing method of the Anycast EIP.
	//
	// Set the value to **PostPaid**, which specifies the pay-as-you-go billing method.
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The metering method of the Anycast EIP.
	//
	// Set the value to **PayByTraffic**, which specifies the pay-by-data-transfer metering method.
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The number of entries to return on each page. Valid values: **20 to 100**. Default value: **50**.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The name of the Anycast EIP.
	//
	// The name must be 0 to 128 characters in length, and can contain digits, hyphens (-), and underscores (\_). The name must start with a letter.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// *   You do not need to specify this parameter for the first request.
	// *   You must specify the token that is obtained from the previous query as the value of NextToken.
	NextToken       *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The access area of the Anycast EIP.
	//
	// Set the value to **international**, which specifies the regions outside the Chinese mainland.
	ServiceLocation *string `json:"ServiceLocation,omitempty" xml:"ServiceLocation,omitempty"`
	// The status of the Anycast EIP. Valid values:
	//
	// *   **Associating**
	// *   **Unassociating**
	// *   **Allocated**
	// *   **Associated**
	// *   **Modifying**
	// *   **Releasing**
	// *   **Released**
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	Tags []*ListAnycastEipAddressesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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

func (s *ListAnycastEipAddressesRequest) SetResourceGroupId(v string) *ListAnycastEipAddressesRequest {
	s.ResourceGroupId = &v
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
	// The tag key of the resource. You can specify up to 20 tag keys. You cannot specify empty strings as tag keys.
	//
	// The key can be up to 64 characters in length and can contain letters, digits, periods (.), underscores (\_), and hyphens (-). The key must start with a letter but cannot start with `aliyun` or `acs:`. The key cannot contain `http://` or `https://`.
	//
	// >  You must specify at least one of **Tag.N** (**Tag.N.Key** and **Tag.N.Value**).
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You can specify up to 20 tag values. It can be an empty string.
	//
	// The value cannot exceed 128 characters in length and can contain digits, periods (.), underscores (\_), and hyphens (-). The value must start with a letter but cannot start with `aliyun` or `acs:`. The value cannot contain `http://` or `https://`.
	//
	// >  You must specify at least one of **Tag.N** (**Tag.N.Key** and **Tag.N.Value**).
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
	// The list of Anycast EIPs.
	AnycastList []*ListAnycastEipAddressesResponseBodyAnycastList `json:"AnycastList,omitempty" xml:"AnycastList,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// *   If **NextToken** is empty, no next page exists.
	// *   If **NextToken** is not empty, the value of NextToken can be used in the next request to retrieve a new page of results.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// The ID of the account to which the Anycast EIP belongs.
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The list of cloud resources with which the Anycast EIPs are associated.
	AnycastEipBindInfoList []*ListAnycastEipAddressesResponseBodyAnycastListAnycastEipBindInfoList `json:"AnycastEipBindInfoList,omitempty" xml:"AnycastEipBindInfoList,omitempty" type:"Repeated"`
	// The ID of the Anycast EIP.
	AnycastId *string `json:"AnycastId,omitempty" xml:"AnycastId,omitempty"`
	// The maximum bandwidth of the Anycast EIP. Unit: Mbit/s.
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The service status of the Anycast EIP. Valid values:
	//
	// *   **Normal**
	// *   **FinancialLocked**
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The time when the Anycast EIP was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the Anycast EIP.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The billing method of the Anycast EIP.
	//
	// **PostPaid**: pay-as-you-go
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The metering method of the Anycast EIP.
	//
	// **PayByTraffic**: pay-by-data-transfer
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The IP address of the Anycast EIP.
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The name of the Anycast EIP.
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The access area of the Anycast EIP.
	//
	// **international**: regions outside the Chinese mainland
	ServiceLocation *string `json:"ServiceLocation,omitempty" xml:"ServiceLocation,omitempty"`
	// Indicates whether the resource is created by the service account.
	//
	// *   **0**: no
	// *   **1**: yes
	ServiceManaged *int32 `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
	// The status of the Anycast EIP.
	//
	// *   **Associating**
	// *   **Unassociating**
	// *   **Allocated**
	// *   **Associated**
	// *   **Modifying**
	// *   **Releasing**
	// *   **Released**
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The information about the tags.
	Tags []*ListAnycastEipAddressesResponseBodyAnycastListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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

func (s *ListAnycastEipAddressesResponseBodyAnycastList) SetResourceGroupId(v string) *ListAnycastEipAddressesResponseBodyAnycastList {
	s.ResourceGroupId = &v
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
	// The ID of the cloud resource with which the Anycast EIP is associated.
	BindInstanceId *string `json:"BindInstanceId,omitempty" xml:"BindInstanceId,omitempty"`
	// The ID of the region where the cloud resource is deployed.
	BindInstanceRegionId *string `json:"BindInstanceRegionId,omitempty" xml:"BindInstanceRegionId,omitempty"`
	// The type of cloud resource with which the Anycast EIP is associated.
	//
	// *   **SlbInstance**: an internal-facing Classic Load Balancer (CLB) instance deployed in a virtual private cloud (VPC). CLB is formerly known as Server Load Balancer (SLB).
	// *   **NetworkInterface**: an elastic network interface (ENI).
	BindInstanceType *string `json:"BindInstanceType,omitempty" xml:"BindInstanceType,omitempty"`
	// The time when the Anycast EIP was associated.
	BindTime *string `json:"BindTime,omitempty" xml:"BindTime,omitempty"`
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
	// The tag key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAnycastEipAddressesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ListTagResourcesRequest struct {
	// The number of entries per page. Valid values: **1** to **50**. Default value: **50**.
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// *   If this is your first query or no next queries are to be sent, ignore this parameter.
	// *   You must specify the token that is obtained from the previous query as the value of **NextToken**.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The resource IDs.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type. Set the value to **ANYCASTEIPADDRESS**.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag information.
	Tag []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetMaxResults(v string) *ListTagResourcesRequest {
	s.MaxResults = &v
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

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	// The key of tag N. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be a up to 128 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `acs:` or `aliyun`.
	//
	// > You must specify **ResourceId.N** or **Tag.N** (**Tag.N.Key** or **Tag.N.Value**).
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// It can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// > You must specify **ResourceId.N** or **Tag.N** (**Tag.N.Key** or **Tag.N.Value**).
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
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// *   If **NextToken** is empty, no next page exists.
	// *   If a value is returned for **NextToken**, the value is the token that determines the start point of the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resources to which the tags are added.
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
	// The resource ID.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type. Only **ANYCASTEIPADDRESS** may be returned.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The key of tag N.
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of tag N.
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

type ModifyAnycastEipAddressAttributeRequest struct {
	// The ID of the Anycast EIP.
	AnycastId *string `json:"AnycastId,omitempty" xml:"AnycastId,omitempty"`
	// The description of the Anycast EIP.
	//
	// The description must be 0 to 256 characters in length and cannot start with `http://` or `https://`.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the Anycast EIP.
	//
	// The name must be 0 to 128 characters in length, and can contain letters, digits, underscores (\_), and hyphens (-). It must start with a letter.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// The request ID.
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
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAnycastEipAddressAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the Anycast EIP.
	AnycastId *string `json:"AnycastId,omitempty" xml:"AnycastId,omitempty"`
	// The maximum bandwidth of the Anycast EIP. Unit: Mbit/s.
	//
	// Valid values: **200** to **1000**.
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
	// The request ID.
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAnycastEipAddressSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the Anycast EIP to be released.
	AnycastId *string `json:"AnycastId,omitempty" xml:"AnycastId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
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
	// The request ID.
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseAnycastEipAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type TagResourcesRequest struct {
	// The resource ID. You can specify at most 20 IDs.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type. Set the value to **ANYCASTEIPADDRESS**.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag information.
	Tag []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	// The key of tag N to add to the resource. You must enter at least one tag key and at most 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length and cannot start with acs: or aliyun. It cannot contain `http://` or `https://`.
	//
	// > When you call this operation, **Tag.N.Key** is required.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the resource. You must enter at least one tag value and at most 20 tag values. The tag value can be an empty string.
	//
	// It can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// > When you call this operation, **Tag.N.Value** is required.
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
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. Valid values:
	//
	// **true**
	//
	// **false**
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s *TagResourcesResponseBody) SetSuccess(v bool) *TagResourcesResponseBody {
	s.Success = &v
	return s
}

type TagResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UnassociateAnycastEipAddressRequest struct {
	// The ID of the Anycast EIP.
	AnycastId *string `json:"AnycastId,omitempty" xml:"AnycastId,omitempty"`
	// The ID of the endpoint from which you want to disassociate the Anycast EIP.
	BindInstanceId *string `json:"BindInstanceId,omitempty" xml:"BindInstanceId,omitempty"`
	// The region where the endpoint is deployed.
	BindInstanceRegionId *string `json:"BindInstanceRegionId,omitempty" xml:"BindInstanceRegionId,omitempty"`
	// The type of endpoint from which you want to disassociate the Anycast EIP. Valid values:
	//
	// *   **SlbInstance**: an internal-facing Server Load Balancer (SLB) instance that is deployed in a virtual private cloud (VPC)
	// *   **NetworkInterface**: elastic network interface (ENI)
	BindInstanceType *string `json:"BindInstanceType,omitempty" xml:"BindInstanceType,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID** as the **client token**. The **request ID** may be different for each request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run, without performing the actual request. Valid values:
	//
	// *   **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	// *   **false** (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The secondary private IP address of the ENI from which you want to disassociate the Anycast EIP.
	//
	// This parameter is valid only when you set **BindInstanceType** to **NetworkInterface**. If you do not specify this parameter, the primary private IP address of the ENI is used.
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
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
	// The request ID.
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnassociateAnycastEipAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type UntagResourcesRequest struct {
	// The resource ID. You can specify up to 20 resource IDs.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type. Set the value to **ANYCASTEIPADDRESS**.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The key of the tag that you want to remove. You can specify at most 20 tag keys. The tag key cannot be an empty string.
	//
	// A tag key can be up to 128 characters in length. It cannot start with aliyun or acs:, and cannot contain `http://` or `https://`.
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. Valid values:
	//
	// **true**
	//
	// **false**
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s *UntagResourcesResponseBody) SetSuccess(v bool) *UntagResourcesResponseBody {
	s.Success = &v
	return s
}

type UntagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpdateAnycastEipAddressAssociationsRequest struct {
	// The ID of the Anycast EIP.
	AnycastId *string `json:"AnycastId,omitempty" xml:"AnycastId,omitempty"`
	// The association mode. Valid values:
	//
	// *   **Default**: the default mode. In this mode, cloud resources to be associated are set as default origin servers.
	// *   **Normal**: the standard mode. In this mode, cloud resources to be associated are set as standard origin servers.
	AssociationMode *string `json:"AssociationMode,omitempty" xml:"AssociationMode,omitempty"`
	// The ID of the cloud resource with which you want to associate the Anycast EIP.
	BindInstanceId *string `json:"BindInstanceId,omitempty" xml:"BindInstanceId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// >  If you do not set this parameter, the system automatically uses **RequestId** as **ClientToken**. **RequestId** may be different for each API request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to only precheck the request. Valid values:
	//
	// *   **true**: prechecks the request without updating the association information. The system checks the required parameters, request syntax, and limits. If the request fails to pass the precheck, an error message is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	// *   **false** (default): sends the API request. If the request passes the precheck, a 2xx HTTP status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The access areas and access points to be added.
	PopLocationAddList []*UpdateAnycastEipAddressAssociationsRequestPopLocationAddList `json:"PopLocationAddList,omitempty" xml:"PopLocationAddList,omitempty" type:"Repeated"`
	// The access areas and access points to be deleted.
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
	// The access points in the access areas to be added.
	//
	// You can call the [DescribeAnycastPopLocations](~~171938~~) operation to query the access points in supported access areas.
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
	// The access points in the access areas to be deleted.
	//
	// >  If the access point in the access area is associated with a default origin server, you cannot delete the access point in the access area.
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
	// The ID of the request.
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
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAnycastEipAddressAssociationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
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

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
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

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
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
		Version:     tea.String("2020-03-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Action:      tea.String("TagResources"),
		Version:     tea.String("2020-03-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2020-03-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
