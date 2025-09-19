// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStrategyDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DescribeStrategyDetailRequest
	GetId() *string
	SetLang(v string) *DescribeStrategyDetailRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeStrategyDetailRequest
	GetSourceIp() *string
}

type DescribeStrategyDetailRequest struct {
	// The ID of the baseline check policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 192.168.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeStrategyDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeStrategyDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeStrategyDetailRequest) GetId() *string {
	return s.Id
}

func (s *DescribeStrategyDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeStrategyDetailRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeStrategyDetailRequest) SetId(v string) *DescribeStrategyDetailRequest {
	s.Id = &v
	return s
}

func (s *DescribeStrategyDetailRequest) SetLang(v string) *DescribeStrategyDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeStrategyDetailRequest) SetSourceIp(v string) *DescribeStrategyDetailRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeStrategyDetailRequest) Validate() error {
	return dara.Validate(s)
}
