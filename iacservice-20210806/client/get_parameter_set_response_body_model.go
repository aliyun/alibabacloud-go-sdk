// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetParameterSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetParameterSet(v *GetParameterSetResponseBodyParameterSet) *GetParameterSetResponseBody
	GetParameterSet() *GetParameterSetResponseBodyParameterSet
	SetRequestId(v string) *GetParameterSetResponseBody
	GetRequestId() *string
}

type GetParameterSetResponseBody struct {
	// The parameter set details.
	ParameterSet *GetParameterSetResponseBodyParameterSet `json:"parameterSet,omitempty" xml:"parameterSet,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 99905C7C-1320-5E7F-A798-3071482EB08E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetParameterSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetParameterSetResponseBody) GoString() string {
	return s.String()
}

func (s *GetParameterSetResponseBody) GetParameterSet() *GetParameterSetResponseBodyParameterSet {
	return s.ParameterSet
}

func (s *GetParameterSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetParameterSetResponseBody) SetParameterSet(v *GetParameterSetResponseBodyParameterSet) *GetParameterSetResponseBody {
	s.ParameterSet = v
	return s
}

func (s *GetParameterSetResponseBody) SetRequestId(v string) *GetParameterSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetParameterSetResponseBody) Validate() error {
	if s.ParameterSet != nil {
		if err := s.ParameterSet.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetParameterSetResponseBodyParameterSet struct {
	// The time when the parameter set was created.
	//
	// example:
	//
	// 2022-01-30T02:14:16Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The description of the parameter set.
	//
	// example:
	//
	// This is parameterSet
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The name of the parameter set.
	//
	// example:
	//
	// demo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ID of the parameter set.
	//
	// example:
	//
	// pts-3b6cb9fa4751afff9c5e4e01624b9
	ParameterSetId *string `json:"parameterSetId,omitempty" xml:"parameterSetId,omitempty"`
	// The parameters in the parameter set.
	Parameters []*GetParameterSetResponseBodyParameterSetParameters `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Repeated"`
	// The association relationships.
	RelationList []*GetParameterSetResponseBodyParameterSetRelationList `json:"relationList,omitempty" xml:"relationList,omitempty" type:"Repeated"`
}

func (s GetParameterSetResponseBodyParameterSet) String() string {
	return dara.Prettify(s)
}

func (s GetParameterSetResponseBodyParameterSet) GoString() string {
	return s.String()
}

func (s *GetParameterSetResponseBodyParameterSet) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetParameterSetResponseBodyParameterSet) GetDescription() *string {
	return s.Description
}

func (s *GetParameterSetResponseBodyParameterSet) GetName() *string {
	return s.Name
}

func (s *GetParameterSetResponseBodyParameterSet) GetParameterSetId() *string {
	return s.ParameterSetId
}

func (s *GetParameterSetResponseBodyParameterSet) GetParameters() []*GetParameterSetResponseBodyParameterSetParameters {
	return s.Parameters
}

func (s *GetParameterSetResponseBodyParameterSet) GetRelationList() []*GetParameterSetResponseBodyParameterSetRelationList {
	return s.RelationList
}

func (s *GetParameterSetResponseBodyParameterSet) SetCreateTime(v string) *GetParameterSetResponseBodyParameterSet {
	s.CreateTime = &v
	return s
}

func (s *GetParameterSetResponseBodyParameterSet) SetDescription(v string) *GetParameterSetResponseBodyParameterSet {
	s.Description = &v
	return s
}

func (s *GetParameterSetResponseBodyParameterSet) SetName(v string) *GetParameterSetResponseBodyParameterSet {
	s.Name = &v
	return s
}

func (s *GetParameterSetResponseBodyParameterSet) SetParameterSetId(v string) *GetParameterSetResponseBodyParameterSet {
	s.ParameterSetId = &v
	return s
}

func (s *GetParameterSetResponseBodyParameterSet) SetParameters(v []*GetParameterSetResponseBodyParameterSetParameters) *GetParameterSetResponseBodyParameterSet {
	s.Parameters = v
	return s
}

func (s *GetParameterSetResponseBodyParameterSet) SetRelationList(v []*GetParameterSetResponseBodyParameterSetRelationList) *GetParameterSetResponseBodyParameterSet {
	s.RelationList = v
	return s
}

