// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlsOpenStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSlsOpenStatusRequest
	GetLang() *string
	SetResourceGroupId(v string) *DescribeSlsOpenStatusRequest
	GetResourceGroupId() *string
	SetSourceIp(v string) *DescribeSlsOpenStatusRequest
	GetSourceIp() *string
}

type DescribeSlsOpenStatusRequest struct {
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

func (s DescribeSlsOpenStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlsOpenStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlsOpenStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSlsOpenStatusRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSlsOpenStatusRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeSlsOpenStatusRequest) SetLang(v string) *DescribeSlsOpenStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSlsOpenStatusRequest) SetResourceGroupId(v string) *DescribeSlsOpenStatusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSlsOpenStatusRequest) SetSourceIp(v string) *DescribeSlsOpenStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSlsOpenStatusRequest) Validate() error {
	return dara.Validate(s)
}
