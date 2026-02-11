// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlertTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlertTemplate(v *CreateAlertTemplateResponseBodyAlertTemplate) *CreateAlertTemplateResponseBody
	GetAlertTemplate() *CreateAlertTemplateResponseBodyAlertTemplate
	SetRequestId(v string) *CreateAlertTemplateResponseBody
	GetRequestId() *string
}

type CreateAlertTemplateResponseBody struct {
	AlertTemplate *CreateAlertTemplateResponseBodyAlertTemplate `json:"AlertTemplate,omitempty" xml:"AlertTemplate,omitempty" type:"Struct"`
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAlertTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAlertTemplateResponseBody) GetAlertTemplate() *CreateAlertTemplateResponseBodyAlertTemplate {
	return s.AlertTemplate
}

func (s *CreateAlertTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAlertTemplateResponseBody) SetAlertTemplate(v *CreateAlertTemplateResponseBodyAlertTemplate) *CreateAlertTemplateResponseBody {
	s.AlertTemplate = v
	return s
}

func (s *CreateAlertTemplateResponseBody) SetRequestId(v string) *CreateAlertTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAlertTemplateResponseBody) Validate() error {
	if s.AlertTemplate != nil {
		if err := s.AlertTemplate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAlertTemplateResponseBodyAlertTemplate struct {
	AlertProvider            *string                                                               `json:"AlertProvider,omitempty" xml:"AlertProvider,omitempty"`
	Annotations              map[string]interface{}                                                `json:"Annotations,omitempty" xml:"Annotations,omitempty"`
	Id                       *int32                                                                `json:"Id,omitempty" xml:"Id,omitempty"`
	LabelMatchExpressionGrid *CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGrid `json:"LabelMatchExpressionGrid,omitempty" xml:"LabelMatchExpressionGrid,omitempty" type:"Struct"`
	Labels                   map[string]interface{}                                                `json:"Labels,omitempty" xml:"Labels,omitempty"`
	Message                  *string                                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	Name                     *string                                                               `json:"Name,omitempty" xml:"Name,omitempty"`
	ParentId                 *int32                                                                `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	Public                   *bool                                                                 `json:"Public,omitempty" xml:"Public,omitempty"`
	Rule                     *string                                                               `json:"Rule,omitempty" xml:"Rule,omitempty"`
	Status                   *bool                                                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	TemplateProvider         *string                                                               `json:"TemplateProvider,omitempty" xml:"TemplateProvider,omitempty"`
	Type                     *string                                                               `json:"Type,omitempty" xml:"Type,omitempty"`
	UserId                   *string                                                               `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateAlertTemplateResponseBodyAlertTemplate) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertTemplateResponseBodyAlertTemplate) GoString() string {
	return s.String()
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) GetAlertProvider() *string {
	return s.AlertProvider
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) GetAnnotations() map[string]interface{} {
	return s.Annotations
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) GetId() *int32 {
	return s.Id
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) GetLabelMatchExpressionGrid() *CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGrid {
	return s.LabelMatchExpressionGrid
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) GetLabels() map[string]interface{} {
	return s.Labels
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) GetMessage() *string {
	return s.Message
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) GetName() *string {
	return s.Name
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) GetParentId() *int32 {
	return s.ParentId
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) GetPublic() *bool {
	return s.Public
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) GetRule() *string {
	return s.Rule
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) GetStatus() *bool {
	return s.Status
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) GetTemplateProvider() *string {
	return s.TemplateProvider
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) GetType() *string {
	return s.Type
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) GetUserId() *string {
	return s.UserId
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) SetAlertProvider(v string) *CreateAlertTemplateResponseBodyAlertTemplate {
	s.AlertProvider = &v
	return s
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) SetAnnotations(v map[string]interface{}) *CreateAlertTemplateResponseBodyAlertTemplate {
	s.Annotations = v
	return s
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) SetId(v int32) *CreateAlertTemplateResponseBodyAlertTemplate {
	s.Id = &v
	return s
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) SetLabelMatchExpressionGrid(v *CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGrid) *CreateAlertTemplateResponseBodyAlertTemplate {
	s.LabelMatchExpressionGrid = v
	return s
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) SetLabels(v map[string]interface{}) *CreateAlertTemplateResponseBodyAlertTemplate {
	s.Labels = v
	return s
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) SetMessage(v string) *CreateAlertTemplateResponseBodyAlertTemplate {
	s.Message = &v
	return s
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) SetName(v string) *CreateAlertTemplateResponseBodyAlertTemplate {
	s.Name = &v
	return s
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) SetParentId(v int32) *CreateAlertTemplateResponseBodyAlertTemplate {
	s.ParentId = &v
	return s
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) SetPublic(v bool) *CreateAlertTemplateResponseBodyAlertTemplate {
	s.Public = &v
	return s
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) SetRule(v string) *CreateAlertTemplateResponseBodyAlertTemplate {
	s.Rule = &v
	return s
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) SetStatus(v bool) *CreateAlertTemplateResponseBodyAlertTemplate {
	s.Status = &v
	return s
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) SetTemplateProvider(v string) *CreateAlertTemplateResponseBodyAlertTemplate {
	s.TemplateProvider = &v
	return s
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) SetType(v string) *CreateAlertTemplateResponseBodyAlertTemplate {
	s.Type = &v
	return s
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) SetUserId(v string) *CreateAlertTemplateResponseBodyAlertTemplate {
	s.UserId = &v
	return s
}

func (s *CreateAlertTemplateResponseBodyAlertTemplate) Validate() error {
	if s.LabelMatchExpressionGrid != nil {
		if err := s.LabelMatchExpressionGrid.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGrid struct {
	LabelMatchExpressionGroups []*CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroups `json:"LabelMatchExpressionGroups,omitempty" xml:"LabelMatchExpressionGroups,omitempty" type:"Repeated"`
}

func (s CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGrid) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGrid) GoString() string {
	return s.String()
}

func (s *CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGrid) GetLabelMatchExpressionGroups() []*CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroups {
	return s.LabelMatchExpressionGroups
}

func (s *CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGrid) SetLabelMatchExpressionGroups(v []*CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroups) *CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGrid {
	s.LabelMatchExpressionGroups = v
	return s
}

func (s *CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGrid) Validate() error {
	if s.LabelMatchExpressionGroups != nil {
		for _, item := range s.LabelMatchExpressionGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroups struct {
	LabelMatchExpressions []*CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions `json:"LabelMatchExpressions,omitempty" xml:"LabelMatchExpressions,omitempty" type:"Repeated"`
}

func (s CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroups) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroups) GoString() string {
	return s.String()
}

func (s *CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroups) GetLabelMatchExpressions() []*CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions {
	return s.LabelMatchExpressions
}

func (s *CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroups) SetLabelMatchExpressions(v []*CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) *CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroups {
	s.LabelMatchExpressions = v
	return s
}

func (s *CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroups) Validate() error {
	if s.LabelMatchExpressions != nil {
		for _, item := range s.LabelMatchExpressions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions struct {
	Key      *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) GoString() string {
	return s.String()
}

func (s *CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) GetKey() *string {
	return s.Key
}

func (s *CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) GetOperator() *string {
	return s.Operator
}

func (s *CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) GetValue() *string {
	return s.Value
}

func (s *CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) SetKey(v string) *CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions {
	s.Key = &v
	return s
}

func (s *CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) SetOperator(v string) *CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions {
	s.Operator = &v
	return s
}

func (s *CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) SetValue(v string) *CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions {
	s.Value = &v
	return s
}

func (s *CreateAlertTemplateResponseBodyAlertTemplateLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) Validate() error {
	return dara.Validate(s)
}
