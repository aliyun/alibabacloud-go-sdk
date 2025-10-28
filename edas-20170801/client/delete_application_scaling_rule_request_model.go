// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationScalingRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteApplicationScalingRuleRequest
	GetAppId() *string
	SetScalingRuleName(v string) *DeleteApplicationScalingRuleRequest
	GetScalingRuleName() *string
}

type DeleteApplicationScalingRuleRequest struct {
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
	// cpu-trigger
	ScalingRuleName *string `json:"ScalingRuleName,omitempty" xml:"ScalingRuleName,omitempty"`
}

func (s DeleteApplicationScalingRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteApplicationScalingRuleRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteApplicationScalingRuleRequest) GetScalingRuleName() *string {
	return s.ScalingRuleName
}

func (s *DeleteApplicationScalingRuleRequest) SetAppId(v string) *DeleteApplicationScalingRuleRequest {
	s.AppId = &v
	return s
}

func (s *DeleteApplicationScalingRuleRequest) SetScalingRuleName(v string) *DeleteApplicationScalingRuleRequest {
	s.ScalingRuleName = &v
	return s
}

func (s *DeleteApplicationScalingRuleRequest) Validate() error {
	return dara.Validate(s)
}
