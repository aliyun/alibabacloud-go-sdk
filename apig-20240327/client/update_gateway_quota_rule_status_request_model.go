// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayQuotaRuleStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClearHistory(v bool) *UpdateGatewayQuotaRuleStatusRequest
	GetClearHistory() *bool
	SetEnable(v bool) *UpdateGatewayQuotaRuleStatusRequest
	GetEnable() *bool
}

type UpdateGatewayQuotaRuleStatusRequest struct {
	// example:
	//
	// false
	ClearHistory *bool `json:"clearHistory,omitempty" xml:"clearHistory,omitempty"`
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
}

func (s UpdateGatewayQuotaRuleStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayQuotaRuleStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayQuotaRuleStatusRequest) GetClearHistory() *bool {
	return s.ClearHistory
}

func (s *UpdateGatewayQuotaRuleStatusRequest) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateGatewayQuotaRuleStatusRequest) SetClearHistory(v bool) *UpdateGatewayQuotaRuleStatusRequest {
	s.ClearHistory = &v
	return s
}

func (s *UpdateGatewayQuotaRuleStatusRequest) SetEnable(v bool) *UpdateGatewayQuotaRuleStatusRequest {
	s.Enable = &v
	return s
}

func (s *UpdateGatewayQuotaRuleStatusRequest) Validate() error {
	return dara.Validate(s)
}
