// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOriginProtectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateOriginProtectionResponseBody
	GetRequestId() *string
}

type CreateOriginProtectionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-A198-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOriginProtectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOriginProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOriginProtectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOriginProtectionResponseBody) SetRequestId(v string) *CreateOriginProtectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOriginProtectionResponseBody) Validate() error {
	return dara.Validate(s)
}
