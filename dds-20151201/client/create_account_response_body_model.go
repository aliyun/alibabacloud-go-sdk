// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateAccountResponseBody
	GetRequestId() *string
}

type CreateAccountResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7FEF1BEF-2842-5758-9126-415A297XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAccountResponseBody) SetRequestId(v string) *CreateAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
