// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagSelector interface {
	dara.Model
	String() string
	GoString() string
	SetConditions(v []*TagCondition) *TagSelector
	GetConditions() []*TagCondition
	SetExpression(v string) *TagSelector
	GetExpression() *string
	SetRelation(v string) *TagSelector
	GetRelation() *string
}

type TagSelector struct {
	Conditions []*TagCondition `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	Expression *string         `json:"expression,omitempty" xml:"expression,omitempty"`
	Relation   *string         `json:"relation,omitempty" xml:"relation,omitempty"`
}

func (s TagSelector) String() string {
	return dara.Prettify(s)
}

func (s TagSelector) GoString() string {
	return s.String()
}

func (s *TagSelector) GetConditions() []*TagCondition {
	return s.Conditions
}

func (s *TagSelector) GetExpression() *string {
	return s.Expression
}

func (s *TagSelector) GetRelation() *string {
	return s.Relation
}

func (s *TagSelector) SetConditions(v []*TagCondition) *TagSelector {
	s.Conditions = v
	return s
}

func (s *TagSelector) SetExpression(v string) *TagSelector {
	s.Expression = &v
	return s
}

func (s *TagSelector) SetRelation(v string) *TagSelector {
	s.Relation = &v
	return s
}

func (s *TagSelector) Validate() error {
	if s.Conditions != nil {
		for _, item := range s.Conditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
