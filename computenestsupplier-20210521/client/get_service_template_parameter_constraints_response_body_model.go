// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceTemplateParameterConstraintsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFamilyConstraints(v []*string) *GetServiceTemplateParameterConstraintsResponseBody
	GetFamilyConstraints() []*string
	SetParameterConstraints(v []*GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) *GetServiceTemplateParameterConstraintsResponseBody
	GetParameterConstraints() []*GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints
	SetRequestId(v string) *GetServiceTemplateParameterConstraintsResponseBody
	GetRequestId() *string
}

type GetServiceTemplateParameterConstraintsResponseBody struct {
	// The constraint families.
	FamilyConstraints []*string `json:"FamilyConstraints,omitempty" xml:"FamilyConstraints,omitempty" type:"Repeated"`
	// The parameters in the template.
	ParameterConstraints []*GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints `json:"ParameterConstraints,omitempty" xml:"ParameterConstraints,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// C81C0732-DEBC-559C-B563-7EB2BEB21088
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetServiceTemplateParameterConstraintsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceTemplateParameterConstraintsResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceTemplateParameterConstraintsResponseBody) GetFamilyConstraints() []*string {
	return s.FamilyConstraints
}

func (s *GetServiceTemplateParameterConstraintsResponseBody) GetParameterConstraints() []*GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints {
	return s.ParameterConstraints
}

func (s *GetServiceTemplateParameterConstraintsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetServiceTemplateParameterConstraintsResponseBody) SetFamilyConstraints(v []*string) *GetServiceTemplateParameterConstraintsResponseBody {
	s.FamilyConstraints = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBody) SetParameterConstraints(v []*GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) *GetServiceTemplateParameterConstraintsResponseBody {
	s.ParameterConstraints = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBody) SetRequestId(v string) *GetServiceTemplateParameterConstraintsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints struct {
	// The valid values of the parameter.
	AllowedValues []*string `json:"AllowedValues,omitempty" xml:"AllowedValues,omitempty" type:"Repeated"`
	// The names of the associated parameters.
	AssociationParameterNames []*string `json:"AssociationParameterNames,omitempty" xml:"AssociationParameterNames,omitempty" type:"Repeated"`
	// The behavior of the parameter. Valid values:
	//
	// 	- NoLimit: The value of this parameter is not limited.
	//
	// 	- NotSupport: The value of this parameter cannot be queried.
	//
	// 	- QueryError: The query failed.
	//
	// >  If AllowedValues is not returned, Behavior and BehaviorReason are returned.
	//
	// example:
	//
	// NoLimit
	Behavior *string `json:"Behavior,omitempty" xml:"Behavior,omitempty"`
	// The reason why the behavior of the parameter is returned.
	//
	// example:
	//
	// none
	BehaviorReason *string `json:"BehaviorReason,omitempty" xml:"BehaviorReason,omitempty"`
	// The original constraint information.
	OriginalConstraints []*GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints `json:"OriginalConstraints,omitempty" xml:"OriginalConstraints,omitempty" type:"Repeated"`
	// The name of the parameter.
	//
	// example:
	//
	// PayType
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The type of the parameter.
	//
	// example:
	//
	// String
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) String() string {
	return dara.Prettify(s)
}

func (s GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) GoString() string {
	return s.String()
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) GetAllowedValues() []*string {
	return s.AllowedValues
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) GetAssociationParameterNames() []*string {
	return s.AssociationParameterNames
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) GetBehavior() *string {
	return s.Behavior
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) GetBehaviorReason() *string {
	return s.BehaviorReason
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) GetOriginalConstraints() []*GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints {
	return s.OriginalConstraints
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) GetType() *string {
	return s.Type
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) SetAllowedValues(v []*string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.AllowedValues = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) SetAssociationParameterNames(v []*string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.AssociationParameterNames = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) SetBehavior(v string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.Behavior = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) SetBehaviorReason(v string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.BehaviorReason = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) SetOriginalConstraints(v []*GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.OriginalConstraints = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) SetParameterKey(v string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.ParameterKey = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) SetType(v string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.Type = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) Validate() error {
	return dara.Validate(s)
}

type GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints struct {
	// The valid values of the parameter.
	AllowedValues []*string `json:"AllowedValues,omitempty" xml:"AllowedValues,omitempty" type:"Repeated"`
	// The property name.
	//
	// example:
	//
	// lnch_Source
	PropertyName *string `json:"PropertyName,omitempty" xml:"PropertyName,omitempty"`
	// The resource name.
	//
	// example:
	//
	// i-8vb0smn1lf6g77md****
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The resource type.
	//
	// example:
	//
	// serviceinstance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) String() string {
	return dara.Prettify(s)
}

func (s GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) GoString() string {
	return s.String()
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) GetAllowedValues() []*string {
	return s.AllowedValues
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) GetPropertyName() *string {
	return s.PropertyName
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) GetResourceName() *string {
	return s.ResourceName
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) SetAllowedValues(v []*string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints {
	s.AllowedValues = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) SetPropertyName(v string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints {
	s.PropertyName = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) SetResourceName(v string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints {
	s.ResourceName = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) SetResourceType(v string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints {
	s.ResourceType = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) Validate() error {
	return dara.Validate(s)
}
