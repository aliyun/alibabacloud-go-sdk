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
	// - zh-CN: Chinese.
	//
	// - en-US (default): English.
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The IP address or domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 223.5.XX.XX
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The attribution information of the address.
	//
	// example:
	//
	// 当前版本不支持传入此参数，请不要传入参数。
	AttributeInfo *string `json:"AttributeInfo,omitempty" xml:"AttributeInfo,omitempty"`
	// The switchover mode for the address when a health check is abnormal. Valid values:
	//
	// - auto: The system automatically manages the address status based on health check results. If an address is unhealthy, DNS resolution for it stops. If the address becomes healthy, DNS resolution resumes.
	//
	// - manual: You manually manage the address status. If you set an address to abnormal, DNS resolution for it stops. It does not resume even if the address becomes healthy. If you set an address to normal, DNS resolution for it resumes. If a healthy address becomes unhealthy, the system sends an alert but does not stop DNS resolution.
	//
	// This parameter is required.
	//
	// example:
	//
	// auto
	AvailableMode *string `json:"AvailableMode,omitempty" xml:"AvailableMode,omitempty"`
	// The client token that is used to ensure the idempotence of the request. Make sure that the token is unique for each request. The token can contain a maximum of 64 ASCII characters.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The status of the address. Valid values:
	//
	// - enable: The address is enabled.
	//
	// - disable: The address is disabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// enable
	EnableStatus *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// The condition for determining the health of the address. This parameter is required if you specify HealthTasks. Valid values:
	//
	// - any_ok: At least one health check is successful.
	//
	// - p30_ok: At least 30% of health checks are successful.
	//
	// - p50_ok: At least 50% of health checks are successful.
	//
	// - p70_ok: At least 70% of health checks are successful.
	//
	// - all_ok: All health checks are successful.
	//
	// This parameter is required.
	//
	// example:
	//
	// p50_ok
	HealthJudgement *string `json:"HealthJudgement,omitempty" xml:"HealthJudgement,omitempty"`
	// The health check tasks for the address.
	HealthTasks []*CreateCloudGtmAddressRequestHealthTasks `json:"HealthTasks,omitempty" xml:"HealthTasks,omitempty" type:"Repeated"`
	// The availability status of the address when the health check-based switchover mode is set to **manual**. Valid values:
	//
	// - available: The address is available. In this state, DNS resolution for the address is normal. If a health check is abnormal, the system only sends an alert and does not stop DNS resolution.
	//
	// - unavailable: The address is unavailable. In this state, DNS resolution for the address is stopped. DNS resolution is not resumed even if a health check is normal.
	//
	// example:
	//
	// available
	ManualAvailableStatus *string `json:"ManualAvailableStatus,omitempty" xml:"ManualAvailableStatus,omitempty"`
	// The name of the address.
	//
	// This parameter is required.
	//
	// example:
	//
	// Address-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The remarks about the address.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The type of the address. Valid values:
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
	// The service port of the destination address for the health check. This parameter is not supported for health checks that use the ping protocol.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the health check template.
	//
	// example:
	//
	// mtp-89518052425100****
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
