// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDelegateAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableDelegateAccountResponseBody
	GetRequestId() *string
}

type DisableDelegateAccountResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 12B3E332-DD16-515B-B695-39BA233AA172
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableDelegateAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableDelegateAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DisableDelegateAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableDelegateAccountResponseBody) SetRequestId(v string) *DisableDelegateAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableDelegateAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
