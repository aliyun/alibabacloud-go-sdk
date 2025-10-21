// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountAuthorityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModifyAccountAuthorityResponseBodyData) *ModifyAccountAuthorityResponseBody
	GetData() *ModifyAccountAuthorityResponseBodyData
	SetRequestId(v string) *ModifyAccountAuthorityResponseBody
	GetRequestId() *string
}

type ModifyAccountAuthorityResponseBody struct {
	// The result returned.
	Data *ModifyAccountAuthorityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 05321590-BB65-4720-8CB6-8218E041CDD0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAccountAuthorityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountAuthorityResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccountAuthorityResponseBody) GetData() *ModifyAccountAuthorityResponseBodyData {
	return s.Data
}

func (s *ModifyAccountAuthorityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAccountAuthorityResponseBody) SetData(v *ModifyAccountAuthorityResponseBodyData) *ModifyAccountAuthorityResponseBody {
	s.Data = v
	return s
}

func (s *ModifyAccountAuthorityResponseBody) SetRequestId(v string) *ModifyAccountAuthorityResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAccountAuthorityResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyAccountAuthorityResponseBodyData struct {
	// The name of the database account.
	//
	// example:
	//
	// test1
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s ModifyAccountAuthorityResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountAuthorityResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyAccountAuthorityResponseBodyData) GetAccount() *string {
	return s.Account
}

func (s *ModifyAccountAuthorityResponseBodyData) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyAccountAuthorityResponseBodyData) SetAccount(v string) *ModifyAccountAuthorityResponseBodyData {
	s.Account = &v
	return s
}

func (s *ModifyAccountAuthorityResponseBodyData) SetDBInstanceId(v string) *ModifyAccountAuthorityResponseBodyData {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyAccountAuthorityResponseBodyData) Validate() error {
	return dara.Validate(s)
}
