// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSystemConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v []*UpdateSystemConfigsRequestConfigs) *UpdateSystemConfigsRequest
	GetConfigs() []*UpdateSystemConfigsRequestConfigs
	SetObjectId(v string) *UpdateSystemConfigsRequest
	GetObjectId() *string
	SetObjectType(v string) *UpdateSystemConfigsRequest
	GetObjectType() *string
}

type UpdateSystemConfigsRequest struct {
	// 配置列表
	Configs []*UpdateSystemConfigsRequestConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// 对象ID
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// 外呼开发时补充参数限制
	//
	// example:
	//
	// INSTANCE
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
}

func (s UpdateSystemConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSystemConfigsRequest) GoString() string {
	return s.String()
}

func (s *UpdateSystemConfigsRequest) GetConfigs() []*UpdateSystemConfigsRequestConfigs {
	return s.Configs
}

func (s *UpdateSystemConfigsRequest) GetObjectId() *string {
	return s.ObjectId
}

func (s *UpdateSystemConfigsRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *UpdateSystemConfigsRequest) SetConfigs(v []*UpdateSystemConfigsRequestConfigs) *UpdateSystemConfigsRequest {
	s.Configs = v
	return s
}

func (s *UpdateSystemConfigsRequest) SetObjectId(v string) *UpdateSystemConfigsRequest {
	s.ObjectId = &v
	return s
}

func (s *UpdateSystemConfigsRequest) SetObjectType(v string) *UpdateSystemConfigsRequest {
	s.ObjectType = &v
	return s
}

func (s *UpdateSystemConfigsRequest) Validate() error {
	if s.Configs != nil {
		for _, item := range s.Configs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateSystemConfigsRequestConfigs struct {
	// example:
	//
	// callableTime
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 5
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateSystemConfigsRequestConfigs) String() string {
	return dara.Prettify(s)
}

func (s UpdateSystemConfigsRequestConfigs) GoString() string {
	return s.String()
}

func (s *UpdateSystemConfigsRequestConfigs) GetName() *string {
	return s.Name
}

func (s *UpdateSystemConfigsRequestConfigs) GetValue() *string {
	return s.Value
}

func (s *UpdateSystemConfigsRequestConfigs) SetName(v string) *UpdateSystemConfigsRequestConfigs {
	s.Name = &v
	return s
}

func (s *UpdateSystemConfigsRequestConfigs) SetValue(v string) *UpdateSystemConfigsRequestConfigs {
	s.Value = &v
	return s
}

func (s *UpdateSystemConfigsRequestConfigs) Validate() error {
	return dara.Validate(s)
}
