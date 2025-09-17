// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteReplicationJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *DeleteReplicationJobRequest
	GetJobId() *string
	SetOwnerId(v int64) *DeleteReplicationJobRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteReplicationJobRequest
	GetResourceOwnerAccount() *string
}

type DeleteReplicationJobRequest struct {
	// The migration job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// j-bp17m1vi6x21qhqk****
	JobId                *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s DeleteReplicationJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteReplicationJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteReplicationJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *DeleteReplicationJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteReplicationJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteReplicationJobRequest) SetJobId(v string) *DeleteReplicationJobRequest {
	s.JobId = &v
	return s
}

func (s *DeleteReplicationJobRequest) SetOwnerId(v int64) *DeleteReplicationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteReplicationJobRequest) SetResourceOwnerAccount(v string) *DeleteReplicationJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteReplicationJobRequest) Validate() error {
	return dara.Validate(s)
}
