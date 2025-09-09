// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatabaseAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseAccountId(v string) *CreateDatabaseAccountResponseBody
	GetDatabaseAccountId() *string
	SetRequestId(v string) *CreateDatabaseAccountResponseBody
	GetRequestId() *string
}

type CreateDatabaseAccountResponseBody struct {
	// The ID of the database account.
	//
	// example:
	//
	// 40
	DatabaseAccountId *string `json:"DatabaseAccountId,omitempty" xml:"DatabaseAccountId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B07C465D-B09F-54DD-8FEC-90788BEABAFC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDatabaseAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDatabaseAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDatabaseAccountResponseBody) GetDatabaseAccountId() *string {
	return s.DatabaseAccountId
}

func (s *CreateDatabaseAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDatabaseAccountResponseBody) SetDatabaseAccountId(v string) *CreateDatabaseAccountResponseBody {
	s.DatabaseAccountId = &v
	return s
}

func (s *CreateDatabaseAccountResponseBody) SetRequestId(v string) *CreateDatabaseAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDatabaseAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
