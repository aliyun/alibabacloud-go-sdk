// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudGtmAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *CreateCloudGtmAddressRequest
	GetAcceptLanguage() *string
	SetAddress(v string) *CreateCloudGtmAddressRequest
	GetAddress() *string
	SetAttributeInfo(v string) *CreateCloudGtmAddressRequest
	GetAttributeInfo() *string
	SetAvailableMode(v string) *CreateCloudGtmAddressRequest
	GetAvailableMode() *string
	SetClientToken(v string) *CreateCloudGtmAddressRequest
	GetClientToken() *string
	SetEnableStatus(v string) *CreateCloudGtmAddressRequest
	GetEnableStatus() *string
	SetHealthJudgement(v string) *CreateCloudGtmAddressRequest
	GetHealthJudgement() *string
	SetHealthTasks(v []*CreateCloudGtmAddressRequestHealthTasks) *CreateCloudGtmAddressRequest
	GetHealthTasks() []*CreateCloudGtmAddressRequestHealthTasks
	SetManualAvailableStatus(v string) *CreateCloudGtmAddressRequest
	GetManualAvailableStatus() *string
	SetName(v string) *CreateCloudGtmAddressRequest
	GetName() *string
	SetRemark(v string) *CreateCloudGtmAddressRequest
	GetRemark() *string
	SetType(v string) *CreateCloudGtmAddressRequest
	GetType() *string
}

type CreateCloudGtmAddressRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh-CN: Chinese
	//
	// 	- en-US (default): English
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// IP address or domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 223.5.XX.XX
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// Address ownership information.
	//
	// example:
	//
	// This parameter is not supported in the version. Do not enter this parameter
	AttributeInfo *string `json:"AttributeInfo,omitempty" xml:"AttributeInfo,omitempty"`
	// The failover mode that is used when address exceptions are identified. Valid values:
	//
	// 	- auto: the automatic mode. The system determines whether to return an address based on the health check results. If the address fails health checks, the system does not return the address. If the address passes health checks, the system returns the address.
	//
	// 	- manual: the manual mode. If an address is in the unavailable state, the address is not returned for DNS requests even if the address passes health checks. If an address is in the available state, the address is returned for DNS requests even if an alert is triggered when the address fails health checks.
	//
	// This parameter is required.
	//
	// example:
	//
	// auto
	AvailableMode *string `json:"AvailableMode,omitempty" xml:"AvailableMode,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24**22
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Indicates the current enabled status of the address:
	//
	// - enable: Enabled status
	//
	// - disable: Disabled status
	//
	// This parameter is required.
	//
	// example:
	//
	// enable
	EnableStatus *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// The condition for determining the health status of the address. This parameter is required when HealthTasks is specified. Valid values:
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
	// This parameter is required.
	//
	// example:
	//
	// p50_ok
	HealthJudgement *string `json:"HealthJudgement,omitempty" xml:"HealthJudgement,omitempty"`
	// The health check tasks associated with the address.
	HealthTasks []*CreateCloudGtmAddressRequestHealthTasks `json:"HealthTasks,omitempty" xml:"HealthTasks,omitempty" type:"Repeated"`
	// The availability state of the address. This parameter is required when AvailableMode is set to **manual**. Valid values:
	//
	// 	- available: The address is normal. In this state, the address is returned for DNS requests even if an alert is triggered when the address fails health checks.
	//
	// 	- unavailable: The address is abnormal. In this state, the address is not returned for DNS requests even if the address passes health checks.
	//
	// example:
	//
	// available
	ManualAvailableStatus *string `json:"ManualAvailableStatus,omitempty" xml:"ManualAvailableStatus,omitempty"`
	// Address name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Address-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Remarks.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// Address type:
	//
	// - IPv4
	//
	// - IPv6
	//
	// - domain
	//
	// This parameter is required.
	//
	// example:
	//
	// IPv4
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateCloudGtmAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudGtmAddressRequest) GoString() string {
	return s.String()
}

func (s *CreateCloudGtmAddressRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *CreateCloudGtmAddressRequest) GetAddress() *string {
	return s.Address
}

func (s *CreateCloudGtmAddressRequest) GetAttributeInfo() *string {
	return s.AttributeInfo
}

func (s *CreateCloudGtmAddressRequest) GetAvailableMode() *string {
	return s.AvailableMode
}

func (s *CreateCloudGtmAddressRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateCloudGtmAddressRequest) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *CreateCloudGtmAddressRequest) GetHealthJudgement() *string {
	return s.HealthJudgement
}

func (s *CreateCloudGtmAddressRequest) GetHealthTasks() []*CreateCloudGtmAddressRequestHealthTasks {
	return s.HealthTasks
}

func (s *CreateCloudGtmAddressRequest) GetManualAvailableStatus() *string {
	return s.ManualAvailableStatus
}

func (s *CreateCloudGtmAddressRequest) GetName() *string {
	return s.Name
}

func (s *CreateCloudGtmAddressRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateCloudGtmAddressRequest) GetType() *string {
	return s.Type
}

func (s *CreateCloudGtmAddressRequest) SetAcceptLanguage(v string) *CreateCloudGtmAddressRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *CreateCloudGtmAddressRequest) SetAddress(v string) *CreateCloudGtmAddressRequest {
	s.Address = &v
	return s
}

func (s *CreateCloudGtmAddressRequest) SetAttributeInfo(v string) *CreateCloudGtmAddressRequest {
	s.AttributeInfo = &v
	return s
}

func (s *CreateCloudGtmAddressRequest) SetAvailableMode(v string) *CreateCloudGtmAddressRequest {
	s.AvailableMode = &v
	return s
}

func (s *CreateCloudGtmAddressRequest) SetClientToken(v string) *CreateCloudGtmAddressRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCloudGtmAddressRequest) SetEnableStatus(v string) *CreateCloudGtmAddressRequest {
	s.EnableStatus = &v
	return s
}

func (s *CreateCloudGtmAddressRequest) SetHealthJudgement(v string) *CreateCloudGtmAddressRequest {
	s.HealthJudgement = &v
	return s
}

func (s *CreateCloudGtmAddressRequest) SetHealthTasks(v []*CreateCloudGtmAddressRequestHealthTasks) *CreateCloudGtmAddressRequest {
	s.HealthTasks = v
	return s
}

func (s *CreateCloudGtmAddressRequest) SetManualAvailableStatus(v string) *CreateCloudGtmAddressRequest {
	s.ManualAvailableStatus = &v
	return s
}

func (s *CreateCloudGtmAddressRequest) SetName(v string) *CreateCloudGtmAddressRequest {
	s.Name = &v
	return s
}

func (s *CreateCloudGtmAddressRequest) SetRemark(v string) *CreateCloudGtmAddressRequest {
	s.Remark = &v
	return s
}

func (s *CreateCloudGtmAddressRequest) SetType(v string) *CreateCloudGtmAddressRequest {
	s.Type = &v
	return s
}

func (s *CreateCloudGtmAddressRequest) Validate() error {
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

type CreateCloudGtmAddressRequestHealthTasks struct {
	// The service port of the address on which health check tasks are performed. If the ping protocol is used for health checks, the configuration of the service port is not supported.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the health check template associated with the address.
	//
	// example:
	//
	// mtp-89518052425100**80
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateCloudGtmAddressRequestHealthTasks) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudGtmAddressRequestHealthTasks) GoString() string {
	return s.String()
}

func (s *CreateCloudGtmAddressRequestHealthTasks) GetPort() *int32 {
	return s.Port
}

func (s *CreateCloudGtmAddressRequestHealthTasks) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateCloudGtmAddressRequestHealthTasks) SetPort(v int32) *CreateCloudGtmAddressRequestHealthTasks {
	s.Port = &v
	return s
}

func (s *CreateCloudGtmAddressRequestHealthTasks) SetTemplateId(v string) *CreateCloudGtmAddressRequestHealthTasks {
	s.TemplateId = &v
	return s
}

func (s *CreateCloudGtmAddressRequestHealthTasks) Validate() error {
	return dara.Validate(s)
}
