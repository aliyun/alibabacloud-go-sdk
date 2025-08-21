// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPortAutoCcStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyPortAutoCcStatusRequest
	GetInstanceId() *string
	SetMode(v string) *ModifyPortAutoCcStatusRequest
	GetMode() *string
	SetSwitch(v string) *ModifyPortAutoCcStatusRequest
	GetSwitch() *string
}

type ModifyPortAutoCcStatusRequest struct {
	// The ID of the instance.
	//
	// > You can call the [DescribeInstanceIds](https://help.aliyun.com/document_detail/157459.html) operation to query the IDs of all instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The mode of the Intelligent Protection policy. Valid values:
	//
	// 	- **normal**
	//
	// 	- **loose**
	//
	// 	- **strict**
	//
	// This parameter is required.
	//
	// example:
	//
	// normal
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// Specifies the status of the Intelligent Protection policy. Valid values:
	//
	// 	- **on**: enables the policy.
	//
	// 	- **off**: disables the policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// on
	Switch *string `json:"Switch,omitempty" xml:"Switch,omitempty"`
}

func (s ModifyPortAutoCcStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPortAutoCcStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyPortAutoCcStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyPortAutoCcStatusRequest) GetMode() *string {
	return s.Mode
}

func (s *ModifyPortAutoCcStatusRequest) GetSwitch() *string {
	return s.Switch
}

func (s *ModifyPortAutoCcStatusRequest) SetInstanceId(v string) *ModifyPortAutoCcStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyPortAutoCcStatusRequest) SetMode(v string) *ModifyPortAutoCcStatusRequest {
	s.Mode = &v
	return s
}

func (s *ModifyPortAutoCcStatusRequest) SetSwitch(v string) *ModifyPortAutoCcStatusRequest {
	s.Switch = &v
	return s
}

func (s *ModifyPortAutoCcStatusRequest) Validate() error {
	return dara.Validate(s)
}
