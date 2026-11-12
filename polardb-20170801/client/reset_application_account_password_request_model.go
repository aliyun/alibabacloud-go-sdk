// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetApplicationAccountPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *ResetApplicationAccountPasswordRequest
	GetAccountName() *string
	SetApplicationId(v string) *ResetApplicationAccountPasswordRequest
	GetApplicationId() *string
	SetNewAccountPassword(v string) *ResetApplicationAccountPasswordRequest
	GetNewAccountPassword() *string
}

type ResetApplicationAccountPasswordRequest struct {
	// The account name.
	//
	// This parameter is required.
	//
	// example:
	//
	// mcp_user
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The application ID. Only applications associated with PolarDB for MySQL clusters are accepted. Before the call, you can obtain the DBClusterId associated with the application by calling DescribeApplicationAttribute, and then query the DBType of the cluster by calling DescribeDBClusterAttribute to confirm that the value is MySQL. Associated applications of other engine types are not applicable.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-bp1abcdef1234567
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The new password.
	//
	// This parameter is required.
	//
	// example:
	//
	// Pw123456
	NewAccountPassword *string `json:"NewAccountPassword,omitempty" xml:"NewAccountPassword,omitempty"`
}

func (s ResetApplicationAccountPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetApplicationAccountPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetApplicationAccountPasswordRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *ResetApplicationAccountPasswordRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ResetApplicationAccountPasswordRequest) GetNewAccountPassword() *string {
	return s.NewAccountPassword
}

func (s *ResetApplicationAccountPasswordRequest) SetAccountName(v string) *ResetApplicationAccountPasswordRequest {
	s.AccountName = &v
	return s
}

func (s *ResetApplicationAccountPasswordRequest) SetApplicationId(v string) *ResetApplicationAccountPasswordRequest {
	s.ApplicationId = &v
	return s
}

func (s *ResetApplicationAccountPasswordRequest) SetNewAccountPassword(v string) *ResetApplicationAccountPasswordRequest {
	s.NewAccountPassword = &v
	return s
}

func (s *ResetApplicationAccountPasswordRequest) Validate() error {
	return dara.Validate(s)
}
