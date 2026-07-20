// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v int64) *GetContactRequest
	GetContactId() *int64
}

type GetContactRequest struct {
	// The contact ID.
	//
	// example:
	//
	// 1397591
	ContactId *int64 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
}

func (s GetContactRequest) String() string {
	return dara.Prettify(s)
}

func (s GetContactRequest) GoString() string {
	return s.String()
}

func (s *GetContactRequest) GetContactId() *int64 {
	return s.ContactId
}

func (s *GetContactRequest) SetContactId(v int64) *GetContactRequest {
	s.ContactId = &v
	return s
}

func (s *GetContactRequest) Validate() error {
	return dara.Validate(s)
}
