// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlsLogstoreInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSlsLogstoreInfoRequest
	GetLang() *string
	SetResourceGroupId(v string) *DescribeSlsLogstoreInfoRequest
	GetResourceGroupId() *string
	SetSourceIp(v string) *DescribeSlsLogstoreInfoRequest
	GetSourceIp() *string
}

type DescribeSlsLogstoreInfoRequest struct {
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

func (s DescribeSlsLogstoreInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlsLogstoreInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlsLogstoreInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSlsLogstoreInfoRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSlsLogstoreInfoRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeSlsLogstoreInfoRequest) SetLang(v string) *DescribeSlsLogstoreInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSlsLogstoreInfoRequest) SetResourceGroupId(v string) *DescribeSlsLogstoreInfoRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSlsLogstoreInfoRequest) SetSourceIp(v string) *DescribeSlsLogstoreInfoRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSlsLogstoreInfoRequest) Validate() error {
	return dara.Validate(s)
}
