// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEgressProxyGatewayAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddListenerProperties(v []*UpdateEgressProxyGatewayAttributeRequestAddListenerProperties) *UpdateEgressProxyGatewayAttributeRequest
	GetAddListenerProperties() []*UpdateEgressProxyGatewayAttributeRequestAddListenerProperties
	SetClientToken(v string) *UpdateEgressProxyGatewayAttributeRequest
	GetClientToken() *string
	SetDefaultEgressProxyIpPoolId(v string) *UpdateEgressProxyGatewayAttributeRequest
	GetDefaultEgressProxyIpPoolId() *string
	SetDescription(v string) *UpdateEgressProxyGatewayAttributeRequest
	GetDescription() *string
	SetDryRun(v bool) *UpdateEgressProxyGatewayAttributeRequest
	GetDryRun() *bool
	SetEgressProxyGatewayId(v string) *UpdateEgressProxyGatewayAttributeRequest
	GetEgressProxyGatewayId() *string
	SetEgressProxyGatewayName(v string) *UpdateEgressProxyGatewayAttributeRequest
	GetEgressProxyGatewayName() *string
	SetRemoveListenerProperties(v []*UpdateEgressProxyGatewayAttributeRequestRemoveListenerProperties) *UpdateEgressProxyGatewayAttributeRequest
	GetRemoveListenerProperties() []*UpdateEgressProxyGatewayAttributeRequestRemoveListenerProperties
	SetTlsInterceptProperties(v *UpdateEgressProxyGatewayAttributeRequestTlsInterceptProperties) *UpdateEgressProxyGatewayAttributeRequest
	GetTlsInterceptProperties() *UpdateEgressProxyGatewayAttributeRequestTlsInterceptProperties
}

type UpdateEgressProxyGatewayAttributeRequest struct {
	AddListenerProperties []*UpdateEgressProxyGatewayAttributeRequestAddListenerProperties `json:"AddListenerProperties,omitempty" xml:"AddListenerProperties,omitempty" type:"Repeated"`
	// example:
	//
	// 123e4567-e89b-12d3-a456-426614174000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// eppool-gw80d96a96ca0b00****
	DefaultEgressProxyIpPoolId *string `json:"DefaultEgressProxyIpPoolId,omitempty" xml:"DefaultEgressProxyIpPoolId,omitempty"`
	// example:
	//
	// testEpg
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// epg-gw80bb6a94f835db****
	EgressProxyGatewayId *string `json:"EgressProxyGatewayId,omitempty" xml:"EgressProxyGatewayId,omitempty"`
	// example:
	//
	// testEpg
	EgressProxyGatewayName   *string                                                             `json:"EgressProxyGatewayName,omitempty" xml:"EgressProxyGatewayName,omitempty"`
	RemoveListenerProperties []*UpdateEgressProxyGatewayAttributeRequestRemoveListenerProperties `json:"RemoveListenerProperties,omitempty" xml:"RemoveListenerProperties,omitempty" type:"Repeated"`
	TlsInterceptProperties   *UpdateEgressProxyGatewayAttributeRequestTlsInterceptProperties     `json:"TlsInterceptProperties,omitempty" xml:"TlsInterceptProperties,omitempty" type:"Struct"`
}

func (s UpdateEgressProxyGatewayAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEgressProxyGatewayAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateEgressProxyGatewayAttributeRequest) GetAddListenerProperties() []*UpdateEgressProxyGatewayAttributeRequestAddListenerProperties {
	return s.AddListenerProperties
}

func (s *UpdateEgressProxyGatewayAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateEgressProxyGatewayAttributeRequest) GetDefaultEgressProxyIpPoolId() *string {
	return s.DefaultEgressProxyIpPoolId
}

func (s *UpdateEgressProxyGatewayAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateEgressProxyGatewayAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateEgressProxyGatewayAttributeRequest) GetEgressProxyGatewayId() *string {
	return s.EgressProxyGatewayId
}

func (s *UpdateEgressProxyGatewayAttributeRequest) GetEgressProxyGatewayName() *string {
	return s.EgressProxyGatewayName
}

func (s *UpdateEgressProxyGatewayAttributeRequest) GetRemoveListenerProperties() []*UpdateEgressProxyGatewayAttributeRequestRemoveListenerProperties {
	return s.RemoveListenerProperties
}

func (s *UpdateEgressProxyGatewayAttributeRequest) GetTlsInterceptProperties() *UpdateEgressProxyGatewayAttributeRequestTlsInterceptProperties {
	return s.TlsInterceptProperties
}

func (s *UpdateEgressProxyGatewayAttributeRequest) SetAddListenerProperties(v []*UpdateEgressProxyGatewayAttributeRequestAddListenerProperties) *UpdateEgressProxyGatewayAttributeRequest {
	s.AddListenerProperties = v
	return s
}

func (s *UpdateEgressProxyGatewayAttributeRequest) SetClientToken(v string) *UpdateEgressProxyGatewayAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateEgressProxyGatewayAttributeRequest) SetDefaultEgressProxyIpPoolId(v string) *UpdateEgressProxyGatewayAttributeRequest {
	s.DefaultEgressProxyIpPoolId = &v
	return s
}

func (s *UpdateEgressProxyGatewayAttributeRequest) SetDescription(v string) *UpdateEgressProxyGatewayAttributeRequest {
	s.Description = &v
	return s
}

func (s *UpdateEgressProxyGatewayAttributeRequest) SetDryRun(v bool) *UpdateEgressProxyGatewayAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateEgressProxyGatewayAttributeRequest) SetEgressProxyGatewayId(v string) *UpdateEgressProxyGatewayAttributeRequest {
	s.EgressProxyGatewayId = &v
	return s
}

