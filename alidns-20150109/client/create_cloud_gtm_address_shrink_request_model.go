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
	HealthTasksShrink *string `json:"HealthTasks,omitempty" xml:"HealthTasks,omitempty"`
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
