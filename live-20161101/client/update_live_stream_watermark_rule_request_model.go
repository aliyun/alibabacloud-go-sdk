// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveStreamWatermarkRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateLiveStreamWatermarkRuleRequest
	GetDescription() *string
	SetName(v string) *UpdateLiveStreamWatermarkRuleRequest
	GetName() *string
	SetOwnerId(v int64) *UpdateLiveStreamWatermarkRuleRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpdateLiveStreamWatermarkRuleRequest
	GetRegionId() *string
	SetRuleId(v string) *UpdateLiveStreamWatermarkRuleRequest
	GetRuleId() *string
	SetTemplateId(v string) *UpdateLiveStreamWatermarkRuleRequest
	GetTemplateId() *string
}

type UpdateLiveStreamWatermarkRuleRequest struct {
	// The description of the custom rule.
	//
	// example:
	//
	// my rule
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the custom rule.
	//
	// example:
	//
	// WatermarkRule****
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the watermark rule.
	//
	// >  You can obtain the rule ID by checking the value of the RuleId parameter that is returned by the [AddLiveStreamWatermarkRule](https://help.aliyun.com/document_detail/2848100.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 445409ec-7eaa-461d-8f29-4bec2eb9****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The ID of the watermark template.
	//
	// >  You can obtain the template ID by checking the value of the TemplateId parameter that is returned by the [AddLiveStreamWatermark](https://help.aliyun.com/document_detail/2848096.html) operation.
	//
	// example:
	//
	// 445409ec-7eaa-461d-8f29-4bec2eb9 ****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s UpdateLiveStreamWatermarkRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveStreamWatermarkRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveStreamWatermarkRuleRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateLiveStreamWatermarkRuleRequest) GetName() *string {
	return s.Name
}

func (s *UpdateLiveStreamWatermarkRuleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateLiveStreamWatermarkRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateLiveStreamWatermarkRuleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *UpdateLiveStreamWatermarkRuleRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateLiveStreamWatermarkRuleRequest) SetDescription(v string) *UpdateLiveStreamWatermarkRuleRequest {
	s.Description = &v
	return s
}

func (s *UpdateLiveStreamWatermarkRuleRequest) SetName(v string) *UpdateLiveStreamWatermarkRuleRequest {
	s.Name = &v
	return s
}

func (s *UpdateLiveStreamWatermarkRuleRequest) SetOwnerId(v int64) *UpdateLiveStreamWatermarkRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateLiveStreamWatermarkRuleRequest) SetRegionId(v string) *UpdateLiveStreamWatermarkRuleRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateLiveStreamWatermarkRuleRequest) SetRuleId(v string) *UpdateLiveStreamWatermarkRuleRequest {
	s.RuleId = &v
	return s
}

func (s *UpdateLiveStreamWatermarkRuleRequest) SetTemplateId(v string) *UpdateLiveStreamWatermarkRuleRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateLiveStreamWatermarkRuleRequest) Validate() error {
	return dara.Validate(s)
}
