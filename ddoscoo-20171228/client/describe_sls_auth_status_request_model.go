// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlsAuthStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSlsAuthStatusRequest
	GetLang() *string
	SetResourceGroupId(v string) *DescribeSlsAuthStatusRequest
	GetResourceGroupId() *string
	SetSourceIp(v string) *DescribeSlsAuthStatusRequest
	GetSourceIp() *string
}

type DescribeSlsAuthStatusRequest struct {
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

func (s DescribeSlsAuthStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlsAuthStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlsAuthStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSlsAuthStatusRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSlsAuthStatusRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeSlsAuthStatusRequest) SetLang(v string) *DescribeSlsAuthStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSlsAuthStatusRequest) SetResourceGroupId(v string) *DescribeSlsAuthStatusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSlsAuthStatusRequest) SetSourceIp(v string) *DescribeSlsAuthStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSlsAuthStatusRequest) Validate() error {
	return dara.Validate(s)
}
