// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AssignLeniPrivateIpAddressRequest struct {
	// The idempotent identifier.
	//
	// example:
	//
	// 3fd79d62-ab1d-11ec-9a53-0242ac110004
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the response code.
	//
	// example:
	//
	// wuhuaiyu
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Lingjun Elastic Network Interface ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// leni-1234****
	ElasticNetworkInterfaceId *string `json:"ElasticNetworkInterfaceId,omitempty" xml:"ElasticNetworkInterfaceId,omitempty"`
	// Lingjun Elastic Network Interface secondary private network IP (automatically assigned by default).
	//
	// example:
	//
	// 10.0.****
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AssignLeniPrivateIpAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s AssignLeniPrivateIpAddressRequest) GoString() string {
	return s.String()
}

func (s *AssignLeniPrivateIpAddressRequest) SetClientToken(v string) *AssignLeniPrivateIpAddressRequest {
	s.ClientToken = &v
	return s
}

func (s *AssignLeniPrivateIpAddressRequest) SetDescription(v string) *AssignLeniPrivateIpAddressRequest {
	s.Description = &v
	return s
}

func (s *AssignLeniPrivateIpAddressRequest) SetElasticNetworkInterfaceId(v string) *AssignLeniPrivateIpAddressRequest {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *AssignLeniPrivateIpAddressRequest) SetPrivateIpAddress(v string) *AssignLeniPrivateIpAddressRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *AssignLeniPrivateIpAddressRequest) SetRegionId(v string) *AssignLeniPrivateIpAddressRequest {
	s.RegionId = &v
	return s
}

type AssignLeniPrivateIpAddressResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Content *AssignLeniPrivateIpAddressResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A88DFED5-24B7-5A3E-87DE-380BF06F3C90
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssignLeniPrivateIpAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssignLeniPrivateIpAddressResponseBody) GoString() string {
	return s.String()
}

func (s *AssignLeniPrivateIpAddressResponseBody) SetAccessDeniedDetail(v string) *AssignLeniPrivateIpAddressResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AssignLeniPrivateIpAddressResponseBody) SetCode(v int32) *AssignLeniPrivateIpAddressResponseBody {
	s.Code = &v
	return s
}

func (s *AssignLeniPrivateIpAddressResponseBody) SetContent(v *AssignLeniPrivateIpAddressResponseBodyContent) *AssignLeniPrivateIpAddressResponseBody {
	s.Content = v
	return s
}

func (s *AssignLeniPrivateIpAddressResponseBody) SetMessage(v string) *AssignLeniPrivateIpAddressResponseBody {
	s.Message = &v
	return s
}

func (s *AssignLeniPrivateIpAddressResponseBody) SetRequestId(v string) *AssignLeniPrivateIpAddressResponseBody {
	s.RequestId = &v
	return s
}

type AssignLeniPrivateIpAddressResponseBodyContent struct {
	// Lingjun Elastic Network Interface ID.
	//
	// example:
	//
	// leni-1234****
	ElasticNetworkInterfaceId *string `json:"ElasticNetworkInterfaceId,omitempty" xml:"ElasticNetworkInterfaceId,omitempty"`
	// Lingjun Elastic Network Interface secondary private IP unique identifier.
	//
	// example:
	//
	// sip-lzwx****
	IpName *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
}

func (s AssignLeniPrivateIpAddressResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s AssignLeniPrivateIpAddressResponseBodyContent) GoString() string {
	return s.String()
}

func (s *AssignLeniPrivateIpAddressResponseBodyContent) SetElasticNetworkInterfaceId(v string) *AssignLeniPrivateIpAddressResponseBodyContent {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *AssignLeniPrivateIpAddressResponseBodyContent) SetIpName(v string) *AssignLeniPrivateIpAddressResponseBodyContent {
	s.IpName = &v
	return s
}

type AssignLeniPrivateIpAddressResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssignLeniPrivateIpAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssignLeniPrivateIpAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s AssignLeniPrivateIpAddressResponse) GoString() string {
	return s.String()
}

func (s *AssignLeniPrivateIpAddressResponse) SetHeaders(v map[string]*string) *AssignLeniPrivateIpAddressResponse {
	s.Headers = v
	return s
}

func (s *AssignLeniPrivateIpAddressResponse) SetStatusCode(v int32) *AssignLeniPrivateIpAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *AssignLeniPrivateIpAddressResponse) SetBody(v *AssignLeniPrivateIpAddressResponseBody) *AssignLeniPrivateIpAddressResponse {
	s.Body = v
	return s
}

type AssignPrivateIpAddressRequest struct {
	// Specifies whether to assign a mac address.
	//
	// example:
	//
	// true
	AssignMac *bool `json:"AssignMac,omitempty" xml:"AssignMac,omitempty"`
	// By default, popApi is not ignored and idempotent
	//
	// example:
	//
	// 3fd79d62-ab1d-11ec-9a53-0242ac110004
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the variable.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the network interface controller.
	//
	// This parameter is required.
	//
	// example:
	//
	// lni-bp18exxqa2rvfn45e5pz
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The secondary private IP address.
	//
	// example:
	//
	// 10.0.6.194
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The default value is false. If you set the value to true, the secondary private IP address application process can be accelerated.
	//
	// >  For more information, submit a ticket.
	//
	// example:
	//
	// false
	SkipConfig *bool `json:"SkipConfig,omitempty" xml:"SkipConfig,omitempty"`
	// It belongs to the Lingjun subnet.
	//
	// This parameter is required.
	//
	// example:
	//
	// subnet-f3zfzmnc
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
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

func (s *AssignPrivateIpAddressRequest) SetClientToken(v string) *AssignPrivateIpAddressRequest {
	s.ClientToken = &v
	return s
}

func (s *AssignPrivateIpAddressRequest) SetDescription(v string) *AssignPrivateIpAddressRequest {
	s.Description = &v
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Content *AssignPrivateIpAddressResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AC8C713A-A9F4-5984-A5E1-76496DF35153
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssignPrivateIpAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssignPrivateIpAddressResponseBody) GoString() string {
	return s.String()
}

func (s *AssignPrivateIpAddressResponseBody) SetAccessDeniedDetail(v string) *AssignPrivateIpAddressResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// The unique IP identifier.
	//
	// example:
	//
	// sip-8exxqa2r
	IpName *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
	// Lingjun network interface controller ID.
	//
	// example:
	//
	// lni-bp18exxqa2rvfn45e5pz
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssignPrivateIpAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type AssociateVpdCidrBlockRequest struct {
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The additional CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// 172.16.0.0/12
	SecondaryCidrBlock *string `json:"SecondaryCidrBlock,omitempty" xml:"SecondaryCidrBlock,omitempty"`
	// The ID of the Lingjun CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpd-omqutbff
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
}

func (s AssociateVpdCidrBlockRequest) String() string {
	return tea.Prettify(s)
}

func (s AssociateVpdCidrBlockRequest) GoString() string {
	return s.String()
}

func (s *AssociateVpdCidrBlockRequest) SetRegionId(v string) *AssociateVpdCidrBlockRequest {
	s.RegionId = &v
	return s
}

func (s *AssociateVpdCidrBlockRequest) SetSecondaryCidrBlock(v string) *AssociateVpdCidrBlockRequest {
	s.SecondaryCidrBlock = &v
	return s
}

func (s *AssociateVpdCidrBlockRequest) SetVpdId(v string) *AssociateVpdCidrBlockRequest {
	s.VpdId = &v
	return s
}

type AssociateVpdCidrBlockResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Content *AssociateVpdCidrBlockResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9C50C9CD-E799-54DA-BA7A-1FAF3DF80857
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateVpdCidrBlockResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssociateVpdCidrBlockResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateVpdCidrBlockResponseBody) SetAccessDeniedDetail(v string) *AssociateVpdCidrBlockResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AssociateVpdCidrBlockResponseBody) SetCode(v int32) *AssociateVpdCidrBlockResponseBody {
	s.Code = &v
	return s
}

func (s *AssociateVpdCidrBlockResponseBody) SetContent(v *AssociateVpdCidrBlockResponseBodyContent) *AssociateVpdCidrBlockResponseBody {
	s.Content = v
	return s
}

func (s *AssociateVpdCidrBlockResponseBody) SetMessage(v string) *AssociateVpdCidrBlockResponseBody {
	s.Message = &v
	return s
}

func (s *AssociateVpdCidrBlockResponseBody) SetRequestId(v string) *AssociateVpdCidrBlockResponseBody {
	s.RequestId = &v
	return s
}

type AssociateVpdCidrBlockResponseBodyContent struct {
	// The ID of the Lingjun CIDR block.
	//
	// example:
	//
	// vpd-eoiy88ju
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
}

func (s AssociateVpdCidrBlockResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s AssociateVpdCidrBlockResponseBodyContent) GoString() string {
	return s.String()
}

func (s *AssociateVpdCidrBlockResponseBodyContent) SetVpdId(v string) *AssociateVpdCidrBlockResponseBodyContent {
	s.VpdId = &v
	return s
}

type AssociateVpdCidrBlockResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateVpdCidrBlockResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateVpdCidrBlockResponse) String() string {
	return tea.Prettify(s)
}

func (s AssociateVpdCidrBlockResponse) GoString() string {
	return s.String()
}

func (s *AssociateVpdCidrBlockResponse) SetHeaders(v map[string]*string) *AssociateVpdCidrBlockResponse {
	s.Headers = v
	return s
}

func (s *AssociateVpdCidrBlockResponse) SetStatusCode(v int32) *AssociateVpdCidrBlockResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateVpdCidrBlockResponse) SetBody(v *AssociateVpdCidrBlockResponseBody) *AssociateVpdCidrBlockResponse {
	s.Body = v
	return s
}

type AttachElasticNetworkInterfaceRequest struct {
	// The ID of the ENI.
	//
	// This parameter is required.
	//
	// example:
	//
	// leni-1234****
	ElasticNetworkInterfaceId *string `json:"ElasticNetworkInterfaceId,omitempty" xml:"ElasticNetworkInterfaceId,omitempty"`
	// The ID of the node.
	//
	// This parameter is required.
	//
	// example:
	//
	// e01-cn-lbj3aej****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The region ID of the disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AttachElasticNetworkInterfaceRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachElasticNetworkInterfaceRequest) GoString() string {
	return s.String()
}

func (s *AttachElasticNetworkInterfaceRequest) SetElasticNetworkInterfaceId(v string) *AttachElasticNetworkInterfaceRequest {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *AttachElasticNetworkInterfaceRequest) SetNodeId(v string) *AttachElasticNetworkInterfaceRequest {
	s.NodeId = &v
	return s
}

func (s *AttachElasticNetworkInterfaceRequest) SetRegionId(v string) *AttachElasticNetworkInterfaceRequest {
	s.RegionId = &v
	return s
}

type AttachElasticNetworkInterfaceResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Response body
	//
	// example:
	//
	// []
	Content interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The message returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 039C3C3A-3C37-5672-80D5-D8CD48C676D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachElasticNetworkInterfaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachElasticNetworkInterfaceResponseBody) GoString() string {
	return s.String()
}

func (s *AttachElasticNetworkInterfaceResponseBody) SetAccessDeniedDetail(v string) *AttachElasticNetworkInterfaceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AttachElasticNetworkInterfaceResponseBody) SetCode(v int32) *AttachElasticNetworkInterfaceResponseBody {
	s.Code = &v
	return s
}

func (s *AttachElasticNetworkInterfaceResponseBody) SetContent(v interface{}) *AttachElasticNetworkInterfaceResponseBody {
	s.Content = v
	return s
}

func (s *AttachElasticNetworkInterfaceResponseBody) SetMessage(v string) *AttachElasticNetworkInterfaceResponseBody {
	s.Message = &v
	return s
}

func (s *AttachElasticNetworkInterfaceResponseBody) SetRequestId(v string) *AttachElasticNetworkInterfaceResponseBody {
	s.RequestId = &v
	return s
}

type AttachElasticNetworkInterfaceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachElasticNetworkInterfaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachElasticNetworkInterfaceResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachElasticNetworkInterfaceResponse) GoString() string {
	return s.String()
}

func (s *AttachElasticNetworkInterfaceResponse) SetHeaders(v map[string]*string) *AttachElasticNetworkInterfaceResponse {
	s.Headers = v
	return s
}

func (s *AttachElasticNetworkInterfaceResponse) SetStatusCode(v int32) *AttachElasticNetworkInterfaceResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachElasticNetworkInterfaceResponse) SetBody(v *AttachElasticNetworkInterfaceResponseBody) *AttachElasticNetworkInterfaceResponse {
	s.Body = v
	return s
}

type CreateElasticNetworkInterfaceRequest struct {
	// The POP API is not ignored by default and is used for idempotence.
	//
	// example:
	//
	// 3fd79d62-ab1d-11ec-9a53-0242ac110004
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the response code.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Whether to enable the jumbo frame capability
	//
	// example:
	//
	// True
	EnableJumboFrame *bool `json:"EnableJumboFrame,omitempty" xml:"EnableJumboFrame,omitempty"`
	// The ID of the Lingjun node.
	//
	// example:
	//
	// e01-cn-lbj3aej****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-wz9fj2s3o21nw2****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-t4nahb0pxck****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC.
	//
	// >  If the NodeId parameter is empty, the VpcId parameter is required. If the NodeId parameter is not empty, the VpcId parameter is optional.
	//
	// example:
	//
	// vpc-uf6aa4ddo97fr****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateElasticNetworkInterfaceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateElasticNetworkInterfaceRequest) GoString() string {
	return s.String()
}

func (s *CreateElasticNetworkInterfaceRequest) SetClientToken(v string) *CreateElasticNetworkInterfaceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateElasticNetworkInterfaceRequest) SetDescription(v string) *CreateElasticNetworkInterfaceRequest {
	s.Description = &v
	return s
}

func (s *CreateElasticNetworkInterfaceRequest) SetEnableJumboFrame(v bool) *CreateElasticNetworkInterfaceRequest {
	s.EnableJumboFrame = &v
	return s
}

func (s *CreateElasticNetworkInterfaceRequest) SetNodeId(v string) *CreateElasticNetworkInterfaceRequest {
	s.NodeId = &v
	return s
}

func (s *CreateElasticNetworkInterfaceRequest) SetRegionId(v string) *CreateElasticNetworkInterfaceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateElasticNetworkInterfaceRequest) SetSecurityGroupId(v string) *CreateElasticNetworkInterfaceRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateElasticNetworkInterfaceRequest) SetVSwitchId(v string) *CreateElasticNetworkInterfaceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateElasticNetworkInterfaceRequest) SetVpcId(v string) *CreateElasticNetworkInterfaceRequest {
	s.VpcId = &v
	return s
}

func (s *CreateElasticNetworkInterfaceRequest) SetZoneId(v string) *CreateElasticNetworkInterfaceRequest {
	s.ZoneId = &v
	return s
}

type CreateElasticNetworkInterfaceResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Content *CreateElasticNetworkInterfaceResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A88DFED5-24B7-5A3E-87DE-380BF06F3C90
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateElasticNetworkInterfaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateElasticNetworkInterfaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateElasticNetworkInterfaceResponseBody) SetAccessDeniedDetail(v string) *CreateElasticNetworkInterfaceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateElasticNetworkInterfaceResponseBody) SetCode(v int32) *CreateElasticNetworkInterfaceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateElasticNetworkInterfaceResponseBody) SetContent(v *CreateElasticNetworkInterfaceResponseBodyContent) *CreateElasticNetworkInterfaceResponseBody {
	s.Content = v
	return s
}

func (s *CreateElasticNetworkInterfaceResponseBody) SetMessage(v string) *CreateElasticNetworkInterfaceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateElasticNetworkInterfaceResponseBody) SetRequestId(v string) *CreateElasticNetworkInterfaceResponseBody {
	s.RequestId = &v
	return s
}

type CreateElasticNetworkInterfaceResponseBodyContent struct {
	// Lingjun Elastic Network Interface ID.
	//
	// example:
	//
	// leni-1fejojjo****
	ElasticNetworkInterfaceId *string `json:"ElasticNetworkInterfaceId,omitempty" xml:"ElasticNetworkInterfaceId,omitempty"`
	// The ID of the Lingjun node.
	//
	// example:
	//
	// e01-cn-lbj3aej****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s CreateElasticNetworkInterfaceResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s CreateElasticNetworkInterfaceResponseBodyContent) GoString() string {
	return s.String()
}

func (s *CreateElasticNetworkInterfaceResponseBodyContent) SetElasticNetworkInterfaceId(v string) *CreateElasticNetworkInterfaceResponseBodyContent {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *CreateElasticNetworkInterfaceResponseBodyContent) SetNodeId(v string) *CreateElasticNetworkInterfaceResponseBodyContent {
	s.NodeId = &v
	return s
}

type CreateElasticNetworkInterfaceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateElasticNetworkInterfaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateElasticNetworkInterfaceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateElasticNetworkInterfaceResponse) GoString() string {
	return s.String()
}

func (s *CreateElasticNetworkInterfaceResponse) SetHeaders(v map[string]*string) *CreateElasticNetworkInterfaceResponse {
	s.Headers = v
	return s
}

func (s *CreateElasticNetworkInterfaceResponse) SetStatusCode(v int32) *CreateElasticNetworkInterfaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateElasticNetworkInterfaceResponse) SetBody(v *CreateElasticNetworkInterfaceResponseBody) *CreateElasticNetworkInterfaceResponse {
	s.Body = v
	return s
}

type CreateErRequest struct {
	// The description of the document.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Lingjun HUB Name
	//
	// This parameter is required.
	//
	// example:
	//
	// er-wulanchabu-main
	ErName *string `json:"ErName,omitempty" xml:"ErName,omitempty"`
	// Primary Zone
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu-b
	MasterZoneId *string `json:"MasterZoneId,omitempty" xml:"MasterZoneId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-acfmyuzlx2iihcy
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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

func (s *CreateErRequest) SetResourceGroupId(v string) *CreateErRequest {
	s.ResourceGroupId = &v
	return s
}

type CreateErResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Content *CreateErResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// AC8C713A-A9F4-5984-A5E1-76496DF35153
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateErResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateErResponseBody) GoString() string {
	return s.String()
}

func (s *CreateErResponseBody) SetAccessDeniedDetail(v string) *CreateErResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// Lingjun HUB ID
	//
	// example:
	//
	// er-aueyxxsy
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
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateErResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// Indicates whether to automatically receive all routes from all instances under the Lingjun HUB. Valid values:
	//
	// 	- **true**: received automatically.
	//
	// 	- **false**: Not received.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	AutoReceiveAllRoute *bool `json:"AutoReceiveAllRoute,omitempty" xml:"AutoReceiveAllRoute,omitempty"`
	// The name of the network instance connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// er-attachemnt-vpd-xksd2obl
	ErAttachmentName *string `json:"ErAttachmentName,omitempty" xml:"ErAttachmentName,omitempty"`
	// Lingjun HUB ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// The ID of the network instance. Valid values: **VPD*	- and **VCC**.
	//
	// For more information, see [What is Lingjun?](https://help.aliyun.com/document_detail/444430.html)
	//
	// You can query **Lingjun CIDR Block*	- and **Lingjun Connection*	- by [ListVpds](https://help.aliyun.com/document_detail/2331077.html) and [ListVccs](https://help.aliyun.com/document_detail/2399526.html?) respectively.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpd-xksd2obl
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The category of the instance. Valid values:
	//
	// 	- **VPD**: indicates the Lingjun CIDR block.
	//
	// 	- **VCC**: indicates a Lingjun connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// VPD
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the tenant to which the resource belongs. This parameter is required for cross-account resources.
	//
	// example:
	//
	// 1620939556166277
	ResourceTenantId *string `json:"ResourceTenantId,omitempty" xml:"ResourceTenantId,omitempty"`
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Content *CreateErAttachmentResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is displayed.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DBAD15D6-3F47-5B36-8A92-57C2919D13D0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateErAttachmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateErAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateErAttachmentResponseBody) SetAccessDeniedDetail(v string) *CreateErAttachmentResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// The ID of the network connection instance.
	//
	// example:
	//
	// er-attachment-ggjbfhqv
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateErAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// Policy description
	//
	// example:
	//
	// terraform-example
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Destination CIDR Block
	//
	// example:
	//
	// 0.0.0.0/0
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// Lingjun HUB ID
	//
	// This parameter is required.
	//
	// example:
	//
	// er-aueyxxsy
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// The ID of the destination instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpd-xlhsvdvt
	ReceptionInstanceId *string `json:"ReceptionInstanceId,omitempty" xml:"ReceptionInstanceId,omitempty"`
	// The tenant to which the route receiving instance belongs.
	//
	// example:
	//
	// 1620939556166277
	ReceptionInstanceOwner *string `json:"ReceptionInstanceOwner,omitempty" xml:"ReceptionInstanceOwner,omitempty"`
	// The type of the destination instance. Valid values:
	//
	// 	- **VPD**: Lingjun network segment.
	//
	// 	- **VCC**: Lingjun Connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// VPD
	ReceptionInstanceType *string `json:"ReceptionInstanceType,omitempty" xml:"ReceptionInstanceType,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Policy behavior; optional values:
	//
	// 	- **permit**: Allow
	//
	// 	- **deny**: Rejected
	//
	// This parameter is required.
	//
	// example:
	//
	// permit
	RouteMapAction *string `json:"RouteMapAction,omitempty" xml:"RouteMapAction,omitempty"`
	// The ID of the policy.
	//
	// A smaller sequence number indicates a lower priority. When a route is matched, a policy with a higher priority is preferentially matched.
	//
	// **Valid values: 1001 to 2000**
	//
	// This parameter is required.
	//
	// example:
	//
	// 1001
	RouteMapNum *int32 `json:"RouteMapNum,omitempty" xml:"RouteMapNum,omitempty"`
	// The ID of the source instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpd-xlsjsdvt
	TransmissionInstanceId *string `json:"TransmissionInstanceId,omitempty" xml:"TransmissionInstanceId,omitempty"`
	// The tenant to which the route publish instance belongs
	//
	// example:
	//
	// 1620939556166277
	TransmissionInstanceOwner *string `json:"TransmissionInstanceOwner,omitempty" xml:"TransmissionInstanceOwner,omitempty"`
	// The type of the source instance. Valid values:
	//
	// 	- **VPD**: Lingjun network segment.
	//
	// 	- **VCC**: Lingjun Connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// VPD
	TransmissionInstanceType *string `json:"TransmissionInstanceType,omitempty" xml:"TransmissionInstanceType,omitempty"`
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Content *CreateErRouteMapResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 039C3C3A-3C37-5672-80D5-D8CD48C676D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateErRouteMapResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateErRouteMapResponseBody) GoString() string {
	return s.String()
}

func (s *CreateErRouteMapResponseBody) SetAccessDeniedDetail(v string) *CreateErRouteMapResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// routing policy ID
	//
	// example:
	//
	// er-rmap-uwglhzom
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateErRouteMapResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The CIDR block of the Subnet.
	//
	// 	- The network segment of the subnet must be a proper subset of the network segment of Lingjun to which it belongs, and the mask must be between 16 bits and 29 bits, which can provide 8 to 65536 addresses. For example, the CIDR block of the Lingjun CIDR block is 192.168.0.0/16, and the CIDR blocks of the subnets under the Lingjun CIDR block are 192.168.0.0/17 to 192.168.0.0/29.
	//
	// 	- The first and last three IP addresses of each subnet segment are reserved by the system. For example, the CIDR blocks of the subnet are 192.168.1.0/24,192.168.1.0, 192.168.1.253, 192.168.1.254, and 192.168.1.255.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.0.0.0/16
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Lingjun subnet instance name
	//
	// This parameter is required.
	//
	// example:
	//
	// subnet-1
	SubnetName *string `json:"SubnetName,omitempty" xml:"SubnetName,omitempty"`
	// The tag information.
	//
	// You can specify up to 20 tags.
	Tag []*CreateSubnetRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Lingjun Subnet Usage Type; optional; optional. Valid values:
	//
	// 	- **If you do not set this field for a common type**
	//
	// 	- **OOB*	- :OOB type
	//
	// 	- **LB**: LB type
	//
	// example:
	//
	// OOB
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the Lingjun CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpd-xcuhjyrj
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// The zone ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
	// The tag key of the VPN attachment.
	//
	// You cannot specify an empty string as a tag key. It can be up to 64 characters in length and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// You can specify at most 20 tag keys in each call.
	//
	// example:
	//
	// tag-subnet
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the VPN connection.
	//
	// The tag value can be empty or a string of up to 128 characters. It cannot start with aliyun or acs:, and cannot contain http:// or https://.
	//
	// Each key-value pair must be unique. You can specify values for at most 20 tag keys in each call.
	//
	// example:
	//
	// subnet-tag-1
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Content *CreateSubnetResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// A88DFED5-24B7-5A3E-87DE-380BF06F3C90
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSubnetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSubnetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSubnetResponseBody) SetAccessDeniedDetail(v string) *CreateSubnetResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// Lingjun subnet instance ID
	//
	// example:
	//
	// subnet-yuvn29bn
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSubnetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// Enabled access to cloud services. Optional values:
	//
	// 	- **true**: Enable access to cloud services
	//
	// 	- **false**: Do not enable access to cloud services
	//
	// example:
	//
	// true
	AccessCouldService *bool `json:"AccessCouldService,omitempty" xml:"AccessCouldService,omitempty"`
	// The bandwidth. Unit: Mbit /s. The minimum value is 1000, representing 1Gbps bandwidth; the maximum value is 400000, representing 400Gbps bandwidth.
	//
	// >  1Gbps = 1000Mbps
	//
	// example:
	//
	// 1000
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// bgp as number
	//
	// example:
	//
	// bgpAsn
	BgpAsn *int64 `json:"BgpAsn,omitempty" xml:"BgpAsn,omitempty"`
	// Internet segment, on-premises input, off-premises default
	//
	// example:
	//
	// 10.0.0.0/24
	BgpCidr *string `json:"BgpCidr,omitempty" xml:"BgpCidr,omitempty"`
	// CEN Instance ID
	//
	// example:
	//
	// cen-bkiw0x1347roekr7f2
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// Account to which cen belongs
	//
	// example:
	//
	// 1511928242963727
	CenOwnerId *string `json:"CenOwnerId,omitempty" xml:"CenOwnerId,omitempty"`
	// The connection mode. Valid values:
	//
	// 	- **VPC**
	//
	// 	- **CEN (CENTR)**
	//
	// example:
	//
	// CENTR
	ConnectionType *string `json:"ConnectionType,omitempty" xml:"ConnectionType,omitempty"`
	// The description of the document.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// For more information about resource groups, see [Resource groups](https://help.aliyun.com/document_detail/94475.htm?spm=a2c4g.11186623.0.0.29e15d7akXhpuu).
	//
	// example:
	//
	// rg-aeky5f3qx6ceapq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tag information.
	//
	// You can specify up to 20 tags.
	Tag []*CreateVccRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the vSwitch. [Virtual Private Cloud VSwitch](https://help.aliyun.com/document_detail/100380.html).
	//
	// You can call the [DescribeVSwitches](https://help.aliyun.com/document_detail/35748.html) operation to query created vSwitches.
	//
	// example:
	//
	// vsw-t4nahb0pxckgktx1kot8q
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the Lingjun connection instance.
	//
	// example:
	//
	// vcc-cn-zvp2w222001
	VccId *string `json:"VccId,omitempty" xml:"VccId,omitempty"`
	// Lingjun Connection Name
	//
	// example:
	//
	// test
	VccName *string `json:"VccName,omitempty" xml:"VccName,omitempty"`
	// Virtual Private Cloud IDs; [What is Virtual Private Cloud](https://help.aliyun.com/document_detail/34217.html)
	//
	// You can call the [DescribeVpcs](https://help.aliyun.com/document_detail/35739.html#demo-0) operation to query the specified VPC.
	//
	// example:
	//
	// vpc-uf6aa4ddo97frj22tgp52
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// Lingjun CIDR block instance ID
	//
	// example:
	//
	// vpd-t2jseldp
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// The zone ID of the disk.
	//
	// example:
	//
	// cn-wulanchabu-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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

func (s *CreateVccRequest) SetBgpAsn(v int64) *CreateVccRequest {
	s.BgpAsn = &v
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

func (s *CreateVccRequest) SetCenOwnerId(v string) *CreateVccRequest {
	s.CenOwnerId = &v
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
	// The tag key of the VPN attachment.
	//
	// You cannot specify an empty string as a tag key. It can be up to 64 characters in length and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// You can specify at most 20 tag keys in each call.
	//
	// example:
	//
	// tag-vcc
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the VPN connection.
	//
	// The tag value can be empty or a string of up to 128 characters. It cannot start with aliyun or acs:, and cannot contain http:// or https://.
	//
	// Each key-value pair must be unique. You can specify values for at most 20 tag keys in each call.
	//
	// example:
	//
	// vcc-group-1
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Content *CreateVccResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// response message, if the success request is
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// 039C3C3A-3C37-5672-80D5-D8CD48C676D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVccResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVccResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVccResponseBody) SetAccessDeniedDetail(v string) *CreateVccResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// The ID of the Lingjun connection instance.
	//
	// example:
	//
	// vcc-cn-zvp2w222001
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVccResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// Lingjun HUB ID
	//
	// This parameter is required.
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// Authorized Tenant ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 1620939556166277
	GrantTenantId *string `json:"GrantTenantId,omitempty" xml:"GrantTenantId,omitempty"`
	// Network Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// vcc-8rgvqazb
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	//
	// example:
	//
	// {}
	Content *CreateVccGrantRuleResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// AC8C713A-A9F4-5984-A5E1-76496DF35153
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVccGrantRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVccGrantRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVccGrantRuleResponseBody) SetAccessDeniedDetail(v string) *CreateVccGrantRuleResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// Authorized resource primary key ID
	//
	// example:
	//
	// grant-rule-8rgvqazb
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVccGrantRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// Destination CIDR block
	//
	// example:
	//
	// 192.168.98.112/28
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Lingjun connection instance.
	//
	// example:
	//
	// vcc-cn-zvp2w222001
	VccId *string `json:"VccId,omitempty" xml:"VccId,omitempty"`
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Content *CreateVccRouteEntryResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 039C3C3A-3C37-5672-80D5-D8CD48C676D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVccRouteEntryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVccRouteEntryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVccRouteEntryResponseBody) SetAccessDeniedDetail(v string) *CreateVccRouteEntryResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// The ID of the route entry.
	//
	// example:
	//
	// vcc-rte-5cey1sap
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVccRouteEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The CIDR block of the VPD.
	//
	// 	- We recommend that you use an RFC private endpoint as the Lingjun CIDR block, such as 10.0.0.0/8,172.16.0.0/12,192.168.0.0/16. In scenarios where the Doringjun CIDR block is connected to each other or where the Lingjun CIDR block is connected to a VPC, make sure that the addresses do not conflict with each other.
	//
	// 	- You can also use a custom CIDR block other than 100.64.0.0/10, 224.0.0.0/4, 127.0.0.0/8, or 169.254.0.0/16 and their subnets as the primary IPv4 CIDR block of the VPD.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.0.0.0/8
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The region ID of the disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// For more information about resource groups, see [Resource groups](https://help.aliyun.com/document_detail/94475.htm?spm=a2c4g.11186623.0.0.29e15d7akXhpuu).
	//
	// example:
	//
	// rg-aek2l4sq6l7unhi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Lingjun subnet information list
	Subnets []*CreateVpdRequestSubnets `json:"Subnets,omitempty" xml:"Subnets,omitempty" type:"Repeated"`
	// A tag.
	Tag []*CreateVpdRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Lingjun CIDR block instance name
	//
	// This parameter is required.
	//
	// example:
	//
	// vpd-1
	VpdName *string `json:"VpdName,omitempty" xml:"VpdName,omitempty"`
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
	// The CIDR block of the Subnet.
	//
	// 	- The network segment of the subnet must be a proper subset of the network segment of Lingjun to which it belongs, and the mask must be between 16 bits and 29 bits, which can provide 8 to 65536 addresses. For example, the CIDR block of the Lingjun CIDR block is 192.168.0.0/16, and the CIDR blocks of the subnets under the Lingjun CIDR block are 192.168.0.0/17 to 192.168.0.0/29.
	//
	// 	- The first and last three IP addresses of each subnet segment are reserved by the system. For example, the CIDR blocks of the subnet are 192.168.1.0/24,192.168.1.0, 192.168.1.253, 192.168.1.254, and 192.168.1.255.
	//
	// example:
	//
	// 10.1.0.0/16
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The region in which the instance resides.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Lingjun subnet instance name
	//
	// example:
	//
	// subnet-1
	SubnetName *string `json:"SubnetName,omitempty" xml:"SubnetName,omitempty"`
	// Lingjun Subnet Usage Type; optional; optional. Valid values:
	//
	// 	- **Generic type is not specified**.
	//
	// 	- **OOB*	- :OOB type
	//
	// 	- **LB**: LB type
	//
	// example:
	//
	// OOB
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The zone ID of the disk.
	//
	// example:
	//
	// cn-wulanchabu-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
	// The tag key of the VPN attachment.
	//
	// You cannot specify an empty string as a tag key. It can be up to 64 characters in length and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// You can specify at most 20 tag keys in each call.
	//
	// example:
	//
	// vpd-wulanchabu
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the VPN connection.
	//
	// The tag value can be empty or a string of up to 128 characters. It cannot start with aliyun or acs:, and cannot contain http:// or https://.
	//
	// Each tag key corresponds to a tag value. You can enter a maximum of 20 tag values at a time.
	//
	// example:
	//
	// wulanchabu-a
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Content *CreateVpdResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DBAD15D6-3F47-5B36-8A92-57C2919D13D0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVpdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVpdResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVpdResponseBody) SetAccessDeniedDetail(v string) *CreateVpdResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// Lingjun subnet ID list
	SubnetIds []*string `json:"SubnetIds,omitempty" xml:"SubnetIds,omitempty" type:"Repeated"`
	// Lingjun CIDR block instance ID
	//
	// example:
	//
	// vpd-eoiy88ju
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVpdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// Lingjun HUB ID
	//
	// This parameter is required.
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// Authorized Tenant ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 1013666993027780
	GrantTenantId *string `json:"GrantTenantId,omitempty" xml:"GrantTenantId,omitempty"`
	// Network Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// vpd-8rgvqazb
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	//
	// example:
	//
	// {}
	Content *CreateVpdGrantRuleResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// DBAD15D6-3F47-5B36-8A92-57C2919D13D0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVpdGrantRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVpdGrantRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVpdGrantRuleResponseBody) SetAccessDeniedDetail(v string) *CreateVpdGrantRuleResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// Authorized resource primary key ID
	//
	// example:
	//
	// grant-rule-hnevjkmw
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVpdGrantRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DeleteElasticNetworkInterfaceRequest struct {
	// By default, popApi is not ignored and idempotent
	//
	// example:
	//
	// 141cccd6-dfbd-11ec-b8e8-0242ac110003
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Lingjun Elastic Network Interface ID
	//
	// This parameter is required.
	//
	// example:
	//
	// leni-1234****
	ElasticNetworkInterfaceId *string `json:"ElasticNetworkInterfaceId,omitempty" xml:"ElasticNetworkInterfaceId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteElasticNetworkInterfaceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteElasticNetworkInterfaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteElasticNetworkInterfaceRequest) SetClientToken(v string) *DeleteElasticNetworkInterfaceRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteElasticNetworkInterfaceRequest) SetElasticNetworkInterfaceId(v string) *DeleteElasticNetworkInterfaceRequest {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *DeleteElasticNetworkInterfaceRequest) SetRegionId(v string) *DeleteElasticNetworkInterfaceRequest {
	s.RegionId = &v
	return s
}

type DeleteElasticNetworkInterfaceResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Content *DeleteElasticNetworkInterfaceResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The return message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// BDBCC783-84CA-5733-8EEA-645C88B9009C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteElasticNetworkInterfaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteElasticNetworkInterfaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteElasticNetworkInterfaceResponseBody) SetAccessDeniedDetail(v string) *DeleteElasticNetworkInterfaceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteElasticNetworkInterfaceResponseBody) SetCode(v int32) *DeleteElasticNetworkInterfaceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteElasticNetworkInterfaceResponseBody) SetContent(v *DeleteElasticNetworkInterfaceResponseBodyContent) *DeleteElasticNetworkInterfaceResponseBody {
	s.Content = v
	return s
}

func (s *DeleteElasticNetworkInterfaceResponseBody) SetMessage(v string) *DeleteElasticNetworkInterfaceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteElasticNetworkInterfaceResponseBody) SetRequestId(v string) *DeleteElasticNetworkInterfaceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteElasticNetworkInterfaceResponseBodyContent struct {
	// Lingjun Elastic Network Interface ID
	//
	// example:
	//
	// leni-1234****
	ElasticNetworkInterfaceId *string `json:"ElasticNetworkInterfaceId,omitempty" xml:"ElasticNetworkInterfaceId,omitempty"`
	// Lingjun Node ID
	//
	// example:
	//
	// e01-cn-lbj3aej****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s DeleteElasticNetworkInterfaceResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s DeleteElasticNetworkInterfaceResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DeleteElasticNetworkInterfaceResponseBodyContent) SetElasticNetworkInterfaceId(v string) *DeleteElasticNetworkInterfaceResponseBodyContent {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *DeleteElasticNetworkInterfaceResponseBodyContent) SetNodeId(v string) *DeleteElasticNetworkInterfaceResponseBodyContent {
	s.NodeId = &v
	return s
}

type DeleteElasticNetworkInterfaceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteElasticNetworkInterfaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteElasticNetworkInterfaceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteElasticNetworkInterfaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteElasticNetworkInterfaceResponse) SetHeaders(v map[string]*string) *DeleteElasticNetworkInterfaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteElasticNetworkInterfaceResponse) SetStatusCode(v int32) *DeleteElasticNetworkInterfaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteElasticNetworkInterfaceResponse) SetBody(v *DeleteElasticNetworkInterfaceResponseBody) *DeleteElasticNetworkInterfaceResponse {
	s.Body = v
	return s
}

type DeleteErRequest struct {
	// Lingjun HUB Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Response body
	//
	// example:
	//
	// {}
	Content interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// 9C50C9CD-E799-54DA-BA7A-1FAF3DF80857
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteErResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteErResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteErResponseBody) SetAccessDeniedDetail(v string) *DeleteErResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteErResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the network connection instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// er-attachment-5n3nsmvl
	ErAttachmentId *string `json:"ErAttachmentId,omitempty" xml:"ErAttachmentId,omitempty"`
	// Lingjun HUB Id
	//
	// This parameter is required.
	//
	// example:
	//
	// er-opy1wrfv
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response content. If a resource has dependent resources, the existing dependent resources are returned.
	//
	// example:
	//
	// {
	//
	//     "ER_RMAP": [
	//
	//         {
	//
	//             "erId": "er-opy1wrfv",
	//
	//             "destinationCidrBlock": "0.0.0.0/0",
	//
	//             "regionId": "cn-wulanchabu",
	//
	//             "routeMapNum": 3000,
	//
	//             "erRouteMapId": "er-rmap-v5lfhmvm",
	//
	//             "action": "permit",
	//
	//             "status": "Available"
	//
	//         },
	//
	//         {
	//
	//             "erId": "er-opy1wrfv",
	//
	//             "destinationCidrBlock": "0.0.0.0/0",
	//
	//             "regionId": "cn-wulanchabu",
	//
	//             "routeMapNum": 3000,
	//
	//             "erRouteMapId": "er-rmap-of3r0ndh",
	//
	//             "action": "permit",
	//
	//             "status": "Available"
	//
	//         }
	//
	//     ]
	//
	// }
	Content interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// response message, if the success request is
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// A88DFED5-24B7-5A3E-87DE-380BF06F3C90
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteErAttachmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteErAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteErAttachmentResponseBody) SetAccessDeniedDetail(v string) *DeleteErAttachmentResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteErAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// Lingjun HUB ID
	//
	// This parameter is required.
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// routing policy Instance ID List
	//
	// This parameter is required.
	ErRouteMapIds []*string `json:"ErRouteMapIds,omitempty" xml:"ErRouteMapIds,omitempty" type:"Repeated"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Response body
	//
	// example:
	//
	// {}
	Content interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AC8C713A-A9F4-5984-A5E1-76496DF35153
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteErRouteMapResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteErRouteMapResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteErRouteMapResponseBody) SetAccessDeniedDetail(v string) *DeleteErRouteMapResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteErRouteMapResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The region ID of the disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Lingjun subnet ID
	//
	// This parameter is required.
	//
	// example:
	//
	// subnet-f3zfzmnc
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	// Lingjun CIDR block ID
	//
	// This parameter is required.
	//
	// example:
	//
	// vpd-iv2zm1qf
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// Zone
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Response content (if the resource has dependent resources, the existing dependent resources will be returned)
	//
	// example:
	//
	// {
	//
	//       "nc": [
	//
	//             {}
	//
	//       ]
	//
	// }
	Content interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// A56F7D3C-8850-5AF4-A342-2D71C9A9D1CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSubnetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubnetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSubnetResponseBody) SetAccessDeniedDetail(v string) *DeleteSubnetResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSubnetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// Lingjun HUB ID
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// Authorization Entry ID
	//
	// This parameter is required.
	//
	// example:
	//
	// grant-rule-jaj34d75h01
	GrantRuleId *string `json:"GrantRuleId,omitempty" xml:"GrantRuleId,omitempty"`
	// Network Instance ID
	//
	// example:
	//
	// vcc-cn-jaj34d75h01
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Response body
	//
	// example:
	//
	// {}
	Content interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// BDBCC783-84CA-5733-8EEA-645C88B9009C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVccGrantRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVccGrantRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVccGrantRuleResponseBody) SetAccessDeniedDetail(v string) *DeleteVccGrantRuleResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVccGrantRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// Destination CIDR block
	//
	// example:
	//
	// 172.16.199.128/25
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Lingjun connection instance.
	//
	// example:
	//
	// vcc-cn-zvp2w222001
	VccId *string `json:"VccId,omitempty" xml:"VccId,omitempty"`
	// The ID of the route entry.
	//
	// example:
	//
	// vcc-rte-5cey1sap
	VccRouteEntryId *string `json:"VccRouteEntryId,omitempty" xml:"VccRouteEntryId,omitempty"`
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Response body
	//
	// example:
	//
	// {}
	Content interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// 0901F411-28FA-5B9C-BAEE-7776463FF0DC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVccRouteEntryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVccRouteEntryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVccRouteEntryResponseBody) SetAccessDeniedDetail(v string) *DeleteVccRouteEntryResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVccRouteEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the region in which the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Lingjun CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpd-zr0farea
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters. (If a dependent resource exists, the existing dependent resource is returned.)
	//
	// example:
	//
	// {
	//
	//       "subnet": [
	//
	//             {
	//
	//                   "tenantId": "1620939556166277",
	//
	//                   "regionId": "cn-wulanchabu",
	//
	//                   "zoneId": "cn",
	//
	//                   "type": null,
	//
	//                   "subnetId": "subnet-zqebaxa0",
	//
	//                   "name": "lql_testVPD"
	//
	//             }
	//
	//       ],
	//
	//       "nc": [
	//
	//             {}
	//
	//       ]
	//
	// }
	Content interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The response message.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BDBCC783-84CA-5733-8EEA-645C88B9009C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVpdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpdResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVpdResponseBody) SetAccessDeniedDetail(v string) *DeleteVpdResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVpdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// Lingjun HUB ID
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// Authorization Entry ID
	//
	// This parameter is required.
	//
	// example:
	//
	// grant-rule-9rgxqazb
	GrantRuleId *string `json:"GrantRuleId,omitempty" xml:"GrantRuleId,omitempty"`
	// Authorized Tenant ID
	//
	// example:
	//
	// 1013666993027780
	GrantTenantId *string `json:"GrantTenantId,omitempty" xml:"GrantTenantId,omitempty"`
	// Network Instance ID
	//
	// example:
	//
	// vpd-8rgvqazb
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Response body
	//
	// example:
	//
	// {}
	Content interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// 0901F411-28FA-5B9C-BAEE-7776463FF0DC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVpdGrantRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpdGrantRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVpdGrantRuleResponseBody) SetAccessDeniedDetail(v string) *DeleteVpdGrantRuleResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVpdGrantRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the resource group to which the RAM instance belongs.
	//
	// example:
	//
	// rg-aeky5f3qx6ceapq
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Content *DescribeSlrResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// 9C50C9CD-E799-54DA-BA7A-1FAF3DF80857
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSlrResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlrResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlrResponseBody) SetAccessDeniedDetail(v string) *DescribeSlrResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// Whether the role exists
	//
	// example:
	//
	// true
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSlrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DetachElasticNetworkInterfaceRequest struct {
	// The ID of the ENI.
	//
	// This parameter is required.
	//
	// example:
	//
	// leni-1234****
	ElasticNetworkInterfaceId *string `json:"ElasticNetworkInterfaceId,omitempty" xml:"ElasticNetworkInterfaceId,omitempty"`
	// The ID of the node.
	//
	// This parameter is required.
	//
	// example:
	//
	// e01-cn-zxu2zp3****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The region ID of the disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DetachElasticNetworkInterfaceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachElasticNetworkInterfaceRequest) GoString() string {
	return s.String()
}

func (s *DetachElasticNetworkInterfaceRequest) SetElasticNetworkInterfaceId(v string) *DetachElasticNetworkInterfaceRequest {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *DetachElasticNetworkInterfaceRequest) SetNodeId(v string) *DetachElasticNetworkInterfaceRequest {
	s.NodeId = &v
	return s
}

func (s *DetachElasticNetworkInterfaceRequest) SetRegionId(v string) *DetachElasticNetworkInterfaceRequest {
	s.RegionId = &v
	return s
}

type DetachElasticNetworkInterfaceResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Response
	//
	// example:
	//
	// You don\\"t have the permission to do this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9C50C9CD-E799-54DA-BA7A-1FAF3DF8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachElasticNetworkInterfaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachElasticNetworkInterfaceResponseBody) GoString() string {
	return s.String()
}

func (s *DetachElasticNetworkInterfaceResponseBody) SetAccessDeniedDetail(v string) *DetachElasticNetworkInterfaceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DetachElasticNetworkInterfaceResponseBody) SetCode(v int32) *DetachElasticNetworkInterfaceResponseBody {
	s.Code = &v
	return s
}

func (s *DetachElasticNetworkInterfaceResponseBody) SetMessage(v string) *DetachElasticNetworkInterfaceResponseBody {
	s.Message = &v
	return s
}

func (s *DetachElasticNetworkInterfaceResponseBody) SetRequestId(v string) *DetachElasticNetworkInterfaceResponseBody {
	s.RequestId = &v
	return s
}

type DetachElasticNetworkInterfaceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachElasticNetworkInterfaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachElasticNetworkInterfaceResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachElasticNetworkInterfaceResponse) GoString() string {
	return s.String()
}

func (s *DetachElasticNetworkInterfaceResponse) SetHeaders(v map[string]*string) *DetachElasticNetworkInterfaceResponse {
	s.Headers = v
	return s
}

func (s *DetachElasticNetworkInterfaceResponse) SetStatusCode(v int32) *DetachElasticNetworkInterfaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachElasticNetworkInterfaceResponse) SetBody(v *DetachElasticNetworkInterfaceResponseBody) *DetachElasticNetworkInterfaceResponse {
	s.Body = v
	return s
}

type GetDestinationCidrBlockRequest struct {
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// vpd-xxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Region ID
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetDestinationCidrBlockRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDestinationCidrBlockRequest) GoString() string {
	return s.String()
}

func (s *GetDestinationCidrBlockRequest) SetInstanceId(v string) *GetDestinationCidrBlockRequest {
	s.InstanceId = &v
	return s
}

func (s *GetDestinationCidrBlockRequest) SetRegionId(v string) *GetDestinationCidrBlockRequest {
	s.RegionId = &v
	return s
}

type GetDestinationCidrBlockResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Response content
	Content *GetDestinationCidrBlockResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// Error message. (Indicates the reason for the anomaly when the instance status is abnormal.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// ID of this request
	//
	// example:
	//
	// D349EE86-AF3F-5F6C-87E2-2A08D3618350
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDestinationCidrBlockResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDestinationCidrBlockResponseBody) GoString() string {
	return s.String()
}

func (s *GetDestinationCidrBlockResponseBody) SetAccessDeniedDetail(v string) *GetDestinationCidrBlockResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetDestinationCidrBlockResponseBody) SetCode(v int32) *GetDestinationCidrBlockResponseBody {
	s.Code = &v
	return s
}

func (s *GetDestinationCidrBlockResponseBody) SetContent(v *GetDestinationCidrBlockResponseBodyContent) *GetDestinationCidrBlockResponseBody {
	s.Content = v
	return s
}

func (s *GetDestinationCidrBlockResponseBody) SetMessage(v string) *GetDestinationCidrBlockResponseBody {
	s.Message = &v
	return s
}

func (s *GetDestinationCidrBlockResponseBody) SetRequestId(v string) *GetDestinationCidrBlockResponseBody {
	s.RequestId = &v
	return s
}

type GetDestinationCidrBlockResponseBodyContent struct {
	// List of destination CIDR block information for the current network instance
	DestinationCidrBlock []*string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty" type:"Repeated"`
}

func (s GetDestinationCidrBlockResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s GetDestinationCidrBlockResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetDestinationCidrBlockResponseBodyContent) SetDestinationCidrBlock(v []*string) *GetDestinationCidrBlockResponseBodyContent {
	s.DestinationCidrBlock = v
	return s
}

type GetDestinationCidrBlockResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDestinationCidrBlockResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDestinationCidrBlockResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDestinationCidrBlockResponse) GoString() string {
	return s.String()
}

func (s *GetDestinationCidrBlockResponse) SetHeaders(v map[string]*string) *GetDestinationCidrBlockResponse {
	s.Headers = v
	return s
}

func (s *GetDestinationCidrBlockResponse) SetStatusCode(v int32) *GetDestinationCidrBlockResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDestinationCidrBlockResponse) SetBody(v *GetDestinationCidrBlockResponseBody) *GetDestinationCidrBlockResponse {
	s.Body = v
	return s
}

type GetElasticNetworkInterfaceRequest struct {
	// Lingjun Elastic Network Interface ID
	//
	// This parameter is required.
	//
	// example:
	//
	// leni-1234****
	ElasticNetworkInterfaceId *string `json:"ElasticNetworkInterfaceId,omitempty" xml:"ElasticNetworkInterfaceId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetElasticNetworkInterfaceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetElasticNetworkInterfaceRequest) GoString() string {
	return s.String()
}

func (s *GetElasticNetworkInterfaceRequest) SetElasticNetworkInterfaceId(v string) *GetElasticNetworkInterfaceRequest {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *GetElasticNetworkInterfaceRequest) SetRegionId(v string) *GetElasticNetworkInterfaceRequest {
	s.RegionId = &v
	return s
}

type GetElasticNetworkInterfaceResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Content *GetElasticNetworkInterfaceResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The return message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0901F411-28FA-5B9C-BAEE-7776463FF0DC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetElasticNetworkInterfaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetElasticNetworkInterfaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetElasticNetworkInterfaceResponseBody) SetAccessDeniedDetail(v string) *GetElasticNetworkInterfaceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBody) SetCode(v int32) *GetElasticNetworkInterfaceResponseBody {
	s.Code = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBody) SetContent(v *GetElasticNetworkInterfaceResponseBodyContent) *GetElasticNetworkInterfaceResponseBody {
	s.Content = v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBody) SetMessage(v string) *GetElasticNetworkInterfaceResponseBody {
	s.Message = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBody) SetRequestId(v string) *GetElasticNetworkInterfaceResponseBody {
	s.RequestId = &v
	return s
}

type GetElasticNetworkInterfaceResponseBodyContent struct {
	// The time when the data address was created.
	//
	// example:
	//
	// 2022-01-13 12:51:41
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The instance description.
	//
	// example:
	//
	// terraform-example
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Lingjun Elastic Network Interface ID
	//
	// example:
	//
	// leni-1234****
	ElasticNetworkInterfaceId *string `json:"ElasticNetworkInterfaceId,omitempty" xml:"ElasticNetworkInterfaceId,omitempty"`
	// Whether to enable the jumboFrame capability
	//
	// example:
	//
	// True
	EnableJumboFrame *bool `json:"EnableJumboFrame,omitempty" xml:"EnableJumboFrame,omitempty"`
	// vswitch gateway address
	//
	// example:
	//
	// 172.16.****
	Gateway *string `json:"Gateway,omitempty" xml:"Gateway,omitempty"`
	// The time when the agent was last modified.
	//
	// example:
	//
	// 2022-01-13 12:51:41
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Elastic Network Interface IP
	//
	// example:
	//
	// 203.107.****
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// IPV6 address
	Ipv6Addresses []*GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses `json:"Ipv6Addresses,omitempty" xml:"Ipv6Addresses,omitempty" type:"Repeated"`
	// mac address
	//
	// example:
	//
	// 00:22:6D:97:**:**
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// vswitch mask bits
	//
	// example:
	//
	// 24
	Mask *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	// The error message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Lingjun Node ID
	//
	// example:
	//
	// e01-cn-lbj3aej****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// Secondary private IP address
	PrivateIpAddresses []*GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses `json:"PrivateIpAddresses,omitempty" xml:"PrivateIpAddresses,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-0jl5s4p4laalruk7****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The state of the private gateway.
	//
	// Valid value:
	//
	// 	- Create Failed: the creation failure.
	//
	// 	- Delete Failed: the that failed to be deleted.
	//
	// 	- Executing
	//
	// 	- Available
	//
	// 	- Deleting
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// NIC Type
	//
	// Valid value:
	//
	// 	- CUSTOM: custom type.
	//
	// 	- DEFAULT: system type.
	//
	// example:
	//
	// DEFAULT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-uf6u8473r84e9****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// VPC ID
	//
	// example:
	//
	// vpc-j6ctp4n75306****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-wulanchabu-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetElasticNetworkInterfaceResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s GetElasticNetworkInterfaceResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) SetCreateTime(v string) *GetElasticNetworkInterfaceResponseBodyContent {
	s.CreateTime = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) SetDescription(v string) *GetElasticNetworkInterfaceResponseBodyContent {
	s.Description = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) SetElasticNetworkInterfaceId(v string) *GetElasticNetworkInterfaceResponseBodyContent {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) SetEnableJumboFrame(v bool) *GetElasticNetworkInterfaceResponseBodyContent {
	s.EnableJumboFrame = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) SetGateway(v string) *GetElasticNetworkInterfaceResponseBodyContent {
	s.Gateway = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) SetGmtModified(v string) *GetElasticNetworkInterfaceResponseBodyContent {
	s.GmtModified = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) SetIp(v string) *GetElasticNetworkInterfaceResponseBodyContent {
	s.Ip = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) SetIpv6Addresses(v []*GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses) *GetElasticNetworkInterfaceResponseBodyContent {
	s.Ipv6Addresses = v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) SetMac(v string) *GetElasticNetworkInterfaceResponseBodyContent {
	s.Mac = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) SetMask(v string) *GetElasticNetworkInterfaceResponseBodyContent {
	s.Mask = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) SetMessage(v string) *GetElasticNetworkInterfaceResponseBodyContent {
	s.Message = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) SetNodeId(v string) *GetElasticNetworkInterfaceResponseBodyContent {
	s.NodeId = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) SetPrivateIpAddresses(v []*GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses) *GetElasticNetworkInterfaceResponseBodyContent {
	s.PrivateIpAddresses = v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) SetRegionId(v string) *GetElasticNetworkInterfaceResponseBodyContent {
	s.RegionId = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) SetSecurityGroupId(v string) *GetElasticNetworkInterfaceResponseBodyContent {
	s.SecurityGroupId = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) SetStatus(v string) *GetElasticNetworkInterfaceResponseBodyContent {
	s.Status = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) SetType(v string) *GetElasticNetworkInterfaceResponseBodyContent {
	s.Type = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) SetVSwitchId(v string) *GetElasticNetworkInterfaceResponseBodyContent {
	s.VSwitchId = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) SetVpcId(v string) *GetElasticNetworkInterfaceResponseBodyContent {
	s.VpcId = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContent) SetZoneId(v string) *GetElasticNetworkInterfaceResponseBodyContent {
	s.ZoneId = &v
	return s
}

type GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses struct {
	// The instance description.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Lingjun Elastic Network Interface ID
	//
	// example:
	//
	// leni-1234****
	ElasticNetworkInterfaceId *string `json:"ElasticNetworkInterfaceId,omitempty" xml:"ElasticNetworkInterfaceId,omitempty"`
	// The time when the data address was created.
	//
	// example:
	//
	// 1585816811000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the cluster was updated.
	//
	// example:
	//
	// 1549012834000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// IPV6 unique identifier
	//
	// example:
	//
	// sip-sg3xabeq
	IpName *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
	// IPV6 address
	//
	// example:
	//
	// 2408:4005:3aa:1000:470d:66fb:56a5:****
	Ipv6Address *string `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the intervention entry. Valid value:
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses) String() string {
	return tea.Prettify(s)
}

func (s GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses) GoString() string {
	return s.String()
}

func (s *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses) SetDescription(v string) *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses {
	s.Description = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses) SetElasticNetworkInterfaceId(v string) *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses) SetGmtCreate(v string) *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses {
	s.GmtCreate = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses) SetGmtModified(v string) *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses {
	s.GmtModified = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses) SetIpName(v string) *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses {
	s.IpName = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses) SetIpv6Address(v string) *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses {
	s.Ipv6Address = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses) SetMessage(v string) *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses {
	s.Message = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses) SetRegionId(v string) *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses {
	s.RegionId = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses) SetStatus(v string) *GetElasticNetworkInterfaceResponseBodyContentIpv6Addresses {
	s.Status = &v
	return s
}

type GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses struct {
	// The instance description.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Lingjun Elastic Network Interface ID
	//
	// example:
	//
	// leni-1234****
	ElasticNetworkInterfaceId *string `json:"ElasticNetworkInterfaceId,omitempty" xml:"ElasticNetworkInterfaceId,omitempty"`
	// The time when the data address was created.
	//
	// example:
	//
	// 1672971789000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the cluster was updated.
	//
	// example:
	//
	// 1672971789000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Lingjun Elastic Network Interface Secondary Private IP Unique Identifier
	//
	// example:
	//
	// sip-ywz****
	IpName *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Lingjun Elastic Network Interface secondary private IP address
	//
	// example:
	//
	// 172.16.****
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the intervention entry. Valid value:
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses) String() string {
	return tea.Prettify(s)
}

func (s GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses) GoString() string {
	return s.String()
}

func (s *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses) SetDescription(v string) *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses {
	s.Description = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses) SetElasticNetworkInterfaceId(v string) *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses) SetGmtCreate(v string) *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses {
	s.GmtCreate = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses) SetGmtModified(v string) *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses {
	s.GmtModified = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses) SetIpName(v string) *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses {
	s.IpName = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses) SetMessage(v string) *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses {
	s.Message = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses) SetPrivateIpAddress(v string) *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses {
	s.PrivateIpAddress = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses) SetRegionId(v string) *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses {
	s.RegionId = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses) SetStatus(v string) *GetElasticNetworkInterfaceResponseBodyContentPrivateIpAddresses {
	s.Status = &v
	return s
}

type GetElasticNetworkInterfaceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetElasticNetworkInterfaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetElasticNetworkInterfaceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetElasticNetworkInterfaceResponse) GoString() string {
	return s.String()
}

func (s *GetElasticNetworkInterfaceResponse) SetHeaders(v map[string]*string) *GetElasticNetworkInterfaceResponse {
	s.Headers = v
	return s
}

func (s *GetElasticNetworkInterfaceResponse) SetStatusCode(v int32) *GetElasticNetworkInterfaceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponse) SetBody(v *GetElasticNetworkInterfaceResponseBody) *GetElasticNetworkInterfaceResponse {
	s.Body = v
	return s
}

type GetErRequest struct {
	// Lingjun HUB ID
	//
	// This parameter is required.
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Content *GetErResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// Information returned when the call fails
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// 308DE9D2-03A6-5B44-A369-67B75D1EE091
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetErResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetErResponseBody) GoString() string {
	return s.String()
}

func (s *GetErResponseBody) SetAccessDeniedDetail(v string) *GetErResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// The time when the data address was created.
	//
	// example:
	//
	// 1644283112720
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Description
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Network instance information list
	ErAttachments []*GetErResponseBodyContentErAttachments `json:"ErAttachments,omitempty" xml:"ErAttachments,omitempty" type:"Repeated"`
	// Lingjun HUB Instance ID
	//
	// example:
	//
	// er-aueyxxsy
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// Lingjun HUB Instance Name
	//
	// example:
	//
	// er-heyuan-main
	ErName *string `json:"ErName,omitempty" xml:"ErName,omitempty"`
	// The list of route entry information.
	ErRouteEntrys []*GetErResponseBodyContentErRouteEntrys `json:"ErRouteEntrys,omitempty" xml:"ErRouteEntrys,omitempty" type:"Repeated"`
	// routing policy information list
	ErRouteMaps []*GetErResponseBodyContentErRouteMaps `json:"ErRouteMaps,omitempty" xml:"ErRouteMaps,omitempty" type:"Repeated"`
	// The time when the agent was last modified.
	//
	// example:
	//
	// 1627545952000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Primary Zone
	//
	// example:
	//
	// cn-wulanchabu-b
	MasterZoneId *string `json:"MasterZoneId,omitempty" xml:"MasterZoneId,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-aekzlki4ehfse4y
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the intervention entry. Valid value:
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 1620939556166277
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
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

func (s *GetErResponseBodyContent) SetResourceGroupId(v string) *GetErResponseBodyContent {
	s.ResourceGroupId = &v
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
	// Cross-account
	//
	// example:
	//
	// false
	Across *bool `json:"Across,omitempty" xml:"Across,omitempty"`
	// Receive all routes automatically
	//
	// example:
	//
	// true
	AutoReceiveAllRoute *bool `json:"AutoReceiveAllRoute,omitempty" xml:"AutoReceiveAllRoute,omitempty"`
	// The time when the data address was created.
	//
	// example:
	//
	// 1644283112720
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The connection ID of the Lingjun HUB network instance.
	//
	// example:
	//
	// er-attachment-f32hxfsu
	ErAttachmentId *string `json:"ErAttachmentId,omitempty" xml:"ErAttachmentId,omitempty"`
	// Network Instance Name
	//
	// example:
	//
	// fudan-egpu
	ErAttachmentName *string `json:"ErAttachmentName,omitempty" xml:"ErAttachmentName,omitempty"`
	// Lingjun HUB Instance ID
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// The time when the agent was last modified.
	//
	// example:
	//
	// 1649303733000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// vpd-kkopgtne
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the ECU.
	//
	// example:
	//
	// zhijiao
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Instance type: VPD and VCC
	//
	// Valid value:
	//
	// 	- VCC: Lingjun Connection.
	//
	// 	- VPD: Lingjun network segment.
	//
	// example:
	//
	// VPD
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The synchronized region where the ECS instances are deployed.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-acfmzzka6bnjvbi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the tenant to which the resource belongs.
	//
	// example:
	//
	// xxxxxxxx
	ResourceTenantId *string `json:"ResourceTenantId,omitempty" xml:"ResourceTenantId,omitempty"`
	// The status of the intervention entry. Valid value:
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 1620939556166277
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
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

func (s *GetErResponseBodyContentErAttachments) SetResourceGroupId(v string) *GetErResponseBodyContentErAttachments {
	s.ResourceGroupId = &v
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
	// Destination CIDR Block
	//
	// example:
	//
	// 10.0.0.0/9
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// Lingjun HUB Instance ID
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// The ID of the route entry.
	//
	// example:
	//
	// er-rte-xnmsd2kl
	ErRouteEntryId *string `json:"ErRouteEntryId,omitempty" xml:"ErRouteEntryId,omitempty"`
	// The time when the cluster was updated.
	//
	// example:
	//
	// 1623317089000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Next Hop Instance
	//
	// example:
	//
	// vcc-xxkmggkw
	NextHopId *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// Next Hop Instance Type
	//
	// example:
	//
	// VCC
	NextHopType *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-acfmyoj5mg3w54y
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the tenant to which the resource belongs.
	//
	// example:
	//
	// 1620939556166277
	ResourceTenantId *string `json:"ResourceTenantId,omitempty" xml:"ResourceTenantId,omitempty"`
	// Route type
	//
	// example:
	//
	// System
	RouteType *string `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
	// The status of the intervention entry. Valid value:
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 1620939556166277
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
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

func (s *GetErResponseBodyContentErRouteEntrys) SetResourceGroupId(v string) *GetErResponseBodyContentErRouteEntrys {
	s.ResourceGroupId = &v
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
	// Policy behavior
	//
	// Valid value:
	//
	// 	- deny: rejects the.
	//
	// 	- permit: The allows.
	//
	// example:
	//
	// permit
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The time when the data address was created.
	//
	// example:
	//
	// 1645766599809
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Policy description
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Destination CIDR Block
	//
	// example:
	//
	// 10.0.0.0/8
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// Lingjun HUB ID
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// routing policy ID
	//
	// example:
	//
	// er-rmap-xkslnmsr
	ErRouteMapId *string `json:"ErRouteMapId,omitempty" xml:"ErRouteMapId,omitempty"`
	// The name of the routing policy.
	//
	// example:
	//
	// route-map-name
	ErRouteMapName *string `json:"ErRouteMapName,omitempty" xml:"ErRouteMapName,omitempty"`
	// The time when the agent was last modified.
	//
	// example:
	//
	// 1623899444000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the destination instance.
	//
	// example:
	//
	// vpd-sdkd2gkx
	ReceptionInstanceId *string `json:"ReceptionInstanceId,omitempty" xml:"ReceptionInstanceId,omitempty"`
	// The name of the destination instance.
	//
	// example:
	//
	// Reception-name
	ReceptionInstanceName *string `json:"ReceptionInstanceName,omitempty" xml:"ReceptionInstanceName,omitempty"`
	// The tenant to which the destination instance belongs.
	//
	// example:
	//
	// 1620939556166277
	ReceptionInstanceOwner *string `json:"ReceptionInstanceOwner,omitempty" xml:"ReceptionInstanceOwner,omitempty"`
	// The type of the destination instance.
	//
	// example:
	//
	// VPD
	ReceptionInstanceType *string `json:"ReceptionInstanceType,omitempty" xml:"ReceptionInstanceType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-acfmzaq3ypaqkdy
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Policy sequence number (1001-2000)
	//
	// example:
	//
	// 1001
	RouteMapNum *int32 `json:"RouteMapNum,omitempty" xml:"RouteMapNum,omitempty"`
	// The status of the intervention entry. Valid value:
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// XXQGPROD-zh_CN
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The ID of the source instance.
	//
	// example:
	//
	// vpd-xmglsymg
	TransmissionInstanceId *string `json:"TransmissionInstanceId,omitempty" xml:"TransmissionInstanceId,omitempty"`
	// Source instance name
	//
	// example:
	//
	// test-transmission
	TransmissionInstanceName *string `json:"TransmissionInstanceName,omitempty" xml:"TransmissionInstanceName,omitempty"`
	// The tenant to which the source instance belongs.
	//
	// example:
	//
	// 1620939556166277
	TransmissionInstanceOwner *string `json:"TransmissionInstanceOwner,omitempty" xml:"TransmissionInstanceOwner,omitempty"`
	// The type of the source instance.
	//
	// example:
	//
	// VPD
	TransmissionInstanceType *string `json:"TransmissionInstanceType,omitempty" xml:"TransmissionInstanceType,omitempty"`
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

func (s *GetErResponseBodyContentErRouteMaps) SetResourceGroupId(v string) *GetErResponseBodyContentErRouteMaps {
	s.ResourceGroupId = &v
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetErResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the Lingjun HUB network connection instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// er-attachment-i1ioibyf
	ErAttachmentId *string `json:"ErAttachmentId,omitempty" xml:"ErAttachmentId,omitempty"`
	// Lingjun HUB ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Content *GetErAttachmentResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is displayed.)
	//
	// example:
	//
	// You don\\"t have the permission to do this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7F0D9440-1F97-5613-87CD-D3047172A93C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetErAttachmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetErAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *GetErAttachmentResponseBody) SetAccessDeniedDetail(v string) *GetErAttachmentResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// Whether cross-account. Valid values:
	//
	// 	- **true**: The network instance is a cross-account resource.
	//
	// 	- **false**: The current network instance is a resource of the current account.
	//
	// example:
	//
	// fasle
	Across *bool `json:"Across,omitempty" xml:"Across,omitempty"`
	// Indicates whether to automatically receive all routes from all instances under the Lingjun HUB. Valid values:
	//
	// 	- **true**: received automatically.
	//
	// 	- **false**: Not received.
	//
	// example:
	//
	// true
	AutoReceiveAllRoute *bool `json:"AutoReceiveAllRoute,omitempty" xml:"AutoReceiveAllRoute,omitempty"`
	// The time when the activation code was created.
	//
	// example:
	//
	// 1648085472000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the Lingjun HUB network instance.
	//
	// example:
	//
	// er-attachment-i1ioibyf
	ErAttachmentId *string `json:"ErAttachmentId,omitempty" xml:"ErAttachmentId,omitempty"`
	// The name of the Lingjun HUB network instance.
	//
	// example:
	//
	// vpd-lxnsj2cx
	ErAttachmentName *string `json:"ErAttachmentName,omitempty" xml:"ErAttachmentName,omitempty"`
	// The ID of the Lingjun HUB instance.
	//
	// example:
	//
	// er-aueyxxsy
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// The time when the O\\&M task was modified.
	//
	// example:
	//
	// 1648085472000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the network instance. Valid values: **VPD*	- and **VCC**.
	//
	// For more information, see [What is Lingjun?](https://help.aliyun.com/document_detail/444430.html)
	//
	// You can query **Lingjun CIDR blocks*	- and **Lingjun connections*	- by [ListVpds](https://help.aliyun.com/document_detail/2331077.html) and [ListVccs](https://help.aliyun.com/document_detail/2399526.html?) respectively.
	//
	// example:
	//
	// vpd-lxnsj2cx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// vpd-wulanchabu-main
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The database type. Valid values:
	//
	// 	- **VPD**: indicates the Lingjun CIDR block.
	//
	// 	- **VCC**: indicates a Lingjun connection.
	//
	// example:
	//
	// VPD
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The returned message.
	//
	// example:
	//
	// test
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-aekzb3n5lgk2ieq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the tenant to which the resource belongs.
	//
	// example:
	//
	// 1620939556166277
	ResourceTenantId *string `json:"ResourceTenantId,omitempty" xml:"ResourceTenantId,omitempty"`
	// The status of the cache reserve instance. Valid values:
	//
	// 	- **Available**: Normal.
	//
	// 	- **Not Available**: Not available.
	//
	// 	- **Executing**: The task is being executed.
	//
	// 	- **Deleting**: The account is being deleted
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 1655449505171
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
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

func (s *GetErAttachmentResponseBodyContent) SetResourceGroupId(v string) *GetErAttachmentResponseBodyContent {
	s.ResourceGroupId = &v
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetErAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// Lingjun HUB ID
	//
	// This parameter is required.
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// The ID of the route entry.
	//
	// This parameter is required.
	//
	// example:
	//
	// er-rte-4q0jbylz
	ErRouteEntryId *string `json:"ErRouteEntryId,omitempty" xml:"ErRouteEntryId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Content *GetErRouteEntryResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// code: 400, Request was denied due to request throttling. request id: 7D177459-C1CF-5690-BB23-321D208B37D5
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// 1F38A2E6-CB47-5369-95D2-96D0C287B4A5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetErRouteEntryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetErRouteEntryResponseBody) GoString() string {
	return s.String()
}

func (s *GetErRouteEntryResponseBody) SetAccessDeniedDetail(v string) *GetErRouteEntryResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// Destination CIDR Block
	//
	// example:
	//
	// 11.0.0.0/16
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// Lingjun HUB Instance ID
	//
	// example:
	//
	// er-aueyxxsy
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// The ID of the route entry.
	//
	// example:
	//
	// er-rte-4q0jbylz
	ErRouteEntryId *string `json:"ErRouteEntryId,omitempty" xml:"ErRouteEntryId,omitempty"`
	// The time when the cluster was updated.
	//
	// example:
	//
	// 1666677783000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Next Hop Instance
	//
	// example:
	//
	// vcc-cn-209300qha01
	NextHopId *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// Next Hop Instance Type
	//
	// example:
	//
	// VCC
	NextHopType *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-aekzb3n5lgk2ieq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Route type
	//
	// example:
	//
	// System
	RouteType *string `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
	// The status of the intervention entry. Valid value:
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 1655449505171
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
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

func (s *GetErRouteEntryResponseBodyContent) SetResourceGroupId(v string) *GetErRouteEntryResponseBodyContent {
	s.ResourceGroupId = &v
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetErRouteEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// Lingjun HUB ID
	//
	// This parameter is required.
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// routing policy ID
	//
	// This parameter is required.
	//
	// example:
	//
	// er-rmap-uwglhzom
	ErRouteMapId *string `json:"ErRouteMapId,omitempty" xml:"ErRouteMapId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Content *GetErRouteMapResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A88DFED5-24B7-5A3E-87DE-380BF06F3C90
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetErRouteMapResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetErRouteMapResponseBody) GoString() string {
	return s.String()
}

func (s *GetErRouteMapResponseBody) SetAccessDeniedDetail(v string) *GetErRouteMapResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// Policy behavior; optional values:
	//
	// 	- **permit**: Allow
	//
	// 	- **deny**: Rejected
	//
	// example:
	//
	// permit
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// Policy description
	//
	// example:
	//
	// ssss
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Destination CIDR block
	//
	// example:
	//
	// 0.0.0.0/0
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// Lingjun HUB ID
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// Lingjun HUB routing policy ID
	//
	// example:
	//
	// er-rmap-uwglhzom
	ErRouteMapId *string `json:"ErRouteMapId,omitempty" xml:"ErRouteMapId,omitempty"`
	// Lingjun HUB routing policy Name
	//
	// example:
	//
	// er-rmap-wulanchabu
	ErRouteMapName *string `json:"ErRouteMapName,omitempty" xml:"ErRouteMapName,omitempty"`
	// The time when the data address was created.
	//
	// example:
	//
	// 1648085472000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the agent was last modified.
	//
	// example:
	//
	// 1648085472000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Receive Instance ID
	//
	// example:
	//
	// vpd-x25vxrb2
	ReceptionInstanceId *string `json:"ReceptionInstanceId,omitempty" xml:"ReceptionInstanceId,omitempty"`
	// Receive Instance Name
	//
	// example:
	//
	// vpd-receprion
	ReceptionInstanceName *string `json:"ReceptionInstanceName,omitempty" xml:"ReceptionInstanceName,omitempty"`
	// The tenant to which the receiving instance belongs
	//
	// example:
	//
	// 1620939556166277
	ReceptionInstanceOwner *string `json:"ReceptionInstanceOwner,omitempty" xml:"ReceptionInstanceOwner,omitempty"`
	// The type of the received instance. Optional values:
	//
	// 	- **VPD**: Lingjun network segment.
	//
	// 	- **VCC**: Lingjun Connection.
	//
	// example:
	//
	// VPD
	ReceptionInstanceType *string `json:"ReceptionInstanceType,omitempty" xml:"ReceptionInstanceType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-aekzlki4ehfse4y
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the policy.
	//
	// A smaller sequence number indicates a lower priority. When a route is matched, a policy with a higher priority is preferentially matched.
	//
	// **Valid values: 1001 to 2000**
	//
	// example:
	//
	// 1001
	RouteMapNum *int32 `json:"RouteMapNum,omitempty" xml:"RouteMapNum,omitempty"`
	// The status of the cache reserve instance. Valid values:
	//
	// 	- **Available**
	//
	// 	- **Not Available**: Unavailable
	//
	// 	- **Executing**: Executing
	//
	// 	- **Deleting**: The node is being deleted.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 1655449505171
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// Release Instance ID
	//
	// example:
	//
	// vpd-xgkb2kl
	TransmissionInstanceId *string `json:"TransmissionInstanceId,omitempty" xml:"TransmissionInstanceId,omitempty"`
	// Release Instance Name
	//
	// example:
	//
	// vpd-transimit
	TransmissionInstanceName *string `json:"TransmissionInstanceName,omitempty" xml:"TransmissionInstanceName,omitempty"`
	// The tenant to which the published instance belongs
	//
	// example:
	//
	// 1620939556166277
	TransmissionInstanceOwner *string `json:"TransmissionInstanceOwner,omitempty" xml:"TransmissionInstanceOwner,omitempty"`
	// Publish instance type; optional values:
	//
	// 	- **VPD**: Lingjun network segment.
	//
	// 	- **VCC**: Lingjun Connection.
	//
	// example:
	//
	// VPD
	TransmissionInstanceType *string `json:"TransmissionInstanceType,omitempty" xml:"TransmissionInstanceType,omitempty"`
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

func (s *GetErRouteMapResponseBodyContent) SetResourceGroupId(v string) *GetErRouteMapResponseBodyContent {
	s.ResourceGroupId = &v
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetErRouteMapResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type GetFabricTopologyRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// i-169263721924****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Lingjun network interface controller ID List
	LniIds []*string `json:"LniIds,omitempty" xml:"LniIds,omitempty" type:"Repeated"`
	// Node ID list
	NodeIds []*string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-k8i0g9fk68t7u0u2w****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// Lingjun CIDR block ID
	//
	// example:
	//
	// vpd-aof7****
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
}

func (s GetFabricTopologyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFabricTopologyRequest) GoString() string {
	return s.String()
}

func (s *GetFabricTopologyRequest) SetClusterId(v string) *GetFabricTopologyRequest {
	s.ClusterId = &v
	return s
}

func (s *GetFabricTopologyRequest) SetLniIds(v []*string) *GetFabricTopologyRequest {
	s.LniIds = v
	return s
}

func (s *GetFabricTopologyRequest) SetNodeIds(v []*string) *GetFabricTopologyRequest {
	s.NodeIds = v
	return s
}

func (s *GetFabricTopologyRequest) SetRegionId(v string) *GetFabricTopologyRequest {
	s.RegionId = &v
	return s
}

func (s *GetFabricTopologyRequest) SetVpcId(v string) *GetFabricTopologyRequest {
	s.VpcId = &v
	return s
}

func (s *GetFabricTopologyRequest) SetVpdId(v string) *GetFabricTopologyRequest {
	s.VpdId = &v
	return s
}

type GetFabricTopologyResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Content *GetFabricTopologyResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// AC8C713A-A9F4-5984-A5E1-76496DF35153
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFabricTopologyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFabricTopologyResponseBody) GoString() string {
	return s.String()
}

func (s *GetFabricTopologyResponseBody) SetAccessDeniedDetail(v string) *GetFabricTopologyResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetFabricTopologyResponseBody) SetCode(v int32) *GetFabricTopologyResponseBody {
	s.Code = &v
	return s
}

func (s *GetFabricTopologyResponseBody) SetContent(v *GetFabricTopologyResponseBodyContent) *GetFabricTopologyResponseBody {
	s.Content = v
	return s
}

func (s *GetFabricTopologyResponseBody) SetMessage(v string) *GetFabricTopologyResponseBody {
	s.Message = &v
	return s
}

func (s *GetFabricTopologyResponseBody) SetRequestId(v string) *GetFabricTopologyResponseBody {
	s.RequestId = &v
	return s
}

type GetFabricTopologyResponseBodyContent struct {
	// The cluster ID.
	//
	// example:
	//
	// cluster-****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// network interface controller Topology Information
	TopoInfo []*GetFabricTopologyResponseBodyContentTopoInfo `json:"TopoInfo,omitempty" xml:"TopoInfo,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-j6ctp4n75306****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// Lingjun CIDR block ID
	//
	// example:
	//
	// vpd-fuli****
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
}

func (s GetFabricTopologyResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s GetFabricTopologyResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetFabricTopologyResponseBodyContent) SetClusterId(v string) *GetFabricTopologyResponseBodyContent {
	s.ClusterId = &v
	return s
}

func (s *GetFabricTopologyResponseBodyContent) SetRegionId(v string) *GetFabricTopologyResponseBodyContent {
	s.RegionId = &v
	return s
}

func (s *GetFabricTopologyResponseBodyContent) SetTopoInfo(v []*GetFabricTopologyResponseBodyContentTopoInfo) *GetFabricTopologyResponseBodyContent {
	s.TopoInfo = v
	return s
}

func (s *GetFabricTopologyResponseBodyContent) SetVpcId(v string) *GetFabricTopologyResponseBodyContent {
	s.VpcId = &v
	return s
}

func (s *GetFabricTopologyResponseBodyContent) SetVpdId(v string) *GetFabricTopologyResponseBodyContent {
	s.VpdId = &v
	return s
}

type GetFabricTopologyResponseBodyContentTopoInfo struct {
	// The resource name.
	//
	// example:
	//
	// core-1
	LayerName *string `json:"LayerName,omitempty" xml:"LayerName,omitempty"`
	// Hierarchical resource types
	//
	// Valid value:
	//
	// 	- core: core layer.
	//
	// 	- node: Lingjun node.
	//
	// 	- lni: lingjun network interface controller.
	//
	// 	- spine: backbone layer.
	//
	// 	- leaf: access layer
	//
	// example:
	//
	// core
	LayerType *string `json:"LayerType,omitempty" xml:"LayerType,omitempty"`
	// Next Level
	NextLayer []interface{} `json:"NextLayer,omitempty" xml:"NextLayer,omitempty" type:"Repeated"`
}

func (s GetFabricTopologyResponseBodyContentTopoInfo) String() string {
	return tea.Prettify(s)
}

func (s GetFabricTopologyResponseBodyContentTopoInfo) GoString() string {
	return s.String()
}

func (s *GetFabricTopologyResponseBodyContentTopoInfo) SetLayerName(v string) *GetFabricTopologyResponseBodyContentTopoInfo {
	s.LayerName = &v
	return s
}

func (s *GetFabricTopologyResponseBodyContentTopoInfo) SetLayerType(v string) *GetFabricTopologyResponseBodyContentTopoInfo {
	s.LayerType = &v
	return s
}

func (s *GetFabricTopologyResponseBodyContentTopoInfo) SetNextLayer(v []interface{}) *GetFabricTopologyResponseBodyContentTopoInfo {
	s.NextLayer = v
	return s
}

type GetFabricTopologyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFabricTopologyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFabricTopologyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFabricTopologyResponse) GoString() string {
	return s.String()
}

func (s *GetFabricTopologyResponse) SetHeaders(v map[string]*string) *GetFabricTopologyResponse {
	s.Headers = v
	return s
}

func (s *GetFabricTopologyResponse) SetStatusCode(v int32) *GetFabricTopologyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFabricTopologyResponse) SetBody(v *GetFabricTopologyResponseBody) *GetFabricTopologyResponse {
	s.Body = v
	return s
}

type GetLeniPrivateIpAddressRequest struct {
	// Lingjun Elastic Network Interface ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// leni-1234****
	ElasticNetworkInterfaceId *string `json:"ElasticNetworkInterfaceId,omitempty" xml:"ElasticNetworkInterfaceId,omitempty"`
	// Lingjun Elastic Network Interface secondary private IP unique identifier.
	//
	// This parameter is required.
	//
	// example:
	//
	// sip-8ylg****
	IpName *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetLeniPrivateIpAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLeniPrivateIpAddressRequest) GoString() string {
	return s.String()
}

func (s *GetLeniPrivateIpAddressRequest) SetElasticNetworkInterfaceId(v string) *GetLeniPrivateIpAddressRequest {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *GetLeniPrivateIpAddressRequest) SetIpName(v string) *GetLeniPrivateIpAddressRequest {
	s.IpName = &v
	return s
}

func (s *GetLeniPrivateIpAddressRequest) SetRegionId(v string) *GetLeniPrivateIpAddressRequest {
	s.RegionId = &v
	return s
}

type GetLeniPrivateIpAddressResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Content *GetLeniPrivateIpAddressResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AC8C713A-A9F4-5984-A5E1-76496DF35153
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLeniPrivateIpAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLeniPrivateIpAddressResponseBody) GoString() string {
	return s.String()
}

func (s *GetLeniPrivateIpAddressResponseBody) SetAccessDeniedDetail(v string) *GetLeniPrivateIpAddressResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetLeniPrivateIpAddressResponseBody) SetCode(v int32) *GetLeniPrivateIpAddressResponseBody {
	s.Code = &v
	return s
}

func (s *GetLeniPrivateIpAddressResponseBody) SetContent(v *GetLeniPrivateIpAddressResponseBodyContent) *GetLeniPrivateIpAddressResponseBody {
	s.Content = v
	return s
}

func (s *GetLeniPrivateIpAddressResponseBody) SetMessage(v string) *GetLeniPrivateIpAddressResponseBody {
	s.Message = &v
	return s
}

func (s *GetLeniPrivateIpAddressResponseBody) SetRequestId(v string) *GetLeniPrivateIpAddressResponseBody {
	s.RequestId = &v
	return s
}

type GetLeniPrivateIpAddressResponseBodyContent struct {
	// The description.
	//
	// example:
	//
	// zhenyuan wdl workflow
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Lingjun Elastic Network Interface ID.
	//
	// example:
	//
	// leni-1234****
	ElasticNetworkInterfaceId *string `json:"ElasticNetworkInterfaceId,omitempty" xml:"ElasticNetworkInterfaceId,omitempty"`
	// The time when the activation code was created.
	//
	// example:
	//
	// 1663722356000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the certificate was updated.
	//
	// example:
	//
	// 1635231890000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Lingjun Elastic Network Interface secondary private IP unique identifier.
	//
	// example:
	//
	// sip-8ylg****
	IpName *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
	// The returned message.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Lingjun Elastic Network Interface secondary private IP address.
	//
	// example:
	//
	// 10.42.****
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The task status.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetLeniPrivateIpAddressResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s GetLeniPrivateIpAddressResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetLeniPrivateIpAddressResponseBodyContent) SetDescription(v string) *GetLeniPrivateIpAddressResponseBodyContent {
	s.Description = &v
	return s
}

func (s *GetLeniPrivateIpAddressResponseBodyContent) SetElasticNetworkInterfaceId(v string) *GetLeniPrivateIpAddressResponseBodyContent {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *GetLeniPrivateIpAddressResponseBodyContent) SetGmtCreate(v string) *GetLeniPrivateIpAddressResponseBodyContent {
	s.GmtCreate = &v
	return s
}

func (s *GetLeniPrivateIpAddressResponseBodyContent) SetGmtModified(v string) *GetLeniPrivateIpAddressResponseBodyContent {
	s.GmtModified = &v
	return s
}

func (s *GetLeniPrivateIpAddressResponseBodyContent) SetIpName(v string) *GetLeniPrivateIpAddressResponseBodyContent {
	s.IpName = &v
	return s
}

func (s *GetLeniPrivateIpAddressResponseBodyContent) SetMessage(v string) *GetLeniPrivateIpAddressResponseBodyContent {
	s.Message = &v
	return s
}

func (s *GetLeniPrivateIpAddressResponseBodyContent) SetPrivateIpAddress(v string) *GetLeniPrivateIpAddressResponseBodyContent {
	s.PrivateIpAddress = &v
	return s
}

func (s *GetLeniPrivateIpAddressResponseBodyContent) SetRegionId(v string) *GetLeniPrivateIpAddressResponseBodyContent {
	s.RegionId = &v
	return s
}

func (s *GetLeniPrivateIpAddressResponseBodyContent) SetStatus(v string) *GetLeniPrivateIpAddressResponseBodyContent {
	s.Status = &v
	return s
}

type GetLeniPrivateIpAddressResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLeniPrivateIpAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLeniPrivateIpAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLeniPrivateIpAddressResponse) GoString() string {
	return s.String()
}

func (s *GetLeniPrivateIpAddressResponse) SetHeaders(v map[string]*string) *GetLeniPrivateIpAddressResponse {
	s.Headers = v
	return s
}

func (s *GetLeniPrivateIpAddressResponse) SetStatusCode(v int32) *GetLeniPrivateIpAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLeniPrivateIpAddressResponse) SetBody(v *GetLeniPrivateIpAddressResponseBody) *GetLeniPrivateIpAddressResponse {
	s.Body = v
	return s
}

type GetLniPrivateIpAddressRequest struct {
	// IP unique identifier
	//
	// This parameter is required.
	//
	// example:
	//
	// sip-xxxxx
	IpName *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
	// Lingjun network interface controller ID
	//
	// This parameter is required.
	//
	// example:
	//
	// lni-bp18exxqa2rvfn45e5pz
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Content *GetLniPrivateIpAddressResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// You don\\"t have the permission of this operation, action=eflo:GetLniPrivateIpAddress, arn=acs:eflo:cn-wulanchabu:1382782317087063:networkinterface/00
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// DBAD15D6-3F47-5B36-8A92-57C2919D13D0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLniPrivateIpAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLniPrivateIpAddressResponseBody) GoString() string {
	return s.String()
}

func (s *GetLniPrivateIpAddressResponseBody) SetAccessDeniedDetail(v string) *GetLniPrivateIpAddressResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// The instance description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the data address was created.
	//
	// example:
	//
	// 2022-12-26 20:16:36
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// MAC address of the secondary private network
	//
	// example:
	//
	// 00-ff-84-15-ba-67
	IpAddressMac *string `json:"IpAddressMac,omitempty" xml:"IpAddressMac,omitempty"`
	// IP unique identifier
	//
	// example:
	//
	// sip-xxxxx
	IpName *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// You don\\"t have the permission of this operation, action=eflo:ListVpdRouteEntries, arn=acs:eflo:cn-wulanchabu:1263399219805497:vpd_rte/*, resourceGroup=null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Lingjun network interface controller ID
	//
	// example:
	//
	// lni-2ze4uww7n6hsfzrwq77y
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The secondary private IP address of the Lingjun network interface controller.
	//
	// example:
	//
	// 10.42.5.92
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The state of the rule.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetLniPrivateIpAddressResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s GetLniPrivateIpAddressResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetLniPrivateIpAddressResponseBodyContent) SetDescription(v string) *GetLniPrivateIpAddressResponseBodyContent {
	s.Description = &v
	return s
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLniPrivateIpAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// Lingjun network interface controller ID
	//
	// This parameter is required.
	//
	// example:
	//
	// lni-bp18exxqa2rvfn45e5pz
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Subnet of Lingjun
	//
	// This parameter is required.
	//
	// example:
	//
	// subnet-f3zfzmnc
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data. (If a resource has dependent resources, the existing dependent resources are returned.)
	Content *GetNetworkInterfaceResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// A88DFED5-24B7-5A3E-87DE-380BF06F3C90
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNetworkInterfaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNetworkInterfaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetNetworkInterfaceResponseBody) SetAccessDeniedDetail(v string) *GetNetworkInterfaceResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// The time when the data address was created.
	//
	// example:
	//
	// 1678273219000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Port
	Ethernet []*string `json:"Ethernet,omitempty" xml:"Ethernet,omitempty" type:"Repeated"`
	// Gateway
	//
	// example:
	//
	// 172.24.20.254
	Gateway *string `json:"Gateway,omitempty" xml:"Gateway,omitempty"`
	// The IP address of the BE cluster.
	//
	// example:
	//
	// 203.107.60.69
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// NC Type
	//
	// Valid value:
	//
	// 	- CUSTOM_LNI_INTEGRATION: two-network one-in-one architecture Lingjun hosting network interface controller.
	//
	// 	- CPU: CPU machine.
	//
	// 	- ELASTIC_6.2: Machine
	//
	// 	- GPU: GPU machine.
	//
	// 	- DEFAULT: the old CPU machine.
	//
	// 	- CUSTOM_LNI: two network separation architecture Lingjun hosting network interface controller.
	//
	// example:
	//
	// DEFAULT
	NcType *string `json:"NcType,omitempty" xml:"NcType,omitempty"`
	// Lingjun network interface controller ID
	//
	// example:
	//
	// lni-f8z4scmfh0u4ewv6vdd8
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// ENI Name
	//
	// example:
	//
	// bond0
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitempty" xml:"NetworkInterfaceName,omitempty"`
	// The ID of the host.
	//
	// example:
	//
	// masterintranett2fdth5fkoocg
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// Secondary Private IP\\&MAC Address Collection
	PrivateIpAddressMacGroup []*GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup `json:"PrivateIpAddressMacGroup,omitempty" xml:"PrivateIpAddressMacGroup,omitempty" type:"Repeated"`
	// network interface controller private secondary IP limit
	//
	// example:
	//
	// 0
	Quota *int32 `json:"Quota,omitempty" xml:"Quota,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Service network interface controller address
	//
	// example:
	//
	// 01-00-5e-00-00-16
	ServiceMac *string `json:"ServiceMac,omitempty" xml:"ServiceMac,omitempty"`
	// The status of the intervention entry. Valid value:
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Lingjun subnet (Subnet) basic information
	SubnetBaseInfo *GetNetworkInterfaceResponseBodyContentSubnetBaseInfo `json:"SubnetBaseInfo,omitempty" xml:"SubnetBaseInfo,omitempty" type:"Struct"`
	// The ID of the tenant.
	//
	// example:
	//
	// 1655449505171
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// Basic information of Lingjun network segment (VPD)
	VpdBaseInfo *GetNetworkInterfaceResponseBodyContentVpdBaseInfo `json:"VpdBaseInfo,omitempty" xml:"VpdBaseInfo,omitempty" type:"Struct"`
	// The zone ID.
	//
	// example:
	//
	// cn-wulanchabu-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
	// The instance description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Secondary private MAC address
	//
	// example:
	//
	// 01-00-5e-00-00-16
	IpAddressMac *string `json:"IpAddressMac,omitempty" xml:"IpAddressMac,omitempty"`
	// IP unique identifier
	//
	// example:
	//
	// sip-xxxx
	IpName *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Secondary private IP address
	//
	// example:
	//
	// 172.23.161.57
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The state of the rule.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup) String() string {
	return tea.Prettify(s)
}

func (s GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup) GoString() string {
	return s.String()
}

func (s *GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup) SetDescription(v string) *GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup {
	s.Description = &v
	return s
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
	// Network address segment
	//
	// example:
	//
	// 116.233.21.57/32
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The time when the data address was created.
	//
	// example:
	//
	// 1678273219000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the Subnet instance.
	//
	// example:
	//
	// subnet-urb01blo
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	// The name of the Subnet instance.
	//
	// example:
	//
	// subnet-1
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
	// The network segment of the Lingjun subnet.
	//
	// 	- The network segment of the subnet must be a proper subset of the network segment of Lingjun to which it belongs, and the mask must be between 16 bits and 29 bits, which can provide 8 to 65536 addresses. For example, the CIDR block of the Lingjun CIDR block is 192.168.0.0/16, and the CIDR blocks of the subnets under the Lingjun CIDR block are 192.168.0.0/17 to 192.168.0.0/29.
	//
	// 	- The first and last three IP addresses of each subnet segment are reserved by the system. For example, the CIDR blocks of the subnet are 192.168.1.0/24,192.168.1.0, 192.168.1.253, 192.168.1.254, and 192.168.1.255.
	//
	// For more information about CIDR blocks, see the [What is CIDR?](https://www.alibabacloud.com/help/doc-detail/40637.htm#title-gu4-uzk-12r) section in the "Network FAQ" topic.
	//
	// This parameter is left empty by default.
	//
	// example:
	//
	// 172.18.0.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The time when the data address was created.
	//
	// example:
	//
	// 1678273219000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the VPD instance.
	//
	// example:
	//
	// vpd-ppdunxzc
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// The name of the VPD instance.
	//
	// example:
	//
	// vpd-1
	VpdName *string `json:"VpdName,omitempty" xml:"VpdName,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNetworkInterfaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type GetNodeInfoForPodRequest struct {
	// The ID of the node for this operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// node-be70****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetNodeInfoForPodRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNodeInfoForPodRequest) GoString() string {
	return s.String()
}

func (s *GetNodeInfoForPodRequest) SetNodeId(v string) *GetNodeInfoForPodRequest {
	s.NodeId = &v
	return s
}

func (s *GetNodeInfoForPodRequest) SetRegionId(v string) *GetNodeInfoForPodRequest {
	s.RegionId = &v
	return s
}

type GetNodeInfoForPodResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Content *GetNodeInfoForPodResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// You don\\"t have the permission of this operation, action=eflo:GetNodeInfoForPod, arn=acs:eflo:cn-wulanchabu:1111156667137893:networkinterface/*, resourceGroup=null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// 9C50C9CD-E799-54DA-BA7A-1FAF3DF80857
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNodeInfoForPodResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNodeInfoForPodResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeInfoForPodResponseBody) SetAccessDeniedDetail(v string) *GetNodeInfoForPodResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetNodeInfoForPodResponseBody) SetCode(v int32) *GetNodeInfoForPodResponseBody {
	s.Code = &v
	return s
}

func (s *GetNodeInfoForPodResponseBody) SetContent(v *GetNodeInfoForPodResponseBodyContent) *GetNodeInfoForPodResponseBody {
	s.Content = v
	return s
}

func (s *GetNodeInfoForPodResponseBody) SetMessage(v string) *GetNodeInfoForPodResponseBody {
	s.Message = &v
	return s
}

func (s *GetNodeInfoForPodResponseBody) SetRequestId(v string) *GetNodeInfoForPodResponseBody {
	s.RequestId = &v
	return s
}

type GetNodeInfoForPodResponseBodyContent struct {
	// The cluster ID.
	//
	// example:
	//
	// cluster-****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Lingjun Gaomi network interface controller quota
	//
	// example:
	//
	// 10
	HdeniQuota *int32 `json:"HdeniQuota,omitempty" xml:"HdeniQuota,omitempty"`
	// Lingjun Elastic Network Interface quota, including system type
	//
	// example:
	//
	// 10
	LeniQuota *int32 `json:"LeniQuota,omitempty" xml:"LeniQuota,omitempty"`
	// Lingjun Elastic Network Interface Secondary Private IP Quota
	//
	// example:
	//
	// 10
	LeniSipQuota *int32 `json:"LeniSipQuota,omitempty" xml:"LeniSipQuota,omitempty"`
	// Lingjun network interface controller Secondary Private IP Quota
	//
	// example:
	//
	// 10
	LniSipQuota *int32 `json:"LniSipQuota,omitempty" xml:"LniSipQuota,omitempty"`
	// The ID of the node for this operation.
	//
	// example:
	//
	// node-be70****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// List of VSwitches that can apply for IP addresses on this node
	VSwitches []*string `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
	// The ID of the Virtual Private Cloud to which the current node belongs.
	//
	// example:
	//
	// vpc-j6ctp4n75306****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-wulanchabu-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetNodeInfoForPodResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s GetNodeInfoForPodResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetNodeInfoForPodResponseBodyContent) SetClusterId(v string) *GetNodeInfoForPodResponseBodyContent {
	s.ClusterId = &v
	return s
}

func (s *GetNodeInfoForPodResponseBodyContent) SetHdeniQuota(v int32) *GetNodeInfoForPodResponseBodyContent {
	s.HdeniQuota = &v
	return s
}

func (s *GetNodeInfoForPodResponseBodyContent) SetLeniQuota(v int32) *GetNodeInfoForPodResponseBodyContent {
	s.LeniQuota = &v
	return s
}

func (s *GetNodeInfoForPodResponseBodyContent) SetLeniSipQuota(v int32) *GetNodeInfoForPodResponseBodyContent {
	s.LeniSipQuota = &v
	return s
}

func (s *GetNodeInfoForPodResponseBodyContent) SetLniSipQuota(v int32) *GetNodeInfoForPodResponseBodyContent {
	s.LniSipQuota = &v
	return s
}

func (s *GetNodeInfoForPodResponseBodyContent) SetNodeId(v string) *GetNodeInfoForPodResponseBodyContent {
	s.NodeId = &v
	return s
}

func (s *GetNodeInfoForPodResponseBodyContent) SetRegionId(v string) *GetNodeInfoForPodResponseBodyContent {
	s.RegionId = &v
	return s
}

func (s *GetNodeInfoForPodResponseBodyContent) SetVSwitches(v []*string) *GetNodeInfoForPodResponseBodyContent {
	s.VSwitches = v
	return s
}

func (s *GetNodeInfoForPodResponseBodyContent) SetVpcId(v string) *GetNodeInfoForPodResponseBodyContent {
	s.VpcId = &v
	return s
}

func (s *GetNodeInfoForPodResponseBodyContent) SetZoneId(v string) *GetNodeInfoForPodResponseBodyContent {
	s.ZoneId = &v
	return s
}

type GetNodeInfoForPodResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNodeInfoForPodResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNodeInfoForPodResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNodeInfoForPodResponse) GoString() string {
	return s.String()
}

func (s *GetNodeInfoForPodResponse) SetHeaders(v map[string]*string) *GetNodeInfoForPodResponse {
	s.Headers = v
	return s
}

func (s *GetNodeInfoForPodResponse) SetStatusCode(v int32) *GetNodeInfoForPodResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNodeInfoForPodResponse) SetBody(v *GetNodeInfoForPodResponseBody) *GetNodeInfoForPodResponse {
	s.Body = v
	return s
}

type GetSubnetRequest struct {
	// The region ID of the data center.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Lingjun subnet instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// subnet-2avf0itf
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	// The ID of the CIDR block to which Lingjun belongs.
	//
	// example:
	//
	// vpd-cxcmdk1m
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Content *GetSubnetResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0901F411-28FA-5B9C-BAEE-7776463FF0DC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSubnetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSubnetResponseBody) GoString() string {
	return s.String()
}

func (s *GetSubnetResponseBody) SetAccessDeniedDetail(v string) *GetSubnetResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// The number of available IP addresses.
	//
	// example:
	//
	// 1024
	AvailableIps *int32 `json:"AvailableIps,omitempty" xml:"AvailableIps,omitempty"`
	// The CIDR block of the Subnet.
	//
	// 	- The network segment of the subnet must be a proper subset of the network segment of Lingjun to which it belongs, and the mask must be between 16 bits and 29 bits, which can provide 8 to 65536 addresses. For example, the CIDR block of the Lingjun CIDR block is 192.168.0.0/16, and the CIDR blocks of the subnets under the Lingjun CIDR block are 192.168.0.0/17 to 192.168.0.0/29.
	//
	// 	- The first and last three IP addresses of each subnet segment are reserved by the system. For example, the CIDR blocks of the subnet are 192.168.1.0/24,192.168.1.0, 192.168.1.253, 192.168.1.254, and 192.168.1.255.
	//
	// example:
	//
	// 10.10.10.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The time when the activation code was created.
	//
	// example:
	//
	// 1648085472000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the O\\&M task was modified.
	//
	// example:
	//
	// 1678273219000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The number of SLB.
	//
	// example:
	//
	// 0
	LbCount *int64 `json:"LbCount,omitempty" xml:"LbCount,omitempty"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// test example
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The number of NCs.
	//
	// example:
	//
	// 16
	NcCount *int32 `json:"NcCount,omitempty" xml:"NcCount,omitempty"`
	// Number of Lingjun network interface controller
	//
	// example:
	//
	// 4
	NetworkInterfaceCount *int32 `json:"NetworkInterfaceCount,omitempty" xml:"NetworkInterfaceCount,omitempty"`
	// The total number of secondary private IP addresses.
	//
	// example:
	//
	// 20
	PrivateIpCount *int64 `json:"PrivateIpCount,omitempty" xml:"PrivateIpCount,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of your Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-aek2l4sq6l7unhi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the cache reserve instance.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the Lingjun subnet instance.
	//
	// example:
	//
	// subnet-aj93mko8
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	// The name of the Lingjun subnet instance.
	//
	// example:
	//
	// subnet-1
	SubnetName *string `json:"SubnetName,omitempty" xml:"SubnetName,omitempty"`
	// The tag information.
	//
	// You can specify up to 20 tags.
	Tags []*GetSubnetResponseBodyContentTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The tenant ID.
	//
	// example:
	//
	// 1655449505171
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// Lingjun Subnet Usage Type; optional; optional. Valid values:
	//
	// 	- **Empty for common data types**
	//
	// 	- **OOB*	- :OOB type
	//
	// 	- **LB**: LB type
	//
	// example:
	//
	// OOB
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The information about the network segment of Lingjun.
	VpdBaseInfo *GetSubnetResponseBodyContentVpdBaseInfo `json:"VpdBaseInfo,omitempty" xml:"VpdBaseInfo,omitempty" type:"Struct"`
	// The ID of the Lingjun CIDR block.
	//
	// example:
	//
	// vpd-ze3na0wf
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-wulanchabu-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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

func (s *GetSubnetResponseBodyContent) SetNcCount(v int32) *GetSubnetResponseBodyContent {
	s.NcCount = &v
	return s
}

func (s *GetSubnetResponseBodyContent) SetNetworkInterfaceCount(v int32) *GetSubnetResponseBodyContent {
	s.NetworkInterfaceCount = &v
	return s
}

func (s *GetSubnetResponseBodyContent) SetPrivateIpCount(v int64) *GetSubnetResponseBodyContent {
	s.PrivateIpCount = &v
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
	// The tag key.
	//
	// You cannot specify an empty string as a tag key. It can be up to 64 characters in length and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// You can specify at most 20 tag keys in each call.
	//
	// example:
	//
	// tag-subnet
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag that is added to the resource.
	//
	// The tag value can be empty or a string of up to 128 characters. It cannot start with aliyun or acs:, and cannot contain http:// or https://.
	//
	// Each key-value pair must be unique. You can specify values for at most 20 tag keys in each call.
	//
	// example:
	//
	// subnet-group-1
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
	// The CIDR block of the VPD.
	//
	// 	- We recommend that you use an RFC private endpoint as the Lingjun CIDR block, such as 10.0.0.0/8,172.16.0.0/12,192.168.0.0/16. In scenarios where the Doringjun CIDR block is connected to each other or where the Lingjun CIDR block is connected to a VPC, make sure that the addresses do not conflict with each other.
	//
	// 	- You can also use a custom CIDR block other than 100.64.0.0/10, 224.0.0.0/4, 127.0.0.0/8, or 169.254.0.0/16 and their subnets as the primary IPv4 CIDR block of the VPD.
	//
	// example:
	//
	// 10.0.0.0/8
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The time when the activation code was created.
	//
	// example:
	//
	// 1678273219000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the Lingjun CIDR block.
	//
	// example:
	//
	// vpd-iv2zm1qf
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// The name of the Lingjun CIDR block.
	//
	// example:
	//
	// vpd-1
	VpdName *string `json:"VpdName,omitempty" xml:"VpdName,omitempty"`
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSubnetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// By default, popApi is not ignored and idempotent
	//
	// example:
	//
	// c5e3130a-d02f-11ec-a7d3-0242ac110005
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Paging Parameters: The current parameters are obsolete.
	//
	// example:
	//
	// false
	EnablePage *bool `json:"EnablePage,omitempty" xml:"EnablePage,omitempty"`
	// Paging Parameters: The current parameters are obsolete.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Paging Parameters: The current parameters are obsolete.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Lingjun connection instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vcc-cn-zvp2w222001
	VccId *string `json:"VccId,omitempty" xml:"VccId,omitempty"`
}

func (s GetVccRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVccRequest) GoString() string {
	return s.String()
}

func (s *GetVccRequest) SetClientToken(v string) *GetVccRequest {
	s.ClientToken = &v
	return s
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Content *GetVccResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// CAD09E47-B651-5206-B2DC-3AB78C8EB446
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVccResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVccResponseBody) GoString() string {
	return s.String()
}

func (s *GetVccResponseBody) SetAccessDeniedDetail(v string) *GetVccResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// Express Connect circuit access point ID:
	//
	// 	- **ap-cn-wulanchabu-jn-ts-A**: Ulanqab-Jining-A
	//
	// 	- **ap-cn-heyuan-yc-ts-SA127**: Heyuan-Yuancheng-A
	//
	// example:
	//
	// ap-cn-wulanchabu-jn-ts-A
	AccessPointId *string `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	// Alibaba Cloud route information list
	AliyunRouterInfo []*GetVccResponseBodyContentAliyunRouterInfo `json:"AliyunRouterInfo,omitempty" xml:"AliyunRouterInfo,omitempty" type:"Repeated"`
	// Whether Lingjun HUB has been bound to a network instance
	//
	// 	- **true**: Bound
	//
	// 	- **false**: unbound
	//
	// example:
	//
	// true
	AttachErStatus *bool `json:"AttachErStatus,omitempty" xml:"AttachErStatus,omitempty"`
	// The bandwidth of the port.
	//
	// example:
	//
	// 1G
	BandwidthStr *string `json:"BandwidthStr,omitempty" xml:"BandwidthStr,omitempty"`
	// BGP AS number
	//
	// example:
	//
	// 45644
	BgpAsn *string `json:"BgpAsn,omitempty" xml:"BgpAsn,omitempty"`
	// BGP CIDR block
	//
	// example:
	//
	// 10.4.0.0/24
	BgpCidr *string `json:"BgpCidr,omitempty" xml:"BgpCidr,omitempty"`
	// The ID of the CEN instance; [What is the CEN?](https://help.aliyun.com/document_detail/181681.html)
	//
	// You can call the [DescribeCens](https://help.aliyun.com/document_detail/468215.htm) to query the information of CEN instances under the current Alibaba Cloud account.
	//
	// example:
	//
	// cen-m2iskbojlvda5w65fp
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// Account to which the CEN belongs
	//
	// example:
	//
	// 1620939556166279
	CenOwnerId *string `json:"CenOwnerId,omitempty" xml:"CenOwnerId,omitempty"`
	// Lingjun Network Routing Information List
	CisRouterInfo []*GetVccResponseBodyContentCisRouterInfo `json:"CisRouterInfo,omitempty" xml:"CisRouterInfo,omitempty" type:"Repeated"`
	// Commodity code
	//
	// example:
	//
	// bccluster_cloudconnectionpre_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The connection mode. Valid values:
	//
	// 	- **VPC**
	//
	// 	- **CENTR**
	//
	// example:
	//
	// CENTR
	ConnectionType *string `json:"ConnectionType,omitempty" xml:"ConnectionType,omitempty"`
	// The time when the data address was created.
	//
	// example:
	//
	// 1648085472000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Current Node
	//
	// example:
	//
	// task-xxx-node-x
	CurrentNode *string `json:"CurrentNode,omitempty" xml:"CurrentNode,omitempty"`
	// Cycle
	//
	// example:
	//
	// 1
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// List of bound Lingjun HUB information
	ErInfos []*GetVccResponseBodyContentErInfos `json:"ErInfos,omitempty" xml:"ErInfos,omitempty" type:"Repeated"`
	// The time when the application expired.
	//
	// example:
	//
	// 1678379917000
	ExpirationDate *string `json:"ExpirationDate,omitempty" xml:"ExpirationDate,omitempty"`
	// The time when the agent was last modified.
	//
	// example:
	//
	// 1648085472000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The billing method for network usage.
	//
	// 	- **PayByTraffic**: pay-by-traffic
	//
	// 	- **PayByBandwidth**: pay-by-bandwidth
	//
	// example:
	//
	// PayByBandwidth
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The connectivity provider of the Express Connect circuit. Valid values:
	//
	// 	- **CO**: other connectivity providers in the Chinese mainland
	//
	// example:
	//
	// CO
	LineOperator *string `json:"LineOperator,omitempty" xml:"LineOperator,omitempty"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// You don\\"t have the permission of this operation, action=eflo:GetVcc, arn=acs:eflo:cn-heyuan:1263399219805497:vcc/vcc-cn-fhh3yxjwe01, resourceGroup=null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- **PREPAY**: subscription
	//
	// 	- **POSTPAY**: pay-as-you-go
	//
	// example:
	//
	// PrePay
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The port type of the Express Connect circuit. Valid values:
	//
	// 	- **100GBase-LR**: 100,000 megabytes of single-mode optical port (10 km)
	//
	// example:
	//
	// 100GBase-LR
	PortType *string `json:"PortType,omitempty" xml:"PortType,omitempty"`
	// The billing cycle. Valid values:
	//
	// 	- **Month**: Billed on a monthly basis
	//
	// 	- **Year**: Billed on an annual basis
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of your Alibaba Cloud resource group.
	//
	// For more information about resource groups, see [Resource groups](https://help.aliyun.com/document_detail/94475.htm).
	//
	// example:
	//
	// rg-aek2l4sq6l7unhi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Specification; value:
	//
	// 	- **Large**: Large
	//
	// example:
	//
	// Large
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The status of the cache reserve instance. Valid values:
	//
	// 	- **Available**: Normal.
	//
	// 	- **Not Available**: Not available.
	//
	// 	- **Executing**: The task is being executed.
	//
	// 	- **Deleting**: The account is being deleted
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag information.
	//
	// You can specify up to 20 tags.
	Tags []*GetVccResponseBodyContentTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the tenant.
	//
	// example:
	//
	// 1620939556166279
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The ID of the vSwitch. [Virtual Private Cloud VSwitch](https://help.aliyun.com/document_detail/100380.html).
	//
	// You can call the [DescribeVSwitches](https://help.aliyun.com/document_detail/35748.html) operation to query created vSwitches.
	//
	// example:
	//
	// vsw-uf6u8473r84e6n1n19he5
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// Information list of border routers
	VbrInfos []*GetVccResponseBodyContentVbrInfos `json:"VbrInfos,omitempty" xml:"VbrInfos,omitempty" type:"Repeated"`
	// The ID of the Lingjun connection instance.
	//
	// example:
	//
	// vcc-cn-cqf2xh40101
	VccId *string `json:"VccId,omitempty" xml:"VccId,omitempty"`
	// The name of the Lingjun connection instance.
	//
	// example:
	//
	// vcc-heyuan-backup
	VccName *string `json:"VccName,omitempty" xml:"VccName,omitempty"`
	// Virtual Private Cloud IDs; [What is Virtual Private Cloud](https://help.aliyun.com/document_detail/34217.html)
	//
	// You can call the [DescribeVpcs](https://help.aliyun.com/document_detail/35739.html#demo-0) operation to query the specified VPC.
	//
	// example:
	//
	// vpc-j6ctp4n75306phv5tmpsm
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// Lingjun network segment information (applicable to the scene where the old version of Lingjun connection is directly bound to Lingjun network segment)
	VpdBaseInfo *GetVccResponseBodyContentVpdBaseInfo `json:"VpdBaseInfo,omitempty" xml:"VpdBaseInfo,omitempty" type:"Struct"`
	// Lingjun CIDR block instance ID
	//
	// example:
	//
	// vpd-d3isyds4
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-wulanchabu-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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

func (s *GetVccResponseBodyContent) SetBgpAsn(v string) *GetVccResponseBodyContent {
	s.BgpAsn = &v
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

func (s *GetVccResponseBodyContent) SetCenOwnerId(v string) *GetVccResponseBodyContent {
	s.CenOwnerId = &v
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

func (s *GetVccResponseBodyContent) SetVbrInfos(v []*GetVccResponseBodyContentVbrInfos) *GetVccResponseBodyContent {
	s.VbrInfos = v
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
	// IPv4 address of Alibaba Cloud-side interconnection
	//
	// example:
	//
	// 169.254.248.30
	LocalGatewayIp *string `json:"LocalGatewayIp,omitempty" xml:"LocalGatewayIp,omitempty"`
	// Masking
	//
	// example:
	//
	// 255.255.255.248
	Mask *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	// Express Connect circuit ID
	//
	// example:
	//
	// pc-0jlof4bphlsnxbdztkvad
	PcId *string `json:"PcId,omitempty" xml:"PcId,omitempty"`
	// Lingjun Side Interconnection IPv4 Address
	//
	// example:
	//
	// 169.254.248.28
	PeerGatewayIp *string `json:"PeerGatewayIp,omitempty" xml:"PeerGatewayIp,omitempty"`
	// The ID of the VBR.
	//
	// example:
	//
	// vbr-2ze4i85p6vb9nwcan5xt0
	VbrId *string `json:"VbrId,omitempty" xml:"VbrId,omitempty"`
	// VLAN ID of the VBR
	//
	// example:
	//
	// 1042
	VlanId *string `json:"VlanId,omitempty" xml:"VlanId,omitempty"`
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
	// Leased Line Information List
	CcInfos []*GetVccResponseBodyContentCisRouterInfoCcInfos `json:"CcInfos,omitempty" xml:"CcInfos,omitempty" type:"Repeated"`
	// The ID of the on-cloud router instance.
	//
	// example:
	//
	// ccr-1ms84am0
	CcrId *string `json:"CcrId,omitempty" xml:"CcrId,omitempty"`
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
	// Leased Line ID
	//
	// example:
	//
	// cc-73aeex5o
	CcId *string `json:"CcId,omitempty" xml:"CcId,omitempty"`
	// Lingjun Side Interconnection IPv4 Address
	//
	// example:
	//
	// 169.254.248.26
	LocalGatewayIp *string `json:"LocalGatewayIp,omitempty" xml:"LocalGatewayIp,omitempty"`
	// Lingjun Side Interconnection IPv4 Address
	//
	// example:
	//
	// 169.254.248.30
	RemoteGatewayIp *string `json:"RemoteGatewayIp,omitempty" xml:"RemoteGatewayIp,omitempty"`
	// The state of the rule.
	//
	// example:
	//
	// Established
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Subnet mask
	//
	// example:
	//
	// 255.255.255.248
	SubnetMask *string `json:"SubnetMask,omitempty" xml:"SubnetMask,omitempty"`
	// Vlan ID of the leased line
	//
	// example:
	//
	// Ethernet1042
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
	// Connections
	//
	// example:
	//
	// 2
	Connections *int64 `json:"Connections,omitempty" xml:"Connections,omitempty"`
	// The time when the data address was created.
	//
	// example:
	//
	// 1678379917000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Description
	//
	// example:
	//
	// this is test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Lingjun HUB ID
	//
	// example:
	//
	// er-p68b0jwn
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// Lingjun HUB Instance Name
	//
	// example:
	//
	// er-1
	ErName *string `json:"ErName,omitempty" xml:"ErName,omitempty"`
	// The time when the agent was last modified.
	//
	// example:
	//
	// 1678379917000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Primary Zone
	//
	// example:
	//
	// cn-wulanchabu-b
	MasterZoneId *string `json:"MasterZoneId,omitempty" xml:"MasterZoneId,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// test message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Lingjun HUB Region Information
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Number of routing policy
	//
	// example:
	//
	// 2
	RouteMaps *int64 `json:"RouteMaps,omitempty" xml:"RouteMaps,omitempty"`
	// The status of the intervention entry. Valid value:
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 1620939556166277
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
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
	// The tag key.
	//
	// You cannot specify an empty string as a tag key. It can be up to 64 characters in length and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// You can specify at most 20 tag keys in each call.
	//
	// example:
	//
	// tag-vcc
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag that is added to the resource.
	//
	// The tag value can be empty or a string of up to 128 characters. It cannot start with aliyun or acs:, and cannot contain http:// or https://.
	//
	// Each key-value pair must be unique. You can specify values for at most 20 tag keys in each call.
	//
	// example:
	//
	// vcc-group-1
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

type GetVccResponseBodyContentVbrInfos struct {
	// CEN ID
	//
	// example:
	//
	// cen-cx0qua8q6cm4z9****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The time when the data address was created.
	//
	// example:
	//
	// 1683250981000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the agent was last modified.
	//
	// example:
	//
	// 1673578603000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The status of the VBR. Valid values:
	//
	// 	- unconfirmed
	//
	// 	- active: The VPN gateway is in a normal state.
	//
	// 	- terminating: The connection is being terminated.
	//
	// 	- terminated: The connection is terminated.
	//
	// 	- recovering: The task is being recovered.
	//
	// 	- deleting: The GDN is being deleted.
	//
	// 	- Available: The service is available.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// BGP neighbor information list
	VbrBgpPeers []*GetVccResponseBodyContentVbrInfosVbrBgpPeers `json:"VbrBgpPeers,omitempty" xml:"VbrBgpPeers,omitempty" type:"Repeated"`
	// The ID of the border router.
	//
	// example:
	//
	// vbr-wz96agu9h3d50z****
	VbrId *string `json:"VbrId,omitempty" xml:"VbrId,omitempty"`
}

func (s GetVccResponseBodyContentVbrInfos) String() string {
	return tea.Prettify(s)
}

func (s GetVccResponseBodyContentVbrInfos) GoString() string {
	return s.String()
}

func (s *GetVccResponseBodyContentVbrInfos) SetCenId(v string) *GetVccResponseBodyContentVbrInfos {
	s.CenId = &v
	return s
}

func (s *GetVccResponseBodyContentVbrInfos) SetGmtCreate(v string) *GetVccResponseBodyContentVbrInfos {
	s.GmtCreate = &v
	return s
}

func (s *GetVccResponseBodyContentVbrInfos) SetGmtModified(v string) *GetVccResponseBodyContentVbrInfos {
	s.GmtModified = &v
	return s
}

func (s *GetVccResponseBodyContentVbrInfos) SetStatus(v string) *GetVccResponseBodyContentVbrInfos {
	s.Status = &v
	return s
}

func (s *GetVccResponseBodyContentVbrInfos) SetVbrBgpPeers(v []*GetVccResponseBodyContentVbrInfosVbrBgpPeers) *GetVccResponseBodyContentVbrInfos {
	s.VbrBgpPeers = v
	return s
}

func (s *GetVccResponseBodyContentVbrInfos) SetVbrId(v string) *GetVccResponseBodyContentVbrInfos {
	s.VbrId = &v
	return s
}

type GetVccResponseBodyContentVbrInfosVbrBgpPeers struct {
	// BGP Group ID
	//
	// example:
	//
	// bgpg-2ze2sit2vakrkapvy****
	BgpGroupId *string `json:"BgpGroupId,omitempty" xml:"BgpGroupId,omitempty"`
	// BGP peer ID
	//
	// example:
	//
	// bgp-uf6heugif9enu48rj****
	BgpPeerId *string `json:"BgpPeerId,omitempty" xml:"BgpPeerId,omitempty"`
	// Peer AS No.
	//
	// example:
	//
	// 98765****
	PeerAsn *string `json:"PeerAsn,omitempty" xml:"PeerAsn,omitempty"`
	// BGP peer IP address
	//
	// example:
	//
	// 169.254.****
	PeerIpAddress *string `json:"PeerIpAddress,omitempty" xml:"PeerIpAddress,omitempty"`
	// The status of the BGP peer. Valid values:
	//
	// 	- Pending: pending
	//
	// 	- Available: The route is available.
	//
	// 	- Modifying: being modified
	//
	// 	- Deleting: The IPv4 gateway is being deleted.
	//
	// 	- Deleted
	//
	// 	- Not Available
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetVccResponseBodyContentVbrInfosVbrBgpPeers) String() string {
	return tea.Prettify(s)
}

func (s GetVccResponseBodyContentVbrInfosVbrBgpPeers) GoString() string {
	return s.String()
}

func (s *GetVccResponseBodyContentVbrInfosVbrBgpPeers) SetBgpGroupId(v string) *GetVccResponseBodyContentVbrInfosVbrBgpPeers {
	s.BgpGroupId = &v
	return s
}

func (s *GetVccResponseBodyContentVbrInfosVbrBgpPeers) SetBgpPeerId(v string) *GetVccResponseBodyContentVbrInfosVbrBgpPeers {
	s.BgpPeerId = &v
	return s
}

func (s *GetVccResponseBodyContentVbrInfosVbrBgpPeers) SetPeerAsn(v string) *GetVccResponseBodyContentVbrInfosVbrBgpPeers {
	s.PeerAsn = &v
	return s
}

func (s *GetVccResponseBodyContentVbrInfosVbrBgpPeers) SetPeerIpAddress(v string) *GetVccResponseBodyContentVbrInfosVbrBgpPeers {
	s.PeerIpAddress = &v
	return s
}

func (s *GetVccResponseBodyContentVbrInfosVbrBgpPeers) SetStatus(v string) *GetVccResponseBodyContentVbrInfosVbrBgpPeers {
	s.Status = &v
	return s
}

type GetVccResponseBodyContentVpdBaseInfo struct {
	// Network address segment
	//
	// example:
	//
	// 172.18.0.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The time when the data address was created.
	//
	// example:
	//
	// 1678379917000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Lingjun CIDR block instance ID
	//
	// example:
	//
	// vpd-ppdunxzc
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// Lingjun CIDR block instance name
	//
	// example:
	//
	// yzp-rg-test3
	VpdName *string `json:"VpdName,omitempty" xml:"VpdName,omitempty"`
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
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVccResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// Lingjun HUB Instance ID
	//
	// example:
	//
	// er-aueyxxsy
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// Authorized Resource Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// grant-rule-jaj34d75h01
	GrantRuleId *string `json:"GrantRuleId,omitempty" xml:"GrantRuleId,omitempty"`
	// Authorized Tenant ID
	//
	// example:
	//
	// 1620939556166277
	GrantTenantId *string `json:"GrantTenantId,omitempty" xml:"GrantTenantId,omitempty"`
	// Authorized Instance ID
	//
	// example:
	//
	// vcc-cn-jaj34d75h01
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region. This parameter is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Content *GetVccGrantRuleResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// 0901F411-28FA-5B9C-BAEE-7776463FF0DC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVccGrantRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVccGrantRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetVccGrantRuleResponseBody) SetAccessDeniedDetail(v string) *GetVccGrantRuleResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// The time when the data address was created.
	//
	// example:
	//
	// 1648085472000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Lingjun HUB Instance ID
	//
	// example:
	//
	// er-aueyxxsy
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// Authorized Resource ID
	//
	// example:
	//
	// grant-rule-jaj34d75h01
	GrantRuleId *string `json:"GrantRuleId,omitempty" xml:"GrantRuleId,omitempty"`
	// Authorized Tenant ID
	//
	// example:
	//
	// 1620939556166277
	GrantTenantId *string `json:"GrantTenantId,omitempty" xml:"GrantTenantId,omitempty"`
	// Network Instance ID
	//
	// example:
	//
	// vcc-cn-jaj34d75h01
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Network Instance Name
	//
	// example:
	//
	// vcc-1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Network Product Code:
	//
	// 	- **VPD**: Lingjun CIDR block
	//
	// 	- **VCC**: Lingjun Connection
	//
	// example:
	//
	// VCC
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-aek2l4sq6l7unhi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 1620939556166279
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// Whether the current authorization information has been used; optional values:
	//
	// 	- **true**: Used
	//
	// 	- **false**: Not used
	//
	// example:
	//
	// false
	Used *bool `json:"Used,omitempty" xml:"Used,omitempty"`
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

func (s *GetVccGrantRuleResponseBodyContent) SetResourceGroupId(v string) *GetVccGrantRuleResponseBodyContent {
	s.ResourceGroupId = &v
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVccGrantRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Lingjun Connection ID
	//
	// This parameter is required.
	//
	// example:
	//
	// vcc-cn-zvp2w222001
	VccId *string `json:"VccId,omitempty" xml:"VccId,omitempty"`
	// The ID of the route entry.
	//
	// This parameter is required.
	//
	// example:
	//
	// vcc-rte-31ocvdhq
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Content *GetVccRouteEntryResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// BDBCC783-84CA-5733-8EEA-645C88B9009C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVccRouteEntryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVccRouteEntryResponseBody) GoString() string {
	return s.String()
}

func (s *GetVccRouteEntryResponseBody) SetAccessDeniedDetail(v string) *GetVccRouteEntryResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// Destination CIDR Block
	//
	// example:
	//
	// 0.0.0.0/0
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The time when the agent was last modified.
	//
	// example:
	//
	// 1648085472000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Next Hop Instance
	//
	// example:
	//
	// local
	NextHopId *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// Next Hop Type
	//
	// example:
	//
	// VCC
	NextHopType *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-aek2l4sq6l7u***
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Route type
	//
	// example:
	//
	// System
	RouteType *string `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
	// The status of the intervention entry. Valid value:
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 1655449505171
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The ID of the Lingjun connection instance.
	//
	// example:
	//
	// vcc-cn-zvp2w222001
	VccId *string `json:"VccId,omitempty" xml:"VccId,omitempty"`
	// The ID of the route entry.
	//
	// example:
	//
	// vcc-rte-31ocvdhq
	VccRouteEntryId *string `json:"VccRouteEntryId,omitempty" xml:"VccRouteEntryId,omitempty"`
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

func (s *GetVccRouteEntryResponseBodyContent) SetMessage(v string) *GetVccRouteEntryResponseBodyContent {
	s.Message = &v
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

func (s *GetVccRouteEntryResponseBodyContent) SetResourceGroupId(v string) *GetVccRouteEntryResponseBodyContent {
	s.ResourceGroupId = &v
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVccRouteEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the VPD instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpd-ze3na0wf
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Content *GetVpdResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A88DFED5-24B7-5A3E-87DE-380BF06F3C90
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVpdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVpdResponseBody) GoString() string {
	return s.String()
}

func (s *GetVpdResponseBody) SetAccessDeniedDetail(v string) *GetVpdResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// Whether the Lingjun HUB(ER) has been bound.
	//
	// 	- **true**: ER is bound.
	//
	// 	- **false**: No ER is bound.
	//
	// example:
	//
	// true
	AttachErStatus *bool `json:"AttachErStatus,omitempty" xml:"AttachErStatus,omitempty"`
	// The CIDR block.
	//
	// example:
	//
	// 10.0.0.0/8
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The time when the activation code was created.
	//
	// example:
	//
	// 1678273219000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The information of the bound Lingjun HUB(ER).
	ErInfos []*GetVpdResponseBodyContentErInfos `json:"ErInfos,omitempty" xml:"ErInfos,omitempty" type:"Repeated"`
	// The time when the O\\&M task was modified.
	//
	// example:
	//
	// 2023-10-25 15:57:16
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The number of NCs.
	//
	// example:
	//
	// 16
	NcCount *int32 `json:"NcCount,omitempty" xml:"NcCount,omitempty"`
	// Number of Lingjun network interface controller.
	//
	// example:
	//
	// 1
	NetworkInterfaceCount *int32 `json:"NetworkInterfaceCount,omitempty" xml:"NetworkInterfaceCount,omitempty"`
	// The total number of secondary private IP addresses.
	//
	// example:
	//
	// 10
	PrivateIpCount *int64 `json:"PrivateIpCount,omitempty" xml:"PrivateIpCount,omitempty"`
	// The total quota information.
	//
	// example:
	//
	// 10
	Quota *int32 `json:"Quota,omitempty" xml:"Quota,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of your Alibaba Cloud resource group.
	//
	// For more information about resource groups, see [Resource groups](https://help.aliyun.com/document_detail/94475.htm?spm=a2c4g.11186623.0.0.29e15d7akXhpuu).
	//
	// example:
	//
	// rg-aek2l4sq6l7unhi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The list of additional CIDR blocks.
	SecondaryCidrBlocks []*string `json:"SecondaryCidrBlocks,omitempty" xml:"SecondaryCidrBlocks,omitempty" type:"Repeated"`
	// Internal Service CIDR block.
	//
	// example:
	//
	// 169.254.252.0/23
	ServiceCidr *string `json:"ServiceCidr,omitempty" xml:"ServiceCidr,omitempty"`
	// The current state of the instance.
	//
	// Valid value:
	//
	// 	- Not Available: Not Available.
	//
	// 	- Available: Normal: Available: Normal.
	//
	// 	- Deleting: Deleting: Deleting: Deleting.
	//
	// 	- Executing: executing: Executing: executing.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of subnets.
	//
	// example:
	//
	// 1
	SubnetCount *int64 `json:"SubnetCount,omitempty" xml:"SubnetCount,omitempty"`
	// The tag information.
	//
	// You can specify up to 20 tags.
	Tags []*GetVpdResponseBodyContentTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The tenant ID.
	//
	// example:
	//
	// 1620939556166279
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The ID of the VPD instance.
	//
	// example:
	//
	// vpd-fuliephf
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// The name of the Lingjun CIDR block.
	//
	// example:
	//
	// vpd-1
	VpdName *string `json:"VpdName,omitempty" xml:"VpdName,omitempty"`
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

func (s *GetVpdResponseBodyContent) SetNcCount(v int32) *GetVpdResponseBodyContent {
	s.NcCount = &v
	return s
}

func (s *GetVpdResponseBodyContent) SetNetworkInterfaceCount(v int32) *GetVpdResponseBodyContent {
	s.NetworkInterfaceCount = &v
	return s
}

func (s *GetVpdResponseBodyContent) SetPrivateIpCount(v int64) *GetVpdResponseBodyContent {
	s.PrivateIpCount = &v
	return s
}

func (s *GetVpdResponseBodyContent) SetQuota(v int32) *GetVpdResponseBodyContent {
	s.Quota = &v
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

func (s *GetVpdResponseBodyContent) SetSecondaryCidrBlocks(v []*string) *GetVpdResponseBodyContent {
	s.SecondaryCidrBlocks = v
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
	// The number of connections.
	//
	// example:
	//
	// 2
	Connections *int64 `json:"Connections,omitempty" xml:"Connections,omitempty"`
	// The time when the activation code was created.
	//
	// example:
	//
	// 1678273219000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the synchronization task.
	//
	// example:
	//
	// Restore verifying
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the Elastic Router (ER) instance.
	//
	// example:
	//
	// er-a7rqv1rq
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// Elastic Router (ER) Instance Name
	//
	// example:
	//
	// er-1
	ErName *string `json:"ErName,omitempty" xml:"ErName,omitempty"`
	// The time when the O\\&M task was modified.
	//
	// example:
	//
	// 1678273219000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The primary zone.
	//
	// example:
	//
	// cn-wulanchabu-b
	MasterZoneId *string `json:"MasterZoneId,omitempty" xml:"MasterZoneId,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the region to which the Elastic Router (ER) belongs.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of routing policy.
	//
	// example:
	//
	// 2
	RouteMaps *int64 `json:"RouteMaps,omitempty" xml:"RouteMaps,omitempty"`
	// The task status.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// t464p4fql1bog
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
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
	// The tag key.
	//
	// You cannot specify an empty string as a tag key. It can be up to 64 characters in length and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// You can specify at most 20 tag keys in each call.
	//
	// example:
	//
	// tag-subent-region
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag that is added to the resource.
	//
	// The tag value can be empty or a string of up to 128 characters. It cannot start with aliyun or acs:, and cannot contain http:// or https://.
	//
	// Each key-value pair must be unique. You can specify values for at most 20 tag keys in each call.
	//
	// example:
	//
	// cn-wulanchabu
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
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVpdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// Lingjun HUB Instance Id
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// Authorized Resource Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// grant-rule-xrgvqazb
	GrantRuleId *string `json:"GrantRuleId,omitempty" xml:"GrantRuleId,omitempty"`
	// Authorized Tenant ID
	//
	// example:
	//
	// 1620939556166277
	GrantTenantId *string `json:"GrantTenantId,omitempty" xml:"GrantTenantId,omitempty"`
	// Authorized Instance ID
	//
	// example:
	//
	// vpd-xxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region. This parameter is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Content *GetVpdGrantRuleResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// 9C50C9CD-E799-54DA-BA7A-1FAF3DF80857
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVpdGrantRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVpdGrantRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetVpdGrantRuleResponseBody) SetAccessDeniedDetail(v string) *GetVpdGrantRuleResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// The time when the data address was created.
	//
	// example:
	//
	// 1648085472000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Lingjun HUB Instance ID
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// Authorized Resource ID
	//
	// example:
	//
	// grant-rule-xxxxxx
	GrantRuleId *string `json:"GrantRuleId,omitempty" xml:"GrantRuleId,omitempty"`
	// Authorized Tenant ID
	//
	// example:
	//
	// 1620939556166277
	GrantTenantId *string `json:"GrantTenantId,omitempty" xml:"GrantTenantId,omitempty"`
	// Network Instance ID
	//
	// example:
	//
	// vpd-xxxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Network Instance Name
	//
	// example:
	//
	// vpd-lingjun
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Network Product Code:
	//
	// 	- **VPD**: Lingjun CIDR block
	//
	// 	- **VCC**: Lingjun Connection
	//
	// example:
	//
	// VPD
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-aek2l4sq6l7u***
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 1655449505171
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// Whether the current authorization information has been used; default is false
	//
	// example:
	//
	// 0
	Used *bool `json:"Used,omitempty" xml:"Used,omitempty"`
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

func (s *GetVpdGrantRuleResponseBodyContent) SetResourceGroupId(v string) *GetVpdGrantRuleResponseBodyContent {
	s.ResourceGroupId = &v
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVpdGrantRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Lingjun CIDR block instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// vpd-iv2zm1qf
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// The ID of the route entry instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpd-rte-toekyqel
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Content *GetVpdRouteEntryResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// 9C50C9CD-E799-54DA-BA7A-1FAF3DF80857
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVpdRouteEntryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVpdRouteEntryResponseBody) GoString() string {
	return s.String()
}

func (s *GetVpdRouteEntryResponseBody) SetAccessDeniedDetail(v string) *GetVpdRouteEntryResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// Destination CIDR block
	//
	// example:
	//
	// 0.0.0.0/0
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The time when the agent was last modified.
	//
	// example:
	//
	// 1678273219000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Next Hop Instance
	//
	// example:
	//
	// er-bmlqiym1
	NextHopId *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// Next Hop Instance Type
	//
	// example:
	//
	// ER
	NextHopType *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-acfmv7mcq63uyhq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Route type
	//
	// example:
	//
	// BGP
	RouteType *string `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
	// The status of the intervention entry. Valid value:
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 1655449505171
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// Lingjun CIDR block instance ID
	//
	// example:
	//
	// vpd-ze3na0wf
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// Lingjun CIDR block route entry ID
	//
	// example:
	//
	// vpd-rte-toekyqel
	VpdRouteEntryId *string `json:"VpdRouteEntryId,omitempty" xml:"VpdRouteEntryId,omitempty"`
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

func (s *GetVpdRouteEntryResponseBodyContent) SetResourceGroupId(v string) *GetVpdRouteEntryResponseBodyContent {
	s.ResourceGroupId = &v
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVpdRouteEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The resource group ID.
	//
	// For more information about resource groups, see [Resource groups](https://help.aliyun.com/document_detail/94475.htm?spm=a2c4g.11186623.0.0.29e15d7akXhpuu).
	//
	// example:
	//
	// rg-acfmxhucx5ewuwy
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Content *InitializeVccResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E30DA7CB-03D0-51EB-8F18-856B99987E18
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InitializeVccResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InitializeVccResponseBody) GoString() string {
	return s.String()
}

func (s *InitializeVccResponseBody) SetAccessDeniedDetail(v string) *InitializeVccResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// The request ID.
	//
	// example:
	//
	// E30DA7CB-03D0-51EB-8F18-856B99987E18
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Linked Role of Lingjun Connection Instance (AliyunServiceRoleForEfloVcc)
	//
	// example:
	//
	// CloudConnectionOperationRole
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitializeVccResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ListElasticNetworkInterfacesRequest struct {
	// Lingjun Elastic Network Interface ID
	//
	// example:
	//
	// leni-1234****
	ElasticNetworkInterfaceId *string `json:"ElasticNetworkInterfaceId,omitempty" xml:"ElasticNetworkInterfaceId,omitempty"`
	// The IP address of the BE cluster.
	//
	// example:
	//
	// 10.0.0.1
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The network type.
	//
	// Valid value:
	//
	// 	- Tenant: Tenant.
	//
	// 	- VPC
	//
	// example:
	//
	// tenant
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// e01-cn-lbj3aej****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The page number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 20.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the enterprise-level snapshot policy.
	//
	// Valid value:
	//
	// 	- Create Failed: the creation failure.
	//
	// 	- Delete Failed: the that failed to be deleted.
	//
	// 	- Executing
	//
	// 	- Available: The template is available.
	//
	// 	- Deleting
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the variable.
	//
	// Valid value:
	//
	// 	- CUSTOM: custom type.
	//
	// 	- DEFAULT: system type.
	//
	// example:
	//
	// DEFAULT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-uf6u8473r84e9****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-uf6aa4ddo97fr****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-wulanchabu-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListElasticNetworkInterfacesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListElasticNetworkInterfacesRequest) GoString() string {
	return s.String()
}

func (s *ListElasticNetworkInterfacesRequest) SetElasticNetworkInterfaceId(v string) *ListElasticNetworkInterfacesRequest {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *ListElasticNetworkInterfacesRequest) SetIp(v string) *ListElasticNetworkInterfacesRequest {
	s.Ip = &v
	return s
}

func (s *ListElasticNetworkInterfacesRequest) SetNetworkType(v string) *ListElasticNetworkInterfacesRequest {
	s.NetworkType = &v
	return s
}

func (s *ListElasticNetworkInterfacesRequest) SetNodeId(v string) *ListElasticNetworkInterfacesRequest {
	s.NodeId = &v
	return s
}

func (s *ListElasticNetworkInterfacesRequest) SetPageNumber(v int32) *ListElasticNetworkInterfacesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListElasticNetworkInterfacesRequest) SetPageSize(v int32) *ListElasticNetworkInterfacesRequest {
	s.PageSize = &v
	return s
}

func (s *ListElasticNetworkInterfacesRequest) SetRegionId(v string) *ListElasticNetworkInterfacesRequest {
	s.RegionId = &v
	return s
}

func (s *ListElasticNetworkInterfacesRequest) SetStatus(v string) *ListElasticNetworkInterfacesRequest {
	s.Status = &v
	return s
}

func (s *ListElasticNetworkInterfacesRequest) SetType(v string) *ListElasticNetworkInterfacesRequest {
	s.Type = &v
	return s
}

func (s *ListElasticNetworkInterfacesRequest) SetVSwitchId(v string) *ListElasticNetworkInterfacesRequest {
	s.VSwitchId = &v
	return s
}

func (s *ListElasticNetworkInterfacesRequest) SetVpcId(v string) *ListElasticNetworkInterfacesRequest {
	s.VpcId = &v
	return s
}

func (s *ListElasticNetworkInterfacesRequest) SetZoneId(v string) *ListElasticNetworkInterfacesRequest {
	s.ZoneId = &v
	return s
}

type ListElasticNetworkInterfacesResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Content *ListElasticNetworkInterfacesResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The return message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// AC8C713A-A9F4-5984-A5E1-76496DF35153
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListElasticNetworkInterfacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListElasticNetworkInterfacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListElasticNetworkInterfacesResponseBody) SetAccessDeniedDetail(v string) *ListElasticNetworkInterfacesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBody) SetCode(v int32) *ListElasticNetworkInterfacesResponseBody {
	s.Code = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBody) SetContent(v *ListElasticNetworkInterfacesResponseBodyContent) *ListElasticNetworkInterfacesResponseBody {
	s.Content = v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBody) SetMessage(v string) *ListElasticNetworkInterfacesResponseBody {
	s.Message = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBody) SetRequestId(v string) *ListElasticNetworkInterfacesResponseBody {
	s.RequestId = &v
	return s
}

type ListElasticNetworkInterfacesResponseBodyContent struct {
	// lingjun Elastic Network Interface information list
	Data []*ListElasticNetworkInterfacesResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListElasticNetworkInterfacesResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s ListElasticNetworkInterfacesResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListElasticNetworkInterfacesResponseBodyContent) SetData(v []*ListElasticNetworkInterfacesResponseBodyContentData) *ListElasticNetworkInterfacesResponseBodyContent {
	s.Data = v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBodyContent) SetTotal(v int64) *ListElasticNetworkInterfacesResponseBodyContent {
	s.Total = &v
	return s
}

type ListElasticNetworkInterfacesResponseBodyContentData struct {
	// The time when the data address was created.
	//
	// example:
	//
	// 1601176751000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The instance description.
	//
	// example:
	//
	// No description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Lingjun Elastic Network Interface ID
	//
	// example:
	//
	// leni-1234****
	ElasticNetworkInterfaceId *string `json:"ElasticNetworkInterfaceId,omitempty" xml:"ElasticNetworkInterfaceId,omitempty"`
	// vswitch gateway address
	//
	// example:
	//
	// 172.16.****
	Gateway *string `json:"Gateway,omitempty" xml:"Gateway,omitempty"`
	// The time when the agent was last modified.
	//
	// example:
	//
	// 1640187007000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The IP address of the BE cluster.
	//
	// example:
	//
	// 10.0.0.13
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// mac address
	//
	// example:
	//
	// E0:01:A6:4A:6A:D0
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// vswitch mask bits
	//
	// example:
	//
	// 24
	Mask *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	// The error message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// e01-cn-uax37m1****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-f8z4wr1b41x3qsc9****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The status of the intervention entry. Valid value:
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// network interface controller type, the default type DEFAULT cannot be manually released
	//
	// Valid value:
	//
	// 	- CUSTOM: custom type.
	//
	// 	- DEFAULT: system type.
	//
	// example:
	//
	// DEFAULT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-uf6u8473r84e9****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-f8ziirfl9k25h2qn7****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-wulanchabu-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListElasticNetworkInterfacesResponseBodyContentData) String() string {
	return tea.Prettify(s)
}

func (s ListElasticNetworkInterfacesResponseBodyContentData) GoString() string {
	return s.String()
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) SetCreateTime(v string) *ListElasticNetworkInterfacesResponseBodyContentData {
	s.CreateTime = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) SetDescription(v string) *ListElasticNetworkInterfacesResponseBodyContentData {
	s.Description = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) SetElasticNetworkInterfaceId(v string) *ListElasticNetworkInterfacesResponseBodyContentData {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) SetGateway(v string) *ListElasticNetworkInterfacesResponseBodyContentData {
	s.Gateway = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) SetGmtModified(v string) *ListElasticNetworkInterfacesResponseBodyContentData {
	s.GmtModified = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) SetIp(v string) *ListElasticNetworkInterfacesResponseBodyContentData {
	s.Ip = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) SetMac(v string) *ListElasticNetworkInterfacesResponseBodyContentData {
	s.Mac = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) SetMask(v string) *ListElasticNetworkInterfacesResponseBodyContentData {
	s.Mask = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) SetMessage(v string) *ListElasticNetworkInterfacesResponseBodyContentData {
	s.Message = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) SetNodeId(v string) *ListElasticNetworkInterfacesResponseBodyContentData {
	s.NodeId = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) SetRegionId(v string) *ListElasticNetworkInterfacesResponseBodyContentData {
	s.RegionId = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) SetSecurityGroupId(v string) *ListElasticNetworkInterfacesResponseBodyContentData {
	s.SecurityGroupId = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) SetStatus(v string) *ListElasticNetworkInterfacesResponseBodyContentData {
	s.Status = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) SetType(v string) *ListElasticNetworkInterfacesResponseBodyContentData {
	s.Type = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) SetVSwitchId(v string) *ListElasticNetworkInterfacesResponseBodyContentData {
	s.VSwitchId = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) SetVpcId(v string) *ListElasticNetworkInterfacesResponseBodyContentData {
	s.VpcId = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponseBodyContentData) SetZoneId(v string) *ListElasticNetworkInterfacesResponseBodyContentData {
	s.ZoneId = &v
	return s
}

type ListElasticNetworkInterfacesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListElasticNetworkInterfacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListElasticNetworkInterfacesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListElasticNetworkInterfacesResponse) GoString() string {
	return s.String()
}

func (s *ListElasticNetworkInterfacesResponse) SetHeaders(v map[string]*string) *ListElasticNetworkInterfacesResponse {
	s.Headers = v
	return s
}

func (s *ListElasticNetworkInterfacesResponse) SetStatusCode(v int32) *ListElasticNetworkInterfacesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponse) SetBody(v *ListElasticNetworkInterfacesResponseBody) *ListElasticNetworkInterfacesResponse {
	s.Body = v
	return s
}

type ListErAttachmentsRequest struct {
	// Whether to automatically receive all routes from all instances under this Lingjun HUB. Valid values:
	//
	// 	- **true**: received automatically.
	//
	// 	- **false**: Not received.
	//
	// example:
	//
	// true
	AutoReceiveAllRoute *bool `json:"AutoReceiveAllRoute,omitempty" xml:"AutoReceiveAllRoute,omitempty"`
	// Specifies whether to enable paged query. Valid values:
	//
	// 	- **true**: enables paged query.
	//
	// 	- **false**: Paged query is not enabled.
	//
	// example:
	//
	// false
	EnablePage *bool `json:"EnablePage,omitempty" xml:"EnablePage,omitempty"`
	// The ID of the network instance connection
	//
	// example:
	//
	// er-attachment-i1ioibyf
	ErAttachmentId *string `json:"ErAttachmentId,omitempty" xml:"ErAttachmentId,omitempty"`
	// The name of the network instance connection.
	//
	// example:
	//
	// vcc-cn-209300qha01
	ErAttachmentName *string `json:"ErAttachmentName,omitempty" xml:"ErAttachmentName,omitempty"`
	// The ID of the Lingjun HUB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// The ID of the network instance. Valid values: **VPD*	- and **VCC**.
	//
	// For more information, see [What is Lingjun?](https://help.aliyun.com/document_detail/444430.html)
	//
	// You can query **Lingjun CIDR blocks*	- and **Lingjun connections*	- by [ListVpds](https://help.aliyun.com/document_detail/2331077.html) and [ListVccs](https://help.aliyun.com/document_detail/2399526.html?) respectively.
	//
	// example:
	//
	// vcc-cn-209300qha01
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The mitigation plan of the instance. Valid values:
	//
	// 	- **VPD**: indicates the Lingjun CIDR block.
	//
	// 	- **VCC**: indicates a Lingjun connection.
	//
	// example:
	//
	// VCC
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The page number to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-aekzb3n5lgk2ieq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the tenant to which the instance belongs.
	//
	// example:
	//
	// 1111156667137893
	ResourceTenantId *string `json:"ResourceTenantId,omitempty" xml:"ResourceTenantId,omitempty"`
	// The status of the CLB instance. Valid values:
	//
	// 	- **Available**: Normal.
	//
	// 	- **Not Available**: Not available.
	//
	// 	- **Executing**: The task is being executed.
	//
	// 	- **Deleting**: The account is being deleted
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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

func (s *ListErAttachmentsRequest) SetResourceGroupId(v string) *ListErAttachmentsRequest {
	s.ResourceGroupId = &v
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Content *ListErAttachmentsResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is displayed.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3D9D6E7B-365B-5200-BFA6-9B79E269058C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListErAttachmentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListErAttachmentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListErAttachmentsResponseBody) SetAccessDeniedDetail(v string) *ListErAttachmentsResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// The list of Lingjun HUB network instances.
	Data []*ListErAttachmentsResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The total number of entries that are returned.
	//
	// example:
	//
	// 0
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
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
	// Whether to cross accounts. Valid values:
	//
	// 	- **true**: The network instance is a cross-account resource.
	//
	// 	- **false**: The current network instance is a resource of the current account.
	//
	// example:
	//
	// false
	Across *bool `json:"Across,omitempty" xml:"Across,omitempty"`
	// Whether to automatically receive all routes from all instances under this Lingjun HUB. Valid values:
	//
	// 	- **true**: received automatically.
	//
	// 	- **false**: Not received.
	//
	// example:
	//
	// true
	AutoReceiveAllRoute *bool `json:"AutoReceiveAllRoute,omitempty" xml:"AutoReceiveAllRoute,omitempty"`
	// The time when the activation code was created.
	//
	// example:
	//
	// 1669734207000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the Lingjun HUB network instance.
	//
	// example:
	//
	// er-attachment-i1ioibyf
	ErAttachmentId *string `json:"ErAttachmentId,omitempty" xml:"ErAttachmentId,omitempty"`
	// The name of the Lingjun HUB network instance.
	//
	// example:
	//
	// vcc-cn-209300qha01
	ErAttachmentName *string `json:"ErAttachmentName,omitempty" xml:"ErAttachmentName,omitempty"`
	// The ID of the Lingjun HUB instance.
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// The time when the O\\&M task was modified.
	//
	// example:
	//
	// 1640187007000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the network instance. Valid values: **VPD*	- and **VCC**.
	//
	// For more information, see [What is Lingjun?](https://help.aliyun.com/document_detail/444430.html)
	//
	// You can query **Lingjun CIDR blocks*	- and **Lingjun connections*	- by [ListVpds](https://help.aliyun.com/document_detail/2331077.html) and [ListVccs](https://help.aliyun.com/document_detail/2399526.html) respectively.
	//
	// example:
	//
	// vcc-cn-209300qha02
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// vcc-wulanchabu-main
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The database type. Valid values:
	//
	// 	- **VPD**: indicates the Lingjun CIDR block.
	//
	// 	- **VCC**: indicates a Lingjun connection.
	//
	// example:
	//
	// VCC
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The returned message.
	//
	// example:
	//
	// test
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Lingjun HUB region information.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-aekzlki4ehfse4y
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the tenant to which the resource belongs.
	//
	// example:
	//
	// 1111156667137893
	ResourceTenantId *string `json:"ResourceTenantId,omitempty" xml:"ResourceTenantId,omitempty"`
	// The status of the cache reserve instance. Valid values:
	//
	// 	- **Available**: Normal.
	//
	// 	- **Not Available**: Not available.
	//
	// 	- **Executing**: The task is being executed.
	//
	// 	- **Deleting**: The account is being deleted
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 1655449505171
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
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

func (s *ListErAttachmentsResponseBodyContentData) SetResourceGroupId(v string) *ListErAttachmentsResponseBodyContentData {
	s.ResourceGroupId = &v
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListErAttachmentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// Destination CIDR Block
	//
	// example:
	//
	// 0.0.0.0/0
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// Specifies whether to enable pagination query.
	//
	// example:
	//
	// false
	EnablePage *bool `json:"EnablePage,omitempty" xml:"EnablePage,omitempty"`
	// Lingjun HUB ID
	//
	// This parameter is required.
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// Filter 32 detailed CIDR blocks. Default value: true
	//
	// example:
	//
	// true
	IgnoreDetailedRouteEntry *bool `json:"IgnoreDetailedRouteEntry,omitempty" xml:"IgnoreDetailedRouteEntry,omitempty"`
	// Network Instance ID
	//
	// example:
	//
	// vcc-cn-209300qha01
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Next Hop Instance
	//
	// example:
	//
	// vcc-cn-209300qha01
	NextHopId *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// Next Hop Instance Type
	//
	// example:
	//
	// VCC
	NextHopType *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	// The page number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-acfmyuzlx2iihcy
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Route type
	//
	// example:
	//
	// VCC
	RouteType *string `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
	// The status of the enterprise-level snapshot policy.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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

func (s *ListErRouteEntriesRequest) SetIgnoreDetailedRouteEntry(v bool) *ListErRouteEntriesRequest {
	s.IgnoreDetailedRouteEntry = &v
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

func (s *ListErRouteEntriesRequest) SetResourceGroupId(v string) *ListErRouteEntriesRequest {
	s.ResourceGroupId = &v
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Content *ListErRouteEntriesResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// A88DFED5-24B7-5A3E-87DE-380BF06F3C90
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListErRouteEntriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListErRouteEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListErRouteEntriesResponseBody) SetAccessDeniedDetail(v string) *ListErRouteEntriesResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// Lingjun HUB Route Entry Information List
	Data []*ListErRouteEntriesResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 0
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
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
	// Destination CIDR Block
	//
	// example:
	//
	// 100.64.1.100/32
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// Lingjun HUB Instance ID
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// The ID of the route entry.
	//
	// example:
	//
	// er-rte-maysfadg
	ErRouteEntryId *string `json:"ErRouteEntryId,omitempty" xml:"ErRouteEntryId,omitempty"`
	// The time when the cluster was updated.
	//
	// example:
	//
	// 1640930901000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Next Hop Instance
	//
	// example:
	//
	// vcc-cn-209300qha01
	NextHopId *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// Next Hop Instance Type
	//
	// example:
	//
	// VCC
	NextHopType *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-aekzb3n5lgk2ieq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the tenant to which the resource belongs.
	//
	// example:
	//
	// 1111156667137893
	ResourceTenantId *string `json:"ResourceTenantId,omitempty" xml:"ResourceTenantId,omitempty"`
	// Route type
	//
	// example:
	//
	// VCC
	RouteType *string `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
	// The task status. Valid values:
	//
	// 	- Synchronizing
	//
	// 	- Available
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 1111156667137893
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
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

func (s *ListErRouteEntriesResponseBodyContentData) SetResourceGroupId(v string) *ListErRouteEntriesResponseBodyContentData {
	s.ResourceGroupId = &v
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListErRouteEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// Destination CIDR Block
	//
	// example:
	//
	// 0.0.0.0/0
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// Specifies whether to enable paged query.
	//
	// example:
	//
	// false
	EnablePage *bool `json:"EnablePage,omitempty" xml:"EnablePage,omitempty"`
	// Elastic Router ID
	//
	// This parameter is required.
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// routing policy ID
	//
	// example:
	//
	// er-rmap-uwglhzom
	ErRouteMapId *string `json:"ErRouteMapId,omitempty" xml:"ErRouteMapId,omitempty"`
	// Policy number (default for automatic creation is 3000; The value range of the policy number manually created by the user is 1001-2000)
	//
	// example:
	//
	// 1001
	ErRouteMapNum *int32 `json:"ErRouteMapNum,omitempty" xml:"ErRouteMapNum,omitempty"`
	// The page number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Receive Instance ID
	//
	// example:
	//
	// vpd-x2lohgpv
	ReceptionInstanceId *string `json:"ReceptionInstanceId,omitempty" xml:"ReceptionInstanceId,omitempty"`
	// Receive Instance Name
	//
	// example:
	//
	// vpd2
	ReceptionInstanceName *string `json:"ReceptionInstanceName,omitempty" xml:"ReceptionInstanceName,omitempty"`
	// The type of the received instance. Optional values:
	//
	// 	- **VPD**: Lingjun network segment.
	//
	// 	- **VCC**: Lingjun Connection.
	//
	// example:
	//
	// VPD
	ReceptionInstanceType *string `json:"ReceptionInstanceType,omitempty" xml:"ReceptionInstanceType,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-acfmzaq3ypaqkdy
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Policy behavior; optional values:
	//
	// 	- **permit**: Allow
	//
	// 	- **deny**: Rejected
	//
	// example:
	//
	// deny
	RouteMapAction *string `json:"RouteMapAction,omitempty" xml:"RouteMapAction,omitempty"`
	// Release Instance ID
	//
	// example:
	//
	// vpd-xsdlg2xb
	TransmissionInstanceId *string `json:"TransmissionInstanceId,omitempty" xml:"TransmissionInstanceId,omitempty"`
	// Release Instance Name
	//
	// example:
	//
	// vpd1
	TransmissionInstanceName *string `json:"TransmissionInstanceName,omitempty" xml:"TransmissionInstanceName,omitempty"`
	// The type of the published instance. Optional values:
	//
	// 	- **VPD**: Lingjun network segment.
	//
	// 	- **VCC**: Lingjun Connection.
	//
	// example:
	//
	// VPD
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

func (s *ListErRouteMapsRequest) SetResourceGroupId(v string) *ListErRouteMapsRequest {
	s.ResourceGroupId = &v
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Content *ListErRouteMapsResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0901F411-28FA-5B9C-BAEE-7776463FF0DC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListErRouteMapsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListErRouteMapsResponseBody) GoString() string {
	return s.String()
}

func (s *ListErRouteMapsResponseBody) SetAccessDeniedDetail(v string) *ListErRouteMapsResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// routing policy information list
	Data []*ListErRouteMapsResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 0
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
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
	// Policy behavior; optional values:
	//
	// 	- **permit**: Allow
	//
	// 	- **deny**: Prohibited
	//
	// example:
	//
	// permit
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The time when the data address was created.
	//
	// example:
	//
	// 1601176751000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Policy description
	//
	// example:
	//
	// No description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Destination CIDR Block
	//
	// example:
	//
	// 0.0.0.0/0
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// Lingjun HUB ID
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// routing policy ID
	//
	// example:
	//
	// er-rmap-uwglhzom
	ErRouteMapId *string `json:"ErRouteMapId,omitempty" xml:"ErRouteMapId,omitempty"`
	// The time when the agent was last modified.
	//
	// example:
	//
	// 1601176751000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Receive Instance ID
	//
	// example:
	//
	// vpd-9rgxqazc
	ReceptionInstanceId *string `json:"ReceptionInstanceId,omitempty" xml:"ReceptionInstanceId,omitempty"`
	// Receive Instance Name
	//
	// example:
	//
	// vpd-reception
	ReceptionInstanceName *string `json:"ReceptionInstanceName,omitempty" xml:"ReceptionInstanceName,omitempty"`
	// The tenant to which the receiving instance belongs
	//
	// example:
	//
	// 1620939556166277
	ReceptionInstanceOwner *string `json:"ReceptionInstanceOwner,omitempty" xml:"ReceptionInstanceOwner,omitempty"`
	// The type of the received instance. Possible values:
	//
	// 	- **VPD**: Lingjun network segment.
	//
	// 	- **VCC**: Lingjun Connection.
	//
	// example:
	//
	// VPD
	ReceptionInstanceType *string `json:"ReceptionInstanceType,omitempty" xml:"ReceptionInstanceType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-aek2l4sq6l7unhi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the policy.
	//
	// A smaller sequence number indicates a lower priority. When a route is matched, a policy with a higher priority is preferentially matched.
	//
	// **Valid values: 1001 to 2000**
	//
	// example:
	//
	// 1001
	RouteMapNum *int32 `json:"RouteMapNum,omitempty" xml:"RouteMapNum,omitempty"`
	// Status The status of the instance. Valid values:
	//
	// 	- **Available**
	//
	// 	- **Not Available**: Unavailable
	//
	// 	- **Executing**: Executing
	//
	// 	- **Deleting**: The node is being deleted.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 1655449505171
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// Release Instance ID
	//
	// example:
	//
	// vpd-8rgvqazb
	TransmissionInstanceId *string `json:"TransmissionInstanceId,omitempty" xml:"TransmissionInstanceId,omitempty"`
	// Release Instance Name
	//
	// example:
	//
	// vpd-transmit
	TransmissionInstanceName *string `json:"TransmissionInstanceName,omitempty" xml:"TransmissionInstanceName,omitempty"`
	// The tenant to which the published instance belongs
	//
	// example:
	//
	// 1620939556166277
	TransmissionInstanceOwner *string `json:"TransmissionInstanceOwner,omitempty" xml:"TransmissionInstanceOwner,omitempty"`
	// The type of the published instance. Possible values:
	//
	// 	- **VPD**: Lingjun network segment.
	//
	// 	- **VCC**: Lingjun Connection.
	//
	// example:
	//
	// VPD
	TransmissionInstanceType *string `json:"TransmissionInstanceType,omitempty" xml:"TransmissionInstanceType,omitempty"`
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

func (s *ListErRouteMapsResponseBodyContentData) SetResourceGroupId(v string) *ListErRouteMapsResponseBodyContentData {
	s.ResourceGroupId = &v
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListErRouteMapsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// Specifies whether to enable paged query. Valid values:
	//
	// 	- true: enables paged query.
	//
	// 	- false: Paged query is disabled.
	//
	// example:
	//
	// false
	EnablePage *bool `json:"EnablePage,omitempty" xml:"EnablePage,omitempty"`
	// The ID of the Lingjun HUB instance.
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// Lingjun HUB name.
	//
	// example:
	//
	// er-heyuan-main
	ErName *string `json:"ErName,omitempty" xml:"ErName,omitempty"`
	// The ID of the network instance.
	//
	// example:
	//
	// vcc-cn-209300qha01
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the attached network instance. Valid values:
	//
	// 	- **VPD**
	//
	// 	- **VCC**
	//
	// example:
	//
	// VCC
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The primary zone.
	//
	// example:
	//
	// cn-wulanchabu-b
	MasterZoneId *string `json:"MasterZoneId,omitempty" xml:"MasterZoneId,omitempty"`
	// The page number to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-acfmwfm33rlt6zi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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

func (s *ListErsRequest) SetResourceGroupId(v string) *ListErsRequest {
	s.ResourceGroupId = &v
	return s
}

type ListErsResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Content *ListErsResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is displayed.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AC8C713A-A9F4-5984-A5E1-76496DF35153
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListErsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListErsResponseBody) GoString() string {
	return s.String()
}

func (s *ListErsResponseBody) SetAccessDeniedDetail(v string) *ListErsResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// lingjun hub information list.
	Data []*ListErsResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
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
	// The number of connections to the Lingjun HUB network instance.
	//
	// example:
	//
	// 2
	Connections *int64 `json:"Connections,omitempty" xml:"Connections,omitempty"`
	// The time when the activation code was created.
	//
	// example:
	//
	// 1640930671000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the synchronization task.
	//
	// example:
	//
	// No description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the Lingjun HUB instance.
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// The name of the Lingjun HUB instance.
	//
	// example:
	//
	// er-wulanchabu-main
	ErName *string `json:"ErName,omitempty" xml:"ErName,omitempty"`
	// The time when the O\\&M task was modified.
	//
	// example:
	//
	// 1640930671000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The primary zone.
	//
	// example:
	//
	// cn-wulanchabu-b
	MasterZoneId *string `json:"MasterZoneId,omitempty" xml:"MasterZoneId,omitempty"`
	// The returned message.
	//
	// example:
	//
	// test
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-acfmv2m2w43japa
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Number of Lingjun HUB routing policy.
	//
	// example:
	//
	// 2
	RouteMaps *int64 `json:"RouteMaps,omitempty" xml:"RouteMaps,omitempty"`
	// The task status.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 1655449505171
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
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

func (s *ListErsResponseBodyContentData) SetResourceGroupId(v string) *ListErsResponseBodyContentData {
	s.ResourceGroupId = &v
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
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListErsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ListInstancesByNcdRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lni-1234****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The parameter that specifies the instance type.
	//
	// Valid value:
	//
	// 	- node: Lingjun node.
	//
	// 	- lni: lingjun network interface controller.
	//
	// This parameter is required.
	//
	// example:
	//
	// lni
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// Maximum network communication distance
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	MaxNcd *int32 `json:"MaxNcd,omitempty" xml:"MaxNcd,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListInstancesByNcdRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesByNcdRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesByNcdRequest) SetInstanceId(v string) *ListInstancesByNcdRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesByNcdRequest) SetInstanceType(v string) *ListInstancesByNcdRequest {
	s.InstanceType = &v
	return s
}

func (s *ListInstancesByNcdRequest) SetMaxNcd(v int32) *ListInstancesByNcdRequest {
	s.MaxNcd = &v
	return s
}

func (s *ListInstancesByNcdRequest) SetRegionId(v string) *ListInstancesByNcdRequest {
	s.RegionId = &v
	return s
}

type ListInstancesByNcdResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Content *ListInstancesByNcdResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// You don\\"t have the permission of this operation, action=eflo:ListInstancesByNcd, arn=acs:eflo:cn-heyuan:1263399219805497:networkinterface/*, resourceGroup=null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// AC8C713A-A9F4-5984-A5E1-76496DF35153
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListInstancesByNcdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesByNcdResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesByNcdResponseBody) SetAccessDeniedDetail(v string) *ListInstancesByNcdResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListInstancesByNcdResponseBody) SetCode(v int32) *ListInstancesByNcdResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstancesByNcdResponseBody) SetContent(v *ListInstancesByNcdResponseBodyContent) *ListInstancesByNcdResponseBody {
	s.Content = v
	return s
}

func (s *ListInstancesByNcdResponseBody) SetMessage(v string) *ListInstancesByNcdResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstancesByNcdResponseBody) SetRequestId(v string) *ListInstancesByNcdResponseBody {
	s.RequestId = &v
	return s
}

type ListInstancesByNcdResponseBodyContent struct {
	// A collection of instances whose network communication distance from the source instance ID does not exceed maxNcd
	InstanceInfos []*ListInstancesByNcdResponseBodyContentInstanceInfos `json:"InstanceInfos,omitempty" xml:"InstanceInfos,omitempty" type:"Repeated"`
	// Instance Type
	//
	// Valid value:
	//
	// 	- node: Lingjun node.
	//
	// 	- lni: lingjun network interface controller.
	//
	// example:
	//
	// lni
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// Maximum communication distance between nodes
	//
	// example:
	//
	// 3
	MaxNcd *int32 `json:"MaxNcd,omitempty" xml:"MaxNcd,omitempty"`
	// The ID of the source instance.
	//
	// example:
	//
	// lni-1234****
	SourceInstanceId *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
}

func (s ListInstancesByNcdResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesByNcdResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListInstancesByNcdResponseBodyContent) SetInstanceInfos(v []*ListInstancesByNcdResponseBodyContentInstanceInfos) *ListInstancesByNcdResponseBodyContent {
	s.InstanceInfos = v
	return s
}

func (s *ListInstancesByNcdResponseBodyContent) SetInstanceType(v string) *ListInstancesByNcdResponseBodyContent {
	s.InstanceType = &v
	return s
}

func (s *ListInstancesByNcdResponseBodyContent) SetMaxNcd(v int32) *ListInstancesByNcdResponseBodyContent {
	s.MaxNcd = &v
	return s
}

func (s *ListInstancesByNcdResponseBodyContent) SetSourceInstanceId(v string) *ListInstancesByNcdResponseBodyContent {
	s.SourceInstanceId = &v
	return s
}

type ListInstancesByNcdResponseBodyContentInstanceInfos struct {
	// The instance ID.
	//
	// example:
	//
	// lni-1235****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// network communication distance
	//
	// example:
	//
	// 2
	Ncd *int32 `json:"Ncd,omitempty" xml:"Ncd,omitempty"`
}

func (s ListInstancesByNcdResponseBodyContentInstanceInfos) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesByNcdResponseBodyContentInstanceInfos) GoString() string {
	return s.String()
}

func (s *ListInstancesByNcdResponseBodyContentInstanceInfos) SetInstanceId(v string) *ListInstancesByNcdResponseBodyContentInstanceInfos {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesByNcdResponseBodyContentInstanceInfos) SetNcd(v int32) *ListInstancesByNcdResponseBodyContentInstanceInfos {
	s.Ncd = &v
	return s
}

type ListInstancesByNcdResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstancesByNcdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstancesByNcdResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesByNcdResponse) GoString() string {
	return s.String()
}

func (s *ListInstancesByNcdResponse) SetHeaders(v map[string]*string) *ListInstancesByNcdResponse {
	s.Headers = v
	return s
}

func (s *ListInstancesByNcdResponse) SetStatusCode(v int32) *ListInstancesByNcdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstancesByNcdResponse) SetBody(v *ListInstancesByNcdResponseBody) *ListInstancesByNcdResponse {
	s.Body = v
	return s
}

type ListLeniPrivateIpAddressesRequest struct {
	// Lingjun Elastic Network Interface ID.
	//
	// example:
	//
	// leni-1234****
	ElasticNetworkInterfaceId *string `json:"ElasticNetworkInterfaceId,omitempty" xml:"ElasticNetworkInterfaceId,omitempty"`
	// Lingjun Elastic Network Interface secondary private IP unique identifier.
	//
	// example:
	//
	// sip-8ylg****
	IpName *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
	// The page number returned.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Lingjun Elastic Network Interface secondary private IP.
	//
	// example:
	//
	// 10.0.****
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the image build command risk.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListLeniPrivateIpAddressesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLeniPrivateIpAddressesRequest) GoString() string {
	return s.String()
}

func (s *ListLeniPrivateIpAddressesRequest) SetElasticNetworkInterfaceId(v string) *ListLeniPrivateIpAddressesRequest {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *ListLeniPrivateIpAddressesRequest) SetIpName(v string) *ListLeniPrivateIpAddressesRequest {
	s.IpName = &v
	return s
}

func (s *ListLeniPrivateIpAddressesRequest) SetPageNumber(v int32) *ListLeniPrivateIpAddressesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListLeniPrivateIpAddressesRequest) SetPageSize(v int32) *ListLeniPrivateIpAddressesRequest {
	s.PageSize = &v
	return s
}

func (s *ListLeniPrivateIpAddressesRequest) SetPrivateIpAddress(v string) *ListLeniPrivateIpAddressesRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *ListLeniPrivateIpAddressesRequest) SetRegionId(v string) *ListLeniPrivateIpAddressesRequest {
	s.RegionId = &v
	return s
}

func (s *ListLeniPrivateIpAddressesRequest) SetStatus(v string) *ListLeniPrivateIpAddressesRequest {
	s.Status = &v
	return s
}

type ListLeniPrivateIpAddressesResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Content *ListLeniPrivateIpAddressesResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// AC8C713A-A9F4-5984-A5E1-76496DF35153
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListLeniPrivateIpAddressesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLeniPrivateIpAddressesResponseBody) GoString() string {
	return s.String()
}

func (s *ListLeniPrivateIpAddressesResponseBody) SetAccessDeniedDetail(v string) *ListLeniPrivateIpAddressesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListLeniPrivateIpAddressesResponseBody) SetCode(v int32) *ListLeniPrivateIpAddressesResponseBody {
	s.Code = &v
	return s
}

func (s *ListLeniPrivateIpAddressesResponseBody) SetContent(v *ListLeniPrivateIpAddressesResponseBodyContent) *ListLeniPrivateIpAddressesResponseBody {
	s.Content = v
	return s
}

func (s *ListLeniPrivateIpAddressesResponseBody) SetMessage(v string) *ListLeniPrivateIpAddressesResponseBody {
	s.Message = &v
	return s
}

func (s *ListLeniPrivateIpAddressesResponseBody) SetRequestId(v string) *ListLeniPrivateIpAddressesResponseBody {
	s.RequestId = &v
	return s
}

type ListLeniPrivateIpAddressesResponseBodyContent struct {
	// The response parameters.
	Data []*ListLeniPrivateIpAddressesResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListLeniPrivateIpAddressesResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s ListLeniPrivateIpAddressesResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListLeniPrivateIpAddressesResponseBodyContent) SetData(v []*ListLeniPrivateIpAddressesResponseBodyContentData) *ListLeniPrivateIpAddressesResponseBodyContent {
	s.Data = v
	return s
}

func (s *ListLeniPrivateIpAddressesResponseBodyContent) SetTotal(v int64) *ListLeniPrivateIpAddressesResponseBodyContent {
	s.Total = &v
	return s
}

type ListLeniPrivateIpAddressesResponseBodyContentData struct {
	// The description.
	//
	// example:
	//
	// test_vpn1_pbr_route_54
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Lingjun Elastic Network Interface ID.
	//
	// example:
	//
	// leni-1234****
	ElasticNetworkInterfaceId *string `json:"ElasticNetworkInterfaceId,omitempty" xml:"ElasticNetworkInterfaceId,omitempty"`
	// The time when the activation code was created.
	//
	// example:
	//
	// 1675929918000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the certificate was updated.
	//
	// example:
	//
	// 1675929918000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Lingjun Elastic Network Interface secondary private IP unique identifier.
	//
	// example:
	//
	// sip-8ylg****
	IpName *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Lingjun Elastic Network Interface secondary private IP address.
	//
	// example:
	//
	// 10.0.****
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The task status.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListLeniPrivateIpAddressesResponseBodyContentData) String() string {
	return tea.Prettify(s)
}

func (s ListLeniPrivateIpAddressesResponseBodyContentData) GoString() string {
	return s.String()
}

func (s *ListLeniPrivateIpAddressesResponseBodyContentData) SetDescription(v string) *ListLeniPrivateIpAddressesResponseBodyContentData {
	s.Description = &v
	return s
}

func (s *ListLeniPrivateIpAddressesResponseBodyContentData) SetElasticNetworkInterfaceId(v string) *ListLeniPrivateIpAddressesResponseBodyContentData {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *ListLeniPrivateIpAddressesResponseBodyContentData) SetGmtCreate(v string) *ListLeniPrivateIpAddressesResponseBodyContentData {
	s.GmtCreate = &v
	return s
}

func (s *ListLeniPrivateIpAddressesResponseBodyContentData) SetGmtModified(v string) *ListLeniPrivateIpAddressesResponseBodyContentData {
	s.GmtModified = &v
	return s
}

func (s *ListLeniPrivateIpAddressesResponseBodyContentData) SetIpName(v string) *ListLeniPrivateIpAddressesResponseBodyContentData {
	s.IpName = &v
	return s
}

func (s *ListLeniPrivateIpAddressesResponseBodyContentData) SetMessage(v string) *ListLeniPrivateIpAddressesResponseBodyContentData {
	s.Message = &v
	return s
}

func (s *ListLeniPrivateIpAddressesResponseBodyContentData) SetPrivateIpAddress(v string) *ListLeniPrivateIpAddressesResponseBodyContentData {
	s.PrivateIpAddress = &v
	return s
}

func (s *ListLeniPrivateIpAddressesResponseBodyContentData) SetRegionId(v string) *ListLeniPrivateIpAddressesResponseBodyContentData {
	s.RegionId = &v
	return s
}

func (s *ListLeniPrivateIpAddressesResponseBodyContentData) SetStatus(v string) *ListLeniPrivateIpAddressesResponseBodyContentData {
	s.Status = &v
	return s
}

type ListLeniPrivateIpAddressesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLeniPrivateIpAddressesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLeniPrivateIpAddressesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLeniPrivateIpAddressesResponse) GoString() string {
	return s.String()
}

func (s *ListLeniPrivateIpAddressesResponse) SetHeaders(v map[string]*string) *ListLeniPrivateIpAddressesResponse {
	s.Headers = v
	return s
}

func (s *ListLeniPrivateIpAddressesResponse) SetStatusCode(v int32) *ListLeniPrivateIpAddressesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLeniPrivateIpAddressesResponse) SetBody(v *ListLeniPrivateIpAddressesResponseBody) *ListLeniPrivateIpAddressesResponse {
	s.Body = v
	return s
}

type ListLniPrivateIpAddressRequest struct {
	// The description of the variable.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Whether pagination is required
	//
	// example:
	//
	// false
	EnablePage *bool `json:"EnablePage,omitempty" xml:"EnablePage,omitempty"`
	// network interface controller IP address
	//
	// example:
	//
	// 10.0.98.10
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// IP unique identifier
	//
	// example:
	//
	// sip-tynhdh2s
	IpName *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
	// Lingjun network interface controller ID
	//
	// example:
	//
	// lni-2ze4uww7n6hsfzrwq77y
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Obtain the index number of the current mouse click for an animation
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListLniPrivateIpAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLniPrivateIpAddressRequest) GoString() string {
	return s.String()
}

func (s *ListLniPrivateIpAddressRequest) SetDescription(v string) *ListLniPrivateIpAddressRequest {
	s.Description = &v
	return s
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Content *ListLniPrivateIpAddressResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// You don\\"t have the permission to do this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A88DFED5-24B7-5A3E-87DE-380BF06F3C90
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListLniPrivateIpAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLniPrivateIpAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ListLniPrivateIpAddressResponseBody) SetAccessDeniedDetail(v string) *ListLniPrivateIpAddressResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// The returned result.
	Data []*ListLniPrivateIpAddressResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 0
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
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
	// The instance description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the data address was created.
	//
	// example:
	//
	// 1651734291000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// MAC address of the secondary private network
	//
	// example:
	//
	// 00-ff-84-15-ba-67
	IpAddressMac *string `json:"IpAddressMac,omitempty" xml:"IpAddressMac,omitempty"`
	// IP unique identifier
	//
	// example:
	//
	// sip-1hq1ql7vz
	IpName *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Lingjun network interface controller ID
	//
	// example:
	//
	// lni-bp11hq1ql7vza3k4xz7q
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// Secondary private IP address of Lingjun network interface controller
	//
	// example:
	//
	// 10.42.5.92
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the intervention entry. Valid value:
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListLniPrivateIpAddressResponseBodyContentData) String() string {
	return tea.Prettify(s)
}

func (s ListLniPrivateIpAddressResponseBodyContentData) GoString() string {
	return s.String()
}

func (s *ListLniPrivateIpAddressResponseBodyContentData) SetDescription(v string) *ListLniPrivateIpAddressResponseBodyContentData {
	s.Description = &v
	return s
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLniPrivateIpAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// Specifies whether pagination is required.
	//
	// example:
	//
	// false
	EnablePage *bool `json:"EnablePage,omitempty" xml:"EnablePage,omitempty"`
	// network interface controller the IP address.
	//
	// example:
	//
	// 203.107.46.227
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// Lingjun network interface controller ID.
	//
	// example:
	//
	// lni-bp18exxqa2rvfn45e5pz
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The ID of the machine to which the instance belongs.
	//
	// example:
	//
	// r-2ze121o4uhr4np3r5t-db-5
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The current number of pages.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the instance to which the Lingjun subnet belongs.
	//
	// example:
	//
	// subnet-anhtskts
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	// The ID of the VPD.
	//
	// example:
	//
	// vpd-iv2zm1qf
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Content *ListNetworkInterfacesResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 039C3C3A-3C37-5672-80D5-D8CD48C676D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListNetworkInterfacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNetworkInterfacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNetworkInterfacesResponseBody) SetAccessDeniedDetail(v string) *ListNetworkInterfacesResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// The response parameters.
	Data []*ListNetworkInterfacesResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The total number of entries that are returned.
	//
	// example:
	//
	// 0
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
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
	// The time when the activation code was created.
	//
	// example:
	//
	// 1669734207000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The port number of the AD server.
	Ethernet []*string `json:"Ethernet,omitempty" xml:"Ethernet,omitempty" type:"Repeated"`
	// The gateway.
	//
	// example:
	//
	// 10.0.0.253
	Gateway *string `json:"Gateway,omitempty" xml:"Gateway,omitempty"`
	// The IP address of the instance.
	//
	// example:
	//
	// 10.0.0.13
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The NC type.
	//
	// Valid value:
	//
	// 	- CUSTOM_LNI_INTEGRATION: two-network one-in-one architecture Lingjun hosting network interface controller.
	//
	// 	- CPU: CPU machine.
	//
	// 	- ELASTIC_6.2: Machine
	//
	// 	- GPU: GPU machine.
	//
	// 	- DEFAULT: the old CPU machine.
	//
	// 	- CUSTOM_LNI: two network separation architecture Lingjun hosting network interface controller.
	//
	// example:
	//
	// GPU
	NcType *string `json:"NcType,omitempty" xml:"NcType,omitempty"`
	// Lingjun network interface controller ID.
	//
	// example:
	//
	// lni-2ze50voovmtswn328ogm
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The port name.
	//
	// example:
	//
	// bond0
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitempty" xml:"NetworkInterfaceName,omitempty"`
	// The ID of the machine to which the instance belongs.
	//
	// example:
	//
	// 2d53f5c204e7476dae69177e7fa6f19c
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// Secondary Private IP\\&MAC Address Collection
	PrivateIpAddressMacGroup []*ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup `json:"PrivateIpAddressMacGroup,omitempty" xml:"PrivateIpAddressMacGroup,omitempty" type:"Repeated"`
	// network interface controller private secondary IP quota.
	//
	// example:
	//
	// 6
	Quota *int32 `json:"Quota,omitempty" xml:"Quota,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The address of the service network interface controller.
	//
	// example:
	//
	// 00-ff-84-15-ba-67
	ServiceMac *string `json:"ServiceMac,omitempty" xml:"ServiceMac,omitempty"`
	// The task status.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Lingjun subnet (Subnet) basic information.
	SubnetBaseInfo *ListNetworkInterfacesResponseBodyContentDataSubnetBaseInfo `json:"SubnetBaseInfo,omitempty" xml:"SubnetBaseInfo,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// example:
	//
	// 1620939556166279
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// Lingjun network segment (VPD) basic information.
	VpdBaseInfo *ListNetworkInterfacesResponseBodyContentDataVpdBaseInfo `json:"VpdBaseInfo,omitempty" xml:"VpdBaseInfo,omitempty" type:"Struct"`
	// The zone ID.
	//
	// example:
	//
	// cn-wulanchabu-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
	// The instance description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Secondary private MAC address.
	//
	// example:
	//
	// 00:25:9d:00:20:20
	IpAddressMac *string `json:"IpAddressMac,omitempty" xml:"IpAddressMac,omitempty"`
	// The unique IP identifier.
	//
	// example:
	//
	// sip-1asjd3xg
	IpName *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The secondary private IP address.
	//
	// example:
	//
	// 10.0.0.14
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The status of the cache reserve instance.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup) String() string {
	return tea.Prettify(s)
}

func (s ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup) GoString() string {
	return s.String()
}

func (s *ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup) SetDescription(v string) *ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup {
	s.Description = &v
	return s
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
	// The network segment of the Subnet.
	//
	// 	- The network segment of the subnet must be a proper subset of the network segment of Lingjun to which it belongs, and the mask must be between 16 bits and 29 bits, which can provide 8 to 65536 addresses. For example, the CIDR block of the Lingjun CIDR block is 192.168.0.0/16, and the CIDR blocks of the subnets under the Lingjun CIDR block are 192.168.0.0/17 to 192.168.0.0/29.
	//
	// 	- The first and last three IP addresses of each subnet segment are reserved by the system. For example, the CIDR blocks of the subnet are 192.168.1.0/24,192.168.1.0, 192.168.1.253, 192.168.1.254, and 192.168.1.255.
	//
	// For more information about CIDR blocks, see the [What is CIDR?](https://www.alibabacloud.com/help/doc-detail/40637.htm#title-gu4-uzk-12r) section in the "Network FAQ" topic.
	//
	// This parameter is left empty by default.
	//
	// example:
	//
	// 10.0.0.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The time when the activation code was created.
	//
	// example:
	//
	// 1623656472000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the Subnet instance.
	//
	// example:
	//
	// subnet-yjnqn5ef
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	// The name of the Subnet instance.
	//
	// example:
	//
	// subnet-1
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
	// The network segment of Lingjun network segment (VPD).
	//
	// 	- We recommend that you use an RFC private endpoint as the Lingjun CIDR block, such as 10.0.0.0/8,172.16.0.0/12,192.168.0.0/16. In scenarios where the Doringjun CIDR block is connected to each other or where the Lingjun CIDR block is connected to a VPC, make sure that the addresses do not conflict with each other.
	//
	// 	- You can also use a custom CIDR block other than 100.64.0.0/10, 224.0.0.0/4, 127.0.0.0/8, or 169.254.0.0/16 and their subnets as the primary IPv4 CIDR block of the VPD. This parameter is left empty by default.
	//
	// example:
	//
	// 10.0.0.0/16
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The time when the activation code was created.
	//
	// example:
	//
	// 1668158213000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the VPD instance.
	//
	// example:
	//
	// vpd-d3isyds4
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// The name of the VPD instance.
	//
	// example:
	//
	// vpd-1
	VpdName *string `json:"VpdName,omitempty" xml:"VpdName,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNetworkInterfacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ListNodeInfosForPodRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// cluster-***
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the node for this operation.
	//
	// example:
	//
	// node-be70****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-wulanchabu-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListNodeInfosForPodRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodeInfosForPodRequest) GoString() string {
	return s.String()
}

func (s *ListNodeInfosForPodRequest) SetClusterId(v string) *ListNodeInfosForPodRequest {
	s.ClusterId = &v
	return s
}

func (s *ListNodeInfosForPodRequest) SetNodeId(v string) *ListNodeInfosForPodRequest {
	s.NodeId = &v
	return s
}

func (s *ListNodeInfosForPodRequest) SetRegionId(v string) *ListNodeInfosForPodRequest {
	s.RegionId = &v
	return s
}

func (s *ListNodeInfosForPodRequest) SetZoneId(v string) *ListNodeInfosForPodRequest {
	s.ZoneId = &v
	return s
}

type ListNodeInfosForPodResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Response body
	Content []*ListNodeInfosForPodResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// You don\\"t have the permission of this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// 0901F411-28FA-5B9C-BAEE-7776463FF0DC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListNodeInfosForPodResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodeInfosForPodResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodeInfosForPodResponseBody) SetAccessDeniedDetail(v string) *ListNodeInfosForPodResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListNodeInfosForPodResponseBody) SetCode(v int32) *ListNodeInfosForPodResponseBody {
	s.Code = &v
	return s
}

func (s *ListNodeInfosForPodResponseBody) SetContent(v []*ListNodeInfosForPodResponseBodyContent) *ListNodeInfosForPodResponseBody {
	s.Content = v
	return s
}

func (s *ListNodeInfosForPodResponseBody) SetMessage(v string) *ListNodeInfosForPodResponseBody {
	s.Message = &v
	return s
}

func (s *ListNodeInfosForPodResponseBody) SetRequestId(v string) *ListNodeInfosForPodResponseBody {
	s.RequestId = &v
	return s
}

type ListNodeInfosForPodResponseBodyContent struct {
	// The cluster ID.
	//
	// example:
	//
	// cluster-****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Lingjun Gaomi network interface controller quota
	//
	// example:
	//
	// 10
	HdeniQuota *int32 `json:"HdeniQuota,omitempty" xml:"HdeniQuota,omitempty"`
	// Lingjun Elastic Network Interface quota, excluding system type
	//
	// example:
	//
	// 10
	LeniQuota *int32 `json:"LeniQuota,omitempty" xml:"LeniQuota,omitempty"`
	// Lingjun Elastic Network Interface Secondary Private IP Quota
	//
	// example:
	//
	// 10
	LeniSipQuota *int32 `json:"LeniSipQuota,omitempty" xml:"LeniSipQuota,omitempty"`
	// Lingjun network interface controller Secondary Private IP Quota
	//
	// example:
	//
	// 10
	LniSipQuota *int32 `json:"LniSipQuota,omitempty" xml:"LniSipQuota,omitempty"`
	// The ID of the node for this operation.
	//
	// example:
	//
	// node-be70****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// List of VSwitches to which IP addresses can be applied for this node
	VSwitches []*string `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
	// The ID of the Virtual Private Cloud to which the current node belongs.
	//
	// example:
	//
	// vpc-j6ctp4n75306****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-wulanchabu-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListNodeInfosForPodResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s ListNodeInfosForPodResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListNodeInfosForPodResponseBodyContent) SetClusterId(v string) *ListNodeInfosForPodResponseBodyContent {
	s.ClusterId = &v
	return s
}

func (s *ListNodeInfosForPodResponseBodyContent) SetHdeniQuota(v int32) *ListNodeInfosForPodResponseBodyContent {
	s.HdeniQuota = &v
	return s
}

func (s *ListNodeInfosForPodResponseBodyContent) SetLeniQuota(v int32) *ListNodeInfosForPodResponseBodyContent {
	s.LeniQuota = &v
	return s
}

func (s *ListNodeInfosForPodResponseBodyContent) SetLeniSipQuota(v int32) *ListNodeInfosForPodResponseBodyContent {
	s.LeniSipQuota = &v
	return s
}

func (s *ListNodeInfosForPodResponseBodyContent) SetLniSipQuota(v int32) *ListNodeInfosForPodResponseBodyContent {
	s.LniSipQuota = &v
	return s
}

func (s *ListNodeInfosForPodResponseBodyContent) SetNodeId(v string) *ListNodeInfosForPodResponseBodyContent {
	s.NodeId = &v
	return s
}

func (s *ListNodeInfosForPodResponseBodyContent) SetRegionId(v string) *ListNodeInfosForPodResponseBodyContent {
	s.RegionId = &v
	return s
}

func (s *ListNodeInfosForPodResponseBodyContent) SetVSwitches(v []*string) *ListNodeInfosForPodResponseBodyContent {
	s.VSwitches = v
	return s
}

func (s *ListNodeInfosForPodResponseBodyContent) SetVpcId(v string) *ListNodeInfosForPodResponseBodyContent {
	s.VpcId = &v
	return s
}

func (s *ListNodeInfosForPodResponseBodyContent) SetZoneId(v string) *ListNodeInfosForPodResponseBodyContent {
	s.ZoneId = &v
	return s
}

type ListNodeInfosForPodResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNodeInfosForPodResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNodeInfosForPodResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodeInfosForPodResponse) GoString() string {
	return s.String()
}

func (s *ListNodeInfosForPodResponse) SetHeaders(v map[string]*string) *ListNodeInfosForPodResponse {
	s.Headers = v
	return s
}

func (s *ListNodeInfosForPodResponse) SetStatusCode(v int32) *ListNodeInfosForPodResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodeInfosForPodResponse) SetBody(v *ListNodeInfosForPodResponseBody) *ListNodeInfosForPodResponse {
	s.Body = v
	return s
}

type ListSubnetsRequest struct {
	// Specifies whether to query by page. Optional values:
	//
	// 	- **true**: Enable pagination query
	//
	// 	- **false**: Pagination query is disabled
	//
	// example:
	//
	// false
	EnablePage *bool `json:"EnablePage,omitempty" xml:"EnablePage,omitempty"`
	// The number of the page to return. The value must be greater than 0. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// For more information about resource groups, see [Resource groups](https://help.aliyun.com/document_detail/94475.htm?spm=a2c4g.11186623.0.0.29e15d7akXhpuu).
	//
	// example:
	//
	// rg-aeky5f3qx6ceapq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the CLB instance. Valid values:
	//
	// 	- **Available**: Normal
	//
	// 	- **Not Available**: Unavailable
	//
	// 	- **Executing**: Executing
	//
	// 	- **Deleting**: The node is being deleted.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Lingjun subnet instance ID
	//
	// example:
	//
	// subnet-anhtskts
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	// Lingjun subnet instance name
	//
	// example:
	//
	// subnet-1
	SubnetName *string `json:"SubnetName,omitempty" xml:"SubnetName,omitempty"`
	// The tag information.
	//
	// You can specify up to 20 tags.
	Tag []*ListSubnetsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Lingjun Subnet Usage Type; optional; optional. Valid values:
	//
	// 	- **If you do not set this field for a common type**
	//
	// 	- **OOB*	- :OOB type
	//
	// 	- **LB**: LB type
	//
	// example:
	//
	// Null
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the Lingjun CIDR block.
	//
	// example:
	//
	// vpd-fuliephf
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// The zone ID of the disk.
	//
	// example:
	//
	// cn-wulanchabu-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
	// The tag key of the VPN attachment.
	//
	// You cannot specify an empty string as a tag key. It can be up to 64 characters in length and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// You can specify at most 20 tag keys in each call.
	//
	// example:
	//
	// rg-subnet
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the VPN connection.
	//
	// The tag value can be empty or a string of up to 128 characters. It cannot start with aliyun or acs:, and cannot contain http:// or https://.
	//
	// Each key-value pair must be unique. You can specify values for at most 20 tag keys in each call.
	//
	// example:
	//
	// subnet-group-1
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Content *ListSubnetsResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// 7F9082CC-3D94-560F-A575-8E8EF6CE2CB8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSubnetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSubnetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubnetsResponseBody) SetAccessDeniedDetail(v string) *ListSubnetsResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// Lingjun subnet information list
	Data []*ListSubnetsResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
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
	// The CIDR block of the Subnet.
	//
	// 	- The network segment of the subnet must be a proper subset of the network segment of Lingjun to which it belongs, and the mask must be between 16 bits and 29 bits, which can provide 8 to 65536 addresses. For example, the CIDR block of the Lingjun CIDR block is 192.168.0.0/16, and the CIDR blocks of the subnets under the Lingjun CIDR block are 192.168.0.0/17 to 192.168.0.0/29.
	//
	// 	- The first and last three IP addresses of each subnet segment are reserved by the system. For example, the CIDR blocks of the subnet are 192.168.1.0/24,192.168.1.0, 192.168.1.253, 192.168.1.254, and 192.168.1.255.
	//
	// example:
	//
	// 172.18.0.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The time when the data address was created.
	//
	// example:
	//
	// 1678273219000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the agent was last modified.
	//
	// example:
	//
	// 1678273219000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Number of NCs
	//
	// example:
	//
	// 16
	NcCount *int32 `json:"NcCount,omitempty" xml:"NcCount,omitempty"`
	// Number of Lingjun network interface controller
	//
	// example:
	//
	// 1
	NetworkInterfaceCount *int32 `json:"NetworkInterfaceCount,omitempty" xml:"NetworkInterfaceCount,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of your Alibaba Cloud resource group.
	//
	// For more information about resource groups, see [Resource groups](https://help.aliyun.com/document_detail/94475.htm?spm=a2c4g.11186623.0.0.29e15d7akXhpuu).
	//
	// example:
	//
	// rg-aek2l4sq6l7unhi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the cache reserve instance. Valid values:
	//
	// 	- **Available**: Normal
	//
	// 	- **Not Available**: Unavailable
	//
	// 	- **Executing**: Executing
	//
	// 	- **Deleting**: The node is being deleted.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Lingjun subnet instance ID
	//
	// example:
	//
	// subnet-c6wci55i
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	// Lingjun subnet instance name
	//
	// example:
	//
	// yzp-rg-test3
	SubnetName *string `json:"SubnetName,omitempty" xml:"SubnetName,omitempty"`
	// The tag information.
	//
	// You can specify up to 20 tags.
	Tags []*ListSubnetsResponseBodyContentDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the tenant.
	//
	// example:
	//
	// 1655449505171
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// Lingjun Subnet Usage Type; optional; optional. Valid values:
	//
	// 	- **If you do not set this field for a common type**
	//
	// 	- **OOB*	- :OOB type
	//
	// 	- **LB**: LB type
	//
	// example:
	//
	// OOB
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// vpd basic information
	VpdBaseInfo *ListSubnetsResponseBodyContentDataVpdBaseInfo `json:"VpdBaseInfo,omitempty" xml:"VpdBaseInfo,omitempty" type:"Struct"`
	// The ID of the Lingjun CIDR block.
	//
	// example:
	//
	// vpd-ze3na0wf
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-wulanchabu-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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

func (s *ListSubnetsResponseBodyContentData) SetNcCount(v int32) *ListSubnetsResponseBodyContentData {
	s.NcCount = &v
	return s
}

func (s *ListSubnetsResponseBodyContentData) SetNetworkInterfaceCount(v int32) *ListSubnetsResponseBodyContentData {
	s.NetworkInterfaceCount = &v
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
	// The tag key.
	//
	// You cannot specify an empty string as a tag key. It can be up to 64 characters in length and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// You can specify at most 20 tag keys in each call.
	//
	// example:
	//
	// tag-subnet
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag that is added to the resource.
	//
	// The tag value can be empty or a string of up to 128 characters. It cannot start with aliyun or acs:, and cannot contain http:// or https://.
	//
	// Each key-value pair must be unique. You can specify values for at most 20 tag keys in each call.
	//
	// example:
	//
	// subnet-group-1
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
	// The CIDR block of the VPD.
	//
	// 	- We recommend that you use an RFC private endpoint as the Lingjun CIDR block, such as 10.0.0.0/8,172.16.0.0/12,192.168.0.0/16. In scenarios where the Doringjun CIDR block is connected to each other or where the Lingjun CIDR block is connected to a VPC, make sure that the addresses do not conflict with each other.
	//
	// 	- You can also use a custom CIDR block other than 100.64.0.0/10, 224.0.0.0/4, 127.0.0.0/8, or 169.254.0.0/16 and their subnets as the primary IPv4 CIDR block of the VPD.
	//
	// example:
	//
	// 10.0.0.0/8
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The time when the data address was created.
	//
	// example:
	//
	// 1678273219000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Lingjun CIDR block instance ID
	//
	// example:
	//
	// vpd-d3isyds4
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// Lingjun CIDR block instance name
	//
	// example:
	//
	// yzp-rg-test3
	VpdName *string `json:"VpdName,omitempty" xml:"VpdName,omitempty"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSubnetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ListVccFlowInfosRequest struct {
	// Direction
	//
	// Valid value:
	//
	// 	- IN: inbound.
	//
	// 	- OUT: the outbound.
	//
	// example:
	//
	// OUT
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The start time. The default value is 5 minutes ago.
	//
	// example:
	//
	// 1667727514000
	From *int64 `json:"From,omitempty" xml:"From,omitempty"`
	// Metric
	//
	// Valid value:
	//
	// 	- totalPacketsRate: Total Packet Rate.
	//
	// 	- dropBytesRate: the of the stream drop rate.
	//
	// 	- dropPacketsRate: Dropped Packet Rate.
	//
	// 	- totalBytesRate: the total streaming rate.
	//
	// 	- passBytesRate: by stream rate.
	//
	// 	- passPacketsRate: by packet rate.
	//
	// example:
	//
	// passBytesRate
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The end time. The default time is the current time.
	//
	// example:
	//
	// 1689749749000
	To *int64 `json:"To,omitempty" xml:"To,omitempty"`
	// Lingjun Connection ID
	//
	// example:
	//
	// vcc-cn-zvp2******
	VccId *string `json:"VccId,omitempty" xml:"VccId,omitempty"`
}

func (s ListVccFlowInfosRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVccFlowInfosRequest) GoString() string {
	return s.String()
}

func (s *ListVccFlowInfosRequest) SetDirection(v string) *ListVccFlowInfosRequest {
	s.Direction = &v
	return s
}

func (s *ListVccFlowInfosRequest) SetFrom(v int64) *ListVccFlowInfosRequest {
	s.From = &v
	return s
}

func (s *ListVccFlowInfosRequest) SetMetricName(v string) *ListVccFlowInfosRequest {
	s.MetricName = &v
	return s
}

func (s *ListVccFlowInfosRequest) SetRegionId(v string) *ListVccFlowInfosRequest {
	s.RegionId = &v
	return s
}

func (s *ListVccFlowInfosRequest) SetTo(v int64) *ListVccFlowInfosRequest {
	s.To = &v
	return s
}

func (s *ListVccFlowInfosRequest) SetVccId(v string) *ListVccFlowInfosRequest {
	s.VccId = &v
	return s
}

type ListVccFlowInfosResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Content *ListVccFlowInfosResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// Response
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// BDBCC783-84CA-5733-8EEA-************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListVccFlowInfosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVccFlowInfosResponseBody) GoString() string {
	return s.String()
}

func (s *ListVccFlowInfosResponseBody) SetAccessDeniedDetail(v string) *ListVccFlowInfosResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListVccFlowInfosResponseBody) SetCode(v int32) *ListVccFlowInfosResponseBody {
	s.Code = &v
	return s
}

func (s *ListVccFlowInfosResponseBody) SetContent(v *ListVccFlowInfosResponseBodyContent) *ListVccFlowInfosResponseBody {
	s.Content = v
	return s
}

func (s *ListVccFlowInfosResponseBody) SetMessage(v string) *ListVccFlowInfosResponseBody {
	s.Message = &v
	return s
}

func (s *ListVccFlowInfosResponseBody) SetRequestId(v string) *ListVccFlowInfosResponseBody {
	s.RequestId = &v
	return s
}

type ListVccFlowInfosResponseBodyContent struct {
	// Lingjun Connection Traffic Information
	Data []*ListVccFlowInfosResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListVccFlowInfosResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s ListVccFlowInfosResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListVccFlowInfosResponseBodyContent) SetData(v []*ListVccFlowInfosResponseBodyContentData) *ListVccFlowInfosResponseBodyContent {
	s.Data = v
	return s
}

func (s *ListVccFlowInfosResponseBodyContent) SetTotal(v int64) *ListVccFlowInfosResponseBodyContent {
	s.Total = &v
	return s
}

type ListVccFlowInfosResponseBodyContentData struct {
	// The direction.
	//
	// example:
	//
	// OUT
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The metric. Valid values:
	//
	// example:
	//
	// passBytesRate
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Time
	//
	// example:
	//
	// 1689749749000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// Value
	//
	// example:
	//
	// 123
	Value *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
	// Lingjun Connection ID
	//
	// example:
	//
	// vcc-cn-zvp2w******
	VccId *string `json:"VccId,omitempty" xml:"VccId,omitempty"`
}

func (s ListVccFlowInfosResponseBodyContentData) String() string {
	return tea.Prettify(s)
}

func (s ListVccFlowInfosResponseBodyContentData) GoString() string {
	return s.String()
}

func (s *ListVccFlowInfosResponseBodyContentData) SetDirection(v string) *ListVccFlowInfosResponseBodyContentData {
	s.Direction = &v
	return s
}

func (s *ListVccFlowInfosResponseBodyContentData) SetMetricName(v string) *ListVccFlowInfosResponseBodyContentData {
	s.MetricName = &v
	return s
}

func (s *ListVccFlowInfosResponseBodyContentData) SetRegionId(v string) *ListVccFlowInfosResponseBodyContentData {
	s.RegionId = &v
	return s
}

func (s *ListVccFlowInfosResponseBodyContentData) SetTimestamp(v int64) *ListVccFlowInfosResponseBodyContentData {
	s.Timestamp = &v
	return s
}

func (s *ListVccFlowInfosResponseBodyContentData) SetValue(v float64) *ListVccFlowInfosResponseBodyContentData {
	s.Value = &v
	return s
}

func (s *ListVccFlowInfosResponseBodyContentData) SetVccId(v string) *ListVccFlowInfosResponseBodyContentData {
	s.VccId = &v
	return s
}

type ListVccFlowInfosResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVccFlowInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVccFlowInfosResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVccFlowInfosResponse) GoString() string {
	return s.String()
}

func (s *ListVccFlowInfosResponse) SetHeaders(v map[string]*string) *ListVccFlowInfosResponse {
	s.Headers = v
	return s
}

func (s *ListVccFlowInfosResponse) SetStatusCode(v int32) *ListVccFlowInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVccFlowInfosResponse) SetBody(v *ListVccFlowInfosResponseBody) *ListVccFlowInfosResponse {
	s.Body = v
	return s
}

type ListVccGrantRulesRequest struct {
	// Specifies whether to enable paged query. Optional values:
	//
	// 	- **true**: Enable pagination query
	//
	// 	- **false**: Pagination query is disabled
	//
	// example:
	//
	// false
	EnablePage *bool `json:"EnablePage,omitempty" xml:"EnablePage,omitempty"`
	// Lingjun HUB ID
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// Use the drop-down box
	//
	// example:
	//
	// true
	ForSelect *bool `json:"ForSelect,omitempty" xml:"ForSelect,omitempty"`
	// Authorization Entry ID
	//
	// example:
	//
	// grant-rule-jaj33d1b804
	GrantRuleId *string `json:"GrantRuleId,omitempty" xml:"GrantRuleId,omitempty"`
	// Authorized Tenant ID
	//
	// example:
	//
	// 1620939556166277
	GrantTenantId *string `json:"GrantTenantId,omitempty" xml:"GrantTenantId,omitempty"`
	// Network Instance ID
	//
	// example:
	//
	// vcc-cn-jaj33d1b804
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Instance name
	//
	// example:
	//
	// vcc-1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The page number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-aekzb3n5lgk2ieq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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

func (s *ListVccGrantRulesRequest) SetResourceGroupId(v string) *ListVccGrantRulesRequest {
	s.ResourceGroupId = &v
	return s
}

type ListVccGrantRulesResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Content *ListVccGrantRulesResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// A56F7D3C-8850-5AF4-A342-2D71C9A9D1CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListVccGrantRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVccGrantRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVccGrantRulesResponseBody) SetAccessDeniedDetail(v string) *ListVccGrantRulesResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// List of cross-account authorization information of Lingjun connection
	Data []*ListVccGrantRulesResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 0
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
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
	// The time when the data address was created.
	//
	// example:
	//
	// 1678273219000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Lingjun HUB ID
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// Cross-account authorization information Instance ID
	//
	// example:
	//
	// grant-rule-jpumgwvp
	GrantRuleId *string `json:"GrantRuleId,omitempty" xml:"GrantRuleId,omitempty"`
	// Authorized Tenant ID
	//
	// example:
	//
	// 1013666993027780
	GrantTenantId *string `json:"GrantTenantId,omitempty" xml:"GrantTenantId,omitempty"`
	// Network Instance ID
	//
	// example:
	//
	// vcc-cn-jaj33d1kb05
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the ECU.
	//
	// example:
	//
	// vcc-1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The type of the authorized product. Valid values:
	//
	// 	- **VPD**: indicates a VPD instance of the Lingjun network segment.
	//
	// 	- **VCC**: indicates that Lingjun connects to the VCC instance.
	//
	// example:
	//
	// VCC
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-aek2l4sq6l7unhi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 1655449505171
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// Whether the current cross-account resource has been bound to the cross-account Lingjun HUB. Valid values:
	//
	// 	- **true**: Used
	//
	// 	- **false**: Not used
	//
	// example:
	//
	// true
	Used *bool `json:"Used,omitempty" xml:"Used,omitempty"`
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

func (s *ListVccGrantRulesResponseBodyContentData) SetResourceGroupId(v string) *ListVccGrantRulesResponseBodyContentData {
	s.ResourceGroupId = &v
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVccGrantRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// Destination CIDR block
	//
	// example:
	//
	// 0.0.0.0/0
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// Specifies whether to enable pagination query.
	//
	// example:
	//
	// false
	EnablePage *bool `json:"EnablePage,omitempty" xml:"EnablePage,omitempty"`
	// Filter 32 detailed CIDR blocks. Default value: true
	//
	// example:
	//
	// true
	IgnoreDetailedRouteEntry *bool `json:"IgnoreDetailedRouteEntry,omitempty" xml:"IgnoreDetailedRouteEntry,omitempty"`
	// Next Hop Instance
	//
	// example:
	//
	// vcc-cn-jaj34d75h01
	NextHopId *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// Next Hop Instance Type
	//
	// example:
	//
	// VCC
	NextHopType *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	// The page number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-aek2l4sq6l7unhi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Route type
	//
	// example:
	//
	// BGP
	RouteType *string `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
	// The status of the enterprise-level snapshot policy.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the Lingjun connection instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vcc-cn-zvp2w222001
	VccId *string `json:"VccId,omitempty" xml:"VccId,omitempty"`
	// Lingjun CIDR block route entry instance ID
	//
	// example:
	//
	// vpd-rte-toekyqel
	VpdRouteEntryId *string `json:"VpdRouteEntryId,omitempty" xml:"VpdRouteEntryId,omitempty"`
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

func (s *ListVccRouteEntriesRequest) SetIgnoreDetailedRouteEntry(v bool) *ListVccRouteEntriesRequest {
	s.IgnoreDetailedRouteEntry = &v
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

func (s *ListVccRouteEntriesRequest) SetResourceGroupId(v string) *ListVccRouteEntriesRequest {
	s.ResourceGroupId = &v
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Content *ListVccRouteEntriesResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// response message, if the success request is
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// A88DFED5-24B7-5A3E-87DE-380BF06F3C90
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListVccRouteEntriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVccRouteEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVccRouteEntriesResponseBody) SetAccessDeniedDetail(v string) *ListVccRouteEntriesResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// List of Lingjun Connection Route Entries
	Data []*ListVccRouteEntriesResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 0
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
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
	// Destination CIDR block
	//
	// example:
	//
	// 10.192.32.0/24
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The time when the cluster was updated.
	//
	// example:
	//
	// 1642745758000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Next Hop Instance
	//
	// example:
	//
	// vcc-cn-zvp2w222001
	NextHopId *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// Next Hop Type
	//
	// example:
	//
	// VCC
	NextHopType *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-aek2l4sq6l7unhi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the tenant to which the resource belongs.
	//
	// example:
	//
	// 1655449505171
	ResourceTenantId *string `json:"ResourceTenantId,omitempty" xml:"ResourceTenantId,omitempty"`
	// Route type
	//
	// example:
	//
	// BGP
	RouteType *string `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
	// The status of the intervention entry. Valid value:
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 1655449505171
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The ID of the Lingjun connection instance.
	//
	// example:
	//
	// vcc-cn-zvp2w222001
	VccId *string `json:"VccId,omitempty" xml:"VccId,omitempty"`
	// The ID of the route entry.
	//
	// example:
	//
	// vcc-rte-maysfadg
	VccRouteEntryId *string `json:"VccRouteEntryId,omitempty" xml:"VccRouteEntryId,omitempty"`
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

func (s *ListVccRouteEntriesResponseBodyContentData) SetMessage(v string) *ListVccRouteEntriesResponseBodyContentData {
	s.Message = &v
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

func (s *ListVccRouteEntriesResponseBodyContentData) SetResourceGroupId(v string) *ListVccRouteEntriesResponseBodyContentData {
	s.ResourceGroupId = &v
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVccRouteEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The peak bandwidth of the Lingjun connection instance. Unit: Mbit/s. Valid values: 1000 to 400000
	//
	// example:
	//
	// 5000
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The ID of the CEN instance; [What is the CEN?](https://help.aliyun.com/document_detail/181681.html)
	//
	// You can call the [DescribeCens](https://help.aliyun.com/document_detail/468215.htm) to query the information of CEN instances under the current Alibaba Cloud account.
	//
	// example:
	//
	// cen-95iwtpyvj3kk1v0ao0
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// Specifies whether to enable paged query. Optional values:
	//
	// 	- **true**: Enable pagination query
	//
	// 	- **false**: Pagination query is disabled
	//
	// example:
	//
	// false
	EnablePage *bool `json:"EnablePage,omitempty" xml:"EnablePage,omitempty"`
	// Excludes all data in the specified status. If the status parameter exists, ExStatus does not take effect.
	//
	// example:
	//
	// Prepaid
	ExStatus *string `json:"ExStatus,omitempty" xml:"ExStatus,omitempty"`
	// Filter queries by Lingjun HUB instance ID
	//
	// example:
	//
	// er-a7rqv1rq
	FilterErId *string `json:"FilterErId,omitempty" xml:"FilterErId,omitempty"`
	// The page number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// For more information about resource groups, see [Resource groups](https://help.aliyun.com/document_detail/94475.htm?spm=a2c4g.11186623.0.0.29e15d7akXhpuu).
	//
	// example:
	//
	// rg-aeky5f3qx6ceapq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The instance status.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag information.
	//
	// You can specify up to 20 tags.
	Tag []*ListVccsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the Lingjun connection instance.
	//
	// example:
	//
	// vcc-cn-zvp2w222001
	VccId *string `json:"VccId,omitempty" xml:"VccId,omitempty"`
	// Virtual Private Cloud IDs; [What is Virtual Private Cloud](https://help.aliyun.com/document_detail/34217.html)
	//
	// You can call the [DescribeVpcs](https://help.aliyun.com/document_detail/35739.html#demo-0) operation to query the specified VPC.
	//
	// example:
	//
	// vpc-bp1nrtkmamy329u6a1z0i
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// Lingjun CIDR block instance ID
	//
	// example:
	//
	// vpd-omqutbff
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
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
	// The tag key of the VPN attachment.
	//
	// You cannot specify an empty string as a tag key. It can be up to 64 characters in length and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// You can specify at most 20 tag keys in each call.
	//
	// example:
	//
	// tag-vcc
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the VPN connection.
	//
	// The tag value can be empty or a string of up to 128 characters. It cannot start with aliyun or acs:, and cannot contain http:// or https://.
	//
	// Each key-value pair must be unique. You can specify values for at most 20 tag keys in each call.
	//
	// example:
	//
	// vcc-group-1
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Content *ListVccsResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// 28451248-7038-5184-B5D3-80F104654BE8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListVccsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVccsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVccsResponseBody) SetAccessDeniedDetail(v string) *ListVccsResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// Lingjun Connection Information List
	Data []*ListVccsResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 0
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
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
	// Express Connect circuit access point ID:
	//
	// 	- **ap-cn-wulanchabu-jn-ts-A**: Ulanqab-Jining-A
	//
	// 	- **ap-cn-heyuan-yc-ts-SA127**: Heyuan-Yuancheng-A
	//
	// example:
	//
	// ap-cn-wulanchabu-jn-ts-A
	AccessPointId *string `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	// The bandwidth of the port.
	//
	// example:
	//
	// 1000
	BandwidthStr *string `json:"BandwidthStr,omitempty" xml:"BandwidthStr,omitempty"`
	// bgp as number
	//
	// example:
	//
	// bgpAsn
	BgpAsn *string `json:"BgpAsn,omitempty" xml:"BgpAsn,omitempty"`
	// bgp network segment
	//
	// example:
	//
	// 172.16.128.0/24
	BgpCidr *string `json:"BgpCidr,omitempty" xml:"BgpCidr,omitempty"`
	// The ID of the CEN instance; [What is the CEN?](https://help.aliyun.com/document_detail/181681.html)
	//
	// You can call the [DescribeCens](https://help.aliyun.com/document_detail/468215.htm) to query the information of CEN instances under the current Alibaba Cloud account.
	//
	// example:
	//
	// cen-w15qot0pfvs83pkckj
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// Account to which cen belongs
	//
	// example:
	//
	// 1238685214107736
	CenOwnerId *string `json:"CenOwnerId,omitempty" xml:"CenOwnerId,omitempty"`
	// Commodity code
	//
	// example:
	//
	// bccluster_cloudconnectionpre_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The connection mode. Valid values:
	//
	// 	- **VPC**
	//
	// 	- **CENTR**
	//
	// example:
	//
	// CENTR
	ConnectionType *string `json:"ConnectionType,omitempty" xml:"ConnectionType,omitempty"`
	// The time when the data address was created.
	//
	// example:
	//
	// 1678273219000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Current process node
	//
	// example:
	//
	// test-xxxx-node-x
	CurrentNode *string `json:"CurrentNode,omitempty" xml:"CurrentNode,omitempty"`
	// List of bound Lingjun HUB information
	ErInfos []*ListVccsResponseBodyContentDataErInfos `json:"ErInfos,omitempty" xml:"ErInfos,omitempty" type:"Repeated"`
	// The time when the application expired.
	//
	// example:
	//
	// 1678273219000
	ExpirationDate *string `json:"ExpirationDate,omitempty" xml:"ExpirationDate,omitempty"`
	// The time when the cluster was updated.
	//
	// example:
	//
	// 1678273219000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The connectivity provider of the Express Connect circuit. Valid values:
	//
	// 	- **CO**: other connectivity providers in the Chinese mainland
	//
	// example:
	//
	// CO
	LineOperator *string `json:"LineOperator,omitempty" xml:"LineOperator,omitempty"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// some message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The port type of the Express Connect circuit. Valid values:
	//
	// 	- **100GBase-LR**: 100,000 megabytes of single-mode optical port (10 km)
	//
	// example:
	//
	// 100GBase-LR
	PortType *string `json:"PortType,omitempty" xml:"PortType,omitempty"`
	// Process progress; value returns 0 to 1; not started is null
	//
	// example:
	//
	// 1
	Rate *float64 `json:"Rate,omitempty" xml:"Rate,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of your Alibaba Cloud resource group.
	//
	// For more information about resource groups, see [Resource groups](https://help.aliyun.com/document_detail/94475.htm?spm=a2c4g.11186623.0.0.29e15d7akXhpuu).
	//
	// example:
	//
	// rg-aek2l4sq6l7unhi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The compute specification.
	//
	// example:
	//
	// Large
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The state of the rule.
	//
	// example:
	//
	// Init
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag information.
	//
	// You can specify up to 20 tags.
	Tags []*ListVccsResponseBodyContentDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The job ID.
	//
	// example:
	//
	// task-cd544092-ed0a-49e9-83eb-e8c94770dccf
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 1620939556166279
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The ID of the Lingjun connection instance.
	//
	// example:
	//
	// vcc-cn-zvp2w222001
	VccId *string `json:"VccId,omitempty" xml:"VccId,omitempty"`
	// The name of the Lingjun connection instance.
	//
	// example:
	//
	// vcc-heyuan-backup
	VccName *string `json:"VccName,omitempty" xml:"VccName,omitempty"`
	// Virtual Private Cloud IDs; [What is Virtual Private Cloud](https://help.aliyun.com/document_detail/34217.html)
	//
	// You can call the [DescribeVpcs](https://help.aliyun.com/document_detail/35739.html#demo-0) operation to query the specified VPC.
	//
	// example:
	//
	// vpc-f8ziirfl9k25h2qn7y4f8
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// Lingjun network segment information (applicable to the scene where the old version of Lingjun connection is directly bound to Lingjun network segment)
	VpdBaseInfo *ListVccsResponseBodyContentDataVpdBaseInfo `json:"VpdBaseInfo,omitempty" xml:"VpdBaseInfo,omitempty" type:"Struct"`
	// Lingjun CIDR block instance ID
	//
	// example:
	//
	// vpd-eoiy88ju
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-wulanchabu-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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

func (s *ListVccsResponseBodyContentData) SetBgpAsn(v string) *ListVccsResponseBodyContentData {
	s.BgpAsn = &v
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

func (s *ListVccsResponseBodyContentData) SetCenOwnerId(v string) *ListVccsResponseBodyContentData {
	s.CenOwnerId = &v
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
	// Connections
	//
	// example:
	//
	// 2
	Connections *int64 `json:"Connections,omitempty" xml:"Connections,omitempty"`
	// The time when the data address was created.
	//
	// example:
	//
	// 1678273219000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Description
	//
	// example:
	//
	// test_api_coverage
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Elastic Router ID
	//
	// example:
	//
	// er-a7rqv1rq
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// ER instance name
	//
	// example:
	//
	// er-1
	ErName *string `json:"ErName,omitempty" xml:"ErName,omitempty"`
	// The time when the agent was last modified.
	//
	// example:
	//
	// 1678273219000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Primary Zone
	//
	// example:
	//
	// cn-wulanchabu-b
	MasterZoneId *string `json:"MasterZoneId,omitempty" xml:"MasterZoneId,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// ER region information
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Number of routing policy
	//
	// example:
	//
	// 2
	RouteMaps *int64 `json:"RouteMaps,omitempty" xml:"RouteMaps,omitempty"`
	// The status of the intervention entry. Valid value:
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 1620939556166277
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
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
	// The tag key.
	//
	// You cannot specify an empty string as a tag key. It can be up to 64 characters in length and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// You can specify at most 20 tag keys in each call.
	//
	// example:
	//
	// tag-vcc
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag that is added to the resource.
	//
	// The tag value can be empty or a string of up to 128 characters. It cannot start with aliyun or acs:, and cannot contain http:// or https://.
	//
	// Each key-value pair must be unique. You can specify values for at most 20 tag keys in each call.
	//
	// example:
	//
	// vcc-group-1
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
	// The CIDR block of the VPD.
	//
	// 	- We recommend that you use an RFC private endpoint as the Lingjun CIDR block, such as 10.0.0.0/8,172.16.0.0/12,192.168.0.0/16. In scenarios where the Doringjun CIDR block is connected to each other or where the Lingjun CIDR block is connected to a VPC, make sure that the addresses do not conflict with each other.
	//
	// 	- You can also use a custom CIDR block other than 100.64.0.0/10, 224.0.0.0/4, 127.0.0.0/8, or 169.254.0.0/16 and their subnets as the primary IPv4 CIDR block of the VPD.
	//
	// example:
	//
	// 10.0.0.0/13
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The time when the data address was created.
	//
	// example:
	//
	// 1668158213000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Lingjun CIDR block instance ID
	//
	// example:
	//
	// vpd-9n7ioqrp
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// Lingjun CIDR block instance name
	//
	// example:
	//
	// yzp-rg-test3
	VpdName *string `json:"VpdName,omitempty" xml:"VpdName,omitempty"`
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
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVccsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// Specifies whether to enable pagination query.
	//
	// example:
	//
	// false
	EnablePage *bool `json:"EnablePage,omitempty" xml:"EnablePage,omitempty"`
	// Lingjun HUB Instance ID
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// Use the drop-down box
	//
	// example:
	//
	// true
	ForSelect *bool `json:"ForSelect,omitempty" xml:"ForSelect,omitempty"`
	// Authorization Entry ID
	//
	// example:
	//
	// grant-rule-8rgvqazb
	GrantRuleId *string `json:"GrantRuleId,omitempty" xml:"GrantRuleId,omitempty"`
	// Authorized Tenant ID
	//
	// example:
	//
	// 1620939556166279
	GrantTenantId *string `json:"GrantTenantId,omitempty" xml:"GrantTenantId,omitempty"`
	// The ID of the network instance that you want to query.
	//
	// example:
	//
	// vpd-8rgvqazb
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Instance name
	//
	// example:
	//
	// vpd-1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The page number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-aek2l4sq6l7unhi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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

func (s *ListVpdGrantRulesRequest) SetResourceGroupId(v string) *ListVpdGrantRulesRequest {
	s.ResourceGroupId = &v
	return s
}

type ListVpdGrantRulesResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Content *ListVpdGrantRulesResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// A56F7D3C-8850-5AF4-A342-2D71C9A9D1CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListVpdGrantRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVpdGrantRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpdGrantRulesResponseBody) SetAccessDeniedDetail(v string) *ListVpdGrantRulesResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// Lingjun CIDR block authorization information list
	Data []*ListVpdGrantRulesResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 0
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
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
	// The time when the data address was created.
	//
	// example:
	//
	// 1643013506000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The current network sample is authorized to the specified Lingjun HUB sample ID.
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// Authorization Entry ID
	//
	// example:
	//
	// grant-rule-8rgvqazb
	GrantRuleId *string `json:"GrantRuleId,omitempty" xml:"GrantRuleId,omitempty"`
	// The ID of the tenant to which the current instance is authorized.
	//
	// example:
	//
	// 1672372231790
	GrantTenantId *string `json:"GrantTenantId,omitempty" xml:"GrantTenantId,omitempty"`
	// Lingjun CIDR block instance ID
	//
	// example:
	//
	// vpd-8rgvqazb
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the ECU.
	//
	// example:
	//
	// vpd-1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The type of the authorized product. Valid values:
	//
	// 	- **VPD**: indicates a VPD instance of the Lingjun network segment.
	//
	// 	- **VCC**: indicates that Lingjun connects to the VCC instance.
	//
	// The caller does not need to specify.
	//
	// example:
	//
	// VPD
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-aek2l4sq6l7unhi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 1655449505171
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// Whether the current authorized instance has been bound
	//
	// example:
	//
	// true
	Used *bool `json:"Used,omitempty" xml:"Used,omitempty"`
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

func (s *ListVpdGrantRulesResponseBodyContentData) SetResourceGroupId(v string) *ListVpdGrantRulesResponseBodyContentData {
	s.ResourceGroupId = &v
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVpdGrantRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// Destination CIDR block
	//
	// example:
	//
	// 0.0.0.0/0
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// Specifies whether to enable paged query. Optional values:
	//
	// 	- **true**: Enable pagination query
	//
	// 	- **false**: Pagination query is disabled
	//
	// example:
	//
	// false
	EnablePage *bool `json:"EnablePage,omitempty" xml:"EnablePage,omitempty"`
	// Filter 32 detailed CIDR blocks. Default value: true
	//
	// example:
	//
	// true
	IgnoreDetailedRouteEntry *bool `json:"IgnoreDetailedRouteEntry,omitempty" xml:"IgnoreDetailedRouteEntry,omitempty"`
	// Next Hop Instance
	//
	// example:
	//
	// vcc-cn-209300qha01
	NextHopId *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// Next Hop Instance Type
	//
	// example:
	//
	// VCC
	NextHopType *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	// The page number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-acfm4mlwqjalz7a
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Route type
	//
	// example:
	//
	// BGP
	RouteType *string `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
	// The status of the enterprise-level snapshot policy.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Lingjun CIDR block instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// vpd-fuliephf
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// Lingjun CIDR block route entry instance ID
	//
	// example:
	//
	// vpd-rte-4r1zbhoh
	VpdRouteEntryId *string `json:"VpdRouteEntryId,omitempty" xml:"VpdRouteEntryId,omitempty"`
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

func (s *ListVpdRouteEntriesRequest) SetIgnoreDetailedRouteEntry(v bool) *ListVpdRouteEntriesRequest {
	s.IgnoreDetailedRouteEntry = &v
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

func (s *ListVpdRouteEntriesRequest) SetResourceGroupId(v string) *ListVpdRouteEntriesRequest {
	s.ResourceGroupId = &v
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Content *ListVpdRouteEntriesResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// 039C3C3A-3C37-5672-80D5-D8CD48C676D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListVpdRouteEntriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVpdRouteEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpdRouteEntriesResponseBody) SetAccessDeniedDetail(v string) *ListVpdRouteEntriesResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// Lingjun CIDR block route entry list
	Data []*ListVpdRouteEntriesResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 0
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
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
	// Destination CIDR block
	//
	// example:
	//
	// 0.0.0.0/0
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The time when the cluster was updated.
	//
	// example:
	//
	// 1678273219000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Next Hop Instance
	//
	// example:
	//
	// er-bmlqiym1
	NextHopId *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// Next Hop Instance Type
	//
	// example:
	//
	// ER
	NextHopType *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-acfmxhucx5ewuwy
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the tenant to which the resource belongs.
	//
	// example:
	//
	// 1655449505171
	ResourceTenantId *string `json:"ResourceTenantId,omitempty" xml:"ResourceTenantId,omitempty"`
	// Route type
	//
	// example:
	//
	// BGP
	RouteType *string `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
	// The status of the intervention entry. Valid value:
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 1655449505171
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// Lingjun CIDR block instance ID
	//
	// example:
	//
	// vpd-eoiy88ju
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// The ID of the route entry.
	//
	// example:
	//
	// vpd-rte-toekyqel
	VpdRouteEntryId *string `json:"VpdRouteEntryId,omitempty" xml:"VpdRouteEntryId,omitempty"`
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

func (s *ListVpdRouteEntriesResponseBodyContentData) SetResourceGroupId(v string) *ListVpdRouteEntriesResponseBodyContentData {
	s.ResourceGroupId = &v
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVpdRouteEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// Specifies whether to enable paged query.
	//
	// example:
	//
	// false
	EnablePage *bool `json:"EnablePage,omitempty" xml:"EnablePage,omitempty"`
	// Queries the network segments of Lingjun that are not bound to a specified Lingjun HUB.
	//
	// example:
	//
	// er-kkopgtne
	FilterErId *string `json:"FilterErId,omitempty" xml:"FilterErId,omitempty"`
	// If you select a drop-down list, only the basic information (including the instance ID and instance name) is returned. Possible values:
	//
	// 	- **true**: Select Query Use from the drop-down list.
	//
	// 	- **false**: Normal queries are used.
	//
	// example:
	//
	// true
	ForSelect *bool `json:"ForSelect,omitempty" xml:"ForSelect,omitempty"`
	// The page number of the page to return. Start value: 1 Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// For more information about resource groups, see [Resource groups](https://help.aliyun.com/document_detail/94475.htm?spm=a2c4g.11186623.0.0.29e15d7akXhpuu).
	//
	// example:
	//
	// rg-aeky5f3qx6ceapq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the CLB instance. Valid values:
	//
	// 	- **Available**: Normal.
	//
	// 	- **Not Available**: Not available.
	//
	// 	- **Executing**: The task is being executed.
	//
	// 	- **Deleting**: The account is being deleted
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag information.
	//
	// You can specify up to 20 tags.
	Tag []*ListVpdsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the VPD instance.
	//
	// example:
	//
	// vpd-fuliephf
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// The name of the VPD instance.
	//
	// example:
	//
	// vpd-1
	VpdName *string `json:"VpdName,omitempty" xml:"VpdName,omitempty"`
	// Specifies whether to include the dependent resource information. We recommend that you do not query the dependent resource information when you query by page. You can query the dependent resource information separately when you delete it. Possible values:
	//
	// 	- **true**: with dependency information.
	//
	// 	- **false**: does not include dependency information.
	//
	// example:
	//
	// false
	WithDependence *bool `json:"WithDependence,omitempty" xml:"WithDependence,omitempty"`
	// Queries the information about a Lingjun CIDR block that is not bound to a Lingjun connection. Possible values:
	//
	// 	- **true**: filters out VPDs that have been bound to VCC
	//
	// 	- **false**: does not filter VPD that has been bound to VCC
	//
	// example:
	//
	// true
	WithoutVcc *bool `json:"WithoutVcc,omitempty" xml:"WithoutVcc,omitempty"`
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
	// The tag key of the VPN attachment.
	//
	// You cannot specify an empty string as a tag key. It can be up to 64 characters in length and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// You can specify at most 20 tag keys in each call.
	//
	// example:
	//
	// tag-vpd-region
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the VPN connection.
	//
	// The tag value can be empty or a string of up to 128 characters. It cannot start with aliyun or acs:, and cannot contain http:// or https://.
	//
	// Each key-value pair must be unique. You can specify values for at most 20 tag keys in each call.
	//
	// example:
	//
	// wulanchabu
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Content *ListVpdsResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9C50C9CD-E799-54DA-BA7A-1FAF3DF80857
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListVpdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVpdsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpdsResponseBody) SetAccessDeniedDetail(v string) *ListVpdsResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// The returned data.
	Data []*ListVpdsResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The total number of entries that are returned.
	//
	// example:
	//
	// 0
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
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
	// The CIDR block of the VPD.
	//
	// 	- We recommend that you use an RFC private endpoint as the Lingjun CIDR block, such as 10.0.0.0/8,172.16.0.0/12,192.168.0.0/16. In scenarios where the Doringjun CIDR block is connected to each other or where the Lingjun CIDR block is connected to a VPC, make sure that the addresses do not conflict with each other.
	//
	// 	- You can also use a custom CIDR block other than 100.64.0.0/10, 224.0.0.0/4, 127.0.0.0/8, or 169.254.0.0/16 and their subnets as the primary IPv4 CIDR block of the VPD.
	//
	// example:
	//
	// 10.0.0.0/8
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The time when the activation code was created.
	//
	// example:
	//
	// 1678273219000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Dependencies.
	Dependence map[string]interface{} `json:"Dependence,omitempty" xml:"Dependence,omitempty"`
	// The information list of the bound Lingjun HUB(ER).
	ErInfos []*ListVpdsResponseBodyContentDataErInfos `json:"ErInfos,omitempty" xml:"ErInfos,omitempty" type:"Repeated"`
	// The time when the O\\&M task was modified.
	//
	// example:
	//
	// 1678273219000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// nc quantity.
	//
	// example:
	//
	// 16
	NcCount *int32 `json:"NcCount,omitempty" xml:"NcCount,omitempty"`
	// Number of Lingjun network interface controller
	//
	// example:
	//
	// 1
	NetworkInterfaceCount *int32 `json:"NetworkInterfaceCount,omitempty" xml:"NetworkInterfaceCount,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of your Alibaba Cloud resource group.
	//
	// For more information about resource groups, see [Resource groups](https://help.aliyun.com/document_detail/94475.htm?spm=a2c4g.11186623.0.0.29e15d7akXhpuu).
	//
	// example:
	//
	// rg-aek2l4sq6l7unhi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The list of additional CIDR blocks.
	SecondaryCidrBlocks []*string `json:"SecondaryCidrBlocks,omitempty" xml:"SecondaryCidrBlocks,omitempty" type:"Repeated"`
	// The Service CIDR block.
	//
	// example:
	//
	// 169.254.252.0/23
	ServiceCidr *string `json:"ServiceCidr,omitempty" xml:"ServiceCidr,omitempty"`
	// The task status.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of subnets.
	//
	// example:
	//
	// 1
	SubnetCount *int32 `json:"SubnetCount,omitempty" xml:"SubnetCount,omitempty"`
	// The tag information.
	//
	// You can specify up to 20 tags.
	Tags []*ListVpdsResponseBodyContentDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The tenant ID.
	//
	// example:
	//
	// 1655449505171
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The ID of the VPD instance.
	//
	// example:
	//
	// vpd-lg4dppgi
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// The name of the VPD.
	//
	// example:
	//
	// vpd-1
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

func (s *ListVpdsResponseBodyContentData) SetNetworkInterfaceCount(v int32) *ListVpdsResponseBodyContentData {
	s.NetworkInterfaceCount = &v
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

func (s *ListVpdsResponseBodyContentData) SetSecondaryCidrBlocks(v []*string) *ListVpdsResponseBodyContentData {
	s.SecondaryCidrBlocks = v
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
	// The number of connections.
	//
	// example:
	//
	// 2
	Connections *int64 `json:"Connections,omitempty" xml:"Connections,omitempty"`
	// The time when the activation code was created.
	//
	// example:
	//
	// 2023-12-26 20:16:36
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the synchronization task.
	//
	// example:
	//
	// No description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the Elastic Router (ER) instance.
	//
	// example:
	//
	// er-63vzm0fw
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// The name of the Lingjun HUB(ER) instance.
	//
	// example:
	//
	// er-1
	ErName *string `json:"ErName,omitempty" xml:"ErName,omitempty"`
	// The time when the O\\&M task was modified.
	//
	// example:
	//
	// 2023-12-26 20:16:36
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The primary zone.
	//
	// example:
	//
	// cn-wulanchabu-b
	MasterZoneId *string `json:"MasterZoneId,omitempty" xml:"MasterZoneId,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The supported region.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of routing policy.
	//
	// example:
	//
	// 2
	RouteMaps *int64 `json:"RouteMaps,omitempty" xml:"RouteMaps,omitempty"`
	// The task status.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 1620939556166277
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
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
	// The tag key.
	//
	// You cannot specify an empty string as a tag key. It can be up to 64 characters in length and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// You can specify at most 20 tag keys in each call.
	//
	// example:
	//
	// tag-vpd-region
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag that is added to the resource.
	//
	// The tag value can be empty or a string of up to 128 characters. It cannot start with aliyun or acs:, and cannot contain http:// or https://.
	//
	// Each key-value pair must be unique. You can specify values for at most 20 tag keys in each call.
	//
	// example:
	//
	// cn-wulanchabu
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
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVpdsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type QueryInstanceNcdRequest struct {
	// Instance 1ID
	//
	// This parameter is required.
	//
	// example:
	//
	// lni-1235****
	InstanceId1 *string `json:"InstanceId1,omitempty" xml:"InstanceId1,omitempty"`
	// Instance 2ID
	//
	// This parameter is required.
	//
	// example:
	//
	// lni-1234****
	InstanceId2 *string `json:"InstanceId2,omitempty" xml:"InstanceId2,omitempty"`
	// The parameter that specifies the instance type.
	//
	// Valid value:
	//
	// 	- node: Lingjun node.
	//
	// 	- lni: lingjun network interface controller.
	//
	// This parameter is required.
	//
	// example:
	//
	// lni
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s QueryInstanceNcdRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryInstanceNcdRequest) GoString() string {
	return s.String()
}

func (s *QueryInstanceNcdRequest) SetInstanceId1(v string) *QueryInstanceNcdRequest {
	s.InstanceId1 = &v
	return s
}

func (s *QueryInstanceNcdRequest) SetInstanceId2(v string) *QueryInstanceNcdRequest {
	s.InstanceId2 = &v
	return s
}

func (s *QueryInstanceNcdRequest) SetInstanceType(v string) *QueryInstanceNcdRequest {
	s.InstanceType = &v
	return s
}

func (s *QueryInstanceNcdRequest) SetRegionId(v string) *QueryInstanceNcdRequest {
	s.RegionId = &v
	return s
}

type QueryInstanceNcdResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Content *QueryInstanceNcdResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// You don\\"t have the permission of this operation, action=eflo:QueryInstanceNcd, arn=acs:eflo:cn-shenzhen:1263399219805497:networkinterface/*, resourceGroup=null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// BDBCC783-84CA-5733-8EEA-645C88B9009C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryInstanceNcdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryInstanceNcdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInstanceNcdResponseBody) SetAccessDeniedDetail(v string) *QueryInstanceNcdResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryInstanceNcdResponseBody) SetCode(v int32) *QueryInstanceNcdResponseBody {
	s.Code = &v
	return s
}

func (s *QueryInstanceNcdResponseBody) SetContent(v *QueryInstanceNcdResponseBodyContent) *QueryInstanceNcdResponseBody {
	s.Content = v
	return s
}

func (s *QueryInstanceNcdResponseBody) SetMessage(v string) *QueryInstanceNcdResponseBody {
	s.Message = &v
	return s
}

func (s *QueryInstanceNcdResponseBody) SetRequestId(v string) *QueryInstanceNcdResponseBody {
	s.RequestId = &v
	return s
}

type QueryInstanceNcdResponseBodyContent struct {
	// Instance 1ID
	//
	// example:
	//
	// lni-1235****
	InstanceId1 *string `json:"InstanceId1,omitempty" xml:"InstanceId1,omitempty"`
	// Instance 2ID
	//
	// example:
	//
	// lni-1234****
	InstanceId2 *string `json:"InstanceId2,omitempty" xml:"InstanceId2,omitempty"`
	// Instance Type
	//
	// Valid value:
	//
	// 	- node: Lingjun node.
	//
	// 	- lni: lingjun network interface controller.
	//
	// example:
	//
	// lni
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// network communication distance between instances
	//
	// example:
	//
	// 1
	Ncd *int32 `json:"Ncd,omitempty" xml:"Ncd,omitempty"`
}

func (s QueryInstanceNcdResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryInstanceNcdResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryInstanceNcdResponseBodyContent) SetInstanceId1(v string) *QueryInstanceNcdResponseBodyContent {
	s.InstanceId1 = &v
	return s
}

func (s *QueryInstanceNcdResponseBodyContent) SetInstanceId2(v string) *QueryInstanceNcdResponseBodyContent {
	s.InstanceId2 = &v
	return s
}

func (s *QueryInstanceNcdResponseBodyContent) SetInstanceType(v string) *QueryInstanceNcdResponseBodyContent {
	s.InstanceType = &v
	return s
}

func (s *QueryInstanceNcdResponseBodyContent) SetNcd(v int32) *QueryInstanceNcdResponseBodyContent {
	s.Ncd = &v
	return s
}

type QueryInstanceNcdResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryInstanceNcdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryInstanceNcdResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryInstanceNcdResponse) GoString() string {
	return s.String()
}

func (s *QueryInstanceNcdResponse) SetHeaders(v map[string]*string) *QueryInstanceNcdResponse {
	s.Headers = v
	return s
}

func (s *QueryInstanceNcdResponse) SetStatusCode(v int32) *QueryInstanceNcdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryInstanceNcdResponse) SetBody(v *QueryInstanceNcdResponseBody) *QueryInstanceNcdResponse {
	s.Body = v
	return s
}

type RefundVccRequest struct {
	// Region
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Lingjun Connection ID
	//
	// example:
	//
	// vcc-cn-zvp2w222001
	VccId *string `json:"VccId,omitempty" xml:"VccId,omitempty"`
}

func (s RefundVccRequest) String() string {
	return tea.Prettify(s)
}

func (s RefundVccRequest) GoString() string {
	return s.String()
}

func (s *RefundVccRequest) SetRegionId(v string) *RefundVccRequest {
	s.RegionId = &v
	return s
}

func (s *RefundVccRequest) SetVccId(v string) *RefundVccRequest {
	s.VccId = &v
	return s
}

type RefundVccResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Response content
	//
	// example:
	//
	// {}
	Content interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// Response message, which is \\"success\\" if the request succeeds
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// AC8C713A-A9F4-5984-A5E1-76496DF35153
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefundVccResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefundVccResponseBody) GoString() string {
	return s.String()
}

func (s *RefundVccResponseBody) SetAccessDeniedDetail(v string) *RefundVccResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *RefundVccResponseBody) SetCode(v int32) *RefundVccResponseBody {
	s.Code = &v
	return s
}

func (s *RefundVccResponseBody) SetContent(v interface{}) *RefundVccResponseBody {
	s.Content = v
	return s
}

func (s *RefundVccResponseBody) SetMessage(v string) *RefundVccResponseBody {
	s.Message = &v
	return s
}

func (s *RefundVccResponseBody) SetRequestId(v string) *RefundVccResponseBody {
	s.RequestId = &v
	return s
}

type RefundVccResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefundVccResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefundVccResponse) String() string {
	return tea.Prettify(s)
}

func (s RefundVccResponse) GoString() string {
	return s.String()
}

func (s *RefundVccResponse) SetHeaders(v map[string]*string) *RefundVccResponse {
	s.Headers = v
	return s
}

func (s *RefundVccResponse) SetStatusCode(v int32) *RefundVccResponse {
	s.StatusCode = &v
	return s
}

func (s *RefundVccResponse) SetBody(v *RefundVccResponseBody) *RefundVccResponse {
	s.Body = v
	return s
}

type RetryVccRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Lingjun Connection ID
	//
	// example:
	//
	// vcc-cn-zvp2w222001
	VccId *string `json:"VccId,omitempty" xml:"VccId,omitempty"`
}

func (s RetryVccRequest) String() string {
	return tea.Prettify(s)
}

func (s RetryVccRequest) GoString() string {
	return s.String()
}

func (s *RetryVccRequest) SetRegionId(v string) *RetryVccRequest {
	s.RegionId = &v
	return s
}

func (s *RetryVccRequest) SetVccId(v string) *RetryVccRequest {
	s.VccId = &v
	return s
}

type RetryVccResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// {}
	Content interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The error message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// 7F9082CC-3D94-560F-A575-8E8EF6CE2CB8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RetryVccResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RetryVccResponseBody) GoString() string {
	return s.String()
}

func (s *RetryVccResponseBody) SetAccessDeniedDetail(v string) *RetryVccResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *RetryVccResponseBody) SetCode(v int32) *RetryVccResponseBody {
	s.Code = &v
	return s
}

func (s *RetryVccResponseBody) SetContent(v interface{}) *RetryVccResponseBody {
	s.Content = v
	return s
}

func (s *RetryVccResponseBody) SetMessage(v string) *RetryVccResponseBody {
	s.Message = &v
	return s
}

func (s *RetryVccResponseBody) SetRequestId(v string) *RetryVccResponseBody {
	s.RequestId = &v
	return s
}

type RetryVccResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetryVccResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetryVccResponse) String() string {
	return tea.Prettify(s)
}

func (s RetryVccResponse) GoString() string {
	return s.String()
}

func (s *RetryVccResponse) SetHeaders(v map[string]*string) *RetryVccResponse {
	s.Headers = v
	return s
}

func (s *RetryVccResponse) SetStatusCode(v int32) *RetryVccResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryVccResponse) SetBody(v *RetryVccResponseBody) *RetryVccResponse {
	s.Body = v
	return s
}

type UnAssignPrivateIpAddressRequest struct {
	// By default, popApi is not ignored and idempotent
	//
	// example:
	//
	// 141cccd6-dfbd-11ec-b8e8-0242ac110003
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// IP unique identifier
	//
	// This parameter is required.
	//
	// example:
	//
	// sip-xxxx
	IpName *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
	// Lingjun network interface controller ID
	//
	// This parameter is required.
	//
	// example:
	//
	// lni-bp18exxqa2rvfn45e5pz
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The private IP address of the instance.
	//
	// example:
	//
	// 10.209.75.242
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// Region
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Subnet
	//
	// This parameter is required.
	//
	// example:
	//
	// subnet-f3zfzmnc
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
}

func (s UnAssignPrivateIpAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s UnAssignPrivateIpAddressRequest) GoString() string {
	return s.String()
}

func (s *UnAssignPrivateIpAddressRequest) SetClientToken(v string) *UnAssignPrivateIpAddressRequest {
	s.ClientToken = &v
	return s
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Content *UnAssignPrivateIpAddressResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// You don\\"t have the permission to do this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// A88DFED5-24B7-5A3E-87DE-380BF06F3C90
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnAssignPrivateIpAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnAssignPrivateIpAddressResponseBody) GoString() string {
	return s.String()
}

func (s *UnAssignPrivateIpAddressResponseBody) SetAccessDeniedDetail(v string) *UnAssignPrivateIpAddressResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// IP unique identifier
	//
	// example:
	//
	// sip-xxxxx
	IpName *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
	// Lingjun network interface controller ID
	//
	// example:
	//
	// lni-bp164jwjpdq4lnsy83s5
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnAssignPrivateIpAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type UnAssociateVpdCidrBlockRequest struct {
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The additional CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.0.0/16
	SecondaryCidrBlock *string `json:"SecondaryCidrBlock,omitempty" xml:"SecondaryCidrBlock,omitempty"`
	// The ID of the Lingjun CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpd-aof7dat1
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
}

func (s UnAssociateVpdCidrBlockRequest) String() string {
	return tea.Prettify(s)
}

func (s UnAssociateVpdCidrBlockRequest) GoString() string {
	return s.String()
}

func (s *UnAssociateVpdCidrBlockRequest) SetRegionId(v string) *UnAssociateVpdCidrBlockRequest {
	s.RegionId = &v
	return s
}

func (s *UnAssociateVpdCidrBlockRequest) SetSecondaryCidrBlock(v string) *UnAssociateVpdCidrBlockRequest {
	s.SecondaryCidrBlock = &v
	return s
}

func (s *UnAssociateVpdCidrBlockRequest) SetVpdId(v string) *UnAssociateVpdCidrBlockRequest {
	s.VpdId = &v
	return s
}

type UnAssociateVpdCidrBlockResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Content *UnAssociateVpdCidrBlockResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// 9C50C9CD-E799-54DA-BA7A-1FAF3DF80857
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnAssociateVpdCidrBlockResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnAssociateVpdCidrBlockResponseBody) GoString() string {
	return s.String()
}

func (s *UnAssociateVpdCidrBlockResponseBody) SetAccessDeniedDetail(v string) *UnAssociateVpdCidrBlockResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UnAssociateVpdCidrBlockResponseBody) SetCode(v int32) *UnAssociateVpdCidrBlockResponseBody {
	s.Code = &v
	return s
}

func (s *UnAssociateVpdCidrBlockResponseBody) SetContent(v *UnAssociateVpdCidrBlockResponseBodyContent) *UnAssociateVpdCidrBlockResponseBody {
	s.Content = v
	return s
}

func (s *UnAssociateVpdCidrBlockResponseBody) SetMessage(v string) *UnAssociateVpdCidrBlockResponseBody {
	s.Message = &v
	return s
}

func (s *UnAssociateVpdCidrBlockResponseBody) SetRequestId(v string) *UnAssociateVpdCidrBlockResponseBody {
	s.RequestId = &v
	return s
}

type UnAssociateVpdCidrBlockResponseBodyContent struct {
	// The ID of the Lingjun CIDR block.
	//
	// example:
	//
	// vpd-ze3na0wf
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
}

func (s UnAssociateVpdCidrBlockResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s UnAssociateVpdCidrBlockResponseBodyContent) GoString() string {
	return s.String()
}

func (s *UnAssociateVpdCidrBlockResponseBodyContent) SetVpdId(v string) *UnAssociateVpdCidrBlockResponseBodyContent {
	s.VpdId = &v
	return s
}

type UnAssociateVpdCidrBlockResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnAssociateVpdCidrBlockResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnAssociateVpdCidrBlockResponse) String() string {
	return tea.Prettify(s)
}

func (s UnAssociateVpdCidrBlockResponse) GoString() string {
	return s.String()
}

func (s *UnAssociateVpdCidrBlockResponse) SetHeaders(v map[string]*string) *UnAssociateVpdCidrBlockResponse {
	s.Headers = v
	return s
}

func (s *UnAssociateVpdCidrBlockResponse) SetStatusCode(v int32) *UnAssociateVpdCidrBlockResponse {
	s.StatusCode = &v
	return s
}

func (s *UnAssociateVpdCidrBlockResponse) SetBody(v *UnAssociateVpdCidrBlockResponseBody) *UnAssociateVpdCidrBlockResponse {
	s.Body = v
	return s
}

type UnassignLeniPrivateIpAddressRequest struct {
	// The idempotent identifier.
	//
	// example:
	//
	// 967e77a2-b61d-11ec-a147-0242c0a80504
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Lingjun Elastic Network Interface ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// leni-1234****
	ElasticNetworkInterfaceId *string `json:"ElasticNetworkInterfaceId,omitempty" xml:"ElasticNetworkInterfaceId,omitempty"`
	// Lingjun Elastic Network Interface secondary private IP unique identifier.
	//
	// This parameter is required.
	//
	// example:
	//
	// sip-8ylg****
	IpName *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UnassignLeniPrivateIpAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s UnassignLeniPrivateIpAddressRequest) GoString() string {
	return s.String()
}

func (s *UnassignLeniPrivateIpAddressRequest) SetClientToken(v string) *UnassignLeniPrivateIpAddressRequest {
	s.ClientToken = &v
	return s
}

func (s *UnassignLeniPrivateIpAddressRequest) SetElasticNetworkInterfaceId(v string) *UnassignLeniPrivateIpAddressRequest {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *UnassignLeniPrivateIpAddressRequest) SetIpName(v string) *UnassignLeniPrivateIpAddressRequest {
	s.IpName = &v
	return s
}

func (s *UnassignLeniPrivateIpAddressRequest) SetRegionId(v string) *UnassignLeniPrivateIpAddressRequest {
	s.RegionId = &v
	return s
}

type UnassignLeniPrivateIpAddressResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	//
	// example:
	//
	// {}
	Content *UnassignLeniPrivateIpAddressResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AC8C713A-A9F4-5984-A5E1-76496DF35153
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnassignLeniPrivateIpAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnassignLeniPrivateIpAddressResponseBody) GoString() string {
	return s.String()
}

func (s *UnassignLeniPrivateIpAddressResponseBody) SetAccessDeniedDetail(v string) *UnassignLeniPrivateIpAddressResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UnassignLeniPrivateIpAddressResponseBody) SetCode(v int32) *UnassignLeniPrivateIpAddressResponseBody {
	s.Code = &v
	return s
}

func (s *UnassignLeniPrivateIpAddressResponseBody) SetContent(v *UnassignLeniPrivateIpAddressResponseBodyContent) *UnassignLeniPrivateIpAddressResponseBody {
	s.Content = v
	return s
}

func (s *UnassignLeniPrivateIpAddressResponseBody) SetMessage(v string) *UnassignLeniPrivateIpAddressResponseBody {
	s.Message = &v
	return s
}

func (s *UnassignLeniPrivateIpAddressResponseBody) SetRequestId(v string) *UnassignLeniPrivateIpAddressResponseBody {
	s.RequestId = &v
	return s
}

type UnassignLeniPrivateIpAddressResponseBodyContent struct {
	// Lingjun Elastic Network Interface ID.
	//
	// example:
	//
	// leni-1234****
	ElasticNetworkInterfaceId *string `json:"ElasticNetworkInterfaceId,omitempty" xml:"ElasticNetworkInterfaceId,omitempty"`
	// Lingjun Elastic Network Interface secondary private IP unique identifier.
	//
	// example:
	//
	// sip-dqvs****
	IpName *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
}

func (s UnassignLeniPrivateIpAddressResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s UnassignLeniPrivateIpAddressResponseBodyContent) GoString() string {
	return s.String()
}

func (s *UnassignLeniPrivateIpAddressResponseBodyContent) SetElasticNetworkInterfaceId(v string) *UnassignLeniPrivateIpAddressResponseBodyContent {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *UnassignLeniPrivateIpAddressResponseBodyContent) SetIpName(v string) *UnassignLeniPrivateIpAddressResponseBodyContent {
	s.IpName = &v
	return s
}

type UnassignLeniPrivateIpAddressResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnassignLeniPrivateIpAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnassignLeniPrivateIpAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s UnassignLeniPrivateIpAddressResponse) GoString() string {
	return s.String()
}

func (s *UnassignLeniPrivateIpAddressResponse) SetHeaders(v map[string]*string) *UnassignLeniPrivateIpAddressResponse {
	s.Headers = v
	return s
}

func (s *UnassignLeniPrivateIpAddressResponse) SetStatusCode(v int32) *UnassignLeniPrivateIpAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *UnassignLeniPrivateIpAddressResponse) SetBody(v *UnassignLeniPrivateIpAddressResponseBody) *UnassignLeniPrivateIpAddressResponse {
	s.Body = v
	return s
}

type UpdateElasticNetworkInterfaceRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 3fd79d62-ab1d-11ec-9a53-0242ac110004
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the variable.
	//
	// example:
	//
	// LHICDOSEExternaluserinquiryP
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Lingjun Elastic Network Interface ID
	//
	// This parameter is required.
	//
	// example:
	//
	// leni-1234****
	ElasticNetworkInterfaceId *string `json:"ElasticNetworkInterfaceId,omitempty" xml:"ElasticNetworkInterfaceId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-wz9fj2s3o21nw2****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s UpdateElasticNetworkInterfaceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateElasticNetworkInterfaceRequest) GoString() string {
	return s.String()
}

func (s *UpdateElasticNetworkInterfaceRequest) SetClientToken(v string) *UpdateElasticNetworkInterfaceRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateElasticNetworkInterfaceRequest) SetDescription(v string) *UpdateElasticNetworkInterfaceRequest {
	s.Description = &v
	return s
}

func (s *UpdateElasticNetworkInterfaceRequest) SetElasticNetworkInterfaceId(v string) *UpdateElasticNetworkInterfaceRequest {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *UpdateElasticNetworkInterfaceRequest) SetRegionId(v string) *UpdateElasticNetworkInterfaceRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateElasticNetworkInterfaceRequest) SetSecurityGroupId(v string) *UpdateElasticNetworkInterfaceRequest {
	s.SecurityGroupId = &v
	return s
}

type UpdateElasticNetworkInterfaceResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Content *UpdateElasticNetworkInterfaceResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The return message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// AC8C713A-A9F4-5984-A5E1-76496DF35153
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateElasticNetworkInterfaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateElasticNetworkInterfaceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateElasticNetworkInterfaceResponseBody) SetAccessDeniedDetail(v string) *UpdateElasticNetworkInterfaceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateElasticNetworkInterfaceResponseBody) SetCode(v int32) *UpdateElasticNetworkInterfaceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateElasticNetworkInterfaceResponseBody) SetContent(v *UpdateElasticNetworkInterfaceResponseBodyContent) *UpdateElasticNetworkInterfaceResponseBody {
	s.Content = v
	return s
}

func (s *UpdateElasticNetworkInterfaceResponseBody) SetMessage(v string) *UpdateElasticNetworkInterfaceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateElasticNetworkInterfaceResponseBody) SetRequestId(v string) *UpdateElasticNetworkInterfaceResponseBody {
	s.RequestId = &v
	return s
}

type UpdateElasticNetworkInterfaceResponseBodyContent struct {
	// Lingjun Elastic Network Interface ID
	//
	// example:
	//
	// leni-1234****
	ElasticNetworkInterfaceId *string `json:"ElasticNetworkInterfaceId,omitempty" xml:"ElasticNetworkInterfaceId,omitempty"`
	// Lingjun Node ID
	//
	// example:
	//
	// e01-cn-lbj3aej****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s UpdateElasticNetworkInterfaceResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s UpdateElasticNetworkInterfaceResponseBodyContent) GoString() string {
	return s.String()
}

func (s *UpdateElasticNetworkInterfaceResponseBodyContent) SetElasticNetworkInterfaceId(v string) *UpdateElasticNetworkInterfaceResponseBodyContent {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *UpdateElasticNetworkInterfaceResponseBodyContent) SetNodeId(v string) *UpdateElasticNetworkInterfaceResponseBodyContent {
	s.NodeId = &v
	return s
}

type UpdateElasticNetworkInterfaceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateElasticNetworkInterfaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateElasticNetworkInterfaceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateElasticNetworkInterfaceResponse) GoString() string {
	return s.String()
}

func (s *UpdateElasticNetworkInterfaceResponse) SetHeaders(v map[string]*string) *UpdateElasticNetworkInterfaceResponse {
	s.Headers = v
	return s
}

func (s *UpdateElasticNetworkInterfaceResponse) SetStatusCode(v int32) *UpdateElasticNetworkInterfaceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateElasticNetworkInterfaceResponse) SetBody(v *UpdateElasticNetworkInterfaceResponseBody) *UpdateElasticNetworkInterfaceResponse {
	s.Body = v
	return s
}

type UpdateErRequest struct {
	// The description of the document.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Lingjun HUB Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// Parameter
	//
	// example:
	//
	// er-wulanchabu-main
	ErName *string `json:"ErName,omitempty" xml:"ErName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// {}
	Content map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// 3D9D6E7B-365B-5200-BFA6-9B79E269058C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateErResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateErResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateErResponseBody) SetAccessDeniedDetail(v string) *UpdateErResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateErResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The connection ID of the Lingjun HUB network instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// er-attachment-i1ioibyf
	ErAttachmentId *string `json:"ErAttachmentId,omitempty" xml:"ErAttachmentId,omitempty"`
	// Lingjun HUB Network Instance Connection Name
	//
	// example:
	//
	// er-attachment-wulanchabu-main
	ErAttachmentName *string `json:"ErAttachmentName,omitempty" xml:"ErAttachmentName,omitempty"`
	// Lingjun HUB ID
	//
	// This parameter is required.
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	//
	// example:
	//
	// {}
	Content map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// 7F9082CC-3D94-560F-A575-8E8EF6CE2CB8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateErAttachmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateErAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateErAttachmentResponseBody) SetAccessDeniedDetail(v string) *UpdateErAttachmentResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateErAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The description of the document.
	//
	// example:
	//
	// test-example
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Lingjun HUB ID
	//
	// This parameter is required.
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// routing policy ID
	//
	// This parameter is required.
	//
	// example:
	//
	// er-rmap-uwglhzom
	ErRouteMapId *string `json:"ErRouteMapId,omitempty" xml:"ErRouteMapId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	//
	// example:
	//
	// {}
	Content map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BDBCC783-84CA-5733-8EEA-645C88B9009C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateErRouteMapResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateErRouteMapResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateErRouteMapResponseBody) SetAccessDeniedDetail(v string) *UpdateErRouteMapResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateErRouteMapResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type UpdateLeniPrivateIpAddressRequest struct {
	// The description of the ECS instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Lingjun Elastic Network Interface ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// leni-1234****
	ElasticNetworkInterfaceId *string `json:"ElasticNetworkInterfaceId,omitempty" xml:"ElasticNetworkInterfaceId,omitempty"`
	// Lingjun Elastic Network Interface secondary private IP unique identifier.
	//
	// This parameter is required.
	//
	// example:
	//
	// sip-8ylg****
	IpName *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateLeniPrivateIpAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLeniPrivateIpAddressRequest) GoString() string {
	return s.String()
}

func (s *UpdateLeniPrivateIpAddressRequest) SetDescription(v string) *UpdateLeniPrivateIpAddressRequest {
	s.Description = &v
	return s
}

func (s *UpdateLeniPrivateIpAddressRequest) SetElasticNetworkInterfaceId(v string) *UpdateLeniPrivateIpAddressRequest {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *UpdateLeniPrivateIpAddressRequest) SetIpName(v string) *UpdateLeniPrivateIpAddressRequest {
	s.IpName = &v
	return s
}

func (s *UpdateLeniPrivateIpAddressRequest) SetRegionId(v string) *UpdateLeniPrivateIpAddressRequest {
	s.RegionId = &v
	return s
}

type UpdateLeniPrivateIpAddressResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Content *UpdateLeniPrivateIpAddressResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A88DFED5-24B7-5A3E-87DE-380BF06F3C90
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLeniPrivateIpAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLeniPrivateIpAddressResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLeniPrivateIpAddressResponseBody) SetAccessDeniedDetail(v string) *UpdateLeniPrivateIpAddressResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateLeniPrivateIpAddressResponseBody) SetCode(v int32) *UpdateLeniPrivateIpAddressResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateLeniPrivateIpAddressResponseBody) SetContent(v *UpdateLeniPrivateIpAddressResponseBodyContent) *UpdateLeniPrivateIpAddressResponseBody {
	s.Content = v
	return s
}

func (s *UpdateLeniPrivateIpAddressResponseBody) SetMessage(v string) *UpdateLeniPrivateIpAddressResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateLeniPrivateIpAddressResponseBody) SetRequestId(v string) *UpdateLeniPrivateIpAddressResponseBody {
	s.RequestId = &v
	return s
}

type UpdateLeniPrivateIpAddressResponseBodyContent struct {
	// Lingjun Elastic Network Interface ID.
	//
	// example:
	//
	// leni-1234****
	ElasticNetworkInterfaceId *string `json:"ElasticNetworkInterfaceId,omitempty" xml:"ElasticNetworkInterfaceId,omitempty"`
	// Lingjun Elastic Network Interface secondary private IP unique identifier.
	//
	// example:
	//
	// sip-8ylg****
	IpName *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
}

func (s UpdateLeniPrivateIpAddressResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s UpdateLeniPrivateIpAddressResponseBodyContent) GoString() string {
	return s.String()
}

func (s *UpdateLeniPrivateIpAddressResponseBodyContent) SetElasticNetworkInterfaceId(v string) *UpdateLeniPrivateIpAddressResponseBodyContent {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *UpdateLeniPrivateIpAddressResponseBodyContent) SetIpName(v string) *UpdateLeniPrivateIpAddressResponseBodyContent {
	s.IpName = &v
	return s
}

type UpdateLeniPrivateIpAddressResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLeniPrivateIpAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLeniPrivateIpAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLeniPrivateIpAddressResponse) GoString() string {
	return s.String()
}

func (s *UpdateLeniPrivateIpAddressResponse) SetHeaders(v map[string]*string) *UpdateLeniPrivateIpAddressResponse {
	s.Headers = v
	return s
}

func (s *UpdateLeniPrivateIpAddressResponse) SetStatusCode(v int32) *UpdateLeniPrivateIpAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLeniPrivateIpAddressResponse) SetBody(v *UpdateLeniPrivateIpAddressResponseBody) *UpdateLeniPrivateIpAddressResponse {
	s.Body = v
	return s
}

type UpdateSubnetRequest struct {
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The subnet instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// subnet-f3zfzmnc
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	// The new name for the subnet instance.
	//
	// example:
	//
	// subnet-1
	SubnetName *string `json:"SubnetName,omitempty" xml:"SubnetName,omitempty"`
	// The ID of the VPD to which the subnet belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpd-aof7dat1
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// The zone ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response content.
	Content *UpdateSubnetResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The message that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3D9D6E7B-365B-5200-BFA6-9B79E269058C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSubnetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubnetResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSubnetResponseBody) SetAccessDeniedDetail(v string) *UpdateSubnetResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// The subnet instance ID.
	//
	// example:
	//
	// subnet-yuvn29bn
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSubnetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The peak bandwidth of the Lingjun connection instance. Unit: Mbit/s. Valid values: 1000 to 400000
	//
	// example:
	//
	// 1000
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The ID of the order placed on the instance.
	//
	// example:
	//
	// 20006627643
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Lingjun connection instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vcc-cn-zvp2w222001
	VccId *string `json:"VccId,omitempty" xml:"VccId,omitempty"`
	// The name of the Lingjun connection instance.
	//
	// example:
	//
	// vcc-heyuan-backup
	VccName *string `json:"VccName,omitempty" xml:"VccName,omitempty"`
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Content *UpdateVccResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F906C4D3-7444-58E2-9819-E3D8563571A3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateVccResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateVccResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVccResponseBody) SetAccessDeniedDetail(v string) *UpdateVccResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// The ID of the Lingjun connection instance.
	//
	// example:
	//
	// vcc-cn-2r42v22cn03
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVccResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the VPD instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpd-omqutbff
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// The name of the VPD instance.
	//
	// example:
	//
	// vpd-lingjun
	VpdName *string `json:"VpdName,omitempty" xml:"VpdName,omitempty"`
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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Content *UpdateVpdResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AC8C713A-A9F4-5984-A5E1-76496DF35153
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateVpdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateVpdResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVpdResponseBody) SetAccessDeniedDetail(v string) *UpdateVpdResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// The ID of the VPD instance.
	//
	// example:
	//
	// vpd-lg4dppgi
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVpdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

// Summary:
//
// Apply for a secondary private IP address for the current Lingjun Elastic Network Interface. You can automatically assign a secondary private IP address.
//
// Description:
//
// Apply for a secondary private IP address for the specified Lingjun Elastic Network Interface.
//
// 	- If the PrivateIp field is empty, a secondary private IP address is automatically assigned and the unique identifier of the IP address is returned.
//
// 	- You can use the GetLeniPrivateIpAddress or ListLeniPrivateIpAddresses interface to check whether the secondary private IP address is assigned.
//
// @param request - AssignLeniPrivateIpAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssignLeniPrivateIpAddressResponse
func (client *Client) AssignLeniPrivateIpAddressWithOptions(request *AssignLeniPrivateIpAddressRequest, runtime *util.RuntimeOptions) (_result *AssignLeniPrivateIpAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticNetworkInterfaceId)) {
		body["ElasticNetworkInterfaceId"] = request.ElasticNetworkInterfaceId
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateIpAddress)) {
		body["PrivateIpAddress"] = request.PrivateIpAddress
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AssignLeniPrivateIpAddress"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &AssignLeniPrivateIpAddressResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &AssignLeniPrivateIpAddressResponse{}
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
// Apply for a secondary private IP address for the current Lingjun Elastic Network Interface. You can automatically assign a secondary private IP address.
//
// Description:
//
// Apply for a secondary private IP address for the specified Lingjun Elastic Network Interface.
//
// 	- If the PrivateIp field is empty, a secondary private IP address is automatically assigned and the unique identifier of the IP address is returned.
//
// 	- You can use the GetLeniPrivateIpAddress or ListLeniPrivateIpAddresses interface to check whether the secondary private IP address is assigned.
//
// @param request - AssignLeniPrivateIpAddressRequest
//
// @return AssignLeniPrivateIpAddressResponse
func (client *Client) AssignLeniPrivateIpAddress(request *AssignLeniPrivateIpAddressRequest) (_result *AssignLeniPrivateIpAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssignLeniPrivateIpAddressResponse{}
	_body, _err := client.AssignLeniPrivateIpAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Applies for a private secondary IP address for the current LNI. You can also call this operation to assign a secondary MAC address to the current LNI.
//
// Description:
//
// >  Apply for secondary private IP addresses
//
// 	- By default, each network interface controller can apply for three secondary private IP addresses. If the quota is exceeded, contact the administrator.
//
// 	- The secondary private IP address is allocated from the Lingjun subnet to which the current network interface controller belongs. The first address and the last two addresses belong to reserved addresses and do not participate in the allocation.
//
// @param request - AssignPrivateIpAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssignPrivateIpAddressResponse
func (client *Client) AssignPrivateIpAddressWithOptions(request *AssignPrivateIpAddressRequest, runtime *util.RuntimeOptions) (_result *AssignPrivateIpAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssignMac)) {
		body["AssignMac"] = request.AssignMac
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &AssignPrivateIpAddressResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &AssignPrivateIpAddressResponse{}
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
// Applies for a private secondary IP address for the current LNI. You can also call this operation to assign a secondary MAC address to the current LNI.
//
// Description:
//
// >  Apply for secondary private IP addresses
//
// 	- By default, each network interface controller can apply for three secondary private IP addresses. If the quota is exceeded, contact the administrator.
//
// 	- The secondary private IP address is allocated from the Lingjun subnet to which the current network interface controller belongs. The first address and the last two addresses belong to reserved addresses and do not participate in the allocation.
//
// @param request - AssignPrivateIpAddressRequest
//
// @return AssignPrivateIpAddressResponse
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

// Summary:
//
// When the VPD primary network segment address is not enough to allocate, you can choose to create an additional network segment as the additional network segment of the VPD primary network segment.
//
// Description:
//
// >  **Add a CIDR block**
//
// 	- The CIDR block cannot start with 0. The subnet mask must be 8 to 28 bits in length.
//
// 	- The secondary IPv4 CIDR block must not overlap with the primary IPv4 CIDR block of the Lingjun CIDR block and the added secondary IPv4 CIDR block.
//
// 	- You cannot use 100.64.0.0/10, 224.0.0.0/4, 127.0.0.0/8, or 169.254.0.0/16 as the CIDR block of Lingjun. Example: In the Lingjun CIDR block whose primary IPv4 CIDR block is 192.168.0.0/16, you cannot add the following CIDR blocks as additional IPv4 CIDR blocks. The CIDR block that is in the same range as 192.168.0.0/16. A CIDR block that is larger than 192.168.0.0/16. Example: 192.168.0.0/8. A CIDR block that is smaller than 192.168.0.0/16. Example: 192.168.0.0/24.
//
// 	- By default, each tenant can create three additional CIDR blocks in each region.
//
// @param request - AssociateVpdCidrBlockRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateVpdCidrBlockResponse
func (client *Client) AssociateVpdCidrBlockWithOptions(request *AssociateVpdCidrBlockRequest, runtime *util.RuntimeOptions) (_result *AssociateVpdCidrBlockResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecondaryCidrBlock)) {
		body["SecondaryCidrBlock"] = request.SecondaryCidrBlock
	}

	if !tea.BoolValue(util.IsUnset(request.VpdId)) {
		body["VpdId"] = request.VpdId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AssociateVpdCidrBlock"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &AssociateVpdCidrBlockResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &AssociateVpdCidrBlockResponse{}
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
// When the VPD primary network segment address is not enough to allocate, you can choose to create an additional network segment as the additional network segment of the VPD primary network segment.
//
// Description:
//
// >  **Add a CIDR block**
//
// 	- The CIDR block cannot start with 0. The subnet mask must be 8 to 28 bits in length.
//
// 	- The secondary IPv4 CIDR block must not overlap with the primary IPv4 CIDR block of the Lingjun CIDR block and the added secondary IPv4 CIDR block.
//
// 	- You cannot use 100.64.0.0/10, 224.0.0.0/4, 127.0.0.0/8, or 169.254.0.0/16 as the CIDR block of Lingjun. Example: In the Lingjun CIDR block whose primary IPv4 CIDR block is 192.168.0.0/16, you cannot add the following CIDR blocks as additional IPv4 CIDR blocks. The CIDR block that is in the same range as 192.168.0.0/16. A CIDR block that is larger than 192.168.0.0/16. Example: 192.168.0.0/8. A CIDR block that is smaller than 192.168.0.0/16. Example: 192.168.0.0/24.
//
// 	- By default, each tenant can create three additional CIDR blocks in each region.
//
// @param request - AssociateVpdCidrBlockRequest
//
// @return AssociateVpdCidrBlockResponse
func (client *Client) AssociateVpdCidrBlock(request *AssociateVpdCidrBlockRequest) (_result *AssociateVpdCidrBlockResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssociateVpdCidrBlockResponse{}
	_body, _err := client.AssociateVpdCidrBlockWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lingjun ENI is bound to NC.
//
// Description:
//
// This interface is an asynchronous interface. You need to use the query interface to wait for the Lingjun Elastic Network Interface to reach the available state.
//
// @param request - AttachElasticNetworkInterfaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachElasticNetworkInterfaceResponse
func (client *Client) AttachElasticNetworkInterfaceWithOptions(request *AttachElasticNetworkInterfaceRequest, runtime *util.RuntimeOptions) (_result *AttachElasticNetworkInterfaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ElasticNetworkInterfaceId)) {
		body["ElasticNetworkInterfaceId"] = request.ElasticNetworkInterfaceId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		body["NodeId"] = request.NodeId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachElasticNetworkInterface"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &AttachElasticNetworkInterfaceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &AttachElasticNetworkInterfaceResponse{}
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
// Lingjun ENI is bound to NC.
//
// Description:
//
// This interface is an asynchronous interface. You need to use the query interface to wait for the Lingjun Elastic Network Interface to reach the available state.
//
// @param request - AttachElasticNetworkInterfaceRequest
//
// @return AttachElasticNetworkInterfaceResponse
func (client *Client) AttachElasticNetworkInterface(request *AttachElasticNetworkInterfaceRequest) (_result *AttachElasticNetworkInterfaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachElasticNetworkInterfaceResponse{}
	_body, _err := client.AttachElasticNetworkInterfaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an LENI.
//
// @param request - CreateElasticNetworkInterfaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateElasticNetworkInterfaceResponse
func (client *Client) CreateElasticNetworkInterfaceWithOptions(request *CreateElasticNetworkInterfaceRequest, runtime *util.RuntimeOptions) (_result *CreateElasticNetworkInterfaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EnableJumboFrame)) {
		body["EnableJumboFrame"] = request.EnableJumboFrame
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		body["NodeId"] = request.NodeId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		body["SecurityGroupId"] = request.SecurityGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		body["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		body["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateElasticNetworkInterface"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateElasticNetworkInterfaceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateElasticNetworkInterfaceResponse{}
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
// Creates an LENI.
//
// @param request - CreateElasticNetworkInterfaceRequest
//
// @return CreateElasticNetworkInterfaceResponse
func (client *Client) CreateElasticNetworkInterface(request *CreateElasticNetworkInterfaceRequest) (_result *CreateElasticNetworkInterfaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateElasticNetworkInterfaceResponse{}
	_body, _err := client.CreateElasticNetworkInterfaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create a Lingjun HUB.
//
// Description:
//
// When you call this operation to create a Lingjun HUB, note that:
//
// 	- Make sure that you have sufficient Lingjun HUB quota.
//
// 	- This interface is an asynchronous interface. After this interface is called, the system will return the ID of a Lingjun HUB. At this time, the Lingjun HUB instance may not be created yet, and the system background creation task is still in progress. You can call the ListErs or GetEr operation to query the status of the Lingjun HUB.
//
//     	- If the status of the Lingjun HUB is Executing, it indicates that it is being created.
//
//     	- If the status of the Lingjun HUB is Available, the creation is successful.
//
// @param request - CreateErRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateErResponse
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

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateErResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateErResponse{}
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
// Create a Lingjun HUB.
//
// Description:
//
// When you call this operation to create a Lingjun HUB, note that:
//
// 	- Make sure that you have sufficient Lingjun HUB quota.
//
// 	- This interface is an asynchronous interface. After this interface is called, the system will return the ID of a Lingjun HUB. At this time, the Lingjun HUB instance may not be created yet, and the system background creation task is still in progress. You can call the ListErs or GetEr operation to query the status of the Lingjun HUB.
//
//     	- If the status of the Lingjun HUB is Executing, it indicates that it is being created.
//
//     	- If the status of the Lingjun HUB is Available, the creation is successful.
//
// @param request - CreateErRequest
//
// @return CreateErResponse
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

// Summary:
//
// Create a network instance connection.
//
// Description:
//
// When you call this operation to create a network instance connection, note that:
//
// 	- Make sure that you have created a Lingjun HUB instance.
//
// 	- Make sure that you have sufficient quota for network instance connections.
//
// 	- This operation is an asynchronous operation. After you call this operation, the system returns the ID of the network instance connection. In this case, the network instance connection may not be created yet, and the system is still creating the network instance in the background. You can query the connection status of a network instance by ListErAttachments or GetErAttachment:
//
//     	- If the connection status of the network instance is Executing, the network instance is being created.
//
//     	- If the connection status of the network instance is Available, the network instance is created.
//
// @param request - CreateErAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateErAttachmentResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateErAttachmentResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateErAttachmentResponse{}
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
// Create a network instance connection.
//
// Description:
//
// When you call this operation to create a network instance connection, note that:
//
// 	- Make sure that you have created a Lingjun HUB instance.
//
// 	- Make sure that you have sufficient quota for network instance connections.
//
// 	- This operation is an asynchronous operation. After you call this operation, the system returns the ID of the network instance connection. In this case, the network instance connection may not be created yet, and the system is still creating the network instance in the background. You can query the connection status of a network instance by ListErAttachments or GetErAttachment:
//
//     	- If the connection status of the network instance is Executing, the network instance is being created.
//
//     	- If the connection status of the network instance is Available, the network instance is created.
//
// @param request - CreateErAttachmentRequest
//
// @return CreateErAttachmentResponse
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

// Summary:
//
// Users can use this API to create routing policy by specifying the network instance connection under Lingjun HUB.
//
// Description:
//
// When you call this operation to create a routing policy, note that:
//
// 	- Make sure that you have created a Lingjun HUB instance.
//
// 	- Make sure that you have created a network instance connection.
//
// 	- This operation is an asynchronous operation. After you call this operation, the system returns the ID of the routing policy. In this case, the routing policy instance may not be created yet, and the system background creation task is still in progress. You can use ListErRouteMaps or GetErRouteMap to query the status of a routing policy.
//
//     	- If the status of the routing policy is Execute, the system is creating the instance.
//
//     	- If the status of the routing policy is Available, the creation is successful.
//
// @param request - CreateErRouteMapRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateErRouteMapResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateErRouteMapResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateErRouteMapResponse{}
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
// Users can use this API to create routing policy by specifying the network instance connection under Lingjun HUB.
//
// Description:
//
// When you call this operation to create a routing policy, note that:
//
// 	- Make sure that you have created a Lingjun HUB instance.
//
// 	- Make sure that you have created a network instance connection.
//
// 	- This operation is an asynchronous operation. After you call this operation, the system returns the ID of the routing policy. In this case, the routing policy instance may not be created yet, and the system background creation task is still in progress. You can use ListErRouteMaps or GetErRouteMap to query the status of a routing policy.
//
//     	- If the status of the routing policy is Execute, the system is creating the instance.
//
//     	- If the status of the routing policy is Available, the creation is successful.
//
// @param request - CreateErRouteMapRequest
//
// @return CreateErRouteMapResponse
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

// Summary:
//
// Users can use this API to create a Lingjun subnet under the Lingjun network segment.
//
// Description:
//
// When you call this operation to create a Lingjun subnet, note that:
//
// 	- You have created a Lingjun CIDR block.
//
// 	- Only one network segment can be specified for a Lingjun subnet.
//
// 	- The network segment cannot be modified after the Lingjun subnet is created.
//
// 	- Make sure that you have sufficient Lingjun subnet quota.
//
// 	- This interface is an asynchronous interface. After calling this interface, the system will return the ID of a Lingjun subnet. At this time, the Lingjun network segment may not be created yet, and the system background creation task is still in progress. You can call the ListSubnets or GetSubnet operation to query the status of the CIDR block of Lingjun.
//
//     	- If the status of the Lingjun subnet is Executed, it indicates that it is being created.
//
//     	- If the status of the Lingjun subnet is Available, the creation is successful.
//
// @param request - CreateSubnetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSubnetResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateSubnetResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateSubnetResponse{}
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
// Users can use this API to create a Lingjun subnet under the Lingjun network segment.
//
// Description:
//
// When you call this operation to create a Lingjun subnet, note that:
//
// 	- You have created a Lingjun CIDR block.
//
// 	- Only one network segment can be specified for a Lingjun subnet.
//
// 	- The network segment cannot be modified after the Lingjun subnet is created.
//
// 	- Make sure that you have sufficient Lingjun subnet quota.
//
// 	- This interface is an asynchronous interface. After calling this interface, the system will return the ID of a Lingjun subnet. At this time, the Lingjun network segment may not be created yet, and the system background creation task is still in progress. You can call the ListSubnets or GetSubnet operation to query the status of the CIDR block of Lingjun.
//
//     	- If the status of the Lingjun subnet is Executed, it indicates that it is being created.
//
//     	- If the status of the Lingjun subnet is Available, the creation is successful.
//
// @param request - CreateSubnetRequest
//
// @return CreateSubnetResponse
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

// Summary:
//
// You can create a Lingjun connection to connect Lingjun network environment and Alibaba Cloud network environment.
//
// Description:
//
// When you call this operation to create a Lingjun connection, note that:
//
// 	- When you specify the vccId parameter, the system will configure the purchased Lingjun connection for you. When the default vccId parameter is set, the system will automatically place an order and configure the Lingjun connection for you.
//
// 	- Make sure that you have called the InitializeVcc operation to grant permissions.
//
// 	- This interface is an asynchronous interface. After this interface is called, the system will return the Lingjun connection ID, but the Lingjun connection instance may not be created yet, and the system background creation task is still in progress. You can call the ListVccs or GetVcc operation to query the status of the Lingjun connection.
//
//     	- If the status of the Lingjun connection is Executed, the Lingjun connection is being created.
//
//     	- If the status of the Lingjun connection is Available, the Lingjun connection is created.
//
// @param request - CreateVccRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVccResponse
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

	if !tea.BoolValue(util.IsUnset(request.BgpAsn)) {
		body["BgpAsn"] = request.BgpAsn
	}

	if !tea.BoolValue(util.IsUnset(request.BgpCidr)) {
		body["BgpCidr"] = request.BgpCidr
	}

	if !tea.BoolValue(util.IsUnset(request.CenId)) {
		body["CenId"] = request.CenId
	}

	if !tea.BoolValue(util.IsUnset(request.CenOwnerId)) {
		body["CenOwnerId"] = request.CenOwnerId
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateVccResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateVccResponse{}
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
// You can create a Lingjun connection to connect Lingjun network environment and Alibaba Cloud network environment.
//
// Description:
//
// When you call this operation to create a Lingjun connection, note that:
//
// 	- When you specify the vccId parameter, the system will configure the purchased Lingjun connection for you. When the default vccId parameter is set, the system will automatically place an order and configure the Lingjun connection for you.
//
// 	- Make sure that you have called the InitializeVcc operation to grant permissions.
//
// 	- This interface is an asynchronous interface. After this interface is called, the system will return the Lingjun connection ID, but the Lingjun connection instance may not be created yet, and the system background creation task is still in progress. You can call the ListVccs or GetVcc operation to query the status of the Lingjun connection.
//
//     	- If the status of the Lingjun connection is Executed, the Lingjun connection is being created.
//
//     	- If the status of the Lingjun connection is Available, the Lingjun connection is created.
//
// @param request - CreateVccRequest
//
// @return CreateVccResponse
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

// Summary:
//
// Users can use this API to connect Lingjun instance to the Lingjun HUB instance of the target account. After authorization, the target account can be associated with your Lingjun connection by using the authorized Lingjun HUB instance.
//
// Description:
//
// When you call this operation to create cross-account authorization for Lingjun HUB, note that:
//
// 	- Make sure that the Alibaba Cloud ID and Lingjun HUB instance that you want to authorize are correct.
//
// 	- If you authorize the account of the other party, the account of the other party can load your local network instance to its Lingjun HUB, and the other party\\"s network will be connected to your network. Please proceed with caution.
//
// @param request - CreateVccGrantRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVccGrantRuleResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateVccGrantRuleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateVccGrantRuleResponse{}
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
// Users can use this API to connect Lingjun instance to the Lingjun HUB instance of the target account. After authorization, the target account can be associated with your Lingjun connection by using the authorized Lingjun HUB instance.
//
// Description:
//
// When you call this operation to create cross-account authorization for Lingjun HUB, note that:
//
// 	- Make sure that the Alibaba Cloud ID and Lingjun HUB instance that you want to authorize are correct.
//
// 	- If you authorize the account of the other party, the account of the other party can load your local network instance to its Lingjun HUB, and the other party\\"s network will be connected to your network. Please proceed with caution.
//
// @param request - CreateVccGrantRuleRequest
//
// @return CreateVccGrantRuleResponse
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

// Summary:
//
// Create a Lingjun connection route entry.
//
// Description:
//
// When you call this operation to create a VBR route entry, take note of the following items:
//
// 	- After you call this operation, static route entries and BGP network announcements are created on the VBR to which the Lingjun connection belongs.
//
// 	- This operation is an asynchronous operation. After you call this operation, the VBR static route entry may not be created yet, and the system still creates the static route entry in the background. You can query the status of VBR static route entries by ListVccRouteEntries or GetVccRouteEntry:
//
//     	- If the VBR static route entry is in the Executing state, it indicates that it is being created.
//
//     	- If the status of the VBR static route entry is Available, the VBR is created.
//
// @param request - CreateVccRouteEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVccRouteEntryResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateVccRouteEntryResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateVccRouteEntryResponse{}
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
// Create a Lingjun connection route entry.
//
// Description:
//
// When you call this operation to create a VBR route entry, take note of the following items:
//
// 	- After you call this operation, static route entries and BGP network announcements are created on the VBR to which the Lingjun connection belongs.
//
// 	- This operation is an asynchronous operation. After you call this operation, the VBR static route entry may not be created yet, and the system still creates the static route entry in the background. You can query the status of VBR static route entries by ListVccRouteEntries or GetVccRouteEntry:
//
//     	- If the VBR static route entry is in the Executing state, it indicates that it is being created.
//
//     	- If the status of the VBR static route entry is Available, the VBR is created.
//
// @param request - CreateVccRouteEntryRequest
//
// @return CreateVccRouteEntryResponse
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

// Summary:
//
// Create a private Lingjun CIDR block. This CIDR block has an independent network environment.
//
// Description:
//
// When you call this operation to create a CIDR block for Lingjun, take note of the following:
//
// 	- A Lingjun network segment can specify an additional network segment in addition to a main network segment.
//
// 	- After the Lingjun network segment is created, the network segment cannot be modified.
//
// 	- Make sure that you have a sufficient quota of Lingjun CIDR blocks.
//
// 	- This interface is an asynchronous interface. After calling this interface, the system will return the ID of a Lingjun network segment. At this time, the Lingjun network segment may not be created yet, and the system background creation task is still in progress. You can call the ListVpds or GetVpd operation to query the status of the CIDR block of Lingjun.
//
//     	- If the status of the Lingjun CIDR block is Executed, the CIDR block is being created.
//
//     	- If the status of the Lingjun CIDR block is Available, the creation is successful.
//
// @param request - CreateVpdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVpdResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateVpdResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateVpdResponse{}
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
// Create a private Lingjun CIDR block. This CIDR block has an independent network environment.
//
// Description:
//
// When you call this operation to create a CIDR block for Lingjun, take note of the following:
//
// 	- A Lingjun network segment can specify an additional network segment in addition to a main network segment.
//
// 	- After the Lingjun network segment is created, the network segment cannot be modified.
//
// 	- Make sure that you have a sufficient quota of Lingjun CIDR blocks.
//
// 	- This interface is an asynchronous interface. After calling this interface, the system will return the ID of a Lingjun network segment. At this time, the Lingjun network segment may not be created yet, and the system background creation task is still in progress. You can call the ListVpds or GetVpd operation to query the status of the CIDR block of Lingjun.
//
//     	- If the status of the Lingjun CIDR block is Executed, the CIDR block is being created.
//
//     	- If the status of the Lingjun CIDR block is Available, the creation is successful.
//
// @param request - CreateVpdRequest
//
// @return CreateVpdResponse
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

// Summary:
//
// Users can use this API to authorize Lingjun HUB instances of the target account. After authorization, the target account can be associated with your Lingjun CIDR block by using the authorized Lingjun HUB instance.
//
// Description:
//
// When you call this operation to create cross-account authorization for Lingjun HUB, note that:
//
// 	- Make sure that the Alibaba Cloud ID and Lingjun HUB instance that you want to authorize are correct.
//
// 	- If you authorize the account of the other party, the account of the other party can load your local network instance to its Lingjun HUB, and the other party\\"s network will be connected to your network. Please proceed with caution.
//
// @param request - CreateVpdGrantRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVpdGrantRuleResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateVpdGrantRuleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateVpdGrantRuleResponse{}
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
// Users can use this API to authorize Lingjun HUB instances of the target account. After authorization, the target account can be associated with your Lingjun CIDR block by using the authorized Lingjun HUB instance.
//
// Description:
//
// When you call this operation to create cross-account authorization for Lingjun HUB, note that:
//
// 	- Make sure that the Alibaba Cloud ID and Lingjun HUB instance that you want to authorize are correct.
//
// 	- If you authorize the account of the other party, the account of the other party can load your local network instance to its Lingjun HUB, and the other party\\"s network will be connected to your network. Please proceed with caution.
//
// @param request - CreateVpdGrantRuleRequest
//
// @return CreateVpdGrantRuleResponse
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

// Summary:
//
// Delete Lingjun Elastic Network Interface. After deletion, all relevant data will be lost and cannot be recovered. Please operate with caution.
//
// @param request - DeleteElasticNetworkInterfaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteElasticNetworkInterfaceResponse
func (client *Client) DeleteElasticNetworkInterfaceWithOptions(request *DeleteElasticNetworkInterfaceRequest, runtime *util.RuntimeOptions) (_result *DeleteElasticNetworkInterfaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticNetworkInterfaceId)) {
		body["ElasticNetworkInterfaceId"] = request.ElasticNetworkInterfaceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteElasticNetworkInterface"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteElasticNetworkInterfaceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteElasticNetworkInterfaceResponse{}
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
// Delete Lingjun Elastic Network Interface. After deletion, all relevant data will be lost and cannot be recovered. Please operate with caution.
//
// @param request - DeleteElasticNetworkInterfaceRequest
//
// @return DeleteElasticNetworkInterfaceResponse
func (client *Client) DeleteElasticNetworkInterface(request *DeleteElasticNetworkInterfaceRequest) (_result *DeleteElasticNetworkInterfaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteElasticNetworkInterfaceResponse{}
	_body, _err := client.DeleteElasticNetworkInterfaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// After you delete a Lingjun HUB instance, the related data is lost and cannot be recovered.
//
// Description:
//
// When you call this operation to delete the Lingjun HUB, note that:
//
// 	- Before you delete the instance, make sure that no network instance is connected to the Lingjun HUB instance.
//
// 	- After deletion, all related data is lost and cannot be recovered. Exercise caution when performing this operation.
//
// 	- This interface is an asynchronous interface. After this interface is called, the Lingjun HUB instance may not be deleted, and the system background deletion task is still in progress. You can call the ListErs or GetEr operation to query the deletion status of the Lingjun HUB.
//
//     	- If the status of the Lingjun HUB is Deleting, the Lingjun HUB instance is being deleted.
//
//     	- If no Lingjun HUB instance is recorded, the Lingjun HUB instance has been deleted.
//
// @param request - DeleteErRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteErResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteErResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteErResponse{}
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
// After you delete a Lingjun HUB instance, the related data is lost and cannot be recovered.
//
// Description:
//
// When you call this operation to delete the Lingjun HUB, note that:
//
// 	- Before you delete the instance, make sure that no network instance is connected to the Lingjun HUB instance.
//
// 	- After deletion, all related data is lost and cannot be recovered. Exercise caution when performing this operation.
//
// 	- This interface is an asynchronous interface. After this interface is called, the Lingjun HUB instance may not be deleted, and the system background deletion task is still in progress. You can call the ListErs or GetEr operation to query the deletion status of the Lingjun HUB.
//
//     	- If the status of the Lingjun HUB is Deleting, the Lingjun HUB instance is being deleted.
//
//     	- If no Lingjun HUB instance is recorded, the Lingjun HUB instance has been deleted.
//
// @param request - DeleteErRequest
//
// @return DeleteErResponse
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

// Summary:
//
// If you delete a network instance that is connected to an instance, the related data is lost and cannot be recovered.
//
// Description:
//
// When you call this operation to delete a network instance connection, take note of the following:
//
// 	- Before you delete the instance, make sure that no routing policy exists under the network instance connection instance.
//
// 	- After deletion, all related data is lost and cannot be recovered. Exercise caution when performing this operation.
//
// 	- This operation is an asynchronous operation. After you call this operation, the network instance that is connected to the instance may not be deleted. The system still deletes the instance in the background. You can call the ListErAttachments or GetErAttachment to query the deletion status of network instance connections:
//
//     	- If the status of the network instance connection is Deleting, the network instance connection is being deleted.
//
//     	- If there is no connection record for the network instance, the connection to the network instance has been deleted.
//
// @param request - DeleteErAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteErAttachmentResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteErAttachmentResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteErAttachmentResponse{}
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
// If you delete a network instance that is connected to an instance, the related data is lost and cannot be recovered.
//
// Description:
//
// When you call this operation to delete a network instance connection, take note of the following:
//
// 	- Before you delete the instance, make sure that no routing policy exists under the network instance connection instance.
//
// 	- After deletion, all related data is lost and cannot be recovered. Exercise caution when performing this operation.
//
// 	- This operation is an asynchronous operation. After you call this operation, the network instance that is connected to the instance may not be deleted. The system still deletes the instance in the background. You can call the ListErAttachments or GetErAttachment to query the deletion status of network instance connections:
//
//     	- If the status of the network instance connection is Deleting, the network instance connection is being deleted.
//
//     	- If there is no connection record for the network instance, the connection to the network instance has been deleted.
//
// @param request - DeleteErAttachmentRequest
//
// @return DeleteErAttachmentResponse
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

// Summary:
//
// If you delete a routing policy instance, the related data is lost and cannot be recovered.
//
// Description:
//
// When you call this operation to delete a routing policy, note that:
//
// 	- After deletion, all related data is lost and cannot be recovered. Exercise caution when performing this operation.
//
// 	- This interface is an asynchronous interface. After this interface is called, the routing policy instance may not be deleted yet, and the system background deletion task is still in progress. You can call the ListErRouteMaps or GetErRouteMap operation to query the deletion status of a routing policy.
//
//     	- If the routing policy is in the Deleting state, the routing policy instance is being deleted.
//
//     	- If no routing policy instance is recorded, the routing policy instance has been deleted.
//
// @param request - DeleteErRouteMapRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteErRouteMapResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteErRouteMapResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteErRouteMapResponse{}
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
// If you delete a routing policy instance, the related data is lost and cannot be recovered.
//
// Description:
//
// When you call this operation to delete a routing policy, note that:
//
// 	- After deletion, all related data is lost and cannot be recovered. Exercise caution when performing this operation.
//
// 	- This interface is an asynchronous interface. After this interface is called, the routing policy instance may not be deleted yet, and the system background deletion task is still in progress. You can call the ListErRouteMaps or GetErRouteMap operation to query the deletion status of a routing policy.
//
//     	- If the routing policy is in the Deleting state, the routing policy instance is being deleted.
//
//     	- If no routing policy instance is recorded, the routing policy instance has been deleted.
//
// @param request - DeleteErRouteMapRequest
//
// @return DeleteErRouteMapResponse
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

// Summary:
//
// If you delete a Lingjun subnet instance, the related data is lost and cannot be recovered.
//
// Description:
//
// When you call this operation to delete a Lingjun subnet, note that:
//
// 	- After deletion, all related data is lost and cannot be recovered. Exercise caution when performing this operation.
//
// 	- This interface is an asynchronous interface. After this interface is called, the Lingjun subnet instance may not be deleted, and the system background deletion task is still in progress. You can call the ListSubnets or GetSubnet operation to query the deletion status of the subnet.
//
//     	- If the status of the Lingjun subnet is Deleting, the Lingjun subnet instance is being deleted.
//
//     	- If there is no record of the Lingjun subnet instance, the Lingjun subnet instance has been deleted.
//
// @param request - DeleteSubnetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSubnetResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteSubnetResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteSubnetResponse{}
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
// If you delete a Lingjun subnet instance, the related data is lost and cannot be recovered.
//
// Description:
//
// When you call this operation to delete a Lingjun subnet, note that:
//
// 	- After deletion, all related data is lost and cannot be recovered. Exercise caution when performing this operation.
//
// 	- This interface is an asynchronous interface. After this interface is called, the Lingjun subnet instance may not be deleted, and the system background deletion task is still in progress. You can call the ListSubnets or GetSubnet operation to query the deletion status of the subnet.
//
//     	- If the status of the Lingjun subnet is Deleting, the Lingjun subnet instance is being deleted.
//
//     	- If there is no record of the Lingjun subnet instance, the Lingjun subnet instance has been deleted.
//
// @param request - DeleteSubnetRequest
//
// @return DeleteSubnetResponse
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

// Summary:
//
// If you delete a Lingjun HUB cross-account authorization that is connected to Lingjun, the related data is lost and cannot be recovered.
//
// @param request - DeleteVccGrantRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVccGrantRuleResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteVccGrantRuleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteVccGrantRuleResponse{}
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
// If you delete a Lingjun HUB cross-account authorization that is connected to Lingjun, the related data is lost and cannot be recovered.
//
// @param request - DeleteVccGrantRuleRequest
//
// @return DeleteVccGrantRuleResponse
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

// Summary:
//
// Delete a Lingjun connection route entry.
//
// Description:
//
// When you call this operation to delete a VBR static route entry, note that:
//
// 	- After deletion, all related data is lost and cannot be recovered. Exercise caution when performing this operation.
//
// 	- This operation is an asynchronous operation. After you call this operation, the VBR static route entries may not be deleted. The system still deletes the VBR static route entries in the background. You can call the ListVccRouteEntries or GetVccRouteEntry to query the deletion status of VBR static route entries:
//
//     	- If the VBR static route entry is in the Deleting state, the VBR static route entry is being deleted.
//
//     	- If no VBR static route entry instance is recorded, the VBR static route entry instance has been deleted.
//
// @param request - DeleteVccRouteEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVccRouteEntryResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteVccRouteEntryResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteVccRouteEntryResponse{}
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
// Delete a Lingjun connection route entry.
//
// Description:
//
// When you call this operation to delete a VBR static route entry, note that:
//
// 	- After deletion, all related data is lost and cannot be recovered. Exercise caution when performing this operation.
//
// 	- This operation is an asynchronous operation. After you call this operation, the VBR static route entries may not be deleted. The system still deletes the VBR static route entries in the background. You can call the ListVccRouteEntries or GetVccRouteEntry to query the deletion status of VBR static route entries:
//
//     	- If the VBR static route entry is in the Deleting state, the VBR static route entry is being deleted.
//
//     	- If no VBR static route entry instance is recorded, the VBR static route entry instance has been deleted.
//
// @param request - DeleteVccRouteEntryRequest
//
// @return DeleteVccRouteEntryResponse
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

// Summary:
//
// After you delete a Lingjun CIDR block, the related data is lost and cannot be recovered.
//
// Description:
//
// When you call this operation to delete a Lingjun CIDR block, take note of the following items:
//
// 	- After deletion, all related data is lost and cannot be recovered. Exercise caution when performing this operation.
//
// 	- Before deleting, make sure that all Lingjun subnet instances under the Lingjun CIDR block have been deleted.
//
// 	- This interface is an asynchronous interface. After this interface is called, the Lingjun network segment instance may not be deleted, and the system background deletion task is still in progress. You can call the ListVpds or GetVpd operation to query the deletion status of the CIDR block.
//
//     	- If the status of the Lingjun CIDR block is Deleting, the Lingjun CIDR block is being deleted.
//
//     	- If there is no record of the Lingjun CIDR block instance, the Lingjun CIDR block instance has been deleted.
//
// @param request - DeleteVpdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVpdResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteVpdResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteVpdResponse{}
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
// After you delete a Lingjun CIDR block, the related data is lost and cannot be recovered.
//
// Description:
//
// When you call this operation to delete a Lingjun CIDR block, take note of the following items:
//
// 	- After deletion, all related data is lost and cannot be recovered. Exercise caution when performing this operation.
//
// 	- Before deleting, make sure that all Lingjun subnet instances under the Lingjun CIDR block have been deleted.
//
// 	- This interface is an asynchronous interface. After this interface is called, the Lingjun network segment instance may not be deleted, and the system background deletion task is still in progress. You can call the ListVpds or GetVpd operation to query the deletion status of the CIDR block.
//
//     	- If the status of the Lingjun CIDR block is Deleting, the Lingjun CIDR block is being deleted.
//
//     	- If there is no record of the Lingjun CIDR block instance, the Lingjun CIDR block instance has been deleted.
//
// @param request - DeleteVpdRequest
//
// @return DeleteVpdResponse
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

// Summary:
//
// Delete the Lingjun HUB cross-account authorization for a Lingjun CIDR block. After the deletion, the related data is lost and cannot be recovered.
//
// @param request - DeleteVpdGrantRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVpdGrantRuleResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteVpdGrantRuleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteVpdGrantRuleResponse{}
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
// Delete the Lingjun HUB cross-account authorization for a Lingjun CIDR block. After the deletion, the related data is lost and cannot be recovered.
//
// @param request - DeleteVpdGrantRuleRequest
//
// @return DeleteVpdGrantRuleResponse
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

// Summary:
//
// Query whether the user has the SLR role-AliyunServiceRoleForEfloVcc required for Lingjun connection.
//
// Description:
//
// You can call this operation to query whether a user account has a **AliyunServiceRoleForEfloVcc*	- role.
//
// >  If you do not have a **AliyunServiceRoleForEfloVcc*	- role, you need to use the initializeVcc interface to complete authorization, otherwise users will not be able to use Lingjun to connect to the product.
//
// @param request - DescribeSlrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSlrResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeSlrResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeSlrResponse{}
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
// Query whether the user has the SLR role-AliyunServiceRoleForEfloVcc required for Lingjun connection.
//
// Description:
//
// You can call this operation to query whether a user account has a **AliyunServiceRoleForEfloVcc*	- role.
//
// >  If you do not have a **AliyunServiceRoleForEfloVcc*	- role, you need to use the initializeVcc interface to complete authorization, otherwise users will not be able to use Lingjun to connect to the product.
//
// @param request - DescribeSlrRequest
//
// @return DescribeSlrResponse
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

// Summary:
//
// Unbind Lingjun ENI from NC.
//
// Description:
//
// This interface is an asynchronous interface, and you need to use the query interface to wait for the Lingjun Elastic Network Interface to reach the unbound state.
//
// @param request - DetachElasticNetworkInterfaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachElasticNetworkInterfaceResponse
func (client *Client) DetachElasticNetworkInterfaceWithOptions(request *DetachElasticNetworkInterfaceRequest, runtime *util.RuntimeOptions) (_result *DetachElasticNetworkInterfaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ElasticNetworkInterfaceId)) {
		body["ElasticNetworkInterfaceId"] = request.ElasticNetworkInterfaceId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		body["NodeId"] = request.NodeId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachElasticNetworkInterface"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DetachElasticNetworkInterfaceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DetachElasticNetworkInterfaceResponse{}
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
// Unbind Lingjun ENI from NC.
//
// Description:
//
// This interface is an asynchronous interface, and you need to use the query interface to wait for the Lingjun Elastic Network Interface to reach the unbound state.
//
// @param request - DetachElasticNetworkInterfaceRequest
//
// @return DetachElasticNetworkInterfaceResponse
func (client *Client) DetachElasticNetworkInterface(request *DetachElasticNetworkInterfaceRequest) (_result *DetachElasticNetworkInterfaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachElasticNetworkInterfaceResponse{}
	_body, _err := client.DetachElasticNetworkInterfaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Users can use this API to query the destination CIDR block of the source policy instance when creating a routing strategy.
//
// @param request - GetDestinationCidrBlockRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDestinationCidrBlockResponse
func (client *Client) GetDestinationCidrBlockWithOptions(request *GetDestinationCidrBlockRequest, runtime *util.RuntimeOptions) (_result *GetDestinationCidrBlockResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
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
		Action:      tea.String("GetDestinationCidrBlock"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetDestinationCidrBlockResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetDestinationCidrBlockResponse{}
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
// Users can use this API to query the destination CIDR block of the source policy instance when creating a routing strategy.
//
// @param request - GetDestinationCidrBlockRequest
//
// @return GetDestinationCidrBlockResponse
func (client *Client) GetDestinationCidrBlock(request *GetDestinationCidrBlockRequest) (_result *GetDestinationCidrBlockResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDestinationCidrBlockResponse{}
	_body, _err := client.GetDestinationCidrBlockWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an LENI.
//
// @param request - GetElasticNetworkInterfaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetElasticNetworkInterfaceResponse
func (client *Client) GetElasticNetworkInterfaceWithOptions(request *GetElasticNetworkInterfaceRequest, runtime *util.RuntimeOptions) (_result *GetElasticNetworkInterfaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ElasticNetworkInterfaceId)) {
		body["ElasticNetworkInterfaceId"] = request.ElasticNetworkInterfaceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetElasticNetworkInterface"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetElasticNetworkInterfaceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetElasticNetworkInterfaceResponse{}
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
// Queries the details of an LENI.
//
// @param request - GetElasticNetworkInterfaceRequest
//
// @return GetElasticNetworkInterfaceResponse
func (client *Client) GetElasticNetworkInterface(request *GetElasticNetworkInterfaceRequest) (_result *GetElasticNetworkInterfaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetElasticNetworkInterfaceResponse{}
	_body, _err := client.GetElasticNetworkInterfaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Lingjun HUB.
//
// @param request - GetErRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetErResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetErResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetErResponse{}
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
// Queries the Lingjun HUB.
//
// @param request - GetErRequest
//
// @return GetErResponse
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

// Summary:
//
// Queries network instance connections.
//
// @param request - GetErAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetErAttachmentResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetErAttachmentResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetErAttachmentResponse{}
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
// Queries network instance connections.
//
// @param request - GetErAttachmentRequest
//
// @return GetErAttachmentResponse
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

// Summary:
//
// Queries the details of Lingjun HUB route entries.
//
// @param request - GetErRouteEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetErRouteEntryResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetErRouteEntryResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetErRouteEntryResponse{}
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
// Queries the details of Lingjun HUB route entries.
//
// @param request - GetErRouteEntryRequest
//
// @return GetErRouteEntryResponse
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

// Summary:
//
// query lingjun hub routing policy details.
//
// @param request - GetErRouteMapRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetErRouteMapResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetErRouteMapResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetErRouteMapResponse{}
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
// query lingjun hub routing policy details.
//
// @param request - GetErRouteMapRequest
//
// @return GetErRouteMapResponse
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

// Summary:
//
// Query the physical topology information of Lingjun network interface controller and Lingjun nodes under VPD.
//
// @param request - GetFabricTopologyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFabricTopologyResponse
func (client *Client) GetFabricTopologyWithOptions(request *GetFabricTopologyRequest, runtime *util.RuntimeOptions) (_result *GetFabricTopologyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.LniIds)) {
		body["LniIds"] = request.LniIds
	}

	if !tea.BoolValue(util.IsUnset(request.NodeIds)) {
		body["NodeIds"] = request.NodeIds
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
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
		Action:      tea.String("GetFabricTopology"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetFabricTopologyResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetFabricTopologyResponse{}
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
// Query the physical topology information of Lingjun network interface controller and Lingjun nodes under VPD.
//
// @param request - GetFabricTopologyRequest
//
// @return GetFabricTopologyResponse
func (client *Client) GetFabricTopology(request *GetFabricTopologyRequest) (_result *GetFabricTopologyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFabricTopologyResponse{}
	_body, _err := client.GetFabricTopologyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the details of the secondary private IP address of a specified Lingjun Elastic Network Interface.
//
// @param request - GetLeniPrivateIpAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLeniPrivateIpAddressResponse
func (client *Client) GetLeniPrivateIpAddressWithOptions(request *GetLeniPrivateIpAddressRequest, runtime *util.RuntimeOptions) (_result *GetLeniPrivateIpAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ElasticNetworkInterfaceId)) {
		body["ElasticNetworkInterfaceId"] = request.ElasticNetworkInterfaceId
	}

	if !tea.BoolValue(util.IsUnset(request.IpName)) {
		body["IpName"] = request.IpName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLeniPrivateIpAddress"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetLeniPrivateIpAddressResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetLeniPrivateIpAddressResponse{}
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
// Obtains the details of the secondary private IP address of a specified Lingjun Elastic Network Interface.
//
// @param request - GetLeniPrivateIpAddressRequest
//
// @return GetLeniPrivateIpAddressResponse
func (client *Client) GetLeniPrivateIpAddress(request *GetLeniPrivateIpAddressRequest) (_result *GetLeniPrivateIpAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLeniPrivateIpAddressResponse{}
	_body, _err := client.GetLeniPrivateIpAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the details about the secondary private IP address.
//
// @param request - GetLniPrivateIpAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLniPrivateIpAddressResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetLniPrivateIpAddressResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetLniPrivateIpAddressResponse{}
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
// Obtains the details about the secondary private IP address.
//
// @param request - GetLniPrivateIpAddressRequest
//
// @return GetLniPrivateIpAddressResponse
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

// Summary:
//
// Queries information about an LNI.
//
// @param request - GetNetworkInterfaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNetworkInterfaceResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetNetworkInterfaceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetNetworkInterfaceResponse{}
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
// Queries information about an LNI.
//
// @param request - GetNetworkInterfaceRequest
//
// @return GetNetworkInterfaceResponse
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

// Summary:
//
// Queries the network information of a node.
//
// @param request - GetNodeInfoForPodRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNodeInfoForPodResponse
func (client *Client) GetNodeInfoForPodWithOptions(request *GetNodeInfoForPodRequest, runtime *util.RuntimeOptions) (_result *GetNodeInfoForPodResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		body["NodeId"] = request.NodeId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNodeInfoForPod"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetNodeInfoForPodResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetNodeInfoForPodResponse{}
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
// Queries the network information of a node.
//
// @param request - GetNodeInfoForPodRequest
//
// @return GetNodeInfoForPodResponse
func (client *Client) GetNodeInfoForPod(request *GetNodeInfoForPodRequest) (_result *GetNodeInfoForPodResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetNodeInfoForPodResponse{}
	_body, _err := client.GetNodeInfoForPodWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a Lingjun subnet, including the type, CIDR block, instance ID, instance status, and number of NCs.
//
// @param request - GetSubnetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSubnetResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetSubnetResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetSubnetResponse{}
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
// Queries the details of a Lingjun subnet, including the type, CIDR block, instance ID, instance status, and number of NCs.
//
// @param request - GetSubnetRequest
//
// @return GetSubnetResponse
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

// Summary:
//
// Queries the details of a Lingjun connection, including the specification, Express Connect circuit access port type, instance status, bandwidth, and BGP CIDR block.
//
// @param request - GetVccRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVccResponse
func (client *Client) GetVccWithOptions(request *GetVccRequest, runtime *util.RuntimeOptions) (_result *GetVccResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetVccResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetVccResponse{}
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
// Queries the details of a Lingjun connection, including the specification, Express Connect circuit access port type, instance status, bandwidth, and BGP CIDR block.
//
// @param request - GetVccRequest
//
// @return GetVccResponse
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

// Summary:
//
// Queries the details of cross-account resource authorization for a Lingjun connection, including the authorized tenant ID, Lingjun HUB instance ID, and network instance ID.
//
// @param request - GetVccGrantRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVccGrantRuleResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetVccGrantRuleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetVccGrantRuleResponse{}
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
// Queries the details of cross-account resource authorization for a Lingjun connection, including the authorized tenant ID, Lingjun HUB instance ID, and network instance ID.
//
// @param request - GetVccGrantRuleRequest
//
// @return GetVccGrantRuleResponse
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

// Summary:
//
// Queries route entries.
//
// @param request - GetVccRouteEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVccRouteEntryResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetVccRouteEntryResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetVccRouteEntryResponse{}
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
// Queries route entries.
//
// @param request - GetVccRouteEntryRequest
//
// @return GetVccRouteEntryResponse
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

// Summary:
//
// Queries the details of a Lingjun CIDR block, including the status of the Lingjun CIDR block, the CIDR block, the number of subnets and NCs.
//
// @param request - GetVpdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVpdResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetVpdResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetVpdResponse{}
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
// Queries the details of a Lingjun CIDR block, including the status of the Lingjun CIDR block, the CIDR block, the number of subnets and NCs.
//
// @param request - GetVpdRequest
//
// @return GetVpdResponse
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

// Summary:
//
// Queries the details of cross-account resource authorization for a Lingjun CIDR block, including the authorized tenant ID, Lingjun HUB instance ID, and network instance ID.
//
// @param request - GetVpdGrantRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVpdGrantRuleResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetVpdGrantRuleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetVpdGrantRuleResponse{}
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
// Queries the details of cross-account resource authorization for a Lingjun CIDR block, including the authorized tenant ID, Lingjun HUB instance ID, and network instance ID.
//
// @param request - GetVpdGrantRuleRequest
//
// @return GetVpdGrantRuleResponse
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

// Summary:
//
// Queries route entries.
//
// @param request - GetVpdRouteEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVpdRouteEntryResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetVpdRouteEntryResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetVpdRouteEntryResponse{}
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
// Queries route entries.
//
// @param request - GetVpdRouteEntryRequest
//
// @return GetVpdRouteEntryResponse
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

// Summary:
//
// Initialize the Lingjun connection and authorize Intelligent Computing Lingjun to create an SLR in your account.
//
// @param request - InitializeVccRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitializeVccResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &InitializeVccResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &InitializeVccResponse{}
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
// Initialize the Lingjun connection and authorize Intelligent Computing Lingjun to create an SLR in your account.
//
// @param request - InitializeVccRequest
//
// @return InitializeVccResponse
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

// Summary:
//
// Queries the LENIs that are associated with a Lingjun node.
//
// @param request - ListElasticNetworkInterfacesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListElasticNetworkInterfacesResponse
func (client *Client) ListElasticNetworkInterfacesWithOptions(request *ListElasticNetworkInterfacesRequest, runtime *util.RuntimeOptions) (_result *ListElasticNetworkInterfacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ElasticNetworkInterfaceId)) {
		body["ElasticNetworkInterfaceId"] = request.ElasticNetworkInterfaceId
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		body["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkType)) {
		body["NetworkType"] = request.NetworkType
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

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		body["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		body["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListElasticNetworkInterfaces"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListElasticNetworkInterfacesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListElasticNetworkInterfacesResponse{}
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
// Queries the LENIs that are associated with a Lingjun node.
//
// @param request - ListElasticNetworkInterfacesRequest
//
// @return ListElasticNetworkInterfacesResponse
func (client *Client) ListElasticNetworkInterfaces(request *ListElasticNetworkInterfacesRequest) (_result *ListElasticNetworkInterfacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListElasticNetworkInterfacesResponse{}
	_body, _err := client.ListElasticNetworkInterfacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries network instance connections.
//
// @param request - ListErAttachmentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListErAttachmentsResponse
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

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListErAttachmentsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListErAttachmentsResponse{}
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
// Queries network instance connections.
//
// @param request - ListErAttachmentsRequest
//
// @return ListErAttachmentsResponse
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

// Summary:
//
// Queries the route entries of the Lingjun HUB.
//
// @param request - ListErRouteEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListErRouteEntriesResponse
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

	if !tea.BoolValue(util.IsUnset(request.IgnoreDetailedRouteEntry)) {
		body["IgnoreDetailedRouteEntry"] = request.IgnoreDetailedRouteEntry
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

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListErRouteEntriesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListErRouteEntriesResponse{}
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
// Queries the route entries of the Lingjun HUB.
//
// @param request - ListErRouteEntriesRequest
//
// @return ListErRouteEntriesResponse
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

// Summary:
//
// Routing policies are queried.
//
// @param request - ListErRouteMapsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListErRouteMapsResponse
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

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListErRouteMapsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListErRouteMapsResponse{}
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
// Routing policies are queried.
//
// @param request - ListErRouteMapsRequest
//
// @return ListErRouteMapsResponse
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

// Summary:
//
// Queries the Lingjun HUB.
//
// @param request - ListErsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListErsResponse
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

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListErsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListErsResponse{}
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
// Queries the Lingjun HUB.
//
// @param request - ListErsRequest
//
// @return ListErsResponse
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

// Summary:
//
// Queries the GPU node list of a specified GPU node whose communication distance does not exceed the specified NCD.
//
// @param request - ListInstancesByNcdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesByNcdResponse
func (client *Client) ListInstancesByNcdWithOptions(request *ListInstancesByNcdRequest, runtime *util.RuntimeOptions) (_result *ListInstancesByNcdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		body["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxNcd)) {
		body["MaxNcd"] = request.MaxNcd
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstancesByNcd"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListInstancesByNcdResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListInstancesByNcdResponse{}
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
// Queries the GPU node list of a specified GPU node whose communication distance does not exceed the specified NCD.
//
// @param request - ListInstancesByNcdRequest
//
// @return ListInstancesByNcdResponse
func (client *Client) ListInstancesByNcd(request *ListInstancesByNcdRequest) (_result *ListInstancesByNcdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstancesByNcdResponse{}
	_body, _err := client.ListInstancesByNcdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of secondary private IP addresses of Lingjun Elastic Network Interface.
//
// @param request - ListLeniPrivateIpAddressesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLeniPrivateIpAddressesResponse
func (client *Client) ListLeniPrivateIpAddressesWithOptions(request *ListLeniPrivateIpAddressesRequest, runtime *util.RuntimeOptions) (_result *ListLeniPrivateIpAddressesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ElasticNetworkInterfaceId)) {
		body["ElasticNetworkInterfaceId"] = request.ElasticNetworkInterfaceId
	}

	if !tea.BoolValue(util.IsUnset(request.IpName)) {
		body["IpName"] = request.IpName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateIpAddress)) {
		body["PrivateIpAddress"] = request.PrivateIpAddress
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLeniPrivateIpAddresses"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListLeniPrivateIpAddressesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListLeniPrivateIpAddressesResponse{}
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
// Queries the list of secondary private IP addresses of Lingjun Elastic Network Interface.
//
// @param request - ListLeniPrivateIpAddressesRequest
//
// @return ListLeniPrivateIpAddressesResponse
func (client *Client) ListLeniPrivateIpAddresses(request *ListLeniPrivateIpAddressesRequest) (_result *ListLeniPrivateIpAddressesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListLeniPrivateIpAddressesResponse{}
	_body, _err := client.ListLeniPrivateIpAddressesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of secondary private IP addresses of Lingjun network interface controller.
//
// @param request - ListLniPrivateIpAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLniPrivateIpAddressResponse
func (client *Client) ListLniPrivateIpAddressWithOptions(request *ListLniPrivateIpAddressRequest, runtime *util.RuntimeOptions) (_result *ListLniPrivateIpAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListLniPrivateIpAddressResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListLniPrivateIpAddressResponse{}
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
// Queries the list of secondary private IP addresses of Lingjun network interface controller.
//
// @param request - ListLniPrivateIpAddressRequest
//
// @return ListLniPrivateIpAddressResponse
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

// Summary:
//
// Queries Lingjun network interfaces (LNIs).
//
// @param request - ListNetworkInterfacesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNetworkInterfacesResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListNetworkInterfacesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListNetworkInterfacesResponse{}
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
// Queries Lingjun network interfaces (LNIs).
//
// @param request - ListNetworkInterfacesRequest
//
// @return ListNetworkInterfacesResponse
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

// Summary:
//
// Queries node network information.
//
// @param request - ListNodeInfosForPodRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodeInfosForPodResponse
func (client *Client) ListNodeInfosForPodWithOptions(request *ListNodeInfosForPodRequest, runtime *util.RuntimeOptions) (_result *ListNodeInfosForPodResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		body["NodeId"] = request.NodeId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		body["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNodeInfosForPod"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListNodeInfosForPodResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListNodeInfosForPodResponse{}
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
// Queries node network information.
//
// @param request - ListNodeInfosForPodRequest
//
// @return ListNodeInfosForPodResponse
func (client *Client) ListNodeInfosForPod(request *ListNodeInfosForPodRequest) (_result *ListNodeInfosForPodResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNodeInfosForPodResponse{}
	_body, _err := client.ListNodeInfosForPodWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call this operation to query the details of one or more Lingjun subnets, including the Lingjun subnet type, network address segment, and instance ID of the Lingjun CIDR block.
//
// @param request - ListSubnetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSubnetsResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListSubnetsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListSubnetsResponse{}
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
// You can call this operation to query the details of one or more Lingjun subnets, including the Lingjun subnet type, network address segment, and instance ID of the Lingjun CIDR block.
//
// @param request - ListSubnetsRequest
//
// @return ListSubnetsResponse
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

// Summary:
//
// Queries the traffic rate of a Lingjun connection.
//
// @param request - ListVccFlowInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVccFlowInfosResponse
func (client *Client) ListVccFlowInfosWithOptions(request *ListVccFlowInfosRequest, runtime *util.RuntimeOptions) (_result *ListVccFlowInfosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Direction)) {
		body["Direction"] = request.Direction
	}

	if !tea.BoolValue(util.IsUnset(request.From)) {
		body["From"] = request.From
	}

	if !tea.BoolValue(util.IsUnset(request.MetricName)) {
		body["MetricName"] = request.MetricName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.To)) {
		body["To"] = request.To
	}

	if !tea.BoolValue(util.IsUnset(request.VccId)) {
		body["VccId"] = request.VccId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListVccFlowInfos"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListVccFlowInfosResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListVccFlowInfosResponse{}
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
// Queries the traffic rate of a Lingjun connection.
//
// @param request - ListVccFlowInfosRequest
//
// @return ListVccFlowInfosResponse
func (client *Client) ListVccFlowInfos(request *ListVccFlowInfosRequest) (_result *ListVccFlowInfosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVccFlowInfosResponse{}
	_body, _err := client.ListVccFlowInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a Lingjun connection authorization, including the authorized tenant ID, region, and Lingjun HUB instance information.
//
// @param request - ListVccGrantRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVccGrantRulesResponse
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

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListVccGrantRulesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListVccGrantRulesResponse{}
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
// Queries the details of a Lingjun connection authorization, including the authorized tenant ID, region, and Lingjun HUB instance information.
//
// @param request - ListVccGrantRulesRequest
//
// @return ListVccGrantRulesResponse
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

// Summary:
//
// Queries Lingjun connection route entries.
//
// @param request - ListVccRouteEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVccRouteEntriesResponse
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

	if !tea.BoolValue(util.IsUnset(request.IgnoreDetailedRouteEntry)) {
		body["IgnoreDetailedRouteEntry"] = request.IgnoreDetailedRouteEntry
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

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListVccRouteEntriesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListVccRouteEntriesResponse{}
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
// Queries Lingjun connection route entries.
//
// @param request - ListVccRouteEntriesRequest
//
// @return ListVccRouteEntriesResponse
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

// Summary:
//
// query the details of one or more lingjun connections, including the specification, Express Connect circuit access port type, instance status, bandwidth, and bgp network segment.
//
// @param request - ListVccsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVccsResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListVccsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListVccsResponse{}
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
// query the details of one or more lingjun connections, including the specification, Express Connect circuit access port type, instance status, bandwidth, and bgp network segment.
//
// @param request - ListVccsRequest
//
// @return ListVccsResponse
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

// Summary:
//
// Queries the details of one or more route entries in the CIDR block of Lingjun, including the route type, route entry status, destination CIDR block, and instance information of the next route entry.
//
// @param request - ListVpdGrantRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVpdGrantRulesResponse
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

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListVpdGrantRulesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListVpdGrantRulesResponse{}
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
// Queries the details of one or more route entries in the CIDR block of Lingjun, including the route type, route entry status, destination CIDR block, and instance information of the next route entry.
//
// @param request - ListVpdGrantRulesRequest
//
// @return ListVpdGrantRulesResponse
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

// Summary:
//
// Queries the route entries of the Lingjun CIDR block.
//
// @param request - ListVpdRouteEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVpdRouteEntriesResponse
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

	if !tea.BoolValue(util.IsUnset(request.IgnoreDetailedRouteEntry)) {
		body["IgnoreDetailedRouteEntry"] = request.IgnoreDetailedRouteEntry
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

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListVpdRouteEntriesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListVpdRouteEntriesResponse{}
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
// Queries the route entries of the Lingjun CIDR block.
//
// @param request - ListVpdRouteEntriesRequest
//
// @return ListVpdRouteEntriesResponse
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

// Summary:
//
// Queries the details of one or more Lingjun CIDR blocks, including the status of Lingjun CIDR blocks, Cidr addresses, service CIDR blocks, and Subnet.
//
// @param request - ListVpdsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVpdsResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListVpdsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListVpdsResponse{}
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
// Queries the details of one or more Lingjun CIDR blocks, including the status of Lingjun CIDR blocks, Cidr addresses, service CIDR blocks, and Subnet.
//
// @param request - ListVpdsRequest
//
// @return ListVpdsResponse
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

// Summary:
//
// Query the network communication distance (Network Communication Distance,NCD) between instances (Lingjun node, Lingjun network interface controller).
//
// @param request - QueryInstanceNcdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryInstanceNcdResponse
func (client *Client) QueryInstanceNcdWithOptions(request *QueryInstanceNcdRequest, runtime *util.RuntimeOptions) (_result *QueryInstanceNcdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId1)) {
		body["InstanceId1"] = request.InstanceId1
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId2)) {
		body["InstanceId2"] = request.InstanceId2
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		body["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryInstanceNcd"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &QueryInstanceNcdResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &QueryInstanceNcdResponse{}
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
// Query the network communication distance (Network Communication Distance,NCD) between instances (Lingjun node, Lingjun network interface controller).
//
// @param request - QueryInstanceNcdRequest
//
// @return QueryInstanceNcdResponse
func (client *Client) QueryInstanceNcd(request *QueryInstanceNcdRequest) (_result *QueryInstanceNcdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryInstanceNcdResponse{}
	_body, _err := client.QueryInstanceNcdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unsubscribe inactive Lingjun connection
//
// Description:
//
// Only unsubscribable for Lingjun connections in the prepayment status.
//
// @param request - RefundVccRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefundVccResponse
func (client *Client) RefundVccWithOptions(request *RefundVccRequest, runtime *util.RuntimeOptions) (_result *RefundVccResponse, _err error) {
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

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RefundVcc"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RefundVccResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RefundVccResponse{}
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
// Unsubscribe inactive Lingjun connection
//
// Description:
//
// Only unsubscribable for Lingjun connections in the prepayment status.
//
// @param request - RefundVccRequest
//
// @return RefundVccResponse
func (client *Client) RefundVcc(request *RefundVccRequest) (_result *RefundVccResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefundVccResponse{}
	_body, _err := client.RefundVccWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retry trying to create /delete a Lingjun connection.
//
// Description:
//
// This operation allows the user to retry the operation if the Lingjun connection creation and deletion processes fail. Only the Lingjun connection in the creation failure and deletion failure state can be retried
//
// @param request - RetryVccRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetryVccResponse
func (client *Client) RetryVccWithOptions(request *RetryVccRequest, runtime *util.RuntimeOptions) (_result *RetryVccResponse, _err error) {
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

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RetryVcc"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RetryVccResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RetryVccResponse{}
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
// Retry trying to create /delete a Lingjun connection.
//
// Description:
//
// This operation allows the user to retry the operation if the Lingjun connection creation and deletion processes fail. Only the Lingjun connection in the creation failure and deletion failure state can be retried
//
// @param request - RetryVccRequest
//
// @return RetryVccResponse
func (client *Client) RetryVcc(request *RetryVccRequest) (_result *RetryVccResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RetryVccResponse{}
	_body, _err := client.RetryVccWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an assigned secondary private IP address.
//
// @param request - UnAssignPrivateIpAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnAssignPrivateIpAddressResponse
func (client *Client) UnAssignPrivateIpAddressWithOptions(request *UnAssignPrivateIpAddressRequest, runtime *util.RuntimeOptions) (_result *UnAssignPrivateIpAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UnAssignPrivateIpAddressResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UnAssignPrivateIpAddressResponse{}
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
// Deletes an assigned secondary private IP address.
//
// @param request - UnAssignPrivateIpAddressRequest
//
// @return UnAssignPrivateIpAddressResponse
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

// Summary:
//
// This function can be used to delete the additional network segment of VPD.
//
// Description:
//
// *
//
// **Warning*	- If the attached CIDR block has Lingjun subnet resources, you must delete the dependent resources before you can delete the attached CIDR block.
//
// @param request - UnAssociateVpdCidrBlockRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnAssociateVpdCidrBlockResponse
func (client *Client) UnAssociateVpdCidrBlockWithOptions(request *UnAssociateVpdCidrBlockRequest, runtime *util.RuntimeOptions) (_result *UnAssociateVpdCidrBlockResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecondaryCidrBlock)) {
		body["SecondaryCidrBlock"] = request.SecondaryCidrBlock
	}

	if !tea.BoolValue(util.IsUnset(request.VpdId)) {
		body["VpdId"] = request.VpdId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UnAssociateVpdCidrBlock"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UnAssociateVpdCidrBlockResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UnAssociateVpdCidrBlockResponse{}
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
// This function can be used to delete the additional network segment of VPD.
//
// Description:
//
// *
//
// **Warning*	- If the attached CIDR block has Lingjun subnet resources, you must delete the dependent resources before you can delete the attached CIDR block.
//
// @param request - UnAssociateVpdCidrBlockRequest
//
// @return UnAssociateVpdCidrBlockResponse
func (client *Client) UnAssociateVpdCidrBlock(request *UnAssociateVpdCidrBlockRequest) (_result *UnAssociateVpdCidrBlockResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnAssociateVpdCidrBlockResponse{}
	_body, _err := client.UnAssociateVpdCidrBlockWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete the assigned secondary private IP address of Lingjun Elastic Network Interface.
//
// @param request - UnassignLeniPrivateIpAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnassignLeniPrivateIpAddressResponse
func (client *Client) UnassignLeniPrivateIpAddressWithOptions(request *UnassignLeniPrivateIpAddressRequest, runtime *util.RuntimeOptions) (_result *UnassignLeniPrivateIpAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticNetworkInterfaceId)) {
		body["ElasticNetworkInterfaceId"] = request.ElasticNetworkInterfaceId
	}

	if !tea.BoolValue(util.IsUnset(request.IpName)) {
		body["IpName"] = request.IpName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UnassignLeniPrivateIpAddress"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UnassignLeniPrivateIpAddressResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UnassignLeniPrivateIpAddressResponse{}
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
// Delete the assigned secondary private IP address of Lingjun Elastic Network Interface.
//
// @param request - UnassignLeniPrivateIpAddressRequest
//
// @return UnassignLeniPrivateIpAddressResponse
func (client *Client) UnassignLeniPrivateIpAddress(request *UnassignLeniPrivateIpAddressRequest) (_result *UnassignLeniPrivateIpAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnassignLeniPrivateIpAddressResponse{}
	_body, _err := client.UnassignLeniPrivateIpAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Update Lingjun Elastic Network Interface information.
//
// @param request - UpdateElasticNetworkInterfaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateElasticNetworkInterfaceResponse
func (client *Client) UpdateElasticNetworkInterfaceWithOptions(request *UpdateElasticNetworkInterfaceRequest, runtime *util.RuntimeOptions) (_result *UpdateElasticNetworkInterfaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticNetworkInterfaceId)) {
		body["ElasticNetworkInterfaceId"] = request.ElasticNetworkInterfaceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		body["SecurityGroupId"] = request.SecurityGroupId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateElasticNetworkInterface"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateElasticNetworkInterfaceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateElasticNetworkInterfaceResponse{}
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
// Update Lingjun Elastic Network Interface information.
//
// @param request - UpdateElasticNetworkInterfaceRequest
//
// @return UpdateElasticNetworkInterfaceResponse
func (client *Client) UpdateElasticNetworkInterface(request *UpdateElasticNetworkInterfaceRequest) (_result *UpdateElasticNetworkInterfaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateElasticNetworkInterfaceResponse{}
	_body, _err := client.UpdateElasticNetworkInterfaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updated Lingjun HUB.
//
// @param request - UpdateErRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateErResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateErResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateErResponse{}
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
// Updated Lingjun HUB.
//
// @param request - UpdateErRequest
//
// @return UpdateErResponse
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

// Summary:
//
// Updates a network instance connection.
//
// @param request - UpdateErAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateErAttachmentResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateErAttachmentResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateErAttachmentResponse{}
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
// Updates a network instance connection.
//
// @param request - UpdateErAttachmentRequest
//
// @return UpdateErAttachmentResponse
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

// Summary:
//
// Update some information about the routing policy, including the description and name of the routing policy.
//
// @param request - UpdateErRouteMapRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateErRouteMapResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateErRouteMapResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateErRouteMapResponse{}
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
// Update some information about the routing policy, including the description and name of the routing policy.
//
// @param request - UpdateErRouteMapRequest
//
// @return UpdateErRouteMapResponse
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

// Summary:
//
// Updated the description of the secondary private network assigned by the Lingjun Elastic Network Interface.
//
// @param request - UpdateLeniPrivateIpAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLeniPrivateIpAddressResponse
func (client *Client) UpdateLeniPrivateIpAddressWithOptions(request *UpdateLeniPrivateIpAddressRequest, runtime *util.RuntimeOptions) (_result *UpdateLeniPrivateIpAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticNetworkInterfaceId)) {
		body["ElasticNetworkInterfaceId"] = request.ElasticNetworkInterfaceId
	}

	if !tea.BoolValue(util.IsUnset(request.IpName)) {
		body["IpName"] = request.IpName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLeniPrivateIpAddress"),
		Version:     tea.String("2022-05-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateLeniPrivateIpAddressResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateLeniPrivateIpAddressResponse{}
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
// Updated the description of the secondary private network assigned by the Lingjun Elastic Network Interface.
//
// @param request - UpdateLeniPrivateIpAddressRequest
//
// @return UpdateLeniPrivateIpAddressResponse
func (client *Client) UpdateLeniPrivateIpAddress(request *UpdateLeniPrivateIpAddressRequest) (_result *UpdateLeniPrivateIpAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLeniPrivateIpAddressResponse{}
	_body, _err := client.UpdateLeniPrivateIpAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates some information about a specified subnet instance, including the name of the subnet instance.
//
// @param request - UpdateSubnetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSubnetResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateSubnetResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateSubnetResponse{}
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
// Updates some information about a specified subnet instance, including the name of the subnet instance.
//
// @param request - UpdateSubnetRequest
//
// @return UpdateSubnetResponse
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

// Summary:
//
// Updates the information about a Lingjun connection instance, including the peak bandwidth and name of the Lingjun connection instance.
//
// @param request - UpdateVccRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVccResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateVccResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateVccResponse{}
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
// Updates the information about a Lingjun connection instance, including the peak bandwidth and name of the Lingjun connection instance.
//
// @param request - UpdateVccRequest
//
// @return UpdateVccResponse
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

// Summary:
//
// Updates the information about the Lingjun CIDR block, including the name of the Lingjun CIDR block.
//
// @param request - UpdateVpdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVpdResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateVpdResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateVpdResponse{}
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
// Updates the information about the Lingjun CIDR block, including the name of the Lingjun CIDR block.
//
// @param request - UpdateVpdRequest
//
// @return UpdateVpdResponse
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
