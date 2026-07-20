// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v int64) *CreateContactResponseBody
	GetContactId() *int64
	SetRequestId(v string) *CreateContactResponseBody
	GetRequestId() *string
}

type CreateContactResponseBody struct {
	// The contact ID.
	//
	// example:
	//
	// 1352570
	ContactId *int64 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 71CE8C5B-3737-52A9-97D0-2A9746059A45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateContactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateContactResponseBody) GoString() string {
	return s.String()
}

func (s *CreateContactResponseBody) GetContactId() *int64 {
	return s.ContactId
}

func (s *CreateContactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateContactResponseBody) SetContactId(v int64) *CreateContactResponseBody {
	s.ContactId = &v
	return s
}

func (s *CreateContactResponseBody) SetRequestId(v string) *CreateContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateContactResponseBody) Validate() error {
	return dara.Validate(s)
}
