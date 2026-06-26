// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserPoolSyncJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSynchronizationJobId(v string) *GetUserPoolSyncJobRequest
	GetSynchronizationJobId() *string
	SetUserPoolName(v string) *GetUserPoolSyncJobRequest
	GetUserPoolName() *string
}

type GetUserPoolSyncJobRequest struct {
	SynchronizationJobId *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	UserPoolName         *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s GetUserPoolSyncJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserPoolSyncJobRequest) GoString() string {
	return s.String()
}

func (s *GetUserPoolSyncJobRequest) GetSynchronizationJobId() *string {
	return s.SynchronizationJobId
}

func (s *GetUserPoolSyncJobRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *GetUserPoolSyncJobRequest) SetSynchronizationJobId(v string) *GetUserPoolSyncJobRequest {
	s.SynchronizationJobId = &v
	return s
}

func (s *GetUserPoolSyncJobRequest) SetUserPoolName(v string) *GetUserPoolSyncJobRequest {
	s.UserPoolName = &v
	return s
}

func (s *GetUserPoolSyncJobRequest) Validate() error {
	return dara.Validate(s)
}
