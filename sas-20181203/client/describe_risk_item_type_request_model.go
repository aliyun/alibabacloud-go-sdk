// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskItemTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeRiskItemTypeRequest
	GetLang() *string
	SetResourceOwnerId(v int64) *DescribeRiskItemTypeRequest
	GetResourceOwnerId() *int64
	SetSourceIp(v string) *DescribeRiskItemTypeRequest
	GetSourceIp() *string
}

type DescribeRiskItemTypeRequest struct {
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 183.237.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeRiskItemTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskItemTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeRiskItemTypeRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeRiskItemTypeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeRiskItemTypeRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeRiskItemTypeRequest) SetLang(v string) *DescribeRiskItemTypeRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRiskItemTypeRequest) SetResourceOwnerId(v int64) *DescribeRiskItemTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRiskItemTypeRequest) SetSourceIp(v string) *DescribeRiskItemTypeRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeRiskItemTypeRequest) Validate() error {
	return dara.Validate(s)
}
