// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContactGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContactGroupId(v string) *CreateContactGroupResponseBody
	GetContactGroupId() *string
	SetRequestId(v string) *CreateContactGroupResponseBody
	GetRequestId() *string
}

type CreateContactGroupResponseBody struct {
	// The ID of the contact group.
	//
	// example:
	//
	// test
	ContactGroupId *string `json:"contactGroupId,omitempty" xml:"contactGroupId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateContactGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateContactGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateContactGroupResponseBody) GetContactGroupId() *string {
	return s.ContactGroupId
}

func (s *CreateContactGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateContactGroupResponseBody) SetContactGroupId(v string) *CreateContactGroupResponseBody {
	s.ContactGroupId = &v
	return s
}

func (s *CreateContactGroupResponseBody) SetRequestId(v string) *CreateContactGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateContactGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
