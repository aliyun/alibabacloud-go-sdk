// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlertContactGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContactGroupId(v string) *CreateAlertContactGroupResponseBody
	GetContactGroupId() *string
	SetRequestId(v string) *CreateAlertContactGroupResponseBody
	GetRequestId() *string
}

type CreateAlertContactGroupResponseBody struct {
	ContactGroupId *string `json:"ContactGroupId,omitempty" xml:"ContactGroupId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAlertContactGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertContactGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAlertContactGroupResponseBody) GetContactGroupId() *string {
	return s.ContactGroupId
}

func (s *CreateAlertContactGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAlertContactGroupResponseBody) SetContactGroupId(v string) *CreateAlertContactGroupResponseBody {
	s.ContactGroupId = &v
	return s
}

func (s *CreateAlertContactGroupResponseBody) SetRequestId(v string) *CreateAlertContactGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAlertContactGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
