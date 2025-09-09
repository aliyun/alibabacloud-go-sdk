// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *ModifyAccountDescriptionRequest
	GetAccountName() *string
	SetDescription(v string) *ModifyAccountDescriptionRequest
	GetDescription() *string
	SetDrdsInstanceId(v string) *ModifyAccountDescriptionRequest
	GetDrdsInstanceId() *string
}

type ModifyAccountDescriptionRequest struct {
	// The name of the member account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The description of the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the ApsaraDB RDS for PostgreSQL instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds************
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s ModifyAccountDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *ModifyAccountDescriptionRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyAccountDescriptionRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *ModifyAccountDescriptionRequest) SetAccountName(v string) *ModifyAccountDescriptionRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetDescription(v string) *ModifyAccountDescriptionRequest {
	s.Description = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetDrdsInstanceId(v string) *ModifyAccountDescriptionRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
