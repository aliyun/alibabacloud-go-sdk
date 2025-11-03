// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomerGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsn(v string) *CreateCustomerGatewayRequest
	GetAsn() *string
	SetAuthKey(v string) *CreateCustomerGatewayRequest
	GetAuthKey() *string
	SetClientToken(v string) *CreateCustomerGatewayRequest
	GetClientToken() *string
	SetDescription(v string) *CreateCustomerGatewayRequest
	GetDescription() *string
	SetIpAddress(v string) *CreateCustomerGatewayRequest
	GetIpAddress() *string
	SetName(v string) *CreateCustomerGatewayRequest
	GetName() *string
	SetOwnerAccount(v string) *CreateCustomerGatewayRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateCustomerGatewayRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateCustomerGatewayRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateCustomerGatewayRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateCustomerGatewayRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateCustomerGatewayRequest
	GetResourceOwnerId() *int64
	SetTags(v []*CreateCustomerGatewayRequestTags) *CreateCustomerGatewayRequest
	GetTags() []*CreateCustomerGatewayRequestTags
}

type CreateCustomerGatewayRequest struct {
	// The autonomous system number (ASN) of the gateway device in your data center. This parameter is required If you want to use Border Gateway Protocol (BGP) for the IPsec-VPN connection. Valid values: 1 to 4294967295. 45104 is not supported.
	//
	// **Asn*	- is a 4-byte number. You can enter it in two segments and separate the first 16 bits from the following 16 bits with a period (.). Enter the number in each segment in decimal format.
	//
	// For example, if you enter 123.456, the ASN is 8061384. The ASN is calculated by using the following formula: 123 × 65536 + 456 = 8061384.
	//
	// > - We recommend that you use a private ASN to establish BGP connections to Alibaba Cloud. For information about the range of private ASNs, see the relevant documentation.
	//
	// > - 45104 is a unique identifier assigned by IANA to Alibaba Cloud. It is used to identify Alibaba Cloud during route selection and data transmission over the Internet.
	//
	// example:
	//
	// 65530
	Asn *string `json:"Asn,omitempty" xml:"Asn,omitempty"`
	// The authentication key of the BGP routing protocol for the gateway device in the data center.
	//
	// The key must be 1 to 64 characters in length. It can contain only ASCII characters and cannot contain spaces or question marks (?).
	//
	// example:
	//
	// AuthKey****
	AuthKey *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the customer gateway.
	//
	// The description must be 1 to 100 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// desctest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The static IP address of the gateway device in the data center.
	//
	// 	- If you want to create a public IPsec-VPN connection, enter a public IP address.
	//
	// 	- If you want to create a private IPsec-VPN connection, enter a private IP address.
	//
	// You cannot use the following IP addresses. Otherwise, a IPsec-VPN connection cannot be established:
	//
	// 	- 100.64.0.0~100.127.255.255
	//
	// 	- 127.0.0.0~127.255.255.255
	//
	// 	- 169.254.0.0~169.254.255.255
	//
	// 	- 224.0.0.0~239.255.255.255
	//
	// 	- 255.0.0.0~255.255.255.255
	//
	// This parameter is required.
	//
	// example:
	//
	// 101.12.XX.XX
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The name of the customer gateway.
	//
	// The name must be 1 to 100 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// nametest
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the customer gateway.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the customer gateway belongs.
	//
	// - You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation to query the resource group list.
	//
	// - If you do not specify a resource group, the customer gateway will belong to the default resource group after being created.
	//
	// example:
	//
	// rg-aek2qo2h4jy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tag value.
	//
	// The tag value can be an empty string and cannot exceed 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// Each tag key corresponds to one tag value. You can specify up to 20 tag values in each call.
	Tags []*CreateCustomerGatewayRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s CreateCustomerGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomerGatewayRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomerGatewayRequest) GetAsn() *string {
	return s.Asn
}

func (s *CreateCustomerGatewayRequest) GetAuthKey() *string {
	return s.AuthKey
}

func (s *CreateCustomerGatewayRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateCustomerGatewayRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateCustomerGatewayRequest) GetIpAddress() *string {
	return s.IpAddress
}

func (s *CreateCustomerGatewayRequest) GetName() *string {
	return s.Name
}

func (s *CreateCustomerGatewayRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateCustomerGatewayRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateCustomerGatewayRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCustomerGatewayRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateCustomerGatewayRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateCustomerGatewayRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateCustomerGatewayRequest) GetTags() []*CreateCustomerGatewayRequestTags {
	return s.Tags
}

func (s *CreateCustomerGatewayRequest) SetAsn(v string) *CreateCustomerGatewayRequest {
	s.Asn = &v
	return s
}

func (s *CreateCustomerGatewayRequest) SetAuthKey(v string) *CreateCustomerGatewayRequest {
	s.AuthKey = &v
	return s
}

func (s *CreateCustomerGatewayRequest) SetClientToken(v string) *CreateCustomerGatewayRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCustomerGatewayRequest) SetDescription(v string) *CreateCustomerGatewayRequest {
	s.Description = &v
	return s
}

func (s *CreateCustomerGatewayRequest) SetIpAddress(v string) *CreateCustomerGatewayRequest {
	s.IpAddress = &v
	return s
}

func (s *CreateCustomerGatewayRequest) SetName(v string) *CreateCustomerGatewayRequest {
	s.Name = &v
	return s
}

func (s *CreateCustomerGatewayRequest) SetOwnerAccount(v string) *CreateCustomerGatewayRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateCustomerGatewayRequest) SetOwnerId(v int64) *CreateCustomerGatewayRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCustomerGatewayRequest) SetRegionId(v string) *CreateCustomerGatewayRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCustomerGatewayRequest) SetResourceGroupId(v string) *CreateCustomerGatewayRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateCustomerGatewayRequest) SetResourceOwnerAccount(v string) *CreateCustomerGatewayRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateCustomerGatewayRequest) SetResourceOwnerId(v int64) *CreateCustomerGatewayRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateCustomerGatewayRequest) SetTags(v []*CreateCustomerGatewayRequestTags) *CreateCustomerGatewayRequest {
	s.Tags = v
	return s
}

func (s *CreateCustomerGatewayRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateCustomerGatewayRequestTags struct {
	// The tag key. The tag key cannot be an empty string.
	//
	// It can be at most 64 characters in length, and cannot contain `http://` or `https://`. It cannot start with `aliyun` or `acs:`.
	//
	// You can specify at most 20 tag keys in each call.
	//
	// example:
	//
	// TagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// The tag value can be an empty string and cannot exceed 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// Each tag key corresponds to one tag value. You can specify at most 20 tag values in each call.
	//
	// example:
	//
	// TagValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateCustomerGatewayRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomerGatewayRequestTags) GoString() string {
	return s.String()
}

func (s *CreateCustomerGatewayRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateCustomerGatewayRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateCustomerGatewayRequestTags) SetKey(v string) *CreateCustomerGatewayRequestTags {
	s.Key = &v
	return s
}

func (s *CreateCustomerGatewayRequestTags) SetValue(v string) *CreateCustomerGatewayRequestTags {
	s.Value = &v
	return s
}

func (s *CreateCustomerGatewayRequestTags) Validate() error {
	return dara.Validate(s)
}
