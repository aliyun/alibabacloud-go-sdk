// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWehookResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v string) *CreateWehookResponseBody
	GetContactId() *string
	SetRequestId(v string) *CreateWehookResponseBody
	GetRequestId() *string
}

type CreateWehookResponseBody struct {
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateWehookResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWehookResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWehookResponseBody) GetContactId() *string {
	return s.ContactId
}

func (s *CreateWehookResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWehookResponseBody) SetContactId(v string) *CreateWehookResponseBody {
	s.ContactId = &v
	return s
}

func (s *CreateWehookResponseBody) SetRequestId(v string) *CreateWehookResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWehookResponseBody) Validate() error {
	return dara.Validate(s)
}
