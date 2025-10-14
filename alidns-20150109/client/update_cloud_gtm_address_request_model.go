// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateCloudGtmAddressRequest
	GetAcceptLanguage() *string
	SetAddress(v string) *UpdateCloudGtmAddressRequest
	GetAddress() *string
	SetAddressId(v string) *UpdateCloudGtmAddressRequest
	GetAddressId() *string
	SetAttributeInfo(v string) *UpdateCloudGtmAddressRequest
	GetAttributeInfo() *string
	SetClientToken(v string) *UpdateCloudGtmAddressRequest
	GetClientToken() *string
	SetHealthJudgement(v string) *UpdateCloudGtmAddressRequest
	GetHealthJudgement() *string
	SetHealthTasks(v []*UpdateCloudGtmAddressRequestHealthTasks) *UpdateCloudGtmAddressRequest
	GetHealthTasks() []*UpdateCloudGtmAddressRequestHealthTasks
	SetName(v string) *UpdateCloudGtmAddressRequest
	GetName() *string
}

type UpdateCloudGtmAddressRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh-CN: Chinese
	//
	// 	- en-US (default): English
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The IP address or domain name.
	//
	// example:
	//
	// 223.5.XX.XX
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The ID of the address. This ID uniquely identifies the address.
	//
	// This parameter is required.
	//
	// example:
	//
	// addr-89518218114368**92
	AddressId *string `json:"AddressId,omitempty" xml:"AddressId,omitempty"`
	// Address Attribution information.
	//
	// example:
	//
	// This parameter is not supported in the current version and does not need to be input.
	AttributeInfo *string `json:"AttributeInfo,omitempty" xml:"AttributeInfo,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can specify a custom value for this parameter, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24**22
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The new condition for determining the health state of the address. Valid values:
	//
	// 	- any_ok: The health check results of at least one health check template are normal.
	//
	// 	- p30_ok: The health check results of at least 30% of health check templates are normal.
	//
	// 	- p50_ok: The health check results of at least 50% of health check templates are normal.
	//
	// 	- p70_ok: The health check results of at least 70% of health check templates are normal.
	//
	// 	- all_ok: The health check results of all health check templates are normal.
	//
	// example:
	//
	// p50_ok
	HealthJudgement *string `json:"HealthJudgement,omitempty" xml:"HealthJudgement,omitempty"`
	// The health check tasks.
	HealthTasks []*UpdateCloudGtmAddressRequestHealthTasks `json:"HealthTasks,omitempty" xml:"HealthTasks,omitempty" type:"Repeated"`
	// The name of the address.
	//
	// example:
	//
	// Address-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateCloudGtmAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmAddressRequest) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmAddressRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateCloudGtmAddressRequest) GetAddress() *string {
	return s.Address
}

func (s *UpdateCloudGtmAddressRequest) GetAddressId() *string {
	return s.AddressId
}

func (s *UpdateCloudGtmAddressRequest) GetAttributeInfo() *string {
	return s.AttributeInfo
}

func (s *UpdateCloudGtmAddressRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCloudGtmAddressRequest) GetHealthJudgement() *string {
	return s.HealthJudgement
}

func (s *UpdateCloudGtmAddressRequest) GetHealthTasks() []*UpdateCloudGtmAddressRequestHealthTasks {
	return s.HealthTasks
}

func (s *UpdateCloudGtmAddressRequest) GetName() *string {
	return s.Name
}

func (s *UpdateCloudGtmAddressRequest) SetAcceptLanguage(v string) *UpdateCloudGtmAddressRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateCloudGtmAddressRequest) SetAddress(v string) *UpdateCloudGtmAddressRequest {
	s.Address = &v
	return s
}

func (s *UpdateCloudGtmAddressRequest) SetAddressId(v string) *UpdateCloudGtmAddressRequest {
	s.AddressId = &v
	return s
}

func (s *UpdateCloudGtmAddressRequest) SetAttributeInfo(v string) *UpdateCloudGtmAddressRequest {
	s.AttributeInfo = &v
	return s
}

func (s *UpdateCloudGtmAddressRequest) SetClientToken(v string) *UpdateCloudGtmAddressRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCloudGtmAddressRequest) SetHealthJudgement(v string) *UpdateCloudGtmAddressRequest {
	s.HealthJudgement = &v
	return s
}

func (s *UpdateCloudGtmAddressRequest) SetHealthTasks(v []*UpdateCloudGtmAddressRequestHealthTasks) *UpdateCloudGtmAddressRequest {
	s.HealthTasks = v
	return s
}

func (s *UpdateCloudGtmAddressRequest) SetName(v string) *UpdateCloudGtmAddressRequest {
	s.Name = &v
	return s
}

func (s *UpdateCloudGtmAddressRequest) Validate() error {
	if s.HealthTasks != nil {
		for _, item := range s.HealthTasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateCloudGtmAddressRequestHealthTasks struct {
	// The service port of the address on which health check tasks are performed. If the ping protocol is used for health checks, the configuration of the service port is not supported.
	//
	// 	- If you leave this parameter empty, the existing service port is deleted.
	//
	// 	- If you specify this parameter, the existing service port is updated based on the value of this parameter.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the health check template that is associated with the address. This parameter is required if you specify a service port of the address for health check tasks.
	//
	// 	- If you leave this parameter empty, the associated health check template is disassociated from the address.
	//
	// 	- If you specify this parameter, the associated health check template is updated based on the value of this parameter.
	//
	// example:
	//
	// mtp-89518052425100**80
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s UpdateCloudGtmAddressRequestHealthTasks) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmAddressRequestHealthTasks) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmAddressRequestHealthTasks) GetPort() *int32 {
	return s.Port
}

func (s *UpdateCloudGtmAddressRequestHealthTasks) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateCloudGtmAddressRequestHealthTasks) SetPort(v int32) *UpdateCloudGtmAddressRequestHealthTasks {
	s.Port = &v
	return s
}

func (s *UpdateCloudGtmAddressRequestHealthTasks) SetTemplateId(v string) *UpdateCloudGtmAddressRequestHealthTasks {
	s.TemplateId = &v
	return s
}

func (s *UpdateCloudGtmAddressRequestHealthTasks) Validate() error {
	return dara.Validate(s)
}
