// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArn(v string) *DescribeNetworkRuleResponseBody
	GetArn() *string
	SetDescription(v string) *DescribeNetworkRuleResponseBody
	GetDescription() *string
	SetRequestId(v string) *DescribeNetworkRuleResponseBody
	GetRequestId() *string
	SetSourcePrivateIp(v string) *DescribeNetworkRuleResponseBody
	GetSourcePrivateIp() *string
	SetType(v string) *DescribeNetworkRuleResponseBody
	GetType() *string
}

type DescribeNetworkRuleResponseBody struct {
	// The ARN of the access control rule.
	//
	// example:
	//
	// acs:kms:cn-hangzhou:119285303511****:network/networkrule_test
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The description.
	//
	// example:
	//
	// Creat by kst-hzz62ee817bvyyr5****
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 3bf02f7a-015b-4f93-be0f-cc043fda2d33
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The private IP address or private CIDR block.
	//
	// example:
	//
	// ["192.10.XX.XX","192.168.XX.XX/24"]
	SourcePrivateIp *string `json:"SourcePrivateIp,omitempty" xml:"SourcePrivateIp,omitempty"`
	// The network type. Only private IP addresses are supported. The value is fixed as Private.
	//
	// example:
	//
	// Private
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeNetworkRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNetworkRuleResponseBody) GetArn() *string {
	return s.Arn
}

func (s *DescribeNetworkRuleResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeNetworkRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNetworkRuleResponseBody) GetSourcePrivateIp() *string {
	return s.SourcePrivateIp
}

func (s *DescribeNetworkRuleResponseBody) GetType() *string {
	return s.Type
}

func (s *DescribeNetworkRuleResponseBody) SetArn(v string) *DescribeNetworkRuleResponseBody {
	s.Arn = &v
	return s
}

func (s *DescribeNetworkRuleResponseBody) SetDescription(v string) *DescribeNetworkRuleResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeNetworkRuleResponseBody) SetRequestId(v string) *DescribeNetworkRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNetworkRuleResponseBody) SetSourcePrivateIp(v string) *DescribeNetworkRuleResponseBody {
	s.SourcePrivateIp = &v
	return s
}

func (s *DescribeNetworkRuleResponseBody) SetType(v string) *DescribeNetworkRuleResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeNetworkRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
