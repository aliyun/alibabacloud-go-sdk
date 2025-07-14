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
	// The ID of the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the trace. The ID is used to query the details of a request.
	//
	// This parameter is required.
	//
	// example:
	//
	// timer-0800-2100
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
