// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlertRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertId(v int64) *UpdateAlertRuleRequest
	GetAlertId() *int64
	SetContactGroupIds(v string) *UpdateAlertRuleRequest
	GetContactGroupIds() *string
	SetIsAutoStart(v bool) *UpdateAlertRuleRequest
	GetIsAutoStart() *bool
	SetRegionId(v string) *UpdateAlertRuleRequest
	GetRegionId() *string
	SetTemplageAlertConfig(v string) *UpdateAlertRuleRequest
	GetTemplageAlertConfig() *string
}

type UpdateAlertRuleRequest struct {
	// This parameter is required.
	AlertId         *int64  `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	ContactGroupIds *string `json:"ContactGroupIds,omitempty" xml:"ContactGroupIds,omitempty"`
	IsAutoStart     *bool   `json:"IsAutoStart,omitempty" xml:"IsAutoStart,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	TemplageAlertConfig *string `json:"TemplageAlertConfig,omitempty" xml:"TemplageAlertConfig,omitempty"`
}

func (s UpdateAlertRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateAlertRuleRequest) GetAlertId() *int64 {
	return s.AlertId
}

func (s *UpdateAlertRuleRequest) GetContactGroupIds() *string {
	return s.ContactGroupIds
}

func (s *UpdateAlertRuleRequest) GetIsAutoStart() *bool {
	return s.IsAutoStart
}

func (s *UpdateAlertRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateAlertRuleRequest) GetTemplageAlertConfig() *string {
	return s.TemplageAlertConfig
}

func (s *UpdateAlertRuleRequest) SetAlertId(v int64) *UpdateAlertRuleRequest {
	s.AlertId = &v
	return s
}

func (s *UpdateAlertRuleRequest) SetContactGroupIds(v string) *UpdateAlertRuleRequest {
	s.ContactGroupIds = &v
	return s
}

func (s *UpdateAlertRuleRequest) SetIsAutoStart(v bool) *UpdateAlertRuleRequest {
	s.IsAutoStart = &v
	return s
}

func (s *UpdateAlertRuleRequest) SetRegionId(v string) *UpdateAlertRuleRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateAlertRuleRequest) SetTemplageAlertConfig(v string) *UpdateAlertRuleRequest {
	s.TemplageAlertConfig = &v
	return s
}

func (s *UpdateAlertRuleRequest) Validate() error {
	return dara.Validate(s)
}
