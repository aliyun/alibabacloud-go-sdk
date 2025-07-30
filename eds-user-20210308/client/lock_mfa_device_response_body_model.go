// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLockMfaDeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *LockMfaDeviceResponseBody
	GetRequestId() *string
}

type LockMfaDeviceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 42FE70D8-4336-471B-8314-CCCFCE4159FD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LockMfaDeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LockMfaDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *LockMfaDeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LockMfaDeviceResponseBody) SetRequestId(v string) *LockMfaDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *LockMfaDeviceResponseBody) Validate() error {
	return dara.Validate(s)
}
