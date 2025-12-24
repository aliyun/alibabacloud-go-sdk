// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableInternalAuthenticationSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthenticationSourceId(v string) *DisableInternalAuthenticationSourceRequest
	GetAuthenticationSourceId() *string
	SetInstanceId(v string) *DisableInternalAuthenticationSourceRequest
	GetInstanceId() *string
}

type DisableInternalAuthenticationSourceRequest struct {
	// 内部认证源ID，比如 ia_password, ia_otp_sms 等
	//
	// example:
	//
	// ia_password
	AuthenticationSourceId *string `json:"AuthenticationSourceId,omitempty" xml:"AuthenticationSourceId,omitempty"`
	// IDaaS EIAM的实例id
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_111xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DisableInternalAuthenticationSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableInternalAuthenticationSourceRequest) GoString() string {
	return s.String()
}

func (s *DisableInternalAuthenticationSourceRequest) GetAuthenticationSourceId() *string {
	return s.AuthenticationSourceId
}

func (s *DisableInternalAuthenticationSourceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableInternalAuthenticationSourceRequest) SetAuthenticationSourceId(v string) *DisableInternalAuthenticationSourceRequest {
	s.AuthenticationSourceId = &v
	return s
}

func (s *DisableInternalAuthenticationSourceRequest) SetInstanceId(v string) *DisableInternalAuthenticationSourceRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableInternalAuthenticationSourceRequest) Validate() error {
	return dara.Validate(s)
}
