// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInventorySynchronizationJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetInventorySynchronizationJobRequest
	GetInstanceId() *string
	SetInventorySynchronizationJobId(v string) *GetInventorySynchronizationJobRequest
	GetInventorySynchronizationJobId() *string
}

type GetInventorySynchronizationJobRequest struct {
	// The ID of the IDaaS EIAM instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the synchronization task.
	//
	// This parameter is required.
	//
	// example:
	//
	// sync_xxxxxxx
	InventorySynchronizationJobId *string `json:"InventorySynchronizationJobId,omitempty" xml:"InventorySynchronizationJobId,omitempty"`
}

func (s GetInventorySynchronizationJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInventorySynchronizationJobRequest) GoString() string {
	return s.String()
}

func (s *GetInventorySynchronizationJobRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInventorySynchronizationJobRequest) GetInventorySynchronizationJobId() *string {
	return s.InventorySynchronizationJobId
}

func (s *GetInventorySynchronizationJobRequest) SetInstanceId(v string) *GetInventorySynchronizationJobRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInventorySynchronizationJobRequest) SetInventorySynchronizationJobId(v string) *GetInventorySynchronizationJobRequest {
	s.InventorySynchronizationJobId = &v
	return s
}

func (s *GetInventorySynchronizationJobRequest) Validate() error {
	return dara.Validate(s)
}