func (s *UpdateEgressProxyGatewayAttributeRequest) SetEgressProxyGatewayName(v string) *UpdateEgressProxyGatewayAttributeRequest {
	s.EgressProxyGatewayName = &v
	return s
}

func (s *UpdateEgressProxyGatewayAttributeRequest) SetRemoveListenerProperties(v []*UpdateEgressProxyGatewayAttributeRequestRemoveListenerProperties) *UpdateEgressProxyGatewayAttributeRequest {
	s.RemoveListenerProperties = v
	return s
}

func (s *UpdateEgressProxyGatewayAttributeRequest) SetTlsInterceptProperties(v *UpdateEgressProxyGatewayAttributeRequestTlsInterceptProperties) *UpdateEgressProxyGatewayAttributeRequest {
	s.TlsInterceptProperties = v
	return s
}

func (s *UpdateEgressProxyGatewayAttributeRequest) Validate() error {
	if s.AddListenerProperties != nil {
		for _, item := range s.AddListenerProperties {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RemoveListenerProperties != nil {
		for _, item := range s.RemoveListenerProperties {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TlsInterceptProperties != nil {
		if err := s.TlsInterceptProperties.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateEgressProxyGatewayAttributeRequestAddListenerProperties struct {
	// This parameter is required.
	//
	// example:
	//
	// 17486
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s UpdateEgressProxyGatewayAttributeRequestAddListenerProperties) String() string {
	return dara.Prettify(s)
}

func (s UpdateEgressProxyGatewayAttributeRequestAddListenerProperties) GoString() string {
	return s.String()
}

func (s *UpdateEgressProxyGatewayAttributeRequestAddListenerProperties) GetPort() *int32 {
	return s.Port
}

func (s *UpdateEgressProxyGatewayAttributeRequestAddListenerProperties) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateEgressProxyGatewayAttributeRequestAddListenerProperties) SetPort(v int32) *UpdateEgressProxyGatewayAttributeRequestAddListenerProperties {
	s.Port = &v
	return s
}

func (s *UpdateEgressProxyGatewayAttributeRequestAddListenerProperties) SetProtocol(v string) *UpdateEgressProxyGatewayAttributeRequestAddListenerProperties {
	s.Protocol = &v
	return s
}

func (s *UpdateEgressProxyGatewayAttributeRequestAddListenerProperties) Validate() error {
	return dara.Validate(s)
}

type UpdateEgressProxyGatewayAttributeRequestRemoveListenerProperties struct {
	// This parameter is required.
	//
	// example:
	//
	// 32720
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s UpdateEgressProxyGatewayAttributeRequestRemoveListenerProperties) String() string {
	return dara.Prettify(s)
}

func (s UpdateEgressProxyGatewayAttributeRequestRemoveListenerProperties) GoString() string {
	return s.String()
}

func (s *UpdateEgressProxyGatewayAttributeRequestRemoveListenerProperties) GetPort() *int32 {
	return s.Port
}

func (s *UpdateEgressProxyGatewayAttributeRequestRemoveListenerProperties) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateEgressProxyGatewayAttributeRequestRemoveListenerProperties) SetPort(v int32) *UpdateEgressProxyGatewayAttributeRequestRemoveListenerProperties {
	s.Port = &v
	return s
}

func (s *UpdateEgressProxyGatewayAttributeRequestRemoveListenerProperties) SetProtocol(v string) *UpdateEgressProxyGatewayAttributeRequestRemoveListenerProperties {
	s.Protocol = &v
	return s
}

func (s *UpdateEgressProxyGatewayAttributeRequestRemoveListenerProperties) Validate() error {
	return dara.Validate(s)
}

type UpdateEgressProxyGatewayAttributeRequestTlsInterceptProperties struct {
	// example:
	//
	// 1f01a6f9-a0e7-64e5-9d3c-29c056****
	PrivateCaCertificateId *string `json:"PrivateCaCertificateId,omitempty" xml:"PrivateCaCertificateId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// enable
	TlsInterceptMode *string `json:"TlsInterceptMode,omitempty" xml:"TlsInterceptMode,omitempty"`
}

func (s UpdateEgressProxyGatewayAttributeRequestTlsInterceptProperties) String() string {
	return dara.Prettify(s)
}

func (s UpdateEgressProxyGatewayAttributeRequestTlsInterceptProperties) GoString() string {
	return s.String()
}

func (s *UpdateEgressProxyGatewayAttributeRequestTlsInterceptProperties) GetPrivateCaCertificateId() *string {
	return s.PrivateCaCertificateId
}

func (s *UpdateEgressProxyGatewayAttributeRequestTlsInterceptProperties) GetTlsInterceptMode() *string {
	return s.TlsInterceptMode
}

func (s *UpdateEgressProxyGatewayAttributeRequestTlsInterceptProperties) SetPrivateCaCertificateId(v string) *UpdateEgressProxyGatewayAttributeRequestTlsInterceptProperties {
	s.PrivateCaCertificateId = &v
	return s
}

func (s *UpdateEgressProxyGatewayAttributeRequestTlsInterceptProperties) SetTlsInterceptMode(v string) *UpdateEgressProxyGatewayAttributeRequestTlsInterceptProperties {
	s.TlsInterceptMode = &v
	return s
}

func (s *UpdateEgressProxyGatewayAttributeRequestTlsInterceptProperties) Validate() error {
	return dara.Validate(s)
}
