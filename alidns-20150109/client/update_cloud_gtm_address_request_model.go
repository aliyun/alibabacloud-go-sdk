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
	// - zh-CN: Chinese
	//
	// - en-US (default): English
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The updated IP address or domain name.
	//
	// example:
	//
	// 223.5.XX.XX
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The unique ID of the address.
	//
	// This parameter is required.
	//
	// example:
	//
	// addr-89518218114368****
	AddressId *string `json:"AddressId,omitempty" xml:"AddressId,omitempty"`
	// The attribution information of the address.
	//
	// example:
	//
	// 当前版本不支持此参数，不需要传入此参数。
	AttributeInfo *string `json:"AttributeInfo,omitempty" xml:"AttributeInfo,omitempty"`
	// A client token that is used to ensure the idempotence of the request. You can specify a custom value for this parameter, but you must make sure that the value is unique among different requests. The value can contain a maximum of 64 ASCII characters.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The updated condition for determining the health status of the address:
	//
	// - any_ok: At least one probe is normal for all health check templates.
	//
	// - p30_ok: At least 30% of the probes are normal for all health check templates.
	//
	// - p50_ok: At least 50% of the probes are normal for all health check templates.
	//
	// - p70_ok: At least 70% of the probes are normal for all health check templates.
	//
	// - all_ok: All probes are normal for all health check templates.
	//
	// example:
	//
	// p50_ok
	HealthJudgement *string `json:"HealthJudgement,omitempty" xml:"HealthJudgement,omitempty"`
	// The list of health check tasks.
	HealthTasks []*UpdateCloudGtmAddressRequestHealthTasks `json:"HealthTasks,omitempty" xml:"HealthTasks,omitempty" type:"Repeated"`
	// The updated name of the address.
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
	// The service port of the target address for the health check. You cannot configure a service port if the health check uses the ping protocol.
	//
	// - If you leave this parameter empty, the currently configured port is deleted.
	//
	// - If you specify a value for this parameter, the port is updated to the specified value.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the health check template associated with the address. This parameter is required if you configure a health check port.
	//
	// - If you leave this parameter empty, the currently configured health check template is deleted.
	//
	// - If you specify a value for this parameter, the health check template is updated to the specified value.
	//
	// example:
	//
	// mtp-89518052425100****
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
