// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterOnlineConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusters(v []*string) *ModifyClusterOnlineConfigRequest
	GetClusters() []*string
	SetConfig(v map[string]*int32) *ModifyClusterOnlineConfigRequest
	GetConfig() map[string]*int32
}

type ModifyClusterOnlineConfigRequest struct {
	// The cluster information.
	Clusters []*string `json:"clusters,omitempty" xml:"clusters,omitempty" type:"Repeated"`
	// The configuration information.
	Config map[string]*int32 `json:"config,omitempty" xml:"config,omitempty"`
}

func (s ModifyClusterOnlineConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterOnlineConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterOnlineConfigRequest) GetClusters() []*string {
	return s.Clusters
}

func (s *ModifyClusterOnlineConfigRequest) GetConfig() map[string]*int32 {
	return s.Config
}

func (s *ModifyClusterOnlineConfigRequest) SetClusters(v []*string) *ModifyClusterOnlineConfigRequest {
	s.Clusters = v
	return s
}

func (s *ModifyClusterOnlineConfigRequest) SetConfig(v map[string]*int32) *ModifyClusterOnlineConfigRequest {
	s.Config = v
	return s
}

func (s *ModifyClusterOnlineConfigRequest) Validate() error {
	return dara.Validate(s)
}