func (s *GetParameterSetResponseBodyParameterSet) Validate() error {
	if s.Parameters != nil {
		for _, item := range s.Parameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RelationList != nil {
		for _, item := range s.RelationList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetParameterSetResponseBodyParameterSetParameters struct {
	// The name of the parameter.
	//
	// example:
	//
	// region
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	Secret *bool   `json:"secret,omitempty" xml:"secret,omitempty"`
	// The status of the parameter. Valid values:
	//
	// - HAS_VALUE (default): A specific value is defined.
	//
	// - EXPLICIT_NULL: The value is explicitly set to null.
	//
	// example:
	//
	// HAS_VALUE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The type of the parameter. Valid values: string, number, bool, map(string), and list(string).
	//
	// example:
	//
	// string
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The value of the parameter.
	//
	// example:
	//
	// cn-hangzhou
	Value interface{} `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetParameterSetResponseBodyParameterSetParameters) String() string {
	return dara.Prettify(s)
}

func (s GetParameterSetResponseBodyParameterSetParameters) GoString() string {
	return s.String()
}

func (s *GetParameterSetResponseBodyParameterSetParameters) GetName() *string {
	return s.Name
}

func (s *GetParameterSetResponseBodyParameterSetParameters) GetSecret() *bool {
	return s.Secret
}

func (s *GetParameterSetResponseBodyParameterSetParameters) GetStatus() *string {
	return s.Status
}

func (s *GetParameterSetResponseBodyParameterSetParameters) GetType() *string {
	return s.Type
}

func (s *GetParameterSetResponseBodyParameterSetParameters) GetValue() interface{} {
	return s.Value
}

func (s *GetParameterSetResponseBodyParameterSetParameters) SetName(v string) *GetParameterSetResponseBodyParameterSetParameters {
	s.Name = &v
	return s
}

func (s *GetParameterSetResponseBodyParameterSetParameters) SetSecret(v bool) *GetParameterSetResponseBodyParameterSetParameters {
	s.Secret = &v
	return s
}

func (s *GetParameterSetResponseBodyParameterSetParameters) SetStatus(v string) *GetParameterSetResponseBodyParameterSetParameters {
	s.Status = &v
	return s
}

func (s *GetParameterSetResponseBodyParameterSetParameters) SetType(v string) *GetParameterSetResponseBodyParameterSetParameters {
	s.Type = &v
	return s
}

func (s *GetParameterSetResponseBodyParameterSetParameters) SetValue(v interface{}) *GetParameterSetResponseBodyParameterSetParameters {
	s.Value = v
	return s
}

func (s *GetParameterSetResponseBodyParameterSetParameters) Validate() error {
	return dara.Validate(s)
}

type GetParameterSetResponseBodyParameterSetRelationList struct {
	// The time when the parameter set was created.
	//
	// example:
	//
	// 2022-04-24T22:58:50Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The resource ID. When the resource type is ModuleVersion, the ID is a concatenation of <moduleId>-<moduleversion>, such as mod-34535345df123fr-v3.
	//
	// example:
	//
	// mod-433aead756057101546eb5d50c1
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The resource type. Valid values:
	//
	// - Module: template
	//
	// - ModuleVersion: template version
	//
	// - Task: task.
	//
	// example:
	//
	// Module
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s GetParameterSetResponseBodyParameterSetRelationList) String() string {
	return dara.Prettify(s)
}

func (s GetParameterSetResponseBodyParameterSetRelationList) GoString() string {
	return s.String()
}

func (s *GetParameterSetResponseBodyParameterSetRelationList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetParameterSetResponseBodyParameterSetRelationList) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetParameterSetResponseBodyParameterSetRelationList) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetParameterSetResponseBodyParameterSetRelationList) SetCreateTime(v string) *GetParameterSetResponseBodyParameterSetRelationList {
	s.CreateTime = &v
	return s
}

func (s *GetParameterSetResponseBodyParameterSetRelationList) SetResourceId(v string) *GetParameterSetResponseBodyParameterSetRelationList {
	s.ResourceId = &v
	return s
}

func (s *GetParameterSetResponseBodyParameterSetRelationList) SetResourceType(v string) *GetParameterSetResponseBodyParameterSetRelationList {
	s.ResourceType = &v
	return s
}

func (s *GetParameterSetResponseBodyParameterSetRelationList) Validate() error {
	return dara.Validate(s)
}
