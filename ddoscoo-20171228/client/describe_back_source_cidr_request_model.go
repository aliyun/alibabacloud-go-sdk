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
	SetSourceIp(v string) *DescribeBackSourceCidrRequest
	GetSourceIp() *string
}

type DescribeBackSourceCidrRequest struct {
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// example:
	//
	// coop-line-001
	Line *string `json:"Line,omitempty" xml:"Line,omitempty"`
	// example:
	//
	// test
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
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

func (s *DescribeBackSourceCidrRequest) GetSourceIp() *string {
	return s.SourceIp
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

func (s *DescribeBackSourceCidrRequest) SetSourceIp(v string) *DescribeBackSourceCidrRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeBackSourceCidrRequest) Validate() error {
	return dara.Validate(s)
}
