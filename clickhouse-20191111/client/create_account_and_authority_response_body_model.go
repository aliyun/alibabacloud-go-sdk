// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccountAndAuthorityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateAccountAndAuthorityResponseBody
	GetRequestId() *string
}

type CreateAccountAndAuthorityResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F5178C10-1407-4987-9133-DE4DC9119F75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAccountAndAuthorityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAccountAndAuthorityResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccountAndAuthorityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAccountAndAuthorityResponseBody) SetRequestId(v string) *CreateAccountAndAuthorityResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAccountAndAuthorityResponseBody) Validate() error {
	return dara.Validate(s)
}
