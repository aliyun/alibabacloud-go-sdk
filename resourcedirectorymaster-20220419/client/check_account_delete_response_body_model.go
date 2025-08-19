// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAccountDeleteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CheckAccountDeleteResponseBody
	GetRequestId() *string
}

type CheckAccountDeleteResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 7CDDDCEF-CDFD-0825-B7D7-217BE0897B22
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckAccountDeleteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckAccountDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *CheckAccountDeleteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckAccountDeleteResponseBody) SetRequestId(v string) *CheckAccountDeleteResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckAccountDeleteResponseBody) Validate() error {
	return dara.Validate(s)
}
