// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSimpleDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *DescribeSimpleDomainsRequest
	GetInstanceIds() []*string
	SetLang(v string) *DescribeSimpleDomainsRequest
	GetLang() *string
	SetResourceGroupId(v string) *DescribeSimpleDomainsRequest
	GetResourceGroupId() *string
	SetSourceIp(v string) *DescribeSimpleDomainsRequest
	GetSourceIp() *string
}

type DescribeSimpleDomainsRequest struct {
	// example:
	//
	// ddoscoo-cn-XXXXXX
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// xx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeSimpleDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSimpleDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSimpleDomainsRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeSimpleDomainsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSimpleDomainsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSimpleDomainsRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeSimpleDomainsRequest) SetInstanceIds(v []*string) *DescribeSimpleDomainsRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeSimpleDomainsRequest) SetLang(v string) *DescribeSimpleDomainsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSimpleDomainsRequest) SetResourceGroupId(v string) *DescribeSimpleDomainsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSimpleDomainsRequest) SetSourceIp(v string) *DescribeSimpleDomainsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSimpleDomainsRequest) Validate() error {
	return dara.Validate(s)
}
