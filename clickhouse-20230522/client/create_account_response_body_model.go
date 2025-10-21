// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateAccountResponseBodyData) *CreateAccountResponseBody
	GetData() *CreateAccountResponseBodyData
	SetRequestId(v string) *CreateAccountResponseBody
	GetRequestId() *string
}

type CreateAccountResponseBody struct {
	// The data returned.
	Data *CreateAccountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 2FED790E-FB61-4721-8C1C-07C627FA5A19
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccountResponseBody) GetData() *CreateAccountResponseBodyData {
	return s.Data
}

func (s *CreateAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAccountResponseBody) SetData(v *CreateAccountResponseBodyData) *CreateAccountResponseBody {
	s.Data = v
	return s
}

func (s *CreateAccountResponseBody) SetRequestId(v string) *CreateAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAccountResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAccountResponseBodyData struct {
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

func (s CreateAccountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateAccountResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAccountResponseBodyData) GetAccount() *string {
	return s.Account
}

func (s *CreateAccountResponseBodyData) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateAccountResponseBodyData) SetAccount(v string) *CreateAccountResponseBodyData {
	s.Account = &v
	return s
}

func (s *CreateAccountResponseBodyData) SetDBInstanceId(v string) *CreateAccountResponseBodyData {
	s.DBInstanceId = &v
	return s
}

func (s *CreateAccountResponseBodyData) Validate() error {
	return dara.Validate(s)
}
