// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePocFunctionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreatePocFunctionResponseBody
	GetRequestId() *string
}

type CreatePocFunctionResponseBody struct {
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePocFunctionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePocFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePocFunctionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePocFunctionResponseBody) SetRequestId(v string) *CreatePocFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePocFunctionResponseBody) Validate() error {
	return dara.Validate(s)
}
