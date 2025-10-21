// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeleteAccountResponseBodyData) *DeleteAccountResponseBody
	GetData() *DeleteAccountResponseBodyData
	SetRequestId(v string) *DeleteAccountResponseBody
	GetRequestId() *string
}

type DeleteAccountResponseBody struct {
	// The data returned.
	Data *DeleteAccountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 21D06907-CEA5-561D-B6B1-198BCCE99ED1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccountResponseBody) GetData() *DeleteAccountResponseBodyData {
	return s.Data
}

func (s *DeleteAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAccountResponseBody) SetData(v *DeleteAccountResponseBodyData) *DeleteAccountResponseBody {
	s.Data = v
	return s
}

func (s *DeleteAccountResponseBody) SetRequestId(v string) *DeleteAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAccountResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteAccountResponseBodyData struct {
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

func (s DeleteAccountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccountResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteAccountResponseBodyData) GetAccount() *string {
	return s.Account
}

func (s *DeleteAccountResponseBodyData) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteAccountResponseBodyData) SetAccount(v string) *DeleteAccountResponseBodyData {
	s.Account = &v
	return s
}

func (s *DeleteAccountResponseBodyData) SetDBInstanceId(v string) *DeleteAccountResponseBodyData {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteAccountResponseBodyData) Validate() error {
	return dara.Validate(s)
}
