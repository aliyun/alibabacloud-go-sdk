// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceServiceConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ModifyInstanceServiceConfigRequest
	GetClusterId() *string
	SetConfigureName(v string) *ModifyInstanceServiceConfigRequest
	GetConfigureName() *string
	SetConfigureValue(v string) *ModifyInstanceServiceConfigRequest
	GetConfigureValue() *string
	SetParameters(v string) *ModifyInstanceServiceConfigRequest
	GetParameters() *string
	SetRestart(v bool) *ModifyInstanceServiceConfigRequest
	GetRestart() *bool
}

type ModifyInstanceServiceConfigRequest struct {
	// The ID of target instance. You can call the [DescribeInstances](https://help.aliyun.com/document_detail/144595.html) operation to obtain target instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// hb-t4naqsay5gn****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// <props="china">The name of the configuration item to modify. You can call the [ListInstanceServiceConfigurations](https://help.aliyun.com/document_detail/201980.html) operation to query the configuration item name.
	//
	// <props="intl">The name of the configuration item to modify.
	//
	// > If you want to modify multiple configuration items, specify the Parameters parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// hbase#hbase-site.xml#hbase.client.keyvalue.maxsize
	ConfigureName *string `json:"ConfigureName,omitempty" xml:"ConfigureName,omitempty"`
	// <props="china">The value of the configuration item to modify. You can call the [ListInstanceServiceConfigurations](https://help.aliyun.com/document_detail/201980.html) operation to query the configuration item value.
	//
	// <props="intl">The value of the configuration item to modify.
	//
	// > If you want to modify multiple configuration items, specify the Parameters parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10485770
	ConfigureValue *string `json:"ConfigureValue,omitempty" xml:"ConfigureValue,omitempty"`
	// The JSON-formatted parameters for modifying multiple configuration items. The key specifies the name of the configuration item, and the value specifies the value of the configuration item.
	//
	// example:
	//
	// {"key1=value1", "key2=value2"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// Specifies whether to restart the instance after the configuration is modified. Valid values:
	//
	// - **true**: Restart the instance.
	//
	// - **false**: Do not restart the instance.
	//
	// example:
	//
	// false
	Restart *bool `json:"Restart,omitempty" xml:"Restart,omitempty"`
}

func (s ModifyInstanceServiceConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceServiceConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceServiceConfigRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyInstanceServiceConfigRequest) GetConfigureName() *string {
	return s.ConfigureName
}

func (s *ModifyInstanceServiceConfigRequest) GetConfigureValue() *string {
	return s.ConfigureValue
}

func (s *ModifyInstanceServiceConfigRequest) GetParameters() *string {
	return s.Parameters
}

func (s *ModifyInstanceServiceConfigRequest) GetRestart() *bool {
	return s.Restart
}

func (s *ModifyInstanceServiceConfigRequest) SetClusterId(v string) *ModifyInstanceServiceConfigRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyInstanceServiceConfigRequest) SetConfigureName(v string) *ModifyInstanceServiceConfigRequest {
	s.ConfigureName = &v
	return s
}

func (s *ModifyInstanceServiceConfigRequest) SetConfigureValue(v string) *ModifyInstanceServiceConfigRequest {
	s.ConfigureValue = &v
	return s
}

func (s *ModifyInstanceServiceConfigRequest) SetParameters(v string) *ModifyInstanceServiceConfigRequest {
	s.Parameters = &v
	return s
}

func (s *ModifyInstanceServiceConfigRequest) SetRestart(v bool) *ModifyInstanceServiceConfigRequest {
	s.Restart = &v
	return s
}

func (s *ModifyInstanceServiceConfigRequest) Validate() error {
	return dara.Validate(s)
}
