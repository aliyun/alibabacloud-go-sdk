// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountDescription(v string) *ModifyAccountDescriptionRequest
	GetAccountDescription() *string
	SetAccountName(v string) *ModifyAccountDescriptionRequest
	GetAccountName() *string
	SetClientToken(v string) *ModifyAccountDescriptionRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *ModifyAccountDescriptionRequest
	GetDBInstanceId() *string
}

type ModifyAccountDescriptionRequest struct {
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
	// testAccoutdescribe
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// The name of the database account.
	//
	// This parameter is required.
	//
	// example:
	//
	// testAccout
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Idempotence check. For more information, see [How to Ensure Idempotence](https://help.aliyun.com/document_detail/327176.html).
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88**********
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the IDs of all AnalyticDB for PostgreSQL instances within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s ModifyAccountDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionRequest) GetAccountDescription() *string {
	return s.AccountDescription
}

func (s *ModifyAccountDescriptionRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *ModifyAccountDescriptionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyAccountDescriptionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyAccountDescriptionRequest) SetAccountDescription(v string) *ModifyAccountDescriptionRequest {
	s.AccountDescription = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetAccountName(v string) *ModifyAccountDescriptionRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetClientToken(v string) *ModifyAccountDescriptionRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetDBInstanceId(v string) *ModifyAccountDescriptionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
