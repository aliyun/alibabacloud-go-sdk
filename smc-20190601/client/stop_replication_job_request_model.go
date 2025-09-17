// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopReplicationJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *StopReplicationJobRequest
	GetJobId() *string
	SetOwnerId(v int64) *StopReplicationJobRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *StopReplicationJobRequest
	GetResourceOwnerAccount() *string
}

type StopReplicationJobRequest struct {
	// The migration job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// j-bw526m1vi6x21qh****
	JobId                *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s StopReplicationJobRequest) String() string {
	return dara.Prettify(s)
}

func (s StopReplicationJobRequest) GoString() string {
	return s.String()
}

func (s *StopReplicationJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *StopReplicationJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StopReplicationJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *StopReplicationJobRequest) SetJobId(v string) *StopReplicationJobRequest {
	s.JobId = &v
	return s
}

func (s *StopReplicationJobRequest) SetOwnerId(v int64) *StopReplicationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *StopReplicationJobRequest) SetResourceOwnerAccount(v string) *StopReplicationJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StopReplicationJobRequest) Validate() error {
	return dara.Validate(s)
}
