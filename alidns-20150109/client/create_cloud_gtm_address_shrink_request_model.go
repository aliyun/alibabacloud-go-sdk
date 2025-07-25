// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudGtmAddressShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *CreateCloudGtmAddressShrinkRequest
	GetAcceptLanguage() *string
	SetAddress(v string) *CreateCloudGtmAddressShrinkRequest
	GetAddress() *string
	SetAttributeInfo(v string) *CreateCloudGtmAddressShrinkRequest
	GetAttributeInfo() *string
	SetAvailableMode(v string) *CreateCloudGtmAddressShrinkRequest
	GetAvailableMode() *string
	SetClientToken(v string) *CreateCloudGtmAddressShrinkRequest
	GetClientToken() *string
	SetEnableStatus(v string) *CreateCloudGtmAddressShrinkRequest
	GetEnableStatus() *string
	SetHealthJudgement(v string) *CreateCloudGtmAddressShrinkRequest
	GetHealthJudgement() *string
	SetHealthTasksShrink(v string) *CreateCloudGtmAddressShrinkRequest
	GetHealthTasksShrink() *string
	SetManualAvailableStatus(v string) *CreateCloudGtmAddressShrinkRequest
	GetManualAvailableStatus() *string
	SetName(v string) *CreateCloudGtmAddressShrinkRequest
	GetName() *string
	SetRemark(v string) *CreateCloudGtmAddressShrinkRequest
	GetRemark() *string
	SetType(v string) *CreateCloudGtmAddressShrinkRequest
	GetType() *string
}

type CreateCloudGtmAddressShrinkRequest struct {
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
	// 当前版本不支持传入此参数，请不要传入参数。
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
	HealthTasksShrink *string `json:"HealthTasks,omitempty" xml:"HealthTasks,omitempty"`
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

func (s CreateCloudGtmAddressShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudGtmAddressShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateCloudGtmAddressShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *CreateCloudGtmAddressShrinkRequest) GetAddress() *string {
	return s.Address
}

func (s *CreateCloudGtmAddressShrinkRequest) GetAttributeInfo() *string {
	return s.AttributeInfo
}

func (s *CreateCloudGtmAddressShrinkRequest) GetAvailableMode() *string {
	return s.AvailableMode
}

func (s *CreateCloudGtmAddressShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateCloudGtmAddressShrinkRequest) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *CreateCloudGtmAddressShrinkRequest) GetHealthJudgement() *string {
	return s.HealthJudgement
}

func (s *CreateCloudGtmAddressShrinkRequest) GetHealthTasksShrink() *string {
	return s.HealthTasksShrink
}

func (s *CreateCloudGtmAddressShrinkRequest) GetManualAvailableStatus() *string {
	return s.ManualAvailableStatus
}

func (s *CreateCloudGtmAddressShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateCloudGtmAddressShrinkRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateCloudGtmAddressShrinkRequest) GetType() *string {
	return s.Type
}

func (s *CreateCloudGtmAddressShrinkRequest) SetAcceptLanguage(v string) *CreateCloudGtmAddressShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *CreateCloudGtmAddressShrinkRequest) SetAddress(v string) *CreateCloudGtmAddressShrinkRequest {
	s.Address = &v
	return s
}

func (s *CreateCloudGtmAddressShrinkRequest) SetAttributeInfo(v string) *CreateCloudGtmAddressShrinkRequest {
	s.AttributeInfo = &v
	return s
}

func (s *CreateCloudGtmAddressShrinkRequest) SetAvailableMode(v string) *CreateCloudGtmAddressShrinkRequest {
	s.AvailableMode = &v
	return s
}

func (s *CreateCloudGtmAddressShrinkRequest) SetClientToken(v string) *CreateCloudGtmAddressShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCloudGtmAddressShrinkRequest) SetEnableStatus(v string) *CreateCloudGtmAddressShrinkRequest {
	s.EnableStatus = &v
	return s
}

func (s *CreateCloudGtmAddressShrinkRequest) SetHealthJudgement(v string) *CreateCloudGtmAddressShrinkRequest {
	s.HealthJudgement = &v
	return s
}

func (s *CreateCloudGtmAddressShrinkRequest) SetHealthTasksShrink(v string) *CreateCloudGtmAddressShrinkRequest {
	s.HealthTasksShrink = &v
	return s
}

func (s *CreateCloudGtmAddressShrinkRequest) SetManualAvailableStatus(v string) *CreateCloudGtmAddressShrinkRequest {
	s.ManualAvailableStatus = &v
	return s
}

func (s *CreateCloudGtmAddressShrinkRequest) SetName(v string) *CreateCloudGtmAddressShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateCloudGtmAddressShrinkRequest) SetRemark(v string) *CreateCloudGtmAddressShrinkRequest {
	s.Remark = &v
	return s
}

func (s *CreateCloudGtmAddressShrinkRequest) SetType(v string) *CreateCloudGtmAddressShrinkRequest {
	s.Type = &v
	return s
}

func (s *CreateCloudGtmAddressShrinkRequest) Validate() error {
	return dara.Validate(s)
}
