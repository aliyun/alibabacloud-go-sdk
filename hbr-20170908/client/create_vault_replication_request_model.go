// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVaultReplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReplicationSourceRegionId(v string) *CreateVaultReplicationRequest
	GetReplicationSourceRegionId() *string
	SetReplicationSourceVaultId(v string) *CreateVaultReplicationRequest
	GetReplicationSourceVaultId() *string
	SetReplicationTargetVaultId(v string) *CreateVaultReplicationRequest
	GetReplicationTargetVaultId() *string
}

type CreateVaultReplicationRequest struct {
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

func (s CreateVaultReplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVaultReplicationRequest) GoString() string {
	return s.String()
}

func (s *CreateVaultReplicationRequest) GetReplicationSourceRegionId() *string {
	return s.ReplicationSourceRegionId
}

func (s *CreateVaultReplicationRequest) GetReplicationSourceVaultId() *string {
	return s.ReplicationSourceVaultId
}

func (s *CreateVaultReplicationRequest) GetReplicationTargetVaultId() *string {
	return s.ReplicationTargetVaultId
}

func (s *CreateVaultReplicationRequest) SetReplicationSourceRegionId(v string) *CreateVaultReplicationRequest {
	s.ReplicationSourceRegionId = &v
	return s
}

func (s *CreateVaultReplicationRequest) SetReplicationSourceVaultId(v string) *CreateVaultReplicationRequest {
	s.ReplicationSourceVaultId = &v
	return s
}

func (s *CreateVaultReplicationRequest) SetReplicationTargetVaultId(v string) *CreateVaultReplicationRequest {
	s.ReplicationTargetVaultId = &v
	return s
}

func (s *CreateVaultReplicationRequest) Validate() error {
	return dara.Validate(s)
}
