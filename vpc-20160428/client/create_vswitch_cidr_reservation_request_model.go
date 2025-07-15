// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVSwitchCidrReservationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateVSwitchCidrReservationRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateVSwitchCidrReservationRequest
	GetDryRun() *bool
	SetIpVersion(v string) *CreateVSwitchCidrReservationRequest
	GetIpVersion() *string
	SetOwnerAccount(v string) *CreateVSwitchCidrReservationRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateVSwitchCidrReservationRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateVSwitchCidrReservationRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateVSwitchCidrReservationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateVSwitchCidrReservationRequest
	GetResourceOwnerId() *int64
	SetTag(v []*CreateVSwitchCidrReservationRequestTag) *CreateVSwitchCidrReservationRequest
	GetTag() []*CreateVSwitchCidrReservationRequestTag
	SetVSwitchCidrReservationCidr(v string) *CreateVSwitchCidrReservationRequest
	GetVSwitchCidrReservationCidr() *string
	SetVSwitchCidrReservationDescription(v string) *CreateVSwitchCidrReservationRequest
	GetVSwitchCidrReservationDescription() *string
	SetVSwitchCidrReservationMask(v string) *CreateVSwitchCidrReservationRequest
	GetVSwitchCidrReservationMask() *string
	SetVSwitchCidrReservationName(v string) *CreateVSwitchCidrReservationRequest
	GetVSwitchCidrReservationName() *string
	SetVSwitchCidrReservationType(v string) *CreateVSwitchCidrReservationRequest
	GetVSwitchCidrReservationType() *string
	SetVSwitchId(v string) *CreateVSwitchCidrReservationRequest
	GetVSwitchId() *string
}

type CreateVSwitchCidrReservationRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run, without performing the actual request. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The IP version of the reserved CIDR block. Valid values:
	//
	// 	- **IPv4*	- (default)
	//
	// 	- **IPv6**
	//
	// example:
	//
	// IPv4
	IpVersion    *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the vSwitch is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Resource tags
	Tag []*CreateVSwitchCidrReservationRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The reserved CIDR block of the vSwitch.
	//
	// 	- When **IpVersion*	- is set to **IPv4**, the reserved CIDR block must be a proper subset of the IPv4 CIDR block of the vSwitch and the subnet mask length of the reserved CIDR block cannot be greater than 28.
	//
	// 	- When **IpVersion*	- is set to **IPv6**, the reserved CIDR block must be a proper subset of the IPv6 CIDR block of the vSwitch and the subnet mask length of the reserved CIDR block cannot be greater than 80.
	//
	// >  You must specify one of **VSwitchCidrReservationMask*	- and **VSwitchCidrReservationCidr**.
	//
	// example:
	//
	// 192.168.1.64/28
	VSwitchCidrReservationCidr *string `json:"VSwitchCidrReservationCidr,omitempty" xml:"VSwitchCidrReservationCidr,omitempty"`
	// The description of the reserved CIDR block. This parameter is empty by default.
	//
	// The description must be 2 to 256 characters in length. It must start with a letter and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// ReservationDescription
	VSwitchCidrReservationDescription *string `json:"VSwitchCidrReservationDescription,omitempty" xml:"VSwitchCidrReservationDescription,omitempty"`
	// The subnet mask of the reserved CIDR block.
	//
	// 	- When **IpVersion*	- is set to **IPv4**, the subnet mask length of the CIDR block must be greater than the IPv4 subnet mask length of the vSwitch and cannot be greater than 28.
	//
	// 	- When **IpVersion*	- is set to **IPv6**, the subnet mask length of the CIDR block must be greater than the IPv6 subnet mask length of the vSwitch and cannot be greater than 80.
	//
	// >  You must specify one of **VSwitchCidrReservationMask*	- and **VSwitchCidrReservationCidr**.
	//
	// example:
	//
	// 28
	VSwitchCidrReservationMask *string `json:"VSwitchCidrReservationMask,omitempty" xml:"VSwitchCidrReservationMask,omitempty"`
	// The name of the reserved CIDR block.
	//
	// The name must be 2 to 128 characters in length and can contain digits, underscores (_), and hyphens (-). It must start with a letter.
	//
	// example:
	//
	// ReservationName
	VSwitchCidrReservationName *string `json:"VSwitchCidrReservationName,omitempty" xml:"VSwitchCidrReservationName,omitempty"`
	// The type of reserved CIDR block. Set the value to **prefix**.
	//
	// >  When a user or a cloud service allocates a CIDR block to an elastic network interface (ENI), the CIDR block must be allocated from the reserved CIDR block. If the reserved CIDR block is exhausted, an error is returned.
	//
	// example:
	//
	// prefix
	VSwitchCidrReservationType *string `json:"VSwitchCidrReservationType,omitempty" xml:"VSwitchCidrReservationType,omitempty"`
	// The ID of the vSwitch to which the reserved CIDR block belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-25navfgbue4g****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateVSwitchCidrReservationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVSwitchCidrReservationRequest) GoString() string {
	return s.String()
}

func (s *CreateVSwitchCidrReservationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateVSwitchCidrReservationRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateVSwitchCidrReservationRequest) GetIpVersion() *string {
	return s.IpVersion
}

func (s *CreateVSwitchCidrReservationRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateVSwitchCidrReservationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateVSwitchCidrReservationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateVSwitchCidrReservationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateVSwitchCidrReservationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateVSwitchCidrReservationRequest) GetTag() []*CreateVSwitchCidrReservationRequestTag {
	return s.Tag
}

func (s *CreateVSwitchCidrReservationRequest) GetVSwitchCidrReservationCidr() *string {
	return s.VSwitchCidrReservationCidr
}

func (s *CreateVSwitchCidrReservationRequest) GetVSwitchCidrReservationDescription() *string {
	return s.VSwitchCidrReservationDescription
}

func (s *CreateVSwitchCidrReservationRequest) GetVSwitchCidrReservationMask() *string {
	return s.VSwitchCidrReservationMask
}

func (s *CreateVSwitchCidrReservationRequest) GetVSwitchCidrReservationName() *string {
	return s.VSwitchCidrReservationName
}

func (s *CreateVSwitchCidrReservationRequest) GetVSwitchCidrReservationType() *string {
	return s.VSwitchCidrReservationType
}

func (s *CreateVSwitchCidrReservationRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateVSwitchCidrReservationRequest) SetClientToken(v string) *CreateVSwitchCidrReservationRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateVSwitchCidrReservationRequest) SetDryRun(v bool) *CreateVSwitchCidrReservationRequest {
	s.DryRun = &v
	return s
}

func (s *CreateVSwitchCidrReservationRequest) SetIpVersion(v string) *CreateVSwitchCidrReservationRequest {
	s.IpVersion = &v
	return s
}

func (s *CreateVSwitchCidrReservationRequest) SetOwnerAccount(v string) *CreateVSwitchCidrReservationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateVSwitchCidrReservationRequest) SetOwnerId(v int64) *CreateVSwitchCidrReservationRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateVSwitchCidrReservationRequest) SetRegionId(v string) *CreateVSwitchCidrReservationRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVSwitchCidrReservationRequest) SetResourceOwnerAccount(v string) *CreateVSwitchCidrReservationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateVSwitchCidrReservationRequest) SetResourceOwnerId(v int64) *CreateVSwitchCidrReservationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateVSwitchCidrReservationRequest) SetTag(v []*CreateVSwitchCidrReservationRequestTag) *CreateVSwitchCidrReservationRequest {
	s.Tag = v
	return s
}

func (s *CreateVSwitchCidrReservationRequest) SetVSwitchCidrReservationCidr(v string) *CreateVSwitchCidrReservationRequest {
	s.VSwitchCidrReservationCidr = &v
	return s
}

func (s *CreateVSwitchCidrReservationRequest) SetVSwitchCidrReservationDescription(v string) *CreateVSwitchCidrReservationRequest {
	s.VSwitchCidrReservationDescription = &v
	return s
}

func (s *CreateVSwitchCidrReservationRequest) SetVSwitchCidrReservationMask(v string) *CreateVSwitchCidrReservationRequest {
	s.VSwitchCidrReservationMask = &v
	return s
}

func (s *CreateVSwitchCidrReservationRequest) SetVSwitchCidrReservationName(v string) *CreateVSwitchCidrReservationRequest {
	s.VSwitchCidrReservationName = &v
	return s
}

func (s *CreateVSwitchCidrReservationRequest) SetVSwitchCidrReservationType(v string) *CreateVSwitchCidrReservationRequest {
	s.VSwitchCidrReservationType = &v
	return s
}

func (s *CreateVSwitchCidrReservationRequest) SetVSwitchId(v string) *CreateVSwitchCidrReservationRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateVSwitchCidrReservationRequest) Validate() error {
	return dara.Validate(s)
}

type CreateVSwitchCidrReservationRequestTag struct {
	// The key of tag N to add to the resource. You can specify at most 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length. It cannot start with aliyun or acs:, and cannot contain http:// or https://.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the resource. You can specify at most 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot start with acs: or aliyun. It cannot contain http:// or https://.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateVSwitchCidrReservationRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateVSwitchCidrReservationRequestTag) GoString() string {
	return s.String()
}

func (s *CreateVSwitchCidrReservationRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateVSwitchCidrReservationRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateVSwitchCidrReservationRequestTag) SetKey(v string) *CreateVSwitchCidrReservationRequestTag {
	s.Key = &v
	return s
}

func (s *CreateVSwitchCidrReservationRequestTag) SetValue(v string) *CreateVSwitchCidrReservationRequestTag {
	s.Value = &v
	return s
}

func (s *CreateVSwitchCidrReservationRequestTag) Validate() error {
	return dara.Validate(s)
}
