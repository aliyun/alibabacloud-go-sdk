// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddConfigList(v string) *ModifyInstanceConfigRequest
	GetAddConfigList() *string
	SetConfigList(v string) *ModifyInstanceConfigRequest
	GetConfigList() *string
	SetDeleteConfigList(v string) *ModifyInstanceConfigRequest
	GetDeleteConfigList() *string
	SetInstanceId(v string) *ModifyInstanceConfigRequest
	GetInstanceId() *string
	SetReason(v string) *ModifyInstanceConfigRequest
	GetReason() *string
	SetConfigsToAdd(v []*InstanceConfigDto) *ModifyInstanceConfigRequest
	GetConfigsToAdd() []*InstanceConfigDto
	SetConfigsToDelete(v []*InstanceConfigDto) *ModifyInstanceConfigRequest
	GetConfigsToDelete() []*InstanceConfigDto
	SetConfigsToUpdate(v []*InstanceConfigDto) *ModifyInstanceConfigRequest
	GetConfigsToUpdate() []*InstanceConfigDto
	SetFastMode(v bool) *ModifyInstanceConfigRequest
	GetFastMode() *bool
	SetRestart(v bool) *ModifyInstanceConfigRequest
	GetRestart() *bool
}

type ModifyInstanceConfigRequest struct {
	// example:
	//
	// []
	AddConfigList *string `json:"AddConfigList,omitempty" xml:"AddConfigList,omitempty"`
	// example:
	//
	// [{\\"type\\":\\"sas_analysis_online-sas-operation-log-sh-sas-event-rasp\\",\\"configItemList\\":[{\\"key\\":\\"item_level\\",\\"valueList\\":[\\"all\\"]},{\\"key\\":\\"alert_type\\",\\"valueList\\":[\\"all\\"]}]}]
	ConfigList *string `json:"ConfigList,omitempty" xml:"ConfigList,omitempty"`
	// example:
	//
	// []
	DeleteConfigList *string `json:"DeleteConfigList,omitempty" xml:"DeleteConfigList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c-991ca6180620****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// test
	Reason          *string              `json:"Reason,omitempty" xml:"Reason,omitempty"`
	ConfigsToAdd    []*InstanceConfigDto `json:"configsToAdd,omitempty" xml:"configsToAdd,omitempty" type:"Repeated"`
	ConfigsToDelete []*InstanceConfigDto `json:"configsToDelete,omitempty" xml:"configsToDelete,omitempty" type:"Repeated"`
	ConfigsToUpdate []*InstanceConfigDto `json:"configsToUpdate,omitempty" xml:"configsToUpdate,omitempty" type:"Repeated"`
	// example:
	//
	// true
	FastMode *bool `json:"fastMode,omitempty" xml:"fastMode,omitempty"`
	// example:
	//
	// true
	Restart *bool `json:"restart,omitempty" xml:"restart,omitempty"`
}

func (s ModifyInstanceConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceConfigRequest) GetAddConfigList() *string {
	return s.AddConfigList
}

func (s *ModifyInstanceConfigRequest) GetConfigList() *string {
	return s.ConfigList
}

func (s *ModifyInstanceConfigRequest) GetDeleteConfigList() *string {
	return s.DeleteConfigList
}

func (s *ModifyInstanceConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceConfigRequest) GetReason() *string {
	return s.Reason
}

func (s *ModifyInstanceConfigRequest) GetConfigsToAdd() []*InstanceConfigDto {
	return s.ConfigsToAdd
}

func (s *ModifyInstanceConfigRequest) GetConfigsToDelete() []*InstanceConfigDto {
	return s.ConfigsToDelete
}

func (s *ModifyInstanceConfigRequest) GetConfigsToUpdate() []*InstanceConfigDto {
	return s.ConfigsToUpdate
}

func (s *ModifyInstanceConfigRequest) GetFastMode() *bool {
	return s.FastMode
}

func (s *ModifyInstanceConfigRequest) GetRestart() *bool {
	return s.Restart
}

func (s *ModifyInstanceConfigRequest) SetAddConfigList(v string) *ModifyInstanceConfigRequest {
	s.AddConfigList = &v
	return s
}

func (s *ModifyInstanceConfigRequest) SetConfigList(v string) *ModifyInstanceConfigRequest {
	s.ConfigList = &v
	return s
}

func (s *ModifyInstanceConfigRequest) SetDeleteConfigList(v string) *ModifyInstanceConfigRequest {
	s.DeleteConfigList = &v
	return s
}

func (s *ModifyInstanceConfigRequest) SetInstanceId(v string) *ModifyInstanceConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceConfigRequest) SetReason(v string) *ModifyInstanceConfigRequest {
	s.Reason = &v
	return s
}

func (s *ModifyInstanceConfigRequest) SetConfigsToAdd(v []*InstanceConfigDto) *ModifyInstanceConfigRequest {
	s.ConfigsToAdd = v
	return s
}

func (s *ModifyInstanceConfigRequest) SetConfigsToDelete(v []*InstanceConfigDto) *ModifyInstanceConfigRequest {
	s.ConfigsToDelete = v
	return s
}

func (s *ModifyInstanceConfigRequest) SetConfigsToUpdate(v []*InstanceConfigDto) *ModifyInstanceConfigRequest {
	s.ConfigsToUpdate = v
	return s
}

func (s *ModifyInstanceConfigRequest) SetFastMode(v bool) *ModifyInstanceConfigRequest {
	s.FastMode = &v
	return s
}

func (s *ModifyInstanceConfigRequest) SetRestart(v bool) *ModifyInstanceConfigRequest {
	s.Restart = &v
	return s
}

func (s *ModifyInstanceConfigRequest) Validate() error {
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
