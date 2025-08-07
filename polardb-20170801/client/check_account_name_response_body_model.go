// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAccountNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CheckAccountNameResponseBody
	GetRequestId() *string
}

type CheckAccountNameResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 925B84D9-CA72-432C-95CF-738C22******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckAccountNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckAccountNameResponseBody) GoString() string {
	return s.String()
}

func (s *CheckAccountNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckAccountNameResponseBody) SetRequestId(v string) *CheckAccountNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckAccountNameResponseBody) Validate() error {
	return dara.Validate(s)
}
