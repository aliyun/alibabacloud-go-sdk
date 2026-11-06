// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetApplicationAccountPasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResetApplicationAccountPasswordResponseBody
	GetRequestId() *string
}

type ResetApplicationAccountPasswordResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CD3FA5F3-FAF3-44CA-AFFF-BAF869******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetApplicationAccountPasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetApplicationAccountPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetApplicationAccountPasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetApplicationAccountPasswordResponseBody) SetRequestId(v string) *ResetApplicationAccountPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetApplicationAccountPasswordResponseBody) Validate() error {
	return dara.Validate(s)
}
