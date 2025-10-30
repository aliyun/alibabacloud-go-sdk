// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFixNodePoolVulsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRestart(v bool) *FixNodePoolVulsRequest
	GetAutoRestart() *bool
	SetNodes(v []*string) *FixNodePoolVulsRequest
	GetNodes() []*string
	SetRolloutPolicy(v *FixNodePoolVulsRequestRolloutPolicy) *FixNodePoolVulsRequest
	GetRolloutPolicy() *FixNodePoolVulsRequestRolloutPolicy
	SetVuls(v []*string) *FixNodePoolVulsRequest
	GetVuls() []*string
}

type FixNodePoolVulsRequest struct {
	// Specifies whether to allow the nodes to restart.
	//
	// example:
	//
	// true
	AutoRestart *bool `json:"auto_restart,omitempty" xml:"auto_restart,omitempty"`
	// The names of the nodes to be patched.
	Nodes []*string `json:"nodes,omitempty" xml:"nodes,omitempty" type:"Repeated"`
	// The batch patching policy.
	RolloutPolicy *FixNodePoolVulsRequestRolloutPolicy `json:"rollout_policy,omitempty" xml:"rollout_policy,omitempty" type:"Struct"`
	// The list of vulnerabilities.
	Vuls []*string `json:"vuls,omitempty" xml:"vuls,omitempty" type:"Repeated"`
}

func (s FixNodePoolVulsRequest) String() string {
	return dara.Prettify(s)
}

func (s FixNodePoolVulsRequest) GoString() string {
	return s.String()
}

func (s *FixNodePoolVulsRequest) GetAutoRestart() *bool {
	return s.AutoRestart
}

func (s *FixNodePoolVulsRequest) GetNodes() []*string {
	return s.Nodes
}

func (s *FixNodePoolVulsRequest) GetRolloutPolicy() *FixNodePoolVulsRequestRolloutPolicy {
	return s.RolloutPolicy
}

func (s *FixNodePoolVulsRequest) GetVuls() []*string {
	return s.Vuls
}

func (s *FixNodePoolVulsRequest) SetAutoRestart(v bool) *FixNodePoolVulsRequest {
	s.AutoRestart = &v
	return s
}

func (s *FixNodePoolVulsRequest) SetNodes(v []*string) *FixNodePoolVulsRequest {
	s.Nodes = v
	return s
}

func (s *FixNodePoolVulsRequest) SetRolloutPolicy(v *FixNodePoolVulsRequestRolloutPolicy) *FixNodePoolVulsRequest {
	s.RolloutPolicy = v
	return s
}

func (s *FixNodePoolVulsRequest) SetVuls(v []*string) *FixNodePoolVulsRequest {
	s.Vuls = v
	return s
}

func (s *FixNodePoolVulsRequest) Validate() error {
	if s.RolloutPolicy != nil {
		if err := s.RolloutPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FixNodePoolVulsRequestRolloutPolicy struct {
	// The maximum concurrency for batch patching. Minimum value: 1. The maximum value equals the number of nodes in the node pool.
	//
	// example:
	//
	// 1
	MaxParallelism *int64 `json:"max_parallelism,omitempty" xml:"max_parallelism,omitempty"`
}

func (s FixNodePoolVulsRequestRolloutPolicy) String() string {
	return dara.Prettify(s)
}

func (s FixNodePoolVulsRequestRolloutPolicy) GoString() string {
	return s.String()
}

func (s *FixNodePoolVulsRequestRolloutPolicy) GetMaxParallelism() *int64 {
	return s.MaxParallelism
}

func (s *FixNodePoolVulsRequestRolloutPolicy) SetMaxParallelism(v int64) *FixNodePoolVulsRequestRolloutPolicy {
	s.MaxParallelism = &v
	return s
}

func (s *FixNodePoolVulsRequestRolloutPolicy) Validate() error {
	return dara.Validate(s)
}
