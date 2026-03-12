// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceConfigPreCheckRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyInstanceConfigPreCheckRequest
	GetInstanceId() *string
	SetConfigsToAdd(v []*InstanceConfigDto) *ModifyInstanceConfigPreCheckRequest
	GetConfigsToAdd() []*InstanceConfigDto
	SetConfigsToDelete(v []*InstanceConfigDto) *ModifyInstanceConfigPreCheckRequest
	GetConfigsToDelete() []*InstanceConfigDto
	SetConfigsToUpdate(v []*InstanceConfigDto) *ModifyInstanceConfigPreCheckRequest
	GetConfigsToUpdate() []*InstanceConfigDto
}

type ModifyInstanceConfigPreCheckRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId      *string              `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ConfigsToAdd    []*InstanceConfigDto `json:"configsToAdd,omitempty" xml:"configsToAdd,omitempty" type:"Repeated"`
	ConfigsToDelete []*InstanceConfigDto `json:"configsToDelete,omitempty" xml:"configsToDelete,omitempty" type:"Repeated"`
	ConfigsToUpdate []*InstanceConfigDto `json:"configsToUpdate,omitempty" xml:"configsToUpdate,omitempty" type:"Repeated"`
}

func (s ModifyInstanceConfigPreCheckRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceConfigPreCheckRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceConfigPreCheckRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceConfigPreCheckRequest) GetConfigsToAdd() []*InstanceConfigDto {
	return s.ConfigsToAdd
}

func (s *ModifyInstanceConfigPreCheckRequest) GetConfigsToDelete() []*InstanceConfigDto {
	return s.ConfigsToDelete
}

func (s *ModifyInstanceConfigPreCheckRequest) GetConfigsToUpdate() []*InstanceConfigDto {
	return s.ConfigsToUpdate
}

func (s *ModifyInstanceConfigPreCheckRequest) SetInstanceId(v string) *ModifyInstanceConfigPreCheckRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceConfigPreCheckRequest) SetConfigsToAdd(v []*InstanceConfigDto) *ModifyInstanceConfigPreCheckRequest {
	s.ConfigsToAdd = v
	return s
}

func (s *ModifyInstanceConfigPreCheckRequest) SetConfigsToDelete(v []*InstanceConfigDto) *ModifyInstanceConfigPreCheckRequest {
	s.ConfigsToDelete = v
	return s
}

func (s *ModifyInstanceConfigPreCheckRequest) SetConfigsToUpdate(v []*InstanceConfigDto) *ModifyInstanceConfigPreCheckRequest {
	s.ConfigsToUpdate = v
	return s
}

func (s *ModifyInstanceConfigPreCheckRequest) Validate() error {
	if s.ConfigsToAdd != nil {
		for _, item := range s.ConfigsToAdd {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ConfigsToDelete != nil {
		for _, item := range s.ConfigsToDelete {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ConfigsToUpdate != nil {
		for _, item := range s.ConfigsToUpdate {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
