// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRemediationTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRemediationTemplates(v []*GetRemediationTemplateResponseBodyRemediationTemplates) *GetRemediationTemplateResponseBody
	GetRemediationTemplates() []*GetRemediationTemplateResponseBodyRemediationTemplates
	SetRequestId(v string) *GetRemediationTemplateResponseBody
	GetRequestId() *string
}

type GetRemediationTemplateResponseBody struct {
	// The information about the automatic remediation template.
	RemediationTemplates []*GetRemediationTemplateResponseBodyRemediationTemplates `json:"RemediationTemplates,omitempty" xml:"RemediationTemplates,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// E232FC35-BD40-51E3-B2EB-09416A23****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRemediationTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRemediationTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetRemediationTemplateResponseBody) GetRemediationTemplates() []*GetRemediationTemplateResponseBodyRemediationTemplates {
	return s.RemediationTemplates
}

func (s *GetRemediationTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRemediationTemplateResponseBody) SetRemediationTemplates(v []*GetRemediationTemplateResponseBodyRemediationTemplates) *GetRemediationTemplateResponseBody {
	s.RemediationTemplates = v
	return s
}

func (s *GetRemediationTemplateResponseBody) SetRequestId(v string) *GetRemediationTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRemediationTemplateResponseBody) Validate() error {
	if s.RemediationTemplates != nil {
		for _, item := range s.RemediationTemplates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRemediationTemplateResponseBodyRemediationTemplates struct {
	// The identifier of the supported rule template.
	//
	// This parameter is required.
	//
	// example:
	//
	// alb-delete-protection-enabled
	ManagedRuleIdentifier *string `json:"ManagedRuleIdentifier,omitempty" xml:"ManagedRuleIdentifier,omitempty"`
	// The type of the automatic remediation template. Valid value: OOS (Operation Orchestration).
	//
	// example:
	//
	// OOS
	RemediationType *string `json:"RemediationType,omitempty" xml:"RemediationType,omitempty"`
	// The definition of the template parameters.
	//
	// example:
	//
	// {\\"Parameters\\":{\\"regionId\\":{\\"AllowValues\\":\\"\\",\\"AssociationProperty\\":\\"RegionId\\",\\"CreateDated\\":\\"2025-08-04T09:54:57\\",\\"Default\\":\\"{regionId}\\",\\"Description\\":{\\"en\\":\\"regionId\\",\\"zh-cn\\":\\"regionId\\"},\\"Id\\":688,\\"MaxKeyLength\\":\\"125\\",\\"MaxValueLength\\":\\"255\\",\\"ModifiedDate\\":\\"2025-08-04T09:54:57\\",\\"Name\\":\\"regionId\\",\\"Optional\\":1,\\"TemplateIdentifier\\":\\"ACS-ALB-BulkyEnableDeletionProtection\\",\\"Type\\":\\"String\\",\\"Version\\":\\"LASTEST\\"},\\"loadBalancerIds\\":{\\"AllowValues\\":\\"\\",\\"AssociationProperty\\":\\"\\",\\"CreateDated\\":\\"2025-08-04T09:54:57\\",\\"Default\\":\\"[\\\\\\"{resourceId}\\\\\\"]\\",\\"Description\\":{\\"en\\":\\"loadBalancerIds\\",\\"zh-cn\\":\\"loadBalancerIds\\"},\\"Id\\":689,\\"MaxKeyLength\\":\\"125\\",\\"MaxValueLength\\":\\"255\\",\\"ModifiedDate\\":\\"2025-08-04T09:54:57\\",\\"Name\\":\\"loadBalancerIds\\",\\"Optional\\":1,\\"TemplateIdentifier\\":\\"ACS-ALB-BulkyEnableDeletionProtection\\",\\"Type\\":\\"ARRAY\\",\\"Version\\":\\"LASTEST\\"}}}
	TemplateDefinition *string `json:"TemplateDefinition,omitempty" xml:"TemplateDefinition,omitempty"`
	// The description of the automatic remediation template.
	//
	// This parameter is required.
	//
	// example:
	//
	// Call the EnableDeletionProtection interface to enable ALB instance deletion protection. Be aware of the risks and exercise caution.
	TemplateDescription *string `json:"TemplateDescription,omitempty" xml:"TemplateDescription,omitempty"`
	// The identifier of the automatic remediation template.
	//
	// example:
	//
	// ACS-ALB-BulkyEnableDeletionProtection
	TemplateIdentifier *string `json:"TemplateIdentifier,omitempty" xml:"TemplateIdentifier,omitempty"`
	// The name of the automatic remediation template.
	//
	// example:
	//
	// Enable ALB instance deletion protection
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s GetRemediationTemplateResponseBodyRemediationTemplates) String() string {
	return dara.Prettify(s)
}

func (s GetRemediationTemplateResponseBodyRemediationTemplates) GoString() string {
	return s.String()
}

func (s *GetRemediationTemplateResponseBodyRemediationTemplates) GetManagedRuleIdentifier() *string {
	return s.ManagedRuleIdentifier
}

func (s *GetRemediationTemplateResponseBodyRemediationTemplates) GetRemediationType() *string {
	return s.RemediationType
}

func (s *GetRemediationTemplateResponseBodyRemediationTemplates) GetTemplateDefinition() *string {
	return s.TemplateDefinition
}

func (s *GetRemediationTemplateResponseBodyRemediationTemplates) GetTemplateDescription() *string {
	return s.TemplateDescription
}

func (s *GetRemediationTemplateResponseBodyRemediationTemplates) GetTemplateIdentifier() *string {
	return s.TemplateIdentifier
}

func (s *GetRemediationTemplateResponseBodyRemediationTemplates) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetRemediationTemplateResponseBodyRemediationTemplates) SetManagedRuleIdentifier(v string) *GetRemediationTemplateResponseBodyRemediationTemplates {
	s.ManagedRuleIdentifier = &v
	return s
}

func (s *GetRemediationTemplateResponseBodyRemediationTemplates) SetRemediationType(v string) *GetRemediationTemplateResponseBodyRemediationTemplates {
	s.RemediationType = &v
	return s
}

func (s *GetRemediationTemplateResponseBodyRemediationTemplates) SetTemplateDefinition(v string) *GetRemediationTemplateResponseBodyRemediationTemplates {
	s.TemplateDefinition = &v
	return s
}

func (s *GetRemediationTemplateResponseBodyRemediationTemplates) SetTemplateDescription(v string) *GetRemediationTemplateResponseBodyRemediationTemplates {
	s.TemplateDescription = &v
	return s
}

func (s *GetRemediationTemplateResponseBodyRemediationTemplates) SetTemplateIdentifier(v string) *GetRemediationTemplateResponseBodyRemediationTemplates {
	s.TemplateIdentifier = &v
	return s
}

func (s *GetRemediationTemplateResponseBodyRemediationTemplates) SetTemplateName(v string) *GetRemediationTemplateResponseBodyRemediationTemplates {
	s.TemplateName = &v
	return s
}

func (s *GetRemediationTemplateResponseBodyRemediationTemplates) Validate() error {
	return dara.Validate(s)
}
