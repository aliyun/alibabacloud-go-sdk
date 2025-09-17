// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCutOverReplicationJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *CutOverReplicationJobRequest
	GetJobId() *string
	SetOwnerId(v int64) *CutOverReplicationJobRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CutOverReplicationJobRequest
	GetResourceOwnerAccount() *string
	SetSyncData(v bool) *CutOverReplicationJobRequest
	GetSyncData() *bool
}

type CutOverReplicationJobRequest struct {
	// The ID of the incremental migration job.
	//
	// This parameter is required.
	//
	// example:
	//
	// j-bp1fnx5y3djc4cop****
	JobId                *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// Specifies whether to migrate full data for the last time. Valid Values:
	//
	// 	- true: migrates full data for the last time.
	//
	// 	- false: does not migrate full data for the last time.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	SyncData *bool `json:"SyncData,omitempty" xml:"SyncData,omitempty"`
}

func (s CutOverReplicationJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CutOverReplicationJobRequest) GoString() string {
	return s.String()
}

func (s *CutOverReplicationJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *CutOverReplicationJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CutOverReplicationJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CutOverReplicationJobRequest) GetSyncData() *bool {
	return s.SyncData
}

func (s *CutOverReplicationJobRequest) SetJobId(v string) *CutOverReplicationJobRequest {
	s.JobId = &v
	return s
}

func (s *CutOverReplicationJobRequest) SetOwnerId(v int64) *CutOverReplicationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *CutOverReplicationJobRequest) SetResourceOwnerAccount(v string) *CutOverReplicationJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CutOverReplicationJobRequest) SetSyncData(v bool) *CutOverReplicationJobRequest {
	s.SyncData = &v
	return s
}

func (s *CutOverReplicationJobRequest) Validate() error {
	return dara.Validate(s)
}
