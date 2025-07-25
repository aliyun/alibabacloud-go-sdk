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
	// The alert group. Only one alert group is supported.
	//
	// >  This parameter is required only for the first modification.
	AlertGroup *string `json:"AlertGroup,omitempty" xml:"AlertGroup,omitempty"`
	// If you set **CnameMode*	- to **CUSTOM**, you must specify the CnameCustomDomainName parameter, which must be set to a primary domain name.
	//
	// example:
	//
	// www.example.com
	CnameCustomDomainName *string `json:"CnameCustomDomainName,omitempty" xml:"CnameCustomDomainName,omitempty"`
	// Specifies whether to use a system-assigned canonical name (CNAME) or a custom CNAME to access GTM. Valid values:
	//
	// 	- **SYSTEM_ASSIGN**: system-assigned CNAME
	//
	// 	- **CUSTOM**: custom CNAME
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
	// instance1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the GTM instance.
	//
	// >  This parameter is required only for the first modification.
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The balancing policy. Valid values:
	//
	// 	- **ALL_RR**: load balancing
	//
	// 	- **RATIO**: weighted round-robin
	//
	// >  This parameter is required only for the first modification.
	//
	// example:
	//
	// RATIO
	LbaStrategy *string `json:"LbaStrategy,omitempty" xml:"LbaStrategy,omitempty"`
	// The global time-to-live (TTL).
	//
	// example:
	//
	// 60
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The primary domain name.
	//
	// >  This parameter is required only for the first modification.
	//
	// example:
	//
	// www.example.com
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
