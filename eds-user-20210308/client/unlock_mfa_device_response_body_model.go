// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnlockMfaDeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnlockMfaDeviceResponseBody
	GetRequestId() *string
}

type UnlockMfaDeviceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9BD39C60-4E38-43BE-BA2F-69136C6C5190
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnlockMfaDeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnlockMfaDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *UnlockMfaDeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnlockMfaDeviceResponseBody) SetRequestId(v string) *UnlockMfaDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnlockMfaDeviceResponseBody) Validate() error {
	return dara.Validate(s)
}
