// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterDelegatedAdministratorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *RegisterDelegatedAdministratorRequest
	GetAccountId() *string
	SetServicePrincipal(v string) *RegisterDelegatedAdministratorRequest
	GetServicePrincipal() *string
}

type RegisterDelegatedAdministratorRequest struct {
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
	// For more information, see the `Trusted service identifier` column in [Supported trusted services](https://help.aliyun.com/document_detail/208133.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cloudfw.aliyuncs.com
	ServicePrincipal *string `json:"ServicePrincipal,omitempty" xml:"ServicePrincipal,omitempty"`
}

func (s RegisterDelegatedAdministratorRequest) String() string {
	return dara.Prettify(s)
}

func (s RegisterDelegatedAdministratorRequest) GoString() string {
	return s.String()
}

func (s *RegisterDelegatedAdministratorRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *RegisterDelegatedAdministratorRequest) GetServicePrincipal() *string {
	return s.ServicePrincipal
}

func (s *RegisterDelegatedAdministratorRequest) SetAccountId(v string) *RegisterDelegatedAdministratorRequest {
	s.AccountId = &v
	return s
}

func (s *RegisterDelegatedAdministratorRequest) SetServicePrincipal(v string) *RegisterDelegatedAdministratorRequest {
	s.ServicePrincipal = &v
	return s
}

func (s *RegisterDelegatedAdministratorRequest) Validate() error {
	return dara.Validate(s)
}
