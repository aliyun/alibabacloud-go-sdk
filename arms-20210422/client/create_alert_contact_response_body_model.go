// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlertContactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v string) *CreateAlertContactResponseBody
	GetContactId() *string
	SetRequestId(v string) *CreateAlertContactResponseBody
	GetRequestId() *string
}

type CreateAlertContactResponseBody struct {
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAlertContactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertContactResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAlertContactResponseBody) GetContactId() *string {
	return s.ContactId
}

func (s *CreateAlertContactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAlertContactResponseBody) SetContactId(v string) *CreateAlertContactResponseBody {
	s.ContactId = &v
	return s
}

func (s *CreateAlertContactResponseBody) SetRequestId(v string) *CreateAlertContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAlertContactResponseBody) Validate() error {
	return dara.Validate(s)
}
