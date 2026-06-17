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
	// The instance ID.
	//
	// > Call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to list all AnalyticDB for PostgreSQL instances in the destination region, including their instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// A description of the service.
	//
	// example:
	//
	// dramatest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// A comma-separated list of IP addresses or CIDR blocks in the IP address whitelist. You can specify up to 1000 entries. To block all external IP access, use 127.0.0.1. Valid formats include the following:
	//
	// - 10.23.12.24 (an IPv4 address)
	//
	// - 10.23.12.24/24 (a CIDR block, where /24 indicates the prefix length, from 1 to 32)
	//
	// > After you create the service, call the ModifyAIServiceSecurityIps operation to update the IP address whitelist.
	//
	// example:
	//
	// 127.0.0.1
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The service account. It must meet these requirements:
	//
	// - Contain only lowercase letters, digits, and underscores.
	//
	// - Start with a lowercase letter and end with a lowercase letter or digit.
	//
	// - Not start with gp.
	//
	// - Be 2 to 16 characters long.
	//
	// This parameter is required.
	//
	// example:
	//
	// dramauser
	ServiceAccount *string `json:"ServiceAccount,omitempty" xml:"ServiceAccount,omitempty"`
	// The password for the service account. It must meet these requirements:
	//
	// - Contain at least three of the following: uppercase letters, lowercase letters, digits, and special characters.
	//
	// - Support these special characters: !@#$%^&\\*()_+-=.
	//
	// - Be 8 to 32 characters long.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456Aa!
	ServiceAccountPassword *string `json:"ServiceAccountPassword,omitempty" xml:"ServiceAccountPassword,omitempty"`
	// The service type. Only drama is supported.
	//
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
