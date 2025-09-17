// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartReplicationJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *StartReplicationJobRequest
	GetJobId() *string
	SetOwnerId(v int64) *StartReplicationJobRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *StartReplicationJobRequest
	GetResourceOwnerAccount() *string
}

type StartReplicationJobRequest struct {
	// The migration job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// j-bw526m1vi6x21q****
	JobId                *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s StartReplicationJobRequest) String() string {
	return dara.Prettify(s)
}

func (s StartReplicationJobRequest) GoString() string {
	return s.String()
}

func (s *StartReplicationJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *StartReplicationJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StartReplicationJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *StartReplicationJobRequest) SetJobId(v string) *StartReplicationJobRequest {
	s.JobId = &v
	return s
}

func (s *StartReplicationJobRequest) SetOwnerId(v int64) *StartReplicationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *StartReplicationJobRequest) SetResourceOwnerAccount(v string) *StartReplicationJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StartReplicationJobRequest) Validate() error {
	return dara.Validate(s)
}
