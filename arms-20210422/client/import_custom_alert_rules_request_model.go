// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportCustomAlertRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactGroupIds(v string) *ImportCustomAlertRulesRequest
	GetContactGroupIds() *string
	SetIsAutoStart(v bool) *ImportCustomAlertRulesRequest
	GetIsAutoStart() *bool
	SetRegionId(v string) *ImportCustomAlertRulesRequest
	GetRegionId() *string
	SetTemplageAlertConfig(v string) *ImportCustomAlertRulesRequest
	GetTemplageAlertConfig() *string
	SetTemplateAlertConfig(v string) *ImportCustomAlertRulesRequest
	GetTemplateAlertConfig() *string
}

type ImportCustomAlertRulesRequest struct {
	ContactGroupIds *string `json:"ContactGroupIds,omitempty" xml:"ContactGroupIds,omitempty"`
	IsAutoStart     *bool   `json:"IsAutoStart,omitempty" xml:"IsAutoStart,omitempty"`
	// This parameter is required.
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TemplageAlertConfig *string `json:"TemplageAlertConfig,omitempty" xml:"TemplageAlertConfig,omitempty"`
	TemplateAlertConfig *string `json:"TemplateAlertConfig,omitempty" xml:"TemplateAlertConfig,omitempty"`
}

func (s ImportCustomAlertRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportCustomAlertRulesRequest) GoString() string {
	return s.String()
}

func (s *ImportCustomAlertRulesRequest) GetContactGroupIds() *string {
	return s.ContactGroupIds
}

func (s *ImportCustomAlertRulesRequest) GetIsAutoStart() *bool {
	return s.IsAutoStart
}

func (s *ImportCustomAlertRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ImportCustomAlertRulesRequest) GetTemplageAlertConfig() *string {
	return s.TemplageAlertConfig
}

func (s *ImportCustomAlertRulesRequest) GetTemplateAlertConfig() *string {
	return s.TemplateAlertConfig
}

func (s *ImportCustomAlertRulesRequest) SetContactGroupIds(v string) *ImportCustomAlertRulesRequest {
	s.ContactGroupIds = &v
	return s
}

func (s *ImportCustomAlertRulesRequest) SetIsAutoStart(v bool) *ImportCustomAlertRulesRequest {
	s.IsAutoStart = &v
	return s
}

func (s *ImportCustomAlertRulesRequest) SetRegionId(v string) *ImportCustomAlertRulesRequest {
	s.RegionId = &v
	return s
}

func (s *ImportCustomAlertRulesRequest) SetTemplageAlertConfig(v string) *ImportCustomAlertRulesRequest {
	s.TemplageAlertConfig = &v
	return s
}

func (s *ImportCustomAlertRulesRequest) SetTemplateAlertConfig(v string) *ImportCustomAlertRulesRequest {
	s.TemplateAlertConfig = &v
	return s
}

func (s *ImportCustomAlertRulesRequest) Validate() error {
	return dara.Validate(s)
}
