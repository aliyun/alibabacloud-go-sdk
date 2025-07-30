// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetPasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResetPasswordResponseBody
	GetRequestId() *string
}

type ResetPasswordResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A00477A5-167F-56D2-A315-EA77E4BD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetPasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetPasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetPasswordResponseBody) SetRequestId(v string) *ResetPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetPasswordResponseBody) Validate() error {
	return dara.Validate(s)
}
