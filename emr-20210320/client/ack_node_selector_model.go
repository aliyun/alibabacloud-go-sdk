// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAckNodeSelector interface {
	dara.Model
	String() string
	GoString() string
	SetLabels(v []*AckNodeSelectorLabels) *AckNodeSelector
	GetLabels() []*AckNodeSelectorLabels
	SetTaints(v []*AckNodeSelectorTaints) *AckNodeSelector
	GetTaints() []*AckNodeSelectorTaints
}

type AckNodeSelector struct {
	// 污点列表。
	Labels []*AckNodeSelectorLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// 污点列表。
	Taints []*AckNodeSelectorTaints `json:"Taints,omitempty" xml:"Taints,omitempty" type:"Repeated"`
}

func (s AckNodeSelector) String() string {
	return dara.Prettify(s)
}

func (s AckNodeSelector) GoString() string {
	return s.String()
}

func (s *AckNodeSelector) GetLabels() []*AckNodeSelectorLabels {
	return s.Labels
}

func (s *AckNodeSelector) GetTaints() []*AckNodeSelectorTaints {
	return s.Taints
}

func (s *AckNodeSelector) SetLabels(v []*AckNodeSelectorLabels) *AckNodeSelector {
	s.Labels = v
	return s
}

func (s *AckNodeSelector) SetTaints(v []*AckNodeSelectorTaints) *AckNodeSelector {
	s.Taints = v
	return s
}

func (s *AckNodeSelector) Validate() error {
	return dara.Validate(s)
}

type AckNodeSelectorLabels struct {
	// 标签键。
	//
	// example:
	//
	// emr
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 标签值。
	//
	// example:
	//
	// true
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AckNodeSelectorLabels) String() string {
	return dara.Prettify(s)
}

func (s AckNodeSelectorLabels) GoString() string {
	return s.String()
}

func (s *AckNodeSelectorLabels) GetKey() *string {
	return s.Key
}

func (s *AckNodeSelectorLabels) GetValue() *string {
	return s.Value
}

func (s *AckNodeSelectorLabels) SetKey(v string) *AckNodeSelectorLabels {
	s.Key = &v
	return s
}

func (s *AckNodeSelectorLabels) SetValue(v string) *AckNodeSelectorLabels {
	s.Value = &v
	return s
}

func (s *AckNodeSelectorLabels) Validate() error {
	return dara.Validate(s)
}

type AckNodeSelectorTaints struct {
	// 污点效果。
	Effect *string `json:"Effect,omitempty" xml:"Effect,omitempty"`
	// 污点键。
	//
	// example:
	//
	// emr
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 污点值。
	//
	// example:
	//
	// true
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AckNodeSelectorTaints) String() string {
	return dara.Prettify(s)
}

func (s AckNodeSelectorTaints) GoString() string {
	return s.String()
}

func (s *AckNodeSelectorTaints) GetEffect() *string {
	return s.Effect
}

func (s *AckNodeSelectorTaints) GetKey() *string {
	return s.Key
}

func (s *AckNodeSelectorTaints) GetValue() *string {
	return s.Value
}

func (s *AckNodeSelectorTaints) SetEffect(v string) *AckNodeSelectorTaints {
	s.Effect = &v
	return s
}

func (s *AckNodeSelectorTaints) SetKey(v string) *AckNodeSelectorTaints {
	s.Key = &v
	return s
}

func (s *AckNodeSelectorTaints) SetValue(v string) *AckNodeSelectorTaints {
	s.Value = &v
	return s
}

func (s *AckNodeSelectorTaints) Validate() error {
	return dara.Validate(s)
}
