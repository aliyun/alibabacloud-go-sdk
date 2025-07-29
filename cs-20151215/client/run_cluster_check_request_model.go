// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunClusterCheckRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOptions(v map[string]*string) *RunClusterCheckRequest
	GetOptions() map[string]*string
	SetTarget(v string) *RunClusterCheckRequest
	GetTarget() *string
	SetType(v string) *RunClusterCheckRequest
	GetType() *string
}

type RunClusterCheckRequest struct {
	// The cluster check parameters.
	Options map[string]*string `json:"options,omitempty" xml:"options,omitempty"`
	// The check target.
	//
	// If you set `type=NodePoolUpgrade`, you must set this parameter to the node pool ID. Otherwise, this parameter is optional.
	//
	// example:
	//
	// np1f6779297c4444a3a1cdd29be8e5****
	Target *string `json:"target,omitempty" xml:"target,omitempty"`
	// The check type.
	//
	// Valid values:
	//
	// 	- ClusterMigrate: cluster migration.
	//
	// 	- MasterUpgrade: control plane upgrade.
	//
	// 	- NodePoolUpgrade: node pool upgrade.
	//
	// 	- ClusterUpgrade: cluster upgrade.
	//
	// This parameter is required.
	//
	// example:
	//
	// ClusterUpgrade
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s RunClusterCheckRequest) String() string {
	return dara.Prettify(s)
}

func (s RunClusterCheckRequest) GoString() string {
	return s.String()
}

func (s *RunClusterCheckRequest) GetOptions() map[string]*string {
	return s.Options
}

func (s *RunClusterCheckRequest) GetTarget() *string {
	return s.Target
}

func (s *RunClusterCheckRequest) GetType() *string {
	return s.Type
}

func (s *RunClusterCheckRequest) SetOptions(v map[string]*string) *RunClusterCheckRequest {
	s.Options = v
	return s
}

func (s *RunClusterCheckRequest) SetTarget(v string) *RunClusterCheckRequest {
	s.Target = &v
	return s
}

func (s *RunClusterCheckRequest) SetType(v string) *RunClusterCheckRequest {
	s.Type = &v
	return s
}

func (s *RunClusterCheckRequest) Validate() error {
	return dara.Validate(s)
}
