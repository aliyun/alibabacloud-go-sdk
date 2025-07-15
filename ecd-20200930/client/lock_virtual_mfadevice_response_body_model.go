// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLockVirtualMFADeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *LockVirtualMFADeviceResponseBody
	GetRequestId() *string
}

type LockVirtualMFADeviceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LockVirtualMFADeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LockVirtualMFADeviceResponseBody) GoString() string {
	return s.String()
}

func (s *LockVirtualMFADeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LockVirtualMFADeviceResponseBody) SetRequestId(v string) *LockVirtualMFADeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *LockVirtualMFADeviceResponseBody) Validate() error {
	return dara.Validate(s)
}
