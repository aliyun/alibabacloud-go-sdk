// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVaultReplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReplicationSourceRegionId(v string) *DeleteVaultReplicationRequest
	GetReplicationSourceRegionId() *string
	SetReplicationSourceVaultId(v string) *DeleteVaultReplicationRequest
	GetReplicationSourceVaultId() *string
	SetReplicationTargetVaultId(v string) *DeleteVaultReplicationRequest
	GetReplicationTargetVaultId() *string
}

type DeleteVaultReplicationRequest struct {
	// example:
	//
	// cn-hangzhou
	ReplicationSourceRegionId *string `json:"ReplicationSourceRegionId,omitempty" xml:"ReplicationSourceRegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v-*********************
	ReplicationSourceVaultId *string `json:"ReplicationSourceVaultId,omitempty" xml:"ReplicationSourceVaultId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// r-*********************
	ReplicationTargetVaultId *string `json:"ReplicationTargetVaultId,omitempty" xml:"ReplicationTargetVaultId,omitempty"`
}

func (s DeleteVaultReplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVaultReplicationRequest) GoString() string {
	return s.String()
}

func (s *DeleteVaultReplicationRequest) GetReplicationSourceRegionId() *string {
	return s.ReplicationSourceRegionId
}

func (s *DeleteVaultReplicationRequest) GetReplicationSourceVaultId() *string {
	return s.ReplicationSourceVaultId
}

func (s *DeleteVaultReplicationRequest) GetReplicationTargetVaultId() *string {
	return s.ReplicationTargetVaultId
}

func (s *DeleteVaultReplicationRequest) SetReplicationSourceRegionId(v string) *DeleteVaultReplicationRequest {
	s.ReplicationSourceRegionId = &v
	return s
}

func (s *DeleteVaultReplicationRequest) SetReplicationSourceVaultId(v string) *DeleteVaultReplicationRequest {
	s.ReplicationSourceVaultId = &v
	return s
}

func (s *DeleteVaultReplicationRequest) SetReplicationTargetVaultId(v string) *DeleteVaultReplicationRequest {
	s.ReplicationTargetVaultId = &v
	return s
}

func (s *DeleteVaultReplicationRequest) Validate() error {
	return dara.Validate(s)
}
