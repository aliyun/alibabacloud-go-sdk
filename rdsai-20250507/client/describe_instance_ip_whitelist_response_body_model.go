// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceIpWhitelistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *DescribeInstanceIpWhitelistResponseBody
	GetInstanceName() *string
	SetIpWhiteListGroups(v []*DescribeInstanceIpWhitelistResponseBodyIpWhiteListGroups) *DescribeInstanceIpWhitelistResponseBody
	GetIpWhiteListGroups() []*DescribeInstanceIpWhitelistResponseBodyIpWhiteListGroups
	SetRequestId(v string) *DescribeInstanceIpWhitelistResponseBody
	GetRequestId() *string
}

type DescribeInstanceIpWhitelistResponseBody struct {
	// The ID of the RDS Supabase instance.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The IP address whitelists.
	IpWhiteListGroups []*DescribeInstanceIpWhitelistResponseBodyIpWhiteListGroups `json:"IpWhiteListGroups,omitempty" xml:"IpWhiteListGroups,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 32DEFB4A-861F-5D1D-ADD5-918E4FD7AB8C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceIpWhitelistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceIpWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceIpWhitelistResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeInstanceIpWhitelistResponseBody) GetIpWhiteListGroups() []*DescribeInstanceIpWhitelistResponseBodyIpWhiteListGroups {
	return s.IpWhiteListGroups
}

func (s *DescribeInstanceIpWhitelistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceIpWhitelistResponseBody) SetInstanceName(v string) *DescribeInstanceIpWhitelistResponseBody {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstanceIpWhitelistResponseBody) SetIpWhiteListGroups(v []*DescribeInstanceIpWhitelistResponseBodyIpWhiteListGroups) *DescribeInstanceIpWhitelistResponseBody {
	s.IpWhiteListGroups = v
	return s
}

func (s *DescribeInstanceIpWhitelistResponseBody) SetRequestId(v string) *DescribeInstanceIpWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceIpWhitelistResponseBody) Validate() error {
	if s.IpWhiteListGroups != nil {
		for _, item := range s.IpWhiteListGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceIpWhitelistResponseBodyIpWhiteListGroups struct {
	// The IP address whitelist name.
	//
	// example:
	//
	// default
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The IP addresses or CIDR blocks in the whitelist.
	//
	// example:
	//
	// 192.168.XXX.XXX/24,10.0.XXX.XXX/24
	IpWhitelist *string `json:"IpWhitelist,omitempty" xml:"IpWhitelist,omitempty"`
}

func (s DescribeInstanceIpWhitelistResponseBodyIpWhiteListGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceIpWhitelistResponseBodyIpWhiteListGroups) GoString() string {
	return s.String()
}

func (s *DescribeInstanceIpWhitelistResponseBodyIpWhiteListGroups) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeInstanceIpWhitelistResponseBodyIpWhiteListGroups) GetIpWhitelist() *string {
	return s.IpWhitelist
}

func (s *DescribeInstanceIpWhitelistResponseBodyIpWhiteListGroups) SetGroupName(v string) *DescribeInstanceIpWhitelistResponseBodyIpWhiteListGroups {
	s.GroupName = &v
	return s
}

func (s *DescribeInstanceIpWhitelistResponseBodyIpWhiteListGroups) SetIpWhitelist(v string) *DescribeInstanceIpWhitelistResponseBodyIpWhiteListGroups {
	s.IpWhitelist = &v
	return s
}

func (s *DescribeInstanceIpWhitelistResponseBodyIpWhiteListGroups) Validate() error {
	return dara.Validate(s)
}
