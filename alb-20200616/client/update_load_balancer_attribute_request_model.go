// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLoadBalancerAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateLoadBalancerAttributeRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateLoadBalancerAttributeRequest
	GetDryRun() *bool
	SetLoadBalancerId(v string) *UpdateLoadBalancerAttributeRequest
	GetLoadBalancerId() *string
	SetLoadBalancerName(v string) *UpdateLoadBalancerAttributeRequest
	GetLoadBalancerName() *string
	SetModificationProtectionConfig(v *UpdateLoadBalancerAttributeRequestModificationProtectionConfig) *UpdateLoadBalancerAttributeRequest
	GetModificationProtectionConfig() *UpdateLoadBalancerAttributeRequestModificationProtectionConfig
}

type UpdateLoadBalancerAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false**: performs a dry run and sends the request. If the request passes the dry run, a `2xx HTTP` status code is returned and the operation is performed. This is the default value.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the ALB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// alb-o9ulmq5hgn68jk****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The name of the ALB instance. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// lb-instance-test
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	// The configuration read-only mode settings.
	ModificationProtectionConfig *UpdateLoadBalancerAttributeRequestModificationProtectionConfig `json:"ModificationProtectionConfig,omitempty" xml:"ModificationProtectionConfig,omitempty" type:"Struct"`
}

func (s UpdateLoadBalancerAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLoadBalancerAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateLoadBalancerAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateLoadBalancerAttributeRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *UpdateLoadBalancerAttributeRequest) GetLoadBalancerName() *string {
	return s.LoadBalancerName
}

func (s *UpdateLoadBalancerAttributeRequest) GetModificationProtectionConfig() *UpdateLoadBalancerAttributeRequestModificationProtectionConfig {
	return s.ModificationProtectionConfig
}

func (s *UpdateLoadBalancerAttributeRequest) SetClientToken(v string) *UpdateLoadBalancerAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateLoadBalancerAttributeRequest) SetDryRun(v bool) *UpdateLoadBalancerAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateLoadBalancerAttributeRequest) SetLoadBalancerId(v string) *UpdateLoadBalancerAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *UpdateLoadBalancerAttributeRequest) SetLoadBalancerName(v string) *UpdateLoadBalancerAttributeRequest {
	s.LoadBalancerName = &v
	return s
}

func (s *UpdateLoadBalancerAttributeRequest) SetModificationProtectionConfig(v *UpdateLoadBalancerAttributeRequestModificationProtectionConfig) *UpdateLoadBalancerAttributeRequest {
	s.ModificationProtectionConfig = v
	return s
}

func (s *UpdateLoadBalancerAttributeRequest) Validate() error {
	if s.ModificationProtectionConfig != nil {
		if err := s.ModificationProtectionConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateLoadBalancerAttributeRequestModificationProtectionConfig struct {
	// The reason for enabling the configuration read-only mode.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
	//
	// This parameter takes effect only when **Status*	- is set to **ConsoleProtection**.
	//
	// example:
	//
	// test
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// Specifies whether to enable the configuration read-only mode. Valid values:
	//
	// 	- **NonProtection**: disables the configuration read-only mode. In this case, the value of the **Reason*	- parameter that you specify does not take effect. If you set the value of **Reason**, the value is cleared.
	//
	// 	- **ConsoleProtection**: enables the configuration read-only mode. In this case, the value of the **Reason*	- parameter that you specify takes effect.****
	//
	// >  If the parameter is set to **ConsoleProtection**, the configuration read-only mode is enabled. You cannot modify the configurations of the ALB instance in the ALB console. However, you can call API operations to modify the configurations of the ALB instance.
	//
	// example:
	//
	// ConsoleProtection
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateLoadBalancerAttributeRequestModificationProtectionConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateLoadBalancerAttributeRequestModificationProtectionConfig) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerAttributeRequestModificationProtectionConfig) GetReason() *string {
	return s.Reason
}

func (s *UpdateLoadBalancerAttributeRequestModificationProtectionConfig) GetStatus() *string {
	return s.Status
}

func (s *UpdateLoadBalancerAttributeRequestModificationProtectionConfig) SetReason(v string) *UpdateLoadBalancerAttributeRequestModificationProtectionConfig {
	s.Reason = &v
	return s
}

func (s *UpdateLoadBalancerAttributeRequestModificationProtectionConfig) SetStatus(v string) *UpdateLoadBalancerAttributeRequestModificationProtectionConfig {
	s.Status = &v
	return s
}

func (s *UpdateLoadBalancerAttributeRequestModificationProtectionConfig) Validate() error {
	return dara.Validate(s)
}
