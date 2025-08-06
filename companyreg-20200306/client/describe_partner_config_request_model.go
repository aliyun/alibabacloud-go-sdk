// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePartnerConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *DescribePartnerConfigRequest
	GetBizType() *string
	SetPartnerCode(v string) *DescribePartnerConfigRequest
	GetPartnerCode() *string
}

type DescribePartnerConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// esp.wangwen
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// jinsan
	PartnerCode *string `json:"PartnerCode,omitempty" xml:"PartnerCode,omitempty"`
}

func (s DescribePartnerConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePartnerConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribePartnerConfigRequest) GetBizType() *string {
	return s.BizType
}

func (s *DescribePartnerConfigRequest) GetPartnerCode() *string {
	return s.PartnerCode
}

func (s *DescribePartnerConfigRequest) SetBizType(v string) *DescribePartnerConfigRequest {
	s.BizType = &v
	return s
}

func (s *DescribePartnerConfigRequest) SetPartnerCode(v string) *DescribePartnerConfigRequest {
	s.PartnerCode = &v
	return s
}

func (s *DescribePartnerConfigRequest) Validate() error {
	return dara.Validate(s)
}
