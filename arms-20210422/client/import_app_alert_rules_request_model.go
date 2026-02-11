// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportAppAlertRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactGroupIds(v string) *ImportAppAlertRulesRequest
	GetContactGroupIds() *string
	SetIsAutoStart(v bool) *ImportAppAlertRulesRequest
	GetIsAutoStart() *bool
	SetPids(v string) *ImportAppAlertRulesRequest
	GetPids() *string
	SetRegionId(v string) *ImportAppAlertRulesRequest
	GetRegionId() *string
	SetTemplageAlertConfig(v string) *ImportAppAlertRulesRequest
	GetTemplageAlertConfig() *string
	SetTemplateAlertId(v string) *ImportAppAlertRulesRequest
	GetTemplateAlertId() *string
}

type ImportAppAlertRulesRequest struct {
	ContactGroupIds *string `json:"ContactGroupIds,omitempty" xml:"ContactGroupIds,omitempty"`
	IsAutoStart     *bool   `json:"IsAutoStart,omitempty" xml:"IsAutoStart,omitempty"`
	// This parameter is required.
	Pids *string `json:"Pids,omitempty" xml:"Pids,omitempty"`
	// This parameter is required.
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TemplageAlertConfig *string `json:"TemplageAlertConfig,omitempty" xml:"TemplageAlertConfig,omitempty"`
	TemplateAlertId     *string `json:"TemplateAlertId,omitempty" xml:"TemplateAlertId,omitempty"`
}

func (s ImportAppAlertRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportAppAlertRulesRequest) GoString() string {
	return s.String()
}

func (s *ImportAppAlertRulesRequest) GetContactGroupIds() *string {
	return s.ContactGroupIds
}

func (s *ImportAppAlertRulesRequest) GetIsAutoStart() *bool {
	return s.IsAutoStart
}

func (s *ImportAppAlertRulesRequest) GetPids() *string {
	return s.Pids
}

func (s *ImportAppAlertRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ImportAppAlertRulesRequest) GetTemplageAlertConfig() *string {
	return s.TemplageAlertConfig
}

func (s *ImportAppAlertRulesRequest) GetTemplateAlertId() *string {
	return s.TemplateAlertId
}

func (s *ImportAppAlertRulesRequest) SetContactGroupIds(v string) *ImportAppAlertRulesRequest {
	s.ContactGroupIds = &v
	return s
}

func (s *ImportAppAlertRulesRequest) SetIsAutoStart(v bool) *ImportAppAlertRulesRequest {
	s.IsAutoStart = &v
	return s
}

func (s *ImportAppAlertRulesRequest) SetPids(v string) *ImportAppAlertRulesRequest {
	s.Pids = &v
	return s
}

func (s *ImportAppAlertRulesRequest) SetRegionId(v string) *ImportAppAlertRulesRequest {
	s.RegionId = &v
	return s
}

func (s *ImportAppAlertRulesRequest) SetTemplageAlertConfig(v string) *ImportAppAlertRulesRequest {
	s.TemplageAlertConfig = &v
	return s
}

func (s *ImportAppAlertRulesRequest) SetTemplateAlertId(v string) *ImportAppAlertRulesRequest {
	s.TemplateAlertId = &v
	return s
}

func (s *ImportAppAlertRulesRequest) Validate() error {
	return dara.Validate(s)
}
