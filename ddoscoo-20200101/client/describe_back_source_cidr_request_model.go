// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackSourceCidrRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpVersion(v string) *DescribeBackSourceCidrRequest
	GetIpVersion() *string
	SetLine(v string) *DescribeBackSourceCidrRequest
	GetLine() *string
	SetResourceGroupId(v string) *DescribeBackSourceCidrRequest
	GetResourceGroupId() *string
}

type DescribeBackSourceCidrRequest struct {
	// The IP version of the back-to-origin CIDR blocks.
	//
	// 	- **Ipv4**
	//
	// 	- **Ipv6**
	//
	// example:
	//
	// IPv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The Internet service provider (ISP) line that you want to query.
	//
	// example:
	//
	// coop-line-001
	Line *string `json:"Line,omitempty" xml:"Line,omitempty"`
	// The ID of the resource group to which the instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeBackSourceCidrRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackSourceCidrRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackSourceCidrRequest) GetIpVersion() *string {
	return s.IpVersion
}

func (s *DescribeBackSourceCidrRequest) GetLine() *string {
	return s.Line
}

func (s *DescribeBackSourceCidrRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeBackSourceCidrRequest) SetIpVersion(v string) *DescribeBackSourceCidrRequest {
	s.IpVersion = &v
	return s
}

func (s *DescribeBackSourceCidrRequest) SetLine(v string) *DescribeBackSourceCidrRequest {
	s.Line = &v
	return s
}

func (s *DescribeBackSourceCidrRequest) SetResourceGroupId(v string) *DescribeBackSourceCidrRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeBackSourceCidrRequest) Validate() error {
	return dara.Validate(s)
}
