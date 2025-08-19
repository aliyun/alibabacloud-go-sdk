// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeregisterDelegatedAdministratorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *DeregisterDelegatedAdministratorRequest
	GetAccountId() *string
	SetServicePrincipal(v string) *DeregisterDelegatedAdministratorRequest
	GetServicePrincipal() *string
}

type DeregisterDelegatedAdministratorRequest struct {
	// The Alibaba Cloud account ID of the member in the resource directory.
	//
	// This parameter is required.
	//
	// example:
	//
	// 181761095690****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The identifier of the trusted service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cloudfw.aliyuncs.com
	ServicePrincipal *string `json:"ServicePrincipal,omitempty" xml:"ServicePrincipal,omitempty"`
}

func (s DeregisterDelegatedAdministratorRequest) String() string {
	return dara.Prettify(s)
}

func (s DeregisterDelegatedAdministratorRequest) GoString() string {
	return s.String()
}

func (s *DeregisterDelegatedAdministratorRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *DeregisterDelegatedAdministratorRequest) GetServicePrincipal() *string {
	return s.ServicePrincipal
}

func (s *DeregisterDelegatedAdministratorRequest) SetAccountId(v string) *DeregisterDelegatedAdministratorRequest {
	s.AccountId = &v
	return s
}

func (s *DeregisterDelegatedAdministratorRequest) SetServicePrincipal(v string) *DeregisterDelegatedAdministratorRequest {
	s.ServicePrincipal = &v
	return s
}

func (s *DeregisterDelegatedAdministratorRequest) Validate() error {
	return dara.Validate(s)
}
