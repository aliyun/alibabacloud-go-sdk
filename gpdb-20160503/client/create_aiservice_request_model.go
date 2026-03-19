// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAIServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *CreateAIServiceRequest
	GetDBInstanceId() *string
	SetDescription(v string) *CreateAIServiceRequest
	GetDescription() *string
	SetSecurityIPList(v string) *CreateAIServiceRequest
	GetSecurityIPList() *string
	SetServiceAccount(v string) *CreateAIServiceRequest
	GetServiceAccount() *string
	SetServiceAccountPassword(v string) *CreateAIServiceRequest
	GetServiceAccountPassword() *string
	SetType(v string) *CreateAIServiceRequest
	GetType() *string
}

type CreateAIServiceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// dramatest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 127.0.0.1
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dramauser
	ServiceAccount *string `json:"ServiceAccount,omitempty" xml:"ServiceAccount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456Aa!
	ServiceAccountPassword *string `json:"ServiceAccountPassword,omitempty" xml:"ServiceAccountPassword,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// drama
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateAIServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAIServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateAIServiceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateAIServiceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAIServiceRequest) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *CreateAIServiceRequest) GetServiceAccount() *string {
	return s.ServiceAccount
}

func (s *CreateAIServiceRequest) GetServiceAccountPassword() *string {
	return s.ServiceAccountPassword
}

func (s *CreateAIServiceRequest) GetType() *string {
	return s.Type
}

func (s *CreateAIServiceRequest) SetDBInstanceId(v string) *CreateAIServiceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateAIServiceRequest) SetDescription(v string) *CreateAIServiceRequest {
	s.Description = &v
	return s
}

func (s *CreateAIServiceRequest) SetSecurityIPList(v string) *CreateAIServiceRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CreateAIServiceRequest) SetServiceAccount(v string) *CreateAIServiceRequest {
	s.ServiceAccount = &v
	return s
}

func (s *CreateAIServiceRequest) SetServiceAccountPassword(v string) *CreateAIServiceRequest {
	s.ServiceAccountPassword = &v
	return s
}

func (s *CreateAIServiceRequest) SetType(v string) *CreateAIServiceRequest {
	s.Type = &v
	return s
}

func (s *CreateAIServiceRequest) Validate() error {
	return dara.Validate(s)
}
