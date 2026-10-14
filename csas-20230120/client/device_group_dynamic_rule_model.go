// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeviceGroupDynamicRule interface {
	dara.Model
	String() string
	GoString() string
	SetArgs(v interface{}) *DeviceGroupDynamicRule
	GetArgs() interface{}
	SetKey(v string) *DeviceGroupDynamicRule
	GetKey() *string
	SetOperator(v string) *DeviceGroupDynamicRule
	GetOperator() *string
}

type DeviceGroupDynamicRule struct {
	// The list of matching parameters. The meaning of the parameters is determined by both Key and Operator. When Key is set to dev_type and Operator is set to in, specify the operating system types to match.
	//
	// example:
	//
	// ["windows"]
	Args interface{} `json:"Args,omitempty" xml:"Args,omitempty"`
	// The user endpoint device property key used for dynamic matching. The example value dev_type indicates the operating system type of the endpoint device.
	//
	// example:
	//
	// dev_type
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The matching operator applied to the property specified by Key. The example value in indicates that the property value belongs to the list of parameters specified by Args.
	//
	// example:
	//
	// in
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
}

func (s DeviceGroupDynamicRule) String() string {
	return dara.Prettify(s)
}

func (s DeviceGroupDynamicRule) GoString() string {
	return s.String()
}

func (s *DeviceGroupDynamicRule) GetArgs() interface{} {
	return s.Args
}

func (s *DeviceGroupDynamicRule) GetKey() *string {
	return s.Key
}

func (s *DeviceGroupDynamicRule) GetOperator() *string {
	return s.Operator
}

func (s *DeviceGroupDynamicRule) SetArgs(v interface{}) *DeviceGroupDynamicRule {
	s.Args = v
	return s
}

func (s *DeviceGroupDynamicRule) SetKey(v string) *DeviceGroupDynamicRule {
	s.Key = &v
	return s
}

func (s *DeviceGroupDynamicRule) SetOperator(v string) *DeviceGroupDynamicRule {
	s.Operator = &v
	return s
}

func (s *DeviceGroupDynamicRule) Validate() error {
	return dara.Validate(s)
}
