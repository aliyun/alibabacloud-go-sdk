// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLayer4RuleAttributesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListeners(v string) *DescribeLayer4RuleAttributesRequest
	GetListeners() *string
	SetSourceIp(v string) *DescribeLayer4RuleAttributesRequest
	GetSourceIp() *string
}

type DescribeLayer4RuleAttributesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// [{"InstanceId":"ddoscoo-cn-XXXXX","Protocol":"tcp","FrontendPort":80}]
	Listeners *string `json:"Listeners,omitempty" xml:"Listeners,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeLayer4RuleAttributesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLayer4RuleAttributesRequest) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RuleAttributesRequest) GetListeners() *string {
	return s.Listeners
}

func (s *DescribeLayer4RuleAttributesRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeLayer4RuleAttributesRequest) SetListeners(v string) *DescribeLayer4RuleAttributesRequest {
	s.Listeners = &v
	return s
}

func (s *DescribeLayer4RuleAttributesRequest) SetSourceIp(v string) *DescribeLayer4RuleAttributesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeLayer4RuleAttributesRequest) Validate() error {
	return dara.Validate(s)
}
