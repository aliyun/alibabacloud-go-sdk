// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceRAGConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyInstanceRAGConfigRequest
	GetClientToken() *string
	SetConfigList(v []*ModifyInstanceRAGConfigRequestConfigList) *ModifyInstanceRAGConfigRequest
	GetConfigList() []*ModifyInstanceRAGConfigRequestConfigList
	SetInstanceName(v string) *ModifyInstanceRAGConfigRequest
	GetInstanceName() *string
	SetRegionId(v string) *ModifyInstanceRAGConfigRequest
	GetRegionId() *string
	SetStatus(v bool) *ModifyInstanceRAGConfigRequest
	GetStatus() *bool
}

type ModifyInstanceRAGConfigRequest struct {
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string                                     `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConfigList  []*ModifyInstanceRAGConfigRequestConfigList `json:"ConfigList,omitempty" xml:"ConfigList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyInstanceRAGConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceRAGConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceRAGConfigRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyInstanceRAGConfigRequest) GetConfigList() []*ModifyInstanceRAGConfigRequestConfigList {
	return s.ConfigList
}

func (s *ModifyInstanceRAGConfigRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyInstanceRAGConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyInstanceRAGConfigRequest) GetStatus() *bool {
	return s.Status
}

func (s *ModifyInstanceRAGConfigRequest) SetClientToken(v string) *ModifyInstanceRAGConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyInstanceRAGConfigRequest) SetConfigList(v []*ModifyInstanceRAGConfigRequestConfigList) *ModifyInstanceRAGConfigRequest {
	s.ConfigList = v
	return s
}

func (s *ModifyInstanceRAGConfigRequest) SetInstanceName(v string) *ModifyInstanceRAGConfigRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyInstanceRAGConfigRequest) SetRegionId(v string) *ModifyInstanceRAGConfigRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceRAGConfigRequest) SetStatus(v bool) *ModifyInstanceRAGConfigRequest {
	s.Status = &v
	return s
}

func (s *ModifyInstanceRAGConfigRequest) Validate() error {
	if s.ConfigList != nil {
		for _, item := range s.ConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyInstanceRAGConfigRequestConfigList struct {
	// example:
	//
	// LLM_MODEL
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// qwen-flash
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyInstanceRAGConfigRequestConfigList) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceRAGConfigRequestConfigList) GoString() string {
	return s.String()
}

func (s *ModifyInstanceRAGConfigRequestConfigList) GetName() *string {
	return s.Name
}

func (s *ModifyInstanceRAGConfigRequestConfigList) GetValue() *string {
	return s.Value
}

func (s *ModifyInstanceRAGConfigRequestConfigList) SetName(v string) *ModifyInstanceRAGConfigRequestConfigList {
	s.Name = &v
	return s
}

func (s *ModifyInstanceRAGConfigRequestConfigList) SetValue(v string) *ModifyInstanceRAGConfigRequestConfigList {
	s.Value = &v
	return s
}

func (s *ModifyInstanceRAGConfigRequestConfigList) Validate() error {
	return dara.Validate(s)
}
