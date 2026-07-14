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
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the details of all AnalyticDB for PostgreSQL instances in a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The description.
	//
	// example:
	//
	// dramatest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The list of IP addresses in IP address whitelist group. You can specify up to 1,000 IP addresses, separated by commas (,). The value 127.0.0.1 indicates that no external IP addresses are allowed to access the instance. The following formats are supported:
	//
	// - 10.23.12.24 (IP address)
	//
	// - 10.23.12.24/24 (CIDR block. The value /24 indicates the length of the prefix in the address, which ranges from 1 to 32.)
	//
	// > After the service is created, you can call the ModifyAIServiceSecurityIps operation to modify IP address whitelist.
	//
	// example:
	//
	// 127.0.0.1
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The service account. The following limits apply:
	//
	// - The account name can contain lowercase letters, digits, and underscores (_).
	//
	// - The account name must start with a lowercase letter and end with a lowercase letter or digit.
	//
	// - The account name cannot start with gp.
	//
	// - The account name must be 2 to 16 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// dramauser
	ServiceAccount *string `json:"ServiceAccount,omitempty" xml:"ServiceAccount,omitempty"`
	// The password of the service account. The following limits apply:
	//
	// - The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// - Supported special characters: !@#$%^&*()_+-=
	//
	// - The password must be 8 to 32 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456Aa!
	ServiceAccountPassword *string `json:"ServiceAccountPassword,omitempty" xml:"ServiceAccountPassword,omitempty"`
	// The service type. Currently, only drama is supported.
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
