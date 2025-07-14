// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableApplicationScalingRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DisableApplicationScalingRuleRequest
	GetAppId() *string
	SetScalingRuleName(v string) *DisableApplicationScalingRuleRequest
	GetScalingRuleName() *string
}

type DisableApplicationScalingRuleRequest struct {
	// timer-0800-2100
	//
	// This parameter is required.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// timer-0800-2100
	ScalingRuleName *string `json:"ScalingRuleName,omitempty" xml:"ScalingRuleName,omitempty"`
}

func (s DisableApplicationScalingRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableApplicationScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *DisableApplicationScalingRuleRequest) GetAppId() *string {
	return s.AppId
}

func (s *DisableApplicationScalingRuleRequest) GetScalingRuleName() *string {
	return s.ScalingRuleName
}

func (s *DisableApplicationScalingRuleRequest) SetAppId(v string) *DisableApplicationScalingRuleRequest {
	s.AppId = &v
	return s
}

func (s *DisableApplicationScalingRuleRequest) SetScalingRuleName(v string) *DisableApplicationScalingRuleRequest {
	s.ScalingRuleName = &v
	return s
}

func (s *DisableApplicationScalingRuleRequest) Validate() error {
	return dara.Validate(s)
}
