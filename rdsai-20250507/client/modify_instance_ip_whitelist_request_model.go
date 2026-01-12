// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceIpWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyInstanceIpWhitelistRequest
	GetClientToken() *string
	SetGroupName(v string) *ModifyInstanceIpWhitelistRequest
	GetGroupName() *string
	SetInstanceName(v string) *ModifyInstanceIpWhitelistRequest
	GetInstanceName() *string
	SetIpWhitelist(v string) *ModifyInstanceIpWhitelistRequest
	GetIpWhitelist() *string
	SetModifyMode(v string) *ModifyInstanceIpWhitelistRequest
	GetModifyMode() *string
	SetRegionId(v string) *ModifyInstanceIpWhitelistRequest
	GetRegionId() *string
}

type ModifyInstanceIpWhitelistRequest struct {
	// The method that is used to modify the IP address whitelist. Valid values:
	//
	// 	- **Cover*	- (default): Uses the IP addresses and CIDR blocks that are specified in the **IpWhitelist*	- parameter to **overwrite*	- the existing ones in the current whitelist.
	//
	// 	- **Append**: **Appends*	- the IP addresses and CIDR blocks that are specified in the **IpWhitelist*	- parameter to the current whitelist.
	//
	// 	- **Delete**: **Deletes*	- the IP addresses and CIDR blocks that are specified in the **IpWhitelist*	- parameter from the current whitelist. You must retain at least one IP address or CIDR block.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The idempotency token. The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.
	//
	// example:
	//
	// default
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The ID of the RDS Supabase instance.
	//
	// example:
	//
	// 10.23.XXX.XXX
	IpWhitelist *string `json:"IpWhitelist,omitempty" xml:"IpWhitelist,omitempty"`
	// The IP address whitelist. Before you modify the IP address whitelist, call the DescribeInstanceIpWhitelist operation to query the existing IP address whitelist of the instance.
	//
	// **Configuration rules**
	//
	// 	- You can configure IP addresses (such as 10.23.XXX.XXX) or CIDR blocks (such as 10.23.XXX.XXX/24).
	//
	// 	- Separate multiple IP addresses or CIDR blocks with commas (,) and do not add spaces preceding or following the commas.
	//
	// 	- You can configure up to 1,000 IP addresses and CIDR blocks in total for each instance. If you want to add a large number of IP addresses, we recommend that you merge the IP addresses into CIDR blocks, such as 10.23.XXX.XXX/24.
	//
	// example:
	//
	// Cover
	ModifyMode *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
	// The operation that you want to perform. Set the value to **ModifyInstanceIpWhitelist**.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyInstanceIpWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceIpWhitelistRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceIpWhitelistRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyInstanceIpWhitelistRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ModifyInstanceIpWhitelistRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyInstanceIpWhitelistRequest) GetIpWhitelist() *string {
	return s.IpWhitelist
}

func (s *ModifyInstanceIpWhitelistRequest) GetModifyMode() *string {
	return s.ModifyMode
}

func (s *ModifyInstanceIpWhitelistRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyInstanceIpWhitelistRequest) SetClientToken(v string) *ModifyInstanceIpWhitelistRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyInstanceIpWhitelistRequest) SetGroupName(v string) *ModifyInstanceIpWhitelistRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyInstanceIpWhitelistRequest) SetInstanceName(v string) *ModifyInstanceIpWhitelistRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyInstanceIpWhitelistRequest) SetIpWhitelist(v string) *ModifyInstanceIpWhitelistRequest {
	s.IpWhitelist = &v
	return s
}

func (s *ModifyInstanceIpWhitelistRequest) SetModifyMode(v string) *ModifyInstanceIpWhitelistRequest {
	s.ModifyMode = &v
	return s
}

func (s *ModifyInstanceIpWhitelistRequest) SetRegionId(v string) *ModifyInstanceIpWhitelistRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceIpWhitelistRequest) Validate() error {
	return dara.Validate(s)
}
