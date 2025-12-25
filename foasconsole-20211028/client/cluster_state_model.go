// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClusterState interface {
	dara.Model
	String() string
	GoString() string
	SetClusterStage(v *ClusterStage) *ClusterState
	GetClusterStage() *ClusterStage
	SetStatus(v string) *ClusterState
	GetStatus() *string
	SetSubStatus(v string) *ClusterState
	GetSubStatus() *string
}

type ClusterState struct {
	ClusterStage *ClusterStage `json:"ClusterStage,omitempty" xml:"ClusterStage,omitempty"`
	Status       *string       `json:"Status,omitempty" xml:"Status,omitempty"`
	SubStatus    *string       `json:"SubStatus,omitempty" xml:"SubStatus,omitempty"`
}

func (s ClusterState) String() string {
	return dara.Prettify(s)
}

func (s ClusterState) GoString() string {
	return s.String()
}

func (s *ClusterState) GetClusterStage() *ClusterStage {
	return s.ClusterStage
}

func (s *ClusterState) GetStatus() *string {
	return s.Status
}

func (s *ClusterState) GetSubStatus() *string {
	return s.SubStatus
}

func (s *ClusterState) SetClusterStage(v *ClusterStage) *ClusterState {
	s.ClusterStage = v
	return s
}

func (s *ClusterState) SetStatus(v string) *ClusterState {
	s.Status = &v
	return s
}

func (s *ClusterState) SetSubStatus(v string) *ClusterState {
	s.SubStatus = &v
	return s
}

func (s *ClusterState) Validate() error {
	if s.ClusterStage != nil {
		if err := s.ClusterStage.Validate(); err != nil {
			return err
		}
	}
	return nil
}
