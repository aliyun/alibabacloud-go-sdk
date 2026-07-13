// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGtmInstanceGlobalConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertGroup(v string) *UpdateGtmInstanceGlobalConfigRequest
	GetAlertGroup() *string
	SetCnameCustomDomainName(v string) *UpdateGtmInstanceGlobalConfigRequest
	GetCnameCustomDomainName() *string
	SetCnameMode(v string) *UpdateGtmInstanceGlobalConfigRequest
	GetCnameMode() *string
	SetInstanceId(v string) *UpdateGtmInstanceGlobalConfigRequest
	GetInstanceId() *string
	SetInstanceName(v string) *UpdateGtmInstanceGlobalConfigRequest
	GetInstanceName() *string
	SetLang(v string) *UpdateGtmInstanceGlobalConfigRequest
	GetLang() *string
	SetLbaStrategy(v string) *UpdateGtmInstanceGlobalConfigRequest
	GetLbaStrategy() *string
	SetTtl(v int32) *UpdateGtmInstanceGlobalConfigRequest
	GetTtl() *int32
	SetUserDomainName(v string) *UpdateGtmInstanceGlobalConfigRequest
	GetUserDomainName() *string
}

type UpdateGtmInstanceGlobalConfigRequest struct {
	// The alert contact group. Only one alert contact group is supported.
	//
	// > This parameter is required when you update the instance for the first time. It is optional for subsequent updates.
	//
	// example:
	//
	// [\\"研发组\\"]
	AlertGroup *string `json:"AlertGroup,omitempty" xml:"AlertGroup,omitempty"`
	// This parameter is required when you set **CnameMode*	- to **CUSTOM**. The value must be the primary domain name.
	//
	// example:
	//
	// dns-example.top
	CnameCustomDomainName *string `json:"CnameCustomDomainName,omitempty" xml:"CnameCustomDomainName,omitempty"`
	// The connection type. Valid values:
	//
	// - **SYSTEM_ASSIGN**: system-assigned
	//
	// - **CUSTOM**: custom
	//
	// example:
	//
	// SYSTEM_ASSIGN
	CnameMode *string `json:"CnameMode,omitempty" xml:"CnameMode,omitempty"`
	// The ID of the GTM instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// gtm-cn-cs02xyk4a**
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance.
	//
	// > This parameter is required when you update the instance for the first time. It is optional for subsequent updates.
	//
	// example:
	//
	// 测试实例
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The load balancing policy. Valid values:
	//
	// - **ALL_RR**: round-robin
	//
	// - **RATIO**: weighted round-robin
	//
	// > This parameter is required when you update the instance for the first time. It is optional for subsequent updates.
	//
	// example:
	//
	// RATIO
	LbaStrategy *string `json:"LbaStrategy,omitempty" xml:"LbaStrategy,omitempty"`
	// The global Time to Live (TTL).
	//
	// example:
	//
	// 60
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The primary domain name.
	//
	// > This parameter is required when you update the instance for the first time. It is optional for subsequent updates.
	//
	// example:
	//
	// dns-example.top
	UserDomainName *string `json:"UserDomainName,omitempty" xml:"UserDomainName,omitempty"`
}

func (s UpdateGtmInstanceGlobalConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGtmInstanceGlobalConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateGtmInstanceGlobalConfigRequest) GetAlertGroup() *string {
	return s.AlertGroup
}

func (s *UpdateGtmInstanceGlobalConfigRequest) GetCnameCustomDomainName() *string {
	return s.CnameCustomDomainName
}

func (s *UpdateGtmInstanceGlobalConfigRequest) GetCnameMode() *string {
	return s.CnameMode
}

func (s *UpdateGtmInstanceGlobalConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateGtmInstanceGlobalConfigRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *UpdateGtmInstanceGlobalConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateGtmInstanceGlobalConfigRequest) GetLbaStrategy() *string {
	return s.LbaStrategy
}

func (s *UpdateGtmInstanceGlobalConfigRequest) GetTtl() *int32 {
	return s.Ttl
}

func (s *UpdateGtmInstanceGlobalConfigRequest) GetUserDomainName() *string {
	return s.UserDomainName
}

func (s *UpdateGtmInstanceGlobalConfigRequest) SetAlertGroup(v string) *UpdateGtmInstanceGlobalConfigRequest {
	s.AlertGroup = &v
	return s
}

func (s *UpdateGtmInstanceGlobalConfigRequest) SetCnameCustomDomainName(v string) *UpdateGtmInstanceGlobalConfigRequest {
	s.CnameCustomDomainName = &v
	return s
}

func (s *UpdateGtmInstanceGlobalConfigRequest) SetCnameMode(v string) *UpdateGtmInstanceGlobalConfigRequest {
	s.CnameMode = &v
	return s
}

func (s *UpdateGtmInstanceGlobalConfigRequest) SetInstanceId(v string) *UpdateGtmInstanceGlobalConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateGtmInstanceGlobalConfigRequest) SetInstanceName(v string) *UpdateGtmInstanceGlobalConfigRequest {
	s.InstanceName = &v
	return s
}

func (s *UpdateGtmInstanceGlobalConfigRequest) SetLang(v string) *UpdateGtmInstanceGlobalConfigRequest {
	s.Lang = &v
	return s
}

func (s *UpdateGtmInstanceGlobalConfigRequest) SetLbaStrategy(v string) *UpdateGtmInstanceGlobalConfigRequest {
	s.LbaStrategy = &v
	return s
}

func (s *UpdateGtmInstanceGlobalConfigRequest) SetTtl(v int32) *UpdateGtmInstanceGlobalConfigRequest {
	s.Ttl = &v
	return s
}

func (s *UpdateGtmInstanceGlobalConfigRequest) SetUserDomainName(v string) *UpdateGtmInstanceGlobalConfigRequest {
	s.UserDomainName = &v
	return s
}

func (s *UpdateGtmInstanceGlobalConfigRequest) Validate() error {
	return dara.Validate(s)
}
