// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFullNatEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ListFullNatEntriesRequest
	GetClientToken() *string
	SetFullNatEntryId(v string) *ListFullNatEntriesRequest
	GetFullNatEntryId() *string
	SetFullNatEntryNames(v []*string) *ListFullNatEntriesRequest
	GetFullNatEntryNames() []*string
	SetFullNatTableId(v string) *ListFullNatEntriesRequest
	GetFullNatTableId() *string
	SetIpProtocol(v string) *ListFullNatEntriesRequest
	GetIpProtocol() *string
	SetMaxResults(v int64) *ListFullNatEntriesRequest
	GetMaxResults() *int64
	SetNatGatewayId(v string) *ListFullNatEntriesRequest
	GetNatGatewayId() *string
	SetNatIp(v string) *ListFullNatEntriesRequest
	GetNatIp() *string
	SetNatIpPort(v string) *ListFullNatEntriesRequest
	GetNatIpPort() *string
	SetNetworkInterfaceIds(v []*string) *ListFullNatEntriesRequest
	GetNetworkInterfaceIds() []*string
	SetNextToken(v string) *ListFullNatEntriesRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListFullNatEntriesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListFullNatEntriesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListFullNatEntriesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListFullNatEntriesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListFullNatEntriesRequest
	GetResourceOwnerId() *int64
}

type ListFullNatEntriesRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate the value from your client and make sure that the value is unique across different requests. The ClientToken parameter supports only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request may be different.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the FULLNAT entry to be queried.
	//
	// example:
	//
	// fullnat-gw8fz23jezpbblf1j****
	FullNatEntryId *string `json:"FullNatEntryId,omitempty" xml:"FullNatEntryId,omitempty"`
	// The names of the FULLNAT entries to be queried. You can query the names of up to 20 FULLNAT entries.
	//
	// The name must be 2 to 128 characters in length, and must start with a letter or a Chinese character. It can contain digits, hyphens (-), and underscores (_).
	FullNatEntryNames []*string `json:"FullNatEntryNames,omitempty" xml:"FullNatEntryNames,omitempty" type:"Repeated"`
	// The ID of the FULLNAT table to which the FULLNAT entries to be queried belong.
	//
	// > You must specify at least one of the **FullNatTableId*	- and **NatGatewayId*	- parameters.
	//
	// example:
	//
	// fulltb-gw88z7hhlv43rmb26****
	FullNatTableId *string `json:"FullNatTableId,omitempty" xml:"FullNatTableId,omitempty"`
	// The protocol type of the Redirection Port. Valid values:
	//
	// - **TCP**: forwards TCP packets.
	//
	// - **UDP**: forwards UDP packets.
	//
	// example:
	//
	// TCP
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The number of entries to return on each page. Valid values: **1*	- to **100**. Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The ID of the NAT gateway to be queried.
	//
	// > You must specify at least one of the **FullNatTableId*	- and **NatGatewayId*	- parameters.
	//
	// example:
	//
	// ngw-bp1uewa15k4iy5770****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The NAT IP address of the FULLNAT entry to be queried.
	//
	// example:
	//
	// 10.0.XX.XX
	NatIp *string `json:"NatIp,omitempty" xml:"NatIp,omitempty"`
	// The frontend port of the FULLNAT entry to be queried. Valid values: 1 to 65535.
	//
	// example:
	//
	// 443
	NatIpPort *string `json:"NatIpPort,omitempty" xml:"NatIpPort,omitempty"`
	// The IDs of the elastic network interfaces (ENIs) to be queried.
	NetworkInterfaceIds []*string `json:"NetworkInterfaceIds,omitempty" xml:"NetworkInterfaceIds,omitempty" type:"Repeated"`
	// The token that is used for the next query. Valid values:
	//
	// - If this is the first query or no next query is to be sent, you do not need to specify this parameter.
	//
	// - If a next query is to be sent, set the value to the **NextToken*	- value returned from the last API call.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the VPC NAT gateway to which the FULLNAT entries to be queried belong.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// eu-central-1
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListFullNatEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFullNatEntriesRequest) GoString() string {
	return s.String()
}

func (s *ListFullNatEntriesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListFullNatEntriesRequest) GetFullNatEntryId() *string {
	return s.FullNatEntryId
}

func (s *ListFullNatEntriesRequest) GetFullNatEntryNames() []*string {
	return s.FullNatEntryNames
}

func (s *ListFullNatEntriesRequest) GetFullNatTableId() *string {
	return s.FullNatTableId
}

func (s *ListFullNatEntriesRequest) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *ListFullNatEntriesRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListFullNatEntriesRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *ListFullNatEntriesRequest) GetNatIp() *string {
	return s.NatIp
}

func (s *ListFullNatEntriesRequest) GetNatIpPort() *string {
	return s.NatIpPort
}

func (s *ListFullNatEntriesRequest) GetNetworkInterfaceIds() []*string {
	return s.NetworkInterfaceIds
}

func (s *ListFullNatEntriesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListFullNatEntriesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListFullNatEntriesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListFullNatEntriesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListFullNatEntriesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListFullNatEntriesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListFullNatEntriesRequest) SetClientToken(v string) *ListFullNatEntriesRequest {
	s.ClientToken = &v
	return s
}

func (s *ListFullNatEntriesRequest) SetFullNatEntryId(v string) *ListFullNatEntriesRequest {
	s.FullNatEntryId = &v
	return s
}

func (s *ListFullNatEntriesRequest) SetFullNatEntryNames(v []*string) *ListFullNatEntriesRequest {
	s.FullNatEntryNames = v
	return s
}

func (s *ListFullNatEntriesRequest) SetFullNatTableId(v string) *ListFullNatEntriesRequest {
	s.FullNatTableId = &v
	return s
}

func (s *ListFullNatEntriesRequest) SetIpProtocol(v string) *ListFullNatEntriesRequest {
	s.IpProtocol = &v
	return s
}

func (s *ListFullNatEntriesRequest) SetMaxResults(v int64) *ListFullNatEntriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListFullNatEntriesRequest) SetNatGatewayId(v string) *ListFullNatEntriesRequest {
	s.NatGatewayId = &v
	return s
}

func (s *ListFullNatEntriesRequest) SetNatIp(v string) *ListFullNatEntriesRequest {
	s.NatIp = &v
	return s
}

func (s *ListFullNatEntriesRequest) SetNatIpPort(v string) *ListFullNatEntriesRequest {
	s.NatIpPort = &v
	return s
}

func (s *ListFullNatEntriesRequest) SetNetworkInterfaceIds(v []*string) *ListFullNatEntriesRequest {
	s.NetworkInterfaceIds = v
	return s
}

func (s *ListFullNatEntriesRequest) SetNextToken(v string) *ListFullNatEntriesRequest {
	s.NextToken = &v
	return s
}

func (s *ListFullNatEntriesRequest) SetOwnerAccount(v string) *ListFullNatEntriesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListFullNatEntriesRequest) SetOwnerId(v int64) *ListFullNatEntriesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListFullNatEntriesRequest) SetRegionId(v string) *ListFullNatEntriesRequest {
	s.RegionId = &v
	return s
}

func (s *ListFullNatEntriesRequest) SetResourceOwnerAccount(v string) *ListFullNatEntriesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListFullNatEntriesRequest) SetResourceOwnerId(v int64) *ListFullNatEntriesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListFullNatEntriesRequest) Validate() error {
	return dara.Validate(s)
}
