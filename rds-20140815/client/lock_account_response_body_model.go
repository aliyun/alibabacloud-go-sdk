// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLockAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *LockAccountResponseBody
	GetRequestId() *string
}

type LockAccountResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LockAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LockAccountResponseBody) GoString() string {
	return s.String()
}

func (s *LockAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LockAccountResponseBody) SetRequestId(v string) *LockAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *LockAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
