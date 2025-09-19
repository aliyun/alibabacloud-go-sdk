// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeRiskTypeRequest
	GetLang() *string
	SetSource(v string) *DescribeRiskTypeRequest
	GetSource() *string
	SetSourceIp(v string) *DescribeRiskTypeRequest
	GetSourceIp() *string
}

type DescribeRiskTypeRequest struct {
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
	// The data source. Valid values:
	//
	// 	- **default**: host baseline
	//
	// 	- **agentless**: agentless baseline
	//
	// example:
	//
	// agentless
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 192.168.X.X
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeRiskTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeRiskTypeRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeRiskTypeRequest) GetSource() *string {
	return s.Source
}

func (s *DescribeRiskTypeRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeRiskTypeRequest) SetLang(v string) *DescribeRiskTypeRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRiskTypeRequest) SetSource(v string) *DescribeRiskTypeRequest {
	s.Source = &v
	return s
}

func (s *DescribeRiskTypeRequest) SetSourceIp(v string) *DescribeRiskTypeRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeRiskTypeRequest) Validate() error {
	return dara.Validate(s)
}
