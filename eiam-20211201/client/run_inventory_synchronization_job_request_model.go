// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunInventorySynchronizationJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *RunInventorySynchronizationJobRequest
	GetInstanceId() *string
	SetInventorySourceId(v string) *RunInventorySynchronizationJobRequest
	GetInventorySourceId() *string
	SetInventoryTypes(v []*string) *RunInventorySynchronizationJobRequest
	GetInventoryTypes() []*string
}

type RunInventorySynchronizationJobRequest struct {
	// The ID of the IDaaS EIAM instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the synchronization target.
	//
	// This parameter is required.
	//
	// example:
	//
	// sync_xxxxxx
	InventorySourceId *string `json:"InventorySourceId,omitempty" xml:"InventorySourceId,omitempty"`
	// The resource types.
	InventoryTypes []*string `json:"InventoryTypes,omitempty" xml:"InventoryTypes,omitempty" type:"Repeated"`
}

func (s RunInventorySynchronizationJobRequest) String() string {
	return dara.Prettify(s)
}

func (s RunInventorySynchronizationJobRequest) GoString() string {
	return s.String()
}

func (s *RunInventorySynchronizationJobRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RunInventorySynchronizationJobRequest) GetInventorySourceId() *string {
	return s.InventorySourceId
}

func (s *RunInventorySynchronizationJobRequest) GetInventoryTypes() []*string {
	return s.InventoryTypes
}

func (s *RunInventorySynchronizationJobRequest) SetInstanceId(v string) *RunInventorySynchronizationJobRequest {
	s.InstanceId = &v
	return s
}

func (s *RunInventorySynchronizationJobRequest) SetInventorySourceId(v string) *RunInventorySynchronizationJobRequest {
	s.InventorySourceId = &v
	return s
}

func (s *RunInventorySynchronizationJobRequest) SetInventoryTypes(v []*string) *RunInventorySynchronizationJobRequest {
	s.InventoryTypes = v
	return s
}

func (s *RunInventorySynchronizationJobRequest) Validate() error {
	return dara.Validate(s)
}
