// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountDescriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModifyAccountDescriptionResponseBodyData) *ModifyAccountDescriptionResponseBody
	GetData() *ModifyAccountDescriptionResponseBodyData
	SetRequestId(v string) *ModifyAccountDescriptionResponseBody
	GetRequestId() *string
}

type ModifyAccountDescriptionResponseBody struct {
	// The returned data.
	Data *ModifyAccountDescriptionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// xxx-xxx-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAccountDescriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionResponseBody) GetData() *ModifyAccountDescriptionResponseBodyData {
	return s.Data
}

func (s *ModifyAccountDescriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAccountDescriptionResponseBody) SetData(v *ModifyAccountDescriptionResponseBodyData) *ModifyAccountDescriptionResponseBody {
	s.Data = v
	return s
}

func (s *ModifyAccountDescriptionResponseBody) SetRequestId(v string) *ModifyAccountDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAccountDescriptionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyAccountDescriptionResponseBodyData struct {
	// The name of the database account.
	//
	// example:
	//
	// testuser
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// cc-xxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s ModifyAccountDescriptionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountDescriptionResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionResponseBodyData) GetAccount() *string {
	return s.Account
}

func (s *ModifyAccountDescriptionResponseBodyData) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyAccountDescriptionResponseBodyData) SetAccount(v string) *ModifyAccountDescriptionResponseBodyData {
	s.Account = &v
	return s
}

func (s *ModifyAccountDescriptionResponseBodyData) SetDBInstanceId(v string) *ModifyAccountDescriptionResponseBodyData {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyAccountDescriptionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
