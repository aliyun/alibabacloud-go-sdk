// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaint interface {
	dara.Model
	String() string
	GoString() string
	SetEffect(v string) *Taint
	GetEffect() *string
	SetKey(v string) *Taint
	GetKey() *string
	SetValue(v string) *Taint
	GetValue() *string
}

type Taint struct {
	// The scheduling policy. Valid values:
	//
	// - `NoSchedule`: Not tolerated. This affects only the scheduling process. Already scheduled pods are not affected. This applies only to newly added pods.
	//
	// - `NoExecute`: Not tolerated. When the taint changes, the pod is evicted.
	//
	// - `PreferNoSchedule`: Soft constraint. Existing pods on the node are not affected. The scheduler tries to avoid scheduling pods that do not tolerate the taint to the node.
	//
	// Default value: `NoSchedule`.
	//
	// example:
	//
	// NoSchedule
	Effect *string `json:"effect,omitempty" xml:"effect,omitempty"`
	// The key of the taint.
	//
	// example:
	//
	// key
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The value of the taint.
	//
	// example:
	//
	// value
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s Taint) String() string {
	return dara.Prettify(s)
}

func (s Taint) GoString() string {
	return s.String()
}

func (s *Taint) GetEffect() *string {
	return s.Effect
}

func (s *Taint) GetKey() *string {
	return s.Key
}

func (s *Taint) GetValue() *string {
	return s.Value
}

func (s *Taint) SetEffect(v string) *Taint {
	s.Effect = &v
	return s
}

func (s *Taint) SetKey(v string) *Taint {
	s.Key = &v
	return s
}

func (s *Taint) SetValue(v string) *Taint {
	s.Value = &v
	return s
}

func (s *Taint) Validate() error {
	return dara.Validate(s)
}
