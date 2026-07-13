// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmAddressShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateCloudGtmAddressShrinkRequest
	GetAcceptLanguage() *string
	SetAddress(v string) *UpdateCloudGtmAddressShrinkRequest
	GetAddress() *string
	SetAddressId(v string) *UpdateCloudGtmAddressShrinkRequest
	GetAddressId() *string
	SetAttributeInfo(v string) *UpdateCloudGtmAddressShrinkRequest
	GetAttributeInfo() *string
	SetClientToken(v string) *UpdateCloudGtmAddressShrinkRequest
	GetClientToken() *string
	SetHealthJudgement(v string) *UpdateCloudGtmAddressShrinkRequest
	GetHealthJudgement() *string
	SetHealthTasksShrink(v string) *UpdateCloudGtmAddressShrinkRequest
	GetHealthTasksShrink() *string
	SetName(v string) *UpdateCloudGtmAddressShrinkRequest
	GetName() *string
}

type UpdateCloudGtmAddressShrinkRequest struct {
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
	HealthTasksShrink *string `json:"HealthTasks,omitempty" xml:"HealthTasks,omitempty"`
	// The updated name of the address.
	//
	// example:
	//
	// Address-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateCloudGtmAddressShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmAddressShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmAddressShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateCloudGtmAddressShrinkRequest) GetAddress() *string {
	return s.Address
}

func (s *UpdateCloudGtmAddressShrinkRequest) GetAddressId() *string {
	return s.AddressId
}

func (s *UpdateCloudGtmAddressShrinkRequest) GetAttributeInfo() *string {
	return s.AttributeInfo
}

func (s *UpdateCloudGtmAddressShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCloudGtmAddressShrinkRequest) GetHealthJudgement() *string {
	return s.HealthJudgement
}

func (s *UpdateCloudGtmAddressShrinkRequest) GetHealthTasksShrink() *string {
	return s.HealthTasksShrink
}

func (s *UpdateCloudGtmAddressShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateCloudGtmAddressShrinkRequest) SetAcceptLanguage(v string) *UpdateCloudGtmAddressShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateCloudGtmAddressShrinkRequest) SetAddress(v string) *UpdateCloudGtmAddressShrinkRequest {
	s.Address = &v
	return s
}

func (s *UpdateCloudGtmAddressShrinkRequest) SetAddressId(v string) *UpdateCloudGtmAddressShrinkRequest {
	s.AddressId = &v
	return s
}

func (s *UpdateCloudGtmAddressShrinkRequest) SetAttributeInfo(v string) *UpdateCloudGtmAddressShrinkRequest {
	s.AttributeInfo = &v
	return s
}

func (s *UpdateCloudGtmAddressShrinkRequest) SetClientToken(v string) *UpdateCloudGtmAddressShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCloudGtmAddressShrinkRequest) SetHealthJudgement(v string) *UpdateCloudGtmAddressShrinkRequest {
	s.HealthJudgement = &v
	return s
}

func (s *UpdateCloudGtmAddressShrinkRequest) SetHealthTasksShrink(v string) *UpdateCloudGtmAddressShrinkRequest {
	s.HealthTasksShrink = &v
	return s
}

func (s *UpdateCloudGtmAddressShrinkRequest) SetName(v string) *UpdateCloudGtmAddressShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateCloudGtmAddressShrinkRequest) Validate() error {
	return dara.Validate(s)
}
