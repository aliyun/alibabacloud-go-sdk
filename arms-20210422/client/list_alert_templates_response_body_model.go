// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlertTemplates(v []*ListAlertTemplatesResponseBodyAlertTemplates) *ListAlertTemplatesResponseBody
	GetAlertTemplates() []*ListAlertTemplatesResponseBodyAlertTemplates
	SetRequestId(v string) *ListAlertTemplatesResponseBody
	GetRequestId() *string
}

type ListAlertTemplatesResponseBody struct {
	AlertTemplates []*ListAlertTemplatesResponseBodyAlertTemplates `json:"AlertTemplates,omitempty" xml:"AlertTemplates,omitempty" type:"Repeated"`
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAlertTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAlertTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlertTemplatesResponseBody) GetAlertTemplates() []*ListAlertTemplatesResponseBodyAlertTemplates {
	return s.AlertTemplates
}

func (s *ListAlertTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAlertTemplatesResponseBody) SetAlertTemplates(v []*ListAlertTemplatesResponseBodyAlertTemplates) *ListAlertTemplatesResponseBody {
	s.AlertTemplates = v
	return s
}

func (s *ListAlertTemplatesResponseBody) SetRequestId(v string) *ListAlertTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlertTemplatesResponseBody) Validate() error {
	if s.AlertTemplates != nil {
		for _, item := range s.AlertTemplates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAlertTemplatesResponseBodyAlertTemplates struct {
	AlertProvider            *string                                                               `json:"AlertProvider,omitempty" xml:"AlertProvider,omitempty"`
	Annotations              map[string]interface{}                                                `json:"Annotations,omitempty" xml:"Annotations,omitempty"`
	Id                       *int32                                                                `json:"Id,omitempty" xml:"Id,omitempty"`
	LabelMatchExpressionGrid *ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGrid `json:"LabelMatchExpressionGrid,omitempty" xml:"LabelMatchExpressionGrid,omitempty" type:"Struct"`
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

func (s ListAlertTemplatesResponseBodyAlertTemplates) String() string {
	return dara.Prettify(s)
}

func (s ListAlertTemplatesResponseBodyAlertTemplates) GoString() string {
	return s.String()
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) GetAlertProvider() *string {
	return s.AlertProvider
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) GetAnnotations() map[string]interface{} {
	return s.Annotations
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) GetId() *int32 {
	return s.Id
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) GetLabelMatchExpressionGrid() *ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGrid {
	return s.LabelMatchExpressionGrid
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) GetLabels() map[string]interface{} {
	return s.Labels
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) GetMessage() *string {
	return s.Message
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) GetName() *string {
	return s.Name
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) GetParentId() *int32 {
	return s.ParentId
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) GetPublic() *bool {
	return s.Public
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) GetRule() *string {
	return s.Rule
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) GetStatus() *bool {
	return s.Status
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) GetTemplateProvider() *string {
	return s.TemplateProvider
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) GetType() *string {
	return s.Type
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) GetUserId() *string {
	return s.UserId
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) SetAlertProvider(v string) *ListAlertTemplatesResponseBodyAlertTemplates {
	s.AlertProvider = &v
	return s
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) SetAnnotations(v map[string]interface{}) *ListAlertTemplatesResponseBodyAlertTemplates {
	s.Annotations = v
	return s
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) SetId(v int32) *ListAlertTemplatesResponseBodyAlertTemplates {
	s.Id = &v
	return s
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) SetLabelMatchExpressionGrid(v *ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGrid) *ListAlertTemplatesResponseBodyAlertTemplates {
	s.LabelMatchExpressionGrid = v
	return s
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) SetLabels(v map[string]interface{}) *ListAlertTemplatesResponseBodyAlertTemplates {
	s.Labels = v
	return s
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) SetMessage(v string) *ListAlertTemplatesResponseBodyAlertTemplates {
	s.Message = &v
	return s
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) SetName(v string) *ListAlertTemplatesResponseBodyAlertTemplates {
	s.Name = &v
	return s
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) SetParentId(v int32) *ListAlertTemplatesResponseBodyAlertTemplates {
	s.ParentId = &v
	return s
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) SetPublic(v bool) *ListAlertTemplatesResponseBodyAlertTemplates {
	s.Public = &v
	return s
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) SetRule(v string) *ListAlertTemplatesResponseBodyAlertTemplates {
	s.Rule = &v
	return s
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) SetStatus(v bool) *ListAlertTemplatesResponseBodyAlertTemplates {
	s.Status = &v
	return s
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) SetTemplateProvider(v string) *ListAlertTemplatesResponseBodyAlertTemplates {
	s.TemplateProvider = &v
	return s
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) SetType(v string) *ListAlertTemplatesResponseBodyAlertTemplates {
	s.Type = &v
	return s
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) SetUserId(v string) *ListAlertTemplatesResponseBodyAlertTemplates {
	s.UserId = &v
	return s
}

func (s *ListAlertTemplatesResponseBodyAlertTemplates) Validate() error {
	if s.LabelMatchExpressionGrid != nil {
		if err := s.LabelMatchExpressionGrid.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGrid struct {
	LabelMatchExpressionGroups []*ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroups `json:"LabelMatchExpressionGroups,omitempty" xml:"LabelMatchExpressionGroups,omitempty" type:"Repeated"`
}

func (s ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGrid) String() string {
	return dara.Prettify(s)
}

func (s ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGrid) GoString() string {
	return s.String()
}

func (s *ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGrid) GetLabelMatchExpressionGroups() []*ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroups {
	return s.LabelMatchExpressionGroups
}

func (s *ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGrid) SetLabelMatchExpressionGroups(v []*ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroups) *ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGrid {
	s.LabelMatchExpressionGroups = v
	return s
}

func (s *ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGrid) Validate() error {
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

type ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroups struct {
	LabelMatchExpressions []*ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions `json:"LabelMatchExpressions,omitempty" xml:"LabelMatchExpressions,omitempty" type:"Repeated"`
}

func (s ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroups) String() string {
	return dara.Prettify(s)
}

func (s ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroups) GoString() string {
	return s.String()
}

func (s *ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroups) GetLabelMatchExpressions() []*ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions {
	return s.LabelMatchExpressions
}

func (s *ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroups) SetLabelMatchExpressions(v []*ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) *ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroups {
	s.LabelMatchExpressions = v
	return s
}

func (s *ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroups) Validate() error {
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

type ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions struct {
	Key      *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) String() string {
	return dara.Prettify(s)
}

func (s ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) GoString() string {
	return s.String()
}

func (s *ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) GetKey() *string {
	return s.Key
}

func (s *ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) GetOperator() *string {
	return s.Operator
}

func (s *ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) GetValue() *string {
	return s.Value
}

func (s *ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) SetKey(v string) *ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions {
	s.Key = &v
	return s
}

func (s *ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) SetOperator(v string) *ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions {
	s.Operator = &v
	return s
}

func (s *ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) SetValue(v string) *ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions {
	s.Value = &v
	return s
}

func (s *ListAlertTemplatesResponseBodyAlertTemplatesLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) Validate() error {
	return dara.Validate(s)
}
