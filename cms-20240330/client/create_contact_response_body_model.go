// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v string) *CreateContactResponseBody
	GetContactId() *string
	SetRequestId(v string) *CreateContactResponseBody
	GetRequestId() *string
}

type CreateContactResponseBody struct {
	// example:
	//
	// testId
	ContactId *string `json:"contactId,omitempty" xml:"contactId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateContactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateContactResponseBody) GoString() string {
	return s.String()
}

func (s *CreateContactResponseBody) GetContactId() *string {
	return s.ContactId
}

func (s *CreateContactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateContactResponseBody) SetContactId(v string) *CreateContactResponseBody {
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
