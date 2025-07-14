// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationScalingRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeApplicationScalingRuleRequest
	GetAppId() *string
	SetScalingRuleName(v string) *DescribeApplicationScalingRuleRequest
	GetScalingRuleName() *string
}

type DescribeApplicationScalingRuleRequest struct {
	// a0d2e04c-159d-40a8-b240-d2f2c263\\*\\*\\*\\*
	//
	// This parameter is required.
	//
	// example:
	//
	// a0d2e04c-159d-40a8-b240-d2f2c263****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// test
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	ScalingRuleName *string `json:"ScalingRuleName,omitempty" xml:"ScalingRuleName,omitempty"`
}

func (s DescribeApplicationScalingRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRuleRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeApplicationScalingRuleRequest) GetScalingRuleName() *string {
	return s.ScalingRuleName
}

func (s *DescribeApplicationScalingRuleRequest) SetAppId(v string) *DescribeApplicationScalingRuleRequest {
	s.AppId = &v
	return s
}

func (s *DescribeApplicationScalingRuleRequest) SetScalingRuleName(v string) *DescribeApplicationScalingRuleRequest {
	s.ScalingRuleName = &v
	return s
}

func (s *DescribeApplicationScalingRuleRequest) Validate() error {
	return dara.Validate(s)
}
