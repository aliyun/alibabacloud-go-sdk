// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeUserPasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ChangeUserPasswordResponseBody
	GetRequestId() *string
}

type ChangeUserPasswordResponseBody struct {
	// example:
	//
	// AA8D67CB-345D-5CDA-986E-FFAC7D0****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeUserPasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeUserPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeUserPasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeUserPasswordResponseBody) SetRequestId(v string) *ChangeUserPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeUserPasswordResponseBody) Validate() error {
	return dara.Validate(s)
}
