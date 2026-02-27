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
	// 	- `NoSchedule`: This taint is not tolerated. However, pods that are already scheduled to the node are not affected.
	//
	// 	- `NoExecute`: Pods that do not tolerate this taint are evicted after this taint is added to the node.
	//
	// 	- `PreferNoSchedule`: This value specifies a soft limit on pods. Existing pods on the node are not affected. The scheduler attempts to avoid scheduling pods that cannot tolerate the taint to the node.
	//
	// Default value: `NoSchedule`.
	//
	// example:
	//
	// NoSchedule
	Effect *string `json:"effect,omitempty" xml:"effect,omitempty"`
	// The `key` of the taint.
	//
	// example:
	//
	// key
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The `value` of the taint.
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
