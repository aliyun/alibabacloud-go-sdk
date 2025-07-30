// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSynchronizationJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetSynchronizationJobRequest
	GetInstanceId() *string
	SetSynchronizationJobId(v string) *GetSynchronizationJobRequest
	GetSynchronizationJobId() *string
}

type GetSynchronizationJobRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the synchronization job.
	//
	// This parameter is required.
	//
	// example:
	//
	// sync_0000347vjovtcf41li0fgsd98gn24q9njxxxxx
	SynchronizationJobId *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
}

func (s GetSynchronizationJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSynchronizationJobRequest) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSynchronizationJobRequest) GetSynchronizationJobId() *string {
	return s.SynchronizationJobId
}

func (s *GetSynchronizationJobRequest) SetInstanceId(v string) *GetSynchronizationJobRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSynchronizationJobRequest) SetSynchronizationJobId(v string) *GetSynchronizationJobRequest {
	s.SynchronizationJobId = &v
	return s
}

func (s *GetSynchronizationJobRequest) Validate() error {
	return dara.Validate(s)
}
