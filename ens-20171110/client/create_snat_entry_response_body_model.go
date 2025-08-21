// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSnatEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSnatEntryResponseBody
	GetRequestId() *string
	SetSnatEntryId(v string) *CreateSnatEntryResponseBody
	GetSnatEntryId() *string
}

type CreateSnatEntryResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 018EED6A-69CA-58C8-A345-498927D5D34E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the SNAT entry.
	//
	// example:
	//
	// snat-5tc08qfj5ecblfdn2rqr9****
	SnatEntryId *string `json:"SnatEntryId,omitempty" xml:"SnatEntryId,omitempty"`
}

func (s CreateSnatEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSnatEntryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSnatEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSnatEntryResponseBody) GetSnatEntryId() *string {
	return s.SnatEntryId
}

func (s *CreateSnatEntryResponseBody) SetRequestId(v string) *CreateSnatEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSnatEntryResponseBody) SetSnatEntryId(v string) *CreateSnatEntryResponseBody {
	s.SnatEntryId = &v
	return s
}

func (s *CreateSnatEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
