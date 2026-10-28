// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSearchConditionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSearchConditionRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeSearchConditionRequest
	GetSourceIp() *string
	SetType(v string) *DescribeSearchConditionRequest
	GetType() *string
}

type DescribeSearchConditionRequest struct {
	// The language of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 117.220.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The Asset Type to query. Valid values:
	//
	// - **ecs**: queries the filter conditions for all ECS servers.
	//
	// - **cloud_product**: queries the filter conditions for all other cloud services.
	//
	// example:
	//
	// ecs
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeSearchConditionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSearchConditionRequest) GoString() string {
	return s.String()
}

func (s *DescribeSearchConditionRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSearchConditionRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeSearchConditionRequest) GetType() *string {
	return s.Type
}

func (s *DescribeSearchConditionRequest) SetLang(v string) *DescribeSearchConditionRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSearchConditionRequest) SetSourceIp(v string) *DescribeSearchConditionRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSearchConditionRequest) SetType(v string) *DescribeSearchConditionRequest {
	s.Type = &v
	return s
}

func (s *DescribeSearchConditionRequest) Validate() error {
	return dara.Validate(s)
}
