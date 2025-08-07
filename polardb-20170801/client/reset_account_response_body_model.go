// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResetAccountResponseBody
	GetRequestId() *string
}

type ResetAccountResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 2FED790E-FB61-4721-8C1C-07C627******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetAccountResponseBody) GoString() string {
	return s.String()
}

func (s *ResetAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetAccountResponseBody) SetRequestId(v string) *ResetAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
