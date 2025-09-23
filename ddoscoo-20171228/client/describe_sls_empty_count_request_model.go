// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlsEmptyCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSlsEmptyCountRequest
	GetLang() *string
	SetResourceGroupId(v string) *DescribeSlsEmptyCountRequest
	GetResourceGroupId() *string
	SetSourceIp(v string) *DescribeSlsEmptyCountRequest
	GetSourceIp() *string
}

type DescribeSlsEmptyCountRequest struct {
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

func (s DescribeSlsEmptyCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlsEmptyCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlsEmptyCountRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSlsEmptyCountRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSlsEmptyCountRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeSlsEmptyCountRequest) SetLang(v string) *DescribeSlsEmptyCountRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSlsEmptyCountRequest) SetResourceGroupId(v string) *DescribeSlsEmptyCountRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSlsEmptyCountRequest) SetSourceIp(v string) *DescribeSlsEmptyCountRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSlsEmptyCountRequest) Validate() error {
	return dara.Validate(s)
}
