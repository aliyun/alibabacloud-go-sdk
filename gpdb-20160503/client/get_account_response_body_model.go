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
	// The modified account description. The description must meet the following requirements:
	//
	// - The description must start with a Chinese character or an English letter.
	//
	// - The description cannot start with `http://` or `https://`.
	//
	// - The description can contain Chinese characters, English characters, underscores (_), hyphens (-), and digits.
	//
	// - The description must be 2 to 256 characters in length.
	//
	// example:
	//
	// The instance used by this account to log in is DBInstanceId. The name used to log in is AccountName.
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// The name of the initial account. The name must meet the following requirements:
	//
	// - The name can contain lowercase letters, digits, and underscores (_).
	//
	// - The name must start with a lowercase letter and end with a lowercase letter or digit.
	//
	// - The name cannot start with gp.
	//
	// - The name must be 2 to 16 characters in length.
	//
	// example:
	//
	// testaccount
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The status of the database account. Valid values:
	//
	// - **0**: Being created.
	//
	// - **1**: In use.
	//
	// - **3**: Being deleted.
	//
	// example:
	//
	// 1
	AccountStatus *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	// The type of the host account. Valid values:
	//
	// - **Normal**: standard account.
	//
	// - **Admin**: administrator account.
	//
	// For more information about the permissions of host accounts, see [Host account permissions](https://help.aliyun.com/document_detail/176240.html).
	//
	// example:
	//
	// Normal
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the details of all AnalyticDB for PostgreSQL instances in a region, including instance IDs.
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
