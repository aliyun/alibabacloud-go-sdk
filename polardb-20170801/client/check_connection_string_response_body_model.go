// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckConnectionStringResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CheckConnectionStringResponseBody
	GetRequestId() *string
}

type CheckConnectionStringResponseBody struct {
	// example:
	//
	// CD3FA5F3-FAF3-44CA-AFFF-BAF869******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckConnectionStringResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckConnectionStringResponseBody) GoString() string {
	return s.String()
}

func (s *CheckConnectionStringResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckConnectionStringResponseBody) SetRequestId(v string) *CheckConnectionStringResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckConnectionStringResponseBody) Validate() error {
	return dara.Validate(s)
}
