// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceTestCaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateServiceTestCaseResponseBody
	GetRequestId() *string
}

type UpdateServiceTestCaseResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// DB1FA13E-1087-5654-84D5-58A0ACAD1B18
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceTestCaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceTestCaseResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceTestCaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateServiceTestCaseResponseBody) SetRequestId(v string) *UpdateServiceTestCaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateServiceTestCaseResponseBody) Validate() error {
	return dara.Validate(s)
}
