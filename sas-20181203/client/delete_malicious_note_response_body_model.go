// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMaliciousNoteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteMaliciousNoteResponseBody
	GetRequestId() *string
}

type DeleteMaliciousNoteResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D65AADFC-1D20-5A6A-8F6A-9FA53CXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMaliciousNoteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMaliciousNoteResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMaliciousNoteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMaliciousNoteResponseBody) SetRequestId(v string) *DeleteMaliciousNoteResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMaliciousNoteResponseBody) Validate() error {
	return dara.Validate(s)
}
