// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMaliciousNoteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateMaliciousNoteResponseBody
	GetRequestId() *string
}

type CreateMaliciousNoteResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F8B6F758-BCD4-597A-8A2C-DA5A552C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMaliciousNoteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMaliciousNoteResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMaliciousNoteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMaliciousNoteResponseBody) SetRequestId(v string) *CreateMaliciousNoteResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMaliciousNoteResponseBody) Validate() error {
	return dara.Validate(s)
}
