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
	// The ID of the application. You can call the ListApplication operation to query the application ID. For more information, see [ListApplication](https://help.aliyun.com/document_detail/149390.html).
	//
	// example:
	//
	// 78194c76-3dca-418e-a263-cccd1ab4****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the auto scaling policy.
	//
	// example:
	//
	// cron-trigger
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
