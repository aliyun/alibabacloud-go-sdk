// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAccountPasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ResetAccountPasswordResponseBodyData) *ResetAccountPasswordResponseBody
	GetData() *ResetAccountPasswordResponseBodyData
	SetRequestId(v string) *ResetAccountPasswordResponseBody
	GetRequestId() *string
}

type ResetAccountPasswordResponseBody struct {
	// The result returned.
	Data *ResetAccountPasswordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 5A6A077A-577C-536E-AC13-8E715D7A34C8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetAccountPasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetAccountPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetAccountPasswordResponseBody) GetData() *ResetAccountPasswordResponseBodyData {
	return s.Data
}

func (s *ResetAccountPasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetAccountPasswordResponseBody) SetData(v *ResetAccountPasswordResponseBodyData) *ResetAccountPasswordResponseBody {
	s.Data = v
	return s
}

func (s *ResetAccountPasswordResponseBody) SetRequestId(v string) *ResetAccountPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetAccountPasswordResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ResetAccountPasswordResponseBodyData struct {
	// The name of the account.
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

func (s ResetAccountPasswordResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ResetAccountPasswordResponseBodyData) GoString() string {
	return s.String()
}

func (s *ResetAccountPasswordResponseBodyData) GetAccount() *string {
	return s.Account
}

func (s *ResetAccountPasswordResponseBodyData) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ResetAccountPasswordResponseBodyData) SetAccount(v string) *ResetAccountPasswordResponseBodyData {
	s.Account = &v
	return s
}

func (s *ResetAccountPasswordResponseBodyData) SetDBInstanceId(v string) *ResetAccountPasswordResponseBodyData {
	s.DBInstanceId = &v
	return s
}

func (s *ResetAccountPasswordResponseBodyData) Validate() error {
	return dara.Validate(s)
}
