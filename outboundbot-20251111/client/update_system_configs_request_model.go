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
	// The list of configurations.
	Configs []*UpdateSystemConfigsRequestConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// The configuration type ID. If ObjectType is set to INSTANCE, this parameter specifies the instance ID. If ObjectType is set to TENANT, this parameter specifies the tenant ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The configuration type. Valid values:
	//
	// - INSTANCE: instance level.
	//
	// - TENANT: tenant level.
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
	// The system configuration name. Valid values:
	//
	// - callableTime: the outbound job window.
	//
	// - calleeDailyAttemptLimit: the maximum number of daily calls to a single callee number.
	//
	// example:
	//
	// callableTime
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The configuration value.
	//
	// - If Name is set to callableTime, a sample Value is [{"beginTime":"09:00:00","endTime":"12:00:00"},{"beginTime":"14:00:00","endTime":"18:00:00"}].
	//
	// - If Name is set to calleeDailyAttemptLimit, the Value is an integer from 1 to 50.
	//
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
