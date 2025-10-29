// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientNodeConfiguration(v *ClientNodeConfiguration) *UpdateInstanceRequest
	GetClientNodeConfiguration() *ClientNodeConfiguration
	SetElasticDataNodeConfiguration(v *ElasticDataNodeConfiguration) *UpdateInstanceRequest
	GetElasticDataNodeConfiguration() *ElasticDataNodeConfiguration
	SetInstanceCategory(v string) *UpdateInstanceRequest
	GetInstanceCategory() *string
	SetKibanaConfiguration(v *KibanaNodeConfiguration) *UpdateInstanceRequest
	GetKibanaConfiguration() *KibanaNodeConfiguration
	SetMasterConfiguration(v *MasterNodeConfiguration) *UpdateInstanceRequest
	GetMasterConfiguration() *MasterNodeConfiguration
	SetNodeAmount(v int32) *UpdateInstanceRequest
	GetNodeAmount() *int32
	SetNodeSpec(v *NodeSpec) *UpdateInstanceRequest
	GetNodeSpec() *NodeSpec
	SetWarmNodeConfiguration(v *WarmNodeConfiguration) *UpdateInstanceRequest
	GetWarmNodeConfiguration() *WarmNodeConfiguration
	SetClientToken(v string) *UpdateInstanceRequest
	GetClientToken() *string
	SetForce(v bool) *UpdateInstanceRequest
	GetForce() *bool
	SetOrderActionType(v string) *UpdateInstanceRequest
	GetOrderActionType() *string
}

type UpdateInstanceRequest struct {
	ClientNodeConfiguration      *ClientNodeConfiguration      `json:"clientNodeConfiguration,omitempty" xml:"clientNodeConfiguration,omitempty"`
	ElasticDataNodeConfiguration *ElasticDataNodeConfiguration `json:"elasticDataNodeConfiguration,omitempty" xml:"elasticDataNodeConfiguration,omitempty"`
	// example:
	//
	// advanced
	InstanceCategory    *string                  `json:"instanceCategory,omitempty" xml:"instanceCategory,omitempty"`
	KibanaConfiguration *KibanaNodeConfiguration `json:"kibanaConfiguration,omitempty" xml:"kibanaConfiguration,omitempty"`
	MasterConfiguration *MasterNodeConfiguration `json:"masterConfiguration,omitempty" xml:"masterConfiguration,omitempty"`
	// example:
	//
	// 3
	NodeAmount            *int32                 `json:"nodeAmount,omitempty" xml:"nodeAmount,omitempty"`
	NodeSpec              *NodeSpec              `json:"nodeSpec,omitempty" xml:"nodeSpec,omitempty"`
	WarmNodeConfiguration *WarmNodeConfiguration `json:"warmNodeConfiguration,omitempty" xml:"warmNodeConfiguration,omitempty"`
	// The result of the request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// example:
	//
	// false
	Force *bool `json:"force,omitempty" xml:"force,omitempty"`
	// The number of data nodes.
	//
	// example:
	//
	// upgrade
	OrderActionType *string `json:"orderActionType,omitempty" xml:"orderActionType,omitempty"`
}

func (s UpdateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequest) GetClientNodeConfiguration() *ClientNodeConfiguration {
	return s.ClientNodeConfiguration
}

func (s *UpdateInstanceRequest) GetElasticDataNodeConfiguration() *ElasticDataNodeConfiguration {
	return s.ElasticDataNodeConfiguration
}

func (s *UpdateInstanceRequest) GetInstanceCategory() *string {
	return s.InstanceCategory
}

func (s *UpdateInstanceRequest) GetKibanaConfiguration() *KibanaNodeConfiguration {
	return s.KibanaConfiguration
}

func (s *UpdateInstanceRequest) GetMasterConfiguration() *MasterNodeConfiguration {
	return s.MasterConfiguration
}

func (s *UpdateInstanceRequest) GetNodeAmount() *int32 {
	return s.NodeAmount
}

func (s *UpdateInstanceRequest) GetNodeSpec() *NodeSpec {
	return s.NodeSpec
}

func (s *UpdateInstanceRequest) GetWarmNodeConfiguration() *WarmNodeConfiguration {
	return s.WarmNodeConfiguration
}

func (s *UpdateInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateInstanceRequest) GetForce() *bool {
	return s.Force
}

func (s *UpdateInstanceRequest) GetOrderActionType() *string {
	return s.OrderActionType
}

func (s *UpdateInstanceRequest) SetClientNodeConfiguration(v *ClientNodeConfiguration) *UpdateInstanceRequest {
	s.ClientNodeConfiguration = v
	return s
}

func (s *UpdateInstanceRequest) SetElasticDataNodeConfiguration(v *ElasticDataNodeConfiguration) *UpdateInstanceRequest {
	s.ElasticDataNodeConfiguration = v
	return s
}

func (s *UpdateInstanceRequest) SetInstanceCategory(v string) *UpdateInstanceRequest {
	s.InstanceCategory = &v
	return s
}

func (s *UpdateInstanceRequest) SetKibanaConfiguration(v *KibanaNodeConfiguration) *UpdateInstanceRequest {
	s.KibanaConfiguration = v
	return s
}

func (s *UpdateInstanceRequest) SetMasterConfiguration(v *MasterNodeConfiguration) *UpdateInstanceRequest {
	s.MasterConfiguration = v
	return s
}

func (s *UpdateInstanceRequest) SetNodeAmount(v int32) *UpdateInstanceRequest {
	s.NodeAmount = &v
	return s
}

func (s *UpdateInstanceRequest) SetNodeSpec(v *NodeSpec) *UpdateInstanceRequest {
	s.NodeSpec = v
	return s
}

func (s *UpdateInstanceRequest) SetWarmNodeConfiguration(v *WarmNodeConfiguration) *UpdateInstanceRequest {
	s.WarmNodeConfiguration = v
	return s
}

func (s *UpdateInstanceRequest) SetClientToken(v string) *UpdateInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateInstanceRequest) SetForce(v bool) *UpdateInstanceRequest {
	s.Force = &v
	return s
}

func (s *UpdateInstanceRequest) SetOrderActionType(v string) *UpdateInstanceRequest {
	s.OrderActionType = &v
	return s
}

func (s *UpdateInstanceRequest) Validate() error {
	if s.ClientNodeConfiguration != nil {
		if err := s.ClientNodeConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.ElasticDataNodeConfiguration != nil {
		if err := s.ElasticDataNodeConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.KibanaConfiguration != nil {
		if err := s.KibanaConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.MasterConfiguration != nil {
		if err := s.MasterConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.NodeSpec != nil {
		if err := s.NodeSpec.Validate(); err != nil {
			return err
		}
	}
	if s.WarmNodeConfiguration != nil {
		if err := s.WarmNodeConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
