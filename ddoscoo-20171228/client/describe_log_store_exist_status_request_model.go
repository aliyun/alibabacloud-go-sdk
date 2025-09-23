// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogStoreExistStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeLogStoreExistStatusRequest
	GetLang() *string
	SetResourceGroupId(v string) *DescribeLogStoreExistStatusRequest
	GetResourceGroupId() *string
	SetSourceIp(v string) *DescribeLogStoreExistStatusRequest
	GetSourceIp() *string
}

type DescribeLogStoreExistStatusRequest struct {
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

func (s DescribeLogStoreExistStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogStoreExistStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogStoreExistStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeLogStoreExistStatusRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeLogStoreExistStatusRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeLogStoreExistStatusRequest) SetLang(v string) *DescribeLogStoreExistStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeLogStoreExistStatusRequest) SetResourceGroupId(v string) *DescribeLogStoreExistStatusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeLogStoreExistStatusRequest) SetSourceIp(v string) *DescribeLogStoreExistStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeLogStoreExistStatusRequest) Validate() error {
	return dara.Validate(s)
}
