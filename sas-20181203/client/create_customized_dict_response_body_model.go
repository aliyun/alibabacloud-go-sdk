// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomizedDictResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateCustomizedDictResponseBody
	GetRequestId() *string
}

type CreateCustomizedDictResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 08571630-26D8-5E07-A4B7-DF8E89CF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCustomizedDictResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomizedDictResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomizedDictResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCustomizedDictResponseBody) SetRequestId(v string) *CreateCustomizedDictResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustomizedDictResponseBody) Validate() error {
	return dara.Validate(s)
}
