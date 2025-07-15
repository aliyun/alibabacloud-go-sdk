// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnlockVirtualMFADeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnlockVirtualMFADeviceResponseBody
	GetRequestId() *string
}

type UnlockVirtualMFADeviceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnlockVirtualMFADeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnlockVirtualMFADeviceResponseBody) GoString() string {
	return s.String()
}

func (s *UnlockVirtualMFADeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnlockVirtualMFADeviceResponseBody) SetRequestId(v string) *UnlockVirtualMFADeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnlockVirtualMFADeviceResponseBody) Validate() error {
	return dara.Validate(s)
}
