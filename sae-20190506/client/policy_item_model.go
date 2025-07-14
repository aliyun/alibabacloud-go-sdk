// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPolicyItem interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *PolicyItem
	GetKey() *string
	SetOperator(v string) *PolicyItem
	GetOperator() *string
	SetType(v string) *PolicyItem
	GetType() *string
	SetValue(v string) *PolicyItem
	GetValue() *string
}

type PolicyItem struct {
	Key      *string `json:"key,omitempty" xml:"key,omitempty"`
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	Type     *string `json:"type,omitempty" xml:"type,omitempty"`
	Value    *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s PolicyItem) String() string {
	return dara.Prettify(s)
}

func (s PolicyItem) GoString() string {
	return s.String()
}

func (s *PolicyItem) GetKey() *string {
	return s.Key
}

func (s *PolicyItem) GetOperator() *string {
	return s.Operator
}

func (s *PolicyItem) GetType() *string {
	return s.Type
}

func (s *PolicyItem) GetValue() *string {
	return s.Value
}

func (s *PolicyItem) SetKey(v string) *PolicyItem {
	s.Key = &v
	return s
}

func (s *PolicyItem) SetOperator(v string) *PolicyItem {
	s.Operator = &v
	return s
}

func (s *PolicyItem) SetType(v string) *PolicyItem {
	s.Type = &v
	return s
}

func (s *PolicyItem) SetValue(v string) *PolicyItem {
	s.Value = &v
	return s
}

func (s *PolicyItem) Validate() error {
	return dara.Validate(s)
}
