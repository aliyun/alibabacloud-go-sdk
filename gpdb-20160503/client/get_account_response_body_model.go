// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccountDescription(v string) *GetAccountResponseBody
	GetAccountDescription() *string
	SetAccountName(v string) *GetAccountResponseBody
	GetAccountName() *string
	SetAccountStatus(v string) *GetAccountResponseBody
	GetAccountStatus() *string
	SetAccountType(v string) *GetAccountResponseBody
	GetAccountType() *string
	SetDBInstanceId(v string) *GetAccountResponseBody
	GetDBInstanceId() *string
	SetRequestId(v string) *GetAccountResponseBody
	GetRequestId() *string
}

type GetAccountResponseBody struct {
	// The new description of the database account.
	//
	// 	- The description must start with a letter.
	//
	// 	- The description cannot start with `http://` or `https://`.
	//
	// 	- The description can contain letters, underscores (_), hyphens (-), and digits.
	//
	// 	- The description must be 2 to 256 characters in length.
	//
	// example:
	//
	// The instance used by this account to log in is DBInstanceId. The name used to log in is AccountName.
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// The name of the initial account.
	//
	// 	- The name can contain lowercase letters, digits, and underscores (_).
	//
	// 	- The name must start with a lowercase letter and end with a lowercase letter or a digit.
	//
	// 	- The name cannot start with gp.
	//
	// 	- The name must be 2 to 16 characters in length.
	//
	// example:
	//
	// testaccount
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The status of the member that you want to query.
	//
	// 	- **enabled**: managed.
	//
	// 	- **disabled**: not managed.
	//
	// 	- **disabling**: being deleted.
	//
	// example:
	//
	// 1
	AccountStatus *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	// 	- Normal: standard account
	//
	// 	- Super: privileged account
	//
	// example:
	//
	// Normal
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The instance ID.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// CA7E4276-E2D5-5F8D-AF06-9EAB3F6C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAccountResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccountResponseBody) GetAccountDescription() *string {
	return s.AccountDescription
}

func (s *GetAccountResponseBody) GetAccountName() *string {
	return s.AccountName
}

func (s *GetAccountResponseBody) GetAccountStatus() *string {
	return s.AccountStatus
}

func (s *GetAccountResponseBody) GetAccountType() *string {
	return s.AccountType
}

func (s *GetAccountResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *GetAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAccountResponseBody) SetAccountDescription(v string) *GetAccountResponseBody {
	s.AccountDescription = &v
	return s
}

func (s *GetAccountResponseBody) SetAccountName(v string) *GetAccountResponseBody {
	s.AccountName = &v
	return s
}

func (s *GetAccountResponseBody) SetAccountStatus(v string) *GetAccountResponseBody {
	s.AccountStatus = &v
	return s
}

func (s *GetAccountResponseBody) SetAccountType(v string) *GetAccountResponseBody {
	s.AccountType = &v
	return s
}

func (s *GetAccountResponseBody) SetDBInstanceId(v string) *GetAccountResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *GetAccountResponseBody) SetRequestId(v string) *GetAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
