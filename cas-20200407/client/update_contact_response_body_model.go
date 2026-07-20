// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v int64) *UpdateContactResponseBody
	GetContactId() *int64
	SetRequestId(v string) *UpdateContactResponseBody
	GetRequestId() *string
}

type UpdateContactResponseBody struct {
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

func (s UpdateContactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateContactResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateContactResponseBody) GetContactId() *int64 {
	return s.ContactId
}

func (s *UpdateContactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateContactResponseBody) SetContactId(v int64) *UpdateContactResponseBody {
	s.ContactId = &v
	return s
}

func (s *UpdateContactResponseBody) SetRequestId(v string) *UpdateContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateContactResponseBody) Validate() error {
	return dara.Validate(s)
}
