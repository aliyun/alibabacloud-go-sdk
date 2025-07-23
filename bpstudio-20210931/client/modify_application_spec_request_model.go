// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApplicationSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ModifyApplicationSpecRequest
	GetApplicationId() *string
	SetInstanceSpec(v []*ModifyApplicationSpecRequestInstanceSpec) *ModifyApplicationSpecRequest
	GetInstanceSpec() []*ModifyApplicationSpecRequestInstanceSpec
}

type ModifyApplicationSpecRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 02S7UU41WKJL7ERR
	ApplicationId *string                                     `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	InstanceSpec  []*ModifyApplicationSpecRequestInstanceSpec `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty" type:"Repeated"`
}

func (s ModifyApplicationSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyApplicationSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyApplicationSpecRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ModifyApplicationSpecRequest) GetInstanceSpec() []*ModifyApplicationSpecRequestInstanceSpec {
	return s.InstanceSpec
}

func (s *ModifyApplicationSpecRequest) SetApplicationId(v string) *ModifyApplicationSpecRequest {
	s.ApplicationId = &v
	return s
}

func (s *ModifyApplicationSpecRequest) SetInstanceSpec(v []*ModifyApplicationSpecRequestInstanceSpec) *ModifyApplicationSpecRequest {
	s.InstanceSpec = v
	return s
}

func (s *ModifyApplicationSpecRequest) Validate() error {
	return dara.Validate(s)
}

type ModifyApplicationSpecRequestInstanceSpec struct {
	Configuration map[string]interface{} `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	// example:
	//
	// rm-2ze8f4ah378a*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ModifyApplicationSpecRequestInstanceSpec) String() string {
	return dara.Prettify(s)
}

func (s ModifyApplicationSpecRequestInstanceSpec) GoString() string {
	return s.String()
}

func (s *ModifyApplicationSpecRequestInstanceSpec) GetConfiguration() map[string]interface{} {
	return s.Configuration
}

func (s *ModifyApplicationSpecRequestInstanceSpec) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyApplicationSpecRequestInstanceSpec) SetConfiguration(v map[string]interface{}) *ModifyApplicationSpecRequestInstanceSpec {
	s.Configuration = v
	return s
}

func (s *ModifyApplicationSpecRequestInstanceSpec) SetInstanceId(v string) *ModifyApplicationSpecRequestInstanceSpec {
	s.InstanceId = &v
	return s
}

func (s *ModifyApplicationSpecRequestInstanceSpec) Validate() error {
	return dara.Validate(s)
}
