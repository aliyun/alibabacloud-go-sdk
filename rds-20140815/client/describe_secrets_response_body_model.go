// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecretsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *DescribeSecretsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeSecretsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeSecretsResponseBody
	GetRequestId() *string
	SetSecrets(v []*DescribeSecretsResponseBodySecrets) *DescribeSecretsResponseBody
	GetSecrets() []*DescribeSecretsResponseBodySecrets
}

type DescribeSecretsResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// DF4961DD-16F5-5B24-BD4C-0C7788F7ADAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the credential.
	Secrets []*DescribeSecretsResponseBodySecrets `json:"Secrets,omitempty" xml:"Secrets,omitempty" type:"Repeated"`
}

func (s DescribeSecretsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecretsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecretsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeSecretsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeSecretsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSecretsResponseBody) GetSecrets() []*DescribeSecretsResponseBodySecrets {
	return s.Secrets
}

func (s *DescribeSecretsResponseBody) SetPageNumber(v int64) *DescribeSecretsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSecretsResponseBody) SetPageSize(v int64) *DescribeSecretsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSecretsResponseBody) SetRequestId(v string) *DescribeSecretsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecretsResponseBody) SetSecrets(v []*DescribeSecretsResponseBodySecrets) *DescribeSecretsResponseBody {
	s.Secrets = v
	return s
}

func (s *DescribeSecretsResponseBody) Validate() error {
	if s.Secrets != nil {
		for _, item := range s.Secrets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSecretsResponseBodySecrets struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 1266348003******
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The description of the credential.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the credential for the created Data API account.
	//
	// example:
	//
	// acs:rds:cn-hangzhou:1335786***:dbInstance/rm-bp1m7l3j63****
	SecretArn *string `json:"SecretArn,omitempty" xml:"SecretArn,omitempty"`
	// The name of the credential.
	//
	// example:
	//
	// Foo
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The username that is used to access the database.
	//
	// example:
	//
	// user_jack
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s DescribeSecretsResponseBodySecrets) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecretsResponseBodySecrets) GoString() string {
	return s.String()
}

func (s *DescribeSecretsResponseBodySecrets) GetAccountId() *string {
	return s.AccountId
}

func (s *DescribeSecretsResponseBodySecrets) GetDescription() *string {
	return s.Description
}

func (s *DescribeSecretsResponseBodySecrets) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSecretsResponseBodySecrets) GetSecretArn() *string {
	return s.SecretArn
}

func (s *DescribeSecretsResponseBodySecrets) GetSecretName() *string {
	return s.SecretName
}

func (s *DescribeSecretsResponseBodySecrets) GetUsername() *string {
	return s.Username
}

func (s *DescribeSecretsResponseBodySecrets) SetAccountId(v string) *DescribeSecretsResponseBodySecrets {
	s.AccountId = &v
	return s
}

func (s *DescribeSecretsResponseBodySecrets) SetDescription(v string) *DescribeSecretsResponseBodySecrets {
	s.Description = &v
	return s
}

func (s *DescribeSecretsResponseBodySecrets) SetRegionId(v string) *DescribeSecretsResponseBodySecrets {
	s.RegionId = &v
	return s
}

func (s *DescribeSecretsResponseBodySecrets) SetSecretArn(v string) *DescribeSecretsResponseBodySecrets {
	s.SecretArn = &v
	return s
}

func (s *DescribeSecretsResponseBodySecrets) SetSecretName(v string) *DescribeSecretsResponseBodySecrets {
	s.SecretName = &v
	return s
}

func (s *DescribeSecretsResponseBodySecrets) SetUsername(v string) *DescribeSecretsResponseBodySecrets {
	s.Username = &v
	return s
}

func (s *DescribeSecretsResponseBodySecrets) Validate() error {
	return dara.Validate(s)
}
